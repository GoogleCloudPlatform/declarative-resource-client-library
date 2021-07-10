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
// Package assuredworkloads contains support code for the assuredworkload service.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Returns the URL of the reource (always a project) with the given index in the workload.
func (r *Workload) projectURL(userBasePath string, index int) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Resources[index].ResourceId),
	}
	return dcl.URL("projects/{{project}}", "https://cloudresourcemanager.googleapis.com/v1/", userBasePath, params), nil
}

// Returns the lifecycle state of the given project resource with the given url in the workload.
func projectLifecycleState(ctx context.Context, client *Client, url string) (string, error) {
	resp, err := dcl.SendRequest(ctx, client.Config, "GET", url, &bytes.Buffer{}, client.Config.RetryProvider)
	if err != nil {
		return "", err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return "", err
	}
	state, ok := m["lifecycleState"].(string)
	if !ok {
		return "", fmt.Errorf("no lifecycle state for project at %q", url)
	}
	return state, nil
}

// The resources to be deleted first befoe deleting Workload
func (r *Workload) workloadDeletePreAction(ctx context.Context, client *Client) error {
	nr, err := client.GetWorkload(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			client.Config.Logger.Infof("Workload not found, returning. Original error: %v", err)
			return nil
		}
		client.Config.Logger.Warningf("GetWorkload checking for existence. error: %v", err)
		return err
	}
	for i, resource := range nr.Resources {
		if resource.ResourceType != nil && *resource.ResourceType != WorkloadResourcesResourceTypeEnum("CONSUMER_PROJECT") {
			// Only support project delete method.
			continue
		}
		u, err := nr.URLNormalized().projectURL(client.Config.BasePath, i)
		if err != nil {
			return err
		}
		state, err := projectLifecycleState(ctx, client, u)
		if err != nil {
			return err
		}
		if state == "DELETE_REQUESTED" {
			// Do not delete already deleted project.
			continue
		}
		_, err = dcl.SendRequest(ctx, client.Config, "DELETE", u, &bytes.Buffer{}, client.Config.RetryProvider)
		if err != nil {
			return fmt.Errorf("failed to delete Workload Resource: %w", err)
		}
	}
	return dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		for i, resource := range nr.Resources {
			if resource.ResourceType != nil && *resource.ResourceType != WorkloadResourcesResourceTypeEnum("CONSUMER_PROJECT") {
				// Only support project delete method.
				continue
			}
			u, err := nr.URLNormalized().projectURL(client.Config.BasePath, i)
			if err != nil {
				return nil, err
			}
			state, err := projectLifecycleState(ctx, client, u)
			if err != nil {
				return nil, err
			}
			if state != "DELETE_REQUESTED" {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{}
			}

		}
		return nil, nil
	}, client.Config.RetryProvider)
}
