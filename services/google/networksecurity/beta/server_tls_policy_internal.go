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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *ServerTlsPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ServerCertificate) {
		if err := r.ServerCertificate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MtlsPolicy) {
		if err := r.MtlsPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServerTlsPolicyServerCertificate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.GrpcEndpoint) {
		if err := r.GrpcEndpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertificateProviderInstance) {
		if err := r.CertificateProviderInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}
func (r *ServerTlsPolicyMtlsPolicy) validate() error {
	if err := dcl.Required(r, "clientValidationCa"); err != nil {
		return err
	}
	return nil
}
func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) validate() error {
	if !dcl.IsEmptyValueIndirect(r.GrpcEndpoint) {
		if err := r.GrpcEndpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertificateProviderInstance) {
		if err := r.CertificateProviderInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}

func serverTlsPolicyGetURL(userBasePath string, r *ServerTlsPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, params), nil
}

func serverTlsPolicyListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, params), nil

}

func serverTlsPolicyCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies?serverTlsPolicyId={{name}}", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, params), nil

}

func serverTlsPolicyDeleteURL(userBasePath string, r *ServerTlsPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, params), nil
}

func (r *ServerTlsPolicy) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}:setIamPolicy", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, fields)
}

func (r *ServerTlsPolicy) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}:getIamPolicy", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, fields)
}

func (r *ServerTlsPolicy) IAMPolicyVersion() int {
	return 3
}

// serverTlsPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serverTlsPolicyApiOperation interface {
	do(context.Context, *ServerTlsPolicy, *Client) error
}

// newUpdateServerTlsPolicyUpdateServerTlsPolicyRequest creates a request for an
// ServerTlsPolicy resource's UpdateServerTlsPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServerTlsPolicyUpdateServerTlsPolicyRequest(ctx context.Context, f *ServerTlsPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.AllowOpen; !dcl.IsEmptyValueIndirect(v) {
		req["allowOpen"] = v
	}
	if v, err := expandServerTlsPolicyServerCertificate(c, f.ServerCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ServerCertificate into serverCertificate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serverCertificate"] = v
	}
	if v, err := expandServerTlsPolicyMtlsPolicy(c, f.MtlsPolicy); err != nil {
		return nil, fmt.Errorf("error expanding MtlsPolicy into mtlsPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mtlsPolicy"] = v
	}
	return req, nil
}

// marshalUpdateServerTlsPolicyUpdateServerTlsPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateServerTlsPolicyUpdateServerTlsPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateServerTlsPolicyUpdateServerTlsPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServerTlsPolicyUpdateServerTlsPolicyOperation) do(ctx context.Context, r *ServerTlsPolicy, c *Client) error {
	_, err := c.GetServerTlsPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateServerTlsPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateServerTlsPolicyUpdateServerTlsPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateServerTlsPolicyUpdateServerTlsPolicyRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1beta1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listServerTlsPolicyRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := serverTlsPolicyListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServerTlsPolicyMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listServerTlsPolicyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listServerTlsPolicy(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*ServerTlsPolicy, string, error) {
	b, err := c.listServerTlsPolicyRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServerTlsPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ServerTlsPolicy
	for _, v := range m.Items {
		res := flattenServerTlsPolicy(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllServerTlsPolicy(ctx context.Context, f func(*ServerTlsPolicy) bool, resources []*ServerTlsPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteServerTlsPolicy(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteServerTlsPolicyOperation struct{}

func (op *deleteServerTlsPolicyOperation) do(ctx context.Context, r *ServerTlsPolicy, c *Client) error {

	_, err := c.GetServerTlsPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ServerTlsPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetServerTlsPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := serverTlsPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetServerTlsPolicy(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createServerTlsPolicyOperation struct {
	response map[string]interface{}
}

func (op *createServerTlsPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServerTlsPolicyOperation) do(ctx context.Context, r *ServerTlsPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := serverTlsPolicyCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetServerTlsPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServerTlsPolicyRaw(ctx context.Context, r *ServerTlsPolicy) ([]byte, error) {

	u, err := serverTlsPolicyGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) serverTlsPolicyDiffsForRawDesired(ctx context.Context, rawDesired *ServerTlsPolicy, opts ...dcl.ApplyOption) (initial, desired *ServerTlsPolicy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ServerTlsPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ServerTlsPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ServerTlsPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetServerTlsPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ServerTlsPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ServerTlsPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that ServerTlsPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServerTlsPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ServerTlsPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ServerTlsPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServerTlsPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ServerTlsPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServerTlsPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ServerTlsPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffServerTlsPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServerTlsPolicyInitialState(rawInitial, rawDesired *ServerTlsPolicy) (*ServerTlsPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServerTlsPolicyDesiredState(rawDesired, rawInitial *ServerTlsPolicy, opts ...dcl.ApplyOption) (*ServerTlsPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ServerCertificate = canonicalizeServerTlsPolicyServerCertificate(rawDesired.ServerCertificate, nil, opts...)
		rawDesired.MtlsPolicy = canonicalizeServerTlsPolicyMtlsPolicy(rawDesired.MtlsPolicy, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.BoolCanonicalize(rawDesired.AllowOpen, rawInitial.AllowOpen) {
		rawDesired.AllowOpen = rawInitial.AllowOpen
	}
	rawDesired.ServerCertificate = canonicalizeServerTlsPolicyServerCertificate(rawDesired.ServerCertificate, rawInitial.ServerCertificate, opts...)
	rawDesired.MtlsPolicy = canonicalizeServerTlsPolicyMtlsPolicy(rawDesired.MtlsPolicy, rawInitial.MtlsPolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeServerTlsPolicyNewState(c *Client, rawNew, rawDesired *ServerTlsPolicy) (*ServerTlsPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AllowOpen) && dcl.IsEmptyValueIndirect(rawDesired.AllowOpen) {
		rawNew.AllowOpen = rawDesired.AllowOpen
	} else {
		if dcl.BoolCanonicalize(rawDesired.AllowOpen, rawNew.AllowOpen) {
			rawNew.AllowOpen = rawDesired.AllowOpen
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerCertificate) && dcl.IsEmptyValueIndirect(rawDesired.ServerCertificate) {
		rawNew.ServerCertificate = rawDesired.ServerCertificate
	} else {
		rawNew.ServerCertificate = canonicalizeNewServerTlsPolicyServerCertificate(c, rawDesired.ServerCertificate, rawNew.ServerCertificate)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MtlsPolicy) && dcl.IsEmptyValueIndirect(rawDesired.MtlsPolicy) {
		rawNew.MtlsPolicy = rawDesired.MtlsPolicy
	} else {
		rawNew.MtlsPolicy = canonicalizeNewServerTlsPolicyMtlsPolicy(c, rawDesired.MtlsPolicy, rawNew.MtlsPolicy)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeServerTlsPolicyServerCertificate(des, initial *ServerTlsPolicyServerCertificate, opts ...dcl.ApplyOption) *ServerTlsPolicyServerCertificate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.GrpcEndpoint = canonicalizeServerTlsPolicyServerCertificateGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	des.CertificateProviderInstance = canonicalizeServerTlsPolicyServerCertificateCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return des
}

func canonicalizeNewServerTlsPolicyServerCertificate(c *Client, des, nw *ServerTlsPolicyServerCertificate) *ServerTlsPolicyServerCertificate {
	if des == nil || nw == nil {
		return nw
	}

	nw.GrpcEndpoint = canonicalizeNewServerTlsPolicyServerCertificateGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewServerTlsPolicyServerCertificateCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewServerTlsPolicyServerCertificateSet(c *Client, des, nw []ServerTlsPolicyServerCertificate) []ServerTlsPolicyServerCertificate {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyServerCertificate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyServerCertificateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyServerCertificateSlice(c *Client, des, nw []ServerTlsPolicyServerCertificate) []ServerTlsPolicyServerCertificate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyServerCertificate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyServerCertificate(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyServerCertificateGrpcEndpoint(des, initial *ServerTlsPolicyServerCertificateGrpcEndpoint, opts ...dcl.ApplyOption) *ServerTlsPolicyServerCertificateGrpcEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.TargetUri, initial.TargetUri) || dcl.IsZeroValue(des.TargetUri) {
		des.TargetUri = initial.TargetUri
	}

	return des
}

func canonicalizeNewServerTlsPolicyServerCertificateGrpcEndpoint(c *Client, des, nw *ServerTlsPolicyServerCertificateGrpcEndpoint) *ServerTlsPolicyServerCertificateGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewServerTlsPolicyServerCertificateGrpcEndpointSet(c *Client, des, nw []ServerTlsPolicyServerCertificateGrpcEndpoint) []ServerTlsPolicyServerCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyServerCertificateGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyServerCertificateGrpcEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyServerCertificateGrpcEndpointSlice(c *Client, des, nw []ServerTlsPolicyServerCertificateGrpcEndpoint) []ServerTlsPolicyServerCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyServerCertificateGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyServerCertificateGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyServerCertificateCertificateProviderInstance(des, initial *ServerTlsPolicyServerCertificateCertificateProviderInstance, opts ...dcl.ApplyOption) *ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.PluginInstance, initial.PluginInstance) || dcl.IsZeroValue(des.PluginInstance) {
		des.PluginInstance = initial.PluginInstance
	}

	return des
}

func canonicalizeNewServerTlsPolicyServerCertificateCertificateProviderInstance(c *Client, des, nw *ServerTlsPolicyServerCertificateCertificateProviderInstance) *ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewServerTlsPolicyServerCertificateCertificateProviderInstanceSet(c *Client, des, nw []ServerTlsPolicyServerCertificateCertificateProviderInstance) []ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyServerCertificateCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyServerCertificateCertificateProviderInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyServerCertificateCertificateProviderInstanceSlice(c *Client, des, nw []ServerTlsPolicyServerCertificateCertificateProviderInstance) []ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyServerCertificateCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyServerCertificateCertificateProviderInstance(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyMtlsPolicy(des, initial *ServerTlsPolicyMtlsPolicy, opts ...dcl.ApplyOption) *ServerTlsPolicyMtlsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ClientValidationCa) {
		des.ClientValidationCa = initial.ClientValidationCa
	}

	return des
}

func canonicalizeNewServerTlsPolicyMtlsPolicy(c *Client, des, nw *ServerTlsPolicyMtlsPolicy) *ServerTlsPolicyMtlsPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.ClientValidationCa = canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaSlice(c, des.ClientValidationCa, nw.ClientValidationCa)

	return nw
}

func canonicalizeNewServerTlsPolicyMtlsPolicySet(c *Client, des, nw []ServerTlsPolicyMtlsPolicy) []ServerTlsPolicyMtlsPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyMtlsPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyMtlsPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyMtlsPolicySlice(c *Client, des, nw []ServerTlsPolicyMtlsPolicy) []ServerTlsPolicyMtlsPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyMtlsPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyMtlsPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyMtlsPolicyClientValidationCa(des, initial *ServerTlsPolicyMtlsPolicyClientValidationCa, opts ...dcl.ApplyOption) *ServerTlsPolicyMtlsPolicyClientValidationCa {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.GrpcEndpoint = canonicalizeServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	des.CertificateProviderInstance = canonicalizeServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return des
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCa(c *Client, des, nw *ServerTlsPolicyMtlsPolicyClientValidationCa) *ServerTlsPolicyMtlsPolicyClientValidationCa {
	if des == nil || nw == nil {
		return nw
	}

	nw.GrpcEndpoint = canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaSet(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCa) []ServerTlsPolicyMtlsPolicyClientValidationCa {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyMtlsPolicyClientValidationCa
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyMtlsPolicyClientValidationCaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaSlice(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCa) []ServerTlsPolicyMtlsPolicyClientValidationCa {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyMtlsPolicyClientValidationCa
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCa(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(des, initial *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint, opts ...dcl.ApplyOption) *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.TargetUri, initial.TargetUri) || dcl.IsZeroValue(des.TargetUri) {
		des.TargetUri = initial.TargetUri
	}

	return des
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c *Client, des, nw *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSet(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSlice(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(des, initial *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance, opts ...dcl.ApplyOption) *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.PluginInstance, initial.PluginInstance) || dcl.IsZeroValue(des.PluginInstance) {
		des.PluginInstance = initial.PluginInstance
	}

	return des
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c *Client, des, nw *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSet(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSlice(c *Client, des, nw []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffServerTlsPolicy(c *Client, desired, actual *ServerTlsPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServerTlsPolicyUpdateServerTlsPolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServerTlsPolicyUpdateServerTlsPolicyOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowOpen, actual.AllowOpen, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServerTlsPolicyUpdateServerTlsPolicyOperation")}, fn.AddNest("AllowOpen")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerCertificate, actual.ServerCertificate, dcl.Info{ObjectFunction: compareServerTlsPolicyServerCertificateNewStyle, OperationSelector: dcl.TriggersOperation("updateServerTlsPolicyUpdateServerTlsPolicyOperation")}, fn.AddNest("ServerCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MtlsPolicy, actual.MtlsPolicy, dcl.Info{ObjectFunction: compareServerTlsPolicyMtlsPolicyNewStyle, OperationSelector: dcl.TriggersOperation("updateServerTlsPolicyUpdateServerTlsPolicyOperation")}, fn.AddNest("MtlsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareServerTlsPolicyServerCertificateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyServerCertificate)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyServerCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificate or *ServerTlsPolicyServerCertificate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyServerCertificate)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyServerCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareServerTlsPolicyServerCertificateGrpcEndpointNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareServerTlsPolicyServerCertificateCertificateProviderInstanceNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyServerCertificateGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyServerCertificateGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyServerCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificateGrpcEndpoint or *ServerTlsPolicyServerCertificateGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyServerCertificateGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyServerCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificateGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyServerCertificateCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyServerCertificateCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyServerCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificateCertificateProviderInstance or *ServerTlsPolicyServerCertificateCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyServerCertificateCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyServerCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyServerCertificateCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyMtlsPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyMtlsPolicy)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyMtlsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicy or *ServerTlsPolicyMtlsPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyMtlsPolicy)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyMtlsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ClientValidationCa, actual.ClientValidationCa, dcl.Info{ObjectFunction: compareServerTlsPolicyMtlsPolicyClientValidationCaNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClientValidationCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyMtlsPolicyClientValidationCaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyMtlsPolicyClientValidationCa)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyMtlsPolicyClientValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCa or *ServerTlsPolicyMtlsPolicyClientValidationCa", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyMtlsPolicyClientValidationCa)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyMtlsPolicyClientValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCa", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint or *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance or *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ServerTlsPolicy) urlNormalized() *ServerTlsPolicy {
	normalized := dcl.Copy(*r).(ServerTlsPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *ServerTlsPolicy) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServerTlsPolicy) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServerTlsPolicy) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServerTlsPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateServerTlsPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ServerTlsPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ServerTlsPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandServerTlsPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ServerTlsPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalServerTlsPolicy decodes JSON responses into the ServerTlsPolicy resource schema.
func unmarshalServerTlsPolicy(b []byte, c *Client) (*ServerTlsPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapServerTlsPolicy(m, c)
}

func unmarshalMapServerTlsPolicy(m map[string]interface{}, c *Client) (*ServerTlsPolicy, error) {

	return flattenServerTlsPolicy(c, m), nil
}

// expandServerTlsPolicy expands ServerTlsPolicy into a JSON request object.
func expandServerTlsPolicy(c *Client, f *ServerTlsPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/serverTlsPolicies/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.AllowOpen; !dcl.IsEmptyValueIndirect(v) {
		m["allowOpen"] = v
	}
	if v, err := expandServerTlsPolicyServerCertificate(c, f.ServerCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ServerCertificate into serverCertificate: %w", err)
	} else if v != nil {
		m["serverCertificate"] = v
	}
	if v, err := expandServerTlsPolicyMtlsPolicy(c, f.MtlsPolicy); err != nil {
		return nil, fmt.Errorf("error expanding MtlsPolicy into mtlsPolicy: %w", err)
	} else if v != nil {
		m["mtlsPolicy"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenServerTlsPolicy flattens ServerTlsPolicy from a JSON request object into the
// ServerTlsPolicy type.
func flattenServerTlsPolicy(c *Client, i interface{}) *ServerTlsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ServerTlsPolicy{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.AllowOpen = dcl.FlattenBool(m["allowOpen"])
	res.ServerCertificate = flattenServerTlsPolicyServerCertificate(c, m["serverCertificate"])
	res.MtlsPolicy = flattenServerTlsPolicyMtlsPolicy(c, m["mtlsPolicy"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandServerTlsPolicyServerCertificateMap expands the contents of ServerTlsPolicyServerCertificate into a JSON
// request object.
func expandServerTlsPolicyServerCertificateMap(c *Client, f map[string]ServerTlsPolicyServerCertificate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyServerCertificate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyServerCertificateSlice expands the contents of ServerTlsPolicyServerCertificate into a JSON
// request object.
func expandServerTlsPolicyServerCertificateSlice(c *Client, f []ServerTlsPolicyServerCertificate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyServerCertificate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyServerCertificateMap flattens the contents of ServerTlsPolicyServerCertificate from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateMap(c *Client, i interface{}) map[string]ServerTlsPolicyServerCertificate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyServerCertificate{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyServerCertificate{}
	}

	items := make(map[string]ServerTlsPolicyServerCertificate)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyServerCertificate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyServerCertificateSlice flattens the contents of ServerTlsPolicyServerCertificate from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateSlice(c *Client, i interface{}) []ServerTlsPolicyServerCertificate {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyServerCertificate{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyServerCertificate{}
	}

	items := make([]ServerTlsPolicyServerCertificate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyServerCertificate(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyServerCertificate expands an instance of ServerTlsPolicyServerCertificate into a JSON
// request object.
func expandServerTlsPolicyServerCertificate(c *Client, f *ServerTlsPolicyServerCertificate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServerTlsPolicyServerCertificateGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandServerTlsPolicyServerCertificateCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyServerCertificate flattens an instance of ServerTlsPolicyServerCertificate from a JSON
// response object.
func flattenServerTlsPolicyServerCertificate(c *Client, i interface{}) *ServerTlsPolicyServerCertificate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyServerCertificate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyServerCertificate
	}
	r.GrpcEndpoint = flattenServerTlsPolicyServerCertificateGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenServerTlsPolicyServerCertificateCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandServerTlsPolicyServerCertificateGrpcEndpointMap expands the contents of ServerTlsPolicyServerCertificateGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyServerCertificateGrpcEndpointMap(c *Client, f map[string]ServerTlsPolicyServerCertificateGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyServerCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyServerCertificateGrpcEndpointSlice expands the contents of ServerTlsPolicyServerCertificateGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyServerCertificateGrpcEndpointSlice(c *Client, f []ServerTlsPolicyServerCertificateGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyServerCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyServerCertificateGrpcEndpointMap flattens the contents of ServerTlsPolicyServerCertificateGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateGrpcEndpointMap(c *Client, i interface{}) map[string]ServerTlsPolicyServerCertificateGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyServerCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyServerCertificateGrpcEndpoint{}
	}

	items := make(map[string]ServerTlsPolicyServerCertificateGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyServerCertificateGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyServerCertificateGrpcEndpointSlice flattens the contents of ServerTlsPolicyServerCertificateGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateGrpcEndpointSlice(c *Client, i interface{}) []ServerTlsPolicyServerCertificateGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyServerCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyServerCertificateGrpcEndpoint{}
	}

	items := make([]ServerTlsPolicyServerCertificateGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyServerCertificateGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyServerCertificateGrpcEndpoint expands an instance of ServerTlsPolicyServerCertificateGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyServerCertificateGrpcEndpoint(c *Client, f *ServerTlsPolicyServerCertificateGrpcEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyServerCertificateGrpcEndpoint flattens an instance of ServerTlsPolicyServerCertificateGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateGrpcEndpoint(c *Client, i interface{}) *ServerTlsPolicyServerCertificateGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyServerCertificateGrpcEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyServerCertificateGrpcEndpoint
	}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandServerTlsPolicyServerCertificateCertificateProviderInstanceMap expands the contents of ServerTlsPolicyServerCertificateCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyServerCertificateCertificateProviderInstanceMap(c *Client, f map[string]ServerTlsPolicyServerCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyServerCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyServerCertificateCertificateProviderInstanceSlice expands the contents of ServerTlsPolicyServerCertificateCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyServerCertificateCertificateProviderInstanceSlice(c *Client, f []ServerTlsPolicyServerCertificateCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyServerCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyServerCertificateCertificateProviderInstanceMap flattens the contents of ServerTlsPolicyServerCertificateCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ServerTlsPolicyServerCertificateCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyServerCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyServerCertificateCertificateProviderInstance{}
	}

	items := make(map[string]ServerTlsPolicyServerCertificateCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyServerCertificateCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyServerCertificateCertificateProviderInstanceSlice flattens the contents of ServerTlsPolicyServerCertificateCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateCertificateProviderInstanceSlice(c *Client, i interface{}) []ServerTlsPolicyServerCertificateCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyServerCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyServerCertificateCertificateProviderInstance{}
	}

	items := make([]ServerTlsPolicyServerCertificateCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyServerCertificateCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyServerCertificateCertificateProviderInstance expands an instance of ServerTlsPolicyServerCertificateCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyServerCertificateCertificateProviderInstance(c *Client, f *ServerTlsPolicyServerCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyServerCertificateCertificateProviderInstance flattens an instance of ServerTlsPolicyServerCertificateCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyServerCertificateCertificateProviderInstance(c *Client, i interface{}) *ServerTlsPolicyServerCertificateCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyServerCertificateCertificateProviderInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyServerCertificateCertificateProviderInstance
	}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// expandServerTlsPolicyMtlsPolicyMap expands the contents of ServerTlsPolicyMtlsPolicy into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyMap(c *Client, f map[string]ServerTlsPolicyMtlsPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyMtlsPolicySlice expands the contents of ServerTlsPolicyMtlsPolicy into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicySlice(c *Client, f []ServerTlsPolicyMtlsPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyMtlsPolicyMap flattens the contents of ServerTlsPolicyMtlsPolicy from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyMap(c *Client, i interface{}) map[string]ServerTlsPolicyMtlsPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyMtlsPolicy{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyMtlsPolicy{}
	}

	items := make(map[string]ServerTlsPolicyMtlsPolicy)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyMtlsPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyMtlsPolicySlice flattens the contents of ServerTlsPolicyMtlsPolicy from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicySlice(c *Client, i interface{}) []ServerTlsPolicyMtlsPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyMtlsPolicy{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyMtlsPolicy{}
	}

	items := make([]ServerTlsPolicyMtlsPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyMtlsPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyMtlsPolicy expands an instance of ServerTlsPolicyMtlsPolicy into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicy(c *Client, f *ServerTlsPolicyMtlsPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServerTlsPolicyMtlsPolicyClientValidationCaSlice(c, f.ClientValidationCa); err != nil {
		return nil, fmt.Errorf("error expanding ClientValidationCa into clientValidationCa: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["clientValidationCa"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyMtlsPolicy flattens an instance of ServerTlsPolicyMtlsPolicy from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicy(c *Client, i interface{}) *ServerTlsPolicyMtlsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyMtlsPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyMtlsPolicy
	}
	r.ClientValidationCa = flattenServerTlsPolicyMtlsPolicyClientValidationCaSlice(c, m["clientValidationCa"])

	return r
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaMap expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCa into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaMap(c *Client, f map[string]ServerTlsPolicyMtlsPolicyClientValidationCa) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCa(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaSlice expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCa into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaSlice(c *Client, f []ServerTlsPolicyMtlsPolicyClientValidationCa) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCa(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaMap flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCa from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaMap(c *Client, i interface{}) map[string]ServerTlsPolicyMtlsPolicyClientValidationCa {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCa{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCa{}
	}

	items := make(map[string]ServerTlsPolicyMtlsPolicyClientValidationCa)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyMtlsPolicyClientValidationCa(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaSlice flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCa from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaSlice(c *Client, i interface{}) []ServerTlsPolicyMtlsPolicyClientValidationCa {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyMtlsPolicyClientValidationCa{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyMtlsPolicyClientValidationCa{}
	}

	items := make([]ServerTlsPolicyMtlsPolicyClientValidationCa, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyMtlsPolicyClientValidationCa(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyMtlsPolicyClientValidationCa expands an instance of ServerTlsPolicyMtlsPolicyClientValidationCa into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCa(c *Client, f *ServerTlsPolicyMtlsPolicyClientValidationCa) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCa flattens an instance of ServerTlsPolicyMtlsPolicyClientValidationCa from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCa(c *Client, i interface{}) *ServerTlsPolicyMtlsPolicyClientValidationCa {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyMtlsPolicyClientValidationCa{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyMtlsPolicyClientValidationCa
	}
	r.GrpcEndpoint = flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointMap expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointMap(c *Client, f map[string]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSlice expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSlice(c *Client, f []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointMap flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointMap(c *Client, i interface{}) map[string]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	}

	items := make(map[string]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSlice flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointSlice(c *Client, i interface{}) []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	}

	items := make([]ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint expands an instance of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c *Client, f *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint flattens an instance of ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(c *Client, i interface{}) *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint
	}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceMap expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceMap(c *Client, f map[string]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSlice expands the contents of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSlice(c *Client, f []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceMap flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	}

	items := make(map[string]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSlice flattens the contents of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceSlice(c *Client, i interface{}) []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	}

	items := make([]ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance expands an instance of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance into a JSON
// request object.
func expandServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c *Client, f *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance flattens an instance of ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(c *Client, i interface{}) *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance
	}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ServerTlsPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalServerTlsPolicy(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type serverTlsPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serverTlsPolicyApiOperation
}

func convertFieldDiffToServerTlsPolicyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serverTlsPolicyDiff, error) {
	var diffs []serverTlsPolicyDiff
	for _, op := range ops {
		diff := serverTlsPolicyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToserverTlsPolicyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToserverTlsPolicyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serverTlsPolicyApiOperation, error) {
	switch op {

	case "updateServerTlsPolicyUpdateServerTlsPolicyOperation":
		return &updateServerTlsPolicyUpdateServerTlsPolicyOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
