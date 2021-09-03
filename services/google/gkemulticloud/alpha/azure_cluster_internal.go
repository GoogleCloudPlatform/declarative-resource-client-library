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

func (r *AzureCluster) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "azureRegion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resourceGroupId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "azureClient"); err != nil {
		return err
	}
	if err := dcl.Required(r, "networking"); err != nil {
		return err
	}
	if err := dcl.Required(r, "controlPlane"); err != nil {
		return err
	}
	if err := dcl.Required(r, "authorization"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Networking) {
		if err := r.Networking.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ControlPlane) {
		if err := r.ControlPlane.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Authorization) {
		if err := r.Authorization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WorkloadIdentityConfig) {
		if err := r.WorkloadIdentityConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AzureClusterNetworking) validate() error {
	if err := dcl.Required(r, "virtualNetworkId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "podAddressCidrBlocks"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceAddressCidrBlocks"); err != nil {
		return err
	}
	return nil
}
func (r *AzureClusterControlPlane) validate() error {
	if err := dcl.Required(r, "version"); err != nil {
		return err
	}
	if err := dcl.Required(r, "subnetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sshConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SshConfig) {
		if err := r.SshConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RootVolume) {
		if err := r.RootVolume.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MainVolume) {
		if err := r.MainVolume.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DatabaseEncryption) {
		if err := r.DatabaseEncryption.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AzureClusterControlPlaneSshConfig) validate() error {
	if err := dcl.Required(r, "authorizedKey"); err != nil {
		return err
	}
	return nil
}
func (r *AzureClusterControlPlaneRootVolume) validate() error {
	return nil
}
func (r *AzureClusterControlPlaneMainVolume) validate() error {
	return nil
}
func (r *AzureClusterControlPlaneDatabaseEncryption) validate() error {
	if err := dcl.Required(r, "resourceGroupId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "kmsKeyIdentifier"); err != nil {
		return err
	}
	return nil
}
func (r *AzureClusterAuthorization) validate() error {
	if err := dcl.Required(r, "adminUsers"); err != nil {
		return err
	}
	return nil
}
func (r *AzureClusterAuthorizationAdminUsers) validate() error {
	if err := dcl.Required(r, "username"); err != nil {
		return err
	}
	return nil
}
func (r *AzureClusterWorkloadIdentityConfig) validate() error {
	return nil
}
func (r *AzureCluster) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-gkemulticloud.googleapis.com/v1", params)
}

func (r *AzureCluster) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AzureCluster) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters", nr.basePath(), userBasePath, params), nil

}

func (r *AzureCluster) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters?azureClusterId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AzureCluster) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

// azureClusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type azureClusterApiOperation interface {
	do(context.Context, *AzureCluster, *Client) error
}

func (c *Client) listAzureClusterRaw(ctx context.Context, r *AzureCluster, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AzureClusterMaxPage {
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

type listAzureClusterOperation struct {
	AzureClusters []map[string]interface{} `json:"azureClusters"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listAzureCluster(ctx context.Context, r *AzureCluster, pageToken string, pageSize int32) ([]*AzureCluster, string, error) {
	b, err := c.listAzureClusterRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAzureClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AzureCluster
	for _, v := range m.AzureClusters {
		res, err := unmarshalMapAzureCluster(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAzureCluster(ctx context.Context, f func(*AzureCluster) bool, resources []*AzureCluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAzureCluster(ctx, res)
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

type deleteAzureClusterOperation struct{}

func (op *deleteAzureClusterOperation) do(ctx context.Context, r *AzureCluster, c *Client) error {
	r, err := c.GetAzureCluster(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "AzureCluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetAzureCluster checking for existence. error: %v", err)
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

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAzureCluster(ctx, r)
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
type createAzureClusterOperation struct {
	response map[string]interface{}
}

func (op *createAzureClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAzureClusterOperation) do(ctx context.Context, r *AzureCluster, c *Client) error {
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

	if _, err := c.GetAzureCluster(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAzureClusterRaw(ctx context.Context, r *AzureCluster) ([]byte, error) {

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

func (c *Client) azureClusterDiffsForRawDesired(ctx context.Context, rawDesired *AzureCluster, opts ...dcl.ApplyOption) (initial, desired *AzureCluster, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AzureCluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AzureCluster); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected AzureCluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAzureCluster(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a AzureCluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AzureCluster resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that AzureCluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAzureClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for AzureCluster: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for AzureCluster: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAzureClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for AzureCluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAzureClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for AzureCluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAzureCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAzureClusterInitialState(rawInitial, rawDesired *AzureCluster) (*AzureCluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAzureClusterDesiredState(rawDesired, rawInitial *AzureCluster, opts ...dcl.ApplyOption) (*AzureCluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Networking = canonicalizeAzureClusterNetworking(rawDesired.Networking, nil, opts...)
		rawDesired.ControlPlane = canonicalizeAzureClusterControlPlane(rawDesired.ControlPlane, nil, opts...)
		rawDesired.Authorization = canonicalizeAzureClusterAuthorization(rawDesired.Authorization, nil, opts...)
		rawDesired.WorkloadIdentityConfig = canonicalizeAzureClusterWorkloadIdentityConfig(rawDesired.WorkloadIdentityConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &AzureCluster{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.AzureRegion, rawInitial.AzureRegion) {
		canonicalDesired.AzureRegion = rawInitial.AzureRegion
	} else {
		canonicalDesired.AzureRegion = rawDesired.AzureRegion
	}
	if dcl.StringCanonicalize(rawDesired.ResourceGroupId, rawInitial.ResourceGroupId) {
		canonicalDesired.ResourceGroupId = rawInitial.ResourceGroupId
	} else {
		canonicalDesired.ResourceGroupId = rawDesired.ResourceGroupId
	}
	if dcl.NameToSelfLink(rawDesired.AzureClient, rawInitial.AzureClient) {
		canonicalDesired.AzureClient = rawInitial.AzureClient
	} else {
		canonicalDesired.AzureClient = rawDesired.AzureClient
	}
	canonicalDesired.Networking = canonicalizeAzureClusterNetworking(rawDesired.Networking, rawInitial.Networking, opts...)
	canonicalDesired.ControlPlane = canonicalizeAzureClusterControlPlane(rawDesired.ControlPlane, rawInitial.ControlPlane, opts...)
	canonicalDesired.Authorization = canonicalizeAzureClusterAuthorization(rawDesired.Authorization, rawInitial.Authorization, opts...)
	if dcl.IsZeroValue(rawDesired.Annotations) {
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
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

func canonicalizeAzureClusterNewState(c *Client, rawNew, rawDesired *AzureCluster) (*AzureCluster, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.AzureRegion) && dcl.IsNotReturnedByServer(rawDesired.AzureRegion) {
		rawNew.AzureRegion = rawDesired.AzureRegion
	} else {
		if dcl.StringCanonicalize(rawDesired.AzureRegion, rawNew.AzureRegion) {
			rawNew.AzureRegion = rawDesired.AzureRegion
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.ResourceGroupId) && dcl.IsNotReturnedByServer(rawDesired.ResourceGroupId) {
		rawNew.ResourceGroupId = rawDesired.ResourceGroupId
	} else {
		if dcl.StringCanonicalize(rawDesired.ResourceGroupId, rawNew.ResourceGroupId) {
			rawNew.ResourceGroupId = rawDesired.ResourceGroupId
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.AzureClient) && dcl.IsNotReturnedByServer(rawDesired.AzureClient) {
		rawNew.AzureClient = rawDesired.AzureClient
	} else {
		if dcl.NameToSelfLink(rawDesired.AzureClient, rawNew.AzureClient) {
			rawNew.AzureClient = rawDesired.AzureClient
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Networking) && dcl.IsNotReturnedByServer(rawDesired.Networking) {
		rawNew.Networking = rawDesired.Networking
	} else {
		rawNew.Networking = canonicalizeNewAzureClusterNetworking(c, rawDesired.Networking, rawNew.Networking)
	}

	if dcl.IsNotReturnedByServer(rawNew.ControlPlane) && dcl.IsNotReturnedByServer(rawDesired.ControlPlane) {
		rawNew.ControlPlane = rawDesired.ControlPlane
	} else {
		rawNew.ControlPlane = canonicalizeNewAzureClusterControlPlane(c, rawDesired.ControlPlane, rawNew.ControlPlane)
	}

	if dcl.IsNotReturnedByServer(rawNew.Authorization) && dcl.IsNotReturnedByServer(rawDesired.Authorization) {
		rawNew.Authorization = rawDesired.Authorization
	} else {
		rawNew.Authorization = canonicalizeNewAzureClusterAuthorization(c, rawDesired.Authorization, rawNew.Authorization)
	}

	if dcl.IsNotReturnedByServer(rawNew.State) && dcl.IsNotReturnedByServer(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Endpoint) && dcl.IsNotReturnedByServer(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.Endpoint, rawNew.Endpoint) {
			rawNew.Endpoint = rawDesired.Endpoint
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Uid) && dcl.IsNotReturnedByServer(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Reconciling) && dcl.IsNotReturnedByServer(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Annotations) && dcl.IsNotReturnedByServer(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.WorkloadIdentityConfig) && dcl.IsNotReturnedByServer(rawDesired.WorkloadIdentityConfig) {
		rawNew.WorkloadIdentityConfig = rawDesired.WorkloadIdentityConfig
	} else {
		rawNew.WorkloadIdentityConfig = canonicalizeNewAzureClusterWorkloadIdentityConfig(c, rawDesired.WorkloadIdentityConfig, rawNew.WorkloadIdentityConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAzureClusterNetworking(des, initial *AzureClusterNetworking, opts ...dcl.ApplyOption) *AzureClusterNetworking {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterNetworking{}

	if dcl.StringCanonicalize(des.VirtualNetworkId, initial.VirtualNetworkId) || dcl.IsZeroValue(des.VirtualNetworkId) {
		cDes.VirtualNetworkId = initial.VirtualNetworkId
	} else {
		cDes.VirtualNetworkId = des.VirtualNetworkId
	}
	if dcl.IsZeroValue(des.PodAddressCidrBlocks) {
		des.PodAddressCidrBlocks = initial.PodAddressCidrBlocks
	} else {
		cDes.PodAddressCidrBlocks = des.PodAddressCidrBlocks
	}
	if dcl.IsZeroValue(des.ServiceAddressCidrBlocks) {
		des.ServiceAddressCidrBlocks = initial.ServiceAddressCidrBlocks
	} else {
		cDes.ServiceAddressCidrBlocks = des.ServiceAddressCidrBlocks
	}

	return cDes
}

func canonicalizeAzureClusterNetworkingSlice(des, initial []AzureClusterNetworking, opts ...dcl.ApplyOption) []AzureClusterNetworking {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterNetworking, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterNetworking(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterNetworking, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterNetworking(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterNetworking(c *Client, des, nw *AzureClusterNetworking) *AzureClusterNetworking {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterNetworking while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.VirtualNetworkId, nw.VirtualNetworkId) {
		nw.VirtualNetworkId = des.VirtualNetworkId
	}

	return nw
}

func canonicalizeNewAzureClusterNetworkingSet(c *Client, des, nw []AzureClusterNetworking) []AzureClusterNetworking {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterNetworking
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterNetworkingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterNetworkingSlice(c *Client, des, nw []AzureClusterNetworking) []AzureClusterNetworking {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterNetworking
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterNetworking(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterControlPlane(des, initial *AzureClusterControlPlane, opts ...dcl.ApplyOption) *AzureClusterControlPlane {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterControlPlane{}

	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.StringCanonicalize(des.SubnetId, initial.SubnetId) || dcl.IsZeroValue(des.SubnetId) {
		cDes.SubnetId = initial.SubnetId
	} else {
		cDes.SubnetId = des.SubnetId
	}
	if dcl.StringCanonicalize(des.VmSize, initial.VmSize) || dcl.IsZeroValue(des.VmSize) {
		cDes.VmSize = initial.VmSize
	} else {
		cDes.VmSize = des.VmSize
	}
	cDes.SshConfig = canonicalizeAzureClusterControlPlaneSshConfig(des.SshConfig, initial.SshConfig, opts...)
	cDes.RootVolume = canonicalizeAzureClusterControlPlaneRootVolume(des.RootVolume, initial.RootVolume, opts...)
	cDes.MainVolume = canonicalizeAzureClusterControlPlaneMainVolume(des.MainVolume, initial.MainVolume, opts...)
	cDes.DatabaseEncryption = canonicalizeAzureClusterControlPlaneDatabaseEncryption(des.DatabaseEncryption, initial.DatabaseEncryption, opts...)
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}

	return cDes
}

func canonicalizeAzureClusterControlPlaneSlice(des, initial []AzureClusterControlPlane, opts ...dcl.ApplyOption) []AzureClusterControlPlane {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterControlPlane, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterControlPlane(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterControlPlane, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterControlPlane(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterControlPlane(c *Client, des, nw *AzureClusterControlPlane) *AzureClusterControlPlane {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterControlPlane while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.SubnetId, nw.SubnetId) {
		nw.SubnetId = des.SubnetId
	}
	if dcl.StringCanonicalize(des.VmSize, nw.VmSize) {
		nw.VmSize = des.VmSize
	}
	nw.SshConfig = canonicalizeNewAzureClusterControlPlaneSshConfig(c, des.SshConfig, nw.SshConfig)
	nw.RootVolume = canonicalizeNewAzureClusterControlPlaneRootVolume(c, des.RootVolume, nw.RootVolume)
	nw.MainVolume = canonicalizeNewAzureClusterControlPlaneMainVolume(c, des.MainVolume, nw.MainVolume)
	nw.DatabaseEncryption = canonicalizeNewAzureClusterControlPlaneDatabaseEncryption(c, des.DatabaseEncryption, nw.DatabaseEncryption)

	return nw
}

func canonicalizeNewAzureClusterControlPlaneSet(c *Client, des, nw []AzureClusterControlPlane) []AzureClusterControlPlane {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterControlPlane
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterControlPlaneNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterControlPlaneSlice(c *Client, des, nw []AzureClusterControlPlane) []AzureClusterControlPlane {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterControlPlane
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterControlPlane(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterControlPlaneSshConfig(des, initial *AzureClusterControlPlaneSshConfig, opts ...dcl.ApplyOption) *AzureClusterControlPlaneSshConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterControlPlaneSshConfig{}

	if dcl.StringCanonicalize(des.AuthorizedKey, initial.AuthorizedKey) || dcl.IsZeroValue(des.AuthorizedKey) {
		cDes.AuthorizedKey = initial.AuthorizedKey
	} else {
		cDes.AuthorizedKey = des.AuthorizedKey
	}

	return cDes
}

func canonicalizeAzureClusterControlPlaneSshConfigSlice(des, initial []AzureClusterControlPlaneSshConfig, opts ...dcl.ApplyOption) []AzureClusterControlPlaneSshConfig {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterControlPlaneSshConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterControlPlaneSshConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterControlPlaneSshConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterControlPlaneSshConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterControlPlaneSshConfig(c *Client, des, nw *AzureClusterControlPlaneSshConfig) *AzureClusterControlPlaneSshConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterControlPlaneSshConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AuthorizedKey, nw.AuthorizedKey) {
		nw.AuthorizedKey = des.AuthorizedKey
	}

	return nw
}

func canonicalizeNewAzureClusterControlPlaneSshConfigSet(c *Client, des, nw []AzureClusterControlPlaneSshConfig) []AzureClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterControlPlaneSshConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterControlPlaneSshConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterControlPlaneSshConfigSlice(c *Client, des, nw []AzureClusterControlPlaneSshConfig) []AzureClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterControlPlaneSshConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterControlPlaneSshConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterControlPlaneRootVolume(des, initial *AzureClusterControlPlaneRootVolume, opts ...dcl.ApplyOption) *AzureClusterControlPlaneRootVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterControlPlaneRootVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}

	return cDes
}

func canonicalizeAzureClusterControlPlaneRootVolumeSlice(des, initial []AzureClusterControlPlaneRootVolume, opts ...dcl.ApplyOption) []AzureClusterControlPlaneRootVolume {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterControlPlaneRootVolume, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterControlPlaneRootVolume(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterControlPlaneRootVolume, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterControlPlaneRootVolume(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterControlPlaneRootVolume(c *Client, des, nw *AzureClusterControlPlaneRootVolume) *AzureClusterControlPlaneRootVolume {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterControlPlaneRootVolume while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewAzureClusterControlPlaneRootVolumeSet(c *Client, des, nw []AzureClusterControlPlaneRootVolume) []AzureClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterControlPlaneRootVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterControlPlaneRootVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterControlPlaneRootVolumeSlice(c *Client, des, nw []AzureClusterControlPlaneRootVolume) []AzureClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterControlPlaneRootVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterControlPlaneRootVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterControlPlaneMainVolume(des, initial *AzureClusterControlPlaneMainVolume, opts ...dcl.ApplyOption) *AzureClusterControlPlaneMainVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterControlPlaneMainVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}

	return cDes
}

func canonicalizeAzureClusterControlPlaneMainVolumeSlice(des, initial []AzureClusterControlPlaneMainVolume, opts ...dcl.ApplyOption) []AzureClusterControlPlaneMainVolume {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterControlPlaneMainVolume, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterControlPlaneMainVolume(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterControlPlaneMainVolume, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterControlPlaneMainVolume(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterControlPlaneMainVolume(c *Client, des, nw *AzureClusterControlPlaneMainVolume) *AzureClusterControlPlaneMainVolume {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterControlPlaneMainVolume while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewAzureClusterControlPlaneMainVolumeSet(c *Client, des, nw []AzureClusterControlPlaneMainVolume) []AzureClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterControlPlaneMainVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterControlPlaneMainVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterControlPlaneMainVolumeSlice(c *Client, des, nw []AzureClusterControlPlaneMainVolume) []AzureClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterControlPlaneMainVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterControlPlaneMainVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterControlPlaneDatabaseEncryption(des, initial *AzureClusterControlPlaneDatabaseEncryption, opts ...dcl.ApplyOption) *AzureClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterControlPlaneDatabaseEncryption{}

	if dcl.StringCanonicalize(des.ResourceGroupId, initial.ResourceGroupId) || dcl.IsZeroValue(des.ResourceGroupId) {
		cDes.ResourceGroupId = initial.ResourceGroupId
	} else {
		cDes.ResourceGroupId = des.ResourceGroupId
	}
	if dcl.StringCanonicalize(des.KmsKeyIdentifier, initial.KmsKeyIdentifier) || dcl.IsZeroValue(des.KmsKeyIdentifier) {
		cDes.KmsKeyIdentifier = initial.KmsKeyIdentifier
	} else {
		cDes.KmsKeyIdentifier = des.KmsKeyIdentifier
	}

	return cDes
}

func canonicalizeAzureClusterControlPlaneDatabaseEncryptionSlice(des, initial []AzureClusterControlPlaneDatabaseEncryption, opts ...dcl.ApplyOption) []AzureClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterControlPlaneDatabaseEncryption, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterControlPlaneDatabaseEncryption(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterControlPlaneDatabaseEncryption, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterControlPlaneDatabaseEncryption(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterControlPlaneDatabaseEncryption(c *Client, des, nw *AzureClusterControlPlaneDatabaseEncryption) *AzureClusterControlPlaneDatabaseEncryption {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterControlPlaneDatabaseEncryption while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceGroupId, nw.ResourceGroupId) {
		nw.ResourceGroupId = des.ResourceGroupId
	}
	if dcl.StringCanonicalize(des.KmsKeyIdentifier, nw.KmsKeyIdentifier) {
		nw.KmsKeyIdentifier = des.KmsKeyIdentifier
	}

	return nw
}

func canonicalizeNewAzureClusterControlPlaneDatabaseEncryptionSet(c *Client, des, nw []AzureClusterControlPlaneDatabaseEncryption) []AzureClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterControlPlaneDatabaseEncryption
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterControlPlaneDatabaseEncryptionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterControlPlaneDatabaseEncryptionSlice(c *Client, des, nw []AzureClusterControlPlaneDatabaseEncryption) []AzureClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterControlPlaneDatabaseEncryption
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterControlPlaneDatabaseEncryption(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterAuthorization(des, initial *AzureClusterAuthorization, opts ...dcl.ApplyOption) *AzureClusterAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterAuthorization{}

	cDes.AdminUsers = canonicalizeAzureClusterAuthorizationAdminUsersSlice(des.AdminUsers, initial.AdminUsers, opts...)

	return cDes
}

func canonicalizeAzureClusterAuthorizationSlice(des, initial []AzureClusterAuthorization, opts ...dcl.ApplyOption) []AzureClusterAuthorization {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterAuthorization, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterAuthorization(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterAuthorization, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterAuthorization(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterAuthorization(c *Client, des, nw *AzureClusterAuthorization) *AzureClusterAuthorization {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterAuthorization while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.AdminUsers = canonicalizeNewAzureClusterAuthorizationAdminUsersSlice(c, des.AdminUsers, nw.AdminUsers)

	return nw
}

func canonicalizeNewAzureClusterAuthorizationSet(c *Client, des, nw []AzureClusterAuthorization) []AzureClusterAuthorization {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterAuthorization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterAuthorizationSlice(c *Client, des, nw []AzureClusterAuthorization) []AzureClusterAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterAuthorizationAdminUsers(des, initial *AzureClusterAuthorizationAdminUsers, opts ...dcl.ApplyOption) *AzureClusterAuthorizationAdminUsers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterAuthorizationAdminUsers{}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		cDes.Username = initial.Username
	} else {
		cDes.Username = des.Username
	}

	return cDes
}

func canonicalizeAzureClusterAuthorizationAdminUsersSlice(des, initial []AzureClusterAuthorizationAdminUsers, opts ...dcl.ApplyOption) []AzureClusterAuthorizationAdminUsers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterAuthorizationAdminUsers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterAuthorizationAdminUsers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterAuthorizationAdminUsers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterAuthorizationAdminUsers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterAuthorizationAdminUsers(c *Client, des, nw *AzureClusterAuthorizationAdminUsers) *AzureClusterAuthorizationAdminUsers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterAuthorizationAdminUsers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}

	return nw
}

func canonicalizeNewAzureClusterAuthorizationAdminUsersSet(c *Client, des, nw []AzureClusterAuthorizationAdminUsers) []AzureClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterAuthorizationAdminUsers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterAuthorizationAdminUsersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterAuthorizationAdminUsersSlice(c *Client, des, nw []AzureClusterAuthorizationAdminUsers) []AzureClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterAuthorizationAdminUsers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterAuthorizationAdminUsers(c, &d, &n))
	}

	return items
}

func canonicalizeAzureClusterWorkloadIdentityConfig(des, initial *AzureClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) *AzureClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureClusterWorkloadIdentityConfig{}

	if dcl.StringCanonicalize(des.IssuerUri, initial.IssuerUri) || dcl.IsZeroValue(des.IssuerUri) {
		cDes.IssuerUri = initial.IssuerUri
	} else {
		cDes.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.WorkloadPool, initial.WorkloadPool) || dcl.IsZeroValue(des.WorkloadPool) {
		cDes.WorkloadPool = initial.WorkloadPool
	} else {
		cDes.WorkloadPool = des.WorkloadPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, initial.IdentityProvider) || dcl.IsZeroValue(des.IdentityProvider) {
		cDes.IdentityProvider = initial.IdentityProvider
	} else {
		cDes.IdentityProvider = des.IdentityProvider
	}

	return cDes
}

func canonicalizeAzureClusterWorkloadIdentityConfigSlice(des, initial []AzureClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) []AzureClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AzureClusterWorkloadIdentityConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAzureClusterWorkloadIdentityConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AzureClusterWorkloadIdentityConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAzureClusterWorkloadIdentityConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAzureClusterWorkloadIdentityConfig(c *Client, des, nw *AzureClusterWorkloadIdentityConfig) *AzureClusterWorkloadIdentityConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for AzureClusterWorkloadIdentityConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IssuerUri, nw.IssuerUri) {
		nw.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.WorkloadPool, nw.WorkloadPool) {
		nw.WorkloadPool = des.WorkloadPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, nw.IdentityProvider) {
		nw.IdentityProvider = des.IdentityProvider
	}

	return nw
}

func canonicalizeNewAzureClusterWorkloadIdentityConfigSet(c *Client, des, nw []AzureClusterWorkloadIdentityConfig) []AzureClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureClusterWorkloadIdentityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureClusterWorkloadIdentityConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureClusterWorkloadIdentityConfigSlice(c *Client, des, nw []AzureClusterWorkloadIdentityConfig) []AzureClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureClusterWorkloadIdentityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureClusterWorkloadIdentityConfig(c, &d, &n))
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
func diffAzureCluster(c *Client, desired, actual *AzureCluster, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AzureRegion, actual.AzureRegion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureRegion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceGroupId, actual.ResourceGroupId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceGroupId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AzureClient, actual.AzureClient, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureClient")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Networking, actual.Networking, dcl.Info{ObjectFunction: compareAzureClusterNetworkingNewStyle, EmptyObject: EmptyAzureClusterNetworking, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Networking")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ControlPlane, actual.ControlPlane, dcl.Info{ObjectFunction: compareAzureClusterControlPlaneNewStyle, EmptyObject: EmptyAzureClusterControlPlane, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ControlPlane")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Authorization, actual.Authorization, dcl.Info{ObjectFunction: compareAzureClusterAuthorizationNewStyle, EmptyObject: EmptyAzureClusterAuthorization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Authorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reconciling, actual.Reconciling, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reconciling")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadIdentityConfig, actual.WorkloadIdentityConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareAzureClusterWorkloadIdentityConfigNewStyle, EmptyObject: EmptyAzureClusterWorkloadIdentityConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityConfig")); len(ds) != 0 || err != nil {
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
func compareAzureClusterNetworkingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterNetworking)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterNetworking or *AzureClusterNetworking", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterNetworking)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterNetworking", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VirtualNetworkId, actual.VirtualNetworkId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VirtualNetworkId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PodAddressCidrBlocks, actual.PodAddressCidrBlocks, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PodAddressCidrBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAddressCidrBlocks, actual.ServiceAddressCidrBlocks, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAddressCidrBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterControlPlaneNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterControlPlane)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlane or *AzureClusterControlPlane", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterControlPlane)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlane", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetId, actual.SubnetId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VmSize, actual.VmSize, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VmSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SshConfig, actual.SshConfig, dcl.Info{ObjectFunction: compareAzureClusterControlPlaneSshConfigNewStyle, EmptyObject: EmptyAzureClusterControlPlaneSshConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SshConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootVolume, actual.RootVolume, dcl.Info{ObjectFunction: compareAzureClusterControlPlaneRootVolumeNewStyle, EmptyObject: EmptyAzureClusterControlPlaneRootVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MainVolume, actual.MainVolume, dcl.Info{ObjectFunction: compareAzureClusterControlPlaneMainVolumeNewStyle, EmptyObject: EmptyAzureClusterControlPlaneMainVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MainVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatabaseEncryption, actual.DatabaseEncryption, dcl.Info{ObjectFunction: compareAzureClusterControlPlaneDatabaseEncryptionNewStyle, EmptyObject: EmptyAzureClusterControlPlaneDatabaseEncryption, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatabaseEncryption")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterControlPlaneSshConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterControlPlaneSshConfig)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneSshConfig or *AzureClusterControlPlaneSshConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterControlPlaneSshConfig)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneSshConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AuthorizedKey, actual.AuthorizedKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthorizedKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterControlPlaneRootVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterControlPlaneRootVolume)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneRootVolume or *AzureClusterControlPlaneRootVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterControlPlaneRootVolume)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneRootVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterControlPlaneMainVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterControlPlaneMainVolume)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneMainVolume or *AzureClusterControlPlaneMainVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterControlPlaneMainVolume)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneMainVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterControlPlaneDatabaseEncryptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterControlPlaneDatabaseEncryption)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneDatabaseEncryption or *AzureClusterControlPlaneDatabaseEncryption", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterControlPlaneDatabaseEncryption)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterControlPlaneDatabaseEncryption", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceGroupId, actual.ResourceGroupId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceGroupId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyIdentifier, actual.KmsKeyIdentifier, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyIdentifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterAuthorization or *AzureClusterAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterAuthorization)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AdminUsers, actual.AdminUsers, dcl.Info{ObjectFunction: compareAzureClusterAuthorizationAdminUsersNewStyle, EmptyObject: EmptyAzureClusterAuthorizationAdminUsers, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdminUsers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterAuthorizationAdminUsersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterAuthorizationAdminUsers)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterAuthorizationAdminUsers or *AzureClusterAuthorizationAdminUsers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterAuthorizationAdminUsers)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterAuthorizationAdminUsers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureClusterWorkloadIdentityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureClusterWorkloadIdentityConfig)
	if !ok {
		desiredNotPointer, ok := d.(AzureClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterWorkloadIdentityConfig or *AzureClusterWorkloadIdentityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureClusterWorkloadIdentityConfig)
	if !ok {
		actualNotPointer, ok := a.(AzureClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureClusterWorkloadIdentityConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IssuerUri, actual.IssuerUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IssuerUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadPool, actual.WorkloadPool, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdentityProvider, actual.IdentityProvider, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdentityProvider")); len(ds) != 0 || err != nil {
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
func (r *AzureCluster) urlNormalized() *AzureCluster {
	normalized := dcl.Copy(*r).(AzureCluster)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AzureRegion = dcl.SelfLinkToName(r.AzureRegion)
	normalized.ResourceGroupId = dcl.SelfLinkToName(r.ResourceGroupId)
	normalized.AzureClient = dcl.SelfLinkToName(r.AzureClient)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AzureCluster) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AzureCluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AzureCluster) marshal(c *Client) ([]byte, error) {
	m, err := expandAzureCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AzureCluster: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAzureCluster decodes JSON responses into the AzureCluster resource schema.
func unmarshalAzureCluster(b []byte, c *Client) (*AzureCluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAzureCluster(m, c)
}

func unmarshalMapAzureCluster(m map[string]interface{}, c *Client) (*AzureCluster, error) {

	return flattenAzureCluster(c, m), nil
}

// expandAzureCluster expands AzureCluster into a JSON request object.
func expandAzureCluster(c *Client, f *AzureCluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/azureClusters/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.AzureRegion; !dcl.IsEmptyValueIndirect(v) {
		m["azureRegion"] = v
	}
	if v := f.ResourceGroupId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroupId"] = v
	}
	if v := f.AzureClient; !dcl.IsEmptyValueIndirect(v) {
		m["azureClient"] = v
	}
	if v, err := expandAzureClusterNetworking(c, f.Networking); err != nil {
		return nil, fmt.Errorf("error expanding Networking into networking: %w", err)
	} else if v != nil {
		m["networking"] = v
	}
	if v, err := expandAzureClusterControlPlane(c, f.ControlPlane); err != nil {
		return nil, fmt.Errorf("error expanding ControlPlane into controlPlane: %w", err)
	} else if v != nil {
		m["controlPlane"] = v
	}
	if v, err := expandAzureClusterAuthorization(c, f.Authorization); err != nil {
		return nil, fmt.Errorf("error expanding Authorization into authorization: %w", err)
	} else if v != nil {
		m["authorization"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
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

// flattenAzureCluster flattens AzureCluster from a JSON request object into the
// AzureCluster type.
func flattenAzureCluster(c *Client, i interface{}) *AzureCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AzureCluster{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.AzureRegion = dcl.FlattenString(m["azureRegion"])
	res.ResourceGroupId = dcl.FlattenString(m["resourceGroupId"])
	res.AzureClient = dcl.FlattenString(m["azureClient"])
	res.Networking = flattenAzureClusterNetworking(c, m["networking"])
	res.ControlPlane = flattenAzureClusterControlPlane(c, m["controlPlane"])
	res.Authorization = flattenAzureClusterAuthorization(c, m["authorization"])
	res.State = flattenAzureClusterStateEnum(m["state"])
	res.Endpoint = dcl.FlattenString(m["endpoint"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.WorkloadIdentityConfig = flattenAzureClusterWorkloadIdentityConfig(c, m["workloadIdentityConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandAzureClusterNetworkingMap expands the contents of AzureClusterNetworking into a JSON
// request object.
func expandAzureClusterNetworkingMap(c *Client, f map[string]AzureClusterNetworking) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterNetworkingSlice expands the contents of AzureClusterNetworking into a JSON
// request object.
func expandAzureClusterNetworkingSlice(c *Client, f []AzureClusterNetworking) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterNetworkingMap flattens the contents of AzureClusterNetworking from a JSON
// response object.
func flattenAzureClusterNetworkingMap(c *Client, i interface{}) map[string]AzureClusterNetworking {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterNetworking{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterNetworking{}
	}

	items := make(map[string]AzureClusterNetworking)
	for k, item := range a {
		items[k] = *flattenAzureClusterNetworking(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterNetworkingSlice flattens the contents of AzureClusterNetworking from a JSON
// response object.
func flattenAzureClusterNetworkingSlice(c *Client, i interface{}) []AzureClusterNetworking {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterNetworking{}
	}

	if len(a) == 0 {
		return []AzureClusterNetworking{}
	}

	items := make([]AzureClusterNetworking, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterNetworking(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterNetworking expands an instance of AzureClusterNetworking into a JSON
// request object.
func expandAzureClusterNetworking(c *Client, f *AzureClusterNetworking) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.VirtualNetworkId; !dcl.IsEmptyValueIndirect(v) {
		m["virtualNetworkId"] = v
	}
	if v := f.PodAddressCidrBlocks; v != nil {
		m["podAddressCidrBlocks"] = v
	}
	if v := f.ServiceAddressCidrBlocks; v != nil {
		m["serviceAddressCidrBlocks"] = v
	}

	return m, nil
}

// flattenAzureClusterNetworking flattens an instance of AzureClusterNetworking from a JSON
// response object.
func flattenAzureClusterNetworking(c *Client, i interface{}) *AzureClusterNetworking {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterNetworking{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterNetworking
	}
	r.VirtualNetworkId = dcl.FlattenString(m["virtualNetworkId"])
	r.PodAddressCidrBlocks = dcl.FlattenStringSlice(m["podAddressCidrBlocks"])
	r.ServiceAddressCidrBlocks = dcl.FlattenStringSlice(m["serviceAddressCidrBlocks"])

	return r
}

// expandAzureClusterControlPlaneMap expands the contents of AzureClusterControlPlane into a JSON
// request object.
func expandAzureClusterControlPlaneMap(c *Client, f map[string]AzureClusterControlPlane) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterControlPlaneSlice expands the contents of AzureClusterControlPlane into a JSON
// request object.
func expandAzureClusterControlPlaneSlice(c *Client, f []AzureClusterControlPlane) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterControlPlaneMap flattens the contents of AzureClusterControlPlane from a JSON
// response object.
func flattenAzureClusterControlPlaneMap(c *Client, i interface{}) map[string]AzureClusterControlPlane {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterControlPlane{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterControlPlane{}
	}

	items := make(map[string]AzureClusterControlPlane)
	for k, item := range a {
		items[k] = *flattenAzureClusterControlPlane(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterControlPlaneSlice flattens the contents of AzureClusterControlPlane from a JSON
// response object.
func flattenAzureClusterControlPlaneSlice(c *Client, i interface{}) []AzureClusterControlPlane {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterControlPlane{}
	}

	if len(a) == 0 {
		return []AzureClusterControlPlane{}
	}

	items := make([]AzureClusterControlPlane, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterControlPlane(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterControlPlane expands an instance of AzureClusterControlPlane into a JSON
// request object.
func expandAzureClusterControlPlane(c *Client, f *AzureClusterControlPlane) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.SubnetId; !dcl.IsEmptyValueIndirect(v) {
		m["subnetId"] = v
	}
	if v := f.VmSize; !dcl.IsEmptyValueIndirect(v) {
		m["vmSize"] = v
	}
	if v, err := expandAzureClusterControlPlaneSshConfig(c, f.SshConfig); err != nil {
		return nil, fmt.Errorf("error expanding SshConfig into sshConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sshConfig"] = v
	}
	if v, err := expandAzureClusterControlPlaneRootVolume(c, f.RootVolume); err != nil {
		return nil, fmt.Errorf("error expanding RootVolume into rootVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rootVolume"] = v
	}
	if v, err := expandAzureClusterControlPlaneMainVolume(c, f.MainVolume); err != nil {
		return nil, fmt.Errorf("error expanding MainVolume into mainVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mainVolume"] = v
	}
	if v, err := expandAzureClusterControlPlaneDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into databaseEncryption: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseEncryption"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}

	return m, nil
}

// flattenAzureClusterControlPlane flattens an instance of AzureClusterControlPlane from a JSON
// response object.
func flattenAzureClusterControlPlane(c *Client, i interface{}) *AzureClusterControlPlane {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterControlPlane{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterControlPlane
	}
	r.Version = dcl.FlattenString(m["version"])
	r.SubnetId = dcl.FlattenString(m["subnetId"])
	r.VmSize = dcl.FlattenString(m["vmSize"])
	r.SshConfig = flattenAzureClusterControlPlaneSshConfig(c, m["sshConfig"])
	r.RootVolume = flattenAzureClusterControlPlaneRootVolume(c, m["rootVolume"])
	r.MainVolume = flattenAzureClusterControlPlaneMainVolume(c, m["mainVolume"])
	r.DatabaseEncryption = flattenAzureClusterControlPlaneDatabaseEncryption(c, m["databaseEncryption"])
	r.Tags = dcl.FlattenKeyValuePairs(m["tags"])

	return r
}

// expandAzureClusterControlPlaneSshConfigMap expands the contents of AzureClusterControlPlaneSshConfig into a JSON
// request object.
func expandAzureClusterControlPlaneSshConfigMap(c *Client, f map[string]AzureClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterControlPlaneSshConfigSlice expands the contents of AzureClusterControlPlaneSshConfig into a JSON
// request object.
func expandAzureClusterControlPlaneSshConfigSlice(c *Client, f []AzureClusterControlPlaneSshConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterControlPlaneSshConfigMap flattens the contents of AzureClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAzureClusterControlPlaneSshConfigMap(c *Client, i interface{}) map[string]AzureClusterControlPlaneSshConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterControlPlaneSshConfig{}
	}

	items := make(map[string]AzureClusterControlPlaneSshConfig)
	for k, item := range a {
		items[k] = *flattenAzureClusterControlPlaneSshConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterControlPlaneSshConfigSlice flattens the contents of AzureClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAzureClusterControlPlaneSshConfigSlice(c *Client, i interface{}) []AzureClusterControlPlaneSshConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return []AzureClusterControlPlaneSshConfig{}
	}

	items := make([]AzureClusterControlPlaneSshConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterControlPlaneSshConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterControlPlaneSshConfig expands an instance of AzureClusterControlPlaneSshConfig into a JSON
// request object.
func expandAzureClusterControlPlaneSshConfig(c *Client, f *AzureClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AuthorizedKey; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedKey"] = v
	}

	return m, nil
}

// flattenAzureClusterControlPlaneSshConfig flattens an instance of AzureClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAzureClusterControlPlaneSshConfig(c *Client, i interface{}) *AzureClusterControlPlaneSshConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterControlPlaneSshConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterControlPlaneSshConfig
	}
	r.AuthorizedKey = dcl.FlattenString(m["authorizedKey"])

	return r
}

// expandAzureClusterControlPlaneRootVolumeMap expands the contents of AzureClusterControlPlaneRootVolume into a JSON
// request object.
func expandAzureClusterControlPlaneRootVolumeMap(c *Client, f map[string]AzureClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterControlPlaneRootVolumeSlice expands the contents of AzureClusterControlPlaneRootVolume into a JSON
// request object.
func expandAzureClusterControlPlaneRootVolumeSlice(c *Client, f []AzureClusterControlPlaneRootVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterControlPlaneRootVolumeMap flattens the contents of AzureClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneRootVolumeMap(c *Client, i interface{}) map[string]AzureClusterControlPlaneRootVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterControlPlaneRootVolume{}
	}

	items := make(map[string]AzureClusterControlPlaneRootVolume)
	for k, item := range a {
		items[k] = *flattenAzureClusterControlPlaneRootVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterControlPlaneRootVolumeSlice flattens the contents of AzureClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneRootVolumeSlice(c *Client, i interface{}) []AzureClusterControlPlaneRootVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return []AzureClusterControlPlaneRootVolume{}
	}

	items := make([]AzureClusterControlPlaneRootVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterControlPlaneRootVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterControlPlaneRootVolume expands an instance of AzureClusterControlPlaneRootVolume into a JSON
// request object.
func expandAzureClusterControlPlaneRootVolume(c *Client, f *AzureClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}

	return m, nil
}

// flattenAzureClusterControlPlaneRootVolume flattens an instance of AzureClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneRootVolume(c *Client, i interface{}) *AzureClusterControlPlaneRootVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterControlPlaneRootVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterControlPlaneRootVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])

	return r
}

// expandAzureClusterControlPlaneMainVolumeMap expands the contents of AzureClusterControlPlaneMainVolume into a JSON
// request object.
func expandAzureClusterControlPlaneMainVolumeMap(c *Client, f map[string]AzureClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterControlPlaneMainVolumeSlice expands the contents of AzureClusterControlPlaneMainVolume into a JSON
// request object.
func expandAzureClusterControlPlaneMainVolumeSlice(c *Client, f []AzureClusterControlPlaneMainVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterControlPlaneMainVolumeMap flattens the contents of AzureClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneMainVolumeMap(c *Client, i interface{}) map[string]AzureClusterControlPlaneMainVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterControlPlaneMainVolume{}
	}

	items := make(map[string]AzureClusterControlPlaneMainVolume)
	for k, item := range a {
		items[k] = *flattenAzureClusterControlPlaneMainVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterControlPlaneMainVolumeSlice flattens the contents of AzureClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneMainVolumeSlice(c *Client, i interface{}) []AzureClusterControlPlaneMainVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return []AzureClusterControlPlaneMainVolume{}
	}

	items := make([]AzureClusterControlPlaneMainVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterControlPlaneMainVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterControlPlaneMainVolume expands an instance of AzureClusterControlPlaneMainVolume into a JSON
// request object.
func expandAzureClusterControlPlaneMainVolume(c *Client, f *AzureClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}

	return m, nil
}

// flattenAzureClusterControlPlaneMainVolume flattens an instance of AzureClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAzureClusterControlPlaneMainVolume(c *Client, i interface{}) *AzureClusterControlPlaneMainVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterControlPlaneMainVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterControlPlaneMainVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])

	return r
}

// expandAzureClusterControlPlaneDatabaseEncryptionMap expands the contents of AzureClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAzureClusterControlPlaneDatabaseEncryptionMap(c *Client, f map[string]AzureClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterControlPlaneDatabaseEncryptionSlice expands the contents of AzureClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAzureClusterControlPlaneDatabaseEncryptionSlice(c *Client, f []AzureClusterControlPlaneDatabaseEncryption) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterControlPlaneDatabaseEncryptionMap flattens the contents of AzureClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAzureClusterControlPlaneDatabaseEncryptionMap(c *Client, i interface{}) map[string]AzureClusterControlPlaneDatabaseEncryption {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterControlPlaneDatabaseEncryption{}
	}

	items := make(map[string]AzureClusterControlPlaneDatabaseEncryption)
	for k, item := range a {
		items[k] = *flattenAzureClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterControlPlaneDatabaseEncryptionSlice flattens the contents of AzureClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAzureClusterControlPlaneDatabaseEncryptionSlice(c *Client, i interface{}) []AzureClusterControlPlaneDatabaseEncryption {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return []AzureClusterControlPlaneDatabaseEncryption{}
	}

	items := make([]AzureClusterControlPlaneDatabaseEncryption, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterControlPlaneDatabaseEncryption expands an instance of AzureClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAzureClusterControlPlaneDatabaseEncryption(c *Client, f *AzureClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceGroupId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroupId"] = v
	}
	if v := f.KmsKeyIdentifier; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyIdentifier"] = v
	}

	return m, nil
}

// flattenAzureClusterControlPlaneDatabaseEncryption flattens an instance of AzureClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAzureClusterControlPlaneDatabaseEncryption(c *Client, i interface{}) *AzureClusterControlPlaneDatabaseEncryption {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterControlPlaneDatabaseEncryption{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterControlPlaneDatabaseEncryption
	}
	r.ResourceGroupId = dcl.FlattenString(m["resourceGroupId"])
	r.KmsKeyIdentifier = dcl.FlattenString(m["kmsKeyIdentifier"])

	return r
}

// expandAzureClusterAuthorizationMap expands the contents of AzureClusterAuthorization into a JSON
// request object.
func expandAzureClusterAuthorizationMap(c *Client, f map[string]AzureClusterAuthorization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterAuthorizationSlice expands the contents of AzureClusterAuthorization into a JSON
// request object.
func expandAzureClusterAuthorizationSlice(c *Client, f []AzureClusterAuthorization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterAuthorizationMap flattens the contents of AzureClusterAuthorization from a JSON
// response object.
func flattenAzureClusterAuthorizationMap(c *Client, i interface{}) map[string]AzureClusterAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterAuthorization{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterAuthorization{}
	}

	items := make(map[string]AzureClusterAuthorization)
	for k, item := range a {
		items[k] = *flattenAzureClusterAuthorization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterAuthorizationSlice flattens the contents of AzureClusterAuthorization from a JSON
// response object.
func flattenAzureClusterAuthorizationSlice(c *Client, i interface{}) []AzureClusterAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterAuthorization{}
	}

	if len(a) == 0 {
		return []AzureClusterAuthorization{}
	}

	items := make([]AzureClusterAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterAuthorization(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterAuthorization expands an instance of AzureClusterAuthorization into a JSON
// request object.
func expandAzureClusterAuthorization(c *Client, f *AzureClusterAuthorization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAzureClusterAuthorizationAdminUsersSlice(c, f.AdminUsers); err != nil {
		return nil, fmt.Errorf("error expanding AdminUsers into adminUsers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["adminUsers"] = v
	}

	return m, nil
}

// flattenAzureClusterAuthorization flattens an instance of AzureClusterAuthorization from a JSON
// response object.
func flattenAzureClusterAuthorization(c *Client, i interface{}) *AzureClusterAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterAuthorization
	}
	r.AdminUsers = flattenAzureClusterAuthorizationAdminUsersSlice(c, m["adminUsers"])

	return r
}

// expandAzureClusterAuthorizationAdminUsersMap expands the contents of AzureClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAzureClusterAuthorizationAdminUsersMap(c *Client, f map[string]AzureClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterAuthorizationAdminUsersSlice expands the contents of AzureClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAzureClusterAuthorizationAdminUsersSlice(c *Client, f []AzureClusterAuthorizationAdminUsers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterAuthorizationAdminUsersMap flattens the contents of AzureClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAzureClusterAuthorizationAdminUsersMap(c *Client, i interface{}) map[string]AzureClusterAuthorizationAdminUsers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterAuthorizationAdminUsers{}
	}

	items := make(map[string]AzureClusterAuthorizationAdminUsers)
	for k, item := range a {
		items[k] = *flattenAzureClusterAuthorizationAdminUsers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterAuthorizationAdminUsersSlice flattens the contents of AzureClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAzureClusterAuthorizationAdminUsersSlice(c *Client, i interface{}) []AzureClusterAuthorizationAdminUsers {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return []AzureClusterAuthorizationAdminUsers{}
	}

	items := make([]AzureClusterAuthorizationAdminUsers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterAuthorizationAdminUsers(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterAuthorizationAdminUsers expands an instance of AzureClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAzureClusterAuthorizationAdminUsers(c *Client, f *AzureClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}

	return m, nil
}

// flattenAzureClusterAuthorizationAdminUsers flattens an instance of AzureClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAzureClusterAuthorizationAdminUsers(c *Client, i interface{}) *AzureClusterAuthorizationAdminUsers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterAuthorizationAdminUsers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterAuthorizationAdminUsers
	}
	r.Username = dcl.FlattenString(m["username"])

	return r
}

// expandAzureClusterWorkloadIdentityConfigMap expands the contents of AzureClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAzureClusterWorkloadIdentityConfigMap(c *Client, f map[string]AzureClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureClusterWorkloadIdentityConfigSlice expands the contents of AzureClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAzureClusterWorkloadIdentityConfigSlice(c *Client, f []AzureClusterWorkloadIdentityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureClusterWorkloadIdentityConfigMap flattens the contents of AzureClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAzureClusterWorkloadIdentityConfigMap(c *Client, i interface{}) map[string]AzureClusterWorkloadIdentityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterWorkloadIdentityConfig{}
	}

	items := make(map[string]AzureClusterWorkloadIdentityConfig)
	for k, item := range a {
		items[k] = *flattenAzureClusterWorkloadIdentityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureClusterWorkloadIdentityConfigSlice flattens the contents of AzureClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAzureClusterWorkloadIdentityConfigSlice(c *Client, i interface{}) []AzureClusterWorkloadIdentityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return []AzureClusterWorkloadIdentityConfig{}
	}

	items := make([]AzureClusterWorkloadIdentityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterWorkloadIdentityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureClusterWorkloadIdentityConfig expands an instance of AzureClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAzureClusterWorkloadIdentityConfig(c *Client, f *AzureClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IssuerUri; !dcl.IsEmptyValueIndirect(v) {
		m["issuerUri"] = v
	}
	if v := f.WorkloadPool; !dcl.IsEmptyValueIndirect(v) {
		m["workloadPool"] = v
	}
	if v := f.IdentityProvider; !dcl.IsEmptyValueIndirect(v) {
		m["identityProvider"] = v
	}

	return m, nil
}

// flattenAzureClusterWorkloadIdentityConfig flattens an instance of AzureClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAzureClusterWorkloadIdentityConfig(c *Client, i interface{}) *AzureClusterWorkloadIdentityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureClusterWorkloadIdentityConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureClusterWorkloadIdentityConfig
	}
	r.IssuerUri = dcl.FlattenString(m["issuerUri"])
	r.WorkloadPool = dcl.FlattenString(m["workloadPool"])
	r.IdentityProvider = dcl.FlattenString(m["identityProvider"])

	return r
}

// flattenAzureClusterStateEnumMap flattens the contents of AzureClusterStateEnum from a JSON
// response object.
func flattenAzureClusterStateEnumMap(c *Client, i interface{}) map[string]AzureClusterStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureClusterStateEnum{}
	}

	if len(a) == 0 {
		return map[string]AzureClusterStateEnum{}
	}

	items := make(map[string]AzureClusterStateEnum)
	for k, item := range a {
		items[k] = *flattenAzureClusterStateEnum(item.(interface{}))
	}

	return items
}

// flattenAzureClusterStateEnumSlice flattens the contents of AzureClusterStateEnum from a JSON
// response object.
func flattenAzureClusterStateEnumSlice(c *Client, i interface{}) []AzureClusterStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureClusterStateEnum{}
	}

	if len(a) == 0 {
		return []AzureClusterStateEnum{}
	}

	items := make([]AzureClusterStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureClusterStateEnum(item.(interface{})))
	}

	return items
}

// flattenAzureClusterStateEnum asserts that an interface is a string, and returns a
// pointer to a *AzureClusterStateEnum with the same value as that string.
func flattenAzureClusterStateEnum(i interface{}) *AzureClusterStateEnum {
	s, ok := i.(string)
	if !ok {
		return AzureClusterStateEnumRef("")
	}

	return AzureClusterStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AzureCluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAzureCluster(b, c)
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

type azureClusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         azureClusterApiOperation
}

func convertFieldDiffsToAzureClusterDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]azureClusterDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []azureClusterDiff
	// For each operation name, create a azureClusterDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := azureClusterDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAzureClusterApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAzureClusterApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (azureClusterApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractAzureClusterFields(r *AzureCluster) error {
	return nil
}
