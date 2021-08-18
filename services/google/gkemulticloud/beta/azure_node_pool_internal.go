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

func (r *AzureNodePool) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "version"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if err := dcl.Required(r, "subnetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "autoscaling"); err != nil {
		return err
	}
	if err := dcl.Required(r, "maxPodsConstraint"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.AzureCluster, "AzureCluster"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Autoscaling) {
		if err := r.Autoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxPodsConstraint) {
		if err := r.MaxPodsConstraint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AzureNodePoolConfig) validate() error {
	if err := dcl.Required(r, "sshConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.RootVolume) {
		if err := r.RootVolume.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SshConfig) {
		if err := r.SshConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AzureNodePoolConfigRootVolume) validate() error {
	return nil
}
func (r *AzureNodePoolConfigSshConfig) validate() error {
	if err := dcl.Required(r, "authorizedKey"); err != nil {
		return err
	}
	return nil
}
func (r *AzureNodePoolAutoscaling) validate() error {
	if err := dcl.Required(r, "minNodeCount"); err != nil {
		return err
	}
	if err := dcl.Required(r, "maxNodeCount"); err != nil {
		return err
	}
	return nil
}
func (r *AzureNodePoolMaxPodsConstraint) validate() error {
	if err := dcl.Required(r, "maxPodsPerNode"); err != nil {
		return err
	}
	return nil
}
func (r *AzureNodePool) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-gkemulticloud.googleapis.com/v1", params)
}

func (r *AzureNodePool) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"azureCluster": dcl.ValueOrEmptyString(nr.AzureCluster),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{azureCluster}}/azureNodePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AzureNodePool) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"azureCluster": dcl.ValueOrEmptyString(nr.AzureCluster),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{azureCluster}}/azureNodePools", nr.basePath(), userBasePath, params), nil

}

func (r *AzureNodePool) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"azureCluster": dcl.ValueOrEmptyString(nr.AzureCluster),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{azureCluster}}/azureNodePools?azureNodePoolId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AzureNodePool) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"azureCluster": dcl.ValueOrEmptyString(nr.AzureCluster),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{azureCluster}}/azureNodePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

// azureNodePoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type azureNodePoolApiOperation interface {
	do(context.Context, *AzureNodePool, *Client) error
}

func (c *Client) listAzureNodePoolRaw(ctx context.Context, r *AzureNodePool, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AzureNodePoolMaxPage {
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

type listAzureNodePoolOperation struct {
	AzureNodePools []map[string]interface{} `json:"azureNodePools"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listAzureNodePool(ctx context.Context, r *AzureNodePool, pageToken string, pageSize int32) ([]*AzureNodePool, string, error) {
	b, err := c.listAzureNodePoolRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAzureNodePoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AzureNodePool
	for _, v := range m.AzureNodePools {
		res, err := unmarshalMapAzureNodePool(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.AzureCluster = r.AzureCluster
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAzureNodePool(ctx context.Context, f func(*AzureNodePool) bool, resources []*AzureNodePool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAzureNodePool(ctx, res)
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

type deleteAzureNodePoolOperation struct{}

func (op *deleteAzureNodePoolOperation) do(ctx context.Context, r *AzureNodePool, c *Client) error {
	r, err := c.GetAzureNodePool(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "AzureNodePool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetAzureNodePool checking for existence. error: %v", err)
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
		_, err = c.GetAzureNodePool(ctx, r)
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
type createAzureNodePoolOperation struct {
	response map[string]interface{}
}

func (op *createAzureNodePoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAzureNodePoolOperation) do(ctx context.Context, r *AzureNodePool, c *Client) error {
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

	if _, err := c.GetAzureNodePool(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAzureNodePoolRaw(ctx context.Context, r *AzureNodePool) ([]byte, error) {

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

func (c *Client) azureNodePoolDiffsForRawDesired(ctx context.Context, rawDesired *AzureNodePool, opts ...dcl.ApplyOption) (initial, desired *AzureNodePool, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AzureNodePool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AzureNodePool); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected AzureNodePool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAzureNodePool(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a AzureNodePool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AzureNodePool resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that AzureNodePool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAzureNodePoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for AzureNodePool: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for AzureNodePool: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAzureNodePoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for AzureNodePool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAzureNodePoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for AzureNodePool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAzureNodePool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAzureNodePoolInitialState(rawInitial, rawDesired *AzureNodePool) (*AzureNodePool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAzureNodePoolDesiredState(rawDesired, rawInitial *AzureNodePool, opts ...dcl.ApplyOption) (*AzureNodePool, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeAzureNodePoolConfig(rawDesired.Config, nil, opts...)
		rawDesired.Autoscaling = canonicalizeAzureNodePoolAutoscaling(rawDesired.Autoscaling, nil, opts...)
		rawDesired.MaxPodsConstraint = canonicalizeAzureNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &AzureNodePool{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Version, rawInitial.Version) {
		canonicalDesired.Version = rawInitial.Version
	} else {
		canonicalDesired.Version = rawDesired.Version
	}
	canonicalDesired.Config = canonicalizeAzureNodePoolConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.StringCanonicalize(rawDesired.SubnetId, rawInitial.SubnetId) {
		canonicalDesired.SubnetId = rawInitial.SubnetId
	} else {
		canonicalDesired.SubnetId = rawDesired.SubnetId
	}
	canonicalDesired.Autoscaling = canonicalizeAzureNodePoolAutoscaling(rawDesired.Autoscaling, rawInitial.Autoscaling, opts...)
	if dcl.IsZeroValue(rawDesired.Annotations) {
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	canonicalDesired.MaxPodsConstraint = canonicalizeAzureNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, rawInitial.MaxPodsConstraint, opts...)
	if dcl.StringCanonicalize(rawDesired.AzureAvailabilityZone, rawInitial.AzureAvailabilityZone) {
		canonicalDesired.AzureAvailabilityZone = rawInitial.AzureAvailabilityZone
	} else {
		canonicalDesired.AzureAvailabilityZone = rawDesired.AzureAvailabilityZone
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
	if dcl.NameToSelfLink(rawDesired.AzureCluster, rawInitial.AzureCluster) {
		canonicalDesired.AzureCluster = rawInitial.AzureCluster
	} else {
		canonicalDesired.AzureCluster = rawDesired.AzureCluster
	}

	return canonicalDesired, nil
}

func canonicalizeAzureNodePoolNewState(c *Client, rawNew, rawDesired *AzureNodePool) (*AzureNodePool, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
	} else {
		if dcl.StringCanonicalize(rawDesired.Version, rawNew.Version) {
			rawNew.Version = rawDesired.Version
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewAzureNodePoolConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubnetId) && dcl.IsEmptyValueIndirect(rawDesired.SubnetId) {
		rawNew.SubnetId = rawDesired.SubnetId
	} else {
		if dcl.StringCanonicalize(rawDesired.SubnetId, rawNew.SubnetId) {
			rawNew.SubnetId = rawDesired.SubnetId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Autoscaling) && dcl.IsEmptyValueIndirect(rawDesired.Autoscaling) {
		rawNew.Autoscaling = rawDesired.Autoscaling
	} else {
		rawNew.Autoscaling = canonicalizeNewAzureNodePoolAutoscaling(c, rawDesired.Autoscaling, rawNew.Autoscaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
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

	if dcl.IsEmptyValueIndirect(rawNew.MaxPodsConstraint) && dcl.IsEmptyValueIndirect(rawDesired.MaxPodsConstraint) {
		rawNew.MaxPodsConstraint = rawDesired.MaxPodsConstraint
	} else {
		rawNew.MaxPodsConstraint = canonicalizeNewAzureNodePoolMaxPodsConstraint(c, rawDesired.MaxPodsConstraint, rawNew.MaxPodsConstraint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AzureAvailabilityZone) && dcl.IsEmptyValueIndirect(rawDesired.AzureAvailabilityZone) {
		rawNew.AzureAvailabilityZone = rawDesired.AzureAvailabilityZone
	} else {
		if dcl.StringCanonicalize(rawDesired.AzureAvailabilityZone, rawNew.AzureAvailabilityZone) {
			rawNew.AzureAvailabilityZone = rawDesired.AzureAvailabilityZone
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.AzureCluster = rawDesired.AzureCluster

	return rawNew, nil
}

func canonicalizeAzureNodePoolConfig(des, initial *AzureNodePoolConfig, opts ...dcl.ApplyOption) *AzureNodePoolConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureNodePoolConfig{}

	if dcl.StringCanonicalize(des.VmSize, initial.VmSize) || dcl.IsZeroValue(des.VmSize) {
		cDes.VmSize = initial.VmSize
	} else {
		cDes.VmSize = des.VmSize
	}
	cDes.RootVolume = canonicalizeAzureNodePoolConfigRootVolume(des.RootVolume, initial.RootVolume, opts...)
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	cDes.SshConfig = canonicalizeAzureNodePoolConfigSshConfig(des.SshConfig, initial.SshConfig, opts...)

	return cDes
}

func canonicalizeNewAzureNodePoolConfig(c *Client, des, nw *AzureNodePoolConfig) *AzureNodePoolConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.VmSize, nw.VmSize) {
		nw.VmSize = des.VmSize
	}
	nw.RootVolume = canonicalizeNewAzureNodePoolConfigRootVolume(c, des.RootVolume, nw.RootVolume)
	nw.SshConfig = canonicalizeNewAzureNodePoolConfigSshConfig(c, des.SshConfig, nw.SshConfig)

	return nw
}

func canonicalizeNewAzureNodePoolConfigSet(c *Client, des, nw []AzureNodePoolConfig) []AzureNodePoolConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureNodePoolConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureNodePoolConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureNodePoolConfigSlice(c *Client, des, nw []AzureNodePoolConfig) []AzureNodePoolConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureNodePoolConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureNodePoolConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAzureNodePoolConfigRootVolume(des, initial *AzureNodePoolConfigRootVolume, opts ...dcl.ApplyOption) *AzureNodePoolConfigRootVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureNodePoolConfigRootVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}

	return cDes
}

func canonicalizeNewAzureNodePoolConfigRootVolume(c *Client, des, nw *AzureNodePoolConfigRootVolume) *AzureNodePoolConfigRootVolume {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAzureNodePoolConfigRootVolumeSet(c *Client, des, nw []AzureNodePoolConfigRootVolume) []AzureNodePoolConfigRootVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureNodePoolConfigRootVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureNodePoolConfigRootVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureNodePoolConfigRootVolumeSlice(c *Client, des, nw []AzureNodePoolConfigRootVolume) []AzureNodePoolConfigRootVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureNodePoolConfigRootVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureNodePoolConfigRootVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAzureNodePoolConfigSshConfig(des, initial *AzureNodePoolConfigSshConfig, opts ...dcl.ApplyOption) *AzureNodePoolConfigSshConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureNodePoolConfigSshConfig{}

	if dcl.StringCanonicalize(des.AuthorizedKey, initial.AuthorizedKey) || dcl.IsZeroValue(des.AuthorizedKey) {
		cDes.AuthorizedKey = initial.AuthorizedKey
	} else {
		cDes.AuthorizedKey = des.AuthorizedKey
	}

	return cDes
}

func canonicalizeNewAzureNodePoolConfigSshConfig(c *Client, des, nw *AzureNodePoolConfigSshConfig) *AzureNodePoolConfigSshConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AuthorizedKey, nw.AuthorizedKey) {
		nw.AuthorizedKey = des.AuthorizedKey
	}

	return nw
}

func canonicalizeNewAzureNodePoolConfigSshConfigSet(c *Client, des, nw []AzureNodePoolConfigSshConfig) []AzureNodePoolConfigSshConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureNodePoolConfigSshConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureNodePoolConfigSshConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureNodePoolConfigSshConfigSlice(c *Client, des, nw []AzureNodePoolConfigSshConfig) []AzureNodePoolConfigSshConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureNodePoolConfigSshConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureNodePoolConfigSshConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAzureNodePoolAutoscaling(des, initial *AzureNodePoolAutoscaling, opts ...dcl.ApplyOption) *AzureNodePoolAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureNodePoolAutoscaling{}

	if dcl.IsZeroValue(des.MinNodeCount) {
		des.MinNodeCount = initial.MinNodeCount
	} else {
		cDes.MinNodeCount = des.MinNodeCount
	}
	if dcl.IsZeroValue(des.MaxNodeCount) {
		des.MaxNodeCount = initial.MaxNodeCount
	} else {
		cDes.MaxNodeCount = des.MaxNodeCount
	}

	return cDes
}

func canonicalizeNewAzureNodePoolAutoscaling(c *Client, des, nw *AzureNodePoolAutoscaling) *AzureNodePoolAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAzureNodePoolAutoscalingSet(c *Client, des, nw []AzureNodePoolAutoscaling) []AzureNodePoolAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureNodePoolAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureNodePoolAutoscalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureNodePoolAutoscalingSlice(c *Client, des, nw []AzureNodePoolAutoscaling) []AzureNodePoolAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureNodePoolAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureNodePoolAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeAzureNodePoolMaxPodsConstraint(des, initial *AzureNodePoolMaxPodsConstraint, opts ...dcl.ApplyOption) *AzureNodePoolMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AzureNodePoolMaxPodsConstraint{}

	if dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	} else {
		cDes.MaxPodsPerNode = des.MaxPodsPerNode
	}

	return cDes
}

func canonicalizeNewAzureNodePoolMaxPodsConstraint(c *Client, des, nw *AzureNodePoolMaxPodsConstraint) *AzureNodePoolMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAzureNodePoolMaxPodsConstraintSet(c *Client, des, nw []AzureNodePoolMaxPodsConstraint) []AzureNodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}
	var reorderedNew []AzureNodePoolMaxPodsConstraint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAzureNodePoolMaxPodsConstraintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAzureNodePoolMaxPodsConstraintSlice(c *Client, des, nw []AzureNodePoolMaxPodsConstraint) []AzureNodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AzureNodePoolMaxPodsConstraint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAzureNodePoolMaxPodsConstraint(c, &d, &n))
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
func diffAzureNodePool(c *Client, desired, actual *AzureNodePool, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareAzureNodePoolConfigNewStyle, EmptyObject: EmptyAzureNodePoolConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetId, actual.SubnetId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autoscaling, actual.Autoscaling, dcl.Info{ObjectFunction: compareAzureNodePoolAutoscalingNewStyle, EmptyObject: EmptyAzureNodePoolAutoscaling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Autoscaling")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.MaxPodsConstraint, actual.MaxPodsConstraint, dcl.Info{ObjectFunction: compareAzureNodePoolMaxPodsConstraintNewStyle, EmptyObject: EmptyAzureNodePoolMaxPodsConstraint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxPodsConstraint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AzureAvailabilityZone, actual.AzureAvailabilityZone, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureAvailabilityZone")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.AzureCluster, actual.AzureCluster, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareAzureNodePoolConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureNodePoolConfig)
	if !ok {
		desiredNotPointer, ok := d.(AzureNodePoolConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfig or *AzureNodePoolConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureNodePoolConfig)
	if !ok {
		actualNotPointer, ok := a.(AzureNodePoolConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VmSize, actual.VmSize, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VmSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootVolume, actual.RootVolume, dcl.Info{ObjectFunction: compareAzureNodePoolConfigRootVolumeNewStyle, EmptyObject: EmptyAzureNodePoolConfigRootVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootVolume")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SshConfig, actual.SshConfig, dcl.Info{ObjectFunction: compareAzureNodePoolConfigSshConfigNewStyle, EmptyObject: EmptyAzureNodePoolConfigSshConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SshConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureNodePoolConfigRootVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureNodePoolConfigRootVolume)
	if !ok {
		desiredNotPointer, ok := d.(AzureNodePoolConfigRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfigRootVolume or *AzureNodePoolConfigRootVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureNodePoolConfigRootVolume)
	if !ok {
		actualNotPointer, ok := a.(AzureNodePoolConfigRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfigRootVolume", a)
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

func compareAzureNodePoolConfigSshConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureNodePoolConfigSshConfig)
	if !ok {
		desiredNotPointer, ok := d.(AzureNodePoolConfigSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfigSshConfig or *AzureNodePoolConfigSshConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureNodePoolConfigSshConfig)
	if !ok {
		actualNotPointer, ok := a.(AzureNodePoolConfigSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolConfigSshConfig", a)
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

func compareAzureNodePoolAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureNodePoolAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(AzureNodePoolAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolAutoscaling or *AzureNodePoolAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureNodePoolAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(AzureNodePoolAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolAutoscaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinNodeCount, actual.MinNodeCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxNodeCount, actual.MaxNodeCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAzureNodePoolMaxPodsConstraintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AzureNodePoolMaxPodsConstraint)
	if !ok {
		desiredNotPointer, ok := d.(AzureNodePoolMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolMaxPodsConstraint or *AzureNodePoolMaxPodsConstraint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AzureNodePoolMaxPodsConstraint)
	if !ok {
		actualNotPointer, ok := a.(AzureNodePoolMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AzureNodePoolMaxPodsConstraint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxPodsPerNode, actual.MaxPodsPerNode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxPodsPerNode")); len(ds) != 0 || err != nil {
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
func (r *AzureNodePool) urlNormalized() *AzureNodePool {
	normalized := dcl.Copy(*r).(AzureNodePool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.SubnetId = dcl.SelfLinkToName(r.SubnetId)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.AzureAvailabilityZone = dcl.SelfLinkToName(r.AzureAvailabilityZone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.AzureCluster = dcl.SelfLinkToName(r.AzureCluster)
	return &normalized
}

func (r *AzureNodePool) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AzureNodePool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AzureNodePool) marshal(c *Client) ([]byte, error) {
	m, err := expandAzureNodePool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AzureNodePool: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAzureNodePool decodes JSON responses into the AzureNodePool resource schema.
func unmarshalAzureNodePool(b []byte, c *Client) (*AzureNodePool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAzureNodePool(m, c)
}

func unmarshalMapAzureNodePool(m map[string]interface{}, c *Client) (*AzureNodePool, error) {

	return flattenAzureNodePool(c, m), nil
}

// expandAzureNodePool expands AzureNodePool into a JSON request object.
func expandAzureNodePool(c *Client, f *AzureNodePool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", f.Name, f.Project, f.Location, f.AzureCluster, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v, err := expandAzureNodePoolConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if v != nil {
		m["config"] = v
	}
	if v := f.SubnetId; !dcl.IsEmptyValueIndirect(v) {
		m["subnetId"] = v
	}
	if v, err := expandAzureNodePoolAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if v != nil {
		m["autoscaling"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandAzureNodePoolMaxPodsConstraint(c, f.MaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding MaxPodsConstraint into maxPodsConstraint: %w", err)
	} else if v != nil {
		m["maxPodsConstraint"] = v
	}
	if v := f.AzureAvailabilityZone; !dcl.IsEmptyValueIndirect(v) {
		m["azureAvailabilityZone"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding AzureCluster into azureCluster: %w", err)
	} else if v != nil {
		m["azureCluster"] = v
	}

	return m, nil
}

// flattenAzureNodePool flattens AzureNodePool from a JSON request object into the
// AzureNodePool type.
func flattenAzureNodePool(c *Client, i interface{}) *AzureNodePool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AzureNodePool{}
	res.Name = dcl.FlattenString(m["name"])
	res.Version = dcl.FlattenString(m["version"])
	res.Config = flattenAzureNodePoolConfig(c, m["config"])
	res.SubnetId = dcl.FlattenString(m["subnetId"])
	res.Autoscaling = flattenAzureNodePoolAutoscaling(c, m["autoscaling"])
	res.State = flattenAzureNodePoolStateEnum(m["state"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.MaxPodsConstraint = flattenAzureNodePoolMaxPodsConstraint(c, m["maxPodsConstraint"])
	res.AzureAvailabilityZone = dcl.FlattenString(m["azureAvailabilityZone"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])
	res.AzureCluster = dcl.FlattenString(m["azureCluster"])

	return res
}

// expandAzureNodePoolConfigMap expands the contents of AzureNodePoolConfig into a JSON
// request object.
func expandAzureNodePoolConfigMap(c *Client, f map[string]AzureNodePoolConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureNodePoolConfigSlice expands the contents of AzureNodePoolConfig into a JSON
// request object.
func expandAzureNodePoolConfigSlice(c *Client, f []AzureNodePoolConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureNodePoolConfigMap flattens the contents of AzureNodePoolConfig from a JSON
// response object.
func flattenAzureNodePoolConfigMap(c *Client, i interface{}) map[string]AzureNodePoolConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolConfig{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolConfig{}
	}

	items := make(map[string]AzureNodePoolConfig)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureNodePoolConfigSlice flattens the contents of AzureNodePoolConfig from a JSON
// response object.
func flattenAzureNodePoolConfigSlice(c *Client, i interface{}) []AzureNodePoolConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolConfig{}
	}

	if len(a) == 0 {
		return []AzureNodePoolConfig{}
	}

	items := make([]AzureNodePoolConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureNodePoolConfig expands an instance of AzureNodePoolConfig into a JSON
// request object.
func expandAzureNodePoolConfig(c *Client, f *AzureNodePoolConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.VmSize; !dcl.IsEmptyValueIndirect(v) {
		m["vmSize"] = v
	}
	if v, err := expandAzureNodePoolConfigRootVolume(c, f.RootVolume); err != nil {
		return nil, fmt.Errorf("error expanding RootVolume into rootVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rootVolume"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v, err := expandAzureNodePoolConfigSshConfig(c, f.SshConfig); err != nil {
		return nil, fmt.Errorf("error expanding SshConfig into sshConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sshConfig"] = v
	}

	return m, nil
}

// flattenAzureNodePoolConfig flattens an instance of AzureNodePoolConfig from a JSON
// response object.
func flattenAzureNodePoolConfig(c *Client, i interface{}) *AzureNodePoolConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureNodePoolConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureNodePoolConfig
	}
	r.VmSize = dcl.FlattenString(m["vmSize"])
	r.RootVolume = flattenAzureNodePoolConfigRootVolume(c, m["rootVolume"])
	r.Tags = dcl.FlattenKeyValuePairs(m["tags"])
	r.SshConfig = flattenAzureNodePoolConfigSshConfig(c, m["sshConfig"])

	return r
}

// expandAzureNodePoolConfigRootVolumeMap expands the contents of AzureNodePoolConfigRootVolume into a JSON
// request object.
func expandAzureNodePoolConfigRootVolumeMap(c *Client, f map[string]AzureNodePoolConfigRootVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureNodePoolConfigRootVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureNodePoolConfigRootVolumeSlice expands the contents of AzureNodePoolConfigRootVolume into a JSON
// request object.
func expandAzureNodePoolConfigRootVolumeSlice(c *Client, f []AzureNodePoolConfigRootVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureNodePoolConfigRootVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureNodePoolConfigRootVolumeMap flattens the contents of AzureNodePoolConfigRootVolume from a JSON
// response object.
func flattenAzureNodePoolConfigRootVolumeMap(c *Client, i interface{}) map[string]AzureNodePoolConfigRootVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolConfigRootVolume{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolConfigRootVolume{}
	}

	items := make(map[string]AzureNodePoolConfigRootVolume)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolConfigRootVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureNodePoolConfigRootVolumeSlice flattens the contents of AzureNodePoolConfigRootVolume from a JSON
// response object.
func flattenAzureNodePoolConfigRootVolumeSlice(c *Client, i interface{}) []AzureNodePoolConfigRootVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolConfigRootVolume{}
	}

	if len(a) == 0 {
		return []AzureNodePoolConfigRootVolume{}
	}

	items := make([]AzureNodePoolConfigRootVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolConfigRootVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureNodePoolConfigRootVolume expands an instance of AzureNodePoolConfigRootVolume into a JSON
// request object.
func expandAzureNodePoolConfigRootVolume(c *Client, f *AzureNodePoolConfigRootVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}

	return m, nil
}

// flattenAzureNodePoolConfigRootVolume flattens an instance of AzureNodePoolConfigRootVolume from a JSON
// response object.
func flattenAzureNodePoolConfigRootVolume(c *Client, i interface{}) *AzureNodePoolConfigRootVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureNodePoolConfigRootVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureNodePoolConfigRootVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])

	return r
}

// expandAzureNodePoolConfigSshConfigMap expands the contents of AzureNodePoolConfigSshConfig into a JSON
// request object.
func expandAzureNodePoolConfigSshConfigMap(c *Client, f map[string]AzureNodePoolConfigSshConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureNodePoolConfigSshConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureNodePoolConfigSshConfigSlice expands the contents of AzureNodePoolConfigSshConfig into a JSON
// request object.
func expandAzureNodePoolConfigSshConfigSlice(c *Client, f []AzureNodePoolConfigSshConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureNodePoolConfigSshConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureNodePoolConfigSshConfigMap flattens the contents of AzureNodePoolConfigSshConfig from a JSON
// response object.
func flattenAzureNodePoolConfigSshConfigMap(c *Client, i interface{}) map[string]AzureNodePoolConfigSshConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolConfigSshConfig{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolConfigSshConfig{}
	}

	items := make(map[string]AzureNodePoolConfigSshConfig)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolConfigSshConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureNodePoolConfigSshConfigSlice flattens the contents of AzureNodePoolConfigSshConfig from a JSON
// response object.
func flattenAzureNodePoolConfigSshConfigSlice(c *Client, i interface{}) []AzureNodePoolConfigSshConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolConfigSshConfig{}
	}

	if len(a) == 0 {
		return []AzureNodePoolConfigSshConfig{}
	}

	items := make([]AzureNodePoolConfigSshConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolConfigSshConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureNodePoolConfigSshConfig expands an instance of AzureNodePoolConfigSshConfig into a JSON
// request object.
func expandAzureNodePoolConfigSshConfig(c *Client, f *AzureNodePoolConfigSshConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AuthorizedKey; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedKey"] = v
	}

	return m, nil
}

// flattenAzureNodePoolConfigSshConfig flattens an instance of AzureNodePoolConfigSshConfig from a JSON
// response object.
func flattenAzureNodePoolConfigSshConfig(c *Client, i interface{}) *AzureNodePoolConfigSshConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureNodePoolConfigSshConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureNodePoolConfigSshConfig
	}
	r.AuthorizedKey = dcl.FlattenString(m["authorizedKey"])

	return r
}

// expandAzureNodePoolAutoscalingMap expands the contents of AzureNodePoolAutoscaling into a JSON
// request object.
func expandAzureNodePoolAutoscalingMap(c *Client, f map[string]AzureNodePoolAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureNodePoolAutoscalingSlice expands the contents of AzureNodePoolAutoscaling into a JSON
// request object.
func expandAzureNodePoolAutoscalingSlice(c *Client, f []AzureNodePoolAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureNodePoolAutoscalingMap flattens the contents of AzureNodePoolAutoscaling from a JSON
// response object.
func flattenAzureNodePoolAutoscalingMap(c *Client, i interface{}) map[string]AzureNodePoolAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolAutoscaling{}
	}

	items := make(map[string]AzureNodePoolAutoscaling)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureNodePoolAutoscalingSlice flattens the contents of AzureNodePoolAutoscaling from a JSON
// response object.
func flattenAzureNodePoolAutoscalingSlice(c *Client, i interface{}) []AzureNodePoolAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return []AzureNodePoolAutoscaling{}
	}

	items := make([]AzureNodePoolAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureNodePoolAutoscaling expands an instance of AzureNodePoolAutoscaling into a JSON
// request object.
func expandAzureNodePoolAutoscaling(c *Client, f *AzureNodePoolAutoscaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["minNodeCount"] = v
	}
	if v := f.MaxNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxNodeCount"] = v
	}

	return m, nil
}

// flattenAzureNodePoolAutoscaling flattens an instance of AzureNodePoolAutoscaling from a JSON
// response object.
func flattenAzureNodePoolAutoscaling(c *Client, i interface{}) *AzureNodePoolAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureNodePoolAutoscaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureNodePoolAutoscaling
	}
	r.MinNodeCount = dcl.FlattenInteger(m["minNodeCount"])
	r.MaxNodeCount = dcl.FlattenInteger(m["maxNodeCount"])

	return r
}

// expandAzureNodePoolMaxPodsConstraintMap expands the contents of AzureNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAzureNodePoolMaxPodsConstraintMap(c *Client, f map[string]AzureNodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAzureNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAzureNodePoolMaxPodsConstraintSlice expands the contents of AzureNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAzureNodePoolMaxPodsConstraintSlice(c *Client, f []AzureNodePoolMaxPodsConstraint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAzureNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAzureNodePoolMaxPodsConstraintMap flattens the contents of AzureNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAzureNodePoolMaxPodsConstraintMap(c *Client, i interface{}) map[string]AzureNodePoolMaxPodsConstraint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolMaxPodsConstraint{}
	}

	items := make(map[string]AzureNodePoolMaxPodsConstraint)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolMaxPodsConstraint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAzureNodePoolMaxPodsConstraintSlice flattens the contents of AzureNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAzureNodePoolMaxPodsConstraintSlice(c *Client, i interface{}) []AzureNodePoolMaxPodsConstraint {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return []AzureNodePoolMaxPodsConstraint{}
	}

	items := make([]AzureNodePoolMaxPodsConstraint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolMaxPodsConstraint(c, item.(map[string]interface{})))
	}

	return items
}

// expandAzureNodePoolMaxPodsConstraint expands an instance of AzureNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAzureNodePoolMaxPodsConstraint(c *Client, f *AzureNodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxPodsPerNode; !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsPerNode"] = v
	}

	return m, nil
}

// flattenAzureNodePoolMaxPodsConstraint flattens an instance of AzureNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAzureNodePoolMaxPodsConstraint(c *Client, i interface{}) *AzureNodePoolMaxPodsConstraint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AzureNodePoolMaxPodsConstraint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAzureNodePoolMaxPodsConstraint
	}
	r.MaxPodsPerNode = dcl.FlattenInteger(m["maxPodsPerNode"])

	return r
}

// flattenAzureNodePoolStateEnumMap flattens the contents of AzureNodePoolStateEnum from a JSON
// response object.
func flattenAzureNodePoolStateEnumMap(c *Client, i interface{}) map[string]AzureNodePoolStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AzureNodePoolStateEnum{}
	}

	if len(a) == 0 {
		return map[string]AzureNodePoolStateEnum{}
	}

	items := make(map[string]AzureNodePoolStateEnum)
	for k, item := range a {
		items[k] = *flattenAzureNodePoolStateEnum(item.(interface{}))
	}

	return items
}

// flattenAzureNodePoolStateEnumSlice flattens the contents of AzureNodePoolStateEnum from a JSON
// response object.
func flattenAzureNodePoolStateEnumSlice(c *Client, i interface{}) []AzureNodePoolStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AzureNodePoolStateEnum{}
	}

	if len(a) == 0 {
		return []AzureNodePoolStateEnum{}
	}

	items := make([]AzureNodePoolStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAzureNodePoolStateEnum(item.(interface{})))
	}

	return items
}

// flattenAzureNodePoolStateEnum asserts that an interface is a string, and returns a
// pointer to a *AzureNodePoolStateEnum with the same value as that string.
func flattenAzureNodePoolStateEnum(i interface{}) *AzureNodePoolStateEnum {
	s, ok := i.(string)
	if !ok {
		return AzureNodePoolStateEnumRef("")
	}

	return AzureNodePoolStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AzureNodePool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAzureNodePool(b, c)
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
		if nr.AzureCluster == nil && ncr.AzureCluster == nil {
			c.Config.Logger.Info("Both AzureCluster fields null - considering equal.")
		} else if nr.AzureCluster == nil || ncr.AzureCluster == nil {
			c.Config.Logger.Info("Only one AzureCluster field is null - considering unequal.")
			return false
		} else if *nr.AzureCluster != *ncr.AzureCluster {
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

type azureNodePoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         azureNodePoolApiOperation
}

func convertFieldDiffsToAzureNodePoolDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]azureNodePoolDiff, error) {
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
	var diffs []azureNodePoolDiff
	// For each operation name, create a azureNodePoolDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := azureNodePoolDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAzureNodePoolApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAzureNodePoolApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (azureNodePoolApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractAzureNodePoolFields(r *AzureNodePool) error {
	return nil
}
