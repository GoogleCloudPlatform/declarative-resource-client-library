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

func (r *AwsCluster) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "networking"); err != nil {
		return err
	}
	if err := dcl.Required(r, "awsRegion"); err != nil {
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
func (r *AwsClusterNetworking) validate() error {
	if err := dcl.Required(r, "vpcId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "podAddressCidrBlocks"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceAddressCidrBlocks"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceLoadBalancerSubnetIds"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterControlPlane) validate() error {
	if err := dcl.Required(r, "version"); err != nil {
		return err
	}
	if err := dcl.Required(r, "subnetIds"); err != nil {
		return err
	}
	if err := dcl.Required(r, "iamInstanceProfile"); err != nil {
		return err
	}
	if err := dcl.Required(r, "databaseEncryption"); err != nil {
		return err
	}
	if err := dcl.Required(r, "awsServicesAuthentication"); err != nil {
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
	if !dcl.IsEmptyValueIndirect(r.AwsServicesAuthentication) {
		if err := r.AwsServicesAuthentication.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AwsClusterControlPlaneSshConfig) validate() error {
	if err := dcl.Required(r, "ec2KeyPair"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterControlPlaneRootVolume) validate() error {
	return nil
}
func (r *AwsClusterControlPlaneMainVolume) validate() error {
	return nil
}
func (r *AwsClusterControlPlaneDatabaseEncryption) validate() error {
	if err := dcl.Required(r, "kmsKeyArn"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterControlPlaneAwsServicesAuthentication) validate() error {
	if err := dcl.Required(r, "roleArn"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterAuthorization) validate() error {
	if err := dcl.Required(r, "adminUsers"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterAuthorizationAdminUsers) validate() error {
	if err := dcl.Required(r, "username"); err != nil {
		return err
	}
	return nil
}
func (r *AwsClusterWorkloadIdentityConfig) validate() error {
	return nil
}
func (r *AwsCluster) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-gkemulticloud.googleapis.com/v1", params)
}

func (r *AwsCluster) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AwsCluster) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters", nr.basePath(), userBasePath, params), nil

}

func (r *AwsCluster) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters?awsClusterId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AwsCluster) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

// awsClusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type awsClusterApiOperation interface {
	do(context.Context, *AwsCluster, *Client) error
}

func (c *Client) listAwsClusterRaw(ctx context.Context, r *AwsCluster, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AwsClusterMaxPage {
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

type listAwsClusterOperation struct {
	AwsClusters []map[string]interface{} `json:"awsClusters"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listAwsCluster(ctx context.Context, r *AwsCluster, pageToken string, pageSize int32) ([]*AwsCluster, string, error) {
	b, err := c.listAwsClusterRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAwsClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AwsCluster
	for _, v := range m.AwsClusters {
		res, err := unmarshalMapAwsCluster(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAwsCluster(ctx context.Context, f func(*AwsCluster) bool, resources []*AwsCluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAwsCluster(ctx, res)
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

type deleteAwsClusterOperation struct{}

func (op *deleteAwsClusterOperation) do(ctx context.Context, r *AwsCluster, c *Client) error {
	r, err := c.GetAwsCluster(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AwsCluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAwsCluster checking for existence. error: %v", err)
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
		_, err = c.GetAwsCluster(ctx, r)
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
type createAwsClusterOperation struct {
	response map[string]interface{}
}

func (op *createAwsClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAwsClusterOperation) do(ctx context.Context, r *AwsCluster, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)
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
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetAwsCluster(ctx, r); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAwsClusterRaw(ctx context.Context, r *AwsCluster) ([]byte, error) {

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

func (c *Client) awsClusterDiffsForRawDesired(ctx context.Context, rawDesired *AwsCluster, opts ...dcl.ApplyOption) (initial, desired *AwsCluster, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AwsCluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AwsCluster); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AwsCluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAwsCluster(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AwsCluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AwsCluster resource: %v", err)
		}
		c.Config.Logger.Info("Found that AwsCluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAwsClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for AwsCluster: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AwsCluster: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAwsClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AwsCluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAwsClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AwsCluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAwsCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAwsClusterInitialState(rawInitial, rawDesired *AwsCluster) (*AwsCluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAwsClusterDesiredState(rawDesired, rawInitial *AwsCluster, opts ...dcl.ApplyOption) (*AwsCluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Networking = canonicalizeAwsClusterNetworking(rawDesired.Networking, nil, opts...)
		rawDesired.ControlPlane = canonicalizeAwsClusterControlPlane(rawDesired.ControlPlane, nil, opts...)
		rawDesired.Authorization = canonicalizeAwsClusterAuthorization(rawDesired.Authorization, nil, opts...)
		rawDesired.WorkloadIdentityConfig = canonicalizeAwsClusterWorkloadIdentityConfig(rawDesired.WorkloadIdentityConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &AwsCluster{}
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
	canonicalDesired.Networking = canonicalizeAwsClusterNetworking(rawDesired.Networking, rawInitial.Networking, opts...)
	if dcl.StringCanonicalize(rawDesired.AwsRegion, rawInitial.AwsRegion) {
		canonicalDesired.AwsRegion = rawInitial.AwsRegion
	} else {
		canonicalDesired.AwsRegion = rawDesired.AwsRegion
	}
	canonicalDesired.ControlPlane = canonicalizeAwsClusterControlPlane(rawDesired.ControlPlane, rawInitial.ControlPlane, opts...)
	canonicalDesired.Authorization = canonicalizeAwsClusterAuthorization(rawDesired.Authorization, rawInitial.Authorization, opts...)
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

func canonicalizeAwsClusterNewState(c *Client, rawNew, rawDesired *AwsCluster) (*AwsCluster, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Networking) && dcl.IsEmptyValueIndirect(rawDesired.Networking) {
		rawNew.Networking = rawDesired.Networking
	} else {
		rawNew.Networking = canonicalizeNewAwsClusterNetworking(c, rawDesired.Networking, rawNew.Networking)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AwsRegion) && dcl.IsEmptyValueIndirect(rawDesired.AwsRegion) {
		rawNew.AwsRegion = rawDesired.AwsRegion
	} else {
		if dcl.StringCanonicalize(rawDesired.AwsRegion, rawNew.AwsRegion) {
			rawNew.AwsRegion = rawDesired.AwsRegion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ControlPlane) && dcl.IsEmptyValueIndirect(rawDesired.ControlPlane) {
		rawNew.ControlPlane = rawDesired.ControlPlane
	} else {
		rawNew.ControlPlane = canonicalizeNewAwsClusterControlPlane(c, rawDesired.ControlPlane, rawNew.ControlPlane)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Authorization) && dcl.IsEmptyValueIndirect(rawDesired.Authorization) {
		rawNew.Authorization = rawDesired.Authorization
	} else {
		rawNew.Authorization = canonicalizeNewAwsClusterAuthorization(c, rawDesired.Authorization, rawNew.Authorization)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Endpoint) && dcl.IsEmptyValueIndirect(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.Endpoint, rawNew.Endpoint) {
			rawNew.Endpoint = rawDesired.Endpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Reconciling) && dcl.IsEmptyValueIndirect(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
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

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.WorkloadIdentityConfig) && dcl.IsEmptyValueIndirect(rawDesired.WorkloadIdentityConfig) {
		rawNew.WorkloadIdentityConfig = rawDesired.WorkloadIdentityConfig
	} else {
		rawNew.WorkloadIdentityConfig = canonicalizeNewAwsClusterWorkloadIdentityConfig(c, rawDesired.WorkloadIdentityConfig, rawNew.WorkloadIdentityConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAwsClusterNetworking(des, initial *AwsClusterNetworking, opts ...dcl.ApplyOption) *AwsClusterNetworking {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterNetworking{}

	if dcl.StringCanonicalize(des.VPCId, initial.VPCId) || dcl.IsZeroValue(des.VPCId) {
		cDes.VPCId = initial.VPCId
	} else {
		cDes.VPCId = des.VPCId
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
	if dcl.IsZeroValue(des.ServiceLoadBalancerSubnetIds) {
		des.ServiceLoadBalancerSubnetIds = initial.ServiceLoadBalancerSubnetIds
	} else {
		cDes.ServiceLoadBalancerSubnetIds = des.ServiceLoadBalancerSubnetIds
	}

	return cDes
}

func canonicalizeNewAwsClusterNetworking(c *Client, des, nw *AwsClusterNetworking) *AwsClusterNetworking {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.VPCId, nw.VPCId) {
		nw.VPCId = des.VPCId
	}

	return nw
}

func canonicalizeNewAwsClusterNetworkingSet(c *Client, des, nw []AwsClusterNetworking) []AwsClusterNetworking {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterNetworking
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterNetworkingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterNetworkingSlice(c *Client, des, nw []AwsClusterNetworking) []AwsClusterNetworking {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterNetworking
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterNetworking(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlane(des, initial *AwsClusterControlPlane, opts ...dcl.ApplyOption) *AwsClusterControlPlane {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlane{}

	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.StringCanonicalize(des.InstanceType, initial.InstanceType) || dcl.IsZeroValue(des.InstanceType) {
		cDes.InstanceType = initial.InstanceType
	} else {
		cDes.InstanceType = des.InstanceType
	}
	cDes.SshConfig = canonicalizeAwsClusterControlPlaneSshConfig(des.SshConfig, initial.SshConfig, opts...)
	if dcl.IsZeroValue(des.SubnetIds) {
		des.SubnetIds = initial.SubnetIds
	} else {
		cDes.SubnetIds = des.SubnetIds
	}
	if dcl.IsZeroValue(des.SecurityGroupIds) {
		des.SecurityGroupIds = initial.SecurityGroupIds
	} else {
		cDes.SecurityGroupIds = des.SecurityGroupIds
	}
	if dcl.StringCanonicalize(des.IamInstanceProfile, initial.IamInstanceProfile) || dcl.IsZeroValue(des.IamInstanceProfile) {
		cDes.IamInstanceProfile = initial.IamInstanceProfile
	} else {
		cDes.IamInstanceProfile = des.IamInstanceProfile
	}
	cDes.RootVolume = canonicalizeAwsClusterControlPlaneRootVolume(des.RootVolume, initial.RootVolume, opts...)
	cDes.MainVolume = canonicalizeAwsClusterControlPlaneMainVolume(des.MainVolume, initial.MainVolume, opts...)
	cDes.DatabaseEncryption = canonicalizeAwsClusterControlPlaneDatabaseEncryption(des.DatabaseEncryption, initial.DatabaseEncryption, opts...)
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	cDes.AwsServicesAuthentication = canonicalizeAwsClusterControlPlaneAwsServicesAuthentication(des.AwsServicesAuthentication, initial.AwsServicesAuthentication, opts...)

	return cDes
}

func canonicalizeNewAwsClusterControlPlane(c *Client, des, nw *AwsClusterControlPlane) *AwsClusterControlPlane {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.InstanceType, nw.InstanceType) {
		nw.InstanceType = des.InstanceType
	}
	nw.SshConfig = canonicalizeNewAwsClusterControlPlaneSshConfig(c, des.SshConfig, nw.SshConfig)
	if dcl.StringCanonicalize(des.IamInstanceProfile, nw.IamInstanceProfile) {
		nw.IamInstanceProfile = des.IamInstanceProfile
	}
	nw.RootVolume = canonicalizeNewAwsClusterControlPlaneRootVolume(c, des.RootVolume, nw.RootVolume)
	nw.MainVolume = canonicalizeNewAwsClusterControlPlaneMainVolume(c, des.MainVolume, nw.MainVolume)
	nw.DatabaseEncryption = canonicalizeNewAwsClusterControlPlaneDatabaseEncryption(c, des.DatabaseEncryption, nw.DatabaseEncryption)
	nw.AwsServicesAuthentication = canonicalizeNewAwsClusterControlPlaneAwsServicesAuthentication(c, des.AwsServicesAuthentication, nw.AwsServicesAuthentication)

	return nw
}

func canonicalizeNewAwsClusterControlPlaneSet(c *Client, des, nw []AwsClusterControlPlane) []AwsClusterControlPlane {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlane
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneSlice(c *Client, des, nw []AwsClusterControlPlane) []AwsClusterControlPlane {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlane
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlane(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlaneSshConfig(des, initial *AwsClusterControlPlaneSshConfig, opts ...dcl.ApplyOption) *AwsClusterControlPlaneSshConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlaneSshConfig{}

	if dcl.StringCanonicalize(des.Ec2KeyPair, initial.Ec2KeyPair) || dcl.IsZeroValue(des.Ec2KeyPair) {
		cDes.Ec2KeyPair = initial.Ec2KeyPair
	} else {
		cDes.Ec2KeyPair = des.Ec2KeyPair
	}

	return cDes
}

func canonicalizeNewAwsClusterControlPlaneSshConfig(c *Client, des, nw *AwsClusterControlPlaneSshConfig) *AwsClusterControlPlaneSshConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Ec2KeyPair, nw.Ec2KeyPair) {
		nw.Ec2KeyPair = des.Ec2KeyPair
	}

	return nw
}

func canonicalizeNewAwsClusterControlPlaneSshConfigSet(c *Client, des, nw []AwsClusterControlPlaneSshConfig) []AwsClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlaneSshConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneSshConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneSshConfigSlice(c *Client, des, nw []AwsClusterControlPlaneSshConfig) []AwsClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlaneSshConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlaneSshConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlaneRootVolume(des, initial *AwsClusterControlPlaneRootVolume, opts ...dcl.ApplyOption) *AwsClusterControlPlaneRootVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlaneRootVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}
	if dcl.IsZeroValue(des.VolumeType) {
		des.VolumeType = initial.VolumeType
	} else {
		cDes.VolumeType = des.VolumeType
	}
	if dcl.IsZeroValue(des.Iops) {
		des.Iops = initial.Iops
	} else {
		cDes.Iops = des.Iops
	}
	if dcl.StringCanonicalize(des.KmsKeyArn, initial.KmsKeyArn) || dcl.IsZeroValue(des.KmsKeyArn) {
		cDes.KmsKeyArn = initial.KmsKeyArn
	} else {
		cDes.KmsKeyArn = des.KmsKeyArn
	}

	return cDes
}

func canonicalizeNewAwsClusterControlPlaneRootVolume(c *Client, des, nw *AwsClusterControlPlaneRootVolume) *AwsClusterControlPlaneRootVolume {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyArn, nw.KmsKeyArn) {
		nw.KmsKeyArn = des.KmsKeyArn
	}

	return nw
}

func canonicalizeNewAwsClusterControlPlaneRootVolumeSet(c *Client, des, nw []AwsClusterControlPlaneRootVolume) []AwsClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlaneRootVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneRootVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneRootVolumeSlice(c *Client, des, nw []AwsClusterControlPlaneRootVolume) []AwsClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlaneRootVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlaneRootVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlaneMainVolume(des, initial *AwsClusterControlPlaneMainVolume, opts ...dcl.ApplyOption) *AwsClusterControlPlaneMainVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlaneMainVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}
	if dcl.IsZeroValue(des.VolumeType) {
		des.VolumeType = initial.VolumeType
	} else {
		cDes.VolumeType = des.VolumeType
	}
	if dcl.IsZeroValue(des.Iops) {
		des.Iops = initial.Iops
	} else {
		cDes.Iops = des.Iops
	}
	if dcl.StringCanonicalize(des.KmsKeyArn, initial.KmsKeyArn) || dcl.IsZeroValue(des.KmsKeyArn) {
		cDes.KmsKeyArn = initial.KmsKeyArn
	} else {
		cDes.KmsKeyArn = des.KmsKeyArn
	}

	return cDes
}

func canonicalizeNewAwsClusterControlPlaneMainVolume(c *Client, des, nw *AwsClusterControlPlaneMainVolume) *AwsClusterControlPlaneMainVolume {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyArn, nw.KmsKeyArn) {
		nw.KmsKeyArn = des.KmsKeyArn
	}

	return nw
}

func canonicalizeNewAwsClusterControlPlaneMainVolumeSet(c *Client, des, nw []AwsClusterControlPlaneMainVolume) []AwsClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlaneMainVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneMainVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneMainVolumeSlice(c *Client, des, nw []AwsClusterControlPlaneMainVolume) []AwsClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlaneMainVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlaneMainVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlaneDatabaseEncryption(des, initial *AwsClusterControlPlaneDatabaseEncryption, opts ...dcl.ApplyOption) *AwsClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlaneDatabaseEncryption{}

	if dcl.StringCanonicalize(des.KmsKeyArn, initial.KmsKeyArn) || dcl.IsZeroValue(des.KmsKeyArn) {
		cDes.KmsKeyArn = initial.KmsKeyArn
	} else {
		cDes.KmsKeyArn = des.KmsKeyArn
	}

	return cDes
}

func canonicalizeNewAwsClusterControlPlaneDatabaseEncryption(c *Client, des, nw *AwsClusterControlPlaneDatabaseEncryption) *AwsClusterControlPlaneDatabaseEncryption {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyArn, nw.KmsKeyArn) {
		nw.KmsKeyArn = des.KmsKeyArn
	}

	return nw
}

func canonicalizeNewAwsClusterControlPlaneDatabaseEncryptionSet(c *Client, des, nw []AwsClusterControlPlaneDatabaseEncryption) []AwsClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlaneDatabaseEncryption
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneDatabaseEncryptionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneDatabaseEncryptionSlice(c *Client, des, nw []AwsClusterControlPlaneDatabaseEncryption) []AwsClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlaneDatabaseEncryption
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlaneDatabaseEncryption(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterControlPlaneAwsServicesAuthentication(des, initial *AwsClusterControlPlaneAwsServicesAuthentication, opts ...dcl.ApplyOption) *AwsClusterControlPlaneAwsServicesAuthentication {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterControlPlaneAwsServicesAuthentication{}

	if dcl.StringCanonicalize(des.RoleArn, initial.RoleArn) || dcl.IsZeroValue(des.RoleArn) {
		cDes.RoleArn = initial.RoleArn
	} else {
		cDes.RoleArn = des.RoleArn
	}
	if dcl.StringCanonicalize(des.RoleSessionName, initial.RoleSessionName) || dcl.IsZeroValue(des.RoleSessionName) {
		cDes.RoleSessionName = initial.RoleSessionName
	} else {
		cDes.RoleSessionName = des.RoleSessionName
	}

	return cDes
}

func canonicalizeNewAwsClusterControlPlaneAwsServicesAuthentication(c *Client, des, nw *AwsClusterControlPlaneAwsServicesAuthentication) *AwsClusterControlPlaneAwsServicesAuthentication {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RoleArn, nw.RoleArn) {
		nw.RoleArn = des.RoleArn
	}
	if dcl.StringCanonicalize(des.RoleSessionName, nw.RoleSessionName) {
		nw.RoleSessionName = des.RoleSessionName
	}

	return nw
}

func canonicalizeNewAwsClusterControlPlaneAwsServicesAuthenticationSet(c *Client, des, nw []AwsClusterControlPlaneAwsServicesAuthentication) []AwsClusterControlPlaneAwsServicesAuthentication {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterControlPlaneAwsServicesAuthentication
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterControlPlaneAwsServicesAuthenticationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterControlPlaneAwsServicesAuthenticationSlice(c *Client, des, nw []AwsClusterControlPlaneAwsServicesAuthentication) []AwsClusterControlPlaneAwsServicesAuthentication {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterControlPlaneAwsServicesAuthentication
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterControlPlaneAwsServicesAuthentication(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterAuthorization(des, initial *AwsClusterAuthorization, opts ...dcl.ApplyOption) *AwsClusterAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterAuthorization{}

	if dcl.IsZeroValue(des.AdminUsers) {
		des.AdminUsers = initial.AdminUsers
	} else {
		cDes.AdminUsers = des.AdminUsers
	}

	return cDes
}

func canonicalizeNewAwsClusterAuthorization(c *Client, des, nw *AwsClusterAuthorization) *AwsClusterAuthorization {
	if des == nil || nw == nil {
		return nw
	}

	nw.AdminUsers = canonicalizeNewAwsClusterAuthorizationAdminUsersSlice(c, des.AdminUsers, nw.AdminUsers)

	return nw
}

func canonicalizeNewAwsClusterAuthorizationSet(c *Client, des, nw []AwsClusterAuthorization) []AwsClusterAuthorization {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterAuthorization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterAuthorizationSlice(c *Client, des, nw []AwsClusterAuthorization) []AwsClusterAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterAuthorizationAdminUsers(des, initial *AwsClusterAuthorizationAdminUsers, opts ...dcl.ApplyOption) *AwsClusterAuthorizationAdminUsers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterAuthorizationAdminUsers{}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		cDes.Username = initial.Username
	} else {
		cDes.Username = des.Username
	}

	return cDes
}

func canonicalizeNewAwsClusterAuthorizationAdminUsers(c *Client, des, nw *AwsClusterAuthorizationAdminUsers) *AwsClusterAuthorizationAdminUsers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}

	return nw
}

func canonicalizeNewAwsClusterAuthorizationAdminUsersSet(c *Client, des, nw []AwsClusterAuthorizationAdminUsers) []AwsClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterAuthorizationAdminUsers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterAuthorizationAdminUsersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterAuthorizationAdminUsersSlice(c *Client, des, nw []AwsClusterAuthorizationAdminUsers) []AwsClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterAuthorizationAdminUsers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterAuthorizationAdminUsers(c, &d, &n))
	}

	return items
}

func canonicalizeAwsClusterWorkloadIdentityConfig(des, initial *AwsClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) *AwsClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsClusterWorkloadIdentityConfig{}

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

func canonicalizeNewAwsClusterWorkloadIdentityConfig(c *Client, des, nw *AwsClusterWorkloadIdentityConfig) *AwsClusterWorkloadIdentityConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeNewAwsClusterWorkloadIdentityConfigSet(c *Client, des, nw []AwsClusterWorkloadIdentityConfig) []AwsClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsClusterWorkloadIdentityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsClusterWorkloadIdentityConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsClusterWorkloadIdentityConfigSlice(c *Client, des, nw []AwsClusterWorkloadIdentityConfig) []AwsClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsClusterWorkloadIdentityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsClusterWorkloadIdentityConfig(c, &d, &n))
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
func diffAwsCluster(c *Client, desired, actual *AwsCluster, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Networking, actual.Networking, dcl.Info{ObjectFunction: compareAwsClusterNetworkingNewStyle, EmptyObject: EmptyAwsClusterNetworking, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Networking")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AwsRegion, actual.AwsRegion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AwsRegion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ControlPlane, actual.ControlPlane, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneNewStyle, EmptyObject: EmptyAwsClusterControlPlane, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ControlPlane")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Authorization, actual.Authorization, dcl.Info{ObjectFunction: compareAwsClusterAuthorizationNewStyle, EmptyObject: EmptyAwsClusterAuthorization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Authorization")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.WorkloadIdentityConfig, actual.WorkloadIdentityConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareAwsClusterWorkloadIdentityConfigNewStyle, EmptyObject: EmptyAwsClusterWorkloadIdentityConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityConfig")); len(ds) != 0 || err != nil {
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
func compareAwsClusterNetworkingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterNetworking)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterNetworking or *AwsClusterNetworking", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterNetworking)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterNetworking", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VPCId, actual.VPCId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VpcId")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ServiceLoadBalancerSubnetIds, actual.ServiceLoadBalancerSubnetIds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceLoadBalancerSubnetIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlane)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlane or *AwsClusterControlPlane", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlane)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlane", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceType, actual.InstanceType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SshConfig, actual.SshConfig, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneSshConfigNewStyle, EmptyObject: EmptyAwsClusterControlPlaneSshConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SshConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetIds, actual.SubnetIds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityGroupIds, actual.SecurityGroupIds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityGroupIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IamInstanceProfile, actual.IamInstanceProfile, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IamInstanceProfile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootVolume, actual.RootVolume, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneRootVolumeNewStyle, EmptyObject: EmptyAwsClusterControlPlaneRootVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MainVolume, actual.MainVolume, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneMainVolumeNewStyle, EmptyObject: EmptyAwsClusterControlPlaneMainVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MainVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatabaseEncryption, actual.DatabaseEncryption, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneDatabaseEncryptionNewStyle, EmptyObject: EmptyAwsClusterControlPlaneDatabaseEncryption, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatabaseEncryption")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.AwsServicesAuthentication, actual.AwsServicesAuthentication, dcl.Info{ObjectFunction: compareAwsClusterControlPlaneAwsServicesAuthenticationNewStyle, EmptyObject: EmptyAwsClusterControlPlaneAwsServicesAuthentication, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AwsServicesAuthentication")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneSshConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlaneSshConfig)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneSshConfig or *AwsClusterControlPlaneSshConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlaneSshConfig)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneSshConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Ec2KeyPair, actual.Ec2KeyPair, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Ec2KeyPair")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneRootVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlaneRootVolume)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneRootVolume or *AwsClusterControlPlaneRootVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlaneRootVolume)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneRootVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeType, actual.VolumeType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VolumeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Iops, actual.Iops, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Iops")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyArn, actual.KmsKeyArn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyArn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneMainVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlaneMainVolume)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneMainVolume or *AwsClusterControlPlaneMainVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlaneMainVolume)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneMainVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeType, actual.VolumeType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VolumeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Iops, actual.Iops, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Iops")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyArn, actual.KmsKeyArn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyArn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneDatabaseEncryptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlaneDatabaseEncryption)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneDatabaseEncryption or *AwsClusterControlPlaneDatabaseEncryption", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlaneDatabaseEncryption)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneDatabaseEncryption", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyArn, actual.KmsKeyArn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyArn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterControlPlaneAwsServicesAuthenticationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterControlPlaneAwsServicesAuthentication)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterControlPlaneAwsServicesAuthentication)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneAwsServicesAuthentication or *AwsClusterControlPlaneAwsServicesAuthentication", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterControlPlaneAwsServicesAuthentication)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterControlPlaneAwsServicesAuthentication)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterControlPlaneAwsServicesAuthentication", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RoleArn, actual.RoleArn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RoleArn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RoleSessionName, actual.RoleSessionName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RoleSessionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterAuthorization or *AwsClusterAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterAuthorization)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AdminUsers, actual.AdminUsers, dcl.Info{ObjectFunction: compareAwsClusterAuthorizationAdminUsersNewStyle, EmptyObject: EmptyAwsClusterAuthorizationAdminUsers, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdminUsers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsClusterAuthorizationAdminUsersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterAuthorizationAdminUsers)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterAuthorizationAdminUsers or *AwsClusterAuthorizationAdminUsers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterAuthorizationAdminUsers)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterAuthorizationAdminUsers", a)
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

func compareAwsClusterWorkloadIdentityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsClusterWorkloadIdentityConfig)
	if !ok {
		desiredNotPointer, ok := d.(AwsClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterWorkloadIdentityConfig or *AwsClusterWorkloadIdentityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsClusterWorkloadIdentityConfig)
	if !ok {
		actualNotPointer, ok := a.(AwsClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsClusterWorkloadIdentityConfig", a)
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
func (r *AwsCluster) urlNormalized() *AwsCluster {
	normalized := dcl.Copy(*r).(AwsCluster)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AwsRegion = dcl.SelfLinkToName(r.AwsRegion)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AwsCluster) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AwsCluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AwsCluster) marshal(c *Client) ([]byte, error) {
	m, err := expandAwsCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AwsCluster: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAwsCluster decodes JSON responses into the AwsCluster resource schema.
func unmarshalAwsCluster(b []byte, c *Client) (*AwsCluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAwsCluster(m, c)
}

func unmarshalMapAwsCluster(m map[string]interface{}, c *Client) (*AwsCluster, error) {

	return flattenAwsCluster(c, m), nil
}

// expandAwsCluster expands AwsCluster into a JSON request object.
func expandAwsCluster(c *Client, f *AwsCluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/awsClusters/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandAwsClusterNetworking(c, f.Networking); err != nil {
		return nil, fmt.Errorf("error expanding Networking into networking: %w", err)
	} else if v != nil {
		m["networking"] = v
	}
	if v := f.AwsRegion; !dcl.IsEmptyValueIndirect(v) {
		m["awsRegion"] = v
	}
	if v, err := expandAwsClusterControlPlane(c, f.ControlPlane); err != nil {
		return nil, fmt.Errorf("error expanding ControlPlane into controlPlane: %w", err)
	} else if v != nil {
		m["controlPlane"] = v
	}
	if v, err := expandAwsClusterAuthorization(c, f.Authorization); err != nil {
		return nil, fmt.Errorf("error expanding Authorization into authorization: %w", err)
	} else if v != nil {
		m["authorization"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Endpoint; !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.Reconciling; !dcl.IsEmptyValueIndirect(v) {
		m["reconciling"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandAwsClusterWorkloadIdentityConfig(c, f.WorkloadIdentityConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkloadIdentityConfig into workloadIdentityConfig: %w", err)
	} else if v != nil {
		m["workloadIdentityConfig"] = v
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

// flattenAwsCluster flattens AwsCluster from a JSON request object into the
// AwsCluster type.
func flattenAwsCluster(c *Client, i interface{}) *AwsCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AwsCluster{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Networking = flattenAwsClusterNetworking(c, m["networking"])
	res.AwsRegion = dcl.FlattenString(m["awsRegion"])
	res.ControlPlane = flattenAwsClusterControlPlane(c, m["controlPlane"])
	res.Authorization = flattenAwsClusterAuthorization(c, m["authorization"])
	res.State = flattenAwsClusterStateEnum(m["state"])
	res.Endpoint = dcl.FlattenString(m["endpoint"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.WorkloadIdentityConfig = flattenAwsClusterWorkloadIdentityConfig(c, m["workloadIdentityConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandAwsClusterNetworkingMap expands the contents of AwsClusterNetworking into a JSON
// request object.
func expandAwsClusterNetworkingMap(c *Client, f map[string]AwsClusterNetworking) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterNetworkingSlice expands the contents of AwsClusterNetworking into a JSON
// request object.
func expandAwsClusterNetworkingSlice(c *Client, f []AwsClusterNetworking) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterNetworkingMap flattens the contents of AwsClusterNetworking from a JSON
// response object.
func flattenAwsClusterNetworkingMap(c *Client, i interface{}) map[string]AwsClusterNetworking {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterNetworking{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterNetworking{}
	}

	items := make(map[string]AwsClusterNetworking)
	for k, item := range a {
		items[k] = *flattenAwsClusterNetworking(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterNetworkingSlice flattens the contents of AwsClusterNetworking from a JSON
// response object.
func flattenAwsClusterNetworkingSlice(c *Client, i interface{}) []AwsClusterNetworking {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterNetworking{}
	}

	if len(a) == 0 {
		return []AwsClusterNetworking{}
	}

	items := make([]AwsClusterNetworking, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterNetworking(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterNetworking expands an instance of AwsClusterNetworking into a JSON
// request object.
func expandAwsClusterNetworking(c *Client, f *AwsClusterNetworking) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.VPCId; !dcl.IsEmptyValueIndirect(v) {
		m["vpcId"] = v
	}
	if v := f.PodAddressCidrBlocks; v != nil {
		m["podAddressCidrBlocks"] = v
	}
	if v := f.ServiceAddressCidrBlocks; v != nil {
		m["serviceAddressCidrBlocks"] = v
	}
	if v := f.ServiceLoadBalancerSubnetIds; v != nil {
		m["serviceLoadBalancerSubnetIds"] = v
	}

	return m, nil
}

// flattenAwsClusterNetworking flattens an instance of AwsClusterNetworking from a JSON
// response object.
func flattenAwsClusterNetworking(c *Client, i interface{}) *AwsClusterNetworking {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterNetworking{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterNetworking
	}
	r.VPCId = dcl.FlattenString(m["vpcId"])
	r.PodAddressCidrBlocks = dcl.FlattenStringSlice(m["podAddressCidrBlocks"])
	r.ServiceAddressCidrBlocks = dcl.FlattenStringSlice(m["serviceAddressCidrBlocks"])
	r.ServiceLoadBalancerSubnetIds = dcl.FlattenStringSlice(m["serviceLoadBalancerSubnetIds"])

	return r
}

// expandAwsClusterControlPlaneMap expands the contents of AwsClusterControlPlane into a JSON
// request object.
func expandAwsClusterControlPlaneMap(c *Client, f map[string]AwsClusterControlPlane) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneSlice expands the contents of AwsClusterControlPlane into a JSON
// request object.
func expandAwsClusterControlPlaneSlice(c *Client, f []AwsClusterControlPlane) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneMap flattens the contents of AwsClusterControlPlane from a JSON
// response object.
func flattenAwsClusterControlPlaneMap(c *Client, i interface{}) map[string]AwsClusterControlPlane {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlane{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlane{}
	}

	items := make(map[string]AwsClusterControlPlane)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlane(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneSlice flattens the contents of AwsClusterControlPlane from a JSON
// response object.
func flattenAwsClusterControlPlaneSlice(c *Client, i interface{}) []AwsClusterControlPlane {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlane{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlane{}
	}

	items := make([]AwsClusterControlPlane, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlane(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlane expands an instance of AwsClusterControlPlane into a JSON
// request object.
func expandAwsClusterControlPlane(c *Client, f *AwsClusterControlPlane) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.InstanceType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceType"] = v
	}
	if v, err := expandAwsClusterControlPlaneSshConfig(c, f.SshConfig); err != nil {
		return nil, fmt.Errorf("error expanding SshConfig into sshConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sshConfig"] = v
	}
	if v := f.SubnetIds; v != nil {
		m["subnetIds"] = v
	}
	if v := f.SecurityGroupIds; v != nil {
		m["securityGroupIds"] = v
	}
	if v := f.IamInstanceProfile; !dcl.IsEmptyValueIndirect(v) {
		m["iamInstanceProfile"] = v
	}
	if v, err := expandAwsClusterControlPlaneRootVolume(c, f.RootVolume); err != nil {
		return nil, fmt.Errorf("error expanding RootVolume into rootVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rootVolume"] = v
	}
	if v, err := expandAwsClusterControlPlaneMainVolume(c, f.MainVolume); err != nil {
		return nil, fmt.Errorf("error expanding MainVolume into mainVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mainVolume"] = v
	}
	if v, err := expandAwsClusterControlPlaneDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into databaseEncryption: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseEncryption"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v, err := expandAwsClusterControlPlaneAwsServicesAuthentication(c, f.AwsServicesAuthentication); err != nil {
		return nil, fmt.Errorf("error expanding AwsServicesAuthentication into awsServicesAuthentication: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["awsServicesAuthentication"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlane flattens an instance of AwsClusterControlPlane from a JSON
// response object.
func flattenAwsClusterControlPlane(c *Client, i interface{}) *AwsClusterControlPlane {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlane{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlane
	}
	r.Version = dcl.FlattenString(m["version"])
	r.InstanceType = dcl.FlattenString(m["instanceType"])
	r.SshConfig = flattenAwsClusterControlPlaneSshConfig(c, m["sshConfig"])
	r.SubnetIds = dcl.FlattenStringSlice(m["subnetIds"])
	r.SecurityGroupIds = dcl.FlattenStringSlice(m["securityGroupIds"])
	r.IamInstanceProfile = dcl.FlattenString(m["iamInstanceProfile"])
	r.RootVolume = flattenAwsClusterControlPlaneRootVolume(c, m["rootVolume"])
	r.MainVolume = flattenAwsClusterControlPlaneMainVolume(c, m["mainVolume"])
	r.DatabaseEncryption = flattenAwsClusterControlPlaneDatabaseEncryption(c, m["databaseEncryption"])
	r.Tags = dcl.FlattenKeyValuePairs(m["tags"])
	r.AwsServicesAuthentication = flattenAwsClusterControlPlaneAwsServicesAuthentication(c, m["awsServicesAuthentication"])

	return r
}

// expandAwsClusterControlPlaneSshConfigMap expands the contents of AwsClusterControlPlaneSshConfig into a JSON
// request object.
func expandAwsClusterControlPlaneSshConfigMap(c *Client, f map[string]AwsClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneSshConfigSlice expands the contents of AwsClusterControlPlaneSshConfig into a JSON
// request object.
func expandAwsClusterControlPlaneSshConfigSlice(c *Client, f []AwsClusterControlPlaneSshConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneSshConfigMap flattens the contents of AwsClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAwsClusterControlPlaneSshConfigMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneSshConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneSshConfig{}
	}

	items := make(map[string]AwsClusterControlPlaneSshConfig)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneSshConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneSshConfigSlice flattens the contents of AwsClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAwsClusterControlPlaneSshConfigSlice(c *Client, i interface{}) []AwsClusterControlPlaneSshConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneSshConfig{}
	}

	items := make([]AwsClusterControlPlaneSshConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneSshConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlaneSshConfig expands an instance of AwsClusterControlPlaneSshConfig into a JSON
// request object.
func expandAwsClusterControlPlaneSshConfig(c *Client, f *AwsClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Ec2KeyPair; !dcl.IsEmptyValueIndirect(v) {
		m["ec2KeyPair"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlaneSshConfig flattens an instance of AwsClusterControlPlaneSshConfig from a JSON
// response object.
func flattenAwsClusterControlPlaneSshConfig(c *Client, i interface{}) *AwsClusterControlPlaneSshConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlaneSshConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlaneSshConfig
	}
	r.Ec2KeyPair = dcl.FlattenString(m["ec2KeyPair"])

	return r
}

// expandAwsClusterControlPlaneRootVolumeMap expands the contents of AwsClusterControlPlaneRootVolume into a JSON
// request object.
func expandAwsClusterControlPlaneRootVolumeMap(c *Client, f map[string]AwsClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneRootVolumeSlice expands the contents of AwsClusterControlPlaneRootVolume into a JSON
// request object.
func expandAwsClusterControlPlaneRootVolumeSlice(c *Client, f []AwsClusterControlPlaneRootVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneRootVolumeMap flattens the contents of AwsClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneRootVolumeMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneRootVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneRootVolume{}
	}

	items := make(map[string]AwsClusterControlPlaneRootVolume)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneRootVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneRootVolumeSlice flattens the contents of AwsClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneRootVolumeSlice(c *Client, i interface{}) []AwsClusterControlPlaneRootVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneRootVolume{}
	}

	items := make([]AwsClusterControlPlaneRootVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneRootVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlaneRootVolume expands an instance of AwsClusterControlPlaneRootVolume into a JSON
// request object.
func expandAwsClusterControlPlaneRootVolume(c *Client, f *AwsClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}
	if v := f.VolumeType; !dcl.IsEmptyValueIndirect(v) {
		m["volumeType"] = v
	}
	if v := f.Iops; !dcl.IsEmptyValueIndirect(v) {
		m["iops"] = v
	}
	if v := f.KmsKeyArn; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyArn"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlaneRootVolume flattens an instance of AwsClusterControlPlaneRootVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneRootVolume(c *Client, i interface{}) *AwsClusterControlPlaneRootVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlaneRootVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlaneRootVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])
	r.VolumeType = flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnum(m["volumeType"])
	r.Iops = dcl.FlattenInteger(m["iops"])
	r.KmsKeyArn = dcl.FlattenString(m["kmsKeyArn"])

	return r
}

// expandAwsClusterControlPlaneMainVolumeMap expands the contents of AwsClusterControlPlaneMainVolume into a JSON
// request object.
func expandAwsClusterControlPlaneMainVolumeMap(c *Client, f map[string]AwsClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneMainVolumeSlice expands the contents of AwsClusterControlPlaneMainVolume into a JSON
// request object.
func expandAwsClusterControlPlaneMainVolumeSlice(c *Client, f []AwsClusterControlPlaneMainVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneMainVolumeMap flattens the contents of AwsClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneMainVolumeMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneMainVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneMainVolume{}
	}

	items := make(map[string]AwsClusterControlPlaneMainVolume)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneMainVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneMainVolumeSlice flattens the contents of AwsClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneMainVolumeSlice(c *Client, i interface{}) []AwsClusterControlPlaneMainVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneMainVolume{}
	}

	items := make([]AwsClusterControlPlaneMainVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneMainVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlaneMainVolume expands an instance of AwsClusterControlPlaneMainVolume into a JSON
// request object.
func expandAwsClusterControlPlaneMainVolume(c *Client, f *AwsClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}
	if v := f.VolumeType; !dcl.IsEmptyValueIndirect(v) {
		m["volumeType"] = v
	}
	if v := f.Iops; !dcl.IsEmptyValueIndirect(v) {
		m["iops"] = v
	}
	if v := f.KmsKeyArn; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyArn"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlaneMainVolume flattens an instance of AwsClusterControlPlaneMainVolume from a JSON
// response object.
func flattenAwsClusterControlPlaneMainVolume(c *Client, i interface{}) *AwsClusterControlPlaneMainVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlaneMainVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlaneMainVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])
	r.VolumeType = flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnum(m["volumeType"])
	r.Iops = dcl.FlattenInteger(m["iops"])
	r.KmsKeyArn = dcl.FlattenString(m["kmsKeyArn"])

	return r
}

// expandAwsClusterControlPlaneDatabaseEncryptionMap expands the contents of AwsClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAwsClusterControlPlaneDatabaseEncryptionMap(c *Client, f map[string]AwsClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneDatabaseEncryptionSlice expands the contents of AwsClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAwsClusterControlPlaneDatabaseEncryptionSlice(c *Client, f []AwsClusterControlPlaneDatabaseEncryption) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneDatabaseEncryptionMap flattens the contents of AwsClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAwsClusterControlPlaneDatabaseEncryptionMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneDatabaseEncryption {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneDatabaseEncryption{}
	}

	items := make(map[string]AwsClusterControlPlaneDatabaseEncryption)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneDatabaseEncryptionSlice flattens the contents of AwsClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAwsClusterControlPlaneDatabaseEncryptionSlice(c *Client, i interface{}) []AwsClusterControlPlaneDatabaseEncryption {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneDatabaseEncryption{}
	}

	items := make([]AwsClusterControlPlaneDatabaseEncryption, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlaneDatabaseEncryption expands an instance of AwsClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandAwsClusterControlPlaneDatabaseEncryption(c *Client, f *AwsClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyArn; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyArn"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlaneDatabaseEncryption flattens an instance of AwsClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenAwsClusterControlPlaneDatabaseEncryption(c *Client, i interface{}) *AwsClusterControlPlaneDatabaseEncryption {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlaneDatabaseEncryption{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlaneDatabaseEncryption
	}
	r.KmsKeyArn = dcl.FlattenString(m["kmsKeyArn"])

	return r
}

// expandAwsClusterControlPlaneAwsServicesAuthenticationMap expands the contents of AwsClusterControlPlaneAwsServicesAuthentication into a JSON
// request object.
func expandAwsClusterControlPlaneAwsServicesAuthenticationMap(c *Client, f map[string]AwsClusterControlPlaneAwsServicesAuthentication) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterControlPlaneAwsServicesAuthentication(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterControlPlaneAwsServicesAuthenticationSlice expands the contents of AwsClusterControlPlaneAwsServicesAuthentication into a JSON
// request object.
func expandAwsClusterControlPlaneAwsServicesAuthenticationSlice(c *Client, f []AwsClusterControlPlaneAwsServicesAuthentication) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterControlPlaneAwsServicesAuthentication(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterControlPlaneAwsServicesAuthenticationMap flattens the contents of AwsClusterControlPlaneAwsServicesAuthentication from a JSON
// response object.
func flattenAwsClusterControlPlaneAwsServicesAuthenticationMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneAwsServicesAuthentication {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneAwsServicesAuthentication{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneAwsServicesAuthentication{}
	}

	items := make(map[string]AwsClusterControlPlaneAwsServicesAuthentication)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneAwsServicesAuthentication(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneAwsServicesAuthenticationSlice flattens the contents of AwsClusterControlPlaneAwsServicesAuthentication from a JSON
// response object.
func flattenAwsClusterControlPlaneAwsServicesAuthenticationSlice(c *Client, i interface{}) []AwsClusterControlPlaneAwsServicesAuthentication {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneAwsServicesAuthentication{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneAwsServicesAuthentication{}
	}

	items := make([]AwsClusterControlPlaneAwsServicesAuthentication, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneAwsServicesAuthentication(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterControlPlaneAwsServicesAuthentication expands an instance of AwsClusterControlPlaneAwsServicesAuthentication into a JSON
// request object.
func expandAwsClusterControlPlaneAwsServicesAuthentication(c *Client, f *AwsClusterControlPlaneAwsServicesAuthentication) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RoleArn; !dcl.IsEmptyValueIndirect(v) {
		m["roleArn"] = v
	}
	if v := f.RoleSessionName; !dcl.IsEmptyValueIndirect(v) {
		m["roleSessionName"] = v
	}

	return m, nil
}

// flattenAwsClusterControlPlaneAwsServicesAuthentication flattens an instance of AwsClusterControlPlaneAwsServicesAuthentication from a JSON
// response object.
func flattenAwsClusterControlPlaneAwsServicesAuthentication(c *Client, i interface{}) *AwsClusterControlPlaneAwsServicesAuthentication {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterControlPlaneAwsServicesAuthentication{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterControlPlaneAwsServicesAuthentication
	}
	r.RoleArn = dcl.FlattenString(m["roleArn"])
	r.RoleSessionName = dcl.FlattenString(m["roleSessionName"])

	return r
}

// expandAwsClusterAuthorizationMap expands the contents of AwsClusterAuthorization into a JSON
// request object.
func expandAwsClusterAuthorizationMap(c *Client, f map[string]AwsClusterAuthorization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterAuthorizationSlice expands the contents of AwsClusterAuthorization into a JSON
// request object.
func expandAwsClusterAuthorizationSlice(c *Client, f []AwsClusterAuthorization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterAuthorizationMap flattens the contents of AwsClusterAuthorization from a JSON
// response object.
func flattenAwsClusterAuthorizationMap(c *Client, i interface{}) map[string]AwsClusterAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterAuthorization{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterAuthorization{}
	}

	items := make(map[string]AwsClusterAuthorization)
	for k, item := range a {
		items[k] = *flattenAwsClusterAuthorization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterAuthorizationSlice flattens the contents of AwsClusterAuthorization from a JSON
// response object.
func flattenAwsClusterAuthorizationSlice(c *Client, i interface{}) []AwsClusterAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterAuthorization{}
	}

	if len(a) == 0 {
		return []AwsClusterAuthorization{}
	}

	items := make([]AwsClusterAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterAuthorization(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterAuthorization expands an instance of AwsClusterAuthorization into a JSON
// request object.
func expandAwsClusterAuthorization(c *Client, f *AwsClusterAuthorization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAwsClusterAuthorizationAdminUsersSlice(c, f.AdminUsers); err != nil {
		return nil, fmt.Errorf("error expanding AdminUsers into adminUsers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["adminUsers"] = v
	}

	return m, nil
}

// flattenAwsClusterAuthorization flattens an instance of AwsClusterAuthorization from a JSON
// response object.
func flattenAwsClusterAuthorization(c *Client, i interface{}) *AwsClusterAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterAuthorization
	}
	r.AdminUsers = flattenAwsClusterAuthorizationAdminUsersSlice(c, m["adminUsers"])

	return r
}

// expandAwsClusterAuthorizationAdminUsersMap expands the contents of AwsClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAwsClusterAuthorizationAdminUsersMap(c *Client, f map[string]AwsClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterAuthorizationAdminUsersSlice expands the contents of AwsClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAwsClusterAuthorizationAdminUsersSlice(c *Client, f []AwsClusterAuthorizationAdminUsers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterAuthorizationAdminUsersMap flattens the contents of AwsClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAwsClusterAuthorizationAdminUsersMap(c *Client, i interface{}) map[string]AwsClusterAuthorizationAdminUsers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterAuthorizationAdminUsers{}
	}

	items := make(map[string]AwsClusterAuthorizationAdminUsers)
	for k, item := range a {
		items[k] = *flattenAwsClusterAuthorizationAdminUsers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterAuthorizationAdminUsersSlice flattens the contents of AwsClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAwsClusterAuthorizationAdminUsersSlice(c *Client, i interface{}) []AwsClusterAuthorizationAdminUsers {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return []AwsClusterAuthorizationAdminUsers{}
	}

	items := make([]AwsClusterAuthorizationAdminUsers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterAuthorizationAdminUsers(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterAuthorizationAdminUsers expands an instance of AwsClusterAuthorizationAdminUsers into a JSON
// request object.
func expandAwsClusterAuthorizationAdminUsers(c *Client, f *AwsClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}

	return m, nil
}

// flattenAwsClusterAuthorizationAdminUsers flattens an instance of AwsClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenAwsClusterAuthorizationAdminUsers(c *Client, i interface{}) *AwsClusterAuthorizationAdminUsers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterAuthorizationAdminUsers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterAuthorizationAdminUsers
	}
	r.Username = dcl.FlattenString(m["username"])

	return r
}

// expandAwsClusterWorkloadIdentityConfigMap expands the contents of AwsClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAwsClusterWorkloadIdentityConfigMap(c *Client, f map[string]AwsClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsClusterWorkloadIdentityConfigSlice expands the contents of AwsClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAwsClusterWorkloadIdentityConfigSlice(c *Client, f []AwsClusterWorkloadIdentityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsClusterWorkloadIdentityConfigMap flattens the contents of AwsClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAwsClusterWorkloadIdentityConfigMap(c *Client, i interface{}) map[string]AwsClusterWorkloadIdentityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterWorkloadIdentityConfig{}
	}

	items := make(map[string]AwsClusterWorkloadIdentityConfig)
	for k, item := range a {
		items[k] = *flattenAwsClusterWorkloadIdentityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsClusterWorkloadIdentityConfigSlice flattens the contents of AwsClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAwsClusterWorkloadIdentityConfigSlice(c *Client, i interface{}) []AwsClusterWorkloadIdentityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return []AwsClusterWorkloadIdentityConfig{}
	}

	items := make([]AwsClusterWorkloadIdentityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterWorkloadIdentityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsClusterWorkloadIdentityConfig expands an instance of AwsClusterWorkloadIdentityConfig into a JSON
// request object.
func expandAwsClusterWorkloadIdentityConfig(c *Client, f *AwsClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
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

// flattenAwsClusterWorkloadIdentityConfig flattens an instance of AwsClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenAwsClusterWorkloadIdentityConfig(c *Client, i interface{}) *AwsClusterWorkloadIdentityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsClusterWorkloadIdentityConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsClusterWorkloadIdentityConfig
	}
	r.IssuerUri = dcl.FlattenString(m["issuerUri"])
	r.WorkloadPool = dcl.FlattenString(m["workloadPool"])
	r.IdentityProvider = dcl.FlattenString(m["identityProvider"])

	return r
}

// flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnumMap flattens the contents of AwsClusterControlPlaneRootVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnumMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneRootVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneRootVolumeVolumeTypeEnum{}
	}

	items := make(map[string]AwsClusterControlPlaneRootVolumeVolumeTypeEnum)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnum(item.(interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnumSlice flattens the contents of AwsClusterControlPlaneRootVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnumSlice(c *Client, i interface{}) []AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneRootVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneRootVolumeVolumeTypeEnum{}
	}

	items := make([]AwsClusterControlPlaneRootVolumeVolumeTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AwsClusterControlPlaneRootVolumeVolumeTypeEnum with the same value as that string.
func flattenAwsClusterControlPlaneRootVolumeVolumeTypeEnum(i interface{}) *AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AwsClusterControlPlaneRootVolumeVolumeTypeEnumRef("")
	}

	return AwsClusterControlPlaneRootVolumeVolumeTypeEnumRef(s)
}

// flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnumMap flattens the contents of AwsClusterControlPlaneMainVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnumMap(c *Client, i interface{}) map[string]AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterControlPlaneMainVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterControlPlaneMainVolumeVolumeTypeEnum{}
	}

	items := make(map[string]AwsClusterControlPlaneMainVolumeVolumeTypeEnum)
	for k, item := range a {
		items[k] = *flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnum(item.(interface{}))
	}

	return items
}

// flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnumSlice flattens the contents of AwsClusterControlPlaneMainVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnumSlice(c *Client, i interface{}) []AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterControlPlaneMainVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return []AwsClusterControlPlaneMainVolumeVolumeTypeEnum{}
	}

	items := make([]AwsClusterControlPlaneMainVolumeVolumeTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AwsClusterControlPlaneMainVolumeVolumeTypeEnum with the same value as that string.
func flattenAwsClusterControlPlaneMainVolumeVolumeTypeEnum(i interface{}) *AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AwsClusterControlPlaneMainVolumeVolumeTypeEnumRef("")
	}

	return AwsClusterControlPlaneMainVolumeVolumeTypeEnumRef(s)
}

// flattenAwsClusterStateEnumMap flattens the contents of AwsClusterStateEnum from a JSON
// response object.
func flattenAwsClusterStateEnumMap(c *Client, i interface{}) map[string]AwsClusterStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsClusterStateEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsClusterStateEnum{}
	}

	items := make(map[string]AwsClusterStateEnum)
	for k, item := range a {
		items[k] = *flattenAwsClusterStateEnum(item.(interface{}))
	}

	return items
}

// flattenAwsClusterStateEnumSlice flattens the contents of AwsClusterStateEnum from a JSON
// response object.
func flattenAwsClusterStateEnumSlice(c *Client, i interface{}) []AwsClusterStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsClusterStateEnum{}
	}

	if len(a) == 0 {
		return []AwsClusterStateEnum{}
	}

	items := make([]AwsClusterStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsClusterStateEnum(item.(interface{})))
	}

	return items
}

// flattenAwsClusterStateEnum asserts that an interface is a string, and returns a
// pointer to a *AwsClusterStateEnum with the same value as that string.
func flattenAwsClusterStateEnum(i interface{}) *AwsClusterStateEnum {
	s, ok := i.(string)
	if !ok {
		return AwsClusterStateEnumRef("")
	}

	return AwsClusterStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AwsCluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAwsCluster(b, c)
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

type awsClusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         awsClusterApiOperation
}

func convertFieldDiffsToAwsClusterDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]awsClusterDiff, error) {
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
	var diffs []awsClusterDiff
	// For each operation name, create a awsClusterDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := awsClusterDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAwsClusterApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAwsClusterApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (awsClusterApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
