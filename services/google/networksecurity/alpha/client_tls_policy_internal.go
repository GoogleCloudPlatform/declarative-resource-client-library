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
package alpha

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

func (r *ClientTlsPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serverValidationCa"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ClientCertificate) {
		if err := r.ClientCertificate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LocalFilepath) {
		if err := r.LocalFilepath.validate(); err != nil {
			return err
		}
	}
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
func (r *ClientTlsPolicyClientCertificateLocalFilepath) validate() error {
	if err := dcl.Required(r, "certificatePath"); err != nil {
		return err
	}
	if err := dcl.Required(r, "privateKeyPath"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyServerValidationCa) validate() error {
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
func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) validate() error {
	if err := dcl.Required(r, "targetUri"); err != nil {
		return err
	}
	return nil
}
func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) validate() error {
	if err := dcl.Required(r, "pluginInstance"); err != nil {
		return err
	}
	return nil
}

func clientTlsPolicyGetURL(userBasePath string, r *ClientTlsPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func clientTlsPolicyListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func clientTlsPolicyCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies?clientTlsPolicyId={{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func clientTlsPolicyDeleteURL(userBasePath string, r *ClientTlsPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func (r *ClientTlsPolicy) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}:setIamPolicy", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields)
}

func (r *ClientTlsPolicy) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}:getIamPolicy", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields)
}

func (r *ClientTlsPolicy) IAMPolicyVersion() int {
	return 3
}

// clientTlsPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clientTlsPolicyApiOperation interface {
	do(context.Context, *ClientTlsPolicy, *Client) error
}

// newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest creates a request for an
// ClientTlsPolicy resource's UpdateClientTlsPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(ctx context.Context, f *ClientTlsPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Sni; !dcl.IsEmptyValueIndirect(v) {
		req["sni"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificate(c, f.ClientCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ClientCertificate into clientCertificate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clientCertificate"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaSlice(c, f.ServerValidationCa); err != nil {
		return nil, fmt.Errorf("error expanding ServerValidationCa into serverValidationCa: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serverValidationCa"] = v
	}
	return req, nil
}

// marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateClientTlsPolicyUpdateClientTlsPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClientTlsPolicyUpdateClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {
	_, err := c.GetClientTlsPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateClientTlsPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClientTlsPolicyUpdateClientTlsPolicyRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listClientTlsPolicyRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := clientTlsPolicyListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ClientTlsPolicyMaxPage {
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

type listClientTlsPolicyOperation struct {
	ClientTlsPolicies []map[string]interface{} `json:"clientTlsPolicies"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listClientTlsPolicy(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*ClientTlsPolicy, string, error) {
	b, err := c.listClientTlsPolicyRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClientTlsPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ClientTlsPolicy
	for _, v := range m.ClientTlsPolicies {
		res := flattenClientTlsPolicy(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllClientTlsPolicy(ctx context.Context, f func(*ClientTlsPolicy) bool, resources []*ClientTlsPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteClientTlsPolicy(ctx, res)
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

type deleteClientTlsPolicyOperation struct{}

func (op *deleteClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {

	_, err := c.GetClientTlsPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ClientTlsPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetClientTlsPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := clientTlsPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetClientTlsPolicy(ctx, r.urlNormalized())
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
type createClientTlsPolicyOperation struct {
	response map[string]interface{}
}

func (op *createClientTlsPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClientTlsPolicyOperation) do(ctx context.Context, r *ClientTlsPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := clientTlsPolicyCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetClientTlsPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getClientTlsPolicyRaw(ctx context.Context, r *ClientTlsPolicy) ([]byte, error) {

	u, err := clientTlsPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) clientTlsPolicyDiffsForRawDesired(ctx context.Context, rawDesired *ClientTlsPolicy, opts ...dcl.ApplyOption) (initial, desired *ClientTlsPolicy, diffs []clientTlsPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ClientTlsPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ClientTlsPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ClientTlsPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetClientTlsPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ClientTlsPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ClientTlsPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that ClientTlsPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ClientTlsPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ClientTlsPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClientTlsPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ClientTlsPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ClientTlsPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffClientTlsPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeClientTlsPolicyInitialState(rawInitial, rawDesired *ClientTlsPolicy) (*ClientTlsPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClientTlsPolicyDesiredState(rawDesired, rawInitial *ClientTlsPolicy, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ClientCertificate = canonicalizeClientTlsPolicyClientCertificate(rawDesired.ClientCertificate, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.Sni, rawInitial.Sni) {
		rawDesired.Sni = rawInitial.Sni
	}
	rawDesired.ClientCertificate = canonicalizeClientTlsPolicyClientCertificate(rawDesired.ClientCertificate, rawInitial.ClientCertificate, opts...)
	if dcl.IsZeroValue(rawDesired.ServerValidationCa) {
		rawDesired.ServerValidationCa = rawInitial.ServerValidationCa
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeClientTlsPolicyNewState(c *Client, rawNew, rawDesired *ClientTlsPolicy) (*ClientTlsPolicy, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Sni) && dcl.IsEmptyValueIndirect(rawDesired.Sni) {
		rawNew.Sni = rawDesired.Sni
	} else {
		if dcl.StringCanonicalize(rawDesired.Sni, rawNew.Sni) {
			rawNew.Sni = rawDesired.Sni
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientCertificate) && dcl.IsEmptyValueIndirect(rawDesired.ClientCertificate) {
		rawNew.ClientCertificate = rawDesired.ClientCertificate
	} else {
		rawNew.ClientCertificate = canonicalizeNewClientTlsPolicyClientCertificate(c, rawDesired.ClientCertificate, rawNew.ClientCertificate)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerValidationCa) && dcl.IsEmptyValueIndirect(rawDesired.ServerValidationCa) {
		rawNew.ServerValidationCa = rawDesired.ServerValidationCa
	} else {
		rawNew.ServerValidationCa = canonicalizeNewClientTlsPolicyServerValidationCaSlice(c, rawDesired.ServerValidationCa, rawNew.ServerValidationCa)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeClientTlsPolicyClientCertificate(des, initial *ClientTlsPolicyClientCertificate, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LocalFilepath = canonicalizeClientTlsPolicyClientCertificateLocalFilepath(des.LocalFilepath, initial.LocalFilepath, opts...)
	des.GrpcEndpoint = canonicalizeClientTlsPolicyClientCertificateGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	des.CertificateProviderInstance = canonicalizeClientTlsPolicyClientCertificateCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return des
}

func canonicalizeNewClientTlsPolicyClientCertificate(c *Client, des, nw *ClientTlsPolicyClientCertificate) *ClientTlsPolicyClientCertificate {
	if des == nil || nw == nil {
		return nw
	}

	nw.LocalFilepath = canonicalizeNewClientTlsPolicyClientCertificateLocalFilepath(c, des.LocalFilepath, nw.LocalFilepath)
	nw.GrpcEndpoint = canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateSet(c *Client, des, nw []ClientTlsPolicyClientCertificate) []ClientTlsPolicyClientCertificate {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyClientCertificate(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyClientCertificateSlice(c *Client, des, nw []ClientTlsPolicyClientCertificate) []ClientTlsPolicyClientCertificate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyClientCertificate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificate(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyClientCertificateLocalFilepath(des, initial *ClientTlsPolicyClientCertificateLocalFilepath, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificateLocalFilepath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CertificatePath, initial.CertificatePath) || dcl.IsZeroValue(des.CertificatePath) {
		des.CertificatePath = initial.CertificatePath
	}
	if dcl.StringCanonicalize(des.PrivateKeyPath, initial.PrivateKeyPath) || dcl.IsZeroValue(des.PrivateKeyPath) {
		des.PrivateKeyPath = initial.PrivateKeyPath
	}

	return des
}

func canonicalizeNewClientTlsPolicyClientCertificateLocalFilepath(c *Client, des, nw *ClientTlsPolicyClientCertificateLocalFilepath) *ClientTlsPolicyClientCertificateLocalFilepath {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CertificatePath, nw.CertificatePath) {
		nw.CertificatePath = des.CertificatePath
	}
	if dcl.StringCanonicalize(des.PrivateKeyPath, nw.PrivateKeyPath) {
		nw.PrivateKeyPath = des.PrivateKeyPath
	}

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateLocalFilepathSet(c *Client, des, nw []ClientTlsPolicyClientCertificateLocalFilepath) []ClientTlsPolicyClientCertificateLocalFilepath {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificateLocalFilepath
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyClientCertificateLocalFilepath(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyClientCertificateLocalFilepathSlice(c *Client, des, nw []ClientTlsPolicyClientCertificateLocalFilepath) []ClientTlsPolicyClientCertificateLocalFilepath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyClientCertificateLocalFilepath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificateLocalFilepath(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyClientCertificateGrpcEndpoint(des, initial *ClientTlsPolicyClientCertificateGrpcEndpoint, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificateGrpcEndpoint {
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

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, des, nw *ClientTlsPolicyClientCertificateGrpcEndpoint) *ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpointSet(c *Client, des, nw []ClientTlsPolicyClientCertificateGrpcEndpoint) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificateGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyClientCertificateGrpcEndpoint(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, des, nw []ClientTlsPolicyClientCertificateGrpcEndpoint) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyClientCertificateGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificateGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyClientCertificateCertificateProviderInstance(des, initial *ClientTlsPolicyClientCertificateCertificateProviderInstance, opts ...dcl.ApplyOption) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
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

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, des, nw *ClientTlsPolicyClientCertificateCertificateProviderInstance) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstanceSet(c *Client, des, nw []ClientTlsPolicyClientCertificateCertificateProviderInstance) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyClientCertificateCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyClientCertificateCertificateProviderInstance(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, des, nw []ClientTlsPolicyClientCertificateCertificateProviderInstance) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyClientCertificateCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyClientCertificateCertificateProviderInstance(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCa(des, initial *ClientTlsPolicyServerValidationCa, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCa {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CaCertPath, initial.CaCertPath) || dcl.IsZeroValue(des.CaCertPath) {
		des.CaCertPath = initial.CaCertPath
	}
	des.GrpcEndpoint = canonicalizeClientTlsPolicyServerValidationCaGrpcEndpoint(des.GrpcEndpoint, initial.GrpcEndpoint, opts...)
	des.CertificateProviderInstance = canonicalizeClientTlsPolicyServerValidationCaCertificateProviderInstance(des.CertificateProviderInstance, initial.CertificateProviderInstance, opts...)

	return des
}

func canonicalizeNewClientTlsPolicyServerValidationCa(c *Client, des, nw *ClientTlsPolicyServerValidationCa) *ClientTlsPolicyServerValidationCa {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CaCertPath, nw.CaCertPath) {
		nw.CaCertPath = des.CaCertPath
	}
	nw.GrpcEndpoint = canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c, des.GrpcEndpoint, nw.GrpcEndpoint)
	nw.CertificateProviderInstance = canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c, des.CertificateProviderInstance, nw.CertificateProviderInstance)

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaSet(c *Client, des, nw []ClientTlsPolicyServerValidationCa) []ClientTlsPolicyServerValidationCa {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCa
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyServerValidationCa(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyServerValidationCaSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCa) []ClientTlsPolicyServerValidationCa {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyServerValidationCa
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCa(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCaGrpcEndpoint(des, initial *ClientTlsPolicyServerValidationCaGrpcEndpoint, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
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

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, des, nw *ClientTlsPolicyServerValidationCaGrpcEndpoint) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TargetUri, nw.TargetUri) {
		nw.TargetUri = des.TargetUri
	}

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpointSet(c *Client, des, nw []ClientTlsPolicyServerValidationCaGrpcEndpoint) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCaGrpcEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyServerValidationCaGrpcEndpoint(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCaGrpcEndpoint) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyServerValidationCaGrpcEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCaGrpcEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeClientTlsPolicyServerValidationCaCertificateProviderInstance(des, initial *ClientTlsPolicyServerValidationCaCertificateProviderInstance, opts ...dcl.ApplyOption) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
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

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, des, nw *ClientTlsPolicyServerValidationCaCertificateProviderInstance) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PluginInstance, nw.PluginInstance) {
		nw.PluginInstance = des.PluginInstance
	}

	return nw
}

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstanceSet(c *Client, des, nw []ClientTlsPolicyServerValidationCaCertificateProviderInstance) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ClientTlsPolicyServerValidationCaCertificateProviderInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &d, &n) {
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

func canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, des, nw []ClientTlsPolicyServerValidationCaCertificateProviderInstance) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClientTlsPolicyServerValidationCaCertificateProviderInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &d, &n))
	}

	return items
}

type clientTlsPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clientTlsPolicyApiOperation
	Diffs            []*dcl.FieldDiff
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffClientTlsPolicy(c *Client, desired, actual *ClientTlsPolicy, opts ...dcl.ApplyOption) ([]clientTlsPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []clientTlsPolicyDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{
			UpdateOp: &updateClientTlsPolicyUpdateClientTlsPolicyOperation{}, Diffs: ds,
			FieldName: "Description",
		})
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CreateTime",
		})
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "UpdateTime",
		})
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{
			UpdateOp: &updateClientTlsPolicyUpdateClientTlsPolicyOperation{}, Diffs: ds,
			FieldName: "Labels",
		})
	}

	if ds, err := dcl.Diff(desired.Sni, actual.Sni, dcl.Info{}, fn.AddNest("Sni")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{
			UpdateOp: &updateClientTlsPolicyUpdateClientTlsPolicyOperation{}, Diffs: ds,
			FieldName: "Sni",
		})
	}

	if ds, err := dcl.Diff(desired.ClientCertificate, actual.ClientCertificate, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateNewStyle}, fn.AddNest("ClientCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{
			UpdateOp: &updateClientTlsPolicyUpdateClientTlsPolicyOperation{}, Diffs: ds,
			FieldName: "ClientCertificate",
		})
	}

	if ds, err := dcl.Diff(desired.ServerValidationCa, actual.ServerValidationCa, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaNewStyle}, fn.AddNest("ServerValidationCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{
			UpdateOp: &updateClientTlsPolicyUpdateClientTlsPolicyOperation{}, Diffs: ds,
			FieldName: "ServerValidationCa",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clientTlsPolicyDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
		})
	}

	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []clientTlsPolicyDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareClientTlsPolicyClientCertificateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificate)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificate or *ClientTlsPolicyClientCertificate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificate)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalFilepath, actual.LocalFilepath, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateLocalFilepathNewStyle}, fn.AddNest("LocalFilepath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateGrpcEndpointNewStyle}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareClientTlsPolicyClientCertificateCertificateProviderInstanceNewStyle}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificate(c *Client, desired, actual *ClientTlsPolicyClientCertificate) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClientTlsPolicyClientCertificateLocalFilepath(c, desired.LocalFilepath, actual.LocalFilepath) && !dcl.IsZeroValue(desired.LocalFilepath) {
		c.Config.Logger.Infof("Diff in LocalFilepath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalFilepath), dcl.SprintResource(actual.LocalFilepath))
		return true
	}
	if compareClientTlsPolicyClientCertificateGrpcEndpoint(c, desired.GrpcEndpoint, actual.GrpcEndpoint) && !dcl.IsZeroValue(desired.GrpcEndpoint) {
		c.Config.Logger.Infof("Diff in GrpcEndpoint.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrpcEndpoint), dcl.SprintResource(actual.GrpcEndpoint))
		return true
	}
	if compareClientTlsPolicyClientCertificateCertificateProviderInstance(c, desired.CertificateProviderInstance, actual.CertificateProviderInstance) && !dcl.IsZeroValue(desired.CertificateProviderInstance) {
		c.Config.Logger.Infof("Diff in CertificateProviderInstance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CertificateProviderInstance), dcl.SprintResource(actual.CertificateProviderInstance))
		return true
	}
	return false
}

func compareClientTlsPolicyClientCertificateSlice(c *Client, desired, actual []ClientTlsPolicyClientCertificate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificate, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyClientCertificate(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificate, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateMap(c *Client, desired, actual map[string]ClientTlsPolicyClientCertificate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificate, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificate, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyClientCertificate(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificate, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateLocalFilepathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificateLocalFilepath)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificateLocalFilepath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateLocalFilepath or *ClientTlsPolicyClientCertificateLocalFilepath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificateLocalFilepath)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificateLocalFilepath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateLocalFilepath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CertificatePath, actual.CertificatePath, dcl.Info{}, fn.AddNest("CertificatePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateKeyPath, actual.PrivateKeyPath, dcl.Info{}, fn.AddNest("PrivateKeyPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificateLocalFilepath(c *Client, desired, actual *ClientTlsPolicyClientCertificateLocalFilepath) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.CertificatePath, actual.CertificatePath) && !dcl.IsZeroValue(desired.CertificatePath) {
		c.Config.Logger.Infof("Diff in CertificatePath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CertificatePath), dcl.SprintResource(actual.CertificatePath))
		return true
	}
	if !dcl.StringCanonicalize(desired.PrivateKeyPath, actual.PrivateKeyPath) && !dcl.IsZeroValue(desired.PrivateKeyPath) {
		c.Config.Logger.Infof("Diff in PrivateKeyPath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivateKeyPath), dcl.SprintResource(actual.PrivateKeyPath))
		return true
	}
	return false
}

func compareClientTlsPolicyClientCertificateLocalFilepathSlice(c *Client, desired, actual []ClientTlsPolicyClientCertificateLocalFilepath) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateLocalFilepath, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyClientCertificateLocalFilepath(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateLocalFilepath, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateLocalFilepathMap(c *Client, desired, actual map[string]ClientTlsPolicyClientCertificateLocalFilepath) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateLocalFilepath, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateLocalFilepath, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyClientCertificateLocalFilepath(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateLocalFilepath, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificateGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateGrpcEndpoint or *ClientTlsPolicyClientCertificateGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificateGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificateGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, desired, actual *ClientTlsPolicyClientCertificateGrpcEndpoint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.TargetUri, actual.TargetUri) && !dcl.IsZeroValue(desired.TargetUri) {
		c.Config.Logger.Infof("Diff in TargetUri.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetUri), dcl.SprintResource(actual.TargetUri))
		return true
	}
	return false
}

func compareClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, desired, actual []ClientTlsPolicyClientCertificateGrpcEndpoint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateGrpcEndpoint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyClientCertificateGrpcEndpoint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateGrpcEndpoint, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateGrpcEndpointMap(c *Client, desired, actual map[string]ClientTlsPolicyClientCertificateGrpcEndpoint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateGrpcEndpoint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateGrpcEndpoint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyClientCertificateGrpcEndpoint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateGrpcEndpoint, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyClientCertificateCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyClientCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateCertificateProviderInstance or *ClientTlsPolicyClientCertificateCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyClientCertificateCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyClientCertificateCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyClientCertificateCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, desired, actual *ClientTlsPolicyClientCertificateCertificateProviderInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.PluginInstance, actual.PluginInstance) && !dcl.IsZeroValue(desired.PluginInstance) {
		c.Config.Logger.Infof("Diff in PluginInstance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PluginInstance), dcl.SprintResource(actual.PluginInstance))
		return true
	}
	return false
}

func compareClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, desired, actual []ClientTlsPolicyClientCertificateCertificateProviderInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateCertificateProviderInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyClientCertificateCertificateProviderInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateCertificateProviderInstance, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyClientCertificateCertificateProviderInstanceMap(c *Client, desired, actual map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyClientCertificateCertificateProviderInstance, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateCertificateProviderInstance, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyClientCertificateCertificateProviderInstance(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyClientCertificateCertificateProviderInstance, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCa)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCa or *ClientTlsPolicyServerValidationCa", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCa)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCa", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CaCertPath, actual.CaCertPath, dcl.Info{}, fn.AddNest("CaCertPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrpcEndpoint, actual.GrpcEndpoint, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaGrpcEndpointNewStyle}, fn.AddNest("GrpcEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificateProviderInstance, actual.CertificateProviderInstance, dcl.Info{ObjectFunction: compareClientTlsPolicyServerValidationCaCertificateProviderInstanceNewStyle}, fn.AddNest("CertificateProviderInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCa(c *Client, desired, actual *ClientTlsPolicyServerValidationCa) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.CaCertPath, actual.CaCertPath) && !dcl.IsZeroValue(desired.CaCertPath) {
		c.Config.Logger.Infof("Diff in CaCertPath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertPath), dcl.SprintResource(actual.CaCertPath))
		return true
	}
	if compareClientTlsPolicyServerValidationCaGrpcEndpoint(c, desired.GrpcEndpoint, actual.GrpcEndpoint) && !dcl.IsZeroValue(desired.GrpcEndpoint) {
		c.Config.Logger.Infof("Diff in GrpcEndpoint.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrpcEndpoint), dcl.SprintResource(actual.GrpcEndpoint))
		return true
	}
	if compareClientTlsPolicyServerValidationCaCertificateProviderInstance(c, desired.CertificateProviderInstance, actual.CertificateProviderInstance) && !dcl.IsZeroValue(desired.CertificateProviderInstance) {
		c.Config.Logger.Infof("Diff in CertificateProviderInstance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CertificateProviderInstance), dcl.SprintResource(actual.CertificateProviderInstance))
		return true
	}
	return false
}

func compareClientTlsPolicyServerValidationCaSlice(c *Client, desired, actual []ClientTlsPolicyServerValidationCa) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCa, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyServerValidationCa(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCa, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaMap(c *Client, desired, actual map[string]ClientTlsPolicyServerValidationCa) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCa, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCa, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyServerValidationCa(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCa, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaGrpcEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCaGrpcEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaGrpcEndpoint or *ClientTlsPolicyServerValidationCaGrpcEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCaGrpcEndpoint)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCaGrpcEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaGrpcEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetUri, actual.TargetUri, dcl.Info{}, fn.AddNest("TargetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, desired, actual *ClientTlsPolicyServerValidationCaGrpcEndpoint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.TargetUri, actual.TargetUri) && !dcl.IsZeroValue(desired.TargetUri) {
		c.Config.Logger.Infof("Diff in TargetUri.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetUri), dcl.SprintResource(actual.TargetUri))
		return true
	}
	return false
}

func compareClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, desired, actual []ClientTlsPolicyServerValidationCaGrpcEndpoint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCaGrpcEndpoint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyServerValidationCaGrpcEndpoint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaGrpcEndpoint, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaGrpcEndpointMap(c *Client, desired, actual map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCaGrpcEndpoint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaGrpcEndpoint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyServerValidationCaGrpcEndpoint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaGrpcEndpoint, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaCertificateProviderInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	if !ok {
		desiredNotPointer, ok := d.(ClientTlsPolicyServerValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaCertificateProviderInstance or *ClientTlsPolicyServerValidationCaCertificateProviderInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	if !ok {
		actualNotPointer, ok := a.(ClientTlsPolicyServerValidationCaCertificateProviderInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClientTlsPolicyServerValidationCaCertificateProviderInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PluginInstance, actual.PluginInstance, dcl.Info{}, fn.AddNest("PluginInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, desired, actual *ClientTlsPolicyServerValidationCaCertificateProviderInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.PluginInstance, actual.PluginInstance) && !dcl.IsZeroValue(desired.PluginInstance) {
		c.Config.Logger.Infof("Diff in PluginInstance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PluginInstance), dcl.SprintResource(actual.PluginInstance))
		return true
	}
	return false
}

func compareClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, desired, actual []ClientTlsPolicyServerValidationCaCertificateProviderInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCaCertificateProviderInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaCertificateProviderInstance, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClientTlsPolicyServerValidationCaCertificateProviderInstanceMap(c *Client, desired, actual map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClientTlsPolicyServerValidationCaCertificateProviderInstance, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaCertificateProviderInstance, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClientTlsPolicyServerValidationCaCertificateProviderInstance, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ClientTlsPolicy) urlNormalized() *ClientTlsPolicy {
	normalized := dcl.Copy(*r).(ClientTlsPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Sni = dcl.SelfLinkToName(r.Sni)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *ClientTlsPolicy) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ClientTlsPolicy) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ClientTlsPolicy) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ClientTlsPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateClientTlsPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ClientTlsPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ClientTlsPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandClientTlsPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ClientTlsPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalClientTlsPolicy decodes JSON responses into the ClientTlsPolicy resource schema.
func unmarshalClientTlsPolicy(b []byte, c *Client) (*ClientTlsPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapClientTlsPolicy(m, c)
}

func unmarshalMapClientTlsPolicy(m map[string]interface{}, c *Client) (*ClientTlsPolicy, error) {

	return flattenClientTlsPolicy(c, m), nil
}

// expandClientTlsPolicy expands ClientTlsPolicy into a JSON request object.
func expandClientTlsPolicy(c *Client, f *ClientTlsPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/clientTlsPolicies/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
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
	if v := f.Sni; !dcl.IsEmptyValueIndirect(v) {
		m["sni"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificate(c, f.ClientCertificate); err != nil {
		return nil, fmt.Errorf("error expanding ClientCertificate into clientCertificate: %w", err)
	} else if v != nil {
		m["clientCertificate"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaSlice(c, f.ServerValidationCa); err != nil {
		return nil, fmt.Errorf("error expanding ServerValidationCa into serverValidationCa: %w", err)
	} else if v != nil {
		m["serverValidationCa"] = v
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

// flattenClientTlsPolicy flattens ClientTlsPolicy from a JSON request object into the
// ClientTlsPolicy type.
func flattenClientTlsPolicy(c *Client, i interface{}) *ClientTlsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &ClientTlsPolicy{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Sni = dcl.FlattenString(m["sni"])
	r.ClientCertificate = flattenClientTlsPolicyClientCertificate(c, m["clientCertificate"])
	r.ServerValidationCa = flattenClientTlsPolicyServerValidationCaSlice(c, m["serverValidationCa"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandClientTlsPolicyClientCertificateMap expands the contents of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificateMap(c *Client, f map[string]ClientTlsPolicyClientCertificate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateSlice expands the contents of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificateSlice(c *Client, f []ClientTlsPolicyClientCertificate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateMap flattens the contents of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificate{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificate{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificate)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateSlice flattens the contents of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificate {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificate{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificate{}
	}

	items := make([]ClientTlsPolicyClientCertificate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificate(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificate expands an instance of ClientTlsPolicyClientCertificate into a JSON
// request object.
func expandClientTlsPolicyClientCertificate(c *Client, f *ClientTlsPolicyClientCertificate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClientTlsPolicyClientCertificateLocalFilepath(c, f.LocalFilepath); err != nil {
		return nil, fmt.Errorf("error expanding LocalFilepath into localFilepath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localFilepath"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificate flattens an instance of ClientTlsPolicyClientCertificate from a JSON
// response object.
func flattenClientTlsPolicyClientCertificate(c *Client, i interface{}) *ClientTlsPolicyClientCertificate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificate{}
	r.LocalFilepath = flattenClientTlsPolicyClientCertificateLocalFilepath(c, m["localFilepath"])
	r.GrpcEndpoint = flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandClientTlsPolicyClientCertificateLocalFilepathMap expands the contents of ClientTlsPolicyClientCertificateLocalFilepath into a JSON
// request object.
func expandClientTlsPolicyClientCertificateLocalFilepathMap(c *Client, f map[string]ClientTlsPolicyClientCertificateLocalFilepath) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificateLocalFilepath(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateLocalFilepathSlice expands the contents of ClientTlsPolicyClientCertificateLocalFilepath into a JSON
// request object.
func expandClientTlsPolicyClientCertificateLocalFilepathSlice(c *Client, f []ClientTlsPolicyClientCertificateLocalFilepath) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificateLocalFilepath(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateLocalFilepathMap flattens the contents of ClientTlsPolicyClientCertificateLocalFilepath from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateLocalFilepathMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificateLocalFilepath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificateLocalFilepath{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificateLocalFilepath{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificateLocalFilepath)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificateLocalFilepath(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateLocalFilepathSlice flattens the contents of ClientTlsPolicyClientCertificateLocalFilepath from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateLocalFilepathSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificateLocalFilepath {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificateLocalFilepath{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificateLocalFilepath{}
	}

	items := make([]ClientTlsPolicyClientCertificateLocalFilepath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificateLocalFilepath(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificateLocalFilepath expands an instance of ClientTlsPolicyClientCertificateLocalFilepath into a JSON
// request object.
func expandClientTlsPolicyClientCertificateLocalFilepath(c *Client, f *ClientTlsPolicyClientCertificateLocalFilepath) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.CertificatePath; !dcl.IsEmptyValueIndirect(v) {
		m["certificatePath"] = v
	}
	if v := f.PrivateKeyPath; !dcl.IsEmptyValueIndirect(v) {
		m["privateKeyPath"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificateLocalFilepath flattens an instance of ClientTlsPolicyClientCertificateLocalFilepath from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateLocalFilepath(c *Client, i interface{}) *ClientTlsPolicyClientCertificateLocalFilepath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificateLocalFilepath{}
	r.CertificatePath = dcl.FlattenString(m["certificatePath"])
	r.PrivateKeyPath = dcl.FlattenString(m["privateKeyPath"])

	return r
}

// expandClientTlsPolicyClientCertificateGrpcEndpointMap expands the contents of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpointMap(c *Client, f map[string]ClientTlsPolicyClientCertificateGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateGrpcEndpointSlice expands the contents of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, f []ClientTlsPolicyClientCertificateGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificateGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateGrpcEndpointMap flattens the contents of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpointMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificateGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificateGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateGrpcEndpointSlice flattens the contents of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpointSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificateGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificateGrpcEndpoint{}
	}

	items := make([]ClientTlsPolicyClientCertificateGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificateGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificateGrpcEndpoint expands an instance of ClientTlsPolicyClientCertificateGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, f *ClientTlsPolicyClientCertificateGrpcEndpoint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificateGrpcEndpoint flattens an instance of ClientTlsPolicyClientCertificateGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateGrpcEndpoint(c *Client, i interface{}) *ClientTlsPolicyClientCertificateGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificateGrpcEndpoint{}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstanceMap expands the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstanceMap(c *Client, f map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstanceSlice expands the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, f []ClientTlsPolicyClientCertificateCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyClientCertificateCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstanceMap flattens the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	items := make(map[string]ClientTlsPolicyClientCertificateCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstanceSlice flattens the contents of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstanceSlice(c *Client, i interface{}) []ClientTlsPolicyClientCertificateCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	}

	items := make([]ClientTlsPolicyClientCertificateCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyClientCertificateCertificateProviderInstance expands an instance of ClientTlsPolicyClientCertificateCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, f *ClientTlsPolicyClientCertificateCertificateProviderInstance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyClientCertificateCertificateProviderInstance flattens an instance of ClientTlsPolicyClientCertificateCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyClientCertificateCertificateProviderInstance(c *Client, i interface{}) *ClientTlsPolicyClientCertificateCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyClientCertificateCertificateProviderInstance{}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// expandClientTlsPolicyServerValidationCaMap expands the contents of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaMap(c *Client, f map[string]ClientTlsPolicyServerValidationCa) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCa(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaSlice expands the contents of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaSlice(c *Client, f []ClientTlsPolicyServerValidationCa) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCa(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaMap flattens the contents of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCa {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCa{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCa{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCa)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCa(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaSlice flattens the contents of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCa {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCa{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCa{}
	}

	items := make([]ClientTlsPolicyServerValidationCa, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCa(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCa expands an instance of ClientTlsPolicyServerValidationCa into a JSON
// request object.
func expandClientTlsPolicyServerValidationCa(c *Client, f *ClientTlsPolicyServerValidationCa) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.CaCertPath; !dcl.IsEmptyValueIndirect(v) {
		m["caCertPath"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, f.GrpcEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding GrpcEndpoint into grpcEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["grpcEndpoint"] = v
	}
	if v, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, f.CertificateProviderInstance); err != nil {
		return nil, fmt.Errorf("error expanding CertificateProviderInstance into certificateProviderInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certificateProviderInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCa flattens an instance of ClientTlsPolicyServerValidationCa from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCa(c *Client, i interface{}) *ClientTlsPolicyServerValidationCa {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCa{}
	r.CaCertPath = dcl.FlattenString(m["caCertPath"])
	r.GrpcEndpoint = flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, m["grpcEndpoint"])
	r.CertificateProviderInstance = flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, m["certificateProviderInstance"])

	return r
}

// expandClientTlsPolicyServerValidationCaGrpcEndpointMap expands the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpointMap(c *Client, f map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaGrpcEndpointSlice expands the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, f []ClientTlsPolicyServerValidationCaGrpcEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaGrpcEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpointMap flattens the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpointMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCaGrpcEndpoint)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpointSlice flattens the contents of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpointSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCaGrpcEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	}

	items := make([]ClientTlsPolicyServerValidationCaGrpcEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCaGrpcEndpoint expands an instance of ClientTlsPolicyServerValidationCaGrpcEndpoint into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, f *ClientTlsPolicyServerValidationCaGrpcEndpoint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.TargetUri; !dcl.IsEmptyValueIndirect(v) {
		m["targetUri"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCaGrpcEndpoint flattens an instance of ClientTlsPolicyServerValidationCaGrpcEndpoint from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaGrpcEndpoint(c *Client, i interface{}) *ClientTlsPolicyServerValidationCaGrpcEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCaGrpcEndpoint{}
	r.TargetUri = dcl.FlattenString(m["targetUri"])

	return r
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstanceMap expands the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstanceMap(c *Client, f map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice expands the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, f []ClientTlsPolicyServerValidationCaCertificateProviderInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceMap flattens the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceMap(c *Client, i interface{}) map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	items := make(map[string]ClientTlsPolicyServerValidationCaCertificateProviderInstance)
	for k, item := range a {
		items[k] = *flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice flattens the contents of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstanceSlice(c *Client, i interface{}) []ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	if len(a) == 0 {
		return []ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	}

	items := make([]ClientTlsPolicyServerValidationCaCertificateProviderInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandClientTlsPolicyServerValidationCaCertificateProviderInstance expands an instance of ClientTlsPolicyServerValidationCaCertificateProviderInstance into a JSON
// request object.
func expandClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, f *ClientTlsPolicyServerValidationCaCertificateProviderInstance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.PluginInstance; !dcl.IsEmptyValueIndirect(v) {
		m["pluginInstance"] = v
	}

	return m, nil
}

// flattenClientTlsPolicyServerValidationCaCertificateProviderInstance flattens an instance of ClientTlsPolicyServerValidationCaCertificateProviderInstance from a JSON
// response object.
func flattenClientTlsPolicyServerValidationCaCertificateProviderInstance(c *Client, i interface{}) *ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	r.PluginInstance = dcl.FlattenString(m["pluginInstance"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ClientTlsPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalClientTlsPolicy(b, c)
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
