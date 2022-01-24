// Copyright 2022 Google LLC. All Rights Reserved.
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
// Package dlp contains types and methods for handling dlp GCP resources.
package beta

import (
	"bytes"
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// encodeDLPCreateRequest encodes the create request for an dlp resource.
func encodeDLPCreateRequest(m map[string]interface{}, resourceName string) map[string]interface{} {
	req := make(map[string]interface{})
	// Put base object into object field.
	dcl.PutMapEntry(req, []string{resourceName}, m)
	return req
}

// encodeDeidentifyTemplateCreateRequest properly encodes the create request for an dlp deidentify template.
func encodeDeidentifyTemplateCreateRequest(m map[string]interface{}) map[string]interface{} {
	return encodeDLPCreateRequest(m, "deidentifyTemplate")
}

// encodeInspectTemplateCreateRequest properly encodes the create request for an dlp inspect template.
func encodeInspectTemplateCreateRequest(m map[string]interface{}) map[string]interface{} {
	return encodeDLPCreateRequest(m, "inspectTemplate")
}

// Update has a custom method because the update mask needs to be in the request body.
func (op *updateDeidentifyTemplateUpdateDeidentifyTemplateOperation) do(ctx context.Context, r *DeidentifyTemplate, c *Client) error {
	_, err := c.GetDeidentifyTemplate(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateDeidentifyTemplate")
	if err != nil {
		return err
	}

	req, err := newUpdateDeidentifyTemplateUpdateDeidentifyTemplateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	mask := dcl.TopLevelUpdateMask(op.FieldDiffs)
	req = map[string]interface{}{
		"deidentifyTemplate": req,
		"updateMask":         mask,
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateDeidentifyTemplateUpdateDeidentifyTemplateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Update has a custom method because the update mask needs to be in the request body.
func (op *updateInspectTemplateUpdateInspectTemplateOperation) do(ctx context.Context, r *InspectTemplate, c *Client) error {
	_, err := c.GetInspectTemplate(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInspectTemplate")
	if err != nil {
		return err
	}

	req, err := newUpdateInspectTemplateUpdateInspectTemplateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	mask := dcl.TopLevelUpdateMask(op.FieldDiffs)
	req = map[string]interface{}{
		"inspectTemplate": req,
		"updateMask":      mask,
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateInspectTemplateUpdateInspectTemplateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}
