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
// Package compute contains handwritten support code for the compute service.
package compute

import (
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// EncodeImageDeprecateRequest properly encodes the image deprecation request for an image.
func EncodeImageDeprecateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	// Deprecate requests involve removing the "deprecated" key.
	if deprecatedVal, ok := m["deprecated"]; ok {
		deprecatedMap := deprecatedVal.(map[string]interface{})
		for key, value := range deprecatedMap {
			req[key] = value
		}
	}

	return req
}

// WrapTargetPoolInstance wraps the instances provided by AddInstance and RemoveInstance
// in their required format.
func WrapTargetPoolInstance(m map[string]interface{}) map[string]interface{} {
	is := m["instances"].([]string)
	wrapped := make([]interface{}, len(is))
	for idx, i := range is {
		wrapped[idx] = map[string]interface{}{
			"instance": i,
		}
	}
	return map[string]interface{}{
		"instances": wrapped,
	}
}

// WrapTargetPoolHealthCheck wraps the instances provided by AddHC and RemoveHC
// in their required format.
func WrapTargetPoolHealthCheck(m map[string]interface{}) map[string]interface{} {
	hcs := m["healthChecks"].([]string)
	wrapped := make([]interface{}, len(hcs))
	for idx, hc := range hcs {
		wrapped[idx] = map[string]interface{}{
			"healthCheck": hc,
		}
	}
	return map[string]interface{}{
		"healthChecks": wrapped,
	}
}

func canonicalizeReservationCPUPlatform(o, n *string) bool {
	if o == nil && n == nil {
		return true
	}
	if o == nil || n == nil {
		return false
	}
	if *o == "automatic" {
		return true
	}
	if *n == "automatic" {
		return true
	}

	return *o == *n
}

func deriveAutoscalerTarget(as *Autoscaler) (*string, error) {
	if !dcl.IsEmptyValueIndirect(as.Location) {
		if dcl.IsZone(as.Location) {
			return dcl.DeriveField("projects/%s/zones/%s/instanceGroupManagers/%s", as.Target, as.Project, dcl.SelfLinkToName(as.Location), dcl.SelfLinkToName(as.Target))
		} else if dcl.IsRegion(as.Location) {
			return dcl.DeriveField("projects/%s/regions/%s/instanceGroupManagers/%s", as.Target, as.Project, dcl.SelfLinkToName(as.Location), dcl.SelfLinkToName(as.Target))
		}
	}
	return nil, errors.New("could not derive autoscaler target - location was neither zone or region")
}
