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
	"bytes"
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
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
	trimmed := dcl.ValueOrEmptyString(selfLink)
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

func (op *createEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := environmentCreateURL(c.Config.BasePath, project, location)

	if err != nil {
		return err
	}

	// Store the pypi packages field value for a later update operation.
	pypiPackages := r.Config.SoftwareConfig.PypiPackages
	r.Config.SoftwareConfig.PypiPackages = nil

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://composer.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if pypiPackages == nil {
		op.response, _ = o.FirstResponse()
	} else {
		// Restore the pypi packages field and perform an update operation.
		r.Config.SoftwareConfig.PypiPackages = pypiPackages
		u, err := r.updateURL(c.Config.BasePath, "PatchEnvironment")
		if err != nil {
			return err
		}
		u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": "config.softwareConfig.pypiPackages"})
		if err != nil {
			return err
		}
		op.response, err = updateEnvironment(ctx, r, c, u)
		if err != nil {
			return err
		}
	}

	if _, err := c.GetEnvironment(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

// do is a custom method because a separate update request must be made for each field in the update mask
func (op *updateEnvironmentPatchEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	_, err := c.GetEnvironment(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchEnvironment")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	for _, maskPart := range strings.Split(mask, ",") {
		u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": maskPart})
		if err != nil {
			return err
		}
		_, err := updateEnvironment(ctx, r, c, u)
		if err != nil {
			return err
		}
	}

	return nil
}

// Perform one update request on the given environment. Return the response.
func updateEnvironment(ctx context.Context, r *Environment, c *Client, u string) (map[string]interface{}, error) {
	req, err := newUpdateEnvironmentPatchEnvironmentRequest(ctx, r, c)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateEnvironmentPatchEnvironmentRequest(c, req)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return nil, err
	}
	err = o.Wait(ctx, c.Config, "https://composer.googleapis.com/v1/", "GET")

	if err != nil {
		return nil, err
	}

	// Return the response for use in the create operation.
	response, _ := o.FirstResponse()
	return response, nil
}
