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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Environment struct {
	Name       *string               `json:"name"`
	Config     *EnvironmentConfig    `json:"config"`
	Uuid       *string               `json:"uuid"`
	State      *EnvironmentStateEnum `json:"state"`
	CreateTime *string               `json:"createTime"`
	UpdateTime *string               `json:"updateTime"`
	Labels     map[string]string     `json:"labels"`
	Project    *string               `json:"project"`
	Location   *string               `json:"location"`
}

func (r *Environment) String() string {
	return dcl.SprintResource(r)
}

// The enum EnvironmentStateEnum.
type EnvironmentStateEnum string

// EnvironmentStateEnumRef returns a *EnvironmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func EnvironmentStateEnumRef(s string) *EnvironmentStateEnum {
	if s == "" {
		return nil
	}

	v := EnvironmentStateEnum(s)
	return &v
}

func (v EnvironmentStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "RUNNING", "UPDATING", "DELETING", "ERROR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EnvironmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type EnvironmentConfig struct {
	empty                         bool                                            `json:"-"`
	GkeCluster                    *string                                         `json:"gkeCluster"`
	DagGcsPrefix                  *string                                         `json:"dagGcsPrefix"`
	NodeCount                     *int64                                          `json:"nodeCount"`
	SoftwareConfig                *EnvironmentConfigSoftwareConfig                `json:"softwareConfig"`
	NodeConfig                    *EnvironmentConfigNodeConfig                    `json:"nodeConfig"`
	PrivateEnvironmentConfig      *EnvironmentConfigPrivateEnvironmentConfig      `json:"privateEnvironmentConfig"`
	WebServerNetworkAccessControl *EnvironmentConfigWebServerNetworkAccessControl `json:"webServerNetworkAccessControl"`
	DatabaseConfig                *EnvironmentConfigDatabaseConfig                `json:"databaseConfig"`
	WebServerConfig               *EnvironmentConfigWebServerConfig               `json:"webServerConfig"`
	EncryptionConfig              *EnvironmentConfigEncryptionConfig              `json:"encryptionConfig"`
	AirflowUri                    *string                                         `json:"airflowUri"`
}

type jsonEnvironmentConfig EnvironmentConfig

func (r *EnvironmentConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfig
	} else {

		r.GkeCluster = res.GkeCluster

		r.DagGcsPrefix = res.DagGcsPrefix

		r.NodeCount = res.NodeCount

		r.SoftwareConfig = res.SoftwareConfig

		r.NodeConfig = res.NodeConfig

		r.PrivateEnvironmentConfig = res.PrivateEnvironmentConfig

		r.WebServerNetworkAccessControl = res.WebServerNetworkAccessControl

		r.DatabaseConfig = res.DatabaseConfig

		r.WebServerConfig = res.WebServerConfig

		r.EncryptionConfig = res.EncryptionConfig

		r.AirflowUri = res.AirflowUri

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfig *EnvironmentConfig = &EnvironmentConfig{empty: true}

func (r *EnvironmentConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigSoftwareConfig struct {
	empty                  bool              `json:"-"`
	ImageVersion           *string           `json:"imageVersion"`
	AirflowConfigOverrides map[string]string `json:"airflowConfigOverrides"`
	PypiPackages           map[string]string `json:"pypiPackages"`
	EnvVariables           map[string]string `json:"envVariables"`
	PythonVersion          *string           `json:"pythonVersion"`
}

type jsonEnvironmentConfigSoftwareConfig EnvironmentConfigSoftwareConfig

func (r *EnvironmentConfigSoftwareConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigSoftwareConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigSoftwareConfig
	} else {

		r.ImageVersion = res.ImageVersion

		r.AirflowConfigOverrides = res.AirflowConfigOverrides

		r.PypiPackages = res.PypiPackages

		r.EnvVariables = res.EnvVariables

		r.PythonVersion = res.PythonVersion

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigSoftwareConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigSoftwareConfig *EnvironmentConfigSoftwareConfig = &EnvironmentConfigSoftwareConfig{empty: true}

func (r *EnvironmentConfigSoftwareConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigSoftwareConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigSoftwareConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigNodeConfig struct {
	empty              bool                                           `json:"-"`
	Location           *string                                        `json:"location"`
	MachineType        *string                                        `json:"machineType"`
	Network            *string                                        `json:"network"`
	Subnetwork         *string                                        `json:"subnetwork"`
	DiskSizeGb         *int64                                         `json:"diskSizeGb"`
	OAuthScopes        []string                                       `json:"oauthScopes"`
	ServiceAccount     *string                                        `json:"serviceAccount"`
	Tags               []string                                       `json:"tags"`
	IPAllocationPolicy *EnvironmentConfigNodeConfigIPAllocationPolicy `json:"ipAllocationPolicy"`
}

type jsonEnvironmentConfigNodeConfig EnvironmentConfigNodeConfig

func (r *EnvironmentConfigNodeConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigNodeConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigNodeConfig
	} else {

		r.Location = res.Location

		r.MachineType = res.MachineType

		r.Network = res.Network

		r.Subnetwork = res.Subnetwork

		r.DiskSizeGb = res.DiskSizeGb

		r.OAuthScopes = res.OAuthScopes

		r.ServiceAccount = res.ServiceAccount

		r.Tags = res.Tags

		r.IPAllocationPolicy = res.IPAllocationPolicy

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigNodeConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigNodeConfig *EnvironmentConfigNodeConfig = &EnvironmentConfigNodeConfig{empty: true}

func (r *EnvironmentConfigNodeConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigNodeConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigNodeConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigNodeConfigIPAllocationPolicy struct {
	empty                      bool    `json:"-"`
	UseIPAliases               *bool   `json:"useIPAliases"`
	ClusterSecondaryRangeName  *string `json:"clusterSecondaryRangeName"`
	ClusterIPv4CidrBlock       *string `json:"clusterIPv4CidrBlock"`
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName"`
	ServicesIPv4CidrBlock      *string `json:"servicesIPv4CidrBlock"`
}

type jsonEnvironmentConfigNodeConfigIPAllocationPolicy EnvironmentConfigNodeConfigIPAllocationPolicy

func (r *EnvironmentConfigNodeConfigIPAllocationPolicy) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigNodeConfigIPAllocationPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigNodeConfigIPAllocationPolicy
	} else {

		r.UseIPAliases = res.UseIPAliases

		r.ClusterSecondaryRangeName = res.ClusterSecondaryRangeName

		r.ClusterIPv4CidrBlock = res.ClusterIPv4CidrBlock

		r.ServicesSecondaryRangeName = res.ServicesSecondaryRangeName

		r.ServicesIPv4CidrBlock = res.ServicesIPv4CidrBlock

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigNodeConfigIPAllocationPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigNodeConfigIPAllocationPolicy *EnvironmentConfigNodeConfigIPAllocationPolicy = &EnvironmentConfigNodeConfigIPAllocationPolicy{empty: true}

func (r *EnvironmentConfigNodeConfigIPAllocationPolicy) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigNodeConfigIPAllocationPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigNodeConfigIPAllocationPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigPrivateEnvironmentConfig struct {
	empty                      bool                                                           `json:"-"`
	EnablePrivateEnvironment   *bool                                                          `json:"enablePrivateEnvironment"`
	PrivateClusterConfig       *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig `json:"privateClusterConfig"`
	WebServerIPv4CidrBlock     *string                                                        `json:"webServerIPv4CidrBlock"`
	CloudSqlIPv4CidrBlock      *string                                                        `json:"cloudSqlIPv4CidrBlock"`
	WebServerIPv4ReservedRange *string                                                        `json:"webServerIPv4ReservedRange"`
}

type jsonEnvironmentConfigPrivateEnvironmentConfig EnvironmentConfigPrivateEnvironmentConfig

func (r *EnvironmentConfigPrivateEnvironmentConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigPrivateEnvironmentConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigPrivateEnvironmentConfig
	} else {

		r.EnablePrivateEnvironment = res.EnablePrivateEnvironment

		r.PrivateClusterConfig = res.PrivateClusterConfig

		r.WebServerIPv4CidrBlock = res.WebServerIPv4CidrBlock

		r.CloudSqlIPv4CidrBlock = res.CloudSqlIPv4CidrBlock

		r.WebServerIPv4ReservedRange = res.WebServerIPv4ReservedRange

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigPrivateEnvironmentConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigPrivateEnvironmentConfig *EnvironmentConfigPrivateEnvironmentConfig = &EnvironmentConfigPrivateEnvironmentConfig{empty: true}

func (r *EnvironmentConfigPrivateEnvironmentConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigPrivateEnvironmentConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigPrivateEnvironmentConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig struct {
	empty                   bool    `json:"-"`
	EnablePrivateEndpoint   *bool   `json:"enablePrivateEndpoint"`
	MasterIPv4CidrBlock     *string `json:"masterIPv4CidrBlock"`
	MasterIPv4ReservedRange *string `json:"masterIPv4ReservedRange"`
}

type jsonEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig

func (r *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig
	} else {

		r.EnablePrivateEndpoint = res.EnablePrivateEndpoint

		r.MasterIPv4CidrBlock = res.MasterIPv4CidrBlock

		r.MasterIPv4ReservedRange = res.MasterIPv4ReservedRange

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig = &EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{empty: true}

func (r *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigWebServerNetworkAccessControl struct {
	empty           bool                                                            `json:"-"`
	AllowedIPRanges []EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges `json:"allowedIPRanges"`
}

type jsonEnvironmentConfigWebServerNetworkAccessControl EnvironmentConfigWebServerNetworkAccessControl

func (r *EnvironmentConfigWebServerNetworkAccessControl) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigWebServerNetworkAccessControl
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigWebServerNetworkAccessControl
	} else {

		r.AllowedIPRanges = res.AllowedIPRanges

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigWebServerNetworkAccessControl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigWebServerNetworkAccessControl *EnvironmentConfigWebServerNetworkAccessControl = &EnvironmentConfigWebServerNetworkAccessControl{empty: true}

func (r *EnvironmentConfigWebServerNetworkAccessControl) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigWebServerNetworkAccessControl) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigWebServerNetworkAccessControl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges struct {
	empty       bool    `json:"-"`
	Value       *string `json:"value"`
	Description *string `json:"description"`
}

type jsonEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges

func (r *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges
	} else {

		r.Value = res.Value

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges = &EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{empty: true}

func (r *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigDatabaseConfig struct {
	empty       bool    `json:"-"`
	MachineType *string `json:"machineType"`
}

type jsonEnvironmentConfigDatabaseConfig EnvironmentConfigDatabaseConfig

func (r *EnvironmentConfigDatabaseConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigDatabaseConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigDatabaseConfig
	} else {

		r.MachineType = res.MachineType

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigDatabaseConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigDatabaseConfig *EnvironmentConfigDatabaseConfig = &EnvironmentConfigDatabaseConfig{empty: true}

func (r *EnvironmentConfigDatabaseConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigDatabaseConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigDatabaseConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigWebServerConfig struct {
	empty       bool    `json:"-"`
	MachineType *string `json:"machineType"`
}

type jsonEnvironmentConfigWebServerConfig EnvironmentConfigWebServerConfig

func (r *EnvironmentConfigWebServerConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigWebServerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigWebServerConfig
	} else {

		r.MachineType = res.MachineType

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigWebServerConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigWebServerConfig *EnvironmentConfigWebServerConfig = &EnvironmentConfigWebServerConfig{empty: true}

func (r *EnvironmentConfigWebServerConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigWebServerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigWebServerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentConfigEncryptionConfig struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonEnvironmentConfigEncryptionConfig EnvironmentConfigEncryptionConfig

func (r *EnvironmentConfigEncryptionConfig) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentConfigEncryptionConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentConfigEncryptionConfig
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentConfigEncryptionConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentConfigEncryptionConfig *EnvironmentConfigEncryptionConfig = &EnvironmentConfigEncryptionConfig{empty: true}

func (r *EnvironmentConfigEncryptionConfig) Empty() bool {
	return r.empty
}

func (r *EnvironmentConfigEncryptionConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentConfigEncryptionConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Environment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "composer",
		Type:    "Environment",
		Version: "composer",
	}
}

const EnvironmentMaxPage = -1

type EnvironmentList struct {
	Items []*Environment

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *EnvironmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EnvironmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEnvironment(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEnvironment(ctx context.Context, project, location string) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEnvironmentWithMaxResults(ctx, project, location, EnvironmentMaxPage)

}

func (c *Client) ListEnvironmentWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listEnvironment(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EnvironmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetEnvironment(ctx context.Context, r *Environment) (*Environment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getEnvironmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEnvironment(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEnvironmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEnvironment(ctx context.Context, r *Environment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Environment resource is nil")
	}
	c.Config.Logger.Info("Deleting Environment...")
	deleteOp := deleteEnvironmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEnvironment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEnvironment(ctx context.Context, project, location string, filter func(*Environment) bool) error {
	listObj, err := c.ListEnvironment(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEnvironment(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (*Environment, error) {

	var resultNewState *Environment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEnvironmentHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyEnvironmentHelper(c *Client, ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (*Environment, error) {
	c.Config.Logger.Info("Beginning ApplyEnvironment...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.environmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToEnvironmentOp(opStrings, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []environmentApiOperation
	if create {
		ops = append(ops, &createEnvironmentOperation{})
	} else if recreate {

		ops = append(ops, &deleteEnvironmentOperation{})

		ops = append(ops, &createEnvironmentOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeEnvironmentDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetEnvironment(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEnvironmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEnvironment(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEnvironmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEnvironmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEnvironmentDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEnvironment(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
