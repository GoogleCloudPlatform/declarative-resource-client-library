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

func (r *UrlMap) validate() error {

	if !dcl.IsEmptyValueIndirect(r.DefaultRouteAction) {
		if err := r.DefaultRouteAction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultUrlRedirect) {
		if err := r.DefaultUrlRedirect.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapDefaultRouteAction) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UrlRewrite) {
		if err := r.UrlRewrite.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Timeout) {
		if err := r.Timeout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryPolicy) {
		if err := r.RetryPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestMirrorPolicy) {
		if err := r.RequestMirrorPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CorsPolicy) {
		if err := r.CorsPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FaultInjectionPolicy) {
		if err := r.FaultInjectionPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapDefaultRouteActionWeightedBackendService) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapHeaderAction) validate() error {
	return nil
}
func (r *UrlMapHeaderActionRequestHeadersToAdd) validate() error {
	return nil
}
func (r *UrlMapHeaderActionResponseHeadersToAdd) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionUrlRewrite) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionTimeout) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionRetryPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PerTryTimeout) {
		if err := r.PerTryTimeout.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionRequestMirrorPolicy) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionCorsPolicy) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionFaultInjectionPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Delay) {
		if err := r.Delay.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Abort) {
		if err := r.Abort.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FixedDelay) {
		if err := r.FixedDelay.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) validate() error {
	return nil
}
func (r *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) validate() error {
	return nil
}
func (r *UrlMapDefaultUrlRedirect) validate() error {
	return nil
}
func (r *UrlMapHostRule) validate() error {
	return nil
}
func (r *UrlMapPathMatcher) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DefaultRouteAction) {
		if err := r.DefaultRouteAction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultUrlRedirect) {
		if err := r.DefaultUrlRedirect.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherDefaultUrlRedirect) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.RouteAction) {
		if err := r.RouteAction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UrlRedirect) {
		if err := r.UrlRedirect.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteAction) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UrlRewrite) {
		if err := r.UrlRewrite.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Timeout) {
		if err := r.Timeout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryPolicy) {
		if err := r.RetryPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestMirrorPolicy) {
		if err := r.RequestMirrorPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CorsPolicy) {
		if err := r.CorsPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FaultInjectionPolicy) {
		if err := r.FaultInjectionPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionTimeout) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PerTryTimeout) {
		if err := r.PerTryTimeout.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Delay) {
		if err := r.Delay.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Abort) {
		if err := r.Abort.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FixedDelay) {
		if err := r.FixedDelay.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) validate() error {
	return nil
}
func (r *UrlMapPathMatcherPathRuleUrlRedirect) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.RouteAction) {
		if err := r.RouteAction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UrlRedirect) {
		if err := r.UrlRedirect.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRule) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) validate() error {
	if !dcl.IsEmptyValueIndirect(r.RangeMatch) {
		if err := r.RangeMatch.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteAction) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UrlRewrite) {
		if err := r.UrlRewrite.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Timeout) {
		if err := r.Timeout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryPolicy) {
		if err := r.RetryPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestMirrorPolicy) {
		if err := r.RequestMirrorPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CorsPolicy) {
		if err := r.CorsPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FaultInjectionPolicy) {
		if err := r.FaultInjectionPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HeaderAction) {
		if err := r.HeaderAction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionTimeout) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PerTryTimeout) {
		if err := r.PerTryTimeout.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Delay) {
		if err := r.Delay.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Abort) {
		if err := r.Abort.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FixedDelay) {
		if err := r.FixedDelay.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) validate() error {
	return nil
}
func (r *UrlMapPathMatcherRouteRuleUrlRedirect) validate() error {
	return nil
}
func (r *UrlMapTest) validate() error {
	return nil
}

func urlMapGetURL(userBasePath string, r *UrlMap) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/urlMaps/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func urlMapListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/urlMaps", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func urlMapCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/urlMaps", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func urlMapDeleteURL(userBasePath string, r *UrlMap) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/urlMaps/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// urlMapApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type urlMapApiOperation interface {
	do(context.Context, *UrlMap, *Client) error
}

// newUpdateUrlMapUpdateRequest creates a request for an
// UrlMap resource's Update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateUrlMapUpdateRequest(ctx context.Context, f *UrlMap, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandUrlMapDefaultRouteAction(c, f.DefaultRouteAction); err != nil {
		return nil, fmt.Errorf("error expanding DefaultRouteAction into defaultRouteAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["defaultRouteAction"] = v
	}
	if v := f.DefaultService; !dcl.IsEmptyValueIndirect(v) {
		req["defaultService"] = v
	}
	if v, err := expandUrlMapDefaultUrlRedirect(c, f.DefaultUrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding DefaultUrlRedirect into defaultUrlRedirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["defaultUrlRedirect"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["headerAction"] = v
	}
	if v, err := expandUrlMapHostRuleSlice(c, f.HostRule); err != nil {
		return nil, fmt.Errorf("error expanding HostRule into hostRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["hostRules"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandUrlMapPathMatcherSlice(c, f.PathMatcher); err != nil {
		return nil, fmt.Errorf("error expanding PathMatcher into pathMatchers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["pathMatchers"] = v
	}
	if v, err := expandUrlMapTestSlice(c, f.Test); err != nil {
		return nil, fmt.Errorf("error expanding Test into tests: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["tests"] = v
	}
	b, err := c.getUrlMapRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"fingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateUrlMapUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateUrlMapUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateUrlMapUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUrlMapUpdateOperation) do(ctx context.Context, r *UrlMap, c *Client) error {
	_, err := c.GetUrlMap(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Update")
	if err != nil {
		return err
	}

	req, err := newUpdateUrlMapUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateUrlMapUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
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

func (c *Client) listUrlMapRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := urlMapListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != UrlMapMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listUrlMapOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listUrlMap(ctx context.Context, project, pageToken string, pageSize int32) ([]*UrlMap, string, error) {
	b, err := c.listUrlMapRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listUrlMapOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*UrlMap
	for _, v := range m.Items {
		res := flattenUrlMap(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllUrlMap(ctx context.Context, f func(*UrlMap) bool, resources []*UrlMap) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteUrlMap(ctx, res)
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

type deleteUrlMapOperation struct{}

func (op *deleteUrlMapOperation) do(ctx context.Context, r *UrlMap, c *Client) error {

	_, err := c.GetUrlMap(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("UrlMap not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUrlMap checking for existence. error: %v", err)
		return err
	}

	u, err := urlMapDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
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
	_, err = c.GetUrlMap(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createUrlMapOperation struct{}

func (op *createUrlMapOperation) do(ctx context.Context, r *UrlMap, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := urlMapCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
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

	if _, err := c.GetUrlMap(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getUrlMapRaw(ctx context.Context, r *UrlMap) ([]byte, error) {

	u, err := urlMapGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

func (c *Client) urlMapDiffsForRawDesired(ctx context.Context, rawDesired *UrlMap, opts ...dcl.ApplyOption) (initial, desired *UrlMap, diffs []urlMapDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *UrlMap
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*UrlMap); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected UrlMap, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetUrlMap(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a UrlMap resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve UrlMap resource: %v", err)
		}
		c.Config.Logger.Info("Found that UrlMap resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeUrlMapDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for UrlMap: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for UrlMap: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeUrlMapInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for UrlMap: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeUrlMapDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for UrlMap: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffUrlMap(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeUrlMapInitialState(rawInitial, rawDesired *UrlMap) (*UrlMap, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeUrlMapDesiredState(rawDesired, rawInitial *UrlMap, opts ...dcl.ApplyOption) (*UrlMap, error) {

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*UrlMap); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected UrlMap, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(rawDesired.DefaultRouteAction, nil, opts...)
		rawDesired.DefaultUrlRedirect = canonicalizeUrlMapDefaultUrlRedirect(rawDesired.DefaultUrlRedirect, nil, opts...)
		rawDesired.HeaderAction = canonicalizeUrlMapHeaderAction(rawDesired.HeaderAction, nil, opts...)

		return rawDesired, nil
	}
	rawDesired.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(rawDesired.DefaultRouteAction, rawInitial.DefaultRouteAction, opts...)
	if dcl.IsZeroValue(rawDesired.DefaultService) {
		rawDesired.DefaultService = rawInitial.DefaultService
	}
	rawDesired.DefaultUrlRedirect = canonicalizeUrlMapDefaultUrlRedirect(rawDesired.DefaultUrlRedirect, rawInitial.DefaultUrlRedirect, opts...)
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.HeaderAction = canonicalizeUrlMapHeaderAction(rawDesired.HeaderAction, rawInitial.HeaderAction, opts...)
	if dcl.IsZeroValue(rawDesired.HostRule) {
		rawDesired.HostRule = rawInitial.HostRule
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.PathMatcher) {
		rawDesired.PathMatcher = rawInitial.PathMatcher
	}
	if dcl.IsZeroValue(rawDesired.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.IsZeroValue(rawDesired.Test) {
		rawDesired.Test = rawInitial.Test
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeUrlMapNewState(c *Client, rawNew, rawDesired *UrlMap) (*UrlMap, error) {

	if dcl.IsEmptyValueIndirect(rawNew.DefaultRouteAction) && dcl.IsEmptyValueIndirect(rawDesired.DefaultRouteAction) {
		rawNew.DefaultRouteAction = rawDesired.DefaultRouteAction
	} else {
		rawNew.DefaultRouteAction = canonicalizeNewUrlMapDefaultRouteAction(c, rawDesired.DefaultRouteAction, rawNew.DefaultRouteAction)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultService) && dcl.IsEmptyValueIndirect(rawDesired.DefaultService) {
		rawNew.DefaultService = rawDesired.DefaultService
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultUrlRedirect) && dcl.IsEmptyValueIndirect(rawDesired.DefaultUrlRedirect) {
		rawNew.DefaultUrlRedirect = rawDesired.DefaultUrlRedirect
	} else {
		rawNew.DefaultUrlRedirect = canonicalizeNewUrlMapDefaultUrlRedirect(c, rawDesired.DefaultUrlRedirect, rawNew.DefaultUrlRedirect)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HeaderAction) && dcl.IsEmptyValueIndirect(rawDesired.HeaderAction) {
		rawNew.HeaderAction = rawDesired.HeaderAction
	} else {
		rawNew.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, rawDesired.HeaderAction, rawNew.HeaderAction)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HostRule) && dcl.IsEmptyValueIndirect(rawDesired.HostRule) {
		rawNew.HostRule = rawDesired.HostRule
	} else {
		rawNew.HostRule = canonicalizeNewUrlMapHostRuleSet(c, rawDesired.HostRule, rawNew.HostRule)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PathMatcher) && dcl.IsEmptyValueIndirect(rawDesired.PathMatcher) {
		rawNew.PathMatcher = rawDesired.PathMatcher
	} else {
		rawNew.PathMatcher = canonicalizeNewUrlMapPathMatcherSet(c, rawDesired.PathMatcher, rawNew.PathMatcher)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Test) && dcl.IsEmptyValueIndirect(rawDesired.Test) {
		rawNew.Test = rawDesired.Test
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeUrlMapDefaultRouteAction(des, initial *UrlMapDefaultRouteAction, opts ...dcl.ApplyOption) *UrlMapDefaultRouteAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	}
	des.UrlRewrite = canonicalizeUrlMapDefaultRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	des.Timeout = canonicalizeUrlMapDefaultRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	des.RetryPolicy = canonicalizeUrlMapDefaultRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	des.RequestMirrorPolicy = canonicalizeUrlMapDefaultRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	des.CorsPolicy = canonicalizeUrlMapDefaultRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	des.FaultInjectionPolicy = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return des
}

func canonicalizeNewUrlMapDefaultRouteAction(c *Client, des, nw *UrlMapDefaultRouteAction) *UrlMapDefaultRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.UrlRewrite = canonicalizeNewUrlMapDefaultRouteActionUrlRewrite(c, des.UrlRewrite, nw.UrlRewrite)
	nw.Timeout = canonicalizeNewUrlMapDefaultRouteActionTimeout(c, des.Timeout, nw.Timeout)
	nw.RetryPolicy = canonicalizeNewUrlMapDefaultRouteActionRetryPolicy(c, des.RetryPolicy, nw.RetryPolicy)
	nw.RequestMirrorPolicy = canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicy(c, des.RequestMirrorPolicy, nw.RequestMirrorPolicy)
	nw.CorsPolicy = canonicalizeNewUrlMapDefaultRouteActionCorsPolicy(c, des.CorsPolicy, nw.CorsPolicy)
	nw.FaultInjectionPolicy = canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicy(c, des.FaultInjectionPolicy, nw.FaultInjectionPolicy)

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionSet(c *Client, des, nw []UrlMapDefaultRouteAction) []UrlMapDefaultRouteAction {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteAction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteAction(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionWeightedBackendService(des, initial *UrlMapDefaultRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	}
	des.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionWeightedBackendService(c *Client, des, nw *UrlMapDefaultRouteActionWeightedBackendService) *UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.BackendService, nw.BackendService) || dcl.IsZeroValue(des.BackendService) {
		nw.BackendService = des.BackendService
	}
	nw.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, des.HeaderAction, nw.HeaderAction)

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionWeightedBackendServiceSet(c *Client, des, nw []UrlMapDefaultRouteActionWeightedBackendService) []UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionWeightedBackendService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionWeightedBackendService(c, &d, &n) {
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

func canonicalizeUrlMapHeaderAction(des, initial *UrlMapHeaderAction, opts ...dcl.ApplyOption) *UrlMapHeaderAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RequestHeadersToRemove) {
		des.RequestHeadersToRemove = initial.RequestHeadersToRemove
	}
	if dcl.IsZeroValue(des.RequestHeadersToAdd) {
		des.RequestHeadersToAdd = initial.RequestHeadersToAdd
	}
	if dcl.IsZeroValue(des.ResponseHeadersToRemove) {
		des.ResponseHeadersToRemove = initial.ResponseHeadersToRemove
	}
	if dcl.IsZeroValue(des.ResponseHeadersToAdd) {
		des.ResponseHeadersToAdd = initial.ResponseHeadersToAdd
	}

	return des
}

func canonicalizeNewUrlMapHeaderAction(c *Client, des, nw *UrlMapHeaderAction) *UrlMapHeaderAction {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapHeaderActionSet(c *Client, des, nw []UrlMapHeaderAction) []UrlMapHeaderAction {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapHeaderAction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapHeaderAction(c, &d, &n) {
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

func canonicalizeUrlMapHeaderActionRequestHeadersToAdd(des, initial *UrlMapHeaderActionRequestHeadersToAdd, opts ...dcl.ApplyOption) *UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HeaderName) {
		des.HeaderName = initial.HeaderName
	}
	if dcl.IsZeroValue(des.HeaderValue) {
		des.HeaderValue = initial.HeaderValue
	}
	if dcl.IsZeroValue(des.Replace) {
		des.Replace = initial.Replace
	}

	return des
}

func canonicalizeNewUrlMapHeaderActionRequestHeadersToAdd(c *Client, des, nw *UrlMapHeaderActionRequestHeadersToAdd) *UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapHeaderActionRequestHeadersToAddSet(c *Client, des, nw []UrlMapHeaderActionRequestHeadersToAdd) []UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapHeaderActionRequestHeadersToAdd
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapHeaderActionRequestHeadersToAdd(c, &d, &n) {
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

func canonicalizeUrlMapHeaderActionResponseHeadersToAdd(des, initial *UrlMapHeaderActionResponseHeadersToAdd, opts ...dcl.ApplyOption) *UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HeaderName) {
		des.HeaderName = initial.HeaderName
	}
	if dcl.IsZeroValue(des.HeaderValue) {
		des.HeaderValue = initial.HeaderValue
	}
	if dcl.IsZeroValue(des.Replace) {
		des.Replace = initial.Replace
	}

	return des
}

func canonicalizeNewUrlMapHeaderActionResponseHeadersToAdd(c *Client, des, nw *UrlMapHeaderActionResponseHeadersToAdd) *UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapHeaderActionResponseHeadersToAddSet(c *Client, des, nw []UrlMapHeaderActionResponseHeadersToAdd) []UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapHeaderActionResponseHeadersToAdd
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapHeaderActionResponseHeadersToAdd(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionUrlRewrite(des, initial *UrlMapDefaultRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PathPrefixRewrite) {
		des.PathPrefixRewrite = initial.PathPrefixRewrite
	}
	if dcl.IsZeroValue(des.HostRewrite) {
		des.HostRewrite = initial.HostRewrite
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionUrlRewrite(c *Client, des, nw *UrlMapDefaultRouteActionUrlRewrite) *UrlMapDefaultRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionUrlRewriteSet(c *Client, des, nw []UrlMapDefaultRouteActionUrlRewrite) []UrlMapDefaultRouteActionUrlRewrite {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionUrlRewrite
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionUrlRewrite(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionTimeout(des, initial *UrlMapDefaultRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionTimeout(c *Client, des, nw *UrlMapDefaultRouteActionTimeout) *UrlMapDefaultRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionTimeoutSet(c *Client, des, nw []UrlMapDefaultRouteActionTimeout) []UrlMapDefaultRouteActionTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionTimeout(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionRetryPolicy(des, initial *UrlMapDefaultRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	}
	des.PerTryTimeout = canonicalizeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicy(c *Client, des, nw *UrlMapDefaultRouteActionRetryPolicy) *UrlMapDefaultRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.PerTryTimeout = canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, des.PerTryTimeout, nw.PerTryTimeout)

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicySet(c *Client, des, nw []UrlMapDefaultRouteActionRetryPolicy) []UrlMapDefaultRouteActionRetryPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionRetryPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionRetryPolicy(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSet(c *Client, des, nw []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionRequestMirrorPolicy(des, initial *UrlMapDefaultRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapDefaultRouteActionRequestMirrorPolicy) *UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicySet(c *Client, des, nw []UrlMapDefaultRouteActionRequestMirrorPolicy) []UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionRequestMirrorPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionRequestMirrorPolicy(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionCorsPolicy(des, initial *UrlMapDefaultRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	}
	if dcl.IsZeroValue(des.AllowCredentials) {
		des.AllowCredentials = initial.AllowCredentials
	}
	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionCorsPolicy(c *Client, des, nw *UrlMapDefaultRouteActionCorsPolicy) *UrlMapDefaultRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionCorsPolicySet(c *Client, des, nw []UrlMapDefaultRouteActionCorsPolicy) []UrlMapDefaultRouteActionCorsPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionCorsPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionCorsPolicy(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicy(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Delay = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	des.Abort = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicy(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicy) *UrlMapDefaultRouteActionFaultInjectionPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.Delay = canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, des.Delay, nw.Delay)
	nw.Abort = canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, des.Abort, nw.Abort)

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicySet(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicy) []UrlMapDefaultRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionFaultInjectionPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionFaultInjectionPolicy(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.FixedDelay = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) *UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelaySet(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyDelay) []UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionFaultInjectionPolicyDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySet(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n) {
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

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) *UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbortSet(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyAbort) []UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultRouteActionFaultInjectionPolicyAbort
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, &d, &n) {
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

func canonicalizeUrlMapDefaultUrlRedirect(des, initial *UrlMapDefaultUrlRedirect, opts ...dcl.ApplyOption) *UrlMapDefaultUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HostRedirect) {
		des.HostRedirect = initial.HostRedirect
	}
	if dcl.IsZeroValue(des.PathRedirect) {
		des.PathRedirect = initial.PathRedirect
	}
	if dcl.IsZeroValue(des.PrefixRedirect) {
		des.PrefixRedirect = initial.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	}
	if dcl.IsZeroValue(des.HttpsRedirect) {
		des.HttpsRedirect = initial.HttpsRedirect
	}
	if dcl.IsZeroValue(des.StripQuery) {
		des.StripQuery = initial.StripQuery
	}

	return des
}

func canonicalizeNewUrlMapDefaultUrlRedirect(c *Client, des, nw *UrlMapDefaultUrlRedirect) *UrlMapDefaultUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapDefaultUrlRedirectSet(c *Client, des, nw []UrlMapDefaultUrlRedirect) []UrlMapDefaultUrlRedirect {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapDefaultUrlRedirect
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapDefaultUrlRedirect(c, &d, &n) {
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

func canonicalizeUrlMapHostRule(des, initial *UrlMapHostRule, opts ...dcl.ApplyOption) *UrlMapHostRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.IsZeroValue(des.PathMatcher) {
		des.PathMatcher = initial.PathMatcher
	}

	return des
}

func canonicalizeNewUrlMapHostRule(c *Client, des, nw *UrlMapHostRule) *UrlMapHostRule {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapHostRuleSet(c *Client, des, nw []UrlMapHostRule) []UrlMapHostRule {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapHostRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapHostRule(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcher(des, initial *UrlMapPathMatcher, opts ...dcl.ApplyOption) *UrlMapPathMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.DefaultService) {
		des.DefaultService = initial.DefaultService
	}
	des.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(des.DefaultRouteAction, initial.DefaultRouteAction, opts...)
	des.DefaultUrlRedirect = canonicalizeUrlMapPathMatcherDefaultUrlRedirect(des.DefaultUrlRedirect, initial.DefaultUrlRedirect, opts...)
	if dcl.IsZeroValue(des.PathRule) {
		des.PathRule = initial.PathRule
	}
	if dcl.IsZeroValue(des.RouteRule) {
		des.RouteRule = initial.RouteRule
	}
	des.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcher(c *Client, des, nw *UrlMapPathMatcher) *UrlMapPathMatcher {
	if des == nil || nw == nil {
		return nw
	}

	nw.DefaultRouteAction = canonicalizeNewUrlMapDefaultRouteAction(c, des.DefaultRouteAction, nw.DefaultRouteAction)
	nw.DefaultUrlRedirect = canonicalizeNewUrlMapPathMatcherDefaultUrlRedirect(c, des.DefaultUrlRedirect, nw.DefaultUrlRedirect)
	nw.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, des.HeaderAction, nw.HeaderAction)

	return nw
}

func canonicalizeNewUrlMapPathMatcherSet(c *Client, des, nw []UrlMapPathMatcher) []UrlMapPathMatcher {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcher
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcher(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherDefaultUrlRedirect(des, initial *UrlMapPathMatcherDefaultUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HostRedirect) {
		des.HostRedirect = initial.HostRedirect
	}
	if dcl.IsZeroValue(des.PathRedirect) {
		des.PathRedirect = initial.PathRedirect
	}
	if dcl.IsZeroValue(des.PrefixRedirect) {
		des.PrefixRedirect = initial.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	}
	if dcl.IsZeroValue(des.HttpsRedirect) {
		des.HttpsRedirect = initial.HttpsRedirect
	}
	if dcl.IsZeroValue(des.StripQuery) {
		des.StripQuery = initial.StripQuery
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherDefaultUrlRedirect(c *Client, des, nw *UrlMapPathMatcherDefaultUrlRedirect) *UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherDefaultUrlRedirectSet(c *Client, des, nw []UrlMapPathMatcherDefaultUrlRedirect) []UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherDefaultUrlRedirect
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherDefaultUrlRedirect(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRule(des, initial *UrlMapPathMatcherPathRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}
	des.RouteAction = canonicalizeUrlMapPathMatcherPathRuleRouteAction(des.RouteAction, initial.RouteAction, opts...)
	des.UrlRedirect = canonicalizeUrlMapPathMatcherPathRuleUrlRedirect(des.UrlRedirect, initial.UrlRedirect, opts...)
	if dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRule(c *Client, des, nw *UrlMapPathMatcherPathRule) *UrlMapPathMatcherPathRule {
	if des == nil || nw == nil {
		return nw
	}

	nw.RouteAction = canonicalizeNewUrlMapPathMatcherPathRuleRouteAction(c, des.RouteAction, nw.RouteAction)
	nw.UrlRedirect = canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirect(c, des.UrlRedirect, nw.UrlRedirect)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleSet(c *Client, des, nw []UrlMapPathMatcherPathRule) []UrlMapPathMatcherPathRule {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRule(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteAction(des, initial *UrlMapPathMatcherPathRuleRouteAction, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	}
	des.UrlRewrite = canonicalizeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	des.Timeout = canonicalizeUrlMapPathMatcherPathRuleRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	des.RetryPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	des.RequestMirrorPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	des.CorsPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	des.FaultInjectionPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteAction(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteAction) *UrlMapPathMatcherPathRuleRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.UrlRewrite = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, des.UrlRewrite, nw.UrlRewrite)
	nw.Timeout = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeout(c, des.Timeout, nw.Timeout)
	nw.RetryPolicy = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, des.RetryPolicy, nw.RetryPolicy)
	nw.RequestMirrorPolicy = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, des.RequestMirrorPolicy, nw.RequestMirrorPolicy)
	nw.CorsPolicy = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, des.CorsPolicy, nw.CorsPolicy)
	nw.FaultInjectionPolicy = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, des.FaultInjectionPolicy, nw.FaultInjectionPolicy)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteAction) []UrlMapPathMatcherPathRuleRouteAction {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteAction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteAction(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(des, initial *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	}
	des.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	nw.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, des.HeaderAction, nw.HeaderAction)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(des, initial *UrlMapPathMatcherPathRuleRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PathPrefixRewrite) {
		des.PathPrefixRewrite = initial.PathPrefixRewrite
	}
	if dcl.IsZeroValue(des.HostRewrite) {
		des.HostRewrite = initial.HostRewrite
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) *UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewriteSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionUrlRewrite) []UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionUrlRewrite
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionTimeout(des, initial *UrlMapPathMatcherPathRuleRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeout(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionTimeout) *UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeoutSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionTimeout) []UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionTimeout(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	}
	des.PerTryTimeout = canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) *UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.PerTryTimeout = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, des.PerTryTimeout, nw.PerTryTimeout)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRetryPolicy) []UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionRetryPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	}
	if dcl.IsZeroValue(des.AllowCredentials) {
		des.AllowCredentials = initial.AllowCredentials
	}
	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) *UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionCorsPolicy) []UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionCorsPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Delay = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	des.Abort = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.Delay = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, des.Delay, nw.Delay)
	nw.Abort = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, des.Abort, nw.Abort)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.FixedDelay = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSet(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherPathRuleUrlRedirect(des, initial *UrlMapPathMatcherPathRuleUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HostRedirect) {
		des.HostRedirect = initial.HostRedirect
	}
	if dcl.IsZeroValue(des.PathRedirect) {
		des.PathRedirect = initial.PathRedirect
	}
	if dcl.IsZeroValue(des.PrefixRedirect) {
		des.PrefixRedirect = initial.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	}
	if dcl.IsZeroValue(des.HttpsRedirect) {
		des.HttpsRedirect = initial.HttpsRedirect
	}
	if dcl.IsZeroValue(des.StripQuery) {
		des.StripQuery = initial.StripQuery
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirect(c *Client, des, nw *UrlMapPathMatcherPathRuleUrlRedirect) *UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirectSet(c *Client, des, nw []UrlMapPathMatcherPathRuleUrlRedirect) []UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherPathRuleUrlRedirect
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherPathRuleUrlRedirect(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRule(des, initial *UrlMapPathMatcherRouteRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Priority) {
		des.Priority = initial.Priority
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.MatchRule) {
		des.MatchRule = initial.MatchRule
	}
	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}
	des.RouteAction = canonicalizeUrlMapPathMatcherRouteRuleRouteAction(des.RouteAction, initial.RouteAction, opts...)
	des.UrlRedirect = canonicalizeUrlMapPathMatcherRouteRuleUrlRedirect(des.UrlRedirect, initial.UrlRedirect, opts...)
	des.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRule(c *Client, des, nw *UrlMapPathMatcherRouteRule) *UrlMapPathMatcherRouteRule {
	if des == nil || nw == nil {
		return nw
	}

	nw.RouteAction = canonicalizeNewUrlMapPathMatcherRouteRuleRouteAction(c, des.RouteAction, nw.RouteAction)
	nw.UrlRedirect = canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirect(c, des.UrlRedirect, nw.UrlRedirect)
	nw.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, des.HeaderAction, nw.HeaderAction)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleSet(c *Client, des, nw []UrlMapPathMatcherRouteRule) []UrlMapPathMatcherRouteRule {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRule(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRule(des, initial *UrlMapPathMatcherRouteRuleMatchRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PrefixMatch) {
		des.PrefixMatch = initial.PrefixMatch
	}
	if dcl.IsZeroValue(des.FullPathMatch) {
		des.FullPathMatch = initial.FullPathMatch
	}
	if dcl.IsZeroValue(des.RegexMatch) {
		des.RegexMatch = initial.RegexMatch
	}
	if dcl.IsZeroValue(des.IgnoreCase) {
		des.IgnoreCase = initial.IgnoreCase
	}
	if dcl.IsZeroValue(des.HeaderMatch) {
		des.HeaderMatch = initial.HeaderMatch
	}
	if dcl.IsZeroValue(des.QueryParameterMatch) {
		des.QueryParameterMatch = initial.QueryParameterMatch
	}
	if dcl.IsZeroValue(des.MetadataFilter) {
		des.MetadataFilter = initial.MetadataFilter
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRule(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRule) *UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRule) []UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRule(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HeaderName) {
		des.HeaderName = initial.HeaderName
	}
	if dcl.IsZeroValue(des.ExactMatch) {
		des.ExactMatch = initial.ExactMatch
	}
	if dcl.IsZeroValue(des.RegexMatch) {
		des.RegexMatch = initial.RegexMatch
	}
	des.RangeMatch = canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(des.RangeMatch, initial.RangeMatch, opts...)
	if dcl.IsZeroValue(des.PresentMatch) {
		des.PresentMatch = initial.PresentMatch
	}
	if dcl.IsZeroValue(des.PrefixMatch) {
		des.PrefixMatch = initial.PrefixMatch
	}
	if dcl.IsZeroValue(des.SuffixMatch) {
		des.SuffixMatch = initial.SuffixMatch
	}
	if dcl.IsZeroValue(des.InvertMatch) {
		des.InvertMatch = initial.InvertMatch
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil || nw == nil {
		return nw
	}

	nw.RangeMatch = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, des.RangeMatch, nw.RangeMatch)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RangeStart) {
		des.RangeStart = initial.RangeStart
	}
	if dcl.IsZeroValue(des.RangeEnd) {
		des.RangeEnd = initial.RangeEnd
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.PresentMatch) {
		des.PresentMatch = initial.PresentMatch
	}
	if dcl.IsZeroValue(des.ExactMatch) {
		des.ExactMatch = initial.ExactMatch
	}
	if dcl.IsZeroValue(des.RegexMatch) {
		des.RegexMatch = initial.RegexMatch
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(des, initial *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.FilterMatchCriteria) {
		des.FilterMatchCriteria = initial.FilterMatchCriteria
	}
	if dcl.IsZeroValue(des.FilterLabel) {
		des.FilterLabel = initial.FilterLabel
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(des, initial *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteAction(des, initial *UrlMapPathMatcherRouteRuleRouteAction, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	}
	des.UrlRewrite = canonicalizeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	des.Timeout = canonicalizeUrlMapPathMatcherRouteRuleRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	des.RetryPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	des.RequestMirrorPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	des.CorsPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	des.FaultInjectionPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteAction(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteAction) *UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.UrlRewrite = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, des.UrlRewrite, nw.UrlRewrite)
	nw.Timeout = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeout(c, des.Timeout, nw.Timeout)
	nw.RetryPolicy = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, des.RetryPolicy, nw.RetryPolicy)
	nw.RequestMirrorPolicy = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, des.RequestMirrorPolicy, nw.RequestMirrorPolicy)
	nw.CorsPolicy = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, des.CorsPolicy, nw.CorsPolicy)
	nw.FaultInjectionPolicy = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, des.FaultInjectionPolicy, nw.FaultInjectionPolicy)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteAction) []UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteAction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteAction(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(des, initial *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	}
	des.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	nw.HeaderAction = canonicalizeNewUrlMapHeaderAction(c, des.HeaderAction, nw.HeaderAction)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(des, initial *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PathPrefixRewrite) {
		des.PathPrefixRewrite = initial.PathPrefixRewrite
	}
	if dcl.IsZeroValue(des.HostRewrite) {
		des.HostRewrite = initial.HostRewrite
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionTimeout(des, initial *UrlMapPathMatcherRouteRuleRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeout(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionTimeout) *UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeoutSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionTimeout) []UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionTimeout(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	}
	des.PerTryTimeout = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.PerTryTimeout = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, des.PerTryTimeout, nw.PerTryTimeout)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BackendService) {
		des.BackendService = initial.BackendService
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	}
	if dcl.IsZeroValue(des.AllowCredentials) {
		des.AllowCredentials = initial.AllowCredentials
	}
	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Delay = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	des.Abort = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.Delay = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, des.Delay, nw.Delay)
	nw.Abort = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, des.Abort, nw.Abort)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.FixedDelay = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, &d, &n) {
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

func canonicalizeUrlMapPathMatcherRouteRuleUrlRedirect(des, initial *UrlMapPathMatcherRouteRuleUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HostRedirect) {
		des.HostRedirect = initial.HostRedirect
	}
	if dcl.IsZeroValue(des.PathRedirect) {
		des.PathRedirect = initial.PathRedirect
	}
	if dcl.IsZeroValue(des.PrefixRedirect) {
		des.PrefixRedirect = initial.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	}
	if dcl.IsZeroValue(des.HttpsRedirect) {
		des.HttpsRedirect = initial.HttpsRedirect
	}
	if dcl.IsZeroValue(des.StripQuery) {
		des.StripQuery = initial.StripQuery
	}

	return des
}

func canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirect(c *Client, des, nw *UrlMapPathMatcherRouteRuleUrlRedirect) *UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirectSet(c *Client, des, nw []UrlMapPathMatcherRouteRuleUrlRedirect) []UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapPathMatcherRouteRuleUrlRedirect
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapPathMatcherRouteRuleUrlRedirect(c, &d, &n) {
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

func canonicalizeUrlMapTest(des, initial *UrlMapTest, opts ...dcl.ApplyOption) *UrlMapTest {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UrlMap)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.ExpectedBackendService) {
		des.ExpectedBackendService = initial.ExpectedBackendService
	}

	return des
}

func canonicalizeNewUrlMapTest(c *Client, des, nw *UrlMapTest) *UrlMapTest {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUrlMapTestSet(c *Client, des, nw []UrlMapTest) []UrlMapTest {
	if des == nil {
		return nw
	}
	var reorderedNew []UrlMapTest
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUrlMapTest(c, &d, &n) {
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

type urlMapDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         urlMapApiOperation
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
func diffUrlMap(c *Client, desired, actual *UrlMap, opts ...dcl.ApplyOption) ([]urlMapDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []urlMapDiff
	if compareUrlMapDefaultRouteAction(c, desired.DefaultRouteAction, actual.DefaultRouteAction) {
		c.Config.Logger.Infof("Detected diff in DefaultRouteAction.\nDESIRED: %v\nACTUAL: %v", desired.DefaultRouteAction, actual.DefaultRouteAction)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "DefaultRouteAction",
		})

	}
	if !dcl.IsZeroValue(desired.DefaultService) && (dcl.IsZeroValue(actual.DefaultService) || !reflect.DeepEqual(*desired.DefaultService, *actual.DefaultService)) {
		c.Config.Logger.Infof("Detected diff in DefaultService.\nDESIRED: %v\nACTUAL: %v", desired.DefaultService, actual.DefaultService)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "DefaultService",
		})

	}
	if compareUrlMapDefaultUrlRedirect(c, desired.DefaultUrlRedirect, actual.DefaultUrlRedirect) {
		c.Config.Logger.Infof("Detected diff in DefaultUrlRedirect.\nDESIRED: %v\nACTUAL: %v", desired.DefaultUrlRedirect, actual.DefaultUrlRedirect)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "DefaultUrlRedirect",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "Description",
		})

	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) {
		c.Config.Logger.Infof("Detected diff in HeaderAction.\nDESIRED: %v\nACTUAL: %v", desired.HeaderAction, actual.HeaderAction)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "HeaderAction",
		})

	}
	if compareUrlMapHostRuleSlice(c, desired.HostRule, actual.HostRule) {
		c.Config.Logger.Infof("Detected diff in HostRule.\nDESIRED: %v\nACTUAL: %v", desired.HostRule, actual.HostRule)

		toAdd, toRemove := compareUrlMapHostRuleSets(c, desired.HostRule, actual.HostRule)
		c.Config.Logger.Infof("diff in HostRule is a set field - recomparing with set logic. \nto add: %#v\nto remove: %#v", toAdd, toRemove)
		if len(toAdd) != 0 || len(toRemove) != 0 {
			c.Config.Logger.Info("diff in HostRule persists after set logic analysis.")
			diffs = append(diffs, urlMapDiff{
				UpdateOp:  &updateUrlMapUpdateOperation{},
				FieldName: "HostRule",
			})
		}

	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "Name",
		})

	}
	if compareUrlMapPathMatcherSlice(c, desired.PathMatcher, actual.PathMatcher) {
		c.Config.Logger.Infof("Detected diff in PathMatcher.\nDESIRED: %v\nACTUAL: %v", desired.PathMatcher, actual.PathMatcher)

		toAdd, toRemove := compareUrlMapPathMatcherSets(c, desired.PathMatcher, actual.PathMatcher)
		c.Config.Logger.Infof("diff in PathMatcher is a set field - recomparing with set logic. \nto add: %#v\nto remove: %#v", toAdd, toRemove)
		if len(toAdd) != 0 || len(toRemove) != 0 {
			c.Config.Logger.Info("diff in PathMatcher persists after set logic analysis.")
			diffs = append(diffs, urlMapDiff{
				UpdateOp:  &updateUrlMapUpdateOperation{},
				FieldName: "PathMatcher",
			})
		}

	}
	if !dcl.IsZeroValue(desired.Region) && (dcl.IsZeroValue(actual.Region) || !reflect.DeepEqual(*desired.Region, *actual.Region)) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, urlMapDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if compareUrlMapTestSlice(c, desired.Test, actual.Test) {
		c.Config.Logger.Infof("Detected diff in Test.\nDESIRED: %v\nACTUAL: %v", desired.Test, actual.Test)

		diffs = append(diffs, urlMapDiff{
			UpdateOp:  &updateUrlMapUpdateOperation{},
			FieldName: "Test",
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
	var deduped []urlMapDiff
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
func compareUrlMapDefaultRouteActionSlice(c *Client, desired, actual []UrlMapDefaultRouteAction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteAction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteAction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteAction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteAction(c *Client, desired, actual *UrlMapDefaultRouteAction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.WeightedBackendService == nil && desired.WeightedBackendService != nil && !dcl.IsEmptyValueIndirect(desired.WeightedBackendService) {
		c.Config.Logger.Infof("desired WeightedBackendService %s - but actually nil", dcl.SprintResource(desired.WeightedBackendService))
		return true
	}
	if compareUrlMapDefaultRouteActionWeightedBackendServiceSlice(c, desired.WeightedBackendService, actual.WeightedBackendService) && !dcl.IsZeroValue(desired.WeightedBackendService) {
		c.Config.Logger.Infof("Diff in WeightedBackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WeightedBackendService), dcl.SprintResource(actual.WeightedBackendService))
		return true
	}
	if actual.UrlRewrite == nil && desired.UrlRewrite != nil && !dcl.IsEmptyValueIndirect(desired.UrlRewrite) {
		c.Config.Logger.Infof("desired UrlRewrite %s - but actually nil", dcl.SprintResource(desired.UrlRewrite))
		return true
	}
	if compareUrlMapDefaultRouteActionUrlRewrite(c, desired.UrlRewrite, actual.UrlRewrite) && !dcl.IsZeroValue(desired.UrlRewrite) {
		c.Config.Logger.Infof("Diff in UrlRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlRewrite), dcl.SprintResource(actual.UrlRewrite))
		return true
	}
	if actual.Timeout == nil && desired.Timeout != nil && !dcl.IsEmptyValueIndirect(desired.Timeout) {
		c.Config.Logger.Infof("desired Timeout %s - but actually nil", dcl.SprintResource(desired.Timeout))
		return true
	}
	if compareUrlMapDefaultRouteActionTimeout(c, desired.Timeout, actual.Timeout) && !dcl.IsZeroValue(desired.Timeout) {
		c.Config.Logger.Infof("Diff in Timeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Timeout), dcl.SprintResource(actual.Timeout))
		return true
	}
	if actual.RetryPolicy == nil && desired.RetryPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RetryPolicy) {
		c.Config.Logger.Infof("desired RetryPolicy %s - but actually nil", dcl.SprintResource(desired.RetryPolicy))
		return true
	}
	if compareUrlMapDefaultRouteActionRetryPolicy(c, desired.RetryPolicy, actual.RetryPolicy) && !dcl.IsZeroValue(desired.RetryPolicy) {
		c.Config.Logger.Infof("Diff in RetryPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryPolicy), dcl.SprintResource(actual.RetryPolicy))
		return true
	}
	if actual.RequestMirrorPolicy == nil && desired.RequestMirrorPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("desired RequestMirrorPolicy %s - but actually nil", dcl.SprintResource(desired.RequestMirrorPolicy))
		return true
	}
	if compareUrlMapDefaultRouteActionRequestMirrorPolicy(c, desired.RequestMirrorPolicy, actual.RequestMirrorPolicy) && !dcl.IsZeroValue(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("Diff in RequestMirrorPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestMirrorPolicy), dcl.SprintResource(actual.RequestMirrorPolicy))
		return true
	}
	if actual.CorsPolicy == nil && desired.CorsPolicy != nil && !dcl.IsEmptyValueIndirect(desired.CorsPolicy) {
		c.Config.Logger.Infof("desired CorsPolicy %s - but actually nil", dcl.SprintResource(desired.CorsPolicy))
		return true
	}
	if compareUrlMapDefaultRouteActionCorsPolicy(c, desired.CorsPolicy, actual.CorsPolicy) && !dcl.IsZeroValue(desired.CorsPolicy) {
		c.Config.Logger.Infof("Diff in CorsPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CorsPolicy), dcl.SprintResource(actual.CorsPolicy))
		return true
	}
	if actual.FaultInjectionPolicy == nil && desired.FaultInjectionPolicy != nil && !dcl.IsEmptyValueIndirect(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("desired FaultInjectionPolicy %s - but actually nil", dcl.SprintResource(desired.FaultInjectionPolicy))
		return true
	}
	if compareUrlMapDefaultRouteActionFaultInjectionPolicy(c, desired.FaultInjectionPolicy, actual.FaultInjectionPolicy) && !dcl.IsZeroValue(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("Diff in FaultInjectionPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FaultInjectionPolicy), dcl.SprintResource(actual.FaultInjectionPolicy))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionWeightedBackendServiceSlice(c *Client, desired, actual []UrlMapDefaultRouteActionWeightedBackendService) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionWeightedBackendService, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionWeightedBackendService(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionWeightedBackendService, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionWeightedBackendService(c *Client, desired, actual *UrlMapDefaultRouteActionWeightedBackendService) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !dcl.NameToSelfLink(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	if actual.Weight == nil && desired.Weight != nil && !dcl.IsEmptyValueIndirect(desired.Weight) {
		c.Config.Logger.Infof("desired Weight %s - but actually nil", dcl.SprintResource(desired.Weight))
		return true
	}
	if !reflect.DeepEqual(desired.Weight, actual.Weight) && !dcl.IsZeroValue(desired.Weight) && !(dcl.IsEmptyValueIndirect(desired.Weight) && dcl.IsZeroValue(actual.Weight)) {
		c.Config.Logger.Infof("Diff in Weight. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Weight), dcl.SprintResource(actual.Weight))
		return true
	}
	if actual.HeaderAction == nil && desired.HeaderAction != nil && !dcl.IsEmptyValueIndirect(desired.HeaderAction) {
		c.Config.Logger.Infof("desired HeaderAction %s - but actually nil", dcl.SprintResource(desired.HeaderAction))
		return true
	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) && !dcl.IsZeroValue(desired.HeaderAction) {
		c.Config.Logger.Infof("Diff in HeaderAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderAction), dcl.SprintResource(actual.HeaderAction))
		return true
	}
	return false
}
func compareUrlMapHeaderActionSlice(c *Client, desired, actual []UrlMapHeaderAction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapHeaderAction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapHeaderAction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapHeaderAction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapHeaderAction(c *Client, desired, actual *UrlMapHeaderAction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RequestHeadersToRemove == nil && desired.RequestHeadersToRemove != nil && !dcl.IsEmptyValueIndirect(desired.RequestHeadersToRemove) {
		c.Config.Logger.Infof("desired RequestHeadersToRemove %s - but actually nil", dcl.SprintResource(desired.RequestHeadersToRemove))
		return true
	}
	if !dcl.SliceEquals(desired.RequestHeadersToRemove, actual.RequestHeadersToRemove) && !dcl.IsZeroValue(desired.RequestHeadersToRemove) {
		c.Config.Logger.Infof("Diff in RequestHeadersToRemove. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestHeadersToRemove), dcl.SprintResource(actual.RequestHeadersToRemove))
		return true
	}
	if actual.RequestHeadersToAdd == nil && desired.RequestHeadersToAdd != nil && !dcl.IsEmptyValueIndirect(desired.RequestHeadersToAdd) {
		c.Config.Logger.Infof("desired RequestHeadersToAdd %s - but actually nil", dcl.SprintResource(desired.RequestHeadersToAdd))
		return true
	}
	if compareUrlMapHeaderActionRequestHeadersToAddSlice(c, desired.RequestHeadersToAdd, actual.RequestHeadersToAdd) && !dcl.IsZeroValue(desired.RequestHeadersToAdd) {
		c.Config.Logger.Infof("Diff in RequestHeadersToAdd. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestHeadersToAdd), dcl.SprintResource(actual.RequestHeadersToAdd))
		return true
	}
	if actual.ResponseHeadersToRemove == nil && desired.ResponseHeadersToRemove != nil && !dcl.IsEmptyValueIndirect(desired.ResponseHeadersToRemove) {
		c.Config.Logger.Infof("desired ResponseHeadersToRemove %s - but actually nil", dcl.SprintResource(desired.ResponseHeadersToRemove))
		return true
	}
	if !dcl.SliceEquals(desired.ResponseHeadersToRemove, actual.ResponseHeadersToRemove) && !dcl.IsZeroValue(desired.ResponseHeadersToRemove) {
		c.Config.Logger.Infof("Diff in ResponseHeadersToRemove. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResponseHeadersToRemove), dcl.SprintResource(actual.ResponseHeadersToRemove))
		return true
	}
	if actual.ResponseHeadersToAdd == nil && desired.ResponseHeadersToAdd != nil && !dcl.IsEmptyValueIndirect(desired.ResponseHeadersToAdd) {
		c.Config.Logger.Infof("desired ResponseHeadersToAdd %s - but actually nil", dcl.SprintResource(desired.ResponseHeadersToAdd))
		return true
	}
	if compareUrlMapHeaderActionResponseHeadersToAddSlice(c, desired.ResponseHeadersToAdd, actual.ResponseHeadersToAdd) && !dcl.IsZeroValue(desired.ResponseHeadersToAdd) {
		c.Config.Logger.Infof("Diff in ResponseHeadersToAdd. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResponseHeadersToAdd), dcl.SprintResource(actual.ResponseHeadersToAdd))
		return true
	}
	return false
}
func compareUrlMapHeaderActionRequestHeadersToAddSlice(c *Client, desired, actual []UrlMapHeaderActionRequestHeadersToAdd) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapHeaderActionRequestHeadersToAdd, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapHeaderActionRequestHeadersToAdd(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapHeaderActionRequestHeadersToAdd, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapHeaderActionRequestHeadersToAdd(c *Client, desired, actual *UrlMapHeaderActionRequestHeadersToAdd) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HeaderName == nil && desired.HeaderName != nil && !dcl.IsEmptyValueIndirect(desired.HeaderName) {
		c.Config.Logger.Infof("desired HeaderName %s - but actually nil", dcl.SprintResource(desired.HeaderName))
		return true
	}
	if !reflect.DeepEqual(desired.HeaderName, actual.HeaderName) && !dcl.IsZeroValue(desired.HeaderName) && !(dcl.IsEmptyValueIndirect(desired.HeaderName) && dcl.IsZeroValue(actual.HeaderName)) {
		c.Config.Logger.Infof("Diff in HeaderName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderName), dcl.SprintResource(actual.HeaderName))
		return true
	}
	if actual.HeaderValue == nil && desired.HeaderValue != nil && !dcl.IsEmptyValueIndirect(desired.HeaderValue) {
		c.Config.Logger.Infof("desired HeaderValue %s - but actually nil", dcl.SprintResource(desired.HeaderValue))
		return true
	}
	if !reflect.DeepEqual(desired.HeaderValue, actual.HeaderValue) && !dcl.IsZeroValue(desired.HeaderValue) && !(dcl.IsEmptyValueIndirect(desired.HeaderValue) && dcl.IsZeroValue(actual.HeaderValue)) {
		c.Config.Logger.Infof("Diff in HeaderValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderValue), dcl.SprintResource(actual.HeaderValue))
		return true
	}
	if actual.Replace == nil && desired.Replace != nil && !dcl.IsEmptyValueIndirect(desired.Replace) {
		c.Config.Logger.Infof("desired Replace %s - but actually nil", dcl.SprintResource(desired.Replace))
		return true
	}
	if !reflect.DeepEqual(desired.Replace, actual.Replace) && !dcl.IsZeroValue(desired.Replace) && !(dcl.IsEmptyValueIndirect(desired.Replace) && dcl.IsZeroValue(actual.Replace)) {
		c.Config.Logger.Infof("Diff in Replace. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Replace), dcl.SprintResource(actual.Replace))
		return true
	}
	return false
}
func compareUrlMapHeaderActionResponseHeadersToAddSlice(c *Client, desired, actual []UrlMapHeaderActionResponseHeadersToAdd) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapHeaderActionResponseHeadersToAdd, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapHeaderActionResponseHeadersToAdd(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapHeaderActionResponseHeadersToAdd, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapHeaderActionResponseHeadersToAdd(c *Client, desired, actual *UrlMapHeaderActionResponseHeadersToAdd) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HeaderName == nil && desired.HeaderName != nil && !dcl.IsEmptyValueIndirect(desired.HeaderName) {
		c.Config.Logger.Infof("desired HeaderName %s - but actually nil", dcl.SprintResource(desired.HeaderName))
		return true
	}
	if !reflect.DeepEqual(desired.HeaderName, actual.HeaderName) && !dcl.IsZeroValue(desired.HeaderName) && !(dcl.IsEmptyValueIndirect(desired.HeaderName) && dcl.IsZeroValue(actual.HeaderName)) {
		c.Config.Logger.Infof("Diff in HeaderName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderName), dcl.SprintResource(actual.HeaderName))
		return true
	}
	if actual.HeaderValue == nil && desired.HeaderValue != nil && !dcl.IsEmptyValueIndirect(desired.HeaderValue) {
		c.Config.Logger.Infof("desired HeaderValue %s - but actually nil", dcl.SprintResource(desired.HeaderValue))
		return true
	}
	if !reflect.DeepEqual(desired.HeaderValue, actual.HeaderValue) && !dcl.IsZeroValue(desired.HeaderValue) && !(dcl.IsEmptyValueIndirect(desired.HeaderValue) && dcl.IsZeroValue(actual.HeaderValue)) {
		c.Config.Logger.Infof("Diff in HeaderValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderValue), dcl.SprintResource(actual.HeaderValue))
		return true
	}
	if actual.Replace == nil && desired.Replace != nil && !dcl.IsEmptyValueIndirect(desired.Replace) {
		c.Config.Logger.Infof("desired Replace %s - but actually nil", dcl.SprintResource(desired.Replace))
		return true
	}
	if !reflect.DeepEqual(desired.Replace, actual.Replace) && !dcl.IsZeroValue(desired.Replace) && !(dcl.IsEmptyValueIndirect(desired.Replace) && dcl.IsZeroValue(actual.Replace)) {
		c.Config.Logger.Infof("Diff in Replace. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Replace), dcl.SprintResource(actual.Replace))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionUrlRewriteSlice(c *Client, desired, actual []UrlMapDefaultRouteActionUrlRewrite) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionUrlRewrite, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionUrlRewrite(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionUrlRewrite, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionUrlRewrite(c *Client, desired, actual *UrlMapDefaultRouteActionUrlRewrite) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PathPrefixRewrite == nil && desired.PathPrefixRewrite != nil && !dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) {
		c.Config.Logger.Infof("desired PathPrefixRewrite %s - but actually nil", dcl.SprintResource(desired.PathPrefixRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.PathPrefixRewrite, actual.PathPrefixRewrite) && !dcl.IsZeroValue(desired.PathPrefixRewrite) && !(dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) && dcl.IsZeroValue(actual.PathPrefixRewrite)) {
		c.Config.Logger.Infof("Diff in PathPrefixRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathPrefixRewrite), dcl.SprintResource(actual.PathPrefixRewrite))
		return true
	}
	if actual.HostRewrite == nil && desired.HostRewrite != nil && !dcl.IsEmptyValueIndirect(desired.HostRewrite) {
		c.Config.Logger.Infof("desired HostRewrite %s - but actually nil", dcl.SprintResource(desired.HostRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.HostRewrite, actual.HostRewrite) && !dcl.IsZeroValue(desired.HostRewrite) && !(dcl.IsEmptyValueIndirect(desired.HostRewrite) && dcl.IsZeroValue(actual.HostRewrite)) {
		c.Config.Logger.Infof("Diff in HostRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRewrite), dcl.SprintResource(actual.HostRewrite))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionTimeoutSlice(c *Client, desired, actual []UrlMapDefaultRouteActionTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionTimeout(c *Client, desired, actual *UrlMapDefaultRouteActionTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionRetryPolicySlice(c *Client, desired, actual []UrlMapDefaultRouteActionRetryPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionRetryPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionRetryPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionRetryPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionRetryPolicy(c *Client, desired, actual *UrlMapDefaultRouteActionRetryPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RetryCondition == nil && desired.RetryCondition != nil && !dcl.IsEmptyValueIndirect(desired.RetryCondition) {
		c.Config.Logger.Infof("desired RetryCondition %s - but actually nil", dcl.SprintResource(desired.RetryCondition))
		return true
	}
	if !dcl.SliceEquals(desired.RetryCondition, actual.RetryCondition) && !dcl.IsZeroValue(desired.RetryCondition) {
		c.Config.Logger.Infof("Diff in RetryCondition. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryCondition), dcl.SprintResource(actual.RetryCondition))
		return true
	}
	if actual.NumRetries == nil && desired.NumRetries != nil && !dcl.IsEmptyValueIndirect(desired.NumRetries) {
		c.Config.Logger.Infof("desired NumRetries %s - but actually nil", dcl.SprintResource(desired.NumRetries))
		return true
	}
	if !reflect.DeepEqual(desired.NumRetries, actual.NumRetries) && !dcl.IsZeroValue(desired.NumRetries) && !(dcl.IsEmptyValueIndirect(desired.NumRetries) && dcl.IsZeroValue(actual.NumRetries)) {
		c.Config.Logger.Infof("Diff in NumRetries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumRetries), dcl.SprintResource(actual.NumRetries))
		return true
	}
	if actual.PerTryTimeout == nil && desired.PerTryTimeout != nil && !dcl.IsEmptyValueIndirect(desired.PerTryTimeout) {
		c.Config.Logger.Infof("desired PerTryTimeout %s - but actually nil", dcl.SprintResource(desired.PerTryTimeout))
		return true
	}
	if compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, desired.PerTryTimeout, actual.PerTryTimeout) && !dcl.IsZeroValue(desired.PerTryTimeout) {
		c.Config.Logger.Infof("Diff in PerTryTimeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerTryTimeout), dcl.SprintResource(actual.PerTryTimeout))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, desired, actual []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionRetryPolicyPerTryTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionRetryPolicyPerTryTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c *Client, desired, actual *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionRequestMirrorPolicySlice(c *Client, desired, actual []UrlMapDefaultRouteActionRequestMirrorPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionRequestMirrorPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionRequestMirrorPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionRequestMirrorPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionRequestMirrorPolicy(c *Client, desired, actual *UrlMapDefaultRouteActionRequestMirrorPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionCorsPolicySlice(c *Client, desired, actual []UrlMapDefaultRouteActionCorsPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionCorsPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionCorsPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionCorsPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionCorsPolicy(c *Client, desired, actual *UrlMapDefaultRouteActionCorsPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AllowOrigin == nil && desired.AllowOrigin != nil && !dcl.IsEmptyValueIndirect(desired.AllowOrigin) {
		c.Config.Logger.Infof("desired AllowOrigin %s - but actually nil", dcl.SprintResource(desired.AllowOrigin))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOrigin, actual.AllowOrigin) && !dcl.IsZeroValue(desired.AllowOrigin) {
		c.Config.Logger.Infof("Diff in AllowOrigin. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOrigin), dcl.SprintResource(actual.AllowOrigin))
		return true
	}
	if actual.AllowOriginRegex == nil && desired.AllowOriginRegex != nil && !dcl.IsEmptyValueIndirect(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("desired AllowOriginRegex %s - but actually nil", dcl.SprintResource(desired.AllowOriginRegex))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOriginRegex, actual.AllowOriginRegex) && !dcl.IsZeroValue(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("Diff in AllowOriginRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOriginRegex), dcl.SprintResource(actual.AllowOriginRegex))
		return true
	}
	if actual.AllowMethod == nil && desired.AllowMethod != nil && !dcl.IsEmptyValueIndirect(desired.AllowMethod) {
		c.Config.Logger.Infof("desired AllowMethod %s - but actually nil", dcl.SprintResource(desired.AllowMethod))
		return true
	}
	if !dcl.SliceEquals(desired.AllowMethod, actual.AllowMethod) && !dcl.IsZeroValue(desired.AllowMethod) {
		c.Config.Logger.Infof("Diff in AllowMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowMethod), dcl.SprintResource(actual.AllowMethod))
		return true
	}
	if actual.AllowHeader == nil && desired.AllowHeader != nil && !dcl.IsEmptyValueIndirect(desired.AllowHeader) {
		c.Config.Logger.Infof("desired AllowHeader %s - but actually nil", dcl.SprintResource(desired.AllowHeader))
		return true
	}
	if !dcl.SliceEquals(desired.AllowHeader, actual.AllowHeader) && !dcl.IsZeroValue(desired.AllowHeader) {
		c.Config.Logger.Infof("Diff in AllowHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowHeader), dcl.SprintResource(actual.AllowHeader))
		return true
	}
	if actual.ExposeHeader == nil && desired.ExposeHeader != nil && !dcl.IsEmptyValueIndirect(desired.ExposeHeader) {
		c.Config.Logger.Infof("desired ExposeHeader %s - but actually nil", dcl.SprintResource(desired.ExposeHeader))
		return true
	}
	if !dcl.SliceEquals(desired.ExposeHeader, actual.ExposeHeader) && !dcl.IsZeroValue(desired.ExposeHeader) {
		c.Config.Logger.Infof("Diff in ExposeHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeHeader), dcl.SprintResource(actual.ExposeHeader))
		return true
	}
	if actual.MaxAge == nil && desired.MaxAge != nil && !dcl.IsEmptyValueIndirect(desired.MaxAge) {
		c.Config.Logger.Infof("desired MaxAge %s - but actually nil", dcl.SprintResource(desired.MaxAge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxAge, actual.MaxAge) && !dcl.IsZeroValue(desired.MaxAge) && !(dcl.IsEmptyValueIndirect(desired.MaxAge) && dcl.IsZeroValue(actual.MaxAge)) {
		c.Config.Logger.Infof("Diff in MaxAge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxAge), dcl.SprintResource(actual.MaxAge))
		return true
	}
	if actual.AllowCredentials == nil && desired.AllowCredentials != nil && !dcl.IsEmptyValueIndirect(desired.AllowCredentials) {
		c.Config.Logger.Infof("desired AllowCredentials %s - but actually nil", dcl.SprintResource(desired.AllowCredentials))
		return true
	}
	if !reflect.DeepEqual(desired.AllowCredentials, actual.AllowCredentials) && !dcl.IsZeroValue(desired.AllowCredentials) && !(dcl.IsEmptyValueIndirect(desired.AllowCredentials) && dcl.IsZeroValue(actual.AllowCredentials)) {
		c.Config.Logger.Infof("Diff in AllowCredentials. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowCredentials), dcl.SprintResource(actual.AllowCredentials))
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionFaultInjectionPolicySlice(c *Client, desired, actual []UrlMapDefaultRouteActionFaultInjectionPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionFaultInjectionPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionFaultInjectionPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionFaultInjectionPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicy(c *Client, desired, actual *UrlMapDefaultRouteActionFaultInjectionPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Delay == nil && desired.Delay != nil && !dcl.IsEmptyValueIndirect(desired.Delay) {
		c.Config.Logger.Infof("desired Delay %s - but actually nil", dcl.SprintResource(desired.Delay))
		return true
	}
	if compareUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, desired.Delay, actual.Delay) && !dcl.IsZeroValue(desired.Delay) {
		c.Config.Logger.Infof("Diff in Delay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Delay), dcl.SprintResource(actual.Delay))
		return true
	}
	if actual.Abort == nil && desired.Abort != nil && !dcl.IsEmptyValueIndirect(desired.Abort) {
		c.Config.Logger.Infof("desired Abort %s - but actually nil", dcl.SprintResource(desired.Abort))
		return true
	}
	if compareUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, desired.Abort, actual.Abort) && !dcl.IsZeroValue(desired.Abort) {
		c.Config.Logger.Infof("Diff in Abort. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Abort), dcl.SprintResource(actual.Abort))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice(c *Client, desired, actual []UrlMapDefaultRouteActionFaultInjectionPolicyDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c *Client, desired, actual *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FixedDelay == nil && desired.FixedDelay != nil && !dcl.IsEmptyValueIndirect(desired.FixedDelay) {
		c.Config.Logger.Infof("desired FixedDelay %s - but actually nil", dcl.SprintResource(desired.FixedDelay))
		return true
	}
	if compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, desired.FixedDelay, actual.FixedDelay) && !dcl.IsZeroValue(desired.FixedDelay) {
		c.Config.Logger.Infof("Diff in FixedDelay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedDelay), dcl.SprintResource(actual.FixedDelay))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, desired, actual []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, desired, actual *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice(c *Client, desired, actual []UrlMapDefaultRouteActionFaultInjectionPolicyAbort) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyAbort, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultRouteActionFaultInjectionPolicyAbort, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c *Client, desired, actual *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HttpStatus == nil && desired.HttpStatus != nil && !dcl.IsEmptyValueIndirect(desired.HttpStatus) {
		c.Config.Logger.Infof("desired HttpStatus %s - but actually nil", dcl.SprintResource(desired.HttpStatus))
		return true
	}
	if !reflect.DeepEqual(desired.HttpStatus, actual.HttpStatus) && !dcl.IsZeroValue(desired.HttpStatus) && !(dcl.IsEmptyValueIndirect(desired.HttpStatus) && dcl.IsZeroValue(actual.HttpStatus)) {
		c.Config.Logger.Infof("Diff in HttpStatus. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpStatus), dcl.SprintResource(actual.HttpStatus))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapDefaultUrlRedirectSlice(c *Client, desired, actual []UrlMapDefaultUrlRedirect) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultUrlRedirect, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultUrlRedirect(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultUrlRedirect, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultUrlRedirect(c *Client, desired, actual *UrlMapDefaultUrlRedirect) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HostRedirect == nil && desired.HostRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HostRedirect) {
		c.Config.Logger.Infof("desired HostRedirect %s - but actually nil", dcl.SprintResource(desired.HostRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HostRedirect, actual.HostRedirect) && !dcl.IsZeroValue(desired.HostRedirect) && !(dcl.IsEmptyValueIndirect(desired.HostRedirect) && dcl.IsZeroValue(actual.HostRedirect)) {
		c.Config.Logger.Infof("Diff in HostRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRedirect), dcl.SprintResource(actual.HostRedirect))
		return true
	}
	if actual.PathRedirect == nil && desired.PathRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PathRedirect) {
		c.Config.Logger.Infof("desired PathRedirect %s - but actually nil", dcl.SprintResource(desired.PathRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PathRedirect, actual.PathRedirect) && !dcl.IsZeroValue(desired.PathRedirect) && !(dcl.IsEmptyValueIndirect(desired.PathRedirect) && dcl.IsZeroValue(actual.PathRedirect)) {
		c.Config.Logger.Infof("Diff in PathRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathRedirect), dcl.SprintResource(actual.PathRedirect))
		return true
	}
	if actual.PrefixRedirect == nil && desired.PrefixRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PrefixRedirect) {
		c.Config.Logger.Infof("desired PrefixRedirect %s - but actually nil", dcl.SprintResource(desired.PrefixRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixRedirect, actual.PrefixRedirect) && !dcl.IsZeroValue(desired.PrefixRedirect) && !(dcl.IsEmptyValueIndirect(desired.PrefixRedirect) && dcl.IsZeroValue(actual.PrefixRedirect)) {
		c.Config.Logger.Infof("Diff in PrefixRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixRedirect), dcl.SprintResource(actual.PrefixRedirect))
		return true
	}
	if actual.RedirectResponseCode == nil && desired.RedirectResponseCode != nil && !dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) {
		c.Config.Logger.Infof("desired RedirectResponseCode %s - but actually nil", dcl.SprintResource(desired.RedirectResponseCode))
		return true
	}
	if !reflect.DeepEqual(desired.RedirectResponseCode, actual.RedirectResponseCode) && !dcl.IsZeroValue(desired.RedirectResponseCode) && !(dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) && dcl.IsZeroValue(actual.RedirectResponseCode)) {
		c.Config.Logger.Infof("Diff in RedirectResponseCode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RedirectResponseCode), dcl.SprintResource(actual.RedirectResponseCode))
		return true
	}
	if actual.HttpsRedirect == nil && desired.HttpsRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HttpsRedirect) {
		c.Config.Logger.Infof("desired HttpsRedirect %s - but actually nil", dcl.SprintResource(desired.HttpsRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HttpsRedirect, actual.HttpsRedirect) && !dcl.IsZeroValue(desired.HttpsRedirect) && !(dcl.IsEmptyValueIndirect(desired.HttpsRedirect) && dcl.IsZeroValue(actual.HttpsRedirect)) {
		c.Config.Logger.Infof("Diff in HttpsRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpsRedirect), dcl.SprintResource(actual.HttpsRedirect))
		return true
	}
	if actual.StripQuery == nil && desired.StripQuery != nil && !dcl.IsEmptyValueIndirect(desired.StripQuery) {
		c.Config.Logger.Infof("desired StripQuery %s - but actually nil", dcl.SprintResource(desired.StripQuery))
		return true
	}
	if !reflect.DeepEqual(desired.StripQuery, actual.StripQuery) && !dcl.IsZeroValue(desired.StripQuery) && !(dcl.IsEmptyValueIndirect(desired.StripQuery) && dcl.IsZeroValue(actual.StripQuery)) {
		c.Config.Logger.Infof("Diff in StripQuery. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StripQuery), dcl.SprintResource(actual.StripQuery))
		return true
	}
	return false
}
func compareUrlMapHostRuleSlice(c *Client, desired, actual []UrlMapHostRule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapHostRule, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapHostRule(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapHostRule, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapHostRuleSets(c *Client, desired, actual []UrlMapHostRule) (toAdd, toRemove []UrlMapHostRule) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]UrlMapHostRule)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]UrlMapHostRule)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]UrlMapHostRule, 0)
	toRemove = make([]UrlMapHostRule, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareUrlMapHostRule(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareUrlMapHostRule(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

func compareUrlMapHostRule(c *Client, desired, actual *UrlMapHostRule) bool {
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
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.Host == nil && desired.Host != nil && !dcl.IsEmptyValueIndirect(desired.Host) {
		c.Config.Logger.Infof("desired Host %s - but actually nil", dcl.SprintResource(desired.Host))
		return true
	}
	if !dcl.SliceEquals(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) {
		c.Config.Logger.Infof("Diff in Host. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	if actual.PathMatcher == nil && desired.PathMatcher != nil && !dcl.IsEmptyValueIndirect(desired.PathMatcher) {
		c.Config.Logger.Infof("desired PathMatcher %s - but actually nil", dcl.SprintResource(desired.PathMatcher))
		return true
	}
	if !reflect.DeepEqual(desired.PathMatcher, actual.PathMatcher) && !dcl.IsZeroValue(desired.PathMatcher) && !(dcl.IsEmptyValueIndirect(desired.PathMatcher) && dcl.IsZeroValue(actual.PathMatcher)) {
		c.Config.Logger.Infof("Diff in PathMatcher. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathMatcher), dcl.SprintResource(actual.PathMatcher))
		return true
	}
	return false
}
func compareUrlMapPathMatcherSlice(c *Client, desired, actual []UrlMapPathMatcher) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcher, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcher(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcher, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherSets(c *Client, desired, actual []UrlMapPathMatcher) (toAdd, toRemove []UrlMapPathMatcher) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]UrlMapPathMatcher)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]UrlMapPathMatcher)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]UrlMapPathMatcher, 0)
	toRemove = make([]UrlMapPathMatcher, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareUrlMapPathMatcher(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareUrlMapPathMatcher(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

func compareUrlMapPathMatcher(c *Client, desired, actual *UrlMapPathMatcher) bool {
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
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.DefaultService == nil && desired.DefaultService != nil && !dcl.IsEmptyValueIndirect(desired.DefaultService) {
		c.Config.Logger.Infof("desired DefaultService %s - but actually nil", dcl.SprintResource(desired.DefaultService))
		return true
	}
	if !reflect.DeepEqual(desired.DefaultService, actual.DefaultService) && !dcl.IsZeroValue(desired.DefaultService) && !(dcl.IsEmptyValueIndirect(desired.DefaultService) && dcl.IsZeroValue(actual.DefaultService)) {
		c.Config.Logger.Infof("Diff in DefaultService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultService), dcl.SprintResource(actual.DefaultService))
		return true
	}
	if actual.DefaultRouteAction == nil && desired.DefaultRouteAction != nil && !dcl.IsEmptyValueIndirect(desired.DefaultRouteAction) {
		c.Config.Logger.Infof("desired DefaultRouteAction %s - but actually nil", dcl.SprintResource(desired.DefaultRouteAction))
		return true
	}
	if compareUrlMapDefaultRouteAction(c, desired.DefaultRouteAction, actual.DefaultRouteAction) && !dcl.IsZeroValue(desired.DefaultRouteAction) {
		c.Config.Logger.Infof("Diff in DefaultRouteAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultRouteAction), dcl.SprintResource(actual.DefaultRouteAction))
		return true
	}
	if actual.DefaultUrlRedirect == nil && desired.DefaultUrlRedirect != nil && !dcl.IsEmptyValueIndirect(desired.DefaultUrlRedirect) {
		c.Config.Logger.Infof("desired DefaultUrlRedirect %s - but actually nil", dcl.SprintResource(desired.DefaultUrlRedirect))
		return true
	}
	if compareUrlMapPathMatcherDefaultUrlRedirect(c, desired.DefaultUrlRedirect, actual.DefaultUrlRedirect) && !dcl.IsZeroValue(desired.DefaultUrlRedirect) {
		c.Config.Logger.Infof("Diff in DefaultUrlRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultUrlRedirect), dcl.SprintResource(actual.DefaultUrlRedirect))
		return true
	}
	if actual.PathRule == nil && desired.PathRule != nil && !dcl.IsEmptyValueIndirect(desired.PathRule) {
		c.Config.Logger.Infof("desired PathRule %s - but actually nil", dcl.SprintResource(desired.PathRule))
		return true
	}
	if compareUrlMapPathMatcherPathRuleSlice(c, desired.PathRule, actual.PathRule) && !dcl.IsZeroValue(desired.PathRule) {
		c.Config.Logger.Infof("Diff in PathRule. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathRule), dcl.SprintResource(actual.PathRule))
		return true
	}
	if actual.RouteRule == nil && desired.RouteRule != nil && !dcl.IsEmptyValueIndirect(desired.RouteRule) {
		c.Config.Logger.Infof("desired RouteRule %s - but actually nil", dcl.SprintResource(desired.RouteRule))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleSlice(c, desired.RouteRule, actual.RouteRule) && !dcl.IsZeroValue(desired.RouteRule) {
		c.Config.Logger.Infof("Diff in RouteRule. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RouteRule), dcl.SprintResource(actual.RouteRule))
		return true
	}
	if actual.HeaderAction == nil && desired.HeaderAction != nil && !dcl.IsEmptyValueIndirect(desired.HeaderAction) {
		c.Config.Logger.Infof("desired HeaderAction %s - but actually nil", dcl.SprintResource(desired.HeaderAction))
		return true
	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) && !dcl.IsZeroValue(desired.HeaderAction) {
		c.Config.Logger.Infof("Diff in HeaderAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderAction), dcl.SprintResource(actual.HeaderAction))
		return true
	}
	return false
}
func compareUrlMapPathMatcherDefaultUrlRedirectSlice(c *Client, desired, actual []UrlMapPathMatcherDefaultUrlRedirect) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherDefaultUrlRedirect, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherDefaultUrlRedirect(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherDefaultUrlRedirect, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherDefaultUrlRedirect(c *Client, desired, actual *UrlMapPathMatcherDefaultUrlRedirect) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HostRedirect == nil && desired.HostRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HostRedirect) {
		c.Config.Logger.Infof("desired HostRedirect %s - but actually nil", dcl.SprintResource(desired.HostRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HostRedirect, actual.HostRedirect) && !dcl.IsZeroValue(desired.HostRedirect) && !(dcl.IsEmptyValueIndirect(desired.HostRedirect) && dcl.IsZeroValue(actual.HostRedirect)) {
		c.Config.Logger.Infof("Diff in HostRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRedirect), dcl.SprintResource(actual.HostRedirect))
		return true
	}
	if actual.PathRedirect == nil && desired.PathRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PathRedirect) {
		c.Config.Logger.Infof("desired PathRedirect %s - but actually nil", dcl.SprintResource(desired.PathRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PathRedirect, actual.PathRedirect) && !dcl.IsZeroValue(desired.PathRedirect) && !(dcl.IsEmptyValueIndirect(desired.PathRedirect) && dcl.IsZeroValue(actual.PathRedirect)) {
		c.Config.Logger.Infof("Diff in PathRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathRedirect), dcl.SprintResource(actual.PathRedirect))
		return true
	}
	if actual.PrefixRedirect == nil && desired.PrefixRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PrefixRedirect) {
		c.Config.Logger.Infof("desired PrefixRedirect %s - but actually nil", dcl.SprintResource(desired.PrefixRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixRedirect, actual.PrefixRedirect) && !dcl.IsZeroValue(desired.PrefixRedirect) && !(dcl.IsEmptyValueIndirect(desired.PrefixRedirect) && dcl.IsZeroValue(actual.PrefixRedirect)) {
		c.Config.Logger.Infof("Diff in PrefixRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixRedirect), dcl.SprintResource(actual.PrefixRedirect))
		return true
	}
	if actual.RedirectResponseCode == nil && desired.RedirectResponseCode != nil && !dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) {
		c.Config.Logger.Infof("desired RedirectResponseCode %s - but actually nil", dcl.SprintResource(desired.RedirectResponseCode))
		return true
	}
	if !reflect.DeepEqual(desired.RedirectResponseCode, actual.RedirectResponseCode) && !dcl.IsZeroValue(desired.RedirectResponseCode) && !(dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) && dcl.IsZeroValue(actual.RedirectResponseCode)) {
		c.Config.Logger.Infof("Diff in RedirectResponseCode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RedirectResponseCode), dcl.SprintResource(actual.RedirectResponseCode))
		return true
	}
	if actual.HttpsRedirect == nil && desired.HttpsRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HttpsRedirect) {
		c.Config.Logger.Infof("desired HttpsRedirect %s - but actually nil", dcl.SprintResource(desired.HttpsRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HttpsRedirect, actual.HttpsRedirect) && !dcl.IsZeroValue(desired.HttpsRedirect) && !(dcl.IsEmptyValueIndirect(desired.HttpsRedirect) && dcl.IsZeroValue(actual.HttpsRedirect)) {
		c.Config.Logger.Infof("Diff in HttpsRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpsRedirect), dcl.SprintResource(actual.HttpsRedirect))
		return true
	}
	if actual.StripQuery == nil && desired.StripQuery != nil && !dcl.IsEmptyValueIndirect(desired.StripQuery) {
		c.Config.Logger.Infof("desired StripQuery %s - but actually nil", dcl.SprintResource(desired.StripQuery))
		return true
	}
	if !reflect.DeepEqual(desired.StripQuery, actual.StripQuery) && !dcl.IsZeroValue(desired.StripQuery) && !(dcl.IsEmptyValueIndirect(desired.StripQuery) && dcl.IsZeroValue(actual.StripQuery)) {
		c.Config.Logger.Infof("Diff in StripQuery. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StripQuery), dcl.SprintResource(actual.StripQuery))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleSlice(c *Client, desired, actual []UrlMapPathMatcherPathRule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRule, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRule(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRule, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRule(c *Client, desired, actual *UrlMapPathMatcherPathRule) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	if actual.RouteAction == nil && desired.RouteAction != nil && !dcl.IsEmptyValueIndirect(desired.RouteAction) {
		c.Config.Logger.Infof("desired RouteAction %s - but actually nil", dcl.SprintResource(desired.RouteAction))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteAction(c, desired.RouteAction, actual.RouteAction) && !dcl.IsZeroValue(desired.RouteAction) {
		c.Config.Logger.Infof("Diff in RouteAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RouteAction), dcl.SprintResource(actual.RouteAction))
		return true
	}
	if actual.UrlRedirect == nil && desired.UrlRedirect != nil && !dcl.IsEmptyValueIndirect(desired.UrlRedirect) {
		c.Config.Logger.Infof("desired UrlRedirect %s - but actually nil", dcl.SprintResource(desired.UrlRedirect))
		return true
	}
	if compareUrlMapPathMatcherPathRuleUrlRedirect(c, desired.UrlRedirect, actual.UrlRedirect) && !dcl.IsZeroValue(desired.UrlRedirect) {
		c.Config.Logger.Infof("Diff in UrlRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlRedirect), dcl.SprintResource(actual.UrlRedirect))
		return true
	}
	if actual.Path == nil && desired.Path != nil && !dcl.IsEmptyValueIndirect(desired.Path) {
		c.Config.Logger.Infof("desired Path %s - but actually nil", dcl.SprintResource(desired.Path))
		return true
	}
	if !dcl.SliceEquals(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteAction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteAction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteAction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteAction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteAction(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteAction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.WeightedBackendService == nil && desired.WeightedBackendService != nil && !dcl.IsEmptyValueIndirect(desired.WeightedBackendService) {
		c.Config.Logger.Infof("desired WeightedBackendService %s - but actually nil", dcl.SprintResource(desired.WeightedBackendService))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c, desired.WeightedBackendService, actual.WeightedBackendService) && !dcl.IsZeroValue(desired.WeightedBackendService) {
		c.Config.Logger.Infof("Diff in WeightedBackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WeightedBackendService), dcl.SprintResource(actual.WeightedBackendService))
		return true
	}
	if actual.UrlRewrite == nil && desired.UrlRewrite != nil && !dcl.IsEmptyValueIndirect(desired.UrlRewrite) {
		c.Config.Logger.Infof("desired UrlRewrite %s - but actually nil", dcl.SprintResource(desired.UrlRewrite))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, desired.UrlRewrite, actual.UrlRewrite) && !dcl.IsZeroValue(desired.UrlRewrite) {
		c.Config.Logger.Infof("Diff in UrlRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlRewrite), dcl.SprintResource(actual.UrlRewrite))
		return true
	}
	if actual.Timeout == nil && desired.Timeout != nil && !dcl.IsEmptyValueIndirect(desired.Timeout) {
		c.Config.Logger.Infof("desired Timeout %s - but actually nil", dcl.SprintResource(desired.Timeout))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionTimeout(c, desired.Timeout, actual.Timeout) && !dcl.IsZeroValue(desired.Timeout) {
		c.Config.Logger.Infof("Diff in Timeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Timeout), dcl.SprintResource(actual.Timeout))
		return true
	}
	if actual.RetryPolicy == nil && desired.RetryPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RetryPolicy) {
		c.Config.Logger.Infof("desired RetryPolicy %s - but actually nil", dcl.SprintResource(desired.RetryPolicy))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, desired.RetryPolicy, actual.RetryPolicy) && !dcl.IsZeroValue(desired.RetryPolicy) {
		c.Config.Logger.Infof("Diff in RetryPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryPolicy), dcl.SprintResource(actual.RetryPolicy))
		return true
	}
	if actual.RequestMirrorPolicy == nil && desired.RequestMirrorPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("desired RequestMirrorPolicy %s - but actually nil", dcl.SprintResource(desired.RequestMirrorPolicy))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, desired.RequestMirrorPolicy, actual.RequestMirrorPolicy) && !dcl.IsZeroValue(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("Diff in RequestMirrorPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestMirrorPolicy), dcl.SprintResource(actual.RequestMirrorPolicy))
		return true
	}
	if actual.CorsPolicy == nil && desired.CorsPolicy != nil && !dcl.IsEmptyValueIndirect(desired.CorsPolicy) {
		c.Config.Logger.Infof("desired CorsPolicy %s - but actually nil", dcl.SprintResource(desired.CorsPolicy))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, desired.CorsPolicy, actual.CorsPolicy) && !dcl.IsZeroValue(desired.CorsPolicy) {
		c.Config.Logger.Infof("Diff in CorsPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CorsPolicy), dcl.SprintResource(actual.CorsPolicy))
		return true
	}
	if actual.FaultInjectionPolicy == nil && desired.FaultInjectionPolicy != nil && !dcl.IsEmptyValueIndirect(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("desired FaultInjectionPolicy %s - but actually nil", dcl.SprintResource(desired.FaultInjectionPolicy))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, desired.FaultInjectionPolicy, actual.FaultInjectionPolicy) && !dcl.IsZeroValue(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("Diff in FaultInjectionPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FaultInjectionPolicy), dcl.SprintResource(actual.FaultInjectionPolicy))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionWeightedBackendService, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionWeightedBackendService, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	if actual.Weight == nil && desired.Weight != nil && !dcl.IsEmptyValueIndirect(desired.Weight) {
		c.Config.Logger.Infof("desired Weight %s - but actually nil", dcl.SprintResource(desired.Weight))
		return true
	}
	if !reflect.DeepEqual(desired.Weight, actual.Weight) && !dcl.IsZeroValue(desired.Weight) && !(dcl.IsEmptyValueIndirect(desired.Weight) && dcl.IsZeroValue(actual.Weight)) {
		c.Config.Logger.Infof("Diff in Weight. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Weight), dcl.SprintResource(actual.Weight))
		return true
	}
	if actual.HeaderAction == nil && desired.HeaderAction != nil && !dcl.IsEmptyValueIndirect(desired.HeaderAction) {
		c.Config.Logger.Infof("desired HeaderAction %s - but actually nil", dcl.SprintResource(desired.HeaderAction))
		return true
	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) && !dcl.IsZeroValue(desired.HeaderAction) {
		c.Config.Logger.Infof("Diff in HeaderAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderAction), dcl.SprintResource(actual.HeaderAction))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionUrlRewrite) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionUrlRewrite, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionUrlRewrite, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PathPrefixRewrite == nil && desired.PathPrefixRewrite != nil && !dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) {
		c.Config.Logger.Infof("desired PathPrefixRewrite %s - but actually nil", dcl.SprintResource(desired.PathPrefixRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.PathPrefixRewrite, actual.PathPrefixRewrite) && !dcl.IsZeroValue(desired.PathPrefixRewrite) && !(dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) && dcl.IsZeroValue(actual.PathPrefixRewrite)) {
		c.Config.Logger.Infof("Diff in PathPrefixRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathPrefixRewrite), dcl.SprintResource(actual.PathPrefixRewrite))
		return true
	}
	if actual.HostRewrite == nil && desired.HostRewrite != nil && !dcl.IsEmptyValueIndirect(desired.HostRewrite) {
		c.Config.Logger.Infof("desired HostRewrite %s - but actually nil", dcl.SprintResource(desired.HostRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.HostRewrite, actual.HostRewrite) && !dcl.IsZeroValue(desired.HostRewrite) && !(dcl.IsEmptyValueIndirect(desired.HostRewrite) && dcl.IsZeroValue(actual.HostRewrite)) {
		c.Config.Logger.Infof("Diff in HostRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRewrite), dcl.SprintResource(actual.HostRewrite))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionTimeoutSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionTimeout(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionRetryPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionRetryPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionRetryPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RetryCondition == nil && desired.RetryCondition != nil && !dcl.IsEmptyValueIndirect(desired.RetryCondition) {
		c.Config.Logger.Infof("desired RetryCondition %s - but actually nil", dcl.SprintResource(desired.RetryCondition))
		return true
	}
	if !dcl.SliceEquals(desired.RetryCondition, actual.RetryCondition) && !dcl.IsZeroValue(desired.RetryCondition) {
		c.Config.Logger.Infof("Diff in RetryCondition. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryCondition), dcl.SprintResource(actual.RetryCondition))
		return true
	}
	if actual.NumRetries == nil && desired.NumRetries != nil && !dcl.IsEmptyValueIndirect(desired.NumRetries) {
		c.Config.Logger.Infof("desired NumRetries %s - but actually nil", dcl.SprintResource(desired.NumRetries))
		return true
	}
	if !reflect.DeepEqual(desired.NumRetries, actual.NumRetries) && !dcl.IsZeroValue(desired.NumRetries) && !(dcl.IsEmptyValueIndirect(desired.NumRetries) && dcl.IsZeroValue(actual.NumRetries)) {
		c.Config.Logger.Infof("Diff in NumRetries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumRetries), dcl.SprintResource(actual.NumRetries))
		return true
	}
	if actual.PerTryTimeout == nil && desired.PerTryTimeout != nil && !dcl.IsEmptyValueIndirect(desired.PerTryTimeout) {
		c.Config.Logger.Infof("desired PerTryTimeout %s - but actually nil", dcl.SprintResource(desired.PerTryTimeout))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, desired.PerTryTimeout, actual.PerTryTimeout) && !dcl.IsZeroValue(desired.PerTryTimeout) {
		c.Config.Logger.Infof("Diff in PerTryTimeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerTryTimeout), dcl.SprintResource(actual.PerTryTimeout))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionCorsPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionCorsPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionCorsPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AllowOrigin == nil && desired.AllowOrigin != nil && !dcl.IsEmptyValueIndirect(desired.AllowOrigin) {
		c.Config.Logger.Infof("desired AllowOrigin %s - but actually nil", dcl.SprintResource(desired.AllowOrigin))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOrigin, actual.AllowOrigin) && !dcl.IsZeroValue(desired.AllowOrigin) {
		c.Config.Logger.Infof("Diff in AllowOrigin. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOrigin), dcl.SprintResource(actual.AllowOrigin))
		return true
	}
	if actual.AllowOriginRegex == nil && desired.AllowOriginRegex != nil && !dcl.IsEmptyValueIndirect(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("desired AllowOriginRegex %s - but actually nil", dcl.SprintResource(desired.AllowOriginRegex))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOriginRegex, actual.AllowOriginRegex) && !dcl.IsZeroValue(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("Diff in AllowOriginRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOriginRegex), dcl.SprintResource(actual.AllowOriginRegex))
		return true
	}
	if actual.AllowMethod == nil && desired.AllowMethod != nil && !dcl.IsEmptyValueIndirect(desired.AllowMethod) {
		c.Config.Logger.Infof("desired AllowMethod %s - but actually nil", dcl.SprintResource(desired.AllowMethod))
		return true
	}
	if !dcl.SliceEquals(desired.AllowMethod, actual.AllowMethod) && !dcl.IsZeroValue(desired.AllowMethod) {
		c.Config.Logger.Infof("Diff in AllowMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowMethod), dcl.SprintResource(actual.AllowMethod))
		return true
	}
	if actual.AllowHeader == nil && desired.AllowHeader != nil && !dcl.IsEmptyValueIndirect(desired.AllowHeader) {
		c.Config.Logger.Infof("desired AllowHeader %s - but actually nil", dcl.SprintResource(desired.AllowHeader))
		return true
	}
	if !dcl.SliceEquals(desired.AllowHeader, actual.AllowHeader) && !dcl.IsZeroValue(desired.AllowHeader) {
		c.Config.Logger.Infof("Diff in AllowHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowHeader), dcl.SprintResource(actual.AllowHeader))
		return true
	}
	if actual.ExposeHeader == nil && desired.ExposeHeader != nil && !dcl.IsEmptyValueIndirect(desired.ExposeHeader) {
		c.Config.Logger.Infof("desired ExposeHeader %s - but actually nil", dcl.SprintResource(desired.ExposeHeader))
		return true
	}
	if !dcl.SliceEquals(desired.ExposeHeader, actual.ExposeHeader) && !dcl.IsZeroValue(desired.ExposeHeader) {
		c.Config.Logger.Infof("Diff in ExposeHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeHeader), dcl.SprintResource(actual.ExposeHeader))
		return true
	}
	if actual.MaxAge == nil && desired.MaxAge != nil && !dcl.IsEmptyValueIndirect(desired.MaxAge) {
		c.Config.Logger.Infof("desired MaxAge %s - but actually nil", dcl.SprintResource(desired.MaxAge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxAge, actual.MaxAge) && !dcl.IsZeroValue(desired.MaxAge) && !(dcl.IsEmptyValueIndirect(desired.MaxAge) && dcl.IsZeroValue(actual.MaxAge)) {
		c.Config.Logger.Infof("Diff in MaxAge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxAge), dcl.SprintResource(actual.MaxAge))
		return true
	}
	if actual.AllowCredentials == nil && desired.AllowCredentials != nil && !dcl.IsEmptyValueIndirect(desired.AllowCredentials) {
		c.Config.Logger.Infof("desired AllowCredentials %s - but actually nil", dcl.SprintResource(desired.AllowCredentials))
		return true
	}
	if !reflect.DeepEqual(desired.AllowCredentials, actual.AllowCredentials) && !dcl.IsZeroValue(desired.AllowCredentials) && !(dcl.IsEmptyValueIndirect(desired.AllowCredentials) && dcl.IsZeroValue(actual.AllowCredentials)) {
		c.Config.Logger.Infof("Diff in AllowCredentials. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowCredentials), dcl.SprintResource(actual.AllowCredentials))
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Delay == nil && desired.Delay != nil && !dcl.IsEmptyValueIndirect(desired.Delay) {
		c.Config.Logger.Infof("desired Delay %s - but actually nil", dcl.SprintResource(desired.Delay))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, desired.Delay, actual.Delay) && !dcl.IsZeroValue(desired.Delay) {
		c.Config.Logger.Infof("Diff in Delay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Delay), dcl.SprintResource(actual.Delay))
		return true
	}
	if actual.Abort == nil && desired.Abort != nil && !dcl.IsEmptyValueIndirect(desired.Abort) {
		c.Config.Logger.Infof("desired Abort %s - but actually nil", dcl.SprintResource(desired.Abort))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, desired.Abort, actual.Abort) && !dcl.IsZeroValue(desired.Abort) {
		c.Config.Logger.Infof("Diff in Abort. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Abort), dcl.SprintResource(actual.Abort))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FixedDelay == nil && desired.FixedDelay != nil && !dcl.IsEmptyValueIndirect(desired.FixedDelay) {
		c.Config.Logger.Infof("desired FixedDelay %s - but actually nil", dcl.SprintResource(desired.FixedDelay))
		return true
	}
	if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, desired.FixedDelay, actual.FixedDelay) && !dcl.IsZeroValue(desired.FixedDelay) {
		c.Config.Logger.Infof("Diff in FixedDelay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedDelay), dcl.SprintResource(actual.FixedDelay))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c *Client, desired, actual *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HttpStatus == nil && desired.HttpStatus != nil && !dcl.IsEmptyValueIndirect(desired.HttpStatus) {
		c.Config.Logger.Infof("desired HttpStatus %s - but actually nil", dcl.SprintResource(desired.HttpStatus))
		return true
	}
	if !reflect.DeepEqual(desired.HttpStatus, actual.HttpStatus) && !dcl.IsZeroValue(desired.HttpStatus) && !(dcl.IsEmptyValueIndirect(desired.HttpStatus) && dcl.IsZeroValue(actual.HttpStatus)) {
		c.Config.Logger.Infof("Diff in HttpStatus. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpStatus), dcl.SprintResource(actual.HttpStatus))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapPathMatcherPathRuleUrlRedirectSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleUrlRedirect) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleUrlRedirect, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleUrlRedirect(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleUrlRedirect, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleUrlRedirect(c *Client, desired, actual *UrlMapPathMatcherPathRuleUrlRedirect) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HostRedirect == nil && desired.HostRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HostRedirect) {
		c.Config.Logger.Infof("desired HostRedirect %s - but actually nil", dcl.SprintResource(desired.HostRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HostRedirect, actual.HostRedirect) && !dcl.IsZeroValue(desired.HostRedirect) && !(dcl.IsEmptyValueIndirect(desired.HostRedirect) && dcl.IsZeroValue(actual.HostRedirect)) {
		c.Config.Logger.Infof("Diff in HostRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRedirect), dcl.SprintResource(actual.HostRedirect))
		return true
	}
	if actual.PathRedirect == nil && desired.PathRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PathRedirect) {
		c.Config.Logger.Infof("desired PathRedirect %s - but actually nil", dcl.SprintResource(desired.PathRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PathRedirect, actual.PathRedirect) && !dcl.IsZeroValue(desired.PathRedirect) && !(dcl.IsEmptyValueIndirect(desired.PathRedirect) && dcl.IsZeroValue(actual.PathRedirect)) {
		c.Config.Logger.Infof("Diff in PathRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathRedirect), dcl.SprintResource(actual.PathRedirect))
		return true
	}
	if actual.PrefixRedirect == nil && desired.PrefixRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PrefixRedirect) {
		c.Config.Logger.Infof("desired PrefixRedirect %s - but actually nil", dcl.SprintResource(desired.PrefixRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixRedirect, actual.PrefixRedirect) && !dcl.IsZeroValue(desired.PrefixRedirect) && !(dcl.IsEmptyValueIndirect(desired.PrefixRedirect) && dcl.IsZeroValue(actual.PrefixRedirect)) {
		c.Config.Logger.Infof("Diff in PrefixRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixRedirect), dcl.SprintResource(actual.PrefixRedirect))
		return true
	}
	if actual.RedirectResponseCode == nil && desired.RedirectResponseCode != nil && !dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) {
		c.Config.Logger.Infof("desired RedirectResponseCode %s - but actually nil", dcl.SprintResource(desired.RedirectResponseCode))
		return true
	}
	if !reflect.DeepEqual(desired.RedirectResponseCode, actual.RedirectResponseCode) && !dcl.IsZeroValue(desired.RedirectResponseCode) && !(dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) && dcl.IsZeroValue(actual.RedirectResponseCode)) {
		c.Config.Logger.Infof("Diff in RedirectResponseCode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RedirectResponseCode), dcl.SprintResource(actual.RedirectResponseCode))
		return true
	}
	if actual.HttpsRedirect == nil && desired.HttpsRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HttpsRedirect) {
		c.Config.Logger.Infof("desired HttpsRedirect %s - but actually nil", dcl.SprintResource(desired.HttpsRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HttpsRedirect, actual.HttpsRedirect) && !dcl.IsZeroValue(desired.HttpsRedirect) && !(dcl.IsEmptyValueIndirect(desired.HttpsRedirect) && dcl.IsZeroValue(actual.HttpsRedirect)) {
		c.Config.Logger.Infof("Diff in HttpsRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpsRedirect), dcl.SprintResource(actual.HttpsRedirect))
		return true
	}
	if actual.StripQuery == nil && desired.StripQuery != nil && !dcl.IsEmptyValueIndirect(desired.StripQuery) {
		c.Config.Logger.Infof("desired StripQuery %s - but actually nil", dcl.SprintResource(desired.StripQuery))
		return true
	}
	if !reflect.DeepEqual(desired.StripQuery, actual.StripQuery) && !dcl.IsZeroValue(desired.StripQuery) && !(dcl.IsEmptyValueIndirect(desired.StripQuery) && dcl.IsZeroValue(actual.StripQuery)) {
		c.Config.Logger.Infof("Diff in StripQuery. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StripQuery), dcl.SprintResource(actual.StripQuery))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRule, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRule(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRule, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRule(c *Client, desired, actual *UrlMapPathMatcherRouteRule) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Priority == nil && desired.Priority != nil && !dcl.IsEmptyValueIndirect(desired.Priority) {
		c.Config.Logger.Infof("desired Priority %s - but actually nil", dcl.SprintResource(desired.Priority))
		return true
	}
	if !reflect.DeepEqual(desired.Priority, actual.Priority) && !dcl.IsZeroValue(desired.Priority) && !(dcl.IsEmptyValueIndirect(desired.Priority) && dcl.IsZeroValue(actual.Priority)) {
		c.Config.Logger.Infof("Diff in Priority. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Priority), dcl.SprintResource(actual.Priority))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.MatchRule == nil && desired.MatchRule != nil && !dcl.IsEmptyValueIndirect(desired.MatchRule) {
		c.Config.Logger.Infof("desired MatchRule %s - but actually nil", dcl.SprintResource(desired.MatchRule))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleSlice(c, desired.MatchRule, actual.MatchRule) && !dcl.IsZeroValue(desired.MatchRule) {
		c.Config.Logger.Infof("Diff in MatchRule. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MatchRule), dcl.SprintResource(actual.MatchRule))
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	if actual.RouteAction == nil && desired.RouteAction != nil && !dcl.IsEmptyValueIndirect(desired.RouteAction) {
		c.Config.Logger.Infof("desired RouteAction %s - but actually nil", dcl.SprintResource(desired.RouteAction))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteAction(c, desired.RouteAction, actual.RouteAction) && !dcl.IsZeroValue(desired.RouteAction) {
		c.Config.Logger.Infof("Diff in RouteAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RouteAction), dcl.SprintResource(actual.RouteAction))
		return true
	}
	if actual.UrlRedirect == nil && desired.UrlRedirect != nil && !dcl.IsEmptyValueIndirect(desired.UrlRedirect) {
		c.Config.Logger.Infof("desired UrlRedirect %s - but actually nil", dcl.SprintResource(desired.UrlRedirect))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleUrlRedirect(c, desired.UrlRedirect, actual.UrlRedirect) && !dcl.IsZeroValue(desired.UrlRedirect) {
		c.Config.Logger.Infof("Diff in UrlRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlRedirect), dcl.SprintResource(actual.UrlRedirect))
		return true
	}
	if actual.HeaderAction == nil && desired.HeaderAction != nil && !dcl.IsEmptyValueIndirect(desired.HeaderAction) {
		c.Config.Logger.Infof("desired HeaderAction %s - but actually nil", dcl.SprintResource(desired.HeaderAction))
		return true
	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) && !dcl.IsZeroValue(desired.HeaderAction) {
		c.Config.Logger.Infof("Diff in HeaderAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderAction), dcl.SprintResource(actual.HeaderAction))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRule, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRule(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRule, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRule(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRule) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PrefixMatch == nil && desired.PrefixMatch != nil && !dcl.IsEmptyValueIndirect(desired.PrefixMatch) {
		c.Config.Logger.Infof("desired PrefixMatch %s - but actually nil", dcl.SprintResource(desired.PrefixMatch))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixMatch, actual.PrefixMatch) && !dcl.IsZeroValue(desired.PrefixMatch) && !(dcl.IsEmptyValueIndirect(desired.PrefixMatch) && dcl.IsZeroValue(actual.PrefixMatch)) {
		c.Config.Logger.Infof("Diff in PrefixMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixMatch), dcl.SprintResource(actual.PrefixMatch))
		return true
	}
	if actual.FullPathMatch == nil && desired.FullPathMatch != nil && !dcl.IsEmptyValueIndirect(desired.FullPathMatch) {
		c.Config.Logger.Infof("desired FullPathMatch %s - but actually nil", dcl.SprintResource(desired.FullPathMatch))
		return true
	}
	if !reflect.DeepEqual(desired.FullPathMatch, actual.FullPathMatch) && !dcl.IsZeroValue(desired.FullPathMatch) && !(dcl.IsEmptyValueIndirect(desired.FullPathMatch) && dcl.IsZeroValue(actual.FullPathMatch)) {
		c.Config.Logger.Infof("Diff in FullPathMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FullPathMatch), dcl.SprintResource(actual.FullPathMatch))
		return true
	}
	if actual.RegexMatch == nil && desired.RegexMatch != nil && !dcl.IsEmptyValueIndirect(desired.RegexMatch) {
		c.Config.Logger.Infof("desired RegexMatch %s - but actually nil", dcl.SprintResource(desired.RegexMatch))
		return true
	}
	if !reflect.DeepEqual(desired.RegexMatch, actual.RegexMatch) && !dcl.IsZeroValue(desired.RegexMatch) && !(dcl.IsEmptyValueIndirect(desired.RegexMatch) && dcl.IsZeroValue(actual.RegexMatch)) {
		c.Config.Logger.Infof("Diff in RegexMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RegexMatch), dcl.SprintResource(actual.RegexMatch))
		return true
	}
	if actual.IgnoreCase == nil && desired.IgnoreCase != nil && !dcl.IsEmptyValueIndirect(desired.IgnoreCase) {
		c.Config.Logger.Infof("desired IgnoreCase %s - but actually nil", dcl.SprintResource(desired.IgnoreCase))
		return true
	}
	if !reflect.DeepEqual(desired.IgnoreCase, actual.IgnoreCase) && !dcl.IsZeroValue(desired.IgnoreCase) && !(dcl.IsEmptyValueIndirect(desired.IgnoreCase) && dcl.IsZeroValue(actual.IgnoreCase)) {
		c.Config.Logger.Infof("Diff in IgnoreCase. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IgnoreCase), dcl.SprintResource(actual.IgnoreCase))
		return true
	}
	if actual.HeaderMatch == nil && desired.HeaderMatch != nil && !dcl.IsEmptyValueIndirect(desired.HeaderMatch) {
		c.Config.Logger.Infof("desired HeaderMatch %s - but actually nil", dcl.SprintResource(desired.HeaderMatch))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c, desired.HeaderMatch, actual.HeaderMatch) && !dcl.IsZeroValue(desired.HeaderMatch) {
		c.Config.Logger.Infof("Diff in HeaderMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderMatch), dcl.SprintResource(actual.HeaderMatch))
		return true
	}
	if actual.QueryParameterMatch == nil && desired.QueryParameterMatch != nil && !dcl.IsEmptyValueIndirect(desired.QueryParameterMatch) {
		c.Config.Logger.Infof("desired QueryParameterMatch %s - but actually nil", dcl.SprintResource(desired.QueryParameterMatch))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c, desired.QueryParameterMatch, actual.QueryParameterMatch) && !dcl.IsZeroValue(desired.QueryParameterMatch) {
		c.Config.Logger.Infof("Diff in QueryParameterMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryParameterMatch), dcl.SprintResource(actual.QueryParameterMatch))
		return true
	}
	if actual.MetadataFilter == nil && desired.MetadataFilter != nil && !dcl.IsEmptyValueIndirect(desired.MetadataFilter) {
		c.Config.Logger.Infof("desired MetadataFilter %s - but actually nil", dcl.SprintResource(desired.MetadataFilter))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c, desired.MetadataFilter, actual.MetadataFilter) && !dcl.IsZeroValue(desired.MetadataFilter) {
		c.Config.Logger.Infof("Diff in MetadataFilter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MetadataFilter), dcl.SprintResource(actual.MetadataFilter))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HeaderName == nil && desired.HeaderName != nil && !dcl.IsEmptyValueIndirect(desired.HeaderName) {
		c.Config.Logger.Infof("desired HeaderName %s - but actually nil", dcl.SprintResource(desired.HeaderName))
		return true
	}
	if !reflect.DeepEqual(desired.HeaderName, actual.HeaderName) && !dcl.IsZeroValue(desired.HeaderName) && !(dcl.IsEmptyValueIndirect(desired.HeaderName) && dcl.IsZeroValue(actual.HeaderName)) {
		c.Config.Logger.Infof("Diff in HeaderName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderName), dcl.SprintResource(actual.HeaderName))
		return true
	}
	if actual.ExactMatch == nil && desired.ExactMatch != nil && !dcl.IsEmptyValueIndirect(desired.ExactMatch) {
		c.Config.Logger.Infof("desired ExactMatch %s - but actually nil", dcl.SprintResource(desired.ExactMatch))
		return true
	}
	if !reflect.DeepEqual(desired.ExactMatch, actual.ExactMatch) && !dcl.IsZeroValue(desired.ExactMatch) && !(dcl.IsEmptyValueIndirect(desired.ExactMatch) && dcl.IsZeroValue(actual.ExactMatch)) {
		c.Config.Logger.Infof("Diff in ExactMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExactMatch), dcl.SprintResource(actual.ExactMatch))
		return true
	}
	if actual.RegexMatch == nil && desired.RegexMatch != nil && !dcl.IsEmptyValueIndirect(desired.RegexMatch) {
		c.Config.Logger.Infof("desired RegexMatch %s - but actually nil", dcl.SprintResource(desired.RegexMatch))
		return true
	}
	if !reflect.DeepEqual(desired.RegexMatch, actual.RegexMatch) && !dcl.IsZeroValue(desired.RegexMatch) && !(dcl.IsEmptyValueIndirect(desired.RegexMatch) && dcl.IsZeroValue(actual.RegexMatch)) {
		c.Config.Logger.Infof("Diff in RegexMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RegexMatch), dcl.SprintResource(actual.RegexMatch))
		return true
	}
	if actual.RangeMatch == nil && desired.RangeMatch != nil && !dcl.IsEmptyValueIndirect(desired.RangeMatch) {
		c.Config.Logger.Infof("desired RangeMatch %s - but actually nil", dcl.SprintResource(desired.RangeMatch))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, desired.RangeMatch, actual.RangeMatch) && !dcl.IsZeroValue(desired.RangeMatch) {
		c.Config.Logger.Infof("Diff in RangeMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RangeMatch), dcl.SprintResource(actual.RangeMatch))
		return true
	}
	if actual.PresentMatch == nil && desired.PresentMatch != nil && !dcl.IsEmptyValueIndirect(desired.PresentMatch) {
		c.Config.Logger.Infof("desired PresentMatch %s - but actually nil", dcl.SprintResource(desired.PresentMatch))
		return true
	}
	if !reflect.DeepEqual(desired.PresentMatch, actual.PresentMatch) && !dcl.IsZeroValue(desired.PresentMatch) && !(dcl.IsEmptyValueIndirect(desired.PresentMatch) && dcl.IsZeroValue(actual.PresentMatch)) {
		c.Config.Logger.Infof("Diff in PresentMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PresentMatch), dcl.SprintResource(actual.PresentMatch))
		return true
	}
	if actual.PrefixMatch == nil && desired.PrefixMatch != nil && !dcl.IsEmptyValueIndirect(desired.PrefixMatch) {
		c.Config.Logger.Infof("desired PrefixMatch %s - but actually nil", dcl.SprintResource(desired.PrefixMatch))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixMatch, actual.PrefixMatch) && !dcl.IsZeroValue(desired.PrefixMatch) && !(dcl.IsEmptyValueIndirect(desired.PrefixMatch) && dcl.IsZeroValue(actual.PrefixMatch)) {
		c.Config.Logger.Infof("Diff in PrefixMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixMatch), dcl.SprintResource(actual.PrefixMatch))
		return true
	}
	if actual.SuffixMatch == nil && desired.SuffixMatch != nil && !dcl.IsEmptyValueIndirect(desired.SuffixMatch) {
		c.Config.Logger.Infof("desired SuffixMatch %s - but actually nil", dcl.SprintResource(desired.SuffixMatch))
		return true
	}
	if !reflect.DeepEqual(desired.SuffixMatch, actual.SuffixMatch) && !dcl.IsZeroValue(desired.SuffixMatch) && !(dcl.IsEmptyValueIndirect(desired.SuffixMatch) && dcl.IsZeroValue(actual.SuffixMatch)) {
		c.Config.Logger.Infof("Diff in SuffixMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SuffixMatch), dcl.SprintResource(actual.SuffixMatch))
		return true
	}
	if actual.InvertMatch == nil && desired.InvertMatch != nil && !dcl.IsEmptyValueIndirect(desired.InvertMatch) {
		c.Config.Logger.Infof("desired InvertMatch %s - but actually nil", dcl.SprintResource(desired.InvertMatch))
		return true
	}
	if !reflect.DeepEqual(desired.InvertMatch, actual.InvertMatch) && !dcl.IsZeroValue(desired.InvertMatch) && !(dcl.IsEmptyValueIndirect(desired.InvertMatch) && dcl.IsZeroValue(actual.InvertMatch)) {
		c.Config.Logger.Infof("Diff in InvertMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InvertMatch), dcl.SprintResource(actual.InvertMatch))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RangeStart == nil && desired.RangeStart != nil && !dcl.IsEmptyValueIndirect(desired.RangeStart) {
		c.Config.Logger.Infof("desired RangeStart %s - but actually nil", dcl.SprintResource(desired.RangeStart))
		return true
	}
	if !reflect.DeepEqual(desired.RangeStart, actual.RangeStart) && !dcl.IsZeroValue(desired.RangeStart) && !(dcl.IsEmptyValueIndirect(desired.RangeStart) && dcl.IsZeroValue(actual.RangeStart)) {
		c.Config.Logger.Infof("Diff in RangeStart. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RangeStart), dcl.SprintResource(actual.RangeStart))
		return true
	}
	if actual.RangeEnd == nil && desired.RangeEnd != nil && !dcl.IsEmptyValueIndirect(desired.RangeEnd) {
		c.Config.Logger.Infof("desired RangeEnd %s - but actually nil", dcl.SprintResource(desired.RangeEnd))
		return true
	}
	if !reflect.DeepEqual(desired.RangeEnd, actual.RangeEnd) && !dcl.IsZeroValue(desired.RangeEnd) && !(dcl.IsEmptyValueIndirect(desired.RangeEnd) && dcl.IsZeroValue(actual.RangeEnd)) {
		c.Config.Logger.Infof("Diff in RangeEnd. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RangeEnd), dcl.SprintResource(actual.RangeEnd))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) bool {
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
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.PresentMatch == nil && desired.PresentMatch != nil && !dcl.IsEmptyValueIndirect(desired.PresentMatch) {
		c.Config.Logger.Infof("desired PresentMatch %s - but actually nil", dcl.SprintResource(desired.PresentMatch))
		return true
	}
	if !reflect.DeepEqual(desired.PresentMatch, actual.PresentMatch) && !dcl.IsZeroValue(desired.PresentMatch) && !(dcl.IsEmptyValueIndirect(desired.PresentMatch) && dcl.IsZeroValue(actual.PresentMatch)) {
		c.Config.Logger.Infof("Diff in PresentMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PresentMatch), dcl.SprintResource(actual.PresentMatch))
		return true
	}
	if actual.ExactMatch == nil && desired.ExactMatch != nil && !dcl.IsEmptyValueIndirect(desired.ExactMatch) {
		c.Config.Logger.Infof("desired ExactMatch %s - but actually nil", dcl.SprintResource(desired.ExactMatch))
		return true
	}
	if !reflect.DeepEqual(desired.ExactMatch, actual.ExactMatch) && !dcl.IsZeroValue(desired.ExactMatch) && !(dcl.IsEmptyValueIndirect(desired.ExactMatch) && dcl.IsZeroValue(actual.ExactMatch)) {
		c.Config.Logger.Infof("Diff in ExactMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExactMatch), dcl.SprintResource(actual.ExactMatch))
		return true
	}
	if actual.RegexMatch == nil && desired.RegexMatch != nil && !dcl.IsEmptyValueIndirect(desired.RegexMatch) {
		c.Config.Logger.Infof("desired RegexMatch %s - but actually nil", dcl.SprintResource(desired.RegexMatch))
		return true
	}
	if !reflect.DeepEqual(desired.RegexMatch, actual.RegexMatch) && !dcl.IsZeroValue(desired.RegexMatch) && !(dcl.IsEmptyValueIndirect(desired.RegexMatch) && dcl.IsZeroValue(actual.RegexMatch)) {
		c.Config.Logger.Infof("Diff in RegexMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RegexMatch), dcl.SprintResource(actual.RegexMatch))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FilterMatchCriteria == nil && desired.FilterMatchCriteria != nil && !dcl.IsEmptyValueIndirect(desired.FilterMatchCriteria) {
		c.Config.Logger.Infof("desired FilterMatchCriteria %s - but actually nil", dcl.SprintResource(desired.FilterMatchCriteria))
		return true
	}
	if !reflect.DeepEqual(desired.FilterMatchCriteria, actual.FilterMatchCriteria) && !dcl.IsZeroValue(desired.FilterMatchCriteria) && !(dcl.IsEmptyValueIndirect(desired.FilterMatchCriteria) && dcl.IsZeroValue(actual.FilterMatchCriteria)) {
		c.Config.Logger.Infof("Diff in FilterMatchCriteria. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FilterMatchCriteria), dcl.SprintResource(actual.FilterMatchCriteria))
		return true
	}
	if actual.FilterLabel == nil && desired.FilterLabel != nil && !dcl.IsEmptyValueIndirect(desired.FilterLabel) {
		c.Config.Logger.Infof("desired FilterLabel %s - but actually nil", dcl.SprintResource(desired.FilterLabel))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c, desired.FilterLabel, actual.FilterLabel) && !dcl.IsZeroValue(desired.FilterLabel) {
		c.Config.Logger.Infof("Diff in FilterLabel. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FilterLabel), dcl.SprintResource(actual.FilterLabel))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) bool {
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
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteAction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteAction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteAction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteAction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteAction(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteAction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.WeightedBackendService == nil && desired.WeightedBackendService != nil && !dcl.IsEmptyValueIndirect(desired.WeightedBackendService) {
		c.Config.Logger.Infof("desired WeightedBackendService %s - but actually nil", dcl.SprintResource(desired.WeightedBackendService))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c, desired.WeightedBackendService, actual.WeightedBackendService) && !dcl.IsZeroValue(desired.WeightedBackendService) {
		c.Config.Logger.Infof("Diff in WeightedBackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WeightedBackendService), dcl.SprintResource(actual.WeightedBackendService))
		return true
	}
	if actual.UrlRewrite == nil && desired.UrlRewrite != nil && !dcl.IsEmptyValueIndirect(desired.UrlRewrite) {
		c.Config.Logger.Infof("desired UrlRewrite %s - but actually nil", dcl.SprintResource(desired.UrlRewrite))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, desired.UrlRewrite, actual.UrlRewrite) && !dcl.IsZeroValue(desired.UrlRewrite) {
		c.Config.Logger.Infof("Diff in UrlRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlRewrite), dcl.SprintResource(actual.UrlRewrite))
		return true
	}
	if actual.Timeout == nil && desired.Timeout != nil && !dcl.IsEmptyValueIndirect(desired.Timeout) {
		c.Config.Logger.Infof("desired Timeout %s - but actually nil", dcl.SprintResource(desired.Timeout))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionTimeout(c, desired.Timeout, actual.Timeout) && !dcl.IsZeroValue(desired.Timeout) {
		c.Config.Logger.Infof("Diff in Timeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Timeout), dcl.SprintResource(actual.Timeout))
		return true
	}
	if actual.RetryPolicy == nil && desired.RetryPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RetryPolicy) {
		c.Config.Logger.Infof("desired RetryPolicy %s - but actually nil", dcl.SprintResource(desired.RetryPolicy))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, desired.RetryPolicy, actual.RetryPolicy) && !dcl.IsZeroValue(desired.RetryPolicy) {
		c.Config.Logger.Infof("Diff in RetryPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryPolicy), dcl.SprintResource(actual.RetryPolicy))
		return true
	}
	if actual.RequestMirrorPolicy == nil && desired.RequestMirrorPolicy != nil && !dcl.IsEmptyValueIndirect(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("desired RequestMirrorPolicy %s - but actually nil", dcl.SprintResource(desired.RequestMirrorPolicy))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, desired.RequestMirrorPolicy, actual.RequestMirrorPolicy) && !dcl.IsZeroValue(desired.RequestMirrorPolicy) {
		c.Config.Logger.Infof("Diff in RequestMirrorPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestMirrorPolicy), dcl.SprintResource(actual.RequestMirrorPolicy))
		return true
	}
	if actual.CorsPolicy == nil && desired.CorsPolicy != nil && !dcl.IsEmptyValueIndirect(desired.CorsPolicy) {
		c.Config.Logger.Infof("desired CorsPolicy %s - but actually nil", dcl.SprintResource(desired.CorsPolicy))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, desired.CorsPolicy, actual.CorsPolicy) && !dcl.IsZeroValue(desired.CorsPolicy) {
		c.Config.Logger.Infof("Diff in CorsPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CorsPolicy), dcl.SprintResource(actual.CorsPolicy))
		return true
	}
	if actual.FaultInjectionPolicy == nil && desired.FaultInjectionPolicy != nil && !dcl.IsEmptyValueIndirect(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("desired FaultInjectionPolicy %s - but actually nil", dcl.SprintResource(desired.FaultInjectionPolicy))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, desired.FaultInjectionPolicy, actual.FaultInjectionPolicy) && !dcl.IsZeroValue(desired.FaultInjectionPolicy) {
		c.Config.Logger.Infof("Diff in FaultInjectionPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FaultInjectionPolicy), dcl.SprintResource(actual.FaultInjectionPolicy))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	if actual.Weight == nil && desired.Weight != nil && !dcl.IsEmptyValueIndirect(desired.Weight) {
		c.Config.Logger.Infof("desired Weight %s - but actually nil", dcl.SprintResource(desired.Weight))
		return true
	}
	if !reflect.DeepEqual(desired.Weight, actual.Weight) && !dcl.IsZeroValue(desired.Weight) && !(dcl.IsEmptyValueIndirect(desired.Weight) && dcl.IsZeroValue(actual.Weight)) {
		c.Config.Logger.Infof("Diff in Weight. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Weight), dcl.SprintResource(actual.Weight))
		return true
	}
	if actual.HeaderAction == nil && desired.HeaderAction != nil && !dcl.IsEmptyValueIndirect(desired.HeaderAction) {
		c.Config.Logger.Infof("desired HeaderAction %s - but actually nil", dcl.SprintResource(desired.HeaderAction))
		return true
	}
	if compareUrlMapHeaderAction(c, desired.HeaderAction, actual.HeaderAction) && !dcl.IsZeroValue(desired.HeaderAction) {
		c.Config.Logger.Infof("Diff in HeaderAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderAction), dcl.SprintResource(actual.HeaderAction))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionUrlRewrite, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionUrlRewrite, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PathPrefixRewrite == nil && desired.PathPrefixRewrite != nil && !dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) {
		c.Config.Logger.Infof("desired PathPrefixRewrite %s - but actually nil", dcl.SprintResource(desired.PathPrefixRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.PathPrefixRewrite, actual.PathPrefixRewrite) && !dcl.IsZeroValue(desired.PathPrefixRewrite) && !(dcl.IsEmptyValueIndirect(desired.PathPrefixRewrite) && dcl.IsZeroValue(actual.PathPrefixRewrite)) {
		c.Config.Logger.Infof("Diff in PathPrefixRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathPrefixRewrite), dcl.SprintResource(actual.PathPrefixRewrite))
		return true
	}
	if actual.HostRewrite == nil && desired.HostRewrite != nil && !dcl.IsEmptyValueIndirect(desired.HostRewrite) {
		c.Config.Logger.Infof("desired HostRewrite %s - but actually nil", dcl.SprintResource(desired.HostRewrite))
		return true
	}
	if !reflect.DeepEqual(desired.HostRewrite, actual.HostRewrite) && !dcl.IsZeroValue(desired.HostRewrite) && !(dcl.IsEmptyValueIndirect(desired.HostRewrite) && dcl.IsZeroValue(actual.HostRewrite)) {
		c.Config.Logger.Infof("Diff in HostRewrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRewrite), dcl.SprintResource(actual.HostRewrite))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionTimeout(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionRetryPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionRetryPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RetryCondition == nil && desired.RetryCondition != nil && !dcl.IsEmptyValueIndirect(desired.RetryCondition) {
		c.Config.Logger.Infof("desired RetryCondition %s - but actually nil", dcl.SprintResource(desired.RetryCondition))
		return true
	}
	if !dcl.SliceEquals(desired.RetryCondition, actual.RetryCondition) && !dcl.IsZeroValue(desired.RetryCondition) {
		c.Config.Logger.Infof("Diff in RetryCondition. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryCondition), dcl.SprintResource(actual.RetryCondition))
		return true
	}
	if actual.NumRetries == nil && desired.NumRetries != nil && !dcl.IsEmptyValueIndirect(desired.NumRetries) {
		c.Config.Logger.Infof("desired NumRetries %s - but actually nil", dcl.SprintResource(desired.NumRetries))
		return true
	}
	if !reflect.DeepEqual(desired.NumRetries, actual.NumRetries) && !dcl.IsZeroValue(desired.NumRetries) && !(dcl.IsEmptyValueIndirect(desired.NumRetries) && dcl.IsZeroValue(actual.NumRetries)) {
		c.Config.Logger.Infof("Diff in NumRetries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumRetries), dcl.SprintResource(actual.NumRetries))
		return true
	}
	if actual.PerTryTimeout == nil && desired.PerTryTimeout != nil && !dcl.IsEmptyValueIndirect(desired.PerTryTimeout) {
		c.Config.Logger.Infof("desired PerTryTimeout %s - but actually nil", dcl.SprintResource(desired.PerTryTimeout))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, desired.PerTryTimeout, actual.PerTryTimeout) && !dcl.IsZeroValue(desired.PerTryTimeout) {
		c.Config.Logger.Infof("Diff in PerTryTimeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerTryTimeout), dcl.SprintResource(actual.PerTryTimeout))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BackendService == nil && desired.BackendService != nil && !dcl.IsEmptyValueIndirect(desired.BackendService) {
		c.Config.Logger.Infof("desired BackendService %s - but actually nil", dcl.SprintResource(desired.BackendService))
		return true
	}
	if !reflect.DeepEqual(desired.BackendService, actual.BackendService) && !dcl.IsZeroValue(desired.BackendService) && !(dcl.IsEmptyValueIndirect(desired.BackendService) && dcl.IsZeroValue(actual.BackendService)) {
		c.Config.Logger.Infof("Diff in BackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackendService), dcl.SprintResource(actual.BackendService))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionCorsPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionCorsPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AllowOrigin == nil && desired.AllowOrigin != nil && !dcl.IsEmptyValueIndirect(desired.AllowOrigin) {
		c.Config.Logger.Infof("desired AllowOrigin %s - but actually nil", dcl.SprintResource(desired.AllowOrigin))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOrigin, actual.AllowOrigin) && !dcl.IsZeroValue(desired.AllowOrigin) {
		c.Config.Logger.Infof("Diff in AllowOrigin. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOrigin), dcl.SprintResource(actual.AllowOrigin))
		return true
	}
	if actual.AllowOriginRegex == nil && desired.AllowOriginRegex != nil && !dcl.IsEmptyValueIndirect(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("desired AllowOriginRegex %s - but actually nil", dcl.SprintResource(desired.AllowOriginRegex))
		return true
	}
	if !dcl.SliceEquals(desired.AllowOriginRegex, actual.AllowOriginRegex) && !dcl.IsZeroValue(desired.AllowOriginRegex) {
		c.Config.Logger.Infof("Diff in AllowOriginRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowOriginRegex), dcl.SprintResource(actual.AllowOriginRegex))
		return true
	}
	if actual.AllowMethod == nil && desired.AllowMethod != nil && !dcl.IsEmptyValueIndirect(desired.AllowMethod) {
		c.Config.Logger.Infof("desired AllowMethod %s - but actually nil", dcl.SprintResource(desired.AllowMethod))
		return true
	}
	if !dcl.SliceEquals(desired.AllowMethod, actual.AllowMethod) && !dcl.IsZeroValue(desired.AllowMethod) {
		c.Config.Logger.Infof("Diff in AllowMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowMethod), dcl.SprintResource(actual.AllowMethod))
		return true
	}
	if actual.AllowHeader == nil && desired.AllowHeader != nil && !dcl.IsEmptyValueIndirect(desired.AllowHeader) {
		c.Config.Logger.Infof("desired AllowHeader %s - but actually nil", dcl.SprintResource(desired.AllowHeader))
		return true
	}
	if !dcl.SliceEquals(desired.AllowHeader, actual.AllowHeader) && !dcl.IsZeroValue(desired.AllowHeader) {
		c.Config.Logger.Infof("Diff in AllowHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowHeader), dcl.SprintResource(actual.AllowHeader))
		return true
	}
	if actual.ExposeHeader == nil && desired.ExposeHeader != nil && !dcl.IsEmptyValueIndirect(desired.ExposeHeader) {
		c.Config.Logger.Infof("desired ExposeHeader %s - but actually nil", dcl.SprintResource(desired.ExposeHeader))
		return true
	}
	if !dcl.SliceEquals(desired.ExposeHeader, actual.ExposeHeader) && !dcl.IsZeroValue(desired.ExposeHeader) {
		c.Config.Logger.Infof("Diff in ExposeHeader. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeHeader), dcl.SprintResource(actual.ExposeHeader))
		return true
	}
	if actual.MaxAge == nil && desired.MaxAge != nil && !dcl.IsEmptyValueIndirect(desired.MaxAge) {
		c.Config.Logger.Infof("desired MaxAge %s - but actually nil", dcl.SprintResource(desired.MaxAge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxAge, actual.MaxAge) && !dcl.IsZeroValue(desired.MaxAge) && !(dcl.IsEmptyValueIndirect(desired.MaxAge) && dcl.IsZeroValue(actual.MaxAge)) {
		c.Config.Logger.Infof("Diff in MaxAge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxAge), dcl.SprintResource(actual.MaxAge))
		return true
	}
	if actual.AllowCredentials == nil && desired.AllowCredentials != nil && !dcl.IsEmptyValueIndirect(desired.AllowCredentials) {
		c.Config.Logger.Infof("desired AllowCredentials %s - but actually nil", dcl.SprintResource(desired.AllowCredentials))
		return true
	}
	if !reflect.DeepEqual(desired.AllowCredentials, actual.AllowCredentials) && !dcl.IsZeroValue(desired.AllowCredentials) && !(dcl.IsEmptyValueIndirect(desired.AllowCredentials) && dcl.IsZeroValue(actual.AllowCredentials)) {
		c.Config.Logger.Infof("Diff in AllowCredentials. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowCredentials), dcl.SprintResource(actual.AllowCredentials))
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Delay == nil && desired.Delay != nil && !dcl.IsEmptyValueIndirect(desired.Delay) {
		c.Config.Logger.Infof("desired Delay %s - but actually nil", dcl.SprintResource(desired.Delay))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, desired.Delay, actual.Delay) && !dcl.IsZeroValue(desired.Delay) {
		c.Config.Logger.Infof("Diff in Delay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Delay), dcl.SprintResource(actual.Delay))
		return true
	}
	if actual.Abort == nil && desired.Abort != nil && !dcl.IsEmptyValueIndirect(desired.Abort) {
		c.Config.Logger.Infof("desired Abort %s - but actually nil", dcl.SprintResource(desired.Abort))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, desired.Abort, actual.Abort) && !dcl.IsZeroValue(desired.Abort) {
		c.Config.Logger.Infof("Diff in Abort. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Abort), dcl.SprintResource(actual.Abort))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FixedDelay == nil && desired.FixedDelay != nil && !dcl.IsEmptyValueIndirect(desired.FixedDelay) {
		c.Config.Logger.Infof("desired FixedDelay %s - but actually nil", dcl.SprintResource(desired.FixedDelay))
		return true
	}
	if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, desired.FixedDelay, actual.FixedDelay) && !dcl.IsZeroValue(desired.FixedDelay) {
		c.Config.Logger.Infof("Diff in FixedDelay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedDelay), dcl.SprintResource(actual.FixedDelay))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c *Client, desired, actual *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HttpStatus == nil && desired.HttpStatus != nil && !dcl.IsEmptyValueIndirect(desired.HttpStatus) {
		c.Config.Logger.Infof("desired HttpStatus %s - but actually nil", dcl.SprintResource(desired.HttpStatus))
		return true
	}
	if !reflect.DeepEqual(desired.HttpStatus, actual.HttpStatus) && !dcl.IsZeroValue(desired.HttpStatus) && !(dcl.IsEmptyValueIndirect(desired.HttpStatus) && dcl.IsZeroValue(actual.HttpStatus)) {
		c.Config.Logger.Infof("Diff in HttpStatus. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpStatus), dcl.SprintResource(actual.HttpStatus))
		return true
	}
	if actual.Percentage == nil && desired.Percentage != nil && !dcl.IsEmptyValueIndirect(desired.Percentage) {
		c.Config.Logger.Infof("desired Percentage %s - but actually nil", dcl.SprintResource(desired.Percentage))
		return true
	}
	if !reflect.DeepEqual(desired.Percentage, actual.Percentage) && !dcl.IsZeroValue(desired.Percentage) && !(dcl.IsEmptyValueIndirect(desired.Percentage) && dcl.IsZeroValue(actual.Percentage)) {
		c.Config.Logger.Infof("Diff in Percentage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percentage), dcl.SprintResource(actual.Percentage))
		return true
	}
	return false
}
func compareUrlMapPathMatcherRouteRuleUrlRedirectSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleUrlRedirect) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleUrlRedirect, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleUrlRedirect(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleUrlRedirect, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleUrlRedirect(c *Client, desired, actual *UrlMapPathMatcherRouteRuleUrlRedirect) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HostRedirect == nil && desired.HostRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HostRedirect) {
		c.Config.Logger.Infof("desired HostRedirect %s - but actually nil", dcl.SprintResource(desired.HostRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HostRedirect, actual.HostRedirect) && !dcl.IsZeroValue(desired.HostRedirect) && !(dcl.IsEmptyValueIndirect(desired.HostRedirect) && dcl.IsZeroValue(actual.HostRedirect)) {
		c.Config.Logger.Infof("Diff in HostRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostRedirect), dcl.SprintResource(actual.HostRedirect))
		return true
	}
	if actual.PathRedirect == nil && desired.PathRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PathRedirect) {
		c.Config.Logger.Infof("desired PathRedirect %s - but actually nil", dcl.SprintResource(desired.PathRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PathRedirect, actual.PathRedirect) && !dcl.IsZeroValue(desired.PathRedirect) && !(dcl.IsEmptyValueIndirect(desired.PathRedirect) && dcl.IsZeroValue(actual.PathRedirect)) {
		c.Config.Logger.Infof("Diff in PathRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PathRedirect), dcl.SprintResource(actual.PathRedirect))
		return true
	}
	if actual.PrefixRedirect == nil && desired.PrefixRedirect != nil && !dcl.IsEmptyValueIndirect(desired.PrefixRedirect) {
		c.Config.Logger.Infof("desired PrefixRedirect %s - but actually nil", dcl.SprintResource(desired.PrefixRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.PrefixRedirect, actual.PrefixRedirect) && !dcl.IsZeroValue(desired.PrefixRedirect) && !(dcl.IsEmptyValueIndirect(desired.PrefixRedirect) && dcl.IsZeroValue(actual.PrefixRedirect)) {
		c.Config.Logger.Infof("Diff in PrefixRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrefixRedirect), dcl.SprintResource(actual.PrefixRedirect))
		return true
	}
	if actual.RedirectResponseCode == nil && desired.RedirectResponseCode != nil && !dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) {
		c.Config.Logger.Infof("desired RedirectResponseCode %s - but actually nil", dcl.SprintResource(desired.RedirectResponseCode))
		return true
	}
	if !reflect.DeepEqual(desired.RedirectResponseCode, actual.RedirectResponseCode) && !dcl.IsZeroValue(desired.RedirectResponseCode) && !(dcl.IsEmptyValueIndirect(desired.RedirectResponseCode) && dcl.IsZeroValue(actual.RedirectResponseCode)) {
		c.Config.Logger.Infof("Diff in RedirectResponseCode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RedirectResponseCode), dcl.SprintResource(actual.RedirectResponseCode))
		return true
	}
	if actual.HttpsRedirect == nil && desired.HttpsRedirect != nil && !dcl.IsEmptyValueIndirect(desired.HttpsRedirect) {
		c.Config.Logger.Infof("desired HttpsRedirect %s - but actually nil", dcl.SprintResource(desired.HttpsRedirect))
		return true
	}
	if !reflect.DeepEqual(desired.HttpsRedirect, actual.HttpsRedirect) && !dcl.IsZeroValue(desired.HttpsRedirect) && !(dcl.IsEmptyValueIndirect(desired.HttpsRedirect) && dcl.IsZeroValue(actual.HttpsRedirect)) {
		c.Config.Logger.Infof("Diff in HttpsRedirect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpsRedirect), dcl.SprintResource(actual.HttpsRedirect))
		return true
	}
	if actual.StripQuery == nil && desired.StripQuery != nil && !dcl.IsEmptyValueIndirect(desired.StripQuery) {
		c.Config.Logger.Infof("desired StripQuery %s - but actually nil", dcl.SprintResource(desired.StripQuery))
		return true
	}
	if !reflect.DeepEqual(desired.StripQuery, actual.StripQuery) && !dcl.IsZeroValue(desired.StripQuery) && !(dcl.IsEmptyValueIndirect(desired.StripQuery) && dcl.IsZeroValue(actual.StripQuery)) {
		c.Config.Logger.Infof("Diff in StripQuery. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StripQuery), dcl.SprintResource(actual.StripQuery))
		return true
	}
	return false
}
func compareUrlMapTestSlice(c *Client, desired, actual []UrlMapTest) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapTest, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapTest(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapTest, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapTest(c *Client, desired, actual *UrlMapTest) bool {
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
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.Host == nil && desired.Host != nil && !dcl.IsEmptyValueIndirect(desired.Host) {
		c.Config.Logger.Infof("desired Host %s - but actually nil", dcl.SprintResource(desired.Host))
		return true
	}
	if !reflect.DeepEqual(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) && !(dcl.IsEmptyValueIndirect(desired.Host) && dcl.IsZeroValue(actual.Host)) {
		c.Config.Logger.Infof("Diff in Host. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	if actual.Path == nil && desired.Path != nil && !dcl.IsEmptyValueIndirect(desired.Path) {
		c.Config.Logger.Infof("desired Path %s - but actually nil", dcl.SprintResource(desired.Path))
		return true
	}
	if !reflect.DeepEqual(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) && !(dcl.IsEmptyValueIndirect(desired.Path) && dcl.IsZeroValue(actual.Path)) {
		c.Config.Logger.Infof("Diff in Path. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if actual.ExpectedBackendService == nil && desired.ExpectedBackendService != nil && !dcl.IsEmptyValueIndirect(desired.ExpectedBackendService) {
		c.Config.Logger.Infof("desired ExpectedBackendService %s - but actually nil", dcl.SprintResource(desired.ExpectedBackendService))
		return true
	}
	if !reflect.DeepEqual(desired.ExpectedBackendService, actual.ExpectedBackendService) && !dcl.IsZeroValue(desired.ExpectedBackendService) && !(dcl.IsEmptyValueIndirect(desired.ExpectedBackendService) && dcl.IsZeroValue(actual.ExpectedBackendService)) {
		c.Config.Logger.Infof("Diff in ExpectedBackendService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExpectedBackendService), dcl.SprintResource(actual.ExpectedBackendService))
		return true
	}
	return false
}
func compareUrlMapDefaultUrlRedirectRedirectResponseCodeEnumSlice(c *Client, desired, actual []UrlMapDefaultUrlRedirectRedirectResponseCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapDefaultUrlRedirectRedirectResponseCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapDefaultUrlRedirectRedirectResponseCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(c *Client, desired, actual *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumSlice(c *Client, desired, actual []UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(c *Client, desired, actual *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumSlice(c *Client, desired, actual []UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(c *Client, desired, actual *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(c *Client, desired, actual *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumSlice(c *Client, desired, actual []UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(c *Client, desired, actual *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *UrlMap) urlNormalized() *UrlMap {
	normalized := deepcopy.Copy(*r).(UrlMap)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *UrlMap) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UrlMap) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *UrlMap) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UrlMap) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/urlMaps/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the UrlMap resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *UrlMap) marshal(c *Client) ([]byte, error) {
	m, err := expandUrlMap(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling UrlMap: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalUrlMap decodes JSON responses into the UrlMap resource schema.
func unmarshalUrlMap(b []byte, c *Client) (*UrlMap, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenUrlMap(c, m), nil
}

// expandUrlMap expands UrlMap into a JSON request object.
func expandUrlMap(c *Client, f *UrlMap) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := expandUrlMapDefaultRouteAction(c, f.DefaultRouteAction); err != nil {
		return nil, fmt.Errorf("error expanding DefaultRouteAction into defaultRouteAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultRouteAction"] = v
	}
	if v := f.DefaultService; !dcl.IsEmptyValueIndirect(v) {
		m["defaultService"] = v
	}
	if v, err := expandUrlMapDefaultUrlRedirect(c, f.DefaultUrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding DefaultUrlRedirect into defaultUrlRedirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultUrlRedirect"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}
	if v, err := expandUrlMapHostRuleSlice(c, f.HostRule); err != nil {
		return nil, fmt.Errorf("error expanding HostRule into hostRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hostRules"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandUrlMapPathMatcherSlice(c, f.PathMatcher); err != nil {
		return nil, fmt.Errorf("error expanding PathMatcher into pathMatchers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pathMatchers"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := expandUrlMapTestSlice(c, f.Test); err != nil {
		return nil, fmt.Errorf("error expanding Test into tests: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tests"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenUrlMap flattens UrlMap from a JSON request object into the
// UrlMap type.
func flattenUrlMap(c *Client, i interface{}) *UrlMap {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &UrlMap{}
	r.DefaultRouteAction = flattenUrlMapDefaultRouteAction(c, m["defaultRouteAction"])
	r.DefaultService = dcl.FlattenString(m["defaultService"])
	r.DefaultUrlRedirect = flattenUrlMapDefaultUrlRedirect(c, m["defaultUrlRedirect"])
	r.Description = dcl.FlattenString(m["description"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])
	r.HostRule = flattenUrlMapHostRuleSlice(c, m["hostRules"])
	r.Name = dcl.FlattenString(m["name"])
	r.PathMatcher = flattenUrlMapPathMatcherSlice(c, m["pathMatchers"])
	r.Region = dcl.FlattenString(m["region"])
	r.Test = flattenUrlMapTestSlice(c, m["tests"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandUrlMapDefaultRouteActionMap expands the contents of UrlMapDefaultRouteAction into a JSON
// request object.
func expandUrlMapDefaultRouteActionMap(c *Client, f map[string]UrlMapDefaultRouteAction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteAction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionSlice expands the contents of UrlMapDefaultRouteAction into a JSON
// request object.
func expandUrlMapDefaultRouteActionSlice(c *Client, f []UrlMapDefaultRouteAction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteAction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionMap flattens the contents of UrlMapDefaultRouteAction from a JSON
// response object.
func flattenUrlMapDefaultRouteActionMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteAction{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteAction{}
	}

	items := make(map[string]UrlMapDefaultRouteAction)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteAction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionSlice flattens the contents of UrlMapDefaultRouteAction from a JSON
// response object.
func flattenUrlMapDefaultRouteActionSlice(c *Client, i interface{}) []UrlMapDefaultRouteAction {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteAction{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteAction{}
	}

	items := make([]UrlMapDefaultRouteAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteAction(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteAction expands an instance of UrlMapDefaultRouteAction into a JSON
// request object.
func expandUrlMapDefaultRouteAction(c *Client, f *UrlMapDefaultRouteAction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapDefaultRouteActionWeightedBackendServiceSlice(c, f.WeightedBackendService); err != nil {
		return nil, fmt.Errorf("error expanding WeightedBackendService into weightedBackendServices: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["weightedBackendServices"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionUrlRewrite(c, f.UrlRewrite); err != nil {
		return nil, fmt.Errorf("error expanding UrlRewrite into urlRewrite: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRewrite"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionTimeout(c, f.Timeout); err != nil {
		return nil, fmt.Errorf("error expanding Timeout into timeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionRetryPolicy(c, f.RetryPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RetryPolicy into retryPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryPolicy"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionRequestMirrorPolicy(c, f.RequestMirrorPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RequestMirrorPolicy into requestMirrorPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestMirrorPolicy"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionCorsPolicy(c, f.CorsPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CorsPolicy into corsPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["corsPolicy"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionFaultInjectionPolicy(c, f.FaultInjectionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding FaultInjectionPolicy into faultInjectionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["faultInjectionPolicy"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteAction flattens an instance of UrlMapDefaultRouteAction from a JSON
// response object.
func flattenUrlMapDefaultRouteAction(c *Client, i interface{}) *UrlMapDefaultRouteAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteAction{}
	r.WeightedBackendService = flattenUrlMapDefaultRouteActionWeightedBackendServiceSlice(c, m["weightedBackendServices"])
	r.UrlRewrite = flattenUrlMapDefaultRouteActionUrlRewrite(c, m["urlRewrite"])
	r.Timeout = flattenUrlMapDefaultRouteActionTimeout(c, m["timeout"])
	r.RetryPolicy = flattenUrlMapDefaultRouteActionRetryPolicy(c, m["retryPolicy"])
	r.RequestMirrorPolicy = flattenUrlMapDefaultRouteActionRequestMirrorPolicy(c, m["requestMirrorPolicy"])
	r.CorsPolicy = flattenUrlMapDefaultRouteActionCorsPolicy(c, m["corsPolicy"])
	r.FaultInjectionPolicy = flattenUrlMapDefaultRouteActionFaultInjectionPolicy(c, m["faultInjectionPolicy"])

	return r
}

// expandUrlMapDefaultRouteActionWeightedBackendServiceMap expands the contents of UrlMapDefaultRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapDefaultRouteActionWeightedBackendServiceMap(c *Client, f map[string]UrlMapDefaultRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionWeightedBackendServiceSlice expands the contents of UrlMapDefaultRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapDefaultRouteActionWeightedBackendServiceSlice(c *Client, f []UrlMapDefaultRouteActionWeightedBackendService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionWeightedBackendServiceMap flattens the contents of UrlMapDefaultRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapDefaultRouteActionWeightedBackendServiceMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionWeightedBackendService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionWeightedBackendService{}
	}

	items := make(map[string]UrlMapDefaultRouteActionWeightedBackendService)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionWeightedBackendService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionWeightedBackendServiceSlice flattens the contents of UrlMapDefaultRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapDefaultRouteActionWeightedBackendServiceSlice(c *Client, i interface{}) []UrlMapDefaultRouteActionWeightedBackendService {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionWeightedBackendService{}
	}

	items := make([]UrlMapDefaultRouteActionWeightedBackendService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionWeightedBackendService(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionWeightedBackendService expands an instance of UrlMapDefaultRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapDefaultRouteActionWeightedBackendService(c *Client, f *UrlMapDefaultRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionWeightedBackendService flattens an instance of UrlMapDefaultRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapDefaultRouteActionWeightedBackendService(c *Client, i interface{}) *UrlMapDefaultRouteActionWeightedBackendService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionWeightedBackendService{}
	r.BackendService = dcl.FlattenString(m["backendService"])
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])

	return r
}

// expandUrlMapHeaderActionMap expands the contents of UrlMapHeaderAction into a JSON
// request object.
func expandUrlMapHeaderActionMap(c *Client, f map[string]UrlMapHeaderAction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapHeaderAction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapHeaderActionSlice expands the contents of UrlMapHeaderAction into a JSON
// request object.
func expandUrlMapHeaderActionSlice(c *Client, f []UrlMapHeaderAction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapHeaderAction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapHeaderActionMap flattens the contents of UrlMapHeaderAction from a JSON
// response object.
func flattenUrlMapHeaderActionMap(c *Client, i interface{}) map[string]UrlMapHeaderAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapHeaderAction{}
	}

	if len(a) == 0 {
		return map[string]UrlMapHeaderAction{}
	}

	items := make(map[string]UrlMapHeaderAction)
	for k, item := range a {
		items[k] = *flattenUrlMapHeaderAction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapHeaderActionSlice flattens the contents of UrlMapHeaderAction from a JSON
// response object.
func flattenUrlMapHeaderActionSlice(c *Client, i interface{}) []UrlMapHeaderAction {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapHeaderAction{}
	}

	if len(a) == 0 {
		return []UrlMapHeaderAction{}
	}

	items := make([]UrlMapHeaderAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapHeaderAction(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapHeaderAction expands an instance of UrlMapHeaderAction into a JSON
// request object.
func expandUrlMapHeaderAction(c *Client, f *UrlMapHeaderAction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RequestHeadersToRemove; !dcl.IsEmptyValueIndirect(v) {
		m["requestHeadersToRemove"] = v
	}
	if v, err := expandUrlMapHeaderActionRequestHeadersToAddSlice(c, f.RequestHeadersToAdd); err != nil {
		return nil, fmt.Errorf("error expanding RequestHeadersToAdd into requestHeadersToAdd: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestHeadersToAdd"] = v
	}
	if v := f.ResponseHeadersToRemove; !dcl.IsEmptyValueIndirect(v) {
		m["responseHeadersToRemove"] = v
	}
	if v, err := expandUrlMapHeaderActionResponseHeadersToAddSlice(c, f.ResponseHeadersToAdd); err != nil {
		return nil, fmt.Errorf("error expanding ResponseHeadersToAdd into responseHeadersToAdd: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["responseHeadersToAdd"] = v
	}

	return m, nil
}

// flattenUrlMapHeaderAction flattens an instance of UrlMapHeaderAction from a JSON
// response object.
func flattenUrlMapHeaderAction(c *Client, i interface{}) *UrlMapHeaderAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapHeaderAction{}
	r.RequestHeadersToRemove = dcl.FlattenStringSlice(m["requestHeadersToRemove"])
	r.RequestHeadersToAdd = flattenUrlMapHeaderActionRequestHeadersToAddSlice(c, m["requestHeadersToAdd"])
	r.ResponseHeadersToRemove = dcl.FlattenStringSlice(m["responseHeadersToRemove"])
	r.ResponseHeadersToAdd = flattenUrlMapHeaderActionResponseHeadersToAddSlice(c, m["responseHeadersToAdd"])

	return r
}

// expandUrlMapHeaderActionRequestHeadersToAddMap expands the contents of UrlMapHeaderActionRequestHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionRequestHeadersToAddMap(c *Client, f map[string]UrlMapHeaderActionRequestHeadersToAdd) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapHeaderActionRequestHeadersToAdd(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapHeaderActionRequestHeadersToAddSlice expands the contents of UrlMapHeaderActionRequestHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionRequestHeadersToAddSlice(c *Client, f []UrlMapHeaderActionRequestHeadersToAdd) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapHeaderActionRequestHeadersToAdd(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapHeaderActionRequestHeadersToAddMap flattens the contents of UrlMapHeaderActionRequestHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionRequestHeadersToAddMap(c *Client, i interface{}) map[string]UrlMapHeaderActionRequestHeadersToAdd {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapHeaderActionRequestHeadersToAdd{}
	}

	if len(a) == 0 {
		return map[string]UrlMapHeaderActionRequestHeadersToAdd{}
	}

	items := make(map[string]UrlMapHeaderActionRequestHeadersToAdd)
	for k, item := range a {
		items[k] = *flattenUrlMapHeaderActionRequestHeadersToAdd(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapHeaderActionRequestHeadersToAddSlice flattens the contents of UrlMapHeaderActionRequestHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionRequestHeadersToAddSlice(c *Client, i interface{}) []UrlMapHeaderActionRequestHeadersToAdd {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapHeaderActionRequestHeadersToAdd{}
	}

	if len(a) == 0 {
		return []UrlMapHeaderActionRequestHeadersToAdd{}
	}

	items := make([]UrlMapHeaderActionRequestHeadersToAdd, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapHeaderActionRequestHeadersToAdd(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapHeaderActionRequestHeadersToAdd expands an instance of UrlMapHeaderActionRequestHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionRequestHeadersToAdd(c *Client, f *UrlMapHeaderActionRequestHeadersToAdd) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}
	if v := f.HeaderValue; !dcl.IsEmptyValueIndirect(v) {
		m["headerValue"] = v
	}
	if v := f.Replace; !dcl.IsEmptyValueIndirect(v) {
		m["replace"] = v
	}

	return m, nil
}

// flattenUrlMapHeaderActionRequestHeadersToAdd flattens an instance of UrlMapHeaderActionRequestHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionRequestHeadersToAdd(c *Client, i interface{}) *UrlMapHeaderActionRequestHeadersToAdd {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapHeaderActionRequestHeadersToAdd{}
	r.HeaderName = dcl.FlattenString(m["headerName"])
	r.HeaderValue = dcl.FlattenString(m["headerValue"])
	r.Replace = dcl.FlattenBool(m["replace"])

	return r
}

// expandUrlMapHeaderActionResponseHeadersToAddMap expands the contents of UrlMapHeaderActionResponseHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionResponseHeadersToAddMap(c *Client, f map[string]UrlMapHeaderActionResponseHeadersToAdd) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapHeaderActionResponseHeadersToAdd(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapHeaderActionResponseHeadersToAddSlice expands the contents of UrlMapHeaderActionResponseHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionResponseHeadersToAddSlice(c *Client, f []UrlMapHeaderActionResponseHeadersToAdd) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapHeaderActionResponseHeadersToAdd(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapHeaderActionResponseHeadersToAddMap flattens the contents of UrlMapHeaderActionResponseHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionResponseHeadersToAddMap(c *Client, i interface{}) map[string]UrlMapHeaderActionResponseHeadersToAdd {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapHeaderActionResponseHeadersToAdd{}
	}

	if len(a) == 0 {
		return map[string]UrlMapHeaderActionResponseHeadersToAdd{}
	}

	items := make(map[string]UrlMapHeaderActionResponseHeadersToAdd)
	for k, item := range a {
		items[k] = *flattenUrlMapHeaderActionResponseHeadersToAdd(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapHeaderActionResponseHeadersToAddSlice flattens the contents of UrlMapHeaderActionResponseHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionResponseHeadersToAddSlice(c *Client, i interface{}) []UrlMapHeaderActionResponseHeadersToAdd {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapHeaderActionResponseHeadersToAdd{}
	}

	if len(a) == 0 {
		return []UrlMapHeaderActionResponseHeadersToAdd{}
	}

	items := make([]UrlMapHeaderActionResponseHeadersToAdd, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapHeaderActionResponseHeadersToAdd(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapHeaderActionResponseHeadersToAdd expands an instance of UrlMapHeaderActionResponseHeadersToAdd into a JSON
// request object.
func expandUrlMapHeaderActionResponseHeadersToAdd(c *Client, f *UrlMapHeaderActionResponseHeadersToAdd) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}
	if v := f.HeaderValue; !dcl.IsEmptyValueIndirect(v) {
		m["headerValue"] = v
	}
	if v := f.Replace; !dcl.IsEmptyValueIndirect(v) {
		m["replace"] = v
	}

	return m, nil
}

// flattenUrlMapHeaderActionResponseHeadersToAdd flattens an instance of UrlMapHeaderActionResponseHeadersToAdd from a JSON
// response object.
func flattenUrlMapHeaderActionResponseHeadersToAdd(c *Client, i interface{}) *UrlMapHeaderActionResponseHeadersToAdd {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapHeaderActionResponseHeadersToAdd{}
	r.HeaderName = dcl.FlattenString(m["headerName"])
	r.HeaderValue = dcl.FlattenString(m["headerValue"])
	r.Replace = dcl.FlattenBool(m["replace"])

	return r
}

// expandUrlMapDefaultRouteActionUrlRewriteMap expands the contents of UrlMapDefaultRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapDefaultRouteActionUrlRewriteMap(c *Client, f map[string]UrlMapDefaultRouteActionUrlRewrite) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionUrlRewriteSlice expands the contents of UrlMapDefaultRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapDefaultRouteActionUrlRewriteSlice(c *Client, f []UrlMapDefaultRouteActionUrlRewrite) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionUrlRewriteMap flattens the contents of UrlMapDefaultRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapDefaultRouteActionUrlRewriteMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionUrlRewrite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionUrlRewrite{}
	}

	items := make(map[string]UrlMapDefaultRouteActionUrlRewrite)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionUrlRewrite(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionUrlRewriteSlice flattens the contents of UrlMapDefaultRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapDefaultRouteActionUrlRewriteSlice(c *Client, i interface{}) []UrlMapDefaultRouteActionUrlRewrite {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionUrlRewrite{}
	}

	items := make([]UrlMapDefaultRouteActionUrlRewrite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionUrlRewrite(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionUrlRewrite expands an instance of UrlMapDefaultRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapDefaultRouteActionUrlRewrite(c *Client, f *UrlMapDefaultRouteActionUrlRewrite) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PathPrefixRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["pathPrefixRewrite"] = v
	}
	if v := f.HostRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["hostRewrite"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionUrlRewrite flattens an instance of UrlMapDefaultRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapDefaultRouteActionUrlRewrite(c *Client, i interface{}) *UrlMapDefaultRouteActionUrlRewrite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionUrlRewrite{}
	r.PathPrefixRewrite = dcl.FlattenString(m["pathPrefixRewrite"])
	r.HostRewrite = dcl.FlattenString(m["hostRewrite"])

	return r
}

// expandUrlMapDefaultRouteActionTimeoutMap expands the contents of UrlMapDefaultRouteActionTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionTimeoutMap(c *Client, f map[string]UrlMapDefaultRouteActionTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionTimeoutSlice expands the contents of UrlMapDefaultRouteActionTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionTimeoutSlice(c *Client, f []UrlMapDefaultRouteActionTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionTimeoutMap flattens the contents of UrlMapDefaultRouteActionTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionTimeoutMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionTimeout{}
	}

	items := make(map[string]UrlMapDefaultRouteActionTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionTimeoutSlice flattens the contents of UrlMapDefaultRouteActionTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionTimeoutSlice(c *Client, i interface{}) []UrlMapDefaultRouteActionTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionTimeout{}
	}

	items := make([]UrlMapDefaultRouteActionTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionTimeout expands an instance of UrlMapDefaultRouteActionTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionTimeout(c *Client, f *UrlMapDefaultRouteActionTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionTimeout flattens an instance of UrlMapDefaultRouteActionTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionTimeout(c *Client, i interface{}) *UrlMapDefaultRouteActionTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapDefaultRouteActionRetryPolicyMap expands the contents of UrlMapDefaultRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicyMap(c *Client, f map[string]UrlMapDefaultRouteActionRetryPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionRetryPolicySlice expands the contents of UrlMapDefaultRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicySlice(c *Client, f []UrlMapDefaultRouteActionRetryPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionRetryPolicyMap flattens the contents of UrlMapDefaultRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicyMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionRetryPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionRetryPolicy{}
	}

	items := make(map[string]UrlMapDefaultRouteActionRetryPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionRetryPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionRetryPolicySlice flattens the contents of UrlMapDefaultRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionRetryPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionRetryPolicy{}
	}

	items := make([]UrlMapDefaultRouteActionRetryPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionRetryPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionRetryPolicy expands an instance of UrlMapDefaultRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicy(c *Client, f *UrlMapDefaultRouteActionRetryPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryCondition; !dcl.IsEmptyValueIndirect(v) {
		m["retryConditions"] = v
	}
	if v := f.NumRetries; !dcl.IsEmptyValueIndirect(v) {
		m["numRetries"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, f.PerTryTimeout); err != nil {
		return nil, fmt.Errorf("error expanding PerTryTimeout into perTryTimeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["perTryTimeout"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionRetryPolicy flattens an instance of UrlMapDefaultRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicy(c *Client, i interface{}) *UrlMapDefaultRouteActionRetryPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionRetryPolicy{}
	r.RetryCondition = dcl.FlattenStringSlice(m["retryConditions"])
	r.NumRetries = dcl.FlattenInteger(m["numRetries"])
	r.PerTryTimeout = flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, m["perTryTimeout"])

	return r
}

// expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutMap expands the contents of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutMap(c *Client, f map[string]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice expands the contents of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, f []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutMap flattens the contents of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make(map[string]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice flattens the contents of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, i interface{}) []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make([]UrlMapDefaultRouteActionRetryPolicyPerTryTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeout expands an instance of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c *Client, f *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeout flattens an instance of UrlMapDefaultRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c *Client, i interface{}) *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapDefaultRouteActionRequestMirrorPolicyMap expands the contents of UrlMapDefaultRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRequestMirrorPolicyMap(c *Client, f map[string]UrlMapDefaultRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionRequestMirrorPolicySlice expands the contents of UrlMapDefaultRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRequestMirrorPolicySlice(c *Client, f []UrlMapDefaultRouteActionRequestMirrorPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionRequestMirrorPolicyMap flattens the contents of UrlMapDefaultRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRequestMirrorPolicyMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionRequestMirrorPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionRequestMirrorPolicy{}
	}

	items := make(map[string]UrlMapDefaultRouteActionRequestMirrorPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionRequestMirrorPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionRequestMirrorPolicySlice flattens the contents of UrlMapDefaultRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRequestMirrorPolicySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionRequestMirrorPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionRequestMirrorPolicy{}
	}

	items := make([]UrlMapDefaultRouteActionRequestMirrorPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionRequestMirrorPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionRequestMirrorPolicy expands an instance of UrlMapDefaultRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionRequestMirrorPolicy(c *Client, f *UrlMapDefaultRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionRequestMirrorPolicy flattens an instance of UrlMapDefaultRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionRequestMirrorPolicy(c *Client, i interface{}) *UrlMapDefaultRouteActionRequestMirrorPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionRequestMirrorPolicy{}
	r.BackendService = dcl.FlattenString(m["backendService"])

	return r
}

// expandUrlMapDefaultRouteActionCorsPolicyMap expands the contents of UrlMapDefaultRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionCorsPolicyMap(c *Client, f map[string]UrlMapDefaultRouteActionCorsPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionCorsPolicySlice expands the contents of UrlMapDefaultRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionCorsPolicySlice(c *Client, f []UrlMapDefaultRouteActionCorsPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionCorsPolicyMap flattens the contents of UrlMapDefaultRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionCorsPolicyMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionCorsPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionCorsPolicy{}
	}

	items := make(map[string]UrlMapDefaultRouteActionCorsPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionCorsPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionCorsPolicySlice flattens the contents of UrlMapDefaultRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionCorsPolicySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionCorsPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionCorsPolicy{}
	}

	items := make([]UrlMapDefaultRouteActionCorsPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionCorsPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionCorsPolicy expands an instance of UrlMapDefaultRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionCorsPolicy(c *Client, f *UrlMapDefaultRouteActionCorsPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowOrigin; !dcl.IsEmptyValueIndirect(v) {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; !dcl.IsEmptyValueIndirect(v) {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; !dcl.IsEmptyValueIndirect(v) {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; !dcl.IsEmptyValueIndirect(v) {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; !dcl.IsEmptyValueIndirect(v) {
		m["exposeHeaders"] = v
	}
	if v := f.MaxAge; !dcl.IsEmptyValueIndirect(v) {
		m["maxAge"] = v
	}
	if v := f.AllowCredentials; !dcl.IsEmptyValueIndirect(v) {
		m["allowCredentials"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionCorsPolicy flattens an instance of UrlMapDefaultRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionCorsPolicy(c *Client, i interface{}) *UrlMapDefaultRouteActionCorsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionCorsPolicy{}
	r.AllowOrigin = dcl.FlattenStringSlice(m["allowOrigins"])
	r.AllowOriginRegex = dcl.FlattenStringSlice(m["allowOriginRegexes"])
	r.AllowMethod = dcl.FlattenStringSlice(m["allowMethods"])
	r.AllowHeader = dcl.FlattenStringSlice(m["allowHeaders"])
	r.ExposeHeader = dcl.FlattenStringSlice(m["exposeHeaders"])
	r.MaxAge = dcl.FlattenInteger(m["maxAge"])
	r.AllowCredentials = dcl.FlattenBool(m["allowCredentials"])
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyMap expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyMap(c *Client, f map[string]UrlMapDefaultRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicySlice expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicySlice(c *Client, f []UrlMapDefaultRouteActionFaultInjectionPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyMap flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionFaultInjectionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicy{}
	}

	items := make(map[string]UrlMapDefaultRouteActionFaultInjectionPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionFaultInjectionPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicySlice flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionFaultInjectionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionFaultInjectionPolicy{}
	}

	items := make([]UrlMapDefaultRouteActionFaultInjectionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionFaultInjectionPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicy expands an instance of UrlMapDefaultRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicy(c *Client, f *UrlMapDefaultRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, f.Delay); err != nil {
		return nil, fmt.Errorf("error expanding Delay into delay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["delay"] = v
	}
	if v, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, f.Abort); err != nil {
		return nil, fmt.Errorf("error expanding Abort into abort: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["abort"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicy flattens an instance of UrlMapDefaultRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicy(c *Client, i interface{}) *UrlMapDefaultRouteActionFaultInjectionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionFaultInjectionPolicy{}
	r.Delay = flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, m["delay"])
	r.Abort = flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, m["abort"])

	return r
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayMap expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayMap(c *Client, f map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice(c *Client, f []UrlMapDefaultRouteActionFaultInjectionPolicyDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayMap flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}
	}

	items := make(map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}
	}

	items := make([]UrlMapDefaultRouteActionFaultInjectionPolicyDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelay expands an instance of UrlMapDefaultRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c *Client, f *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, f.FixedDelay); err != nil {
		return nil, fmt.Errorf("error expanding FixedDelay into fixedDelay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixedDelay"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelay flattens an instance of UrlMapDefaultRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c *Client, i interface{}) *UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}
	r.FixedDelay = flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, m["fixedDelay"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayMap expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, f map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, f []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayMap flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make(map[string]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, i interface{}) []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make([]UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay expands an instance of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, f *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay flattens an instance of UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, i interface{}) *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyAbortMap expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyAbortMap(c *Client, f map[string]UrlMapDefaultRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice expands the contents of UrlMapDefaultRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice(c *Client, f []UrlMapDefaultRouteActionFaultInjectionPolicyAbort) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbortMap flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbortMap(c *Client, i interface{}) map[string]UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}
	}

	items := make(map[string]UrlMapDefaultRouteActionFaultInjectionPolicyAbort)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice flattens the contents of UrlMapDefaultRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice(c *Client, i interface{}) []UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}
	}

	items := make([]UrlMapDefaultRouteActionFaultInjectionPolicyAbort, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultRouteActionFaultInjectionPolicyAbort expands an instance of UrlMapDefaultRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c *Client, f *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpStatus; !dcl.IsEmptyValueIndirect(v) {
		m["httpStatus"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbort flattens an instance of UrlMapDefaultRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c *Client, i interface{}) *UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}
	r.HttpStatus = dcl.FlattenInteger(m["httpStatus"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapDefaultUrlRedirectMap expands the contents of UrlMapDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapDefaultUrlRedirectMap(c *Client, f map[string]UrlMapDefaultUrlRedirect) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapDefaultUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapDefaultUrlRedirectSlice expands the contents of UrlMapDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapDefaultUrlRedirectSlice(c *Client, f []UrlMapDefaultUrlRedirect) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapDefaultUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapDefaultUrlRedirectMap flattens the contents of UrlMapDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapDefaultUrlRedirectMap(c *Client, i interface{}) map[string]UrlMapDefaultUrlRedirect {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultUrlRedirect{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultUrlRedirect{}
	}

	items := make(map[string]UrlMapDefaultUrlRedirect)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultUrlRedirect(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapDefaultUrlRedirectSlice flattens the contents of UrlMapDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapDefaultUrlRedirectSlice(c *Client, i interface{}) []UrlMapDefaultUrlRedirect {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultUrlRedirect{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultUrlRedirect{}
	}

	items := make([]UrlMapDefaultUrlRedirect, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultUrlRedirect(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapDefaultUrlRedirect expands an instance of UrlMapDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapDefaultUrlRedirect(c *Client, f *UrlMapDefaultUrlRedirect) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["hostRedirect"] = v
	}
	if v := f.PathRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["pathRedirect"] = v
	}
	if v := f.PrefixRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["prefixRedirect"] = v
	}
	if v := f.RedirectResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["redirectResponseCode"] = v
	}
	if v := f.HttpsRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["httpsRedirect"] = v
	}
	if v := f.StripQuery; !dcl.IsEmptyValueIndirect(v) {
		m["stripQuery"] = v
	}

	return m, nil
}

// flattenUrlMapDefaultUrlRedirect flattens an instance of UrlMapDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapDefaultUrlRedirect(c *Client, i interface{}) *UrlMapDefaultUrlRedirect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapDefaultUrlRedirect{}
	r.HostRedirect = dcl.FlattenString(m["hostRedirect"])
	r.PathRedirect = dcl.FlattenString(m["pathRedirect"])
	r.PrefixRedirect = dcl.FlattenString(m["prefixRedirect"])
	r.RedirectResponseCode = flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(m["redirectResponseCode"])
	r.HttpsRedirect = dcl.FlattenBool(m["httpsRedirect"])
	r.StripQuery = dcl.FlattenBool(m["stripQuery"])

	return r
}

// expandUrlMapHostRuleMap expands the contents of UrlMapHostRule into a JSON
// request object.
func expandUrlMapHostRuleMap(c *Client, f map[string]UrlMapHostRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapHostRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapHostRuleSlice expands the contents of UrlMapHostRule into a JSON
// request object.
func expandUrlMapHostRuleSlice(c *Client, f []UrlMapHostRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapHostRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapHostRuleMap flattens the contents of UrlMapHostRule from a JSON
// response object.
func flattenUrlMapHostRuleMap(c *Client, i interface{}) map[string]UrlMapHostRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapHostRule{}
	}

	if len(a) == 0 {
		return map[string]UrlMapHostRule{}
	}

	items := make(map[string]UrlMapHostRule)
	for k, item := range a {
		items[k] = *flattenUrlMapHostRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapHostRuleSlice flattens the contents of UrlMapHostRule from a JSON
// response object.
func flattenUrlMapHostRuleSlice(c *Client, i interface{}) []UrlMapHostRule {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapHostRule{}
	}

	if len(a) == 0 {
		return []UrlMapHostRule{}
	}

	items := make([]UrlMapHostRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapHostRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapHostRule expands an instance of UrlMapHostRule into a JSON
// request object.
func expandUrlMapHostRule(c *Client, f *UrlMapHostRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["hosts"] = v
	}
	if v := f.PathMatcher; !dcl.IsEmptyValueIndirect(v) {
		m["pathMatcher"] = v
	}

	return m, nil
}

// flattenUrlMapHostRule flattens an instance of UrlMapHostRule from a JSON
// response object.
func flattenUrlMapHostRule(c *Client, i interface{}) *UrlMapHostRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapHostRule{}
	r.Description = dcl.FlattenString(m["description"])
	r.Host = dcl.FlattenStringSlice(m["hosts"])
	r.PathMatcher = dcl.FlattenString(m["pathMatcher"])

	return r
}

// expandUrlMapPathMatcherMap expands the contents of UrlMapPathMatcher into a JSON
// request object.
func expandUrlMapPathMatcherMap(c *Client, f map[string]UrlMapPathMatcher) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcher(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherSlice expands the contents of UrlMapPathMatcher into a JSON
// request object.
func expandUrlMapPathMatcherSlice(c *Client, f []UrlMapPathMatcher) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcher(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherMap flattens the contents of UrlMapPathMatcher from a JSON
// response object.
func flattenUrlMapPathMatcherMap(c *Client, i interface{}) map[string]UrlMapPathMatcher {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcher{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcher{}
	}

	items := make(map[string]UrlMapPathMatcher)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcher(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherSlice flattens the contents of UrlMapPathMatcher from a JSON
// response object.
func flattenUrlMapPathMatcherSlice(c *Client, i interface{}) []UrlMapPathMatcher {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcher{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcher{}
	}

	items := make([]UrlMapPathMatcher, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcher(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcher expands an instance of UrlMapPathMatcher into a JSON
// request object.
func expandUrlMapPathMatcher(c *Client, f *UrlMapPathMatcher) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DefaultService; !dcl.IsEmptyValueIndirect(v) {
		m["defaultService"] = v
	}
	if v, err := expandUrlMapDefaultRouteAction(c, f.DefaultRouteAction); err != nil {
		return nil, fmt.Errorf("error expanding DefaultRouteAction into defaultRouteAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultRouteAction"] = v
	}
	if v, err := expandUrlMapPathMatcherDefaultUrlRedirect(c, f.DefaultUrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding DefaultUrlRedirect into defaultUrlRedirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultUrlRedirect"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleSlice(c, f.PathRule); err != nil {
		return nil, fmt.Errorf("error expanding PathRule into pathRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pathRules"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleSlice(c, f.RouteRule); err != nil {
		return nil, fmt.Errorf("error expanding RouteRule into routeRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["routeRules"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcher flattens an instance of UrlMapPathMatcher from a JSON
// response object.
func flattenUrlMapPathMatcher(c *Client, i interface{}) *UrlMapPathMatcher {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcher{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.DefaultService = dcl.FlattenString(m["defaultService"])
	r.DefaultRouteAction = flattenUrlMapDefaultRouteAction(c, m["defaultRouteAction"])
	r.DefaultUrlRedirect = flattenUrlMapPathMatcherDefaultUrlRedirect(c, m["defaultUrlRedirect"])
	r.PathRule = flattenUrlMapPathMatcherPathRuleSlice(c, m["pathRules"])
	r.RouteRule = flattenUrlMapPathMatcherRouteRuleSlice(c, m["routeRules"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])

	return r
}

// expandUrlMapPathMatcherDefaultUrlRedirectMap expands the contents of UrlMapPathMatcherDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherDefaultUrlRedirectMap(c *Client, f map[string]UrlMapPathMatcherDefaultUrlRedirect) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherDefaultUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherDefaultUrlRedirectSlice expands the contents of UrlMapPathMatcherDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherDefaultUrlRedirectSlice(c *Client, f []UrlMapPathMatcherDefaultUrlRedirect) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherDefaultUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherDefaultUrlRedirectMap flattens the contents of UrlMapPathMatcherDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherDefaultUrlRedirectMap(c *Client, i interface{}) map[string]UrlMapPathMatcherDefaultUrlRedirect {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherDefaultUrlRedirect{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherDefaultUrlRedirect{}
	}

	items := make(map[string]UrlMapPathMatcherDefaultUrlRedirect)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherDefaultUrlRedirect(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherDefaultUrlRedirectSlice flattens the contents of UrlMapPathMatcherDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherDefaultUrlRedirectSlice(c *Client, i interface{}) []UrlMapPathMatcherDefaultUrlRedirect {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherDefaultUrlRedirect{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherDefaultUrlRedirect{}
	}

	items := make([]UrlMapPathMatcherDefaultUrlRedirect, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherDefaultUrlRedirect(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherDefaultUrlRedirect expands an instance of UrlMapPathMatcherDefaultUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherDefaultUrlRedirect(c *Client, f *UrlMapPathMatcherDefaultUrlRedirect) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["hostRedirect"] = v
	}
	if v := f.PathRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["pathRedirect"] = v
	}
	if v := f.PrefixRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["prefixRedirect"] = v
	}
	if v := f.RedirectResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["redirectResponseCode"] = v
	}
	if v := f.HttpsRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["httpsRedirect"] = v
	}
	if v := f.StripQuery; !dcl.IsEmptyValueIndirect(v) {
		m["stripQuery"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherDefaultUrlRedirect flattens an instance of UrlMapPathMatcherDefaultUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherDefaultUrlRedirect(c *Client, i interface{}) *UrlMapPathMatcherDefaultUrlRedirect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherDefaultUrlRedirect{}
	r.HostRedirect = dcl.FlattenString(m["hostRedirect"])
	r.PathRedirect = dcl.FlattenString(m["pathRedirect"])
	r.PrefixRedirect = dcl.FlattenString(m["prefixRedirect"])
	r.RedirectResponseCode = flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(m["redirectResponseCode"])
	r.HttpsRedirect = dcl.FlattenBool(m["httpsRedirect"])
	r.StripQuery = dcl.FlattenBool(m["stripQuery"])

	return r
}

// expandUrlMapPathMatcherPathRuleMap expands the contents of UrlMapPathMatcherPathRule into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleMap(c *Client, f map[string]UrlMapPathMatcherPathRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleSlice expands the contents of UrlMapPathMatcherPathRule into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleSlice(c *Client, f []UrlMapPathMatcherPathRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleMap flattens the contents of UrlMapPathMatcherPathRule from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRule{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRule{}
	}

	items := make(map[string]UrlMapPathMatcherPathRule)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleSlice flattens the contents of UrlMapPathMatcherPathRule from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRule {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRule{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRule{}
	}

	items := make([]UrlMapPathMatcherPathRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRule expands an instance of UrlMapPathMatcherPathRule into a JSON
// request object.
func expandUrlMapPathMatcherPathRule(c *Client, f *UrlMapPathMatcherPathRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteAction(c, f.RouteAction); err != nil {
		return nil, fmt.Errorf("error expanding RouteAction into routeAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["routeAction"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleUrlRedirect(c, f.UrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding UrlRedirect into urlRedirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRedirect"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["paths"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRule flattens an instance of UrlMapPathMatcherPathRule from a JSON
// response object.
func flattenUrlMapPathMatcherPathRule(c *Client, i interface{}) *UrlMapPathMatcherPathRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRule{}
	r.BackendService = dcl.FlattenString(m["service"])
	r.RouteAction = flattenUrlMapPathMatcherPathRuleRouteAction(c, m["routeAction"])
	r.UrlRedirect = flattenUrlMapPathMatcherPathRuleUrlRedirect(c, m["urlRedirect"])
	r.Path = dcl.FlattenStringSlice(m["paths"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionMap expands the contents of UrlMapPathMatcherPathRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteAction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteAction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionSlice expands the contents of UrlMapPathMatcherPathRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteAction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteAction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionMap flattens the contents of UrlMapPathMatcherPathRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteAction{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteAction{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteAction)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteAction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionSlice flattens the contents of UrlMapPathMatcherPathRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteAction {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteAction{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteAction{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteAction(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteAction expands an instance of UrlMapPathMatcherPathRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteAction(c *Client, f *UrlMapPathMatcherPathRuleRouteAction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c, f.WeightedBackendService); err != nil {
		return nil, fmt.Errorf("error expanding WeightedBackendService into weightedBackendServices: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["weightedBackendServices"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, f.UrlRewrite); err != nil {
		return nil, fmt.Errorf("error expanding UrlRewrite into urlRewrite: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRewrite"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionTimeout(c, f.Timeout); err != nil {
		return nil, fmt.Errorf("error expanding Timeout into timeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, f.RetryPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RetryPolicy into retryPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, f.RequestMirrorPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RequestMirrorPolicy into requestMirrorPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestMirrorPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, f.CorsPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CorsPolicy into corsPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["corsPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, f.FaultInjectionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding FaultInjectionPolicy into faultInjectionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["faultInjectionPolicy"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteAction flattens an instance of UrlMapPathMatcherPathRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteAction(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteAction{}
	r.WeightedBackendService = flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c, m["weightedBackendServices"])
	r.UrlRewrite = flattenUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, m["urlRewrite"])
	r.Timeout = flattenUrlMapPathMatcherPathRuleRouteActionTimeout(c, m["timeout"])
	r.RetryPolicy = flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, m["retryPolicy"])
	r.RequestMirrorPolicy = flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, m["requestMirrorPolicy"])
	r.CorsPolicy = flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, m["corsPolicy"])
	r.FaultInjectionPolicy = flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, m["faultInjectionPolicy"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceMap expands the contents of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice expands the contents of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionWeightedBackendService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendService expands an instance of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c *Client, f *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendService flattens an instance of UrlMapPathMatcherPathRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}
	r.BackendService = dcl.FlattenString(m["backendService"])
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionUrlRewriteMap expands the contents of UrlMapPathMatcherPathRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionUrlRewriteMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionUrlRewrite) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice expands the contents of UrlMapPathMatcherPathRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionUrlRewrite) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionUrlRewriteMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionUrlRewriteMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionUrlRewrite)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionUrlRewrite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionUrlRewrite expands an instance of UrlMapPathMatcherPathRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c *Client, f *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PathPrefixRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["pathPrefixRewrite"] = v
	}
	if v := f.HostRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["hostRewrite"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionUrlRewrite flattens an instance of UrlMapPathMatcherPathRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}
	r.PathPrefixRewrite = dcl.FlattenString(m["pathPrefixRewrite"])
	r.HostRewrite = dcl.FlattenString(m["hostRewrite"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionTimeoutMap expands the contents of UrlMapPathMatcherPathRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionTimeoutMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionTimeoutSlice expands the contents of UrlMapPathMatcherPathRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionTimeoutSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionTimeoutMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionTimeoutMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionTimeout{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionTimeoutSlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionTimeoutSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionTimeout{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionTimeout expands an instance of UrlMapPathMatcherPathRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionTimeout(c *Client, f *UrlMapPathMatcherPathRuleRouteActionTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionTimeout flattens an instance of UrlMapPathMatcherPathRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionTimeout(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyMap expands the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionRetryPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionRetryPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicy expands an instance of UrlMapPathMatcherPathRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c *Client, f *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryCondition; !dcl.IsEmptyValueIndirect(v) {
		m["retryConditions"] = v
	}
	if v := f.NumRetries; !dcl.IsEmptyValueIndirect(v) {
		m["numRetries"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, f.PerTryTimeout); err != nil {
		return nil, fmt.Errorf("error expanding PerTryTimeout into perTryTimeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["perTryTimeout"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicy flattens an instance of UrlMapPathMatcherPathRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}
	r.RetryCondition = dcl.FlattenStringSlice(m["retryConditions"])
	r.NumRetries = dcl.FlattenInteger(m["numRetries"])
	r.PerTryTimeout = flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, m["perTryTimeout"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutMap expands the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice expands the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout expands an instance of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c *Client, f *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout flattens an instance of UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyMap expands the contents of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy expands an instance of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c *Client, f *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy flattens an instance of UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}
	r.BackendService = dcl.FlattenString(m["backendService"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionCorsPolicyMap expands the contents of UrlMapPathMatcherPathRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionCorsPolicyMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionCorsPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionCorsPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicyMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionCorsPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionCorsPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionCorsPolicy expands an instance of UrlMapPathMatcherPathRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c *Client, f *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowOrigin; !dcl.IsEmptyValueIndirect(v) {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; !dcl.IsEmptyValueIndirect(v) {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; !dcl.IsEmptyValueIndirect(v) {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; !dcl.IsEmptyValueIndirect(v) {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; !dcl.IsEmptyValueIndirect(v) {
		m["exposeHeaders"] = v
	}
	if v := f.MaxAge; !dcl.IsEmptyValueIndirect(v) {
		m["maxAge"] = v
	}
	if v := f.AllowCredentials; !dcl.IsEmptyValueIndirect(v) {
		m["allowCredentials"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicy flattens an instance of UrlMapPathMatcherPathRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}
	r.AllowOrigin = dcl.FlattenStringSlice(m["allowOrigins"])
	r.AllowOriginRegex = dcl.FlattenStringSlice(m["allowOriginRegexes"])
	r.AllowMethod = dcl.FlattenStringSlice(m["allowMethods"])
	r.AllowHeader = dcl.FlattenStringSlice(m["allowHeaders"])
	r.ExposeHeader = dcl.FlattenStringSlice(m["exposeHeaders"])
	r.MaxAge = dcl.FlattenInteger(m["maxAge"])
	r.AllowCredentials = dcl.FlattenBool(m["allowCredentials"])
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyMap expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy expands an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c *Client, f *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, f.Delay); err != nil {
		return nil, fmt.Errorf("error expanding Delay into delay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["delay"] = v
	}
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, f.Abort); err != nil {
		return nil, fmt.Errorf("error expanding Abort into abort: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["abort"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy flattens an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}
	r.Delay = flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, m["delay"])
	r.Abort = flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, m["abort"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayMap expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay expands an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c *Client, f *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, f.FixedDelay); err != nil {
		return nil, fmt.Errorf("error expanding FixedDelay into fixedDelay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixedDelay"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay flattens an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}
	r.FixedDelay = flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, m["fixedDelay"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay expands an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, f *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay flattens an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortMap expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortMap(c *Client, f map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice expands the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, f []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortMap flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice flattens the contents of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}
	}

	items := make([]UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort expands an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c *Client, f *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpStatus; !dcl.IsEmptyValueIndirect(v) {
		m["httpStatus"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort flattens an instance of UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c *Client, i interface{}) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}
	r.HttpStatus = dcl.FlattenInteger(m["httpStatus"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapPathMatcherPathRuleUrlRedirectMap expands the contents of UrlMapPathMatcherPathRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleUrlRedirectMap(c *Client, f map[string]UrlMapPathMatcherPathRuleUrlRedirect) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherPathRuleUrlRedirectSlice expands the contents of UrlMapPathMatcherPathRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleUrlRedirectSlice(c *Client, f []UrlMapPathMatcherPathRuleUrlRedirect) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherPathRuleUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherPathRuleUrlRedirectMap flattens the contents of UrlMapPathMatcherPathRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleUrlRedirectMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleUrlRedirect {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleUrlRedirect{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleUrlRedirect{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleUrlRedirect)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleUrlRedirect(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleUrlRedirectSlice flattens the contents of UrlMapPathMatcherPathRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleUrlRedirectSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleUrlRedirect {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleUrlRedirect{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleUrlRedirect{}
	}

	items := make([]UrlMapPathMatcherPathRuleUrlRedirect, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleUrlRedirect(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherPathRuleUrlRedirect expands an instance of UrlMapPathMatcherPathRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherPathRuleUrlRedirect(c *Client, f *UrlMapPathMatcherPathRuleUrlRedirect) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["hostRedirect"] = v
	}
	if v := f.PathRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["pathRedirect"] = v
	}
	if v := f.PrefixRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["prefixRedirect"] = v
	}
	if v := f.RedirectResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["redirectResponseCode"] = v
	}
	if v := f.HttpsRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["httpsRedirect"] = v
	}
	if v := f.StripQuery; !dcl.IsEmptyValueIndirect(v) {
		m["stripQuery"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherPathRuleUrlRedirect flattens an instance of UrlMapPathMatcherPathRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleUrlRedirect(c *Client, i interface{}) *UrlMapPathMatcherPathRuleUrlRedirect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherPathRuleUrlRedirect{}
	r.HostRedirect = dcl.FlattenString(m["hostRedirect"])
	r.PathRedirect = dcl.FlattenString(m["pathRedirect"])
	r.PrefixRedirect = dcl.FlattenString(m["prefixRedirect"])
	r.RedirectResponseCode = flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(m["redirectResponseCode"])
	r.HttpsRedirect = dcl.FlattenBool(m["httpsRedirect"])
	r.StripQuery = dcl.FlattenBool(m["stripQuery"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMap expands the contents of UrlMapPathMatcherRouteRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMap(c *Client, f map[string]UrlMapPathMatcherRouteRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleSlice expands the contents of UrlMapPathMatcherRouteRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleSlice(c *Client, f []UrlMapPathMatcherRouteRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMap flattens the contents of UrlMapPathMatcherRouteRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRule{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRule{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRule)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleSlice flattens the contents of UrlMapPathMatcherRouteRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRule {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRule{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRule{}
	}

	items := make([]UrlMapPathMatcherRouteRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRule expands an instance of UrlMapPathMatcherRouteRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRule(c *Client, f *UrlMapPathMatcherRouteRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		m["priority"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleSlice(c, f.MatchRule); err != nil {
		return nil, fmt.Errorf("error expanding MatchRule into matchRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["matchRules"] = v
	}
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteAction(c, f.RouteAction); err != nil {
		return nil, fmt.Errorf("error expanding RouteAction into routeAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["routeAction"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleUrlRedirect(c, f.UrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding UrlRedirect into urlRedirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRedirect"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRule flattens an instance of UrlMapPathMatcherRouteRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRule(c *Client, i interface{}) *UrlMapPathMatcherRouteRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRule{}
	r.Priority = dcl.FlattenInteger(m["priority"])
	r.Description = dcl.FlattenString(m["description"])
	r.MatchRule = flattenUrlMapPathMatcherRouteRuleMatchRuleSlice(c, m["matchRules"])
	r.BackendService = dcl.FlattenString(m["service"])
	r.RouteAction = flattenUrlMapPathMatcherRouteRuleRouteAction(c, m["routeAction"])
	r.UrlRedirect = flattenUrlMapPathMatcherRouteRuleUrlRedirect(c, m["urlRedirect"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMap expands the contents of UrlMapPathMatcherRouteRuleMatchRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRule{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRule{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRule)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRule {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRule{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRule{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRule expands an instance of UrlMapPathMatcherRouteRuleMatchRule into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRule(c *Client, f *UrlMapPathMatcherRouteRuleMatchRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PrefixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["prefixMatch"] = v
	}
	if v := f.FullPathMatch; !dcl.IsEmptyValueIndirect(v) {
		m["fullPathMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}
	if v := f.IgnoreCase; !dcl.IsEmptyValueIndirect(v) {
		m["ignoreCase"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c, f.HeaderMatch); err != nil {
		return nil, fmt.Errorf("error expanding HeaderMatch into headerMatches: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerMatches"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c, f.QueryParameterMatch); err != nil {
		return nil, fmt.Errorf("error expanding QueryParameterMatch into queryParameterMatches: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["queryParameterMatches"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c, f.MetadataFilter); err != nil {
		return nil, fmt.Errorf("error expanding MetadataFilter into metadataFilters: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataFilters"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRule flattens an instance of UrlMapPathMatcherRouteRuleMatchRule from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRule(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRule{}
	r.PrefixMatch = dcl.FlattenString(m["prefixMatch"])
	r.FullPathMatch = dcl.FlattenString(m["fullPathMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])
	r.IgnoreCase = dcl.FlattenBool(m["ignoreCase"])
	r.HeaderMatch = flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c, m["headerMatches"])
	r.QueryParameterMatch = flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c, m["queryParameterMatches"])
	r.MetadataFilter = flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c, m["metadataFilters"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchMap expands the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch expands an instance of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c *Client, f *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}
	if v := f.ExactMatch; !dcl.IsEmptyValueIndirect(v) {
		m["exactMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, f.RangeMatch); err != nil {
		return nil, fmt.Errorf("error expanding RangeMatch into rangeMatch: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rangeMatch"] = v
	}
	if v := f.PresentMatch; !dcl.IsEmptyValueIndirect(v) {
		m["presentMatch"] = v
	}
	if v := f.PrefixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["prefixMatch"] = v
	}
	if v := f.SuffixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["suffixMatch"] = v
	}
	if v := f.InvertMatch; !dcl.IsEmptyValueIndirect(v) {
		m["invertMatch"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch flattens an instance of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}
	r.HeaderName = dcl.FlattenString(m["headerName"])
	r.ExactMatch = dcl.FlattenString(m["exactMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])
	r.RangeMatch = flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, m["rangeMatch"])
	r.PresentMatch = dcl.FlattenBool(m["presentMatch"])
	r.PrefixMatch = dcl.FlattenString(m["prefixMatch"])
	r.SuffixMatch = dcl.FlattenString(m["suffixMatch"])
	r.InvertMatch = dcl.FlattenBool(m["invertMatch"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchMap expands the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch expands an instance of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c *Client, f *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RangeStart; !dcl.IsEmptyValueIndirect(v) {
		m["rangeStart"] = v
	}
	if v := f.RangeEnd; !dcl.IsEmptyValueIndirect(v) {
		m["rangeEnd"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch flattens an instance of UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}
	r.RangeStart = dcl.FlattenInteger(m["rangeStart"])
	r.RangeEnd = dcl.FlattenInteger(m["rangeEnd"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchMap expands the contents of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch expands an instance of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c *Client, f *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.PresentMatch; !dcl.IsEmptyValueIndirect(v) {
		m["presentMatch"] = v
	}
	if v := f.ExactMatch; !dcl.IsEmptyValueIndirect(v) {
		m["exactMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch flattens an instance of UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}
	r.Name = dcl.FlattenString(m["name"])
	r.PresentMatch = dcl.FlattenBool(m["presentMatch"])
	r.ExactMatch = dcl.FlattenString(m["exactMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterMap expands the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter expands an instance of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c *Client, f *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FilterMatchCriteria; !dcl.IsEmptyValueIndirect(v) {
		m["filterMatchCriteria"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c, f.FilterLabel); err != nil {
		return nil, fmt.Errorf("error expanding FilterLabel into filterLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["filterLabels"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter flattens an instance of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}
	r.FilterMatchCriteria = flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(m["filterMatchCriteria"])
	r.FilterLabel = flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c, m["filterLabels"])

	return r
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelMap expands the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice expands the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c *Client, f []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel expands an instance of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c *Client, f *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel flattens an instance of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionMap expands the contents of UrlMapPathMatcherRouteRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteAction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteAction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionSlice expands the contents of UrlMapPathMatcherRouteRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteAction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteAction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionMap flattens the contents of UrlMapPathMatcherRouteRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteAction{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteAction{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteAction)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteAction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteAction {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteAction{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteAction{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteAction(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteAction expands an instance of UrlMapPathMatcherRouteRuleRouteAction into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteAction(c *Client, f *UrlMapPathMatcherRouteRuleRouteAction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c, f.WeightedBackendService); err != nil {
		return nil, fmt.Errorf("error expanding WeightedBackendService into weightedBackendServices: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["weightedBackendServices"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, f.UrlRewrite); err != nil {
		return nil, fmt.Errorf("error expanding UrlRewrite into urlRewrite: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRewrite"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionTimeout(c, f.Timeout); err != nil {
		return nil, fmt.Errorf("error expanding Timeout into timeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, f.RetryPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RetryPolicy into retryPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, f.RequestMirrorPolicy); err != nil {
		return nil, fmt.Errorf("error expanding RequestMirrorPolicy into requestMirrorPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestMirrorPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, f.CorsPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CorsPolicy into corsPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["corsPolicy"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, f.FaultInjectionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding FaultInjectionPolicy into faultInjectionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["faultInjectionPolicy"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteAction flattens an instance of UrlMapPathMatcherRouteRuleRouteAction from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteAction(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteAction{}
	r.WeightedBackendService = flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c, m["weightedBackendServices"])
	r.UrlRewrite = flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, m["urlRewrite"])
	r.Timeout = flattenUrlMapPathMatcherRouteRuleRouteActionTimeout(c, m["timeout"])
	r.RetryPolicy = flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, m["retryPolicy"])
	r.RequestMirrorPolicy = flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, m["requestMirrorPolicy"])
	r.CorsPolicy = flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, m["corsPolicy"])
	r.FaultInjectionPolicy = flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, m["faultInjectionPolicy"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService expands an instance of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["headerAction"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService flattens an instance of UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}
	r.BackendService = dcl.FlattenString(m["backendService"])
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionUrlRewriteMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionUrlRewriteMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewriteMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewriteMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionUrlRewrite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionUrlRewrite expands an instance of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PathPrefixRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["pathPrefixRewrite"] = v
	}
	if v := f.HostRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["hostRewrite"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewrite flattens an instance of UrlMapPathMatcherRouteRuleRouteActionUrlRewrite from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}
	r.PathPrefixRewrite = dcl.FlattenString(m["pathPrefixRewrite"])
	r.HostRewrite = dcl.FlattenString(m["hostRewrite"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionTimeoutMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionTimeoutMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionTimeoutMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionTimeoutMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionTimeout{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionTimeout{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionTimeout expands an instance of UrlMapPathMatcherRouteRuleRouteActionTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionTimeout(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionTimeout flattens an instance of UrlMapPathMatcherRouteRuleRouteActionTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionTimeout(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionRetryPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicy expands an instance of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryCondition; !dcl.IsEmptyValueIndirect(v) {
		m["retryConditions"] = v
	}
	if v := f.NumRetries; !dcl.IsEmptyValueIndirect(v) {
		m["numRetries"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, f.PerTryTimeout); err != nil {
		return nil, fmt.Errorf("error expanding PerTryTimeout into perTryTimeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["perTryTimeout"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicy flattens an instance of UrlMapPathMatcherRouteRuleRouteActionRetryPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}
	r.RetryCondition = dcl.FlattenStringSlice(m["retryConditions"])
	r.NumRetries = dcl.FlattenInteger(m["numRetries"])
	r.PerTryTimeout = flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, m["perTryTimeout"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout expands an instance of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout flattens an instance of UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy expands an instance of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy flattens an instance of UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}
	r.BackendService = dcl.FlattenString(m["backendService"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicyMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicyMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicyMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionCorsPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicy expands an instance of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowOrigin; !dcl.IsEmptyValueIndirect(v) {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; !dcl.IsEmptyValueIndirect(v) {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; !dcl.IsEmptyValueIndirect(v) {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; !dcl.IsEmptyValueIndirect(v) {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; !dcl.IsEmptyValueIndirect(v) {
		m["exposeHeaders"] = v
	}
	if v := f.MaxAge; !dcl.IsEmptyValueIndirect(v) {
		m["maxAge"] = v
	}
	if v := f.AllowCredentials; !dcl.IsEmptyValueIndirect(v) {
		m["allowCredentials"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicy flattens an instance of UrlMapPathMatcherRouteRuleRouteActionCorsPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}
	r.AllowOrigin = dcl.FlattenStringSlice(m["allowOrigins"])
	r.AllowOriginRegex = dcl.FlattenStringSlice(m["allowOriginRegexes"])
	r.AllowMethod = dcl.FlattenStringSlice(m["allowMethods"])
	r.AllowHeader = dcl.FlattenStringSlice(m["allowHeaders"])
	r.ExposeHeader = dcl.FlattenStringSlice(m["exposeHeaders"])
	r.MaxAge = dcl.FlattenInteger(m["maxAge"])
	r.AllowCredentials = dcl.FlattenBool(m["allowCredentials"])
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy expands an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, f.Delay); err != nil {
		return nil, fmt.Errorf("error expanding Delay into delay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["delay"] = v
	}
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, f.Abort); err != nil {
		return nil, fmt.Errorf("error expanding Abort into abort: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["abort"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy flattens an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}
	r.Delay = flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, m["delay"])
	r.Abort = flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, m["abort"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay expands an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, f.FixedDelay); err != nil {
		return nil, fmt.Errorf("error expanding FixedDelay into fixedDelay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixedDelay"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay flattens an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}
	r.FixedDelay = flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, m["fixedDelay"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay expands an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay flattens an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortMap expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice expands the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, f []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortMap flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice flattens the contents of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}
	}

	items := make([]UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort expands an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c *Client, f *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpStatus; !dcl.IsEmptyValueIndirect(v) {
		m["httpStatus"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort flattens an instance of UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}
	r.HttpStatus = dcl.FlattenInteger(m["httpStatus"])
	r.Percentage = dcl.FlattenDouble(m["percentage"])

	return r
}

// expandUrlMapPathMatcherRouteRuleUrlRedirectMap expands the contents of UrlMapPathMatcherRouteRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleUrlRedirectMap(c *Client, f map[string]UrlMapPathMatcherRouteRuleUrlRedirect) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapPathMatcherRouteRuleUrlRedirectSlice expands the contents of UrlMapPathMatcherRouteRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleUrlRedirectSlice(c *Client, f []UrlMapPathMatcherRouteRuleUrlRedirect) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapPathMatcherRouteRuleUrlRedirect(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapPathMatcherRouteRuleUrlRedirectMap flattens the contents of UrlMapPathMatcherRouteRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleUrlRedirectMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleUrlRedirect {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleUrlRedirect{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleUrlRedirect{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleUrlRedirect)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleUrlRedirect(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleUrlRedirectSlice flattens the contents of UrlMapPathMatcherRouteRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleUrlRedirectSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleUrlRedirect {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleUrlRedirect{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleUrlRedirect{}
	}

	items := make([]UrlMapPathMatcherRouteRuleUrlRedirect, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleUrlRedirect(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapPathMatcherRouteRuleUrlRedirect expands an instance of UrlMapPathMatcherRouteRuleUrlRedirect into a JSON
// request object.
func expandUrlMapPathMatcherRouteRuleUrlRedirect(c *Client, f *UrlMapPathMatcherRouteRuleUrlRedirect) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["hostRedirect"] = v
	}
	if v := f.PathRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["pathRedirect"] = v
	}
	if v := f.PrefixRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["prefixRedirect"] = v
	}
	if v := f.RedirectResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["redirectResponseCode"] = v
	}
	if v := f.HttpsRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["httpsRedirect"] = v
	}
	if v := f.StripQuery; !dcl.IsEmptyValueIndirect(v) {
		m["stripQuery"] = v
	}

	return m, nil
}

// flattenUrlMapPathMatcherRouteRuleUrlRedirect flattens an instance of UrlMapPathMatcherRouteRuleUrlRedirect from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleUrlRedirect(c *Client, i interface{}) *UrlMapPathMatcherRouteRuleUrlRedirect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapPathMatcherRouteRuleUrlRedirect{}
	r.HostRedirect = dcl.FlattenString(m["hostRedirect"])
	r.PathRedirect = dcl.FlattenString(m["pathRedirect"])
	r.PrefixRedirect = dcl.FlattenString(m["prefixRedirect"])
	r.RedirectResponseCode = flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(m["redirectResponseCode"])
	r.HttpsRedirect = dcl.FlattenBool(m["httpsRedirect"])
	r.StripQuery = dcl.FlattenBool(m["stripQuery"])

	return r
}

// expandUrlMapTestMap expands the contents of UrlMapTest into a JSON
// request object.
func expandUrlMapTestMap(c *Client, f map[string]UrlMapTest) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUrlMapTest(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUrlMapTestSlice expands the contents of UrlMapTest into a JSON
// request object.
func expandUrlMapTestSlice(c *Client, f []UrlMapTest) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUrlMapTest(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUrlMapTestMap flattens the contents of UrlMapTest from a JSON
// response object.
func flattenUrlMapTestMap(c *Client, i interface{}) map[string]UrlMapTest {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapTest{}
	}

	if len(a) == 0 {
		return map[string]UrlMapTest{}
	}

	items := make(map[string]UrlMapTest)
	for k, item := range a {
		items[k] = *flattenUrlMapTest(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUrlMapTestSlice flattens the contents of UrlMapTest from a JSON
// response object.
func flattenUrlMapTestSlice(c *Client, i interface{}) []UrlMapTest {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapTest{}
	}

	if len(a) == 0 {
		return []UrlMapTest{}
	}

	items := make([]UrlMapTest, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapTest(c, item.(map[string]interface{})))
	}

	return items
}

// expandUrlMapTest expands an instance of UrlMapTest into a JSON
// request object.
func expandUrlMapTest(c *Client, f *UrlMapTest) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.ExpectedBackendService; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}

	return m, nil
}

// flattenUrlMapTest flattens an instance of UrlMapTest from a JSON
// response object.
func flattenUrlMapTest(c *Client, i interface{}) *UrlMapTest {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UrlMapTest{}
	r.Description = dcl.FlattenString(m["description"])
	r.Host = dcl.FlattenString(m["host"])
	r.Path = dcl.FlattenString(m["path"])
	r.ExpectedBackendService = dcl.FlattenString(m["service"])

	return r
}

// flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnumSlice flattens the contents of UrlMapDefaultUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnumSlice(c *Client, i interface{}) []UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []UrlMapDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make([]UrlMapDefaultUrlRedirectRedirectResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum with the same value as that string.
func flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(i interface{}) *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return UrlMapDefaultUrlRedirectRedirectResponseCodeEnumRef("")
	}

	return UrlMapDefaultUrlRedirectRedirectResponseCodeEnumRef(s)
}

// flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumSlice flattens the contents of UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumSlice(c *Client, i interface{}) []UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make([]UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum with the same value as that string.
func flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(i interface{}) *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumRef("")
	}

	return UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumRef(s)
}

// flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumSlice flattens the contents of UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumSlice(c *Client, i interface{}) []UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make([]UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum with the same value as that string.
func flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(i interface{}) *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumRef("")
	}

	return UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumRef(s)
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumSlice flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	items := make([]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum asserts that an interface is a string, and returns a
// pointer to a *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum with the same value as that string.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(i interface{}) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	s, ok := i.(string)
	if !ok {
		return UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumRef("")
	}

	return UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumRef(s)
}

// flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumSlice flattens the contents of UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumSlice(c *Client, i interface{}) []UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make([]UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum with the same value as that string.
func flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(i interface{}) *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumRef("")
	}

	return UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *UrlMap) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalUrlMap(b, c)
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
