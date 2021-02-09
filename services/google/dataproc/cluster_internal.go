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
package dataproc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
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
func (r *ClusterClusterConfig) validate() error {
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
func (r *ClusterClusterConfigGceClusterConfig) validate() error {
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
func (r *ClusterClusterConfigGceClusterConfigReservationAffinity) validate() error {
	return nil
}
func (r *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) validate() error {
	if err := dcl.Required(r, "nodeGroup"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterInstanceGroupConfig) validate() error {
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
func (r *ClusterInstanceGroupConfigDiskConfig) validate() error {
	return nil
}
func (r *ClusterInstanceGroupConfigManagedGroupConfig) validate() error {
	return nil
}
func (r *ClusterInstanceGroupConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterClusterConfigSoftwareConfig) validate() error {
	return nil
}
func (r *ClusterClusterConfigInitializationActions) validate() error {
	if err := dcl.Required(r, "executableFile"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterClusterConfigEncryptionConfig) validate() error {
	return nil
}
func (r *ClusterClusterConfigAutoscalingConfig) validate() error {
	return nil
}
func (r *ClusterClusterConfigSecurityConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KerberosConfig) {
		if err := r.KerberosConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterClusterConfigSecurityConfigKerberosConfig) validate() error {
	return nil
}
func (r *ClusterClusterConfigLifecycleConfig) validate() error {
	return nil
}
func (r *ClusterClusterConfigEndpointConfig) validate() error {
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

func clusterGetURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
}

func clusterListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func clusterCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func clusterDeleteURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCluster")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateClusterRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateClusterRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://dataproc.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listClusterRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := clusterListURL(c.Config.BasePath, project, location)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listClusterOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listCluster(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Cluster, string, error) {
	b, err := c.listClusterRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Cluster
	for _, v := range m.Items {
		res := flattenCluster(c, v)
		res.Project = &project
		res.Location = &location
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

	_, err := c.GetCluster(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Cluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetCluster checking for existence. error: %v", err)
		return err
	}

	u, err := clusterDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://dataproc.googleapis.com/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetCluster(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createClusterOperation struct{}

func (op *createClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := clusterCreateURL(c.Config.BasePath, project, location)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://dataproc.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetCluster(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getClusterRaw(ctx context.Context, r *Cluster) ([]byte, error) {

	u, err := clusterGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

func (c *Client) clusterDiffsForRawDesired(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (initial, desired *Cluster, diffs []clusterDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Cluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCluster(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Cluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Cluster resource: %v", err)
		}
		c.Config.Logger.Info("Found that Cluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Cluster: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Cluster: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Cluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Cluster: %v", desired)

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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeClusterClusterConfig(rawDesired.Config, nil, opts...)
		rawDesired.Status = canonicalizeClusterStatus(rawDesired.Status, nil, opts...)
		rawDesired.Metrics = canonicalizeClusterMetrics(rawDesired.Metrics, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.Config = canonicalizeClusterClusterConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.Status = canonicalizeClusterStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.IsZeroValue(rawDesired.StatusHistory) {
		rawDesired.StatusHistory = rawInitial.StatusHistory
	}
	if dcl.IsZeroValue(rawDesired.ClusterUuid) {
		rawDesired.ClusterUuid = rawInitial.ClusterUuid
	}
	rawDesired.Metrics = canonicalizeClusterMetrics(rawDesired.Metrics, rawInitial.Metrics, opts...)
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewClusterClusterConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewClusterStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusHistory) && dcl.IsEmptyValueIndirect(rawDesired.StatusHistory) {
		rawNew.StatusHistory = rawDesired.StatusHistory
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClusterUuid) && dcl.IsEmptyValueIndirect(rawDesired.ClusterUuid) {
		rawNew.ClusterUuid = rawDesired.ClusterUuid
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Metrics) && dcl.IsEmptyValueIndirect(rawDesired.Metrics) {
		rawNew.Metrics = rawDesired.Metrics
	} else {
		rawNew.Metrics = canonicalizeNewClusterMetrics(c, rawDesired.Metrics, rawNew.Metrics)
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeClusterClusterConfig(des, initial *ClusterClusterConfig, opts ...dcl.ApplyOption) *ClusterClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.StagingBucket, initial.StagingBucket) || dcl.IsZeroValue(des.StagingBucket) {
		des.StagingBucket = initial.StagingBucket
	}
	if dcl.NameToSelfLink(des.TempBucket, initial.TempBucket) || dcl.IsZeroValue(des.TempBucket) {
		des.TempBucket = initial.TempBucket
	}
	des.GceClusterConfig = canonicalizeClusterClusterConfigGceClusterConfig(des.GceClusterConfig, initial.GceClusterConfig, opts...)
	des.MasterConfig = canonicalizeClusterInstanceGroupConfig(des.MasterConfig, initial.MasterConfig, opts...)
	des.WorkerConfig = canonicalizeClusterInstanceGroupConfig(des.WorkerConfig, initial.WorkerConfig, opts...)
	des.SecondaryWorkerConfig = canonicalizeClusterInstanceGroupConfig(des.SecondaryWorkerConfig, initial.SecondaryWorkerConfig, opts...)
	des.SoftwareConfig = canonicalizeClusterClusterConfigSoftwareConfig(des.SoftwareConfig, initial.SoftwareConfig, opts...)
	if dcl.IsZeroValue(des.InitializationActions) {
		des.InitializationActions = initial.InitializationActions
	}
	des.EncryptionConfig = canonicalizeClusterClusterConfigEncryptionConfig(des.EncryptionConfig, initial.EncryptionConfig, opts...)
	des.AutoscalingConfig = canonicalizeClusterClusterConfigAutoscalingConfig(des.AutoscalingConfig, initial.AutoscalingConfig, opts...)
	des.SecurityConfig = canonicalizeClusterClusterConfigSecurityConfig(des.SecurityConfig, initial.SecurityConfig, opts...)
	des.LifecycleConfig = canonicalizeClusterClusterConfigLifecycleConfig(des.LifecycleConfig, initial.LifecycleConfig, opts...)
	des.EndpointConfig = canonicalizeClusterClusterConfigEndpointConfig(des.EndpointConfig, initial.EndpointConfig, opts...)

	return des
}

func canonicalizeNewClusterClusterConfig(c *Client, des, nw *ClusterClusterConfig) *ClusterClusterConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.StagingBucket, nw.StagingBucket) || dcl.IsZeroValue(des.StagingBucket) {
		nw.StagingBucket = des.StagingBucket
	}
	if dcl.NameToSelfLink(des.TempBucket, nw.TempBucket) || dcl.IsZeroValue(des.TempBucket) {
		nw.TempBucket = des.TempBucket
	}
	nw.GceClusterConfig = canonicalizeNewClusterClusterConfigGceClusterConfig(c, des.GceClusterConfig, nw.GceClusterConfig)
	nw.MasterConfig = canonicalizeNewClusterInstanceGroupConfig(c, des.MasterConfig, nw.MasterConfig)
	nw.WorkerConfig = canonicalizeNewClusterInstanceGroupConfig(c, des.WorkerConfig, nw.WorkerConfig)
	nw.SecondaryWorkerConfig = canonicalizeNewClusterInstanceGroupConfig(c, des.SecondaryWorkerConfig, nw.SecondaryWorkerConfig)
	nw.SoftwareConfig = canonicalizeNewClusterClusterConfigSoftwareConfig(c, des.SoftwareConfig, nw.SoftwareConfig)
	nw.EncryptionConfig = canonicalizeNewClusterClusterConfigEncryptionConfig(c, des.EncryptionConfig, nw.EncryptionConfig)
	nw.AutoscalingConfig = canonicalizeNewClusterClusterConfigAutoscalingConfig(c, des.AutoscalingConfig, nw.AutoscalingConfig)
	nw.SecurityConfig = canonicalizeNewClusterClusterConfigSecurityConfig(c, des.SecurityConfig, nw.SecurityConfig)
	nw.LifecycleConfig = canonicalizeNewClusterClusterConfigLifecycleConfig(c, des.LifecycleConfig, nw.LifecycleConfig)
	nw.EndpointConfig = canonicalizeNewClusterClusterConfigEndpointConfig(c, des.EndpointConfig, nw.EndpointConfig)

	return nw
}

func canonicalizeNewClusterClusterConfigSet(c *Client, des, nw []ClusterClusterConfig) []ClusterClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigGceClusterConfig(des, initial *ClusterClusterConfigGceClusterConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigGceClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		des.Zone = initial.Zone
	}
	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, initial.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		des.Subnetwork = initial.Subnetwork
	}
	if dcl.IsZeroValue(des.InternalIPOnly) {
		des.InternalIPOnly = initial.InternalIPOnly
	}
	if dcl.IsZeroValue(des.PrivateIPv6GoogleAccess) {
		des.PrivateIPv6GoogleAccess = initial.PrivateIPv6GoogleAccess
	}
	if dcl.NameToSelfLink(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	if dcl.IsZeroValue(des.ServiceAccountScopes) {
		des.ServiceAccountScopes = initial.ServiceAccountScopes
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}
	des.ReservationAffinity = canonicalizeClusterClusterConfigGceClusterConfigReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	des.NodeGroupAffinity = canonicalizeClusterClusterConfigGceClusterConfigNodeGroupAffinity(des.NodeGroupAffinity, initial.NodeGroupAffinity, opts...)

	return des
}

func canonicalizeNewClusterClusterConfigGceClusterConfig(c *Client, des, nw *ClusterClusterConfigGceClusterConfig) *ClusterClusterConfigGceClusterConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Zone, nw.Zone) || dcl.IsZeroValue(des.Zone) {
		nw.Zone = des.Zone
	}
	if dcl.NameToSelfLink(des.Network, nw.Network) || dcl.IsZeroValue(des.Network) {
		nw.Network = des.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, nw.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}
	if dcl.NameToSelfLink(des.ServiceAccount, nw.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	nw.ReservationAffinity = canonicalizeNewClusterClusterConfigGceClusterConfigReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.NodeGroupAffinity = canonicalizeNewClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, des.NodeGroupAffinity, nw.NodeGroupAffinity)

	return nw
}

func canonicalizeNewClusterClusterConfigGceClusterConfigSet(c *Client, des, nw []ClusterClusterConfigGceClusterConfig) []ClusterClusterConfigGceClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigGceClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigGceClusterConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigGceClusterConfigReservationAffinity(des, initial *ClusterClusterConfigGceClusterConfigReservationAffinity, opts ...dcl.ApplyOption) *ClusterClusterConfigGceClusterConfigReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ConsumeReservationType) {
		des.ConsumeReservationType = initial.ConsumeReservationType
	}
	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewClusterClusterConfigGceClusterConfigReservationAffinity(c *Client, des, nw *ClusterClusterConfigGceClusterConfigReservationAffinity) *ClusterClusterConfigGceClusterConfigReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterClusterConfigGceClusterConfigReservationAffinitySet(c *Client, des, nw []ClusterClusterConfigGceClusterConfigReservationAffinity) []ClusterClusterConfigGceClusterConfigReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigGceClusterConfigReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigGceClusterConfigReservationAffinity(c, &d, &n) {
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

func canonicalizeClusterClusterConfigGceClusterConfigNodeGroupAffinity(des, initial *ClusterClusterConfigGceClusterConfigNodeGroupAffinity, opts ...dcl.ApplyOption) *ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.NodeGroup, initial.NodeGroup) || dcl.IsZeroValue(des.NodeGroup) {
		des.NodeGroup = initial.NodeGroup
	}

	return des
}

func canonicalizeNewClusterClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, des, nw *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NodeGroup, nw.NodeGroup) || dcl.IsZeroValue(des.NodeGroup) {
		nw.NodeGroup = des.NodeGroup
	}

	return nw
}

func canonicalizeNewClusterClusterConfigGceClusterConfigNodeGroupAffinitySet(c *Client, des, nw []ClusterClusterConfigGceClusterConfigNodeGroupAffinity) []ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigGceClusterConfigNodeGroupAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, &d, &n) {
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

func canonicalizeClusterInstanceGroupConfig(des, initial *ClusterInstanceGroupConfig, opts ...dcl.ApplyOption) *ClusterInstanceGroupConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumInstances) {
		des.NumInstances = initial.NumInstances
	}
	if dcl.IsZeroValue(des.InstanceNames) {
		des.InstanceNames = initial.InstanceNames
	}
	if dcl.NameToSelfLink(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		des.Image = initial.Image
	}
	if dcl.NameToSelfLink(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	des.DiskConfig = canonicalizeClusterInstanceGroupConfigDiskConfig(des.DiskConfig, initial.DiskConfig, opts...)
	if dcl.IsZeroValue(des.IsPreemptible) {
		des.IsPreemptible = initial.IsPreemptible
	}
	if dcl.IsZeroValue(des.Preemptibility) {
		des.Preemptibility = initial.Preemptibility
	}
	des.ManagedGroupConfig = canonicalizeClusterInstanceGroupConfigManagedGroupConfig(des.ManagedGroupConfig, initial.ManagedGroupConfig, opts...)
	if dcl.IsZeroValue(des.Accelerators) {
		des.Accelerators = initial.Accelerators
	}
	if dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}

	return des
}

func canonicalizeNewClusterInstanceGroupConfig(c *Client, des, nw *ClusterInstanceGroupConfig) *ClusterInstanceGroupConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Image, nw.Image) || dcl.IsZeroValue(des.Image) {
		nw.Image = des.Image
	}
	if dcl.NameToSelfLink(des.MachineType, nw.MachineType) || dcl.IsZeroValue(des.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.DiskConfig = canonicalizeNewClusterInstanceGroupConfigDiskConfig(c, des.DiskConfig, nw.DiskConfig)
	nw.ManagedGroupConfig = canonicalizeNewClusterInstanceGroupConfigManagedGroupConfig(c, des.ManagedGroupConfig, nw.ManagedGroupConfig)

	return nw
}

func canonicalizeNewClusterInstanceGroupConfigSet(c *Client, des, nw []ClusterInstanceGroupConfig) []ClusterInstanceGroupConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterInstanceGroupConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterInstanceGroupConfig(c, &d, &n) {
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

func canonicalizeClusterInstanceGroupConfigDiskConfig(des, initial *ClusterInstanceGroupConfigDiskConfig, opts ...dcl.ApplyOption) *ClusterInstanceGroupConfigDiskConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BootDiskType) {
		des.BootDiskType = initial.BootDiskType
	}
	if dcl.IsZeroValue(des.BootDiskSizeGb) {
		des.BootDiskSizeGb = initial.BootDiskSizeGb
	}
	if dcl.IsZeroValue(des.NumLocalSsds) {
		des.NumLocalSsds = initial.NumLocalSsds
	}

	return des
}

func canonicalizeNewClusterInstanceGroupConfigDiskConfig(c *Client, des, nw *ClusterInstanceGroupConfigDiskConfig) *ClusterInstanceGroupConfigDiskConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterInstanceGroupConfigDiskConfigSet(c *Client, des, nw []ClusterInstanceGroupConfigDiskConfig) []ClusterInstanceGroupConfigDiskConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterInstanceGroupConfigDiskConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterInstanceGroupConfigDiskConfig(c, &d, &n) {
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

func canonicalizeClusterInstanceGroupConfigManagedGroupConfig(des, initial *ClusterInstanceGroupConfigManagedGroupConfig, opts ...dcl.ApplyOption) *ClusterInstanceGroupConfigManagedGroupConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.InstanceTemplateName) {
		des.InstanceTemplateName = initial.InstanceTemplateName
	}
	if dcl.IsZeroValue(des.InstanceGroupManagerName) {
		des.InstanceGroupManagerName = initial.InstanceGroupManagerName
	}

	return des
}

func canonicalizeNewClusterInstanceGroupConfigManagedGroupConfig(c *Client, des, nw *ClusterInstanceGroupConfigManagedGroupConfig) *ClusterInstanceGroupConfigManagedGroupConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterInstanceGroupConfigManagedGroupConfigSet(c *Client, des, nw []ClusterInstanceGroupConfigManagedGroupConfig) []ClusterInstanceGroupConfigManagedGroupConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterInstanceGroupConfigManagedGroupConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterInstanceGroupConfigManagedGroupConfig(c, &d, &n) {
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

func canonicalizeClusterInstanceGroupConfigAccelerators(des, initial *ClusterInstanceGroupConfigAccelerators, opts ...dcl.ApplyOption) *ClusterInstanceGroupConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}
	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}

	return des
}

func canonicalizeNewClusterInstanceGroupConfigAccelerators(c *Client, des, nw *ClusterInstanceGroupConfigAccelerators) *ClusterInstanceGroupConfigAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.AcceleratorType, nw.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterInstanceGroupConfigAcceleratorsSet(c *Client, des, nw []ClusterInstanceGroupConfigAccelerators) []ClusterInstanceGroupConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterInstanceGroupConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterInstanceGroupConfigAccelerators(c, &d, &n) {
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

func canonicalizeClusterClusterConfigSoftwareConfig(des, initial *ClusterClusterConfigSoftwareConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigSoftwareConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ImageVersion) {
		des.ImageVersion = initial.ImageVersion
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	if dcl.IsZeroValue(des.OptionalComponents) {
		des.OptionalComponents = initial.OptionalComponents
	}

	return des
}

func canonicalizeNewClusterClusterConfigSoftwareConfig(c *Client, des, nw *ClusterClusterConfigSoftwareConfig) *ClusterClusterConfigSoftwareConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterClusterConfigSoftwareConfigSet(c *Client, des, nw []ClusterClusterConfigSoftwareConfig) []ClusterClusterConfigSoftwareConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigSoftwareConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigSoftwareConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigInitializationActions(des, initial *ClusterClusterConfigInitializationActions, opts ...dcl.ApplyOption) *ClusterClusterConfigInitializationActions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ExecutableFile) {
		des.ExecutableFile = initial.ExecutableFile
	}
	if dcl.IsZeroValue(des.ExecutionTimeout) {
		des.ExecutionTimeout = initial.ExecutionTimeout
	}

	return des
}

func canonicalizeNewClusterClusterConfigInitializationActions(c *Client, des, nw *ClusterClusterConfigInitializationActions) *ClusterClusterConfigInitializationActions {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterClusterConfigInitializationActionsSet(c *Client, des, nw []ClusterClusterConfigInitializationActions) []ClusterClusterConfigInitializationActions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigInitializationActions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigInitializationActions(c, &d, &n) {
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

func canonicalizeClusterClusterConfigEncryptionConfig(des, initial *ClusterClusterConfigEncryptionConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigEncryptionConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.GcePdKmsKeyName, initial.GcePdKmsKeyName) || dcl.IsZeroValue(des.GcePdKmsKeyName) {
		des.GcePdKmsKeyName = initial.GcePdKmsKeyName
	}

	return des
}

func canonicalizeNewClusterClusterConfigEncryptionConfig(c *Client, des, nw *ClusterClusterConfigEncryptionConfig) *ClusterClusterConfigEncryptionConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.GcePdKmsKeyName, nw.GcePdKmsKeyName) || dcl.IsZeroValue(des.GcePdKmsKeyName) {
		nw.GcePdKmsKeyName = des.GcePdKmsKeyName
	}

	return nw
}

func canonicalizeNewClusterClusterConfigEncryptionConfigSet(c *Client, des, nw []ClusterClusterConfigEncryptionConfig) []ClusterClusterConfigEncryptionConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigEncryptionConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigEncryptionConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigAutoscalingConfig(des, initial *ClusterClusterConfigAutoscalingConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigAutoscalingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Policy, initial.Policy) || dcl.IsZeroValue(des.Policy) {
		des.Policy = initial.Policy
	}

	return des
}

func canonicalizeNewClusterClusterConfigAutoscalingConfig(c *Client, des, nw *ClusterClusterConfigAutoscalingConfig) *ClusterClusterConfigAutoscalingConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Policy, nw.Policy) || dcl.IsZeroValue(des.Policy) {
		nw.Policy = des.Policy
	}

	return nw
}

func canonicalizeNewClusterClusterConfigAutoscalingConfigSet(c *Client, des, nw []ClusterClusterConfigAutoscalingConfig) []ClusterClusterConfigAutoscalingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigAutoscalingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigAutoscalingConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigSecurityConfig(des, initial *ClusterClusterConfigSecurityConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigSecurityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.KerberosConfig = canonicalizeClusterClusterConfigSecurityConfigKerberosConfig(des.KerberosConfig, initial.KerberosConfig, opts...)

	return des
}

func canonicalizeNewClusterClusterConfigSecurityConfig(c *Client, des, nw *ClusterClusterConfigSecurityConfig) *ClusterClusterConfigSecurityConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.KerberosConfig = canonicalizeNewClusterClusterConfigSecurityConfigKerberosConfig(c, des.KerberosConfig, nw.KerberosConfig)

	return nw
}

func canonicalizeNewClusterClusterConfigSecurityConfigSet(c *Client, des, nw []ClusterClusterConfigSecurityConfig) []ClusterClusterConfigSecurityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigSecurityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigSecurityConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigSecurityConfigKerberosConfig(des, initial *ClusterClusterConfigSecurityConfigKerberosConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigSecurityConfigKerberosConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EnableKerberos) {
		des.EnableKerberos = initial.EnableKerberos
	}
	if dcl.NameToSelfLink(des.RootPrincipalPassword, initial.RootPrincipalPassword) || dcl.IsZeroValue(des.RootPrincipalPassword) {
		des.RootPrincipalPassword = initial.RootPrincipalPassword
	}
	if dcl.NameToSelfLink(des.KmsKey, initial.KmsKey) || dcl.IsZeroValue(des.KmsKey) {
		des.KmsKey = initial.KmsKey
	}
	if dcl.NameToSelfLink(des.Keystore, initial.Keystore) || dcl.IsZeroValue(des.Keystore) {
		des.Keystore = initial.Keystore
	}
	if dcl.NameToSelfLink(des.Truststore, initial.Truststore) || dcl.IsZeroValue(des.Truststore) {
		des.Truststore = initial.Truststore
	}
	if dcl.NameToSelfLink(des.KeystorePassword, initial.KeystorePassword) || dcl.IsZeroValue(des.KeystorePassword) {
		des.KeystorePassword = initial.KeystorePassword
	}
	if dcl.NameToSelfLink(des.KeyPassword, initial.KeyPassword) || dcl.IsZeroValue(des.KeyPassword) {
		des.KeyPassword = initial.KeyPassword
	}
	if dcl.NameToSelfLink(des.TruststorePassword, initial.TruststorePassword) || dcl.IsZeroValue(des.TruststorePassword) {
		des.TruststorePassword = initial.TruststorePassword
	}
	if dcl.IsZeroValue(des.CrossRealmTrustRealm) {
		des.CrossRealmTrustRealm = initial.CrossRealmTrustRealm
	}
	if dcl.IsZeroValue(des.CrossRealmTrustKdc) {
		des.CrossRealmTrustKdc = initial.CrossRealmTrustKdc
	}
	if dcl.IsZeroValue(des.CrossRealmTrustAdminServer) {
		des.CrossRealmTrustAdminServer = initial.CrossRealmTrustAdminServer
	}
	if dcl.NameToSelfLink(des.CrossRealmTrustSharedPassword, initial.CrossRealmTrustSharedPassword) || dcl.IsZeroValue(des.CrossRealmTrustSharedPassword) {
		des.CrossRealmTrustSharedPassword = initial.CrossRealmTrustSharedPassword
	}
	if dcl.NameToSelfLink(des.KdcDbKey, initial.KdcDbKey) || dcl.IsZeroValue(des.KdcDbKey) {
		des.KdcDbKey = initial.KdcDbKey
	}
	if dcl.IsZeroValue(des.TgtLifetimeHours) {
		des.TgtLifetimeHours = initial.TgtLifetimeHours
	}
	if dcl.IsZeroValue(des.Realm) {
		des.Realm = initial.Realm
	}

	return des
}

func canonicalizeNewClusterClusterConfigSecurityConfigKerberosConfig(c *Client, des, nw *ClusterClusterConfigSecurityConfigKerberosConfig) *ClusterClusterConfigSecurityConfigKerberosConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.RootPrincipalPassword, nw.RootPrincipalPassword) || dcl.IsZeroValue(des.RootPrincipalPassword) {
		nw.RootPrincipalPassword = des.RootPrincipalPassword
	}
	if dcl.NameToSelfLink(des.KmsKey, nw.KmsKey) || dcl.IsZeroValue(des.KmsKey) {
		nw.KmsKey = des.KmsKey
	}
	if dcl.NameToSelfLink(des.Keystore, nw.Keystore) || dcl.IsZeroValue(des.Keystore) {
		nw.Keystore = des.Keystore
	}
	if dcl.NameToSelfLink(des.Truststore, nw.Truststore) || dcl.IsZeroValue(des.Truststore) {
		nw.Truststore = des.Truststore
	}
	if dcl.NameToSelfLink(des.KeystorePassword, nw.KeystorePassword) || dcl.IsZeroValue(des.KeystorePassword) {
		nw.KeystorePassword = des.KeystorePassword
	}
	if dcl.NameToSelfLink(des.KeyPassword, nw.KeyPassword) || dcl.IsZeroValue(des.KeyPassword) {
		nw.KeyPassword = des.KeyPassword
	}
	if dcl.NameToSelfLink(des.TruststorePassword, nw.TruststorePassword) || dcl.IsZeroValue(des.TruststorePassword) {
		nw.TruststorePassword = des.TruststorePassword
	}
	if dcl.NameToSelfLink(des.CrossRealmTrustSharedPassword, nw.CrossRealmTrustSharedPassword) || dcl.IsZeroValue(des.CrossRealmTrustSharedPassword) {
		nw.CrossRealmTrustSharedPassword = des.CrossRealmTrustSharedPassword
	}
	if dcl.NameToSelfLink(des.KdcDbKey, nw.KdcDbKey) || dcl.IsZeroValue(des.KdcDbKey) {
		nw.KdcDbKey = des.KdcDbKey
	}

	return nw
}

func canonicalizeNewClusterClusterConfigSecurityConfigKerberosConfigSet(c *Client, des, nw []ClusterClusterConfigSecurityConfigKerberosConfig) []ClusterClusterConfigSecurityConfigKerberosConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigSecurityConfigKerberosConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigSecurityConfigKerberosConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigLifecycleConfig(des, initial *ClusterClusterConfigLifecycleConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigLifecycleConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IdleDeleteTtl) {
		des.IdleDeleteTtl = initial.IdleDeleteTtl
	}
	if dcl.IsZeroValue(des.AutoDeleteTime) {
		des.AutoDeleteTime = initial.AutoDeleteTime
	}
	if dcl.IsZeroValue(des.AutoDeleteTtl) {
		des.AutoDeleteTtl = initial.AutoDeleteTtl
	}
	if dcl.IsZeroValue(des.IdleStartTime) {
		des.IdleStartTime = initial.IdleStartTime
	}

	return des
}

func canonicalizeNewClusterClusterConfigLifecycleConfig(c *Client, des, nw *ClusterClusterConfigLifecycleConfig) *ClusterClusterConfigLifecycleConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterClusterConfigLifecycleConfigSet(c *Client, des, nw []ClusterClusterConfigLifecycleConfig) []ClusterClusterConfigLifecycleConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigLifecycleConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigLifecycleConfig(c, &d, &n) {
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

func canonicalizeClusterClusterConfigEndpointConfig(des, initial *ClusterClusterConfigEndpointConfig, opts ...dcl.ApplyOption) *ClusterClusterConfigEndpointConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpPorts) {
		des.HttpPorts = initial.HttpPorts
	}
	if dcl.IsZeroValue(des.EnableHttpPortAccess) {
		des.EnableHttpPortAccess = initial.EnableHttpPortAccess
	}

	return des
}

func canonicalizeNewClusterClusterConfigEndpointConfig(c *Client, des, nw *ClusterClusterConfigEndpointConfig) *ClusterClusterConfigEndpointConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterClusterConfigEndpointConfigSet(c *Client, des, nw []ClusterClusterConfigEndpointConfig) []ClusterClusterConfigEndpointConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterClusterConfigEndpointConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterClusterConfigEndpointConfig(c, &d, &n) {
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

func canonicalizeClusterStatus(des, initial *ClusterStatus, opts ...dcl.ApplyOption) *ClusterStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.Detail) {
		des.Detail = initial.Detail
	}
	if dcl.IsZeroValue(des.StateStartTime) {
		des.StateStartTime = initial.StateStartTime
	}
	if dcl.IsZeroValue(des.Substate) {
		des.Substate = initial.Substate
	}

	return des
}

func canonicalizeNewClusterStatus(c *Client, des, nw *ClusterStatus) *ClusterStatus {
	if des == nil || nw == nil {
		return nw
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
			if !compareClusterStatus(c, &d, &n) {
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

func canonicalizeClusterStatusHistory(des, initial *ClusterStatusHistory, opts ...dcl.ApplyOption) *ClusterStatusHistory {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.Detail) {
		des.Detail = initial.Detail
	}
	if dcl.IsZeroValue(des.StateStartTime) {
		des.StateStartTime = initial.StateStartTime
	}
	if dcl.IsZeroValue(des.Substate) {
		des.Substate = initial.Substate
	}

	return des
}

func canonicalizeNewClusterStatusHistory(c *Client, des, nw *ClusterStatusHistory) *ClusterStatusHistory {
	if des == nil || nw == nil {
		return nw
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
			if !compareClusterStatusHistory(c, &d, &n) {
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

func canonicalizeClusterMetrics(des, initial *ClusterMetrics, opts ...dcl.ApplyOption) *ClusterMetrics {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HdfsMetrics) {
		des.HdfsMetrics = initial.HdfsMetrics
	}
	if dcl.IsZeroValue(des.YarnMetrics) {
		des.YarnMetrics = initial.YarnMetrics
	}

	return des
}

func canonicalizeNewClusterMetrics(c *Client, des, nw *ClusterMetrics) *ClusterMetrics {
	if des == nil || nw == nil {
		return nw
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
			if !compareClusterMetrics(c, &d, &n) {
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

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
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
func diffCluster(c *Client, desired, actual *Cluster, opts ...dcl.ApplyOption) ([]clusterDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []clusterDiff
	if !dcl.IsZeroValue(desired.Project) && (dcl.IsZeroValue(actual.Project) || !reflect.DeepEqual(*desired.Project, *actual.Project)) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %v\nACTUAL: %v", desired.Project, actual.Project)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareClusterClusterConfig(c, desired.Config, actual.Config) {
		c.Config.Logger.Infof("Detected diff in Config.\nDESIRED: %v\nACTUAL: %v", desired.Config, actual.Config)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Config",
		})
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateClusterOperation{},
			FieldName: "Labels",
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
	var deduped []clusterDiff
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
func compareClusterClusterConfigSlice(c *Client, desired, actual []ClusterClusterConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfig(c *Client, desired, actual *ClusterClusterConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StagingBucket == nil && desired.StagingBucket != nil && !dcl.IsEmptyValueIndirect(desired.StagingBucket) {
		c.Config.Logger.Infof("desired StagingBucket %s - but actually nil", dcl.SprintResource(desired.StagingBucket))
		return true
	}
	if !dcl.NameToSelfLink(desired.StagingBucket, actual.StagingBucket) && !dcl.IsZeroValue(desired.StagingBucket) {
		c.Config.Logger.Infof("Diff in StagingBucket. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StagingBucket), dcl.SprintResource(actual.StagingBucket))
		return true
	}
	if actual.TempBucket == nil && desired.TempBucket != nil && !dcl.IsEmptyValueIndirect(desired.TempBucket) {
		c.Config.Logger.Infof("desired TempBucket %s - but actually nil", dcl.SprintResource(desired.TempBucket))
		return true
	}
	if !dcl.NameToSelfLink(desired.TempBucket, actual.TempBucket) && !dcl.IsZeroValue(desired.TempBucket) {
		c.Config.Logger.Infof("Diff in TempBucket. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TempBucket), dcl.SprintResource(actual.TempBucket))
		return true
	}
	if actual.GceClusterConfig == nil && desired.GceClusterConfig != nil && !dcl.IsEmptyValueIndirect(desired.GceClusterConfig) {
		c.Config.Logger.Infof("desired GceClusterConfig %s - but actually nil", dcl.SprintResource(desired.GceClusterConfig))
		return true
	}
	if compareClusterClusterConfigGceClusterConfig(c, desired.GceClusterConfig, actual.GceClusterConfig) && !dcl.IsZeroValue(desired.GceClusterConfig) {
		c.Config.Logger.Infof("Diff in GceClusterConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GceClusterConfig), dcl.SprintResource(actual.GceClusterConfig))
		return true
	}
	if actual.MasterConfig == nil && desired.MasterConfig != nil && !dcl.IsEmptyValueIndirect(desired.MasterConfig) {
		c.Config.Logger.Infof("desired MasterConfig %s - but actually nil", dcl.SprintResource(desired.MasterConfig))
		return true
	}
	if compareClusterInstanceGroupConfig(c, desired.MasterConfig, actual.MasterConfig) && !dcl.IsZeroValue(desired.MasterConfig) {
		c.Config.Logger.Infof("Diff in MasterConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterConfig), dcl.SprintResource(actual.MasterConfig))
		return true
	}
	if actual.WorkerConfig == nil && desired.WorkerConfig != nil && !dcl.IsEmptyValueIndirect(desired.WorkerConfig) {
		c.Config.Logger.Infof("desired WorkerConfig %s - but actually nil", dcl.SprintResource(desired.WorkerConfig))
		return true
	}
	if compareClusterInstanceGroupConfig(c, desired.WorkerConfig, actual.WorkerConfig) && !dcl.IsZeroValue(desired.WorkerConfig) {
		c.Config.Logger.Infof("Diff in WorkerConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkerConfig), dcl.SprintResource(actual.WorkerConfig))
		return true
	}
	if actual.SecondaryWorkerConfig == nil && desired.SecondaryWorkerConfig != nil && !dcl.IsEmptyValueIndirect(desired.SecondaryWorkerConfig) {
		c.Config.Logger.Infof("desired SecondaryWorkerConfig %s - but actually nil", dcl.SprintResource(desired.SecondaryWorkerConfig))
		return true
	}
	if compareClusterInstanceGroupConfig(c, desired.SecondaryWorkerConfig, actual.SecondaryWorkerConfig) && !dcl.IsZeroValue(desired.SecondaryWorkerConfig) {
		c.Config.Logger.Infof("Diff in SecondaryWorkerConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecondaryWorkerConfig), dcl.SprintResource(actual.SecondaryWorkerConfig))
		return true
	}
	if actual.SoftwareConfig == nil && desired.SoftwareConfig != nil && !dcl.IsEmptyValueIndirect(desired.SoftwareConfig) {
		c.Config.Logger.Infof("desired SoftwareConfig %s - but actually nil", dcl.SprintResource(desired.SoftwareConfig))
		return true
	}
	if compareClusterClusterConfigSoftwareConfig(c, desired.SoftwareConfig, actual.SoftwareConfig) && !dcl.IsZeroValue(desired.SoftwareConfig) {
		c.Config.Logger.Infof("Diff in SoftwareConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SoftwareConfig), dcl.SprintResource(actual.SoftwareConfig))
		return true
	}
	if actual.InitializationActions == nil && desired.InitializationActions != nil && !dcl.IsEmptyValueIndirect(desired.InitializationActions) {
		c.Config.Logger.Infof("desired InitializationActions %s - but actually nil", dcl.SprintResource(desired.InitializationActions))
		return true
	}
	if compareClusterClusterConfigInitializationActionsSlice(c, desired.InitializationActions, actual.InitializationActions) && !dcl.IsZeroValue(desired.InitializationActions) {
		c.Config.Logger.Infof("Diff in InitializationActions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitializationActions), dcl.SprintResource(actual.InitializationActions))
		return true
	}
	if actual.EncryptionConfig == nil && desired.EncryptionConfig != nil && !dcl.IsEmptyValueIndirect(desired.EncryptionConfig) {
		c.Config.Logger.Infof("desired EncryptionConfig %s - but actually nil", dcl.SprintResource(desired.EncryptionConfig))
		return true
	}
	if compareClusterClusterConfigEncryptionConfig(c, desired.EncryptionConfig, actual.EncryptionConfig) && !dcl.IsZeroValue(desired.EncryptionConfig) {
		c.Config.Logger.Infof("Diff in EncryptionConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EncryptionConfig), dcl.SprintResource(actual.EncryptionConfig))
		return true
	}
	if actual.AutoscalingConfig == nil && desired.AutoscalingConfig != nil && !dcl.IsEmptyValueIndirect(desired.AutoscalingConfig) {
		c.Config.Logger.Infof("desired AutoscalingConfig %s - but actually nil", dcl.SprintResource(desired.AutoscalingConfig))
		return true
	}
	if compareClusterClusterConfigAutoscalingConfig(c, desired.AutoscalingConfig, actual.AutoscalingConfig) && !dcl.IsZeroValue(desired.AutoscalingConfig) {
		c.Config.Logger.Infof("Diff in AutoscalingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoscalingConfig), dcl.SprintResource(actual.AutoscalingConfig))
		return true
	}
	if actual.SecurityConfig == nil && desired.SecurityConfig != nil && !dcl.IsEmptyValueIndirect(desired.SecurityConfig) {
		c.Config.Logger.Infof("desired SecurityConfig %s - but actually nil", dcl.SprintResource(desired.SecurityConfig))
		return true
	}
	if compareClusterClusterConfigSecurityConfig(c, desired.SecurityConfig, actual.SecurityConfig) && !dcl.IsZeroValue(desired.SecurityConfig) {
		c.Config.Logger.Infof("Diff in SecurityConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecurityConfig), dcl.SprintResource(actual.SecurityConfig))
		return true
	}
	if actual.LifecycleConfig == nil && desired.LifecycleConfig != nil && !dcl.IsEmptyValueIndirect(desired.LifecycleConfig) {
		c.Config.Logger.Infof("desired LifecycleConfig %s - but actually nil", dcl.SprintResource(desired.LifecycleConfig))
		return true
	}
	if compareClusterClusterConfigLifecycleConfig(c, desired.LifecycleConfig, actual.LifecycleConfig) && !dcl.IsZeroValue(desired.LifecycleConfig) {
		c.Config.Logger.Infof("Diff in LifecycleConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LifecycleConfig), dcl.SprintResource(actual.LifecycleConfig))
		return true
	}
	if actual.EndpointConfig == nil && desired.EndpointConfig != nil && !dcl.IsEmptyValueIndirect(desired.EndpointConfig) {
		c.Config.Logger.Infof("desired EndpointConfig %s - but actually nil", dcl.SprintResource(desired.EndpointConfig))
		return true
	}
	if compareClusterClusterConfigEndpointConfig(c, desired.EndpointConfig, actual.EndpointConfig) && !dcl.IsZeroValue(desired.EndpointConfig) {
		c.Config.Logger.Infof("Diff in EndpointConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndpointConfig), dcl.SprintResource(actual.EndpointConfig))
		return true
	}
	return false
}
func compareClusterClusterConfigGceClusterConfigSlice(c *Client, desired, actual []ClusterClusterConfigGceClusterConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigGceClusterConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigGceClusterConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigGceClusterConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigGceClusterConfig(c *Client, desired, actual *ClusterClusterConfigGceClusterConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Zone == nil && desired.Zone != nil && !dcl.IsEmptyValueIndirect(desired.Zone) {
		c.Config.Logger.Infof("desired Zone %s - but actually nil", dcl.SprintResource(desired.Zone))
		return true
	}
	if !dcl.NameToSelfLink(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) {
		c.Config.Logger.Infof("Diff in Zone. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
		return true
	}
	if actual.Network == nil && desired.Network != nil && !dcl.IsEmptyValueIndirect(desired.Network) {
		c.Config.Logger.Infof("desired Network %s - but actually nil", dcl.SprintResource(desired.Network))
		return true
	}
	if !dcl.NameToSelfLink(desired.Network, actual.Network) && !dcl.IsZeroValue(desired.Network) {
		c.Config.Logger.Infof("Diff in Network. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Network), dcl.SprintResource(actual.Network))
		return true
	}
	if actual.Subnetwork == nil && desired.Subnetwork != nil && !dcl.IsEmptyValueIndirect(desired.Subnetwork) {
		c.Config.Logger.Infof("desired Subnetwork %s - but actually nil", dcl.SprintResource(desired.Subnetwork))
		return true
	}
	if !dcl.NameToSelfLink(desired.Subnetwork, actual.Subnetwork) && !dcl.IsZeroValue(desired.Subnetwork) {
		c.Config.Logger.Infof("Diff in Subnetwork. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Subnetwork), dcl.SprintResource(actual.Subnetwork))
		return true
	}
	if actual.InternalIPOnly == nil && desired.InternalIPOnly != nil && !dcl.IsEmptyValueIndirect(desired.InternalIPOnly) {
		c.Config.Logger.Infof("desired InternalIPOnly %s - but actually nil", dcl.SprintResource(desired.InternalIPOnly))
		return true
	}
	if !reflect.DeepEqual(desired.InternalIPOnly, actual.InternalIPOnly) && !dcl.IsZeroValue(desired.InternalIPOnly) && !(dcl.IsEmptyValueIndirect(desired.InternalIPOnly) && dcl.IsZeroValue(actual.InternalIPOnly)) {
		c.Config.Logger.Infof("Diff in InternalIPOnly. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InternalIPOnly), dcl.SprintResource(actual.InternalIPOnly))
		return true
	}
	if actual.PrivateIPv6GoogleAccess == nil && desired.PrivateIPv6GoogleAccess != nil && !dcl.IsEmptyValueIndirect(desired.PrivateIPv6GoogleAccess) {
		c.Config.Logger.Infof("desired PrivateIPv6GoogleAccess %s - but actually nil", dcl.SprintResource(desired.PrivateIPv6GoogleAccess))
		return true
	}
	if !reflect.DeepEqual(desired.PrivateIPv6GoogleAccess, actual.PrivateIPv6GoogleAccess) && !dcl.IsZeroValue(desired.PrivateIPv6GoogleAccess) && !(dcl.IsEmptyValueIndirect(desired.PrivateIPv6GoogleAccess) && dcl.IsZeroValue(actual.PrivateIPv6GoogleAccess)) {
		c.Config.Logger.Infof("Diff in PrivateIPv6GoogleAccess. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivateIPv6GoogleAccess), dcl.SprintResource(actual.PrivateIPv6GoogleAccess))
		return true
	}
	if actual.ServiceAccount == nil && desired.ServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccount) {
		c.Config.Logger.Infof("desired ServiceAccount %s - but actually nil", dcl.SprintResource(desired.ServiceAccount))
		return true
	}
	if !dcl.NameToSelfLink(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) {
		c.Config.Logger.Infof("Diff in ServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if actual.ServiceAccountScopes == nil && desired.ServiceAccountScopes != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccountScopes) {
		c.Config.Logger.Infof("desired ServiceAccountScopes %s - but actually nil", dcl.SprintResource(desired.ServiceAccountScopes))
		return true
	}
	if !dcl.SliceEquals(desired.ServiceAccountScopes, actual.ServiceAccountScopes) && !dcl.IsZeroValue(desired.ServiceAccountScopes) {
		c.Config.Logger.Infof("Diff in ServiceAccountScopes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountScopes), dcl.SprintResource(actual.ServiceAccountScopes))
		return true
	}
	if actual.Tags == nil && desired.Tags != nil && !dcl.IsEmptyValueIndirect(desired.Tags) {
		c.Config.Logger.Infof("desired Tags %s - but actually nil", dcl.SprintResource(desired.Tags))
		return true
	}
	if toAdd, toRemove := dcl.CompareStringSets(desired.Tags, actual.Tags); len(toAdd)+len(toRemove) > 0 {
		c.Config.Logger.Infof("Diff in Tags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	if actual.Metadata == nil && desired.Metadata != nil && !dcl.IsEmptyValueIndirect(desired.Metadata) {
		c.Config.Logger.Infof("desired Metadata %s - but actually nil", dcl.SprintResource(desired.Metadata))
		return true
	}
	if !reflect.DeepEqual(desired.Metadata, actual.Metadata) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if actual.ReservationAffinity == nil && desired.ReservationAffinity != nil && !dcl.IsEmptyValueIndirect(desired.ReservationAffinity) {
		c.Config.Logger.Infof("desired ReservationAffinity %s - but actually nil", dcl.SprintResource(desired.ReservationAffinity))
		return true
	}
	if compareClusterClusterConfigGceClusterConfigReservationAffinity(c, desired.ReservationAffinity, actual.ReservationAffinity) && !dcl.IsZeroValue(desired.ReservationAffinity) {
		c.Config.Logger.Infof("Diff in ReservationAffinity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReservationAffinity), dcl.SprintResource(actual.ReservationAffinity))
		return true
	}
	if actual.NodeGroupAffinity == nil && desired.NodeGroupAffinity != nil && !dcl.IsEmptyValueIndirect(desired.NodeGroupAffinity) {
		c.Config.Logger.Infof("desired NodeGroupAffinity %s - but actually nil", dcl.SprintResource(desired.NodeGroupAffinity))
		return true
	}
	if compareClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, desired.NodeGroupAffinity, actual.NodeGroupAffinity) && !dcl.IsZeroValue(desired.NodeGroupAffinity) {
		c.Config.Logger.Infof("Diff in NodeGroupAffinity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeGroupAffinity), dcl.SprintResource(actual.NodeGroupAffinity))
		return true
	}
	return false
}
func compareClusterClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, desired, actual []ClusterClusterConfigGceClusterConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigGceClusterConfigReservationAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigGceClusterConfigReservationAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigGceClusterConfigReservationAffinity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigGceClusterConfigReservationAffinity(c *Client, desired, actual *ClusterClusterConfigGceClusterConfigReservationAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ConsumeReservationType == nil && desired.ConsumeReservationType != nil && !dcl.IsEmptyValueIndirect(desired.ConsumeReservationType) {
		c.Config.Logger.Infof("desired ConsumeReservationType %s - but actually nil", dcl.SprintResource(desired.ConsumeReservationType))
		return true
	}
	if !reflect.DeepEqual(desired.ConsumeReservationType, actual.ConsumeReservationType) && !dcl.IsZeroValue(desired.ConsumeReservationType) && !(dcl.IsEmptyValueIndirect(desired.ConsumeReservationType) && dcl.IsZeroValue(actual.ConsumeReservationType)) {
		c.Config.Logger.Infof("Diff in ConsumeReservationType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumeReservationType), dcl.SprintResource(actual.ConsumeReservationType))
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Values == nil && desired.Values != nil && !dcl.IsEmptyValueIndirect(desired.Values) {
		c.Config.Logger.Infof("desired Values %s - but actually nil", dcl.SprintResource(desired.Values))
		return true
	}
	if !dcl.SliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}
func compareClusterClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, desired, actual []ClusterClusterConfigGceClusterConfigNodeGroupAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigGceClusterConfigNodeGroupAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigGceClusterConfigNodeGroupAffinity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, desired, actual *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NodeGroup == nil && desired.NodeGroup != nil && !dcl.IsEmptyValueIndirect(desired.NodeGroup) {
		c.Config.Logger.Infof("desired NodeGroup %s - but actually nil", dcl.SprintResource(desired.NodeGroup))
		return true
	}
	if !dcl.NameToSelfLink(desired.NodeGroup, actual.NodeGroup) && !dcl.IsZeroValue(desired.NodeGroup) {
		c.Config.Logger.Infof("Diff in NodeGroup. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeGroup), dcl.SprintResource(actual.NodeGroup))
		return true
	}
	return false
}
func compareClusterInstanceGroupConfigSlice(c *Client, desired, actual []ClusterInstanceGroupConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterInstanceGroupConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterInstanceGroupConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterInstanceGroupConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterInstanceGroupConfig(c *Client, desired, actual *ClusterInstanceGroupConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumInstances == nil && desired.NumInstances != nil && !dcl.IsEmptyValueIndirect(desired.NumInstances) {
		c.Config.Logger.Infof("desired NumInstances %s - but actually nil", dcl.SprintResource(desired.NumInstances))
		return true
	}
	if !reflect.DeepEqual(desired.NumInstances, actual.NumInstances) && !dcl.IsZeroValue(desired.NumInstances) && !(dcl.IsEmptyValueIndirect(desired.NumInstances) && dcl.IsZeroValue(actual.NumInstances)) {
		c.Config.Logger.Infof("Diff in NumInstances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumInstances), dcl.SprintResource(actual.NumInstances))
		return true
	}
	if actual.Image == nil && desired.Image != nil && !dcl.IsEmptyValueIndirect(desired.Image) {
		c.Config.Logger.Infof("desired Image %s - but actually nil", dcl.SprintResource(desired.Image))
		return true
	}
	if !dcl.NameToSelfLink(desired.Image, actual.Image) && !dcl.IsZeroValue(desired.Image) {
		c.Config.Logger.Infof("Diff in Image. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Image), dcl.SprintResource(actual.Image))
		return true
	}
	if actual.MachineType == nil && desired.MachineType != nil && !dcl.IsEmptyValueIndirect(desired.MachineType) {
		c.Config.Logger.Infof("desired MachineType %s - but actually nil", dcl.SprintResource(desired.MachineType))
		return true
	}
	if !dcl.NameToSelfLink(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) {
		c.Config.Logger.Infof("Diff in MachineType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if actual.DiskConfig == nil && desired.DiskConfig != nil && !dcl.IsEmptyValueIndirect(desired.DiskConfig) {
		c.Config.Logger.Infof("desired DiskConfig %s - but actually nil", dcl.SprintResource(desired.DiskConfig))
		return true
	}
	if compareClusterInstanceGroupConfigDiskConfig(c, desired.DiskConfig, actual.DiskConfig) && !dcl.IsZeroValue(desired.DiskConfig) {
		c.Config.Logger.Infof("Diff in DiskConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskConfig), dcl.SprintResource(actual.DiskConfig))
		return true
	}
	if actual.Preemptibility == nil && desired.Preemptibility != nil && !dcl.IsEmptyValueIndirect(desired.Preemptibility) {
		c.Config.Logger.Infof("desired Preemptibility %s - but actually nil", dcl.SprintResource(desired.Preemptibility))
		return true
	}
	if !reflect.DeepEqual(desired.Preemptibility, actual.Preemptibility) && !dcl.IsZeroValue(desired.Preemptibility) && !(dcl.IsEmptyValueIndirect(desired.Preemptibility) && dcl.IsZeroValue(actual.Preemptibility)) {
		c.Config.Logger.Infof("Diff in Preemptibility. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptibility), dcl.SprintResource(actual.Preemptibility))
		return true
	}
	if actual.Accelerators == nil && desired.Accelerators != nil && !dcl.IsEmptyValueIndirect(desired.Accelerators) {
		c.Config.Logger.Infof("desired Accelerators %s - but actually nil", dcl.SprintResource(desired.Accelerators))
		return true
	}
	if compareClusterInstanceGroupConfigAcceleratorsSlice(c, desired.Accelerators, actual.Accelerators) && !dcl.IsZeroValue(desired.Accelerators) {
		c.Config.Logger.Infof("Diff in Accelerators. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Accelerators), dcl.SprintResource(actual.Accelerators))
		return true
	}
	if actual.MinCpuPlatform == nil && desired.MinCpuPlatform != nil && !dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("desired MinCpuPlatform %s - but actually nil", dcl.SprintResource(desired.MinCpuPlatform))
		return true
	}
	if !reflect.DeepEqual(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) && !(dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) && dcl.IsZeroValue(actual.MinCpuPlatform)) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	return false
}
func compareClusterInstanceGroupConfigDiskConfigSlice(c *Client, desired, actual []ClusterInstanceGroupConfigDiskConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterInstanceGroupConfigDiskConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterInstanceGroupConfigDiskConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterInstanceGroupConfigDiskConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterInstanceGroupConfigDiskConfig(c *Client, desired, actual *ClusterInstanceGroupConfigDiskConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BootDiskType == nil && desired.BootDiskType != nil && !dcl.IsEmptyValueIndirect(desired.BootDiskType) {
		c.Config.Logger.Infof("desired BootDiskType %s - but actually nil", dcl.SprintResource(desired.BootDiskType))
		return true
	}
	if !reflect.DeepEqual(desired.BootDiskType, actual.BootDiskType) && !dcl.IsZeroValue(desired.BootDiskType) && !(dcl.IsEmptyValueIndirect(desired.BootDiskType) && dcl.IsZeroValue(actual.BootDiskType)) {
		c.Config.Logger.Infof("Diff in BootDiskType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BootDiskType), dcl.SprintResource(actual.BootDiskType))
		return true
	}
	if actual.BootDiskSizeGb == nil && desired.BootDiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.BootDiskSizeGb) {
		c.Config.Logger.Infof("desired BootDiskSizeGb %s - but actually nil", dcl.SprintResource(desired.BootDiskSizeGb))
		return true
	}
	if !reflect.DeepEqual(desired.BootDiskSizeGb, actual.BootDiskSizeGb) && !dcl.IsZeroValue(desired.BootDiskSizeGb) && !(dcl.IsEmptyValueIndirect(desired.BootDiskSizeGb) && dcl.IsZeroValue(actual.BootDiskSizeGb)) {
		c.Config.Logger.Infof("Diff in BootDiskSizeGb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BootDiskSizeGb), dcl.SprintResource(actual.BootDiskSizeGb))
		return true
	}
	if actual.NumLocalSsds == nil && desired.NumLocalSsds != nil && !dcl.IsEmptyValueIndirect(desired.NumLocalSsds) {
		c.Config.Logger.Infof("desired NumLocalSsds %s - but actually nil", dcl.SprintResource(desired.NumLocalSsds))
		return true
	}
	if !reflect.DeepEqual(desired.NumLocalSsds, actual.NumLocalSsds) && !dcl.IsZeroValue(desired.NumLocalSsds) && !(dcl.IsEmptyValueIndirect(desired.NumLocalSsds) && dcl.IsZeroValue(actual.NumLocalSsds)) {
		c.Config.Logger.Infof("Diff in NumLocalSsds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumLocalSsds), dcl.SprintResource(actual.NumLocalSsds))
		return true
	}
	return false
}
func compareClusterInstanceGroupConfigManagedGroupConfigSlice(c *Client, desired, actual []ClusterInstanceGroupConfigManagedGroupConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterInstanceGroupConfigManagedGroupConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterInstanceGroupConfigManagedGroupConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterInstanceGroupConfigManagedGroupConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterInstanceGroupConfigManagedGroupConfig(c *Client, desired, actual *ClusterInstanceGroupConfigManagedGroupConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}
func compareClusterInstanceGroupConfigAcceleratorsSlice(c *Client, desired, actual []ClusterInstanceGroupConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterInstanceGroupConfigAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterInstanceGroupConfigAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterInstanceGroupConfigAccelerators, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterInstanceGroupConfigAccelerators(c *Client, desired, actual *ClusterInstanceGroupConfigAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AcceleratorType == nil && desired.AcceleratorType != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorType) {
		c.Config.Logger.Infof("desired AcceleratorType %s - but actually nil", dcl.SprintResource(desired.AcceleratorType))
		return true
	}
	if !dcl.NameToSelfLink(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) {
		c.Config.Logger.Infof("Diff in AcceleratorType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	if actual.AcceleratorCount == nil && desired.AcceleratorCount != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorCount) {
		c.Config.Logger.Infof("desired AcceleratorCount %s - but actually nil", dcl.SprintResource(desired.AcceleratorCount))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) && !(dcl.IsEmptyValueIndirect(desired.AcceleratorCount) && dcl.IsZeroValue(actual.AcceleratorCount)) {
		c.Config.Logger.Infof("Diff in AcceleratorCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	return false
}
func compareClusterClusterConfigSoftwareConfigSlice(c *Client, desired, actual []ClusterClusterConfigSoftwareConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigSoftwareConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigSoftwareConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigSoftwareConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigSoftwareConfig(c *Client, desired, actual *ClusterClusterConfigSoftwareConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ImageVersion == nil && desired.ImageVersion != nil && !dcl.IsEmptyValueIndirect(desired.ImageVersion) {
		c.Config.Logger.Infof("desired ImageVersion %s - but actually nil", dcl.SprintResource(desired.ImageVersion))
		return true
	}
	if !reflect.DeepEqual(desired.ImageVersion, actual.ImageVersion) && !dcl.IsZeroValue(desired.ImageVersion) && !(dcl.IsEmptyValueIndirect(desired.ImageVersion) && dcl.IsZeroValue(actual.ImageVersion)) {
		c.Config.Logger.Infof("Diff in ImageVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImageVersion), dcl.SprintResource(actual.ImageVersion))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !reflect.DeepEqual(desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.OptionalComponents == nil && desired.OptionalComponents != nil && !dcl.IsEmptyValueIndirect(desired.OptionalComponents) {
		c.Config.Logger.Infof("desired OptionalComponents %s - but actually nil", dcl.SprintResource(desired.OptionalComponents))
		return true
	}
	if compareClusterClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c, desired.OptionalComponents, actual.OptionalComponents) && !dcl.IsZeroValue(desired.OptionalComponents) {
		c.Config.Logger.Infof("Diff in OptionalComponents. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OptionalComponents), dcl.SprintResource(actual.OptionalComponents))
		return true
	}
	return false
}
func compareClusterClusterConfigInitializationActionsSlice(c *Client, desired, actual []ClusterClusterConfigInitializationActions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigInitializationActions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigInitializationActions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigInitializationActions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigInitializationActions(c *Client, desired, actual *ClusterClusterConfigInitializationActions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ExecutableFile == nil && desired.ExecutableFile != nil && !dcl.IsEmptyValueIndirect(desired.ExecutableFile) {
		c.Config.Logger.Infof("desired ExecutableFile %s - but actually nil", dcl.SprintResource(desired.ExecutableFile))
		return true
	}
	if !reflect.DeepEqual(desired.ExecutableFile, actual.ExecutableFile) && !dcl.IsZeroValue(desired.ExecutableFile) && !(dcl.IsEmptyValueIndirect(desired.ExecutableFile) && dcl.IsZeroValue(actual.ExecutableFile)) {
		c.Config.Logger.Infof("Diff in ExecutableFile. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExecutableFile), dcl.SprintResource(actual.ExecutableFile))
		return true
	}
	if actual.ExecutionTimeout == nil && desired.ExecutionTimeout != nil && !dcl.IsEmptyValueIndirect(desired.ExecutionTimeout) {
		c.Config.Logger.Infof("desired ExecutionTimeout %s - but actually nil", dcl.SprintResource(desired.ExecutionTimeout))
		return true
	}
	if !reflect.DeepEqual(desired.ExecutionTimeout, actual.ExecutionTimeout) && !dcl.IsZeroValue(desired.ExecutionTimeout) && !(dcl.IsEmptyValueIndirect(desired.ExecutionTimeout) && dcl.IsZeroValue(actual.ExecutionTimeout)) {
		c.Config.Logger.Infof("Diff in ExecutionTimeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExecutionTimeout), dcl.SprintResource(actual.ExecutionTimeout))
		return true
	}
	return false
}
func compareClusterClusterConfigEncryptionConfigSlice(c *Client, desired, actual []ClusterClusterConfigEncryptionConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigEncryptionConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigEncryptionConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigEncryptionConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigEncryptionConfig(c *Client, desired, actual *ClusterClusterConfigEncryptionConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.GcePdKmsKeyName == nil && desired.GcePdKmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.GcePdKmsKeyName) {
		c.Config.Logger.Infof("desired GcePdKmsKeyName %s - but actually nil", dcl.SprintResource(desired.GcePdKmsKeyName))
		return true
	}
	if !dcl.NameToSelfLink(desired.GcePdKmsKeyName, actual.GcePdKmsKeyName) && !dcl.IsZeroValue(desired.GcePdKmsKeyName) {
		c.Config.Logger.Infof("Diff in GcePdKmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GcePdKmsKeyName), dcl.SprintResource(actual.GcePdKmsKeyName))
		return true
	}
	return false
}
func compareClusterClusterConfigAutoscalingConfigSlice(c *Client, desired, actual []ClusterClusterConfigAutoscalingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigAutoscalingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigAutoscalingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigAutoscalingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigAutoscalingConfig(c *Client, desired, actual *ClusterClusterConfigAutoscalingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Policy == nil && desired.Policy != nil && !dcl.IsEmptyValueIndirect(desired.Policy) {
		c.Config.Logger.Infof("desired Policy %s - but actually nil", dcl.SprintResource(desired.Policy))
		return true
	}
	if !dcl.NameToSelfLink(desired.Policy, actual.Policy) && !dcl.IsZeroValue(desired.Policy) {
		c.Config.Logger.Infof("Diff in Policy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Policy), dcl.SprintResource(actual.Policy))
		return true
	}
	return false
}
func compareClusterClusterConfigSecurityConfigSlice(c *Client, desired, actual []ClusterClusterConfigSecurityConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigSecurityConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigSecurityConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigSecurityConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigSecurityConfig(c *Client, desired, actual *ClusterClusterConfigSecurityConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.KerberosConfig == nil && desired.KerberosConfig != nil && !dcl.IsEmptyValueIndirect(desired.KerberosConfig) {
		c.Config.Logger.Infof("desired KerberosConfig %s - but actually nil", dcl.SprintResource(desired.KerberosConfig))
		return true
	}
	if compareClusterClusterConfigSecurityConfigKerberosConfig(c, desired.KerberosConfig, actual.KerberosConfig) && !dcl.IsZeroValue(desired.KerberosConfig) {
		c.Config.Logger.Infof("Diff in KerberosConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KerberosConfig), dcl.SprintResource(actual.KerberosConfig))
		return true
	}
	return false
}
func compareClusterClusterConfigSecurityConfigKerberosConfigSlice(c *Client, desired, actual []ClusterClusterConfigSecurityConfigKerberosConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigSecurityConfigKerberosConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigSecurityConfigKerberosConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigSecurityConfigKerberosConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigSecurityConfigKerberosConfig(c *Client, desired, actual *ClusterClusterConfigSecurityConfigKerberosConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableKerberos == nil && desired.EnableKerberos != nil && !dcl.IsEmptyValueIndirect(desired.EnableKerberos) {
		c.Config.Logger.Infof("desired EnableKerberos %s - but actually nil", dcl.SprintResource(desired.EnableKerberos))
		return true
	}
	if !reflect.DeepEqual(desired.EnableKerberos, actual.EnableKerberos) && !dcl.IsZeroValue(desired.EnableKerberos) && !(dcl.IsEmptyValueIndirect(desired.EnableKerberos) && dcl.IsZeroValue(actual.EnableKerberos)) {
		c.Config.Logger.Infof("Diff in EnableKerberos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableKerberos), dcl.SprintResource(actual.EnableKerberos))
		return true
	}
	if actual.RootPrincipalPassword == nil && desired.RootPrincipalPassword != nil && !dcl.IsEmptyValueIndirect(desired.RootPrincipalPassword) {
		c.Config.Logger.Infof("desired RootPrincipalPassword %s - but actually nil", dcl.SprintResource(desired.RootPrincipalPassword))
		return true
	}
	if !dcl.NameToSelfLink(desired.RootPrincipalPassword, actual.RootPrincipalPassword) && !dcl.IsZeroValue(desired.RootPrincipalPassword) {
		c.Config.Logger.Infof("Diff in RootPrincipalPassword. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RootPrincipalPassword), dcl.SprintResource(actual.RootPrincipalPassword))
		return true
	}
	if actual.KmsKey == nil && desired.KmsKey != nil && !dcl.IsEmptyValueIndirect(desired.KmsKey) {
		c.Config.Logger.Infof("desired KmsKey %s - but actually nil", dcl.SprintResource(desired.KmsKey))
		return true
	}
	if !dcl.NameToSelfLink(desired.KmsKey, actual.KmsKey) && !dcl.IsZeroValue(desired.KmsKey) {
		c.Config.Logger.Infof("Diff in KmsKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKey), dcl.SprintResource(actual.KmsKey))
		return true
	}
	if actual.Keystore == nil && desired.Keystore != nil && !dcl.IsEmptyValueIndirect(desired.Keystore) {
		c.Config.Logger.Infof("desired Keystore %s - but actually nil", dcl.SprintResource(desired.Keystore))
		return true
	}
	if !dcl.NameToSelfLink(desired.Keystore, actual.Keystore) && !dcl.IsZeroValue(desired.Keystore) {
		c.Config.Logger.Infof("Diff in Keystore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Keystore), dcl.SprintResource(actual.Keystore))
		return true
	}
	if actual.Truststore == nil && desired.Truststore != nil && !dcl.IsEmptyValueIndirect(desired.Truststore) {
		c.Config.Logger.Infof("desired Truststore %s - but actually nil", dcl.SprintResource(desired.Truststore))
		return true
	}
	if !dcl.NameToSelfLink(desired.Truststore, actual.Truststore) && !dcl.IsZeroValue(desired.Truststore) {
		c.Config.Logger.Infof("Diff in Truststore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Truststore), dcl.SprintResource(actual.Truststore))
		return true
	}
	if actual.KeystorePassword == nil && desired.KeystorePassword != nil && !dcl.IsEmptyValueIndirect(desired.KeystorePassword) {
		c.Config.Logger.Infof("desired KeystorePassword %s - but actually nil", dcl.SprintResource(desired.KeystorePassword))
		return true
	}
	if !dcl.NameToSelfLink(desired.KeystorePassword, actual.KeystorePassword) && !dcl.IsZeroValue(desired.KeystorePassword) {
		c.Config.Logger.Infof("Diff in KeystorePassword. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeystorePassword), dcl.SprintResource(actual.KeystorePassword))
		return true
	}
	if actual.KeyPassword == nil && desired.KeyPassword != nil && !dcl.IsEmptyValueIndirect(desired.KeyPassword) {
		c.Config.Logger.Infof("desired KeyPassword %s - but actually nil", dcl.SprintResource(desired.KeyPassword))
		return true
	}
	if !dcl.NameToSelfLink(desired.KeyPassword, actual.KeyPassword) && !dcl.IsZeroValue(desired.KeyPassword) {
		c.Config.Logger.Infof("Diff in KeyPassword. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeyPassword), dcl.SprintResource(actual.KeyPassword))
		return true
	}
	if actual.TruststorePassword == nil && desired.TruststorePassword != nil && !dcl.IsEmptyValueIndirect(desired.TruststorePassword) {
		c.Config.Logger.Infof("desired TruststorePassword %s - but actually nil", dcl.SprintResource(desired.TruststorePassword))
		return true
	}
	if !dcl.NameToSelfLink(desired.TruststorePassword, actual.TruststorePassword) && !dcl.IsZeroValue(desired.TruststorePassword) {
		c.Config.Logger.Infof("Diff in TruststorePassword. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TruststorePassword), dcl.SprintResource(actual.TruststorePassword))
		return true
	}
	if actual.CrossRealmTrustRealm == nil && desired.CrossRealmTrustRealm != nil && !dcl.IsEmptyValueIndirect(desired.CrossRealmTrustRealm) {
		c.Config.Logger.Infof("desired CrossRealmTrustRealm %s - but actually nil", dcl.SprintResource(desired.CrossRealmTrustRealm))
		return true
	}
	if !reflect.DeepEqual(desired.CrossRealmTrustRealm, actual.CrossRealmTrustRealm) && !dcl.IsZeroValue(desired.CrossRealmTrustRealm) && !(dcl.IsEmptyValueIndirect(desired.CrossRealmTrustRealm) && dcl.IsZeroValue(actual.CrossRealmTrustRealm)) {
		c.Config.Logger.Infof("Diff in CrossRealmTrustRealm. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossRealmTrustRealm), dcl.SprintResource(actual.CrossRealmTrustRealm))
		return true
	}
	if actual.CrossRealmTrustKdc == nil && desired.CrossRealmTrustKdc != nil && !dcl.IsEmptyValueIndirect(desired.CrossRealmTrustKdc) {
		c.Config.Logger.Infof("desired CrossRealmTrustKdc %s - but actually nil", dcl.SprintResource(desired.CrossRealmTrustKdc))
		return true
	}
	if !reflect.DeepEqual(desired.CrossRealmTrustKdc, actual.CrossRealmTrustKdc) && !dcl.IsZeroValue(desired.CrossRealmTrustKdc) && !(dcl.IsEmptyValueIndirect(desired.CrossRealmTrustKdc) && dcl.IsZeroValue(actual.CrossRealmTrustKdc)) {
		c.Config.Logger.Infof("Diff in CrossRealmTrustKdc. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossRealmTrustKdc), dcl.SprintResource(actual.CrossRealmTrustKdc))
		return true
	}
	if actual.CrossRealmTrustAdminServer == nil && desired.CrossRealmTrustAdminServer != nil && !dcl.IsEmptyValueIndirect(desired.CrossRealmTrustAdminServer) {
		c.Config.Logger.Infof("desired CrossRealmTrustAdminServer %s - but actually nil", dcl.SprintResource(desired.CrossRealmTrustAdminServer))
		return true
	}
	if !reflect.DeepEqual(desired.CrossRealmTrustAdminServer, actual.CrossRealmTrustAdminServer) && !dcl.IsZeroValue(desired.CrossRealmTrustAdminServer) && !(dcl.IsEmptyValueIndirect(desired.CrossRealmTrustAdminServer) && dcl.IsZeroValue(actual.CrossRealmTrustAdminServer)) {
		c.Config.Logger.Infof("Diff in CrossRealmTrustAdminServer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossRealmTrustAdminServer), dcl.SprintResource(actual.CrossRealmTrustAdminServer))
		return true
	}
	if actual.CrossRealmTrustSharedPassword == nil && desired.CrossRealmTrustSharedPassword != nil && !dcl.IsEmptyValueIndirect(desired.CrossRealmTrustSharedPassword) {
		c.Config.Logger.Infof("desired CrossRealmTrustSharedPassword %s - but actually nil", dcl.SprintResource(desired.CrossRealmTrustSharedPassword))
		return true
	}
	if !dcl.NameToSelfLink(desired.CrossRealmTrustSharedPassword, actual.CrossRealmTrustSharedPassword) && !dcl.IsZeroValue(desired.CrossRealmTrustSharedPassword) {
		c.Config.Logger.Infof("Diff in CrossRealmTrustSharedPassword. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossRealmTrustSharedPassword), dcl.SprintResource(actual.CrossRealmTrustSharedPassword))
		return true
	}
	if actual.KdcDbKey == nil && desired.KdcDbKey != nil && !dcl.IsEmptyValueIndirect(desired.KdcDbKey) {
		c.Config.Logger.Infof("desired KdcDbKey %s - but actually nil", dcl.SprintResource(desired.KdcDbKey))
		return true
	}
	if !dcl.NameToSelfLink(desired.KdcDbKey, actual.KdcDbKey) && !dcl.IsZeroValue(desired.KdcDbKey) {
		c.Config.Logger.Infof("Diff in KdcDbKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KdcDbKey), dcl.SprintResource(actual.KdcDbKey))
		return true
	}
	if actual.TgtLifetimeHours == nil && desired.TgtLifetimeHours != nil && !dcl.IsEmptyValueIndirect(desired.TgtLifetimeHours) {
		c.Config.Logger.Infof("desired TgtLifetimeHours %s - but actually nil", dcl.SprintResource(desired.TgtLifetimeHours))
		return true
	}
	if !reflect.DeepEqual(desired.TgtLifetimeHours, actual.TgtLifetimeHours) && !dcl.IsZeroValue(desired.TgtLifetimeHours) && !(dcl.IsEmptyValueIndirect(desired.TgtLifetimeHours) && dcl.IsZeroValue(actual.TgtLifetimeHours)) {
		c.Config.Logger.Infof("Diff in TgtLifetimeHours. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TgtLifetimeHours), dcl.SprintResource(actual.TgtLifetimeHours))
		return true
	}
	if actual.Realm == nil && desired.Realm != nil && !dcl.IsEmptyValueIndirect(desired.Realm) {
		c.Config.Logger.Infof("desired Realm %s - but actually nil", dcl.SprintResource(desired.Realm))
		return true
	}
	if !reflect.DeepEqual(desired.Realm, actual.Realm) && !dcl.IsZeroValue(desired.Realm) && !(dcl.IsEmptyValueIndirect(desired.Realm) && dcl.IsZeroValue(actual.Realm)) {
		c.Config.Logger.Infof("Diff in Realm. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Realm), dcl.SprintResource(actual.Realm))
		return true
	}
	return false
}
func compareClusterClusterConfigLifecycleConfigSlice(c *Client, desired, actual []ClusterClusterConfigLifecycleConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigLifecycleConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigLifecycleConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigLifecycleConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigLifecycleConfig(c *Client, desired, actual *ClusterClusterConfigLifecycleConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IdleDeleteTtl == nil && desired.IdleDeleteTtl != nil && !dcl.IsEmptyValueIndirect(desired.IdleDeleteTtl) {
		c.Config.Logger.Infof("desired IdleDeleteTtl %s - but actually nil", dcl.SprintResource(desired.IdleDeleteTtl))
		return true
	}
	if !reflect.DeepEqual(desired.IdleDeleteTtl, actual.IdleDeleteTtl) && !dcl.IsZeroValue(desired.IdleDeleteTtl) && !(dcl.IsEmptyValueIndirect(desired.IdleDeleteTtl) && dcl.IsZeroValue(actual.IdleDeleteTtl)) {
		c.Config.Logger.Infof("Diff in IdleDeleteTtl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IdleDeleteTtl), dcl.SprintResource(actual.IdleDeleteTtl))
		return true
	}
	if actual.AutoDeleteTime == nil && desired.AutoDeleteTime != nil && !dcl.IsEmptyValueIndirect(desired.AutoDeleteTime) {
		c.Config.Logger.Infof("desired AutoDeleteTime %s - but actually nil", dcl.SprintResource(desired.AutoDeleteTime))
		return true
	}
	if !reflect.DeepEqual(desired.AutoDeleteTime, actual.AutoDeleteTime) && !dcl.IsZeroValue(desired.AutoDeleteTime) && !(dcl.IsEmptyValueIndirect(desired.AutoDeleteTime) && dcl.IsZeroValue(actual.AutoDeleteTime)) {
		c.Config.Logger.Infof("Diff in AutoDeleteTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoDeleteTime), dcl.SprintResource(actual.AutoDeleteTime))
		return true
	}
	if actual.AutoDeleteTtl == nil && desired.AutoDeleteTtl != nil && !dcl.IsEmptyValueIndirect(desired.AutoDeleteTtl) {
		c.Config.Logger.Infof("desired AutoDeleteTtl %s - but actually nil", dcl.SprintResource(desired.AutoDeleteTtl))
		return true
	}
	if !reflect.DeepEqual(desired.AutoDeleteTtl, actual.AutoDeleteTtl) && !dcl.IsZeroValue(desired.AutoDeleteTtl) && !(dcl.IsEmptyValueIndirect(desired.AutoDeleteTtl) && dcl.IsZeroValue(actual.AutoDeleteTtl)) {
		c.Config.Logger.Infof("Diff in AutoDeleteTtl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoDeleteTtl), dcl.SprintResource(actual.AutoDeleteTtl))
		return true
	}
	return false
}
func compareClusterClusterConfigEndpointConfigSlice(c *Client, desired, actual []ClusterClusterConfigEndpointConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigEndpointConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigEndpointConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigEndpointConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigEndpointConfig(c *Client, desired, actual *ClusterClusterConfigEndpointConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableHttpPortAccess == nil && desired.EnableHttpPortAccess != nil && !dcl.IsEmptyValueIndirect(desired.EnableHttpPortAccess) {
		c.Config.Logger.Infof("desired EnableHttpPortAccess %s - but actually nil", dcl.SprintResource(desired.EnableHttpPortAccess))
		return true
	}
	if !reflect.DeepEqual(desired.EnableHttpPortAccess, actual.EnableHttpPortAccess) && !dcl.IsZeroValue(desired.EnableHttpPortAccess) && !(dcl.IsEmptyValueIndirect(desired.EnableHttpPortAccess) && dcl.IsZeroValue(actual.EnableHttpPortAccess)) {
		c.Config.Logger.Infof("Diff in EnableHttpPortAccess. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableHttpPortAccess), dcl.SprintResource(actual.EnableHttpPortAccess))
		return true
	}
	return false
}
func compareClusterStatusSlice(c *Client, desired, actual []ClusterStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatus, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatus(c *Client, desired, actual *ClusterStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}
func compareClusterStatusHistorySlice(c *Client, desired, actual []ClusterStatusHistory) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatusHistory, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatusHistory(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatusHistory, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatusHistory(c *Client, desired, actual *ClusterStatusHistory) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}
func compareClusterMetricsSlice(c *Client, desired, actual []ClusterMetrics) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMetrics, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMetrics(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMetrics, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMetrics(c *Client, desired, actual *ClusterMetrics) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HdfsMetrics == nil && desired.HdfsMetrics != nil && !dcl.IsEmptyValueIndirect(desired.HdfsMetrics) {
		c.Config.Logger.Infof("desired HdfsMetrics %s - but actually nil", dcl.SprintResource(desired.HdfsMetrics))
		return true
	}
	if !reflect.DeepEqual(desired.HdfsMetrics, actual.HdfsMetrics) && !dcl.IsZeroValue(desired.HdfsMetrics) {
		c.Config.Logger.Infof("Diff in HdfsMetrics. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HdfsMetrics), dcl.SprintResource(actual.HdfsMetrics))
		return true
	}
	if actual.YarnMetrics == nil && desired.YarnMetrics != nil && !dcl.IsEmptyValueIndirect(desired.YarnMetrics) {
		c.Config.Logger.Infof("desired YarnMetrics %s - but actually nil", dcl.SprintResource(desired.YarnMetrics))
		return true
	}
	if !reflect.DeepEqual(desired.YarnMetrics, actual.YarnMetrics) && !dcl.IsZeroValue(desired.YarnMetrics) {
		c.Config.Logger.Infof("Diff in YarnMetrics. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.YarnMetrics), dcl.SprintResource(actual.YarnMetrics))
		return true
	}
	return false
}
func compareClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumSlice(c *Client, desired, actual []ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(c *Client, desired, actual *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, desired, actual []ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(c *Client, desired, actual *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterInstanceGroupConfigPreemptibilityEnumSlice(c *Client, desired, actual []ClusterInstanceGroupConfigPreemptibilityEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterInstanceGroupConfigPreemptibilityEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterInstanceGroupConfigPreemptibilityEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterInstanceGroupConfigPreemptibilityEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterInstanceGroupConfigPreemptibilityEnum(c *Client, desired, actual *ClusterInstanceGroupConfigPreemptibilityEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c *Client, desired, actual []ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterClusterConfigSoftwareConfigOptionalComponentsEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterClusterConfigSoftwareConfigOptionalComponentsEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterClusterConfigSoftwareConfigOptionalComponentsEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterClusterConfigSoftwareConfigOptionalComponentsEnum(c *Client, desired, actual *ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterStatusStateEnumSlice(c *Client, desired, actual []ClusterStatusStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatusStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatusStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatusStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatusStateEnum(c *Client, desired, actual *ClusterStatusStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterStatusSubstateEnumSlice(c *Client, desired, actual []ClusterStatusSubstateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatusSubstateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatusSubstateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatusSubstateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatusSubstateEnum(c *Client, desired, actual *ClusterStatusSubstateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterStatusHistoryStateEnumSlice(c *Client, desired, actual []ClusterStatusHistoryStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatusHistoryStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatusHistoryStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatusHistoryStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatusHistoryStateEnum(c *Client, desired, actual *ClusterStatusHistoryStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterStatusHistorySubstateEnumSlice(c *Client, desired, actual []ClusterStatusHistorySubstateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterStatusHistorySubstateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterStatusHistorySubstateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterStatusHistorySubstateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterStatusHistorySubstateEnum(c *Client, desired, actual *ClusterStatusHistorySubstateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Cluster) urlNormalized() *Cluster {
	normalized := deepcopy.Copy(*r).(Cluster)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Cluster) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Cluster) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Cluster) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Cluster) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateCluster" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{location}}/clusters/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, fields), nil

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

	return flattenCluster(c, m), nil
}

// expandCluster expands Cluster into a JSON request object.
func expandCluster(c *Client, f *Cluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v, err := expandClusterClusterConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandClusterStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v, err := expandClusterStatusHistorySlice(c, f.StatusHistory); err != nil {
		return nil, fmt.Errorf("error expanding StatusHistory into statusHistory: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["statusHistory"] = v
	}
	if v := f.ClusterUuid; !dcl.IsEmptyValueIndirect(v) {
		m["clusterUuid"] = v
	}
	if v, err := expandClusterMetrics(c, f.Metrics); err != nil {
		return nil, fmt.Errorf("error expanding Metrics into metrics: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metrics"] = v
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

	r := &Cluster{}
	r.Project = dcl.FlattenString(m["projectId"])
	r.Name = dcl.FlattenString(m["clusterName"])
	r.Config = flattenClusterClusterConfig(c, m["config"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Status = flattenClusterStatus(c, m["status"])
	r.StatusHistory = flattenClusterStatusHistorySlice(c, m["statusHistory"])
	r.ClusterUuid = dcl.FlattenString(m["clusterUuid"])
	r.Metrics = flattenClusterMetrics(c, m["metrics"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandClusterClusterConfigMap expands the contents of ClusterClusterConfig into a JSON
// request object.
func expandClusterClusterConfigMap(c *Client, f map[string]ClusterClusterConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigSlice expands the contents of ClusterClusterConfig into a JSON
// request object.
func expandClusterClusterConfigSlice(c *Client, f []ClusterClusterConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigMap flattens the contents of ClusterClusterConfig from a JSON
// response object.
func flattenClusterClusterConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfig{}
	}

	items := make(map[string]ClusterClusterConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigSlice flattens the contents of ClusterClusterConfig from a JSON
// response object.
func flattenClusterClusterConfigSlice(c *Client, i interface{}) []ClusterClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfig{}
	}

	items := make([]ClusterClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfig expands an instance of ClusterClusterConfig into a JSON
// request object.
func expandClusterClusterConfig(c *Client, f *ClusterClusterConfig) (map[string]interface{}, error) {
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
	if v, err := expandClusterClusterConfigGceClusterConfig(c, f.GceClusterConfig); err != nil {
		return nil, fmt.Errorf("error expanding GceClusterConfig into gceClusterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gceClusterConfig"] = v
	}
	if v, err := expandClusterInstanceGroupConfig(c, f.MasterConfig); err != nil {
		return nil, fmt.Errorf("error expanding MasterConfig into masterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterConfig"] = v
	}
	if v, err := expandClusterInstanceGroupConfig(c, f.WorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workerConfig"] = v
	}
	if v, err := expandClusterInstanceGroupConfig(c, f.SecondaryWorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryWorkerConfig into secondaryWorkerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryWorkerConfig"] = v
	}
	if v, err := expandClusterClusterConfigSoftwareConfig(c, f.SoftwareConfig); err != nil {
		return nil, fmt.Errorf("error expanding SoftwareConfig into softwareConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["softwareConfig"] = v
	}
	if v, err := expandClusterClusterConfigInitializationActionsSlice(c, f.InitializationActions); err != nil {
		return nil, fmt.Errorf("error expanding InitializationActions into initializationActions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["initializationActions"] = v
	}
	if v, err := expandClusterClusterConfigEncryptionConfig(c, f.EncryptionConfig); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionConfig into encryptionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionConfig"] = v
	}
	if v, err := expandClusterClusterConfigAutoscalingConfig(c, f.AutoscalingConfig); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingConfig into autoscalingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscalingConfig"] = v
	}
	if v, err := expandClusterClusterConfigSecurityConfig(c, f.SecurityConfig); err != nil {
		return nil, fmt.Errorf("error expanding SecurityConfig into securityConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["securityConfig"] = v
	}
	if v, err := expandClusterClusterConfigLifecycleConfig(c, f.LifecycleConfig); err != nil {
		return nil, fmt.Errorf("error expanding LifecycleConfig into lifecycleConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["lifecycleConfig"] = v
	}
	if v, err := expandClusterClusterConfigEndpointConfig(c, f.EndpointConfig); err != nil {
		return nil, fmt.Errorf("error expanding EndpointConfig into endpointConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpointConfig"] = v
	}

	return m, nil
}

// flattenClusterClusterConfig flattens an instance of ClusterClusterConfig from a JSON
// response object.
func flattenClusterClusterConfig(c *Client, i interface{}) *ClusterClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfig{}
	r.StagingBucket = dcl.FlattenString(m["configBucket"])
	r.TempBucket = dcl.FlattenString(m["tempBucket"])
	r.GceClusterConfig = flattenClusterClusterConfigGceClusterConfig(c, m["gceClusterConfig"])
	r.MasterConfig = flattenClusterInstanceGroupConfig(c, m["masterConfig"])
	r.WorkerConfig = flattenClusterInstanceGroupConfig(c, m["workerConfig"])
	r.SecondaryWorkerConfig = flattenClusterInstanceGroupConfig(c, m["secondaryWorkerConfig"])
	r.SoftwareConfig = flattenClusterClusterConfigSoftwareConfig(c, m["softwareConfig"])
	r.InitializationActions = flattenClusterClusterConfigInitializationActionsSlice(c, m["initializationActions"])
	r.EncryptionConfig = flattenClusterClusterConfigEncryptionConfig(c, m["encryptionConfig"])
	r.AutoscalingConfig = flattenClusterClusterConfigAutoscalingConfig(c, m["autoscalingConfig"])
	r.SecurityConfig = flattenClusterClusterConfigSecurityConfig(c, m["securityConfig"])
	r.LifecycleConfig = flattenClusterClusterConfigLifecycleConfig(c, m["lifecycleConfig"])
	r.EndpointConfig = flattenClusterClusterConfigEndpointConfig(c, m["endpointConfig"])

	return r
}

// expandClusterClusterConfigGceClusterConfigMap expands the contents of ClusterClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigMap(c *Client, f map[string]ClusterClusterConfigGceClusterConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigGceClusterConfigSlice expands the contents of ClusterClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigSlice(c *Client, f []ClusterClusterConfigGceClusterConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigGceClusterConfigMap flattens the contents of ClusterClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigGceClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigGceClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigGceClusterConfig{}
	}

	items := make(map[string]ClusterClusterConfigGceClusterConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigGceClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigGceClusterConfigSlice flattens the contents of ClusterClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigSlice(c *Client, i interface{}) []ClusterClusterConfigGceClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigGceClusterConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigGceClusterConfig{}
	}

	items := make([]ClusterClusterConfigGceClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigGceClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigGceClusterConfig expands an instance of ClusterClusterConfigGceClusterConfig into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfig(c *Client, f *ClusterClusterConfigGceClusterConfig) (map[string]interface{}, error) {
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
	if v := f.ServiceAccountScopes; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountScopes"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := expandClusterClusterConfigGceClusterConfigReservationAffinity(c, f.ReservationAffinity); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, f.NodeGroupAffinity); err != nil {
		return nil, fmt.Errorf("error expanding NodeGroupAffinity into nodeGroupAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroupAffinity"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigGceClusterConfig flattens an instance of ClusterClusterConfigGceClusterConfig from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfig(c *Client, i interface{}) *ClusterClusterConfigGceClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigGceClusterConfig{}
	r.Zone = dcl.FlattenString(m["zoneUri"])
	r.Network = dcl.FlattenString(m["networkUri"])
	r.Subnetwork = dcl.FlattenString(m["subnetworkUri"])
	r.InternalIPOnly = dcl.FlattenBool(m["internalIpOnly"])
	r.PrivateIPv6GoogleAccess = flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(m["privateIpv6GoogleAccess"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.ServiceAccountScopes = dcl.FlattenStringSlice(m["serviceAccountScopes"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ReservationAffinity = flattenClusterClusterConfigGceClusterConfigReservationAffinity(c, m["reservationAffinity"])
	r.NodeGroupAffinity = flattenClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, m["nodeGroupAffinity"])

	return r
}

// expandClusterClusterConfigGceClusterConfigReservationAffinityMap expands the contents of ClusterClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigReservationAffinityMap(c *Client, f map[string]ClusterClusterConfigGceClusterConfigReservationAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigGceClusterConfigReservationAffinitySlice expands the contents of ClusterClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, f []ClusterClusterConfigGceClusterConfigReservationAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigGceClusterConfigReservationAffinityMap flattens the contents of ClusterClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigReservationAffinityMap(c *Client, i interface{}) map[string]ClusterClusterConfigGceClusterConfigReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigGceClusterConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigGceClusterConfigReservationAffinity{}
	}

	items := make(map[string]ClusterClusterConfigGceClusterConfigReservationAffinity)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigGceClusterConfigReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigGceClusterConfigReservationAffinitySlice flattens the contents of ClusterClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigReservationAffinitySlice(c *Client, i interface{}) []ClusterClusterConfigGceClusterConfigReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigGceClusterConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigGceClusterConfigReservationAffinity{}
	}

	items := make([]ClusterClusterConfigGceClusterConfigReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigGceClusterConfigReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigGceClusterConfigReservationAffinity expands an instance of ClusterClusterConfigGceClusterConfigReservationAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigReservationAffinity(c *Client, f *ClusterClusterConfigGceClusterConfigReservationAffinity) (map[string]interface{}, error) {
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
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigGceClusterConfigReservationAffinity flattens an instance of ClusterClusterConfigGceClusterConfigReservationAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigReservationAffinity(c *Client, i interface{}) *ClusterClusterConfigGceClusterConfigReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigGceClusterConfigReservationAffinity{}
	r.ConsumeReservationType = flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(m["consumeReservationType"])
	r.Key = dcl.FlattenString(m["key"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandClusterClusterConfigGceClusterConfigNodeGroupAffinityMap expands the contents of ClusterClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigNodeGroupAffinityMap(c *Client, f map[string]ClusterClusterConfigGceClusterConfigNodeGroupAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigGceClusterConfigNodeGroupAffinitySlice expands the contents of ClusterClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, f []ClusterClusterConfigGceClusterConfigNodeGroupAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigGceClusterConfigNodeGroupAffinityMap flattens the contents of ClusterClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigNodeGroupAffinityMap(c *Client, i interface{}) map[string]ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	items := make(map[string]ClusterClusterConfigGceClusterConfigNodeGroupAffinity)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigGceClusterConfigNodeGroupAffinitySlice flattens the contents of ClusterClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigNodeGroupAffinitySlice(c *Client, i interface{}) []ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	}

	items := make([]ClusterClusterConfigGceClusterConfigNodeGroupAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigGceClusterConfigNodeGroupAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigGceClusterConfigNodeGroupAffinity expands an instance of ClusterClusterConfigGceClusterConfigNodeGroupAffinity into a JSON
// request object.
func expandClusterClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, f *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NodeGroup; !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroupUri"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigGceClusterConfigNodeGroupAffinity flattens an instance of ClusterClusterConfigGceClusterConfigNodeGroupAffinity from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigNodeGroupAffinity(c *Client, i interface{}) *ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	r.NodeGroup = dcl.FlattenString(m["nodeGroupUri"])

	return r
}

// expandClusterInstanceGroupConfigMap expands the contents of ClusterInstanceGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigMap(c *Client, f map[string]ClusterInstanceGroupConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterInstanceGroupConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterInstanceGroupConfigSlice expands the contents of ClusterInstanceGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigSlice(c *Client, f []ClusterInstanceGroupConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterInstanceGroupConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterInstanceGroupConfigMap flattens the contents of ClusterInstanceGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigMap(c *Client, i interface{}) map[string]ClusterInstanceGroupConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterInstanceGroupConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterInstanceGroupConfig{}
	}

	items := make(map[string]ClusterInstanceGroupConfig)
	for k, item := range a {
		items[k] = *flattenClusterInstanceGroupConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterInstanceGroupConfigSlice flattens the contents of ClusterInstanceGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigSlice(c *Client, i interface{}) []ClusterInstanceGroupConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterInstanceGroupConfig{}
	}

	if len(a) == 0 {
		return []ClusterInstanceGroupConfig{}
	}

	items := make([]ClusterInstanceGroupConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterInstanceGroupConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterInstanceGroupConfig expands an instance of ClusterInstanceGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfig(c *Client, f *ClusterInstanceGroupConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumInstances; !dcl.IsEmptyValueIndirect(v) {
		m["numInstances"] = v
	}
	if v := f.InstanceNames; !dcl.IsEmptyValueIndirect(v) {
		m["instanceNames"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["imageUri"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineTypeUri"] = v
	}
	if v, err := expandClusterInstanceGroupConfigDiskConfig(c, f.DiskConfig); err != nil {
		return nil, fmt.Errorf("error expanding DiskConfig into diskConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskConfig"] = v
	}
	if v := f.IsPreemptible; !dcl.IsEmptyValueIndirect(v) {
		m["isPreemptible"] = v
	}
	if v := f.Preemptibility; !dcl.IsEmptyValueIndirect(v) {
		m["preemptibility"] = v
	}
	if v, err := expandClusterInstanceGroupConfigManagedGroupConfig(c, f.ManagedGroupConfig); err != nil {
		return nil, fmt.Errorf("error expanding ManagedGroupConfig into managedGroupConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["managedGroupConfig"] = v
	}
	if v, err := expandClusterInstanceGroupConfigAcceleratorsSlice(c, f.Accelerators); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accelerators"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}

	return m, nil
}

// flattenClusterInstanceGroupConfig flattens an instance of ClusterInstanceGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfig(c *Client, i interface{}) *ClusterInstanceGroupConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterInstanceGroupConfig{}
	r.NumInstances = dcl.FlattenInteger(m["numInstances"])
	r.InstanceNames = dcl.FlattenStringSlice(m["instanceNames"])
	r.Image = dcl.FlattenString(m["imageUri"])
	r.MachineType = dcl.FlattenString(m["machineTypeUri"])
	r.DiskConfig = flattenClusterInstanceGroupConfigDiskConfig(c, m["diskConfig"])
	r.IsPreemptible = dcl.FlattenBool(m["isPreemptible"])
	r.Preemptibility = flattenClusterInstanceGroupConfigPreemptibilityEnum(m["preemptibility"])
	r.ManagedGroupConfig = flattenClusterInstanceGroupConfigManagedGroupConfig(c, m["managedGroupConfig"])
	r.Accelerators = flattenClusterInstanceGroupConfigAcceleratorsSlice(c, m["accelerators"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])

	return r
}

// expandClusterInstanceGroupConfigDiskConfigMap expands the contents of ClusterInstanceGroupConfigDiskConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigDiskConfigMap(c *Client, f map[string]ClusterInstanceGroupConfigDiskConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterInstanceGroupConfigDiskConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterInstanceGroupConfigDiskConfigSlice expands the contents of ClusterInstanceGroupConfigDiskConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigDiskConfigSlice(c *Client, f []ClusterInstanceGroupConfigDiskConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterInstanceGroupConfigDiskConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterInstanceGroupConfigDiskConfigMap flattens the contents of ClusterInstanceGroupConfigDiskConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigDiskConfigMap(c *Client, i interface{}) map[string]ClusterInstanceGroupConfigDiskConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterInstanceGroupConfigDiskConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterInstanceGroupConfigDiskConfig{}
	}

	items := make(map[string]ClusterInstanceGroupConfigDiskConfig)
	for k, item := range a {
		items[k] = *flattenClusterInstanceGroupConfigDiskConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterInstanceGroupConfigDiskConfigSlice flattens the contents of ClusterInstanceGroupConfigDiskConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigDiskConfigSlice(c *Client, i interface{}) []ClusterInstanceGroupConfigDiskConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterInstanceGroupConfigDiskConfig{}
	}

	if len(a) == 0 {
		return []ClusterInstanceGroupConfigDiskConfig{}
	}

	items := make([]ClusterInstanceGroupConfigDiskConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterInstanceGroupConfigDiskConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterInstanceGroupConfigDiskConfig expands an instance of ClusterInstanceGroupConfigDiskConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigDiskConfig(c *Client, f *ClusterInstanceGroupConfigDiskConfig) (map[string]interface{}, error) {
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

// flattenClusterInstanceGroupConfigDiskConfig flattens an instance of ClusterInstanceGroupConfigDiskConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigDiskConfig(c *Client, i interface{}) *ClusterInstanceGroupConfigDiskConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterInstanceGroupConfigDiskConfig{}
	r.BootDiskType = dcl.FlattenString(m["bootDiskType"])
	r.BootDiskSizeGb = dcl.FlattenInteger(m["bootDiskSizeGb"])
	r.NumLocalSsds = dcl.FlattenInteger(m["numLocalSsds"])

	return r
}

// expandClusterInstanceGroupConfigManagedGroupConfigMap expands the contents of ClusterInstanceGroupConfigManagedGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigManagedGroupConfigMap(c *Client, f map[string]ClusterInstanceGroupConfigManagedGroupConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterInstanceGroupConfigManagedGroupConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterInstanceGroupConfigManagedGroupConfigSlice expands the contents of ClusterInstanceGroupConfigManagedGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigManagedGroupConfigSlice(c *Client, f []ClusterInstanceGroupConfigManagedGroupConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterInstanceGroupConfigManagedGroupConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterInstanceGroupConfigManagedGroupConfigMap flattens the contents of ClusterInstanceGroupConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigManagedGroupConfigMap(c *Client, i interface{}) map[string]ClusterInstanceGroupConfigManagedGroupConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterInstanceGroupConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterInstanceGroupConfigManagedGroupConfig{}
	}

	items := make(map[string]ClusterInstanceGroupConfigManagedGroupConfig)
	for k, item := range a {
		items[k] = *flattenClusterInstanceGroupConfigManagedGroupConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterInstanceGroupConfigManagedGroupConfigSlice flattens the contents of ClusterInstanceGroupConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigManagedGroupConfigSlice(c *Client, i interface{}) []ClusterInstanceGroupConfigManagedGroupConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterInstanceGroupConfigManagedGroupConfig{}
	}

	if len(a) == 0 {
		return []ClusterInstanceGroupConfigManagedGroupConfig{}
	}

	items := make([]ClusterInstanceGroupConfigManagedGroupConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterInstanceGroupConfigManagedGroupConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterInstanceGroupConfigManagedGroupConfig expands an instance of ClusterInstanceGroupConfigManagedGroupConfig into a JSON
// request object.
func expandClusterInstanceGroupConfigManagedGroupConfig(c *Client, f *ClusterInstanceGroupConfigManagedGroupConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.InstanceTemplateName; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTemplateName"] = v
	}
	if v := f.InstanceGroupManagerName; !dcl.IsEmptyValueIndirect(v) {
		m["instanceGroupManagerName"] = v
	}

	return m, nil
}

// flattenClusterInstanceGroupConfigManagedGroupConfig flattens an instance of ClusterInstanceGroupConfigManagedGroupConfig from a JSON
// response object.
func flattenClusterInstanceGroupConfigManagedGroupConfig(c *Client, i interface{}) *ClusterInstanceGroupConfigManagedGroupConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterInstanceGroupConfigManagedGroupConfig{}
	r.InstanceTemplateName = dcl.FlattenString(m["instanceTemplateName"])
	r.InstanceGroupManagerName = dcl.FlattenString(m["instanceGroupManagerName"])

	return r
}

// expandClusterInstanceGroupConfigAcceleratorsMap expands the contents of ClusterInstanceGroupConfigAccelerators into a JSON
// request object.
func expandClusterInstanceGroupConfigAcceleratorsMap(c *Client, f map[string]ClusterInstanceGroupConfigAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterInstanceGroupConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterInstanceGroupConfigAcceleratorsSlice expands the contents of ClusterInstanceGroupConfigAccelerators into a JSON
// request object.
func expandClusterInstanceGroupConfigAcceleratorsSlice(c *Client, f []ClusterInstanceGroupConfigAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterInstanceGroupConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterInstanceGroupConfigAcceleratorsMap flattens the contents of ClusterInstanceGroupConfigAccelerators from a JSON
// response object.
func flattenClusterInstanceGroupConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterInstanceGroupConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterInstanceGroupConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterInstanceGroupConfigAccelerators{}
	}

	items := make(map[string]ClusterInstanceGroupConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterInstanceGroupConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterInstanceGroupConfigAcceleratorsSlice flattens the contents of ClusterInstanceGroupConfigAccelerators from a JSON
// response object.
func flattenClusterInstanceGroupConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterInstanceGroupConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterInstanceGroupConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterInstanceGroupConfigAccelerators{}
	}

	items := make([]ClusterInstanceGroupConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterInstanceGroupConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterInstanceGroupConfigAccelerators expands an instance of ClusterInstanceGroupConfigAccelerators into a JSON
// request object.
func expandClusterInstanceGroupConfigAccelerators(c *Client, f *ClusterInstanceGroupConfigAccelerators) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
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

// flattenClusterInstanceGroupConfigAccelerators flattens an instance of ClusterInstanceGroupConfigAccelerators from a JSON
// response object.
func flattenClusterInstanceGroupConfigAccelerators(c *Client, i interface{}) *ClusterInstanceGroupConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterInstanceGroupConfigAccelerators{}
	r.AcceleratorType = dcl.FlattenString(m["acceleratorTypeUri"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandClusterClusterConfigSoftwareConfigMap expands the contents of ClusterClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterClusterConfigSoftwareConfigMap(c *Client, f map[string]ClusterClusterConfigSoftwareConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigSoftwareConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigSoftwareConfigSlice expands the contents of ClusterClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterClusterConfigSoftwareConfigSlice(c *Client, f []ClusterClusterConfigSoftwareConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigSoftwareConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigSoftwareConfigMap flattens the contents of ClusterClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterClusterConfigSoftwareConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigSoftwareConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigSoftwareConfig{}
	}

	items := make(map[string]ClusterClusterConfigSoftwareConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigSoftwareConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigSoftwareConfigSlice flattens the contents of ClusterClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterClusterConfigSoftwareConfigSlice(c *Client, i interface{}) []ClusterClusterConfigSoftwareConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigSoftwareConfig{}
	}

	items := make([]ClusterClusterConfigSoftwareConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigSoftwareConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigSoftwareConfig expands an instance of ClusterClusterConfigSoftwareConfig into a JSON
// request object.
func expandClusterClusterConfigSoftwareConfig(c *Client, f *ClusterClusterConfigSoftwareConfig) (map[string]interface{}, error) {
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
	if v := f.OptionalComponents; !dcl.IsEmptyValueIndirect(v) {
		m["optionalComponents"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigSoftwareConfig flattens an instance of ClusterClusterConfigSoftwareConfig from a JSON
// response object.
func flattenClusterClusterConfigSoftwareConfig(c *Client, i interface{}) *ClusterClusterConfigSoftwareConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigSoftwareConfig{}
	r.ImageVersion = dcl.FlattenString(m["imageVersion"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.OptionalComponents = flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c, m["optionalComponents"])

	return r
}

// expandClusterClusterConfigInitializationActionsMap expands the contents of ClusterClusterConfigInitializationActions into a JSON
// request object.
func expandClusterClusterConfigInitializationActionsMap(c *Client, f map[string]ClusterClusterConfigInitializationActions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigInitializationActions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigInitializationActionsSlice expands the contents of ClusterClusterConfigInitializationActions into a JSON
// request object.
func expandClusterClusterConfigInitializationActionsSlice(c *Client, f []ClusterClusterConfigInitializationActions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigInitializationActions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigInitializationActionsMap flattens the contents of ClusterClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterClusterConfigInitializationActionsMap(c *Client, i interface{}) map[string]ClusterClusterConfigInitializationActions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigInitializationActions{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigInitializationActions{}
	}

	items := make(map[string]ClusterClusterConfigInitializationActions)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigInitializationActions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigInitializationActionsSlice flattens the contents of ClusterClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterClusterConfigInitializationActionsSlice(c *Client, i interface{}) []ClusterClusterConfigInitializationActions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigInitializationActions{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigInitializationActions{}
	}

	items := make([]ClusterClusterConfigInitializationActions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigInitializationActions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigInitializationActions expands an instance of ClusterClusterConfigInitializationActions into a JSON
// request object.
func expandClusterClusterConfigInitializationActions(c *Client, f *ClusterClusterConfigInitializationActions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
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

// flattenClusterClusterConfigInitializationActions flattens an instance of ClusterClusterConfigInitializationActions from a JSON
// response object.
func flattenClusterClusterConfigInitializationActions(c *Client, i interface{}) *ClusterClusterConfigInitializationActions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigInitializationActions{}
	r.ExecutableFile = dcl.FlattenString(m["executableFile"])
	r.ExecutionTimeout = dcl.FlattenString(m["executionTimeout"])

	return r
}

// expandClusterClusterConfigEncryptionConfigMap expands the contents of ClusterClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterClusterConfigEncryptionConfigMap(c *Client, f map[string]ClusterClusterConfigEncryptionConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigEncryptionConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigEncryptionConfigSlice expands the contents of ClusterClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterClusterConfigEncryptionConfigSlice(c *Client, f []ClusterClusterConfigEncryptionConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigEncryptionConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigEncryptionConfigMap flattens the contents of ClusterClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterClusterConfigEncryptionConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigEncryptionConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigEncryptionConfig{}
	}

	items := make(map[string]ClusterClusterConfigEncryptionConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigEncryptionConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigEncryptionConfigSlice flattens the contents of ClusterClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterClusterConfigEncryptionConfigSlice(c *Client, i interface{}) []ClusterClusterConfigEncryptionConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigEncryptionConfig{}
	}

	items := make([]ClusterClusterConfigEncryptionConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigEncryptionConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigEncryptionConfig expands an instance of ClusterClusterConfigEncryptionConfig into a JSON
// request object.
func expandClusterClusterConfigEncryptionConfig(c *Client, f *ClusterClusterConfigEncryptionConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.GcePdKmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["gcePdKmsKeyName"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigEncryptionConfig flattens an instance of ClusterClusterConfigEncryptionConfig from a JSON
// response object.
func flattenClusterClusterConfigEncryptionConfig(c *Client, i interface{}) *ClusterClusterConfigEncryptionConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigEncryptionConfig{}
	r.GcePdKmsKeyName = dcl.FlattenString(m["gcePdKmsKeyName"])

	return r
}

// expandClusterClusterConfigAutoscalingConfigMap expands the contents of ClusterClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterClusterConfigAutoscalingConfigMap(c *Client, f map[string]ClusterClusterConfigAutoscalingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigAutoscalingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigAutoscalingConfigSlice expands the contents of ClusterClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterClusterConfigAutoscalingConfigSlice(c *Client, f []ClusterClusterConfigAutoscalingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigAutoscalingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigAutoscalingConfigMap flattens the contents of ClusterClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterClusterConfigAutoscalingConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigAutoscalingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigAutoscalingConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigAutoscalingConfig{}
	}

	items := make(map[string]ClusterClusterConfigAutoscalingConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigAutoscalingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigAutoscalingConfigSlice flattens the contents of ClusterClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterClusterConfigAutoscalingConfigSlice(c *Client, i interface{}) []ClusterClusterConfigAutoscalingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigAutoscalingConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigAutoscalingConfig{}
	}

	items := make([]ClusterClusterConfigAutoscalingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigAutoscalingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigAutoscalingConfig expands an instance of ClusterClusterConfigAutoscalingConfig into a JSON
// request object.
func expandClusterClusterConfigAutoscalingConfig(c *Client, f *ClusterClusterConfigAutoscalingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Policy; !dcl.IsEmptyValueIndirect(v) {
		m["policyUri"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigAutoscalingConfig flattens an instance of ClusterClusterConfigAutoscalingConfig from a JSON
// response object.
func flattenClusterClusterConfigAutoscalingConfig(c *Client, i interface{}) *ClusterClusterConfigAutoscalingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigAutoscalingConfig{}
	r.Policy = dcl.FlattenString(m["policyUri"])

	return r
}

// expandClusterClusterConfigSecurityConfigMap expands the contents of ClusterClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfigMap(c *Client, f map[string]ClusterClusterConfigSecurityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigSecurityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigSecurityConfigSlice expands the contents of ClusterClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfigSlice(c *Client, f []ClusterClusterConfigSecurityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigSecurityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigSecurityConfigMap flattens the contents of ClusterClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigSecurityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigSecurityConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigSecurityConfig{}
	}

	items := make(map[string]ClusterClusterConfigSecurityConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigSecurityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigSecurityConfigSlice flattens the contents of ClusterClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfigSlice(c *Client, i interface{}) []ClusterClusterConfigSecurityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigSecurityConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigSecurityConfig{}
	}

	items := make([]ClusterClusterConfigSecurityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigSecurityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigSecurityConfig expands an instance of ClusterClusterConfigSecurityConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfig(c *Client, f *ClusterClusterConfigSecurityConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandClusterClusterConfigSecurityConfigKerberosConfig(c, f.KerberosConfig); err != nil {
		return nil, fmt.Errorf("error expanding KerberosConfig into kerberosConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kerberosConfig"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigSecurityConfig flattens an instance of ClusterClusterConfigSecurityConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfig(c *Client, i interface{}) *ClusterClusterConfigSecurityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigSecurityConfig{}
	r.KerberosConfig = flattenClusterClusterConfigSecurityConfigKerberosConfig(c, m["kerberosConfig"])

	return r
}

// expandClusterClusterConfigSecurityConfigKerberosConfigMap expands the contents of ClusterClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfigKerberosConfigMap(c *Client, f map[string]ClusterClusterConfigSecurityConfigKerberosConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigSecurityConfigKerberosConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigSecurityConfigKerberosConfigSlice expands the contents of ClusterClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfigKerberosConfigSlice(c *Client, f []ClusterClusterConfigSecurityConfigKerberosConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigSecurityConfigKerberosConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigSecurityConfigKerberosConfigMap flattens the contents of ClusterClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfigKerberosConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigSecurityConfigKerberosConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigSecurityConfigKerberosConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigSecurityConfigKerberosConfig{}
	}

	items := make(map[string]ClusterClusterConfigSecurityConfigKerberosConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigSecurityConfigKerberosConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigSecurityConfigKerberosConfigSlice flattens the contents of ClusterClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfigKerberosConfigSlice(c *Client, i interface{}) []ClusterClusterConfigSecurityConfigKerberosConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigSecurityConfigKerberosConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigSecurityConfigKerberosConfig{}
	}

	items := make([]ClusterClusterConfigSecurityConfigKerberosConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigSecurityConfigKerberosConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigSecurityConfigKerberosConfig expands an instance of ClusterClusterConfigSecurityConfigKerberosConfig into a JSON
// request object.
func expandClusterClusterConfigSecurityConfigKerberosConfig(c *Client, f *ClusterClusterConfigSecurityConfigKerberosConfig) (map[string]interface{}, error) {
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

// flattenClusterClusterConfigSecurityConfigKerberosConfig flattens an instance of ClusterClusterConfigSecurityConfigKerberosConfig from a JSON
// response object.
func flattenClusterClusterConfigSecurityConfigKerberosConfig(c *Client, i interface{}) *ClusterClusterConfigSecurityConfigKerberosConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigSecurityConfigKerberosConfig{}
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

// expandClusterClusterConfigLifecycleConfigMap expands the contents of ClusterClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterClusterConfigLifecycleConfigMap(c *Client, f map[string]ClusterClusterConfigLifecycleConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigLifecycleConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigLifecycleConfigSlice expands the contents of ClusterClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterClusterConfigLifecycleConfigSlice(c *Client, f []ClusterClusterConfigLifecycleConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigLifecycleConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigLifecycleConfigMap flattens the contents of ClusterClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterClusterConfigLifecycleConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigLifecycleConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigLifecycleConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigLifecycleConfig{}
	}

	items := make(map[string]ClusterClusterConfigLifecycleConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigLifecycleConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigLifecycleConfigSlice flattens the contents of ClusterClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterClusterConfigLifecycleConfigSlice(c *Client, i interface{}) []ClusterClusterConfigLifecycleConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigLifecycleConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigLifecycleConfig{}
	}

	items := make([]ClusterClusterConfigLifecycleConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigLifecycleConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigLifecycleConfig expands an instance of ClusterClusterConfigLifecycleConfig into a JSON
// request object.
func expandClusterClusterConfigLifecycleConfig(c *Client, f *ClusterClusterConfigLifecycleConfig) (map[string]interface{}, error) {
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
	if v := f.IdleStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["idleStartTime"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigLifecycleConfig flattens an instance of ClusterClusterConfigLifecycleConfig from a JSON
// response object.
func flattenClusterClusterConfigLifecycleConfig(c *Client, i interface{}) *ClusterClusterConfigLifecycleConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigLifecycleConfig{}
	r.IdleDeleteTtl = dcl.FlattenString(m["idleDeleteTtl"])
	r.AutoDeleteTime = dcl.FlattenString(m["autoDeleteTime"])
	r.AutoDeleteTtl = dcl.FlattenString(m["autoDeleteTtl"])
	r.IdleStartTime = dcl.FlattenString(m["idleStartTime"])

	return r
}

// expandClusterClusterConfigEndpointConfigMap expands the contents of ClusterClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterClusterConfigEndpointConfigMap(c *Client, f map[string]ClusterClusterConfigEndpointConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterClusterConfigEndpointConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterClusterConfigEndpointConfigSlice expands the contents of ClusterClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterClusterConfigEndpointConfigSlice(c *Client, f []ClusterClusterConfigEndpointConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterClusterConfigEndpointConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterClusterConfigEndpointConfigMap flattens the contents of ClusterClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterClusterConfigEndpointConfigMap(c *Client, i interface{}) map[string]ClusterClusterConfigEndpointConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterClusterConfigEndpointConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterClusterConfigEndpointConfig{}
	}

	items := make(map[string]ClusterClusterConfigEndpointConfig)
	for k, item := range a {
		items[k] = *flattenClusterClusterConfigEndpointConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterClusterConfigEndpointConfigSlice flattens the contents of ClusterClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterClusterConfigEndpointConfigSlice(c *Client, i interface{}) []ClusterClusterConfigEndpointConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigEndpointConfig{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigEndpointConfig{}
	}

	items := make([]ClusterClusterConfigEndpointConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigEndpointConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterClusterConfigEndpointConfig expands an instance of ClusterClusterConfigEndpointConfig into a JSON
// request object.
func expandClusterClusterConfigEndpointConfig(c *Client, f *ClusterClusterConfigEndpointConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpPorts; !dcl.IsEmptyValueIndirect(v) {
		m["httpPorts"] = v
	}
	if v := f.EnableHttpPortAccess; !dcl.IsEmptyValueIndirect(v) {
		m["enableHttpPortAccess"] = v
	}

	return m, nil
}

// flattenClusterClusterConfigEndpointConfig flattens an instance of ClusterClusterConfigEndpointConfig from a JSON
// response object.
func flattenClusterClusterConfigEndpointConfig(c *Client, i interface{}) *ClusterClusterConfigEndpointConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterClusterConfigEndpointConfig{}
	r.HttpPorts = dcl.FlattenKeyValuePairs(m["httpPorts"])
	r.EnableHttpPortAccess = dcl.FlattenBool(m["enableHttpPortAccess"])

	return r
}

// expandClusterStatusMap expands the contents of ClusterStatus into a JSON
// request object.
func expandClusterStatusMap(c *Client, f map[string]ClusterStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterStatus(c, &item)
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
func expandClusterStatusSlice(c *Client, f []ClusterStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterStatus(c, &item)
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
func expandClusterStatus(c *Client, f *ClusterStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Detail; !dcl.IsEmptyValueIndirect(v) {
		m["detail"] = v
	}
	if v := f.StateStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["stateStartTime"] = v
	}
	if v := f.Substate; !dcl.IsEmptyValueIndirect(v) {
		m["substate"] = v
	}

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
	r.State = flattenClusterStatusStateEnum(m["state"])
	r.Detail = dcl.FlattenString(m["detail"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenClusterStatusSubstateEnum(m["substate"])

	return r
}

// expandClusterStatusHistoryMap expands the contents of ClusterStatusHistory into a JSON
// request object.
func expandClusterStatusHistoryMap(c *Client, f map[string]ClusterStatusHistory) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterStatusHistory(c, &item)
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
func expandClusterStatusHistorySlice(c *Client, f []ClusterStatusHistory) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterStatusHistory(c, &item)
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
func expandClusterStatusHistory(c *Client, f *ClusterStatusHistory) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Detail; !dcl.IsEmptyValueIndirect(v) {
		m["detail"] = v
	}
	if v := f.StateStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["stateStartTime"] = v
	}
	if v := f.Substate; !dcl.IsEmptyValueIndirect(v) {
		m["substate"] = v
	}

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
	r.State = flattenClusterStatusHistoryStateEnum(m["state"])
	r.Detail = dcl.FlattenString(m["detail"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenClusterStatusHistorySubstateEnum(m["substate"])

	return r
}

// expandClusterMetricsMap expands the contents of ClusterMetrics into a JSON
// request object.
func expandClusterMetricsMap(c *Client, f map[string]ClusterMetrics) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMetrics(c, &item)
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
func expandClusterMetricsSlice(c *Client, f []ClusterMetrics) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMetrics(c, &item)
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
func expandClusterMetrics(c *Client, f *ClusterMetrics) (map[string]interface{}, error) {
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
	r.HdfsMetrics = dcl.FlattenKeyValuePairs(m["hdfsMetrics"])
	r.YarnMetrics = dcl.FlattenKeyValuePairs(m["yarnMetrics"])

	return r
}

// flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumSlice flattens the contents of ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumSlice(c *Client, i interface{}) []ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum{}
	}

	items := make([]ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum with the same value as that string.
func flattenClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(i interface{}) *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef("")
	}

	return ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s)
}

// flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumSlice flattens the contents of ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, i interface{}) []ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make([]ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum with the same value as that string.
func flattenClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(i interface{}) *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef("")
	}

	return ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s)
}

// flattenClusterInstanceGroupConfigPreemptibilityEnumSlice flattens the contents of ClusterInstanceGroupConfigPreemptibilityEnum from a JSON
// response object.
func flattenClusterInstanceGroupConfigPreemptibilityEnumSlice(c *Client, i interface{}) []ClusterInstanceGroupConfigPreemptibilityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterInstanceGroupConfigPreemptibilityEnum{}
	}

	if len(a) == 0 {
		return []ClusterInstanceGroupConfigPreemptibilityEnum{}
	}

	items := make([]ClusterInstanceGroupConfigPreemptibilityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterInstanceGroupConfigPreemptibilityEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterInstanceGroupConfigPreemptibilityEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterInstanceGroupConfigPreemptibilityEnum with the same value as that string.
func flattenClusterInstanceGroupConfigPreemptibilityEnum(i interface{}) *ClusterInstanceGroupConfigPreemptibilityEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterInstanceGroupConfigPreemptibilityEnumRef("")
	}

	return ClusterInstanceGroupConfigPreemptibilityEnumRef(s)
}

// flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnumSlice flattens the contents of ClusterClusterConfigSoftwareConfigOptionalComponentsEnum from a JSON
// response object.
func flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnumSlice(c *Client, i interface{}) []ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	if len(a) == 0 {
		return []ClusterClusterConfigSoftwareConfigOptionalComponentsEnum{}
	}

	items := make([]ClusterClusterConfigSoftwareConfigOptionalComponentsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterClusterConfigSoftwareConfigOptionalComponentsEnum with the same value as that string.
func flattenClusterClusterConfigSoftwareConfigOptionalComponentsEnum(i interface{}) *ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterClusterConfigSoftwareConfigOptionalComponentsEnumRef("")
	}

	return ClusterClusterConfigSoftwareConfigOptionalComponentsEnumRef(s)
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
		items = append(items, *flattenClusterStatusStateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterStatusStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusStateEnum with the same value as that string.
func flattenClusterStatusStateEnum(i interface{}) *ClusterStatusStateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterStatusStateEnumRef("")
	}

	return ClusterStatusStateEnumRef(s)
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
		items = append(items, *flattenClusterStatusSubstateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterStatusSubstateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusSubstateEnum with the same value as that string.
func flattenClusterStatusSubstateEnum(i interface{}) *ClusterStatusSubstateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterStatusSubstateEnumRef("")
	}

	return ClusterStatusSubstateEnumRef(s)
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
		items = append(items, *flattenClusterStatusHistoryStateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterStatusHistoryStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusHistoryStateEnum with the same value as that string.
func flattenClusterStatusHistoryStateEnum(i interface{}) *ClusterStatusHistoryStateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterStatusHistoryStateEnumRef("")
	}

	return ClusterStatusHistoryStateEnumRef(s)
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
		items = append(items, *flattenClusterStatusHistorySubstateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenClusterStatusHistorySubstateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStatusHistorySubstateEnum with the same value as that string.
func flattenClusterStatusHistorySubstateEnum(i interface{}) *ClusterStatusHistorySubstateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterStatusHistorySubstateEnumRef("")
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
