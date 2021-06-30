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
// Package binaryauthorization provides types and functiosn for handling binaryauthorization GCP resources.
package binaryauthorization

import (
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Returns a copy of the given map with the string "spiffe://" removed from its keys.
func withoutSpiffe(isiar map[string]PolicyAdmissionRule) map[string]PolicyAdmissionRule {
	resultISIAR := make(map[string]PolicyAdmissionRule, len(isiar))
	for k, v := range isiar {
		resultISIAR[strings.TrimPrefix(k, "spiffe://")] = v
	}
	return resultISIAR
}

func equalsPolicyISIAR(m, n map[string]PolicyAdmissionRule) bool {
	if m == nil && n == nil {
		return true
	}
	m = withoutSpiffe(m)
	n = withoutSpiffe(n)
	ds, err := dcl.Diff(m, n, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, dcl.FieldName{})
	return len(ds) == 0 && err == nil
}

// Compares two values of istioServiceIdentity
func canonicalizePolicyISIAR(m, n interface{}) bool {
	mVal, _ := m.(map[string]PolicyAdmissionRule)
	nVal, _ := n.(map[string]PolicyAdmissionRule)
	return equalsPolicyISIAR(mVal, nVal)
}
