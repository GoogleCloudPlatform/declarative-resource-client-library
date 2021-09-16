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
// Package orgpolicy defines types and methods for working with orgpolicy GCP resources.
package alpha

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func expandPolicyName(r *Policy, name *string) (*string, error) {
	nameParts := strings.Split(dcl.ValueOrEmptyString(name), "/")
	if len(nameParts) == 4 {
		fullName := strings.Join(nameParts, "/")
		return &fullName, nil
	}
	shortName := nameParts[len(nameParts)-1]
	fullName := fmt.Sprintf("%s/policies/%s", dcl.ValueOrEmptyString(r.Parent), shortName)
	return &fullName, nil
}

func equalsPolicyName(m, n *string) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil || n == nil {
		return false
	}
	return *dcl.SelfLinkToName(m) == *dcl.SelfLinkToName(n)
}

// Compares two values of policy name. Custom diff function required because API returns project numbers.
func canonicalizePolicyName(m, n interface{}) bool {
	mString, ok := m.(*string)
	if !ok {
		return false
	}
	nString, ok := n.(*string)
	if !ok {
		return false
	}
	return equalsPolicyName(mString, nString)
}

func equalsPolicyRulesConditionExpression(m, n *string) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil || n == nil {
		return false
	}
	mReplaced := strings.ReplaceAll(strings.ReplaceAll(*m, "Labels", "TagId"), "label", "tag")
	nReplaced := strings.ReplaceAll(strings.ReplaceAll(*n, "Labels", "TagId"), "label", "tag")
	return mReplaced == nReplaced
}

// Compares two values of policy rules condition expression. Custom diff function required due to API substitutions.
func canonicalizePolicyRulesConditionExpression(m, n interface{}) bool {
	mString, ok := m.(*string)
	if !ok {
		return false
	}
	nString, ok := n.(*string)
	if !ok {
		return false
	}
	return equalsPolicyRulesConditionExpression(mString, nString)
}

// TODO(magic-modules-eng): Convert this behavior to a field trait so that it can be reused.
// expandPolicyRulesAllowAll expands the AllowAll enum into a boolean to be sent to the API.
func expandPolicyRulesAllowAll(_ *PolicySpecRules, v *PolicySpecRulesAllowAllEnum) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	switch *v {
	case *PolicySpecRulesAllowAllEnumRef(""):
		return nil, nil
	case *PolicySpecRulesAllowAllEnumRef("TRUE"):
		return dcl.Bool(true), nil
	case *PolicySpecRulesAllowAllEnumRef("FALSE"):
		return dcl.Bool(false), nil
	}
	return nil, fmt.Errorf("unknown value for AllowAll: %v", *v)
}

// flattenPolicyRulesAllowAll converts the boolean value returned by the API into the
// three option enum value for AllowAll.
func flattenPolicyRulesAllowAll(i interface{}) *PolicySpecRulesAllowAllEnum {
	v, ok := i.(bool)
	if !ok {
		return nil
	}
	if v == true {
		return PolicySpecRulesAllowAllEnumRef("TRUE")
	}
	return PolicySpecRulesAllowAllEnumRef("FALSE")
}

// expandPolicyRulesDenyAll expands the DenyAll enum into a boolean to be sent to the API.
func expandPolicyRulesDenyAll(_ *PolicySpecRules, v *PolicySpecRulesDenyAllEnum) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	switch *v {
	case *PolicySpecRulesDenyAllEnumRef(""):
		return nil, nil
	case *PolicySpecRulesDenyAllEnumRef("TRUE"):
		return dcl.Bool(true), nil
	case *PolicySpecRulesDenyAllEnumRef("FALSE"):
		return dcl.Bool(false), nil
	}
	return nil, fmt.Errorf("unknown value for DenyAll: %v", *v)
}

// flattenPolicyRulesDenyAll converts the boolean value returned by the API into the
// three option enum value for DenyAll.
func flattenPolicyRulesDenyAll(i interface{}) *PolicySpecRulesDenyAllEnum {
	v, ok := i.(bool)
	if !ok {
		return nil
	}
	if v == true {
		return PolicySpecRulesDenyAllEnumRef("TRUE")
	}
	return PolicySpecRulesDenyAllEnumRef("FALSE")
}

// expandPolicyRulesEnforce expands the Enforce enum into a boolean to be sent to the API.
func expandPolicyRulesEnforce(_ *PolicySpecRules, v *PolicySpecRulesEnforceEnum) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	switch *v {
	case *PolicySpecRulesEnforceEnumRef(""):
		return nil, nil
	case *PolicySpecRulesEnforceEnumRef("TRUE"):
		return dcl.Bool(true), nil
	case *PolicySpecRulesEnforceEnumRef("FALSE"):
		return dcl.Bool(false), nil
	}
	return nil, fmt.Errorf("unknown value for Enforce: %v", *v)
}

// flattenPolicyRulesEnforce converts the boolean value returned by the API into the
// three option enum value for Enforce.
func flattenPolicyRulesEnforce(i interface{}) *PolicySpecRulesEnforceEnum {
	v, ok := i.(bool)
	if !ok {
		return nil
	}
	if v == true {
		return PolicySpecRulesEnforceEnumRef("TRUE")
	}
	return PolicySpecRulesEnforceEnumRef("FALSE")
}
