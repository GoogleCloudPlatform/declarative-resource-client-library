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
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type GuestPolicy struct {
	Name                *string                          `json:"name"`
	Description         *string                          `json:"description"`
	CreateTime          *string                          `json:"createTime"`
	UpdateTime          *string                          `json:"updateTime"`
	Assignment          *GuestPolicyAssignment           `json:"assignment"`
	Packages            []GuestPolicyPackages            `json:"packages"`
	PackageRepositories []GuestPolicyPackageRepositories `json:"packageRepositories"`
	Recipes             []GuestPolicyRecipes             `json:"recipes"`
	Etag                *string                          `json:"etag"`
	Project             *string                          `json:"project"`
}

func (r *GuestPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum GuestPolicyPackagesDesiredStateEnum.
type GuestPolicyPackagesDesiredStateEnum string

// GuestPolicyPackagesDesiredStateEnumRef returns a *GuestPolicyPackagesDesiredStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackagesDesiredStateEnumRef(s string) *GuestPolicyPackagesDesiredStateEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyPackagesDesiredStateEnum(s)
	return &v
}

func (v GuestPolicyPackagesDesiredStateEnum) Validate() error {
	for _, s := range []string{"DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackagesDesiredStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyPackagesManagerEnum.
type GuestPolicyPackagesManagerEnum string

// GuestPolicyPackagesManagerEnumRef returns a *GuestPolicyPackagesManagerEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackagesManagerEnumRef(s string) *GuestPolicyPackagesManagerEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyPackagesManagerEnum(s)
	return &v
}

func (v GuestPolicyPackagesManagerEnum) Validate() error {
	for _, s := range []string{"MANAGER_UNSPECIFIED", "ANY", "APT", "YUM", "ZYPPER", "GOO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackagesManagerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyPackageRepositoriesAptArchiveTypeEnum.
type GuestPolicyPackageRepositoriesAptArchiveTypeEnum string

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef returns a *GuestPolicyPackageRepositoriesAptArchiveTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef(s string) *GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyPackageRepositoriesAptArchiveTypeEnum(s)
	return &v
}

func (v GuestPolicyPackageRepositoriesAptArchiveTypeEnum) Validate() error {
	for _, s := range []string{"ARCHIVE_TYPE_UNSPECIFIED", "DEB", "DEB_SRC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackageRepositoriesAptArchiveTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.
type GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum string

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef returns a *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef(s string) *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(s)
	return &v
}

func (v GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "VALIDATION", "DESIRED_STATE_CHECK", "DESIRED_STATE_ENFORCEMENT", "DESIRED_STATE_CHECK_POST_ENFORCEMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.
type GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum string

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef returns a *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef(s string) *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(s)
	return &v
}

func (v GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) Validate() error {
	for _, s := range []string{"INTERPRETER_UNSPECIFIED", "NONE", "SHELL", "POWERSHELL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.
type GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum string

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef returns a *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef(s string) *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(s)
	return &v
}

func (v GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "VALIDATION", "DESIRED_STATE_CHECK", "DESIRED_STATE_ENFORCEMENT", "DESIRED_STATE_CHECK_POST_ENFORCEMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.
type GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum string

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef returns a *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef(s string) *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(s)
	return &v
}

func (v GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) Validate() error {
	for _, s := range []string{"INTERPRETER_UNSPECIFIED", "NONE", "SHELL", "POWERSHELL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesDesiredStateEnum.
type GuestPolicyRecipesDesiredStateEnum string

// GuestPolicyRecipesDesiredStateEnumRef returns a *GuestPolicyRecipesDesiredStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesDesiredStateEnumRef(s string) *GuestPolicyRecipesDesiredStateEnum {
	if s == "" {
		return nil
	}

	v := GuestPolicyRecipesDesiredStateEnum(s)
	return &v
}

func (v GuestPolicyRecipesDesiredStateEnum) Validate() error {
	for _, s := range []string{"DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesDesiredStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type GuestPolicyAssignment struct {
	empty                bool                               `json:"-"`
	GroupLabels          []GuestPolicyAssignmentGroupLabels `json:"groupLabels"`
	Zones                []string                           `json:"zones"`
	Instances            []string                           `json:"instances"`
	InstanceNamePrefixes []string                           `json:"instanceNamePrefixes"`
	OsTypes              []GuestPolicyAssignmentOsTypes     `json:"osTypes"`
}

// This object is used to assert a desired state where this GuestPolicyAssignment is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignment *GuestPolicyAssignment = &GuestPolicyAssignment{empty: true}

func (r *GuestPolicyAssignment) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyAssignmentGroupLabels struct {
	empty  bool              `json:"-"`
	Labels map[string]string `json:"labels"`
}

// This object is used to assert a desired state where this GuestPolicyAssignmentGroupLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignmentGroupLabels *GuestPolicyAssignmentGroupLabels = &GuestPolicyAssignmentGroupLabels{empty: true}

func (r *GuestPolicyAssignmentGroupLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignmentGroupLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyAssignmentOsTypes struct {
	empty          bool    `json:"-"`
	OsShortName    *string `json:"osShortName"`
	OsVersion      *string `json:"osVersion"`
	OsArchitecture *string `json:"osArchitecture"`
}

// This object is used to assert a desired state where this GuestPolicyAssignmentOsTypes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignmentOsTypes *GuestPolicyAssignmentOsTypes = &GuestPolicyAssignmentOsTypes{empty: true}

func (r *GuestPolicyAssignmentOsTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignmentOsTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackages struct {
	empty        bool                                 `json:"-"`
	Name         *string                              `json:"name"`
	DesiredState *GuestPolicyPackagesDesiredStateEnum `json:"desiredState"`
	Manager      *GuestPolicyPackagesManagerEnum      `json:"manager"`
}

// This object is used to assert a desired state where this GuestPolicyPackages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackages *GuestPolicyPackages = &GuestPolicyPackages{empty: true}

func (r *GuestPolicyPackages) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositories struct {
	empty  bool                                  `json:"-"`
	Apt    *GuestPolicyPackageRepositoriesApt    `json:"apt"`
	Yum    *GuestPolicyPackageRepositoriesYum    `json:"yum"`
	Zypper *GuestPolicyPackageRepositoriesZypper `json:"zypper"`
	Goo    *GuestPolicyPackageRepositoriesGoo    `json:"goo"`
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositories is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositories *GuestPolicyPackageRepositories = &GuestPolicyPackageRepositories{empty: true}

func (r *GuestPolicyPackageRepositories) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositories) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesApt struct {
	empty        bool                                              `json:"-"`
	ArchiveType  *GuestPolicyPackageRepositoriesAptArchiveTypeEnum `json:"archiveType"`
	Uri          *string                                           `json:"uri"`
	Distribution *string                                           `json:"distribution"`
	Components   []string                                          `json:"components"`
	GpgKey       *string                                           `json:"gpgKey"`
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesApt is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesApt *GuestPolicyPackageRepositoriesApt = &GuestPolicyPackageRepositoriesApt{empty: true}

func (r *GuestPolicyPackageRepositoriesApt) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesApt) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesYum struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesYum is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesYum *GuestPolicyPackageRepositoriesYum = &GuestPolicyPackageRepositoriesYum{empty: true}

func (r *GuestPolicyPackageRepositoriesYum) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesYum) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesZypper struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesZypper is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesZypper *GuestPolicyPackageRepositoriesZypper = &GuestPolicyPackageRepositoriesZypper{empty: true}

func (r *GuestPolicyPackageRepositoriesZypper) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesZypper) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesGoo struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesGoo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesGoo *GuestPolicyPackageRepositoriesGoo = &GuestPolicyPackageRepositoriesGoo{empty: true}

func (r *GuestPolicyPackageRepositoriesGoo) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesGoo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipes struct {
	empty        bool                                `json:"-"`
	Name         *string                             `json:"name"`
	Version      *string                             `json:"version"`
	Artifacts    []GuestPolicyRecipesArtifacts       `json:"artifacts"`
	InstallSteps []GuestPolicyRecipesInstallSteps    `json:"installSteps"`
	UpdateSteps  []GuestPolicyRecipesUpdateSteps     `json:"updateSteps"`
	DesiredState *GuestPolicyRecipesDesiredStateEnum `json:"desiredState"`
}

// This object is used to assert a desired state where this GuestPolicyRecipes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipes *GuestPolicyRecipes = &GuestPolicyRecipes{empty: true}

func (r *GuestPolicyRecipes) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifacts struct {
	empty         bool                               `json:"-"`
	Id            *string                            `json:"id"`
	Remote        *GuestPolicyRecipesArtifactsRemote `json:"remote"`
	Gcs           *GuestPolicyRecipesArtifactsGcs    `json:"gcs"`
	AllowInsecure *bool                              `json:"allowInsecure"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifacts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifacts *GuestPolicyRecipesArtifacts = &GuestPolicyRecipesArtifacts{empty: true}

func (r *GuestPolicyRecipesArtifacts) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifacts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifactsRemote struct {
	empty    bool    `json:"-"`
	Uri      *string `json:"uri"`
	Checksum *string `json:"checksum"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifactsRemote is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifactsRemote *GuestPolicyRecipesArtifactsRemote = &GuestPolicyRecipesArtifactsRemote{empty: true}

func (r *GuestPolicyRecipesArtifactsRemote) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifactsRemote) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifactsGcs struct {
	empty      bool    `json:"-"`
	Bucket     *string `json:"bucket"`
	Object     *string `json:"object"`
	Generation *int64  `json:"generation"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifactsGcs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifactsGcs *GuestPolicyRecipesArtifactsGcs = &GuestPolicyRecipesArtifactsGcs{empty: true}

func (r *GuestPolicyRecipesArtifactsGcs) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifactsGcs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallSteps struct {
	empty             bool                                             `json:"-"`
	FileCopy          *GuestPolicyRecipesInstallStepsFileCopy          `json:"fileCopy"`
	ArchiveExtraction *GuestPolicyRecipesInstallStepsArchiveExtraction `json:"archiveExtraction"`
	MsiInstallation   *GuestPolicyRecipesInstallStepsMsiInstallation   `json:"msiInstallation"`
	DpkgInstallation  *GuestPolicyRecipesInstallStepsDpkgInstallation  `json:"dpkgInstallation"`
	RpmInstallation   *GuestPolicyRecipesInstallStepsRpmInstallation   `json:"rpmInstallation"`
	FileExec          *GuestPolicyRecipesInstallStepsFileExec          `json:"fileExec"`
	ScriptRun         *GuestPolicyRecipesInstallStepsScriptRun         `json:"scriptRun"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallSteps *GuestPolicyRecipesInstallSteps = &GuestPolicyRecipesInstallSteps{empty: true}

func (r *GuestPolicyRecipesInstallSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsFileCopy struct {
	empty       bool    `json:"-"`
	ArtifactId  *string `json:"artifactId"`
	Destination *string `json:"destination"`
	Overwrite   *bool   `json:"overwrite"`
	Permissions *string `json:"permissions"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsFileCopy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsFileCopy *GuestPolicyRecipesInstallStepsFileCopy = &GuestPolicyRecipesInstallStepsFileCopy{empty: true}

func (r *GuestPolicyRecipesInstallStepsFileCopy) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsFileCopy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsArchiveExtraction struct {
	empty       bool                                                     `json:"-"`
	ArtifactId  *string                                                  `json:"artifactId"`
	Destination *string                                                  `json:"destination"`
	Type        *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsArchiveExtraction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsArchiveExtraction *GuestPolicyRecipesInstallStepsArchiveExtraction = &GuestPolicyRecipesInstallStepsArchiveExtraction{empty: true}

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsMsiInstallation struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	Flags            []string `json:"flags"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsMsiInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsMsiInstallation *GuestPolicyRecipesInstallStepsMsiInstallation = &GuestPolicyRecipesInstallStepsMsiInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsDpkgInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsDpkgInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsDpkgInstallation *GuestPolicyRecipesInstallStepsDpkgInstallation = &GuestPolicyRecipesInstallStepsDpkgInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsRpmInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsRpmInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsRpmInstallation *GuestPolicyRecipesInstallStepsRpmInstallation = &GuestPolicyRecipesInstallStepsRpmInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsFileExec struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	LocalPath        *string  `json:"localPath"`
	Args             []string `json:"args"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsFileExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsFileExec *GuestPolicyRecipesInstallStepsFileExec = &GuestPolicyRecipesInstallStepsFileExec{empty: true}

func (r *GuestPolicyRecipesInstallStepsFileExec) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsFileExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsScriptRun struct {
	empty            bool                                                    `json:"-"`
	Script           *string                                                 `json:"script"`
	AllowedExitCodes []int64                                                 `json:"allowedExitCodes"`
	Interpreter      *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum `json:"interpreter"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsScriptRun is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsScriptRun *GuestPolicyRecipesInstallStepsScriptRun = &GuestPolicyRecipesInstallStepsScriptRun{empty: true}

func (r *GuestPolicyRecipesInstallStepsScriptRun) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsScriptRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateSteps struct {
	empty             bool                                            `json:"-"`
	FileCopy          *GuestPolicyRecipesUpdateStepsFileCopy          `json:"fileCopy"`
	ArchiveExtraction *GuestPolicyRecipesUpdateStepsArchiveExtraction `json:"archiveExtraction"`
	MsiInstallation   *GuestPolicyRecipesUpdateStepsMsiInstallation   `json:"msiInstallation"`
	DpkgInstallation  *GuestPolicyRecipesUpdateStepsDpkgInstallation  `json:"dpkgInstallation"`
	RpmInstallation   *GuestPolicyRecipesUpdateStepsRpmInstallation   `json:"rpmInstallation"`
	FileExec          *GuestPolicyRecipesUpdateStepsFileExec          `json:"fileExec"`
	ScriptRun         *GuestPolicyRecipesUpdateStepsScriptRun         `json:"scriptRun"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateSteps *GuestPolicyRecipesUpdateSteps = &GuestPolicyRecipesUpdateSteps{empty: true}

func (r *GuestPolicyRecipesUpdateSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsFileCopy struct {
	empty       bool    `json:"-"`
	ArtifactId  *string `json:"artifactId"`
	Destination *string `json:"destination"`
	Overwrite   *bool   `json:"overwrite"`
	Permissions *string `json:"permissions"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsFileCopy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsFileCopy *GuestPolicyRecipesUpdateStepsFileCopy = &GuestPolicyRecipesUpdateStepsFileCopy{empty: true}

func (r *GuestPolicyRecipesUpdateStepsFileCopy) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsFileCopy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsArchiveExtraction struct {
	empty       bool                                                    `json:"-"`
	ArtifactId  *string                                                 `json:"artifactId"`
	Destination *string                                                 `json:"destination"`
	Type        *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsArchiveExtraction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction *GuestPolicyRecipesUpdateStepsArchiveExtraction = &GuestPolicyRecipesUpdateStepsArchiveExtraction{empty: true}

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsMsiInstallation struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	Flags            []string `json:"flags"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsMsiInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsMsiInstallation *GuestPolicyRecipesUpdateStepsMsiInstallation = &GuestPolicyRecipesUpdateStepsMsiInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsDpkgInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsDpkgInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation *GuestPolicyRecipesUpdateStepsDpkgInstallation = &GuestPolicyRecipesUpdateStepsDpkgInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsRpmInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsRpmInstallation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsRpmInstallation *GuestPolicyRecipesUpdateStepsRpmInstallation = &GuestPolicyRecipesUpdateStepsRpmInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsFileExec struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	LocalPath        *string  `json:"localPath"`
	Args             []string `json:"args"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsFileExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsFileExec *GuestPolicyRecipesUpdateStepsFileExec = &GuestPolicyRecipesUpdateStepsFileExec{empty: true}

func (r *GuestPolicyRecipesUpdateStepsFileExec) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsFileExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsScriptRun struct {
	empty            bool                                                   `json:"-"`
	Script           *string                                                `json:"script"`
	AllowedExitCodes []int64                                                `json:"allowedExitCodes"`
	Interpreter      *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum `json:"interpreter"`
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsScriptRun is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsScriptRun *GuestPolicyRecipesUpdateStepsScriptRun = &GuestPolicyRecipesUpdateStepsScriptRun{empty: true}

func (r *GuestPolicyRecipesUpdateStepsScriptRun) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsScriptRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *GuestPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "osconfig",
		Type:    "GuestPolicy",
		Version: "beta",
	}
}

const GuestPolicyMaxPage = -1

type GuestPolicyList struct {
	Items []*GuestPolicy

	nextToken string

	pageSize int32

	project string
}

func (l *GuestPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *GuestPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listGuestPolicy(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListGuestPolicy(ctx context.Context, project string) (*GuestPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListGuestPolicyWithMaxResults(ctx, project, GuestPolicyMaxPage)

}

func (c *Client) ListGuestPolicyWithMaxResults(ctx context.Context, project string, pageSize int32) (*GuestPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listGuestPolicy(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &GuestPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetGuestPolicy(ctx context.Context, r *GuestPolicy) (*GuestPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getGuestPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalGuestPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeGuestPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteGuestPolicy(ctx context.Context, r *GuestPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("GuestPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting GuestPolicy...")
	deleteOp := deleteGuestPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllGuestPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllGuestPolicy(ctx context.Context, project string, filter func(*GuestPolicy) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListGuestPolicy(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllGuestPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllGuestPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyGuestPolicy(ctx context.Context, rawDesired *GuestPolicy, opts ...dcl.ApplyOption) (*GuestPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyGuestPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.guestPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
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
	var ops []guestPolicyApiOperation
	if create {
		ops = append(ops, &createGuestPolicyOperation{})
	} else if recreate {

		ops = append(ops, &deleteGuestPolicyOperation{})

		ops = append(ops, &createGuestPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeGuestPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetGuestPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createGuestPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapGuestPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeGuestPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeGuestPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeGuestPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffGuestPolicy(c, newDesired, newState)
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
