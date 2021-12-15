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
// Package logging provides types and functiosn for handling logging GCP resources.
package beta

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// do makes a request to create a new log bucket if the name of the desired bucket is not
// "_Default" or "_Required"
func (op *createLogBucketOperation) do(ctx context.Context, r *LogBucket, c *Client) error {
	if r.Name != nil && (*r.Name == "_Default" || *r.Name == "_Required") {
		return nil
	}
	c.Config.Logger.Infof("Attempting to create %v", r)

	u, err := r.createURL(c.Config.BasePath)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetLogBucket(ctx, r); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

// do makes a request to delete a log bucket if the name of the bucket is not
// "_Default" or "_Required"
func (op *deleteLogBucketOperation) do(ctx context.Context, r *LogBucket, c *Client) error {

	_, err := c.GetLogBucket(ctx, r)

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("LogBucket not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetLogBucket checking for existence. error: %v", err)
		return err
	}

	if r.Name != nil && (*r.Name == "_Default" || *r.Name == "_Required") {
		return nil
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete LogBucket: %w", err)
	}
	return nil
}
