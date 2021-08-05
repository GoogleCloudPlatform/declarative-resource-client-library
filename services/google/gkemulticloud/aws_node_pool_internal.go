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
package gkemulticloud

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

func (r *AwsNodePool) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "version"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if err := dcl.Required(r, "autoscaling"); err != nil {
		return err
	}
	if err := dcl.Required(r, "subnetId"); err != nil {
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
	if err := dcl.RequiredParameter(r.AwsCluster, "AwsCluster"); err != nil {
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
func (r *AwsNodePoolConfig) validate() error {
	if err := dcl.Required(r, "iamInstanceProfile"); err != nil {
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
func (r *AwsNodePoolConfigRootVolume) validate() error {
	return nil
}
func (r *AwsNodePoolConfigTaints) validate() error {
	if err := dcl.Required(r, "key"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if err := dcl.Required(r, "effect"); err != nil {
		return err
	}
	return nil
}
func (r *AwsNodePoolConfigSshConfig) validate() error {
	if err := dcl.Required(r, "ec2KeyPair"); err != nil {
		return err
	}
	return nil
}
func (r *AwsNodePoolAutoscaling) validate() error {
	if err := dcl.Required(r, "minNodeCount"); err != nil {
		return err
	}
	if err := dcl.Required(r, "maxNodeCount"); err != nil {
		return err
	}
	return nil
}
func (r *AwsNodePoolMaxPodsConstraint) validate() error {
	if err := dcl.Required(r, "maxPodsPerNode"); err != nil {
		return err
	}
	return nil
}
func (r *AwsNodePool) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-gkemulticloud.googleapis.com/v1", params)
}

func (r *AwsNodePool) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":    dcl.ValueOrEmptyString(nr.Project),
		"location":   dcl.ValueOrEmptyString(nr.Location),
		"awsCluster": dcl.ValueOrEmptyString(nr.AwsCluster),
		"name":       dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{awsCluster}}/awsNodePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AwsNodePool) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":    dcl.ValueOrEmptyString(nr.Project),
		"location":   dcl.ValueOrEmptyString(nr.Location),
		"awsCluster": dcl.ValueOrEmptyString(nr.AwsCluster),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{awsCluster}}/awsNodePools", nr.basePath(), userBasePath, params), nil

}

func (r *AwsNodePool) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":    dcl.ValueOrEmptyString(nr.Project),
		"location":   dcl.ValueOrEmptyString(nr.Location),
		"awsCluster": dcl.ValueOrEmptyString(nr.AwsCluster),
		"name":       dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{awsCluster}}/awsNodePools?awsNodePoolId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AwsNodePool) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":    dcl.ValueOrEmptyString(nr.Project),
		"location":   dcl.ValueOrEmptyString(nr.Location),
		"awsCluster": dcl.ValueOrEmptyString(nr.AwsCluster),
		"name":       dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/awsClusters/{{awsCluster}}/awsNodePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

// awsNodePoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type awsNodePoolApiOperation interface {
	do(context.Context, *AwsNodePool, *Client) error
}

func (c *Client) listAwsNodePoolRaw(ctx context.Context, r *AwsNodePool, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AwsNodePoolMaxPage {
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

type listAwsNodePoolOperation struct {
	AwsNodePools []map[string]interface{} `json:"awsNodePools"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listAwsNodePool(ctx context.Context, r *AwsNodePool, pageToken string, pageSize int32) ([]*AwsNodePool, string, error) {
	b, err := c.listAwsNodePoolRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAwsNodePoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AwsNodePool
	for _, v := range m.AwsNodePools {
		res, err := unmarshalMapAwsNodePool(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.AwsCluster = r.AwsCluster
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAwsNodePool(ctx context.Context, f func(*AwsNodePool) bool, resources []*AwsNodePool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAwsNodePool(ctx, res)
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

type deleteAwsNodePoolOperation struct{}

func (op *deleteAwsNodePoolOperation) do(ctx context.Context, r *AwsNodePool, c *Client) error {
	r, err := c.GetAwsNodePool(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AwsNodePool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAwsNodePool checking for existence. error: %v", err)
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
	if err := o.Wait(ctx, c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAwsNodePool(ctx, r)
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
type createAwsNodePoolOperation struct {
	response map[string]interface{}
}

func (op *createAwsNodePoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAwsNodePoolOperation) do(ctx context.Context, r *AwsNodePool, c *Client) error {
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
	if err := o.Wait(ctx, c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetAwsNodePool(ctx, r); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAwsNodePoolRaw(ctx context.Context, r *AwsNodePool) ([]byte, error) {

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

func (c *Client) awsNodePoolDiffsForRawDesired(ctx context.Context, rawDesired *AwsNodePool, opts ...dcl.ApplyOption) (initial, desired *AwsNodePool, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AwsNodePool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AwsNodePool); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AwsNodePool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAwsNodePool(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AwsNodePool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AwsNodePool resource: %v", err)
		}
		c.Config.Logger.Info("Found that AwsNodePool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAwsNodePoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for AwsNodePool: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AwsNodePool: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAwsNodePoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AwsNodePool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAwsNodePoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AwsNodePool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAwsNodePool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAwsNodePoolInitialState(rawInitial, rawDesired *AwsNodePool) (*AwsNodePool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAwsNodePoolDesiredState(rawDesired, rawInitial *AwsNodePool, opts ...dcl.ApplyOption) (*AwsNodePool, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeAwsNodePoolConfig(rawDesired.Config, nil, opts...)
		rawDesired.Autoscaling = canonicalizeAwsNodePoolAutoscaling(rawDesired.Autoscaling, nil, opts...)
		rawDesired.MaxPodsConstraint = canonicalizeAwsNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &AwsNodePool{}
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
	canonicalDesired.Config = canonicalizeAwsNodePoolConfig(rawDesired.Config, rawInitial.Config, opts...)
	canonicalDesired.Autoscaling = canonicalizeAwsNodePoolAutoscaling(rawDesired.Autoscaling, rawInitial.Autoscaling, opts...)
	if dcl.StringCanonicalize(rawDesired.SubnetId, rawInitial.SubnetId) {
		canonicalDesired.SubnetId = rawInitial.SubnetId
	} else {
		canonicalDesired.SubnetId = rawDesired.SubnetId
	}
	if dcl.IsZeroValue(rawDesired.Annotations) {
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	canonicalDesired.MaxPodsConstraint = canonicalizeAwsNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, rawInitial.MaxPodsConstraint, opts...)
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
	if dcl.NameToSelfLink(rawDesired.AwsCluster, rawInitial.AwsCluster) {
		canonicalDesired.AwsCluster = rawInitial.AwsCluster
	} else {
		canonicalDesired.AwsCluster = rawDesired.AwsCluster
	}

	return canonicalDesired, nil
}

func canonicalizeAwsNodePoolNewState(c *Client, rawNew, rawDesired *AwsNodePool) (*AwsNodePool, error) {

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
		rawNew.Config = canonicalizeNewAwsNodePoolConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Autoscaling) && dcl.IsEmptyValueIndirect(rawDesired.Autoscaling) {
		rawNew.Autoscaling = rawDesired.Autoscaling
	} else {
		rawNew.Autoscaling = canonicalizeNewAwsNodePoolAutoscaling(c, rawDesired.Autoscaling, rawNew.Autoscaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubnetId) && dcl.IsEmptyValueIndirect(rawDesired.SubnetId) {
		rawNew.SubnetId = rawDesired.SubnetId
	} else {
		if dcl.StringCanonicalize(rawDesired.SubnetId, rawNew.SubnetId) {
			rawNew.SubnetId = rawDesired.SubnetId
		}
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
		rawNew.MaxPodsConstraint = canonicalizeNewAwsNodePoolMaxPodsConstraint(c, rawDesired.MaxPodsConstraint, rawNew.MaxPodsConstraint)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.AwsCluster = rawDesired.AwsCluster

	return rawNew, nil
}

func canonicalizeAwsNodePoolConfig(des, initial *AwsNodePoolConfig, opts ...dcl.ApplyOption) *AwsNodePoolConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolConfig{}

	if dcl.StringCanonicalize(des.InstanceType, initial.InstanceType) || dcl.IsZeroValue(des.InstanceType) {
		cDes.InstanceType = initial.InstanceType
	} else {
		cDes.InstanceType = des.InstanceType
	}
	cDes.RootVolume = canonicalizeAwsNodePoolConfigRootVolume(des.RootVolume, initial.RootVolume, opts...)
	if dcl.IsZeroValue(des.Taints) {
		des.Taints = initial.Taints
	} else {
		cDes.Taints = des.Taints
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	} else {
		cDes.Labels = des.Labels
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	if dcl.StringCanonicalize(des.IamInstanceProfile, initial.IamInstanceProfile) || dcl.IsZeroValue(des.IamInstanceProfile) {
		cDes.IamInstanceProfile = initial.IamInstanceProfile
	} else {
		cDes.IamInstanceProfile = des.IamInstanceProfile
	}
	cDes.SshConfig = canonicalizeAwsNodePoolConfigSshConfig(des.SshConfig, initial.SshConfig, opts...)
	if dcl.IsZeroValue(des.SecurityGroupIds) {
		des.SecurityGroupIds = initial.SecurityGroupIds
	} else {
		cDes.SecurityGroupIds = des.SecurityGroupIds
	}

	return cDes
}

func canonicalizeNewAwsNodePoolConfig(c *Client, des, nw *AwsNodePoolConfig) *AwsNodePoolConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.InstanceType, nw.InstanceType) {
		nw.InstanceType = des.InstanceType
	}
	nw.RootVolume = canonicalizeNewAwsNodePoolConfigRootVolume(c, des.RootVolume, nw.RootVolume)
	nw.Taints = canonicalizeNewAwsNodePoolConfigTaintsSlice(c, des.Taints, nw.Taints)
	if dcl.StringCanonicalize(des.IamInstanceProfile, nw.IamInstanceProfile) {
		nw.IamInstanceProfile = des.IamInstanceProfile
	}
	nw.SshConfig = canonicalizeNewAwsNodePoolConfigSshConfig(c, des.SshConfig, nw.SshConfig)

	return nw
}

func canonicalizeNewAwsNodePoolConfigSet(c *Client, des, nw []AwsNodePoolConfig) []AwsNodePoolConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolConfigSlice(c *Client, des, nw []AwsNodePoolConfig) []AwsNodePoolConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAwsNodePoolConfigRootVolume(des, initial *AwsNodePoolConfigRootVolume, opts ...dcl.ApplyOption) *AwsNodePoolConfigRootVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolConfigRootVolume{}

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

func canonicalizeNewAwsNodePoolConfigRootVolume(c *Client, des, nw *AwsNodePoolConfigRootVolume) *AwsNodePoolConfigRootVolume {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyArn, nw.KmsKeyArn) {
		nw.KmsKeyArn = des.KmsKeyArn
	}

	return nw
}

func canonicalizeNewAwsNodePoolConfigRootVolumeSet(c *Client, des, nw []AwsNodePoolConfigRootVolume) []AwsNodePoolConfigRootVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolConfigRootVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolConfigRootVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolConfigRootVolumeSlice(c *Client, des, nw []AwsNodePoolConfigRootVolume) []AwsNodePoolConfigRootVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolConfigRootVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolConfigRootVolume(c, &d, &n))
	}

	return items
}

func canonicalizeAwsNodePoolConfigTaints(des, initial *AwsNodePoolConfigTaints, opts ...dcl.ApplyOption) *AwsNodePoolConfigTaints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolConfigTaints{}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		cDes.Key = initial.Key
	} else {
		cDes.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	if dcl.IsZeroValue(des.Effect) {
		des.Effect = initial.Effect
	} else {
		cDes.Effect = des.Effect
	}

	return cDes
}

func canonicalizeNewAwsNodePoolConfigTaints(c *Client, des, nw *AwsNodePoolConfigTaints) *AwsNodePoolConfigTaints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewAwsNodePoolConfigTaintsSet(c *Client, des, nw []AwsNodePoolConfigTaints) []AwsNodePoolConfigTaints {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolConfigTaints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolConfigTaintsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolConfigTaintsSlice(c *Client, des, nw []AwsNodePoolConfigTaints) []AwsNodePoolConfigTaints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolConfigTaints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolConfigTaints(c, &d, &n))
	}

	return items
}

func canonicalizeAwsNodePoolConfigSshConfig(des, initial *AwsNodePoolConfigSshConfig, opts ...dcl.ApplyOption) *AwsNodePoolConfigSshConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolConfigSshConfig{}

	if dcl.StringCanonicalize(des.Ec2KeyPair, initial.Ec2KeyPair) || dcl.IsZeroValue(des.Ec2KeyPair) {
		cDes.Ec2KeyPair = initial.Ec2KeyPair
	} else {
		cDes.Ec2KeyPair = des.Ec2KeyPair
	}

	return cDes
}

func canonicalizeNewAwsNodePoolConfigSshConfig(c *Client, des, nw *AwsNodePoolConfigSshConfig) *AwsNodePoolConfigSshConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Ec2KeyPair, nw.Ec2KeyPair) {
		nw.Ec2KeyPair = des.Ec2KeyPair
	}

	return nw
}

func canonicalizeNewAwsNodePoolConfigSshConfigSet(c *Client, des, nw []AwsNodePoolConfigSshConfig) []AwsNodePoolConfigSshConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolConfigSshConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolConfigSshConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolConfigSshConfigSlice(c *Client, des, nw []AwsNodePoolConfigSshConfig) []AwsNodePoolConfigSshConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolConfigSshConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolConfigSshConfig(c, &d, &n))
	}

	return items
}

func canonicalizeAwsNodePoolAutoscaling(des, initial *AwsNodePoolAutoscaling, opts ...dcl.ApplyOption) *AwsNodePoolAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolAutoscaling{}

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

func canonicalizeNewAwsNodePoolAutoscaling(c *Client, des, nw *AwsNodePoolAutoscaling) *AwsNodePoolAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAwsNodePoolAutoscalingSet(c *Client, des, nw []AwsNodePoolAutoscaling) []AwsNodePoolAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolAutoscalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolAutoscalingSlice(c *Client, des, nw []AwsNodePoolAutoscaling) []AwsNodePoolAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeAwsNodePoolMaxPodsConstraint(des, initial *AwsNodePoolMaxPodsConstraint, opts ...dcl.ApplyOption) *AwsNodePoolMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AwsNodePoolMaxPodsConstraint{}

	if dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	} else {
		cDes.MaxPodsPerNode = des.MaxPodsPerNode
	}

	return cDes
}

func canonicalizeNewAwsNodePoolMaxPodsConstraint(c *Client, des, nw *AwsNodePoolMaxPodsConstraint) *AwsNodePoolMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAwsNodePoolMaxPodsConstraintSet(c *Client, des, nw []AwsNodePoolMaxPodsConstraint) []AwsNodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}
	var reorderedNew []AwsNodePoolMaxPodsConstraint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAwsNodePoolMaxPodsConstraintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAwsNodePoolMaxPodsConstraintSlice(c *Client, des, nw []AwsNodePoolMaxPodsConstraint) []AwsNodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AwsNodePoolMaxPodsConstraint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAwsNodePoolMaxPodsConstraint(c, &d, &n))
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
func diffAwsNodePool(c *Client, desired, actual *AwsNodePool, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareAwsNodePoolConfigNewStyle, EmptyObject: EmptyAwsNodePoolConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autoscaling, actual.Autoscaling, dcl.Info{ObjectFunction: compareAwsNodePoolAutoscalingNewStyle, EmptyObject: EmptyAwsNodePoolAutoscaling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Autoscaling")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.MaxPodsConstraint, actual.MaxPodsConstraint, dcl.Info{ObjectFunction: compareAwsNodePoolMaxPodsConstraintNewStyle, EmptyObject: EmptyAwsNodePoolMaxPodsConstraint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxPodsConstraint")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.AwsCluster, actual.AwsCluster, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AwsCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareAwsNodePoolConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolConfig)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfig or *AwsNodePoolConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolConfig)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceType, actual.InstanceType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootVolume, actual.RootVolume, dcl.Info{ObjectFunction: compareAwsNodePoolConfigRootVolumeNewStyle, EmptyObject: EmptyAwsNodePoolConfigRootVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Taints, actual.Taints, dcl.Info{ObjectFunction: compareAwsNodePoolConfigTaintsNewStyle, EmptyObject: EmptyAwsNodePoolConfigTaints, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Taints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.IamInstanceProfile, actual.IamInstanceProfile, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IamInstanceProfile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SshConfig, actual.SshConfig, dcl.Info{ObjectFunction: compareAwsNodePoolConfigSshConfigNewStyle, EmptyObject: EmptyAwsNodePoolConfigSshConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SshConfig")); len(ds) != 0 || err != nil {
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
	return diffs, nil
}

func compareAwsNodePoolConfigRootVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolConfigRootVolume)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolConfigRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigRootVolume or *AwsNodePoolConfigRootVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolConfigRootVolume)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolConfigRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigRootVolume", a)
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

func compareAwsNodePoolConfigTaintsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolConfigTaints)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigTaints or *AwsNodePoolConfigTaints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolConfigTaints)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigTaints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Effect, actual.Effect, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Effect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAwsNodePoolConfigSshConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolConfigSshConfig)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolConfigSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigSshConfig or *AwsNodePoolConfigSshConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolConfigSshConfig)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolConfigSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolConfigSshConfig", a)
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

func compareAwsNodePoolAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolAutoscaling or *AwsNodePoolAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolAutoscaling", a)
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

func compareAwsNodePoolMaxPodsConstraintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AwsNodePoolMaxPodsConstraint)
	if !ok {
		desiredNotPointer, ok := d.(AwsNodePoolMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolMaxPodsConstraint or *AwsNodePoolMaxPodsConstraint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AwsNodePoolMaxPodsConstraint)
	if !ok {
		actualNotPointer, ok := a.(AwsNodePoolMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AwsNodePoolMaxPodsConstraint", a)
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
func (r *AwsNodePool) urlNormalized() *AwsNodePool {
	normalized := dcl.Copy(*r).(AwsNodePool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.SubnetId = dcl.SelfLinkToName(r.SubnetId)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.AwsCluster = dcl.SelfLinkToName(r.AwsCluster)
	return &normalized
}

func (r *AwsNodePool) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AwsNodePool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AwsNodePool) marshal(c *Client) ([]byte, error) {
	m, err := expandAwsNodePool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AwsNodePool: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAwsNodePool decodes JSON responses into the AwsNodePool resource schema.
func unmarshalAwsNodePool(b []byte, c *Client) (*AwsNodePool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAwsNodePool(m, c)
}

func unmarshalMapAwsNodePool(m map[string]interface{}, c *Client) (*AwsNodePool, error) {

	return flattenAwsNodePool(c, m), nil
}

// expandAwsNodePool expands AwsNodePool into a JSON request object.
func expandAwsNodePool(c *Client, f *AwsNodePool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", f.Name, f.Project, f.Location, f.AwsCluster, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v, err := expandAwsNodePoolConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if v != nil {
		m["config"] = v
	}
	if v, err := expandAwsNodePoolAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if v != nil {
		m["autoscaling"] = v
	}
	if v := f.SubnetId; !dcl.IsEmptyValueIndirect(v) {
		m["subnetId"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
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
	if v, err := expandAwsNodePoolMaxPodsConstraint(c, f.MaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding MaxPodsConstraint into maxPodsConstraint: %w", err)
	} else if v != nil {
		m["maxPodsConstraint"] = v
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
		return nil, fmt.Errorf("error expanding AwsCluster into awsCluster: %w", err)
	} else if v != nil {
		m["awsCluster"] = v
	}

	return m, nil
}

// flattenAwsNodePool flattens AwsNodePool from a JSON request object into the
// AwsNodePool type.
func flattenAwsNodePool(c *Client, i interface{}) *AwsNodePool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AwsNodePool{}
	res.Name = dcl.FlattenString(m["name"])
	res.Version = dcl.FlattenString(m["version"])
	res.Config = flattenAwsNodePoolConfig(c, m["config"])
	res.Autoscaling = flattenAwsNodePoolAutoscaling(c, m["autoscaling"])
	res.SubnetId = dcl.FlattenString(m["subnetId"])
	res.State = flattenAwsNodePoolStateEnum(m["state"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.MaxPodsConstraint = flattenAwsNodePoolMaxPodsConstraint(c, m["maxPodsConstraint"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])
	res.AwsCluster = dcl.FlattenString(m["awsCluster"])

	return res
}

// expandAwsNodePoolConfigMap expands the contents of AwsNodePoolConfig into a JSON
// request object.
func expandAwsNodePoolConfigMap(c *Client, f map[string]AwsNodePoolConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolConfigSlice expands the contents of AwsNodePoolConfig into a JSON
// request object.
func expandAwsNodePoolConfigSlice(c *Client, f []AwsNodePoolConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolConfigMap flattens the contents of AwsNodePoolConfig from a JSON
// response object.
func flattenAwsNodePoolConfigMap(c *Client, i interface{}) map[string]AwsNodePoolConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfig{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfig{}
	}

	items := make(map[string]AwsNodePoolConfig)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigSlice flattens the contents of AwsNodePoolConfig from a JSON
// response object.
func flattenAwsNodePoolConfigSlice(c *Client, i interface{}) []AwsNodePoolConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfig{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfig{}
	}

	items := make([]AwsNodePoolConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolConfig expands an instance of AwsNodePoolConfig into a JSON
// request object.
func expandAwsNodePoolConfig(c *Client, f *AwsNodePoolConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.InstanceType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceType"] = v
	}
	if v, err := expandAwsNodePoolConfigRootVolume(c, f.RootVolume); err != nil {
		return nil, fmt.Errorf("error expanding RootVolume into rootVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rootVolume"] = v
	}
	if v, err := expandAwsNodePoolConfigTaintsSlice(c, f.Taints); err != nil {
		return nil, fmt.Errorf("error expanding Taints into taints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["taints"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.IamInstanceProfile; !dcl.IsEmptyValueIndirect(v) {
		m["iamInstanceProfile"] = v
	}
	if v, err := expandAwsNodePoolConfigSshConfig(c, f.SshConfig); err != nil {
		return nil, fmt.Errorf("error expanding SshConfig into sshConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sshConfig"] = v
	}
	if v := f.SecurityGroupIds; v != nil {
		m["securityGroupIds"] = v
	}

	return m, nil
}

// flattenAwsNodePoolConfig flattens an instance of AwsNodePoolConfig from a JSON
// response object.
func flattenAwsNodePoolConfig(c *Client, i interface{}) *AwsNodePoolConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolConfig
	}
	r.InstanceType = dcl.FlattenString(m["instanceType"])
	r.RootVolume = flattenAwsNodePoolConfigRootVolume(c, m["rootVolume"])
	r.Taints = flattenAwsNodePoolConfigTaintsSlice(c, m["taints"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Tags = dcl.FlattenKeyValuePairs(m["tags"])
	r.IamInstanceProfile = dcl.FlattenString(m["iamInstanceProfile"])
	r.SshConfig = flattenAwsNodePoolConfigSshConfig(c, m["sshConfig"])
	r.SecurityGroupIds = dcl.FlattenStringSlice(m["securityGroupIds"])

	return r
}

// expandAwsNodePoolConfigRootVolumeMap expands the contents of AwsNodePoolConfigRootVolume into a JSON
// request object.
func expandAwsNodePoolConfigRootVolumeMap(c *Client, f map[string]AwsNodePoolConfigRootVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolConfigRootVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolConfigRootVolumeSlice expands the contents of AwsNodePoolConfigRootVolume into a JSON
// request object.
func expandAwsNodePoolConfigRootVolumeSlice(c *Client, f []AwsNodePoolConfigRootVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolConfigRootVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolConfigRootVolumeMap flattens the contents of AwsNodePoolConfigRootVolume from a JSON
// response object.
func flattenAwsNodePoolConfigRootVolumeMap(c *Client, i interface{}) map[string]AwsNodePoolConfigRootVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfigRootVolume{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfigRootVolume{}
	}

	items := make(map[string]AwsNodePoolConfigRootVolume)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfigRootVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigRootVolumeSlice flattens the contents of AwsNodePoolConfigRootVolume from a JSON
// response object.
func flattenAwsNodePoolConfigRootVolumeSlice(c *Client, i interface{}) []AwsNodePoolConfigRootVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfigRootVolume{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfigRootVolume{}
	}

	items := make([]AwsNodePoolConfigRootVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfigRootVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolConfigRootVolume expands an instance of AwsNodePoolConfigRootVolume into a JSON
// request object.
func expandAwsNodePoolConfigRootVolume(c *Client, f *AwsNodePoolConfigRootVolume) (map[string]interface{}, error) {
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

// flattenAwsNodePoolConfigRootVolume flattens an instance of AwsNodePoolConfigRootVolume from a JSON
// response object.
func flattenAwsNodePoolConfigRootVolume(c *Client, i interface{}) *AwsNodePoolConfigRootVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolConfigRootVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolConfigRootVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])
	r.VolumeType = flattenAwsNodePoolConfigRootVolumeVolumeTypeEnum(m["volumeType"])
	r.Iops = dcl.FlattenInteger(m["iops"])
	r.KmsKeyArn = dcl.FlattenString(m["kmsKeyArn"])

	return r
}

// expandAwsNodePoolConfigTaintsMap expands the contents of AwsNodePoolConfigTaints into a JSON
// request object.
func expandAwsNodePoolConfigTaintsMap(c *Client, f map[string]AwsNodePoolConfigTaints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolConfigTaintsSlice expands the contents of AwsNodePoolConfigTaints into a JSON
// request object.
func expandAwsNodePoolConfigTaintsSlice(c *Client, f []AwsNodePoolConfigTaints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolConfigTaintsMap flattens the contents of AwsNodePoolConfigTaints from a JSON
// response object.
func flattenAwsNodePoolConfigTaintsMap(c *Client, i interface{}) map[string]AwsNodePoolConfigTaints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfigTaints{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfigTaints{}
	}

	items := make(map[string]AwsNodePoolConfigTaints)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfigTaints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigTaintsSlice flattens the contents of AwsNodePoolConfigTaints from a JSON
// response object.
func flattenAwsNodePoolConfigTaintsSlice(c *Client, i interface{}) []AwsNodePoolConfigTaints {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfigTaints{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfigTaints{}
	}

	items := make([]AwsNodePoolConfigTaints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfigTaints(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolConfigTaints expands an instance of AwsNodePoolConfigTaints into a JSON
// request object.
func expandAwsNodePoolConfigTaints(c *Client, f *AwsNodePoolConfigTaints) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Effect; !dcl.IsEmptyValueIndirect(v) {
		m["effect"] = v
	}

	return m, nil
}

// flattenAwsNodePoolConfigTaints flattens an instance of AwsNodePoolConfigTaints from a JSON
// response object.
func flattenAwsNodePoolConfigTaints(c *Client, i interface{}) *AwsNodePoolConfigTaints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolConfigTaints{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolConfigTaints
	}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])
	r.Effect = flattenAwsNodePoolConfigTaintsEffectEnum(m["effect"])

	return r
}

// expandAwsNodePoolConfigSshConfigMap expands the contents of AwsNodePoolConfigSshConfig into a JSON
// request object.
func expandAwsNodePoolConfigSshConfigMap(c *Client, f map[string]AwsNodePoolConfigSshConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolConfigSshConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolConfigSshConfigSlice expands the contents of AwsNodePoolConfigSshConfig into a JSON
// request object.
func expandAwsNodePoolConfigSshConfigSlice(c *Client, f []AwsNodePoolConfigSshConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolConfigSshConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolConfigSshConfigMap flattens the contents of AwsNodePoolConfigSshConfig from a JSON
// response object.
func flattenAwsNodePoolConfigSshConfigMap(c *Client, i interface{}) map[string]AwsNodePoolConfigSshConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfigSshConfig{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfigSshConfig{}
	}

	items := make(map[string]AwsNodePoolConfigSshConfig)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfigSshConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigSshConfigSlice flattens the contents of AwsNodePoolConfigSshConfig from a JSON
// response object.
func flattenAwsNodePoolConfigSshConfigSlice(c *Client, i interface{}) []AwsNodePoolConfigSshConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfigSshConfig{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfigSshConfig{}
	}

	items := make([]AwsNodePoolConfigSshConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfigSshConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolConfigSshConfig expands an instance of AwsNodePoolConfigSshConfig into a JSON
// request object.
func expandAwsNodePoolConfigSshConfig(c *Client, f *AwsNodePoolConfigSshConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Ec2KeyPair; !dcl.IsEmptyValueIndirect(v) {
		m["ec2KeyPair"] = v
	}

	return m, nil
}

// flattenAwsNodePoolConfigSshConfig flattens an instance of AwsNodePoolConfigSshConfig from a JSON
// response object.
func flattenAwsNodePoolConfigSshConfig(c *Client, i interface{}) *AwsNodePoolConfigSshConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolConfigSshConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolConfigSshConfig
	}
	r.Ec2KeyPair = dcl.FlattenString(m["ec2KeyPair"])

	return r
}

// expandAwsNodePoolAutoscalingMap expands the contents of AwsNodePoolAutoscaling into a JSON
// request object.
func expandAwsNodePoolAutoscalingMap(c *Client, f map[string]AwsNodePoolAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolAutoscalingSlice expands the contents of AwsNodePoolAutoscaling into a JSON
// request object.
func expandAwsNodePoolAutoscalingSlice(c *Client, f []AwsNodePoolAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolAutoscalingMap flattens the contents of AwsNodePoolAutoscaling from a JSON
// response object.
func flattenAwsNodePoolAutoscalingMap(c *Client, i interface{}) map[string]AwsNodePoolAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolAutoscaling{}
	}

	items := make(map[string]AwsNodePoolAutoscaling)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolAutoscalingSlice flattens the contents of AwsNodePoolAutoscaling from a JSON
// response object.
func flattenAwsNodePoolAutoscalingSlice(c *Client, i interface{}) []AwsNodePoolAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return []AwsNodePoolAutoscaling{}
	}

	items := make([]AwsNodePoolAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolAutoscaling expands an instance of AwsNodePoolAutoscaling into a JSON
// request object.
func expandAwsNodePoolAutoscaling(c *Client, f *AwsNodePoolAutoscaling) (map[string]interface{}, error) {
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

// flattenAwsNodePoolAutoscaling flattens an instance of AwsNodePoolAutoscaling from a JSON
// response object.
func flattenAwsNodePoolAutoscaling(c *Client, i interface{}) *AwsNodePoolAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolAutoscaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolAutoscaling
	}
	r.MinNodeCount = dcl.FlattenInteger(m["minNodeCount"])
	r.MaxNodeCount = dcl.FlattenInteger(m["maxNodeCount"])

	return r
}

// expandAwsNodePoolMaxPodsConstraintMap expands the contents of AwsNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAwsNodePoolMaxPodsConstraintMap(c *Client, f map[string]AwsNodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAwsNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAwsNodePoolMaxPodsConstraintSlice expands the contents of AwsNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAwsNodePoolMaxPodsConstraintSlice(c *Client, f []AwsNodePoolMaxPodsConstraint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAwsNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAwsNodePoolMaxPodsConstraintMap flattens the contents of AwsNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAwsNodePoolMaxPodsConstraintMap(c *Client, i interface{}) map[string]AwsNodePoolMaxPodsConstraint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolMaxPodsConstraint{}
	}

	items := make(map[string]AwsNodePoolMaxPodsConstraint)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolMaxPodsConstraint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAwsNodePoolMaxPodsConstraintSlice flattens the contents of AwsNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAwsNodePoolMaxPodsConstraintSlice(c *Client, i interface{}) []AwsNodePoolMaxPodsConstraint {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return []AwsNodePoolMaxPodsConstraint{}
	}

	items := make([]AwsNodePoolMaxPodsConstraint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolMaxPodsConstraint(c, item.(map[string]interface{})))
	}

	return items
}

// expandAwsNodePoolMaxPodsConstraint expands an instance of AwsNodePoolMaxPodsConstraint into a JSON
// request object.
func expandAwsNodePoolMaxPodsConstraint(c *Client, f *AwsNodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxPodsPerNode; !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsPerNode"] = v
	}

	return m, nil
}

// flattenAwsNodePoolMaxPodsConstraint flattens an instance of AwsNodePoolMaxPodsConstraint from a JSON
// response object.
func flattenAwsNodePoolMaxPodsConstraint(c *Client, i interface{}) *AwsNodePoolMaxPodsConstraint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AwsNodePoolMaxPodsConstraint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAwsNodePoolMaxPodsConstraint
	}
	r.MaxPodsPerNode = dcl.FlattenInteger(m["maxPodsPerNode"])

	return r
}

// flattenAwsNodePoolConfigRootVolumeVolumeTypeEnumMap flattens the contents of AwsNodePoolConfigRootVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsNodePoolConfigRootVolumeVolumeTypeEnumMap(c *Client, i interface{}) map[string]AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfigRootVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfigRootVolumeVolumeTypeEnum{}
	}

	items := make(map[string]AwsNodePoolConfigRootVolumeVolumeTypeEnum)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfigRootVolumeVolumeTypeEnum(item.(interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigRootVolumeVolumeTypeEnumSlice flattens the contents of AwsNodePoolConfigRootVolumeVolumeTypeEnum from a JSON
// response object.
func flattenAwsNodePoolConfigRootVolumeVolumeTypeEnumSlice(c *Client, i interface{}) []AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfigRootVolumeVolumeTypeEnum{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfigRootVolumeVolumeTypeEnum{}
	}

	items := make([]AwsNodePoolConfigRootVolumeVolumeTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfigRootVolumeVolumeTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAwsNodePoolConfigRootVolumeVolumeTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AwsNodePoolConfigRootVolumeVolumeTypeEnum with the same value as that string.
func flattenAwsNodePoolConfigRootVolumeVolumeTypeEnum(i interface{}) *AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AwsNodePoolConfigRootVolumeVolumeTypeEnumRef("")
	}

	return AwsNodePoolConfigRootVolumeVolumeTypeEnumRef(s)
}

// flattenAwsNodePoolConfigTaintsEffectEnumMap flattens the contents of AwsNodePoolConfigTaintsEffectEnum from a JSON
// response object.
func flattenAwsNodePoolConfigTaintsEffectEnumMap(c *Client, i interface{}) map[string]AwsNodePoolConfigTaintsEffectEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolConfigTaintsEffectEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolConfigTaintsEffectEnum{}
	}

	items := make(map[string]AwsNodePoolConfigTaintsEffectEnum)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolConfigTaintsEffectEnum(item.(interface{}))
	}

	return items
}

// flattenAwsNodePoolConfigTaintsEffectEnumSlice flattens the contents of AwsNodePoolConfigTaintsEffectEnum from a JSON
// response object.
func flattenAwsNodePoolConfigTaintsEffectEnumSlice(c *Client, i interface{}) []AwsNodePoolConfigTaintsEffectEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolConfigTaintsEffectEnum{}
	}

	if len(a) == 0 {
		return []AwsNodePoolConfigTaintsEffectEnum{}
	}

	items := make([]AwsNodePoolConfigTaintsEffectEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolConfigTaintsEffectEnum(item.(interface{})))
	}

	return items
}

// flattenAwsNodePoolConfigTaintsEffectEnum asserts that an interface is a string, and returns a
// pointer to a *AwsNodePoolConfigTaintsEffectEnum with the same value as that string.
func flattenAwsNodePoolConfigTaintsEffectEnum(i interface{}) *AwsNodePoolConfigTaintsEffectEnum {
	s, ok := i.(string)
	if !ok {
		return AwsNodePoolConfigTaintsEffectEnumRef("")
	}

	return AwsNodePoolConfigTaintsEffectEnumRef(s)
}

// flattenAwsNodePoolStateEnumMap flattens the contents of AwsNodePoolStateEnum from a JSON
// response object.
func flattenAwsNodePoolStateEnumMap(c *Client, i interface{}) map[string]AwsNodePoolStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AwsNodePoolStateEnum{}
	}

	if len(a) == 0 {
		return map[string]AwsNodePoolStateEnum{}
	}

	items := make(map[string]AwsNodePoolStateEnum)
	for k, item := range a {
		items[k] = *flattenAwsNodePoolStateEnum(item.(interface{}))
	}

	return items
}

// flattenAwsNodePoolStateEnumSlice flattens the contents of AwsNodePoolStateEnum from a JSON
// response object.
func flattenAwsNodePoolStateEnumSlice(c *Client, i interface{}) []AwsNodePoolStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AwsNodePoolStateEnum{}
	}

	if len(a) == 0 {
		return []AwsNodePoolStateEnum{}
	}

	items := make([]AwsNodePoolStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAwsNodePoolStateEnum(item.(interface{})))
	}

	return items
}

// flattenAwsNodePoolStateEnum asserts that an interface is a string, and returns a
// pointer to a *AwsNodePoolStateEnum with the same value as that string.
func flattenAwsNodePoolStateEnum(i interface{}) *AwsNodePoolStateEnum {
	s, ok := i.(string)
	if !ok {
		return AwsNodePoolStateEnumRef("")
	}

	return AwsNodePoolStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AwsNodePool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAwsNodePool(b, c)
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
		if nr.AwsCluster == nil && ncr.AwsCluster == nil {
			c.Config.Logger.Info("Both AwsCluster fields null - considering equal.")
		} else if nr.AwsCluster == nil || ncr.AwsCluster == nil {
			c.Config.Logger.Info("Only one AwsCluster field is null - considering unequal.")
			return false
		} else if *nr.AwsCluster != *ncr.AwsCluster {
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

type awsNodePoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         awsNodePoolApiOperation
}

func convertFieldDiffsToAwsNodePoolDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]awsNodePoolDiff, error) {
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
	var diffs []awsNodePoolDiff
	// For each operation name, create a awsNodePoolDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := awsNodePoolDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAwsNodePoolApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAwsNodePoolApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (awsNodePoolApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
