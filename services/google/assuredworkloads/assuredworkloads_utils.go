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
package assuredworkloads

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// The resources to be deleted first befoe deleting Workload
func (r *Workload) workloadDeletePreAction(ctx context.Context, client *Client) error {
	_, err := client.GetWorkload(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			client.Config.Logger.Infof("Workload not found, returning. Original error: %v", err)
			return nil
		}
		client.Config.Logger.Warningf("GetWorkload checking for existence. error: %v", err)
		return err
	}
	for i := 0; i < len(r.Resources); i++ {
		u, err := workloadResourceDeleteURL(client.Config.BasePath, r.urlNormalized(), i)
		if err != nil {
			return err
		}
		// Delete should never have a body
		body := &bytes.Buffer{}
		_, err = dcl.SendRequest(ctx, client.Config, "DELETE", u, body, client.Config.RetryProvider)
		if err != nil {
			return fmt.Errorf("failed to delete Workload Resource: %w", err)
		}
	}
	return nil
}

// The URL of the reource to be deleted
func workloadResourceDeleteURL(userBasePath string, r *Workload, index int) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Resources[index].ResourceId),
	}
	return dcl.URL("projects/{{project}}", "https://cloudresourcemanager.googleapis.com/v1/", userBasePath, params), nil
}
