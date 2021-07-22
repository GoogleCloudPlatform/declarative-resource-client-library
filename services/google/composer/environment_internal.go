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
package composer

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

func (r *Environment) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
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
	return nil
}
func (r *EnvironmentConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SoftwareConfig) {
		if err := r.SoftwareConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NodeConfig) {
		if err := r.NodeConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PrivateEnvironmentConfig) {
		if err := r.PrivateEnvironmentConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WebServerNetworkAccessControl) {
		if err := r.WebServerNetworkAccessControl.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DatabaseConfig) {
		if err := r.DatabaseConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WebServerConfig) {
		if err := r.WebServerConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionConfig) {
		if err := r.EncryptionConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EnvironmentConfigSoftwareConfig) validate() error {
	return nil
}
func (r *EnvironmentConfigNodeConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.IPAllocationPolicy) {
		if err := r.IPAllocationPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EnvironmentConfigNodeConfigIPAllocationPolicy) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ClusterSecondaryRangeName", "ClusterIPv4CidrBlock"}, r.ClusterSecondaryRangeName, r.ClusterIPv4CidrBlock); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ServicesSecondaryRangeName", "ServicesIPv4CidrBlock"}, r.ServicesSecondaryRangeName, r.ServicesIPv4CidrBlock); err != nil {
		return err
	}
	return nil
}
func (r *EnvironmentConfigPrivateEnvironmentConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PrivateClusterConfig) {
		if err := r.PrivateClusterConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) validate() error {
	return nil
}
func (r *EnvironmentConfigWebServerNetworkAccessControl) validate() error {
	return nil
}
func (r *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) validate() error {
	return nil
}
func (r *EnvironmentConfigDatabaseConfig) validate() error {
	return nil
}
func (r *EnvironmentConfigWebServerConfig) validate() error {
	return nil
}
func (r *EnvironmentConfigEncryptionConfig) validate() error {
	return nil
}

func environmentGetURL(userBasePath string, r *Environment) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/environments/{{name}}", "https://composer.googleapis.com/v1/", userBasePath, params), nil
}

func environmentListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/environments", "https://composer.googleapis.com/v1/", userBasePath, params), nil

}

func environmentCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/environments", "https://composer.googleapis.com/v1/", userBasePath, params), nil

}

func environmentDeleteURL(userBasePath string, r *Environment) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/environments/{{name}}", "https://composer.googleapis.com/v1/", userBasePath, params), nil
}

// environmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type environmentApiOperation interface {
	do(context.Context, *Environment, *Client) error
}

func (c *Client) listEnvironmentRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := environmentListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvironmentMaxPage {
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

type listEnvironmentOperation struct {
	Environments []map[string]interface{} `json:"environments"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listEnvironment(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Environment, string, error) {
	b, err := c.listEnvironmentRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEnvironmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Environment
	for _, v := range m.Environments {
		res, err := unmarshalMapEnvironment(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEnvironment(ctx context.Context, f func(*Environment) bool, resources []*Environment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvironment(ctx, res)
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

type deleteEnvironmentOperation struct{}

func (op *deleteEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	r, err := c.GetEnvironment(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Environment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetEnvironment checking for existence. error: %v", err)
		return err
	}

	u, err := environmentDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://composer.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetEnvironment(ctx, r.URLNormalized())
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
type createEnvironmentOperation struct {
	response map[string]interface{}
}

func (op *createEnvironmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := environmentCreateURL(c.Config.BasePath, project, location)

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
	if err := o.Wait(ctx, c.Config, "https://composer.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetEnvironment(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvironmentRaw(ctx context.Context, r *Environment) ([]byte, error) {

	u, err := environmentGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) environmentDiffsForRawDesired(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (initial, desired *Environment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Environment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Environment); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Environment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvironment(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Environment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Environment resource: %v", err)
		}
		c.Config.Logger.Info("Found that Environment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Environment: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Environment: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvironmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Environment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Environment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvironment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvironmentInitialState(rawInitial, rawDesired *Environment) (*Environment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvironmentDesiredState(rawDesired, rawInitial *Environment, opts ...dcl.ApplyOption) (*Environment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeEnvironmentConfig(rawDesired.Config, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Environment{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.Config = canonicalizeEnvironmentConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
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

func canonicalizeEnvironmentNewState(c *Client, rawNew, rawDesired *Environment) (*Environment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewEnvironmentConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uuid) && dcl.IsEmptyValueIndirect(rawDesired.Uuid) {
		rawNew.Uuid = rawDesired.Uuid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uuid, rawNew.Uuid) {
			rawNew.Uuid = rawDesired.Uuid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
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

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeEnvironmentConfig(des, initial *EnvironmentConfig, opts ...dcl.ApplyOption) *EnvironmentConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfig{}

	if dcl.IsZeroValue(des.NodeCount) {
		des.NodeCount = initial.NodeCount
	} else {
		cDes.NodeCount = des.NodeCount
	}
	cDes.SoftwareConfig = canonicalizeEnvironmentConfigSoftwareConfig(des.SoftwareConfig, initial.SoftwareConfig, opts...)
	cDes.NodeConfig = canonicalizeEnvironmentConfigNodeConfig(des.NodeConfig, initial.NodeConfig, opts...)
	cDes.PrivateEnvironmentConfig = canonicalizeEnvironmentConfigPrivateEnvironmentConfig(des.PrivateEnvironmentConfig, initial.PrivateEnvironmentConfig, opts...)
	cDes.WebServerNetworkAccessControl = canonicalizeEnvironmentConfigWebServerNetworkAccessControl(des.WebServerNetworkAccessControl, initial.WebServerNetworkAccessControl, opts...)
	cDes.DatabaseConfig = canonicalizeEnvironmentConfigDatabaseConfig(des.DatabaseConfig, initial.DatabaseConfig, opts...)
	cDes.WebServerConfig = canonicalizeEnvironmentConfigWebServerConfig(des.WebServerConfig, initial.WebServerConfig, opts...)
	cDes.EncryptionConfig = canonicalizeEnvironmentConfigEncryptionConfig(des.EncryptionConfig, initial.EncryptionConfig, opts...)

	return cDes
}

func canonicalizeNewEnvironmentConfig(c *Client, des, nw *EnvironmentConfig) *EnvironmentConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.GkeCluster, nw.GkeCluster) {
		nw.GkeCluster = des.GkeCluster
	}
	if dcl.StringCanonicalize(des.DagGcsPrefix, nw.DagGcsPrefix) {
		nw.DagGcsPrefix = des.DagGcsPrefix
	}
	if dcl.IsZeroValue(nw.NodeCount) {
		nw.NodeCount = des.NodeCount
	}
	nw.SoftwareConfig = canonicalizeNewEnvironmentConfigSoftwareConfig(c, des.SoftwareConfig, nw.SoftwareConfig)
	nw.NodeConfig = canonicalizeNewEnvironmentConfigNodeConfig(c, des.NodeConfig, nw.NodeConfig)
	nw.PrivateEnvironmentConfig = canonicalizeNewEnvironmentConfigPrivateEnvironmentConfig(c, des.PrivateEnvironmentConfig, nw.PrivateEnvironmentConfig)
	nw.WebServerNetworkAccessControl = canonicalizeNewEnvironmentConfigWebServerNetworkAccessControl(c, des.WebServerNetworkAccessControl, nw.WebServerNetworkAccessControl)
	nw.DatabaseConfig = canonicalizeNewEnvironmentConfigDatabaseConfig(c, des.DatabaseConfig, nw.DatabaseConfig)
	nw.WebServerConfig = canonicalizeNewEnvironmentConfigWebServerConfig(c, des.WebServerConfig, nw.WebServerConfig)
	nw.EncryptionConfig = canonicalizeNewEnvironmentConfigEncryptionConfig(c, des.EncryptionConfig, nw.EncryptionConfig)
	if dcl.StringCanonicalize(des.AirflowUri, nw.AirflowUri) {
		nw.AirflowUri = des.AirflowUri
	}

	return nw
}

func canonicalizeNewEnvironmentConfigSet(c *Client, des, nw []EnvironmentConfig) []EnvironmentConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigSlice(c *Client, des, nw []EnvironmentConfig) []EnvironmentConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigSoftwareConfig(des, initial *EnvironmentConfigSoftwareConfig, opts ...dcl.ApplyOption) *EnvironmentConfigSoftwareConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigSoftwareConfig{}

	if dcl.StringCanonicalize(des.ImageVersion, initial.ImageVersion) || dcl.IsZeroValue(des.ImageVersion) {
		cDes.ImageVersion = initial.ImageVersion
	} else {
		cDes.ImageVersion = des.ImageVersion
	}
	if dcl.IsZeroValue(des.AirflowConfigOverrides) {
		des.AirflowConfigOverrides = initial.AirflowConfigOverrides
	} else {
		cDes.AirflowConfigOverrides = des.AirflowConfigOverrides
	}
	if dcl.IsZeroValue(des.PypiPackages) {
		des.PypiPackages = initial.PypiPackages
	} else {
		cDes.PypiPackages = des.PypiPackages
	}
	if dcl.IsZeroValue(des.EnvVariables) {
		des.EnvVariables = initial.EnvVariables
	} else {
		cDes.EnvVariables = des.EnvVariables
	}
	if dcl.StringCanonicalize(des.PythonVersion, initial.PythonVersion) || dcl.IsZeroValue(des.PythonVersion) {
		cDes.PythonVersion = initial.PythonVersion
	} else {
		cDes.PythonVersion = des.PythonVersion
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigSoftwareConfig(c *Client, des, nw *EnvironmentConfigSoftwareConfig) *EnvironmentConfigSoftwareConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ImageVersion, nw.ImageVersion) {
		nw.ImageVersion = des.ImageVersion
	}
	if dcl.IsZeroValue(nw.AirflowConfigOverrides) {
		nw.AirflowConfigOverrides = des.AirflowConfigOverrides
	}
	if dcl.IsZeroValue(nw.PypiPackages) {
		nw.PypiPackages = des.PypiPackages
	}
	if dcl.IsZeroValue(nw.EnvVariables) {
		nw.EnvVariables = des.EnvVariables
	}
	if dcl.StringCanonicalize(des.PythonVersion, nw.PythonVersion) {
		nw.PythonVersion = des.PythonVersion
	}

	return nw
}

func canonicalizeNewEnvironmentConfigSoftwareConfigSet(c *Client, des, nw []EnvironmentConfigSoftwareConfig) []EnvironmentConfigSoftwareConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigSoftwareConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigSoftwareConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigSoftwareConfigSlice(c *Client, des, nw []EnvironmentConfigSoftwareConfig) []EnvironmentConfigSoftwareConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigSoftwareConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigSoftwareConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigNodeConfig(des, initial *EnvironmentConfigNodeConfig, opts ...dcl.ApplyOption) *EnvironmentConfigNodeConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigNodeConfig{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
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
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	} else {
		cDes.DiskSizeGb = des.DiskSizeGb
	}
	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	} else {
		cDes.OAuthScopes = des.OAuthScopes
	}
	if dcl.NameToSelfLink(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	cDes.IPAllocationPolicy = canonicalizeEnvironmentConfigNodeConfigIPAllocationPolicy(des.IPAllocationPolicy, initial.IPAllocationPolicy, opts...)

	return cDes
}

func canonicalizeNewEnvironmentConfigNodeConfig(c *Client, des, nw *EnvironmentConfigNodeConfig) *EnvironmentConfigNodeConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.NameToSelfLink(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, nw.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}
	if dcl.IsZeroValue(nw.DiskSizeGb) {
		nw.DiskSizeGb = des.DiskSizeGb
	}
	if dcl.IsZeroValue(nw.OAuthScopes) {
		nw.OAuthScopes = des.OAuthScopes
	}
	if dcl.NameToSelfLink(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.IsZeroValue(nw.Tags) {
		nw.Tags = des.Tags
	}
	nw.IPAllocationPolicy = canonicalizeNewEnvironmentConfigNodeConfigIPAllocationPolicy(c, des.IPAllocationPolicy, nw.IPAllocationPolicy)

	return nw
}

func canonicalizeNewEnvironmentConfigNodeConfigSet(c *Client, des, nw []EnvironmentConfigNodeConfig) []EnvironmentConfigNodeConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigNodeConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigNodeConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigNodeConfigSlice(c *Client, des, nw []EnvironmentConfigNodeConfig) []EnvironmentConfigNodeConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigNodeConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigNodeConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigNodeConfigIPAllocationPolicy(des, initial *EnvironmentConfigNodeConfigIPAllocationPolicy, opts ...dcl.ApplyOption) *EnvironmentConfigNodeConfigIPAllocationPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.ClusterSecondaryRangeName != nil || (initial != nil && initial.ClusterSecondaryRangeName != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ClusterIPv4CidrBlock) {
			des.ClusterSecondaryRangeName = nil
			if initial != nil {
				initial.ClusterSecondaryRangeName = nil
			}
		}
	}

	if des.ClusterIPv4CidrBlock != nil || (initial != nil && initial.ClusterIPv4CidrBlock != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ClusterSecondaryRangeName) {
			des.ClusterIPv4CidrBlock = nil
			if initial != nil {
				initial.ClusterIPv4CidrBlock = nil
			}
		}
	}

	if des.ServicesSecondaryRangeName != nil || (initial != nil && initial.ServicesSecondaryRangeName != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ServicesIPv4CidrBlock) {
			des.ServicesSecondaryRangeName = nil
			if initial != nil {
				initial.ServicesSecondaryRangeName = nil
			}
		}
	}

	if des.ServicesIPv4CidrBlock != nil || (initial != nil && initial.ServicesIPv4CidrBlock != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ServicesSecondaryRangeName) {
			des.ServicesIPv4CidrBlock = nil
			if initial != nil {
				initial.ServicesIPv4CidrBlock = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigNodeConfigIPAllocationPolicy{}

	if dcl.BoolCanonicalize(des.UseIPAliases, initial.UseIPAliases) || dcl.IsZeroValue(des.UseIPAliases) {
		cDes.UseIPAliases = initial.UseIPAliases
	} else {
		cDes.UseIPAliases = des.UseIPAliases
	}
	if dcl.StringCanonicalize(des.ClusterSecondaryRangeName, initial.ClusterSecondaryRangeName) || dcl.IsZeroValue(des.ClusterSecondaryRangeName) {
		cDes.ClusterSecondaryRangeName = initial.ClusterSecondaryRangeName
	} else {
		cDes.ClusterSecondaryRangeName = des.ClusterSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ClusterIPv4CidrBlock, initial.ClusterIPv4CidrBlock) || dcl.IsZeroValue(des.ClusterIPv4CidrBlock) {
		cDes.ClusterIPv4CidrBlock = initial.ClusterIPv4CidrBlock
	} else {
		cDes.ClusterIPv4CidrBlock = des.ClusterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesSecondaryRangeName, initial.ServicesSecondaryRangeName) || dcl.IsZeroValue(des.ServicesSecondaryRangeName) {
		cDes.ServicesSecondaryRangeName = initial.ServicesSecondaryRangeName
	} else {
		cDes.ServicesSecondaryRangeName = des.ServicesSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ServicesIPv4CidrBlock, initial.ServicesIPv4CidrBlock) || dcl.IsZeroValue(des.ServicesIPv4CidrBlock) {
		cDes.ServicesIPv4CidrBlock = initial.ServicesIPv4CidrBlock
	} else {
		cDes.ServicesIPv4CidrBlock = des.ServicesIPv4CidrBlock
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigNodeConfigIPAllocationPolicy(c *Client, des, nw *EnvironmentConfigNodeConfigIPAllocationPolicy) *EnvironmentConfigNodeConfigIPAllocationPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.UseIPAliases, nw.UseIPAliases) {
		nw.UseIPAliases = des.UseIPAliases
	}
	if dcl.StringCanonicalize(des.ClusterSecondaryRangeName, nw.ClusterSecondaryRangeName) {
		nw.ClusterSecondaryRangeName = des.ClusterSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ClusterIPv4CidrBlock, nw.ClusterIPv4CidrBlock) {
		nw.ClusterIPv4CidrBlock = des.ClusterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesSecondaryRangeName, nw.ServicesSecondaryRangeName) {
		nw.ServicesSecondaryRangeName = des.ServicesSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ServicesIPv4CidrBlock, nw.ServicesIPv4CidrBlock) {
		nw.ServicesIPv4CidrBlock = des.ServicesIPv4CidrBlock
	}

	return nw
}

func canonicalizeNewEnvironmentConfigNodeConfigIPAllocationPolicySet(c *Client, des, nw []EnvironmentConfigNodeConfigIPAllocationPolicy) []EnvironmentConfigNodeConfigIPAllocationPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigNodeConfigIPAllocationPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigNodeConfigIPAllocationPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigNodeConfigIPAllocationPolicySlice(c *Client, des, nw []EnvironmentConfigNodeConfigIPAllocationPolicy) []EnvironmentConfigNodeConfigIPAllocationPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigNodeConfigIPAllocationPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigNodeConfigIPAllocationPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigPrivateEnvironmentConfig(des, initial *EnvironmentConfigPrivateEnvironmentConfig, opts ...dcl.ApplyOption) *EnvironmentConfigPrivateEnvironmentConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigPrivateEnvironmentConfig{}

	if dcl.BoolCanonicalize(des.EnablePrivateEnvironment, initial.EnablePrivateEnvironment) || dcl.IsZeroValue(des.EnablePrivateEnvironment) {
		cDes.EnablePrivateEnvironment = initial.EnablePrivateEnvironment
	} else {
		cDes.EnablePrivateEnvironment = des.EnablePrivateEnvironment
	}
	cDes.PrivateClusterConfig = canonicalizeEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(des.PrivateClusterConfig, initial.PrivateClusterConfig, opts...)
	if dcl.StringCanonicalize(des.WebServerIPv4CidrBlock, initial.WebServerIPv4CidrBlock) || dcl.IsZeroValue(des.WebServerIPv4CidrBlock) {
		cDes.WebServerIPv4CidrBlock = initial.WebServerIPv4CidrBlock
	} else {
		cDes.WebServerIPv4CidrBlock = des.WebServerIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.CloudSqlIPv4CidrBlock, initial.CloudSqlIPv4CidrBlock) || dcl.IsZeroValue(des.CloudSqlIPv4CidrBlock) {
		cDes.CloudSqlIPv4CidrBlock = initial.CloudSqlIPv4CidrBlock
	} else {
		cDes.CloudSqlIPv4CidrBlock = des.CloudSqlIPv4CidrBlock
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfig(c *Client, des, nw *EnvironmentConfigPrivateEnvironmentConfig) *EnvironmentConfigPrivateEnvironmentConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnablePrivateEnvironment, nw.EnablePrivateEnvironment) {
		nw.EnablePrivateEnvironment = des.EnablePrivateEnvironment
	}
	nw.PrivateClusterConfig = canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, des.PrivateClusterConfig, nw.PrivateClusterConfig)
	if dcl.StringCanonicalize(des.WebServerIPv4CidrBlock, nw.WebServerIPv4CidrBlock) {
		nw.WebServerIPv4CidrBlock = des.WebServerIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.CloudSqlIPv4CidrBlock, nw.CloudSqlIPv4CidrBlock) {
		nw.CloudSqlIPv4CidrBlock = des.CloudSqlIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.WebServerIPv4ReservedRange, nw.WebServerIPv4ReservedRange) {
		nw.WebServerIPv4ReservedRange = des.WebServerIPv4ReservedRange
	}

	return nw
}

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigSet(c *Client, des, nw []EnvironmentConfigPrivateEnvironmentConfig) []EnvironmentConfigPrivateEnvironmentConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigPrivateEnvironmentConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigPrivateEnvironmentConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigSlice(c *Client, des, nw []EnvironmentConfigPrivateEnvironmentConfig) []EnvironmentConfigPrivateEnvironmentConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigPrivateEnvironmentConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigPrivateEnvironmentConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(des, initial *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, opts ...dcl.ApplyOption) *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}

	if dcl.BoolCanonicalize(des.EnablePrivateEndpoint, initial.EnablePrivateEndpoint) || dcl.IsZeroValue(des.EnablePrivateEndpoint) {
		cDes.EnablePrivateEndpoint = initial.EnablePrivateEndpoint
	} else {
		cDes.EnablePrivateEndpoint = des.EnablePrivateEndpoint
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, initial.MasterIPv4CidrBlock) || dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		cDes.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	} else {
		cDes.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c *Client, des, nw *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnablePrivateEndpoint, nw.EnablePrivateEndpoint) {
		nw.EnablePrivateEndpoint = des.EnablePrivateEndpoint
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, nw.MasterIPv4CidrBlock) {
		nw.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.MasterIPv4ReservedRange, nw.MasterIPv4ReservedRange) {
		nw.MasterIPv4ReservedRange = des.MasterIPv4ReservedRange
	}

	return nw
}

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSet(c *Client, des, nw []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSlice(c *Client, des, nw []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigWebServerNetworkAccessControl(des, initial *EnvironmentConfigWebServerNetworkAccessControl, opts ...dcl.ApplyOption) *EnvironmentConfigWebServerNetworkAccessControl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigWebServerNetworkAccessControl{}

	if dcl.IsZeroValue(des.AllowedIPRanges) {
		des.AllowedIPRanges = initial.AllowedIPRanges
	} else {
		cDes.AllowedIPRanges = des.AllowedIPRanges
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControl(c *Client, des, nw *EnvironmentConfigWebServerNetworkAccessControl) *EnvironmentConfigWebServerNetworkAccessControl {
	if des == nil || nw == nil {
		return nw
	}

	nw.AllowedIPRanges = canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c, des.AllowedIPRanges, nw.AllowedIPRanges)

	return nw
}

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlSet(c *Client, des, nw []EnvironmentConfigWebServerNetworkAccessControl) []EnvironmentConfigWebServerNetworkAccessControl {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigWebServerNetworkAccessControl
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigWebServerNetworkAccessControlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlSlice(c *Client, des, nw []EnvironmentConfigWebServerNetworkAccessControl) []EnvironmentConfigWebServerNetworkAccessControl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigWebServerNetworkAccessControl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigWebServerNetworkAccessControl(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(des, initial *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, opts ...dcl.ApplyOption) *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}

	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c *Client, des, nw *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSet(c *Client, des, nw []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c *Client, des, nw []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigDatabaseConfig(des, initial *EnvironmentConfigDatabaseConfig, opts ...dcl.ApplyOption) *EnvironmentConfigDatabaseConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigDatabaseConfig{}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigDatabaseConfig(c *Client, des, nw *EnvironmentConfigDatabaseConfig) *EnvironmentConfigDatabaseConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}

	return nw
}

func canonicalizeNewEnvironmentConfigDatabaseConfigSet(c *Client, des, nw []EnvironmentConfigDatabaseConfig) []EnvironmentConfigDatabaseConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigDatabaseConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigDatabaseConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigDatabaseConfigSlice(c *Client, des, nw []EnvironmentConfigDatabaseConfig) []EnvironmentConfigDatabaseConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigDatabaseConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigDatabaseConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigWebServerConfig(des, initial *EnvironmentConfigWebServerConfig, opts ...dcl.ApplyOption) *EnvironmentConfigWebServerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigWebServerConfig{}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigWebServerConfig(c *Client, des, nw *EnvironmentConfigWebServerConfig) *EnvironmentConfigWebServerConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}

	return nw
}

func canonicalizeNewEnvironmentConfigWebServerConfigSet(c *Client, des, nw []EnvironmentConfigWebServerConfig) []EnvironmentConfigWebServerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigWebServerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigWebServerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigWebServerConfigSlice(c *Client, des, nw []EnvironmentConfigWebServerConfig) []EnvironmentConfigWebServerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigWebServerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigWebServerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentConfigEncryptionConfig(des, initial *EnvironmentConfigEncryptionConfig, opts ...dcl.ApplyOption) *EnvironmentConfigEncryptionConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EnvironmentConfigEncryptionConfig{}

	if dcl.NameToSelfLink(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeNewEnvironmentConfigEncryptionConfig(c *Client, des, nw *EnvironmentConfigEncryptionConfig) *EnvironmentConfigEncryptionConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}

	return nw
}

func canonicalizeNewEnvironmentConfigEncryptionConfigSet(c *Client, des, nw []EnvironmentConfigEncryptionConfig) []EnvironmentConfigEncryptionConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentConfigEncryptionConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentConfigEncryptionConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentConfigEncryptionConfigSlice(c *Client, des, nw []EnvironmentConfigEncryptionConfig) []EnvironmentConfigEncryptionConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentConfigEncryptionConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentConfigEncryptionConfig(c, &d, &n))
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
func diffEnvironment(c *Client, desired, actual *Environment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareEnvironmentConfigNewStyle, EmptyObject: EmptyEnvironmentConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uuid, actual.Uuid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uuid")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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
func compareEnvironmentConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfig or *EnvironmentConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GkeCluster, actual.GkeCluster, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GkeCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DagGcsPrefix, actual.DagGcsPrefix, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DagGcsPrefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SoftwareConfig, actual.SoftwareConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigSoftwareConfigNewStyle, EmptyObject: EmptyEnvironmentConfigSoftwareConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SoftwareConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeConfig, actual.NodeConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigNodeConfigNewStyle, EmptyObject: EmptyEnvironmentConfigNodeConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateEnvironmentConfig, actual.PrivateEnvironmentConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigPrivateEnvironmentConfigNewStyle, EmptyObject: EmptyEnvironmentConfigPrivateEnvironmentConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateEnvironmentConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WebServerNetworkAccessControl, actual.WebServerNetworkAccessControl, dcl.Info{ObjectFunction: compareEnvironmentConfigWebServerNetworkAccessControlNewStyle, EmptyObject: EmptyEnvironmentConfigWebServerNetworkAccessControl, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WebServerNetworkAccessControl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatabaseConfig, actual.DatabaseConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigDatabaseConfigNewStyle, EmptyObject: EmptyEnvironmentConfigDatabaseConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatabaseConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WebServerConfig, actual.WebServerConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigWebServerConfigNewStyle, EmptyObject: EmptyEnvironmentConfigWebServerConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WebServerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionConfig, actual.EncryptionConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigEncryptionConfigNewStyle, EmptyObject: EmptyEnvironmentConfigEncryptionConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AirflowUri, actual.AirflowUri, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AirflowUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigSoftwareConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigSoftwareConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigSoftwareConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigSoftwareConfig or *EnvironmentConfigSoftwareConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigSoftwareConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigSoftwareConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigSoftwareConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ImageVersion, actual.ImageVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AirflowConfigOverrides, actual.AirflowConfigOverrides, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AirflowConfigOverrides")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PypiPackages, actual.PypiPackages, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PypiPackages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnvVariables, actual.EnvVariables, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnvVariables")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PythonVersion, actual.PythonVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PythonVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigNodeConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigNodeConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigNodeConfig or *EnvironmentConfigNodeConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigNodeConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigNodeConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subnetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuthScopes, actual.OAuthScopes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OauthScopes")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAllocationPolicy, actual.IPAllocationPolicy, dcl.Info{ObjectFunction: compareEnvironmentConfigNodeConfigIPAllocationPolicyNewStyle, EmptyObject: EmptyEnvironmentConfigNodeConfigIPAllocationPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAllocationPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigNodeConfigIPAllocationPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigNodeConfigIPAllocationPolicy)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigNodeConfigIPAllocationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigNodeConfigIPAllocationPolicy or *EnvironmentConfigNodeConfigIPAllocationPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigNodeConfigIPAllocationPolicy)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigNodeConfigIPAllocationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigNodeConfigIPAllocationPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseIPAliases, actual.UseIPAliases, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UseIpAliases")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterSecondaryRangeName, actual.ClusterSecondaryRangeName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterSecondaryRangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterIPv4CidrBlock, actual.ClusterIPv4CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesSecondaryRangeName, actual.ServicesSecondaryRangeName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesSecondaryRangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesIPv4CidrBlock, actual.ServicesIPv4CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigPrivateEnvironmentConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigPrivateEnvironmentConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigPrivateEnvironmentConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigPrivateEnvironmentConfig or *EnvironmentConfigPrivateEnvironmentConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigPrivateEnvironmentConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigPrivateEnvironmentConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigPrivateEnvironmentConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnablePrivateEnvironment, actual.EnablePrivateEnvironment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnablePrivateEnvironment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateClusterConfig, actual.PrivateClusterConfig, dcl.Info{ObjectFunction: compareEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigNewStyle, EmptyObject: EmptyEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateClusterConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WebServerIPv4CidrBlock, actual.WebServerIPv4CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WebServerIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudSqlIPv4CidrBlock, actual.CloudSqlIPv4CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudSqlIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WebServerIPv4ReservedRange, actual.WebServerIPv4ReservedRange, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WebServerIpv4ReservedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig or *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnablePrivateEndpoint, actual.EnablePrivateEndpoint, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnablePrivateEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MasterIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4ReservedRange, actual.MasterIPv4ReservedRange, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MasterIpv4ReservedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigWebServerNetworkAccessControlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigWebServerNetworkAccessControl)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigWebServerNetworkAccessControl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerNetworkAccessControl or *EnvironmentConfigWebServerNetworkAccessControl", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigWebServerNetworkAccessControl)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigWebServerNetworkAccessControl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerNetworkAccessControl", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedIPRanges, actual.AllowedIPRanges, dcl.Info{ObjectFunction: compareEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesNewStyle, EmptyObject: EmptyEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedIpRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges or *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigDatabaseConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigDatabaseConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigDatabaseConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigDatabaseConfig or *EnvironmentConfigDatabaseConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigDatabaseConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigDatabaseConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigDatabaseConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigWebServerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigWebServerConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigWebServerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerConfig or *EnvironmentConfigWebServerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigWebServerConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigWebServerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigWebServerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentConfigEncryptionConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentConfigEncryptionConfig)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentConfigEncryptionConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigEncryptionConfig or *EnvironmentConfigEncryptionConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentConfigEncryptionConfig)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentConfigEncryptionConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentConfigEncryptionConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Environment) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Environment) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Environment) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Environment) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Environment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Environment) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvironment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Environment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvironment decodes JSON responses into the Environment resource schema.
func unmarshalEnvironment(b []byte, c *Client) (*Environment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvironment(m, c)
}

func unmarshalMapEnvironment(m map[string]interface{}, c *Client) (*Environment, error) {

	return flattenEnvironment(c, m), nil
}

// expandEnvironment expands Environment into a JSON request object.
func expandEnvironment(c *Client, f *Environment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	// Alias full resource as res to distinguish it from nested objects.
	res := f
	if v, err := dcl.DeriveField("projects/%s/locations/%s/environments/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := expandEnvironmentConfig(c, f.Config, res); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if v != nil {
		m["config"] = v
	}
	if v := f.Uuid; !dcl.IsEmptyValueIndirect(v) {
		m["uuid"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
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

// flattenEnvironment flattens Environment from a JSON request object into the
// Environment type.
func flattenEnvironment(c *Client, i interface{}) *Environment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Environment{}
	res.Name = dcl.FlattenString(m["name"])
	res.Config = flattenEnvironmentConfig(c, m["config"])
	res.Uuid = dcl.FlattenString(m["uuid"])
	res.State = flattenEnvironmentStateEnum(m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandEnvironmentConfigMap expands the contents of EnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfigMap(c *Client, f map[string]EnvironmentConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigSlice expands the contents of EnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfigSlice(c *Client, f []EnvironmentConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigMap flattens the contents of EnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfigMap(c *Client, i interface{}) map[string]EnvironmentConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfig{}
	}

	items := make(map[string]EnvironmentConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigSlice flattens the contents of EnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfigSlice(c *Client, i interface{}) []EnvironmentConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfig{}
	}

	items := make([]EnvironmentConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfig expands an instance of EnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfig(c *Client, f *EnvironmentConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.GkeCluster; !dcl.IsEmptyValueIndirect(v) {
		m["gkeCluster"] = v
	}
	if v := f.DagGcsPrefix; !dcl.IsEmptyValueIndirect(v) {
		m["dagGcsPrefix"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["nodeCount"] = v
	}
	if v, err := expandEnvironmentConfigSoftwareConfig(c, f.SoftwareConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SoftwareConfig into softwareConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["softwareConfig"] = v
	}
	if v, err := expandEnvironmentConfigNodeConfig(c, f.NodeConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NodeConfig into nodeConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nodeConfig"] = v
	}
	if v, err := expandEnvironmentConfigPrivateEnvironmentConfig(c, f.PrivateEnvironmentConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding PrivateEnvironmentConfig into privateEnvironmentConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["privateEnvironmentConfig"] = v
	}
	if v, err := expandEnvironmentConfigWebServerNetworkAccessControl(c, f.WebServerNetworkAccessControl, res); err != nil {
		return nil, fmt.Errorf("error expanding WebServerNetworkAccessControl into webServerNetworkAccessControl: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["webServerNetworkAccessControl"] = v
	}
	if v, err := expandEnvironmentConfigDatabaseConfig(c, f.DatabaseConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseConfig into databaseConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseConfig"] = v
	}
	if v, err := expandEnvironmentConfigWebServerConfig(c, f.WebServerConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding WebServerConfig into webServerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["webServerConfig"] = v
	}
	if v, err := expandEnvironmentConfigEncryptionConfig(c, f.EncryptionConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionConfig into encryptionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionConfig"] = v
	}
	if v := f.AirflowUri; !dcl.IsEmptyValueIndirect(v) {
		m["airflowUri"] = v
	}

	return m, nil
}

// flattenEnvironmentConfig flattens an instance of EnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfig(c *Client, i interface{}) *EnvironmentConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfig
	}
	r.GkeCluster = dcl.FlattenString(m["gkeCluster"])
	r.DagGcsPrefix = dcl.FlattenString(m["dagGcsPrefix"])
	r.NodeCount = dcl.FlattenInteger(m["nodeCount"])
	r.SoftwareConfig = flattenEnvironmentConfigSoftwareConfig(c, m["softwareConfig"])
	r.NodeConfig = flattenEnvironmentConfigNodeConfig(c, m["nodeConfig"])
	r.PrivateEnvironmentConfig = flattenEnvironmentConfigPrivateEnvironmentConfig(c, m["privateEnvironmentConfig"])
	r.WebServerNetworkAccessControl = flattenEnvironmentConfigWebServerNetworkAccessControl(c, m["webServerNetworkAccessControl"])
	r.DatabaseConfig = flattenEnvironmentConfigDatabaseConfig(c, m["databaseConfig"])
	r.WebServerConfig = flattenEnvironmentConfigWebServerConfig(c, m["webServerConfig"])
	r.EncryptionConfig = flattenEnvironmentConfigEncryptionConfig(c, m["encryptionConfig"])
	r.AirflowUri = dcl.FlattenString(m["airflowUri"])

	return r
}

// expandEnvironmentConfigSoftwareConfigMap expands the contents of EnvironmentConfigSoftwareConfig into a JSON
// request object.
func expandEnvironmentConfigSoftwareConfigMap(c *Client, f map[string]EnvironmentConfigSoftwareConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigSoftwareConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigSoftwareConfigSlice expands the contents of EnvironmentConfigSoftwareConfig into a JSON
// request object.
func expandEnvironmentConfigSoftwareConfigSlice(c *Client, f []EnvironmentConfigSoftwareConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigSoftwareConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigSoftwareConfigMap flattens the contents of EnvironmentConfigSoftwareConfig from a JSON
// response object.
func flattenEnvironmentConfigSoftwareConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigSoftwareConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigSoftwareConfig{}
	}

	items := make(map[string]EnvironmentConfigSoftwareConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigSoftwareConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigSoftwareConfigSlice flattens the contents of EnvironmentConfigSoftwareConfig from a JSON
// response object.
func flattenEnvironmentConfigSoftwareConfigSlice(c *Client, i interface{}) []EnvironmentConfigSoftwareConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigSoftwareConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigSoftwareConfig{}
	}

	items := make([]EnvironmentConfigSoftwareConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigSoftwareConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigSoftwareConfig expands an instance of EnvironmentConfigSoftwareConfig into a JSON
// request object.
func expandEnvironmentConfigSoftwareConfig(c *Client, f *EnvironmentConfigSoftwareConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ImageVersion; !dcl.IsEmptyValueIndirect(v) {
		m["imageVersion"] = v
	}
	if v := f.AirflowConfigOverrides; !dcl.IsEmptyValueIndirect(v) {
		m["airflowConfigOverrides"] = v
	}
	if v := f.PypiPackages; !dcl.IsEmptyValueIndirect(v) {
		m["pypiPackages"] = v
	}
	if v := f.EnvVariables; !dcl.IsEmptyValueIndirect(v) {
		m["envVariables"] = v
	}
	if v := f.PythonVersion; !dcl.IsEmptyValueIndirect(v) {
		m["pythonVersion"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigSoftwareConfig flattens an instance of EnvironmentConfigSoftwareConfig from a JSON
// response object.
func flattenEnvironmentConfigSoftwareConfig(c *Client, i interface{}) *EnvironmentConfigSoftwareConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigSoftwareConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigSoftwareConfig
	}
	r.ImageVersion = dcl.FlattenString(m["imageVersion"])
	r.AirflowConfigOverrides = dcl.FlattenKeyValuePairs(m["airflowConfigOverrides"])
	r.PypiPackages = dcl.FlattenKeyValuePairs(m["pypiPackages"])
	r.EnvVariables = dcl.FlattenKeyValuePairs(m["envVariables"])
	r.PythonVersion = dcl.FlattenString(m["pythonVersion"])

	return r
}

// expandEnvironmentConfigNodeConfigMap expands the contents of EnvironmentConfigNodeConfig into a JSON
// request object.
func expandEnvironmentConfigNodeConfigMap(c *Client, f map[string]EnvironmentConfigNodeConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigNodeConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigNodeConfigSlice expands the contents of EnvironmentConfigNodeConfig into a JSON
// request object.
func expandEnvironmentConfigNodeConfigSlice(c *Client, f []EnvironmentConfigNodeConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigNodeConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigNodeConfigMap flattens the contents of EnvironmentConfigNodeConfig from a JSON
// response object.
func flattenEnvironmentConfigNodeConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigNodeConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigNodeConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigNodeConfig{}
	}

	items := make(map[string]EnvironmentConfigNodeConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigNodeConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigNodeConfigSlice flattens the contents of EnvironmentConfigNodeConfig from a JSON
// response object.
func flattenEnvironmentConfigNodeConfigSlice(c *Client, i interface{}) []EnvironmentConfigNodeConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigNodeConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigNodeConfig{}
	}

	items := make([]EnvironmentConfigNodeConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigNodeConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigNodeConfig expands an instance of EnvironmentConfigNodeConfig into a JSON
// request object.
func expandEnvironmentConfigNodeConfig(c *Client, f *EnvironmentConfigNodeConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandComposerEnvironmentLocation(res); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := expandComposerEnvironmentMachineType(res); err != nil {
		return nil, fmt.Errorf("error expanding MachineType into machineType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v, err := expandComposerEnvironmentNetwork(res); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v, err := expandComposerEnvironmentSubnetwork(res); err != nil {
		return nil, fmt.Errorf("error expanding Subnetwork into subnetwork: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.OAuthScopes; !dcl.IsEmptyValueIndirect(v) {
		m["oauthScopes"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v, err := expandEnvironmentConfigNodeConfigIPAllocationPolicy(c, f.IPAllocationPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding IPAllocationPolicy into ipAllocationPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ipAllocationPolicy"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigNodeConfig flattens an instance of EnvironmentConfigNodeConfig from a JSON
// response object.
func flattenEnvironmentConfigNodeConfig(c *Client, i interface{}) *EnvironmentConfigNodeConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigNodeConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigNodeConfig
	}
	r.Location = dcl.FlattenString(m["location"])
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.Network = dcl.FlattenString(m["network"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.OAuthScopes = dcl.FlattenStringSlice(m["oauthScopes"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.IPAllocationPolicy = flattenEnvironmentConfigNodeConfigIPAllocationPolicy(c, m["ipAllocationPolicy"])

	return r
}

// expandEnvironmentConfigNodeConfigIPAllocationPolicyMap expands the contents of EnvironmentConfigNodeConfigIPAllocationPolicy into a JSON
// request object.
func expandEnvironmentConfigNodeConfigIPAllocationPolicyMap(c *Client, f map[string]EnvironmentConfigNodeConfigIPAllocationPolicy, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigNodeConfigIPAllocationPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigNodeConfigIPAllocationPolicySlice expands the contents of EnvironmentConfigNodeConfigIPAllocationPolicy into a JSON
// request object.
func expandEnvironmentConfigNodeConfigIPAllocationPolicySlice(c *Client, f []EnvironmentConfigNodeConfigIPAllocationPolicy, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigNodeConfigIPAllocationPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigNodeConfigIPAllocationPolicyMap flattens the contents of EnvironmentConfigNodeConfigIPAllocationPolicy from a JSON
// response object.
func flattenEnvironmentConfigNodeConfigIPAllocationPolicyMap(c *Client, i interface{}) map[string]EnvironmentConfigNodeConfigIPAllocationPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigNodeConfigIPAllocationPolicy{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigNodeConfigIPAllocationPolicy{}
	}

	items := make(map[string]EnvironmentConfigNodeConfigIPAllocationPolicy)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigNodeConfigIPAllocationPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigNodeConfigIPAllocationPolicySlice flattens the contents of EnvironmentConfigNodeConfigIPAllocationPolicy from a JSON
// response object.
func flattenEnvironmentConfigNodeConfigIPAllocationPolicySlice(c *Client, i interface{}) []EnvironmentConfigNodeConfigIPAllocationPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigNodeConfigIPAllocationPolicy{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigNodeConfigIPAllocationPolicy{}
	}

	items := make([]EnvironmentConfigNodeConfigIPAllocationPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigNodeConfigIPAllocationPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigNodeConfigIPAllocationPolicy expands an instance of EnvironmentConfigNodeConfigIPAllocationPolicy into a JSON
// request object.
func expandEnvironmentConfigNodeConfigIPAllocationPolicy(c *Client, f *EnvironmentConfigNodeConfigIPAllocationPolicy, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseIPAliases; !dcl.IsEmptyValueIndirect(v) {
		m["useIpAliases"] = v
	}
	if v := f.ClusterSecondaryRangeName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterSecondaryRangeName"] = v
	}
	if v := f.ClusterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["clusterIpv4CidrBlock"] = v
	}
	if v := f.ServicesSecondaryRangeName; !dcl.IsEmptyValueIndirect(v) {
		m["servicesSecondaryRangeName"] = v
	}
	if v := f.ServicesIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIpv4CidrBlock"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigNodeConfigIPAllocationPolicy flattens an instance of EnvironmentConfigNodeConfigIPAllocationPolicy from a JSON
// response object.
func flattenEnvironmentConfigNodeConfigIPAllocationPolicy(c *Client, i interface{}) *EnvironmentConfigNodeConfigIPAllocationPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigNodeConfigIPAllocationPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigNodeConfigIPAllocationPolicy
	}
	r.UseIPAliases = dcl.FlattenBool(m["useIpAliases"])
	r.ClusterSecondaryRangeName = dcl.FlattenString(m["clusterSecondaryRangeName"])
	r.ClusterIPv4CidrBlock = dcl.FlattenString(m["clusterIpv4CidrBlock"])
	r.ServicesSecondaryRangeName = dcl.FlattenString(m["servicesSecondaryRangeName"])
	r.ServicesIPv4CidrBlock = dcl.FlattenString(m["servicesIpv4CidrBlock"])

	return r
}

// expandEnvironmentConfigPrivateEnvironmentConfigMap expands the contents of EnvironmentConfigPrivateEnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfigMap(c *Client, f map[string]EnvironmentConfigPrivateEnvironmentConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigPrivateEnvironmentConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigPrivateEnvironmentConfigSlice expands the contents of EnvironmentConfigPrivateEnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfigSlice(c *Client, f []EnvironmentConfigPrivateEnvironmentConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigPrivateEnvironmentConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigPrivateEnvironmentConfigMap flattens the contents of EnvironmentConfigPrivateEnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigPrivateEnvironmentConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigPrivateEnvironmentConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigPrivateEnvironmentConfig{}
	}

	items := make(map[string]EnvironmentConfigPrivateEnvironmentConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigPrivateEnvironmentConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigPrivateEnvironmentConfigSlice flattens the contents of EnvironmentConfigPrivateEnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfigSlice(c *Client, i interface{}) []EnvironmentConfigPrivateEnvironmentConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigPrivateEnvironmentConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigPrivateEnvironmentConfig{}
	}

	items := make([]EnvironmentConfigPrivateEnvironmentConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigPrivateEnvironmentConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigPrivateEnvironmentConfig expands an instance of EnvironmentConfigPrivateEnvironmentConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfig(c *Client, f *EnvironmentConfigPrivateEnvironmentConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnablePrivateEnvironment; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateEnvironment"] = v
	}
	if v, err := expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, f.PrivateClusterConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding PrivateClusterConfig into privateClusterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["privateClusterConfig"] = v
	}
	if v := f.WebServerIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["webServerIpv4CidrBlock"] = v
	}
	if v := f.CloudSqlIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["cloudSqlIpv4CidrBlock"] = v
	}
	if v := f.WebServerIPv4ReservedRange; !dcl.IsEmptyValueIndirect(v) {
		m["webServerIpv4ReservedRange"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigPrivateEnvironmentConfig flattens an instance of EnvironmentConfigPrivateEnvironmentConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfig(c *Client, i interface{}) *EnvironmentConfigPrivateEnvironmentConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigPrivateEnvironmentConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigPrivateEnvironmentConfig
	}
	r.EnablePrivateEnvironment = dcl.FlattenBool(m["enablePrivateEnvironment"])
	r.PrivateClusterConfig = flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, m["privateClusterConfig"])
	r.WebServerIPv4CidrBlock = dcl.FlattenString(m["webServerIpv4CidrBlock"])
	r.CloudSqlIPv4CidrBlock = dcl.FlattenString(m["cloudSqlIpv4CidrBlock"])
	r.WebServerIPv4ReservedRange = dcl.FlattenString(m["webServerIpv4ReservedRange"])

	return r
}

// expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigMap expands the contents of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigMap(c *Client, f map[string]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSlice expands the contents of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSlice(c *Client, f []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigMap flattens the contents of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}
	}

	items := make(map[string]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSlice flattens the contents of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigSlice(c *Client, i interface{}) []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}
	}

	items := make([]EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig expands an instance of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig into a JSON
// request object.
func expandEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c *Client, f *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnablePrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateEndpoint"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4CidrBlock"] = v
	}
	if v := f.MasterIPv4ReservedRange; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4ReservedRange"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig flattens an instance of EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig from a JSON
// response object.
func flattenEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(c *Client, i interface{}) *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig
	}
	r.EnablePrivateEndpoint = dcl.FlattenBool(m["enablePrivateEndpoint"])
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIpv4CidrBlock"])
	r.MasterIPv4ReservedRange = dcl.FlattenString(m["masterIpv4ReservedRange"])

	return r
}

// expandEnvironmentConfigWebServerNetworkAccessControlMap expands the contents of EnvironmentConfigWebServerNetworkAccessControl into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControlMap(c *Client, f map[string]EnvironmentConfigWebServerNetworkAccessControl, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigWebServerNetworkAccessControl(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigWebServerNetworkAccessControlSlice expands the contents of EnvironmentConfigWebServerNetworkAccessControl into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControlSlice(c *Client, f []EnvironmentConfigWebServerNetworkAccessControl, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigWebServerNetworkAccessControl(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigWebServerNetworkAccessControlMap flattens the contents of EnvironmentConfigWebServerNetworkAccessControl from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControlMap(c *Client, i interface{}) map[string]EnvironmentConfigWebServerNetworkAccessControl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigWebServerNetworkAccessControl{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigWebServerNetworkAccessControl{}
	}

	items := make(map[string]EnvironmentConfigWebServerNetworkAccessControl)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigWebServerNetworkAccessControl(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigWebServerNetworkAccessControlSlice flattens the contents of EnvironmentConfigWebServerNetworkAccessControl from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControlSlice(c *Client, i interface{}) []EnvironmentConfigWebServerNetworkAccessControl {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigWebServerNetworkAccessControl{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigWebServerNetworkAccessControl{}
	}

	items := make([]EnvironmentConfigWebServerNetworkAccessControl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigWebServerNetworkAccessControl(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigWebServerNetworkAccessControl expands an instance of EnvironmentConfigWebServerNetworkAccessControl into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControl(c *Client, f *EnvironmentConfigWebServerNetworkAccessControl, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c, f.AllowedIPRanges, res); err != nil {
		return nil, fmt.Errorf("error expanding AllowedIPRanges into allowedIpRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedIpRanges"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigWebServerNetworkAccessControl flattens an instance of EnvironmentConfigWebServerNetworkAccessControl from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControl(c *Client, i interface{}) *EnvironmentConfigWebServerNetworkAccessControl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigWebServerNetworkAccessControl{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigWebServerNetworkAccessControl
	}
	r.AllowedIPRanges = flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c, m["allowedIpRanges"])

	return r
}

// expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesMap expands the contents of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesMap(c *Client, f map[string]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice expands the contents of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c *Client, f []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesMap flattens the contents of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesMap(c *Client, i interface{}) map[string]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}
	}

	items := make(map[string]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice flattens the contents of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesSlice(c *Client, i interface{}) []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}
	}

	items := make([]EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges expands an instance of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges into a JSON
// request object.
func expandEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c *Client, f *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges flattens an instance of EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges from a JSON
// response object.
func flattenEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(c *Client, i interface{}) *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges
	}
	r.Value = dcl.FlattenString(m["value"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandEnvironmentConfigDatabaseConfigMap expands the contents of EnvironmentConfigDatabaseConfig into a JSON
// request object.
func expandEnvironmentConfigDatabaseConfigMap(c *Client, f map[string]EnvironmentConfigDatabaseConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigDatabaseConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigDatabaseConfigSlice expands the contents of EnvironmentConfigDatabaseConfig into a JSON
// request object.
func expandEnvironmentConfigDatabaseConfigSlice(c *Client, f []EnvironmentConfigDatabaseConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigDatabaseConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigDatabaseConfigMap flattens the contents of EnvironmentConfigDatabaseConfig from a JSON
// response object.
func flattenEnvironmentConfigDatabaseConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigDatabaseConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigDatabaseConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigDatabaseConfig{}
	}

	items := make(map[string]EnvironmentConfigDatabaseConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigDatabaseConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigDatabaseConfigSlice flattens the contents of EnvironmentConfigDatabaseConfig from a JSON
// response object.
func flattenEnvironmentConfigDatabaseConfigSlice(c *Client, i interface{}) []EnvironmentConfigDatabaseConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigDatabaseConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigDatabaseConfig{}
	}

	items := make([]EnvironmentConfigDatabaseConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigDatabaseConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigDatabaseConfig expands an instance of EnvironmentConfigDatabaseConfig into a JSON
// request object.
func expandEnvironmentConfigDatabaseConfig(c *Client, f *EnvironmentConfigDatabaseConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigDatabaseConfig flattens an instance of EnvironmentConfigDatabaseConfig from a JSON
// response object.
func flattenEnvironmentConfigDatabaseConfig(c *Client, i interface{}) *EnvironmentConfigDatabaseConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigDatabaseConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigDatabaseConfig
	}
	r.MachineType = dcl.FlattenString(m["machineType"])

	return r
}

// expandEnvironmentConfigWebServerConfigMap expands the contents of EnvironmentConfigWebServerConfig into a JSON
// request object.
func expandEnvironmentConfigWebServerConfigMap(c *Client, f map[string]EnvironmentConfigWebServerConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigWebServerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigWebServerConfigSlice expands the contents of EnvironmentConfigWebServerConfig into a JSON
// request object.
func expandEnvironmentConfigWebServerConfigSlice(c *Client, f []EnvironmentConfigWebServerConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigWebServerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigWebServerConfigMap flattens the contents of EnvironmentConfigWebServerConfig from a JSON
// response object.
func flattenEnvironmentConfigWebServerConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigWebServerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigWebServerConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigWebServerConfig{}
	}

	items := make(map[string]EnvironmentConfigWebServerConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigWebServerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigWebServerConfigSlice flattens the contents of EnvironmentConfigWebServerConfig from a JSON
// response object.
func flattenEnvironmentConfigWebServerConfigSlice(c *Client, i interface{}) []EnvironmentConfigWebServerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigWebServerConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigWebServerConfig{}
	}

	items := make([]EnvironmentConfigWebServerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigWebServerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigWebServerConfig expands an instance of EnvironmentConfigWebServerConfig into a JSON
// request object.
func expandEnvironmentConfigWebServerConfig(c *Client, f *EnvironmentConfigWebServerConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigWebServerConfig flattens an instance of EnvironmentConfigWebServerConfig from a JSON
// response object.
func flattenEnvironmentConfigWebServerConfig(c *Client, i interface{}) *EnvironmentConfigWebServerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigWebServerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigWebServerConfig
	}
	r.MachineType = dcl.FlattenString(m["machineType"])

	return r
}

// expandEnvironmentConfigEncryptionConfigMap expands the contents of EnvironmentConfigEncryptionConfig into a JSON
// request object.
func expandEnvironmentConfigEncryptionConfigMap(c *Client, f map[string]EnvironmentConfigEncryptionConfig, res *Environment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentConfigEncryptionConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentConfigEncryptionConfigSlice expands the contents of EnvironmentConfigEncryptionConfig into a JSON
// request object.
func expandEnvironmentConfigEncryptionConfigSlice(c *Client, f []EnvironmentConfigEncryptionConfig, res *Environment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentConfigEncryptionConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentConfigEncryptionConfigMap flattens the contents of EnvironmentConfigEncryptionConfig from a JSON
// response object.
func flattenEnvironmentConfigEncryptionConfigMap(c *Client, i interface{}) map[string]EnvironmentConfigEncryptionConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentConfigEncryptionConfig{}
	}

	items := make(map[string]EnvironmentConfigEncryptionConfig)
	for k, item := range a {
		items[k] = *flattenEnvironmentConfigEncryptionConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentConfigEncryptionConfigSlice flattens the contents of EnvironmentConfigEncryptionConfig from a JSON
// response object.
func flattenEnvironmentConfigEncryptionConfigSlice(c *Client, i interface{}) []EnvironmentConfigEncryptionConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentConfigEncryptionConfig{}
	}

	if len(a) == 0 {
		return []EnvironmentConfigEncryptionConfig{}
	}

	items := make([]EnvironmentConfigEncryptionConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentConfigEncryptionConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentConfigEncryptionConfig expands an instance of EnvironmentConfigEncryptionConfig into a JSON
// request object.
func expandEnvironmentConfigEncryptionConfig(c *Client, f *EnvironmentConfigEncryptionConfig, res *Environment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenEnvironmentConfigEncryptionConfig flattens an instance of EnvironmentConfigEncryptionConfig from a JSON
// response object.
func flattenEnvironmentConfigEncryptionConfig(c *Client, i interface{}) *EnvironmentConfigEncryptionConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentConfigEncryptionConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentConfigEncryptionConfig
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// flattenEnvironmentStateEnumMap flattens the contents of EnvironmentStateEnum from a JSON
// response object.
func flattenEnvironmentStateEnumMap(c *Client, i interface{}) map[string]EnvironmentStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentStateEnum{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentStateEnum{}
	}

	items := make(map[string]EnvironmentStateEnum)
	for k, item := range a {
		items[k] = *flattenEnvironmentStateEnum(item.(interface{}))
	}

	return items
}

// flattenEnvironmentStateEnumSlice flattens the contents of EnvironmentStateEnum from a JSON
// response object.
func flattenEnvironmentStateEnumSlice(c *Client, i interface{}) []EnvironmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentStateEnum{}
	}

	if len(a) == 0 {
		return []EnvironmentStateEnum{}
	}

	items := make([]EnvironmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenEnvironmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *EnvironmentStateEnum with the same value as that string.
func flattenEnvironmentStateEnum(i interface{}) *EnvironmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return EnvironmentStateEnumRef("")
	}

	return EnvironmentStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Environment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvironment(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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

type environmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         environmentApiOperation
}

func convertFieldDiffsToEnvironmentDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]environmentDiff, error) {
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
	var diffs []environmentDiff
	// For each operation name, create a environmentDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := environmentDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEnvironmentApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEnvironmentApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (environmentApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
