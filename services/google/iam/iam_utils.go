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
package iam

import (
	"regexp"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// EncodeIAMCreateRequest encodes the create request for an iam resource.
func EncodeIAMCreateRequest(m map[string]interface{}, resourceName, idField string) map[string]interface{} {
	req := make(map[string]interface{})
	// Put base object into object field.
	dcl.PutMapEntry(req, []string{resourceName}, m)
	// Move name field from from nested object to id field.
	dcl.MoveMapEntry(req, []string{resourceName, "name"}, []string{idField})
	// Delete projectID field.
	delete(req[resourceName].(map[string]interface{}), "projectId")
	// Change value of id field to only last part after / and before @.
	idParts := regexp.MustCompile("([^@/]+)[^/]*$").FindStringSubmatch(*req[idField].(*string))
	if len(idParts) < 2 {
		return req
	}
	req[idField] = idParts[1]
	return req
}

// EncodeRoleCreateRequest properly encodes the create request for an iam role.
func EncodeRoleCreateRequest(m map[string]interface{}) map[string]interface{} {
	return EncodeIAMCreateRequest(m, "role", "roleId")
}

// EncodeServiceAccountCreateRequest properly encodes the create request for an iam service account.
func EncodeServiceAccountCreateRequest(m map[string]interface{}) map[string]interface{} {
	return EncodeIAMCreateRequest(m, "serviceAccount", "accountId")
}
