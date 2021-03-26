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
// Package gkehub includes all of the code for gkehub.
package beta

import (
	"strings"
)

func expandHubReferenceLink(f *MembershipEndpointGkeCluster, val *string) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	v := *val

	if strings.HasPrefix(v, "https:") {
		return strings.Replace(strings.Replace(strings.Replace(*val, "https:", "", 1), "v1/", "", 1), "v1beta1/", "", 1), nil
	}
	return "//container.googleapis.com/" + v, nil
}

func flattenHubReferenceLink(config interface{}) *string {
	v, ok := config.(string)
	if !ok {
		return nil
	}

	v = strings.Replace(v, "//container.googleapis.com/", "", 1)

	return &v
}
