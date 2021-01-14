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
// Package tpu contains handwritten support code for the tpu service.
package tpu

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Operation can be parsed from the returned API operation and waited on.
type Operation struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Done  bool            `json:"done"`
	Error *OperationError `json:"error"`

	config *dcl.Config
}

// OperationError is an error in a TPU operation.
type OperationError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Wait waits for an Operation to complete by fetching the operation until it completes.
func (op *Operation) Wait(ctx context.Context, c *dcl.Config) error {
	c.Logger.Infof("Waiting on: %v", op)
	op.config = c

	return dcl.Do(ctx, op.operate, c.Retry)
}

func (op *Operation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	u := dcl.URL(op.Name, "https://tpu.googleapis.com/v1/", op.config.BasePath, nil)
	resp, err := dcl.SendRequest(ctx, op.config, "GET", u, &bytes.Buffer{}, nil)
	if err != nil {
		if dcl.IsRetryableRequestError(op.config, err, false) {
			return nil, dcl.OperationNotDone{}
		}
		return nil, err
	}
	if err := dcl.ParseResponse(resp.Response, op); err != nil {
		return nil, err
	}
	if !op.Done {
		return nil, dcl.OperationNotDone{}
	}
	if op.Error != nil {
		return nil, fmt.Errorf("operation received error: %+v", op.Error)
	}
	return resp, nil
}
