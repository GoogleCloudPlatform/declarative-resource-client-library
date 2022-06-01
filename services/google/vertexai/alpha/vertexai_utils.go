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
// Package vertexai defines types and methods for handling vertex GCP resources.
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func encodeVertexRequest(m map[string]interface{}, fieldName string) map[string]interface{} {
	req := make(map[string]interface{})
	// Put base object into object field.
	dcl.PutMapEntry(req, []string{fieldName}, m)
	return req
}

// encodeUploadModelRequest encodes the create request for a vertex model resource.
func encodeUploadModelRequest(m map[string]interface{}) map[string]interface{} {
	return encodeVertexRequest(m, "model")
}

// encodeDeployModelRequest encodes the deploy request for a vertex model resource.
func encodeDeployModelRequest(m map[string]interface{}) map[string]interface{} {
	return encodeVertexRequest(m, "deployedModel")
}

// Model has a custom create method because the resource name has a different field name in the operation.
func (op *createModelOperation) do(ctx context.Context, r *Model, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
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
	// Wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["model"].(string)
	if !ok {
		return fmt.Errorf("expected model to be a string in %v, was %T", op.response, op.response["model"])
	}
	r.Name = &name

	if _, err := c.GetModel(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

// Model deployment has a custom delete method because the request should have a body.
func (op *deleteModelDeploymentOperation) do(ctx context.Context, r *ModelDeployment, c *Client) error {
	r, err := c.GetModelDeployment(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "ModelDeployment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetModelDeployment checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := json.Marshal(map[string]interface{}{
		"deployedModelId": dcl.ValueOrEmptyString(r.Id),
	})
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(req)
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// Wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetModelDeployment(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// expands the list of key-value pairs into a map
func endpointTrafficSplitSliceToMap(_ *Client, objectSlice []EndpointTrafficSplitTrafficSplit, resource *EndpointTrafficSplit) (map[string]int64, error) {
	m := make(map[string]int64, len(objectSlice))
	for _, object := range objectSlice {
		if deployedModelID := dcl.ValueOrEmptyString(object.DeployedModelId); deployedModelID != "" {
			m[deployedModelID] = dcl.ValueOrEmptyInt64(object.TrafficPercentage)
		} else {
			return m, fmt.Errorf("deployed model id not found on endpoint %s with percentage %d", dcl.ValueOrEmptyString(resource.Endpoint), dcl.ValueOrEmptyInt64(object.TrafficPercentage))
		}
	}
	return m, nil
}

// flattens a map into a list of key-value pairs
func endpointTrafficSplitMapToSlice(c *Client, i interface{}, resource *EndpointTrafficSplit) []EndpointTrafficSplitTrafficSplit {
	if i == nil {
		return nil
	}
	m, ok := i.(map[string]interface{})
	if !ok {
		c.Config.Logger.Warningf("could not cast %T into a map for endpoint %s", i, dcl.ValueOrEmptyString(resource.Endpoint))
		return nil
	}
	items := make([]EndpointTrafficSplitTrafficSplit, len(m))
	j := 0
	for key, value := range m {
		floatValue, ok := value.(float64)
		if !ok {
			c.Config.Logger.Warningf("could not cast %T into a float for endpoint %s", value, dcl.ValueOrEmptyString(resource.Endpoint))
			return nil
		}
		intValue := int64(floatValue)
		items[j] = EndpointTrafficSplitTrafficSplit{
			DeployedModelId:   dcl.String(key),
			TrafficPercentage: dcl.Int64(intValue),
		}
		j++
	}
	return items
}
