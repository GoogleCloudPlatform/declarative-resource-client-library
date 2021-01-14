// Copyright 2020 Google LLC. All Rights Reserved.
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
package sql

import (
	"strings"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// IsNotFoundOrInstanceNotRunning returns true if the resource is not found or if the resource returns
// 400 bad request and contains a string which indicates that the issue is that the instance is down.
func IsNotFoundOrInstanceNotRunning(err error) bool {
	if dcl.IsNotFound(err) {
		return true
	}
	if gErr, ok := err.(*googleapi.Error); ok && gErr.Code == 400 && strings.Contains(gErr.Message, "Invalid request since instance is not running") {
		return true
	}
	return false
}
