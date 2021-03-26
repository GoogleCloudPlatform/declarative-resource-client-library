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
package logging

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (op *deleteLogBucketOperation) do(ctx context.Context, r *LogBucket, c *Client) error {

	_, err := c.GetLogBucket(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("LogBucket not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetLogBucket checking for existence. error: %v", err)
		return err
	}

	if *r.Name != "_Default" && *r.Name != "_Required" {
		u, err := logBucketDeleteURL(c.Config.BasePath, r.urlNormalized())
		if err != nil {
			return err
		}

		// Delete should never have a body
		body := &bytes.Buffer{}
		_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
		if err != nil {
			return fmt.Errorf("failed to delete LogBucket: %w", err)
		}
	}
	return nil
}
