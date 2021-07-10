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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type OsPolicyAssignment struct {
	Name               *string                             `json:"name"`
	Description        *string                             `json:"description"`
	OsPolicies         []OsPolicyAssignmentOsPolicies      `json:"osPolicies"`
	InstanceFilter     *OsPolicyAssignmentInstanceFilter   `json:"instanceFilter"`
	Rollout            *OsPolicyAssignmentRollout          `json:"rollout"`
	RevisionId         *string                             `json:"revisionId"`
	RevisionCreateTime *string                             `json:"revisionCreateTime"`
	RolloutState       *OsPolicyAssignmentRolloutStateEnum `json:"rolloutState"`
	Baseline           *bool                               `json:"baseline"`
	Deleted            *bool                               `json:"deleted"`
	Reconciling        *bool                               `json:"reconciling"`
	Uid                *string                             `json:"uid"`
	Project            *string                             `json:"project"`
	Location           *string                             `json:"location"`
}

func (r *OsPolicyAssignment) String() string {
	return dcl.SprintResource(r)
}

// The enum OsPolicyAssignmentOsPoliciesModeEnum.
type OsPolicyAssignmentOsPoliciesModeEnum string

// OsPolicyAssignmentOsPoliciesModeEnumRef returns a *OsPolicyAssignmentOsPoliciesModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentOsPoliciesModeEnumRef(s string) *OsPolicyAssignmentOsPoliciesModeEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentOsPoliciesModeEnum(s)
	return &v
}

func (v OsPolicyAssignmentOsPoliciesModeEnum) Validate() error {
	for _, s := range []string{"MODE_UNSPECIFIED", "VALIDATION", "ENFORCEMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentOsPoliciesModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.
type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum string

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumRef returns a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumRef(s string) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(s)
	return &v
}

func (v OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum) Validate() error {
	for _, s := range []string{"DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.
type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum string

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumRef returns a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumRef(s string) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(s)
	return &v
}

func (v OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) Validate() error {
	for _, s := range []string{"ARCHIVE_TYPE_UNSPECIFIED", "DEB", "DEB_SRC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OsPolicyAssignmentExecInterpreterEnum.
type OsPolicyAssignmentExecInterpreterEnum string

// OsPolicyAssignmentExecInterpreterEnumRef returns a *OsPolicyAssignmentExecInterpreterEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentExecInterpreterEnumRef(s string) *OsPolicyAssignmentExecInterpreterEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentExecInterpreterEnum(s)
	return &v
}

func (v OsPolicyAssignmentExecInterpreterEnum) Validate() error {
	for _, s := range []string{"INTERPRETER_UNSPECIFIED", "NONE", "SHELL", "POWERSHELL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentExecInterpreterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.
type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum string

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumRef returns a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumRef(s string) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(s)
	return &v
}

func (v OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum) Validate() error {
	for _, s := range []string{"OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED", "COMPLIANT", "NON_COMPLIANT", "UNKNOWN", "NO_OS_POLICIES_APPLICABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OsPolicyAssignmentRolloutStateEnum.
type OsPolicyAssignmentRolloutStateEnum string

// OsPolicyAssignmentRolloutStateEnumRef returns a *OsPolicyAssignmentRolloutStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func OsPolicyAssignmentRolloutStateEnumRef(s string) *OsPolicyAssignmentRolloutStateEnum {
	if s == "" {
		return nil
	}

	v := OsPolicyAssignmentRolloutStateEnum(s)
	return &v
}

func (v OsPolicyAssignmentRolloutStateEnum) Validate() error {
	for _, s := range []string{"ROLLOUT_STATE_UNSPECIFIED", "IN_PROGRESS", "CANCELLING", "CANCELLED", "SUCCEEDED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OsPolicyAssignmentRolloutStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type OsPolicyAssignmentOsPolicies struct {
	empty                     bool                                         `json:"-"`
	Id                        *string                                      `json:"id"`
	Description               *string                                      `json:"description"`
	Mode                      *OsPolicyAssignmentOsPoliciesModeEnum        `json:"mode"`
	ResourceGroups            []OsPolicyAssignmentOsPoliciesResourceGroups `json:"resourceGroups"`
	AllowNoResourceGroupMatch *bool                                        `json:"allowNoResourceGroupMatch"`
}

type jsonOsPolicyAssignmentOsPolicies OsPolicyAssignmentOsPolicies

func (r *OsPolicyAssignmentOsPolicies) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPolicies
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPolicies
	} else {

		r.Id = res.Id

		r.Description = res.Description

		r.Mode = res.Mode

		r.ResourceGroups = res.ResourceGroups

		r.AllowNoResourceGroupMatch = res.AllowNoResourceGroupMatch

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPolicies is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPolicies *OsPolicyAssignmentOsPolicies = &OsPolicyAssignmentOsPolicies{empty: true}

func (r *OsPolicyAssignmentOsPolicies) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPolicies) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPolicies) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroups struct {
	empty     bool                                                  `json:"-"`
	OsFilter  *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter   `json:"osFilter"`
	Resources []OsPolicyAssignmentOsPoliciesResourceGroupsResources `json:"resources"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroups OsPolicyAssignmentOsPoliciesResourceGroups

func (r *OsPolicyAssignmentOsPoliciesResourceGroups) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroups
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroups
	} else {

		r.OsFilter = res.OsFilter

		r.Resources = res.Resources

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroups is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroups *OsPolicyAssignmentOsPoliciesResourceGroups = &OsPolicyAssignmentOsPoliciesResourceGroups{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroups) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroups) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroups) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter struct {
	empty       bool    `json:"-"`
	OsShortName *string `json:"osShortName"`
	OsVersion   *string `json:"osVersion"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter
	} else {

		r.OsShortName = res.OsShortName

		r.OsVersion = res.OsVersion

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter = &OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResources struct {
	empty      bool                                                           `json:"-"`
	Id         *string                                                        `json:"id"`
	Pkg        *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg        `json:"pkg"`
	Repository *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository `json:"repository"`
	Exec       *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec       `json:"exec"`
	File       *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile       `json:"file"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResources OsPolicyAssignmentOsPoliciesResourceGroupsResources

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResources) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResources
	} else {

		r.Id = res.Id

		r.Pkg = res.Pkg

		r.Repository = res.Repository

		r.Exec = res.Exec

		r.File = res.File

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResources *OsPolicyAssignmentOsPoliciesResourceGroupsResources = &OsPolicyAssignmentOsPoliciesResourceGroupsResources{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResources) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResources) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg struct {
	empty        bool                                                                    `json:"-"`
	DesiredState *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum `json:"desiredState"`
	Apt          *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt              `json:"apt"`
	Deb          *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb              `json:"deb"`
	Yum          *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum              `json:"yum"`
	Zypper       *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper           `json:"zypper"`
	Rpm          *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm              `json:"rpm"`
	Googet       *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget           `json:"googet"`
	Msi          *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi              `json:"msi"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	} else {

		r.DesiredState = res.DesiredState

		r.Apt = res.Apt

		r.Deb = res.Deb

		r.Yum = res.Yum

		r.Zypper = res.Zypper

		r.Rpm = res.Rpm

		r.Googet = res.Googet

		r.Msi = res.Msi

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb struct {
	empty    bool                                                             `json:"-"`
	Source   *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource `json:"source"`
	PullDeps *bool                                                            `json:"pullDeps"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	} else {

		r.Source = res.Source

		r.PullDeps = res.PullDeps

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource struct {
	empty         bool                                                                   `json:"-"`
	Remote        *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote `json:"remote"`
	Gcs           *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs    `json:"gcs"`
	LocalPath     *string                                                                `json:"localPath"`
	AllowInsecure *bool                                                                  `json:"allowInsecure"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource
	} else {

		r.Remote = res.Remote

		r.Gcs = res.Gcs

		r.LocalPath = res.LocalPath

		r.AllowInsecure = res.AllowInsecure

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote struct {
	empty          bool    `json:"-"`
	Uri            *string `json:"uri"`
	Sha256Checksum *string `json:"sha256Checksum"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	} else {

		r.Uri = res.Uri

		r.Sha256Checksum = res.Sha256Checksum

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs struct {
	empty      bool    `json:"-"`
	Bucket     *string `json:"bucket"`
	Object     *string `json:"object"`
	Generation *int64  `json:"generation"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	} else {

		r.Bucket = res.Bucket

		r.Object = res.Object

		r.Generation = res.Generation

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm struct {
	empty    bool                    `json:"-"`
	Source   *OsPolicyAssignmentFile `json:"source"`
	PullDeps *bool                   `json:"pullDeps"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	} else {

		r.Source = res.Source

		r.PullDeps = res.PullDeps

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentFile struct {
	empty         bool                          `json:"-"`
	Remote        *OsPolicyAssignmentFileRemote `json:"remote"`
	Gcs           *OsPolicyAssignmentFileGcs    `json:"gcs"`
	LocalPath     *string                       `json:"localPath"`
	AllowInsecure *bool                         `json:"allowInsecure"`
}

type jsonOsPolicyAssignmentFile OsPolicyAssignmentFile

func (r *OsPolicyAssignmentFile) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentFile
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentFile
	} else {

		r.Remote = res.Remote

		r.Gcs = res.Gcs

		r.LocalPath = res.LocalPath

		r.AllowInsecure = res.AllowInsecure

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentFile is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentFile *OsPolicyAssignmentFile = &OsPolicyAssignmentFile{empty: true}

func (r *OsPolicyAssignmentFile) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentFile) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentFile) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentFileRemote struct {
	empty          bool    `json:"-"`
	Uri            *string `json:"uri"`
	Sha256Checksum *string `json:"sha256Checksum"`
}

type jsonOsPolicyAssignmentFileRemote OsPolicyAssignmentFileRemote

func (r *OsPolicyAssignmentFileRemote) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentFileRemote
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentFileRemote
	} else {

		r.Uri = res.Uri

		r.Sha256Checksum = res.Sha256Checksum

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentFileRemote is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentFileRemote *OsPolicyAssignmentFileRemote = &OsPolicyAssignmentFileRemote{empty: true}

func (r *OsPolicyAssignmentFileRemote) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentFileRemote) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentFileRemote) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentFileGcs struct {
	empty      bool    `json:"-"`
	Bucket     *string `json:"bucket"`
	Object     *string `json:"object"`
	Generation *int64  `json:"generation"`
}

type jsonOsPolicyAssignmentFileGcs OsPolicyAssignmentFileGcs

func (r *OsPolicyAssignmentFileGcs) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentFileGcs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentFileGcs
	} else {

		r.Bucket = res.Bucket

		r.Object = res.Object

		r.Generation = res.Generation

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentFileGcs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentFileGcs *OsPolicyAssignmentFileGcs = &OsPolicyAssignmentFileGcs{empty: true}

func (r *OsPolicyAssignmentFileGcs) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentFileGcs) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentFileGcs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi struct {
	empty      bool                    `json:"-"`
	Source     *OsPolicyAssignmentFile `json:"source"`
	Properties []string                `json:"properties"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	} else {

		r.Source = res.Source

		r.Properties = res.Properties

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository struct {
	empty  bool                                                                 `json:"-"`
	Apt    *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt    `json:"apt"`
	Yum    *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum    `json:"yum"`
	Zypper *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper `json:"zypper"`
	Goo    *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo    `json:"goo"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	} else {

		r.Apt = res.Apt

		r.Yum = res.Yum

		r.Zypper = res.Zypper

		r.Goo = res.Goo

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	empty        bool                                                                             `json:"-"`
	ArchiveType  *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum `json:"archiveType"`
	Uri          *string                                                                          `json:"uri"`
	Distribution *string                                                                          `json:"distribution"`
	Components   []string                                                                         `json:"components"`
	GpgKey       *string                                                                          `json:"gpgKey"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	} else {

		r.ArchiveType = res.ArchiveType

		r.Uri = res.Uri

		r.Distribution = res.Distribution

		r.Components = res.Components

		r.GpgKey = res.GpgKey

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	} else {

		r.Id = res.Id

		r.DisplayName = res.DisplayName

		r.BaseUrl = res.BaseUrl

		r.GpgKeys = res.GpgKeys

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	} else {

		r.Id = res.Id

		r.DisplayName = res.DisplayName

		r.BaseUrl = res.BaseUrl

		r.GpgKeys = res.GpgKeys

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	} else {

		r.Name = res.Name

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec struct {
	empty    bool                    `json:"-"`
	Validate *OsPolicyAssignmentExec `json:"validate"`
	Enforce  *OsPolicyAssignmentExec `json:"enforce"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	} else {

		r.Validate = res.Validate

		r.Enforce = res.Enforce

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentExec struct {
	empty       bool                                   `json:"-"`
	File        *OsPolicyAssignmentFile                `json:"file"`
	Script      *string                                `json:"script"`
	Args        []string                               `json:"args"`
	Interpreter *OsPolicyAssignmentExecInterpreterEnum `json:"interpreter"`
}

type jsonOsPolicyAssignmentExec OsPolicyAssignmentExec

func (r *OsPolicyAssignmentExec) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentExec
	} else {

		r.File = res.File

		r.Script = res.Script

		r.Args = res.Args

		r.Interpreter = res.Interpreter

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentExec *OsPolicyAssignmentExec = &OsPolicyAssignmentExec{empty: true}

func (r *OsPolicyAssignmentExec) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentExec) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile struct {
	empty       bool                                                              `json:"-"`
	File        *OsPolicyAssignmentFile                                           `json:"file"`
	Content     *string                                                           `json:"content"`
	Path        *string                                                           `json:"path"`
	State       *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum `json:"state"`
	Permissions *string                                                           `json:"permissions"`
}

type jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	} else {

		r.File = res.File

		r.Content = res.Content

		r.Path = res.Path

		r.State = res.State

		r.Permissions = res.Permissions

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile = &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{empty: true}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentInstanceFilter struct {
	empty           bool                                              `json:"-"`
	All             *bool                                             `json:"all"`
	OsShortNames    []string                                          `json:"osShortNames"`
	InclusionLabels []OsPolicyAssignmentInstanceFilterInclusionLabels `json:"inclusionLabels"`
	ExclusionLabels []OsPolicyAssignmentInstanceFilterExclusionLabels `json:"exclusionLabels"`
}

type jsonOsPolicyAssignmentInstanceFilter OsPolicyAssignmentInstanceFilter

func (r *OsPolicyAssignmentInstanceFilter) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentInstanceFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentInstanceFilter
	} else {

		r.All = res.All

		r.OsShortNames = res.OsShortNames

		r.InclusionLabels = res.InclusionLabels

		r.ExclusionLabels = res.ExclusionLabels

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentInstanceFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentInstanceFilter *OsPolicyAssignmentInstanceFilter = &OsPolicyAssignmentInstanceFilter{empty: true}

func (r *OsPolicyAssignmentInstanceFilter) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentInstanceFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentInstanceFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentInstanceFilterInclusionLabels struct {
	empty  bool              `json:"-"`
	Labels map[string]string `json:"labels"`
}

type jsonOsPolicyAssignmentInstanceFilterInclusionLabels OsPolicyAssignmentInstanceFilterInclusionLabels

func (r *OsPolicyAssignmentInstanceFilterInclusionLabels) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentInstanceFilterInclusionLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentInstanceFilterInclusionLabels
	} else {

		r.Labels = res.Labels

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentInstanceFilterInclusionLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentInstanceFilterInclusionLabels *OsPolicyAssignmentInstanceFilterInclusionLabels = &OsPolicyAssignmentInstanceFilterInclusionLabels{empty: true}

func (r *OsPolicyAssignmentInstanceFilterInclusionLabels) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentInstanceFilterInclusionLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentInstanceFilterInclusionLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentInstanceFilterExclusionLabels struct {
	empty  bool              `json:"-"`
	Labels map[string]string `json:"labels"`
}

type jsonOsPolicyAssignmentInstanceFilterExclusionLabels OsPolicyAssignmentInstanceFilterExclusionLabels

func (r *OsPolicyAssignmentInstanceFilterExclusionLabels) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentInstanceFilterExclusionLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentInstanceFilterExclusionLabels
	} else {

		r.Labels = res.Labels

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentInstanceFilterExclusionLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentInstanceFilterExclusionLabels *OsPolicyAssignmentInstanceFilterExclusionLabels = &OsPolicyAssignmentInstanceFilterExclusionLabels{empty: true}

func (r *OsPolicyAssignmentInstanceFilterExclusionLabels) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentInstanceFilterExclusionLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentInstanceFilterExclusionLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentRollout struct {
	empty            bool                                       `json:"-"`
	DisruptionBudget *OsPolicyAssignmentRolloutDisruptionBudget `json:"disruptionBudget"`
	MinWaitDuration  *string                                    `json:"minWaitDuration"`
}

type jsonOsPolicyAssignmentRollout OsPolicyAssignmentRollout

func (r *OsPolicyAssignmentRollout) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentRollout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentRollout
	} else {

		r.DisruptionBudget = res.DisruptionBudget

		r.MinWaitDuration = res.MinWaitDuration

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentRollout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentRollout *OsPolicyAssignmentRollout = &OsPolicyAssignmentRollout{empty: true}

func (r *OsPolicyAssignmentRollout) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentRollout) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentRollout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OsPolicyAssignmentRolloutDisruptionBudget struct {
	empty   bool   `json:"-"`
	Fixed   *int64 `json:"fixed"`
	Percent *int64 `json:"percent"`
}

type jsonOsPolicyAssignmentRolloutDisruptionBudget OsPolicyAssignmentRolloutDisruptionBudget

func (r *OsPolicyAssignmentRolloutDisruptionBudget) UnmarshalJSON(data []byte) error {
	var res jsonOsPolicyAssignmentRolloutDisruptionBudget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOsPolicyAssignmentRolloutDisruptionBudget
	} else {

		r.Fixed = res.Fixed

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this OsPolicyAssignmentRolloutDisruptionBudget is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOsPolicyAssignmentRolloutDisruptionBudget *OsPolicyAssignmentRolloutDisruptionBudget = &OsPolicyAssignmentRolloutDisruptionBudget{empty: true}

func (r *OsPolicyAssignmentRolloutDisruptionBudget) Empty() bool {
	return r.empty
}

func (r *OsPolicyAssignmentRolloutDisruptionBudget) String() string {
	return dcl.SprintResource(r)
}

func (r *OsPolicyAssignmentRolloutDisruptionBudget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *OsPolicyAssignment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "os_config",
		Type:    "OsPolicyAssignment",
		Version: "alpha",
	}
}

const OsPolicyAssignmentMaxPage = -1

type OsPolicyAssignmentList struct {
	Items []*OsPolicyAssignment

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *OsPolicyAssignmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *OsPolicyAssignmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listOsPolicyAssignment(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListOsPolicyAssignment(ctx context.Context, project, location string) (*OsPolicyAssignmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListOsPolicyAssignmentWithMaxResults(ctx, project, location, OsPolicyAssignmentMaxPage)

}

func (c *Client) ListOsPolicyAssignmentWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*OsPolicyAssignmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listOsPolicyAssignment(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &OsPolicyAssignmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *OsPolicyAssignment) URLNormalized() *OsPolicyAssignment {
	normalized := dcl.Copy(*r).(OsPolicyAssignment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.RevisionId = dcl.SelfLinkToName(r.RevisionId)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetOsPolicyAssignment(ctx context.Context, r *OsPolicyAssignment) (*OsPolicyAssignment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getOsPolicyAssignmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalOsPolicyAssignment(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeOsPolicyAssignmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteOsPolicyAssignment(ctx context.Context, r *OsPolicyAssignment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("OsPolicyAssignment resource is nil")
	}
	c.Config.Logger.Info("Deleting OsPolicyAssignment...")
	deleteOp := deleteOsPolicyAssignmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllOsPolicyAssignment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllOsPolicyAssignment(ctx context.Context, project, location string, filter func(*OsPolicyAssignment) bool) error {
	listObj, err := c.ListOsPolicyAssignment(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllOsPolicyAssignment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllOsPolicyAssignment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyOsPolicyAssignment(ctx context.Context, rawDesired *OsPolicyAssignment, opts ...dcl.ApplyOption) (*OsPolicyAssignment, error) {

	var resultNewState *OsPolicyAssignment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyOsPolicyAssignmentHelper(c, ctx, rawDesired, opts...)
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

func applyOsPolicyAssignmentHelper(c *Client, ctx context.Context, rawDesired *OsPolicyAssignment, opts ...dcl.ApplyOption) (*OsPolicyAssignment, error) {
	c.Config.Logger.Info("Beginning ApplyOsPolicyAssignment...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.osPolicyAssignmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToOsPolicyAssignmentOp(opStrings, fieldDiffs, opts)
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
	var ops []osPolicyAssignmentApiOperation
	if create {
		ops = append(ops, &createOsPolicyAssignmentOperation{})
	} else if recreate {
		ops = append(ops, &deleteOsPolicyAssignmentOperation{})
		ops = append(ops, &createOsPolicyAssignmentOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeOsPolicyAssignmentDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetOsPolicyAssignment(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createOsPolicyAssignmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapOsPolicyAssignment(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeOsPolicyAssignmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeOsPolicyAssignmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeOsPolicyAssignmentDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffOsPolicyAssignment(c, newDesired, newState)
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
