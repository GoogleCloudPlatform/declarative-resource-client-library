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
// Package apigee contains methods and types for handling apigee GCP resources.
package apigee

import (
	"bytes"
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.
func (op *updateOrganizationUpdateOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {
	_, err := c.GetOrganization(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateOrganization")
	if err != nil {
		return err
	}

	req, err := newUpdateOrganizationUpdateOrganizationRequest(ctx, r, c)
	if err != nil {
		return err
	}
	dcl.PutMapEntry(req, []string{"name"}, r.Name)

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateOrganizationUpdateOrganizationRequest(c, req)
	if err != nil {
		return err
	}
	if _, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider); err != nil {
		return err
	}

	return nil
}
