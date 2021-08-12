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

type AwsCluster struct {
	Name                   *string                           `json:"name"`
	Description            *string                           `json:"description"`
	Networking             *AwsClusterNetworking             `json:"networking"`
	AwsRegion              *string                           `json:"awsRegion"`
	ControlPlane           *AwsClusterControlPlane           `json:"controlPlane"`
	Authorization          *AwsClusterAuthorization          `json:"authorization"`
	State                  *AwsClusterStateEnum              `json:"state"`
	Endpoint               *string                           `json:"endpoint"`
	Uid                    *string                           `json:"uid"`
	Reconciling            *bool                             `json:"reconciling"`
	CreateTime             *string                           `json:"createTime"`
	UpdateTime             *string                           `json:"updateTime"`
	Etag                   *string                           `json:"etag"`
	Annotations            map[string]string                 `json:"annotations"`
	WorkloadIdentityConfig *AwsClusterWorkloadIdentityConfig `json:"workloadIdentityConfig"`
	Project                *string                           `json:"project"`
	Location               *string                           `json:"location"`
}

func (r *AwsCluster) String() string {
	return dcl.SprintResource(r)
}

// The enum AwsClusterControlPlaneRootVolumeVolumeTypeEnum.
type AwsClusterControlPlaneRootVolumeVolumeTypeEnum string

// AwsClusterControlPlaneRootVolumeVolumeTypeEnumRef returns a *AwsClusterControlPlaneRootVolumeVolumeTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsClusterControlPlaneRootVolumeVolumeTypeEnumRef(s string) *AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if s == "" {
		return nil
	}

	v := AwsClusterControlPlaneRootVolumeVolumeTypeEnum(s)
	return &v
}

func (v AwsClusterControlPlaneRootVolumeVolumeTypeEnum) Validate() error {
	for _, s := range []string{"VOLUME_TYPE_UNSPECIFIED", "GP2", "GP3"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsClusterControlPlaneRootVolumeVolumeTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AwsClusterControlPlaneMainVolumeVolumeTypeEnum.
type AwsClusterControlPlaneMainVolumeVolumeTypeEnum string

// AwsClusterControlPlaneMainVolumeVolumeTypeEnumRef returns a *AwsClusterControlPlaneMainVolumeVolumeTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsClusterControlPlaneMainVolumeVolumeTypeEnumRef(s string) *AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if s == "" {
		return nil
	}

	v := AwsClusterControlPlaneMainVolumeVolumeTypeEnum(s)
	return &v
}

func (v AwsClusterControlPlaneMainVolumeVolumeTypeEnum) Validate() error {
	for _, s := range []string{"VOLUME_TYPE_UNSPECIFIED", "GP2", "GP3"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsClusterControlPlaneMainVolumeVolumeTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AwsClusterStateEnum.
type AwsClusterStateEnum string

// AwsClusterStateEnumRef returns a *AwsClusterStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsClusterStateEnumRef(s string) *AwsClusterStateEnum {
	if s == "" {
		return nil
	}

	v := AwsClusterStateEnum(s)
	return &v
}

func (v AwsClusterStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PROVISIONING", "RUNNING", "RECONCILING", "STOPPING", "ERROR", "DEGRADED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsClusterStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AwsClusterNetworking struct {
	empty                        bool     `json:"-"`
	VPCId                        *string  `json:"vpcId"`
	PodAddressCidrBlocks         []string `json:"podAddressCidrBlocks"`
	ServiceAddressCidrBlocks     []string `json:"serviceAddressCidrBlocks"`
	ServiceLoadBalancerSubnetIds []string `json:"serviceLoadBalancerSubnetIds"`
}

type jsonAwsClusterNetworking AwsClusterNetworking

func (r *AwsClusterNetworking) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterNetworking
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterNetworking
	} else {

		r.VPCId = res.VPCId

		r.PodAddressCidrBlocks = res.PodAddressCidrBlocks

		r.ServiceAddressCidrBlocks = res.ServiceAddressCidrBlocks

		r.ServiceLoadBalancerSubnetIds = res.ServiceLoadBalancerSubnetIds

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterNetworking is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterNetworking *AwsClusterNetworking = &AwsClusterNetworking{empty: true}

func (r *AwsClusterNetworking) Empty() bool {
	return r.empty
}

func (r *AwsClusterNetworking) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterNetworking) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlane struct {
	empty                     bool                                             `json:"-"`
	Version                   *string                                          `json:"version"`
	InstanceType              *string                                          `json:"instanceType"`
	SshConfig                 *AwsClusterControlPlaneSshConfig                 `json:"sshConfig"`
	SubnetIds                 []string                                         `json:"subnetIds"`
	SecurityGroupIds          []string                                         `json:"securityGroupIds"`
	IamInstanceProfile        *string                                          `json:"iamInstanceProfile"`
	RootVolume                *AwsClusterControlPlaneRootVolume                `json:"rootVolume"`
	MainVolume                *AwsClusterControlPlaneMainVolume                `json:"mainVolume"`
	DatabaseEncryption        *AwsClusterControlPlaneDatabaseEncryption        `json:"databaseEncryption"`
	Tags                      map[string]string                                `json:"tags"`
	AwsServicesAuthentication *AwsClusterControlPlaneAwsServicesAuthentication `json:"awsServicesAuthentication"`
}

type jsonAwsClusterControlPlane AwsClusterControlPlane

func (r *AwsClusterControlPlane) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlane
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlane
	} else {

		r.Version = res.Version

		r.InstanceType = res.InstanceType

		r.SshConfig = res.SshConfig

		r.SubnetIds = res.SubnetIds

		r.SecurityGroupIds = res.SecurityGroupIds

		r.IamInstanceProfile = res.IamInstanceProfile

		r.RootVolume = res.RootVolume

		r.MainVolume = res.MainVolume

		r.DatabaseEncryption = res.DatabaseEncryption

		r.Tags = res.Tags

		r.AwsServicesAuthentication = res.AwsServicesAuthentication

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlane is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlane *AwsClusterControlPlane = &AwsClusterControlPlane{empty: true}

func (r *AwsClusterControlPlane) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlane) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlane) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlaneSshConfig struct {
	empty      bool    `json:"-"`
	Ec2KeyPair *string `json:"ec2KeyPair"`
}

type jsonAwsClusterControlPlaneSshConfig AwsClusterControlPlaneSshConfig

func (r *AwsClusterControlPlaneSshConfig) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlaneSshConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlaneSshConfig
	} else {

		r.Ec2KeyPair = res.Ec2KeyPair

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlaneSshConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlaneSshConfig *AwsClusterControlPlaneSshConfig = &AwsClusterControlPlaneSshConfig{empty: true}

func (r *AwsClusterControlPlaneSshConfig) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlaneSshConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlaneSshConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlaneRootVolume struct {
	empty      bool                                            `json:"-"`
	SizeGib    *int64                                          `json:"sizeGib"`
	VolumeType *AwsClusterControlPlaneRootVolumeVolumeTypeEnum `json:"volumeType"`
	Iops       *int64                                          `json:"iops"`
	KmsKeyArn  *string                                         `json:"kmsKeyArn"`
}

type jsonAwsClusterControlPlaneRootVolume AwsClusterControlPlaneRootVolume

func (r *AwsClusterControlPlaneRootVolume) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlaneRootVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlaneRootVolume
	} else {

		r.SizeGib = res.SizeGib

		r.VolumeType = res.VolumeType

		r.Iops = res.Iops

		r.KmsKeyArn = res.KmsKeyArn

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlaneRootVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlaneRootVolume *AwsClusterControlPlaneRootVolume = &AwsClusterControlPlaneRootVolume{empty: true}

func (r *AwsClusterControlPlaneRootVolume) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlaneRootVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlaneRootVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlaneMainVolume struct {
	empty      bool                                            `json:"-"`
	SizeGib    *int64                                          `json:"sizeGib"`
	VolumeType *AwsClusterControlPlaneMainVolumeVolumeTypeEnum `json:"volumeType"`
	Iops       *int64                                          `json:"iops"`
	KmsKeyArn  *string                                         `json:"kmsKeyArn"`
}

type jsonAwsClusterControlPlaneMainVolume AwsClusterControlPlaneMainVolume

func (r *AwsClusterControlPlaneMainVolume) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlaneMainVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlaneMainVolume
	} else {

		r.SizeGib = res.SizeGib

		r.VolumeType = res.VolumeType

		r.Iops = res.Iops

		r.KmsKeyArn = res.KmsKeyArn

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlaneMainVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlaneMainVolume *AwsClusterControlPlaneMainVolume = &AwsClusterControlPlaneMainVolume{empty: true}

func (r *AwsClusterControlPlaneMainVolume) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlaneMainVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlaneMainVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlaneDatabaseEncryption struct {
	empty     bool    `json:"-"`
	KmsKeyArn *string `json:"kmsKeyArn"`
}

type jsonAwsClusterControlPlaneDatabaseEncryption AwsClusterControlPlaneDatabaseEncryption

func (r *AwsClusterControlPlaneDatabaseEncryption) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlaneDatabaseEncryption
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlaneDatabaseEncryption
	} else {

		r.KmsKeyArn = res.KmsKeyArn

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlaneDatabaseEncryption is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlaneDatabaseEncryption *AwsClusterControlPlaneDatabaseEncryption = &AwsClusterControlPlaneDatabaseEncryption{empty: true}

func (r *AwsClusterControlPlaneDatabaseEncryption) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlaneDatabaseEncryption) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlaneDatabaseEncryption) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterControlPlaneAwsServicesAuthentication struct {
	empty           bool    `json:"-"`
	RoleArn         *string `json:"roleArn"`
	RoleSessionName *string `json:"roleSessionName"`
}

type jsonAwsClusterControlPlaneAwsServicesAuthentication AwsClusterControlPlaneAwsServicesAuthentication

func (r *AwsClusterControlPlaneAwsServicesAuthentication) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterControlPlaneAwsServicesAuthentication
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterControlPlaneAwsServicesAuthentication
	} else {

		r.RoleArn = res.RoleArn

		r.RoleSessionName = res.RoleSessionName

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterControlPlaneAwsServicesAuthentication is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterControlPlaneAwsServicesAuthentication *AwsClusterControlPlaneAwsServicesAuthentication = &AwsClusterControlPlaneAwsServicesAuthentication{empty: true}

func (r *AwsClusterControlPlaneAwsServicesAuthentication) Empty() bool {
	return r.empty
}

func (r *AwsClusterControlPlaneAwsServicesAuthentication) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterControlPlaneAwsServicesAuthentication) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterAuthorization struct {
	empty      bool                                `json:"-"`
	AdminUsers []AwsClusterAuthorizationAdminUsers `json:"adminUsers"`
}

type jsonAwsClusterAuthorization AwsClusterAuthorization

func (r *AwsClusterAuthorization) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterAuthorization
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterAuthorization
	} else {

		r.AdminUsers = res.AdminUsers

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterAuthorization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterAuthorization *AwsClusterAuthorization = &AwsClusterAuthorization{empty: true}

func (r *AwsClusterAuthorization) Empty() bool {
	return r.empty
}

func (r *AwsClusterAuthorization) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterAuthorization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterAuthorizationAdminUsers struct {
	empty    bool    `json:"-"`
	Username *string `json:"username"`
}

type jsonAwsClusterAuthorizationAdminUsers AwsClusterAuthorizationAdminUsers

func (r *AwsClusterAuthorizationAdminUsers) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterAuthorizationAdminUsers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterAuthorizationAdminUsers
	} else {

		r.Username = res.Username

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterAuthorizationAdminUsers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterAuthorizationAdminUsers *AwsClusterAuthorizationAdminUsers = &AwsClusterAuthorizationAdminUsers{empty: true}

func (r *AwsClusterAuthorizationAdminUsers) Empty() bool {
	return r.empty
}

func (r *AwsClusterAuthorizationAdminUsers) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterAuthorizationAdminUsers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsClusterWorkloadIdentityConfig struct {
	empty            bool    `json:"-"`
	IssuerUri        *string `json:"issuerUri"`
	WorkloadPool     *string `json:"workloadPool"`
	IdentityProvider *string `json:"identityProvider"`
}

type jsonAwsClusterWorkloadIdentityConfig AwsClusterWorkloadIdentityConfig

func (r *AwsClusterWorkloadIdentityConfig) UnmarshalJSON(data []byte) error {
	var res jsonAwsClusterWorkloadIdentityConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsClusterWorkloadIdentityConfig
	} else {

		r.IssuerUri = res.IssuerUri

		r.WorkloadPool = res.WorkloadPool

		r.IdentityProvider = res.IdentityProvider

	}
	return nil
}

// This object is used to assert a desired state where this AwsClusterWorkloadIdentityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsClusterWorkloadIdentityConfig *AwsClusterWorkloadIdentityConfig = &AwsClusterWorkloadIdentityConfig{empty: true}

func (r *AwsClusterWorkloadIdentityConfig) Empty() bool {
	return r.empty
}

func (r *AwsClusterWorkloadIdentityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsClusterWorkloadIdentityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AwsCluster) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gkemulticloud",
		Type:    "AwsCluster",
		Version: "gkemulticloud",
	}
}

func (r *AwsCluster) ID() (string, error) {
	if err := extractAwsClusterFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                   dcl.ValueOrEmptyString(nr.Name),
		"description":            dcl.ValueOrEmptyString(nr.Description),
		"networking":             dcl.ValueOrEmptyString(nr.Networking),
		"awsRegion":              dcl.ValueOrEmptyString(nr.AwsRegion),
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
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/awsClusters/{{name}}", params), nil
}

const AwsClusterMaxPage = -1

type AwsClusterList struct {
	Items []*AwsCluster

	nextToken string

	pageSize int32

	resource *AwsCluster
}

func (l *AwsClusterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AwsClusterList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAwsCluster(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAwsCluster(ctx context.Context, r *AwsCluster) (*AwsClusterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAwsClusterWithMaxResults(ctx, r, AwsClusterMaxPage)

}

func (c *Client) ListAwsClusterWithMaxResults(ctx context.Context, r *AwsCluster, pageSize int32) (*AwsClusterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAwsCluster(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AwsClusterList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAwsCluster(ctx context.Context, r *AwsCluster) (*AwsCluster, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAwsClusterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAwsCluster(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAwsClusterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAwsCluster(ctx context.Context, r *AwsCluster) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AwsCluster resource is nil")
	}
	c.Config.Logger.Info("Deleting AwsCluster...")
	deleteOp := deleteAwsClusterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAwsCluster deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAwsCluster(ctx context.Context, project, location string, filter func(*AwsCluster) bool) error {
	r := &AwsCluster{

		Project: &project,

		Location: &location,
	}
	listObj, err := c.ListAwsCluster(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAwsCluster(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAwsCluster(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAwsCluster(ctx context.Context, rawDesired *AwsCluster, opts ...dcl.ApplyOption) (*AwsCluster, error) {
	var resultNewState *AwsCluster
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAwsClusterHelper(c, ctx, rawDesired, opts...)
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

func applyAwsClusterHelper(c *Client, ctx context.Context, rawDesired *AwsCluster, opts ...dcl.ApplyOption) (*AwsCluster, error) {
	c.Config.Logger.Info("Beginning ApplyAwsCluster...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractAwsClusterFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.awsClusterDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAwsClusterDiffs(c.Config, fieldDiffs, opts)
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
	var ops []awsClusterApiOperation
	if create {
		ops = append(ops, &createAwsClusterOperation{})
	} else if recreate {
		ops = append(ops, &deleteAwsClusterOperation{})
		ops = append(ops, &createAwsClusterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAwsClusterDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAwsCluster(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAwsClusterOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAwsCluster(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAwsClusterNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAwsClusterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAwsClusterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAwsCluster(c, newDesired, newState)
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
