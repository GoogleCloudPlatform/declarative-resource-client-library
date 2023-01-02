// Copyright 2023 Google LLC. All Rights Reserved.
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
// Package vmware contains methods and types for handling vmware GCP resources.
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Compare two private cloud network field values by only considering the networks' short names and
// not their partial urls.
func canonicalizePrivateCloudNetwork(n, m interface{}) bool {
	mVal, ok := m.(*string)
	if !ok {
		return false
	}
	nVal, ok := n.(*string)
	if !ok {
		return false
	}
	if mVal == nil && nVal == nil {
		return true
	}
	if mVal == nil || nVal == nil {
		return false
	}
	return dcl.ValueOrEmptyString(dcl.SelfLinkToName(mVal)) == dcl.ValueOrEmptyString(dcl.SelfLinkToName(nVal))
}
