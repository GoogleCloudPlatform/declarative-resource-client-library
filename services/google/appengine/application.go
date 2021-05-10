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
package appengine

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Application struct {
	AuthDomain      *string                       `json:"authDomain"`
	CodeBucket      *string                       `json:"codeBucket"`
	DatabaseType    *ApplicationDatabaseTypeEnum  `json:"databaseType"`
	DefaultBucket   *string                       `json:"defaultBucket"`
	DefaultHostname *string                       `json:"defaultHostname"`
	DispatchRules   []ApplicationDispatchRules    `json:"dispatchRules"`
	FeatureSettings *ApplicationFeatureSettings   `json:"featureSettings"`
	GcrDomain       *string                       `json:"gcrDomain"`
	Iap             *ApplicationIap               `json:"iap"`
	Name            *string                       `json:"name"`
	Location        *string                       `json:"location"`
	ServingStatus   *ApplicationServingStatusEnum `json:"servingStatus"`
}

func (r *Application) String() string {
	return dcl.SprintResource(r)
}

// The enum ApplicationDatabaseTypeEnum.
type ApplicationDatabaseTypeEnum string

// ApplicationDatabaseTypeEnumRef returns a *ApplicationDatabaseTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ApplicationDatabaseTypeEnumRef(s string) *ApplicationDatabaseTypeEnum {
	if s == "" {
		return nil
	}

	v := ApplicationDatabaseTypeEnum(s)
	return &v
}

func (v ApplicationDatabaseTypeEnum) Validate() error {
	for _, s := range []string{"DATABASE_TYPE_UNSPECIFIED", "CLOUD_DATASTORE", "CLOUD_FIRESTORE", "CLOUD_DATASTORE_COMPATIBILITY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ApplicationDatabaseTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ApplicationServingStatusEnum.
type ApplicationServingStatusEnum string

// ApplicationServingStatusEnumRef returns a *ApplicationServingStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func ApplicationServingStatusEnumRef(s string) *ApplicationServingStatusEnum {
	if s == "" {
		return nil
	}

	v := ApplicationServingStatusEnum(s)
	return &v
}

func (v ApplicationServingStatusEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "SERVING", "USER_DISABLED", "SYSTEM_DISABLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ApplicationServingStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ApplicationDispatchRules struct {
	empty   bool    `json:"-"`
	Domain  *string `json:"domain"`
	Path    *string `json:"path"`
	Service *string `json:"service"`
}

type jsonApplicationDispatchRules ApplicationDispatchRules

func (r *ApplicationDispatchRules) UnmarshalJSON(data []byte) error {
	var res jsonApplicationDispatchRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyApplicationDispatchRules
	} else {

		r.Domain = res.Domain

		r.Path = res.Path

		r.Service = res.Service

	}
	return nil
}

// This object is used to assert a desired state where this ApplicationDispatchRules is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyApplicationDispatchRules *ApplicationDispatchRules = &ApplicationDispatchRules{empty: true}

func (r *ApplicationDispatchRules) String() string {
	return dcl.SprintResource(r)
}

func (r *ApplicationDispatchRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ApplicationFeatureSettings struct {
	empty                   bool  `json:"-"`
	SplitHealthChecks       *bool `json:"splitHealthChecks"`
	UseContainerOptimizedOs *bool `json:"useContainerOptimizedOs"`
}

type jsonApplicationFeatureSettings ApplicationFeatureSettings

func (r *ApplicationFeatureSettings) UnmarshalJSON(data []byte) error {
	var res jsonApplicationFeatureSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyApplicationFeatureSettings
	} else {

		r.SplitHealthChecks = res.SplitHealthChecks

		r.UseContainerOptimizedOs = res.UseContainerOptimizedOs

	}
	return nil
}

// This object is used to assert a desired state where this ApplicationFeatureSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyApplicationFeatureSettings *ApplicationFeatureSettings = &ApplicationFeatureSettings{empty: true}

func (r *ApplicationFeatureSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *ApplicationFeatureSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ApplicationIap struct {
	empty                    bool    `json:"-"`
	Enabled                  *bool   `json:"enabled"`
	OAuth2ClientId           *string `json:"oauth2ClientId"`
	OAuth2ClientSecret       *string `json:"oauth2ClientSecret"`
	OAuth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256"`
}

type jsonApplicationIap ApplicationIap

func (r *ApplicationIap) UnmarshalJSON(data []byte) error {
	var res jsonApplicationIap
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyApplicationIap
	} else {

		r.Enabled = res.Enabled

		r.OAuth2ClientId = res.OAuth2ClientId

		r.OAuth2ClientSecret = res.OAuth2ClientSecret

		r.OAuth2ClientSecretSha256 = res.OAuth2ClientSecretSha256

	}
	return nil
}

// This object is used to assert a desired state where this ApplicationIap is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyApplicationIap *ApplicationIap = &ApplicationIap{empty: true}

func (r *ApplicationIap) String() string {
	return dcl.SprintResource(r)
}

func (r *ApplicationIap) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Application) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "app_engine",
		Type:    "Application",
		Version: "appengine",
	}
}

const ApplicationMaxPage = -1

type ApplicationList struct {
	Items []*Application

	nextToken string
}

func (c *Client) GetApplication(ctx context.Context, r *Application) (*Application, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getApplicationRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalApplication(b, c)
	if err != nil {
		return nil, err
	}
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeApplicationNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyApplication(ctx context.Context, rawDesired *Application, opts ...dcl.ApplyOption) (*Application, error) {
	c.Config.Logger.Info("Beginning ApplyApplication...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.applicationDiffsForRawDesired(ctx, rawDesired, opts...)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []applicationApiOperation
	if create {
		ops = append(ops, &createApplicationOperation{})
	} else if recreate {

		ops = append(ops, &createApplicationOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeApplicationDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetApplication(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createApplicationOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapApplication(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeApplicationNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeApplicationNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeApplicationDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffApplication(c, newDesired, newState)
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
