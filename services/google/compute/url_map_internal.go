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
	"strings"
	"time"

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
	b, err := c.getUrlMapRaw(ctx, f.URLNormalized())
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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUrlMapUpdateOperation) do(ctx context.Context, r *UrlMap, c *Client) error {
	_, err := c.GetUrlMap(ctx, r.URLNormalized())
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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
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
		res, err := unmarshalMapUrlMap(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetUrlMap(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("UrlMap not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUrlMap checking for existence. error: %v", err)
		return err
	}

	u, err := urlMapDeleteURL(c.Config.BasePath, r.URLNormalized())
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

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetUrlMap(ctx, r.URLNormalized())
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
type createUrlMapOperation struct {
	response map[string]interface{}
}

func (op *createUrlMapOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

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

	if _, err := c.GetUrlMap(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getUrlMapRaw(ctx context.Context, r *UrlMap) ([]byte, error) {

	u, err := urlMapGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) urlMapDiffsForRawDesired(ctx context.Context, rawDesired *UrlMap, opts ...dcl.ApplyOption) (initial, desired *UrlMap, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetUrlMap(ctx, fetchState.URLNormalized())
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

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(rawDesired.DefaultRouteAction, nil, opts...)
		rawDesired.DefaultUrlRedirect = canonicalizeUrlMapDefaultUrlRedirect(rawDesired.DefaultUrlRedirect, nil, opts...)
		rawDesired.HeaderAction = canonicalizeUrlMapHeaderAction(rawDesired.HeaderAction, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &UrlMap{}
	canonicalDesired.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(rawDesired.DefaultRouteAction, rawInitial.DefaultRouteAction, opts...)
	if dcl.StringCanonicalize(rawDesired.DefaultService, rawInitial.DefaultService) {
		canonicalDesired.DefaultService = rawInitial.DefaultService
	} else {
		canonicalDesired.DefaultService = rawDesired.DefaultService
	}
	canonicalDesired.DefaultUrlRedirect = canonicalizeUrlMapDefaultUrlRedirect(rawDesired.DefaultUrlRedirect, rawInitial.DefaultUrlRedirect, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.HeaderAction = canonicalizeUrlMapHeaderAction(rawDesired.HeaderAction, rawInitial.HeaderAction, opts...)
	if dcl.IsZeroValue(rawDesired.HostRule) {
		canonicalDesired.HostRule = rawInitial.HostRule
	} else {
		canonicalDesired.HostRule = rawDesired.HostRule
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.PathMatcher) {
		canonicalDesired.PathMatcher = rawInitial.PathMatcher
	} else {
		canonicalDesired.PathMatcher = rawDesired.PathMatcher
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		canonicalDesired.Region = rawInitial.Region
	} else {
		canonicalDesired.Region = rawDesired.Region
	}
	if dcl.IsZeroValue(rawDesired.Test) {
		canonicalDesired.Test = rawInitial.Test
	} else {
		canonicalDesired.Test = rawDesired.Test
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
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
		if dcl.StringCanonicalize(rawDesired.DefaultService, rawNew.DefaultService) {
			rawNew.DefaultService = rawDesired.DefaultService
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultUrlRedirect) && dcl.IsEmptyValueIndirect(rawDesired.DefaultUrlRedirect) {
		rawNew.DefaultUrlRedirect = rawDesired.DefaultUrlRedirect
	} else {
		rawNew.DefaultUrlRedirect = canonicalizeNewUrlMapDefaultUrlRedirect(c, rawDesired.DefaultUrlRedirect, rawNew.DefaultUrlRedirect)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
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
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PathMatcher) && dcl.IsEmptyValueIndirect(rawDesired.PathMatcher) {
		rawNew.PathMatcher = rawDesired.PathMatcher
	} else {
		rawNew.PathMatcher = canonicalizeNewUrlMapPathMatcherSet(c, rawDesired.PathMatcher, rawNew.PathMatcher)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Test) && dcl.IsEmptyValueIndirect(rawDesired.Test) {
		rawNew.Test = rawDesired.Test
	} else {
		rawNew.Test = canonicalizeNewUrlMapTestSlice(c, rawDesired.Test, rawNew.Test)
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

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteAction{}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	} else {
		cDes.WeightedBackendService = des.WeightedBackendService
	}
	cDes.UrlRewrite = canonicalizeUrlMapDefaultRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	cDes.Timeout = canonicalizeUrlMapDefaultRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	cDes.RetryPolicy = canonicalizeUrlMapDefaultRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	cDes.RequestMirrorPolicy = canonicalizeUrlMapDefaultRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	cDes.CorsPolicy = canonicalizeUrlMapDefaultRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	cDes.FaultInjectionPolicy = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteAction(c *Client, des, nw *UrlMapDefaultRouteAction) *UrlMapDefaultRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.WeightedBackendService = canonicalizeNewUrlMapDefaultRouteActionWeightedBackendServiceSlice(c, des.WeightedBackendService, nw.WeightedBackendService)
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
			if diffs, _ := compareUrlMapDefaultRouteActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionSlice(c *Client, des, nw []UrlMapDefaultRouteAction) []UrlMapDefaultRouteAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteAction(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionWeightedBackendService(des, initial *UrlMapDefaultRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionWeightedBackendService{}

	if dcl.NameToSelfLink(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	cDes.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionWeightedBackendService(c *Client, des, nw *UrlMapDefaultRouteActionWeightedBackendService) *UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(nw.Weight) {
		nw.Weight = des.Weight
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
			if diffs, _ := compareUrlMapDefaultRouteActionWeightedBackendServiceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionWeightedBackendServiceSlice(c *Client, des, nw []UrlMapDefaultRouteActionWeightedBackendService) []UrlMapDefaultRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionWeightedBackendService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionWeightedBackendService(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapHeaderAction(des, initial *UrlMapHeaderAction, opts ...dcl.ApplyOption) *UrlMapHeaderAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapHeaderAction{}

	if dcl.IsZeroValue(des.RequestHeadersToRemove) {
		des.RequestHeadersToRemove = initial.RequestHeadersToRemove
	} else {
		cDes.RequestHeadersToRemove = des.RequestHeadersToRemove
	}
	if dcl.IsZeroValue(des.RequestHeadersToAdd) {
		des.RequestHeadersToAdd = initial.RequestHeadersToAdd
	} else {
		cDes.RequestHeadersToAdd = des.RequestHeadersToAdd
	}
	if dcl.IsZeroValue(des.ResponseHeadersToRemove) {
		des.ResponseHeadersToRemove = initial.ResponseHeadersToRemove
	} else {
		cDes.ResponseHeadersToRemove = des.ResponseHeadersToRemove
	}
	if dcl.IsZeroValue(des.ResponseHeadersToAdd) {
		des.ResponseHeadersToAdd = initial.ResponseHeadersToAdd
	} else {
		cDes.ResponseHeadersToAdd = des.ResponseHeadersToAdd
	}

	return cDes
}

func canonicalizeNewUrlMapHeaderAction(c *Client, des, nw *UrlMapHeaderAction) *UrlMapHeaderAction {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RequestHeadersToRemove) {
		nw.RequestHeadersToRemove = des.RequestHeadersToRemove
	}
	nw.RequestHeadersToAdd = canonicalizeNewUrlMapHeaderActionRequestHeadersToAddSlice(c, des.RequestHeadersToAdd, nw.RequestHeadersToAdd)
	if dcl.IsZeroValue(nw.ResponseHeadersToRemove) {
		nw.ResponseHeadersToRemove = des.ResponseHeadersToRemove
	}
	nw.ResponseHeadersToAdd = canonicalizeNewUrlMapHeaderActionResponseHeadersToAddSlice(c, des.ResponseHeadersToAdd, nw.ResponseHeadersToAdd)

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
			if diffs, _ := compareUrlMapHeaderActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapHeaderActionSlice(c *Client, des, nw []UrlMapHeaderAction) []UrlMapHeaderAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapHeaderAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapHeaderAction(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapHeaderActionRequestHeadersToAdd(des, initial *UrlMapHeaderActionRequestHeadersToAdd, opts ...dcl.ApplyOption) *UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapHeaderActionRequestHeadersToAdd{}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		cDes.HeaderName = initial.HeaderName
	} else {
		cDes.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.HeaderValue, initial.HeaderValue) || dcl.IsZeroValue(des.HeaderValue) {
		cDes.HeaderValue = initial.HeaderValue
	} else {
		cDes.HeaderValue = des.HeaderValue
	}
	if dcl.BoolCanonicalize(des.Replace, initial.Replace) || dcl.IsZeroValue(des.Replace) {
		cDes.Replace = initial.Replace
	} else {
		cDes.Replace = des.Replace
	}

	return cDes
}

func canonicalizeNewUrlMapHeaderActionRequestHeadersToAdd(c *Client, des, nw *UrlMapHeaderActionRequestHeadersToAdd) *UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.HeaderValue, nw.HeaderValue) {
		nw.HeaderValue = des.HeaderValue
	}
	if dcl.BoolCanonicalize(des.Replace, nw.Replace) {
		nw.Replace = des.Replace
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
			if diffs, _ := compareUrlMapHeaderActionRequestHeadersToAddNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapHeaderActionRequestHeadersToAddSlice(c *Client, des, nw []UrlMapHeaderActionRequestHeadersToAdd) []UrlMapHeaderActionRequestHeadersToAdd {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapHeaderActionRequestHeadersToAdd
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapHeaderActionRequestHeadersToAdd(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapHeaderActionResponseHeadersToAdd(des, initial *UrlMapHeaderActionResponseHeadersToAdd, opts ...dcl.ApplyOption) *UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapHeaderActionResponseHeadersToAdd{}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		cDes.HeaderName = initial.HeaderName
	} else {
		cDes.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.HeaderValue, initial.HeaderValue) || dcl.IsZeroValue(des.HeaderValue) {
		cDes.HeaderValue = initial.HeaderValue
	} else {
		cDes.HeaderValue = des.HeaderValue
	}
	if dcl.BoolCanonicalize(des.Replace, initial.Replace) || dcl.IsZeroValue(des.Replace) {
		cDes.Replace = initial.Replace
	} else {
		cDes.Replace = des.Replace
	}

	return cDes
}

func canonicalizeNewUrlMapHeaderActionResponseHeadersToAdd(c *Client, des, nw *UrlMapHeaderActionResponseHeadersToAdd) *UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.HeaderValue, nw.HeaderValue) {
		nw.HeaderValue = des.HeaderValue
	}
	if dcl.BoolCanonicalize(des.Replace, nw.Replace) {
		nw.Replace = des.Replace
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
			if diffs, _ := compareUrlMapHeaderActionResponseHeadersToAddNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapHeaderActionResponseHeadersToAddSlice(c *Client, des, nw []UrlMapHeaderActionResponseHeadersToAdd) []UrlMapHeaderActionResponseHeadersToAdd {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapHeaderActionResponseHeadersToAdd
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapHeaderActionResponseHeadersToAdd(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionUrlRewrite(des, initial *UrlMapDefaultRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionUrlRewrite{}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, initial.PathPrefixRewrite) || dcl.IsZeroValue(des.PathPrefixRewrite) {
		cDes.PathPrefixRewrite = initial.PathPrefixRewrite
	} else {
		cDes.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, initial.HostRewrite) || dcl.IsZeroValue(des.HostRewrite) {
		cDes.HostRewrite = initial.HostRewrite
	} else {
		cDes.HostRewrite = des.HostRewrite
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionUrlRewrite(c *Client, des, nw *UrlMapDefaultRouteActionUrlRewrite) *UrlMapDefaultRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, nw.PathPrefixRewrite) {
		nw.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, nw.HostRewrite) {
		nw.HostRewrite = des.HostRewrite
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
			if diffs, _ := compareUrlMapDefaultRouteActionUrlRewriteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionUrlRewriteSlice(c *Client, des, nw []UrlMapDefaultRouteActionUrlRewrite) []UrlMapDefaultRouteActionUrlRewrite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionUrlRewrite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionUrlRewrite(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionTimeout(des, initial *UrlMapDefaultRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionTimeout{}

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

func canonicalizeNewUrlMapDefaultRouteActionTimeout(c *Client, des, nw *UrlMapDefaultRouteActionTimeout) *UrlMapDefaultRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapDefaultRouteActionTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionTimeoutSlice(c *Client, des, nw []UrlMapDefaultRouteActionTimeout) []UrlMapDefaultRouteActionTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionRetryPolicy(des, initial *UrlMapDefaultRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionRetryPolicy{}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	} else {
		cDes.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	} else {
		cDes.NumRetries = des.NumRetries
	}
	cDes.PerTryTimeout = canonicalizeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicy(c *Client, des, nw *UrlMapDefaultRouteActionRetryPolicy) *UrlMapDefaultRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RetryCondition) {
		nw.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(nw.NumRetries) {
		nw.NumRetries = des.NumRetries
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
			if diffs, _ := compareUrlMapDefaultRouteActionRetryPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicySlice(c *Client, des, nw []UrlMapDefaultRouteActionRetryPolicy) []UrlMapDefaultRouteActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionRetryPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionRetryPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{}

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

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, des, nw []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionRetryPolicyPerTryTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionRequestMirrorPolicy(des, initial *UrlMapDefaultRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionRequestMirrorPolicy{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapDefaultRouteActionRequestMirrorPolicy) *UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
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
			if diffs, _ := compareUrlMapDefaultRouteActionRequestMirrorPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicySlice(c *Client, des, nw []UrlMapDefaultRouteActionRequestMirrorPolicy) []UrlMapDefaultRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionRequestMirrorPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionRequestMirrorPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionCorsPolicy(des, initial *UrlMapDefaultRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionCorsPolicy{}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	} else {
		cDes.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	} else {
		cDes.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	} else {
		cDes.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	} else {
		cDes.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	} else {
		cDes.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	} else {
		cDes.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, initial.AllowCredentials) || dcl.IsZeroValue(des.AllowCredentials) {
		cDes.AllowCredentials = initial.AllowCredentials
	} else {
		cDes.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		cDes.Disabled = initial.Disabled
	} else {
		cDes.Disabled = des.Disabled
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionCorsPolicy(c *Client, des, nw *UrlMapDefaultRouteActionCorsPolicy) *UrlMapDefaultRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowOrigin) {
		nw.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(nw.AllowOriginRegex) {
		nw.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(nw.AllowMethod) {
		nw.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(nw.AllowHeader) {
		nw.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(nw.ExposeHeader) {
		nw.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(nw.MaxAge) {
		nw.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, nw.AllowCredentials) {
		nw.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
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
			if diffs, _ := compareUrlMapDefaultRouteActionCorsPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionCorsPolicySlice(c *Client, des, nw []UrlMapDefaultRouteActionCorsPolicy) []UrlMapDefaultRouteActionCorsPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionCorsPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionCorsPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicy(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionFaultInjectionPolicy{}

	cDes.Delay = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	cDes.Abort = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return cDes
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
			if diffs, _ := compareUrlMapDefaultRouteActionFaultInjectionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicySlice(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicy) []UrlMapDefaultRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionFaultInjectionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionFaultInjectionPolicyDelay{}

	cDes.FixedDelay = canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) *UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
	}

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
			if diffs, _ := compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelaySlice(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyDelay) []UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionFaultInjectionPolicyDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{}

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

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapDefaultRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultRouteActionFaultInjectionPolicyAbort{}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	} else {
		cDes.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) *UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.HttpStatus) {
		nw.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
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
			if diffs, _ := compareUrlMapDefaultRouteActionFaultInjectionPolicyAbortNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbortSlice(c *Client, des, nw []UrlMapDefaultRouteActionFaultInjectionPolicyAbort) []UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultRouteActionFaultInjectionPolicyAbort
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultRouteActionFaultInjectionPolicyAbort(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapDefaultUrlRedirect(des, initial *UrlMapDefaultUrlRedirect, opts ...dcl.ApplyOption) *UrlMapDefaultUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapDefaultUrlRedirect{}

	if dcl.StringCanonicalize(des.HostRedirect, initial.HostRedirect) || dcl.IsZeroValue(des.HostRedirect) {
		cDes.HostRedirect = initial.HostRedirect
	} else {
		cDes.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, initial.PathRedirect) || dcl.IsZeroValue(des.PathRedirect) {
		cDes.PathRedirect = initial.PathRedirect
	} else {
		cDes.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, initial.PrefixRedirect) || dcl.IsZeroValue(des.PrefixRedirect) {
		cDes.PrefixRedirect = initial.PrefixRedirect
	} else {
		cDes.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	} else {
		cDes.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, initial.HttpsRedirect) || dcl.IsZeroValue(des.HttpsRedirect) {
		cDes.HttpsRedirect = initial.HttpsRedirect
	} else {
		cDes.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, initial.StripQuery) || dcl.IsZeroValue(des.StripQuery) {
		cDes.StripQuery = initial.StripQuery
	} else {
		cDes.StripQuery = des.StripQuery
	}

	return cDes
}

func canonicalizeNewUrlMapDefaultUrlRedirect(c *Client, des, nw *UrlMapDefaultUrlRedirect) *UrlMapDefaultUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HostRedirect, nw.HostRedirect) {
		nw.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, nw.PathRedirect) {
		nw.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, nw.PrefixRedirect) {
		nw.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(nw.RedirectResponseCode) {
		nw.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, nw.HttpsRedirect) {
		nw.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, nw.StripQuery) {
		nw.StripQuery = des.StripQuery
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
			if diffs, _ := compareUrlMapDefaultUrlRedirectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapDefaultUrlRedirectSlice(c *Client, des, nw []UrlMapDefaultUrlRedirect) []UrlMapDefaultUrlRedirect {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapDefaultUrlRedirect
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapDefaultUrlRedirect(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapHostRule(des, initial *UrlMapHostRule, opts ...dcl.ApplyOption) *UrlMapHostRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapHostRule{}

	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	} else {
		cDes.Host = des.Host
	}
	if dcl.StringCanonicalize(des.PathMatcher, initial.PathMatcher) || dcl.IsZeroValue(des.PathMatcher) {
		cDes.PathMatcher = initial.PathMatcher
	} else {
		cDes.PathMatcher = des.PathMatcher
	}

	return cDes
}

func canonicalizeNewUrlMapHostRule(c *Client, des, nw *UrlMapHostRule) *UrlMapHostRule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.IsZeroValue(nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.PathMatcher, nw.PathMatcher) {
		nw.PathMatcher = des.PathMatcher
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
			if diffs, _ := compareUrlMapHostRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapHostRuleSlice(c *Client, des, nw []UrlMapHostRule) []UrlMapHostRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapHostRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapHostRule(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcher(des, initial *UrlMapPathMatcher, opts ...dcl.ApplyOption) *UrlMapPathMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcher{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.StringCanonicalize(des.DefaultService, initial.DefaultService) || dcl.IsZeroValue(des.DefaultService) {
		cDes.DefaultService = initial.DefaultService
	} else {
		cDes.DefaultService = des.DefaultService
	}
	cDes.DefaultRouteAction = canonicalizeUrlMapDefaultRouteAction(des.DefaultRouteAction, initial.DefaultRouteAction, opts...)
	cDes.DefaultUrlRedirect = canonicalizeUrlMapPathMatcherDefaultUrlRedirect(des.DefaultUrlRedirect, initial.DefaultUrlRedirect, opts...)
	if dcl.IsZeroValue(des.PathRule) {
		des.PathRule = initial.PathRule
	} else {
		cDes.PathRule = des.PathRule
	}
	if dcl.IsZeroValue(des.RouteRule) {
		des.RouteRule = initial.RouteRule
	} else {
		cDes.RouteRule = des.RouteRule
	}
	cDes.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcher(c *Client, des, nw *UrlMapPathMatcher) *UrlMapPathMatcher {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.DefaultService, nw.DefaultService) {
		nw.DefaultService = des.DefaultService
	}
	nw.DefaultRouteAction = canonicalizeNewUrlMapDefaultRouteAction(c, des.DefaultRouteAction, nw.DefaultRouteAction)
	nw.DefaultUrlRedirect = canonicalizeNewUrlMapPathMatcherDefaultUrlRedirect(c, des.DefaultUrlRedirect, nw.DefaultUrlRedirect)
	nw.PathRule = canonicalizeNewUrlMapPathMatcherPathRuleSlice(c, des.PathRule, nw.PathRule)
	nw.RouteRule = canonicalizeNewUrlMapPathMatcherRouteRuleSlice(c, des.RouteRule, nw.RouteRule)
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
			if diffs, _ := compareUrlMapPathMatcherNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherSlice(c *Client, des, nw []UrlMapPathMatcher) []UrlMapPathMatcher {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcher
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcher(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherDefaultUrlRedirect(des, initial *UrlMapPathMatcherDefaultUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherDefaultUrlRedirect{}

	if dcl.StringCanonicalize(des.HostRedirect, initial.HostRedirect) || dcl.IsZeroValue(des.HostRedirect) {
		cDes.HostRedirect = initial.HostRedirect
	} else {
		cDes.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, initial.PathRedirect) || dcl.IsZeroValue(des.PathRedirect) {
		cDes.PathRedirect = initial.PathRedirect
	} else {
		cDes.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, initial.PrefixRedirect) || dcl.IsZeroValue(des.PrefixRedirect) {
		cDes.PrefixRedirect = initial.PrefixRedirect
	} else {
		cDes.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	} else {
		cDes.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, initial.HttpsRedirect) || dcl.IsZeroValue(des.HttpsRedirect) {
		cDes.HttpsRedirect = initial.HttpsRedirect
	} else {
		cDes.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, initial.StripQuery) || dcl.IsZeroValue(des.StripQuery) {
		cDes.StripQuery = initial.StripQuery
	} else {
		cDes.StripQuery = des.StripQuery
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherDefaultUrlRedirect(c *Client, des, nw *UrlMapPathMatcherDefaultUrlRedirect) *UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HostRedirect, nw.HostRedirect) {
		nw.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, nw.PathRedirect) {
		nw.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, nw.PrefixRedirect) {
		nw.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(nw.RedirectResponseCode) {
		nw.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, nw.HttpsRedirect) {
		nw.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, nw.StripQuery) {
		nw.StripQuery = des.StripQuery
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
			if diffs, _ := compareUrlMapPathMatcherDefaultUrlRedirectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherDefaultUrlRedirectSlice(c *Client, des, nw []UrlMapPathMatcherDefaultUrlRedirect) []UrlMapPathMatcherDefaultUrlRedirect {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherDefaultUrlRedirect
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherDefaultUrlRedirect(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRule(des, initial *UrlMapPathMatcherPathRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRule{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}
	cDes.RouteAction = canonicalizeUrlMapPathMatcherPathRuleRouteAction(des.RouteAction, initial.RouteAction, opts...)
	cDes.UrlRedirect = canonicalizeUrlMapPathMatcherPathRuleUrlRedirect(des.UrlRedirect, initial.UrlRedirect, opts...)
	if dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRule(c *Client, des, nw *UrlMapPathMatcherPathRule) *UrlMapPathMatcherPathRule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
	}
	nw.RouteAction = canonicalizeNewUrlMapPathMatcherPathRuleRouteAction(c, des.RouteAction, nw.RouteAction)
	nw.UrlRedirect = canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirect(c, des.UrlRedirect, nw.UrlRedirect)
	if dcl.IsZeroValue(nw.Path) {
		nw.Path = des.Path
	}

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
			if diffs, _ := compareUrlMapPathMatcherPathRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleSlice(c *Client, des, nw []UrlMapPathMatcherPathRule) []UrlMapPathMatcherPathRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRule(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteAction(des, initial *UrlMapPathMatcherPathRuleRouteAction, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteAction{}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	} else {
		cDes.WeightedBackendService = des.WeightedBackendService
	}
	cDes.UrlRewrite = canonicalizeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	cDes.Timeout = canonicalizeUrlMapPathMatcherPathRuleRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	cDes.RetryPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	cDes.RequestMirrorPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	cDes.CorsPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	cDes.FaultInjectionPolicy = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteAction(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteAction) *UrlMapPathMatcherPathRuleRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.WeightedBackendService = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c, des.WeightedBackendService, nw.WeightedBackendService)
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteAction) []UrlMapPathMatcherPathRuleRouteAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteAction(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(des, initial *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	cDes.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(nw.Weight) {
		nw.Weight = des.Weight
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(des, initial *UrlMapPathMatcherPathRuleRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionUrlRewrite{}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, initial.PathPrefixRewrite) || dcl.IsZeroValue(des.PathPrefixRewrite) {
		cDes.PathPrefixRewrite = initial.PathPrefixRewrite
	} else {
		cDes.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, initial.HostRewrite) || dcl.IsZeroValue(des.HostRewrite) {
		cDes.HostRewrite = initial.HostRewrite
	} else {
		cDes.HostRewrite = des.HostRewrite
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) *UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, nw.PathPrefixRewrite) {
		nw.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, nw.HostRewrite) {
		nw.HostRewrite = des.HostRewrite
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionUrlRewriteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewriteSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionUrlRewrite) []UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionUrlRewrite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionUrlRewrite(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionTimeout(des, initial *UrlMapPathMatcherPathRuleRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionTimeout{}

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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeout(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionTimeout) *UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeoutSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionTimeout) []UrlMapPathMatcherPathRuleRouteActionTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionRetryPolicy{}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	} else {
		cDes.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	} else {
		cDes.NumRetries = des.NumRetries
	}
	cDes.PerTryTimeout = canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) *UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RetryCondition) {
		nw.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(nw.NumRetries) {
		nw.NumRetries = des.NumRetries
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRetryPolicy) []UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionRetryPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{}

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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionCorsPolicy{}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	} else {
		cDes.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	} else {
		cDes.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	} else {
		cDes.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	} else {
		cDes.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	} else {
		cDes.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	} else {
		cDes.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, initial.AllowCredentials) || dcl.IsZeroValue(des.AllowCredentials) {
		cDes.AllowCredentials = initial.AllowCredentials
	} else {
		cDes.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		cDes.Disabled = initial.Disabled
	} else {
		cDes.Disabled = des.Disabled
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) *UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowOrigin) {
		nw.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(nw.AllowOriginRegex) {
		nw.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(nw.AllowMethod) {
		nw.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(nw.AllowHeader) {
		nw.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(nw.ExposeHeader) {
		nw.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(nw.MaxAge) {
		nw.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, nw.AllowCredentials) {
		nw.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionCorsPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionCorsPolicy) []UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionCorsPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionCorsPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{}

	cDes.Delay = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	cDes.Abort = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return cDes
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{}

	cDes.FixedDelay = canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
	}

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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}

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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	} else {
		cDes.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.HttpStatus) {
		nw.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherPathRuleUrlRedirect(des, initial *UrlMapPathMatcherPathRuleUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherPathRuleUrlRedirect{}

	if dcl.StringCanonicalize(des.HostRedirect, initial.HostRedirect) || dcl.IsZeroValue(des.HostRedirect) {
		cDes.HostRedirect = initial.HostRedirect
	} else {
		cDes.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, initial.PathRedirect) || dcl.IsZeroValue(des.PathRedirect) {
		cDes.PathRedirect = initial.PathRedirect
	} else {
		cDes.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, initial.PrefixRedirect) || dcl.IsZeroValue(des.PrefixRedirect) {
		cDes.PrefixRedirect = initial.PrefixRedirect
	} else {
		cDes.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	} else {
		cDes.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, initial.HttpsRedirect) || dcl.IsZeroValue(des.HttpsRedirect) {
		cDes.HttpsRedirect = initial.HttpsRedirect
	} else {
		cDes.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, initial.StripQuery) || dcl.IsZeroValue(des.StripQuery) {
		cDes.StripQuery = initial.StripQuery
	} else {
		cDes.StripQuery = des.StripQuery
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirect(c *Client, des, nw *UrlMapPathMatcherPathRuleUrlRedirect) *UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HostRedirect, nw.HostRedirect) {
		nw.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, nw.PathRedirect) {
		nw.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, nw.PrefixRedirect) {
		nw.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(nw.RedirectResponseCode) {
		nw.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, nw.HttpsRedirect) {
		nw.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, nw.StripQuery) {
		nw.StripQuery = des.StripQuery
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
			if diffs, _ := compareUrlMapPathMatcherPathRuleUrlRedirectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirectSlice(c *Client, des, nw []UrlMapPathMatcherPathRuleUrlRedirect) []UrlMapPathMatcherPathRuleUrlRedirect {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherPathRuleUrlRedirect
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherPathRuleUrlRedirect(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRule(des, initial *UrlMapPathMatcherRouteRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRule{}

	if dcl.IsZeroValue(des.Priority) {
		des.Priority = initial.Priority
	} else {
		cDes.Priority = des.Priority
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.IsZeroValue(des.MatchRule) {
		des.MatchRule = initial.MatchRule
	} else {
		cDes.MatchRule = des.MatchRule
	}
	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}
	cDes.RouteAction = canonicalizeUrlMapPathMatcherRouteRuleRouteAction(des.RouteAction, initial.RouteAction, opts...)
	cDes.UrlRedirect = canonicalizeUrlMapPathMatcherRouteRuleUrlRedirect(des.UrlRedirect, initial.UrlRedirect, opts...)
	cDes.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRule(c *Client, des, nw *UrlMapPathMatcherRouteRule) *UrlMapPathMatcherRouteRule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Priority) {
		nw.Priority = des.Priority
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	nw.MatchRule = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleSlice(c, des.MatchRule, nw.MatchRule)
	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleSlice(c *Client, des, nw []UrlMapPathMatcherRouteRule) []UrlMapPathMatcherRouteRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRule(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRule(des, initial *UrlMapPathMatcherRouteRuleMatchRule, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRule{}

	if dcl.StringCanonicalize(des.PrefixMatch, initial.PrefixMatch) || dcl.IsZeroValue(des.PrefixMatch) {
		cDes.PrefixMatch = initial.PrefixMatch
	} else {
		cDes.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.FullPathMatch, initial.FullPathMatch) || dcl.IsZeroValue(des.FullPathMatch) {
		cDes.FullPathMatch = initial.FullPathMatch
	} else {
		cDes.FullPathMatch = des.FullPathMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.IgnoreCase, initial.IgnoreCase) || dcl.IsZeroValue(des.IgnoreCase) {
		cDes.IgnoreCase = initial.IgnoreCase
	} else {
		cDes.IgnoreCase = des.IgnoreCase
	}
	if dcl.IsZeroValue(des.HeaderMatch) {
		des.HeaderMatch = initial.HeaderMatch
	} else {
		cDes.HeaderMatch = des.HeaderMatch
	}
	if dcl.IsZeroValue(des.QueryParameterMatch) {
		des.QueryParameterMatch = initial.QueryParameterMatch
	} else {
		cDes.QueryParameterMatch = des.QueryParameterMatch
	}
	if dcl.IsZeroValue(des.MetadataFilter) {
		des.MetadataFilter = initial.MetadataFilter
	} else {
		cDes.MetadataFilter = des.MetadataFilter
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRule(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRule) *UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PrefixMatch, nw.PrefixMatch) {
		nw.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.FullPathMatch, nw.FullPathMatch) {
		nw.FullPathMatch = des.FullPathMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.IgnoreCase, nw.IgnoreCase) {
		nw.IgnoreCase = des.IgnoreCase
	}
	nw.HeaderMatch = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c, des.HeaderMatch, nw.HeaderMatch)
	nw.QueryParameterMatch = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c, des.QueryParameterMatch, nw.QueryParameterMatch)
	nw.MetadataFilter = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c, des.MetadataFilter, nw.MetadataFilter)

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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRule) []UrlMapPathMatcherRouteRuleMatchRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRule(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		cDes.HeaderName = initial.HeaderName
	} else {
		cDes.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.ExactMatch, initial.ExactMatch) || dcl.IsZeroValue(des.ExactMatch) {
		cDes.ExactMatch = initial.ExactMatch
	} else {
		cDes.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}
	cDes.RangeMatch = canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(des.RangeMatch, initial.RangeMatch, opts...)
	if dcl.BoolCanonicalize(des.PresentMatch, initial.PresentMatch) || dcl.IsZeroValue(des.PresentMatch) {
		cDes.PresentMatch = initial.PresentMatch
	} else {
		cDes.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, initial.PrefixMatch) || dcl.IsZeroValue(des.PrefixMatch) {
		cDes.PrefixMatch = initial.PrefixMatch
	} else {
		cDes.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.SuffixMatch, initial.SuffixMatch) || dcl.IsZeroValue(des.SuffixMatch) {
		cDes.SuffixMatch = initial.SuffixMatch
	} else {
		cDes.SuffixMatch = des.SuffixMatch
	}
	if dcl.BoolCanonicalize(des.InvertMatch, initial.InvertMatch) || dcl.IsZeroValue(des.InvertMatch) {
		cDes.InvertMatch = initial.InvertMatch
	} else {
		cDes.InvertMatch = des.InvertMatch
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.ExactMatch, nw.ExactMatch) {
		nw.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}
	nw.RangeMatch = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, des.RangeMatch, nw.RangeMatch)
	if dcl.BoolCanonicalize(des.PresentMatch, nw.PresentMatch) {
		nw.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, nw.PrefixMatch) {
		nw.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.SuffixMatch, nw.SuffixMatch) {
		nw.SuffixMatch = des.SuffixMatch
	}
	if dcl.BoolCanonicalize(des.InvertMatch, nw.InvertMatch) {
		nw.InvertMatch = des.InvertMatch
	}

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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{}

	if dcl.IsZeroValue(des.RangeStart) {
		des.RangeStart = initial.RangeStart
	} else {
		cDes.RangeStart = des.RangeStart
	}
	if dcl.IsZeroValue(des.RangeEnd) {
		des.RangeEnd = initial.RangeEnd
	} else {
		cDes.RangeEnd = des.RangeEnd
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RangeStart) {
		nw.RangeStart = des.RangeStart
	}
	if dcl.IsZeroValue(nw.RangeEnd) {
		nw.RangeEnd = des.RangeEnd
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(des, initial *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.BoolCanonicalize(des.PresentMatch, initial.PresentMatch) || dcl.IsZeroValue(des.PresentMatch) {
		cDes.PresentMatch = initial.PresentMatch
	} else {
		cDes.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.ExactMatch, initial.ExactMatch) || dcl.IsZeroValue(des.ExactMatch) {
		cDes.ExactMatch = initial.ExactMatch
	} else {
		cDes.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.BoolCanonicalize(des.PresentMatch, nw.PresentMatch) {
		nw.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.ExactMatch, nw.ExactMatch) {
		nw.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(des, initial *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{}

	if dcl.IsZeroValue(des.FilterMatchCriteria) {
		des.FilterMatchCriteria = initial.FilterMatchCriteria
	} else {
		cDes.FilterMatchCriteria = des.FilterMatchCriteria
	}
	if dcl.IsZeroValue(des.FilterLabel) {
		des.FilterLabel = initial.FilterLabel
	} else {
		cDes.FilterLabel = des.FilterLabel
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.FilterMatchCriteria) {
		nw.FilterMatchCriteria = des.FilterMatchCriteria
	}
	nw.FilterLabel = canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c, des.FilterLabel, nw.FilterLabel)

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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(des, initial *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c *Client, des, nw *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteAction(des, initial *UrlMapPathMatcherRouteRuleRouteAction, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteAction{}

	if dcl.IsZeroValue(des.WeightedBackendService) {
		des.WeightedBackendService = initial.WeightedBackendService
	} else {
		cDes.WeightedBackendService = des.WeightedBackendService
	}
	cDes.UrlRewrite = canonicalizeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	cDes.Timeout = canonicalizeUrlMapPathMatcherRouteRuleRouteActionTimeout(des.Timeout, initial.Timeout, opts...)
	cDes.RetryPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	cDes.RequestMirrorPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	cDes.CorsPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)
	cDes.FaultInjectionPolicy = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteAction(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteAction) *UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil || nw == nil {
		return nw
	}

	nw.WeightedBackendService = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c, des.WeightedBackendService, nw.WeightedBackendService)
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteAction) []UrlMapPathMatcherRouteRuleRouteAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteAction(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(des, initial *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	cDes.HeaderAction = canonicalizeUrlMapHeaderAction(des.HeaderAction, initial.HeaderAction, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
	}
	if dcl.IsZeroValue(nw.Weight) {
		nw.Weight = des.Weight
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(des, initial *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, initial.PathPrefixRewrite) || dcl.IsZeroValue(des.PathPrefixRewrite) {
		cDes.PathPrefixRewrite = initial.PathPrefixRewrite
	} else {
		cDes.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, initial.HostRewrite) || dcl.IsZeroValue(des.HostRewrite) {
		cDes.HostRewrite = initial.HostRewrite
	} else {
		cDes.HostRewrite = des.HostRewrite
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, nw.PathPrefixRewrite) {
		nw.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, nw.HostRewrite) {
		nw.HostRewrite = des.HostRewrite
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionUrlRewriteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewriteSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionUrlRewrite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionTimeout(des, initial *UrlMapPathMatcherRouteRuleRouteActionTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionTimeout{}

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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeout(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionTimeout) *UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeoutSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionTimeout) []UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{}

	if dcl.IsZeroValue(des.RetryCondition) {
		des.RetryCondition = initial.RetryCondition
	} else {
		cDes.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(des.NumRetries) {
		des.NumRetries = initial.NumRetries
	} else {
		cDes.NumRetries = des.NumRetries
	}
	cDes.PerTryTimeout = canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(des.PerTryTimeout, initial.PerTryTimeout, opts...)

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RetryCondition) {
		nw.RetryCondition = des.RetryCondition
	}
	if dcl.IsZeroValue(nw.NumRetries) {
		nw.NumRetries = des.NumRetries
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionRetryPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(des, initial *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{}

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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{}

	if dcl.StringCanonicalize(des.BackendService, initial.BackendService) || dcl.IsZeroValue(des.BackendService) {
		cDes.BackendService = initial.BackendService
	} else {
		cDes.BackendService = des.BackendService
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BackendService, nw.BackendService) {
		nw.BackendService = des.BackendService
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{}

	if dcl.IsZeroValue(des.AllowOrigin) {
		des.AllowOrigin = initial.AllowOrigin
	} else {
		cDes.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(des.AllowOriginRegex) {
		des.AllowOriginRegex = initial.AllowOriginRegex
	} else {
		cDes.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(des.AllowMethod) {
		des.AllowMethod = initial.AllowMethod
	} else {
		cDes.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(des.AllowHeader) {
		des.AllowHeader = initial.AllowHeader
	} else {
		cDes.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(des.ExposeHeader) {
		des.ExposeHeader = initial.ExposeHeader
	} else {
		cDes.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(des.MaxAge) {
		des.MaxAge = initial.MaxAge
	} else {
		cDes.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, initial.AllowCredentials) || dcl.IsZeroValue(des.AllowCredentials) {
		cDes.AllowCredentials = initial.AllowCredentials
	} else {
		cDes.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		cDes.Disabled = initial.Disabled
	} else {
		cDes.Disabled = des.Disabled
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowOrigin) {
		nw.AllowOrigin = des.AllowOrigin
	}
	if dcl.IsZeroValue(nw.AllowOriginRegex) {
		nw.AllowOriginRegex = des.AllowOriginRegex
	}
	if dcl.IsZeroValue(nw.AllowMethod) {
		nw.AllowMethod = des.AllowMethod
	}
	if dcl.IsZeroValue(nw.AllowHeader) {
		nw.AllowHeader = des.AllowHeader
	}
	if dcl.IsZeroValue(nw.ExposeHeader) {
		nw.ExposeHeader = des.ExposeHeader
	}
	if dcl.IsZeroValue(nw.MaxAge) {
		nw.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, nw.AllowCredentials) {
		nw.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionCorsPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{}

	cDes.Delay = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	cDes.Abort = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return cDes
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{}

	cDes.FixedDelay = canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des.FixedDelay, initial.FixedDelay, opts...)
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil || nw == nil {
		return nw
	}

	nw.FixedDelay = canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, des.FixedDelay, nw.FixedDelay)
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
	}

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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelaySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{}

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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelaySlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(des, initial *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{}

	if dcl.IsZeroValue(des.HttpStatus) {
		des.HttpStatus = initial.HttpStatus
	} else {
		cDes.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) {
		des.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c *Client, des, nw *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.HttpStatus) {
		nw.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(nw.Percentage) {
		nw.Percentage = des.Percentage
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapPathMatcherRouteRuleUrlRedirect(des, initial *UrlMapPathMatcherRouteRuleUrlRedirect, opts ...dcl.ApplyOption) *UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapPathMatcherRouteRuleUrlRedirect{}

	if dcl.StringCanonicalize(des.HostRedirect, initial.HostRedirect) || dcl.IsZeroValue(des.HostRedirect) {
		cDes.HostRedirect = initial.HostRedirect
	} else {
		cDes.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, initial.PathRedirect) || dcl.IsZeroValue(des.PathRedirect) {
		cDes.PathRedirect = initial.PathRedirect
	} else {
		cDes.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, initial.PrefixRedirect) || dcl.IsZeroValue(des.PrefixRedirect) {
		cDes.PrefixRedirect = initial.PrefixRedirect
	} else {
		cDes.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(des.RedirectResponseCode) {
		des.RedirectResponseCode = initial.RedirectResponseCode
	} else {
		cDes.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, initial.HttpsRedirect) || dcl.IsZeroValue(des.HttpsRedirect) {
		cDes.HttpsRedirect = initial.HttpsRedirect
	} else {
		cDes.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, initial.StripQuery) || dcl.IsZeroValue(des.StripQuery) {
		cDes.StripQuery = initial.StripQuery
	} else {
		cDes.StripQuery = des.StripQuery
	}

	return cDes
}

func canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirect(c *Client, des, nw *UrlMapPathMatcherRouteRuleUrlRedirect) *UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HostRedirect, nw.HostRedirect) {
		nw.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, nw.PathRedirect) {
		nw.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRedirect, nw.PrefixRedirect) {
		nw.PrefixRedirect = des.PrefixRedirect
	}
	if dcl.IsZeroValue(nw.RedirectResponseCode) {
		nw.RedirectResponseCode = des.RedirectResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, nw.HttpsRedirect) {
		nw.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, nw.StripQuery) {
		nw.StripQuery = des.StripQuery
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
			if diffs, _ := compareUrlMapPathMatcherRouteRuleUrlRedirectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirectSlice(c *Client, des, nw []UrlMapPathMatcherRouteRuleUrlRedirect) []UrlMapPathMatcherRouteRuleUrlRedirect {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapPathMatcherRouteRuleUrlRedirect
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapPathMatcherRouteRuleUrlRedirect(c, &d, &n))
	}

	return items
}

func canonicalizeUrlMapTest(des, initial *UrlMapTest, opts ...dcl.ApplyOption) *UrlMapTest {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &UrlMapTest{}

	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		cDes.Host = initial.Host
	} else {
		cDes.Host = des.Host
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.StringCanonicalize(des.ExpectedBackendService, initial.ExpectedBackendService) || dcl.IsZeroValue(des.ExpectedBackendService) {
		cDes.ExpectedBackendService = initial.ExpectedBackendService
	} else {
		cDes.ExpectedBackendService = des.ExpectedBackendService
	}

	return cDes
}

func canonicalizeNewUrlMapTest(c *Client, des, nw *UrlMapTest) *UrlMapTest {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.ExpectedBackendService, nw.ExpectedBackendService) {
		nw.ExpectedBackendService = des.ExpectedBackendService
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
			if diffs, _ := compareUrlMapTestNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewUrlMapTestSlice(c *Client, des, nw []UrlMapTest) []UrlMapTest {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []UrlMapTest
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUrlMapTest(c, &d, &n))
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
func diffUrlMap(c *Client, desired, actual *UrlMap, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.DefaultRouteAction, actual.DefaultRouteAction, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionNewStyle, EmptyObject: EmptyUrlMapDefaultRouteAction, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("DefaultRouteAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultService, actual.DefaultService, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("DefaultService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultUrlRedirect, actual.DefaultUrlRedirect, dcl.Info{ObjectFunction: compareUrlMapDefaultUrlRedirectNewStyle, EmptyObject: EmptyUrlMapDefaultUrlRedirect, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("DefaultUrlRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HostRule, actual.HostRule, dcl.Info{Type: "Set", ObjectFunction: compareUrlMapHostRuleNewStyle, EmptyObject: EmptyUrlMapHostRule, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("HostRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathMatcher, actual.PathMatcher, dcl.Info{Type: "Set", ObjectFunction: compareUrlMapPathMatcherNewStyle, EmptyObject: EmptyUrlMapPathMatcher, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("PathMatchers")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Test, actual.Test, dcl.Info{ObjectFunction: compareUrlMapTestNewStyle, EmptyObject: EmptyUrlMapTest, OperationSelector: dcl.TriggersOperation("updateUrlMapUpdateOperation")}, fn.AddNest("Tests")); len(ds) != 0 || err != nil {
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
func compareUrlMapDefaultRouteActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteAction)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteAction or *UrlMapDefaultRouteAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteAction)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WeightedBackendService, actual.WeightedBackendService, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionWeightedBackendServiceNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionWeightedBackendService, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WeightedBackendServices")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRewrite, actual.UrlRewrite, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionUrlRewriteNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionUrlRewrite, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionTimeoutNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryPolicy, actual.RetryPolicy, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionRetryPolicyNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionRetryPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestMirrorPolicy, actual.RequestMirrorPolicy, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionRequestMirrorPolicyNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionRequestMirrorPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestMirrorPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CorsPolicy, actual.CorsPolicy, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionCorsPolicyNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionCorsPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CorsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FaultInjectionPolicy, actual.FaultInjectionPolicy, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionFaultInjectionPolicyNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionFaultInjectionPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FaultInjectionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionWeightedBackendServiceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionWeightedBackendService)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionWeightedBackendService or *UrlMapDefaultRouteActionWeightedBackendService", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionWeightedBackendService)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionWeightedBackendService", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapHeaderActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapHeaderAction)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapHeaderAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderAction or *UrlMapHeaderAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapHeaderAction)
	if !ok {
		actualNotPointer, ok := a.(UrlMapHeaderAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RequestHeadersToRemove, actual.RequestHeadersToRemove, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestHeadersToRemove")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestHeadersToAdd, actual.RequestHeadersToAdd, dcl.Info{ObjectFunction: compareUrlMapHeaderActionRequestHeadersToAddNewStyle, EmptyObject: EmptyUrlMapHeaderActionRequestHeadersToAdd, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestHeadersToAdd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseHeadersToRemove, actual.ResponseHeadersToRemove, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResponseHeadersToRemove")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseHeadersToAdd, actual.ResponseHeadersToAdd, dcl.Info{ObjectFunction: compareUrlMapHeaderActionResponseHeadersToAddNewStyle, EmptyObject: EmptyUrlMapHeaderActionResponseHeadersToAdd, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResponseHeadersToAdd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapHeaderActionRequestHeadersToAddNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapHeaderActionRequestHeadersToAdd)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapHeaderActionRequestHeadersToAdd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderActionRequestHeadersToAdd or *UrlMapHeaderActionRequestHeadersToAdd", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapHeaderActionRequestHeadersToAdd)
	if !ok {
		actualNotPointer, ok := a.(UrlMapHeaderActionRequestHeadersToAdd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderActionRequestHeadersToAdd", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HeaderName, actual.HeaderName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderValue, actual.HeaderValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Replace, actual.Replace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Replace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapHeaderActionResponseHeadersToAddNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapHeaderActionResponseHeadersToAdd)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapHeaderActionResponseHeadersToAdd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderActionResponseHeadersToAdd or *UrlMapHeaderActionResponseHeadersToAdd", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapHeaderActionResponseHeadersToAdd)
	if !ok {
		actualNotPointer, ok := a.(UrlMapHeaderActionResponseHeadersToAdd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHeaderActionResponseHeadersToAdd", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HeaderName, actual.HeaderName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderValue, actual.HeaderValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Replace, actual.Replace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Replace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionUrlRewriteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionUrlRewrite)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionUrlRewrite or *UrlMapDefaultRouteActionUrlRewrite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionUrlRewrite)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionUrlRewrite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PathPrefixRewrite, actual.PathPrefixRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathPrefixRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HostRewrite, actual.HostRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionTimeout or *UrlMapDefaultRouteActionTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapDefaultRouteActionRetryPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionRetryPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRetryPolicy or *UrlMapDefaultRouteActionRetryPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionRetryPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRetryPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryCondition, actual.RetryCondition, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryConditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRetries, actual.NumRetries, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerTryTimeout, actual.PerTryTimeout, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionRetryPolicyPerTryTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerTryTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout or *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapDefaultRouteActionRequestMirrorPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionRequestMirrorPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRequestMirrorPolicy or *UrlMapDefaultRouteActionRequestMirrorPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionRequestMirrorPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionRequestMirrorPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionCorsPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionCorsPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionCorsPolicy or *UrlMapDefaultRouteActionCorsPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionCorsPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionCorsPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowOrigin, actual.AllowOrigin, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOrigins")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowOriginRegex, actual.AllowOriginRegex, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOriginRegexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowMethod, actual.AllowMethod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowMethods")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowHeader, actual.AllowHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExposeHeader, actual.ExposeHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExposeHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxAge, actual.MaxAge, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxAge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowCredentials, actual.AllowCredentials, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowCredentials")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionFaultInjectionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicy or *UrlMapDefaultRouteActionFaultInjectionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionFaultInjectionPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Delay, actual.Delay, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Delay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abort, actual.Abort, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionFaultInjectionPolicyAbortNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionFaultInjectionPolicyAbort, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Abort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionFaultInjectionPolicyDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyDelay or *UrlMapDefaultRouteActionFaultInjectionPolicyDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionFaultInjectionPolicyDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyDelay", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedDelay, actual.FixedDelay, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle, EmptyObject: EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FixedDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay or *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapDefaultRouteActionFaultInjectionPolicyAbortNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultRouteActionFaultInjectionPolicyAbort)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyAbort or *UrlMapDefaultRouteActionFaultInjectionPolicyAbort", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultRouteActionFaultInjectionPolicyAbort)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultRouteActionFaultInjectionPolicyAbort", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpStatus, actual.HttpStatus, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapDefaultUrlRedirectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapDefaultUrlRedirect)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapDefaultUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultUrlRedirect or *UrlMapDefaultUrlRedirect", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapDefaultUrlRedirect)
	if !ok {
		actualNotPointer, ok := a.(UrlMapDefaultUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapDefaultUrlRedirect", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostRedirect, actual.HostRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRedirect, actual.PathRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixRedirect, actual.PrefixRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedirectResponseCode, actual.RedirectResponseCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedirectResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsRedirect, actual.HttpsRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StripQuery, actual.StripQuery, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StripQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapHostRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapHostRule)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapHostRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHostRule or *UrlMapHostRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapHostRule)
	if !ok {
		actualNotPointer, ok := a.(UrlMapHostRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapHostRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Hosts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathMatcher, actual.PathMatcher, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathMatcher")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcher)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcher or *UrlMapPathMatcher", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcher)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcher", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DefaultService, actual.DefaultService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultRouteAction, actual.DefaultRouteAction, dcl.Info{ObjectFunction: compareUrlMapDefaultRouteActionNewStyle, EmptyObject: EmptyUrlMapDefaultRouteAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultRouteAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultUrlRedirect, actual.DefaultUrlRedirect, dcl.Info{ObjectFunction: compareUrlMapPathMatcherDefaultUrlRedirectNewStyle, EmptyObject: EmptyUrlMapPathMatcherDefaultUrlRedirect, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultUrlRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRule, actual.PathRule, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRule, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RouteRule, actual.RouteRule, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRule, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RouteRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherDefaultUrlRedirectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherDefaultUrlRedirect)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherDefaultUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherDefaultUrlRedirect or *UrlMapPathMatcherDefaultUrlRedirect", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherDefaultUrlRedirect)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherDefaultUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherDefaultUrlRedirect", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostRedirect, actual.HostRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRedirect, actual.PathRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixRedirect, actual.PrefixRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedirectResponseCode, actual.RedirectResponseCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedirectResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsRedirect, actual.HttpsRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StripQuery, actual.StripQuery, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StripQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRule)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRule or *UrlMapPathMatcherPathRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRule)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RouteAction, actual.RouteAction, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RouteAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRedirect, actual.UrlRedirect, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleUrlRedirectNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleUrlRedirect, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Paths")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteAction)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteAction or *UrlMapPathMatcherPathRuleRouteAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteAction)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WeightedBackendService, actual.WeightedBackendService, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionWeightedBackendService, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WeightedBackendServices")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRewrite, actual.UrlRewrite, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionUrlRewriteNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionUrlRewrite, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionTimeoutNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryPolicy, actual.RetryPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestMirrorPolicy, actual.RequestMirrorPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestMirrorPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CorsPolicy, actual.CorsPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionCorsPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionCorsPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CorsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FaultInjectionPolicy, actual.FaultInjectionPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FaultInjectionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionWeightedBackendService)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService or *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionWeightedBackendService)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionUrlRewriteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionUrlRewrite)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionUrlRewrite or *UrlMapPathMatcherPathRuleRouteActionUrlRewrite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionUrlRewrite)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionUrlRewrite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PathPrefixRewrite, actual.PathPrefixRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathPrefixRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HostRewrite, actual.HostRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionTimeout or *UrlMapPathMatcherPathRuleRouteActionTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionRetryPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRetryPolicy or *UrlMapPathMatcherPathRuleRouteActionRetryPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionRetryPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRetryPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryCondition, actual.RetryCondition, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryConditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRetries, actual.NumRetries, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerTryTimeout, actual.PerTryTimeout, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerTryTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout or *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy or *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionCorsPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionCorsPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionCorsPolicy or *UrlMapPathMatcherPathRuleRouteActionCorsPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionCorsPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionCorsPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowOrigin, actual.AllowOrigin, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOrigins")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowOriginRegex, actual.AllowOriginRegex, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOriginRegexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowMethod, actual.AllowMethod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowMethods")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowHeader, actual.AllowHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExposeHeader, actual.ExposeHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExposeHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxAge, actual.MaxAge, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxAge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowCredentials, actual.AllowCredentials, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowCredentials")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy or *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Delay, actual.Delay, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Delay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abort, actual.Abort, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Abort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay or *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedDelay, actual.FixedDelay, dcl.Info{ObjectFunction: compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle, EmptyObject: EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FixedDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay or *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort or *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpStatus, actual.HttpStatus, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherPathRuleUrlRedirectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherPathRuleUrlRedirect)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherPathRuleUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleUrlRedirect or *UrlMapPathMatcherPathRuleUrlRedirect", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherPathRuleUrlRedirect)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherPathRuleUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherPathRuleUrlRedirect", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostRedirect, actual.HostRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRedirect, actual.PathRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixRedirect, actual.PrefixRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedirectResponseCode, actual.RedirectResponseCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedirectResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsRedirect, actual.HttpsRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StripQuery, actual.StripQuery, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StripQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRule)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRule or *UrlMapPathMatcherRouteRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRule)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Priority, actual.Priority, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Priority")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.MatchRule, actual.MatchRule, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRule, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MatchRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RouteAction, actual.RouteAction, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RouteAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRedirect, actual.UrlRedirect, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleUrlRedirectNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleUrlRedirect, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRule)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRule or *UrlMapPathMatcherRouteRuleMatchRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRule)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PrefixMatch, actual.PrefixMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullPathMatch, actual.FullPathMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullPathMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoreCase, actual.IgnoreCase, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IgnoreCase")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderMatch, actual.HeaderMatch, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderMatches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QueryParameterMatch, actual.QueryParameterMatch, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("QueryParameterMatches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetadataFilter, actual.MetadataFilter, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetadataFilters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch or *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HeaderName, actual.HeaderName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExactMatch, actual.ExactMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExactMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RangeMatch, actual.RangeMatch, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RangeMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PresentMatch, actual.PresentMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PresentMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixMatch, actual.PrefixMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuffixMatch, actual.SuffixMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SuffixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertMatch, actual.InvertMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InvertMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch or *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RangeStart, actual.RangeStart, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RangeStart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RangeEnd, actual.RangeEnd, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RangeEnd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch or *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PresentMatch, actual.PresentMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PresentMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExactMatch, actual.ExactMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExactMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter or *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FilterMatchCriteria, actual.FilterMatchCriteria, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FilterMatchCriteria")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FilterLabel, actual.FilterLabel, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FilterLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel or *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteAction)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteAction or *UrlMapPathMatcherRouteRuleRouteAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteAction)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WeightedBackendService, actual.WeightedBackendService, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WeightedBackendServices")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRewrite, actual.UrlRewrite, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionUrlRewriteNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionUrlRewrite, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionTimeoutNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryPolicy, actual.RetryPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestMirrorPolicy, actual.RequestMirrorPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestMirrorPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CorsPolicy, actual.CorsPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionCorsPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CorsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FaultInjectionPolicy, actual.FaultInjectionPolicy, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FaultInjectionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService or *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HeaderAction, actual.HeaderAction, dcl.Info{ObjectFunction: compareUrlMapHeaderActionNewStyle, EmptyObject: EmptyUrlMapHeaderAction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HeaderAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionUrlRewriteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionUrlRewrite)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite or *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionUrlRewrite)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PathPrefixRewrite, actual.PathPrefixRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathPrefixRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HostRewrite, actual.HostRewrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionTimeout or *UrlMapPathMatcherRouteRuleRouteActionTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionRetryPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy or *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionRetryPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryCondition, actual.RetryCondition, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetryConditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRetries, actual.NumRetries, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerTryTimeout, actual.PerTryTimeout, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerTryTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout or *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy or *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BackendService, actual.BackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BackendService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionCorsPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionCorsPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy or *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionCorsPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowOrigin, actual.AllowOrigin, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOrigins")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowOriginRegex, actual.AllowOriginRegex, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowOriginRegexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowMethod, actual.AllowMethod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowMethods")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowHeader, actual.AllowHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExposeHeader, actual.ExposeHeader, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExposeHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxAge, actual.MaxAge, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxAge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowCredentials, actual.AllowCredentials, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowCredentials")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy or *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Delay, actual.Delay, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Delay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abort, actual.Abort, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Abort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay or *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedDelay, actual.FixedDelay, dcl.Info{ObjectFunction: compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle, EmptyObject: EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FixedDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay or *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay", a)
		}
		actual = &actualNotPointer
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

func compareUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort or *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpStatus, actual.HttpStatus, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapPathMatcherRouteRuleUrlRedirectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapPathMatcherRouteRuleUrlRedirect)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapPathMatcherRouteRuleUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleUrlRedirect or *UrlMapPathMatcherRouteRuleUrlRedirect", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapPathMatcherRouteRuleUrlRedirect)
	if !ok {
		actualNotPointer, ok := a.(UrlMapPathMatcherRouteRuleUrlRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapPathMatcherRouteRuleUrlRedirect", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostRedirect, actual.HostRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRedirect, actual.PathRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PathRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixRedirect, actual.PrefixRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrefixRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedirectResponseCode, actual.RedirectResponseCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedirectResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsRedirect, actual.HttpsRedirect, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StripQuery, actual.StripQuery, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StripQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareUrlMapTestNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UrlMapTest)
	if !ok {
		desiredNotPointer, ok := d.(UrlMapTest)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapTest or *UrlMapTest", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UrlMapTest)
	if !ok {
		actualNotPointer, ok := a.(UrlMapTest)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UrlMapTest", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpectedBackendService, actual.ExpectedBackendService, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *UrlMap) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UrlMap) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *UrlMap) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UrlMap) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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
	return unmarshalMapUrlMap(m, c)
}

func unmarshalMapUrlMap(m map[string]interface{}, c *Client) (*UrlMap, error) {

	return flattenUrlMap(c, m), nil
}

// expandUrlMap expands UrlMap into a JSON request object.
func expandUrlMap(c *Client, f *UrlMap) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := expandUrlMapDefaultRouteAction(c, f.DefaultRouteAction); err != nil {
		return nil, fmt.Errorf("error expanding DefaultRouteAction into defaultRouteAction: %w", err)
	} else if v != nil {
		m["defaultRouteAction"] = v
	}
	if v := f.DefaultService; !dcl.IsEmptyValueIndirect(v) {
		m["defaultService"] = v
	}
	if v, err := expandUrlMapDefaultUrlRedirect(c, f.DefaultUrlRedirect); err != nil {
		return nil, fmt.Errorf("error expanding DefaultUrlRedirect into defaultUrlRedirect: %w", err)
	} else if v != nil {
		m["defaultUrlRedirect"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := expandUrlMapHeaderAction(c, f.HeaderAction); err != nil {
		return nil, fmt.Errorf("error expanding HeaderAction into headerAction: %w", err)
	} else if v != nil {
		m["headerAction"] = v
	}
	if v, err := expandUrlMapHostRuleSlice(c, f.HostRule); err != nil {
		return nil, fmt.Errorf("error expanding HostRule into hostRules: %w", err)
	} else {
		m["hostRules"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandUrlMapPathMatcherSlice(c, f.PathMatcher); err != nil {
		return nil, fmt.Errorf("error expanding PathMatcher into pathMatchers: %w", err)
	} else {
		m["pathMatchers"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := expandUrlMapTestSlice(c, f.Test); err != nil {
		return nil, fmt.Errorf("error expanding Test into tests: %w", err)
	} else {
		m["tests"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
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

	res := &UrlMap{}
	res.DefaultRouteAction = flattenUrlMapDefaultRouteAction(c, m["defaultRouteAction"])
	res.DefaultService = dcl.FlattenString(m["defaultService"])
	res.DefaultUrlRedirect = flattenUrlMapDefaultUrlRedirect(c, m["defaultUrlRedirect"])
	res.Description = dcl.FlattenString(m["description"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.HeaderAction = flattenUrlMapHeaderAction(c, m["headerAction"])
	res.HostRule = flattenUrlMapHostRuleSlice(c, m["hostRules"])
	res.Name = dcl.FlattenString(m["name"])
	res.PathMatcher = flattenUrlMapPathMatcherSlice(c, m["pathMatchers"])
	res.Region = dcl.FlattenString(m["region"])
	res.Test = flattenUrlMapTestSlice(c, m["tests"])
	res.Project = dcl.FlattenString(m["project"])

	return res
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteAction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionWeightedBackendService
	}
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
	if v := f.RequestHeadersToRemove; v != nil {
		m["requestHeadersToRemove"] = v
	}
	if v, err := expandUrlMapHeaderActionRequestHeadersToAddSlice(c, f.RequestHeadersToAdd); err != nil {
		return nil, fmt.Errorf("error expanding RequestHeadersToAdd into requestHeadersToAdd: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestHeadersToAdd"] = v
	}
	if v := f.ResponseHeadersToRemove; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapHeaderAction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapHeaderActionRequestHeadersToAdd
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapHeaderActionResponseHeadersToAdd
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionUrlRewrite
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionTimeout
	}
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
	if v := f.RetryCondition; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionRetryPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionRetryPolicyPerTryTimeout
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionRequestMirrorPolicy
	}
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
	if v := f.AllowOrigin; v != nil {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; v != nil {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; v != nil {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; v != nil {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionCorsPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionFaultInjectionPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultRouteActionFaultInjectionPolicyAbort
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapDefaultUrlRedirect
	}
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
	if v := f.Host; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapHostRule
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcher
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherDefaultUrlRedirect
	}
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
	if v := f.Path; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRule
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteAction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionWeightedBackendService
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionUrlRewrite
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionTimeout
	}
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
	if v := f.RetryCondition; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	}
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
	if v := f.AllowOrigin; v != nil {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; v != nil {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; v != nil {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; v != nil {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionCorsPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherPathRuleUrlRedirect
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRule
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRule
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteAction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionUrlRewrite
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionTimeout
	}
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
	if v := f.RetryCondition; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy
	}
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
	if v := f.AllowOrigin; v != nil {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegex; v != nil {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethod; v != nil {
		m["allowMethods"] = v
	}
	if v := f.AllowHeader; v != nil {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeader; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionCorsPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapPathMatcherRouteRuleUrlRedirect
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUrlMapTest
	}
	r.Description = dcl.FlattenString(m["description"])
	r.Host = dcl.FlattenString(m["host"])
	r.Path = dcl.FlattenString(m["path"])
	r.ExpectedBackendService = dcl.FlattenString(m["service"])

	return r
}

// flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnumMap flattens the contents of UrlMapDefaultUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnumMap(c *Client, i interface{}) map[string]UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]UrlMapDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make(map[string]UrlMapDefaultUrlRedirectRedirectResponseCodeEnum)
	for k, item := range a {
		items[k] = *flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(item.(interface{}))
	}

	return items
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
		items = append(items, *flattenUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(item.(interface{})))
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

// flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumMap flattens the contents of UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumMap(c *Client, i interface{}) map[string]UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make(map[string]UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(item.(interface{}))
	}

	return items
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
		items = append(items, *flattenUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(item.(interface{})))
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

// flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumMap flattens the contents of UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumMap(c *Client, i interface{}) map[string]UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make(map[string]UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(item.(interface{}))
	}

	return items
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
		items = append(items, *flattenUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(item.(interface{})))
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

// flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumMap flattens the contents of UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(item.(interface{}))
	}

	return items
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
		items = append(items, *flattenUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(item.(interface{})))
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

// flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumMap flattens the contents of UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum from a JSON
// response object.
func flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumMap(c *Client, i interface{}) map[string]UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum{}
	}

	items := make(map[string]UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum)
	for k, item := range a {
		items[k] = *flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(item.(interface{}))
	}

	return items
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
		items = append(items, *flattenUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(item.(interface{})))
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

type urlMapDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         urlMapApiOperation
}

func convertFieldDiffsToUrlMapDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]urlMapDiff, error) {
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
	var diffs []urlMapDiff
	// For each operation name, create a urlMapDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := urlMapDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToUrlMapApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToUrlMapApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (urlMapApiOperation, error) {
	switch opName {

	case "updateUrlMapUpdateOperation":
		return &updateUrlMapUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
