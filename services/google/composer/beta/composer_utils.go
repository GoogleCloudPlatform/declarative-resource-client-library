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
// Package composer defines types and methods for managing cloud composer GCP resources.
package beta

import (
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Return either zones or regions depending on the type of the given location.
func locationType(location *string) *string {
	var locationType string
	if dcl.IsZone(location) {
		locationType = "zones"
	} else {
		locationType = "regions"
	}
	return &locationType
}

// Remove the prefix of the given compute self link to generate a relative resource name.
func trimSelfLink(selfLink *string) *string {
	trimmed := *selfLink
	trimmed = strings.TrimPrefix(trimmed, "https://www.googleapis.com/compute/v1/")
	return &trimmed
}

// Expand the given location into a relative resource name.
func expandComposerEnvironmentLocation(r *Environment) (*string, error) {
	return dcl.DeriveField("projects/%s/%s/%s", r.Config.NodeConfig.Location, r.Project, locationType(r.Config.NodeConfig.Location), r.Config.NodeConfig.Location)
}

// Flatten the location relative resource name down to its name.
func flattenComposerEnvironmentLocation(r *Environment) *string {
	return dcl.SelfLinkToName(r.Config.NodeConfig.Location)
}

// Expand the given machine type into a relative resource name.
func expandComposerEnvironmentMachineType(r *Environment) (*string, error) {
	return dcl.DeriveField("projects/%s/%s/%s/machineTypes/%s", r.Config.NodeConfig.MachineType, r.Project, locationType(r.Config.NodeConfig.Location), r.Config.NodeConfig.Location, r.Config.NodeConfig.MachineType)
}

// Flatten the machine type relative resource name down to its name.
func flattenComposerEnvironmentMachineType(r *Environment) *string {
	return dcl.SelfLinkToName(r.Config.NodeConfig.MachineType)
}

// Expand the network self link to a relative resource name.
func expandComposerEnvironmentNetwork(r *Environment) (*string, error) {
	return trimSelfLink(r.Config.NodeConfig.Network), nil
}

// Flatten the network relative resource name down to its name.
func flattenComposerEnvironmentNetwork(r *Environment) *string {
	return dcl.SelfLinkToName(r.Config.NodeConfig.Network)
}

// Expand the subnetwork self link to a relative resource name.
func expandComposerEnvironmentSubnetwork(r *Environment) (*string, error) {
	return trimSelfLink(r.Config.NodeConfig.Subnetwork), nil
}

// Flatten the subnetwork relative resource name down to its name.
func flattenComposerEnvironmentSubnetwork(r *Environment) *string {
	return dcl.SelfLinkToName(r.Config.NodeConfig.Subnetwork)
}
