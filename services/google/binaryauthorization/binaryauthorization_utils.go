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

// Removes the string "spiffe://" from the keys of the given map and returns the map.
func removeSpiffe(isiar map[string]PolicyAdmissionRule) {
	for k, v := range isiar {
		if strings.HasPrefix(k, "spiffe://") {
			isiar[k[9:len(k)]] = v
			delete(isiar, k)
		}
	}
}

func equalsPolicyISIAR(m, n map[string]PolicyAdmissionRule) bool {
	if m == nil && n == nil {
		return true
	}
	removeSpiffe(m)
	removeSpiffe(n)
	if ds, err := dcl.Diff(m, n, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, dcl.FieldName{}); len(ds) != 0 || err != nil {
		return false
	}
	return true
}

func canonicalizePolicyISIAR(m, n interface{}) bool {
	mVal, _ := m.(map[string]PolicyAdmissionRule)
	nVal, _ := n.(map[string]PolicyAdmissionRule)
	return equalsPolicyISIAR(mVal, nVal)
}
