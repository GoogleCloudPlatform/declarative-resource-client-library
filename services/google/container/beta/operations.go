// Copyright 2020 Google LLC. All Rights Reserved.
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
package beta

import (
	"bytes"
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Operation can be parsed from the returned API operation and waited on.
type Operation struct {
	ID         string `json:"id"`
	SelfLink   string `json:"selfLink"`
	Status     string `json:"status"`
	TargetLink string `json:"targetLink"`
	// other irrelevant fields omitted

	config *dcl.Config
}

// Wait waits for an Operation to complete by fetching the operation until it completes.
func (op *Operation) Wait(ctx context.Context, c *dcl.Config) error {
	c.Logger.Infof("Waiting on: %v", op)
	op.config = c

	return dcl.Do(ctx, op.operate, c.Retry)
}

func (op *Operation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	resp, err := dcl.SendRequest(ctx, op.config, "GET", op.SelfLink, &bytes.Buffer{}, nil)
	if err != nil {
		if dcl.IsRetryableRequestError(op.config, err, false) {
			return nil, dcl.OperationNotDone{}
		}
		return nil, err
	}
	if err := dcl.ParseResponse(resp.Response, op); err != nil {
		return nil, err
	}
	if op.Status != "DONE" {
		return nil, dcl.OperationNotDone{}
	}
	return resp, nil
}
