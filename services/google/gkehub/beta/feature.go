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

type Feature struct {
	Name             *string               `json:"name"`
	Labels           map[string]string     `json:"labels"`
	ResourceState    *FeatureResourceState `json:"resourceState"`
	Spec             *FeatureSpec          `json:"spec"`
	MembershipSpecs  map[string]string     `json:"membershipSpecs"`
	State            *FeatureState         `json:"state"`
	MembershipStates map[string]string     `json:"membershipStates"`
	CreateTime       *string               `json:"createTime"`
	UpdateTime       *string               `json:"updateTime"`
	DeleteTime       *string               `json:"deleteTime"`
	Project          *string               `json:"project"`
	Location         *string               `json:"location"`
}

func (r *Feature) String() string {
	return dcl.SprintResource(r)
}

// The enum FeatureResourceStateStateEnum.
type FeatureResourceStateStateEnum string

// FeatureResourceStateStateEnumRef returns a *FeatureResourceStateStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureResourceStateStateEnumRef(s string) *FeatureResourceStateStateEnum {
	if s == "" {
		return nil
	}

	v := FeatureResourceStateStateEnum(s)
	return &v
}

func (v FeatureResourceStateStateEnum) Validate() error {
	for _, s := range []string{"DEPLOYMENT_STATE_UNSPECIFIED", "NOT_INSTALLED", "INSTALLED", "ERROR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureResourceStateStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureSpecMulticlusteringressBillingEnum.
type FeatureSpecMulticlusteringressBillingEnum string

// FeatureSpecMulticlusteringressBillingEnumRef returns a *FeatureSpecMulticlusteringressBillingEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureSpecMulticlusteringressBillingEnumRef(s string) *FeatureSpecMulticlusteringressBillingEnum {
	if s == "" {
		return nil
	}

	v := FeatureSpecMulticlusteringressBillingEnum(s)
	return &v
}

func (v FeatureSpecMulticlusteringressBillingEnum) Validate() error {
	for _, s := range []string{"BILLING_UNSPECIFIED", "PAY_AS_YOU_GO", "ANTHOS_LICENSE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureSpecMulticlusteringressBillingEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureSpecHelloworldFeatureTestThirdEnum.
type FeatureSpecHelloworldFeatureTestThirdEnum string

// FeatureSpecHelloworldFeatureTestThirdEnumRef returns a *FeatureSpecHelloworldFeatureTestThirdEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureSpecHelloworldFeatureTestThirdEnumRef(s string) *FeatureSpecHelloworldFeatureTestThirdEnum {
	if s == "" {
		return nil
	}

	v := FeatureSpecHelloworldFeatureTestThirdEnum(s)
	return &v
}

func (v FeatureSpecHelloworldFeatureTestThirdEnum) Validate() error {
	for _, s := range []string{"BAR_UNSPECIFIED", "FIRST", "SECOND"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureSpecHelloworldFeatureTestThirdEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureStateStateCodeEnum.
type FeatureStateStateCodeEnum string

// FeatureStateStateCodeEnumRef returns a *FeatureStateStateCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureStateStateCodeEnumRef(s string) *FeatureStateStateCodeEnum {
	if s == "" {
		return nil
	}

	v := FeatureStateStateCodeEnum(s)
	return &v
}

func (v FeatureStateStateCodeEnum) Validate() error {
	for _, s := range []string{"CODE_UNSPECIFIED", "CREATING", "READY", "DELETING", "UPDATING", "SERVICE_UPDATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureStateStateCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FeatureResourceState struct {
	empty bool                           `json:"-"`
	State *FeatureResourceStateStateEnum `json:"state"`
}

// This object is used to assert a desired state where this FeatureResourceState is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureResourceState *FeatureResourceState = &FeatureResourceState{empty: true}

func (r *FeatureResourceState) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureResourceState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpec struct {
	empty               bool                            `json:"-"`
	Multiclusteringress *FeatureSpecMulticlusteringress `json:"multiclusteringress"`
	Helloworld          *FeatureSpecHelloworld          `json:"helloworld"`
}

// This object is used to assert a desired state where this FeatureSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpec *FeatureSpec = &FeatureSpec{empty: true}

func (r *FeatureSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpecMulticlusteringress struct {
	empty            bool                                       `json:"-"`
	ConfigMembership *string                                    `json:"configMembership"`
	Billing          *FeatureSpecMulticlusteringressBillingEnum `json:"billing"`
}

// This object is used to assert a desired state where this FeatureSpecMulticlusteringress is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpecMulticlusteringress *FeatureSpecMulticlusteringress = &FeatureSpecMulticlusteringress{empty: true}

func (r *FeatureSpecMulticlusteringress) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpecMulticlusteringress) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpecHelloworld struct {
	empty        bool                              `json:"-"`
	FeatureTest  *FeatureSpecHelloworldFeatureTest `json:"featureTest"`
	CustomConfig *string                           `json:"customConfig"`
}

// This object is used to assert a desired state where this FeatureSpecHelloworld is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpecHelloworld *FeatureSpecHelloworld = &FeatureSpecHelloworld{empty: true}

func (r *FeatureSpecHelloworld) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpecHelloworld) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpecHelloworldFeatureTest struct {
	empty   bool                                       `json:"-"`
	First   *string                                    `json:"first"`
	Second  *int64                                     `json:"second"`
	Third   *FeatureSpecHelloworldFeatureTestThirdEnum `json:"third"`
	Fourth  *string                                    `json:"fourth"`
	Fifth   *FeatureSpecHelloworldFeatureTestFifth     `json:"fifth"`
	Sixth   *int64                                     `json:"sixth"`
	Seventh *string                                    `json:"seventh"`
	Eighth  []FeatureSpecHelloworldFeatureTestEighth   `json:"eighth"`
	Ninth   map[string]string                          `json:"ninth"`
}

// This object is used to assert a desired state where this FeatureSpecHelloworldFeatureTest is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpecHelloworldFeatureTest *FeatureSpecHelloworldFeatureTest = &FeatureSpecHelloworldFeatureTest{empty: true}

func (r *FeatureSpecHelloworldFeatureTest) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpecHelloworldFeatureTest) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpecHelloworldFeatureTestFifth struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this FeatureSpecHelloworldFeatureTestFifth is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpecHelloworldFeatureTestFifth *FeatureSpecHelloworldFeatureTestFifth = &FeatureSpecHelloworldFeatureTestFifth{empty: true}

func (r *FeatureSpecHelloworldFeatureTestFifth) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpecHelloworldFeatureTestFifth) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureSpecHelloworldFeatureTestEighth struct {
	empty  bool    `json:"-"`
	First  *string `json:"first"`
	Second *int64  `json:"second"`
}

// This object is used to assert a desired state where this FeatureSpecHelloworldFeatureTestEighth is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureSpecHelloworldFeatureTestEighth *FeatureSpecHelloworldFeatureTestEighth = &FeatureSpecHelloworldFeatureTestEighth{empty: true}

func (r *FeatureSpecHelloworldFeatureTestEighth) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureSpecHelloworldFeatureTestEighth) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureState struct {
	empty      bool                    `json:"-"`
	State      *FeatureStateState      `json:"state"`
	Helloworld *FeatureStateHelloworld `json:"helloworld"`
}

// This object is used to assert a desired state where this FeatureState is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureState *FeatureState = &FeatureState{empty: true}

func (r *FeatureState) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureStateState struct {
	empty       bool                       `json:"-"`
	Code        *FeatureStateStateCodeEnum `json:"code"`
	Description *string                    `json:"description"`
	UpdateTime  *string                    `json:"updateTime"`
}

// This object is used to assert a desired state where this FeatureStateState is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureStateState *FeatureStateState = &FeatureStateState{empty: true}

func (r *FeatureStateState) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureStateState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureStateHelloworld struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this FeatureStateHelloworld is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFeatureStateHelloworld *FeatureStateHelloworld = &FeatureStateHelloworld{empty: true}

func (r *FeatureStateHelloworld) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureStateHelloworld) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Feature) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gkehub",
		Type:    "Feature",
		Version: "beta",
	}
}

const FeatureMaxPage = -1

type FeatureList struct {
	Items []*Feature

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *FeatureList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FeatureList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFeature(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFeature(ctx context.Context, project, location string) (*FeatureList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListFeatureWithMaxResults(ctx, project, location, FeatureMaxPage)

}

func (c *Client) ListFeatureWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*FeatureList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listFeature(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FeatureList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetFeature(ctx context.Context, r *Feature) (*Feature, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getFeatureRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFeature(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFeatureNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFeature(ctx context.Context, r *Feature) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Feature resource is nil")
	}
	c.Config.Logger.Info("Deleting Feature...")
	deleteOp := deleteFeatureOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFeature deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFeature(ctx context.Context, project, location string, filter func(*Feature) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListFeature(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllFeature(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFeature(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFeature(ctx context.Context, rawDesired *Feature, opts ...dcl.ApplyOption) (*Feature, error) {
	c.Config.Logger.Info("Beginning ApplyFeature...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.featureDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []featureApiOperation
	if create {
		ops = append(ops, &createFeatureOperation{})
	} else if recreate {

		ops = append(ops, &deleteFeatureOperation{})

		ops = append(ops, &createFeatureOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFeatureDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetFeature(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFeatureOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFeature(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFeatureNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFeatureNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFeatureDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFeature(c, newDesired, newState)
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
