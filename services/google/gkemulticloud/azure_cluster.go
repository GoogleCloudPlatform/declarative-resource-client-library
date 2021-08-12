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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AzureCluster struct {
	Name                   *string                             `json:"name"`
	Description            *string                             `json:"description"`
	AzureRegion            *string                             `json:"azureRegion"`
	ResourceGroupId        *string                             `json:"resourceGroupId"`
	AzureClient            *string                             `json:"azureClient"`
	Networking             *AzureClusterNetworking             `json:"networking"`
	ControlPlane           *AzureClusterControlPlane           `json:"controlPlane"`
	Authorization          *AzureClusterAuthorization          `json:"authorization"`
	State                  *AzureClusterStateEnum              `json:"state"`
	Endpoint               *string                             `json:"endpoint"`
	Uid                    *string                             `json:"uid"`
	Reconciling            *bool                               `json:"reconciling"`
	CreateTime             *string                             `json:"createTime"`
	UpdateTime             *string                             `json:"updateTime"`
	Etag                   *string                             `json:"etag"`
	Annotations            map[string]string                   `json:"annotations"`
	WorkloadIdentityConfig *AzureClusterWorkloadIdentityConfig `json:"workloadIdentityConfig"`
	Project                *string                             `json:"project"`
	Location               *string                             `json:"location"`
}

func (r *AzureCluster) String() string {
	return dcl.SprintResource(r)
}

// The enum AzureClusterStateEnum.
type AzureClusterStateEnum string

// AzureClusterStateEnumRef returns a *AzureClusterStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func AzureClusterStateEnumRef(s string) *AzureClusterStateEnum {
	if s == "" {
		return nil
	}

	v := AzureClusterStateEnum(s)
	return &v
}

func (v AzureClusterStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PROVISIONING", "RUNNING", "RECONCILING", "STOPPING", "ERROR", "DEGRADED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AzureClusterStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AzureClusterNetworking struct {
	empty                    bool     `json:"-"`
	VirtualNetworkId         *string  `json:"virtualNetworkId"`
	PodAddressCidrBlocks     []string `json:"podAddressCidrBlocks"`
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks"`
}

type jsonAzureClusterNetworking AzureClusterNetworking

func (r *AzureClusterNetworking) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterNetworking
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterNetworking
	} else {

		r.VirtualNetworkId = res.VirtualNetworkId

		r.PodAddressCidrBlocks = res.PodAddressCidrBlocks

		r.ServiceAddressCidrBlocks = res.ServiceAddressCidrBlocks

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterNetworking is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterNetworking *AzureClusterNetworking = &AzureClusterNetworking{empty: true}

func (r *AzureClusterNetworking) Empty() bool {
	return r.empty
}

func (r *AzureClusterNetworking) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterNetworking) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterControlPlane struct {
	empty              bool                                        `json:"-"`
	Version            *string                                     `json:"version"`
	SubnetId           *string                                     `json:"subnetId"`
	VmSize             *string                                     `json:"vmSize"`
	SshConfig          *AzureClusterControlPlaneSshConfig          `json:"sshConfig"`
	RootVolume         *AzureClusterControlPlaneRootVolume         `json:"rootVolume"`
	MainVolume         *AzureClusterControlPlaneMainVolume         `json:"mainVolume"`
	DatabaseEncryption *AzureClusterControlPlaneDatabaseEncryption `json:"databaseEncryption"`
	Tags               map[string]string                           `json:"tags"`
}

type jsonAzureClusterControlPlane AzureClusterControlPlane

func (r *AzureClusterControlPlane) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterControlPlane
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterControlPlane
	} else {

		r.Version = res.Version

		r.SubnetId = res.SubnetId

		r.VmSize = res.VmSize

		r.SshConfig = res.SshConfig

		r.RootVolume = res.RootVolume

		r.MainVolume = res.MainVolume

		r.DatabaseEncryption = res.DatabaseEncryption

		r.Tags = res.Tags

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterControlPlane is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterControlPlane *AzureClusterControlPlane = &AzureClusterControlPlane{empty: true}

func (r *AzureClusterControlPlane) Empty() bool {
	return r.empty
}

func (r *AzureClusterControlPlane) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterControlPlane) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterControlPlaneSshConfig struct {
	empty         bool    `json:"-"`
	AuthorizedKey *string `json:"authorizedKey"`
}

type jsonAzureClusterControlPlaneSshConfig AzureClusterControlPlaneSshConfig

func (r *AzureClusterControlPlaneSshConfig) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterControlPlaneSshConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterControlPlaneSshConfig
	} else {

		r.AuthorizedKey = res.AuthorizedKey

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterControlPlaneSshConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterControlPlaneSshConfig *AzureClusterControlPlaneSshConfig = &AzureClusterControlPlaneSshConfig{empty: true}

func (r *AzureClusterControlPlaneSshConfig) Empty() bool {
	return r.empty
}

func (r *AzureClusterControlPlaneSshConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterControlPlaneSshConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterControlPlaneRootVolume struct {
	empty   bool   `json:"-"`
	SizeGib *int64 `json:"sizeGib"`
}

type jsonAzureClusterControlPlaneRootVolume AzureClusterControlPlaneRootVolume

func (r *AzureClusterControlPlaneRootVolume) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterControlPlaneRootVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterControlPlaneRootVolume
	} else {

		r.SizeGib = res.SizeGib

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterControlPlaneRootVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterControlPlaneRootVolume *AzureClusterControlPlaneRootVolume = &AzureClusterControlPlaneRootVolume{empty: true}

func (r *AzureClusterControlPlaneRootVolume) Empty() bool {
	return r.empty
}

func (r *AzureClusterControlPlaneRootVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterControlPlaneRootVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterControlPlaneMainVolume struct {
	empty   bool   `json:"-"`
	SizeGib *int64 `json:"sizeGib"`
}

type jsonAzureClusterControlPlaneMainVolume AzureClusterControlPlaneMainVolume

func (r *AzureClusterControlPlaneMainVolume) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterControlPlaneMainVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterControlPlaneMainVolume
	} else {

		r.SizeGib = res.SizeGib

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterControlPlaneMainVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterControlPlaneMainVolume *AzureClusterControlPlaneMainVolume = &AzureClusterControlPlaneMainVolume{empty: true}

func (r *AzureClusterControlPlaneMainVolume) Empty() bool {
	return r.empty
}

func (r *AzureClusterControlPlaneMainVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterControlPlaneMainVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterControlPlaneDatabaseEncryption struct {
	empty            bool    `json:"-"`
	ResourceGroupId  *string `json:"resourceGroupId"`
	KmsKeyIdentifier *string `json:"kmsKeyIdentifier"`
}

type jsonAzureClusterControlPlaneDatabaseEncryption AzureClusterControlPlaneDatabaseEncryption

func (r *AzureClusterControlPlaneDatabaseEncryption) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterControlPlaneDatabaseEncryption
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterControlPlaneDatabaseEncryption
	} else {

		r.ResourceGroupId = res.ResourceGroupId

		r.KmsKeyIdentifier = res.KmsKeyIdentifier

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterControlPlaneDatabaseEncryption is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterControlPlaneDatabaseEncryption *AzureClusterControlPlaneDatabaseEncryption = &AzureClusterControlPlaneDatabaseEncryption{empty: true}

func (r *AzureClusterControlPlaneDatabaseEncryption) Empty() bool {
	return r.empty
}

func (r *AzureClusterControlPlaneDatabaseEncryption) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterControlPlaneDatabaseEncryption) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterAuthorization struct {
	empty      bool                                  `json:"-"`
	AdminUsers []AzureClusterAuthorizationAdminUsers `json:"adminUsers"`
}

type jsonAzureClusterAuthorization AzureClusterAuthorization

func (r *AzureClusterAuthorization) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterAuthorization
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterAuthorization
	} else {

		r.AdminUsers = res.AdminUsers

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterAuthorization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterAuthorization *AzureClusterAuthorization = &AzureClusterAuthorization{empty: true}

func (r *AzureClusterAuthorization) Empty() bool {
	return r.empty
}

func (r *AzureClusterAuthorization) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterAuthorization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterAuthorizationAdminUsers struct {
	empty    bool    `json:"-"`
	Username *string `json:"username"`
}

type jsonAzureClusterAuthorizationAdminUsers AzureClusterAuthorizationAdminUsers

func (r *AzureClusterAuthorizationAdminUsers) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterAuthorizationAdminUsers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterAuthorizationAdminUsers
	} else {

		r.Username = res.Username

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterAuthorizationAdminUsers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterAuthorizationAdminUsers *AzureClusterAuthorizationAdminUsers = &AzureClusterAuthorizationAdminUsers{empty: true}

func (r *AzureClusterAuthorizationAdminUsers) Empty() bool {
	return r.empty
}

func (r *AzureClusterAuthorizationAdminUsers) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterAuthorizationAdminUsers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureClusterWorkloadIdentityConfig struct {
	empty            bool    `json:"-"`
	IssuerUri        *string `json:"issuerUri"`
	WorkloadPool     *string `json:"workloadPool"`
	IdentityProvider *string `json:"identityProvider"`
}

type jsonAzureClusterWorkloadIdentityConfig AzureClusterWorkloadIdentityConfig

func (r *AzureClusterWorkloadIdentityConfig) UnmarshalJSON(data []byte) error {
	var res jsonAzureClusterWorkloadIdentityConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureClusterWorkloadIdentityConfig
	} else {

		r.IssuerUri = res.IssuerUri

		r.WorkloadPool = res.WorkloadPool

		r.IdentityProvider = res.IdentityProvider

	}
	return nil
}

// This object is used to assert a desired state where this AzureClusterWorkloadIdentityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureClusterWorkloadIdentityConfig *AzureClusterWorkloadIdentityConfig = &AzureClusterWorkloadIdentityConfig{empty: true}

func (r *AzureClusterWorkloadIdentityConfig) Empty() bool {
	return r.empty
}

func (r *AzureClusterWorkloadIdentityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureClusterWorkloadIdentityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AzureCluster) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gkemulticloud",
		Type:    "AzureCluster",
		Version: "gkemulticloud",
	}
}

func (r *AzureCluster) ID() (string, error) {
	if err := extractAzureClusterFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                   dcl.ValueOrEmptyString(nr.Name),
		"description":            dcl.ValueOrEmptyString(nr.Description),
		"azureRegion":            dcl.ValueOrEmptyString(nr.AzureRegion),
		"resourceGroupId":        dcl.ValueOrEmptyString(nr.ResourceGroupId),
		"azureClient":            dcl.ValueOrEmptyString(nr.AzureClient),
		"networking":             dcl.ValueOrEmptyString(nr.Networking),
		"controlPlane":           dcl.ValueOrEmptyString(nr.ControlPlane),
		"authorization":          dcl.ValueOrEmptyString(nr.Authorization),
		"state":                  dcl.ValueOrEmptyString(nr.State),
		"endpoint":               dcl.ValueOrEmptyString(nr.Endpoint),
		"uid":                    dcl.ValueOrEmptyString(nr.Uid),
		"reconciling":            dcl.ValueOrEmptyString(nr.Reconciling),
		"createTime":             dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":             dcl.ValueOrEmptyString(nr.UpdateTime),
		"etag":                   dcl.ValueOrEmptyString(nr.Etag),
		"annotations":            dcl.ValueOrEmptyString(nr.Annotations),
		"workloadIdentityConfig": dcl.ValueOrEmptyString(nr.WorkloadIdentityConfig),
		"project":                dcl.ValueOrEmptyString(nr.Project),
		"location":               dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/azureClusters/{{name}}", params), nil
}

const AzureClusterMaxPage = -1

type AzureClusterList struct {
	Items []*AzureCluster

	nextToken string

	pageSize int32

	resource *AzureCluster
}

func (l *AzureClusterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AzureClusterList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAzureCluster(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAzureCluster(ctx context.Context, r *AzureCluster) (*AzureClusterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAzureClusterWithMaxResults(ctx, r, AzureClusterMaxPage)

}

func (c *Client) ListAzureClusterWithMaxResults(ctx context.Context, r *AzureCluster, pageSize int32) (*AzureClusterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAzureCluster(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AzureClusterList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAzureCluster(ctx context.Context, r *AzureCluster) (*AzureCluster, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAzureClusterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAzureCluster(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAzureClusterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAzureCluster(ctx context.Context, r *AzureCluster) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AzureCluster resource is nil")
	}
	c.Config.Logger.Info("Deleting AzureCluster...")
	deleteOp := deleteAzureClusterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAzureCluster deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAzureCluster(ctx context.Context, project, location string, filter func(*AzureCluster) bool) error {
	r := &AzureCluster{

		Project: &project,

		Location: &location,
	}
	listObj, err := c.ListAzureCluster(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAzureCluster(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAzureCluster(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAzureCluster(ctx context.Context, rawDesired *AzureCluster, opts ...dcl.ApplyOption) (*AzureCluster, error) {
	var resultNewState *AzureCluster
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAzureClusterHelper(c, ctx, rawDesired, opts...)
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

func applyAzureClusterHelper(c *Client, ctx context.Context, rawDesired *AzureCluster, opts ...dcl.ApplyOption) (*AzureCluster, error) {
	c.Config.Logger.Info("Beginning ApplyAzureCluster...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractAzureClusterFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.azureClusterDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAzureClusterDiffs(c.Config, fieldDiffs, opts)
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
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []azureClusterApiOperation
	if create {
		ops = append(ops, &createAzureClusterOperation{})
	} else if recreate {
		ops = append(ops, &deleteAzureClusterOperation{})
		ops = append(ops, &createAzureClusterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAzureClusterDesiredState(rawDesired, nil)
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
		c.Config.Logger.Infof("Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetAzureCluster(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAzureClusterOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAzureCluster(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAzureClusterNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAzureClusterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAzureClusterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAzureCluster(c, newDesired, newState)
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
