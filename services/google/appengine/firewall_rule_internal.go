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
package appengine

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

func (r *FirewallRule) validate() error {

	if err := dcl.RequiredParameter(r.App, "App"); err != nil {
		return err
	}
	return nil
}

func firewallRuleGetURL(userBasePath string, r *FirewallRule) (string, error) {
	params := map[string]interface{}{
		"app":      dcl.ValueOrEmptyString(r.App),
		"priority": dcl.ValueOrEmptyString(r.Priority),
	}
	return dcl.URL("apps/{{app}}/firewall/ingressRules/{{priority}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

func firewallRuleListURL(userBasePath, app string) (string, error) {
	params := map[string]interface{}{
		"app": app,
	}
	return dcl.URL("apps/{{app}}/firewall/ingressRules", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func firewallRuleCreateURL(userBasePath, app string) (string, error) {
	params := map[string]interface{}{
		"app": app,
	}
	return dcl.URL("apps/{{app}}/firewall/ingressRules", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func firewallRuleDeleteURL(userBasePath string, r *FirewallRule) (string, error) {
	params := map[string]interface{}{
		"app":      dcl.ValueOrEmptyString(r.App),
		"priority": dcl.ValueOrEmptyString(r.Priority),
	}
	return dcl.URL("apps/{{app}}/firewall/ingressRules/{{priority}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

// firewallRuleApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type firewallRuleApiOperation interface {
	do(context.Context, *FirewallRule, *Client) error
}

// newUpdateFirewallRulePatchFirewallRuleRequest creates a request for an
// FirewallRule resource's PatchFirewallRule update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFirewallRulePatchFirewallRuleRequest(ctx context.Context, f *FirewallRule, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateFirewallRulePatchFirewallRuleRequest converts the update into
// the final JSON request body.
func marshalUpdateFirewallRulePatchFirewallRuleRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFirewallRulePatchFirewallRuleOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFirewallRulePatchFirewallRuleOperation) do(ctx context.Context, r *FirewallRule, c *Client) error {
	_, err := c.GetFirewallRule(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchFirewallRule")
	if err != nil {
		return err
	}

	req, err := newUpdateFirewallRulePatchFirewallRuleRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFirewallRulePatchFirewallRuleRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listFirewallRuleRaw(ctx context.Context, app, pageToken string, pageSize int32) ([]byte, error) {
	u, err := firewallRuleListURL(c.Config.BasePath, app)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FirewallRuleMaxPage {
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

type listFirewallRuleOperation struct {
	IngressRules []map[string]interface{} `json:"ingressRules"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listFirewallRule(ctx context.Context, app, pageToken string, pageSize int32) ([]*FirewallRule, string, error) {
	b, err := c.listFirewallRuleRaw(ctx, app, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFirewallRuleOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*FirewallRule
	for _, v := range m.IngressRules {
		res := flattenFirewallRule(c, v)
		res.App = &app
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFirewallRule(ctx context.Context, f func(*FirewallRule) bool, resources []*FirewallRule) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFirewallRule(ctx, res)
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

type deleteFirewallRuleOperation struct{}

func (op *deleteFirewallRuleOperation) do(ctx context.Context, r *FirewallRule, c *Client) error {

	_, err := c.GetFirewallRule(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("FirewallRule not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFirewallRule checking for existence. error: %v", err)
		return err
	}

	u, err := firewallRuleDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete FirewallRule: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetFirewallRule(ctx, r.urlNormalized())
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
type createFirewallRuleOperation struct {
	response map[string]interface{}
}

func (op *createFirewallRuleOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFirewallRuleOperation) do(ctx context.Context, r *FirewallRule, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	app := r.createFields()
	u, err := firewallRuleCreateURL(c.Config.BasePath, app)

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

	if _, err := c.GetFirewallRule(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFirewallRuleRaw(ctx context.Context, r *FirewallRule) ([]byte, error) {

	u, err := firewallRuleGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) firewallRuleDiffsForRawDesired(ctx context.Context, rawDesired *FirewallRule, opts ...dcl.ApplyOption) (initial, desired *FirewallRule, diffs []firewallRuleDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FirewallRule
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FirewallRule); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected FirewallRule, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFirewallRule(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a FirewallRule resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FirewallRule resource: %v", err)
		}
		c.Config.Logger.Info("Found that FirewallRule resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFirewallRuleDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for FirewallRule: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for FirewallRule: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFirewallRuleInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for FirewallRule: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFirewallRuleDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for FirewallRule: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFirewallRule(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFirewallRuleInitialState(rawInitial, rawDesired *FirewallRule) (*FirewallRule, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFirewallRuleDesiredState(rawDesired, rawInitial *FirewallRule, opts ...dcl.ApplyOption) (*FirewallRule, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Action) {
		rawDesired.Action = rawInitial.Action
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Priority) {
		rawDesired.Priority = rawInitial.Priority
	}
	if dcl.StringCanonicalize(rawDesired.SourceRange, rawInitial.SourceRange) {
		rawDesired.SourceRange = rawInitial.SourceRange
	}
	if dcl.NameToSelfLink(rawDesired.App, rawInitial.App) {
		rawDesired.App = rawInitial.App
	}

	return rawDesired, nil
}

func canonicalizeFirewallRuleNewState(c *Client, rawNew, rawDesired *FirewallRule) (*FirewallRule, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Action) && dcl.IsEmptyValueIndirect(rawDesired.Action) {
		rawNew.Action = rawDesired.Action
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Priority) && dcl.IsEmptyValueIndirect(rawDesired.Priority) {
		rawNew.Priority = rawDesired.Priority
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceRange) && dcl.IsEmptyValueIndirect(rawDesired.SourceRange) {
		rawNew.SourceRange = rawDesired.SourceRange
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceRange, rawNew.SourceRange) {
			rawNew.SourceRange = rawDesired.SourceRange
		}
	}

	rawNew.App = rawDesired.App

	return rawNew, nil
}

type firewallRuleDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         firewallRuleApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffFirewallRule(c *Client, desired, actual *FirewallRule, opts ...dcl.ApplyOption) ([]firewallRuleDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []firewallRuleDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFirewallRuleDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFirewallRuleDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Priority, actual.Priority, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Priority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFirewallRuleDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SourceRange, actual.SourceRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFirewallRuleDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.App, actual.App, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("App")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFirewallRuleDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
	var deduped []firewallRuleDiff
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *FirewallRule) urlNormalized() *FirewallRule {
	normalized := dcl.Copy(*r).(FirewallRule)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceRange = dcl.SelfLinkToName(r.SourceRange)
	normalized.App = dcl.SelfLinkToName(r.App)
	return &normalized
}

func (r *FirewallRule) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Priority)
}

func (r *FirewallRule) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App)
}

func (r *FirewallRule) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Priority)
}

func (r *FirewallRule) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "PatchFirewallRule" {
		fields := map[string]interface{}{
			"app":      dcl.ValueOrEmptyString(n.App),
			"priority": dcl.ValueOrEmptyString(n.Priority),
		}
		return dcl.URL("apps/{{app}}/firewall/ingressRules/{{priority}}", "https://appengine.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FirewallRule resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FirewallRule) marshal(c *Client) ([]byte, error) {
	m, err := expandFirewallRule(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FirewallRule: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFirewallRule decodes JSON responses into the FirewallRule resource schema.
func unmarshalFirewallRule(b []byte, c *Client) (*FirewallRule, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFirewallRule(m, c)
}

func unmarshalMapFirewallRule(m map[string]interface{}, c *Client) (*FirewallRule, error) {

	return flattenFirewallRule(c, m), nil
}

// expandFirewallRule expands FirewallRule into a JSON request object.
func expandFirewallRule(c *Client, f *FirewallRule) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		m["priority"] = v
	}
	if v := f.SourceRange; !dcl.IsEmptyValueIndirect(v) {
		m["sourceRange"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding App into app: %w", err)
	} else if v != nil {
		m["app"] = v
	}

	return m, nil
}

// flattenFirewallRule flattens FirewallRule from a JSON request object into the
// FirewallRule type.
func flattenFirewallRule(c *Client, i interface{}) *FirewallRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &FirewallRule{}
	res.Action = flattenFirewallRuleActionEnum(m["action"])
	res.Description = dcl.FlattenString(m["description"])
	res.Priority = dcl.FlattenInteger(m["priority"])
	res.SourceRange = dcl.FlattenString(m["sourceRange"])
	res.App = dcl.FlattenString(m["app"])

	return res
}

// flattenFirewallRuleActionEnumSlice flattens the contents of FirewallRuleActionEnum from a JSON
// response object.
func flattenFirewallRuleActionEnumSlice(c *Client, i interface{}) []FirewallRuleActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallRuleActionEnum{}
	}

	if len(a) == 0 {
		return []FirewallRuleActionEnum{}
	}

	items := make([]FirewallRuleActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallRuleActionEnum(item.(interface{})))
	}

	return items
}

// flattenFirewallRuleActionEnum asserts that an interface is a string, and returns a
// pointer to a *FirewallRuleActionEnum with the same value as that string.
func flattenFirewallRuleActionEnum(i interface{}) *FirewallRuleActionEnum {
	s, ok := i.(string)
	if !ok {
		return FirewallRuleActionEnumRef("")
	}

	return FirewallRuleActionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FirewallRule) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFirewallRule(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.App == nil && ncr.App == nil {
			c.Config.Logger.Info("Both App fields null - considering equal.")
		} else if nr.App == nil || ncr.App == nil {
			c.Config.Logger.Info("Only one App field is null - considering unequal.")
			return false
		} else if *nr.App != *ncr.App {
			return false
		}
		if nr.Priority == nil && ncr.Priority == nil {
			c.Config.Logger.Info("Both Priority fields null - considering equal.")
		} else if nr.Priority == nil || ncr.Priority == nil {
			c.Config.Logger.Info("Only one Priority field is null - considering unequal.")
			return false
		} else if *nr.Priority != *ncr.Priority {
			return false
		}
		return true
	}
}

func convertFieldDiffToFirewallRuleDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]firewallRuleDiff, error) {
	var diffs []firewallRuleDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := firewallRuleDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTofirewallRuleApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameTofirewallRuleApiOperation(op string, opts ...dcl.ApplyOption) (firewallRuleApiOperation, error) {
	switch op {

	case "updateFirewallRulePatchFirewallRuleOperation":
		return &updateFirewallRulePatchFirewallRuleOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
