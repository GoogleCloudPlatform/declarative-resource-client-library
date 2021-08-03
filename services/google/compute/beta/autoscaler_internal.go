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
	if v, err := deriveAutoscalerTarget(f, f.Target); err != nil {
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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAutoscalerUpdateOperation) do(ctx context.Context, r *Autoscaler, c *Client) error {
	_, err := c.GetAutoscaler(ctx, r.URLNormalized())
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
		res, err := unmarshalMapAutoscaler(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetAutoscaler(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Autoscaler not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAutoscaler checking for existence. error: %v", err)
		return err
	}

	u, err := autoscalerDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetAutoscaler(ctx, r.URLNormalized())
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

	if _, err := c.GetAutoscaler(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAutoscalerRaw(ctx context.Context, r *Autoscaler) ([]byte, error) {

	u, err := autoscalerGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) autoscalerDiffsForRawDesired(ctx context.Context, rawDesired *Autoscaler, opts ...dcl.ApplyOption) (initial, desired *Autoscaler, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetAutoscaler(ctx, fetchState.URLNormalized())
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
	canonicalDesired := &Autoscaler{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.NameToSelfLink(rawDesired.Target, rawInitial.Target) {
		canonicalDesired.Target = rawInitial.Target
	} else {
		canonicalDesired.Target = rawDesired.Target
	}
	canonicalDesired.AutoscalingPolicy = canonicalizeAutoscalerAutoscalingPolicy(rawDesired.AutoscalingPolicy, rawInitial.AutoscalingPolicy, opts...)
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		canonicalDesired.Zone = rawInitial.Zone
	} else {
		canonicalDesired.Zone = rawDesired.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		canonicalDesired.Region = rawInitial.Region
	} else {
		canonicalDesired.Region = rawDesired.Region
	}
	if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawInitial.SelfLinkWithId) {
		canonicalDesired.SelfLinkWithId = rawInitial.SelfLinkWithId
	} else {
		canonicalDesired.SelfLinkWithId = rawDesired.SelfLinkWithId
	}
	if dcl.IsZeroValue(rawDesired.ScalingScheduleStatus) {
		canonicalDesired.ScalingScheduleStatus = rawInitial.ScalingScheduleStatus
	} else {
		canonicalDesired.ScalingScheduleStatus = rawDesired.ScalingScheduleStatus
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

	cDes := &AutoscalerAutoscalingPolicy{}

	if dcl.IsZeroValue(des.MinNumReplicas) {
		des.MinNumReplicas = initial.MinNumReplicas
	} else {
		cDes.MinNumReplicas = des.MinNumReplicas
	}
	if dcl.IsZeroValue(des.MaxNumReplicas) {
		des.MaxNumReplicas = initial.MaxNumReplicas
	} else {
		cDes.MaxNumReplicas = des.MaxNumReplicas
	}
	cDes.ScaleInControl = canonicalizeAutoscalerAutoscalingPolicyScaleInControl(des.ScaleInControl, initial.ScaleInControl, opts...)
	if dcl.IsZeroValue(des.CoolDownPeriodSec) {
		des.CoolDownPeriodSec = initial.CoolDownPeriodSec
	} else {
		cDes.CoolDownPeriodSec = des.CoolDownPeriodSec
	}
	cDes.CpuUtilization = canonicalizeAutoscalerAutoscalingPolicyCpuUtilization(des.CpuUtilization, initial.CpuUtilization, opts...)
	if dcl.IsZeroValue(des.CustomMetricUtilizations) {
		des.CustomMetricUtilizations = initial.CustomMetricUtilizations
	} else {
		cDes.CustomMetricUtilizations = des.CustomMetricUtilizations
	}
	cDes.LoadBalancingUtilization = canonicalizeAutoscalerAutoscalingPolicyLoadBalancingUtilization(des.LoadBalancingUtilization, initial.LoadBalancingUtilization, opts...)
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerAutoscalingPolicyScaleInControl{}

	cDes.MaxScaledInReplicas = canonicalizeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(des.MaxScaledInReplicas, initial.MaxScaledInReplicas, opts...)
	if dcl.IsZeroValue(des.TimeWindowSec) {
		des.TimeWindowSec = initial.TimeWindowSec
	} else {
		cDes.TimeWindowSec = des.TimeWindowSec
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyScaleInControlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	} else {
		cDes.Fixed = des.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerAutoscalingPolicyCpuUtilization{}

	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	} else {
		cDes.UtilizationTarget = des.UtilizationTarget
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyCpuUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerAutoscalingPolicyCustomMetricUtilizations{}

	if dcl.StringCanonicalize(des.Metric, initial.Metric) || dcl.IsZeroValue(des.Metric) {
		cDes.Metric = initial.Metric
	} else {
		cDes.Metric = des.Metric
	}
	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	} else {
		cDes.UtilizationTarget = des.UtilizationTarget
	}
	if dcl.IsZeroValue(des.UtilizationTargetType) {
		des.UtilizationTargetType = initial.UtilizationTargetType
	} else {
		cDes.UtilizationTargetType = des.UtilizationTargetType
	}
	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	if dcl.IsZeroValue(des.SingleInstanceAssignment) {
		des.SingleInstanceAssignment = initial.SingleInstanceAssignment
	} else {
		cDes.SingleInstanceAssignment = des.SingleInstanceAssignment
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerAutoscalingPolicyLoadBalancingUtilization{}

	if dcl.IsZeroValue(des.UtilizationTarget) {
		des.UtilizationTarget = initial.UtilizationTarget
	} else {
		cDes.UtilizationTarget = des.UtilizationTarget
	}

	return cDes
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
			if diffs, _ := compareAutoscalerAutoscalingPolicyLoadBalancingUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &AutoscalerStatusDetails{}

	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}

	return cDes
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
			if diffs, _ := compareAutoscalerStatusDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []AutoscalerStatusDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalerStatusDetails(c, &d, &n))
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
func diffAutoscaler(c *Client, desired, actual *Autoscaler, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Target, actual.Target, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Target")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoscalingPolicy, actual.AutoscalingPolicy, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicy, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("AutoscalingPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusDetails, actual.StatusDetails, dcl.Info{OutputOnly: true, ObjectFunction: compareAutoscalerStatusDetailsNewStyle, EmptyObject: EmptyAutoscalerStatusDetails, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StatusDetails")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RecommendedSize, actual.RecommendedSize, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RecommendedSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLinkWithId, actual.SelfLinkWithId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLinkWithId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScalingScheduleStatus, actual.ScalingScheduleStatus, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScalingScheduleStatus")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
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
func compareAutoscalerAutoscalingPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicy)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicy or *AutoscalerAutoscalingPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicy)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinNumReplicas, actual.MinNumReplicas, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("MinNumReplicas")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxNumReplicas, actual.MaxNumReplicas, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("MaxNumReplicas")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScaleInControl, actual.ScaleInControl, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyScaleInControlNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicyScaleInControl, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("ScaleInControl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CoolDownPeriodSec, actual.CoolDownPeriodSec, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CoolDownPeriodSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuUtilization, actual.CpuUtilization, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyCpuUtilizationNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicyCpuUtilization, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("CpuUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomMetricUtilizations, actual.CustomMetricUtilizations, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicyCustomMetricUtilizations, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("CustomMetricUtilizations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LoadBalancingUtilization, actual.LoadBalancingUtilization, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyLoadBalancingUtilizationNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicyLoadBalancingUtilization, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("LoadBalancingUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerAutoscalingPolicyScaleInControlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicyScaleInControl)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicyScaleInControl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyScaleInControl or *AutoscalerAutoscalingPolicyScaleInControl", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicyScaleInControl)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicyScaleInControl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyScaleInControl", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxScaledInReplicas, actual.MaxScaledInReplicas, dcl.Info{ObjectFunction: compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasNewStyle, EmptyObject: EmptyAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas, OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("MaxScaledInReplicas")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeWindowSec, actual.TimeWindowSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("TimeWindowSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas or *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fixed, actual.Fixed, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Fixed")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Calculated, actual.Calculated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Calculated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerAutoscalingPolicyCpuUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicyCpuUtilization)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicyCpuUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyCpuUtilization or *AutoscalerAutoscalingPolicyCpuUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicyCpuUtilization)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicyCpuUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyCpuUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UtilizationTarget, actual.UtilizationTarget, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("UtilizationTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerAutoscalingPolicyCustomMetricUtilizationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicyCustomMetricUtilizations)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicyCustomMetricUtilizations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyCustomMetricUtilizations or *AutoscalerAutoscalingPolicyCustomMetricUtilizations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicyCustomMetricUtilizations)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicyCustomMetricUtilizations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyCustomMetricUtilizations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Metric, actual.Metric, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Metric")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UtilizationTarget, actual.UtilizationTarget, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("UtilizationTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UtilizationTargetType, actual.UtilizationTargetType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("UtilizationTargetType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SingleInstanceAssignment, actual.SingleInstanceAssignment, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("SingleInstanceAssignment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerAutoscalingPolicyLoadBalancingUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerAutoscalingPolicyLoadBalancingUtilization)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerAutoscalingPolicyLoadBalancingUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyLoadBalancingUtilization or *AutoscalerAutoscalingPolicyLoadBalancingUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerAutoscalingPolicyLoadBalancingUtilization)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerAutoscalingPolicyLoadBalancingUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerAutoscalingPolicyLoadBalancingUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UtilizationTarget, actual.UtilizationTarget, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAutoscalerUpdateOperation")}, fn.AddNest("UtilizationTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAutoscalerStatusDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AutoscalerStatusDetails)
	if !ok {
		desiredNotPointer, ok := d.(AutoscalerStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerStatusDetails or *AutoscalerStatusDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AutoscalerStatusDetails)
	if !ok {
		actualNotPointer, ok := a.(AutoscalerStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AutoscalerStatusDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
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

func (r *Autoscaler) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Autoscaler) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Autoscaler) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Autoscaler) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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
	if v, err := deriveAutoscalerTarget(f, f.Target); err != nil {
		return nil, fmt.Errorf("error expanding Target into target: %w", err)
	} else if v != nil {
		m["target"] = v
	}
	if v, err := expandAutoscalerAutoscalingPolicy(c, f.AutoscalingPolicy); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingPolicy into autoscalingPolicy: %w", err)
	} else if v != nil {
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
	} else {
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
	} else if v != nil {
		m["scalingScheduleStatus"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
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

	res := &Autoscaler{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Target = dcl.FlattenString(m["target"])
	res.AutoscalingPolicy = flattenAutoscalerAutoscalingPolicy(c, m["autoscalingPolicy"])
	res.Zone = dcl.FlattenString(m["zone"])
	res.Region = dcl.FlattenString(m["region"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Status = flattenAutoscalerStatusEnum(m["status"])
	res.StatusDetails = flattenAutoscalerStatusDetailsSlice(c, m["statusDetails"])
	res.RecommendedSize = dcl.FlattenInteger(m["recommendedSize"])
	res.SelfLinkWithId = dcl.FlattenString(m["selfLinkWithId"])
	res.ScalingScheduleStatus = dcl.FlattenKeyValuePairs(m["scalingScheduleStatus"])
	res.Project = dcl.FlattenString(m["project"])
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Location = dcl.FlattenString(m["location"])

	return res
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
	if v := f.MinNumReplicas; v != nil {
		m["minNumReplicas"] = v
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicyScaleInControl
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicyCpuUtilization
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicyCustomMetricUtilizations
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerAutoscalingPolicyLoadBalancingUtilization
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAutoscalerStatusDetails
	}
	r.Message = dcl.FlattenString(m["message"])
	r.Type = flattenAutoscalerStatusDetailsTypeEnum(m["type"])

	return r
}

// flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumMap flattens the contents of AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(item.(interface{}))
	}

	return items
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

// flattenAutoscalerAutoscalingPolicyModeEnumMap flattens the contents of AutoscalerAutoscalingPolicyModeEnum from a JSON
// response object.
func flattenAutoscalerAutoscalingPolicyModeEnumMap(c *Client, i interface{}) map[string]AutoscalerAutoscalingPolicyModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerAutoscalingPolicyModeEnum{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerAutoscalingPolicyModeEnum{}
	}

	items := make(map[string]AutoscalerAutoscalingPolicyModeEnum)
	for k, item := range a {
		items[k] = *flattenAutoscalerAutoscalingPolicyModeEnum(item.(interface{}))
	}

	return items
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

// flattenAutoscalerStatusEnumMap flattens the contents of AutoscalerStatusEnum from a JSON
// response object.
func flattenAutoscalerStatusEnumMap(c *Client, i interface{}) map[string]AutoscalerStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerStatusEnum{}
	}

	items := make(map[string]AutoscalerStatusEnum)
	for k, item := range a {
		items[k] = *flattenAutoscalerStatusEnum(item.(interface{}))
	}

	return items
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

// flattenAutoscalerStatusDetailsTypeEnumMap flattens the contents of AutoscalerStatusDetailsTypeEnum from a JSON
// response object.
func flattenAutoscalerStatusDetailsTypeEnumMap(c *Client, i interface{}) map[string]AutoscalerStatusDetailsTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalerStatusDetailsTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AutoscalerStatusDetailsTypeEnum{}
	}

	items := make(map[string]AutoscalerStatusDetailsTypeEnum)
	for k, item := range a {
		items[k] = *flattenAutoscalerStatusDetailsTypeEnum(item.(interface{}))
	}

	return items
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

type autoscalerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         autoscalerApiOperation
}

func convertFieldDiffsToAutoscalerDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]autoscalerDiff, error) {
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
	var diffs []autoscalerDiff
	// For each operation name, create a autoscalerDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := autoscalerDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAutoscalerApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAutoscalerApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (autoscalerApiOperation, error) {
	switch opName {

	case "updateAutoscalerUpdateOperation":
		return &updateAutoscalerUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
