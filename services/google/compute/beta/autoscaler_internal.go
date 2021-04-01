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
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Autoscaler) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "target"); err != nil {
		return err
	}
	if err := dcl.Required(r, "autoscalingPolicy"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.AutoscalingPolicy) {
		if err := r.AutoscalingPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AutoscalerAutoscalingPolicy) validate() error {
	if err := dcl.Required(r, "minNumReplicas"); err != nil {
		return err
	}
	if err := dcl.Required(r, "maxNumReplicas"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ScaleInControl) {
		if err := r.ScaleInControl.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CpuUtilization) {
		if err := r.CpuUtilization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LoadBalancingUtilization) {
		if err := r.LoadBalancingUtilization.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AutoscalerAutoscalingPolicyScaleInControl) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MaxScaledInReplicas) {
		if err := r.MaxScaledInReplicas.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) validate() error {
	return nil
}
func (r *AutoscalerAutoscalingPolicyCpuUtilization) validate() error {
	if err := dcl.Required(r, "utilizationTarget"); err != nil {
		return err
	}
	return nil
}
func (r *AutoscalerAutoscalingPolicyCustomMetricUtilizations) validate() error {
	if err := dcl.Required(r, "metric"); err != nil {
		return err
	}
	return nil
}
func (r *AutoscalerAutoscalingPolicyLoadBalancingUtilization) validate() error {
	if err := dcl.Required(r, "utilizationTarget"); err != nil {
		return err
	}
	return nil
}
func (r *AutoscalerStatusDetails) validate() error {
	return nil
}

func autoscalerGetURL(userBasePath string, r *Autoscaler) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/autoscalers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/autoscalers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func autoscalerListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/autoscalers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/autoscalers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func autoscalerCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/autoscalers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/autoscalers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func autoscalerDeleteURL(userBasePath string, r *Autoscaler) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/autoscalers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/autoscalers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// autoscalerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type autoscalerApiOperation interface {
	do(context.Context, *Autoscaler, *Client) error
}

// newUpdateAutoscalerUpdateRequest creates a request for an
// Autoscaler resource's Update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAutoscalerUpdateRequest(ctx context.Context, f *Autoscaler, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := deriveAutoscalerTarget(f); err != nil {
		return nil, fmt.Errorf("error expanding Target into target: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["target"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicy(c, f.AutoscalingPolicy); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingPolicy into autoscalingPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["autoscalingPolicy"] = v
	}
	req["name"] = fmt.Sprintf("%s", *f.Name)

	return req, nil
}

// marshalUpdateAutoscalerUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateAutoscalerUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAutoscalerUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAutoscalerUpdateOperation) do(ctx context.Context, r *Autoscaler, c *Client) error {
	_, err := c.GetAutoscaler(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Update")
	if err != nil {
		return err
	}

	req, err := newUpdateAutoscalerUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAutoscalerUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listAutoscalerRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := autoscalerListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AutoscalerMaxPage {
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

type listAutoscalerOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listAutoscaler(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Autoscaler, string, error) {
	b, err := c.listAutoscalerRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAutoscalerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Autoscaler
	for _, v := range m.Items {
		res := flattenAutoscaler(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAutoscaler(ctx context.Context, f func(*Autoscaler) bool, resources []*Autoscaler) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAutoscaler(ctx, res)
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

type deleteAutoscalerOperation struct{}

func (op *deleteAutoscalerOperation) do(ctx context.Context, r *Autoscaler, c *Client) error {

	_, err := c.GetAutoscaler(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Autoscaler not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAutoscaler checking for existence. error: %v", err)
		return err
	}

	u, err := autoscalerDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetAutoscaler(ctx, r.urlNormalized())
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
type createAutoscalerOperation struct {
	response map[string]interface{}
}

func (op *createAutoscalerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAutoscalerOperation) do(ctx context.Context, r *Autoscaler, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := autoscalerCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetAutoscaler(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAutoscalerRaw(ctx context.Context, r *Autoscaler) ([]byte, error) {

	u, err := autoscalerGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) autoscalerDiffsForRawDesired(ctx context.Context, rawDesired *Autoscaler, opts ...dcl.ApplyOption) (initial, desired *Autoscaler, diffs []autoscalerDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Autoscaler
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Autoscaler); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Autoscaler, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAutoscaler(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Autoscaler resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Autoscaler resource: %v", err)
		}
		c.Config.Logger.Info("Found that Autoscaler resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAutoscalerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Autoscaler: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Autoscaler: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAutoscalerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Autoscaler: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAutoscalerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Autoscaler: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAutoscaler(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAutoscalerInitialState(rawInitial, rawDesired *Autoscaler) (*Autoscaler, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAutoscalerDesiredState(rawDesired, rawInitial *Autoscaler, opts ...dcl.ApplyOption) (*Autoscaler, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.AutoscalingPolicy = canonicalizeAutoscalerAutoscalingPolicy(rawDesired.AutoscalingPolicy, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.NameToSelfLink(rawDesired.Target, rawInitial.Target) {
		rawDesired.Target = rawInitial.Target
	}
	rawDesired.AutoscalingPolicy = canonicalizeAutoscalerAutoscalingPolicy(rawDesired.AutoscalingPolicy, rawInitial.AutoscalingPolicy, opts...)
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.IsZeroValue(rawDesired.StatusDetails) {
		rawDesired.StatusDetails = rawInitial.StatusDetails
	}
	if dcl.IsZeroValue(rawDesired.RecommendedSize) {
		rawDesired.RecommendedSize = rawInitial.RecommendedSize
	}
	if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawInitial.SelfLinkWithId) {
		rawDesired.SelfLinkWithId = rawInitial.SelfLinkWithId
	}
	if dcl.IsZeroValue(rawDesired.ScalingScheduleStatus) {
		rawDesired.ScalingScheduleStatus = rawInitial.ScalingScheduleStatus
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawInitial.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeAutoscalerNewState(c *Client, rawNew, rawDesired *Autoscaler) (*Autoscaler, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Target) && dcl.IsEmptyValueIndirect(rawDesired.Target) {
		rawNew.Target = rawDesired.Target
	} else {
		if dcl.NameToSelfLink(rawDesired.Target, rawNew.Target) {
			rawNew.Target = rawDesired.Target
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AutoscalingPolicy) && dcl.IsEmptyValueIndirect(rawDesired.AutoscalingPolicy) {
		rawNew.AutoscalingPolicy = rawDesired.AutoscalingPolicy
	} else {
		rawNew.AutoscalingPolicy = canonicalizeNewAutoscalerAutoscalingPolicy(c, rawDesired.AutoscalingPolicy, rawNew.AutoscalingPolicy)
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

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusDetails) && dcl.IsEmptyValueIndirect(rawDesired.StatusDetails) {
		rawNew.StatusDetails = rawDesired.StatusDetails
	} else {
		rawNew.StatusDetails = canonicalizeNewAutoscalerStatusDetailsSlice(c, rawDesired.StatusDetails, rawNew.StatusDetails)
	}

	if dcl.IsEmptyValueIndirect(rawNew.RecommendedSize) && dcl.IsEmptyValueIndirect(rawDesired.RecommendedSize) {
		rawNew.RecommendedSize = rawDesired.RecommendedSize
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLinkWithId) && dcl.IsEmptyValueIndirect(rawDesired.SelfLinkWithId) {
		rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawNew.SelfLinkWithId) {
			rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScalingScheduleStatus) && dcl.IsEmptyValueIndirect(rawDesired.ScalingScheduleStatus) {
		rawNew.ScalingScheduleStatus = rawDesired.ScalingScheduleStatus
	} else {
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAutoscalerAutoscalingPolicy(des, initial *AutoscalerAutoscalingPolicy, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.CoolDownPeriodSec) {
		des.CoolDownPeriodSec = dcl.Int64(60)
	}

	if dcl.IsZeroValue(des.Mode) {
		des.Mode = AutoscalerAutoscalingPolicyModeEnumRef("ON")
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinNumReplicas) {
		des.MinNumReplicas = initial.MinNumReplicas
	}
	if dcl.IsZeroValue(des.MaxNumReplicas) {
		des.MaxNumReplicas = initial.MaxNumReplicas
	}
	des.ScaleInControl = canonicalizeAutoscalerAutoscalingPolicyScaleInControl(des.ScaleInControl, initial.ScaleInControl, opts...)
	if dcl.IsZeroValue(des.CoolDownPeriodSec) {
		des.CoolDownPeriodSec = initial.CoolDownPeriodSec
	}
	des.CpuUtilization = canonicalizeAutoscalerAutoscalingPolicyCpuUtilization(des.CpuUtilization, initial.CpuUtilization, opts...)
	if dcl.IsZeroValue(des.CustomMetricUtilizations) {
		des.CustomMetricUtilizations = initial.CustomMetricUtilizations
	}
	des.LoadBalancingUtilization = canonicalizeAutoscalerAutoscalingPolicyLoadBalancingUtilization(des.LoadBalancingUtilization, initial.LoadBalancingUtilization, opts...)
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}

	return des
}

func canonicalizeNewAutoscalerAutoscalingPolicy(c *Client, des, nw *AutoscalerAutoscalingPolicy) *AutoscalerAutoscalingPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.CoolDownPeriodSec) {
		nw.CoolDownPeriodSec = dcl.Int64(60)
	}

	if dcl.IsZeroValue(nw.Mode) {
		nw.Mode = AutoscalerAutoscalingPolicyModeEnumRef("ON")
	}

	nw.ScaleInControl = canonicalizeNewAutoscalerAutoscalingPolicyScaleInControl(c, des.ScaleInControl, nw.ScaleInControl)
	nw.CpuUtilization = canonicalizeNewAutoscalerAutoscalingPolicyCpuUtilization(c, des.CpuUtilization, nw.CpuUtilization)
	nw.CustomMetricUtilizations = canonicalizeNewAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c, des.CustomMetricUtilizations, nw.CustomMetricUtilizations)
	nw.LoadBalancingUtilization = canonicalizeNewAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, des.LoadBalancingUtilization, nw.LoadBalancingUtilization)

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicySet(c *Client, des, nw []AutoscalerAutoscalingPolicy) []AutoscalerAutoscalingPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicy(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicySlice(c *Client, des, nw []AutoscalerAutoscalingPolicy) []AutoscalerAutoscalingPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerAutoscalingPolicyScaleInControl(des, initial *AutoscalerAutoscalingPolicyScaleInControl, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicyScaleInControl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.MaxScaledInReplicas = canonicalizeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(des.MaxScaledInReplicas, initial.MaxScaledInReplicas, opts...)
	if dcl.IsZeroValue(des.TimeWindowSec) {
		des.TimeWindowSec = initial.TimeWindowSec
	}

	return des
}

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControl(c *Client, des, nw *AutoscalerAutoscalingPolicyScaleInControl) *AutoscalerAutoscalingPolicyScaleInControl {
	if des == nil || nw == nil {
		return nw
	}

	nw.MaxScaledInReplicas = canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, des.MaxScaledInReplicas, nw.MaxScaledInReplicas)

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlSet(c *Client, des, nw []AutoscalerAutoscalingPolicyScaleInControl) []AutoscalerAutoscalingPolicyScaleInControl {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicyScaleInControl
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicyScaleInControl(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlSlice(c *Client, des, nw []AutoscalerAutoscalingPolicyScaleInControl) []AutoscalerAutoscalingPolicyScaleInControl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicyScaleInControl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicyScaleInControl(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(des, initial *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
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

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c *Client, des, nw *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSet(c *Client, des, nw []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice(c *Client, des, nw []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerAutoscalingPolicyCpuUtilization(des, initial *AutoscalerAutoscalingPolicyCpuUtilization, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicyCpuUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	}

	return des
}

func canonicalizeNewAutoscalerAutoscalingPolicyCpuUtilization(c *Client, des, nw *AutoscalerAutoscalingPolicyCpuUtilization) *AutoscalerAutoscalingPolicyCpuUtilization {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicyCpuUtilizationSet(c *Client, des, nw []AutoscalerAutoscalingPolicyCpuUtilization) []AutoscalerAutoscalingPolicyCpuUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicyCpuUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicyCpuUtilization(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicyCpuUtilizationSlice(c *Client, des, nw []AutoscalerAutoscalingPolicyCpuUtilization) []AutoscalerAutoscalingPolicyCpuUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicyCpuUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicyCpuUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerAutoscalingPolicyCustomMetricUtilizations(des, initial *AutoscalerAutoscalingPolicyCustomMetricUtilizations, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Metric, initial.Metric) || dcl.IsZeroValue(des.Metric) {
		des.Metric = initial.Metric
	}
	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	}
	if dcl.IsZeroValue(des.UtilizationTargetType) {
		des.UtilizationTargetType = initial.UtilizationTargetType
	}
	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.SingleInstanceAssignment) {
		des.SingleInstanceAssignment = initial.SingleInstanceAssignment
	}

	return des
}

func canonicalizeNewAutoscalerAutoscalingPolicyCustomMetricUtilizations(c *Client, des, nw *AutoscalerAutoscalingPolicyCustomMetricUtilizations) *AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Metric, nw.Metric) {
		nw.Metric = des.Metric
	}
	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicyCustomMetricUtilizationsSet(c *Client, des, nw []AutoscalerAutoscalingPolicyCustomMetricUtilizations) []AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicyCustomMetricUtilizations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c *Client, des, nw []AutoscalerAutoscalingPolicyCustomMetricUtilizations) []AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicyCustomMetricUtilizations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerAutoscalingPolicyLoadBalancingUtilization(des, initial *AutoscalerAutoscalingPolicyLoadBalancingUtilization, opts ...dcl.ApplyOption) *AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	}

	return des
}

func canonicalizeNewAutoscalerAutoscalingPolicyLoadBalancingUtilization(c *Client, des, nw *AutoscalerAutoscalingPolicyLoadBalancingUtilization) *AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAutoscalerAutoscalingPolicyLoadBalancingUtilizationSet(c *Client, des, nw []AutoscalerAutoscalingPolicyLoadBalancingUtilization) []AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerAutoscalingPolicyLoadBalancingUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &d, &n) {
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

func canonicalizeNewAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice(c *Client, des, nw []AutoscalerAutoscalingPolicyLoadBalancingUtilization) []AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerAutoscalingPolicyLoadBalancingUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalerStatusDetails(des, initial *AutoscalerStatusDetails, opts ...dcl.ApplyOption) *AutoscalerStatusDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewAutoscalerStatusDetails(c *Client, des, nw *AutoscalerStatusDetails) *AutoscalerStatusDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewAutoscalerStatusDetailsSet(c *Client, des, nw []AutoscalerStatusDetails) []AutoscalerStatusDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalerStatusDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalerStatusDetails(c, &d, &n) {
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

func canonicalizeNewAutoscalerStatusDetailsSlice(c *Client, des, nw []AutoscalerStatusDetails) []AutoscalerStatusDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalerStatusDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerStatusDetails(c, &d, &n))
	}

	return items
}

type autoscalerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         autoscalerApiOperation
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
func diffAutoscaler(c *Client, desired, actual *Autoscaler, opts ...dcl.ApplyOption) ([]autoscalerDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []autoscalerDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, autoscalerDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, autoscalerDiff{
			UpdateOp:  &updateAutoscalerUpdateOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.IsZeroValue(desired.Target) && !dcl.NameToSelfLink(desired.Target, actual.Target) {
		c.Config.Logger.Infof("Detected diff in Target.\nDESIRED: %v\nACTUAL: %v", desired.Target, actual.Target)

		diffs = append(diffs, autoscalerDiff{
			UpdateOp:  &updateAutoscalerUpdateOperation{},
			FieldName: "Target",
		})

	}
	if compareAutoscalerAutoscalingPolicy(c, desired.AutoscalingPolicy, actual.AutoscalingPolicy) {
		c.Config.Logger.Infof("Detected diff in AutoscalingPolicy.\nDESIRED: %v\nACTUAL: %v", desired.AutoscalingPolicy, actual.AutoscalingPolicy)

		diffs = append(diffs, autoscalerDiff{
			UpdateOp:  &updateAutoscalerUpdateOperation{},
			FieldName: "AutoscalingPolicy",
		})

	}
	if !dcl.IsZeroValue(desired.Zone) && !dcl.StringCanonicalize(desired.Zone, actual.Zone) {
		c.Config.Logger.Infof("Detected diff in Zone.\nDESIRED: %v\nACTUAL: %v", desired.Zone, actual.Zone)
		diffs = append(diffs, autoscalerDiff{
			RequiresRecreate: true,
			FieldName:        "Zone",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.StringCanonicalize(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, autoscalerDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.SelfLinkWithId) && !dcl.StringCanonicalize(desired.SelfLinkWithId, actual.SelfLinkWithId) {
		c.Config.Logger.Infof("Detected diff in SelfLinkWithId.\nDESIRED: %v\nACTUAL: %v", desired.SelfLinkWithId, actual.SelfLinkWithId)
		diffs = append(diffs, autoscalerDiff{
			RequiresRecreate: true,
			FieldName:        "SelfLinkWithId",
		})
	}
	if !dcl.MapEquals(desired.ScalingScheduleStatus, actual.ScalingScheduleStatus, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in ScalingScheduleStatus.\nDESIRED: %v\nACTUAL: %v", desired.ScalingScheduleStatus, actual.ScalingScheduleStatus)
		diffs = append(diffs, autoscalerDiff{
			RequiresRecreate: true,
			FieldName:        "ScalingScheduleStatus",
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
	var deduped []autoscalerDiff
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
func compareAutoscalerAutoscalingPolicy(c *Client, desired, actual *AutoscalerAutoscalingPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinNumReplicas == nil && desired.MinNumReplicas != nil && !dcl.IsEmptyValueIndirect(desired.MinNumReplicas) {
		c.Config.Logger.Infof("desired MinNumReplicas %s - but actually nil", dcl.SprintResource(desired.MinNumReplicas))
		return true
	}
	if !reflect.DeepEqual(desired.MinNumReplicas, actual.MinNumReplicas) && !dcl.IsZeroValue(desired.MinNumReplicas) {
		c.Config.Logger.Infof("Diff in MinNumReplicas. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinNumReplicas), dcl.SprintResource(actual.MinNumReplicas))
		return true
	}
	if actual.MaxNumReplicas == nil && desired.MaxNumReplicas != nil && !dcl.IsEmptyValueIndirect(desired.MaxNumReplicas) {
		c.Config.Logger.Infof("desired MaxNumReplicas %s - but actually nil", dcl.SprintResource(desired.MaxNumReplicas))
		return true
	}
	if !reflect.DeepEqual(desired.MaxNumReplicas, actual.MaxNumReplicas) && !dcl.IsZeroValue(desired.MaxNumReplicas) {
		c.Config.Logger.Infof("Diff in MaxNumReplicas. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxNumReplicas), dcl.SprintResource(actual.MaxNumReplicas))
		return true
	}
	if actual.ScaleInControl == nil && desired.ScaleInControl != nil && !dcl.IsEmptyValueIndirect(desired.ScaleInControl) {
		c.Config.Logger.Infof("desired ScaleInControl %s - but actually nil", dcl.SprintResource(desired.ScaleInControl))
		return true
	}
	if compareAutoscalerAutoscalingPolicyScaleInControl(c, desired.ScaleInControl, actual.ScaleInControl) && !dcl.IsZeroValue(desired.ScaleInControl) {
		c.Config.Logger.Infof("Diff in ScaleInControl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScaleInControl), dcl.SprintResource(actual.ScaleInControl))
		return true
	}
	if actual.CoolDownPeriodSec == nil && desired.CoolDownPeriodSec != nil && !dcl.IsEmptyValueIndirect(desired.CoolDownPeriodSec) {
		c.Config.Logger.Infof("desired CoolDownPeriodSec %s - but actually nil", dcl.SprintResource(desired.CoolDownPeriodSec))
		return true
	}
	if !reflect.DeepEqual(desired.CoolDownPeriodSec, actual.CoolDownPeriodSec) && !dcl.IsZeroValue(desired.CoolDownPeriodSec) {
		c.Config.Logger.Infof("Diff in CoolDownPeriodSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CoolDownPeriodSec), dcl.SprintResource(actual.CoolDownPeriodSec))
		return true
	}
	if actual.CpuUtilization == nil && desired.CpuUtilization != nil && !dcl.IsEmptyValueIndirect(desired.CpuUtilization) {
		c.Config.Logger.Infof("desired CpuUtilization %s - but actually nil", dcl.SprintResource(desired.CpuUtilization))
		return true
	}
	if compareAutoscalerAutoscalingPolicyCpuUtilization(c, desired.CpuUtilization, actual.CpuUtilization) && !dcl.IsZeroValue(desired.CpuUtilization) {
		c.Config.Logger.Infof("Diff in CpuUtilization. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuUtilization), dcl.SprintResource(actual.CpuUtilization))
		return true
	}
	if actual.CustomMetricUtilizations == nil && desired.CustomMetricUtilizations != nil && !dcl.IsEmptyValueIndirect(desired.CustomMetricUtilizations) {
		c.Config.Logger.Infof("desired CustomMetricUtilizations %s - but actually nil", dcl.SprintResource(desired.CustomMetricUtilizations))
		return true
	}
	if compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c, desired.CustomMetricUtilizations, actual.CustomMetricUtilizations) && !dcl.IsZeroValue(desired.CustomMetricUtilizations) {
		c.Config.Logger.Infof("Diff in CustomMetricUtilizations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CustomMetricUtilizations), dcl.SprintResource(actual.CustomMetricUtilizations))
		return true
	}
	if actual.LoadBalancingUtilization == nil && desired.LoadBalancingUtilization != nil && !dcl.IsEmptyValueIndirect(desired.LoadBalancingUtilization) {
		c.Config.Logger.Infof("desired LoadBalancingUtilization %s - but actually nil", dcl.SprintResource(desired.LoadBalancingUtilization))
		return true
	}
	if compareAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, desired.LoadBalancingUtilization, actual.LoadBalancingUtilization) && !dcl.IsZeroValue(desired.LoadBalancingUtilization) {
		c.Config.Logger.Infof("Diff in LoadBalancingUtilization. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoadBalancingUtilization), dcl.SprintResource(actual.LoadBalancingUtilization))
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
	return false
}

func compareAutoscalerAutoscalingPolicySlice(c *Client, desired, actual []AutoscalerAutoscalingPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyScaleInControl(c *Client, desired, actual *AutoscalerAutoscalingPolicyScaleInControl) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxScaledInReplicas == nil && desired.MaxScaledInReplicas != nil && !dcl.IsEmptyValueIndirect(desired.MaxScaledInReplicas) {
		c.Config.Logger.Infof("desired MaxScaledInReplicas %s - but actually nil", dcl.SprintResource(desired.MaxScaledInReplicas))
		return true
	}
	if compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, desired.MaxScaledInReplicas, actual.MaxScaledInReplicas) && !dcl.IsZeroValue(desired.MaxScaledInReplicas) {
		c.Config.Logger.Infof("Diff in MaxScaledInReplicas. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxScaledInReplicas), dcl.SprintResource(actual.MaxScaledInReplicas))
		return true
	}
	if actual.TimeWindowSec == nil && desired.TimeWindowSec != nil && !dcl.IsEmptyValueIndirect(desired.TimeWindowSec) {
		c.Config.Logger.Infof("desired TimeWindowSec %s - but actually nil", dcl.SprintResource(desired.TimeWindowSec))
		return true
	}
	if !reflect.DeepEqual(desired.TimeWindowSec, actual.TimeWindowSec) && !dcl.IsZeroValue(desired.TimeWindowSec) {
		c.Config.Logger.Infof("Diff in TimeWindowSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeWindowSec), dcl.SprintResource(actual.TimeWindowSec))
		return true
	}
	return false
}

func compareAutoscalerAutoscalingPolicyScaleInControlSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyScaleInControl) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyScaleInControl, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyScaleInControl(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControl, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyScaleInControlMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicyScaleInControl) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyScaleInControl, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControl, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicyScaleInControl(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControl, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c *Client, desired, actual *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) bool {
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

func compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCpuUtilization(c *Client, desired, actual *AutoscalerAutoscalingPolicyCpuUtilization) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.UtilizationTarget == nil && desired.UtilizationTarget != nil && !dcl.IsEmptyValueIndirect(desired.UtilizationTarget) {
		c.Config.Logger.Infof("desired UtilizationTarget %s - but actually nil", dcl.SprintResource(desired.UtilizationTarget))
		return true
	}
	if !reflect.DeepEqual(desired.UtilizationTarget, actual.UtilizationTarget) && !dcl.IsZeroValue(desired.UtilizationTarget) {
		c.Config.Logger.Infof("Diff in UtilizationTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UtilizationTarget), dcl.SprintResource(actual.UtilizationTarget))
		return true
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCpuUtilizationSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyCpuUtilization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyCpuUtilization, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyCpuUtilization(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCpuUtilization, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCpuUtilizationMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicyCpuUtilization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyCpuUtilization, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCpuUtilization, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicyCpuUtilization(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCpuUtilization, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizations(c *Client, desired, actual *AutoscalerAutoscalingPolicyCustomMetricUtilizations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Metric == nil && desired.Metric != nil && !dcl.IsEmptyValueIndirect(desired.Metric) {
		c.Config.Logger.Infof("desired Metric %s - but actually nil", dcl.SprintResource(desired.Metric))
		return true
	}
	if !dcl.StringCanonicalize(desired.Metric, actual.Metric) && !dcl.IsZeroValue(desired.Metric) {
		c.Config.Logger.Infof("Diff in Metric. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metric), dcl.SprintResource(actual.Metric))
		return true
	}
	if actual.UtilizationTarget == nil && desired.UtilizationTarget != nil && !dcl.IsEmptyValueIndirect(desired.UtilizationTarget) {
		c.Config.Logger.Infof("desired UtilizationTarget %s - but actually nil", dcl.SprintResource(desired.UtilizationTarget))
		return true
	}
	if !reflect.DeepEqual(desired.UtilizationTarget, actual.UtilizationTarget) && !dcl.IsZeroValue(desired.UtilizationTarget) {
		c.Config.Logger.Infof("Diff in UtilizationTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UtilizationTarget), dcl.SprintResource(actual.UtilizationTarget))
		return true
	}
	if actual.UtilizationTargetType == nil && desired.UtilizationTargetType != nil && !dcl.IsEmptyValueIndirect(desired.UtilizationTargetType) {
		c.Config.Logger.Infof("desired UtilizationTargetType %s - but actually nil", dcl.SprintResource(desired.UtilizationTargetType))
		return true
	}
	if !reflect.DeepEqual(desired.UtilizationTargetType, actual.UtilizationTargetType) && !dcl.IsZeroValue(desired.UtilizationTargetType) {
		c.Config.Logger.Infof("Diff in UtilizationTargetType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UtilizationTargetType), dcl.SprintResource(actual.UtilizationTargetType))
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !dcl.StringCanonicalize(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.SingleInstanceAssignment == nil && desired.SingleInstanceAssignment != nil && !dcl.IsEmptyValueIndirect(desired.SingleInstanceAssignment) {
		c.Config.Logger.Infof("desired SingleInstanceAssignment %s - but actually nil", dcl.SprintResource(desired.SingleInstanceAssignment))
		return true
	}
	if !reflect.DeepEqual(desired.SingleInstanceAssignment, actual.SingleInstanceAssignment) && !dcl.IsZeroValue(desired.SingleInstanceAssignment) {
		c.Config.Logger.Infof("Diff in SingleInstanceAssignment. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SingleInstanceAssignment), dcl.SprintResource(actual.SingleInstanceAssignment))
		return true
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyCustomMetricUtilizations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizations, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizations, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizations, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyLoadBalancingUtilization(c *Client, desired, actual *AutoscalerAutoscalingPolicyLoadBalancingUtilization) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.UtilizationTarget == nil && desired.UtilizationTarget != nil && !dcl.IsEmptyValueIndirect(desired.UtilizationTarget) {
		c.Config.Logger.Infof("desired UtilizationTarget %s - but actually nil", dcl.SprintResource(desired.UtilizationTarget))
		return true
	}
	if !reflect.DeepEqual(desired.UtilizationTarget, actual.UtilizationTarget) && !dcl.IsZeroValue(desired.UtilizationTarget) {
		c.Config.Logger.Infof("Diff in UtilizationTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UtilizationTarget), dcl.SprintResource(actual.UtilizationTarget))
		return true
	}
	return false
}

func compareAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyLoadBalancingUtilization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyLoadBalancingUtilization, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyLoadBalancingUtilization, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyLoadBalancingUtilizationMap(c *Client, desired, actual map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyLoadBalancingUtilization, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyLoadBalancingUtilization, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyLoadBalancingUtilization, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerStatusDetails(c *Client, desired, actual *AutoscalerStatusDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Message == nil && desired.Message != nil && !dcl.IsEmptyValueIndirect(desired.Message) {
		c.Config.Logger.Infof("desired Message %s - but actually nil", dcl.SprintResource(desired.Message))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
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

func compareAutoscalerStatusDetailsSlice(c *Client, desired, actual []AutoscalerStatusDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerStatusDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerStatusDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerStatusDetails, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerStatusDetailsMap(c *Client, desired, actual map[string]AutoscalerStatusDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerStatusDetails, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalerStatusDetails, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalerStatusDetails(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalerStatusDetails, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(c *Client, desired, actual *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAutoscalerAutoscalingPolicyModeEnumSlice(c *Client, desired, actual []AutoscalerAutoscalingPolicyModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerAutoscalingPolicyModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerAutoscalingPolicyModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerAutoscalingPolicyModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerAutoscalingPolicyModeEnum(c *Client, desired, actual *AutoscalerAutoscalingPolicyModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAutoscalerStatusEnumSlice(c *Client, desired, actual []AutoscalerStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerStatusEnum(c *Client, desired, actual *AutoscalerStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAutoscalerStatusDetailsTypeEnumSlice(c *Client, desired, actual []AutoscalerStatusDetailsTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalerStatusDetailsTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalerStatusDetailsTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalerStatusDetailsTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalerStatusDetailsTypeEnum(c *Client, desired, actual *AutoscalerStatusDetailsTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Autoscaler) urlNormalized() *Autoscaler {
	normalized := deepcopy.Copy(*r).(Autoscaler)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Target = dcl.SelfLinkToName(r.Target)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.SelfLinkWithId = dcl.SelfLinkToName(r.SelfLinkWithId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Autoscaler) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Autoscaler) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Autoscaler) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Autoscaler) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/autoscalers?autoscaler={{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		if dcl.IsZone(r.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/autoscalers?autoscaler={{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Autoscaler resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Autoscaler) marshal(c *Client) ([]byte, error) {
	m, err := expandAutoscaler(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Autoscaler: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAutoscaler decodes JSON responses into the Autoscaler resource schema.
func unmarshalAutoscaler(b []byte, c *Client) (*Autoscaler, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAutoscaler(m, c)
}

func unmarshalMapAutoscaler(m map[string]interface{}, c *Client) (*Autoscaler, error) {
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"scaling_schedule_status", "items"}); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"scaling_schedule_status"},
			v,
		)
	}

	return flattenAutoscaler(c, m), nil
}

// expandAutoscaler expands Autoscaler into a JSON request object.
func expandAutoscaler(c *Client, f *Autoscaler) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := deriveAutoscalerTarget(f); err != nil {
		return nil, fmt.Errorf("error expanding Target into target: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["target"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicy(c, f.AutoscalingPolicy); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingPolicy into autoscalingPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscalingPolicy"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v, err := expandAutoscalerStatusDetailsSlice(c, f.StatusDetails); err != nil {
		return nil, fmt.Errorf("error expanding StatusDetails into statusDetails: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["statusDetails"] = v
	}
	if v := f.RecommendedSize; !dcl.IsEmptyValueIndirect(v) {
		m["recommendedSize"] = v
	}
	if v := f.SelfLinkWithId; !dcl.IsEmptyValueIndirect(v) {
		m["selfLinkWithId"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInItemsStruct(f.ScalingScheduleStatus); err != nil {
		return nil, fmt.Errorf("error expanding ScalingScheduleStatus into scalingScheduleStatus: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scalingScheduleStatus"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenAutoscaler flattens Autoscaler from a JSON request object into the
// Autoscaler type.
func flattenAutoscaler(c *Client, i interface{}) *Autoscaler {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Autoscaler{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Target = dcl.FlattenString(m["target"])
	r.AutoscalingPolicy = flattenAutoscalerAutoscalingPolicy(c, m["autoscalingPolicy"])
	r.Zone = dcl.FlattenString(m["zone"])
	r.Region = dcl.FlattenString(m["region"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Status = flattenAutoscalerStatusEnum(m["status"])
	r.StatusDetails = flattenAutoscalerStatusDetailsSlice(c, m["statusDetails"])
	r.RecommendedSize = dcl.FlattenInteger(m["recommendedSize"])
	r.SelfLinkWithId = dcl.FlattenString(m["selfLinkWithId"])
	r.ScalingScheduleStatus = dcl.FlattenKeyValuePairs(m["scalingScheduleStatus"])
	r.Project = dcl.FlattenString(m["project"])
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandAutoscalerAutoscalingPolicyMap expands the contents of AutoscalerAutoscalingPolicy into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyMap(c *Client, f map[string]AutoscalerAutoscalingPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicySlice expands the contents of AutoscalerAutoscalingPolicy into a JSON
// request object.
func expandAutoscalerAutoscalingPolicySlice(c *Client, f []AutoscalerAutoscalingPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyMap flattens the contents of AutoscalerAutoscalingPolicy from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicy{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicy{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicy)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicySlice flattens the contents of AutoscalerAutoscalingPolicy from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicySlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicy{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicy{}
	}

	items := make([]AutoscalerAutoscalingPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicy expands an instance of AutoscalerAutoscalingPolicy into a JSON
// request object.
func expandAutoscalerAutoscalingPolicy(c *Client, f *AutoscalerAutoscalingPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	m["minNumReplicas"] = f.MinNumReplicas
	if v := f.MaxNumReplicas; !dcl.IsEmptyValueIndirect(v) {
		m["maxNumReplicas"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicyScaleInControl(c, f.ScaleInControl); err != nil {
		return nil, fmt.Errorf("error expanding ScaleInControl into scaleInControl: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scaleInControl"] = v
	}
	if v := f.CoolDownPeriodSec; !dcl.IsEmptyValueIndirect(v) {
		m["coolDownPeriodSec"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicyCpuUtilization(c, f.CpuUtilization); err != nil {
		return nil, fmt.Errorf("error expanding CpuUtilization into cpuUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cpuUtilization"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c, f.CustomMetricUtilizations); err != nil {
		return nil, fmt.Errorf("error expanding CustomMetricUtilizations into customMetricUtilizations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customMetricUtilizations"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, f.LoadBalancingUtilization); err != nil {
		return nil, fmt.Errorf("error expanding LoadBalancingUtilization into loadBalancingUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loadBalancingUtilization"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenAutoscalerAutoscalingPolicy flattens an instance of AutoscalerAutoscalingPolicy from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicy(c *Client, i interface{}) *AutoscalerAutoscalingPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicy{}
	r.MinNumReplicas = dcl.FlattenInteger(m["minNumReplicas"])
	r.MaxNumReplicas = dcl.FlattenInteger(m["maxNumReplicas"])
	r.ScaleInControl = flattenAutoscalerAutoscalingPolicyScaleInControl(c, m["scaleInControl"])
	r.CoolDownPeriodSec = dcl.FlattenInteger(m["coolDownPeriodSec"])
	if dcl.IsEmptyValueIndirect(m["coolDownPeriodSec"]) {
		c.Config.Logger.Info("Using default value for coolDownPeriodSec.")
		r.CoolDownPeriodSec = dcl.Int64(60)
	}
	r.CpuUtilization = flattenAutoscalerAutoscalingPolicyCpuUtilization(c, m["cpuUtilization"])
	r.CustomMetricUtilizations = flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c, m["customMetricUtilizations"])
	r.LoadBalancingUtilization = flattenAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, m["loadBalancingUtilization"])
	r.Mode = flattenAutoscalerAutoscalingPolicyModeEnum(m["mode"])
	if dcl.IsEmptyValueIndirect(m["mode"]) {
		c.Config.Logger.Info("Using default value for mode.")
		r.Mode = AutoscalerAutoscalingPolicyModeEnumRef("ON")
	}

	return r
}

// expandAutoscalerAutoscalingPolicyScaleInControlMap expands the contents of AutoscalerAutoscalingPolicyScaleInControl into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControlMap(c *Client, f map[string]AutoscalerAutoscalingPolicyScaleInControl) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyScaleInControl(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicyScaleInControlSlice expands the contents of AutoscalerAutoscalingPolicyScaleInControl into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControlSlice(c *Client, f []AutoscalerAutoscalingPolicyScaleInControl) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyScaleInControl(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyScaleInControlMap flattens the contents of AutoscalerAutoscalingPolicyScaleInControl from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControlMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyScaleInControl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyScaleInControl{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyScaleInControl{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyScaleInControl)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyScaleInControl(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyScaleInControlSlice flattens the contents of AutoscalerAutoscalingPolicyScaleInControl from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControlSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyScaleInControl {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyScaleInControl{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyScaleInControl{}
	}

	items := make([]AutoscalerAutoscalingPolicyScaleInControl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyScaleInControl(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicyScaleInControl expands an instance of AutoscalerAutoscalingPolicyScaleInControl into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControl(c *Client, f *AutoscalerAutoscalingPolicyScaleInControl) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, f.MaxScaledInReplicas); err != nil {
		return nil, fmt.Errorf("error expanding MaxScaledInReplicas into maxScaledInReplicas: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxScaledInReplicas"] = v
	}
	if v := f.TimeWindowSec; !dcl.IsEmptyValueIndirect(v) {
		m["timeWindowSec"] = v
	}

	return m, nil
}

// flattenAutoscalerAutoscalingPolicyScaleInControl flattens an instance of AutoscalerAutoscalingPolicyScaleInControl from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControl(c *Client, i interface{}) *AutoscalerAutoscalingPolicyScaleInControl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicyScaleInControl{}
	r.MaxScaledInReplicas = flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, m["maxScaledInReplicas"])
	r.TimeWindowSec = dcl.FlattenInteger(m["timeWindowSec"])

	return r
}

// expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasMap expands the contents of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasMap(c *Client, f map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice expands the contents of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice(c *Client, f []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasMap flattens the contents of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice flattens the contents of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}
	}

	items := make([]AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas expands an instance of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c *Client, f *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) (map[string]interface{}, error) {
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

// flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas flattens an instance of AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(c *Client, i interface{}) *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Calculated = dcl.FlattenInteger(m["calculated"])

	return r
}

// expandAutoscalerAutoscalingPolicyCpuUtilizationMap expands the contents of AutoscalerAutoscalingPolicyCpuUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCpuUtilizationMap(c *Client, f map[string]AutoscalerAutoscalingPolicyCpuUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyCpuUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicyCpuUtilizationSlice expands the contents of AutoscalerAutoscalingPolicyCpuUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCpuUtilizationSlice(c *Client, f []AutoscalerAutoscalingPolicyCpuUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyCpuUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyCpuUtilizationMap flattens the contents of AutoscalerAutoscalingPolicyCpuUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCpuUtilizationMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyCpuUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyCpuUtilization{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyCpuUtilization{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyCpuUtilization)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyCpuUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyCpuUtilizationSlice flattens the contents of AutoscalerAutoscalingPolicyCpuUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCpuUtilizationSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyCpuUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyCpuUtilization{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyCpuUtilization{}
	}

	items := make([]AutoscalerAutoscalingPolicyCpuUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyCpuUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicyCpuUtilization expands an instance of AutoscalerAutoscalingPolicyCpuUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCpuUtilization(c *Client, f *AutoscalerAutoscalingPolicyCpuUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UtilizationTarget; !dcl.IsEmptyValueIndirect(v) {
		m["utilizationTarget"] = v
	}

	return m, nil
}

// flattenAutoscalerAutoscalingPolicyCpuUtilization flattens an instance of AutoscalerAutoscalingPolicyCpuUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCpuUtilization(c *Client, i interface{}) *AutoscalerAutoscalingPolicyCpuUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicyCpuUtilization{}
	r.UtilizationTarget = dcl.FlattenDouble(m["utilizationTarget"])

	return r
}

// expandAutoscalerAutoscalingPolicyCustomMetricUtilizationsMap expands the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizations into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCustomMetricUtilizationsMap(c *Client, f map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice expands the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizations into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c *Client, f []AutoscalerAutoscalingPolicyCustomMetricUtilizations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsMap flattens the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizations from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizations)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice flattens the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizations from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyCustomMetricUtilizations{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyCustomMetricUtilizations{}
	}

	items := make([]AutoscalerAutoscalingPolicyCustomMetricUtilizations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyCustomMetricUtilizations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicyCustomMetricUtilizations expands an instance of AutoscalerAutoscalingPolicyCustomMetricUtilizations into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyCustomMetricUtilizations(c *Client, f *AutoscalerAutoscalingPolicyCustomMetricUtilizations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Metric; !dcl.IsEmptyValueIndirect(v) {
		m["metric"] = v
	}
	if v := f.UtilizationTarget; !dcl.IsEmptyValueIndirect(v) {
		m["utilizationTarget"] = v
	}
	if v := f.UtilizationTargetType; !dcl.IsEmptyValueIndirect(v) {
		m["utilizationTargetType"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.SingleInstanceAssignment; !dcl.IsEmptyValueIndirect(v) {
		m["singleInstanceAssignment"] = v
	}

	return m, nil
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizations flattens an instance of AutoscalerAutoscalingPolicyCustomMetricUtilizations from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizations(c *Client, i interface{}) *AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicyCustomMetricUtilizations{}
	r.Metric = dcl.FlattenString(m["metric"])
	r.UtilizationTarget = dcl.FlattenDouble(m["utilizationTarget"])
	r.UtilizationTargetType = flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(m["utilizationTargetType"])
	r.Filter = dcl.FlattenString(m["filter"])
	r.SingleInstanceAssignment = dcl.FlattenDouble(m["singleInstanceAssignment"])

	return r
}

// expandAutoscalerAutoscalingPolicyLoadBalancingUtilizationMap expands the contents of AutoscalerAutoscalingPolicyLoadBalancingUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyLoadBalancingUtilizationMap(c *Client, f map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice expands the contents of AutoscalerAutoscalingPolicyLoadBalancingUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice(c *Client, f []AutoscalerAutoscalingPolicyLoadBalancingUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerAutoscalingPolicyLoadBalancingUtilizationMap flattens the contents of AutoscalerAutoscalingPolicyLoadBalancingUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyLoadBalancingUtilizationMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyLoadBalancingUtilization)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice flattens the contents of AutoscalerAutoscalingPolicyLoadBalancingUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyLoadBalancingUtilizationSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyLoadBalancingUtilization{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyLoadBalancingUtilization{}
	}

	items := make([]AutoscalerAutoscalingPolicyLoadBalancingUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyLoadBalancingUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerAutoscalingPolicyLoadBalancingUtilization expands an instance of AutoscalerAutoscalingPolicyLoadBalancingUtilization into a JSON
// request object.
func expandAutoscalerAutoscalingPolicyLoadBalancingUtilization(c *Client, f *AutoscalerAutoscalingPolicyLoadBalancingUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UtilizationTarget; !dcl.IsEmptyValueIndirect(v) {
		m["utilizationTarget"] = v
	}

	return m, nil
}

// flattenAutoscalerAutoscalingPolicyLoadBalancingUtilization flattens an instance of AutoscalerAutoscalingPolicyLoadBalancingUtilization from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyLoadBalancingUtilization(c *Client, i interface{}) *AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerAutoscalingPolicyLoadBalancingUtilization{}
	r.UtilizationTarget = dcl.FlattenDouble(m["utilizationTarget"])

	return r
}

// expandAutoscalerStatusDetailsMap expands the contents of AutoscalerStatusDetails into a JSON
// request object.
func expandAutoscalerStatusDetailsMap(c *Client, f map[string]AutoscalerStatusDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalerStatusDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalerStatusDetailsSlice expands the contents of AutoscalerStatusDetails into a JSON
// request object.
func expandAutoscalerStatusDetailsSlice(c *Client, f []AutoscalerStatusDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalerStatusDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalerStatusDetailsMap flattens the contents of AutoscalerStatusDetails from a JSON
// response object.
func flattenAutoscalerStatusDetailsMap(c *Client, i interface{}) map[string]AutoscalerStatusDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerStatusDetails{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerStatusDetails{}
	}

	items := make(map[string]AutoscalerStatusDetails)
	for k, item := range a {
		items[k] = *flattenAutoscalerStatusDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalerStatusDetailsSlice flattens the contents of AutoscalerStatusDetails from a JSON
// response object.
func flattenAutoscalerStatusDetailsSlice(c *Client, i interface{}) []AutoscalerStatusDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerStatusDetails{}
	}

	if len(a) == 0 {
		return []AutoscalerStatusDetails{}
	}

	items := make([]AutoscalerStatusDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerStatusDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalerStatusDetails expands an instance of AutoscalerStatusDetails into a JSON
// request object.
func expandAutoscalerStatusDetails(c *Client, f *AutoscalerStatusDetails) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenAutoscalerStatusDetails flattens an instance of AutoscalerStatusDetails from a JSON
// response object.
func flattenAutoscalerStatusDetails(c *Client, i interface{}) *AutoscalerStatusDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalerStatusDetails{}
	r.Message = dcl.FlattenString(m["message"])
	r.Type = flattenAutoscalerStatusDetailsTypeEnum(m["type"])

	return r
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumSlice flattens the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum{}
	}

	items := make([]AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum with the same value as that string.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(i interface{}) *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumRef("")
	}

	return AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumRef(s)
}

// flattenAutoscalerAutoscalingPolicyModeEnumSlice flattens the contents of AutoscalerAutoscalingPolicyModeEnum from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyModeEnumSlice(c *Client, i interface{}) []AutoscalerAutoscalingPolicyModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerAutoscalingPolicyModeEnum{}
	}

	if len(a) == 0 {
		return []AutoscalerAutoscalingPolicyModeEnum{}
	}

	items := make([]AutoscalerAutoscalingPolicyModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerAutoscalingPolicyModeEnum(item.(interface{})))
	}

	return items
}

// flattenAutoscalerAutoscalingPolicyModeEnum asserts that an interface is a string, and returns a
// pointer to a *AutoscalerAutoscalingPolicyModeEnum with the same value as that string.
func flattenAutoscalerAutoscalingPolicyModeEnum(i interface{}) *AutoscalerAutoscalingPolicyModeEnum {
	s, ok := i.(string)
	if !ok {
		return AutoscalerAutoscalingPolicyModeEnumRef("")
	}

	return AutoscalerAutoscalingPolicyModeEnumRef(s)
}

// flattenAutoscalerStatusEnumSlice flattens the contents of AutoscalerStatusEnum from a JSON
// response object.
func flattenAutoscalerStatusEnumSlice(c *Client, i interface{}) []AutoscalerStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerStatusEnum{}
	}

	if len(a) == 0 {
		return []AutoscalerStatusEnum{}
	}

	items := make([]AutoscalerStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerStatusEnum(item.(interface{})))
	}

	return items
}

// flattenAutoscalerStatusEnum asserts that an interface is a string, and returns a
// pointer to a *AutoscalerStatusEnum with the same value as that string.
func flattenAutoscalerStatusEnum(i interface{}) *AutoscalerStatusEnum {
	s, ok := i.(string)
	if !ok {
		return AutoscalerStatusEnumRef("")
	}

	return AutoscalerStatusEnumRef(s)
}

// flattenAutoscalerStatusDetailsTypeEnumSlice flattens the contents of AutoscalerStatusDetailsTypeEnum from a JSON
// response object.
func flattenAutoscalerStatusDetailsTypeEnumSlice(c *Client, i interface{}) []AutoscalerStatusDetailsTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalerStatusDetailsTypeEnum{}
	}

	if len(a) == 0 {
		return []AutoscalerStatusDetailsTypeEnum{}
	}

	items := make([]AutoscalerStatusDetailsTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalerStatusDetailsTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAutoscalerStatusDetailsTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AutoscalerStatusDetailsTypeEnum with the same value as that string.
func flattenAutoscalerStatusDetailsTypeEnum(i interface{}) *AutoscalerStatusDetailsTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AutoscalerStatusDetailsTypeEnumRef("")
	}

	return AutoscalerStatusDetailsTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Autoscaler) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAutoscaler(b, c)
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
