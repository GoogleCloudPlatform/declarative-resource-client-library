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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Policy struct {
	AdmissionWhitelistPatterns             []PolicyAdmissionWhitelistPatterns    `json:"admissionWhitelistPatterns"`
	ClusterAdmissionRules                  map[string]PolicyAdmissionRule        `json:"clusterAdmissionRules"`
	KubernetesNamespaceAdmissionRules      map[string]PolicyAdmissionRule        `json:"kubernetesNamespaceAdmissionRules"`
	KubernetesServiceAccountAdmissionRules map[string]PolicyAdmissionRule        `json:"kubernetesServiceAccountAdmissionRules"`
	IstioServiceIdentityAdmissionRules     map[string]PolicyAdmissionRule        `json:"istioServiceIdentityAdmissionRules"`
	DefaultAdmissionRule                   *PolicyAdmissionRule                  `json:"defaultAdmissionRule"`
	Description                            *string                               `json:"description"`
	GlobalPolicyEvaluationMode             *PolicyGlobalPolicyEvaluationModeEnum `json:"globalPolicyEvaluationMode"`
	SelfLink                               *string                               `json:"selfLink"`
	Project                                *string                               `json:"project"`
	UpdateTime                             *string                               `json:"updateTime"`
}

func (r *Policy) String() string {
	return dcl.SprintResource(r)
}

// The enum PolicyAdmissionRuleEvaluationModeEnum.
type PolicyAdmissionRuleEvaluationModeEnum string

// PolicyAdmissionRuleEvaluationModeEnumRef returns a *PolicyAdmissionRuleEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyAdmissionRuleEvaluationModeEnumRef(s string) *PolicyAdmissionRuleEvaluationModeEnum {
	if s == "" {
		return nil
	}

	v := PolicyAdmissionRuleEvaluationModeEnum(s)
	return &v
}

func (v PolicyAdmissionRuleEvaluationModeEnum) Validate() error {
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyAdmissionRuleEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyAdmissionRuleEnforcementModeEnum.
type PolicyAdmissionRuleEnforcementModeEnum string

// PolicyAdmissionRuleEnforcementModeEnumRef returns a *PolicyAdmissionRuleEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyAdmissionRuleEnforcementModeEnumRef(s string) *PolicyAdmissionRuleEnforcementModeEnum {
	if s == "" {
		return nil
	}

	v := PolicyAdmissionRuleEnforcementModeEnum(s)
	return &v
}

func (v PolicyAdmissionRuleEnforcementModeEnum) Validate() error {
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyAdmissionRuleEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyGlobalPolicyEvaluationModeEnum.
type PolicyGlobalPolicyEvaluationModeEnum string

// PolicyGlobalPolicyEvaluationModeEnumRef returns a *PolicyGlobalPolicyEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyGlobalPolicyEvaluationModeEnumRef(s string) *PolicyGlobalPolicyEvaluationModeEnum {
	if s == "" {
		return nil
	}

	v := PolicyGlobalPolicyEvaluationModeEnum(s)
	return &v
}

func (v PolicyGlobalPolicyEvaluationModeEnum) Validate() error {
	for _, s := range []string{"GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED", "ENABLE", "DISABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyGlobalPolicyEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type PolicyAdmissionWhitelistPatterns struct {
	empty       bool    `json:"-"`
	NamePattern *string `json:"namePattern"`
}

type jsonPolicyAdmissionWhitelistPatterns PolicyAdmissionWhitelistPatterns

func (r *PolicyAdmissionWhitelistPatterns) UnmarshalJSON(data []byte) error {
	var res jsonPolicyAdmissionWhitelistPatterns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyAdmissionWhitelistPatterns
	} else {

		r.NamePattern = res.NamePattern

	}
	return nil
}

// This object is used to assert a desired state where this PolicyAdmissionWhitelistPatterns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyPolicyAdmissionWhitelistPatterns *PolicyAdmissionWhitelistPatterns = &PolicyAdmissionWhitelistPatterns{empty: true}

func (r *PolicyAdmissionWhitelistPatterns) Empty() bool {
	return r.empty
}

func (r *PolicyAdmissionWhitelistPatterns) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyAdmissionWhitelistPatterns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyAdmissionRule struct {
	empty                 bool                                    `json:"-"`
	EvaluationMode        *PolicyAdmissionRuleEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyAdmissionRuleEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyAdmissionRule PolicyAdmissionRule

func (r *PolicyAdmissionRule) UnmarshalJSON(data []byte) error {
	var res jsonPolicyAdmissionRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyAdmissionRule
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyAdmissionRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyPolicyAdmissionRule *PolicyAdmissionRule = &PolicyAdmissionRule{empty: true}

func (r *PolicyAdmissionRule) Empty() bool {
	return r.empty
}

func (r *PolicyAdmissionRule) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyAdmissionRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Policy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "binary_authorization",
		Type:    "Policy",
		Version: "beta",
	}
}

const PolicyMaxPage = -1

type PolicyList struct {
	Items []*Policy

	nextToken string

	resource *Policy
}

func (c *Client) GetPolicy(ctx context.Context, r *Policy) (*Policy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizePolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyPolicy(ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (*Policy, error) {
	var resultNewState *Policy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyPolicyHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyPolicyHelper(c *Client, ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (*Policy, error) {
	c.Config.Logger.Info("Beginning ApplyPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.policyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToPolicyDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		return nil, dcl.ApplyInfeasibleError{Message: "No initial state found for singleton resource."}
	} else {
		for _, d := range diffs {
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}
	var ops []policyApiOperation
	for _, d := range diffs {
		ops = append(ops, d.UpdateOp)
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizePolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizePolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffPolicy(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
func (r *Policy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
