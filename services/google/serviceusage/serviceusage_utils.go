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
// Package serviceusage contains methods and objects for managing serviceusage GCP resources.
package serviceusage

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// Do creates a create request and creates a service.
func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	u, err := serviceCreateURL(c.Config.BasePath, *r.Project, *r.Name)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer([]byte{}), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	// The initial request may return an operation already in the "done" state.
	if !o.Done {
		if err := o.Wait(ctx, c.Config, "https://serviceusage.googleapis.com/v1", "GET"); err != nil {
			c.Config.Logger.Warningf("Enable %q failed after waiting for operation: %v", dcl.ValueOrEmptyString(r.Name), err)
			return err
		}
	}

	err = r.waitForServiceEnabled(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

// Do creates a delete request and deletes a service.
func (op *deleteServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	u, err := serviceDeleteURL(c.Config.BasePath, r)
	if err != nil {
		return err
	}
	body, err := json.Marshal(map[string]bool{"disableDependentServices": true})
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// contains returns true if a service with the given name exists in the service list.
func (list *ServiceList) hasEnabled(name string, c *Client) bool {
	for _, service := range list.Items {
		selfLinkName := dcl.SelfLinkToName(service.Name)
		if selfLinkName != nil && *selfLinkName == name {
			state := dcl.ValueOrEmptyString(service.State)
			if state == "ENABLED" {
				c.Config.Logger.Infof("Service %q is enabled", name)
				return true
			}
			c.Config.Logger.Infof("Service %q is not yet enabled: state = %q", name, state)
			return false
		}
	}
	return false
}

// waitForServiceEnabled waits for the service to appear in the list of enabled services.
func (r *Service) waitForServiceEnabled(ctx context.Context, c *Client) error {
	return dcl.Do(ctx, func(ctc context.Context) (*dcl.RetryDetails, error) {
		list, err := c.ListService(ctx, *r.Project)
		if err != nil {
			return nil, err
		}
		if list.hasEnabled(*r.Name, c) {
			return nil, nil
		}
		for list.HasNext() {
			err = list.Next(ctx, c)
			if err != nil {
				return nil, err
			}
			if list.hasEnabled(*r.Name, c) {
				return nil, nil
			}
		}
		return &dcl.RetryDetails{}, dcl.OperationNotDone{}
	}, c.Config.RetryProvider)
}
