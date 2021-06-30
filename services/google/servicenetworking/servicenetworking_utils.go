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
// Package servicenetworking contains handwritten code for dealing with service networking configs.
package servicenetworking

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// Do removes the link from the network specified.
func (op *deleteConnectionOperation) do(ctx context.Context, r *Connection, c *Client) error {
	un := r.URLNormalized()
	u := fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/%s/global/networks/%s/removePeering", *un.Project, *un.Network)
	body := fmt.Sprintf(`{"name": %q}`, *un.Name)
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBufferString(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://compute.googleapis.com/v1/", "GET")
	return err
}
