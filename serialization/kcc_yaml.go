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
// Package serialization contains functions for serializing and deserializing DCL resources to and from KCC and Terraform formats.
package serialization

import (
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

type kccComputeSSLPolicy struct {
	APIVersion string                      `yaml:"apiVersion"`
	Kind       string                      `yaml:"kind"`
	Metadata   kccComputeSSLPolicyMetadata `yaml:"metadata"`
	Spec       kccComputeSSLPolicySpec     `yaml:"spec"`
}

type kccComputeSSLPolicyMetadata struct {
	Name string `yaml:"name"`
}

type kccComputeSSLPolicySpec struct {
	CustomFeatures []string `yaml:"customFeatures,omitempty"`
	Description    string   `yaml:"description,omitempty"`
	MinTLSVersion  string   `yaml:"minTlsVersion,omitempty"`
	Profile        string   `yaml:"profile,omitempty"`
}

// YAMLToDCL takes in a full kcc-formatted YAML as a slice of bytes
// and returns a DCL resource or an error.
func YAMLToDCL(b []byte) (dcl.Resource, error) {
	var m map[string]interface{}
	if err := yaml.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	if k, ok := m["kind"]; ok && k == "ComputeSSLPolicy" {
		s := kccComputeSSLPolicy{}
		if err := yaml.UnmarshalStrict(b, &s); err != nil {
			return nil, err
		}
		return &compute.SslPolicy{
			CustomFeature: s.Spec.CustomFeatures,
			Description:   dcl.StringOrNil(s.Spec.Description),
			MinTlsVersion: compute.SslPolicyMinTlsVersionEnumRef(s.Spec.MinTLSVersion),
			Profile:       compute.SslPolicyProfileEnumRef(s.Spec.Profile),
			Name:          dcl.String(s.Metadata.Name),
		}, nil
	}
	return nil, fmt.Errorf("could not parse kind %q", m["kind"])
}

// ComputeSslPolicyToYAML converts a compute.SslPolicy to KCC YAML format, serialized as a slice of bytes.
func ComputeSslPolicyToYAML(p compute.SslPolicy) ([]byte, error) {
	sp := kccComputeSSLPolicy{
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
		Kind:       "ComputeSSLPolicy",
		Metadata:   kccComputeSSLPolicyMetadata{Name: *p.Name},
		Spec:       kccComputeSSLPolicySpec{},
	}
	if len(p.CustomFeature) > 0 {
		sp.Spec.CustomFeatures = p.CustomFeature
	}
	if p.Description != nil {
		sp.Spec.Description = *p.Description
	}
	if p.MinTlsVersion != nil {
		sp.Spec.MinTLSVersion = string(*p.MinTlsVersion)
	}
	if p.Profile != nil {
		sp.Spec.Profile = string(*p.Profile)
	}
	return yaml.Marshal(sp)
}
