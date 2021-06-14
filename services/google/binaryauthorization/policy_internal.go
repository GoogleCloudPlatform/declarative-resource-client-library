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
package binaryauthorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Policy) validate() error {

	if err := dcl.Required(r, "defaultAdmissionRule"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultAdmissionRule) {
		if err := r.DefaultAdmissionRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *PolicyAdmissionWhitelistPatterns) validate() error {
	return nil
}
func (r *PolicyClusterAdmissionRules) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *PolicyDefaultAdmissionRule) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}

func policyGetURL(userBasePath string, r *Policy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
	}
	return dcl.URL("projects/{{project}}/policy", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil
}

func policyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/policy", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil

}

func policyCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/policy", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil

}

// policyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type policyApiOperation interface {
	do(context.Context, *Policy, *Client) error
}

func (c *Client) listPolicyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := policyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != PolicyMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listPolicyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listPolicy(ctx context.Context, project, pageToken string, pageSize int32) ([]*Policy, string, error) {
	b, err := c.listPolicyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Policy
	for _, v := range m.Items {
		res, err := unmarshalMapPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createPolicyOperation struct {
	response map[string]interface{}
}

func (op *createPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createPolicyOperation) do(ctx context.Context, r *Policy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := policyCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getPolicyRaw(ctx context.Context, r *Policy) ([]byte, error) {

	u, err := policyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) policyDiffsForRawDesired(ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (initial, desired *Policy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Policy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Policy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Policy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// Simulate the resource not existing because create operation should be called anyway.
	rawInitial := &Policy{}
	desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial)
	return nil, desired, nil, err

	c.Config.Logger.Infof("Found initial state for Policy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Policy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizePolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Policy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Policy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizePolicyInitialState(rawInitial, rawDesired *Policy) (*Policy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizePolicyDesiredState(rawDesired, rawInitial *Policy, opts ...dcl.ApplyOption) (*Policy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultAdmissionRule = canonicalizePolicyDefaultAdmissionRule(rawDesired.DefaultAdmissionRule, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.AdmissionWhitelistPatterns) {
		rawDesired.AdmissionWhitelistPatterns = rawInitial.AdmissionWhitelistPatterns
	}
	if dcl.IsZeroValue(rawDesired.ClusterAdmissionRules) {
		rawDesired.ClusterAdmissionRules = rawInitial.ClusterAdmissionRules
	}
	rawDesired.DefaultAdmissionRule = canonicalizePolicyDefaultAdmissionRule(rawDesired.DefaultAdmissionRule, rawInitial.DefaultAdmissionRule, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.GlobalPolicyEvaluationMode) {
		rawDesired.GlobalPolicyEvaluationMode = rawInitial.GlobalPolicyEvaluationMode
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizePolicyNewState(c *Client, rawNew, rawDesired *Policy) (*Policy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.AdmissionWhitelistPatterns) && dcl.IsEmptyValueIndirect(rawDesired.AdmissionWhitelistPatterns) {
		rawNew.AdmissionWhitelistPatterns = rawDesired.AdmissionWhitelistPatterns
	} else {
		rawNew.AdmissionWhitelistPatterns = canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c, rawDesired.AdmissionWhitelistPatterns, rawNew.AdmissionWhitelistPatterns)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClusterAdmissionRules) && dcl.IsEmptyValueIndirect(rawDesired.ClusterAdmissionRules) {
		rawNew.ClusterAdmissionRules = rawDesired.ClusterAdmissionRules
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultAdmissionRule) && dcl.IsEmptyValueIndirect(rawDesired.DefaultAdmissionRule) {
		rawNew.DefaultAdmissionRule = rawDesired.DefaultAdmissionRule
	} else {
		rawNew.DefaultAdmissionRule = canonicalizeNewPolicyDefaultAdmissionRule(c, rawDesired.DefaultAdmissionRule, rawNew.DefaultAdmissionRule)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GlobalPolicyEvaluationMode) && dcl.IsEmptyValueIndirect(rawDesired.GlobalPolicyEvaluationMode) {
		rawNew.GlobalPolicyEvaluationMode = rawDesired.GlobalPolicyEvaluationMode
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	return rawNew, nil
}

func canonicalizePolicyAdmissionWhitelistPatterns(des, initial *PolicyAdmissionWhitelistPatterns, opts ...dcl.ApplyOption) *PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.NamePattern, initial.NamePattern) || dcl.IsZeroValue(des.NamePattern) {
		des.NamePattern = initial.NamePattern
	}

	return des
}

func canonicalizeNewPolicyAdmissionWhitelistPatterns(c *Client, des, nw *PolicyAdmissionWhitelistPatterns) *PolicyAdmissionWhitelistPatterns {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.NamePattern, nw.NamePattern) {
		nw.NamePattern = des.NamePattern
	}

	return nw
}

func canonicalizeNewPolicyAdmissionWhitelistPatternsSet(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}
	var reorderedNew []PolicyAdmissionWhitelistPatterns
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePolicyAdmissionWhitelistPatternsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyAdmissionWhitelistPatterns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyAdmissionWhitelistPatterns(c, &d, &n))
	}

	return items
}

func canonicalizePolicyClusterAdmissionRules(des, initial *PolicyClusterAdmissionRules, opts ...dcl.ApplyOption) *PolicyClusterAdmissionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EvaluationMode) {
		des.EvaluationMode = initial.EvaluationMode
	}
	if dcl.IsZeroValue(des.RequireAttestationsBy) {
		des.RequireAttestationsBy = initial.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) {
		des.EnforcementMode = initial.EnforcementMode
	}

	return des
}

func canonicalizeNewPolicyClusterAdmissionRules(c *Client, des, nw *PolicyClusterAdmissionRules) *PolicyClusterAdmissionRules {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.EvaluationMode) {
		nw.EvaluationMode = des.EvaluationMode
	}
	if dcl.IsZeroValue(nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(nw.EnforcementMode) {
		nw.EnforcementMode = des.EnforcementMode
	}

	return nw
}

func canonicalizeNewPolicyClusterAdmissionRulesSet(c *Client, des, nw []PolicyClusterAdmissionRules) []PolicyClusterAdmissionRules {
	if des == nil {
		return nw
	}
	var reorderedNew []PolicyClusterAdmissionRules
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePolicyClusterAdmissionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewPolicyClusterAdmissionRulesSlice(c *Client, des, nw []PolicyClusterAdmissionRules) []PolicyClusterAdmissionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyClusterAdmissionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyClusterAdmissionRules(c, &d, &n))
	}

	return items
}

func canonicalizePolicyDefaultAdmissionRule(des, initial *PolicyDefaultAdmissionRule, opts ...dcl.ApplyOption) *PolicyDefaultAdmissionRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EvaluationMode) {
		des.EvaluationMode = initial.EvaluationMode
	}
	if dcl.IsZeroValue(des.RequireAttestationsBy) {
		des.RequireAttestationsBy = initial.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) {
		des.EnforcementMode = initial.EnforcementMode
	}

	return des
}

func canonicalizeNewPolicyDefaultAdmissionRule(c *Client, des, nw *PolicyDefaultAdmissionRule) *PolicyDefaultAdmissionRule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.EvaluationMode) {
		nw.EvaluationMode = des.EvaluationMode
	}
	if dcl.IsZeroValue(nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(nw.EnforcementMode) {
		nw.EnforcementMode = des.EnforcementMode
	}

	return nw
}

func canonicalizeNewPolicyDefaultAdmissionRuleSet(c *Client, des, nw []PolicyDefaultAdmissionRule) []PolicyDefaultAdmissionRule {
	if des == nil {
		return nw
	}
	var reorderedNew []PolicyDefaultAdmissionRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePolicyDefaultAdmissionRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewPolicyDefaultAdmissionRuleSlice(c *Client, des, nw []PolicyDefaultAdmissionRule) []PolicyDefaultAdmissionRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyDefaultAdmissionRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyDefaultAdmissionRule(c, &d, &n))
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
func diffPolicy(c *Client, desired, actual *Policy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.AdmissionWhitelistPatterns, actual.AdmissionWhitelistPatterns, dcl.Info{ObjectFunction: comparePolicyAdmissionWhitelistPatternsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdmissionWhitelistPatterns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultAdmissionRule, actual.DefaultAdmissionRule, dcl.Info{ObjectFunction: comparePolicyDefaultAdmissionRuleNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultAdmissionRule")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.GlobalPolicyEvaluationMode, actual.GlobalPolicyEvaluationMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GlobalPolicyEvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func comparePolicyAdmissionWhitelistPatternsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		desiredNotPointer, ok := d.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns or *PolicyAdmissionWhitelistPatterns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		actualNotPointer, ok := a.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NamePattern, actual.NamePattern, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NamePattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyClusterAdmissionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyClusterAdmissionRules)
	if !ok {
		desiredNotPointer, ok := d.(PolicyClusterAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyClusterAdmissionRules or *PolicyClusterAdmissionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyClusterAdmissionRules)
	if !ok {
		actualNotPointer, ok := a.(PolicyClusterAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyClusterAdmissionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyDefaultAdmissionRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyDefaultAdmissionRule)
	if !ok {
		desiredNotPointer, ok := d.(PolicyDefaultAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyDefaultAdmissionRule or *PolicyDefaultAdmissionRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyDefaultAdmissionRule)
	if !ok {
		actualNotPointer, ok := a.(PolicyDefaultAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyDefaultAdmissionRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
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
func (r *Policy) urlNormalized() *Policy {
	normalized := dcl.Copy(*r).(Policy)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Policy) getFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Policy) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Policy) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Policy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Policy) marshal(c *Client) ([]byte, error) {
	m, err := expandPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Policy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalPolicy decodes JSON responses into the Policy resource schema.
func unmarshalPolicy(b []byte, c *Client) (*Policy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapPolicy(m, c)
}

func unmarshalMapPolicy(m map[string]interface{}, c *Client) (*Policy, error) {

	return flattenPolicy(c, m), nil
}

// expandPolicy expands Policy into a JSON request object.
func expandPolicy(c *Client, f *Policy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := expandPolicyAdmissionWhitelistPatternsSlice(c, f.AdmissionWhitelistPatterns); err != nil {
		return nil, fmt.Errorf("error expanding AdmissionWhitelistPatterns into admissionWhitelistPatterns: %w", err)
	} else if v != nil {
		m["admissionWhitelistPatterns"] = v
	}
	if v, err := expandPolicyClusterAdmissionRulesMap(c, f.ClusterAdmissionRules); err != nil {
		return nil, fmt.Errorf("error expanding ClusterAdmissionRules into clusterAdmissionRules: %w", err)
	} else if v != nil {
		m["clusterAdmissionRules"] = v
	}
	if v, err := expandPolicyDefaultAdmissionRule(c, f.DefaultAdmissionRule); err != nil {
		return nil, fmt.Errorf("error expanding DefaultAdmissionRule into defaultAdmissionRule: %w", err)
	} else if v != nil {
		m["defaultAdmissionRule"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.GlobalPolicyEvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["globalPolicyEvaluationMode"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/policy", f.SelfLink, f.Project); err != nil {
		return nil, fmt.Errorf("error expanding SelfLink into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenPolicy flattens Policy from a JSON request object into the
// Policy type.
func flattenPolicy(c *Client, i interface{}) *Policy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Policy{}
	res.AdmissionWhitelistPatterns = flattenPolicyAdmissionWhitelistPatternsSlice(c, m["admissionWhitelistPatterns"])
	res.ClusterAdmissionRules = flattenPolicyClusterAdmissionRulesMap(c, m["clusterAdmissionRules"])
	res.DefaultAdmissionRule = flattenPolicyDefaultAdmissionRule(c, m["defaultAdmissionRule"])
	res.Description = dcl.FlattenString(m["description"])
	res.GlobalPolicyEvaluationMode = flattenPolicyGlobalPolicyEvaluationModeEnum(m["globalPolicyEvaluationMode"])
	res.SelfLink = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])

	return res
}

// expandPolicyAdmissionWhitelistPatternsMap expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsMap(c *Client, f map[string]PolicyAdmissionWhitelistPatterns) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyAdmissionWhitelistPatternsSlice expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsSlice(c *Client, f []PolicyAdmissionWhitelistPatterns) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyAdmissionWhitelistPatternsMap flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsMap(c *Client, i interface{}) map[string]PolicyAdmissionWhitelistPatterns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	items := make(map[string]PolicyAdmissionWhitelistPatterns)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{}))
	}

	return items
}

// flattenPolicyAdmissionWhitelistPatternsSlice flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsSlice(c *Client, i interface{}) []PolicyAdmissionWhitelistPatterns {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	items := make([]PolicyAdmissionWhitelistPatterns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{})))
	}

	return items
}

// expandPolicyAdmissionWhitelistPatterns expands an instance of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatterns(c *Client, f *PolicyAdmissionWhitelistPatterns) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NamePattern; !dcl.IsEmptyValueIndirect(v) {
		m["namePattern"] = v
	}

	return m, nil
}

// flattenPolicyAdmissionWhitelistPatterns flattens an instance of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatterns(c *Client, i interface{}) *PolicyAdmissionWhitelistPatterns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyAdmissionWhitelistPatterns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyAdmissionWhitelistPatterns
	}
	r.NamePattern = dcl.FlattenString(m["namePattern"])

	return r
}

// expandPolicyClusterAdmissionRulesMap expands the contents of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRulesMap(c *Client, f map[string]PolicyClusterAdmissionRules) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyClusterAdmissionRules(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyClusterAdmissionRulesSlice expands the contents of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRulesSlice(c *Client, f []PolicyClusterAdmissionRules) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyClusterAdmissionRules(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyClusterAdmissionRulesMap flattens the contents of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesMap(c *Client, i interface{}) map[string]PolicyClusterAdmissionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyClusterAdmissionRules{}
	}

	if len(a) == 0 {
		return map[string]PolicyClusterAdmissionRules{}
	}

	items := make(map[string]PolicyClusterAdmissionRules)
	for k, item := range a {
		items[k] = *flattenPolicyClusterAdmissionRules(c, item.(map[string]interface{}))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesSlice flattens the contents of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesSlice(c *Client, i interface{}) []PolicyClusterAdmissionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRules{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRules{}
	}

	items := make([]PolicyClusterAdmissionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRules(c, item.(map[string]interface{})))
	}

	return items
}

// expandPolicyClusterAdmissionRules expands an instance of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRules(c *Client, f *PolicyClusterAdmissionRules) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; !dcl.IsEmptyValueIndirect(v) {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyClusterAdmissionRules flattens an instance of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRules(c *Client, i interface{}) *PolicyClusterAdmissionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyClusterAdmissionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyClusterAdmissionRules
	}
	r.EvaluationMode = flattenPolicyClusterAdmissionRulesEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyClusterAdmissionRulesEnforcementModeEnum(m["enforcementMode"])

	return r
}

// expandPolicyDefaultAdmissionRuleMap expands the contents of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRuleMap(c *Client, f map[string]PolicyDefaultAdmissionRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyDefaultAdmissionRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyDefaultAdmissionRuleSlice expands the contents of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRuleSlice(c *Client, f []PolicyDefaultAdmissionRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyDefaultAdmissionRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyDefaultAdmissionRuleMap flattens the contents of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleMap(c *Client, i interface{}) map[string]PolicyDefaultAdmissionRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyDefaultAdmissionRule{}
	}

	if len(a) == 0 {
		return map[string]PolicyDefaultAdmissionRule{}
	}

	items := make(map[string]PolicyDefaultAdmissionRule)
	for k, item := range a {
		items[k] = *flattenPolicyDefaultAdmissionRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleSlice flattens the contents of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleSlice(c *Client, i interface{}) []PolicyDefaultAdmissionRule {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRule{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRule{}
	}

	items := make([]PolicyDefaultAdmissionRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandPolicyDefaultAdmissionRule expands an instance of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRule(c *Client, f *PolicyDefaultAdmissionRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; !dcl.IsEmptyValueIndirect(v) {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyDefaultAdmissionRule flattens an instance of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRule(c *Client, i interface{}) *PolicyDefaultAdmissionRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyDefaultAdmissionRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyDefaultAdmissionRule
	}
	r.EvaluationMode = flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(m["enforcementMode"])

	return r
}

// flattenPolicyClusterAdmissionRulesEvaluationModeEnumSlice flattens the contents of PolicyClusterAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEvaluationModeEnumSlice(c *Client, i interface{}) []PolicyClusterAdmissionRulesEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	items := make([]PolicyClusterAdmissionRulesEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRulesEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyClusterAdmissionRulesEvaluationModeEnum with the same value as that string.
func flattenPolicyClusterAdmissionRulesEvaluationModeEnum(i interface{}) *PolicyClusterAdmissionRulesEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyClusterAdmissionRulesEvaluationModeEnumRef("")
	}

	return PolicyClusterAdmissionRulesEvaluationModeEnumRef(s)
}

// flattenPolicyClusterAdmissionRulesEnforcementModeEnumSlice flattens the contents of PolicyClusterAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEnforcementModeEnumSlice(c *Client, i interface{}) []PolicyClusterAdmissionRulesEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	items := make([]PolicyClusterAdmissionRulesEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRulesEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyClusterAdmissionRulesEnforcementModeEnum with the same value as that string.
func flattenPolicyClusterAdmissionRulesEnforcementModeEnum(i interface{}) *PolicyClusterAdmissionRulesEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyClusterAdmissionRulesEnforcementModeEnumRef("")
	}

	return PolicyClusterAdmissionRulesEnforcementModeEnumRef(s)
}

// flattenPolicyDefaultAdmissionRuleEvaluationModeEnumSlice flattens the contents of PolicyDefaultAdmissionRuleEvaluationModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEvaluationModeEnumSlice(c *Client, i interface{}) []PolicyDefaultAdmissionRuleEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	items := make([]PolicyDefaultAdmissionRuleEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyDefaultAdmissionRuleEvaluationModeEnum with the same value as that string.
func flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(i interface{}) *PolicyDefaultAdmissionRuleEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyDefaultAdmissionRuleEvaluationModeEnumRef("")
	}

	return PolicyDefaultAdmissionRuleEvaluationModeEnumRef(s)
}

// flattenPolicyDefaultAdmissionRuleEnforcementModeEnumSlice flattens the contents of PolicyDefaultAdmissionRuleEnforcementModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEnforcementModeEnumSlice(c *Client, i interface{}) []PolicyDefaultAdmissionRuleEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	items := make([]PolicyDefaultAdmissionRuleEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyDefaultAdmissionRuleEnforcementModeEnum with the same value as that string.
func flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(i interface{}) *PolicyDefaultAdmissionRuleEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyDefaultAdmissionRuleEnforcementModeEnumRef("")
	}

	return PolicyDefaultAdmissionRuleEnforcementModeEnumRef(s)
}

// flattenPolicyGlobalPolicyEvaluationModeEnumSlice flattens the contents of PolicyGlobalPolicyEvaluationModeEnum from a JSON
// response object.
func flattenPolicyGlobalPolicyEvaluationModeEnumSlice(c *Client, i interface{}) []PolicyGlobalPolicyEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	items := make([]PolicyGlobalPolicyEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyGlobalPolicyEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyGlobalPolicyEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyGlobalPolicyEvaluationModeEnum with the same value as that string.
func flattenPolicyGlobalPolicyEvaluationModeEnum(i interface{}) *PolicyGlobalPolicyEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyGlobalPolicyEvaluationModeEnumRef("")
	}

	return PolicyGlobalPolicyEvaluationModeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Policy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalPolicy(b, c)
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
		return true
	}
}

type policyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         policyApiOperation
}

func convertFieldDiffToPolicyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]policyDiff, error) {
	var diffs []policyDiff
	for _, op := range ops {
		diff := policyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTopolicyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTopolicyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (policyApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
