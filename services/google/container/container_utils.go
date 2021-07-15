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
// Package container contains handwritten support code for the container service.
package container

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// DeleteDefaultNodePool deletes the default node pool on a cluster,
// since clusters should have zero node pools to be more declarative.
func (r *Cluster) DeleteDefaultNodePool(ctx context.Context, c *Client) error {
	// Don't delete in case of Autopilot clusters as Autogke clusters do not
	// support mutating node pools.
	if r.Autopilot != nil && dcl.ValueOrEmptyBool(r.Autopilot.Enabled) {
		return nil
	}
	o := deleteDefaultNodePoolOperation{Cluster: r, Client: c}
	return dcl.Do(ctx, o.operate, c.Config.RetryProvider)
}

type deleteDefaultNodePoolOperation struct {
	Cluster *Cluster
	Client  *Client
}

func (d *deleteDefaultNodePoolOperation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	if err := deleteDefaultNodePool(ctx, d.Cluster, d.Client); err != nil {
		// Mutates to a GKE cluster fail with 400 if an operation is deemed in
		// progress. They're also sometimes based on outdated statuses, so we need
		// to just retry a few times.
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 400 {
			return nil, dcl.OperationNotDone{}
		}

		if dcl.IsRetryableRequestError(d.Client.Config, err, false) {
			return nil, dcl.OperationNotDone{}
		}

		return nil, err
	}

	return nil, nil
}

func deleteDefaultNodePool(ctx context.Context, cluster *Cluster, client *Client) error {
	u := cluster.ClusterDeleteDefaultNodePoolURL()

	// Delete should never have a body
	body := &bytes.Buffer{}

	resp, err := dcl.SendRequest(ctx, client.Config, "DELETE", u, body, client.Config.RetryProvider)
	if err != nil {
		return err
	}

	// Wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}

	if err := o.Wait(ctx, client.Config, "", ""); err != nil {
		return err
	}

	return nil
}

// ClusterDeleteDefaultNodePoolURL is the URL for deleting the default node pool.
func (r *Cluster) ClusterDeleteDefaultNodePoolURL() string {
	return fmt.Sprintf("https://container.googleapis.com/v1/projects/%s/locations/%s/clusters/%s/nodePools/default-pool", *r.Project, *r.Location, *r.Name)
}

// EncodeClusterCreateRequest properly encodes the create request for a k8s cluster.
func EncodeClusterCreateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	// Create requests involving a master version have to be sent under the "initialMasterVersion" key.
	if val, ok := m["currentMasterVersion"]; ok {
		m["initialClusterVersion"] = val
		delete(m, "currentMasterVersion")
	}

	dcl.PutMapEntry(req, []string{"cluster"}, m)
	return req
}

// EncodeClusterUpdateRequest places all fields with "desired" prefix and places them in a request under the "update" block.
func EncodeClusterUpdateRequest(m map[string]interface{}) map[string]interface{} {
	update := make(map[string]interface{})
	for k, v := range m {
		if strings.HasPrefix(k, "desired") {
			update[k] = v
		}
	}
	return map[string]interface{}{"update": update}
}

// SetAddonsConfig takes any empty values in AddonsConfig and sets them to {'disabled': false} to match DCL's expectations
func SetAddonsConfig(m map[string]interface{}) (map[string]interface{}, error) {
	if val, ok := m["addonsConfig"]; ok {
		i := val.(map[string]interface{})
		for k, v := range i {
			if _, ok := v.(map[string]interface{})["disabled"]; !ok {
				val.(map[string]interface{})[k] = map[string]bool{"disabled": false}
			}
		}
		return val.(map[string]interface{}), nil
	}
	return make(map[string]interface{}), nil
}

// EncodeClusterLegacyAbacUpdateRequest sends only the "enabled" field.
func EncodeClusterLegacyAbacUpdateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	// TODO(b/150883761): Check and return the error on the GetMapEntry() call.
	i, _ := dcl.GetMapEntry(m, []string{"legacyAbac", "enabled"})
	dcl.PutMapEntry(req, []string{"enabled"}, i)
	return req
}

// EncodeNodePoolCreateRequest properly encodes the create request for a k8s node pool.
func EncodeNodePoolCreateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	v, err := dcl.GetMapEntry(m, []string{"autoscaling", "min_node_count"})
	if err == nil {
		dcl.PutMapEntry(m, []string{"initial_node_count"}, v)
	}

	dcl.PutMapEntry(req, []string{"nodePool"}, m)
	return req
}
