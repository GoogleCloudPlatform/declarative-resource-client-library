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
package monitoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *UptimeCheckConfig) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"MonitoredResource", "ResourceGroup"}, r.MonitoredResource, r.ResourceGroup); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HttpCheck", "TcpCheck"}, r.HttpCheck, r.TcpCheck); err != nil {
		return err
	}
	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "timeout"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MonitoredResource) {
		if err := r.MonitoredResource.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceGroup) {
		if err := r.ResourceGroup.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpCheck) {
		if err := r.HttpCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TcpCheck) {
		if err := r.TcpCheck.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UptimeCheckConfigMonitoredResource) validate() error {
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filterLabels"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigResourceGroup) validate() error {
	return nil
}
func (r *UptimeCheckConfigHttpCheck) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AuthInfo) {
		if err := r.AuthInfo.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UptimeCheckConfigHttpCheckAuthInfo) validate() error {
	if err := dcl.Required(r, "username"); err != nil {
		return err
	}
	if err := dcl.Required(r, "password"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigTcpCheck) validate() error {
	if err := dcl.Required(r, "port"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigContentMatchers) validate() error {
	if err := dcl.Required(r, "content"); err != nil {
		return err
	}
	return nil
}

func uptimeCheckConfigGetURL(userBasePath string, r *UptimeCheckConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/v3/", userBasePath, params), nil
}

func uptimeCheckConfigListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/uptimeCheckConfigs", "https://monitoring.googleapis.com/v3/", userBasePath, params), nil

}

func uptimeCheckConfigCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/uptimeCheckConfigs", "https://monitoring.googleapis.com/v3/", userBasePath, params), nil

}

func uptimeCheckConfigDeleteURL(userBasePath string, r *UptimeCheckConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/v3/", userBasePath, params), nil
}

// uptimeCheckConfigApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type uptimeCheckConfigApiOperation interface {
	do(context.Context, *UptimeCheckConfig, *Client) error
}

// newUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest creates a request for an
// UptimeCheckConfig resource's UpdateUptimeCheckConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest(ctx context.Context, f *UptimeCheckConfig, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheck(c, f.HttpCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpCheck into httpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpCheck"] = v
	}
	if v, err := expandUptimeCheckConfigTcpCheck(c, f.TcpCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpCheck into tcpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["tcpCheck"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		req["timeout"] = v
	}
	if v, err := expandUptimeCheckConfigContentMatchersSlice(c, f.ContentMatchers); err != nil {
		return nil, fmt.Errorf("error expanding ContentMatchers into contentMatchers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["contentMatchers"] = v
	}
	if v := f.PrivateCheckers; !dcl.IsEmptyValueIndirect(v) {
		req["privateCheckers"] = v
	}
	if v := f.SelectedRegions; !dcl.IsEmptyValueIndirect(v) {
		req["selectedRegions"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", *f.Project, *f.Name)

	return req, nil
}

// marshalUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateUptimeCheckConfigUpdateUptimeCheckConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUptimeCheckConfigUpdateUptimeCheckConfigOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {
	_, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateUptimeCheckConfig")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateUptimeCheckConfigUpdateUptimeCheckConfigRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listUptimeCheckConfigRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := uptimeCheckConfigListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != UptimeCheckConfigMaxPage {
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

type listUptimeCheckConfigOperation struct {
	UptimeCheckConfigs []map[string]interface{} `json:"uptimeCheckConfigs"`
	Token              string                   `json:"nextPageToken"`
}

func (c *Client) listUptimeCheckConfig(ctx context.Context, project, pageToken string, pageSize int32) ([]*UptimeCheckConfig, string, error) {
	b, err := c.listUptimeCheckConfigRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listUptimeCheckConfigOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*UptimeCheckConfig
	for _, v := range m.UptimeCheckConfigs {
		res, err := unmarshalMapUptimeCheckConfig(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllUptimeCheckConfig(ctx context.Context, f func(*UptimeCheckConfig) bool, resources []*UptimeCheckConfig) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteUptimeCheckConfig(ctx, res)
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

type deleteUptimeCheckConfigOperation struct{}

func (op *deleteUptimeCheckConfigOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {

	_, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("UptimeCheckConfig not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUptimeCheckConfig checking for existence. error: %v", err)
		return err
	}

	u, err := uptimeCheckConfigDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete UptimeCheckConfig: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetUptimeCheckConfig(ctx, r.urlNormalized())
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
type createUptimeCheckConfigOperation struct {
	response map[string]interface{}
}

func (op *createUptimeCheckConfigOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createUptimeCheckConfigOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := uptimeCheckConfigCreateURL(c.Config.BasePath, project)

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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getUptimeCheckConfigRaw(ctx context.Context, r *UptimeCheckConfig) ([]byte, error) {
	if dcl.IsZeroValue(r.Period) {
		r.Period = dcl.String("60s")
	}

	u, err := uptimeCheckConfigGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) uptimeCheckConfigDiffsForRawDesired(ctx context.Context, rawDesired *UptimeCheckConfig, opts ...dcl.ApplyOption) (initial, desired *UptimeCheckConfig, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *UptimeCheckConfig
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*UptimeCheckConfig); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected UptimeCheckConfig, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeUptimeCheckConfigDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetUptimeCheckConfig(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a UptimeCheckConfig resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve UptimeCheckConfig resource: %v", err)
		}
		c.Config.Logger.Info("Found that UptimeCheckConfig resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for UptimeCheckConfig: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for UptimeCheckConfig: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeUptimeCheckConfigInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for UptimeCheckConfig: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for UptimeCheckConfig: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffUptimeCheckConfig(c, desired, initial, opts...)
	fmt.Printf("newDiffs: %v\n", diffs)
	return initial, desired, diffs, err
}

func canonicalizeUptimeCheckConfigInitialState(rawInitial, rawDesired *UptimeCheckConfig) (*UptimeCheckConfig, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.MonitoredResource) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.ResourceGroup) {
			rawInitial.MonitoredResource = EmptyUptimeCheckConfigMonitoredResource
		}
	}

	if dcl.IsZeroValue(rawInitial.ResourceGroup) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.MonitoredResource) {
			rawInitial.ResourceGroup = EmptyUptimeCheckConfigResourceGroup
		}
	}

	if dcl.IsZeroValue(rawInitial.HttpCheck) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.TcpCheck) {
			rawInitial.HttpCheck = EmptyUptimeCheckConfigHttpCheck
		}
	}

	if dcl.IsZeroValue(rawInitial.TcpCheck) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HttpCheck) {
			rawInitial.TcpCheck = EmptyUptimeCheckConfigTcpCheck
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial *UptimeCheckConfig, opts ...dcl.ApplyOption) (*UptimeCheckConfig, error) {

	if dcl.IsZeroValue(rawDesired.Period) {
		rawDesired.Period = dcl.String("60s")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MonitoredResource = canonicalizeUptimeCheckConfigMonitoredResource(rawDesired.MonitoredResource, nil, opts...)
		rawDesired.ResourceGroup = canonicalizeUptimeCheckConfigResourceGroup(rawDesired.ResourceGroup, nil, opts...)
		rawDesired.HttpCheck = canonicalizeUptimeCheckConfigHttpCheck(rawDesired.HttpCheck, nil, opts...)
		rawDesired.TcpCheck = canonicalizeUptimeCheckConfigTcpCheck(rawDesired.TcpCheck, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.MonitoredResource != nil || rawInitial.MonitoredResource != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.ResourceGroup) {
			rawDesired.MonitoredResource = nil
			rawInitial.MonitoredResource = nil
		}
	}

	if rawDesired.ResourceGroup != nil || rawInitial.ResourceGroup != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.MonitoredResource) {
			rawDesired.ResourceGroup = nil
			rawInitial.ResourceGroup = nil
		}
	}

	if rawDesired.HttpCheck != nil || rawInitial.HttpCheck != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.TcpCheck) {
			rawDesired.HttpCheck = nil
			rawInitial.HttpCheck = nil
		}
	}

	if rawDesired.TcpCheck != nil || rawInitial.TcpCheck != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HttpCheck) {
			rawDesired.TcpCheck = nil
			rawInitial.TcpCheck = nil
		}
	}

	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.MonitoredResource = canonicalizeUptimeCheckConfigMonitoredResource(rawDesired.MonitoredResource, rawInitial.MonitoredResource, opts...)
	rawDesired.ResourceGroup = canonicalizeUptimeCheckConfigResourceGroup(rawDesired.ResourceGroup, rawInitial.ResourceGroup, opts...)
	rawDesired.HttpCheck = canonicalizeUptimeCheckConfigHttpCheck(rawDesired.HttpCheck, rawInitial.HttpCheck, opts...)
	rawDesired.TcpCheck = canonicalizeUptimeCheckConfigTcpCheck(rawDesired.TcpCheck, rawInitial.TcpCheck, opts...)
	if dcl.StringCanonicalize(rawDesired.Period, rawInitial.Period) {
		rawDesired.Period = rawInitial.Period
	}
	if dcl.StringCanonicalize(rawDesired.Timeout, rawInitial.Timeout) {
		rawDesired.Timeout = rawInitial.Timeout
	}
	if dcl.IsZeroValue(rawDesired.ContentMatchers) {
		rawDesired.ContentMatchers = rawInitial.ContentMatchers
	}
	if dcl.IsZeroValue(rawDesired.PrivateCheckers) {
		rawDesired.PrivateCheckers = rawInitial.PrivateCheckers
	}
	if dcl.IsZeroValue(rawDesired.SelectedRegions) {
		rawDesired.SelectedRegions = rawInitial.SelectedRegions
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeUptimeCheckConfigNewState(c *Client, rawNew, rawDesired *UptimeCheckConfig) (*UptimeCheckConfig, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MonitoredResource) && dcl.IsEmptyValueIndirect(rawDesired.MonitoredResource) {
		rawNew.MonitoredResource = rawDesired.MonitoredResource
	} else {
		rawNew.MonitoredResource = canonicalizeNewUptimeCheckConfigMonitoredResource(c, rawDesired.MonitoredResource, rawNew.MonitoredResource)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceGroup) && dcl.IsEmptyValueIndirect(rawDesired.ResourceGroup) {
		rawNew.ResourceGroup = rawDesired.ResourceGroup
	} else {
		rawNew.ResourceGroup = canonicalizeNewUptimeCheckConfigResourceGroup(c, rawDesired.ResourceGroup, rawNew.ResourceGroup)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpCheck) && dcl.IsEmptyValueIndirect(rawDesired.HttpCheck) {
		rawNew.HttpCheck = rawDesired.HttpCheck
	} else {
		rawNew.HttpCheck = canonicalizeNewUptimeCheckConfigHttpCheck(c, rawDesired.HttpCheck, rawNew.HttpCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TcpCheck) && dcl.IsEmptyValueIndirect(rawDesired.TcpCheck) {
		rawNew.TcpCheck = rawDesired.TcpCheck
	} else {
		rawNew.TcpCheck = canonicalizeNewUptimeCheckConfigTcpCheck(c, rawDesired.TcpCheck, rawNew.TcpCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Period) && dcl.IsEmptyValueIndirect(rawDesired.Period) {
		rawNew.Period = rawDesired.Period
	} else {
		if dcl.StringCanonicalize(rawDesired.Period, rawNew.Period) {
			rawNew.Period = rawDesired.Period
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Timeout) && dcl.IsEmptyValueIndirect(rawDesired.Timeout) {
		rawNew.Timeout = rawDesired.Timeout
	} else {
		if dcl.StringCanonicalize(rawDesired.Timeout, rawNew.Timeout) {
			rawNew.Timeout = rawDesired.Timeout
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ContentMatchers) && dcl.IsEmptyValueIndirect(rawDesired.ContentMatchers) {
		rawNew.ContentMatchers = rawDesired.ContentMatchers
	} else {
		rawNew.ContentMatchers = canonicalizeNewUptimeCheckConfigContentMatchersSlice(c, rawDesired.ContentMatchers, rawNew.ContentMatchers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateCheckers) && dcl.IsEmptyValueIndirect(rawDesired.PrivateCheckers) {
		rawNew.PrivateCheckers = rawDesired.PrivateCheckers
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelectedRegions) && dcl.IsEmptyValueIndirect(rawDesired.SelectedRegions) {
		rawNew.SelectedRegions = rawDesired.SelectedRegions
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeUptimeCheckConfigMonitoredResource(des, initial *UptimeCheckConfigMonitoredResource, opts ...dcl.ApplyOption) *UptimeCheckConfigMonitoredResource {
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
	if dcl.IsZeroValue(des.FilterLabels) {
		des.FilterLabels = initial.FilterLabels
	}

	return des
}

func canonicalizeNewUptimeCheckConfigMonitoredResource(c *Client, des, nw *UptimeCheckConfigMonitoredResource) *UptimeCheckConfigMonitoredResource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.IsZeroValue(nw.FilterLabels) {
		nw.FilterLabels = des.FilterLabels
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigMonitoredResourceSet(c *Client, des, nw []UptimeCheckConfigMonitoredResource) []UptimeCheckConfigMonitoredResource {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigMonitoredResource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigMonitoredResourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigMonitoredResourceSlice(c *Client, des, nw []UptimeCheckConfigMonitoredResource) []UptimeCheckConfigMonitoredResource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigMonitoredResource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigMonitoredResource(c, &d, &n))
	}

	return items
}

func canonicalizeUptimeCheckConfigResourceGroup(des, initial *UptimeCheckConfigResourceGroup, opts ...dcl.ApplyOption) *UptimeCheckConfigResourceGroup {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.GroupId, initial.GroupId) || dcl.IsZeroValue(des.GroupId) {
		des.GroupId = initial.GroupId
	}
	if dcl.IsZeroValue(des.ResourceType) {
		des.ResourceType = initial.ResourceType
	}

	return des
}

func canonicalizeNewUptimeCheckConfigResourceGroup(c *Client, des, nw *UptimeCheckConfigResourceGroup) *UptimeCheckConfigResourceGroup {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.GroupId, nw.GroupId) {
		nw.GroupId = des.GroupId
	}
	if dcl.IsZeroValue(nw.ResourceType) {
		nw.ResourceType = des.ResourceType
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigResourceGroupSet(c *Client, des, nw []UptimeCheckConfigResourceGroup) []UptimeCheckConfigResourceGroup {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigResourceGroup
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigResourceGroupNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigResourceGroupSlice(c *Client, des, nw []UptimeCheckConfigResourceGroup) []UptimeCheckConfigResourceGroup {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigResourceGroup
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigResourceGroup(c, &d, &n))
	}

	return items
}

func canonicalizeUptimeCheckConfigHttpCheck(des, initial *UptimeCheckConfigHttpCheck, opts ...dcl.ApplyOption) *UptimeCheckConfigHttpCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.RequestMethod) {
		des.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}

	if dcl.IsZeroValue(des.Path) {
		des.Path = dcl.String("/")
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RequestMethod) {
		des.RequestMethod = initial.RequestMethod
	}
	if dcl.BoolCanonicalize(des.UseSsl, initial.UseSsl) || dcl.IsZeroValue(des.UseSsl) {
		des.UseSsl = initial.UseSsl
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	des.AuthInfo = canonicalizeUptimeCheckConfigHttpCheckAuthInfo(des.AuthInfo, initial.AuthInfo, opts...)
	if dcl.BoolCanonicalize(des.MaskHeaders, initial.MaskHeaders) || dcl.IsZeroValue(des.MaskHeaders) {
		des.MaskHeaders = initial.MaskHeaders
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.IsZeroValue(des.ContentType) {
		des.ContentType = initial.ContentType
	}
	if dcl.BoolCanonicalize(des.ValidateSsl, initial.ValidateSsl) || dcl.IsZeroValue(des.ValidateSsl) {
		des.ValidateSsl = initial.ValidateSsl
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		des.Body = initial.Body
	}

	return des
}

func canonicalizeNewUptimeCheckConfigHttpCheck(c *Client, des, nw *UptimeCheckConfigHttpCheck) *UptimeCheckConfigHttpCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RequestMethod) {
		nw.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}

	if dcl.IsZeroValue(nw.Path) {
		nw.Path = dcl.String("/")
	}

	if dcl.IsZeroValue(nw.RequestMethod) {
		nw.RequestMethod = des.RequestMethod
	}
	if dcl.BoolCanonicalize(des.UseSsl, nw.UseSsl) {
		nw.UseSsl = des.UseSsl
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	nw.AuthInfo = canonicalizeNewUptimeCheckConfigHttpCheckAuthInfo(c, des.AuthInfo, nw.AuthInfo)
	if dcl.BoolCanonicalize(des.MaskHeaders, nw.MaskHeaders) {
		nw.MaskHeaders = des.MaskHeaders
	}
	if dcl.IsZeroValue(nw.Headers) {
		nw.Headers = des.Headers
	}
	if dcl.IsZeroValue(nw.ContentType) {
		nw.ContentType = des.ContentType
	}
	if dcl.BoolCanonicalize(des.ValidateSsl, nw.ValidateSsl) {
		nw.ValidateSsl = des.ValidateSsl
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigHttpCheckSet(c *Client, des, nw []UptimeCheckConfigHttpCheck) []UptimeCheckConfigHttpCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigHttpCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigHttpCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigHttpCheckSlice(c *Client, des, nw []UptimeCheckConfigHttpCheck) []UptimeCheckConfigHttpCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigHttpCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigHttpCheck(c, &d, &n))
	}

	return items
}

func canonicalizeUptimeCheckConfigHttpCheckAuthInfo(des, initial *UptimeCheckConfigHttpCheckAuthInfo, opts ...dcl.ApplyOption) *UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}

	return des
}

func canonicalizeNewUptimeCheckConfigHttpCheckAuthInfo(c *Client, des, nw *UptimeCheckConfigHttpCheckAuthInfo) *UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	nw.Password = des.Password

	return nw
}

func canonicalizeNewUptimeCheckConfigHttpCheckAuthInfoSet(c *Client, des, nw []UptimeCheckConfigHttpCheckAuthInfo) []UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigHttpCheckAuthInfo
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigHttpCheckAuthInfoNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, des, nw []UptimeCheckConfigHttpCheckAuthInfo) []UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigHttpCheckAuthInfo
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigHttpCheckAuthInfo(c, &d, &n))
	}

	return items
}

func canonicalizeUptimeCheckConfigTcpCheck(des, initial *UptimeCheckConfigTcpCheck, opts ...dcl.ApplyOption) *UptimeCheckConfigTcpCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}

	return des
}

func canonicalizeNewUptimeCheckConfigTcpCheck(c *Client, des, nw *UptimeCheckConfigTcpCheck) *UptimeCheckConfigTcpCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigTcpCheckSet(c *Client, des, nw []UptimeCheckConfigTcpCheck) []UptimeCheckConfigTcpCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigTcpCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigTcpCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigTcpCheckSlice(c *Client, des, nw []UptimeCheckConfigTcpCheck) []UptimeCheckConfigTcpCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigTcpCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigTcpCheck(c, &d, &n))
	}

	return items
}

func canonicalizeUptimeCheckConfigContentMatchers(des, initial *UptimeCheckConfigContentMatchers, opts ...dcl.ApplyOption) *UptimeCheckConfigContentMatchers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.Matcher) {
		des.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.Matcher) {
		des.Matcher = initial.Matcher
	}

	return des
}

func canonicalizeNewUptimeCheckConfigContentMatchers(c *Client, des, nw *UptimeCheckConfigContentMatchers) *UptimeCheckConfigContentMatchers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Matcher) {
		nw.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.IsZeroValue(nw.Matcher) {
		nw.Matcher = des.Matcher
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigContentMatchersSet(c *Client, des, nw []UptimeCheckConfigContentMatchers) []UptimeCheckConfigContentMatchers {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigContentMatchers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareUptimeCheckConfigContentMatchersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUptimeCheckConfigContentMatchersSlice(c *Client, des, nw []UptimeCheckConfigContentMatchers) []UptimeCheckConfigContentMatchers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UptimeCheckConfigContentMatchers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUptimeCheckConfigContentMatchers(c, &d, &n))
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
func diffUptimeCheckConfig(c *Client, desired, actual *UptimeCheckConfig, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MonitoredResource, actual.MonitoredResource, dcl.Info{ObjectFunction: compareUptimeCheckConfigMonitoredResourceNewStyle, EmptyObject: EmptyUptimeCheckConfigMonitoredResource, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MonitoredResource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceGroup, actual.ResourceGroup, dcl.Info{ObjectFunction: compareUptimeCheckConfigResourceGroupNewStyle, EmptyObject: EmptyUptimeCheckConfigResourceGroup, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpCheck, actual.HttpCheck, dcl.Info{ObjectFunction: compareUptimeCheckConfigHttpCheckNewStyle, EmptyObject: EmptyUptimeCheckConfigHttpCheck, OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("HttpCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TcpCheck, actual.TcpCheck, dcl.Info{ObjectFunction: compareUptimeCheckConfigTcpCheckNewStyle, EmptyObject: EmptyUptimeCheckConfigTcpCheck, OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("TcpCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Period, actual.Period, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Period")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentMatchers, actual.ContentMatchers, dcl.Info{ObjectFunction: compareUptimeCheckConfigContentMatchersNewStyle, EmptyObject: EmptyUptimeCheckConfigContentMatchers, OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("ContentMatchers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateCheckers, actual.PrivateCheckers, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("PrivateCheckers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelectedRegions, actual.SelectedRegions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("SelectedRegions")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareUptimeCheckConfigMonitoredResourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigMonitoredResource)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigMonitoredResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigMonitoredResource or *UptimeCheckConfigMonitoredResource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigMonitoredResource)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigMonitoredResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigMonitoredResource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FilterLabels, actual.FilterLabels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUptimeCheckConfigResourceGroupNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigResourceGroup)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigResourceGroup)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigResourceGroup or *UptimeCheckConfigResourceGroup", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigResourceGroup)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigResourceGroup)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigResourceGroup", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GroupId, actual.GroupId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUptimeCheckConfigHttpCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigHttpCheck)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigHttpCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigHttpCheck or *UptimeCheckConfigHttpCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigHttpCheck)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigHttpCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigHttpCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RequestMethod, actual.RequestMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseSsl, actual.UseSsl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("UseSsl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthInfo, actual.AuthInfo, dcl.Info{ObjectFunction: compareUptimeCheckConfigHttpCheckAuthInfoNewStyle, EmptyObject: EmptyUptimeCheckConfigHttpCheckAuthInfo, OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("AuthInfo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaskHeaders, actual.MaskHeaders, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("MaskHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentType, actual.ContentType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("ContentType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValidateSsl, actual.ValidateSsl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("ValidateSsl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUptimeCheckConfigHttpCheckAuthInfoNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigHttpCheckAuthInfo)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigHttpCheckAuthInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigHttpCheckAuthInfo or *UptimeCheckConfigHttpCheckAuthInfo", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigHttpCheckAuthInfo)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigHttpCheckAuthInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigHttpCheckAuthInfo", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Password")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUptimeCheckConfigTcpCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigTcpCheck)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigTcpCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigTcpCheck or *UptimeCheckConfigTcpCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigTcpCheck)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigTcpCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigTcpCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUptimeCheckConfigContentMatchersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UptimeCheckConfigContentMatchers)
	if !ok {
		desiredNotPointer, ok := d.(UptimeCheckConfigContentMatchers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigContentMatchers or *UptimeCheckConfigContentMatchers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UptimeCheckConfigContentMatchers)
	if !ok {
		actualNotPointer, ok := a.(UptimeCheckConfigContentMatchers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UptimeCheckConfigContentMatchers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Matcher, actual.Matcher, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateUptimeCheckConfigUpdateUptimeCheckConfigOperation")}, fn.AddNest("Matcher")); len(ds) != 0 || err != nil {
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
func (r *UptimeCheckConfig) urlNormalized() *UptimeCheckConfig {
	normalized := dcl.Copy(*r).(UptimeCheckConfig)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Period = dcl.SelfLinkToName(r.Period)
	normalized.Timeout = dcl.SelfLinkToName(r.Timeout)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *UptimeCheckConfig) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UptimeCheckConfig) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *UptimeCheckConfig) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UptimeCheckConfig) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateUptimeCheckConfig" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/v3/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the UptimeCheckConfig resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *UptimeCheckConfig) marshal(c *Client) ([]byte, error) {
	m, err := expandUptimeCheckConfig(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling UptimeCheckConfig: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalUptimeCheckConfig decodes JSON responses into the UptimeCheckConfig resource schema.
func unmarshalUptimeCheckConfig(b []byte, c *Client) (*UptimeCheckConfig, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapUptimeCheckConfig(m, c)
}

func unmarshalMapUptimeCheckConfig(m map[string]interface{}, c *Client) (*UptimeCheckConfig, error) {

	return flattenUptimeCheckConfig(c, m), nil
}

// expandUptimeCheckConfig expands UptimeCheckConfig into a JSON request object.
func expandUptimeCheckConfig(c *Client, f *UptimeCheckConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandUptimeCheckConfigMonitoredResource(c, f.MonitoredResource); err != nil {
		return nil, fmt.Errorf("error expanding MonitoredResource into monitoredResource: %w", err)
	} else if v != nil {
		m["monitoredResource"] = v
	}
	if v, err := expandUptimeCheckConfigResourceGroup(c, f.ResourceGroup); err != nil {
		return nil, fmt.Errorf("error expanding ResourceGroup into resourceGroup: %w", err)
	} else if v != nil {
		m["resourceGroup"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheck(c, f.HttpCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpCheck into httpCheck: %w", err)
	} else if v != nil {
		m["httpCheck"] = v
	}
	if v, err := expandUptimeCheckConfigTcpCheck(c, f.TcpCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpCheck into tcpCheck: %w", err)
	} else if v != nil {
		m["tcpCheck"] = v
	}
	if v := f.Period; !dcl.IsEmptyValueIndirect(v) {
		m["period"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandUptimeCheckConfigContentMatchersSlice(c, f.ContentMatchers); err != nil {
		return nil, fmt.Errorf("error expanding ContentMatchers into contentMatchers: %w", err)
	} else if v != nil {
		m["contentMatchers"] = v
	}
	if v := f.PrivateCheckers; !dcl.IsEmptyValueIndirect(v) {
		m["privateCheckers"] = v
	}
	if v := f.SelectedRegions; !dcl.IsEmptyValueIndirect(v) {
		m["selectedRegions"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfig flattens UptimeCheckConfig from a JSON request object into the
// UptimeCheckConfig type.
func flattenUptimeCheckConfig(c *Client, i interface{}) *UptimeCheckConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &UptimeCheckConfig{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.MonitoredResource = flattenUptimeCheckConfigMonitoredResource(c, m["monitoredResource"])
	res.ResourceGroup = flattenUptimeCheckConfigResourceGroup(c, m["resourceGroup"])
	res.HttpCheck = flattenUptimeCheckConfigHttpCheck(c, m["httpCheck"])
	res.TcpCheck = flattenUptimeCheckConfigTcpCheck(c, m["tcpCheck"])
	res.Period = dcl.FlattenString(m["period"])
	if _, ok := m["period"]; !ok {
		c.Config.Logger.Info("Using default value for period")
		res.Period = dcl.String("60s")
	}
	res.Timeout = dcl.FlattenString(m["timeout"])
	res.ContentMatchers = flattenUptimeCheckConfigContentMatchersSlice(c, m["contentMatchers"])
	res.PrivateCheckers = dcl.FlattenStringSlice(m["privateCheckers"])
	res.SelectedRegions = dcl.FlattenStringSlice(m["selectedRegions"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandUptimeCheckConfigMonitoredResourceMap expands the contents of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResourceMap(c *Client, f map[string]UptimeCheckConfigMonitoredResource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigMonitoredResource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigMonitoredResourceSlice expands the contents of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResourceSlice(c *Client, f []UptimeCheckConfigMonitoredResource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigMonitoredResource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigMonitoredResourceMap flattens the contents of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResourceMap(c *Client, i interface{}) map[string]UptimeCheckConfigMonitoredResource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigMonitoredResource{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigMonitoredResource{}
	}

	items := make(map[string]UptimeCheckConfigMonitoredResource)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigMonitoredResource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigMonitoredResourceSlice flattens the contents of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResourceSlice(c *Client, i interface{}) []UptimeCheckConfigMonitoredResource {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigMonitoredResource{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigMonitoredResource{}
	}

	items := make([]UptimeCheckConfigMonitoredResource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigMonitoredResource(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigMonitoredResource expands an instance of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResource(c *Client, f *UptimeCheckConfigMonitoredResource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.FilterLabels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigMonitoredResource flattens an instance of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResource(c *Client, i interface{}) *UptimeCheckConfigMonitoredResource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigMonitoredResource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigMonitoredResource
	}
	r.Type = dcl.FlattenString(m["type"])
	r.FilterLabels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandUptimeCheckConfigResourceGroupMap expands the contents of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroupMap(c *Client, f map[string]UptimeCheckConfigResourceGroup) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigResourceGroup(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigResourceGroupSlice expands the contents of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroupSlice(c *Client, f []UptimeCheckConfigResourceGroup) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigResourceGroup(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigResourceGroupMap flattens the contents of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupMap(c *Client, i interface{}) map[string]UptimeCheckConfigResourceGroup {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigResourceGroup{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigResourceGroup{}
	}

	items := make(map[string]UptimeCheckConfigResourceGroup)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigResourceGroup(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigResourceGroupSlice flattens the contents of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupSlice(c *Client, i interface{}) []UptimeCheckConfigResourceGroup {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigResourceGroup{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigResourceGroup{}
	}

	items := make([]UptimeCheckConfigResourceGroup, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigResourceGroup(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigResourceGroup expands an instance of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroup(c *Client, f *UptimeCheckConfigResourceGroup) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding GroupId into groupId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["groupId"] = v
	}
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigResourceGroup flattens an instance of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroup(c *Client, i interface{}) *UptimeCheckConfigResourceGroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigResourceGroup{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigResourceGroup
	}
	r.GroupId = dcl.FlattenString(m["groupId"])
	r.ResourceType = flattenUptimeCheckConfigResourceGroupResourceTypeEnum(m["resourceType"])

	return r
}

// expandUptimeCheckConfigHttpCheckMap expands the contents of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckMap(c *Client, f map[string]UptimeCheckConfigHttpCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigHttpCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigHttpCheckSlice expands the contents of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckSlice(c *Client, f []UptimeCheckConfigHttpCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigHttpCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigHttpCheckMap flattens the contents of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckMap(c *Client, i interface{}) map[string]UptimeCheckConfigHttpCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigHttpCheck{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigHttpCheck{}
	}

	items := make(map[string]UptimeCheckConfigHttpCheck)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigHttpCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckSlice flattens the contents of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheck{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheck{}
	}

	items := make([]UptimeCheckConfigHttpCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigHttpCheck expands an instance of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheck(c *Client, f *UptimeCheckConfigHttpCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RequestMethod; !dcl.IsEmptyValueIndirect(v) {
		m["requestMethod"] = v
	}
	if v := f.UseSsl; !dcl.IsEmptyValueIndirect(v) {
		m["useSsl"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, f.AuthInfo); err != nil {
		return nil, fmt.Errorf("error expanding AuthInfo into authInfo: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authInfo"] = v
	}
	if v := f.MaskHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["maskHeaders"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.ContentType; !dcl.IsEmptyValueIndirect(v) {
		m["contentType"] = v
	}
	if v := f.ValidateSsl; !dcl.IsEmptyValueIndirect(v) {
		m["validateSsl"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigHttpCheck flattens an instance of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheck(c *Client, i interface{}) *UptimeCheckConfigHttpCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigHttpCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigHttpCheck
	}
	r.RequestMethod = flattenUptimeCheckConfigHttpCheckRequestMethodEnum(m["requestMethod"])
	if dcl.IsEmptyValueIndirect(m["requestMethod"]) {
		c.Config.Logger.Info("Using default value for requestMethod.")
		r.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}
	r.UseSsl = dcl.FlattenBool(m["useSsl"])
	r.Path = dcl.FlattenString(m["path"])
	if dcl.IsEmptyValueIndirect(m["path"]) {
		c.Config.Logger.Info("Using default value for path.")
		r.Path = dcl.String("/")
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.AuthInfo = flattenUptimeCheckConfigHttpCheckAuthInfo(c, m["authInfo"])
	r.MaskHeaders = dcl.FlattenBool(m["maskHeaders"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.ContentType = flattenUptimeCheckConfigHttpCheckContentTypeEnum(m["contentType"])
	r.ValidateSsl = dcl.FlattenBool(m["validateSsl"])
	r.Body = dcl.FlattenString(m["body"])

	return r
}

// expandUptimeCheckConfigHttpCheckAuthInfoMap expands the contents of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfoMap(c *Client, f map[string]UptimeCheckConfigHttpCheckAuthInfo) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigHttpCheckAuthInfoSlice expands the contents of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, f []UptimeCheckConfigHttpCheckAuthInfo) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigHttpCheckAuthInfoMap flattens the contents of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfoMap(c *Client, i interface{}) map[string]UptimeCheckConfigHttpCheckAuthInfo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigHttpCheckAuthInfo{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigHttpCheckAuthInfo{}
	}

	items := make(map[string]UptimeCheckConfigHttpCheckAuthInfo)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigHttpCheckAuthInfo(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckAuthInfoSlice flattens the contents of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckAuthInfo {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckAuthInfo{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckAuthInfo{}
	}

	items := make([]UptimeCheckConfigHttpCheckAuthInfo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckAuthInfo(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigHttpCheckAuthInfo expands an instance of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfo(c *Client, f *UptimeCheckConfigHttpCheckAuthInfo) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigHttpCheckAuthInfo flattens an instance of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfo(c *Client, i interface{}) *UptimeCheckConfigHttpCheckAuthInfo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigHttpCheckAuthInfo{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigHttpCheckAuthInfo
	}
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])

	return r
}

// expandUptimeCheckConfigTcpCheckMap expands the contents of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheckMap(c *Client, f map[string]UptimeCheckConfigTcpCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigTcpCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigTcpCheckSlice expands the contents of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheckSlice(c *Client, f []UptimeCheckConfigTcpCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigTcpCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigTcpCheckMap flattens the contents of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheckMap(c *Client, i interface{}) map[string]UptimeCheckConfigTcpCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigTcpCheck{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigTcpCheck{}
	}

	items := make(map[string]UptimeCheckConfigTcpCheck)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigTcpCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigTcpCheckSlice flattens the contents of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheckSlice(c *Client, i interface{}) []UptimeCheckConfigTcpCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigTcpCheck{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigTcpCheck{}
	}

	items := make([]UptimeCheckConfigTcpCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigTcpCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigTcpCheck expands an instance of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheck(c *Client, f *UptimeCheckConfigTcpCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigTcpCheck flattens an instance of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheck(c *Client, i interface{}) *UptimeCheckConfigTcpCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigTcpCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigTcpCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])

	return r
}

// expandUptimeCheckConfigContentMatchersMap expands the contents of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchersMap(c *Client, f map[string]UptimeCheckConfigContentMatchers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigContentMatchers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigContentMatchersSlice expands the contents of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchersSlice(c *Client, f []UptimeCheckConfigContentMatchers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigContentMatchers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigContentMatchersMap flattens the contents of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersMap(c *Client, i interface{}) map[string]UptimeCheckConfigContentMatchers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigContentMatchers{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigContentMatchers{}
	}

	items := make(map[string]UptimeCheckConfigContentMatchers)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigContentMatchers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigContentMatchersSlice flattens the contents of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersSlice(c *Client, i interface{}) []UptimeCheckConfigContentMatchers {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigContentMatchers{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigContentMatchers{}
	}

	items := make([]UptimeCheckConfigContentMatchers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigContentMatchers(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigContentMatchers expands an instance of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchers(c *Client, f *UptimeCheckConfigContentMatchers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.Matcher; !dcl.IsEmptyValueIndirect(v) {
		m["matcher"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigContentMatchers flattens an instance of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchers(c *Client, i interface{}) *UptimeCheckConfigContentMatchers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigContentMatchers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUptimeCheckConfigContentMatchers
	}
	r.Content = dcl.FlattenString(m["content"])
	r.Matcher = flattenUptimeCheckConfigContentMatchersMatcherEnum(m["matcher"])
	if dcl.IsEmptyValueIndirect(m["matcher"]) {
		c.Config.Logger.Info("Using default value for matcher.")
		r.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	return r
}

// flattenUptimeCheckConfigResourceGroupResourceTypeEnumSlice flattens the contents of UptimeCheckConfigResourceGroupResourceTypeEnum from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupResourceTypeEnumSlice(c *Client, i interface{}) []UptimeCheckConfigResourceGroupResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigResourceGroupResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigResourceGroupResourceTypeEnum{}
	}

	items := make([]UptimeCheckConfigResourceGroupResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigResourceGroupResourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenUptimeCheckConfigResourceGroupResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigResourceGroupResourceTypeEnum with the same value as that string.
func flattenUptimeCheckConfigResourceGroupResourceTypeEnum(i interface{}) *UptimeCheckConfigResourceGroupResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigResourceGroupResourceTypeEnumRef("")
	}

	return UptimeCheckConfigResourceGroupResourceTypeEnumRef(s)
}

// flattenUptimeCheckConfigHttpCheckRequestMethodEnumSlice flattens the contents of UptimeCheckConfigHttpCheckRequestMethodEnum from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckRequestMethodEnumSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckRequestMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckRequestMethodEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckRequestMethodEnum{}
	}

	items := make([]UptimeCheckConfigHttpCheckRequestMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckRequestMethodEnum(item.(interface{})))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckRequestMethodEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigHttpCheckRequestMethodEnum with the same value as that string.
func flattenUptimeCheckConfigHttpCheckRequestMethodEnum(i interface{}) *UptimeCheckConfigHttpCheckRequestMethodEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigHttpCheckRequestMethodEnumRef("")
	}

	return UptimeCheckConfigHttpCheckRequestMethodEnumRef(s)
}

// flattenUptimeCheckConfigHttpCheckContentTypeEnumSlice flattens the contents of UptimeCheckConfigHttpCheckContentTypeEnum from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckContentTypeEnumSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckContentTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckContentTypeEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckContentTypeEnum{}
	}

	items := make([]UptimeCheckConfigHttpCheckContentTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckContentTypeEnum(item.(interface{})))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckContentTypeEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigHttpCheckContentTypeEnum with the same value as that string.
func flattenUptimeCheckConfigHttpCheckContentTypeEnum(i interface{}) *UptimeCheckConfigHttpCheckContentTypeEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigHttpCheckContentTypeEnumRef("")
	}

	return UptimeCheckConfigHttpCheckContentTypeEnumRef(s)
}

// flattenUptimeCheckConfigContentMatchersMatcherEnumSlice flattens the contents of UptimeCheckConfigContentMatchersMatcherEnum from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersMatcherEnumSlice(c *Client, i interface{}) []UptimeCheckConfigContentMatchersMatcherEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigContentMatchersMatcherEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigContentMatchersMatcherEnum{}
	}

	items := make([]UptimeCheckConfigContentMatchersMatcherEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigContentMatchersMatcherEnum(item.(interface{})))
	}

	return items
}

// flattenUptimeCheckConfigContentMatchersMatcherEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigContentMatchersMatcherEnum with the same value as that string.
func flattenUptimeCheckConfigContentMatchersMatcherEnum(i interface{}) *UptimeCheckConfigContentMatchersMatcherEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigContentMatchersMatcherEnumRef("")
	}

	return UptimeCheckConfigContentMatchersMatcherEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *UptimeCheckConfig) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalUptimeCheckConfig(b, c)
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

type uptimeCheckConfigDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         uptimeCheckConfigApiOperation
}

func convertFieldDiffToUptimeCheckConfigOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]uptimeCheckConfigDiff, error) {
	var diffs []uptimeCheckConfigDiff
	for _, op := range ops {
		diff := uptimeCheckConfigDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTouptimeCheckConfigApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTouptimeCheckConfigApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (uptimeCheckConfigApiOperation, error) {
	switch op {

	case "updateUptimeCheckConfigUpdateUptimeCheckConfigOperation":
		return &updateUptimeCheckConfigUpdateUptimeCheckConfigOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
