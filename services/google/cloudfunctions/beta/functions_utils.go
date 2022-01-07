// Copyright 2022 Google LLC. All Rights Reserved.
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
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// ExpandFunctionEventResource handles the special case for StorageBucket
// and makes sure resource field has format projects/{{project}}/buckets/{{name}}.
// http://b/210923505
func ExpandFunctionEventResource(f *Function) (*string, error) {
	if f.EventTrigger == nil || f.EventTrigger.Resource == nil {
		return nil, nil
	}

	if dcl.ValueOrEmptyString(f.EventTrigger.EventType) == "providers/cloud.storage/eventTypes/object.change" {
		return dcl.DeriveField("projects/%s/buckets/%s", f.EventTrigger.Resource, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.EventTrigger.Resource))
	}

	return f.EventTrigger.Resource, nil
}

// ExpandFunctionEventRetry inverts the FlattenFunctionEventRetry transformation.
func ExpandFunctionEventRetry(f *Function) (interface{}, error) {
	if f.EventTrigger == nil || f.EventTrigger.FailurePolicy == nil || !*(f.EventTrigger.FailurePolicy) {
		return nil, nil
	}
	return map[string]interface{}{
		"retry": map[string]interface{}{},
	}, nil
}

// FlattenFunctionEventRetry converts the API reprensentation of an event
// trigger retry policy, which is true or false based on the presence or absence
// of an empty object, to an actual bool for convenience purposes.
func FlattenFunctionEventRetry(i interface{}) *bool {
	if _, ok := i.(map[string]interface{}); ok {
		return dcl.Bool(true)
	}
	return nil
}

// CanonicalizeFunctionSourceRepoURL compares source repo url because /paths/ can be omitted but will be added by the API.
func CanonicalizeFunctionSourceRepoURL(m interface{}, n interface{}) bool {
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
	// Both values are pointers to strings. Compare values with suffix trimmed.
	return strings.TrimSuffix(*mVal, "/paths/") == strings.TrimSuffix(*nVal, "/paths/")
}
