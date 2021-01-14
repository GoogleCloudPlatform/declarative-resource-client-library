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
package operations

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// StandardGCPOperation can be parsed from the returned API operation and waited on.
// This is the typical GCP operation.
type StandardGCPOperation struct {
	Name     string                     `json:"name"`
	Error    *StandardGCPOperationError `json:"error"`
	Done     bool                       `json:"done"`
	Response map[string]interface{}     `json:"response"`
	// other irrelevant fields omitted

	config   *dcl.Config
	basePath string
	verb     string
}

// StandardGCPOperationError is the GCP operation's Error body.
type StandardGCPOperationError struct {
	Errors []*StandardGCPOperationErrorError `json:"errors"`
}

// String formats the StandardGCPOperationError as an error string.
func (e *StandardGCPOperationError) String() string {
	var b strings.Builder
	for _, err := range e.Errors {
		fmt.Fprintf(&b, "error code %q, message: %s\n", err.Code, err.Message)
	}

	return b.String()
}

// StandardGCPOperationErrorError is a singular error in a GCP operation.
type StandardGCPOperationErrorError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// FetchResponseValue fetches a top-level field from the Response object of a
// completed operation.  The response object is usually only present in a
// completed operation, so it is unlikely that this will return anything useful
// if called on an operation in progress.
func (op *StandardGCPOperation) FetchResponseValue(val string) (*string, error) {
	if v, ok := op.Response[val]; ok {
		if vs, ok := v.(string); ok {
			return dcl.String(vs), nil
		}
		return nil, fmt.Errorf("could not cast %v - value at %q, to string", v, val)
	}
	return nil, fmt.Errorf("could not find value at %q", val)
}

// Wait waits for an StandardGCPOperation to complete by fetching the operation until it completes.
func (op *StandardGCPOperation) Wait(ctx context.Context, c *dcl.Config, basePath, verb string) error {
	c.Logger.Infof("Waiting on: %v", op)
	op.config = c
	op.basePath = basePath
	op.verb = verb

	return dcl.Do(ctx, op.operate, c.Retry)
}

func (op *StandardGCPOperation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	u := dcl.URL(op.Name, op.basePath, op.config.BasePath, nil)
	resp, err := dcl.SendRequest(ctx, op.config, op.verb, u, &bytes.Buffer{}, nil)
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
		return nil, fmt.Errorf("operation received error: %s", op.Error)
	}

	return resp, nil
}
