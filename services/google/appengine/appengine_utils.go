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
// Package appengine contains handwritten utilities for the appengine library.
package appengine

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	rm "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
)

// ApplicationDeleteURL returns the URL to send a DELETE request for Application.  It is notably to the
// resourcemanager API - Applications cannot be deleted, only their underlying projects can.
func applicationDeleteURL(userBasePath string, a *Application) (string, error) {
	return dcl.URL(dcl.ValueOrEmptyString(a.Name), "https://cloudresourcemanager.googleapis.com/v1/projects/", userBasePath, nil), nil
}

// PlanProjectStatus adds a create for the underlying project (and/or a delete for the underlying project)
// based on the existing plan.
func PlanProjectStatus(inOps []applicationApiOperation) (outOps []applicationApiOperation, err error) {
	for _, op := range inOps {
		if _, ok := op.(*createApplicationOperation); ok {
			outOps = append(outOps, &createUnderlyingProjectOperation{})
		}
		if _, ok := op.(*deleteApplicationOperation); ok {
			outOps = append(outOps, &deleteUnderlyingProjectOperation{})
		} else {
			// Do not include a DeleteApplicationOperation, ever - you can't
			// delete an App.
			outOps = append(outOps, op)
		}
	}
	return
}

type createUnderlyingProjectOperation struct{}

func (*createUnderlyingProjectOperation) do(ctx context.Context, a *Application, c *Client) error {
	rmCl := rm.NewClient(c.Config)
	// We can turn the appengine parent into a resourcemanager parent via the json
	// encoding.
	parent := &rm.ProjectParent{}
	if a.Parent != nil {
		b, err := json.Marshal(a.Parent)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, parent); err != nil {
			return err
		}
	}
	_, err := rmCl.ApplyProject(ctx, &rm.Project{
		Name:   a.Name,
		Parent: parent,
	})
	return err
}

type deleteUnderlyingProjectOperation struct{}

func (*deleteUnderlyingProjectOperation) do(ctx context.Context, a *Application, c *Client) error {
	rmCl := rm.NewClient(c.Config)
	if a.Name == nil {
		return errors.New("cannot delete underlying project of application with nil name")
	}

	return rmCl.DeleteProject(ctx, &rm.Project{Name: a.Name})
}

func (c *Client) listApplicationRaw(ctx context.Context, key string) ([]byte, error) {
	return nil, errors.New("appengine does not have a list endpoint")
}
