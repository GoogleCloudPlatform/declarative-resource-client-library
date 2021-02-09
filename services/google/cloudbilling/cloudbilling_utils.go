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
// Package cloudbilling contains handwritten code for cloudbilling.
package cloudbilling

import (
	"bytes"
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Do creates an Update request and deletes a Project Billing Info.
func (op *deleteProjectBillingInfoOperation) do(ctx context.Context, r *ProjectBillingInfo, c *Client) error {
	// ProjectBillingInfo has no DELETE route, so we will instead delete using a PUT request with billing account set to an empty string.

	projectID := r.createFields()
	u, err := projectBillingInfoCreateURL(c.Config.BasePath, projectID)
	if err != nil {
		return err
	}
	r.BillingAccountName = nil
	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}

	return nil
}
