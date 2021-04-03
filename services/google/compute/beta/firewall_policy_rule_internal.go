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

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *FirewallPolicyRule) validate() error {

	if !dcl.IsEmptyValueIndirect(r.Match) {
		if err := r.Match.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FirewallPolicyRuleMatch) validate() error {
	return nil
}
func (r *FirewallPolicyRuleMatchLayer4Configs) validate() error {
	return nil
}

func firewallPolicyRuleGetURL(userBasePath string, r *FirewallPolicyRule) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": dcl.ValueOrEmptyString(r.FirewallPolicy),
		"priority":       dcl.ValueOrEmptyString(r.Priority),
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/getRule?priority={{priority}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func firewallPolicyRuleListURL(userBasePath, firewallPolicy string) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": firewallPolicy,
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func firewallPolicyRuleCreateURL(userBasePath, firewallPolicy string) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": firewallPolicy,
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/addRule", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func firewallPolicyRuleDeleteURL(userBasePath string, r *FirewallPolicyRule) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": dcl.ValueOrEmptyString(r.FirewallPolicy),
		"priority":       dcl.ValueOrEmptyString(r.Priority),
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/removeRule?priority={{priority}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// firewallPolicyRuleApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type firewallPolicyRuleApiOperation interface {
	do(context.Context, *FirewallPolicyRule, *Client) error
}

// newUpdateFirewallPolicyRulePatchRuleRequest creates a request for an
// FirewallPolicyRule resource's PatchRule update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFirewallPolicyRulePatchRuleRequest(ctx context.Context, f *FirewallPolicyRule, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandFirewallPolicyRuleMatch(c, f.Match); err != nil {
		return nil, fmt.Errorf("error expanding Match into match: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["match"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		req["action"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		req["direction"] = v
	}
	if v := f.TargetResources; !dcl.IsEmptyValueIndirect(v) {
		req["targetResources"] = v
	}
	if v := f.EnableLogging; !dcl.IsEmptyValueIndirect(v) {
		req["enableLogging"] = v
	}
	if v := f.RuleTupleCount; !dcl.IsEmptyValueIndirect(v) {
		req["ruleTupleCount"] = v
	}
	if v := f.TargetServiceAccounts; !dcl.IsEmptyValueIndirect(v) {
		req["targetServiceAccounts"] = v
	}
	if v := f.TargetSecureLabels; !dcl.IsEmptyValueIndirect(v) {
		req["targetSecureLabels"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	return req, nil
}

// marshalUpdateFirewallPolicyRulePatchRuleRequest converts the update into
// the final JSON request body.
func marshalUpdateFirewallPolicyRulePatchRuleRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFirewallPolicyRulePatchRuleOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listFirewallPolicyRuleRaw(ctx context.Context, firewallPolicy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := firewallPolicyRuleListURL(c.Config.BasePath, firewallPolicy)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FirewallPolicyRuleMaxPage {
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

type listFirewallPolicyRuleOperation struct {
	Rules []map[string]interface{} `json:"rules"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listFirewallPolicyRule(ctx context.Context, firewallPolicy, pageToken string, pageSize int32) ([]*FirewallPolicyRule, string, error) {
	b, err := c.listFirewallPolicyRuleRaw(ctx, firewallPolicy, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFirewallPolicyRuleOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*FirewallPolicyRule
	for _, v := range m.Rules {
		res := flattenFirewallPolicyRule(c, v)
		res.FirewallPolicy = &firewallPolicy
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFirewallPolicyRule(ctx context.Context, f func(*FirewallPolicyRule) bool, resources []*FirewallPolicyRule) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFirewallPolicyRule(ctx, res)
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

type deleteFirewallPolicyRuleOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFirewallPolicyRuleOperation struct {
	response map[string]interface{}
}

func (op *createFirewallPolicyRuleOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getFirewallPolicyRuleRaw(ctx context.Context, r *FirewallPolicyRule) ([]byte, error) {

	u, err := firewallPolicyRuleGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) firewallPolicyRuleDiffsForRawDesired(ctx context.Context, rawDesired *FirewallPolicyRule, opts ...dcl.ApplyOption) (initial, desired *FirewallPolicyRule, diffs []firewallPolicyRuleDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FirewallPolicyRule
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FirewallPolicyRule); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected FirewallPolicyRule, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFirewallPolicyRule(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 400) {
			c.Config.Logger.Warningf("Failed to retrieve whether a FirewallPolicyRule resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FirewallPolicyRule resource: %v", err)
		}
		c.Config.Logger.Info("Found that FirewallPolicyRule resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFirewallPolicyRuleDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for FirewallPolicyRule: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for FirewallPolicyRule: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFirewallPolicyRuleInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for FirewallPolicyRule: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFirewallPolicyRuleDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for FirewallPolicyRule: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFirewallPolicyRule(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFirewallPolicyRuleInitialState(rawInitial, rawDesired *FirewallPolicyRule) (*FirewallPolicyRule, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFirewallPolicyRuleDesiredState(rawDesired, rawInitial *FirewallPolicyRule, opts ...dcl.ApplyOption) (*FirewallPolicyRule, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Match = canonicalizeFirewallPolicyRuleMatch(rawDesired.Match, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Priority) {
		rawDesired.Priority = rawInitial.Priority
	}
	rawDesired.Match = canonicalizeFirewallPolicyRuleMatch(rawDesired.Match, rawInitial.Match, opts...)
	if dcl.StringCanonicalize(rawDesired.Action, rawInitial.Action) {
		rawDesired.Action = rawInitial.Action
	}
	if dcl.IsZeroValue(rawDesired.Direction) {
		rawDesired.Direction = rawInitial.Direction
	}
	if dcl.IsZeroValue(rawDesired.TargetResources) {
		rawDesired.TargetResources = rawInitial.TargetResources
	}
	if dcl.BoolCanonicalize(rawDesired.EnableLogging, rawInitial.EnableLogging) {
		rawDesired.EnableLogging = rawInitial.EnableLogging
	}
	if dcl.IsZeroValue(rawDesired.RuleTupleCount) {
		rawDesired.RuleTupleCount = rawInitial.RuleTupleCount
	}
	if dcl.IsZeroValue(rawDesired.TargetServiceAccounts) {
		rawDesired.TargetServiceAccounts = rawInitial.TargetServiceAccounts
	}
	if dcl.IsZeroValue(rawDesired.TargetSecureLabels) {
		rawDesired.TargetSecureLabels = rawInitial.TargetSecureLabels
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		rawDesired.Disabled = rawInitial.Disabled
	}
	if dcl.StringCanonicalize(rawDesired.Kind, rawInitial.Kind) {
		rawDesired.Kind = rawInitial.Kind
	}
	if dcl.NameToSelfLink(rawDesired.FirewallPolicy, rawInitial.FirewallPolicy) {
		rawDesired.FirewallPolicy = rawInitial.FirewallPolicy
	}

	return rawDesired, nil
}

func canonicalizeFirewallPolicyRuleNewState(c *Client, rawNew, rawDesired *FirewallPolicyRule) (*FirewallPolicyRule, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Match) && dcl.IsEmptyValueIndirect(rawDesired.Match) {
		rawNew.Match = rawDesired.Match
	} else {
		rawNew.Match = canonicalizeNewFirewallPolicyRuleMatch(c, rawDesired.Match, rawNew.Match)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Action) && dcl.IsEmptyValueIndirect(rawDesired.Action) {
		rawNew.Action = rawDesired.Action
	} else {
		if dcl.StringCanonicalize(rawDesired.Action, rawNew.Action) {
			rawNew.Action = rawDesired.Action
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Direction) && dcl.IsEmptyValueIndirect(rawDesired.Direction) {
		rawNew.Direction = rawDesired.Direction
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetResources) && dcl.IsEmptyValueIndirect(rawDesired.TargetResources) {
		rawNew.TargetResources = rawDesired.TargetResources
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableLogging) && dcl.IsEmptyValueIndirect(rawDesired.EnableLogging) {
		rawNew.EnableLogging = rawDesired.EnableLogging
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableLogging, rawNew.EnableLogging) {
			rawNew.EnableLogging = rawDesired.EnableLogging
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuleTupleCount) && dcl.IsEmptyValueIndirect(rawDesired.RuleTupleCount) {
		rawNew.RuleTupleCount = rawDesired.RuleTupleCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetServiceAccounts) && dcl.IsEmptyValueIndirect(rawDesired.TargetServiceAccounts) {
		rawNew.TargetServiceAccounts = rawDesired.TargetServiceAccounts
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetSecureLabels) && dcl.IsEmptyValueIndirect(rawDesired.TargetSecureLabels) {
		rawNew.TargetSecureLabels = rawDesired.TargetSecureLabels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Kind) && dcl.IsEmptyValueIndirect(rawDesired.Kind) {
		rawNew.Kind = rawDesired.Kind
	} else {
		if dcl.StringCanonicalize(rawDesired.Kind, rawNew.Kind) {
			rawNew.Kind = rawDesired.Kind
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.FirewallPolicy) && dcl.IsEmptyValueIndirect(rawDesired.FirewallPolicy) {
		rawNew.FirewallPolicy = rawDesired.FirewallPolicy
	} else {
		if dcl.NameToSelfLink(rawDesired.FirewallPolicy, rawNew.FirewallPolicy) {
			rawNew.FirewallPolicy = rawDesired.FirewallPolicy
		}
	}

	return rawNew, nil
}

func canonicalizeFirewallPolicyRuleMatch(des, initial *FirewallPolicyRuleMatch, opts ...dcl.ApplyOption) *FirewallPolicyRuleMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.SrcIPRanges) {
		des.SrcIPRanges = initial.SrcIPRanges
	}
	if dcl.IsZeroValue(des.DestIPRanges) {
		des.DestIPRanges = initial.DestIPRanges
	}
	if dcl.IsZeroValue(des.Layer4Configs) {
		des.Layer4Configs = initial.Layer4Configs
	}
	if dcl.IsZeroValue(des.SrcSecureLabels) {
		des.SrcSecureLabels = initial.SrcSecureLabels
	}

	return des
}

func canonicalizeNewFirewallPolicyRuleMatch(c *Client, des, nw *FirewallPolicyRuleMatch) *FirewallPolicyRuleMatch {
	if des == nil || nw == nil {
		return nw
	}

	nw.Layer4Configs = canonicalizeNewFirewallPolicyRuleMatchLayer4ConfigsSlice(c, des.Layer4Configs, nw.Layer4Configs)

	return nw
}

func canonicalizeNewFirewallPolicyRuleMatchSet(c *Client, des, nw []FirewallPolicyRuleMatch) []FirewallPolicyRuleMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []FirewallPolicyRuleMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFirewallPolicyRuleMatch(c, &d, &n) {
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

func canonicalizeNewFirewallPolicyRuleMatchSlice(c *Client, des, nw []FirewallPolicyRuleMatch) []FirewallPolicyRuleMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FirewallPolicyRuleMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirewallPolicyRuleMatch(c, &d, &n))
	}

	return items
}

func canonicalizeFirewallPolicyRuleMatchLayer4Configs(des, initial *FirewallPolicyRuleMatchLayer4Configs, opts ...dcl.ApplyOption) *FirewallPolicyRuleMatchLayer4Configs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.IPProtocol, initial.IPProtocol) || dcl.IsZeroValue(des.IPProtocol) {
		des.IPProtocol = initial.IPProtocol
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	}

	return des
}

func canonicalizeNewFirewallPolicyRuleMatchLayer4Configs(c *Client, des, nw *FirewallPolicyRuleMatchLayer4Configs) *FirewallPolicyRuleMatchLayer4Configs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPProtocol, nw.IPProtocol) {
		nw.IPProtocol = des.IPProtocol
	}

	return nw
}

func canonicalizeNewFirewallPolicyRuleMatchLayer4ConfigsSet(c *Client, des, nw []FirewallPolicyRuleMatchLayer4Configs) []FirewallPolicyRuleMatchLayer4Configs {
	if des == nil {
		return nw
	}
	var reorderedNew []FirewallPolicyRuleMatchLayer4Configs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFirewallPolicyRuleMatchLayer4Configs(c, &d, &n) {
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

func canonicalizeNewFirewallPolicyRuleMatchLayer4ConfigsSlice(c *Client, des, nw []FirewallPolicyRuleMatchLayer4Configs) []FirewallPolicyRuleMatchLayer4Configs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FirewallPolicyRuleMatchLayer4Configs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirewallPolicyRuleMatchLayer4Configs(c, &d, &n))
	}

	return items
}

type firewallPolicyRuleDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         firewallPolicyRuleApiOperation
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
func diffFirewallPolicyRule(c *Client, desired, actual *FirewallPolicyRule, opts ...dcl.ApplyOption) ([]firewallPolicyRuleDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []firewallPolicyRuleDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp: &updateFirewallPolicyRulePatchRuleOperation{}, FieldName: "Description",
		})
	}

	if d, err := dcl.Diff(desired.Priority, actual.Priority, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{RequiresRecreate: true, FieldName: "Priority"})
	}

	if d, err := dcl.Diff(desired.Action, actual.Action, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp: &updateFirewallPolicyRulePatchRuleOperation{}, FieldName: "Action",
		})
	}

	if d, err := dcl.Diff(desired.EnableLogging, actual.EnableLogging, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp: &updateFirewallPolicyRulePatchRuleOperation{}, FieldName: "EnableLogging",
		})
	}

	if d, err := dcl.Diff(desired.RuleTupleCount, actual.RuleTupleCount, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp: &updateFirewallPolicyRulePatchRuleOperation{}, FieldName: "RuleTupleCount",
		})
	}

	if d, err := dcl.Diff(desired.Disabled, actual.Disabled, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp: &updateFirewallPolicyRulePatchRuleOperation{}, FieldName: "Disabled",
		})
	}

	if d, err := dcl.Diff(desired.Kind, actual.Kind, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{RequiresRecreate: true, FieldName: "Kind"})
	}

	if d, err := dcl.Diff(desired.FirewallPolicy, actual.FirewallPolicy, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, firewallPolicyRuleDiff{RequiresRecreate: true, FieldName: "FirewallPolicy"})
	}

	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "Description",
		})

	}
	if !reflect.DeepEqual(desired.Priority, actual.Priority) {
		c.Config.Logger.Infof("Detected diff in Priority.\nDESIRED: %v\nACTUAL: %v", desired.Priority, actual.Priority)
		diffs = append(diffs, firewallPolicyRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Priority",
		})
	}
	if compareFirewallPolicyRuleMatch(c, desired.Match, actual.Match) {
		c.Config.Logger.Infof("Detected diff in Match.\nDESIRED: %v\nACTUAL: %v", desired.Match, actual.Match)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "Match",
		})

	}
	if !dcl.IsZeroValue(desired.Action) && !dcl.StringCanonicalize(desired.Action, actual.Action) {
		c.Config.Logger.Infof("Detected diff in Action.\nDESIRED: %v\nACTUAL: %v", desired.Action, actual.Action)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "Action",
		})

	}
	if !reflect.DeepEqual(desired.Direction, actual.Direction) {
		c.Config.Logger.Infof("Detected diff in Direction.\nDESIRED: %v\nACTUAL: %v", desired.Direction, actual.Direction)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "Direction",
		})

	}
	if !dcl.StringSliceEqualsWithSelfLink(desired.TargetResources, actual.TargetResources) {
		c.Config.Logger.Infof("Detected diff in TargetResources.\nDESIRED: %v\nACTUAL: %v", desired.TargetResources, actual.TargetResources)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "TargetResources",
		})

	}
	if !dcl.IsZeroValue(desired.EnableLogging) && !dcl.BoolCanonicalize(desired.EnableLogging, actual.EnableLogging) {
		c.Config.Logger.Infof("Detected diff in EnableLogging.\nDESIRED: %v\nACTUAL: %v", desired.EnableLogging, actual.EnableLogging)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "EnableLogging",
		})

	}
	if !dcl.StringSliceEqualsWithSelfLink(desired.TargetServiceAccounts, actual.TargetServiceAccounts) {
		c.Config.Logger.Infof("Detected diff in TargetServiceAccounts.\nDESIRED: %v\nACTUAL: %v", desired.TargetServiceAccounts, actual.TargetServiceAccounts)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "TargetServiceAccounts",
		})

	}
	if !dcl.StringSliceEquals(desired.TargetSecureLabels, actual.TargetSecureLabels) {
		c.Config.Logger.Infof("Detected diff in TargetSecureLabels.\nDESIRED: %v\nACTUAL: %v", desired.TargetSecureLabels, actual.TargetSecureLabels)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "TargetSecureLabels",
		})

	}
	if !dcl.IsZeroValue(desired.Disabled) && !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) {
		c.Config.Logger.Infof("Detected diff in Disabled.\nDESIRED: %v\nACTUAL: %v", desired.Disabled, actual.Disabled)

		diffs = append(diffs, firewallPolicyRuleDiff{
			UpdateOp:  &updateFirewallPolicyRulePatchRuleOperation{},
			FieldName: "Disabled",
		})

	}
	if !dcl.IsZeroValue(desired.Kind) && !dcl.StringCanonicalize(desired.Kind, actual.Kind) {
		c.Config.Logger.Infof("Detected diff in Kind.\nDESIRED: %v\nACTUAL: %v", desired.Kind, actual.Kind)
		diffs = append(diffs, firewallPolicyRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Kind",
		})
	}
	if !dcl.IsZeroValue(desired.FirewallPolicy) && !dcl.NameToSelfLink(desired.FirewallPolicy, actual.FirewallPolicy) {
		c.Config.Logger.Infof("Detected diff in FirewallPolicy.\nDESIRED: %v\nACTUAL: %v", desired.FirewallPolicy, actual.FirewallPolicy)
		diffs = append(diffs, firewallPolicyRuleDiff{
			RequiresRecreate: true,
			FieldName:        "FirewallPolicy",
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
	var deduped []firewallPolicyRuleDiff
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
func compareFirewallPolicyRuleMatch(c *Client, desired, actual *FirewallPolicyRuleMatch) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.SrcIPRanges == nil && desired.SrcIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.SrcIPRanges) {
		c.Config.Logger.Infof("desired SrcIPRanges %s - but actually nil", dcl.SprintResource(desired.SrcIPRanges))
		return true
	}
	if !dcl.StringSliceEquals(desired.SrcIPRanges, actual.SrcIPRanges) && !dcl.IsZeroValue(desired.SrcIPRanges) {
		c.Config.Logger.Infof("Diff in SrcIPRanges. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SrcIPRanges), dcl.SprintResource(actual.SrcIPRanges))
		return true
	}
	if actual.DestIPRanges == nil && desired.DestIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.DestIPRanges) {
		c.Config.Logger.Infof("desired DestIPRanges %s - but actually nil", dcl.SprintResource(desired.DestIPRanges))
		return true
	}
	if !dcl.StringSliceEquals(desired.DestIPRanges, actual.DestIPRanges) && !dcl.IsZeroValue(desired.DestIPRanges) {
		c.Config.Logger.Infof("Diff in DestIPRanges. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DestIPRanges), dcl.SprintResource(actual.DestIPRanges))
		return true
	}
	if actual.Layer4Configs == nil && desired.Layer4Configs != nil && !dcl.IsEmptyValueIndirect(desired.Layer4Configs) {
		c.Config.Logger.Infof("desired Layer4Configs %s - but actually nil", dcl.SprintResource(desired.Layer4Configs))
		return true
	}
	if compareFirewallPolicyRuleMatchLayer4ConfigsSlice(c, desired.Layer4Configs, actual.Layer4Configs) && !dcl.IsZeroValue(desired.Layer4Configs) {
		c.Config.Logger.Infof("Diff in Layer4Configs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Layer4Configs), dcl.SprintResource(actual.Layer4Configs))
		return true
	}
	if actual.SrcSecureLabels == nil && desired.SrcSecureLabels != nil && !dcl.IsEmptyValueIndirect(desired.SrcSecureLabels) {
		c.Config.Logger.Infof("desired SrcSecureLabels %s - but actually nil", dcl.SprintResource(desired.SrcSecureLabels))
		return true
	}
	if !dcl.StringSliceEquals(desired.SrcSecureLabels, actual.SrcSecureLabels) && !dcl.IsZeroValue(desired.SrcSecureLabels) {
		c.Config.Logger.Infof("Diff in SrcSecureLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SrcSecureLabels), dcl.SprintResource(actual.SrcSecureLabels))
		return true
	}
	return false
}

func compareFirewallPolicyRuleMatchSlice(c *Client, desired, actual []FirewallPolicyRuleMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FirewallPolicyRuleMatch, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFirewallPolicyRuleMatch(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatch, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFirewallPolicyRuleMatchMap(c *Client, desired, actual map[string]FirewallPolicyRuleMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FirewallPolicyRuleMatch, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatch, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFirewallPolicyRuleMatch(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatch, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFirewallPolicyRuleMatchLayer4Configs(c *Client, desired, actual *FirewallPolicyRuleMatchLayer4Configs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IPProtocol == nil && desired.IPProtocol != nil && !dcl.IsEmptyValueIndirect(desired.IPProtocol) {
		c.Config.Logger.Infof("desired IPProtocol %s - but actually nil", dcl.SprintResource(desired.IPProtocol))
		return true
	}
	if !dcl.StringCanonicalize(desired.IPProtocol, actual.IPProtocol) && !dcl.IsZeroValue(desired.IPProtocol) {
		c.Config.Logger.Infof("Diff in IPProtocol. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPProtocol), dcl.SprintResource(actual.IPProtocol))
		return true
	}
	if actual.Ports == nil && desired.Ports != nil && !dcl.IsEmptyValueIndirect(desired.Ports) {
		c.Config.Logger.Infof("desired Ports %s - but actually nil", dcl.SprintResource(desired.Ports))
		return true
	}
	if !dcl.StringSliceEquals(desired.Ports, actual.Ports) && !dcl.IsZeroValue(desired.Ports) {
		c.Config.Logger.Infof("Diff in Ports. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Ports), dcl.SprintResource(actual.Ports))
		return true
	}
	return false
}

func compareFirewallPolicyRuleMatchLayer4ConfigsSlice(c *Client, desired, actual []FirewallPolicyRuleMatchLayer4Configs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FirewallPolicyRuleMatchLayer4Configs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFirewallPolicyRuleMatchLayer4Configs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatchLayer4Configs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFirewallPolicyRuleMatchLayer4ConfigsMap(c *Client, desired, actual map[string]FirewallPolicyRuleMatchLayer4Configs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FirewallPolicyRuleMatchLayer4Configs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatchLayer4Configs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFirewallPolicyRuleMatchLayer4Configs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleMatchLayer4Configs, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFirewallPolicyRuleDirectionEnumSlice(c *Client, desired, actual []FirewallPolicyRuleDirectionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FirewallPolicyRuleDirectionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFirewallPolicyRuleDirectionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FirewallPolicyRuleDirectionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFirewallPolicyRuleDirectionEnum(c *Client, desired, actual *FirewallPolicyRuleDirectionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *FirewallPolicyRule) urlNormalized() *FirewallPolicyRule {
	normalized := deepcopy.Copy(*r).(FirewallPolicyRule)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Action = dcl.SelfLinkToName(r.Action)
	normalized.Kind = dcl.SelfLinkToName(r.Kind)
	normalized.FirewallPolicy = dcl.SelfLinkToName(r.FirewallPolicy)
	return &normalized
}

func (r *FirewallPolicyRule) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy), dcl.ValueOrEmptyString(n.Priority)
}

func (r *FirewallPolicyRule) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy)
}

func (r *FirewallPolicyRule) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy), dcl.ValueOrEmptyString(n.Priority)
}

func (r *FirewallPolicyRule) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "PatchRule" {
		fields := map[string]interface{}{
			"firewallPolicy": dcl.ValueOrEmptyString(n.FirewallPolicy),
			"priority":       dcl.ValueOrEmptyString(n.Priority),
		}
		return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/patchRule?priority={{priority}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FirewallPolicyRule resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FirewallPolicyRule) marshal(c *Client) ([]byte, error) {
	m, err := expandFirewallPolicyRule(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FirewallPolicyRule: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFirewallPolicyRule decodes JSON responses into the FirewallPolicyRule resource schema.
func unmarshalFirewallPolicyRule(b []byte, c *Client) (*FirewallPolicyRule, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFirewallPolicyRule(m, c)
}

func unmarshalMapFirewallPolicyRule(m map[string]interface{}, c *Client) (*FirewallPolicyRule, error) {

	return flattenFirewallPolicyRule(c, m), nil
}

// expandFirewallPolicyRule expands FirewallPolicyRule into a JSON request object.
func expandFirewallPolicyRule(c *Client, f *FirewallPolicyRule) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		m["priority"] = v
	}
	if v, err := expandFirewallPolicyRuleMatch(c, f.Match); err != nil {
		return nil, fmt.Errorf("error expanding Match into match: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["match"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}
	if v := f.TargetResources; !dcl.IsEmptyValueIndirect(v) {
		m["targetResources"] = v
	}
	if v := f.EnableLogging; !dcl.IsEmptyValueIndirect(v) {
		m["enableLogging"] = v
	}
	if v := f.RuleTupleCount; !dcl.IsEmptyValueIndirect(v) {
		m["ruleTupleCount"] = v
	}
	if v := f.TargetServiceAccounts; !dcl.IsEmptyValueIndirect(v) {
		m["targetServiceAccounts"] = v
	}
	if v := f.TargetSecureLabels; !dcl.IsEmptyValueIndirect(v) {
		m["targetSecureLabels"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FirewallPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["firewallPolicy"] = v
	}

	return m, nil
}

// flattenFirewallPolicyRule flattens FirewallPolicyRule from a JSON request object into the
// FirewallPolicyRule type.
func flattenFirewallPolicyRule(c *Client, i interface{}) *FirewallPolicyRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &FirewallPolicyRule{}
	r.Description = dcl.FlattenString(m["description"])
	r.Priority = dcl.FlattenInteger(m["priority"])
	r.Match = flattenFirewallPolicyRuleMatch(c, m["match"])
	r.Action = dcl.FlattenString(m["action"])
	r.Direction = flattenFirewallPolicyRuleDirectionEnum(m["direction"])
	r.TargetResources = dcl.FlattenStringSlice(m["targetResources"])
	r.EnableLogging = dcl.FlattenBool(m["enableLogging"])
	r.RuleTupleCount = dcl.FlattenInteger(m["ruleTupleCount"])
	r.TargetServiceAccounts = dcl.FlattenStringSlice(m["targetServiceAccounts"])
	r.TargetSecureLabels = dcl.FlattenStringSlice(m["targetSecureLabels"])
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.FirewallPolicy = dcl.FlattenString(m["firewallPolicy"])

	return r
}

// expandFirewallPolicyRuleMatchMap expands the contents of FirewallPolicyRuleMatch into a JSON
// request object.
func expandFirewallPolicyRuleMatchMap(c *Client, f map[string]FirewallPolicyRuleMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirewallPolicyRuleMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirewallPolicyRuleMatchSlice expands the contents of FirewallPolicyRuleMatch into a JSON
// request object.
func expandFirewallPolicyRuleMatchSlice(c *Client, f []FirewallPolicyRuleMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirewallPolicyRuleMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirewallPolicyRuleMatchMap flattens the contents of FirewallPolicyRuleMatch from a JSON
// response object.
func flattenFirewallPolicyRuleMatchMap(c *Client, i interface{}) map[string]FirewallPolicyRuleMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallPolicyRuleMatch{}
	}

	if len(a) == 0 {
		return map[string]FirewallPolicyRuleMatch{}
	}

	items := make(map[string]FirewallPolicyRuleMatch)
	for k, item := range a {
		items[k] = *flattenFirewallPolicyRuleMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFirewallPolicyRuleMatchSlice flattens the contents of FirewallPolicyRuleMatch from a JSON
// response object.
func flattenFirewallPolicyRuleMatchSlice(c *Client, i interface{}) []FirewallPolicyRuleMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallPolicyRuleMatch{}
	}

	if len(a) == 0 {
		return []FirewallPolicyRuleMatch{}
	}

	items := make([]FirewallPolicyRuleMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallPolicyRuleMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandFirewallPolicyRuleMatch expands an instance of FirewallPolicyRuleMatch into a JSON
// request object.
func expandFirewallPolicyRuleMatch(c *Client, f *FirewallPolicyRuleMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SrcIPRanges; !dcl.IsEmptyValueIndirect(v) {
		m["srcIpRanges"] = v
	}
	if v := f.DestIPRanges; !dcl.IsEmptyValueIndirect(v) {
		m["destIpRanges"] = v
	}
	if v, err := expandFirewallPolicyRuleMatchLayer4ConfigsSlice(c, f.Layer4Configs); err != nil {
		return nil, fmt.Errorf("error expanding Layer4Configs into layer4Configs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["layer4Configs"] = v
	}
	if v := f.SrcSecureLabels; !dcl.IsEmptyValueIndirect(v) {
		m["srcSecureLabels"] = v
	}

	return m, nil
}

// flattenFirewallPolicyRuleMatch flattens an instance of FirewallPolicyRuleMatch from a JSON
// response object.
func flattenFirewallPolicyRuleMatch(c *Client, i interface{}) *FirewallPolicyRuleMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirewallPolicyRuleMatch{}
	r.SrcIPRanges = dcl.FlattenStringSlice(m["srcIpRanges"])
	r.DestIPRanges = dcl.FlattenStringSlice(m["destIpRanges"])
	r.Layer4Configs = flattenFirewallPolicyRuleMatchLayer4ConfigsSlice(c, m["layer4Configs"])
	r.SrcSecureLabels = dcl.FlattenStringSlice(m["srcSecureLabels"])

	return r
}

// expandFirewallPolicyRuleMatchLayer4ConfigsMap expands the contents of FirewallPolicyRuleMatchLayer4Configs into a JSON
// request object.
func expandFirewallPolicyRuleMatchLayer4ConfigsMap(c *Client, f map[string]FirewallPolicyRuleMatchLayer4Configs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirewallPolicyRuleMatchLayer4Configs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirewallPolicyRuleMatchLayer4ConfigsSlice expands the contents of FirewallPolicyRuleMatchLayer4Configs into a JSON
// request object.
func expandFirewallPolicyRuleMatchLayer4ConfigsSlice(c *Client, f []FirewallPolicyRuleMatchLayer4Configs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirewallPolicyRuleMatchLayer4Configs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirewallPolicyRuleMatchLayer4ConfigsMap flattens the contents of FirewallPolicyRuleMatchLayer4Configs from a JSON
// response object.
func flattenFirewallPolicyRuleMatchLayer4ConfigsMap(c *Client, i interface{}) map[string]FirewallPolicyRuleMatchLayer4Configs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallPolicyRuleMatchLayer4Configs{}
	}

	if len(a) == 0 {
		return map[string]FirewallPolicyRuleMatchLayer4Configs{}
	}

	items := make(map[string]FirewallPolicyRuleMatchLayer4Configs)
	for k, item := range a {
		items[k] = *flattenFirewallPolicyRuleMatchLayer4Configs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFirewallPolicyRuleMatchLayer4ConfigsSlice flattens the contents of FirewallPolicyRuleMatchLayer4Configs from a JSON
// response object.
func flattenFirewallPolicyRuleMatchLayer4ConfigsSlice(c *Client, i interface{}) []FirewallPolicyRuleMatchLayer4Configs {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallPolicyRuleMatchLayer4Configs{}
	}

	if len(a) == 0 {
		return []FirewallPolicyRuleMatchLayer4Configs{}
	}

	items := make([]FirewallPolicyRuleMatchLayer4Configs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallPolicyRuleMatchLayer4Configs(c, item.(map[string]interface{})))
	}

	return items
}

// expandFirewallPolicyRuleMatchLayer4Configs expands an instance of FirewallPolicyRuleMatchLayer4Configs into a JSON
// request object.
func expandFirewallPolicyRuleMatchLayer4Configs(c *Client, f *FirewallPolicyRuleMatchLayer4Configs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPProtocol; !dcl.IsEmptyValueIndirect(v) {
		m["ipProtocol"] = v
	}
	if v := f.Ports; !dcl.IsEmptyValueIndirect(v) {
		m["ports"] = v
	}

	return m, nil
}

// flattenFirewallPolicyRuleMatchLayer4Configs flattens an instance of FirewallPolicyRuleMatchLayer4Configs from a JSON
// response object.
func flattenFirewallPolicyRuleMatchLayer4Configs(c *Client, i interface{}) *FirewallPolicyRuleMatchLayer4Configs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirewallPolicyRuleMatchLayer4Configs{}
	r.IPProtocol = dcl.FlattenString(m["ipProtocol"])
	r.Ports = dcl.FlattenStringSlice(m["ports"])

	return r
}

// flattenFirewallPolicyRuleDirectionEnumSlice flattens the contents of FirewallPolicyRuleDirectionEnum from a JSON
// response object.
func flattenFirewallPolicyRuleDirectionEnumSlice(c *Client, i interface{}) []FirewallPolicyRuleDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallPolicyRuleDirectionEnum{}
	}

	if len(a) == 0 {
		return []FirewallPolicyRuleDirectionEnum{}
	}

	items := make([]FirewallPolicyRuleDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallPolicyRuleDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenFirewallPolicyRuleDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *FirewallPolicyRuleDirectionEnum with the same value as that string.
func flattenFirewallPolicyRuleDirectionEnum(i interface{}) *FirewallPolicyRuleDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return FirewallPolicyRuleDirectionEnumRef("")
	}

	return FirewallPolicyRuleDirectionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FirewallPolicyRule) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFirewallPolicyRule(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.FirewallPolicy == nil && ncr.FirewallPolicy == nil {
			c.Config.Logger.Info("Both FirewallPolicy fields null - considering equal.")
		} else if nr.FirewallPolicy == nil || ncr.FirewallPolicy == nil {
			c.Config.Logger.Info("Only one FirewallPolicy field is null - considering unequal.")
			return false
		} else if *nr.FirewallPolicy != *ncr.FirewallPolicy {
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
