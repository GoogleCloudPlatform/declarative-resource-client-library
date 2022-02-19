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
package dataproc

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

func (r *Cluster) validate() error {

	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Metrics) {
		if err := r.Metrics.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.GceClusterConfig) {
		if err := r.GceClusterConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MasterConfig) {
		if err := r.MasterConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WorkerConfig) {
		if err := r.WorkerConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryWorkerConfig) {
		if err := r.SecondaryWorkerConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SoftwareConfig) {
		if err := r.SoftwareConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionConfig) {
		if err := r.EncryptionConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AutoscalingConfig) {
		if err := r.AutoscalingConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecurityConfig) {
		if err := r.SecurityConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LifecycleConfig) {
		if err := r.LifecycleConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EndpointConfig) {
		if err := r.EndpointConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigGceClusterConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReservationAffinity) {
		if err := r.ReservationAffinity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NodeGroupAffinity) {
		if err := r.NodeGroupAffinity.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigGceClusterConfigReservationAffinity) validate() error {
	return nil
}
func (r *ClusterConfigGceClusterConfigNodeGroupAffinity) validate() error {
	if err := dcl.Required(r, "nodeGroup"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterConfigMasterConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DiskConfig) {
		if err := r.DiskConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagedGroupConfig) {
		if err := r.ManagedGroupConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigMasterConfigDiskConfig) validate() error {
	return nil
}
func (r *ClusterConfigMasterConfigManagedGroupConfig) validate() error {
	return nil
}
func (r *ClusterConfigMasterConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterConfigWorkerConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DiskConfig) {
		if err := r.DiskConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagedGroupConfig) {
		if err := r.ManagedGroupConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigWorkerConfigDiskConfig) validate() error {
	return nil
}
func (r *ClusterConfigWorkerConfigManagedGroupConfig) validate() error {
	return nil
}
func (r *ClusterConfigWorkerConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterConfigSecondaryWorkerConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DiskConfig) {
		if err := r.DiskConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagedGroupConfig) {
		if err := r.ManagedGroupConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigSecondaryWorkerConfigDiskConfig) validate() error {
	return nil
}
func (r *ClusterConfigSecondaryWorkerConfigManagedGroupConfig) validate() error {
	return nil
}
func (r *ClusterConfigSecondaryWorkerConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterConfigSoftwareConfig) validate() error {
	return nil
}
func (r *ClusterConfigInitializationActions) validate() error {
	if err := dcl.Required(r, "executableFile"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterConfigEncryptionConfig) validate() error {
	return nil
}
func (r *ClusterConfigAutoscalingConfig) validate() error {
	return nil
}
func (r *ClusterConfigSecurityConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KerberosConfig) {
		if err := r.KerberosConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterConfigSecurityConfigKerberosConfig) validate() error {
	return nil
}
func (r *ClusterConfigLifecycleConfig) validate() error {
	return nil
}
func (r *ClusterConfigEndpointConfig) validate() error {
	return nil
}
func (r *ClusterStatus) validate() error {
	return nil
}
func (r *ClusterStatusHistory) validate() error {
	return nil
}
func (r *ClusterMetrics) validate() error {
	return nil
}
func (r *Cluster) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://dataproc.googleapis.com/v1/", params)
}

func (r *Cluster) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Cluster) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Cluster) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Cluster) SetPolicyVerb() string {
	return "POST"
}

func (r *Cluster) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Cluster) IAMPolicyVersion() int {
	return 3
}

// clusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clusterApiOperation interface {
	do(context.Context, *Cluster, *Client) error
}

// newUpdateClusterUpdateClusterRequest creates a request for an
// Cluster resource's UpdateCluster update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateClusterRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateClusterRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateClusterRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateClusterUpdateClusterOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCluster")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateClusterRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateClusterRequest(c, req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listClusterRaw(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ClusterMaxPage {
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

type listClusterOperation struct {
	Clusters []map[string]interface{} `json:"clusters"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listCluster(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]*Cluster, string, error) {
	b, err := c.listClusterRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Cluster
	for _, v := range m.Clusters {
		res, err := unmarshalMapCluster(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCluster(ctx context.Context, f func(*Cluster) bool, resources []*Cluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCluster(ctx, res)
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

type deleteClusterOperation struct{}

func (op *deleteClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	r, err := c.GetCluster(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Cluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetCluster checking for existence. error: %v", err)
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
		_, err = c.GetCluster(ctx, r)
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
type createClusterOperation struct {
	response map[string]interface{}
}

func (op *createClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
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

	if _, err := c.GetCluster(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getClusterRaw(ctx context.Context, r *Cluster) ([]byte, error) {

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

func (c *Client) clusterDiffsForRawDesired(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (initial, desired *Cluster, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Cluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCluster(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Cluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Cluster resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Cluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Cluster: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Cluster: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractClusterFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Cluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Cluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeClusterInitialState(rawInitial, rawDesired *Cluster) (*Cluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClusterDesiredState(rawDesired, rawInitial *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeClusterConfig(rawDesired.Config, nil, opts...)
		rawDesired.Status = canonicalizeClusterStatus(rawDesired.Status, nil, opts...)
		rawDesired.Metrics = canonicalizeClusterMetrics(rawDesired.Metrics, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Cluster{}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.Config = canonicalizeClusterConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	if dcl.IsNotReturnedByServer(rawNew.Project) && dcl.IsNotReturnedByServer(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Config) && dcl.IsNotReturnedByServer(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewClusterConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsNotReturnedByServer(rawNew.Labels) && dcl.IsNotReturnedByServer(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Status) && dcl.IsNotReturnedByServer(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewClusterStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsNotReturnedByServer(rawNew.StatusHistory) && dcl.IsNotReturnedByServer(rawDesired.StatusHistory) {
		rawNew.StatusHistory = rawDesired.StatusHistory
	} else {
		rawNew.StatusHistory = canonicalizeNewClusterStatusHistorySlice(c, rawDesired.StatusHistory, rawNew.StatusHistory)
	}

	if dcl.IsNotReturnedByServer(rawNew.ClusterUuid) && dcl.IsNotReturnedByServer(rawDesired.ClusterUuid) {
		rawNew.ClusterUuid = rawDesired.ClusterUuid
	} else {
		if dcl.StringCanonicalize(rawDesired.ClusterUuid, rawNew.ClusterUuid) {
			rawNew.ClusterUuid = rawDesired.ClusterUuid
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Metrics) && dcl.IsNotReturnedByServer(rawDesired.Metrics) {
		rawNew.Metrics = rawDesired.Metrics
	} else {
		rawNew.Metrics = canonicalizeNewClusterMetrics(c, rawDesired.Metrics, rawNew.Metrics)
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeClusterConfig(des, initial *ClusterConfig, opts ...dcl.ApplyOption) *ClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfig{}

	if dcl.NameToSelfLink(des.StagingBucket, initial.StagingBucket) || dcl.IsZeroValue(des.StagingBucket) {
		cDes.StagingBucket = initial.StagingBucket
	} else {
		cDes.StagingBucket = des.StagingBucket
	}
	if dcl.NameToSelfLink(des.TempBucket, initial.TempBucket) || dcl.IsZeroValue(des.TempBucket) {
		cDes.TempBucket = initial.TempBucket
	} else {
		cDes.TempBucket = des.TempBucket
	}
	cDes.GceClusterConfig = canonicalizeClusterConfigGceClusterConfig(des.GceClusterConfig, initial.GceClusterConfig, opts...)
	cDes.MasterConfig = canonicalizeClusterConfigMasterConfig(des.MasterConfig, initial.MasterConfig, opts...)
	cDes.WorkerConfig = canonicalizeClusterConfigWorkerConfig(des.WorkerConfig, initial.WorkerConfig, opts...)
	cDes.SecondaryWorkerConfig = canonicalizeClusterConfigSecondaryWorkerConfig(des.SecondaryWorkerConfig, initial.SecondaryWorkerConfig, opts...)
	cDes.SoftwareConfig = canonicalizeClusterConfigSoftwareConfig(des.SoftwareConfig, initial.SoftwareConfig, opts...)
	cDes.InitializationActions = canonicalizeClusterConfigInitializationActionsSlice(des.InitializationActions, initial.InitializationActions, opts...)
	cDes.EncryptionConfig = canonicalizeClusterConfigEncryptionConfig(des.EncryptionConfig, initial.EncryptionConfig, opts...)
	cDes.AutoscalingConfig = canonicalizeClusterConfigAutoscalingConfig(des.AutoscalingConfig, initial.AutoscalingConfig, opts...)
	cDes.SecurityConfig = canonicalizeClusterConfigSecurityConfig(des.SecurityConfig, initial.SecurityConfig, opts...)
	cDes.LifecycleConfig = canonicalizeClusterConfigLifecycleConfig(des.LifecycleConfig, initial.LifecycleConfig, opts...)
	cDes.EndpointConfig = canonicalizeClusterConfigEndpointConfig(des.EndpointConfig, initial.EndpointConfig, opts...)

	return cDes
}

func canonicalizeClusterConfigSlice(des, initial []ClusterConfig, opts ...dcl.ApplyOption) []ClusterConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfig(c *Client, des, nw *ClusterConfig) *ClusterConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.StagingBucket, nw.StagingBucket) {
		nw.StagingBucket = des.StagingBucket
	}
	if dcl.NameToSelfLink(des.TempBucket, nw.TempBucket) {
		nw.TempBucket = des.TempBucket
	}
	nw.GceClusterConfig = canonicalizeNewClusterConfigGceClusterConfig(c, des.GceClusterConfig, nw.GceClusterConfig)
	nw.MasterConfig = canonicalizeNewClusterConfigMasterConfig(c, des.MasterConfig, nw.MasterConfig)
	nw.WorkerConfig = canonicalizeNewClusterConfigWorkerConfig(c, des.WorkerConfig, nw.WorkerConfig)
	nw.SecondaryWorkerConfig = canonicalizeNewClusterConfigSecondaryWorkerConfig(c, des.SecondaryWorkerConfig, nw.SecondaryWorkerConfig)
	nw.SoftwareConfig = canonicalizeNewClusterConfigSoftwareConfig(c, des.SoftwareConfig, nw.SoftwareConfig)
	nw.InitializationActions = canonicalizeNewClusterConfigInitializationActionsSlice(c, des.InitializationActions, nw.InitializationActions)
	nw.EncryptionConfig = canonicalizeNewClusterConfigEncryptionConfig(c, des.EncryptionConfig, nw.EncryptionConfig)
	nw.AutoscalingConfig = canonicalizeNewClusterConfigAutoscalingConfig(c, des.AutoscalingConfig, nw.AutoscalingConfig)
	nw.SecurityConfig = canonicalizeNewClusterConfigSecurityConfig(c, des.SecurityConfig, nw.SecurityConfig)
	nw.LifecycleConfig = canonicalizeNewClusterConfigLifecycleConfig(c, des.LifecycleConfig, nw.LifecycleConfig)
	nw.EndpointConfig = canonicalizeNewClusterConfigEndpointConfig(c, des.EndpointConfig, nw.EndpointConfig)

	return nw
}

func canonicalizeNewClusterConfigSet(c *Client, des, nw []ClusterConfig) []ClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSlice(c *Client, des, nw []ClusterConfig) []ClusterConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigGceClusterConfig(des, initial *ClusterConfigGceClusterConfig, opts ...dcl.ApplyOption) *ClusterConfigGceClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigGceClusterConfig{}

	if dcl.StringCanonicalize(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		cDes.Zone = initial.Zone
	} else {
		cDes.Zone = des.Zone
	}
	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, initial.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		cDes.Subnetwork = initial.Subnetwork
	} else {
		cDes.Subnetwork = des.Subnetwork
	}
	if dcl.BoolCanonicalize(des.InternalIPOnly, initial.InternalIPOnly) || dcl.IsZeroValue(des.InternalIPOnly) {
		cDes.InternalIPOnly = initial.InternalIPOnly
	} else {
		cDes.InternalIPOnly = des.InternalIPOnly
	}
	if dcl.IsZeroValue(des.PrivateIPv6GoogleAccess) || (dcl.IsEmptyValueIndirect(des.PrivateIPv6GoogleAccess) && dcl.IsEmptyValueIndirect(initial.PrivateIPv6GoogleAccess)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PrivateIPv6GoogleAccess = initial.PrivateIPv6GoogleAccess
	} else {
		cDes.PrivateIPv6GoogleAccess = des.PrivateIPv6GoogleAccess
	}
	if dcl.NameToSelfLink(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringArrayCanonicalize(des.ServiceAccountScopes, initial.ServiceAccountScopes) {
		cDes.ServiceAccountScopes = initial.ServiceAccountScopes
	} else {
		cDes.ServiceAccountScopes = des.ServiceAccountScopes
	}
	if dcl.StringArrayCanonicalize(des.Tags, initial.Tags) {
		cDes.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	if dcl.IsZeroValue(des.Metadata) || (dcl.IsEmptyValueIndirect(des.Metadata) && dcl.IsEmptyValueIndirect(initial.Metadata)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Metadata = initial.Metadata
	} else {
		cDes.Metadata = des.Metadata
	}
	cDes.ReservationAffinity = canonicalizeClusterConfigGceClusterConfigReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	cDes.NodeGroupAffinity = canonicalizeClusterConfigGceClusterConfigNodeGroupAffinity(des.NodeGroupAffinity, initial.NodeGroupAffinity, opts...)

	return cDes
}

func canonicalizeClusterConfigGceClusterConfigSlice(des, initial []ClusterConfigGceClusterConfig, opts ...dcl.ApplyOption) []ClusterConfigGceClusterConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigGceClusterConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigGceClusterConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigGceClusterConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigGceClusterConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigGceClusterConfig(c *Client, des, nw *ClusterConfigGceClusterConfig) *ClusterConfigGceClusterConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigGceClusterConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Zone, nw.Zone) {
		nw.Zone = des.Zone
	}
	if dcl.NameToSelfLink(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, nw.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}
	if dcl.BoolCanonicalize(des.InternalIPOnly, nw.InternalIPOnly) {
		nw.InternalIPOnly = des.InternalIPOnly
	}
	if dcl.NameToSelfLink(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringArrayCanonicalize(des.ServiceAccountScopes, nw.ServiceAccountScopes) {
		nw.ServiceAccountScopes = des.ServiceAccountScopes
	}
	if dcl.StringArrayCanonicalize(des.Tags, nw.Tags) {
		nw.Tags = des.Tags
	}
	nw.ReservationAffinity = canonicalizeNewClusterConfigGceClusterConfigReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.NodeGroupAffinity = canonicalizeNewClusterConfigGceClusterConfigNodeGroupAffinity(c, des.NodeGroupAffinity, nw.NodeGroupAffinity)

	return nw
}

func canonicalizeNewClusterConfigGceClusterConfigSet(c *Client, des, nw []ClusterConfigGceClusterConfig) []ClusterConfigGceClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigGceClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigGceClusterConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigGceClusterConfigSlice(c *Client, des, nw []ClusterConfigGceClusterConfig) []ClusterConfigGceClusterConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigGceClusterConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigGceClusterConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigGceClusterConfigReservationAffinity(des, initial *ClusterConfigGceClusterConfigReservationAffinity, opts ...dcl.ApplyOption) *ClusterConfigGceClusterConfigReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigGceClusterConfigReservationAffinity{}

	if dcl.IsZeroValue(des.ConsumeReservationType) || (dcl.IsEmptyValueIndirect(des.ConsumeReservationType) && dcl.IsEmptyValueIndirect(initial.ConsumeReservationType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ConsumeReservationType = initial.ConsumeReservationType
	} else {
		cDes.ConsumeReservationType = des.ConsumeReservationType
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		cDes.Key = initial.Key
	} else {
		cDes.Key = des.Key
	}
	if dcl.StringArrayCanonicalize(des.Values, initial.Values) {
		cDes.Values = initial.Values
	} else {
		cDes.Values = des.Values
	}

	return cDes
}

func canonicalizeClusterConfigGceClusterConfigReservationAffinitySlice(des, initial []ClusterConfigGceClusterConfigReservationAffinity, opts ...dcl.ApplyOption) []ClusterConfigGceClusterConfigReservationAffinity {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigGceClusterConfigReservationAffinity, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigGceClusterConfigReservationAffinity(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigGceClusterConfigReservationAffinity, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigGceClusterConfigReservationAffinity(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigGceClusterConfigReservationAffinity(c *Client, des, nw *ClusterConfigGceClusterConfigReservationAffinity) *ClusterConfigGceClusterConfigReservationAffinity {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigGceClusterConfigReservationAffinity while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringArrayCanonicalize(des.Values, nw.Values) {
		nw.Values = des.Values
	}

	return nw
}

func canonicalizeNewClusterConfigGceClusterConfigReservationAffinitySet(c *Client, des, nw []ClusterConfigGceClusterConfigReservationAffinity) []ClusterConfigGceClusterConfigReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigGceClusterConfigReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigGceClusterConfigReservationAffinityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, des, nw []ClusterConfigGceClusterConfigReservationAffinity) []ClusterConfigGceClusterConfigReservationAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigGceClusterConfigReservationAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigGceClusterConfigReservationAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigGceClusterConfigNodeGroupAffinity(des, initial *ClusterConfigGceClusterConfigNodeGroupAffinity, opts ...dcl.ApplyOption) *ClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigGceClusterConfigNodeGroupAffinity{}

	if dcl.NameToSelfLink(des.NodeGroup, initial.NodeGroup) || dcl.IsZeroValue(des.NodeGroup) {
		cDes.NodeGroup = initial.NodeGroup
	} else {
		cDes.NodeGroup = des.NodeGroup
	}

	return cDes
}

func canonicalizeClusterConfigGceClusterConfigNodeGroupAffinitySlice(des, initial []ClusterConfigGceClusterConfigNodeGroupAffinity, opts ...dcl.ApplyOption) []ClusterConfigGceClusterConfigNodeGroupAffinity {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigGceClusterConfigNodeGroupAffinity, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigGceClusterConfigNodeGroupAffinity(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigGceClusterConfigNodeGroupAffinity, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigGceClusterConfigNodeGroupAffinity(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, des, nw *ClusterConfigGceClusterConfigNodeGroupAffinity) *ClusterConfigGceClusterConfigNodeGroupAffinity {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigGceClusterConfigNodeGroupAffinity while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.NodeGroup, nw.NodeGroup) {
		nw.NodeGroup = des.NodeGroup
	}

	return nw
}

func canonicalizeNewClusterConfigGceClusterConfigNodeGroupAffinitySet(c *Client, des, nw []ClusterConfigGceClusterConfigNodeGroupAffinity) []ClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigGceClusterConfigNodeGroupAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigGceClusterConfigNodeGroupAffinityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, des, nw []ClusterConfigGceClusterConfigNodeGroupAffinity) []ClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigGceClusterConfigNodeGroupAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigGceClusterConfigNodeGroupAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigMasterConfig(des, initial *ClusterConfigMasterConfig, opts ...dcl.ApplyOption) *ClusterConfigMasterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigMasterConfig{}

	if dcl.IsZeroValue(des.NumInstances) || (dcl.IsEmptyValueIndirect(des.NumInstances) && dcl.IsEmptyValueIndirect(initial.NumInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumInstances = initial.NumInstances
	} else {
		cDes.NumInstances = des.NumInstances
	}
	if dcl.NameToSelfLink(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		cDes.Image = initial.Image
	} else {
		cDes.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}
	cDes.DiskConfig = canonicalizeClusterConfigMasterConfigDiskConfig(des.DiskConfig, initial.DiskConfig, opts...)
	if dcl.IsZeroValue(des.Preemptibility) || (dcl.IsEmptyValueIndirect(des.Preemptibility) && dcl.IsEmptyValueIndirect(initial.Preemptibility)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Preemptibility = initial.Preemptibility
	} else {
		cDes.Preemptibility = des.Preemptibility
	}
	cDes.Accelerators = canonicalizeClusterConfigMasterConfigAcceleratorsSlice(des.Accelerators, initial.Accelerators, opts...)
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		cDes.MinCpuPlatform = initial.MinCpuPlatform
	} else {
		cDes.MinCpuPlatform = des.MinCpuPlatform
	}

	return cDes
}

func canonicalizeClusterConfigMasterConfigSlice(des, initial []ClusterConfigMasterConfig, opts ...dcl.ApplyOption) []ClusterConfigMasterConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigMasterConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigMasterConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigMasterConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigMasterConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigMasterConfig(c *Client, des, nw *ClusterConfigMasterConfig) *ClusterConfigMasterConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigMasterConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.InstanceNames, nw.InstanceNames) {
		nw.InstanceNames = des.InstanceNames
	}
	if dcl.NameToSelfLink(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.DiskConfig = canonicalizeNewClusterConfigMasterConfigDiskConfig(c, des.DiskConfig, nw.DiskConfig)
	if dcl.BoolCanonicalize(des.IsPreemptible, nw.IsPreemptible) {
		nw.IsPreemptible = des.IsPreemptible
	}
	nw.ManagedGroupConfig = canonicalizeNewClusterConfigMasterConfigManagedGroupConfig(c, des.ManagedGroupConfig, nw.ManagedGroupConfig)
	nw.Accelerators = canonicalizeNewClusterConfigMasterConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}

	return nw
}

func canonicalizeNewClusterConfigMasterConfigSet(c *Client, des, nw []ClusterConfigMasterConfig) []ClusterConfigMasterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigMasterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigMasterConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigMasterConfigSlice(c *Client, des, nw []ClusterConfigMasterConfig) []ClusterConfigMasterConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigMasterConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigMasterConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigMasterConfigDiskConfig(des, initial *ClusterConfigMasterConfigDiskConfig, opts ...dcl.ApplyOption) *ClusterConfigMasterConfigDiskConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigMasterConfigDiskConfig{}

	if dcl.StringCanonicalize(des.BootDiskType, initial.BootDiskType) || dcl.IsZeroValue(des.BootDiskType) {
		cDes.BootDiskType = initial.BootDiskType
	} else {
		cDes.BootDiskType = des.BootDiskType
	}
	if dcl.IsZeroValue(des.BootDiskSizeGb) || (dcl.IsEmptyValueIndirect(des.BootDiskSizeGb) && dcl.IsEmptyValueIndirect(initial.BootDiskSizeGb)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BootDiskSizeGb = initial.BootDiskSizeGb
	} else {
		cDes.BootDiskSizeGb = des.BootDiskSizeGb
	}
	if dcl.IsZeroValue(des.NumLocalSsds) || (dcl.IsEmptyValueIndirect(des.NumLocalSsds) && dcl.IsEmptyValueIndirect(initial.NumLocalSsds)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumLocalSsds = initial.NumLocalSsds
	} else {
		cDes.NumLocalSsds = des.NumLocalSsds
	}

	return cDes
}

func canonicalizeClusterConfigMasterConfigDiskConfigSlice(des, initial []ClusterConfigMasterConfigDiskConfig, opts ...dcl.ApplyOption) []ClusterConfigMasterConfigDiskConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigMasterConfigDiskConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigMasterConfigDiskConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigMasterConfigDiskConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigMasterConfigDiskConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigMasterConfigDiskConfig(c *Client, des, nw *ClusterConfigMasterConfigDiskConfig) *ClusterConfigMasterConfigDiskConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigMasterConfigDiskConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.BootDiskType, nw.BootDiskType) {
		nw.BootDiskType = des.BootDiskType
	}

	return nw
}

func canonicalizeNewClusterConfigMasterConfigDiskConfigSet(c *Client, des, nw []ClusterConfigMasterConfigDiskConfig) []ClusterConfigMasterConfigDiskConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigMasterConfigDiskConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigMasterConfigDiskConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigMasterConfigDiskConfigSlice(c *Client, des, nw []ClusterConfigMasterConfigDiskConfig) []ClusterConfigMasterConfigDiskConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigMasterConfigDiskConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigMasterConfigDiskConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigMasterConfigManagedGroupConfig(des, initial *ClusterConfigMasterConfigManagedGroupConfig, opts ...dcl.ApplyOption) *ClusterConfigMasterConfigManagedGroupConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigMasterConfigManagedGroupConfig{}

	return cDes
}

func canonicalizeClusterConfigMasterConfigManagedGroupConfigSlice(des, initial []ClusterConfigMasterConfigManagedGroupConfig, opts ...dcl.ApplyOption) []ClusterConfigMasterConfigManagedGroupConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigMasterConfigManagedGroupConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigMasterConfigManagedGroupConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigMasterConfigManagedGroupConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigMasterConfigManagedGroupConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigMasterConfigManagedGroupConfig(c *Client, des, nw *ClusterConfigMasterConfigManagedGroupConfig) *ClusterConfigMasterConfigManagedGroupConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigMasterConfigManagedGroupConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.InstanceTemplateName, nw.InstanceTemplateName) {
		nw.InstanceTemplateName = des.InstanceTemplateName
	}
	if dcl.StringCanonicalize(des.InstanceGroupManagerName, nw.InstanceGroupManagerName) {
		nw.InstanceGroupManagerName = des.InstanceGroupManagerName
	}

	return nw
}

func canonicalizeNewClusterConfigMasterConfigManagedGroupConfigSet(c *Client, des, nw []ClusterConfigMasterConfigManagedGroupConfig) []ClusterConfigMasterConfigManagedGroupConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigMasterConfigManagedGroupConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigMasterConfigManagedGroupConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigMasterConfigManagedGroupConfigSlice(c *Client, des, nw []ClusterConfigMasterConfigManagedGroupConfig) []ClusterConfigMasterConfigManagedGroupConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigMasterConfigManagedGroupConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigMasterConfigManagedGroupConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigMasterConfigAccelerators(des, initial *ClusterConfigMasterConfigAccelerators, opts ...dcl.ApplyOption) *ClusterConfigMasterConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigMasterConfigAccelerators{}

	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
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

func canonicalizeClusterConfigMasterConfigAcceleratorsSlice(des, initial []ClusterConfigMasterConfigAccelerators, opts ...dcl.ApplyOption) []ClusterConfigMasterConfigAccelerators {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigMasterConfigAccelerators, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigMasterConfigAccelerators(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigMasterConfigAccelerators, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigMasterConfigAccelerators(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigMasterConfigAccelerators(c *Client, des, nw *ClusterConfigMasterConfigAccelerators) *ClusterConfigMasterConfigAccelerators {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigMasterConfigAccelerators while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterConfigMasterConfigAcceleratorsSet(c *Client, des, nw []ClusterConfigMasterConfigAccelerators) []ClusterConfigMasterConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigMasterConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigMasterConfigAcceleratorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigMasterConfigAcceleratorsSlice(c *Client, des, nw []ClusterConfigMasterConfigAccelerators) []ClusterConfigMasterConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigMasterConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigMasterConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigWorkerConfig(des, initial *ClusterConfigWorkerConfig, opts ...dcl.ApplyOption) *ClusterConfigWorkerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigWorkerConfig{}

	if dcl.IsZeroValue(des.NumInstances) || (dcl.IsEmptyValueIndirect(des.NumInstances) && dcl.IsEmptyValueIndirect(initial.NumInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumInstances = initial.NumInstances
	} else {
		cDes.NumInstances = des.NumInstances
	}
	if dcl.NameToSelfLink(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		cDes.Image = initial.Image
	} else {
		cDes.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}
	cDes.DiskConfig = canonicalizeClusterConfigWorkerConfigDiskConfig(des.DiskConfig, initial.DiskConfig, opts...)
	if dcl.IsZeroValue(des.Preemptibility) || (dcl.IsEmptyValueIndirect(des.Preemptibility) && dcl.IsEmptyValueIndirect(initial.Preemptibility)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Preemptibility = initial.Preemptibility
	} else {
		cDes.Preemptibility = des.Preemptibility
	}
	cDes.Accelerators = canonicalizeClusterConfigWorkerConfigAcceleratorsSlice(des.Accelerators, initial.Accelerators, opts...)
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		cDes.MinCpuPlatform = initial.MinCpuPlatform
	} else {
		cDes.MinCpuPlatform = des.MinCpuPlatform
	}

	return cDes
}

func canonicalizeClusterConfigWorkerConfigSlice(des, initial []ClusterConfigWorkerConfig, opts ...dcl.ApplyOption) []ClusterConfigWorkerConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigWorkerConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigWorkerConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigWorkerConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigWorkerConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigWorkerConfig(c *Client, des, nw *ClusterConfigWorkerConfig) *ClusterConfigWorkerConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigWorkerConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.InstanceNames, nw.InstanceNames) {
		nw.InstanceNames = des.InstanceNames
	}
	if dcl.NameToSelfLink(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.DiskConfig = canonicalizeNewClusterConfigWorkerConfigDiskConfig(c, des.DiskConfig, nw.DiskConfig)
	if dcl.BoolCanonicalize(des.IsPreemptible, nw.IsPreemptible) {
		nw.IsPreemptible = des.IsPreemptible
	}
	nw.ManagedGroupConfig = canonicalizeNewClusterConfigWorkerConfigManagedGroupConfig(c, des.ManagedGroupConfig, nw.ManagedGroupConfig)
	nw.Accelerators = canonicalizeNewClusterConfigWorkerConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}

	return nw
}

func canonicalizeNewClusterConfigWorkerConfigSet(c *Client, des, nw []ClusterConfigWorkerConfig) []ClusterConfigWorkerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigWorkerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigWorkerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigWorkerConfigSlice(c *Client, des, nw []ClusterConfigWorkerConfig) []ClusterConfigWorkerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigWorkerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigWorkerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigWorkerConfigDiskConfig(des, initial *ClusterConfigWorkerConfigDiskConfig, opts ...dcl.ApplyOption) *ClusterConfigWorkerConfigDiskConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigWorkerConfigDiskConfig{}

	if dcl.StringCanonicalize(des.BootDiskType, initial.BootDiskType) || dcl.IsZeroValue(des.BootDiskType) {
		cDes.BootDiskType = initial.BootDiskType
	} else {
		cDes.BootDiskType = des.BootDiskType
	}
	if dcl.IsZeroValue(des.BootDiskSizeGb) || (dcl.IsEmptyValueIndirect(des.BootDiskSizeGb) && dcl.IsEmptyValueIndirect(initial.BootDiskSizeGb)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BootDiskSizeGb = initial.BootDiskSizeGb
	} else {
		cDes.BootDiskSizeGb = des.BootDiskSizeGb
	}
	if dcl.IsZeroValue(des.NumLocalSsds) || (dcl.IsEmptyValueIndirect(des.NumLocalSsds) && dcl.IsEmptyValueIndirect(initial.NumLocalSsds)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumLocalSsds = initial.NumLocalSsds
	} else {
		cDes.NumLocalSsds = des.NumLocalSsds
	}

	return cDes
}

func canonicalizeClusterConfigWorkerConfigDiskConfigSlice(des, initial []ClusterConfigWorkerConfigDiskConfig, opts ...dcl.ApplyOption) []ClusterConfigWorkerConfigDiskConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigWorkerConfigDiskConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigWorkerConfigDiskConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigWorkerConfigDiskConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigWorkerConfigDiskConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigWorkerConfigDiskConfig(c *Client, des, nw *ClusterConfigWorkerConfigDiskConfig) *ClusterConfigWorkerConfigDiskConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigWorkerConfigDiskConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.BootDiskType, nw.BootDiskType) {
		nw.BootDiskType = des.BootDiskType
	}

	return nw
}

func canonicalizeNewClusterConfigWorkerConfigDiskConfigSet(c *Client, des, nw []ClusterConfigWorkerConfigDiskConfig) []ClusterConfigWorkerConfigDiskConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigWorkerConfigDiskConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigWorkerConfigDiskConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigWorkerConfigDiskConfigSlice(c *Client, des, nw []ClusterConfigWorkerConfigDiskConfig) []ClusterConfigWorkerConfigDiskConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigWorkerConfigDiskConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigWorkerConfigDiskConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigWorkerConfigManagedGroupConfig(des, initial *ClusterConfigWorkerConfigManagedGroupConfig, opts ...dcl.ApplyOption) *ClusterConfigWorkerConfigManagedGroupConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigWorkerConfigManagedGroupConfig{}

	return cDes
}

func canonicalizeClusterConfigWorkerConfigManagedGroupConfigSlice(des, initial []ClusterConfigWorkerConfigManagedGroupConfig, opts ...dcl.ApplyOption) []ClusterConfigWorkerConfigManagedGroupConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigWorkerConfigManagedGroupConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigWorkerConfigManagedGroupConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigWorkerConfigManagedGroupConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigWorkerConfigManagedGroupConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigWorkerConfigManagedGroupConfig(c *Client, des, nw *ClusterConfigWorkerConfigManagedGroupConfig) *ClusterConfigWorkerConfigManagedGroupConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigWorkerConfigManagedGroupConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.InstanceTemplateName, nw.InstanceTemplateName) {
		nw.InstanceTemplateName = des.InstanceTemplateName
	}
	if dcl.StringCanonicalize(des.InstanceGroupManagerName, nw.InstanceGroupManagerName) {
		nw.InstanceGroupManagerName = des.InstanceGroupManagerName
	}

	return nw
}

func canonicalizeNewClusterConfigWorkerConfigManagedGroupConfigSet(c *Client, des, nw []ClusterConfigWorkerConfigManagedGroupConfig) []ClusterConfigWorkerConfigManagedGroupConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigWorkerConfigManagedGroupConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigWorkerConfigManagedGroupConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigWorkerConfigManagedGroupConfigSlice(c *Client, des, nw []ClusterConfigWorkerConfigManagedGroupConfig) []ClusterConfigWorkerConfigManagedGroupConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigWorkerConfigManagedGroupConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigWorkerConfigManagedGroupConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigWorkerConfigAccelerators(des, initial *ClusterConfigWorkerConfigAccelerators, opts ...dcl.ApplyOption) *ClusterConfigWorkerConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigWorkerConfigAccelerators{}

	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
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

func canonicalizeClusterConfigWorkerConfigAcceleratorsSlice(des, initial []ClusterConfigWorkerConfigAccelerators, opts ...dcl.ApplyOption) []ClusterConfigWorkerConfigAccelerators {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigWorkerConfigAccelerators, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigWorkerConfigAccelerators(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigWorkerConfigAccelerators, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigWorkerConfigAccelerators(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigWorkerConfigAccelerators(c *Client, des, nw *ClusterConfigWorkerConfigAccelerators) *ClusterConfigWorkerConfigAccelerators {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigWorkerConfigAccelerators while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterConfigWorkerConfigAcceleratorsSet(c *Client, des, nw []ClusterConfigWorkerConfigAccelerators) []ClusterConfigWorkerConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigWorkerConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigWorkerConfigAcceleratorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigWorkerConfigAcceleratorsSlice(c *Client, des, nw []ClusterConfigWorkerConfigAccelerators) []ClusterConfigWorkerConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigWorkerConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigWorkerConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecondaryWorkerConfig(des, initial *ClusterConfigSecondaryWorkerConfig, opts ...dcl.ApplyOption) *ClusterConfigSecondaryWorkerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecondaryWorkerConfig{}

	if dcl.IsZeroValue(des.NumInstances) || (dcl.IsEmptyValueIndirect(des.NumInstances) && dcl.IsEmptyValueIndirect(initial.NumInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumInstances = initial.NumInstances
	} else {
		cDes.NumInstances = des.NumInstances
	}
	if dcl.NameToSelfLink(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		cDes.Image = initial.Image
	} else {
		cDes.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}
	cDes.DiskConfig = canonicalizeClusterConfigSecondaryWorkerConfigDiskConfig(des.DiskConfig, initial.DiskConfig, opts...)
	if dcl.IsZeroValue(des.Preemptibility) || (dcl.IsEmptyValueIndirect(des.Preemptibility) && dcl.IsEmptyValueIndirect(initial.Preemptibility)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Preemptibility = initial.Preemptibility
	} else {
		cDes.Preemptibility = des.Preemptibility
	}
	cDes.Accelerators = canonicalizeClusterConfigSecondaryWorkerConfigAcceleratorsSlice(des.Accelerators, initial.Accelerators, opts...)
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		cDes.MinCpuPlatform = initial.MinCpuPlatform
	} else {
		cDes.MinCpuPlatform = des.MinCpuPlatform
	}

	return cDes
}

func canonicalizeClusterConfigSecondaryWorkerConfigSlice(des, initial []ClusterConfigSecondaryWorkerConfig, opts ...dcl.ApplyOption) []ClusterConfigSecondaryWorkerConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecondaryWorkerConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecondaryWorkerConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecondaryWorkerConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecondaryWorkerConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecondaryWorkerConfig(c *Client, des, nw *ClusterConfigSecondaryWorkerConfig) *ClusterConfigSecondaryWorkerConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecondaryWorkerConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.InstanceNames, nw.InstanceNames) {
		nw.InstanceNames = des.InstanceNames
	}
	if dcl.NameToSelfLink(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.DiskConfig = canonicalizeNewClusterConfigSecondaryWorkerConfigDiskConfig(c, des.DiskConfig, nw.DiskConfig)
	if dcl.BoolCanonicalize(des.IsPreemptible, nw.IsPreemptible) {
		nw.IsPreemptible = des.IsPreemptible
	}
	nw.ManagedGroupConfig = canonicalizeNewClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, des.ManagedGroupConfig, nw.ManagedGroupConfig)
	nw.Accelerators = canonicalizeNewClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}

	return nw
}

func canonicalizeNewClusterConfigSecondaryWorkerConfigSet(c *Client, des, nw []ClusterConfigSecondaryWorkerConfig) []ClusterConfigSecondaryWorkerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecondaryWorkerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecondaryWorkerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecondaryWorkerConfigSlice(c *Client, des, nw []ClusterConfigSecondaryWorkerConfig) []ClusterConfigSecondaryWorkerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecondaryWorkerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecondaryWorkerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecondaryWorkerConfigDiskConfig(des, initial *ClusterConfigSecondaryWorkerConfigDiskConfig, opts ...dcl.ApplyOption) *ClusterConfigSecondaryWorkerConfigDiskConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecondaryWorkerConfigDiskConfig{}

	if dcl.StringCanonicalize(des.BootDiskType, initial.BootDiskType) || dcl.IsZeroValue(des.BootDiskType) {
		cDes.BootDiskType = initial.BootDiskType
	} else {
		cDes.BootDiskType = des.BootDiskType
	}
	if dcl.IsZeroValue(des.BootDiskSizeGb) || (dcl.IsEmptyValueIndirect(des.BootDiskSizeGb) && dcl.IsEmptyValueIndirect(initial.BootDiskSizeGb)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BootDiskSizeGb = initial.BootDiskSizeGb
	} else {
		cDes.BootDiskSizeGb = des.BootDiskSizeGb
	}
	if dcl.IsZeroValue(des.NumLocalSsds) || (dcl.IsEmptyValueIndirect(des.NumLocalSsds) && dcl.IsEmptyValueIndirect(initial.NumLocalSsds)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumLocalSsds = initial.NumLocalSsds
	} else {
		cDes.NumLocalSsds = des.NumLocalSsds
	}

	return cDes
}

func canonicalizeClusterConfigSecondaryWorkerConfigDiskConfigSlice(des, initial []ClusterConfigSecondaryWorkerConfigDiskConfig, opts ...dcl.ApplyOption) []ClusterConfigSecondaryWorkerConfigDiskConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecondaryWorkerConfigDiskConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecondaryWorkerConfigDiskConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecondaryWorkerConfigDiskConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecondaryWorkerConfigDiskConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecondaryWorkerConfigDiskConfig(c *Client, des, nw *ClusterConfigSecondaryWorkerConfigDiskConfig) *ClusterConfigSecondaryWorkerConfigDiskConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecondaryWorkerConfigDiskConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.BootDiskType, nw.BootDiskType) {
		nw.BootDiskType = des.BootDiskType
	}

	return nw
}

func canonicalizeNewClusterConfigSecondaryWorkerConfigDiskConfigSet(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigDiskConfig) []ClusterConfigSecondaryWorkerConfigDiskConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecondaryWorkerConfigDiskConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecondaryWorkerConfigDiskConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecondaryWorkerConfigDiskConfigSlice(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigDiskConfig) []ClusterConfigSecondaryWorkerConfigDiskConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecondaryWorkerConfigDiskConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecondaryWorkerConfigDiskConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecondaryWorkerConfigManagedGroupConfig(des, initial *ClusterConfigSecondaryWorkerConfigManagedGroupConfig, opts ...dcl.ApplyOption) *ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}

	return cDes
}

func canonicalizeClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice(des, initial []ClusterConfigSecondaryWorkerConfigManagedGroupConfig, opts ...dcl.ApplyOption) []ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecondaryWorkerConfigManagedGroupConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecondaryWorkerConfigManagedGroupConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecondaryWorkerConfigManagedGroupConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecondaryWorkerConfigManagedGroupConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecondaryWorkerConfigManagedGroupConfig(c *Client, des, nw *ClusterConfigSecondaryWorkerConfigManagedGroupConfig) *ClusterConfigSecondaryWorkerConfigManagedGroupConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecondaryWorkerConfigManagedGroupConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.InstanceTemplateName, nw.InstanceTemplateName) {
		nw.InstanceTemplateName = des.InstanceTemplateName
	}
	if dcl.StringCanonicalize(des.InstanceGroupManagerName, nw.InstanceGroupManagerName) {
		nw.InstanceGroupManagerName = des.InstanceGroupManagerName
	}

	return nw
}

func canonicalizeNewClusterConfigSecondaryWorkerConfigManagedGroupConfigSet(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigManagedGroupConfig) []ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecondaryWorkerConfigManagedGroupConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecondaryWorkerConfigManagedGroupConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigManagedGroupConfig) []ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecondaryWorkerConfigManagedGroupConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecondaryWorkerConfigAccelerators(des, initial *ClusterConfigSecondaryWorkerConfigAccelerators, opts ...dcl.ApplyOption) *ClusterConfigSecondaryWorkerConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecondaryWorkerConfigAccelerators{}

	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
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

func canonicalizeClusterConfigSecondaryWorkerConfigAcceleratorsSlice(des, initial []ClusterConfigSecondaryWorkerConfigAccelerators, opts ...dcl.ApplyOption) []ClusterConfigSecondaryWorkerConfigAccelerators {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecondaryWorkerConfigAccelerators, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecondaryWorkerConfigAccelerators(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecondaryWorkerConfigAccelerators, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecondaryWorkerConfigAccelerators(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecondaryWorkerConfigAccelerators(c *Client, des, nw *ClusterConfigSecondaryWorkerConfigAccelerators) *ClusterConfigSecondaryWorkerConfigAccelerators {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecondaryWorkerConfigAccelerators while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterConfigSecondaryWorkerConfigAcceleratorsSet(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigAccelerators) []ClusterConfigSecondaryWorkerConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecondaryWorkerConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecondaryWorkerConfigAcceleratorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c *Client, des, nw []ClusterConfigSecondaryWorkerConfigAccelerators) []ClusterConfigSecondaryWorkerConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecondaryWorkerConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecondaryWorkerConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSoftwareConfig(des, initial *ClusterConfigSoftwareConfig, opts ...dcl.ApplyOption) *ClusterConfigSoftwareConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSoftwareConfig{}

	if dcl.StringCanonicalize(des.ImageVersion, initial.ImageVersion) || dcl.IsZeroValue(des.ImageVersion) {
		cDes.ImageVersion = initial.ImageVersion
	} else {
		cDes.ImageVersion = des.ImageVersion
	}
	if dcl.IsZeroValue(des.Properties) || (dcl.IsEmptyValueIndirect(des.Properties) && dcl.IsEmptyValueIndirect(initial.Properties)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Properties = initial.Properties
	} else {
		cDes.Properties = des.Properties
	}
	if dcl.IsZeroValue(des.OptionalComponents) || (dcl.IsEmptyValueIndirect(des.OptionalComponents) && dcl.IsEmptyValueIndirect(initial.OptionalComponents)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.OptionalComponents = initial.OptionalComponents
	} else {
		cDes.OptionalComponents = des.OptionalComponents
	}

	return cDes
}

func canonicalizeClusterConfigSoftwareConfigSlice(des, initial []ClusterConfigSoftwareConfig, opts ...dcl.ApplyOption) []ClusterConfigSoftwareConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSoftwareConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSoftwareConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSoftwareConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSoftwareConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSoftwareConfig(c *Client, des, nw *ClusterConfigSoftwareConfig) *ClusterConfigSoftwareConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSoftwareConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ImageVersion, nw.ImageVersion) {
		nw.ImageVersion = des.ImageVersion
	}

	return nw
}

func canonicalizeNewClusterConfigSoftwareConfigSet(c *Client, des, nw []ClusterConfigSoftwareConfig) []ClusterConfigSoftwareConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSoftwareConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSoftwareConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSoftwareConfigSlice(c *Client, des, nw []ClusterConfigSoftwareConfig) []ClusterConfigSoftwareConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSoftwareConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSoftwareConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigInitializationActions(des, initial *ClusterConfigInitializationActions, opts ...dcl.ApplyOption) *ClusterConfigInitializationActions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigInitializationActions{}

	if dcl.StringCanonicalize(des.ExecutableFile, initial.ExecutableFile) || dcl.IsZeroValue(des.ExecutableFile) {
		cDes.ExecutableFile = initial.ExecutableFile
	} else {
		cDes.ExecutableFile = des.ExecutableFile
	}
	if dcl.StringCanonicalize(des.ExecutionTimeout, initial.ExecutionTimeout) || dcl.IsZeroValue(des.ExecutionTimeout) {
		cDes.ExecutionTimeout = initial.ExecutionTimeout
	} else {
		cDes.ExecutionTimeout = des.ExecutionTimeout
	}

	return cDes
}

func canonicalizeClusterConfigInitializationActionsSlice(des, initial []ClusterConfigInitializationActions, opts ...dcl.ApplyOption) []ClusterConfigInitializationActions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigInitializationActions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigInitializationActions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigInitializationActions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigInitializationActions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigInitializationActions(c *Client, des, nw *ClusterConfigInitializationActions) *ClusterConfigInitializationActions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigInitializationActions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ExecutableFile, nw.ExecutableFile) {
		nw.ExecutableFile = des.ExecutableFile
	}
	if dcl.StringCanonicalize(des.ExecutionTimeout, nw.ExecutionTimeout) {
		nw.ExecutionTimeout = des.ExecutionTimeout
	}

	return nw
}

func canonicalizeNewClusterConfigInitializationActionsSet(c *Client, des, nw []ClusterConfigInitializationActions) []ClusterConfigInitializationActions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigInitializationActions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigInitializationActionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigInitializationActionsSlice(c *Client, des, nw []ClusterConfigInitializationActions) []ClusterConfigInitializationActions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigInitializationActions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigInitializationActions(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigEncryptionConfig(des, initial *ClusterConfigEncryptionConfig, opts ...dcl.ApplyOption) *ClusterConfigEncryptionConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigEncryptionConfig{}

	if dcl.NameToSelfLink(des.GcePdKmsKeyName, initial.GcePdKmsKeyName) || dcl.IsZeroValue(des.GcePdKmsKeyName) {
		cDes.GcePdKmsKeyName = initial.GcePdKmsKeyName
	} else {
		cDes.GcePdKmsKeyName = des.GcePdKmsKeyName
	}

	return cDes
}

func canonicalizeClusterConfigEncryptionConfigSlice(des, initial []ClusterConfigEncryptionConfig, opts ...dcl.ApplyOption) []ClusterConfigEncryptionConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigEncryptionConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigEncryptionConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigEncryptionConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigEncryptionConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigEncryptionConfig(c *Client, des, nw *ClusterConfigEncryptionConfig) *ClusterConfigEncryptionConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigEncryptionConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.GcePdKmsKeyName, nw.GcePdKmsKeyName) {
		nw.GcePdKmsKeyName = des.GcePdKmsKeyName
	}

	return nw
}

func canonicalizeNewClusterConfigEncryptionConfigSet(c *Client, des, nw []ClusterConfigEncryptionConfig) []ClusterConfigEncryptionConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigEncryptionConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigEncryptionConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigEncryptionConfigSlice(c *Client, des, nw []ClusterConfigEncryptionConfig) []ClusterConfigEncryptionConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigEncryptionConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigEncryptionConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigAutoscalingConfig(des, initial *ClusterConfigAutoscalingConfig, opts ...dcl.ApplyOption) *ClusterConfigAutoscalingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigAutoscalingConfig{}

	if dcl.NameToSelfLink(des.Policy, initial.Policy) || dcl.IsZeroValue(des.Policy) {
		cDes.Policy = initial.Policy
	} else {
		cDes.Policy = des.Policy
	}

	return cDes
}

func canonicalizeClusterConfigAutoscalingConfigSlice(des, initial []ClusterConfigAutoscalingConfig, opts ...dcl.ApplyOption) []ClusterConfigAutoscalingConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigAutoscalingConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigAutoscalingConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigAutoscalingConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigAutoscalingConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigAutoscalingConfig(c *Client, des, nw *ClusterConfigAutoscalingConfig) *ClusterConfigAutoscalingConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigAutoscalingConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Policy, nw.Policy) {
		nw.Policy = des.Policy
	}

	return nw
}

func canonicalizeNewClusterConfigAutoscalingConfigSet(c *Client, des, nw []ClusterConfigAutoscalingConfig) []ClusterConfigAutoscalingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigAutoscalingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigAutoscalingConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigAutoscalingConfigSlice(c *Client, des, nw []ClusterConfigAutoscalingConfig) []ClusterConfigAutoscalingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigAutoscalingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigAutoscalingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecurityConfig(des, initial *ClusterConfigSecurityConfig, opts ...dcl.ApplyOption) *ClusterConfigSecurityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecurityConfig{}

	cDes.KerberosConfig = canonicalizeClusterConfigSecurityConfigKerberosConfig(des.KerberosConfig, initial.KerberosConfig, opts...)

	return cDes
}

func canonicalizeClusterConfigSecurityConfigSlice(des, initial []ClusterConfigSecurityConfig, opts ...dcl.ApplyOption) []ClusterConfigSecurityConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecurityConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecurityConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecurityConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecurityConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecurityConfig(c *Client, des, nw *ClusterConfigSecurityConfig) *ClusterConfigSecurityConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecurityConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.KerberosConfig = canonicalizeNewClusterConfigSecurityConfigKerberosConfig(c, des.KerberosConfig, nw.KerberosConfig)

	return nw
}

func canonicalizeNewClusterConfigSecurityConfigSet(c *Client, des, nw []ClusterConfigSecurityConfig) []ClusterConfigSecurityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecurityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecurityConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecurityConfigSlice(c *Client, des, nw []ClusterConfigSecurityConfig) []ClusterConfigSecurityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecurityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecurityConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigSecurityConfigKerberosConfig(des, initial *ClusterConfigSecurityConfigKerberosConfig, opts ...dcl.ApplyOption) *ClusterConfigSecurityConfigKerberosConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigSecurityConfigKerberosConfig{}

	if dcl.BoolCanonicalize(des.EnableKerberos, initial.EnableKerberos) || dcl.IsZeroValue(des.EnableKerberos) {
		cDes.EnableKerberos = initial.EnableKerberos
	} else {
		cDes.EnableKerberos = des.EnableKerberos
	}
	if dcl.StringCanonicalize(des.RootPrincipalPassword, initial.RootPrincipalPassword) || dcl.IsZeroValue(des.RootPrincipalPassword) {
		cDes.RootPrincipalPassword = initial.RootPrincipalPassword
	} else {
		cDes.RootPrincipalPassword = des.RootPrincipalPassword
	}
	if dcl.NameToSelfLink(des.KmsKey, initial.KmsKey) || dcl.IsZeroValue(des.KmsKey) {
		cDes.KmsKey = initial.KmsKey
	} else {
		cDes.KmsKey = des.KmsKey
	}
	if dcl.StringCanonicalize(des.Keystore, initial.Keystore) || dcl.IsZeroValue(des.Keystore) {
		cDes.Keystore = initial.Keystore
	} else {
		cDes.Keystore = des.Keystore
	}
	if dcl.StringCanonicalize(des.Truststore, initial.Truststore) || dcl.IsZeroValue(des.Truststore) {
		cDes.Truststore = initial.Truststore
	} else {
		cDes.Truststore = des.Truststore
	}
	if dcl.StringCanonicalize(des.KeystorePassword, initial.KeystorePassword) || dcl.IsZeroValue(des.KeystorePassword) {
		cDes.KeystorePassword = initial.KeystorePassword
	} else {
		cDes.KeystorePassword = des.KeystorePassword
	}
	if dcl.StringCanonicalize(des.KeyPassword, initial.KeyPassword) || dcl.IsZeroValue(des.KeyPassword) {
		cDes.KeyPassword = initial.KeyPassword
	} else {
		cDes.KeyPassword = des.KeyPassword
	}
	if dcl.StringCanonicalize(des.TruststorePassword, initial.TruststorePassword) || dcl.IsZeroValue(des.TruststorePassword) {
		cDes.TruststorePassword = initial.TruststorePassword
	} else {
		cDes.TruststorePassword = des.TruststorePassword
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustRealm, initial.CrossRealmTrustRealm) || dcl.IsZeroValue(des.CrossRealmTrustRealm) {
		cDes.CrossRealmTrustRealm = initial.CrossRealmTrustRealm
	} else {
		cDes.CrossRealmTrustRealm = des.CrossRealmTrustRealm
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustKdc, initial.CrossRealmTrustKdc) || dcl.IsZeroValue(des.CrossRealmTrustKdc) {
		cDes.CrossRealmTrustKdc = initial.CrossRealmTrustKdc
	} else {
		cDes.CrossRealmTrustKdc = des.CrossRealmTrustKdc
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustAdminServer, initial.CrossRealmTrustAdminServer) || dcl.IsZeroValue(des.CrossRealmTrustAdminServer) {
		cDes.CrossRealmTrustAdminServer = initial.CrossRealmTrustAdminServer
	} else {
		cDes.CrossRealmTrustAdminServer = des.CrossRealmTrustAdminServer
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustSharedPassword, initial.CrossRealmTrustSharedPassword) || dcl.IsZeroValue(des.CrossRealmTrustSharedPassword) {
		cDes.CrossRealmTrustSharedPassword = initial.CrossRealmTrustSharedPassword
	} else {
		cDes.CrossRealmTrustSharedPassword = des.CrossRealmTrustSharedPassword
	}
	if dcl.StringCanonicalize(des.KdcDbKey, initial.KdcDbKey) || dcl.IsZeroValue(des.KdcDbKey) {
		cDes.KdcDbKey = initial.KdcDbKey
	} else {
		cDes.KdcDbKey = des.KdcDbKey
	}
	if dcl.IsZeroValue(des.TgtLifetimeHours) || (dcl.IsEmptyValueIndirect(des.TgtLifetimeHours) && dcl.IsEmptyValueIndirect(initial.TgtLifetimeHours)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TgtLifetimeHours = initial.TgtLifetimeHours
	} else {
		cDes.TgtLifetimeHours = des.TgtLifetimeHours
	}
	if dcl.StringCanonicalize(des.Realm, initial.Realm) || dcl.IsZeroValue(des.Realm) {
		cDes.Realm = initial.Realm
	} else {
		cDes.Realm = des.Realm
	}

	return cDes
}

func canonicalizeClusterConfigSecurityConfigKerberosConfigSlice(des, initial []ClusterConfigSecurityConfigKerberosConfig, opts ...dcl.ApplyOption) []ClusterConfigSecurityConfigKerberosConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigSecurityConfigKerberosConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigSecurityConfigKerberosConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigSecurityConfigKerberosConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigSecurityConfigKerberosConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigSecurityConfigKerberosConfig(c *Client, des, nw *ClusterConfigSecurityConfigKerberosConfig) *ClusterConfigSecurityConfigKerberosConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigSecurityConfigKerberosConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.EnableKerberos, nw.EnableKerberos) {
		nw.EnableKerberos = des.EnableKerberos
	}
	if dcl.StringCanonicalize(des.RootPrincipalPassword, nw.RootPrincipalPassword) {
		nw.RootPrincipalPassword = des.RootPrincipalPassword
	}
	if dcl.NameToSelfLink(des.KmsKey, nw.KmsKey) {
		nw.KmsKey = des.KmsKey
	}
	if dcl.StringCanonicalize(des.Keystore, nw.Keystore) {
		nw.Keystore = des.Keystore
	}
	if dcl.StringCanonicalize(des.Truststore, nw.Truststore) {
		nw.Truststore = des.Truststore
	}
	if dcl.StringCanonicalize(des.KeystorePassword, nw.KeystorePassword) {
		nw.KeystorePassword = des.KeystorePassword
	}
	if dcl.StringCanonicalize(des.KeyPassword, nw.KeyPassword) {
		nw.KeyPassword = des.KeyPassword
	}
	if dcl.StringCanonicalize(des.TruststorePassword, nw.TruststorePassword) {
		nw.TruststorePassword = des.TruststorePassword
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustRealm, nw.CrossRealmTrustRealm) {
		nw.CrossRealmTrustRealm = des.CrossRealmTrustRealm
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustKdc, nw.CrossRealmTrustKdc) {
		nw.CrossRealmTrustKdc = des.CrossRealmTrustKdc
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustAdminServer, nw.CrossRealmTrustAdminServer) {
		nw.CrossRealmTrustAdminServer = des.CrossRealmTrustAdminServer
	}
	if dcl.StringCanonicalize(des.CrossRealmTrustSharedPassword, nw.CrossRealmTrustSharedPassword) {
		nw.CrossRealmTrustSharedPassword = des.CrossRealmTrustSharedPassword
	}
	if dcl.StringCanonicalize(des.KdcDbKey, nw.KdcDbKey) {
		nw.KdcDbKey = des.KdcDbKey
	}
	if dcl.StringCanonicalize(des.Realm, nw.Realm) {
		nw.Realm = des.Realm
	}

	return nw
}

func canonicalizeNewClusterConfigSecurityConfigKerberosConfigSet(c *Client, des, nw []ClusterConfigSecurityConfigKerberosConfig) []ClusterConfigSecurityConfigKerberosConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigSecurityConfigKerberosConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigSecurityConfigKerberosConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigSecurityConfigKerberosConfigSlice(c *Client, des, nw []ClusterConfigSecurityConfigKerberosConfig) []ClusterConfigSecurityConfigKerberosConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigSecurityConfigKerberosConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigSecurityConfigKerberosConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigLifecycleConfig(des, initial *ClusterConfigLifecycleConfig, opts ...dcl.ApplyOption) *ClusterConfigLifecycleConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigLifecycleConfig{}

	if dcl.StringCanonicalize(des.IdleDeleteTtl, initial.IdleDeleteTtl) || dcl.IsZeroValue(des.IdleDeleteTtl) {
		cDes.IdleDeleteTtl = initial.IdleDeleteTtl
	} else {
		cDes.IdleDeleteTtl = des.IdleDeleteTtl
	}
	if dcl.IsZeroValue(des.AutoDeleteTime) || (dcl.IsEmptyValueIndirect(des.AutoDeleteTime) && dcl.IsEmptyValueIndirect(initial.AutoDeleteTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AutoDeleteTime = initial.AutoDeleteTime
	} else {
		cDes.AutoDeleteTime = des.AutoDeleteTime
	}
	if dcl.StringCanonicalize(des.AutoDeleteTtl, initial.AutoDeleteTtl) || dcl.IsZeroValue(des.AutoDeleteTtl) {
		cDes.AutoDeleteTtl = initial.AutoDeleteTtl
	} else {
		cDes.AutoDeleteTtl = des.AutoDeleteTtl
	}

	return cDes
}

func canonicalizeClusterConfigLifecycleConfigSlice(des, initial []ClusterConfigLifecycleConfig, opts ...dcl.ApplyOption) []ClusterConfigLifecycleConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigLifecycleConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigLifecycleConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigLifecycleConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigLifecycleConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigLifecycleConfig(c *Client, des, nw *ClusterConfigLifecycleConfig) *ClusterConfigLifecycleConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigLifecycleConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IdleDeleteTtl, nw.IdleDeleteTtl) {
		nw.IdleDeleteTtl = des.IdleDeleteTtl
	}
	if dcl.StringCanonicalize(des.AutoDeleteTtl, nw.AutoDeleteTtl) {
		nw.AutoDeleteTtl = des.AutoDeleteTtl
	}

	return nw
}

func canonicalizeNewClusterConfigLifecycleConfigSet(c *Client, des, nw []ClusterConfigLifecycleConfig) []ClusterConfigLifecycleConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigLifecycleConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigLifecycleConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigLifecycleConfigSlice(c *Client, des, nw []ClusterConfigLifecycleConfig) []ClusterConfigLifecycleConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigLifecycleConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigLifecycleConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfigEndpointConfig(des, initial *ClusterConfigEndpointConfig, opts ...dcl.ApplyOption) *ClusterConfigEndpointConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterConfigEndpointConfig{}

	if dcl.BoolCanonicalize(des.EnableHttpPortAccess, initial.EnableHttpPortAccess) || dcl.IsZeroValue(des.EnableHttpPortAccess) {
		cDes.EnableHttpPortAccess = initial.EnableHttpPortAccess
	} else {
		cDes.EnableHttpPortAccess = des.EnableHttpPortAccess
	}

	return cDes
}

func canonicalizeClusterConfigEndpointConfigSlice(des, initial []ClusterConfigEndpointConfig, opts ...dcl.ApplyOption) []ClusterConfigEndpointConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterConfigEndpointConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterConfigEndpointConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterConfigEndpointConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterConfigEndpointConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterConfigEndpointConfig(c *Client, des, nw *ClusterConfigEndpointConfig) *ClusterConfigEndpointConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterConfigEndpointConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.EnableHttpPortAccess, nw.EnableHttpPortAccess) {
		nw.EnableHttpPortAccess = des.EnableHttpPortAccess
	}

	return nw
}

func canonicalizeNewClusterConfigEndpointConfigSet(c *Client, des, nw []ClusterConfigEndpointConfig) []ClusterConfigEndpointConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfigEndpointConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterConfigEndpointConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterConfigEndpointConfigSlice(c *Client, des, nw []ClusterConfigEndpointConfig) []ClusterConfigEndpointConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterConfigEndpointConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfigEndpointConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterStatus(des, initial *ClusterStatus, opts ...dcl.ApplyOption) *ClusterStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterStatus{}

	return cDes
}

func canonicalizeClusterStatusSlice(des, initial []ClusterStatus, opts ...dcl.ApplyOption) []ClusterStatus {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterStatus(c *Client, des, nw *ClusterStatus) *ClusterStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Detail, nw.Detail) {
		nw.Detail = des.Detail
	}

	return nw
}

func canonicalizeNewClusterStatusSet(c *Client, des, nw []ClusterStatus) []ClusterStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterStatusSlice(c *Client, des, nw []ClusterStatus) []ClusterStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterStatus(c, &d, &n))
	}

	return items
}

func canonicalizeClusterStatusHistory(des, initial *ClusterStatusHistory, opts ...dcl.ApplyOption) *ClusterStatusHistory {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterStatusHistory{}

	return cDes
}

func canonicalizeClusterStatusHistorySlice(des, initial []ClusterStatusHistory, opts ...dcl.ApplyOption) []ClusterStatusHistory {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterStatusHistory, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterStatusHistory(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterStatusHistory, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterStatusHistory(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterStatusHistory(c *Client, des, nw *ClusterStatusHistory) *ClusterStatusHistory {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterStatusHistory while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Detail, nw.Detail) {
		nw.Detail = des.Detail
	}

	return nw
}

func canonicalizeNewClusterStatusHistorySet(c *Client, des, nw []ClusterStatusHistory) []ClusterStatusHistory {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterStatusHistory
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterStatusHistoryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterStatusHistorySlice(c *Client, des, nw []ClusterStatusHistory) []ClusterStatusHistory {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterStatusHistory
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterStatusHistory(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMetrics(des, initial *ClusterMetrics, opts ...dcl.ApplyOption) *ClusterMetrics {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterMetrics{}

	if dcl.IsZeroValue(des.HdfsMetrics) || (dcl.IsEmptyValueIndirect(des.HdfsMetrics) && dcl.IsEmptyValueIndirect(initial.HdfsMetrics)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.HdfsMetrics = initial.HdfsMetrics
	} else {
		cDes.HdfsMetrics = des.HdfsMetrics
	}
	if dcl.IsZeroValue(des.YarnMetrics) || (dcl.IsEmptyValueIndirect(des.YarnMetrics) && dcl.IsEmptyValueIndirect(initial.YarnMetrics)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.YarnMetrics = initial.YarnMetrics
	} else {
		cDes.YarnMetrics = des.YarnMetrics
	}

	return cDes
}

func canonicalizeClusterMetricsSlice(des, initial []ClusterMetrics, opts ...dcl.ApplyOption) []ClusterMetrics {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterMetrics, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterMetrics(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterMetrics, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterMetrics(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterMetrics(c *Client, des, nw *ClusterMetrics) *ClusterMetrics {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterMetrics while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewClusterMetricsSet(c *Client, des, nw []ClusterMetrics) []ClusterMetrics {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMetrics
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterMetricsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterMetricsSlice(c *Client, des, nw []ClusterMetrics) []ClusterMetrics {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterMetrics
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMetrics(c, &d, &n))
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
func diffCluster(c *Client, desired, actual *Cluster, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareClusterConfigNewStyle, EmptyObject: EmptyClusterConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{IgnoredPrefixes: []string{"goog-dataproc-"}, OperationSelector: dcl.TriggersOperation("updateClusterUpdateClusterOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterStatusNewStyle, EmptyObject: EmptyClusterStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusHistory, actual.StatusHistory, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterStatusHistoryNewStyle, EmptyObject: EmptyClusterStatusHistory, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StatusHistory")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterUuid, actual.ClusterUuid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterUuid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metrics, actual.Metrics, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterMetricsNewStyle, EmptyObject: EmptyClusterMetrics, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metrics")); len(ds) != 0 || err != nil {
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
func compareClusterConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfig or *ClusterConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StagingBucket, actual.StagingBucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TempBucket, actual.TempBucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TempBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GceClusterConfig, actual.GceClusterConfig, dcl.Info{ObjectFunction: compareClusterConfigGceClusterConfigNewStyle, EmptyObject: EmptyClusterConfigGceClusterConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GceClusterConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterConfig, actual.MasterConfig, dcl.Info{ObjectFunction: compareClusterConfigMasterConfigNewStyle, EmptyObject: EmptyClusterConfigMasterConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MasterConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkerConfig, actual.WorkerConfig, dcl.Info{ObjectFunction: compareClusterConfigWorkerConfigNewStyle, EmptyObject: EmptyClusterConfigWorkerConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryWorkerConfig, actual.SecondaryWorkerConfig, dcl.Info{ObjectFunction: compareClusterConfigSecondaryWorkerConfigNewStyle, EmptyObject: EmptyClusterConfigSecondaryWorkerConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecondaryWorkerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SoftwareConfig, actual.SoftwareConfig, dcl.Info{ObjectFunction: compareClusterConfigSoftwareConfigNewStyle, EmptyObject: EmptyClusterConfigSoftwareConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SoftwareConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitializationActions, actual.InitializationActions, dcl.Info{ObjectFunction: compareClusterConfigInitializationActionsNewStyle, EmptyObject: EmptyClusterConfigInitializationActions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InitializationActions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionConfig, actual.EncryptionConfig, dcl.Info{ObjectFunction: compareClusterConfigEncryptionConfigNewStyle, EmptyObject: EmptyClusterConfigEncryptionConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoscalingConfig, actual.AutoscalingConfig, dcl.Info{ObjectFunction: compareClusterConfigAutoscalingConfigNewStyle, EmptyObject: EmptyClusterConfigAutoscalingConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutoscalingConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityConfig, actual.SecurityConfig, dcl.Info{ObjectFunction: compareClusterConfigSecurityConfigNewStyle, EmptyObject: EmptyClusterConfigSecurityConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LifecycleConfig, actual.LifecycleConfig, dcl.Info{ObjectFunction: compareClusterConfigLifecycleConfigNewStyle, EmptyObject: EmptyClusterConfigLifecycleConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LifecycleConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndpointConfig, actual.EndpointConfig, dcl.Info{ObjectFunction: compareClusterConfigEndpointConfigNewStyle, EmptyObject: EmptyClusterConfigEndpointConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndpointConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigGceClusterConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigGceClusterConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigGceClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfig or *ClusterConfigGceClusterConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigGceClusterConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigGceClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ZoneUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetworkUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIPOnly, actual.InternalIPOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InternalIpOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateIPv6GoogleAccess, actual.PrivateIPv6GoogleAccess, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateIpv6GoogleAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccountScopes, actual.ServiceAccountScopes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccountScopes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{Type: "Set", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservationAffinity, actual.ReservationAffinity, dcl.Info{ObjectFunction: compareClusterConfigGceClusterConfigReservationAffinityNewStyle, EmptyObject: EmptyClusterConfigGceClusterConfigReservationAffinity, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReservationAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeGroupAffinity, actual.NodeGroupAffinity, dcl.Info{ObjectFunction: compareClusterConfigGceClusterConfigNodeGroupAffinityNewStyle, EmptyObject: EmptyClusterConfigGceClusterConfigNodeGroupAffinity, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeGroupAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigGceClusterConfigReservationAffinityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigGceClusterConfigReservationAffinity)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigGceClusterConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfigReservationAffinity or *ClusterConfigGceClusterConfigReservationAffinity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigGceClusterConfigReservationAffinity)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigGceClusterConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfigReservationAffinity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConsumeReservationType, actual.ConsumeReservationType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConsumeReservationType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigGceClusterConfigNodeGroupAffinityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigGceClusterConfigNodeGroupAffinity)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigGceClusterConfigNodeGroupAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfigNodeGroupAffinity or *ClusterConfigGceClusterConfigNodeGroupAffinity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigGceClusterConfigNodeGroupAffinity)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigGceClusterConfigNodeGroupAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigGceClusterConfigNodeGroupAffinity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NodeGroup, actual.NodeGroup, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeGroupUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigMasterConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigMasterConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigMasterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfig or *ClusterConfigMasterConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigMasterConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigMasterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumInstances, actual.NumInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceNames, actual.InstanceNames, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskConfig, actual.DiskConfig, dcl.Info{ObjectFunction: compareClusterConfigMasterConfigDiskConfigNewStyle, EmptyObject: EmptyClusterConfigMasterConfigDiskConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IsPreemptible, actual.IsPreemptible, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsPreemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptibility, actual.Preemptibility, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Preemptibility")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagedGroupConfig, actual.ManagedGroupConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterConfigMasterConfigManagedGroupConfigNewStyle, EmptyObject: EmptyClusterConfigMasterConfigManagedGroupConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagedGroupConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Accelerators, actual.Accelerators, dcl.Info{ObjectFunction: compareClusterConfigMasterConfigAcceleratorsNewStyle, EmptyObject: EmptyClusterConfigMasterConfigAccelerators, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Accelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigMasterConfigDiskConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigMasterConfigDiskConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigMasterConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigDiskConfig or *ClusterConfigMasterConfigDiskConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigMasterConfigDiskConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigMasterConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigDiskConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BootDiskType, actual.BootDiskType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskSizeGb, actual.BootDiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumLocalSsds, actual.NumLocalSsds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumLocalSsds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigMasterConfigManagedGroupConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigMasterConfigManagedGroupConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigMasterConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigManagedGroupConfig or *ClusterConfigMasterConfigManagedGroupConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigMasterConfigManagedGroupConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigMasterConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigManagedGroupConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceTemplateName, actual.InstanceTemplateName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceTemplateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceGroupManagerName, actual.InstanceGroupManagerName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceGroupManagerName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigMasterConfigAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigMasterConfigAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigMasterConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigAccelerators or *ClusterConfigMasterConfigAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigMasterConfigAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigMasterConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigMasterConfigAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigWorkerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigWorkerConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfig or *ClusterConfigWorkerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigWorkerConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumInstances, actual.NumInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceNames, actual.InstanceNames, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskConfig, actual.DiskConfig, dcl.Info{ObjectFunction: compareClusterConfigWorkerConfigDiskConfigNewStyle, EmptyObject: EmptyClusterConfigWorkerConfigDiskConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IsPreemptible, actual.IsPreemptible, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsPreemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptibility, actual.Preemptibility, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Preemptibility")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagedGroupConfig, actual.ManagedGroupConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterConfigWorkerConfigManagedGroupConfigNewStyle, EmptyObject: EmptyClusterConfigWorkerConfigManagedGroupConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagedGroupConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Accelerators, actual.Accelerators, dcl.Info{ObjectFunction: compareClusterConfigWorkerConfigAcceleratorsNewStyle, EmptyObject: EmptyClusterConfigWorkerConfigAccelerators, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Accelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigWorkerConfigDiskConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigWorkerConfigDiskConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigWorkerConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigDiskConfig or *ClusterConfigWorkerConfigDiskConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigWorkerConfigDiskConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigWorkerConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigDiskConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BootDiskType, actual.BootDiskType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskSizeGb, actual.BootDiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumLocalSsds, actual.NumLocalSsds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumLocalSsds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigWorkerConfigManagedGroupConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigWorkerConfigManagedGroupConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigWorkerConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigManagedGroupConfig or *ClusterConfigWorkerConfigManagedGroupConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigWorkerConfigManagedGroupConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigWorkerConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigManagedGroupConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceTemplateName, actual.InstanceTemplateName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceTemplateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceGroupManagerName, actual.InstanceGroupManagerName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceGroupManagerName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigWorkerConfigAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigWorkerConfigAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigWorkerConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigAccelerators or *ClusterConfigWorkerConfigAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigWorkerConfigAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigWorkerConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigWorkerConfigAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecondaryWorkerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecondaryWorkerConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecondaryWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfig or *ClusterConfigSecondaryWorkerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecondaryWorkerConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecondaryWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumInstances, actual.NumInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceNames, actual.InstanceNames, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskConfig, actual.DiskConfig, dcl.Info{ObjectFunction: compareClusterConfigSecondaryWorkerConfigDiskConfigNewStyle, EmptyObject: EmptyClusterConfigSecondaryWorkerConfigDiskConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IsPreemptible, actual.IsPreemptible, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsPreemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptibility, actual.Preemptibility, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Preemptibility")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagedGroupConfig, actual.ManagedGroupConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterConfigSecondaryWorkerConfigManagedGroupConfigNewStyle, EmptyObject: EmptyClusterConfigSecondaryWorkerConfigManagedGroupConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagedGroupConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Accelerators, actual.Accelerators, dcl.Info{ObjectFunction: compareClusterConfigSecondaryWorkerConfigAcceleratorsNewStyle, EmptyObject: EmptyClusterConfigSecondaryWorkerConfigAccelerators, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Accelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecondaryWorkerConfigDiskConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecondaryWorkerConfigDiskConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecondaryWorkerConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigDiskConfig or *ClusterConfigSecondaryWorkerConfigDiskConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecondaryWorkerConfigDiskConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecondaryWorkerConfigDiskConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigDiskConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BootDiskType, actual.BootDiskType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskSizeGb, actual.BootDiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BootDiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumLocalSsds, actual.NumLocalSsds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumLocalSsds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecondaryWorkerConfigManagedGroupConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecondaryWorkerConfigManagedGroupConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecondaryWorkerConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigManagedGroupConfig or *ClusterConfigSecondaryWorkerConfigManagedGroupConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecondaryWorkerConfigManagedGroupConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecondaryWorkerConfigManagedGroupConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigManagedGroupConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceTemplateName, actual.InstanceTemplateName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceTemplateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceGroupManagerName, actual.InstanceGroupManagerName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceGroupManagerName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecondaryWorkerConfigAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecondaryWorkerConfigAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecondaryWorkerConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigAccelerators or *ClusterConfigSecondaryWorkerConfigAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecondaryWorkerConfigAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecondaryWorkerConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecondaryWorkerConfigAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorTypeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSoftwareConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSoftwareConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSoftwareConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSoftwareConfig or *ClusterConfigSoftwareConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSoftwareConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSoftwareConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSoftwareConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ImageVersion, actual.ImageVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OptionalComponents, actual.OptionalComponents, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OptionalComponents")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigInitializationActionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigInitializationActions)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigInitializationActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigInitializationActions or *ClusterConfigInitializationActions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigInitializationActions)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigInitializationActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigInitializationActions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ExecutableFile, actual.ExecutableFile, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExecutableFile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionTimeout, actual.ExecutionTimeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExecutionTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigEncryptionConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigEncryptionConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigEncryptionConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigEncryptionConfig or *ClusterConfigEncryptionConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigEncryptionConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigEncryptionConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigEncryptionConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GcePdKmsKeyName, actual.GcePdKmsKeyName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GcePdKmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigAutoscalingConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigAutoscalingConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigAutoscalingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigAutoscalingConfig or *ClusterConfigAutoscalingConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigAutoscalingConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigAutoscalingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigAutoscalingConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Policy, actual.Policy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecurityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecurityConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecurityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecurityConfig or *ClusterConfigSecurityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecurityConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecurityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecurityConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KerberosConfig, actual.KerberosConfig, dcl.Info{ObjectFunction: compareClusterConfigSecurityConfigKerberosConfigNewStyle, EmptyObject: EmptyClusterConfigSecurityConfigKerberosConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KerberosConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigSecurityConfigKerberosConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigSecurityConfigKerberosConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigSecurityConfigKerberosConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecurityConfigKerberosConfig or *ClusterConfigSecurityConfigKerberosConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigSecurityConfigKerberosConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigSecurityConfigKerberosConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigSecurityConfigKerberosConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableKerberos, actual.EnableKerberos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableKerberos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootPrincipalPassword, actual.RootPrincipalPassword, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootPrincipalPasswordUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKey, actual.KmsKey, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Keystore, actual.Keystore, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeystoreUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Truststore, actual.Truststore, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TruststoreUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeystorePassword, actual.KeystorePassword, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeystorePasswordUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyPassword, actual.KeyPassword, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyPasswordUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TruststorePassword, actual.TruststorePassword, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TruststorePasswordUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossRealmTrustRealm, actual.CrossRealmTrustRealm, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossRealmTrustRealm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossRealmTrustKdc, actual.CrossRealmTrustKdc, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossRealmTrustKdc")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossRealmTrustAdminServer, actual.CrossRealmTrustAdminServer, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossRealmTrustAdminServer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossRealmTrustSharedPassword, actual.CrossRealmTrustSharedPassword, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossRealmTrustSharedPasswordUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KdcDbKey, actual.KdcDbKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KdcDbKeyUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TgtLifetimeHours, actual.TgtLifetimeHours, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TgtLifetimeHours")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Realm, actual.Realm, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Realm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigLifecycleConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigLifecycleConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigLifecycleConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigLifecycleConfig or *ClusterConfigLifecycleConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigLifecycleConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigLifecycleConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigLifecycleConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IdleDeleteTtl, actual.IdleDeleteTtl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdleDeleteTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoDeleteTime, actual.AutoDeleteTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutoDeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoDeleteTtl, actual.AutoDeleteTtl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutoDeleteTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdleStartTime, actual.IdleStartTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdleStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfigEndpointConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfigEndpointConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfigEndpointConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigEndpointConfig or *ClusterConfigEndpointConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfigEndpointConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfigEndpointConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfigEndpointConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpPorts, actual.HttpPorts, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpPorts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableHttpPortAccess, actual.EnableHttpPortAccess, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableHttpPortAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterStatus)
	if !ok {
		desiredNotPointer, ok := d.(ClusterStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterStatus or *ClusterStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterStatus)
	if !ok {
		actualNotPointer, ok := a.(ClusterStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Detail, actual.Detail, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Detail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StateStartTime, actual.StateStartTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StateStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Substate, actual.Substate, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Substate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterStatusHistoryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterStatusHistory)
	if !ok {
		desiredNotPointer, ok := d.(ClusterStatusHistory)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterStatusHistory or *ClusterStatusHistory", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterStatusHistory)
	if !ok {
		actualNotPointer, ok := a.(ClusterStatusHistory)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterStatusHistory", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Detail, actual.Detail, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Detail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StateStartTime, actual.StateStartTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StateStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Substate, actual.Substate, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Substate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMetricsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMetrics)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMetrics)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMetrics or *ClusterMetrics", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMetrics)
	if !ok {
		actualNotPointer, ok := a.(ClusterMetrics)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMetrics", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HdfsMetrics, actual.HdfsMetrics, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HdfsMetrics")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.YarnMetrics, actual.YarnMetrics, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("YarnMetrics")); len(ds) != 0 || err != nil {
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
func (r *Cluster) urlNormalized() *Cluster {
	normalized := dcl.Copy(*r).(Cluster)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ClusterUuid = dcl.SelfLinkToName(r.ClusterUuid)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Cluster) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateCluster" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Cluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Cluster) marshal(c *Client) ([]byte, error) {
	m, err := expandCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Cluster: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCluster decodes JSON responses into the Cluster resource schema.
func unmarshalCluster(b []byte, c *Client) (*Cluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCluster(m, c)
}

func unmarshalMapCluster(m map[string]interface{}, c *Client) (*Cluster, error) {

	flattened := flattenCluster(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandCluster expands Cluster into a JSON request object.
func expandCluster(c *Client, f *Cluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := expandClusterProject(c, f.Project, res); err != nil {
		return nil, fmt.Errorf("error expanding Project into projectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["clusterName"] = v
	}
	if v, err := expandClusterConfig(c, f.Config, res); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenCluster flattens Cluster from a JSON request object into the
// Cluster type.
func flattenCluster(c *Client, i interface{}) *Cluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Cluster{}
	res.Project = dcl.FlattenString(m["projectId"])
	res.Name = dcl.FlattenString(m["clusterName"])
	res.Config = flattenClusterConfig(c, m["config"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Status = flattenClusterStatus(c, m["status"])
	res.StatusHistory = flattenClusterStatusHistorySlice(c, m["statusHistory"])
	res.ClusterUuid = dcl.FlattenString(m["clusterUuid"])
	res.Metrics = flattenClusterMetrics(c, m["metrics"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandClusterConfigMap expands the contents of ClusterConfig into a JSON
// request object.
func expandClusterConfigMap(c *Client, f map[string]ClusterConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSlice expands the contents of ClusterConfig into a JSON
// request object.
func expandClusterConfigSlice(c *Client, f []ClusterConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigMap flattens the contents of ClusterConfig from a JSON
// response object.
func flattenClusterConfigMap(c *Client, i interface{}) map[string]ClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfig{}
	}

	items := make(map[string]ClusterConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSlice flattens the contents of ClusterConfig from a JSON
// response object.
func flattenClusterConfigSlice(c *Client, i interface{}) []ClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfig{}
	}

	items := make([]ClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfig expands an instance of ClusterConfig into a JSON
// request object.
func expandClusterConfig(c *Client, f *ClusterConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StagingBucket; !dcl.IsEmptyValueIndirect(v) {
		m["configBucket"] = v
	}
	if v := f.TempBucket; !dcl.IsEmptyValueIndirect(v) {
		m["tempBucket"] = v
	}
	if v, err := expandClusterConfigGceClusterConfig(c, f.GceClusterConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding GceClusterConfig into gceClusterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gceClusterConfig"] = v
	}
	if v, err := expandClusterConfigMasterConfig(c, f.MasterConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding MasterConfig into masterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterConfig"] = v
	}
	if v, err := expandClusterConfigWorkerConfig(c, f.WorkerConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workerConfig"] = v
	}
	if v, err := expandClusterConfigSecondaryWorkerConfig(c, f.SecondaryWorkerConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryWorkerConfig into secondaryWorkerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryWorkerConfig"] = v
	}
	if v, err := expandClusterConfigSoftwareConfig(c, f.SoftwareConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SoftwareConfig into softwareConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["softwareConfig"] = v
	}
	if v, err := expandClusterConfigInitializationActionsSlice(c, f.InitializationActions, res); err != nil {
		return nil, fmt.Errorf("error expanding InitializationActions into initializationActions: %w", err)
	} else if v != nil {
		m["initializationActions"] = v
	}
	if v, err := expandClusterConfigEncryptionConfig(c, f.EncryptionConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionConfig into encryptionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionConfig"] = v
	}
	if v, err := expandClusterConfigAutoscalingConfig(c, f.AutoscalingConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingConfig into autoscalingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscalingConfig"] = v
	}
	if v, err := expandClusterConfigSecurityConfig(c, f.SecurityConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SecurityConfig into securityConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["securityConfig"] = v
	}
	if v, err := expandClusterConfigLifecycleConfig(c, f.LifecycleConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding LifecycleConfig into lifecycleConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["lifecycleConfig"] = v
	}
	if v, err := expandClusterConfigEndpointConfig(c, f.EndpointConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding EndpointConfig into endpointConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpointConfig"] = v
	}

	return m, nil
}

// flattenClusterConfig flattens an instance of ClusterConfig from a JSON
// response object.
func flattenClusterConfig(c *Client, i interface{}) *ClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfig
	}
	r.StagingBucket = dcl.FlattenString(m["configBucket"])
	r.TempBucket = dcl.FlattenString(m["tempBucket"])
	r.GceClusterConfig = flattenClusterConfigGceClusterConfig(c, m["gceClusterConfig"])
	r.MasterConfig = flattenClusterConfigMasterConfig(c, m["masterConfig"])
	r.WorkerConfig = flattenClusterConfigWorkerConfig(c, m["workerConfig"])
	r.SecondaryWorkerConfig = flattenClusterConfigSecondaryWorkerConfig(c, m["secondaryWorkerConfig"])
	r.SoftwareConfig = flattenClusterConfigSoftwareConfig(c, m["softwareConfig"])
	r.InitializationActions = flattenClusterConfigInitializationActionsSlice(c, m["initializationActions"])
	r.EncryptionConfig = flattenClusterConfigEncryptionConfig(c, m["encryptionConfig"])
	r.AutoscalingConfig = flattenClusterConfigAutoscalingConfig(c, m["autoscalingConfig"])
	r.SecurityConfig = flattenClusterConfigSecurityConfig(c, m["securityConfig"])
	r.LifecycleConfig = flattenClusterConfigLifecycleConfig(c, m["lifecycleConfig"])
	r.EndpointConfig = flattenClusterConfigEndpointConfig(c, m["endpointConfig"])

	return r
}

// expandClusterConfigGceClusterConfigMap expands the contents of ClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterConfigGceClusterConfigMap(c *Client, f map[string]ClusterConfigGceClusterConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigGceClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigGceClusterConfigSlice expands the contents of ClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterConfigGceClusterConfigSlice(c *Client, f []ClusterConfigGceClusterConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigGceClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigGceClusterConfigMap flattens the contents of ClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterConfigGceClusterConfigMap(c *Client, i interface{}) map[string]ClusterConfigGceClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigGceClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigGceClusterConfig{}
	}

	items := make(map[string]ClusterConfigGceClusterConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigGceClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigGceClusterConfigSlice flattens the contents of ClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterConfigGceClusterConfigSlice(c *Client, i interface{}) []ClusterConfigGceClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigGceClusterConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigGceClusterConfig{}
	}

	items := make([]ClusterConfigGceClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigGceClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigGceClusterConfig expands an instance of ClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterConfigGceClusterConfig(c *Client, f *ClusterConfigGceClusterConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zoneUri"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["networkUri"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetworkUri"] = v
	}
	if v := f.InternalIPOnly; !dcl.IsEmptyValueIndirect(v) {
		m["internalIpOnly"] = v
	}
	if v := f.PrivateIPv6GoogleAccess; !dcl.IsEmptyValueIndirect(v) {
		m["privateIpv6GoogleAccess"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.ServiceAccountScopes; v != nil {
		m["serviceAccountScopes"] = v
	}
	if v := f.Tags; v != nil {
		m["tags"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := expandClusterConfigGceClusterConfigReservationAffinity(c, f.ReservationAffinity, res); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandClusterConfigGceClusterConfigNodeGroupAffinity(c, f.NodeGroupAffinity, res); err != nil {
		return nil, fmt.Errorf("error expanding NodeGroupAffinity into nodeGroupAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroupAffinity"] = v
	}

	return m, nil
}

// flattenClusterConfigGceClusterConfig flattens an instance of ClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterConfigGceClusterConfig(c *Client, i interface{}) *ClusterConfigGceClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigGceClusterConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigGceClusterConfig
	}
	r.Zone = dcl.FlattenString(m["zoneUri"])
	r.Network = dcl.FlattenString(m["networkUri"])
	r.Subnetwork = dcl.FlattenString(m["subnetworkUri"])
	r.InternalIPOnly = dcl.FlattenBool(m["internalIpOnly"])
	r.PrivateIPv6GoogleAccess = flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(m["privateIpv6GoogleAccess"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.ServiceAccountScopes = dcl.FlattenStringSlice(m["serviceAccountScopes"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ReservationAffinity = flattenClusterConfigGceClusterConfigReservationAffinity(c, m["reservationAffinity"])
	r.NodeGroupAffinity = flattenClusterConfigGceClusterConfigNodeGroupAffinity(c, m["nodeGroupAffinity"])

	return r
}

// expandClusterConfigGceClusterConfigReservationAffinityMap expands the contents of ClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigReservationAffinityMap(c *Client, f map[string]ClusterConfigGceClusterConfigReservationAffinity, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigGceClusterConfigReservationAffinity(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigGceClusterConfigReservationAffinitySlice expands the contents of ClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, f []ClusterConfigGceClusterConfigReservationAffinity, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigGceClusterConfigReservationAffinity(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigGceClusterConfigReservationAffinityMap flattens the contents of ClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigReservationAffinityMap(c *Client, i interface{}) map[string]ClusterConfigGceClusterConfigReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigGceClusterConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigGceClusterConfigReservationAffinity{}
	}

	items := make(map[string]ClusterConfigGceClusterConfigReservationAffinity)
	for k, item := range a {
		items[k] = *flattenClusterConfigGceClusterConfigReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigGceClusterConfigReservationAffinitySlice flattens the contents of ClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, i interface{}) []ClusterConfigGceClusterConfigReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigGceClusterConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return []ClusterConfigGceClusterConfigReservationAffinity{}
	}

	items := make([]ClusterConfigGceClusterConfigReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigGceClusterConfigReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigGceClusterConfigReservationAffinity expands an instance of ClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigReservationAffinity(c *Client, f *ClusterConfigGceClusterConfigReservationAffinity, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConsumeReservationType; !dcl.IsEmptyValueIndirect(v) {
		m["consumeReservationType"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Values; v != nil {
		m["values"] = v
	}

	return m, nil
}

// flattenClusterConfigGceClusterConfigReservationAffinity flattens an instance of ClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigReservationAffinity(c *Client, i interface{}) *ClusterConfigGceClusterConfigReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigGceClusterConfigReservationAffinity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigGceClusterConfigReservationAffinity
	}
	r.ConsumeReservationType = flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(m["consumeReservationType"])
	r.Key = dcl.FlattenString(m["key"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandClusterConfigGceClusterConfigNodeGroupAffinityMap expands the contents of ClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigNodeGroupAffinityMap(c *Client, f map[string]ClusterConfigGceClusterConfigNodeGroupAffinity, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigGceClusterConfigNodeGroupAffinity(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigGceClusterConfigNodeGroupAffinitySlice expands the contents of ClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, f []ClusterConfigGceClusterConfigNodeGroupAffinity, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigGceClusterConfigNodeGroupAffinity(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigGceClusterConfigNodeGroupAffinityMap flattens the contents of ClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigNodeGroupAffinityMap(c *Client, i interface{}) map[string]ClusterConfigGceClusterConfigNodeGroupAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	items := make(map[string]ClusterConfigGceClusterConfigNodeGroupAffinity)
	for k, item := range a {
		items[k] = *flattenClusterConfigGceClusterConfigNodeGroupAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigGceClusterConfigNodeGroupAffinitySlice flattens the contents of ClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, i interface{}) []ClusterConfigGceClusterConfigNodeGroupAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	if len(a) == 0 {
		return []ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	items := make([]ClusterConfigGceClusterConfigNodeGroupAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigGceClusterConfigNodeGroupAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigGceClusterConfigNodeGroupAffinity expands an instance of ClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, f *ClusterConfigGceClusterConfigNodeGroupAffinity, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NodeGroup; !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroupUri"] = v
	}

	return m, nil
}

// flattenClusterConfigGceClusterConfigNodeGroupAffinity flattens an instance of ClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, i interface{}) *ClusterConfigGceClusterConfigNodeGroupAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigGceClusterConfigNodeGroupAffinity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigGceClusterConfigNodeGroupAffinity
	}
	r.NodeGroup = dcl.FlattenString(m["nodeGroupUri"])

	return r
}

// expandClusterConfigMasterConfigMap expands the contents of ClusterConfigMasterConfig into a JSON
// request object.
func expandClusterConfigMasterConfigMap(c *Client, f map[string]ClusterConfigMasterConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigMasterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigMasterConfigSlice expands the contents of ClusterConfigMasterConfig into a JSON
// request object.
func expandClusterConfigMasterConfigSlice(c *Client, f []ClusterConfigMasterConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigMasterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigMasterConfigMap flattens the contents of ClusterConfigMasterConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigMap(c *Client, i interface{}) map[string]ClusterConfigMasterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigMasterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigMasterConfig{}
	}

	items := make(map[string]ClusterConfigMasterConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigMasterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigMasterConfigSlice flattens the contents of ClusterConfigMasterConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigSlice(c *Client, i interface{}) []ClusterConfigMasterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigMasterConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigMasterConfig{}
	}

	items := make([]ClusterConfigMasterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigMasterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigMasterConfig expands an instance of ClusterConfigMasterConfig into a JSON
// request object.
func expandClusterConfigMasterConfig(c *Client, f *ClusterConfigMasterConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumInstances; !dcl.IsEmptyValueIndirect(v) {
		m["numInstances"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["imageUri"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineTypeUri"] = v
	}
	if v, err := expandClusterConfigMasterConfigDiskConfig(c, f.DiskConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding DiskConfig into diskConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskConfig"] = v
	}
	if v := f.Preemptibility; !dcl.IsEmptyValueIndirect(v) {
		m["preemptibility"] = v
	}
	if v, err := expandClusterConfigMasterConfigAcceleratorsSlice(c, f.Accelerators, res); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if v != nil {
		m["accelerators"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}

	return m, nil
}

// flattenClusterConfigMasterConfig flattens an instance of ClusterConfigMasterConfig from a JSON
// response object.
func flattenClusterConfigMasterConfig(c *Client, i interface{}) *ClusterConfigMasterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigMasterConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigMasterConfig
	}
	r.NumInstances = dcl.FlattenInteger(m["numInstances"])
	r.InstanceNames = dcl.FlattenStringSlice(m["instanceNames"])
	r.Image = dcl.FlattenString(m["imageUri"])
	r.MachineType = dcl.FlattenString(m["machineTypeUri"])
	r.DiskConfig = flattenClusterConfigMasterConfigDiskConfig(c, m["diskConfig"])
	r.IsPreemptible = dcl.FlattenBool(m["isPreemptible"])
	r.Preemptibility = flattenClusterConfigMasterConfigPreemptibilityEnum(m["preemptibility"])
	r.ManagedGroupConfig = flattenClusterConfigMasterConfigManagedGroupConfig(c, m["managedGroupConfig"])
	r.Accelerators = flattenClusterConfigMasterConfigAcceleratorsSlice(c, m["accelerators"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])

	return r
}

// expandClusterConfigMasterConfigDiskConfigMap expands the contents of ClusterConfigMasterConfigDiskConfig into a JSON
// request object.
func expandClusterConfigMasterConfigDiskConfigMap(c *Client, f map[string]ClusterConfigMasterConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigMasterConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigMasterConfigDiskConfigSlice expands the contents of ClusterConfigMasterConfigDiskConfig into a JSON
// request object.
func expandClusterConfigMasterConfigDiskConfigSlice(c *Client, f []ClusterConfigMasterConfigDiskConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigMasterConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigMasterConfigDiskConfigMap flattens the contents of ClusterConfigMasterConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigDiskConfigMap(c *Client, i interface{}) map[string]ClusterConfigMasterConfigDiskConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigMasterConfigDiskConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigMasterConfigDiskConfig{}
	}

	items := make(map[string]ClusterConfigMasterConfigDiskConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigMasterConfigDiskConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigMasterConfigDiskConfigSlice flattens the contents of ClusterConfigMasterConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigDiskConfigSlice(c *Client, i interface{}) []ClusterConfigMasterConfigDiskConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigMasterConfigDiskConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigMasterConfigDiskConfig{}
	}

	items := make([]ClusterConfigMasterConfigDiskConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigMasterConfigDiskConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigMasterConfigDiskConfig expands an instance of ClusterConfigMasterConfigDiskConfig into a JSON
// request object.
func expandClusterConfigMasterConfigDiskConfig(c *Client, f *ClusterConfigMasterConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BootDiskType; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskType"] = v
	}
	if v := f.BootDiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskSizeGb"] = v
	}
	if v := f.NumLocalSsds; !dcl.IsEmptyValueIndirect(v) {
		m["numLocalSsds"] = v
	}

	return m, nil
}

// flattenClusterConfigMasterConfigDiskConfig flattens an instance of ClusterConfigMasterConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigDiskConfig(c *Client, i interface{}) *ClusterConfigMasterConfigDiskConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigMasterConfigDiskConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigMasterConfigDiskConfig
	}
	r.BootDiskType = dcl.FlattenString(m["bootDiskType"])
	r.BootDiskSizeGb = dcl.FlattenInteger(m["bootDiskSizeGb"])
	r.NumLocalSsds = dcl.FlattenInteger(m["numLocalSsds"])

	return r
}

// expandClusterConfigMasterConfigManagedGroupConfigMap expands the contents of ClusterConfigMasterConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigMasterConfigManagedGroupConfigMap(c *Client, f map[string]ClusterConfigMasterConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigMasterConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigMasterConfigManagedGroupConfigSlice expands the contents of ClusterConfigMasterConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigMasterConfigManagedGroupConfigSlice(c *Client, f []ClusterConfigMasterConfigManagedGroupConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigMasterConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigMasterConfigManagedGroupConfigMap flattens the contents of ClusterConfigMasterConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigManagedGroupConfigMap(c *Client, i interface{}) map[string]ClusterConfigMasterConfigManagedGroupConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigMasterConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigMasterConfigManagedGroupConfig{}
	}

	items := make(map[string]ClusterConfigMasterConfigManagedGroupConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigMasterConfigManagedGroupConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigMasterConfigManagedGroupConfigSlice flattens the contents of ClusterConfigMasterConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigManagedGroupConfigSlice(c *Client, i interface{}) []ClusterConfigMasterConfigManagedGroupConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigMasterConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigMasterConfigManagedGroupConfig{}
	}

	items := make([]ClusterConfigMasterConfigManagedGroupConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigMasterConfigManagedGroupConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigMasterConfigManagedGroupConfig expands an instance of ClusterConfigMasterConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigMasterConfigManagedGroupConfig(c *Client, f *ClusterConfigMasterConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenClusterConfigMasterConfigManagedGroupConfig flattens an instance of ClusterConfigMasterConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigMasterConfigManagedGroupConfig(c *Client, i interface{}) *ClusterConfigMasterConfigManagedGroupConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigMasterConfigManagedGroupConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigMasterConfigManagedGroupConfig
	}
	r.InstanceTemplateName = dcl.FlattenString(m["instanceTemplateName"])
	r.InstanceGroupManagerName = dcl.FlattenString(m["instanceGroupManagerName"])

	return r
}

// expandClusterConfigMasterConfigAcceleratorsMap expands the contents of ClusterConfigMasterConfigAccelerators into a JSON
// request object.
func expandClusterConfigMasterConfigAcceleratorsMap(c *Client, f map[string]ClusterConfigMasterConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigMasterConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigMasterConfigAcceleratorsSlice expands the contents of ClusterConfigMasterConfigAccelerators into a JSON
// request object.
func expandClusterConfigMasterConfigAcceleratorsSlice(c *Client, f []ClusterConfigMasterConfigAccelerators, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigMasterConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigMasterConfigAcceleratorsMap flattens the contents of ClusterConfigMasterConfigAccelerators from a JSON
// response object.
func flattenClusterConfigMasterConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterConfigMasterConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigMasterConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigMasterConfigAccelerators{}
	}

	items := make(map[string]ClusterConfigMasterConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterConfigMasterConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigMasterConfigAcceleratorsSlice flattens the contents of ClusterConfigMasterConfigAccelerators from a JSON
// response object.
func flattenClusterConfigMasterConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterConfigMasterConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigMasterConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterConfigMasterConfigAccelerators{}
	}

	items := make([]ClusterConfigMasterConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigMasterConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigMasterConfigAccelerators expands an instance of ClusterConfigMasterConfigAccelerators into a JSON
// request object.
func expandClusterConfigMasterConfigAccelerators(c *Client, f *ClusterConfigMasterConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorTypeUri"] = v
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}

	return m, nil
}

// flattenClusterConfigMasterConfigAccelerators flattens an instance of ClusterConfigMasterConfigAccelerators from a JSON
// response object.
func flattenClusterConfigMasterConfigAccelerators(c *Client, i interface{}) *ClusterConfigMasterConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigMasterConfigAccelerators{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigMasterConfigAccelerators
	}
	r.AcceleratorType = dcl.FlattenString(m["acceleratorTypeUri"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandClusterConfigWorkerConfigMap expands the contents of ClusterConfigWorkerConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigMap(c *Client, f map[string]ClusterConfigWorkerConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigWorkerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigWorkerConfigSlice expands the contents of ClusterConfigWorkerConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigSlice(c *Client, f []ClusterConfigWorkerConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigWorkerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigWorkerConfigMap flattens the contents of ClusterConfigWorkerConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigMap(c *Client, i interface{}) map[string]ClusterConfigWorkerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigWorkerConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigWorkerConfig{}
	}

	items := make(map[string]ClusterConfigWorkerConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigWorkerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigWorkerConfigSlice flattens the contents of ClusterConfigWorkerConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigSlice(c *Client, i interface{}) []ClusterConfigWorkerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigWorkerConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigWorkerConfig{}
	}

	items := make([]ClusterConfigWorkerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigWorkerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigWorkerConfig expands an instance of ClusterConfigWorkerConfig into a JSON
// request object.
func expandClusterConfigWorkerConfig(c *Client, f *ClusterConfigWorkerConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumInstances; !dcl.IsEmptyValueIndirect(v) {
		m["numInstances"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["imageUri"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineTypeUri"] = v
	}
	if v, err := expandClusterConfigWorkerConfigDiskConfig(c, f.DiskConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding DiskConfig into diskConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskConfig"] = v
	}
	if v := f.Preemptibility; !dcl.IsEmptyValueIndirect(v) {
		m["preemptibility"] = v
	}
	if v, err := expandClusterConfigWorkerConfigAcceleratorsSlice(c, f.Accelerators, res); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if v != nil {
		m["accelerators"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}

	return m, nil
}

// flattenClusterConfigWorkerConfig flattens an instance of ClusterConfigWorkerConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfig(c *Client, i interface{}) *ClusterConfigWorkerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigWorkerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigWorkerConfig
	}
	r.NumInstances = dcl.FlattenInteger(m["numInstances"])
	r.InstanceNames = dcl.FlattenStringSlice(m["instanceNames"])
	r.Image = dcl.FlattenString(m["imageUri"])
	r.MachineType = dcl.FlattenString(m["machineTypeUri"])
	r.DiskConfig = flattenClusterConfigWorkerConfigDiskConfig(c, m["diskConfig"])
	r.IsPreemptible = dcl.FlattenBool(m["isPreemptible"])
	r.Preemptibility = flattenClusterConfigWorkerConfigPreemptibilityEnum(m["preemptibility"])
	r.ManagedGroupConfig = flattenClusterConfigWorkerConfigManagedGroupConfig(c, m["managedGroupConfig"])
	r.Accelerators = flattenClusterConfigWorkerConfigAcceleratorsSlice(c, m["accelerators"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])

	return r
}

// expandClusterConfigWorkerConfigDiskConfigMap expands the contents of ClusterConfigWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigDiskConfigMap(c *Client, f map[string]ClusterConfigWorkerConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigWorkerConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigWorkerConfigDiskConfigSlice expands the contents of ClusterConfigWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigDiskConfigSlice(c *Client, f []ClusterConfigWorkerConfigDiskConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigWorkerConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigWorkerConfigDiskConfigMap flattens the contents of ClusterConfigWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigDiskConfigMap(c *Client, i interface{}) map[string]ClusterConfigWorkerConfigDiskConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigWorkerConfigDiskConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigWorkerConfigDiskConfig{}
	}

	items := make(map[string]ClusterConfigWorkerConfigDiskConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigWorkerConfigDiskConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigWorkerConfigDiskConfigSlice flattens the contents of ClusterConfigWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigDiskConfigSlice(c *Client, i interface{}) []ClusterConfigWorkerConfigDiskConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigWorkerConfigDiskConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigWorkerConfigDiskConfig{}
	}

	items := make([]ClusterConfigWorkerConfigDiskConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigWorkerConfigDiskConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigWorkerConfigDiskConfig expands an instance of ClusterConfigWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigDiskConfig(c *Client, f *ClusterConfigWorkerConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BootDiskType; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskType"] = v
	}
	if v := f.BootDiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskSizeGb"] = v
	}
	if v := f.NumLocalSsds; !dcl.IsEmptyValueIndirect(v) {
		m["numLocalSsds"] = v
	}

	return m, nil
}

// flattenClusterConfigWorkerConfigDiskConfig flattens an instance of ClusterConfigWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigDiskConfig(c *Client, i interface{}) *ClusterConfigWorkerConfigDiskConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigWorkerConfigDiskConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigWorkerConfigDiskConfig
	}
	r.BootDiskType = dcl.FlattenString(m["bootDiskType"])
	r.BootDiskSizeGb = dcl.FlattenInteger(m["bootDiskSizeGb"])
	r.NumLocalSsds = dcl.FlattenInteger(m["numLocalSsds"])

	return r
}

// expandClusterConfigWorkerConfigManagedGroupConfigMap expands the contents of ClusterConfigWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigManagedGroupConfigMap(c *Client, f map[string]ClusterConfigWorkerConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigWorkerConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigWorkerConfigManagedGroupConfigSlice expands the contents of ClusterConfigWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigManagedGroupConfigSlice(c *Client, f []ClusterConfigWorkerConfigManagedGroupConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigWorkerConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigWorkerConfigManagedGroupConfigMap flattens the contents of ClusterConfigWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigManagedGroupConfigMap(c *Client, i interface{}) map[string]ClusterConfigWorkerConfigManagedGroupConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigWorkerConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigWorkerConfigManagedGroupConfig{}
	}

	items := make(map[string]ClusterConfigWorkerConfigManagedGroupConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigWorkerConfigManagedGroupConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigWorkerConfigManagedGroupConfigSlice flattens the contents of ClusterConfigWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigManagedGroupConfigSlice(c *Client, i interface{}) []ClusterConfigWorkerConfigManagedGroupConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigWorkerConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigWorkerConfigManagedGroupConfig{}
	}

	items := make([]ClusterConfigWorkerConfigManagedGroupConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigWorkerConfigManagedGroupConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigWorkerConfigManagedGroupConfig expands an instance of ClusterConfigWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigWorkerConfigManagedGroupConfig(c *Client, f *ClusterConfigWorkerConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenClusterConfigWorkerConfigManagedGroupConfig flattens an instance of ClusterConfigWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigWorkerConfigManagedGroupConfig(c *Client, i interface{}) *ClusterConfigWorkerConfigManagedGroupConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigWorkerConfigManagedGroupConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigWorkerConfigManagedGroupConfig
	}
	r.InstanceTemplateName = dcl.FlattenString(m["instanceTemplateName"])
	r.InstanceGroupManagerName = dcl.FlattenString(m["instanceGroupManagerName"])

	return r
}

// expandClusterConfigWorkerConfigAcceleratorsMap expands the contents of ClusterConfigWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigWorkerConfigAcceleratorsMap(c *Client, f map[string]ClusterConfigWorkerConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigWorkerConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigWorkerConfigAcceleratorsSlice expands the contents of ClusterConfigWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigWorkerConfigAcceleratorsSlice(c *Client, f []ClusterConfigWorkerConfigAccelerators, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigWorkerConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigWorkerConfigAcceleratorsMap flattens the contents of ClusterConfigWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigWorkerConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterConfigWorkerConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigWorkerConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigWorkerConfigAccelerators{}
	}

	items := make(map[string]ClusterConfigWorkerConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterConfigWorkerConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigWorkerConfigAcceleratorsSlice flattens the contents of ClusterConfigWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigWorkerConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterConfigWorkerConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigWorkerConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterConfigWorkerConfigAccelerators{}
	}

	items := make([]ClusterConfigWorkerConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigWorkerConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigWorkerConfigAccelerators expands an instance of ClusterConfigWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigWorkerConfigAccelerators(c *Client, f *ClusterConfigWorkerConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorTypeUri"] = v
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}

	return m, nil
}

// flattenClusterConfigWorkerConfigAccelerators flattens an instance of ClusterConfigWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigWorkerConfigAccelerators(c *Client, i interface{}) *ClusterConfigWorkerConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigWorkerConfigAccelerators{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigWorkerConfigAccelerators
	}
	r.AcceleratorType = dcl.FlattenString(m["acceleratorTypeUri"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandClusterConfigSecondaryWorkerConfigMap expands the contents of ClusterConfigSecondaryWorkerConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigMap(c *Client, f map[string]ClusterConfigSecondaryWorkerConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecondaryWorkerConfigSlice expands the contents of ClusterConfigSecondaryWorkerConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigSlice(c *Client, f []ClusterConfigSecondaryWorkerConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecondaryWorkerConfigMap flattens the contents of ClusterConfigSecondaryWorkerConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigMap(c *Client, i interface{}) map[string]ClusterConfigSecondaryWorkerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecondaryWorkerConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecondaryWorkerConfig{}
	}

	items := make(map[string]ClusterConfigSecondaryWorkerConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecondaryWorkerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigSlice flattens the contents of ClusterConfigSecondaryWorkerConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigSlice(c *Client, i interface{}) []ClusterConfigSecondaryWorkerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecondaryWorkerConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecondaryWorkerConfig{}
	}

	items := make([]ClusterConfigSecondaryWorkerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecondaryWorkerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecondaryWorkerConfig expands an instance of ClusterConfigSecondaryWorkerConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfig(c *Client, f *ClusterConfigSecondaryWorkerConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumInstances; !dcl.IsEmptyValueIndirect(v) {
		m["numInstances"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["imageUri"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineTypeUri"] = v
	}
	if v, err := expandClusterConfigSecondaryWorkerConfigDiskConfig(c, f.DiskConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding DiskConfig into diskConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskConfig"] = v
	}
	if v := f.Preemptibility; !dcl.IsEmptyValueIndirect(v) {
		m["preemptibility"] = v
	}
	if v, err := expandClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c, f.Accelerators, res); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if v != nil {
		m["accelerators"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}

	return m, nil
}

// flattenClusterConfigSecondaryWorkerConfig flattens an instance of ClusterConfigSecondaryWorkerConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfig(c *Client, i interface{}) *ClusterConfigSecondaryWorkerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecondaryWorkerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecondaryWorkerConfig
	}
	r.NumInstances = dcl.FlattenInteger(m["numInstances"])
	r.InstanceNames = dcl.FlattenStringSlice(m["instanceNames"])
	r.Image = dcl.FlattenString(m["imageUri"])
	r.MachineType = dcl.FlattenString(m["machineTypeUri"])
	r.DiskConfig = flattenClusterConfigSecondaryWorkerConfigDiskConfig(c, m["diskConfig"])
	r.IsPreemptible = dcl.FlattenBool(m["isPreemptible"])
	r.Preemptibility = flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnum(m["preemptibility"])
	r.ManagedGroupConfig = flattenClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, m["managedGroupConfig"])
	r.Accelerators = flattenClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c, m["accelerators"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])

	return r
}

// expandClusterConfigSecondaryWorkerConfigDiskConfigMap expands the contents of ClusterConfigSecondaryWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigDiskConfigMap(c *Client, f map[string]ClusterConfigSecondaryWorkerConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecondaryWorkerConfigDiskConfigSlice expands the contents of ClusterConfigSecondaryWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigDiskConfigSlice(c *Client, f []ClusterConfigSecondaryWorkerConfigDiskConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigDiskConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecondaryWorkerConfigDiskConfigMap flattens the contents of ClusterConfigSecondaryWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigDiskConfigMap(c *Client, i interface{}) map[string]ClusterConfigSecondaryWorkerConfigDiskConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}

	items := make(map[string]ClusterConfigSecondaryWorkerConfigDiskConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecondaryWorkerConfigDiskConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigDiskConfigSlice flattens the contents of ClusterConfigSecondaryWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigDiskConfigSlice(c *Client, i interface{}) []ClusterConfigSecondaryWorkerConfigDiskConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}

	items := make([]ClusterConfigSecondaryWorkerConfigDiskConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecondaryWorkerConfigDiskConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecondaryWorkerConfigDiskConfig expands an instance of ClusterConfigSecondaryWorkerConfigDiskConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigDiskConfig(c *Client, f *ClusterConfigSecondaryWorkerConfigDiskConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BootDiskType; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskType"] = v
	}
	if v := f.BootDiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskSizeGb"] = v
	}
	if v := f.NumLocalSsds; !dcl.IsEmptyValueIndirect(v) {
		m["numLocalSsds"] = v
	}

	return m, nil
}

// flattenClusterConfigSecondaryWorkerConfigDiskConfig flattens an instance of ClusterConfigSecondaryWorkerConfigDiskConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigDiskConfig(c *Client, i interface{}) *ClusterConfigSecondaryWorkerConfigDiskConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecondaryWorkerConfigDiskConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecondaryWorkerConfigDiskConfig
	}
	r.BootDiskType = dcl.FlattenString(m["bootDiskType"])
	r.BootDiskSizeGb = dcl.FlattenInteger(m["bootDiskSizeGb"])
	r.NumLocalSsds = dcl.FlattenInteger(m["numLocalSsds"])

	return r
}

// expandClusterConfigSecondaryWorkerConfigManagedGroupConfigMap expands the contents of ClusterConfigSecondaryWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigManagedGroupConfigMap(c *Client, f map[string]ClusterConfigSecondaryWorkerConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice expands the contents of ClusterConfigSecondaryWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice(c *Client, f []ClusterConfigSecondaryWorkerConfigManagedGroupConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecondaryWorkerConfigManagedGroupConfigMap flattens the contents of ClusterConfigSecondaryWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigManagedGroupConfigMap(c *Client, i interface{}) map[string]ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}

	items := make(map[string]ClusterConfigSecondaryWorkerConfigManagedGroupConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice flattens the contents of ClusterConfigSecondaryWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigManagedGroupConfigSlice(c *Client, i interface{}) []ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}

	items := make([]ClusterConfigSecondaryWorkerConfigManagedGroupConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecondaryWorkerConfigManagedGroupConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecondaryWorkerConfigManagedGroupConfig expands an instance of ClusterConfigSecondaryWorkerConfigManagedGroupConfig into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigManagedGroupConfig(c *Client, f *ClusterConfigSecondaryWorkerConfigManagedGroupConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenClusterConfigSecondaryWorkerConfigManagedGroupConfig flattens an instance of ClusterConfigSecondaryWorkerConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigManagedGroupConfig(c *Client, i interface{}) *ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecondaryWorkerConfigManagedGroupConfig
	}
	r.InstanceTemplateName = dcl.FlattenString(m["instanceTemplateName"])
	r.InstanceGroupManagerName = dcl.FlattenString(m["instanceGroupManagerName"])

	return r
}

// expandClusterConfigSecondaryWorkerConfigAcceleratorsMap expands the contents of ClusterConfigSecondaryWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigAcceleratorsMap(c *Client, f map[string]ClusterConfigSecondaryWorkerConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecondaryWorkerConfigAcceleratorsSlice expands the contents of ClusterConfigSecondaryWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c *Client, f []ClusterConfigSecondaryWorkerConfigAccelerators, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecondaryWorkerConfigAccelerators(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecondaryWorkerConfigAcceleratorsMap flattens the contents of ClusterConfigSecondaryWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterConfigSecondaryWorkerConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecondaryWorkerConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecondaryWorkerConfigAccelerators{}
	}

	items := make(map[string]ClusterConfigSecondaryWorkerConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecondaryWorkerConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigAcceleratorsSlice flattens the contents of ClusterConfigSecondaryWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterConfigSecondaryWorkerConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecondaryWorkerConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecondaryWorkerConfigAccelerators{}
	}

	items := make([]ClusterConfigSecondaryWorkerConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecondaryWorkerConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecondaryWorkerConfigAccelerators expands an instance of ClusterConfigSecondaryWorkerConfigAccelerators into a JSON
// request object.
func expandClusterConfigSecondaryWorkerConfigAccelerators(c *Client, f *ClusterConfigSecondaryWorkerConfigAccelerators, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorTypeUri"] = v
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}

	return m, nil
}

// flattenClusterConfigSecondaryWorkerConfigAccelerators flattens an instance of ClusterConfigSecondaryWorkerConfigAccelerators from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigAccelerators(c *Client, i interface{}) *ClusterConfigSecondaryWorkerConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecondaryWorkerConfigAccelerators{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecondaryWorkerConfigAccelerators
	}
	r.AcceleratorType = dcl.FlattenString(m["acceleratorTypeUri"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandClusterConfigSoftwareConfigMap expands the contents of ClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterConfigSoftwareConfigMap(c *Client, f map[string]ClusterConfigSoftwareConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSoftwareConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSoftwareConfigSlice expands the contents of ClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterConfigSoftwareConfigSlice(c *Client, f []ClusterConfigSoftwareConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSoftwareConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSoftwareConfigMap flattens the contents of ClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterConfigSoftwareConfigMap(c *Client, i interface{}) map[string]ClusterConfigSoftwareConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSoftwareConfig{}
	}

	items := make(map[string]ClusterConfigSoftwareConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSoftwareConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSoftwareConfigSlice flattens the contents of ClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterConfigSoftwareConfigSlice(c *Client, i interface{}) []ClusterConfigSoftwareConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSoftwareConfig{}
	}

	items := make([]ClusterConfigSoftwareConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSoftwareConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSoftwareConfig expands an instance of ClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterConfigSoftwareConfig(c *Client, f *ClusterConfigSoftwareConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ImageVersion; !dcl.IsEmptyValueIndirect(v) {
		m["imageVersion"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.OptionalComponents; v != nil {
		m["optionalComponents"] = v
	}

	return m, nil
}

// flattenClusterConfigSoftwareConfig flattens an instance of ClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterConfigSoftwareConfig(c *Client, i interface{}) *ClusterConfigSoftwareConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSoftwareConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSoftwareConfig
	}
	r.ImageVersion = dcl.FlattenString(m["imageVersion"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.OptionalComponents = flattenClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c, m["optionalComponents"])

	return r
}

// expandClusterConfigInitializationActionsMap expands the contents of ClusterConfigInitializationActions into a JSON
// request object.
func expandClusterConfigInitializationActionsMap(c *Client, f map[string]ClusterConfigInitializationActions, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigInitializationActions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigInitializationActionsSlice expands the contents of ClusterConfigInitializationActions into a JSON
// request object.
func expandClusterConfigInitializationActionsSlice(c *Client, f []ClusterConfigInitializationActions, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigInitializationActions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigInitializationActionsMap flattens the contents of ClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterConfigInitializationActionsMap(c *Client, i interface{}) map[string]ClusterConfigInitializationActions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigInitializationActions{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigInitializationActions{}
	}

	items := make(map[string]ClusterConfigInitializationActions)
	for k, item := range a {
		items[k] = *flattenClusterConfigInitializationActions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigInitializationActionsSlice flattens the contents of ClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterConfigInitializationActionsSlice(c *Client, i interface{}) []ClusterConfigInitializationActions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigInitializationActions{}
	}

	if len(a) == 0 {
		return []ClusterConfigInitializationActions{}
	}

	items := make([]ClusterConfigInitializationActions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigInitializationActions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigInitializationActions expands an instance of ClusterConfigInitializationActions into a JSON
// request object.
func expandClusterConfigInitializationActions(c *Client, f *ClusterConfigInitializationActions, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ExecutableFile; !dcl.IsEmptyValueIndirect(v) {
		m["executableFile"] = v
	}
	if v := f.ExecutionTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["executionTimeout"] = v
	}

	return m, nil
}

// flattenClusterConfigInitializationActions flattens an instance of ClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterConfigInitializationActions(c *Client, i interface{}) *ClusterConfigInitializationActions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigInitializationActions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigInitializationActions
	}
	r.ExecutableFile = dcl.FlattenString(m["executableFile"])
	r.ExecutionTimeout = dcl.FlattenString(m["executionTimeout"])

	return r
}

// expandClusterConfigEncryptionConfigMap expands the contents of ClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterConfigEncryptionConfigMap(c *Client, f map[string]ClusterConfigEncryptionConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigEncryptionConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigEncryptionConfigSlice expands the contents of ClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterConfigEncryptionConfigSlice(c *Client, f []ClusterConfigEncryptionConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigEncryptionConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigEncryptionConfigMap flattens the contents of ClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterConfigEncryptionConfigMap(c *Client, i interface{}) map[string]ClusterConfigEncryptionConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigEncryptionConfig{}
	}

	items := make(map[string]ClusterConfigEncryptionConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigEncryptionConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigEncryptionConfigSlice flattens the contents of ClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterConfigEncryptionConfigSlice(c *Client, i interface{}) []ClusterConfigEncryptionConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigEncryptionConfig{}
	}

	items := make([]ClusterConfigEncryptionConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigEncryptionConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigEncryptionConfig expands an instance of ClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterConfigEncryptionConfig(c *Client, f *ClusterConfigEncryptionConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.GcePdKmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["gcePdKmsKeyName"] = v
	}

	return m, nil
}

// flattenClusterConfigEncryptionConfig flattens an instance of ClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterConfigEncryptionConfig(c *Client, i interface{}) *ClusterConfigEncryptionConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigEncryptionConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigEncryptionConfig
	}
	r.GcePdKmsKeyName = dcl.FlattenString(m["gcePdKmsKeyName"])

	return r
}

// expandClusterConfigAutoscalingConfigMap expands the contents of ClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterConfigAutoscalingConfigMap(c *Client, f map[string]ClusterConfigAutoscalingConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigAutoscalingConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigAutoscalingConfigSlice expands the contents of ClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterConfigAutoscalingConfigSlice(c *Client, f []ClusterConfigAutoscalingConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigAutoscalingConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigAutoscalingConfigMap flattens the contents of ClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterConfigAutoscalingConfigMap(c *Client, i interface{}) map[string]ClusterConfigAutoscalingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigAutoscalingConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigAutoscalingConfig{}
	}

	items := make(map[string]ClusterConfigAutoscalingConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigAutoscalingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigAutoscalingConfigSlice flattens the contents of ClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterConfigAutoscalingConfigSlice(c *Client, i interface{}) []ClusterConfigAutoscalingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigAutoscalingConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigAutoscalingConfig{}
	}

	items := make([]ClusterConfigAutoscalingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigAutoscalingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigAutoscalingConfig expands an instance of ClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterConfigAutoscalingConfig(c *Client, f *ClusterConfigAutoscalingConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Policy; !dcl.IsEmptyValueIndirect(v) {
		m["policyUri"] = v
	}

	return m, nil
}

// flattenClusterConfigAutoscalingConfig flattens an instance of ClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterConfigAutoscalingConfig(c *Client, i interface{}) *ClusterConfigAutoscalingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigAutoscalingConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigAutoscalingConfig
	}
	r.Policy = dcl.FlattenString(m["policyUri"])

	return r
}

// expandClusterConfigSecurityConfigMap expands the contents of ClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterConfigSecurityConfigMap(c *Client, f map[string]ClusterConfigSecurityConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecurityConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecurityConfigSlice expands the contents of ClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterConfigSecurityConfigSlice(c *Client, f []ClusterConfigSecurityConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecurityConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecurityConfigMap flattens the contents of ClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfigMap(c *Client, i interface{}) map[string]ClusterConfigSecurityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecurityConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecurityConfig{}
	}

	items := make(map[string]ClusterConfigSecurityConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecurityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecurityConfigSlice flattens the contents of ClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfigSlice(c *Client, i interface{}) []ClusterConfigSecurityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecurityConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecurityConfig{}
	}

	items := make([]ClusterConfigSecurityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecurityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecurityConfig expands an instance of ClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterConfigSecurityConfig(c *Client, f *ClusterConfigSecurityConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandClusterConfigSecurityConfigKerberosConfig(c, f.KerberosConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding KerberosConfig into kerberosConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kerberosConfig"] = v
	}

	return m, nil
}

// flattenClusterConfigSecurityConfig flattens an instance of ClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfig(c *Client, i interface{}) *ClusterConfigSecurityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecurityConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecurityConfig
	}
	r.KerberosConfig = flattenClusterConfigSecurityConfigKerberosConfig(c, m["kerberosConfig"])

	return r
}

// expandClusterConfigSecurityConfigKerberosConfigMap expands the contents of ClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterConfigSecurityConfigKerberosConfigMap(c *Client, f map[string]ClusterConfigSecurityConfigKerberosConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigSecurityConfigKerberosConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigSecurityConfigKerberosConfigSlice expands the contents of ClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterConfigSecurityConfigKerberosConfigSlice(c *Client, f []ClusterConfigSecurityConfigKerberosConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigSecurityConfigKerberosConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigSecurityConfigKerberosConfigMap flattens the contents of ClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfigKerberosConfigMap(c *Client, i interface{}) map[string]ClusterConfigSecurityConfigKerberosConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecurityConfigKerberosConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecurityConfigKerberosConfig{}
	}

	items := make(map[string]ClusterConfigSecurityConfigKerberosConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecurityConfigKerberosConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigSecurityConfigKerberosConfigSlice flattens the contents of ClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfigKerberosConfigSlice(c *Client, i interface{}) []ClusterConfigSecurityConfigKerberosConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecurityConfigKerberosConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecurityConfigKerberosConfig{}
	}

	items := make([]ClusterConfigSecurityConfigKerberosConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecurityConfigKerberosConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigSecurityConfigKerberosConfig expands an instance of ClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterConfigSecurityConfigKerberosConfig(c *Client, f *ClusterConfigSecurityConfigKerberosConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableKerberos; !dcl.IsEmptyValueIndirect(v) {
		m["enableKerberos"] = v
	}
	if v := f.RootPrincipalPassword; !dcl.IsEmptyValueIndirect(v) {
		m["rootPrincipalPasswordUri"] = v
	}
	if v := f.KmsKey; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyUri"] = v
	}
	if v := f.Keystore; !dcl.IsEmptyValueIndirect(v) {
		m["keystoreUri"] = v
	}
	if v := f.Truststore; !dcl.IsEmptyValueIndirect(v) {
		m["truststoreUri"] = v
	}
	if v := f.KeystorePassword; !dcl.IsEmptyValueIndirect(v) {
		m["keystorePasswordUri"] = v
	}
	if v := f.KeyPassword; !dcl.IsEmptyValueIndirect(v) {
		m["keyPasswordUri"] = v
	}
	if v := f.TruststorePassword; !dcl.IsEmptyValueIndirect(v) {
		m["truststorePasswordUri"] = v
	}
	if v := f.CrossRealmTrustRealm; !dcl.IsEmptyValueIndirect(v) {
		m["crossRealmTrustRealm"] = v
	}
	if v := f.CrossRealmTrustKdc; !dcl.IsEmptyValueIndirect(v) {
		m["crossRealmTrustKdc"] = v
	}
	if v := f.CrossRealmTrustAdminServer; !dcl.IsEmptyValueIndirect(v) {
		m["crossRealmTrustAdminServer"] = v
	}
	if v := f.CrossRealmTrustSharedPassword; !dcl.IsEmptyValueIndirect(v) {
		m["crossRealmTrustSharedPasswordUri"] = v
	}
	if v := f.KdcDbKey; !dcl.IsEmptyValueIndirect(v) {
		m["kdcDbKeyUri"] = v
	}
	if v := f.TgtLifetimeHours; !dcl.IsEmptyValueIndirect(v) {
		m["tgtLifetimeHours"] = v
	}
	if v := f.Realm; !dcl.IsEmptyValueIndirect(v) {
		m["realm"] = v
	}

	return m, nil
}

// flattenClusterConfigSecurityConfigKerberosConfig flattens an instance of ClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterConfigSecurityConfigKerberosConfig(c *Client, i interface{}) *ClusterConfigSecurityConfigKerberosConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigSecurityConfigKerberosConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigSecurityConfigKerberosConfig
	}
	r.EnableKerberos = dcl.FlattenBool(m["enableKerberos"])
	r.RootPrincipalPassword = dcl.FlattenString(m["rootPrincipalPasswordUri"])
	r.KmsKey = dcl.FlattenString(m["kmsKeyUri"])
	r.Keystore = dcl.FlattenString(m["keystoreUri"])
	r.Truststore = dcl.FlattenString(m["truststoreUri"])
	r.KeystorePassword = dcl.FlattenString(m["keystorePasswordUri"])
	r.KeyPassword = dcl.FlattenString(m["keyPasswordUri"])
	r.TruststorePassword = dcl.FlattenString(m["truststorePasswordUri"])
	r.CrossRealmTrustRealm = dcl.FlattenString(m["crossRealmTrustRealm"])
	r.CrossRealmTrustKdc = dcl.FlattenString(m["crossRealmTrustKdc"])
	r.CrossRealmTrustAdminServer = dcl.FlattenString(m["crossRealmTrustAdminServer"])
	r.CrossRealmTrustSharedPassword = dcl.FlattenString(m["crossRealmTrustSharedPasswordUri"])
	r.KdcDbKey = dcl.FlattenString(m["kdcDbKeyUri"])
	r.TgtLifetimeHours = dcl.FlattenInteger(m["tgtLifetimeHours"])
	r.Realm = dcl.FlattenString(m["realm"])

	return r
}

// expandClusterConfigLifecycleConfigMap expands the contents of ClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterConfigLifecycleConfigMap(c *Client, f map[string]ClusterConfigLifecycleConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigLifecycleConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigLifecycleConfigSlice expands the contents of ClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterConfigLifecycleConfigSlice(c *Client, f []ClusterConfigLifecycleConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigLifecycleConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigLifecycleConfigMap flattens the contents of ClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterConfigLifecycleConfigMap(c *Client, i interface{}) map[string]ClusterConfigLifecycleConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigLifecycleConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigLifecycleConfig{}
	}

	items := make(map[string]ClusterConfigLifecycleConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigLifecycleConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigLifecycleConfigSlice flattens the contents of ClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterConfigLifecycleConfigSlice(c *Client, i interface{}) []ClusterConfigLifecycleConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigLifecycleConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigLifecycleConfig{}
	}

	items := make([]ClusterConfigLifecycleConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigLifecycleConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigLifecycleConfig expands an instance of ClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterConfigLifecycleConfig(c *Client, f *ClusterConfigLifecycleConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IdleDeleteTtl; !dcl.IsEmptyValueIndirect(v) {
		m["idleDeleteTtl"] = v
	}
	if v := f.AutoDeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["autoDeleteTime"] = v
	}
	if v := f.AutoDeleteTtl; !dcl.IsEmptyValueIndirect(v) {
		m["autoDeleteTtl"] = v
	}

	return m, nil
}

// flattenClusterConfigLifecycleConfig flattens an instance of ClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterConfigLifecycleConfig(c *Client, i interface{}) *ClusterConfigLifecycleConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigLifecycleConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigLifecycleConfig
	}
	r.IdleDeleteTtl = dcl.FlattenString(m["idleDeleteTtl"])
	r.AutoDeleteTime = dcl.FlattenString(m["autoDeleteTime"])
	r.AutoDeleteTtl = dcl.FlattenString(m["autoDeleteTtl"])
	r.IdleStartTime = dcl.FlattenString(m["idleStartTime"])

	return r
}

// expandClusterConfigEndpointConfigMap expands the contents of ClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterConfigEndpointConfigMap(c *Client, f map[string]ClusterConfigEndpointConfig, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfigEndpointConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfigEndpointConfigSlice expands the contents of ClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterConfigEndpointConfigSlice(c *Client, f []ClusterConfigEndpointConfig, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfigEndpointConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfigEndpointConfigMap flattens the contents of ClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterConfigEndpointConfigMap(c *Client, i interface{}) map[string]ClusterConfigEndpointConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigEndpointConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigEndpointConfig{}
	}

	items := make(map[string]ClusterConfigEndpointConfig)
	for k, item := range a {
		items[k] = *flattenClusterConfigEndpointConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfigEndpointConfigSlice flattens the contents of ClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterConfigEndpointConfigSlice(c *Client, i interface{}) []ClusterConfigEndpointConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigEndpointConfig{}
	}

	if len(a) == 0 {
		return []ClusterConfigEndpointConfig{}
	}

	items := make([]ClusterConfigEndpointConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigEndpointConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfigEndpointConfig expands an instance of ClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterConfigEndpointConfig(c *Client, f *ClusterConfigEndpointConfig, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableHttpPortAccess; !dcl.IsEmptyValueIndirect(v) {
		m["enableHttpPortAccess"] = v
	}

	return m, nil
}

// flattenClusterConfigEndpointConfig flattens an instance of ClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterConfigEndpointConfig(c *Client, i interface{}) *ClusterConfigEndpointConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfigEndpointConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterConfigEndpointConfig
	}
	r.HttpPorts = dcl.FlattenKeyValuePairs(m["httpPorts"])
	r.EnableHttpPortAccess = dcl.FlattenBool(m["enableHttpPortAccess"])

	return r
}

// expandClusterStatusMap expands the contents of ClusterStatus into a JSON
// request object.
func expandClusterStatusMap(c *Client, f map[string]ClusterStatus, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterStatus(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterStatusSlice expands the contents of ClusterStatus into a JSON
// request object.
func expandClusterStatusSlice(c *Client, f []ClusterStatus, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterStatus(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterStatusMap flattens the contents of ClusterStatus from a JSON
// response object.
func flattenClusterStatusMap(c *Client, i interface{}) map[string]ClusterStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatus{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatus{}
	}

	items := make(map[string]ClusterStatus)
	for k, item := range a {
		items[k] = *flattenClusterStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterStatusSlice flattens the contents of ClusterStatus from a JSON
// response object.
func flattenClusterStatusSlice(c *Client, i interface{}) []ClusterStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatus{}
	}

	if len(a) == 0 {
		return []ClusterStatus{}
	}

	items := make([]ClusterStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterStatus expands an instance of ClusterStatus into a JSON
// request object.
func expandClusterStatus(c *Client, f *ClusterStatus, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenClusterStatus flattens an instance of ClusterStatus from a JSON
// response object.
func flattenClusterStatus(c *Client, i interface{}) *ClusterStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterStatus
	}
	r.State = flattenClusterStatusStateEnum(m["state"])
	r.Detail = dcl.FlattenString(m["detail"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenClusterStatusSubstateEnum(m["substate"])

	return r
}

// expandClusterStatusHistoryMap expands the contents of ClusterStatusHistory into a JSON
// request object.
func expandClusterStatusHistoryMap(c *Client, f map[string]ClusterStatusHistory, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterStatusHistory(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterStatusHistorySlice expands the contents of ClusterStatusHistory into a JSON
// request object.
func expandClusterStatusHistorySlice(c *Client, f []ClusterStatusHistory, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterStatusHistory(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterStatusHistoryMap flattens the contents of ClusterStatusHistory from a JSON
// response object.
func flattenClusterStatusHistoryMap(c *Client, i interface{}) map[string]ClusterStatusHistory {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatusHistory{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatusHistory{}
	}

	items := make(map[string]ClusterStatusHistory)
	for k, item := range a {
		items[k] = *flattenClusterStatusHistory(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterStatusHistorySlice flattens the contents of ClusterStatusHistory from a JSON
// response object.
func flattenClusterStatusHistorySlice(c *Client, i interface{}) []ClusterStatusHistory {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatusHistory{}
	}

	if len(a) == 0 {
		return []ClusterStatusHistory{}
	}

	items := make([]ClusterStatusHistory, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatusHistory(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterStatusHistory expands an instance of ClusterStatusHistory into a JSON
// request object.
func expandClusterStatusHistory(c *Client, f *ClusterStatusHistory, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenClusterStatusHistory flattens an instance of ClusterStatusHistory from a JSON
// response object.
func flattenClusterStatusHistory(c *Client, i interface{}) *ClusterStatusHistory {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterStatusHistory{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterStatusHistory
	}
	r.State = flattenClusterStatusHistoryStateEnum(m["state"])
	r.Detail = dcl.FlattenString(m["detail"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenClusterStatusHistorySubstateEnum(m["substate"])

	return r
}

// expandClusterMetricsMap expands the contents of ClusterMetrics into a JSON
// request object.
func expandClusterMetricsMap(c *Client, f map[string]ClusterMetrics, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMetrics(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMetricsSlice expands the contents of ClusterMetrics into a JSON
// request object.
func expandClusterMetricsSlice(c *Client, f []ClusterMetrics, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMetrics(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMetricsMap flattens the contents of ClusterMetrics from a JSON
// response object.
func flattenClusterMetricsMap(c *Client, i interface{}) map[string]ClusterMetrics {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMetrics{}
	}

	if len(a) == 0 {
		return map[string]ClusterMetrics{}
	}

	items := make(map[string]ClusterMetrics)
	for k, item := range a {
		items[k] = *flattenClusterMetrics(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMetricsSlice flattens the contents of ClusterMetrics from a JSON
// response object.
func flattenClusterMetricsSlice(c *Client, i interface{}) []ClusterMetrics {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMetrics{}
	}

	if len(a) == 0 {
		return []ClusterMetrics{}
	}

	items := make([]ClusterMetrics, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMetrics(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMetrics expands an instance of ClusterMetrics into a JSON
// request object.
func expandClusterMetrics(c *Client, f *ClusterMetrics, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HdfsMetrics; !dcl.IsEmptyValueIndirect(v) {
		m["hdfsMetrics"] = v
	}
	if v := f.YarnMetrics; !dcl.IsEmptyValueIndirect(v) {
		m["yarnMetrics"] = v
	}

	return m, nil
}

// flattenClusterMetrics flattens an instance of ClusterMetrics from a JSON
// response object.
func flattenClusterMetrics(c *Client, i interface{}) *ClusterMetrics {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMetrics{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterMetrics
	}
	r.HdfsMetrics = dcl.FlattenKeyValuePairs(m["hdfsMetrics"])
	r.YarnMetrics = dcl.FlattenKeyValuePairs(m["yarnMetrics"])

	return r
}

// flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumMap flattens the contents of ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum from a JSON
// response object.
func flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumMap(c *Client, i interface{}) map[string]ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	items := make(map[string]ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumSlice flattens the contents of ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum from a JSON
// response object.
func flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumSlice(c *Client, i interface{}) []ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	items := make([]ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum with the same value as that string.
func flattenClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(i interface{}) *ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s)
}

// flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumMap flattens the contents of ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumMap(c *Client, i interface{}) map[string]ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make(map[string]ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumSlice flattens the contents of ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, i interface{}) []ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make([]ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum with the same value as that string.
func flattenClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(i interface{}) *ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s)
}

// flattenClusterConfigMasterConfigPreemptibilityEnumMap flattens the contents of ClusterConfigMasterConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigMasterConfigPreemptibilityEnumMap(c *Client, i interface{}) map[string]ClusterConfigMasterConfigPreemptibilityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigMasterConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigMasterConfigPreemptibilityEnum{}
	}

	items := make(map[string]ClusterConfigMasterConfigPreemptibilityEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigMasterConfigPreemptibilityEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigMasterConfigPreemptibilityEnumSlice flattens the contents of ClusterConfigMasterConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigMasterConfigPreemptibilityEnumSlice(c *Client, i interface{}) []ClusterConfigMasterConfigPreemptibilityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigMasterConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigMasterConfigPreemptibilityEnum{}
	}

	items := make([]ClusterConfigMasterConfigPreemptibilityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigMasterConfigPreemptibilityEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigMasterConfigPreemptibilityEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigMasterConfigPreemptibilityEnum with the same value as that string.
func flattenClusterConfigMasterConfigPreemptibilityEnum(i interface{}) *ClusterConfigMasterConfigPreemptibilityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigMasterConfigPreemptibilityEnumRef(s)
}

// flattenClusterConfigWorkerConfigPreemptibilityEnumMap flattens the contents of ClusterConfigWorkerConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigWorkerConfigPreemptibilityEnumMap(c *Client, i interface{}) map[string]ClusterConfigWorkerConfigPreemptibilityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigWorkerConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigWorkerConfigPreemptibilityEnum{}
	}

	items := make(map[string]ClusterConfigWorkerConfigPreemptibilityEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigWorkerConfigPreemptibilityEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigWorkerConfigPreemptibilityEnumSlice flattens the contents of ClusterConfigWorkerConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigWorkerConfigPreemptibilityEnumSlice(c *Client, i interface{}) []ClusterConfigWorkerConfigPreemptibilityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigWorkerConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigWorkerConfigPreemptibilityEnum{}
	}

	items := make([]ClusterConfigWorkerConfigPreemptibilityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigWorkerConfigPreemptibilityEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigWorkerConfigPreemptibilityEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigWorkerConfigPreemptibilityEnum with the same value as that string.
func flattenClusterConfigWorkerConfigPreemptibilityEnum(i interface{}) *ClusterConfigWorkerConfigPreemptibilityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigWorkerConfigPreemptibilityEnumRef(s)
}

// flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnumMap flattens the contents of ClusterConfigSecondaryWorkerConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnumMap(c *Client, i interface{}) map[string]ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSecondaryWorkerConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSecondaryWorkerConfigPreemptibilityEnum{}
	}

	items := make(map[string]ClusterConfigSecondaryWorkerConfigPreemptibilityEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnumSlice flattens the contents of ClusterConfigSecondaryWorkerConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnumSlice(c *Client, i interface{}) []ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSecondaryWorkerConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigSecondaryWorkerConfigPreemptibilityEnum{}
	}

	items := make([]ClusterConfigSecondaryWorkerConfigPreemptibilityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigSecondaryWorkerConfigPreemptibilityEnum with the same value as that string.
func flattenClusterConfigSecondaryWorkerConfigPreemptibilityEnum(i interface{}) *ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigSecondaryWorkerConfigPreemptibilityEnumRef(s)
}

// flattenClusterConfigSoftwareConfigOptionalComponentsEnumMap flattens the contents of ClusterConfigSoftwareConfigOptionalComponentsEnum from a JSON
// response object.
func flattenClusterConfigSoftwareConfigOptionalComponentsEnumMap(c *Client, i interface{}) map[string]ClusterConfigSoftwareConfigOptionalComponentsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	items := make(map[string]ClusterConfigSoftwareConfigOptionalComponentsEnum)
	for k, item := range a {
		items[k] = *flattenClusterConfigSoftwareConfigOptionalComponentsEnum(item.(interface{}))
	}

	return items
}

// flattenClusterConfigSoftwareConfigOptionalComponentsEnumSlice flattens the contents of ClusterConfigSoftwareConfigOptionalComponentsEnum from a JSON
// response object.
func flattenClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c *Client, i interface{}) []ClusterConfigSoftwareConfigOptionalComponentsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	if len(a) == 0 {
		return []ClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	items := make([]ClusterConfigSoftwareConfigOptionalComponentsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfigSoftwareConfigOptionalComponentsEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConfigSoftwareConfigOptionalComponentsEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConfigSoftwareConfigOptionalComponentsEnum with the same value as that string.
func flattenClusterConfigSoftwareConfigOptionalComponentsEnum(i interface{}) *ClusterConfigSoftwareConfigOptionalComponentsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterConfigSoftwareConfigOptionalComponentsEnumRef(s)
}

// flattenClusterStatusStateEnumMap flattens the contents of ClusterStatusStateEnum from a JSON
// response object.
func flattenClusterStatusStateEnumMap(c *Client, i interface{}) map[string]ClusterStatusStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatusStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatusStateEnum{}
	}

	items := make(map[string]ClusterStatusStateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStatusStateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStatusStateEnumSlice flattens the contents of ClusterStatusStateEnum from a JSON
// response object.
func flattenClusterStatusStateEnumSlice(c *Client, i interface{}) []ClusterStatusStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatusStateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStatusStateEnum{}
	}

	items := make([]ClusterStatusStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatusStateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStatusStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusStateEnum with the same value as that string.
func flattenClusterStatusStateEnum(i interface{}) *ClusterStatusStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterStatusStateEnumRef(s)
}

// flattenClusterStatusSubstateEnumMap flattens the contents of ClusterStatusSubstateEnum from a JSON
// response object.
func flattenClusterStatusSubstateEnumMap(c *Client, i interface{}) map[string]ClusterStatusSubstateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatusSubstateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatusSubstateEnum{}
	}

	items := make(map[string]ClusterStatusSubstateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStatusSubstateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStatusSubstateEnumSlice flattens the contents of ClusterStatusSubstateEnum from a JSON
// response object.
func flattenClusterStatusSubstateEnumSlice(c *Client, i interface{}) []ClusterStatusSubstateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatusSubstateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStatusSubstateEnum{}
	}

	items := make([]ClusterStatusSubstateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatusSubstateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStatusSubstateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusSubstateEnum with the same value as that string.
func flattenClusterStatusSubstateEnum(i interface{}) *ClusterStatusSubstateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterStatusSubstateEnumRef(s)
}

// flattenClusterStatusHistoryStateEnumMap flattens the contents of ClusterStatusHistoryStateEnum from a JSON
// response object.
func flattenClusterStatusHistoryStateEnumMap(c *Client, i interface{}) map[string]ClusterStatusHistoryStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatusHistoryStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatusHistoryStateEnum{}
	}

	items := make(map[string]ClusterStatusHistoryStateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStatusHistoryStateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStatusHistoryStateEnumSlice flattens the contents of ClusterStatusHistoryStateEnum from a JSON
// response object.
func flattenClusterStatusHistoryStateEnumSlice(c *Client, i interface{}) []ClusterStatusHistoryStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatusHistoryStateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStatusHistoryStateEnum{}
	}

	items := make([]ClusterStatusHistoryStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatusHistoryStateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStatusHistoryStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusHistoryStateEnum with the same value as that string.
func flattenClusterStatusHistoryStateEnum(i interface{}) *ClusterStatusHistoryStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterStatusHistoryStateEnumRef(s)
}

// flattenClusterStatusHistorySubstateEnumMap flattens the contents of ClusterStatusHistorySubstateEnum from a JSON
// response object.
func flattenClusterStatusHistorySubstateEnumMap(c *Client, i interface{}) map[string]ClusterStatusHistorySubstateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStatusHistorySubstateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStatusHistorySubstateEnum{}
	}

	items := make(map[string]ClusterStatusHistorySubstateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStatusHistorySubstateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStatusHistorySubstateEnumSlice flattens the contents of ClusterStatusHistorySubstateEnum from a JSON
// response object.
func flattenClusterStatusHistorySubstateEnumSlice(c *Client, i interface{}) []ClusterStatusHistorySubstateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStatusHistorySubstateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStatusHistorySubstateEnum{}
	}

	items := make([]ClusterStatusHistorySubstateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStatusHistorySubstateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStatusHistorySubstateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusHistorySubstateEnum with the same value as that string.
func flattenClusterStatusHistorySubstateEnum(i interface{}) *ClusterStatusHistorySubstateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterStatusHistorySubstateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Cluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCluster(b, c)
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

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
}

func convertFieldDiffsToClusterDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]clusterDiff, error) {
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
	var diffs []clusterDiff
	// For each operation name, create a clusterDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := clusterDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToClusterApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToClusterApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (clusterApiOperation, error) {
	switch opName {

	case "updateClusterUpdateClusterOperation":
		return &updateClusterUpdateClusterOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractClusterFields(r *Cluster) error {
	vConfig := r.Config
	if vConfig == nil {
		// note: explicitly not the empty object.
		vConfig = &ClusterConfig{}
	}
	if err := extractClusterConfigFields(r, vConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vConfig) {
		r.Config = vConfig
	}
	vStatus := r.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &ClusterStatus{}
	}
	if err := extractClusterStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStatus) {
		r.Status = vStatus
	}
	vMetrics := r.Metrics
	if vMetrics == nil {
		// note: explicitly not the empty object.
		vMetrics = &ClusterMetrics{}
	}
	if err := extractClusterMetricsFields(r, vMetrics); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMetrics) {
		r.Metrics = vMetrics
	}
	return nil
}
func extractClusterConfigFields(r *Cluster, o *ClusterConfig) error {
	vGceClusterConfig := o.GceClusterConfig
	if vGceClusterConfig == nil {
		// note: explicitly not the empty object.
		vGceClusterConfig = &ClusterConfigGceClusterConfig{}
	}
	if err := extractClusterConfigGceClusterConfigFields(r, vGceClusterConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGceClusterConfig) {
		o.GceClusterConfig = vGceClusterConfig
	}
	vMasterConfig := o.MasterConfig
	if vMasterConfig == nil {
		// note: explicitly not the empty object.
		vMasterConfig = &ClusterConfigMasterConfig{}
	}
	if err := extractClusterConfigMasterConfigFields(r, vMasterConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMasterConfig) {
		o.MasterConfig = vMasterConfig
	}
	vWorkerConfig := o.WorkerConfig
	if vWorkerConfig == nil {
		// note: explicitly not the empty object.
		vWorkerConfig = &ClusterConfigWorkerConfig{}
	}
	if err := extractClusterConfigWorkerConfigFields(r, vWorkerConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vWorkerConfig) {
		o.WorkerConfig = vWorkerConfig
	}
	vSecondaryWorkerConfig := o.SecondaryWorkerConfig
	if vSecondaryWorkerConfig == nil {
		// note: explicitly not the empty object.
		vSecondaryWorkerConfig = &ClusterConfigSecondaryWorkerConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigFields(r, vSecondaryWorkerConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryWorkerConfig) {
		o.SecondaryWorkerConfig = vSecondaryWorkerConfig
	}
	vSoftwareConfig := o.SoftwareConfig
	if vSoftwareConfig == nil {
		// note: explicitly not the empty object.
		vSoftwareConfig = &ClusterConfigSoftwareConfig{}
	}
	if err := extractClusterConfigSoftwareConfigFields(r, vSoftwareConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSoftwareConfig) {
		o.SoftwareConfig = vSoftwareConfig
	}
	vEncryptionConfig := o.EncryptionConfig
	if vEncryptionConfig == nil {
		// note: explicitly not the empty object.
		vEncryptionConfig = &ClusterConfigEncryptionConfig{}
	}
	if err := extractClusterConfigEncryptionConfigFields(r, vEncryptionConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEncryptionConfig) {
		o.EncryptionConfig = vEncryptionConfig
	}
	vAutoscalingConfig := o.AutoscalingConfig
	if vAutoscalingConfig == nil {
		// note: explicitly not the empty object.
		vAutoscalingConfig = &ClusterConfigAutoscalingConfig{}
	}
	if err := extractClusterConfigAutoscalingConfigFields(r, vAutoscalingConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAutoscalingConfig) {
		o.AutoscalingConfig = vAutoscalingConfig
	}
	vSecurityConfig := o.SecurityConfig
	if vSecurityConfig == nil {
		// note: explicitly not the empty object.
		vSecurityConfig = &ClusterConfigSecurityConfig{}
	}
	if err := extractClusterConfigSecurityConfigFields(r, vSecurityConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecurityConfig) {
		o.SecurityConfig = vSecurityConfig
	}
	vLifecycleConfig := o.LifecycleConfig
	if vLifecycleConfig == nil {
		// note: explicitly not the empty object.
		vLifecycleConfig = &ClusterConfigLifecycleConfig{}
	}
	if err := extractClusterConfigLifecycleConfigFields(r, vLifecycleConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vLifecycleConfig) {
		o.LifecycleConfig = vLifecycleConfig
	}
	vEndpointConfig := o.EndpointConfig
	if vEndpointConfig == nil {
		// note: explicitly not the empty object.
		vEndpointConfig = &ClusterConfigEndpointConfig{}
	}
	if err := extractClusterConfigEndpointConfigFields(r, vEndpointConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEndpointConfig) {
		o.EndpointConfig = vEndpointConfig
	}
	return nil
}
func extractClusterConfigGceClusterConfigFields(r *Cluster, o *ClusterConfigGceClusterConfig) error {
	vReservationAffinity := o.ReservationAffinity
	if vReservationAffinity == nil {
		// note: explicitly not the empty object.
		vReservationAffinity = &ClusterConfigGceClusterConfigReservationAffinity{}
	}
	if err := extractClusterConfigGceClusterConfigReservationAffinityFields(r, vReservationAffinity); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vReservationAffinity) {
		o.ReservationAffinity = vReservationAffinity
	}
	vNodeGroupAffinity := o.NodeGroupAffinity
	if vNodeGroupAffinity == nil {
		// note: explicitly not the empty object.
		vNodeGroupAffinity = &ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}
	if err := extractClusterConfigGceClusterConfigNodeGroupAffinityFields(r, vNodeGroupAffinity); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNodeGroupAffinity) {
		o.NodeGroupAffinity = vNodeGroupAffinity
	}
	return nil
}
func extractClusterConfigGceClusterConfigReservationAffinityFields(r *Cluster, o *ClusterConfigGceClusterConfigReservationAffinity) error {
	return nil
}
func extractClusterConfigGceClusterConfigNodeGroupAffinityFields(r *Cluster, o *ClusterConfigGceClusterConfigNodeGroupAffinity) error {
	return nil
}
func extractClusterConfigMasterConfigFields(r *Cluster, o *ClusterConfigMasterConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigMasterConfigDiskConfig{}
	}
	if err := extractClusterConfigMasterConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigMasterConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigMasterConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func extractClusterConfigMasterConfigDiskConfigFields(r *Cluster, o *ClusterConfigMasterConfigDiskConfig) error {
	return nil
}
func extractClusterConfigMasterConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigMasterConfigManagedGroupConfig) error {
	return nil
}
func extractClusterConfigMasterConfigAcceleratorsFields(r *Cluster, o *ClusterConfigMasterConfigAccelerators) error {
	return nil
}
func extractClusterConfigWorkerConfigFields(r *Cluster, o *ClusterConfigWorkerConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigWorkerConfigDiskConfig{}
	}
	if err := extractClusterConfigWorkerConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigWorkerConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigWorkerConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func extractClusterConfigWorkerConfigDiskConfigFields(r *Cluster, o *ClusterConfigWorkerConfigDiskConfig) error {
	return nil
}
func extractClusterConfigWorkerConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigWorkerConfigManagedGroupConfig) error {
	return nil
}
func extractClusterConfigWorkerConfigAcceleratorsFields(r *Cluster, o *ClusterConfigWorkerConfigAccelerators) error {
	return nil
}
func extractClusterConfigSecondaryWorkerConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func extractClusterConfigSecondaryWorkerConfigDiskConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigDiskConfig) error {
	return nil
}
func extractClusterConfigSecondaryWorkerConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigManagedGroupConfig) error {
	return nil
}
func extractClusterConfigSecondaryWorkerConfigAcceleratorsFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigAccelerators) error {
	return nil
}
func extractClusterConfigSoftwareConfigFields(r *Cluster, o *ClusterConfigSoftwareConfig) error {
	return nil
}
func extractClusterConfigInitializationActionsFields(r *Cluster, o *ClusterConfigInitializationActions) error {
	return nil
}
func extractClusterConfigEncryptionConfigFields(r *Cluster, o *ClusterConfigEncryptionConfig) error {
	return nil
}
func extractClusterConfigAutoscalingConfigFields(r *Cluster, o *ClusterConfigAutoscalingConfig) error {
	return nil
}
func extractClusterConfigSecurityConfigFields(r *Cluster, o *ClusterConfigSecurityConfig) error {
	vKerberosConfig := o.KerberosConfig
	if vKerberosConfig == nil {
		// note: explicitly not the empty object.
		vKerberosConfig = &ClusterConfigSecurityConfigKerberosConfig{}
	}
	if err := extractClusterConfigSecurityConfigKerberosConfigFields(r, vKerberosConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vKerberosConfig) {
		o.KerberosConfig = vKerberosConfig
	}
	return nil
}
func extractClusterConfigSecurityConfigKerberosConfigFields(r *Cluster, o *ClusterConfigSecurityConfigKerberosConfig) error {
	return nil
}
func extractClusterConfigLifecycleConfigFields(r *Cluster, o *ClusterConfigLifecycleConfig) error {
	return nil
}
func extractClusterConfigEndpointConfigFields(r *Cluster, o *ClusterConfigEndpointConfig) error {
	return nil
}
func extractClusterStatusFields(r *Cluster, o *ClusterStatus) error {
	return nil
}
func extractClusterStatusHistoryFields(r *Cluster, o *ClusterStatusHistory) error {
	return nil
}
func extractClusterMetricsFields(r *Cluster, o *ClusterMetrics) error {
	return nil
}

func postReadExtractClusterFields(r *Cluster) error {
	vConfig := r.Config
	if vConfig == nil {
		// note: explicitly not the empty object.
		vConfig = &ClusterConfig{}
	}
	if err := postReadExtractClusterConfigFields(r, vConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vConfig) {
		r.Config = vConfig
	}
	vStatus := r.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &ClusterStatus{}
	}
	if err := postReadExtractClusterStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStatus) {
		r.Status = vStatus
	}
	vMetrics := r.Metrics
	if vMetrics == nil {
		// note: explicitly not the empty object.
		vMetrics = &ClusterMetrics{}
	}
	if err := postReadExtractClusterMetricsFields(r, vMetrics); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMetrics) {
		r.Metrics = vMetrics
	}
	return nil
}
func postReadExtractClusterConfigFields(r *Cluster, o *ClusterConfig) error {
	vGceClusterConfig := o.GceClusterConfig
	if vGceClusterConfig == nil {
		// note: explicitly not the empty object.
		vGceClusterConfig = &ClusterConfigGceClusterConfig{}
	}
	if err := extractClusterConfigGceClusterConfigFields(r, vGceClusterConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGceClusterConfig) {
		o.GceClusterConfig = vGceClusterConfig
	}
	vMasterConfig := o.MasterConfig
	if vMasterConfig == nil {
		// note: explicitly not the empty object.
		vMasterConfig = &ClusterConfigMasterConfig{}
	}
	if err := extractClusterConfigMasterConfigFields(r, vMasterConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMasterConfig) {
		o.MasterConfig = vMasterConfig
	}
	vWorkerConfig := o.WorkerConfig
	if vWorkerConfig == nil {
		// note: explicitly not the empty object.
		vWorkerConfig = &ClusterConfigWorkerConfig{}
	}
	if err := extractClusterConfigWorkerConfigFields(r, vWorkerConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vWorkerConfig) {
		o.WorkerConfig = vWorkerConfig
	}
	vSecondaryWorkerConfig := o.SecondaryWorkerConfig
	if vSecondaryWorkerConfig == nil {
		// note: explicitly not the empty object.
		vSecondaryWorkerConfig = &ClusterConfigSecondaryWorkerConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigFields(r, vSecondaryWorkerConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryWorkerConfig) {
		o.SecondaryWorkerConfig = vSecondaryWorkerConfig
	}
	vSoftwareConfig := o.SoftwareConfig
	if vSoftwareConfig == nil {
		// note: explicitly not the empty object.
		vSoftwareConfig = &ClusterConfigSoftwareConfig{}
	}
	if err := extractClusterConfigSoftwareConfigFields(r, vSoftwareConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSoftwareConfig) {
		o.SoftwareConfig = vSoftwareConfig
	}
	vEncryptionConfig := o.EncryptionConfig
	if vEncryptionConfig == nil {
		// note: explicitly not the empty object.
		vEncryptionConfig = &ClusterConfigEncryptionConfig{}
	}
	if err := extractClusterConfigEncryptionConfigFields(r, vEncryptionConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEncryptionConfig) {
		o.EncryptionConfig = vEncryptionConfig
	}
	vAutoscalingConfig := o.AutoscalingConfig
	if vAutoscalingConfig == nil {
		// note: explicitly not the empty object.
		vAutoscalingConfig = &ClusterConfigAutoscalingConfig{}
	}
	if err := extractClusterConfigAutoscalingConfigFields(r, vAutoscalingConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAutoscalingConfig) {
		o.AutoscalingConfig = vAutoscalingConfig
	}
	vSecurityConfig := o.SecurityConfig
	if vSecurityConfig == nil {
		// note: explicitly not the empty object.
		vSecurityConfig = &ClusterConfigSecurityConfig{}
	}
	if err := extractClusterConfigSecurityConfigFields(r, vSecurityConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecurityConfig) {
		o.SecurityConfig = vSecurityConfig
	}
	vLifecycleConfig := o.LifecycleConfig
	if vLifecycleConfig == nil {
		// note: explicitly not the empty object.
		vLifecycleConfig = &ClusterConfigLifecycleConfig{}
	}
	if err := extractClusterConfigLifecycleConfigFields(r, vLifecycleConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vLifecycleConfig) {
		o.LifecycleConfig = vLifecycleConfig
	}
	vEndpointConfig := o.EndpointConfig
	if vEndpointConfig == nil {
		// note: explicitly not the empty object.
		vEndpointConfig = &ClusterConfigEndpointConfig{}
	}
	if err := extractClusterConfigEndpointConfigFields(r, vEndpointConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEndpointConfig) {
		o.EndpointConfig = vEndpointConfig
	}
	return nil
}
func postReadExtractClusterConfigGceClusterConfigFields(r *Cluster, o *ClusterConfigGceClusterConfig) error {
	vReservationAffinity := o.ReservationAffinity
	if vReservationAffinity == nil {
		// note: explicitly not the empty object.
		vReservationAffinity = &ClusterConfigGceClusterConfigReservationAffinity{}
	}
	if err := extractClusterConfigGceClusterConfigReservationAffinityFields(r, vReservationAffinity); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vReservationAffinity) {
		o.ReservationAffinity = vReservationAffinity
	}
	vNodeGroupAffinity := o.NodeGroupAffinity
	if vNodeGroupAffinity == nil {
		// note: explicitly not the empty object.
		vNodeGroupAffinity = &ClusterConfigGceClusterConfigNodeGroupAffinity{}
	}
	if err := extractClusterConfigGceClusterConfigNodeGroupAffinityFields(r, vNodeGroupAffinity); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNodeGroupAffinity) {
		o.NodeGroupAffinity = vNodeGroupAffinity
	}
	return nil
}
func postReadExtractClusterConfigGceClusterConfigReservationAffinityFields(r *Cluster, o *ClusterConfigGceClusterConfigReservationAffinity) error {
	return nil
}
func postReadExtractClusterConfigGceClusterConfigNodeGroupAffinityFields(r *Cluster, o *ClusterConfigGceClusterConfigNodeGroupAffinity) error {
	return nil
}
func postReadExtractClusterConfigMasterConfigFields(r *Cluster, o *ClusterConfigMasterConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigMasterConfigDiskConfig{}
	}
	if err := extractClusterConfigMasterConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigMasterConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigMasterConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func postReadExtractClusterConfigMasterConfigDiskConfigFields(r *Cluster, o *ClusterConfigMasterConfigDiskConfig) error {
	return nil
}
func postReadExtractClusterConfigMasterConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigMasterConfigManagedGroupConfig) error {
	return nil
}
func postReadExtractClusterConfigMasterConfigAcceleratorsFields(r *Cluster, o *ClusterConfigMasterConfigAccelerators) error {
	return nil
}
func postReadExtractClusterConfigWorkerConfigFields(r *Cluster, o *ClusterConfigWorkerConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigWorkerConfigDiskConfig{}
	}
	if err := extractClusterConfigWorkerConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigWorkerConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigWorkerConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func postReadExtractClusterConfigWorkerConfigDiskConfigFields(r *Cluster, o *ClusterConfigWorkerConfigDiskConfig) error {
	return nil
}
func postReadExtractClusterConfigWorkerConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigWorkerConfigManagedGroupConfig) error {
	return nil
}
func postReadExtractClusterConfigWorkerConfigAcceleratorsFields(r *Cluster, o *ClusterConfigWorkerConfigAccelerators) error {
	return nil
}
func postReadExtractClusterConfigSecondaryWorkerConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfig) error {
	vDiskConfig := o.DiskConfig
	if vDiskConfig == nil {
		// note: explicitly not the empty object.
		vDiskConfig = &ClusterConfigSecondaryWorkerConfigDiskConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigDiskConfigFields(r, vDiskConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiskConfig) {
		o.DiskConfig = vDiskConfig
	}
	vManagedGroupConfig := o.ManagedGroupConfig
	if vManagedGroupConfig == nil {
		// note: explicitly not the empty object.
		vManagedGroupConfig = &ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	}
	if err := extractClusterConfigSecondaryWorkerConfigManagedGroupConfigFields(r, vManagedGroupConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagedGroupConfig) {
		o.ManagedGroupConfig = vManagedGroupConfig
	}
	return nil
}
func postReadExtractClusterConfigSecondaryWorkerConfigDiskConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigDiskConfig) error {
	return nil
}
func postReadExtractClusterConfigSecondaryWorkerConfigManagedGroupConfigFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigManagedGroupConfig) error {
	return nil
}
func postReadExtractClusterConfigSecondaryWorkerConfigAcceleratorsFields(r *Cluster, o *ClusterConfigSecondaryWorkerConfigAccelerators) error {
	return nil
}
func postReadExtractClusterConfigSoftwareConfigFields(r *Cluster, o *ClusterConfigSoftwareConfig) error {
	return nil
}
func postReadExtractClusterConfigInitializationActionsFields(r *Cluster, o *ClusterConfigInitializationActions) error {
	return nil
}
func postReadExtractClusterConfigEncryptionConfigFields(r *Cluster, o *ClusterConfigEncryptionConfig) error {
	return nil
}
func postReadExtractClusterConfigAutoscalingConfigFields(r *Cluster, o *ClusterConfigAutoscalingConfig) error {
	return nil
}
func postReadExtractClusterConfigSecurityConfigFields(r *Cluster, o *ClusterConfigSecurityConfig) error {
	vKerberosConfig := o.KerberosConfig
	if vKerberosConfig == nil {
		// note: explicitly not the empty object.
		vKerberosConfig = &ClusterConfigSecurityConfigKerberosConfig{}
	}
	if err := extractClusterConfigSecurityConfigKerberosConfigFields(r, vKerberosConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vKerberosConfig) {
		o.KerberosConfig = vKerberosConfig
	}
	return nil
}
func postReadExtractClusterConfigSecurityConfigKerberosConfigFields(r *Cluster, o *ClusterConfigSecurityConfigKerberosConfig) error {
	return nil
}
func postReadExtractClusterConfigLifecycleConfigFields(r *Cluster, o *ClusterConfigLifecycleConfig) error {
	return nil
}
func postReadExtractClusterConfigEndpointConfigFields(r *Cluster, o *ClusterConfigEndpointConfig) error {
	return nil
}
func postReadExtractClusterStatusFields(r *Cluster, o *ClusterStatus) error {
	return nil
}
func postReadExtractClusterStatusHistoryFields(r *Cluster, o *ClusterStatusHistory) error {
	return nil
}
func postReadExtractClusterMetricsFields(r *Cluster, o *ClusterMetrics) error {
	return nil
}
