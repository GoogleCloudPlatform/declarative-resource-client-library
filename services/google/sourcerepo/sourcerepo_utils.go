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
// Package sourcerepo contains handwritten support code for the SourceRepo service.
package sourcerepo

import (
	"fmt"

	glog "github.com/golang/glog"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// expandRepoPubsubConfigs modifies a map[string]string to a nested map keyed off the name of the pubsub topic
func expandRepoPubsubConfig(repo *Repo, configs []RepoPubsubConfigs) (interface{}, error) {
	transformed := make(map[string]interface{})
	for _, config := range configs {
		topic := config.Topic
		subMap := map[string]string{"topic": *topic}
		if config.MessageFormat != nil {
			subMap["messageFormat"] = *config.MessageFormat
		}
		if config.ServiceAccountEmail != nil {
			subMap["serviceAccountEmail"] = *config.ServiceAccountEmail
		}
		transformed[*topic] = subMap
	}
	glog.Infof(fmt.Sprintf("Transformed: %#v", transformed))
	glog.Infof(fmt.Sprintf("Configs: %#v", configs))
	return transformed, nil
}

// flattenRepoPubsubConfigs flattens the response into DCL objects
// The pubsub config response is structured as a map[string]map[string]string
// where the first key is the name of the topic
func flattenRepoPubsubConfig(v interface{}) []RepoPubsubConfigs {
	if v == nil {
		return nil
	}
	l, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	var transformed []RepoPubsubConfigs
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, RepoPubsubConfigs{
			Topic:               dcl.FlattenString(k),
			MessageFormat:       dcl.FlattenString(original["messageFormat"]),
			ServiceAccountEmail: dcl.FlattenString(original["serviceAccountEmail"]),
		})
	}
	return transformed
}

// createPubsubConfigs adds a patch to apply PubsubConfigs after create (if applicable).
func createPubsubConfigs(inOps []repoApiOperation) ([]repoApiOperation, error) {
	for _, op := range inOps {
		if _, ok := op.(*createRepoOperation); ok {
			return append(inOps, &updateRepoUpdateRepoOperation{Diffs: []*dcl.FieldDiff{{FieldName: "pubsubConfigs"}}}), nil
		}
	}
	return inOps, nil
}

// EncodeRepoUpdateRequest encapsulates fields into a repo block as expected by the API.
func EncodeRepoUpdateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	dcl.PutMapEntry(req, []string{"repo"}, m)
	return req
}
