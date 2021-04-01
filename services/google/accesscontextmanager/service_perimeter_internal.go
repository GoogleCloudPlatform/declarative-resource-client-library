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
package accesscontextmanager

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

func (r *ServicePerimeter) validate() error {

	if err := dcl.Required(r, "title"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Policy, "Policy"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServicePerimeterStatus) validate() error {
	if !dcl.IsEmptyValueIndirect(r.VPCAccessibleServices) {
		if err := r.VPCAccessibleServices.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServicePerimeterStatusVPCAccessibleServices) validate() error {
	return nil
}
func (r *ServicePerimeterSpec) validate() error {
	if !dcl.IsEmptyValueIndirect(r.VPCAccessibleServices) {
		if err := r.VPCAccessibleServices.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServicePerimeterSpecVPCAccessibleServices) validate() error {
	return nil
}

func servicePerimeterGetURL(userBasePath string, r *ServicePerimeter) (string, error) {
	params := map[string]interface{}{
		"policy": dcl.ValueOrEmptyString(r.Policy),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("accessPolicies/{{policy}}/servicePerimeters/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

func servicePerimeterListURL(userBasePath, policy string) (string, error) {
	params := map[string]interface{}{
		"policy": policy,
	}
	return dcl.URL("{{policy}}/servicePerimeters", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil

}

func servicePerimeterCreateURL(userBasePath, policy string) (string, error) {
	params := map[string]interface{}{
		"policy": policy,
	}
	return dcl.URL("accessPolicies/{{policy}}/servicePerimeters", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil

}

func servicePerimeterDeleteURL(userBasePath string, r *ServicePerimeter) (string, error) {
	params := map[string]interface{}{
		"policy": dcl.ValueOrEmptyString(r.Policy),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("accessPolicies/{{policy}}/servicePerimeters/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

// servicePerimeterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type servicePerimeterApiOperation interface {
	do(context.Context, *ServicePerimeter, *Client) error
}

// newUpdateServicePerimeterUpdateRequest creates a request for an
// ServicePerimeter resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServicePerimeterUpdateRequest(ctx context.Context, f *ServicePerimeter, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		req["title"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandServicePerimeterStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["status"] = v
	}
	if v, err := expandServicePerimeterSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["spec"] = v
	}
	return req, nil
}

// marshalUpdateServicePerimeterUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateServicePerimeterUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateServicePerimeterUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServicePerimeterUpdateOperation) do(ctx context.Context, r *ServicePerimeter, c *Client) error {
	_, err := c.GetServicePerimeter(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"title", "description", "status", "spec"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateServicePerimeterUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateServicePerimeterUpdateRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listServicePerimeterRaw(ctx context.Context, policy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := servicePerimeterListURL(c.Config.BasePath, policy)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServicePerimeterMaxPage {
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

type listServicePerimeterOperation struct {
	ServicePerimeters []map[string]interface{} `json:"servicePerimeters"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listServicePerimeter(ctx context.Context, policy, pageToken string, pageSize int32) ([]*ServicePerimeter, string, error) {
	b, err := c.listServicePerimeterRaw(ctx, policy, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServicePerimeterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ServicePerimeter
	for _, v := range m.ServicePerimeters {
		res := flattenServicePerimeter(c, v)
		res.Policy = &policy
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllServicePerimeter(ctx context.Context, f func(*ServicePerimeter) bool, resources []*ServicePerimeter) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteServicePerimeter(ctx, res)
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

type deleteServicePerimeterOperation struct{}

func (op *deleteServicePerimeterOperation) do(ctx context.Context, r *ServicePerimeter, c *Client) error {

	_, err := c.GetServicePerimeter(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ServicePerimeter not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetServicePerimeter checking for existence. error: %v", err)
		return err
	}

	u, err := servicePerimeterDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetServicePerimeter(ctx, r.urlNormalized())
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
type createServicePerimeterOperation struct {
	response map[string]interface{}
}

func (op *createServicePerimeterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServicePerimeterOperation) do(ctx context.Context, r *ServicePerimeter, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	policy := r.createFields()
	u, err := servicePerimeterCreateURL(c.Config.BasePath, policy)

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
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetServicePerimeter(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServicePerimeterRaw(ctx context.Context, r *ServicePerimeter) ([]byte, error) {
	if dcl.IsZeroValue(r.PerimeterType) {
		r.PerimeterType = ServicePerimeterPerimeterTypeEnumRef("PERIMETER_TYPE_REGULAR")
	}

	u, err := servicePerimeterGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) servicePerimeterDiffsForRawDesired(ctx context.Context, rawDesired *ServicePerimeter, opts ...dcl.ApplyOption) (initial, desired *ServicePerimeter, diffs []servicePerimeterDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ServicePerimeter
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ServicePerimeter); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ServicePerimeter, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetServicePerimeter(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ServicePerimeter resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ServicePerimeter resource: %v", err)
		}
		c.Config.Logger.Info("Found that ServicePerimeter resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServicePerimeterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ServicePerimeter: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ServicePerimeter: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServicePerimeterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ServicePerimeter: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServicePerimeterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ServicePerimeter: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffServicePerimeter(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServicePerimeterInitialState(rawInitial, rawDesired *ServicePerimeter) (*ServicePerimeter, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServicePerimeterDesiredState(rawDesired, rawInitial *ServicePerimeter, opts ...dcl.ApplyOption) (*ServicePerimeter, error) {

	if dcl.IsZeroValue(rawDesired.PerimeterType) {
		rawDesired.PerimeterType = ServicePerimeterPerimeterTypeEnumRef("PERIMETER_TYPE_REGULAR")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Status = canonicalizeServicePerimeterStatus(rawDesired.Status, nil, opts...)
		rawDesired.Spec = canonicalizeServicePerimeterSpec(rawDesired.Spec, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Title, rawInitial.Title) {
		rawDesired.Title = rawInitial.Title
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.PerimeterType) {
		rawDesired.PerimeterType = rawInitial.PerimeterType
	}
	rawDesired.Status = canonicalizeServicePerimeterStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.NameToSelfLink(rawDesired.Policy, rawInitial.Policy) {
		rawDesired.Policy = rawInitial.Policy
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.BoolCanonicalize(rawDesired.UseExplicitDryRunSpec, rawInitial.UseExplicitDryRunSpec) {
		rawDesired.UseExplicitDryRunSpec = rawInitial.UseExplicitDryRunSpec
	}
	rawDesired.Spec = canonicalizeServicePerimeterSpec(rawDesired.Spec, rawInitial.Spec, opts...)

	return rawDesired, nil
}

func canonicalizeServicePerimeterNewState(c *Client, rawNew, rawDesired *ServicePerimeter) (*ServicePerimeter, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Title) && dcl.IsEmptyValueIndirect(rawDesired.Title) {
		rawNew.Title = rawDesired.Title
	} else {
		if dcl.StringCanonicalize(rawDesired.Title, rawNew.Title) {
			rawNew.Title = rawDesired.Title
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PerimeterType) && dcl.IsEmptyValueIndirect(rawDesired.PerimeterType) {
		rawNew.PerimeterType = rawDesired.PerimeterType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewServicePerimeterStatus(c, rawDesired.Status, rawNew.Status)
	}

	rawNew.Policy = rawDesired.Policy

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UseExplicitDryRunSpec) && dcl.IsEmptyValueIndirect(rawDesired.UseExplicitDryRunSpec) {
		rawNew.UseExplicitDryRunSpec = rawDesired.UseExplicitDryRunSpec
	} else {
		if dcl.BoolCanonicalize(rawDesired.UseExplicitDryRunSpec, rawNew.UseExplicitDryRunSpec) {
			rawNew.UseExplicitDryRunSpec = rawDesired.UseExplicitDryRunSpec
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Spec) && dcl.IsEmptyValueIndirect(rawDesired.Spec) {
		rawNew.Spec = rawDesired.Spec
	} else {
		rawNew.Spec = canonicalizeNewServicePerimeterSpec(c, rawDesired.Spec, rawNew.Spec)
	}

	return rawNew, nil
}

func canonicalizeServicePerimeterStatus(des, initial *ServicePerimeterStatus, opts ...dcl.ApplyOption) *ServicePerimeterStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Resources) {
		des.Resources = initial.Resources
	}
	if dcl.IsZeroValue(des.AccessLevels) {
		des.AccessLevels = initial.AccessLevels
	}
	if dcl.IsZeroValue(des.RestrictedServices) {
		des.RestrictedServices = initial.RestrictedServices
	}
	des.VPCAccessibleServices = canonicalizeServicePerimeterStatusVPCAccessibleServices(des.VPCAccessibleServices, initial.VPCAccessibleServices, opts...)

	return des
}

func canonicalizeNewServicePerimeterStatus(c *Client, des, nw *ServicePerimeterStatus) *ServicePerimeterStatus {
	if des == nil || nw == nil {
		return nw
	}

	nw.VPCAccessibleServices = canonicalizeNewServicePerimeterStatusVPCAccessibleServices(c, des.VPCAccessibleServices, nw.VPCAccessibleServices)

	return nw
}

func canonicalizeNewServicePerimeterStatusSet(c *Client, des, nw []ServicePerimeterStatus) []ServicePerimeterStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []ServicePerimeterStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServicePerimeterStatus(c, &d, &n) {
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

func canonicalizeNewServicePerimeterStatusSlice(c *Client, des, nw []ServicePerimeterStatus) []ServicePerimeterStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServicePerimeterStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServicePerimeterStatus(c, &d, &n))
	}

	return items
}

func canonicalizeServicePerimeterStatusVPCAccessibleServices(des, initial *ServicePerimeterStatusVPCAccessibleServices, opts ...dcl.ApplyOption) *ServicePerimeterStatusVPCAccessibleServices {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableRestriction, initial.EnableRestriction) || dcl.IsZeroValue(des.EnableRestriction) {
		des.EnableRestriction = initial.EnableRestriction
	}
	if dcl.IsZeroValue(des.AllowedServices) {
		des.AllowedServices = initial.AllowedServices
	}

	return des
}

func canonicalizeNewServicePerimeterStatusVPCAccessibleServices(c *Client, des, nw *ServicePerimeterStatusVPCAccessibleServices) *ServicePerimeterStatusVPCAccessibleServices {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableRestriction, nw.EnableRestriction) {
		nw.EnableRestriction = des.EnableRestriction
	}

	return nw
}

func canonicalizeNewServicePerimeterStatusVPCAccessibleServicesSet(c *Client, des, nw []ServicePerimeterStatusVPCAccessibleServices) []ServicePerimeterStatusVPCAccessibleServices {
	if des == nil {
		return nw
	}
	var reorderedNew []ServicePerimeterStatusVPCAccessibleServices
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServicePerimeterStatusVPCAccessibleServices(c, &d, &n) {
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

func canonicalizeNewServicePerimeterStatusVPCAccessibleServicesSlice(c *Client, des, nw []ServicePerimeterStatusVPCAccessibleServices) []ServicePerimeterStatusVPCAccessibleServices {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServicePerimeterStatusVPCAccessibleServices
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServicePerimeterStatusVPCAccessibleServices(c, &d, &n))
	}

	return items
}

func canonicalizeServicePerimeterSpec(des, initial *ServicePerimeterSpec, opts ...dcl.ApplyOption) *ServicePerimeterSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Resources) {
		des.Resources = initial.Resources
	}
	if dcl.IsZeroValue(des.AccessLevels) {
		des.AccessLevels = initial.AccessLevels
	}
	if dcl.IsZeroValue(des.RestrictedServices) {
		des.RestrictedServices = initial.RestrictedServices
	}
	des.VPCAccessibleServices = canonicalizeServicePerimeterSpecVPCAccessibleServices(des.VPCAccessibleServices, initial.VPCAccessibleServices, opts...)

	return des
}

func canonicalizeNewServicePerimeterSpec(c *Client, des, nw *ServicePerimeterSpec) *ServicePerimeterSpec {
	if des == nil || nw == nil {
		return nw
	}

	nw.VPCAccessibleServices = canonicalizeNewServicePerimeterSpecVPCAccessibleServices(c, des.VPCAccessibleServices, nw.VPCAccessibleServices)

	return nw
}

func canonicalizeNewServicePerimeterSpecSet(c *Client, des, nw []ServicePerimeterSpec) []ServicePerimeterSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ServicePerimeterSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServicePerimeterSpec(c, &d, &n) {
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

func canonicalizeNewServicePerimeterSpecSlice(c *Client, des, nw []ServicePerimeterSpec) []ServicePerimeterSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServicePerimeterSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServicePerimeterSpec(c, &d, &n))
	}

	return items
}

func canonicalizeServicePerimeterSpecVPCAccessibleServices(des, initial *ServicePerimeterSpecVPCAccessibleServices, opts ...dcl.ApplyOption) *ServicePerimeterSpecVPCAccessibleServices {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableRestriction, initial.EnableRestriction) || dcl.IsZeroValue(des.EnableRestriction) {
		des.EnableRestriction = initial.EnableRestriction
	}
	if dcl.IsZeroValue(des.AllowedServices) {
		des.AllowedServices = initial.AllowedServices
	}

	return des
}

func canonicalizeNewServicePerimeterSpecVPCAccessibleServices(c *Client, des, nw *ServicePerimeterSpecVPCAccessibleServices) *ServicePerimeterSpecVPCAccessibleServices {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableRestriction, nw.EnableRestriction) {
		nw.EnableRestriction = des.EnableRestriction
	}

	return nw
}

func canonicalizeNewServicePerimeterSpecVPCAccessibleServicesSet(c *Client, des, nw []ServicePerimeterSpecVPCAccessibleServices) []ServicePerimeterSpecVPCAccessibleServices {
	if des == nil {
		return nw
	}
	var reorderedNew []ServicePerimeterSpecVPCAccessibleServices
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServicePerimeterSpecVPCAccessibleServices(c, &d, &n) {
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

func canonicalizeNewServicePerimeterSpecVPCAccessibleServicesSlice(c *Client, des, nw []ServicePerimeterSpecVPCAccessibleServices) []ServicePerimeterSpecVPCAccessibleServices {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServicePerimeterSpecVPCAccessibleServices
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServicePerimeterSpecVPCAccessibleServices(c, &d, &n))
	}

	return items
}

type servicePerimeterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         servicePerimeterApiOperation
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
func diffServicePerimeter(c *Client, desired, actual *ServicePerimeter, opts ...dcl.ApplyOption) ([]servicePerimeterDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []servicePerimeterDiff
	if !dcl.IsZeroValue(desired.Title) && !dcl.StringCanonicalize(desired.Title, actual.Title) {
		c.Config.Logger.Infof("Detected diff in Title.\nDESIRED: %v\nACTUAL: %v", desired.Title, actual.Title)

		diffs = append(diffs, servicePerimeterDiff{
			UpdateOp:  &updateServicePerimeterUpdateOperation{},
			FieldName: "Title",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, servicePerimeterDiff{
			UpdateOp:  &updateServicePerimeterUpdateOperation{},
			FieldName: "Description",
		})

	}
	if !reflect.DeepEqual(desired.PerimeterType, actual.PerimeterType) {
		c.Config.Logger.Infof("Detected diff in PerimeterType.\nDESIRED: %v\nACTUAL: %v", desired.PerimeterType, actual.PerimeterType)
		diffs = append(diffs, servicePerimeterDiff{
			RequiresRecreate: true,
			FieldName:        "PerimeterType",
		})
	}
	if compareServicePerimeterStatus(c, desired.Status, actual.Status) {
		c.Config.Logger.Infof("Detected diff in Status.\nDESIRED: %v\nACTUAL: %v", desired.Status, actual.Status)

		diffs = append(diffs, servicePerimeterDiff{
			UpdateOp:  &updateServicePerimeterUpdateOperation{},
			FieldName: "Status",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, servicePerimeterDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.UseExplicitDryRunSpec) && !dcl.BoolCanonicalize(desired.UseExplicitDryRunSpec, actual.UseExplicitDryRunSpec) {
		c.Config.Logger.Infof("Detected diff in UseExplicitDryRunSpec.\nDESIRED: %v\nACTUAL: %v", desired.UseExplicitDryRunSpec, actual.UseExplicitDryRunSpec)
		diffs = append(diffs, servicePerimeterDiff{
			RequiresRecreate: true,
			FieldName:        "UseExplicitDryRunSpec",
		})
	}
	if compareServicePerimeterSpec(c, desired.Spec, actual.Spec) {
		c.Config.Logger.Infof("Detected diff in Spec.\nDESIRED: %v\nACTUAL: %v", desired.Spec, actual.Spec)

		diffs = append(diffs, servicePerimeterDiff{
			UpdateOp:  &updateServicePerimeterUpdateOperation{},
			FieldName: "Spec",
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
	var deduped []servicePerimeterDiff
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
func compareServicePerimeterStatus(c *Client, desired, actual *ServicePerimeterStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Resources == nil && desired.Resources != nil && !dcl.IsEmptyValueIndirect(desired.Resources) {
		c.Config.Logger.Infof("desired Resources %s - but actually nil", dcl.SprintResource(desired.Resources))
		return true
	}
	if !dcl.StringSliceEquals(desired.Resources, actual.Resources) && !dcl.IsZeroValue(desired.Resources) {
		c.Config.Logger.Infof("Diff in Resources. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Resources), dcl.SprintResource(actual.Resources))
		return true
	}
	if actual.AccessLevels == nil && desired.AccessLevels != nil && !dcl.IsEmptyValueIndirect(desired.AccessLevels) {
		c.Config.Logger.Infof("desired AccessLevels %s - but actually nil", dcl.SprintResource(desired.AccessLevels))
		return true
	}
	if !dcl.StringSliceEquals(desired.AccessLevels, actual.AccessLevels) && !dcl.IsZeroValue(desired.AccessLevels) {
		c.Config.Logger.Infof("Diff in AccessLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AccessLevels), dcl.SprintResource(actual.AccessLevels))
		return true
	}
	if actual.RestrictedServices == nil && desired.RestrictedServices != nil && !dcl.IsEmptyValueIndirect(desired.RestrictedServices) {
		c.Config.Logger.Infof("desired RestrictedServices %s - but actually nil", dcl.SprintResource(desired.RestrictedServices))
		return true
	}
	if !dcl.StringSliceEquals(desired.RestrictedServices, actual.RestrictedServices) && !dcl.IsZeroValue(desired.RestrictedServices) {
		c.Config.Logger.Infof("Diff in RestrictedServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RestrictedServices), dcl.SprintResource(actual.RestrictedServices))
		return true
	}
	if actual.VPCAccessibleServices == nil && desired.VPCAccessibleServices != nil && !dcl.IsEmptyValueIndirect(desired.VPCAccessibleServices) {
		c.Config.Logger.Infof("desired VPCAccessibleServices %s - but actually nil", dcl.SprintResource(desired.VPCAccessibleServices))
		return true
	}
	if compareServicePerimeterStatusVPCAccessibleServices(c, desired.VPCAccessibleServices, actual.VPCAccessibleServices) && !dcl.IsZeroValue(desired.VPCAccessibleServices) {
		c.Config.Logger.Infof("Diff in VPCAccessibleServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VPCAccessibleServices), dcl.SprintResource(actual.VPCAccessibleServices))
		return true
	}
	return false
}

func compareServicePerimeterStatusSlice(c *Client, desired, actual []ServicePerimeterStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServicePerimeterStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatus, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServicePerimeterStatusMap(c *Client, desired, actual map[string]ServicePerimeterStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServicePerimeterStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatus, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServicePerimeterStatusVPCAccessibleServices(c *Client, desired, actual *ServicePerimeterStatusVPCAccessibleServices) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableRestriction == nil && desired.EnableRestriction != nil && !dcl.IsEmptyValueIndirect(desired.EnableRestriction) {
		c.Config.Logger.Infof("desired EnableRestriction %s - but actually nil", dcl.SprintResource(desired.EnableRestriction))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableRestriction, actual.EnableRestriction) && !dcl.IsZeroValue(desired.EnableRestriction) {
		c.Config.Logger.Infof("Diff in EnableRestriction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableRestriction), dcl.SprintResource(actual.EnableRestriction))
		return true
	}
	if actual.AllowedServices == nil && desired.AllowedServices != nil && !dcl.IsEmptyValueIndirect(desired.AllowedServices) {
		c.Config.Logger.Infof("desired AllowedServices %s - but actually nil", dcl.SprintResource(desired.AllowedServices))
		return true
	}
	if !dcl.StringSliceEquals(desired.AllowedServices, actual.AllowedServices) && !dcl.IsZeroValue(desired.AllowedServices) {
		c.Config.Logger.Infof("Diff in AllowedServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedServices), dcl.SprintResource(actual.AllowedServices))
		return true
	}
	return false
}

func compareServicePerimeterStatusVPCAccessibleServicesSlice(c *Client, desired, actual []ServicePerimeterStatusVPCAccessibleServices) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterStatusVPCAccessibleServices, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServicePerimeterStatusVPCAccessibleServices(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatusVPCAccessibleServices, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServicePerimeterStatusVPCAccessibleServicesMap(c *Client, desired, actual map[string]ServicePerimeterStatusVPCAccessibleServices) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterStatusVPCAccessibleServices, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatusVPCAccessibleServices, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServicePerimeterStatusVPCAccessibleServices(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServicePerimeterStatusVPCAccessibleServices, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServicePerimeterSpec(c *Client, desired, actual *ServicePerimeterSpec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Resources == nil && desired.Resources != nil && !dcl.IsEmptyValueIndirect(desired.Resources) {
		c.Config.Logger.Infof("desired Resources %s - but actually nil", dcl.SprintResource(desired.Resources))
		return true
	}
	if !dcl.StringSliceEquals(desired.Resources, actual.Resources) && !dcl.IsZeroValue(desired.Resources) {
		c.Config.Logger.Infof("Diff in Resources. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Resources), dcl.SprintResource(actual.Resources))
		return true
	}
	if actual.AccessLevels == nil && desired.AccessLevels != nil && !dcl.IsEmptyValueIndirect(desired.AccessLevels) {
		c.Config.Logger.Infof("desired AccessLevels %s - but actually nil", dcl.SprintResource(desired.AccessLevels))
		return true
	}
	if !dcl.StringSliceEquals(desired.AccessLevels, actual.AccessLevels) && !dcl.IsZeroValue(desired.AccessLevels) {
		c.Config.Logger.Infof("Diff in AccessLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AccessLevels), dcl.SprintResource(actual.AccessLevels))
		return true
	}
	if actual.RestrictedServices == nil && desired.RestrictedServices != nil && !dcl.IsEmptyValueIndirect(desired.RestrictedServices) {
		c.Config.Logger.Infof("desired RestrictedServices %s - but actually nil", dcl.SprintResource(desired.RestrictedServices))
		return true
	}
	if !dcl.StringSliceEquals(desired.RestrictedServices, actual.RestrictedServices) && !dcl.IsZeroValue(desired.RestrictedServices) {
		c.Config.Logger.Infof("Diff in RestrictedServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RestrictedServices), dcl.SprintResource(actual.RestrictedServices))
		return true
	}
	if actual.VPCAccessibleServices == nil && desired.VPCAccessibleServices != nil && !dcl.IsEmptyValueIndirect(desired.VPCAccessibleServices) {
		c.Config.Logger.Infof("desired VPCAccessibleServices %s - but actually nil", dcl.SprintResource(desired.VPCAccessibleServices))
		return true
	}
	if compareServicePerimeterSpecVPCAccessibleServices(c, desired.VPCAccessibleServices, actual.VPCAccessibleServices) && !dcl.IsZeroValue(desired.VPCAccessibleServices) {
		c.Config.Logger.Infof("Diff in VPCAccessibleServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VPCAccessibleServices), dcl.SprintResource(actual.VPCAccessibleServices))
		return true
	}
	return false
}

func compareServicePerimeterSpecSlice(c *Client, desired, actual []ServicePerimeterSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterSpec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServicePerimeterSpec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpec, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServicePerimeterSpecMap(c *Client, desired, actual map[string]ServicePerimeterSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterSpec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServicePerimeterSpec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpec, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServicePerimeterSpecVPCAccessibleServices(c *Client, desired, actual *ServicePerimeterSpecVPCAccessibleServices) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableRestriction == nil && desired.EnableRestriction != nil && !dcl.IsEmptyValueIndirect(desired.EnableRestriction) {
		c.Config.Logger.Infof("desired EnableRestriction %s - but actually nil", dcl.SprintResource(desired.EnableRestriction))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableRestriction, actual.EnableRestriction) && !dcl.IsZeroValue(desired.EnableRestriction) {
		c.Config.Logger.Infof("Diff in EnableRestriction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableRestriction), dcl.SprintResource(actual.EnableRestriction))
		return true
	}
	if actual.AllowedServices == nil && desired.AllowedServices != nil && !dcl.IsEmptyValueIndirect(desired.AllowedServices) {
		c.Config.Logger.Infof("desired AllowedServices %s - but actually nil", dcl.SprintResource(desired.AllowedServices))
		return true
	}
	if !dcl.StringSliceEquals(desired.AllowedServices, actual.AllowedServices) && !dcl.IsZeroValue(desired.AllowedServices) {
		c.Config.Logger.Infof("Diff in AllowedServices. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedServices), dcl.SprintResource(actual.AllowedServices))
		return true
	}
	return false
}

func compareServicePerimeterSpecVPCAccessibleServicesSlice(c *Client, desired, actual []ServicePerimeterSpecVPCAccessibleServices) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterSpecVPCAccessibleServices, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServicePerimeterSpecVPCAccessibleServices(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpecVPCAccessibleServices, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServicePerimeterSpecVPCAccessibleServicesMap(c *Client, desired, actual map[string]ServicePerimeterSpecVPCAccessibleServices) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterSpecVPCAccessibleServices, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpecVPCAccessibleServices, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServicePerimeterSpecVPCAccessibleServices(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServicePerimeterSpecVPCAccessibleServices, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServicePerimeterPerimeterTypeEnumSlice(c *Client, desired, actual []ServicePerimeterPerimeterTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServicePerimeterPerimeterTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServicePerimeterPerimeterTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServicePerimeterPerimeterTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServicePerimeterPerimeterTypeEnum(c *Client, desired, actual *ServicePerimeterPerimeterTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ServicePerimeter) urlNormalized() *ServicePerimeter {
	normalized := deepcopy.Copy(*r).(ServicePerimeter)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Policy = dcl.SelfLinkToName(r.Policy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	return &normalized
}

func (r *ServicePerimeter) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServicePerimeter) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy)
}

func (r *ServicePerimeter) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServicePerimeter) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"policy": dcl.ValueOrEmptyString(n.Policy),
			"name":   dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("accessPolicies/{{policy}}/servicePerimeters/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ServicePerimeter resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ServicePerimeter) marshal(c *Client) ([]byte, error) {
	m, err := expandServicePerimeter(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ServicePerimeter: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalServicePerimeter decodes JSON responses into the ServicePerimeter resource schema.
func unmarshalServicePerimeter(b []byte, c *Client) (*ServicePerimeter, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapServicePerimeter(m, c)
}

func unmarshalMapServicePerimeter(m map[string]interface{}, c *Client) (*ServicePerimeter, error) {

	return flattenServicePerimeter(c, m), nil
}

// expandServicePerimeter expands ServicePerimeter into a JSON request object.
func expandServicePerimeter(c *Client, f *ServicePerimeter) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		m["title"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.PerimeterType; !dcl.IsEmptyValueIndirect(v) {
		m["perimeterType"] = v
	}
	if v, err := expandServicePerimeterStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Policy into policy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policy"] = v
	}
	if v, err := dcl.DeriveField("%s/servicePerimeters/%s", f.Name, f.Policy, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.UseExplicitDryRunSpec; !dcl.IsEmptyValueIndirect(v) {
		m["useExplicitDryRunSpec"] = v
	}
	if v, err := expandServicePerimeterSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["spec"] = v
	}

	return m, nil
}

// flattenServicePerimeter flattens ServicePerimeter from a JSON request object into the
// ServicePerimeter type.
func flattenServicePerimeter(c *Client, i interface{}) *ServicePerimeter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &ServicePerimeter{}
	r.Title = dcl.FlattenString(m["title"])
	r.Description = dcl.FlattenString(m["description"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.PerimeterType = flattenServicePerimeterPerimeterTypeEnum(m["perimeterType"])
	if _, ok := m["perimeterType"]; !ok {
		c.Config.Logger.Info("Using default value for perimeterType")
		r.PerimeterType = ServicePerimeterPerimeterTypeEnumRef("PERIMETER_TYPE_REGULAR")
	}
	r.Status = flattenServicePerimeterStatus(c, m["status"])
	r.Policy = dcl.FlattenString(m["policy"])
	r.Name = dcl.FlattenString(m["name"])
	r.UseExplicitDryRunSpec = dcl.FlattenBool(m["useExplicitDryRunSpec"])
	r.Spec = flattenServicePerimeterSpec(c, m["spec"])

	return r
}

// expandServicePerimeterStatusMap expands the contents of ServicePerimeterStatus into a JSON
// request object.
func expandServicePerimeterStatusMap(c *Client, f map[string]ServicePerimeterStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServicePerimeterStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServicePerimeterStatusSlice expands the contents of ServicePerimeterStatus into a JSON
// request object.
func expandServicePerimeterStatusSlice(c *Client, f []ServicePerimeterStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServicePerimeterStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServicePerimeterStatusMap flattens the contents of ServicePerimeterStatus from a JSON
// response object.
func flattenServicePerimeterStatusMap(c *Client, i interface{}) map[string]ServicePerimeterStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServicePerimeterStatus{}
	}

	if len(a) == 0 {
		return map[string]ServicePerimeterStatus{}
	}

	items := make(map[string]ServicePerimeterStatus)
	for k, item := range a {
		items[k] = *flattenServicePerimeterStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServicePerimeterStatusSlice flattens the contents of ServicePerimeterStatus from a JSON
// response object.
func flattenServicePerimeterStatusSlice(c *Client, i interface{}) []ServicePerimeterStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []ServicePerimeterStatus{}
	}

	if len(a) == 0 {
		return []ServicePerimeterStatus{}
	}

	items := make([]ServicePerimeterStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServicePerimeterStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandServicePerimeterStatus expands an instance of ServicePerimeterStatus into a JSON
// request object.
func expandServicePerimeterStatus(c *Client, f *ServicePerimeterStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Resources; !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v := f.AccessLevels; !dcl.IsEmptyValueIndirect(v) {
		m["accessLevels"] = v
	}
	if v := f.RestrictedServices; !dcl.IsEmptyValueIndirect(v) {
		m["restrictedServices"] = v
	}
	if v, err := expandServicePerimeterStatusVPCAccessibleServices(c, f.VPCAccessibleServices); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccessibleServices into vpcAccessibleServices: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vpcAccessibleServices"] = v
	}

	return m, nil
}

// flattenServicePerimeterStatus flattens an instance of ServicePerimeterStatus from a JSON
// response object.
func flattenServicePerimeterStatus(c *Client, i interface{}) *ServicePerimeterStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServicePerimeterStatus{}
	r.Resources = dcl.FlattenStringSlice(m["resources"])
	r.AccessLevels = dcl.FlattenStringSlice(m["accessLevels"])
	r.RestrictedServices = dcl.FlattenStringSlice(m["restrictedServices"])
	r.VPCAccessibleServices = flattenServicePerimeterStatusVPCAccessibleServices(c, m["vpcAccessibleServices"])

	return r
}

// expandServicePerimeterStatusVPCAccessibleServicesMap expands the contents of ServicePerimeterStatusVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterStatusVPCAccessibleServicesMap(c *Client, f map[string]ServicePerimeterStatusVPCAccessibleServices) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServicePerimeterStatusVPCAccessibleServices(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServicePerimeterStatusVPCAccessibleServicesSlice expands the contents of ServicePerimeterStatusVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterStatusVPCAccessibleServicesSlice(c *Client, f []ServicePerimeterStatusVPCAccessibleServices) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServicePerimeterStatusVPCAccessibleServices(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServicePerimeterStatusVPCAccessibleServicesMap flattens the contents of ServicePerimeterStatusVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterStatusVPCAccessibleServicesMap(c *Client, i interface{}) map[string]ServicePerimeterStatusVPCAccessibleServices {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServicePerimeterStatusVPCAccessibleServices{}
	}

	if len(a) == 0 {
		return map[string]ServicePerimeterStatusVPCAccessibleServices{}
	}

	items := make(map[string]ServicePerimeterStatusVPCAccessibleServices)
	for k, item := range a {
		items[k] = *flattenServicePerimeterStatusVPCAccessibleServices(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServicePerimeterStatusVPCAccessibleServicesSlice flattens the contents of ServicePerimeterStatusVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterStatusVPCAccessibleServicesSlice(c *Client, i interface{}) []ServicePerimeterStatusVPCAccessibleServices {
	a, ok := i.([]interface{})
	if !ok {
		return []ServicePerimeterStatusVPCAccessibleServices{}
	}

	if len(a) == 0 {
		return []ServicePerimeterStatusVPCAccessibleServices{}
	}

	items := make([]ServicePerimeterStatusVPCAccessibleServices, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServicePerimeterStatusVPCAccessibleServices(c, item.(map[string]interface{})))
	}

	return items
}

// expandServicePerimeterStatusVPCAccessibleServices expands an instance of ServicePerimeterStatusVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterStatusVPCAccessibleServices(c *Client, f *ServicePerimeterStatusVPCAccessibleServices) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableRestriction; !dcl.IsEmptyValueIndirect(v) {
		m["enableRestriction"] = v
	}
	if v := f.AllowedServices; !dcl.IsEmptyValueIndirect(v) {
		m["allowedServices"] = v
	}

	return m, nil
}

// flattenServicePerimeterStatusVPCAccessibleServices flattens an instance of ServicePerimeterStatusVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterStatusVPCAccessibleServices(c *Client, i interface{}) *ServicePerimeterStatusVPCAccessibleServices {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServicePerimeterStatusVPCAccessibleServices{}
	r.EnableRestriction = dcl.FlattenBool(m["enableRestriction"])
	r.AllowedServices = dcl.FlattenStringSlice(m["allowedServices"])

	return r
}

// expandServicePerimeterSpecMap expands the contents of ServicePerimeterSpec into a JSON
// request object.
func expandServicePerimeterSpecMap(c *Client, f map[string]ServicePerimeterSpec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServicePerimeterSpec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServicePerimeterSpecSlice expands the contents of ServicePerimeterSpec into a JSON
// request object.
func expandServicePerimeterSpecSlice(c *Client, f []ServicePerimeterSpec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServicePerimeterSpec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServicePerimeterSpecMap flattens the contents of ServicePerimeterSpec from a JSON
// response object.
func flattenServicePerimeterSpecMap(c *Client, i interface{}) map[string]ServicePerimeterSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServicePerimeterSpec{}
	}

	if len(a) == 0 {
		return map[string]ServicePerimeterSpec{}
	}

	items := make(map[string]ServicePerimeterSpec)
	for k, item := range a {
		items[k] = *flattenServicePerimeterSpec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServicePerimeterSpecSlice flattens the contents of ServicePerimeterSpec from a JSON
// response object.
func flattenServicePerimeterSpecSlice(c *Client, i interface{}) []ServicePerimeterSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ServicePerimeterSpec{}
	}

	if len(a) == 0 {
		return []ServicePerimeterSpec{}
	}

	items := make([]ServicePerimeterSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServicePerimeterSpec(c, item.(map[string]interface{})))
	}

	return items
}

// expandServicePerimeterSpec expands an instance of ServicePerimeterSpec into a JSON
// request object.
func expandServicePerimeterSpec(c *Client, f *ServicePerimeterSpec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Resources; !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v := f.AccessLevels; !dcl.IsEmptyValueIndirect(v) {
		m["accessLevels"] = v
	}
	if v := f.RestrictedServices; !dcl.IsEmptyValueIndirect(v) {
		m["restrictedServices"] = v
	}
	if v, err := expandServicePerimeterSpecVPCAccessibleServices(c, f.VPCAccessibleServices); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccessibleServices into vpcAccessibleServices: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vpcAccessibleServices"] = v
	}

	return m, nil
}

// flattenServicePerimeterSpec flattens an instance of ServicePerimeterSpec from a JSON
// response object.
func flattenServicePerimeterSpec(c *Client, i interface{}) *ServicePerimeterSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServicePerimeterSpec{}
	r.Resources = dcl.FlattenStringSlice(m["resources"])
	r.AccessLevels = dcl.FlattenStringSlice(m["accessLevels"])
	r.RestrictedServices = dcl.FlattenStringSlice(m["restrictedServices"])
	r.VPCAccessibleServices = flattenServicePerimeterSpecVPCAccessibleServices(c, m["vpcAccessibleServices"])

	return r
}

// expandServicePerimeterSpecVPCAccessibleServicesMap expands the contents of ServicePerimeterSpecVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterSpecVPCAccessibleServicesMap(c *Client, f map[string]ServicePerimeterSpecVPCAccessibleServices) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServicePerimeterSpecVPCAccessibleServices(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServicePerimeterSpecVPCAccessibleServicesSlice expands the contents of ServicePerimeterSpecVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterSpecVPCAccessibleServicesSlice(c *Client, f []ServicePerimeterSpecVPCAccessibleServices) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServicePerimeterSpecVPCAccessibleServices(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServicePerimeterSpecVPCAccessibleServicesMap flattens the contents of ServicePerimeterSpecVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterSpecVPCAccessibleServicesMap(c *Client, i interface{}) map[string]ServicePerimeterSpecVPCAccessibleServices {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServicePerimeterSpecVPCAccessibleServices{}
	}

	if len(a) == 0 {
		return map[string]ServicePerimeterSpecVPCAccessibleServices{}
	}

	items := make(map[string]ServicePerimeterSpecVPCAccessibleServices)
	for k, item := range a {
		items[k] = *flattenServicePerimeterSpecVPCAccessibleServices(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServicePerimeterSpecVPCAccessibleServicesSlice flattens the contents of ServicePerimeterSpecVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterSpecVPCAccessibleServicesSlice(c *Client, i interface{}) []ServicePerimeterSpecVPCAccessibleServices {
	a, ok := i.([]interface{})
	if !ok {
		return []ServicePerimeterSpecVPCAccessibleServices{}
	}

	if len(a) == 0 {
		return []ServicePerimeterSpecVPCAccessibleServices{}
	}

	items := make([]ServicePerimeterSpecVPCAccessibleServices, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServicePerimeterSpecVPCAccessibleServices(c, item.(map[string]interface{})))
	}

	return items
}

// expandServicePerimeterSpecVPCAccessibleServices expands an instance of ServicePerimeterSpecVPCAccessibleServices into a JSON
// request object.
func expandServicePerimeterSpecVPCAccessibleServices(c *Client, f *ServicePerimeterSpecVPCAccessibleServices) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableRestriction; !dcl.IsEmptyValueIndirect(v) {
		m["enableRestriction"] = v
	}
	if v := f.AllowedServices; !dcl.IsEmptyValueIndirect(v) {
		m["allowedServices"] = v
	}

	return m, nil
}

// flattenServicePerimeterSpecVPCAccessibleServices flattens an instance of ServicePerimeterSpecVPCAccessibleServices from a JSON
// response object.
func flattenServicePerimeterSpecVPCAccessibleServices(c *Client, i interface{}) *ServicePerimeterSpecVPCAccessibleServices {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServicePerimeterSpecVPCAccessibleServices{}
	r.EnableRestriction = dcl.FlattenBool(m["enableRestriction"])
	r.AllowedServices = dcl.FlattenStringSlice(m["allowedServices"])

	return r
}

// flattenServicePerimeterPerimeterTypeEnumSlice flattens the contents of ServicePerimeterPerimeterTypeEnum from a JSON
// response object.
func flattenServicePerimeterPerimeterTypeEnumSlice(c *Client, i interface{}) []ServicePerimeterPerimeterTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServicePerimeterPerimeterTypeEnum{}
	}

	if len(a) == 0 {
		return []ServicePerimeterPerimeterTypeEnum{}
	}

	items := make([]ServicePerimeterPerimeterTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServicePerimeterPerimeterTypeEnum(item.(interface{})))
	}

	return items
}

// flattenServicePerimeterPerimeterTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ServicePerimeterPerimeterTypeEnum with the same value as that string.
func flattenServicePerimeterPerimeterTypeEnum(i interface{}) *ServicePerimeterPerimeterTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ServicePerimeterPerimeterTypeEnumRef("")
	}

	return ServicePerimeterPerimeterTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ServicePerimeter) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalServicePerimeter(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Policy == nil && ncr.Policy == nil {
			c.Config.Logger.Info("Both Policy fields null - considering equal.")
		} else if nr.Policy == nil || ncr.Policy == nil {
			c.Config.Logger.Info("Only one Policy field is null - considering unequal.")
			return false
		} else if *nr.Policy != *ncr.Policy {
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
