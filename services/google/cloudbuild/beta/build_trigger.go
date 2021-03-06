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
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type BuildTrigger struct {
	Name            *string                      `json:"name"`
	Description     *string                      `json:"description"`
	Tags            []string                     `json:"tags"`
	Disabled        *bool                        `json:"disabled"`
	Substitutions   map[string]string            `json:"substitutions"`
	Filename        *string                      `json:"filename"`
	IgnoredFiles    []string                     `json:"ignoredFiles"`
	IncludedFiles   []string                     `json:"includedFiles"`
	TriggerTemplate *BuildTriggerTriggerTemplate `json:"triggerTemplate"`
	Github          *BuildTriggerGithub          `json:"github"`
	Project         *string                      `json:"project"`
	Build           *BuildTriggerBuild           `json:"build"`
	Id              *string                      `json:"id"`
	CreateTime      *string                      `json:"createTime"`
}

func (r *BuildTrigger) String() string {
	return dcl.SprintResource(r)
}

// The enum BuildTriggerGithubPullRequestCommentControlEnum.
type BuildTriggerGithubPullRequestCommentControlEnum string

// BuildTriggerGithubPullRequestCommentControlEnumRef returns a *BuildTriggerGithubPullRequestCommentControlEnum with the value of string s
// If the empty string is provided, nil is returned.
func BuildTriggerGithubPullRequestCommentControlEnumRef(s string) *BuildTriggerGithubPullRequestCommentControlEnum {
	if s == "" {
		return nil
	}

	v := BuildTriggerGithubPullRequestCommentControlEnum(s)
	return &v
}

func (v BuildTriggerGithubPullRequestCommentControlEnum) Validate() error {
	for _, s := range []string{"COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BuildTriggerGithubPullRequestCommentControlEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BuildTriggerBuildStepsStatusEnum.
type BuildTriggerBuildStepsStatusEnum string

// BuildTriggerBuildStepsStatusEnumRef returns a *BuildTriggerBuildStepsStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func BuildTriggerBuildStepsStatusEnumRef(s string) *BuildTriggerBuildStepsStatusEnum {
	if s == "" {
		return nil
	}

	v := BuildTriggerBuildStepsStatusEnum(s)
	return &v
}

func (v BuildTriggerBuildStepsStatusEnum) Validate() error {
	for _, s := range []string{"STATUS_UNKNOWN", "QUEUED", "WORKING", "SUCCESS", "FAILURE", "INTERNAL_ERROR", "TIMEOUT", "CANCELLED", "EXPIRED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BuildTriggerBuildStepsStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type BuildTriggerTriggerTemplate struct {
	empty       bool    `json:"-"`
	ProjectId   *string `json:"projectId"`
	RepoName    *string `json:"repoName"`
	BranchName  *string `json:"branchName"`
	TagName     *string `json:"tagName"`
	CommitSha   *string `json:"commitSha"`
	Dir         *string `json:"dir"`
	InvertRegex *bool   `json:"invertRegex"`
}

type jsonBuildTriggerTriggerTemplate BuildTriggerTriggerTemplate

func (r *BuildTriggerTriggerTemplate) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerTriggerTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerTriggerTemplate
	} else {

		r.ProjectId = res.ProjectId

		r.RepoName = res.RepoName

		r.BranchName = res.BranchName

		r.TagName = res.TagName

		r.CommitSha = res.CommitSha

		r.Dir = res.Dir

		r.InvertRegex = res.InvertRegex

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerTriggerTemplate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerTriggerTemplate *BuildTriggerTriggerTemplate = &BuildTriggerTriggerTemplate{empty: true}

func (r *BuildTriggerTriggerTemplate) Empty() bool {
	return r.empty
}

func (r *BuildTriggerTriggerTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerTriggerTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerGithub struct {
	empty       bool                           `json:"-"`
	Owner       *string                        `json:"owner"`
	Name        *string                        `json:"name"`
	PullRequest *BuildTriggerGithubPullRequest `json:"pullRequest"`
	Push        *BuildTriggerGithubPush        `json:"push"`
}

type jsonBuildTriggerGithub BuildTriggerGithub

func (r *BuildTriggerGithub) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerGithub
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerGithub
	} else {

		r.Owner = res.Owner

		r.Name = res.Name

		r.PullRequest = res.PullRequest

		r.Push = res.Push

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerGithub is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerGithub *BuildTriggerGithub = &BuildTriggerGithub{empty: true}

func (r *BuildTriggerGithub) Empty() bool {
	return r.empty
}

func (r *BuildTriggerGithub) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerGithub) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerGithubPullRequest struct {
	empty          bool                                             `json:"-"`
	Branch         *string                                          `json:"branch"`
	CommentControl *BuildTriggerGithubPullRequestCommentControlEnum `json:"commentControl"`
	InvertRegex    *bool                                            `json:"invertRegex"`
}

type jsonBuildTriggerGithubPullRequest BuildTriggerGithubPullRequest

func (r *BuildTriggerGithubPullRequest) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerGithubPullRequest
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerGithubPullRequest
	} else {

		r.Branch = res.Branch

		r.CommentControl = res.CommentControl

		r.InvertRegex = res.InvertRegex

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerGithubPullRequest is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerGithubPullRequest *BuildTriggerGithubPullRequest = &BuildTriggerGithubPullRequest{empty: true}

func (r *BuildTriggerGithubPullRequest) Empty() bool {
	return r.empty
}

func (r *BuildTriggerGithubPullRequest) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerGithubPullRequest) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerGithubPush struct {
	empty       bool    `json:"-"`
	Branch      *string `json:"branch"`
	Tag         *string `json:"tag"`
	InvertRegex *bool   `json:"invertRegex"`
}

type jsonBuildTriggerGithubPush BuildTriggerGithubPush

func (r *BuildTriggerGithubPush) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerGithubPush
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerGithubPush
	} else {

		r.Branch = res.Branch

		r.Tag = res.Tag

		r.InvertRegex = res.InvertRegex

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerGithubPush is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerGithubPush *BuildTriggerGithubPush = &BuildTriggerGithubPush{empty: true}

func (r *BuildTriggerGithubPush) Empty() bool {
	return r.empty
}

func (r *BuildTriggerGithubPush) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerGithubPush) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuild struct {
	empty         bool                       `json:"-"`
	Tags          []string                   `json:"tags"`
	Images        []string                   `json:"images"`
	Substitutions map[string]string          `json:"substitutions"`
	QueueTtl      *string                    `json:"queueTtl"`
	LogsBucket    *string                    `json:"logsBucket"`
	Timeout       *string                    `json:"timeout"`
	Secrets       []BuildTriggerBuildSecrets `json:"secrets"`
	Steps         []BuildTriggerBuildSteps   `json:"steps"`
	Source        *BuildTriggerBuildSource   `json:"source"`
}

type jsonBuildTriggerBuild BuildTriggerBuild

func (r *BuildTriggerBuild) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuild
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuild
	} else {

		r.Tags = res.Tags

		r.Images = res.Images

		r.Substitutions = res.Substitutions

		r.QueueTtl = res.QueueTtl

		r.LogsBucket = res.LogsBucket

		r.Timeout = res.Timeout

		r.Secrets = res.Secrets

		r.Steps = res.Steps

		r.Source = res.Source

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuild is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuild *BuildTriggerBuild = &BuildTriggerBuild{empty: true}

func (r *BuildTriggerBuild) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuild) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuild) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildSecrets struct {
	empty      bool              `json:"-"`
	KmsKeyName *string           `json:"kmsKeyName"`
	SecretEnv  map[string]string `json:"secretEnv"`
}

type jsonBuildTriggerBuildSecrets BuildTriggerBuildSecrets

func (r *BuildTriggerBuildSecrets) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildSecrets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildSecrets
	} else {

		r.KmsKeyName = res.KmsKeyName

		r.SecretEnv = res.SecretEnv

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildSecrets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildSecrets *BuildTriggerBuildSecrets = &BuildTriggerBuildSecrets{empty: true}

func (r *BuildTriggerBuildSecrets) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildSecrets) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildSecrets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildSteps struct {
	empty      bool                              `json:"-"`
	Name       *string                           `json:"name"`
	Env        []string                          `json:"env"`
	Args       []string                          `json:"args"`
	Dir        *string                           `json:"dir"`
	Id         *string                           `json:"id"`
	WaitFor    []string                          `json:"waitFor"`
	Entrypoint *string                           `json:"entrypoint"`
	SecretEnv  []string                          `json:"secretEnv"`
	Volumes    []BuildTriggerBuildStepsVolumes   `json:"volumes"`
	Timing     *BuildTriggerBuildStepsTiming     `json:"timing"`
	PullTiming *BuildTriggerBuildStepsPullTiming `json:"pullTiming"`
	Timeout    *string                           `json:"timeout"`
	Status     *BuildTriggerBuildStepsStatusEnum `json:"status"`
}

type jsonBuildTriggerBuildSteps BuildTriggerBuildSteps

func (r *BuildTriggerBuildSteps) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildSteps
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildSteps
	} else {

		r.Name = res.Name

		r.Env = res.Env

		r.Args = res.Args

		r.Dir = res.Dir

		r.Id = res.Id

		r.WaitFor = res.WaitFor

		r.Entrypoint = res.Entrypoint

		r.SecretEnv = res.SecretEnv

		r.Volumes = res.Volumes

		r.Timing = res.Timing

		r.PullTiming = res.PullTiming

		r.Timeout = res.Timeout

		r.Status = res.Status

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildSteps *BuildTriggerBuildSteps = &BuildTriggerBuildSteps{empty: true}

func (r *BuildTriggerBuildSteps) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildStepsVolumes struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Path  *string `json:"path"`
}

type jsonBuildTriggerBuildStepsVolumes BuildTriggerBuildStepsVolumes

func (r *BuildTriggerBuildStepsVolumes) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildStepsVolumes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildStepsVolumes
	} else {

		r.Name = res.Name

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildStepsVolumes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildStepsVolumes *BuildTriggerBuildStepsVolumes = &BuildTriggerBuildStepsVolumes{empty: true}

func (r *BuildTriggerBuildStepsVolumes) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildStepsVolumes) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildStepsVolumes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildStepsTiming struct {
	empty     bool    `json:"-"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

type jsonBuildTriggerBuildStepsTiming BuildTriggerBuildStepsTiming

func (r *BuildTriggerBuildStepsTiming) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildStepsTiming
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildStepsTiming
	} else {

		r.StartTime = res.StartTime

		r.EndTime = res.EndTime

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildStepsTiming is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildStepsTiming *BuildTriggerBuildStepsTiming = &BuildTriggerBuildStepsTiming{empty: true}

func (r *BuildTriggerBuildStepsTiming) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildStepsTiming) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildStepsTiming) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildStepsPullTiming struct {
	empty     bool    `json:"-"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

type jsonBuildTriggerBuildStepsPullTiming BuildTriggerBuildStepsPullTiming

func (r *BuildTriggerBuildStepsPullTiming) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildStepsPullTiming
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildStepsPullTiming
	} else {

		r.StartTime = res.StartTime

		r.EndTime = res.EndTime

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildStepsPullTiming is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildStepsPullTiming *BuildTriggerBuildStepsPullTiming = &BuildTriggerBuildStepsPullTiming{empty: true}

func (r *BuildTriggerBuildStepsPullTiming) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildStepsPullTiming) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildStepsPullTiming) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildSource struct {
	empty         bool                                  `json:"-"`
	StorageSource *BuildTriggerBuildSourceStorageSource `json:"storageSource"`
	RepoSource    *BuildTriggerBuildSourceRepoSource    `json:"repoSource"`
}

type jsonBuildTriggerBuildSource BuildTriggerBuildSource

func (r *BuildTriggerBuildSource) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildSource
	} else {

		r.StorageSource = res.StorageSource

		r.RepoSource = res.RepoSource

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildSource is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildSource *BuildTriggerBuildSource = &BuildTriggerBuildSource{empty: true}

func (r *BuildTriggerBuildSource) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildSource) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildSourceStorageSource struct {
	empty      bool    `json:"-"`
	Bucket     *string `json:"bucket"`
	Object     *string `json:"object"`
	Generation *string `json:"generation"`
}

type jsonBuildTriggerBuildSourceStorageSource BuildTriggerBuildSourceStorageSource

func (r *BuildTriggerBuildSourceStorageSource) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildSourceStorageSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildSourceStorageSource
	} else {

		r.Bucket = res.Bucket

		r.Object = res.Object

		r.Generation = res.Generation

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildSourceStorageSource is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildSourceStorageSource *BuildTriggerBuildSourceStorageSource = &BuildTriggerBuildSourceStorageSource{empty: true}

func (r *BuildTriggerBuildSourceStorageSource) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildSourceStorageSource) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildSourceStorageSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BuildTriggerBuildSourceRepoSource struct {
	empty         bool              `json:"-"`
	ProjectId     *string           `json:"projectId"`
	RepoName      *string           `json:"repoName"`
	BranchName    *string           `json:"branchName"`
	TagName       *string           `json:"tagName"`
	CommitSha     *string           `json:"commitSha"`
	Dir           *string           `json:"dir"`
	InvertRegex   *bool             `json:"invertRegex"`
	Substitutions map[string]string `json:"substitutions"`
}

type jsonBuildTriggerBuildSourceRepoSource BuildTriggerBuildSourceRepoSource

func (r *BuildTriggerBuildSourceRepoSource) UnmarshalJSON(data []byte) error {
	var res jsonBuildTriggerBuildSourceRepoSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBuildTriggerBuildSourceRepoSource
	} else {

		r.ProjectId = res.ProjectId

		r.RepoName = res.RepoName

		r.BranchName = res.BranchName

		r.TagName = res.TagName

		r.CommitSha = res.CommitSha

		r.Dir = res.Dir

		r.InvertRegex = res.InvertRegex

		r.Substitutions = res.Substitutions

	}
	return nil
}

// This object is used to assert a desired state where this BuildTriggerBuildSourceRepoSource is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBuildTriggerBuildSourceRepoSource *BuildTriggerBuildSourceRepoSource = &BuildTriggerBuildSourceRepoSource{empty: true}

func (r *BuildTriggerBuildSourceRepoSource) Empty() bool {
	return r.empty
}

func (r *BuildTriggerBuildSourceRepoSource) String() string {
	return dcl.SprintResource(r)
}

func (r *BuildTriggerBuildSourceRepoSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *BuildTrigger) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloud_build",
		Type:    "BuildTrigger",
		Version: "beta",
	}
}

const BuildTriggerMaxPage = -1

type BuildTriggerList struct {
	Items []*BuildTrigger

	nextToken string

	pageSize int32

	project string
}

func (l *BuildTriggerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BuildTriggerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBuildTrigger(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBuildTrigger(ctx context.Context, project string) (*BuildTriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListBuildTriggerWithMaxResults(ctx, project, BuildTriggerMaxPage)

}

func (c *Client) ListBuildTriggerWithMaxResults(ctx context.Context, project string, pageSize int32) (*BuildTriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listBuildTrigger(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BuildTriggerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *BuildTrigger) URLNormalized() *BuildTrigger {
	normalized := dcl.Copy(*r).(BuildTrigger)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Filename = dcl.SelfLinkToName(r.Filename)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.CreateTime = dcl.SelfLinkToName(r.CreateTime)
	return &normalized
}

func (c *Client) GetBuildTrigger(ctx context.Context, r *BuildTrigger) (*BuildTrigger, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getBuildTriggerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBuildTrigger(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBuildTriggerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBuildTrigger(ctx context.Context, r *BuildTrigger) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("BuildTrigger resource is nil")
	}
	c.Config.Logger.Info("Deleting BuildTrigger...")
	deleteOp := deleteBuildTriggerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBuildTrigger deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBuildTrigger(ctx context.Context, project string, filter func(*BuildTrigger) bool) error {
	listObj, err := c.ListBuildTrigger(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllBuildTrigger(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBuildTrigger(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBuildTrigger(ctx context.Context, rawDesired *BuildTrigger, opts ...dcl.ApplyOption) (*BuildTrigger, error) {
	var resultNewState *BuildTrigger
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyBuildTriggerHelper(c, ctx, rawDesired, opts...)
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

func applyBuildTriggerHelper(c *Client, ctx context.Context, rawDesired *BuildTrigger, opts ...dcl.ApplyOption) (*BuildTrigger, error) {
	c.Config.Logger.Info("Beginning ApplyBuildTrigger...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.buildTriggerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToBuildTriggerDiffs(c.Config, fieldDiffs, opts)
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
	var ops []buildTriggerApiOperation
	if create {
		ops = append(ops, &createBuildTriggerOperation{})
	} else if recreate {
		ops = append(ops, &deleteBuildTriggerOperation{})
		ops = append(ops, &createBuildTriggerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeBuildTriggerDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetBuildTrigger(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBuildTriggerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBuildTrigger(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBuildTriggerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBuildTriggerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBuildTriggerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBuildTrigger(c, newDesired, newState)
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
