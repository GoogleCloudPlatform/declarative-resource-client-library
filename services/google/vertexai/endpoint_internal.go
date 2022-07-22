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
package vertexai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Endpoint) validate() error {

	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionSpec) {
		if err := r.EncryptionSpec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointDeployedModels) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"DedicatedResources", "AutomaticResources"}, r.DedicatedResources, r.AutomaticResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DedicatedResources) {
		if err := r.DedicatedResources.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AutomaticResources) {
		if err := r.AutomaticResources.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PrivateEndpoints) {
		if err := r.PrivateEndpoints.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointDeployedModelsDedicatedResources) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MachineSpec) {
		if err := r.MachineSpec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointDeployedModelsDedicatedResourcesMachineSpec) validate() error {
	return nil
}
func (r *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) validate() error {
	return nil
}
func (r *EndpointDeployedModelsAutomaticResources) validate() error {
	return nil
}
func (r *EndpointDeployedModelsPrivateEndpoints) validate() error {
	return nil
}
func (r *EndpointEncryptionSpec) validate() error {
	if err := dcl.Required(r, "kmsKeyName"); err != nil {
		return err
	}
	return nil
}
func (r *Endpoint) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1/", params)
}

func (r *Endpoint) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Endpoint) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints", nr.basePath(), userBasePath, params), nil

}

func (r *Endpoint) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints", nr.basePath(), userBasePath, params), nil

}

func (r *Endpoint) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{name}}", nr.basePath(), userBasePath, params), nil
}

// endpointApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type endpointApiOperation interface {
	do(context.Context, *Endpoint, *Client) error
}

// newUpdateEndpointUpdateEndpointRequest creates a request for an
// Endpoint resource's UpdateEndpoint update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEndpointUpdateEndpointRequest(ctx context.Context, f *Endpoint, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getEndpointRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateEndpointUpdateEndpointRequest converts the update into
// the final JSON request body.
func marshalUpdateEndpointUpdateEndpointRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEndpointUpdateEndpointOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEndpointUpdateEndpointOperation) do(ctx context.Context, r *Endpoint, c *Client) error {
	_, err := c.GetEndpoint(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEndpoint")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateEndpointUpdateEndpointRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEndpointUpdateEndpointRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listEndpointRaw(ctx context.Context, r *Endpoint, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EndpointMaxPage {
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

type listEndpointOperation struct {
	Endpoints []map[string]interface{} `json:"endpoints"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listEndpoint(ctx context.Context, r *Endpoint, pageToken string, pageSize int32) ([]*Endpoint, string, error) {
	b, err := c.listEndpointRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEndpointOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Endpoint
	for _, v := range m.Endpoints {
		res, err := unmarshalMapEndpoint(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEndpoint(ctx context.Context, f func(*Endpoint) bool, resources []*Endpoint) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEndpoint(ctx, res)
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

type deleteEndpointOperation struct{}

func (op *deleteEndpointOperation) do(ctx context.Context, r *Endpoint, c *Client) error {
	r, err := c.GetEndpoint(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Endpoint not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetEndpoint checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetEndpoint(ctx, r)
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

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createEndpointOperation struct {
	response map[string]interface{}
}

func (op *createEndpointOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEndpointOperation) do(ctx context.Context, r *Endpoint, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	if r.Name != nil {
		// Allowing creation to continue with Name set could result in a Endpoint with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetEndpoint(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEndpointRaw(ctx context.Context, r *Endpoint) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) endpointDiffsForRawDesired(ctx context.Context, rawDesired *Endpoint, opts ...dcl.ApplyOption) (initial, desired *Endpoint, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Endpoint
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Endpoint); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Endpoint, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeEndpointDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEndpoint(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Endpoint resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Endpoint resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Endpoint resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEndpointDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Endpoint: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Endpoint: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEndpointFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEndpointInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Endpoint: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEndpointDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Endpoint: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEndpoint(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEndpointInitialState(rawInitial, rawDesired *Endpoint) (*Endpoint, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEndpointDesiredState(rawDesired, rawInitial *Endpoint, opts ...dcl.ApplyOption) (*Endpoint, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.EncryptionSpec = canonicalizeEndpointEncryptionSpec(rawDesired.EncryptionSpec, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Endpoint{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.EncryptionSpec = canonicalizeEndpointEncryptionSpec(rawDesired.EncryptionSpec, rawInitial.EncryptionSpec, opts...)
	if dcl.IsZeroValue(rawDesired.Network) || (dcl.IsEmptyValueIndirect(rawDesired.Network) && dcl.IsEmptyValueIndirect(rawInitial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Network = rawInitial.Network
	} else {
		canonicalDesired.Network = rawDesired.Network
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeEndpointNewState(c *Client, rawNew, rawDesired *Endpoint) (*Endpoint, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeployedModels) && dcl.IsEmptyValueIndirect(rawDesired.DeployedModels) {
		rawNew.DeployedModels = rawDesired.DeployedModels
	} else {
		rawNew.DeployedModels = canonicalizeNewEndpointDeployedModelsSlice(c, rawDesired.DeployedModels, rawNew.DeployedModels)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EncryptionSpec) && dcl.IsEmptyValueIndirect(rawDesired.EncryptionSpec) {
		rawNew.EncryptionSpec = rawDesired.EncryptionSpec
	} else {
		rawNew.EncryptionSpec = canonicalizeNewEndpointEncryptionSpec(c, rawDesired.EncryptionSpec, rawNew.EncryptionSpec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ModelDeploymentMonitoringJob) && dcl.IsEmptyValueIndirect(rawDesired.ModelDeploymentMonitoringJob) {
		rawNew.ModelDeploymentMonitoringJob = rawDesired.ModelDeploymentMonitoringJob
	} else {
		if dcl.StringCanonicalize(rawDesired.ModelDeploymentMonitoringJob, rawNew.ModelDeploymentMonitoringJob) {
			rawNew.ModelDeploymentMonitoringJob = rawDesired.ModelDeploymentMonitoringJob
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeEndpointDeployedModels(des, initial *EndpointDeployedModels, opts ...dcl.ApplyOption) *EndpointDeployedModels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.DedicatedResources != nil || (initial != nil && initial.DedicatedResources != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.AutomaticResources) {
			des.DedicatedResources = nil
			if initial != nil {
				initial.DedicatedResources = nil
			}
		}
	}

	if des.AutomaticResources != nil || (initial != nil && initial.AutomaticResources != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DedicatedResources) {
			des.AutomaticResources = nil
			if initial != nil {
				initial.AutomaticResources = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModels{}

	cDes.DedicatedResources = canonicalizeEndpointDeployedModelsDedicatedResources(des.DedicatedResources, initial.DedicatedResources, opts...)
	cDes.AutomaticResources = canonicalizeEndpointDeployedModelsAutomaticResources(des.AutomaticResources, initial.AutomaticResources, opts...)
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.IsZeroValue(des.Model) || (dcl.IsEmptyValueIndirect(des.Model) && dcl.IsEmptyValueIndirect(initial.Model)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Model = initial.Model
	} else {
		cDes.Model = des.Model
	}
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		cDes.DisplayName = initial.DisplayName
	} else {
		cDes.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	if dcl.BoolCanonicalize(des.DisableContainerLogging, initial.DisableContainerLogging) || dcl.IsZeroValue(des.DisableContainerLogging) {
		cDes.DisableContainerLogging = initial.DisableContainerLogging
	} else {
		cDes.DisableContainerLogging = des.DisableContainerLogging
	}
	if dcl.BoolCanonicalize(des.EnableAccessLogging, initial.EnableAccessLogging) || dcl.IsZeroValue(des.EnableAccessLogging) {
		cDes.EnableAccessLogging = initial.EnableAccessLogging
	} else {
		cDes.EnableAccessLogging = des.EnableAccessLogging
	}

	return cDes
}

func canonicalizeEndpointDeployedModelsSlice(des, initial []EndpointDeployedModels, opts ...dcl.ApplyOption) []EndpointDeployedModels {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModels, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModels(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModels, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModels(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModels(c *Client, des, nw *EndpointDeployedModels) *EndpointDeployedModels {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModels while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.DedicatedResources = canonicalizeNewEndpointDeployedModelsDedicatedResources(c, des.DedicatedResources, nw.DedicatedResources)
	nw.AutomaticResources = canonicalizeNewEndpointDeployedModelsAutomaticResources(c, des.AutomaticResources, nw.AutomaticResources)
	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.ModelVersionId, nw.ModelVersionId) {
		nw.ModelVersionId = des.ModelVersionId
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.BoolCanonicalize(des.DisableContainerLogging, nw.DisableContainerLogging) {
		nw.DisableContainerLogging = des.DisableContainerLogging
	}
	if dcl.BoolCanonicalize(des.EnableAccessLogging, nw.EnableAccessLogging) {
		nw.EnableAccessLogging = des.EnableAccessLogging
	}
	nw.PrivateEndpoints = canonicalizeNewEndpointDeployedModelsPrivateEndpoints(c, des.PrivateEndpoints, nw.PrivateEndpoints)

	return nw
}

func canonicalizeNewEndpointDeployedModelsSet(c *Client, des, nw []EndpointDeployedModels) []EndpointDeployedModels {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsSlice(c *Client, des, nw []EndpointDeployedModels) []EndpointDeployedModels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModels(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointDeployedModelsDedicatedResources(des, initial *EndpointDeployedModelsDedicatedResources, opts ...dcl.ApplyOption) *EndpointDeployedModelsDedicatedResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModelsDedicatedResources{}

	cDes.MachineSpec = canonicalizeEndpointDeployedModelsDedicatedResourcesMachineSpec(des.MachineSpec, initial.MachineSpec, opts...)
	if dcl.IsZeroValue(des.MinReplicaCount) || (dcl.IsEmptyValueIndirect(des.MinReplicaCount) && dcl.IsEmptyValueIndirect(initial.MinReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinReplicaCount = initial.MinReplicaCount
	} else {
		cDes.MinReplicaCount = des.MinReplicaCount
	}
	if dcl.IsZeroValue(des.MaxReplicaCount) || (dcl.IsEmptyValueIndirect(des.MaxReplicaCount) && dcl.IsEmptyValueIndirect(initial.MaxReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxReplicaCount = initial.MaxReplicaCount
	} else {
		cDes.MaxReplicaCount = des.MaxReplicaCount
	}
	cDes.AutoscalingMetricSpecs = canonicalizeEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(des.AutoscalingMetricSpecs, initial.AutoscalingMetricSpecs, opts...)

	return cDes
}

func canonicalizeEndpointDeployedModelsDedicatedResourcesSlice(des, initial []EndpointDeployedModelsDedicatedResources, opts ...dcl.ApplyOption) []EndpointDeployedModelsDedicatedResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModelsDedicatedResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModelsDedicatedResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModelsDedicatedResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModelsDedicatedResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModelsDedicatedResources(c *Client, des, nw *EndpointDeployedModelsDedicatedResources) *EndpointDeployedModelsDedicatedResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModelsDedicatedResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MachineSpec = canonicalizeNewEndpointDeployedModelsDedicatedResourcesMachineSpec(c, des.MachineSpec, nw.MachineSpec)
	nw.AutoscalingMetricSpecs = canonicalizeNewEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c, des.AutoscalingMetricSpecs, nw.AutoscalingMetricSpecs)

	return nw
}

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesSet(c *Client, des, nw []EndpointDeployedModelsDedicatedResources) []EndpointDeployedModelsDedicatedResources {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModelsDedicatedResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsDedicatedResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesSlice(c *Client, des, nw []EndpointDeployedModelsDedicatedResources) []EndpointDeployedModelsDedicatedResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModelsDedicatedResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModelsDedicatedResources(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointDeployedModelsDedicatedResourcesMachineSpec(des, initial *EndpointDeployedModelsDedicatedResourcesMachineSpec, opts ...dcl.ApplyOption) *EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModelsDedicatedResourcesMachineSpec{}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}
	if dcl.IsZeroValue(des.AcceleratorType) || (dcl.IsEmptyValueIndirect(des.AcceleratorType) && dcl.IsEmptyValueIndirect(initial.AcceleratorType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AcceleratorType = initial.AcceleratorType
	} else {
		cDes.AcceleratorType = des.AcceleratorType
	}
	if dcl.IsZeroValue(des.AcceleratorCount) || (dcl.IsEmptyValueIndirect(des.AcceleratorCount) && dcl.IsEmptyValueIndirect(initial.AcceleratorCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AcceleratorCount = initial.AcceleratorCount
	} else {
		cDes.AcceleratorCount = des.AcceleratorCount
	}

	return cDes
}

func canonicalizeEndpointDeployedModelsDedicatedResourcesMachineSpecSlice(des, initial []EndpointDeployedModelsDedicatedResourcesMachineSpec, opts ...dcl.ApplyOption) []EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModelsDedicatedResourcesMachineSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModelsDedicatedResourcesMachineSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModelsDedicatedResourcesMachineSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModelsDedicatedResourcesMachineSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesMachineSpec(c *Client, des, nw *EndpointDeployedModelsDedicatedResourcesMachineSpec) *EndpointDeployedModelsDedicatedResourcesMachineSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModelsDedicatedResourcesMachineSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}

	return nw
}

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesMachineSpecSet(c *Client, des, nw []EndpointDeployedModelsDedicatedResourcesMachineSpec) []EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModelsDedicatedResourcesMachineSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsDedicatedResourcesMachineSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesMachineSpecSlice(c *Client, des, nw []EndpointDeployedModelsDedicatedResourcesMachineSpec) []EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModelsDedicatedResourcesMachineSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModelsDedicatedResourcesMachineSpec(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(des, initial *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, opts ...dcl.ApplyOption) *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}

	if dcl.StringCanonicalize(des.MetricName, initial.MetricName) || dcl.IsZeroValue(des.MetricName) {
		cDes.MetricName = initial.MetricName
	} else {
		cDes.MetricName = des.MetricName
	}
	if dcl.IsZeroValue(des.Target) || (dcl.IsEmptyValueIndirect(des.Target) && dcl.IsEmptyValueIndirect(initial.Target)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Target = initial.Target
	} else {
		cDes.Target = des.Target
	}

	return cDes
}

func canonicalizeEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(des, initial []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, opts ...dcl.ApplyOption) []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c *Client, des, nw *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MetricName, nw.MetricName) {
		nw.MetricName = des.MetricName
	}

	return nw
}

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSet(c *Client, des, nw []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c *Client, des, nw []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointDeployedModelsAutomaticResources(des, initial *EndpointDeployedModelsAutomaticResources, opts ...dcl.ApplyOption) *EndpointDeployedModelsAutomaticResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModelsAutomaticResources{}

	if dcl.IsZeroValue(des.MinReplicaCount) || (dcl.IsEmptyValueIndirect(des.MinReplicaCount) && dcl.IsEmptyValueIndirect(initial.MinReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinReplicaCount = initial.MinReplicaCount
	} else {
		cDes.MinReplicaCount = des.MinReplicaCount
	}
	if dcl.IsZeroValue(des.MaxReplicaCount) || (dcl.IsEmptyValueIndirect(des.MaxReplicaCount) && dcl.IsEmptyValueIndirect(initial.MaxReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxReplicaCount = initial.MaxReplicaCount
	} else {
		cDes.MaxReplicaCount = des.MaxReplicaCount
	}

	return cDes
}

func canonicalizeEndpointDeployedModelsAutomaticResourcesSlice(des, initial []EndpointDeployedModelsAutomaticResources, opts ...dcl.ApplyOption) []EndpointDeployedModelsAutomaticResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModelsAutomaticResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModelsAutomaticResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModelsAutomaticResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModelsAutomaticResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModelsAutomaticResources(c *Client, des, nw *EndpointDeployedModelsAutomaticResources) *EndpointDeployedModelsAutomaticResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModelsAutomaticResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewEndpointDeployedModelsAutomaticResourcesSet(c *Client, des, nw []EndpointDeployedModelsAutomaticResources) []EndpointDeployedModelsAutomaticResources {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModelsAutomaticResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsAutomaticResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsAutomaticResourcesSlice(c *Client, des, nw []EndpointDeployedModelsAutomaticResources) []EndpointDeployedModelsAutomaticResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModelsAutomaticResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModelsAutomaticResources(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointDeployedModelsPrivateEndpoints(des, initial *EndpointDeployedModelsPrivateEndpoints, opts ...dcl.ApplyOption) *EndpointDeployedModelsPrivateEndpoints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointDeployedModelsPrivateEndpoints{}

	return cDes
}

func canonicalizeEndpointDeployedModelsPrivateEndpointsSlice(des, initial []EndpointDeployedModelsPrivateEndpoints, opts ...dcl.ApplyOption) []EndpointDeployedModelsPrivateEndpoints {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointDeployedModelsPrivateEndpoints, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointDeployedModelsPrivateEndpoints(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointDeployedModelsPrivateEndpoints, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointDeployedModelsPrivateEndpoints(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointDeployedModelsPrivateEndpoints(c *Client, des, nw *EndpointDeployedModelsPrivateEndpoints) *EndpointDeployedModelsPrivateEndpoints {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointDeployedModelsPrivateEndpoints while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.PredictHttpUri, nw.PredictHttpUri) {
		nw.PredictHttpUri = des.PredictHttpUri
	}
	if dcl.StringCanonicalize(des.ExplainHttpUri, nw.ExplainHttpUri) {
		nw.ExplainHttpUri = des.ExplainHttpUri
	}
	if dcl.StringCanonicalize(des.HealthHttpUri, nw.HealthHttpUri) {
		nw.HealthHttpUri = des.HealthHttpUri
	}
	if dcl.StringCanonicalize(des.ServiceAttachment, nw.ServiceAttachment) {
		nw.ServiceAttachment = des.ServiceAttachment
	}

	return nw
}

func canonicalizeNewEndpointDeployedModelsPrivateEndpointsSet(c *Client, des, nw []EndpointDeployedModelsPrivateEndpoints) []EndpointDeployedModelsPrivateEndpoints {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointDeployedModelsPrivateEndpoints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointDeployedModelsPrivateEndpointsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointDeployedModelsPrivateEndpointsSlice(c *Client, des, nw []EndpointDeployedModelsPrivateEndpoints) []EndpointDeployedModelsPrivateEndpoints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointDeployedModelsPrivateEndpoints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointDeployedModelsPrivateEndpoints(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointEncryptionSpec(des, initial *EndpointEncryptionSpec, opts ...dcl.ApplyOption) *EndpointEncryptionSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointEncryptionSpec{}

	if dcl.IsZeroValue(des.KmsKeyName) || (dcl.IsEmptyValueIndirect(des.KmsKeyName) && dcl.IsEmptyValueIndirect(initial.KmsKeyName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeEndpointEncryptionSpecSlice(des, initial []EndpointEncryptionSpec, opts ...dcl.ApplyOption) []EndpointEncryptionSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointEncryptionSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointEncryptionSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointEncryptionSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointEncryptionSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointEncryptionSpec(c *Client, des, nw *EndpointEncryptionSpec) *EndpointEncryptionSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointEncryptionSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewEndpointEncryptionSpecSet(c *Client, des, nw []EndpointEncryptionSpec) []EndpointEncryptionSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointEncryptionSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointEncryptionSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEndpointEncryptionSpecSlice(c *Client, des, nw []EndpointEncryptionSpec) []EndpointEncryptionSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointEncryptionSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointEncryptionSpec(c, &d, &n))
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
func diffEndpoint(c *Client, desired, actual *Endpoint, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointUpdateEndpointOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointUpdateEndpointOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployedModels, actual.DeployedModels, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareEndpointDeployedModelsNewStyle, EmptyObject: EmptyEndpointDeployedModels, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedModels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointUpdateEndpointOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionSpec, actual.EncryptionSpec, dcl.DiffInfo{ObjectFunction: compareEndpointEncryptionSpecNewStyle, EmptyObject: EmptyEndpointEncryptionSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ModelDeploymentMonitoringJob, actual.ModelDeploymentMonitoringJob, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ModelDeploymentMonitoringJob")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareEndpointDeployedModelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModels)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModels or *EndpointDeployedModels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModels)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DedicatedResources, actual.DedicatedResources, dcl.DiffInfo{ObjectFunction: compareEndpointDeployedModelsDedicatedResourcesNewStyle, EmptyObject: EmptyEndpointDeployedModelsDedicatedResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DedicatedResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutomaticResources, actual.AutomaticResources, dcl.DiffInfo{ObjectFunction: compareEndpointDeployedModelsAutomaticResourcesNewStyle, EmptyObject: EmptyEndpointDeployedModelsAutomaticResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutomaticResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Model, actual.Model, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Model")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ModelVersionId, actual.ModelVersionId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ModelVersionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableContainerLogging, actual.DisableContainerLogging, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisableContainerLogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableAccessLogging, actual.EnableAccessLogging, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableAccessLogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateEndpoints, actual.PrivateEndpoints, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareEndpointDeployedModelsPrivateEndpointsNewStyle, EmptyObject: EmptyEndpointDeployedModelsPrivateEndpoints, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateEndpoints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointDeployedModelsDedicatedResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModelsDedicatedResources)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModelsDedicatedResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResources or *EndpointDeployedModelsDedicatedResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModelsDedicatedResources)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModelsDedicatedResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineSpec, actual.MachineSpec, dcl.DiffInfo{ObjectFunction: compareEndpointDeployedModelsDedicatedResourcesMachineSpecNewStyle, EmptyObject: EmptyEndpointDeployedModelsDedicatedResourcesMachineSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinReplicaCount, actual.MinReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxReplicaCount, actual.MaxReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoscalingMetricSpecs, actual.AutoscalingMetricSpecs, dcl.DiffInfo{ObjectFunction: compareEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsNewStyle, EmptyObject: EmptyEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutoscalingMetricSpecs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointDeployedModelsDedicatedResourcesMachineSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModelsDedicatedResourcesMachineSpec)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModelsDedicatedResourcesMachineSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResourcesMachineSpec or *EndpointDeployedModelsDedicatedResourcesMachineSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModelsDedicatedResourcesMachineSpec)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModelsDedicatedResourcesMachineSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResourcesMachineSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs or *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MetricName, actual.MetricName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetricName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Target, actual.Target, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Target")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointDeployedModelsAutomaticResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModelsAutomaticResources)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModelsAutomaticResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsAutomaticResources or *EndpointDeployedModelsAutomaticResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModelsAutomaticResources)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModelsAutomaticResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsAutomaticResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinReplicaCount, actual.MinReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxReplicaCount, actual.MaxReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointDeployedModelsPrivateEndpointsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointDeployedModelsPrivateEndpoints)
	if !ok {
		desiredNotPointer, ok := d.(EndpointDeployedModelsPrivateEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsPrivateEndpoints or *EndpointDeployedModelsPrivateEndpoints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointDeployedModelsPrivateEndpoints)
	if !ok {
		actualNotPointer, ok := a.(EndpointDeployedModelsPrivateEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointDeployedModelsPrivateEndpoints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PredictHttpUri, actual.PredictHttpUri, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PredictHttpUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplainHttpUri, actual.ExplainHttpUri, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplainHttpUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthHttpUri, actual.HealthHttpUri, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HealthHttpUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAttachment, actual.ServiceAttachment, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAttachment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointEncryptionSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointEncryptionSpec)
	if !ok {
		desiredNotPointer, ok := d.(EndpointEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointEncryptionSpec or *EndpointEncryptionSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointEncryptionSpec)
	if !ok {
		actualNotPointer, ok := a.(EndpointEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointEncryptionSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
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
func (r *Endpoint) urlNormalized() *Endpoint {
	normalized := dcl.Copy(*r).(Endpoint)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.ModelDeploymentMonitoringJob = dcl.SelfLinkToName(r.ModelDeploymentMonitoringJob)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Endpoint) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEndpoint" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Endpoint resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Endpoint) marshal(c *Client) ([]byte, error) {
	m, err := expandEndpoint(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Endpoint: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEndpoint decodes JSON responses into the Endpoint resource schema.
func unmarshalEndpoint(b []byte, c *Client, res *Endpoint) (*Endpoint, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEndpoint(m, c, res)
}

func unmarshalMapEndpoint(m map[string]interface{}, c *Client, res *Endpoint) (*Endpoint, error) {

	flattened := flattenEndpoint(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEndpoint expands Endpoint into a JSON request object.
func expandEndpoint(c *Client, f *Endpoint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/endpoints/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandEndpointEncryptionSpec(c, f.EncryptionSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionSpec into encryptionSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionSpec"] = v
	}
	if v, err := dcl.ExpandProjectIDsToNumbers(c.Config, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenEndpoint flattens Endpoint from a JSON request object into the
// Endpoint type.
func flattenEndpoint(c *Client, i interface{}, res *Endpoint) *Endpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Endpoint{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.DeployedModels = flattenEndpointDeployedModelsSlice(c, m["deployedModels"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.EncryptionSpec = flattenEndpointEncryptionSpec(c, m["encryptionSpec"], res)
	resultRes.Network = dcl.FlattenProjectNumbersToIDs(c.Config, dcl.FlattenString(m["network"]))
	resultRes.ModelDeploymentMonitoringJob = dcl.FlattenString(m["modelDeploymentMonitoringJob"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandEndpointDeployedModelsMap expands the contents of EndpointDeployedModels into a JSON
// request object.
func expandEndpointDeployedModelsMap(c *Client, f map[string]EndpointDeployedModels, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModels(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsSlice expands the contents of EndpointDeployedModels into a JSON
// request object.
func expandEndpointDeployedModelsSlice(c *Client, f []EndpointDeployedModels, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModels(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsMap flattens the contents of EndpointDeployedModels from a JSON
// response object.
func flattenEndpointDeployedModelsMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModels{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModels{}
	}

	items := make(map[string]EndpointDeployedModels)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModels(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsSlice flattens the contents of EndpointDeployedModels from a JSON
// response object.
func flattenEndpointDeployedModelsSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModels {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModels{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModels{}
	}

	items := make([]EndpointDeployedModels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModels(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModels expands an instance of EndpointDeployedModels into a JSON
// request object.
func expandEndpointDeployedModels(c *Client, f *EndpointDeployedModels, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEndpointDeployedModelsDedicatedResources(c, f.DedicatedResources, res); err != nil {
		return nil, fmt.Errorf("error expanding DedicatedResources into dedicatedResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dedicatedResources"] = v
	}
	if v, err := expandEndpointDeployedModelsAutomaticResources(c, f.AutomaticResources, res); err != nil {
		return nil, fmt.Errorf("error expanding AutomaticResources into automaticResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["automaticResources"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Model; !dcl.IsEmptyValueIndirect(v) {
		m["model"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.DisableContainerLogging; !dcl.IsEmptyValueIndirect(v) {
		m["disableContainerLogging"] = v
	}
	if v := f.EnableAccessLogging; !dcl.IsEmptyValueIndirect(v) {
		m["enableAccessLogging"] = v
	}

	return m, nil
}

// flattenEndpointDeployedModels flattens an instance of EndpointDeployedModels from a JSON
// response object.
func flattenEndpointDeployedModels(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModels
	}
	r.DedicatedResources = flattenEndpointDeployedModelsDedicatedResources(c, m["dedicatedResources"], res)
	r.AutomaticResources = flattenEndpointDeployedModelsAutomaticResources(c, m["automaticResources"], res)
	r.Id = dcl.FlattenString(m["id"])
	r.Model = dcl.FlattenString(m["model"])
	r.ModelVersionId = dcl.FlattenString(m["modelVersionId"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.DisableContainerLogging = dcl.FlattenBool(m["disableContainerLogging"])
	r.EnableAccessLogging = dcl.FlattenBool(m["enableAccessLogging"])
	r.PrivateEndpoints = flattenEndpointDeployedModelsPrivateEndpoints(c, m["privateEndpoints"], res)

	return r
}

// expandEndpointDeployedModelsDedicatedResourcesMap expands the contents of EndpointDeployedModelsDedicatedResources into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesMap(c *Client, f map[string]EndpointDeployedModelsDedicatedResources, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsDedicatedResourcesSlice expands the contents of EndpointDeployedModelsDedicatedResources into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesSlice(c *Client, f []EndpointDeployedModelsDedicatedResources, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsDedicatedResourcesMap flattens the contents of EndpointDeployedModelsDedicatedResources from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsDedicatedResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsDedicatedResources{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsDedicatedResources{}
	}

	items := make(map[string]EndpointDeployedModelsDedicatedResources)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsDedicatedResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsDedicatedResourcesSlice flattens the contents of EndpointDeployedModelsDedicatedResources from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsDedicatedResources {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsDedicatedResources{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsDedicatedResources{}
	}

	items := make([]EndpointDeployedModelsDedicatedResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsDedicatedResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModelsDedicatedResources expands an instance of EndpointDeployedModelsDedicatedResources into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResources(c *Client, f *EndpointDeployedModelsDedicatedResources, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEndpointDeployedModelsDedicatedResourcesMachineSpec(c, f.MachineSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding MachineSpec into machineSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["machineSpec"] = v
	}
	if v := f.MinReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["minReplicaCount"] = v
	}
	if v := f.MaxReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxReplicaCount"] = v
	}
	if v, err := expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c, f.AutoscalingMetricSpecs, res); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingMetricSpecs into autoscalingMetricSpecs: %w", err)
	} else if v != nil {
		m["autoscalingMetricSpecs"] = v
	}

	return m, nil
}

// flattenEndpointDeployedModelsDedicatedResources flattens an instance of EndpointDeployedModelsDedicatedResources from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResources(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModelsDedicatedResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModelsDedicatedResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModelsDedicatedResources
	}
	r.MachineSpec = flattenEndpointDeployedModelsDedicatedResourcesMachineSpec(c, m["machineSpec"], res)
	r.MinReplicaCount = dcl.FlattenInteger(m["minReplicaCount"])
	r.MaxReplicaCount = dcl.FlattenInteger(m["maxReplicaCount"])
	r.AutoscalingMetricSpecs = flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c, m["autoscalingMetricSpecs"], res)

	return r
}

// expandEndpointDeployedModelsDedicatedResourcesMachineSpecMap expands the contents of EndpointDeployedModelsDedicatedResourcesMachineSpec into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesMachineSpecMap(c *Client, f map[string]EndpointDeployedModelsDedicatedResourcesMachineSpec, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResourcesMachineSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsDedicatedResourcesMachineSpecSlice expands the contents of EndpointDeployedModelsDedicatedResourcesMachineSpec into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesMachineSpecSlice(c *Client, f []EndpointDeployedModelsDedicatedResourcesMachineSpec, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResourcesMachineSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpecMap flattens the contents of EndpointDeployedModelsDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpecMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsDedicatedResourcesMachineSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}

	items := make(map[string]EndpointDeployedModelsDedicatedResourcesMachineSpec)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsDedicatedResourcesMachineSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpecSlice flattens the contents of EndpointDeployedModelsDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpecSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsDedicatedResourcesMachineSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}

	items := make([]EndpointDeployedModelsDedicatedResourcesMachineSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsDedicatedResourcesMachineSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModelsDedicatedResourcesMachineSpec expands an instance of EndpointDeployedModelsDedicatedResourcesMachineSpec into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesMachineSpec(c *Client, f *EndpointDeployedModelsDedicatedResourcesMachineSpec, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}

	return m, nil
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpec flattens an instance of EndpointDeployedModelsDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpec(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModelsDedicatedResourcesMachineSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModelsDedicatedResourcesMachineSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModelsDedicatedResourcesMachineSpec
	}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.AcceleratorType = flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(m["acceleratorType"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsMap expands the contents of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsMap(c *Client, f map[string]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice expands the contents of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c *Client, f []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsMap flattens the contents of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	}

	items := make(map[string]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice flattens the contents of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	}

	items := make([]EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs expands an instance of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs into a JSON
// request object.
func expandEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c *Client, f *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MetricName; !dcl.IsEmptyValueIndirect(v) {
		m["metricName"] = v
	}
	if v := f.Target; !dcl.IsEmptyValueIndirect(v) {
		m["target"] = v
	}

	return m, nil
}

// flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs flattens an instance of EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
	}
	r.MetricName = dcl.FlattenString(m["metricName"])
	r.Target = dcl.FlattenInteger(m["target"])

	return r
}

// expandEndpointDeployedModelsAutomaticResourcesMap expands the contents of EndpointDeployedModelsAutomaticResources into a JSON
// request object.
func expandEndpointDeployedModelsAutomaticResourcesMap(c *Client, f map[string]EndpointDeployedModelsAutomaticResources, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModelsAutomaticResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsAutomaticResourcesSlice expands the contents of EndpointDeployedModelsAutomaticResources into a JSON
// request object.
func expandEndpointDeployedModelsAutomaticResourcesSlice(c *Client, f []EndpointDeployedModelsAutomaticResources, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModelsAutomaticResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsAutomaticResourcesMap flattens the contents of EndpointDeployedModelsAutomaticResources from a JSON
// response object.
func flattenEndpointDeployedModelsAutomaticResourcesMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsAutomaticResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsAutomaticResources{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsAutomaticResources{}
	}

	items := make(map[string]EndpointDeployedModelsAutomaticResources)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsAutomaticResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsAutomaticResourcesSlice flattens the contents of EndpointDeployedModelsAutomaticResources from a JSON
// response object.
func flattenEndpointDeployedModelsAutomaticResourcesSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsAutomaticResources {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsAutomaticResources{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsAutomaticResources{}
	}

	items := make([]EndpointDeployedModelsAutomaticResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsAutomaticResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModelsAutomaticResources expands an instance of EndpointDeployedModelsAutomaticResources into a JSON
// request object.
func expandEndpointDeployedModelsAutomaticResources(c *Client, f *EndpointDeployedModelsAutomaticResources, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["minReplicaCount"] = v
	}
	if v := f.MaxReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxReplicaCount"] = v
	}

	return m, nil
}

// flattenEndpointDeployedModelsAutomaticResources flattens an instance of EndpointDeployedModelsAutomaticResources from a JSON
// response object.
func flattenEndpointDeployedModelsAutomaticResources(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModelsAutomaticResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModelsAutomaticResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModelsAutomaticResources
	}
	r.MinReplicaCount = dcl.FlattenInteger(m["minReplicaCount"])
	r.MaxReplicaCount = dcl.FlattenInteger(m["maxReplicaCount"])

	return r
}

// expandEndpointDeployedModelsPrivateEndpointsMap expands the contents of EndpointDeployedModelsPrivateEndpoints into a JSON
// request object.
func expandEndpointDeployedModelsPrivateEndpointsMap(c *Client, f map[string]EndpointDeployedModelsPrivateEndpoints, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointDeployedModelsPrivateEndpoints(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointDeployedModelsPrivateEndpointsSlice expands the contents of EndpointDeployedModelsPrivateEndpoints into a JSON
// request object.
func expandEndpointDeployedModelsPrivateEndpointsSlice(c *Client, f []EndpointDeployedModelsPrivateEndpoints, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointDeployedModelsPrivateEndpoints(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointDeployedModelsPrivateEndpointsMap flattens the contents of EndpointDeployedModelsPrivateEndpoints from a JSON
// response object.
func flattenEndpointDeployedModelsPrivateEndpointsMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsPrivateEndpoints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsPrivateEndpoints{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsPrivateEndpoints{}
	}

	items := make(map[string]EndpointDeployedModelsPrivateEndpoints)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsPrivateEndpoints(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointDeployedModelsPrivateEndpointsSlice flattens the contents of EndpointDeployedModelsPrivateEndpoints from a JSON
// response object.
func flattenEndpointDeployedModelsPrivateEndpointsSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsPrivateEndpoints {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsPrivateEndpoints{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsPrivateEndpoints{}
	}

	items := make([]EndpointDeployedModelsPrivateEndpoints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsPrivateEndpoints(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointDeployedModelsPrivateEndpoints expands an instance of EndpointDeployedModelsPrivateEndpoints into a JSON
// request object.
func expandEndpointDeployedModelsPrivateEndpoints(c *Client, f *EndpointDeployedModelsPrivateEndpoints, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenEndpointDeployedModelsPrivateEndpoints flattens an instance of EndpointDeployedModelsPrivateEndpoints from a JSON
// response object.
func flattenEndpointDeployedModelsPrivateEndpoints(c *Client, i interface{}, res *Endpoint) *EndpointDeployedModelsPrivateEndpoints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointDeployedModelsPrivateEndpoints{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointDeployedModelsPrivateEndpoints
	}
	r.PredictHttpUri = dcl.FlattenString(m["predictHttpUri"])
	r.ExplainHttpUri = dcl.FlattenString(m["explainHttpUri"])
	r.HealthHttpUri = dcl.FlattenString(m["healthHttpUri"])
	r.ServiceAttachment = dcl.FlattenString(m["serviceAttachment"])

	return r
}

// expandEndpointEncryptionSpecMap expands the contents of EndpointEncryptionSpec into a JSON
// request object.
func expandEndpointEncryptionSpecMap(c *Client, f map[string]EndpointEncryptionSpec, res *Endpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointEncryptionSpecSlice expands the contents of EndpointEncryptionSpec into a JSON
// request object.
func expandEndpointEncryptionSpecSlice(c *Client, f []EndpointEncryptionSpec, res *Endpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointEncryptionSpecMap flattens the contents of EndpointEncryptionSpec from a JSON
// response object.
func flattenEndpointEncryptionSpecMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointEncryptionSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointEncryptionSpec{}
	}

	if len(a) == 0 {
		return map[string]EndpointEncryptionSpec{}
	}

	items := make(map[string]EndpointEncryptionSpec)
	for k, item := range a {
		items[k] = *flattenEndpointEncryptionSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointEncryptionSpecSlice flattens the contents of EndpointEncryptionSpec from a JSON
// response object.
func flattenEndpointEncryptionSpecSlice(c *Client, i interface{}, res *Endpoint) []EndpointEncryptionSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointEncryptionSpec{}
	}

	if len(a) == 0 {
		return []EndpointEncryptionSpec{}
	}

	items := make([]EndpointEncryptionSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointEncryptionSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointEncryptionSpec expands an instance of EndpointEncryptionSpec into a JSON
// request object.
func expandEndpointEncryptionSpec(c *Client, f *EndpointEncryptionSpec, res *Endpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenEndpointEncryptionSpec flattens an instance of EndpointEncryptionSpec from a JSON
// response object.
func flattenEndpointEncryptionSpec(c *Client, i interface{}, res *Endpoint) *EndpointEncryptionSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointEncryptionSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointEncryptionSpec
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumMap flattens the contents of EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumMap(c *Client, i interface{}, res *Endpoint) map[string]EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum{}
	}

	items := make(map[string]EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum)
	for k, item := range a {
		items[k] = *flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(item.(interface{}))
	}

	return items
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumSlice flattens the contents of EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum from a JSON
// response object.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumSlice(c *Client, i interface{}, res *Endpoint) []EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum{}
	}

	if len(a) == 0 {
		return []EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum{}
	}

	items := make([]EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(item.(interface{})))
	}

	return items
}

// flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum asserts that an interface is a string, and returns a
// pointer to a *EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum with the same value as that string.
func flattenEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(i interface{}) *EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Endpoint) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEndpoint(b, c, r)
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

type endpointDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         endpointApiOperation
}

func convertFieldDiffsToEndpointDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]endpointDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []endpointDiff
	// For each operation name, create a endpointDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := endpointDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEndpointApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEndpointApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (endpointApiOperation, error) {
	switch opName {

	case "updateEndpointUpdateEndpointOperation":
		return &updateEndpointUpdateEndpointOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEndpointFields(r *Endpoint) error {
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &EndpointEncryptionSpec{}
	}
	if err := extractEndpointEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	return nil
}
func extractEndpointDeployedModelsFields(r *Endpoint, o *EndpointDeployedModels) error {
	vDedicatedResources := o.DedicatedResources
	if vDedicatedResources == nil {
		// note: explicitly not the empty object.
		vDedicatedResources = &EndpointDeployedModelsDedicatedResources{}
	}
	if err := extractEndpointDeployedModelsDedicatedResourcesFields(r, vDedicatedResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDedicatedResources) {
		o.DedicatedResources = vDedicatedResources
	}
	vAutomaticResources := o.AutomaticResources
	if vAutomaticResources == nil {
		// note: explicitly not the empty object.
		vAutomaticResources = &EndpointDeployedModelsAutomaticResources{}
	}
	if err := extractEndpointDeployedModelsAutomaticResourcesFields(r, vAutomaticResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAutomaticResources) {
		o.AutomaticResources = vAutomaticResources
	}
	vPrivateEndpoints := o.PrivateEndpoints
	if vPrivateEndpoints == nil {
		// note: explicitly not the empty object.
		vPrivateEndpoints = &EndpointDeployedModelsPrivateEndpoints{}
	}
	if err := extractEndpointDeployedModelsPrivateEndpointsFields(r, vPrivateEndpoints); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPrivateEndpoints) {
		o.PrivateEndpoints = vPrivateEndpoints
	}
	return nil
}
func extractEndpointDeployedModelsDedicatedResourcesFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResources) error {
	vMachineSpec := o.MachineSpec
	if vMachineSpec == nil {
		// note: explicitly not the empty object.
		vMachineSpec = &EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}
	if err := extractEndpointDeployedModelsDedicatedResourcesMachineSpecFields(r, vMachineSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMachineSpec) {
		o.MachineSpec = vMachineSpec
	}
	return nil
}
func extractEndpointDeployedModelsDedicatedResourcesMachineSpecFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResourcesMachineSpec) error {
	return nil
}
func extractEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) error {
	return nil
}
func extractEndpointDeployedModelsAutomaticResourcesFields(r *Endpoint, o *EndpointDeployedModelsAutomaticResources) error {
	return nil
}
func extractEndpointDeployedModelsPrivateEndpointsFields(r *Endpoint, o *EndpointDeployedModelsPrivateEndpoints) error {
	return nil
}
func extractEndpointEncryptionSpecFields(r *Endpoint, o *EndpointEncryptionSpec) error {
	return nil
}

func postReadExtractEndpointFields(r *Endpoint) error {
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &EndpointEncryptionSpec{}
	}
	if err := postReadExtractEndpointEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	return nil
}
func postReadExtractEndpointDeployedModelsFields(r *Endpoint, o *EndpointDeployedModels) error {
	vDedicatedResources := o.DedicatedResources
	if vDedicatedResources == nil {
		// note: explicitly not the empty object.
		vDedicatedResources = &EndpointDeployedModelsDedicatedResources{}
	}
	if err := extractEndpointDeployedModelsDedicatedResourcesFields(r, vDedicatedResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDedicatedResources) {
		o.DedicatedResources = vDedicatedResources
	}
	vAutomaticResources := o.AutomaticResources
	if vAutomaticResources == nil {
		// note: explicitly not the empty object.
		vAutomaticResources = &EndpointDeployedModelsAutomaticResources{}
	}
	if err := extractEndpointDeployedModelsAutomaticResourcesFields(r, vAutomaticResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAutomaticResources) {
		o.AutomaticResources = vAutomaticResources
	}
	vPrivateEndpoints := o.PrivateEndpoints
	if vPrivateEndpoints == nil {
		// note: explicitly not the empty object.
		vPrivateEndpoints = &EndpointDeployedModelsPrivateEndpoints{}
	}
	if err := extractEndpointDeployedModelsPrivateEndpointsFields(r, vPrivateEndpoints); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPrivateEndpoints) {
		o.PrivateEndpoints = vPrivateEndpoints
	}
	return nil
}
func postReadExtractEndpointDeployedModelsDedicatedResourcesFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResources) error {
	vMachineSpec := o.MachineSpec
	if vMachineSpec == nil {
		// note: explicitly not the empty object.
		vMachineSpec = &EndpointDeployedModelsDedicatedResourcesMachineSpec{}
	}
	if err := extractEndpointDeployedModelsDedicatedResourcesMachineSpecFields(r, vMachineSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMachineSpec) {
		o.MachineSpec = vMachineSpec
	}
	return nil
}
func postReadExtractEndpointDeployedModelsDedicatedResourcesMachineSpecFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResourcesMachineSpec) error {
	return nil
}
func postReadExtractEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsFields(r *Endpoint, o *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) error {
	return nil
}
func postReadExtractEndpointDeployedModelsAutomaticResourcesFields(r *Endpoint, o *EndpointDeployedModelsAutomaticResources) error {
	return nil
}
func postReadExtractEndpointDeployedModelsPrivateEndpointsFields(r *Endpoint, o *EndpointDeployedModelsPrivateEndpoints) error {
	return nil
}
func postReadExtractEndpointEncryptionSpecFields(r *Endpoint, o *EndpointEncryptionSpec) error {
	return nil
}
