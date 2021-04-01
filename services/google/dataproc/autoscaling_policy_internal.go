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
package dataproc

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
)

func (r *AutoscalingPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "basicAlgorithm"); err != nil {
		return err
	}
	if err := dcl.Required(r, "workerConfig"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BasicAlgorithm) {
		if err := r.BasicAlgorithm.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WorkerConfig) {
		if err := r.WorkerConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryWorkerConfig) {
		if err := r.SecondaryWorkerConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AutoscalingPolicyBasicAlgorithm) validate() error {
	if err := dcl.Required(r, "yarnConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.YarnConfig) {
		if err := r.YarnConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AutoscalingPolicyBasicAlgorithmYarnConfig) validate() error {
	if err := dcl.Required(r, "gracefulDecommissionTimeout"); err != nil {
		return err
	}
	if err := dcl.Required(r, "scaleUpFactor"); err != nil {
		return err
	}
	if err := dcl.Required(r, "scaleDownFactor"); err != nil {
		return err
	}
	return nil
}
func (r *AutoscalingPolicyWorkerConfig) validate() error {
	if err := dcl.Required(r, "maxInstances"); err != nil {
		return err
	}
	return nil
}
func (r *AutoscalingPolicySecondaryWorkerConfig) validate() error {
	return nil
}

func autoscalingPolicyGetURL(userBasePath string, r *AutoscalingPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
}

func autoscalingPolicyListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/autoscalingPolicies", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func autoscalingPolicyCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/autoscalingPolicies", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func autoscalingPolicyDeleteURL(userBasePath string, r *AutoscalingPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
}

// autoscalingPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type autoscalingPolicyApiOperation interface {
	do(context.Context, *AutoscalingPolicy, *Client) error
}

// newUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest creates a request for an
// AutoscalingPolicy resource's UpdateAutoscalingPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest(ctx context.Context, f *AutoscalingPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandAutoscalingPolicyBasicAlgorithm(c, f.BasicAlgorithm); err != nil {
		return nil, fmt.Errorf("error expanding BasicAlgorithm into basicAlgorithm: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["basicAlgorithm"] = v
	}
	if v, err := expandAutoscalingPolicyWorkerConfig(c, f.WorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["workerConfig"] = v
	}
	if v, err := expandAutoscalingPolicySecondaryWorkerConfig(c, f.SecondaryWorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryWorkerConfig into secondaryWorkerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["secondaryWorkerConfig"] = v
	}
	if v, err := dcl.DeriveField("%s", f.Name); err != nil {
		return nil, err
	} else {
		req["id"] = v
	}

	return req, nil
}

// marshalUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAutoscalingPolicyUpdateAutoscalingPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAutoscalingPolicyUpdateAutoscalingPolicyOperation) do(ctx context.Context, r *AutoscalingPolicy, c *Client) error {
	_, err := c.GetAutoscalingPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAutoscalingPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAutoscalingPolicyUpdateAutoscalingPolicyRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAutoscalingPolicyRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := autoscalingPolicyListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AutoscalingPolicyMaxPage {
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

type listAutoscalingPolicyOperation struct {
	Policies []map[string]interface{} `json:"policies"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listAutoscalingPolicy(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*AutoscalingPolicy, string, error) {
	b, err := c.listAutoscalingPolicyRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAutoscalingPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AutoscalingPolicy
	for _, v := range m.Policies {
		res := flattenAutoscalingPolicy(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAutoscalingPolicy(ctx context.Context, f func(*AutoscalingPolicy) bool, resources []*AutoscalingPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAutoscalingPolicy(ctx, res)
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

type deleteAutoscalingPolicyOperation struct{}

func (op *deleteAutoscalingPolicyOperation) do(ctx context.Context, r *AutoscalingPolicy, c *Client) error {

	_, err := c.GetAutoscalingPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AutoscalingPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAutoscalingPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := autoscalingPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete AutoscalingPolicy: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAutoscalingPolicy(ctx, r.urlNormalized())
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
type createAutoscalingPolicyOperation struct {
	response map[string]interface{}
}

func (op *createAutoscalingPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAutoscalingPolicyOperation) do(ctx context.Context, r *AutoscalingPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := autoscalingPolicyCreateURL(c.Config.BasePath, project, location)

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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetAutoscalingPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAutoscalingPolicyRaw(ctx context.Context, r *AutoscalingPolicy) ([]byte, error) {

	u, err := autoscalingPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) autoscalingPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AutoscalingPolicy, opts ...dcl.ApplyOption) (initial, desired *AutoscalingPolicy, diffs []autoscalingPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AutoscalingPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AutoscalingPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AutoscalingPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAutoscalingPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AutoscalingPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AutoscalingPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that AutoscalingPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAutoscalingPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for AutoscalingPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AutoscalingPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAutoscalingPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AutoscalingPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAutoscalingPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AutoscalingPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAutoscalingPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAutoscalingPolicyInitialState(rawInitial, rawDesired *AutoscalingPolicy) (*AutoscalingPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAutoscalingPolicyDesiredState(rawDesired, rawInitial *AutoscalingPolicy, opts ...dcl.ApplyOption) (*AutoscalingPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BasicAlgorithm = canonicalizeAutoscalingPolicyBasicAlgorithm(rawDesired.BasicAlgorithm, nil, opts...)
		rawDesired.WorkerConfig = canonicalizeAutoscalingPolicyWorkerConfig(rawDesired.WorkerConfig, nil, opts...)
		rawDesired.SecondaryWorkerConfig = canonicalizeAutoscalingPolicySecondaryWorkerConfig(rawDesired.SecondaryWorkerConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.BasicAlgorithm = canonicalizeAutoscalingPolicyBasicAlgorithm(rawDesired.BasicAlgorithm, rawInitial.BasicAlgorithm, opts...)
	rawDesired.WorkerConfig = canonicalizeAutoscalingPolicyWorkerConfig(rawDesired.WorkerConfig, rawInitial.WorkerConfig, opts...)
	rawDesired.SecondaryWorkerConfig = canonicalizeAutoscalingPolicySecondaryWorkerConfig(rawDesired.SecondaryWorkerConfig, rawInitial.SecondaryWorkerConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeAutoscalingPolicyNewState(c *Client, rawNew, rawDesired *AutoscalingPolicy) (*AutoscalingPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.BasicAlgorithm) && dcl.IsEmptyValueIndirect(rawDesired.BasicAlgorithm) {
		rawNew.BasicAlgorithm = rawDesired.BasicAlgorithm
	} else {
		rawNew.BasicAlgorithm = canonicalizeNewAutoscalingPolicyBasicAlgorithm(c, rawDesired.BasicAlgorithm, rawNew.BasicAlgorithm)
	}

	if dcl.IsEmptyValueIndirect(rawNew.WorkerConfig) && dcl.IsEmptyValueIndirect(rawDesired.WorkerConfig) {
		rawNew.WorkerConfig = rawDesired.WorkerConfig
	} else {
		rawNew.WorkerConfig = canonicalizeNewAutoscalingPolicyWorkerConfig(c, rawDesired.WorkerConfig, rawNew.WorkerConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SecondaryWorkerConfig) && dcl.IsEmptyValueIndirect(rawDesired.SecondaryWorkerConfig) {
		rawNew.SecondaryWorkerConfig = rawDesired.SecondaryWorkerConfig
	} else {
		rawNew.SecondaryWorkerConfig = canonicalizeNewAutoscalingPolicySecondaryWorkerConfig(c, rawDesired.SecondaryWorkerConfig, rawNew.SecondaryWorkerConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAutoscalingPolicyBasicAlgorithm(des, initial *AutoscalingPolicyBasicAlgorithm, opts ...dcl.ApplyOption) *AutoscalingPolicyBasicAlgorithm {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.YarnConfig = canonicalizeAutoscalingPolicyBasicAlgorithmYarnConfig(des.YarnConfig, initial.YarnConfig, opts...)
	if dcl.StringCanonicalize(des.CooldownPeriod, initial.CooldownPeriod) || dcl.IsZeroValue(des.CooldownPeriod) {
		des.CooldownPeriod = initial.CooldownPeriod
	}

	return des
}

func canonicalizeNewAutoscalingPolicyBasicAlgorithm(c *Client, des, nw *AutoscalingPolicyBasicAlgorithm) *AutoscalingPolicyBasicAlgorithm {
	if des == nil || nw == nil {
		return nw
	}

	nw.YarnConfig = canonicalizeNewAutoscalingPolicyBasicAlgorithmYarnConfig(c, des.YarnConfig, nw.YarnConfig)
	if dcl.StringCanonicalize(des.CooldownPeriod, nw.CooldownPeriod) {
		nw.CooldownPeriod = des.CooldownPeriod
	}

	return nw
}

func canonicalizeNewAutoscalingPolicyBasicAlgorithmSet(c *Client, des, nw []AutoscalingPolicyBasicAlgorithm) []AutoscalingPolicyBasicAlgorithm {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalingPolicyBasicAlgorithm
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalingPolicyBasicAlgorithm(c, &d, &n) {
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

func canonicalizeNewAutoscalingPolicyBasicAlgorithmSlice(c *Client, des, nw []AutoscalingPolicyBasicAlgorithm) []AutoscalingPolicyBasicAlgorithm {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalingPolicyBasicAlgorithm
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalingPolicyBasicAlgorithm(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalingPolicyBasicAlgorithmYarnConfig(des, initial *AutoscalingPolicyBasicAlgorithmYarnConfig, opts ...dcl.ApplyOption) *AutoscalingPolicyBasicAlgorithmYarnConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.GracefulDecommissionTimeout, initial.GracefulDecommissionTimeout) || dcl.IsZeroValue(des.GracefulDecommissionTimeout) {
		des.GracefulDecommissionTimeout = initial.GracefulDecommissionTimeout
	}
	if dcl.IsZeroValue(des.ScaleUpFactor) {
		des.ScaleUpFactor = initial.ScaleUpFactor
	}
	if dcl.IsZeroValue(des.ScaleDownFactor) {
		des.ScaleDownFactor = initial.ScaleDownFactor
	}
	if dcl.IsZeroValue(des.ScaleUpMinWorkerFraction) {
		des.ScaleUpMinWorkerFraction = initial.ScaleUpMinWorkerFraction
	}
	if dcl.IsZeroValue(des.ScaleDownMinWorkerFraction) {
		des.ScaleDownMinWorkerFraction = initial.ScaleDownMinWorkerFraction
	}

	return des
}

func canonicalizeNewAutoscalingPolicyBasicAlgorithmYarnConfig(c *Client, des, nw *AutoscalingPolicyBasicAlgorithmYarnConfig) *AutoscalingPolicyBasicAlgorithmYarnConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.GracefulDecommissionTimeout, nw.GracefulDecommissionTimeout) {
		nw.GracefulDecommissionTimeout = des.GracefulDecommissionTimeout
	}

	return nw
}

func canonicalizeNewAutoscalingPolicyBasicAlgorithmYarnConfigSet(c *Client, des, nw []AutoscalingPolicyBasicAlgorithmYarnConfig) []AutoscalingPolicyBasicAlgorithmYarnConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalingPolicyBasicAlgorithmYarnConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalingPolicyBasicAlgorithmYarnConfig(c, &d, &n) {
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

func canonicalizeNewAutoscalingPolicyBasicAlgorithmYarnConfigSlice(c *Client, des, nw []AutoscalingPolicyBasicAlgorithmYarnConfig) []AutoscalingPolicyBasicAlgorithmYarnConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalingPolicyBasicAlgorithmYarnConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalingPolicyBasicAlgorithmYarnConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalingPolicyWorkerConfig(des, initial *AutoscalingPolicyWorkerConfig, opts ...dcl.ApplyOption) *AutoscalingPolicyWorkerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinInstances) {
		des.MinInstances = initial.MinInstances
	}
	if dcl.IsZeroValue(des.MaxInstances) {
		des.MaxInstances = initial.MaxInstances
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	}

	return des
}

func canonicalizeNewAutoscalingPolicyWorkerConfig(c *Client, des, nw *AutoscalingPolicyWorkerConfig) *AutoscalingPolicyWorkerConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAutoscalingPolicyWorkerConfigSet(c *Client, des, nw []AutoscalingPolicyWorkerConfig) []AutoscalingPolicyWorkerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalingPolicyWorkerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalingPolicyWorkerConfig(c, &d, &n) {
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

func canonicalizeNewAutoscalingPolicyWorkerConfigSlice(c *Client, des, nw []AutoscalingPolicyWorkerConfig) []AutoscalingPolicyWorkerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalingPolicyWorkerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalingPolicyWorkerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAutoscalingPolicySecondaryWorkerConfig(des, initial *AutoscalingPolicySecondaryWorkerConfig, opts ...dcl.ApplyOption) *AutoscalingPolicySecondaryWorkerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinInstances) {
		des.MinInstances = initial.MinInstances
	}
	if dcl.IsZeroValue(des.MaxInstances) {
		des.MaxInstances = initial.MaxInstances
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	}

	return des
}

func canonicalizeNewAutoscalingPolicySecondaryWorkerConfig(c *Client, des, nw *AutoscalingPolicySecondaryWorkerConfig) *AutoscalingPolicySecondaryWorkerConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAutoscalingPolicySecondaryWorkerConfigSet(c *Client, des, nw []AutoscalingPolicySecondaryWorkerConfig) []AutoscalingPolicySecondaryWorkerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AutoscalingPolicySecondaryWorkerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAutoscalingPolicySecondaryWorkerConfig(c, &d, &n) {
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

func canonicalizeNewAutoscalingPolicySecondaryWorkerConfigSlice(c *Client, des, nw []AutoscalingPolicySecondaryWorkerConfig) []AutoscalingPolicySecondaryWorkerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AutoscalingPolicySecondaryWorkerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAutoscalingPolicySecondaryWorkerConfig(c, &d, &n))
	}

	return items
}

type autoscalingPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         autoscalingPolicyApiOperation
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
func diffAutoscalingPolicy(c *Client, desired, actual *AutoscalingPolicy, opts ...dcl.ApplyOption) ([]autoscalingPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []autoscalingPolicyDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, autoscalingPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareAutoscalingPolicyBasicAlgorithm(c, desired.BasicAlgorithm, actual.BasicAlgorithm) {
		c.Config.Logger.Infof("Detected diff in BasicAlgorithm.\nDESIRED: %v\nACTUAL: %v", desired.BasicAlgorithm, actual.BasicAlgorithm)

		diffs = append(diffs, autoscalingPolicyDiff{
			UpdateOp:  &updateAutoscalingPolicyUpdateAutoscalingPolicyOperation{},
			FieldName: "BasicAlgorithm",
		})

	}
	if compareAutoscalingPolicyWorkerConfig(c, desired.WorkerConfig, actual.WorkerConfig) {
		c.Config.Logger.Infof("Detected diff in WorkerConfig.\nDESIRED: %v\nACTUAL: %v", desired.WorkerConfig, actual.WorkerConfig)

		diffs = append(diffs, autoscalingPolicyDiff{
			UpdateOp:  &updateAutoscalingPolicyUpdateAutoscalingPolicyOperation{},
			FieldName: "WorkerConfig",
		})

	}
	if compareAutoscalingPolicySecondaryWorkerConfig(c, desired.SecondaryWorkerConfig, actual.SecondaryWorkerConfig) {
		c.Config.Logger.Infof("Detected diff in SecondaryWorkerConfig.\nDESIRED: %v\nACTUAL: %v", desired.SecondaryWorkerConfig, actual.SecondaryWorkerConfig)

		diffs = append(diffs, autoscalingPolicyDiff{
			UpdateOp:  &updateAutoscalingPolicyUpdateAutoscalingPolicyOperation{},
			FieldName: "SecondaryWorkerConfig",
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
	var deduped []autoscalingPolicyDiff
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
func compareAutoscalingPolicyBasicAlgorithm(c *Client, desired, actual *AutoscalingPolicyBasicAlgorithm) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.YarnConfig == nil && desired.YarnConfig != nil && !dcl.IsEmptyValueIndirect(desired.YarnConfig) {
		c.Config.Logger.Infof("desired YarnConfig %s - but actually nil", dcl.SprintResource(desired.YarnConfig))
		return true
	}
	if compareAutoscalingPolicyBasicAlgorithmYarnConfig(c, desired.YarnConfig, actual.YarnConfig) && !dcl.IsZeroValue(desired.YarnConfig) {
		c.Config.Logger.Infof("Diff in YarnConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.YarnConfig), dcl.SprintResource(actual.YarnConfig))
		return true
	}
	if actual.CooldownPeriod == nil && desired.CooldownPeriod != nil && !dcl.IsEmptyValueIndirect(desired.CooldownPeriod) {
		c.Config.Logger.Infof("desired CooldownPeriod %s - but actually nil", dcl.SprintResource(desired.CooldownPeriod))
		return true
	}
	if !dcl.StringCanonicalize(desired.CooldownPeriod, actual.CooldownPeriod) && !dcl.IsZeroValue(desired.CooldownPeriod) {
		c.Config.Logger.Infof("Diff in CooldownPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CooldownPeriod), dcl.SprintResource(actual.CooldownPeriod))
		return true
	}
	return false
}

func compareAutoscalingPolicyBasicAlgorithmSlice(c *Client, desired, actual []AutoscalingPolicyBasicAlgorithm) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyBasicAlgorithm, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalingPolicyBasicAlgorithm(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithm, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicyBasicAlgorithmMap(c *Client, desired, actual map[string]AutoscalingPolicyBasicAlgorithm) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyBasicAlgorithm, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithm, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalingPolicyBasicAlgorithm(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithm, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicyBasicAlgorithmYarnConfig(c *Client, desired, actual *AutoscalingPolicyBasicAlgorithmYarnConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.GracefulDecommissionTimeout == nil && desired.GracefulDecommissionTimeout != nil && !dcl.IsEmptyValueIndirect(desired.GracefulDecommissionTimeout) {
		c.Config.Logger.Infof("desired GracefulDecommissionTimeout %s - but actually nil", dcl.SprintResource(desired.GracefulDecommissionTimeout))
		return true
	}
	if !dcl.StringCanonicalize(desired.GracefulDecommissionTimeout, actual.GracefulDecommissionTimeout) && !dcl.IsZeroValue(desired.GracefulDecommissionTimeout) {
		c.Config.Logger.Infof("Diff in GracefulDecommissionTimeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GracefulDecommissionTimeout), dcl.SprintResource(actual.GracefulDecommissionTimeout))
		return true
	}
	if actual.ScaleUpFactor == nil && desired.ScaleUpFactor != nil && !dcl.IsEmptyValueIndirect(desired.ScaleUpFactor) {
		c.Config.Logger.Infof("desired ScaleUpFactor %s - but actually nil", dcl.SprintResource(desired.ScaleUpFactor))
		return true
	}
	if !reflect.DeepEqual(desired.ScaleUpFactor, actual.ScaleUpFactor) && !dcl.IsZeroValue(desired.ScaleUpFactor) {
		c.Config.Logger.Infof("Diff in ScaleUpFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScaleUpFactor), dcl.SprintResource(actual.ScaleUpFactor))
		return true
	}
	if actual.ScaleDownFactor == nil && desired.ScaleDownFactor != nil && !dcl.IsEmptyValueIndirect(desired.ScaleDownFactor) {
		c.Config.Logger.Infof("desired ScaleDownFactor %s - but actually nil", dcl.SprintResource(desired.ScaleDownFactor))
		return true
	}
	if !reflect.DeepEqual(desired.ScaleDownFactor, actual.ScaleDownFactor) && !dcl.IsZeroValue(desired.ScaleDownFactor) {
		c.Config.Logger.Infof("Diff in ScaleDownFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScaleDownFactor), dcl.SprintResource(actual.ScaleDownFactor))
		return true
	}
	if actual.ScaleUpMinWorkerFraction == nil && desired.ScaleUpMinWorkerFraction != nil && !dcl.IsEmptyValueIndirect(desired.ScaleUpMinWorkerFraction) {
		c.Config.Logger.Infof("desired ScaleUpMinWorkerFraction %s - but actually nil", dcl.SprintResource(desired.ScaleUpMinWorkerFraction))
		return true
	}
	if !reflect.DeepEqual(desired.ScaleUpMinWorkerFraction, actual.ScaleUpMinWorkerFraction) && !dcl.IsZeroValue(desired.ScaleUpMinWorkerFraction) {
		c.Config.Logger.Infof("Diff in ScaleUpMinWorkerFraction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScaleUpMinWorkerFraction), dcl.SprintResource(actual.ScaleUpMinWorkerFraction))
		return true
	}
	if actual.ScaleDownMinWorkerFraction == nil && desired.ScaleDownMinWorkerFraction != nil && !dcl.IsEmptyValueIndirect(desired.ScaleDownMinWorkerFraction) {
		c.Config.Logger.Infof("desired ScaleDownMinWorkerFraction %s - but actually nil", dcl.SprintResource(desired.ScaleDownMinWorkerFraction))
		return true
	}
	if !reflect.DeepEqual(desired.ScaleDownMinWorkerFraction, actual.ScaleDownMinWorkerFraction) && !dcl.IsZeroValue(desired.ScaleDownMinWorkerFraction) {
		c.Config.Logger.Infof("Diff in ScaleDownMinWorkerFraction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScaleDownMinWorkerFraction), dcl.SprintResource(actual.ScaleDownMinWorkerFraction))
		return true
	}
	return false
}

func compareAutoscalingPolicyBasicAlgorithmYarnConfigSlice(c *Client, desired, actual []AutoscalingPolicyBasicAlgorithmYarnConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyBasicAlgorithmYarnConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalingPolicyBasicAlgorithmYarnConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithmYarnConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicyBasicAlgorithmYarnConfigMap(c *Client, desired, actual map[string]AutoscalingPolicyBasicAlgorithmYarnConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyBasicAlgorithmYarnConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithmYarnConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalingPolicyBasicAlgorithmYarnConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyBasicAlgorithmYarnConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicyWorkerConfig(c *Client, desired, actual *AutoscalingPolicyWorkerConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinInstances == nil && desired.MinInstances != nil && !dcl.IsEmptyValueIndirect(desired.MinInstances) {
		c.Config.Logger.Infof("desired MinInstances %s - but actually nil", dcl.SprintResource(desired.MinInstances))
		return true
	}
	if !reflect.DeepEqual(desired.MinInstances, actual.MinInstances) && !dcl.IsZeroValue(desired.MinInstances) {
		c.Config.Logger.Infof("Diff in MinInstances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinInstances), dcl.SprintResource(actual.MinInstances))
		return true
	}
	if actual.MaxInstances == nil && desired.MaxInstances != nil && !dcl.IsEmptyValueIndirect(desired.MaxInstances) {
		c.Config.Logger.Infof("desired MaxInstances %s - but actually nil", dcl.SprintResource(desired.MaxInstances))
		return true
	}
	if !reflect.DeepEqual(desired.MaxInstances, actual.MaxInstances) && !dcl.IsZeroValue(desired.MaxInstances) {
		c.Config.Logger.Infof("Diff in MaxInstances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxInstances), dcl.SprintResource(actual.MaxInstances))
		return true
	}
	if actual.Weight == nil && desired.Weight != nil && !dcl.IsEmptyValueIndirect(desired.Weight) {
		c.Config.Logger.Infof("desired Weight %s - but actually nil", dcl.SprintResource(desired.Weight))
		return true
	}
	if !reflect.DeepEqual(desired.Weight, actual.Weight) && !dcl.IsZeroValue(desired.Weight) {
		c.Config.Logger.Infof("Diff in Weight. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Weight), dcl.SprintResource(actual.Weight))
		return true
	}
	return false
}

func compareAutoscalingPolicyWorkerConfigSlice(c *Client, desired, actual []AutoscalingPolicyWorkerConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyWorkerConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalingPolicyWorkerConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyWorkerConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicyWorkerConfigMap(c *Client, desired, actual map[string]AutoscalingPolicyWorkerConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicyWorkerConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyWorkerConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalingPolicyWorkerConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicyWorkerConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicySecondaryWorkerConfig(c *Client, desired, actual *AutoscalingPolicySecondaryWorkerConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinInstances == nil && desired.MinInstances != nil && !dcl.IsEmptyValueIndirect(desired.MinInstances) {
		c.Config.Logger.Infof("desired MinInstances %s - but actually nil", dcl.SprintResource(desired.MinInstances))
		return true
	}
	if !reflect.DeepEqual(desired.MinInstances, actual.MinInstances) && !dcl.IsZeroValue(desired.MinInstances) {
		c.Config.Logger.Infof("Diff in MinInstances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinInstances), dcl.SprintResource(actual.MinInstances))
		return true
	}
	if actual.MaxInstances == nil && desired.MaxInstances != nil && !dcl.IsEmptyValueIndirect(desired.MaxInstances) {
		c.Config.Logger.Infof("desired MaxInstances %s - but actually nil", dcl.SprintResource(desired.MaxInstances))
		return true
	}
	if !reflect.DeepEqual(desired.MaxInstances, actual.MaxInstances) && !dcl.IsZeroValue(desired.MaxInstances) {
		c.Config.Logger.Infof("Diff in MaxInstances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxInstances), dcl.SprintResource(actual.MaxInstances))
		return true
	}
	if actual.Weight == nil && desired.Weight != nil && !dcl.IsEmptyValueIndirect(desired.Weight) {
		c.Config.Logger.Infof("desired Weight %s - but actually nil", dcl.SprintResource(desired.Weight))
		return true
	}
	if !reflect.DeepEqual(desired.Weight, actual.Weight) && !dcl.IsZeroValue(desired.Weight) {
		c.Config.Logger.Infof("Diff in Weight. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Weight), dcl.SprintResource(actual.Weight))
		return true
	}
	return false
}

func compareAutoscalingPolicySecondaryWorkerConfigSlice(c *Client, desired, actual []AutoscalingPolicySecondaryWorkerConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicySecondaryWorkerConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAutoscalingPolicySecondaryWorkerConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicySecondaryWorkerConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAutoscalingPolicySecondaryWorkerConfigMap(c *Client, desired, actual map[string]AutoscalingPolicySecondaryWorkerConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AutoscalingPolicySecondaryWorkerConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AutoscalingPolicySecondaryWorkerConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAutoscalingPolicySecondaryWorkerConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AutoscalingPolicySecondaryWorkerConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AutoscalingPolicy) urlNormalized() *AutoscalingPolicy {
	normalized := deepcopy.Copy(*r).(AutoscalingPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AutoscalingPolicy) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *AutoscalingPolicy) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *AutoscalingPolicy) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *AutoscalingPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateAutoscalingPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AutoscalingPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AutoscalingPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandAutoscalingPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AutoscalingPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAutoscalingPolicy decodes JSON responses into the AutoscalingPolicy resource schema.
func unmarshalAutoscalingPolicy(b []byte, c *Client) (*AutoscalingPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAutoscalingPolicy(m, c)
}

func unmarshalMapAutoscalingPolicy(m map[string]interface{}, c *Client) (*AutoscalingPolicy, error) {

	return flattenAutoscalingPolicy(c, m), nil
}

// expandAutoscalingPolicy expands AutoscalingPolicy into a JSON request object.
func expandAutoscalingPolicy(c *Client, f *AutoscalingPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandAutoscalingPolicyBasicAlgorithm(c, f.BasicAlgorithm); err != nil {
		return nil, fmt.Errorf("error expanding BasicAlgorithm into basicAlgorithm: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["basicAlgorithm"] = v
	}
	if v, err := expandAutoscalingPolicyWorkerConfig(c, f.WorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workerConfig"] = v
	}
	if v, err := expandAutoscalingPolicySecondaryWorkerConfig(c, f.SecondaryWorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryWorkerConfig into secondaryWorkerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryWorkerConfig"] = v
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

// flattenAutoscalingPolicy flattens AutoscalingPolicy from a JSON request object into the
// AutoscalingPolicy type.
func flattenAutoscalingPolicy(c *Client, i interface{}) *AutoscalingPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &AutoscalingPolicy{}
	r.Name = dcl.FlattenString(m["id"])
	r.BasicAlgorithm = flattenAutoscalingPolicyBasicAlgorithm(c, m["basicAlgorithm"])
	r.WorkerConfig = flattenAutoscalingPolicyWorkerConfig(c, m["workerConfig"])
	r.SecondaryWorkerConfig = flattenAutoscalingPolicySecondaryWorkerConfig(c, m["secondaryWorkerConfig"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandAutoscalingPolicyBasicAlgorithmMap expands the contents of AutoscalingPolicyBasicAlgorithm into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithmMap(c *Client, f map[string]AutoscalingPolicyBasicAlgorithm) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalingPolicyBasicAlgorithm(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalingPolicyBasicAlgorithmSlice expands the contents of AutoscalingPolicyBasicAlgorithm into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithmSlice(c *Client, f []AutoscalingPolicyBasicAlgorithm) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalingPolicyBasicAlgorithm(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalingPolicyBasicAlgorithmMap flattens the contents of AutoscalingPolicyBasicAlgorithm from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithmMap(c *Client, i interface{}) map[string]AutoscalingPolicyBasicAlgorithm {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalingPolicyBasicAlgorithm{}
	}

	if len(a) == 0 {
		return map[string]AutoscalingPolicyBasicAlgorithm{}
	}

	items := make(map[string]AutoscalingPolicyBasicAlgorithm)
	for k, item := range a {
		items[k] = *flattenAutoscalingPolicyBasicAlgorithm(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalingPolicyBasicAlgorithmSlice flattens the contents of AutoscalingPolicyBasicAlgorithm from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithmSlice(c *Client, i interface{}) []AutoscalingPolicyBasicAlgorithm {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalingPolicyBasicAlgorithm{}
	}

	if len(a) == 0 {
		return []AutoscalingPolicyBasicAlgorithm{}
	}

	items := make([]AutoscalingPolicyBasicAlgorithm, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalingPolicyBasicAlgorithm(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalingPolicyBasicAlgorithm expands an instance of AutoscalingPolicyBasicAlgorithm into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithm(c *Client, f *AutoscalingPolicyBasicAlgorithm) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAutoscalingPolicyBasicAlgorithmYarnConfig(c, f.YarnConfig); err != nil {
		return nil, fmt.Errorf("error expanding YarnConfig into yarnConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yarnConfig"] = v
	}
	if v := f.CooldownPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["cooldownPeriod"] = v
	}

	return m, nil
}

// flattenAutoscalingPolicyBasicAlgorithm flattens an instance of AutoscalingPolicyBasicAlgorithm from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithm(c *Client, i interface{}) *AutoscalingPolicyBasicAlgorithm {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalingPolicyBasicAlgorithm{}
	r.YarnConfig = flattenAutoscalingPolicyBasicAlgorithmYarnConfig(c, m["yarnConfig"])
	r.CooldownPeriod = dcl.FlattenString(m["cooldownPeriod"])

	return r
}

// expandAutoscalingPolicyBasicAlgorithmYarnConfigMap expands the contents of AutoscalingPolicyBasicAlgorithmYarnConfig into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithmYarnConfigMap(c *Client, f map[string]AutoscalingPolicyBasicAlgorithmYarnConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalingPolicyBasicAlgorithmYarnConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalingPolicyBasicAlgorithmYarnConfigSlice expands the contents of AutoscalingPolicyBasicAlgorithmYarnConfig into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithmYarnConfigSlice(c *Client, f []AutoscalingPolicyBasicAlgorithmYarnConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalingPolicyBasicAlgorithmYarnConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalingPolicyBasicAlgorithmYarnConfigMap flattens the contents of AutoscalingPolicyBasicAlgorithmYarnConfig from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithmYarnConfigMap(c *Client, i interface{}) map[string]AutoscalingPolicyBasicAlgorithmYarnConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalingPolicyBasicAlgorithmYarnConfig{}
	}

	if len(a) == 0 {
		return map[string]AutoscalingPolicyBasicAlgorithmYarnConfig{}
	}

	items := make(map[string]AutoscalingPolicyBasicAlgorithmYarnConfig)
	for k, item := range a {
		items[k] = *flattenAutoscalingPolicyBasicAlgorithmYarnConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalingPolicyBasicAlgorithmYarnConfigSlice flattens the contents of AutoscalingPolicyBasicAlgorithmYarnConfig from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithmYarnConfigSlice(c *Client, i interface{}) []AutoscalingPolicyBasicAlgorithmYarnConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalingPolicyBasicAlgorithmYarnConfig{}
	}

	if len(a) == 0 {
		return []AutoscalingPolicyBasicAlgorithmYarnConfig{}
	}

	items := make([]AutoscalingPolicyBasicAlgorithmYarnConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalingPolicyBasicAlgorithmYarnConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalingPolicyBasicAlgorithmYarnConfig expands an instance of AutoscalingPolicyBasicAlgorithmYarnConfig into a JSON
// request object.
func expandAutoscalingPolicyBasicAlgorithmYarnConfig(c *Client, f *AutoscalingPolicyBasicAlgorithmYarnConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.GracefulDecommissionTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["gracefulDecommissionTimeout"] = v
	}
	if v := f.ScaleUpFactor; !dcl.IsEmptyValueIndirect(v) {
		m["scaleUpFactor"] = v
	}
	if v := f.ScaleDownFactor; !dcl.IsEmptyValueIndirect(v) {
		m["scaleDownFactor"] = v
	}
	if v := f.ScaleUpMinWorkerFraction; !dcl.IsEmptyValueIndirect(v) {
		m["scaleUpMinWorkerFraction"] = v
	}
	if v := f.ScaleDownMinWorkerFraction; !dcl.IsEmptyValueIndirect(v) {
		m["scaleDownMinWorkerFraction"] = v
	}

	return m, nil
}

// flattenAutoscalingPolicyBasicAlgorithmYarnConfig flattens an instance of AutoscalingPolicyBasicAlgorithmYarnConfig from a JSON
// response object.
func flattenAutoscalingPolicyBasicAlgorithmYarnConfig(c *Client, i interface{}) *AutoscalingPolicyBasicAlgorithmYarnConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalingPolicyBasicAlgorithmYarnConfig{}
	r.GracefulDecommissionTimeout = dcl.FlattenString(m["gracefulDecommissionTimeout"])
	r.ScaleUpFactor = dcl.FlattenDouble(m["scaleUpFactor"])
	r.ScaleDownFactor = dcl.FlattenDouble(m["scaleDownFactor"])
	r.ScaleUpMinWorkerFraction = dcl.FlattenDouble(m["scaleUpMinWorkerFraction"])
	r.ScaleDownMinWorkerFraction = dcl.FlattenDouble(m["scaleDownMinWorkerFraction"])

	return r
}

// expandAutoscalingPolicyWorkerConfigMap expands the contents of AutoscalingPolicyWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicyWorkerConfigMap(c *Client, f map[string]AutoscalingPolicyWorkerConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalingPolicyWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalingPolicyWorkerConfigSlice expands the contents of AutoscalingPolicyWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicyWorkerConfigSlice(c *Client, f []AutoscalingPolicyWorkerConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalingPolicyWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalingPolicyWorkerConfigMap flattens the contents of AutoscalingPolicyWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicyWorkerConfigMap(c *Client, i interface{}) map[string]AutoscalingPolicyWorkerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalingPolicyWorkerConfig{}
	}

	if len(a) == 0 {
		return map[string]AutoscalingPolicyWorkerConfig{}
	}

	items := make(map[string]AutoscalingPolicyWorkerConfig)
	for k, item := range a {
		items[k] = *flattenAutoscalingPolicyWorkerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalingPolicyWorkerConfigSlice flattens the contents of AutoscalingPolicyWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicyWorkerConfigSlice(c *Client, i interface{}) []AutoscalingPolicyWorkerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalingPolicyWorkerConfig{}
	}

	if len(a) == 0 {
		return []AutoscalingPolicyWorkerConfig{}
	}

	items := make([]AutoscalingPolicyWorkerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalingPolicyWorkerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalingPolicyWorkerConfig expands an instance of AutoscalingPolicyWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicyWorkerConfig(c *Client, f *AutoscalingPolicyWorkerConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minInstances"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}

	return m, nil
}

// flattenAutoscalingPolicyWorkerConfig flattens an instance of AutoscalingPolicyWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicyWorkerConfig(c *Client, i interface{}) *AutoscalingPolicyWorkerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalingPolicyWorkerConfig{}
	r.MinInstances = dcl.FlattenInteger(m["minInstances"])
	r.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	r.Weight = dcl.FlattenInteger(m["weight"])

	return r
}

// expandAutoscalingPolicySecondaryWorkerConfigMap expands the contents of AutoscalingPolicySecondaryWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicySecondaryWorkerConfigMap(c *Client, f map[string]AutoscalingPolicySecondaryWorkerConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAutoscalingPolicySecondaryWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAutoscalingPolicySecondaryWorkerConfigSlice expands the contents of AutoscalingPolicySecondaryWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicySecondaryWorkerConfigSlice(c *Client, f []AutoscalingPolicySecondaryWorkerConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAutoscalingPolicySecondaryWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAutoscalingPolicySecondaryWorkerConfigMap flattens the contents of AutoscalingPolicySecondaryWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicySecondaryWorkerConfigMap(c *Client, i interface{}) map[string]AutoscalingPolicySecondaryWorkerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AutoscalingPolicySecondaryWorkerConfig{}
	}

	if len(a) == 0 {
		return map[string]AutoscalingPolicySecondaryWorkerConfig{}
	}

	items := make(map[string]AutoscalingPolicySecondaryWorkerConfig)
	for k, item := range a {
		items[k] = *flattenAutoscalingPolicySecondaryWorkerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAutoscalingPolicySecondaryWorkerConfigSlice flattens the contents of AutoscalingPolicySecondaryWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicySecondaryWorkerConfigSlice(c *Client, i interface{}) []AutoscalingPolicySecondaryWorkerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AutoscalingPolicySecondaryWorkerConfig{}
	}

	if len(a) == 0 {
		return []AutoscalingPolicySecondaryWorkerConfig{}
	}

	items := make([]AutoscalingPolicySecondaryWorkerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAutoscalingPolicySecondaryWorkerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAutoscalingPolicySecondaryWorkerConfig expands an instance of AutoscalingPolicySecondaryWorkerConfig into a JSON
// request object.
func expandAutoscalingPolicySecondaryWorkerConfig(c *Client, f *AutoscalingPolicySecondaryWorkerConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minInstances"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}

	return m, nil
}

// flattenAutoscalingPolicySecondaryWorkerConfig flattens an instance of AutoscalingPolicySecondaryWorkerConfig from a JSON
// response object.
func flattenAutoscalingPolicySecondaryWorkerConfig(c *Client, i interface{}) *AutoscalingPolicySecondaryWorkerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AutoscalingPolicySecondaryWorkerConfig{}
	r.MinInstances = dcl.FlattenInteger(m["minInstances"])
	r.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	r.Weight = dcl.FlattenInteger(m["weight"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AutoscalingPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAutoscalingPolicy(b, c)
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
