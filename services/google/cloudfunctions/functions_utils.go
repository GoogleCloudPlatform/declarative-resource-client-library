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
package cloudfunctions

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// ExpandCloudFunctionTimeout converts the DCL representation of a Cloud Function timeout
// (an integer representing the number of seconds) to the API representation (a string
// containing a float representing the number of seconds followed by the letter 's').
func ExpandCloudFunctionTimeout(f *CloudFunction, t *int64) (string, error) {
	if t == nil {
		return "", nil
	}
	return fmt.Sprintf("%ds", *t), nil
}

// FlattenCloudFunctionTimeout reverses ExpandCloudFunctionTimeout(), converting the
// string representation to the integer representation. Any fractional part of the
// timeout is stripped, e.g., "7.5s" becomes 7.
func FlattenCloudFunctionTimeout(i interface{}) *int64 {
	if s, ok := i.(string); ok {
		var trimmed string
		for _, r := range s {
			if !unicode.IsNumber(r) {
				break
			}
			trimmed += string(r)
		}
		if trimmed != "" {
			if n, err := strconv.ParseInt(trimmed, 10, 64); err == nil {
				return &n
			}
		}
	}
	return nil
}

// ExpandCloudFunctionEventRetry inverts the FlattenCloudFunctionEventRetry transformation.
func ExpandCloudFunctionEventRetry(f *CloudFunctionEventTrigger, t *bool) (interface{}, error) {
	if t == nil || !*t {
		return nil, nil
	}
	return map[string]interface{}{
		"retry": map[string]interface{}{},
	}, nil
}

// FlattenCloudFunctionEventRetry converts the API reprensentation of an event
// trigger retry policy, which is true or false based on the presence or absence
// of an empty object, to an actual bool for convenience purposes.
func FlattenCloudFunctionEventRetry(i interface{}) *bool {
	if _, ok := i.(map[string]interface{}); ok {
		return dcl.Bool(true)
	}
	return nil
}
