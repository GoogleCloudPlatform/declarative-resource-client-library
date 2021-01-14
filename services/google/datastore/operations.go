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
// Package datastore contains handwritten support code for the datastore service.
package datastore

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Operation can be parsed from the returned API operation and waited on.
type Operation struct {
	Name     string             `json:"name"`
	Done     bool               `json:"done"`
	Metadata *OperationMetadata `json:"metadata"`
	Error    *OperationError    `json:"error"`
	config   *dcl.Config
}

// OperationMetadata is an error in a datastore operation.
type OperationMetadata struct {
	IndexID string `json:"indexId"`
}

// OperationError is an error in a datastore operation.
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
	u := dcl.URL(op.Name, "https://datastore.googleapis.com/v1/", op.config.BasePath, nil)
	resp, err := dcl.SendRequest(ctx, op.config, "GET", u, &bytes.Buffer{}, nil)
	if err != nil {
		if IsDatastoreRetryableError(op.config, err) || dcl.IsRetryableRequestError(op.config, err, true) {
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

// FetchIndexID will fetch the operation and return the indexId of the resource created.
// Datastore index uses a machine generated id as a param.
// It must be called after the resource has been created.
func (op *Operation) FetchIndexID() (*string, error) {
	if op.Metadata.IndexID == "" {
		return nil, fmt.Errorf("this operation (%s) has no IndexID and probably hasn't been run before", op.Name)
	}
	return &op.Metadata.IndexID, nil
}

// IsDatastoreRetryableError checks for additional retryable errors that are
// specific to Datastore.
func IsDatastoreRetryableError(c *dcl.Config, err error) bool {
	if gerr, ok := err.(*googleapi.Error); ok {
		if gerr.Code == 409 && strings.Contains(gerr.Body, "too much contention") {
			c.Logger.Infof("Error is retryable: too much contention - waiting for less activity")
			return true
		}
	}
	return false
}
