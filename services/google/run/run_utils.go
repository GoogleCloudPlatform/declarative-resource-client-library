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
// Package run contains DCL for Google Cloud Run.
package run

import (
	"bytes"
	"context"
	"io/ioutil"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {
	if dcl.IsZeroValue(r.ApiVersion) {
		r.ApiVersion = dcl.String("serving.knative.dev/v1")
	}
	if dcl.IsZeroValue(r.Kind) {
		r.Kind = dcl.String("Service")
	}

	u, err := serviceGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) GetService(ctx context.Context, r *Service) (*Service, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(60*time.Second))
	defer cancel()

	u, err := serviceGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}

	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}

	var o operations.KNativeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return nil, err
	}

	if err := o.Wait(ctx, c.Config, "https://{{location}}-run.googleapis.com/", "GET"); err != nil {
		c.Config.Logger.Warningf("Get failed after waiting for operation: %v", err)
		return nil, err
	}

	b, err := c.getServiceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}

	result, err := unmarshalService(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name
	if dcl.IsZeroValue(result.ApiVersion) {
		result.ApiVersion = dcl.String("serving.knative.dev/v1")
	}
	if dcl.IsZeroValue(result.Kind) {
		result.Kind = dcl.String("Service")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServiceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}
