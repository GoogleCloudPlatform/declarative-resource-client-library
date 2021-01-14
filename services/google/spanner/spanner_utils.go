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
package spanner

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// EncodeSpannerInstanceCreateRequest encapsulates fields other than instanceId
// into an instance {} block, as expected by https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances/create
func EncodeSpannerInstanceCreateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	dcl.PutMapEntry(req, []string{"instance"}, m)
	dcl.MoveMapEntry(req, []string{"instance", "instanceId"}, []string{"instanceId"})
	return req
}

// EncodeSpannerInstanceUpdateRequest encapsulates fields other than instanceI
// and nodeCount
// into an instance {} block, as expected by https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances/patch
// Additionally, it adds a nonstandard update mask called field_mask aliased to
// all of the updatable fields.
func EncodeSpannerInstanceUpdateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	dcl.PutMapEntry(req, []string{"instance"}, m)
	// Hardcode all fields- we can't tell what specific fields are affected
	// inside an encoder.
	req["field_mask"] = "config,displayName,labels,nodeCount"
	return req
}
