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
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Do creates a create request and creates a service.
func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	u, err := serviceCreateURL(c.Config.BasePath, *r.Project, *r.Name)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer([]byte{}), c.Config.Retry)
	if err != nil {
		return err
	}
	err = r.waitForServiceEnabled(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

// Do deletes a delete request and deletes a service.
func (op *deleteServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	u, err := serviceDeleteURL(c.Config.BasePath, r)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer([]byte{}), c.Config.Retry)
	if err != nil {
		return err
	}

	return nil
}

// contains returns true if a service with the given name exists in the service list.
func (list *ServiceList) contains(name string) bool {
	for _, service := range list.Items {
		if *dcl.SelfLinkToName(service.Name) == name {
			return true
		}
	}
	return false
}

// waitForServiceEnabled waits for the service to appear in the list of enabled services.
func (r *Service) waitForServiceEnabled(ctx context.Context, c *Client) error {
	for {
		list, err := c.ListService(ctx, *r.Project)
		if err != nil {
			return err
		}
		if list.contains(*r.Name) {
			return nil
		}
		for list.HasNext() {
			err = list.Next(ctx, c)
			if err != nil {
				return err
			}
			if list.contains(*r.Name) {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
}
