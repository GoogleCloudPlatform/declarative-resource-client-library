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

// AcmOperation can be parsed from the returned API operation and waited on.
type AcmOperation struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Done     bool     `json:"done"`
	Response response `json:"response"`

	ResourceName string

	config *dcl.Config
}

type response struct {
	Name string `json:"name"`
}

// Wait waits for an AcmOperation to complete by fetching the operation until it completes.
func (op *AcmOperation) Wait(ctx context.Context, c *dcl.Config, _, _ string) error {
	c.Logger.Infof("Waiting on: %v", op)
	op.config = c

	// The first operation query happens outside of this Wait.
	// We need to read the response immediately in case the first operation call has the response in it.
	if op.Response.Name != "" {
		// Names come in the form "accessPolicies/{{name}}"
		op.ResourceName = strings.Split(op.Response.Name, "/")[1]
	}

	return dcl.Do(ctx, op.operate, c.Retry)
}

func (op *AcmOperation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	u := dcl.URL(op.Name, "https://accesscontextmanager.googleapis.com/v1/", op.config.BasePath, nil)
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

	if op.Response.Name != "" {
		op.ResourceName = strings.Split(op.Response.Name, "/")[1]
	}
	return resp, nil
}

// FetchName will fetch the operation and return the name of the resource created.
// AccessContextManager creates resources with machine generated names.
// It must be called after the resource has been created.
func (op *AcmOperation) FetchName() (*string, error) {
	if op.ResourceName == "" {
		return nil, fmt.Errorf("this operation (%s) has no name and probably hasn't been run before", op.Name)
	}
	return &op.ResourceName, nil
}
