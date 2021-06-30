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
// Package privateca contains methods and objects for handling privateca GCP resources.
package beta

import (
	"bytes"
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func waitForCertificateAuthorityEnabled(ctx context.Context, r *CertificateAuthority, c *Client) error {
	return dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		r, err := c.GetCertificateAuthority(ctx, r)
		if err != nil {
			return nil, err
		}
		if *r.State == *CertificateAuthorityStateEnumRef("ENABLED") {
			return nil, nil
		}
		return &dcl.RetryDetails{}, dcl.OperationNotDone{}
	}, c.Config.RetryProvider)
}

func (op *createCertificateAuthorityOperation) do(ctx context.Context, r *CertificateAuthority, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := certificateAuthorityCreateURL(c.Config.BasePath, project, location, name)

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
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://privateca.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()
	if err := waitForCertificateAuthorityEnabled(ctx, r, c); err != nil {
		return err
	}

	return nil
}

func (op *deleteCertificateAuthorityOperation) do(ctx context.Context, r *CertificateAuthority, c *Client) error {

	_, err := c.GetCertificateAuthority(ctx, r.URLNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("CertificateAuthority not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetCertificateAuthority checking for existence. error: %v", err)
		return err
	}

	n := r.URLNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(n.Project),
		"location": dcl.ValueOrEmptyString(n.Location),
		"name":     dcl.ValueOrEmptyString(n.Name),
	}
	u := dcl.URL("projects/{{project}}/locations/{{location}}/certificateAuthorities/{{name}}:disable", "https://privateca.googleapis.com/v1beta1/", c.Config.BasePath, params)
	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be disabled.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://privateca.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	u, err = certificateAuthorityDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	resp, err = dcl.SendRequest(ctx, c.Config, "POST", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://privateca.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}
	return nil
}
