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
package monitoring

import (
	"context"
	"crypto/sha256"
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type UptimeCheckConfig struct {
	Name              *string                             `json:"name"`
	DisplayName       *string                             `json:"displayName"`
	MonitoredResource *UptimeCheckConfigMonitoredResource `json:"monitoredResource"`
	ResourceGroup     *UptimeCheckConfigResourceGroup     `json:"resourceGroup"`
	HttpCheck         *UptimeCheckConfigHttpCheck         `json:"httpCheck"`
	TcpCheck          *UptimeCheckConfigTcpCheck          `json:"tcpCheck"`
	Period            *string                             `json:"period"`
	Timeout           *string                             `json:"timeout"`
	ContentMatchers   []UptimeCheckConfigContentMatchers  `json:"contentMatchers"`
	PrivateCheckers   []string                            `json:"privateCheckers"`
	SelectedRegions   []string                            `json:"selectedRegions"`
	Project           *string                             `json:"project"`
}

func (r *UptimeCheckConfig) String() string {
	return dcl.SprintResource(r)
}

// The enum UptimeCheckConfigResourceGroupResourceTypeEnum.
type UptimeCheckConfigResourceGroupResourceTypeEnum string

// UptimeCheckConfigResourceGroupResourceTypeEnumRef returns a *UptimeCheckConfigResourceGroupResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UptimeCheckConfigResourceGroupResourceTypeEnumRef(s string) *UptimeCheckConfigResourceGroupResourceTypeEnum {
	if s == "" {
		return nil
	}

	v := UptimeCheckConfigResourceGroupResourceTypeEnum(s)
	return &v
}

func (v UptimeCheckConfigResourceGroupResourceTypeEnum) Validate() error {
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "INSTANCE", "AWS_ELB_LOAD_BALANCER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UptimeCheckConfigResourceGroupResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UptimeCheckConfigHttpCheckRequestMethodEnum.
type UptimeCheckConfigHttpCheckRequestMethodEnum string

// UptimeCheckConfigHttpCheckRequestMethodEnumRef returns a *UptimeCheckConfigHttpCheckRequestMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func UptimeCheckConfigHttpCheckRequestMethodEnumRef(s string) *UptimeCheckConfigHttpCheckRequestMethodEnum {
	if s == "" {
		return nil
	}

	v := UptimeCheckConfigHttpCheckRequestMethodEnum(s)
	return &v
}

func (v UptimeCheckConfigHttpCheckRequestMethodEnum) Validate() error {
	for _, s := range []string{"METHOD_UNSPECIFIED", "GET", "POST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UptimeCheckConfigHttpCheckRequestMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UptimeCheckConfigHttpCheckContentTypeEnum.
type UptimeCheckConfigHttpCheckContentTypeEnum string

// UptimeCheckConfigHttpCheckContentTypeEnumRef returns a *UptimeCheckConfigHttpCheckContentTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UptimeCheckConfigHttpCheckContentTypeEnumRef(s string) *UptimeCheckConfigHttpCheckContentTypeEnum {
	if s == "" {
		return nil
	}

	v := UptimeCheckConfigHttpCheckContentTypeEnum(s)
	return &v
}

func (v UptimeCheckConfigHttpCheckContentTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "URL_ENCODED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UptimeCheckConfigHttpCheckContentTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UptimeCheckConfigContentMatchersMatcherEnum.
type UptimeCheckConfigContentMatchersMatcherEnum string

// UptimeCheckConfigContentMatchersMatcherEnumRef returns a *UptimeCheckConfigContentMatchersMatcherEnum with the value of string s
// If the empty string is provided, nil is returned.
func UptimeCheckConfigContentMatchersMatcherEnumRef(s string) *UptimeCheckConfigContentMatchersMatcherEnum {
	if s == "" {
		return nil
	}

	v := UptimeCheckConfigContentMatchersMatcherEnum(s)
	return &v
}

func (v UptimeCheckConfigContentMatchersMatcherEnum) Validate() error {
	for _, s := range []string{"CONTENT_MATCHER_OPTION_UNSPECIFIED", "CONTAINS_STRING", "NOT_CONTAINS_STRING", "MATCHES_REGEX", "NOT_MATCHES_REGEX"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UptimeCheckConfigContentMatchersMatcherEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type UptimeCheckConfigMonitoredResource struct {
	empty        bool              `json:"-"`
	Type         *string           `json:"type"`
	FilterLabels map[string]string `json:"filterLabels"`
}

// This object is used to assert a desired state where this UptimeCheckConfigMonitoredResource is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigMonitoredResource *UptimeCheckConfigMonitoredResource = &UptimeCheckConfigMonitoredResource{empty: true}

func (r *UptimeCheckConfigMonitoredResource) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigMonitoredResource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UptimeCheckConfigResourceGroup struct {
	empty        bool                                            `json:"-"`
	GroupId      *string                                         `json:"groupId"`
	ResourceType *UptimeCheckConfigResourceGroupResourceTypeEnum `json:"resourceType"`
}

// This object is used to assert a desired state where this UptimeCheckConfigResourceGroup is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigResourceGroup *UptimeCheckConfigResourceGroup = &UptimeCheckConfigResourceGroup{empty: true}

func (r *UptimeCheckConfigResourceGroup) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigResourceGroup) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UptimeCheckConfigHttpCheck struct {
	empty         bool                                         `json:"-"`
	RequestMethod *UptimeCheckConfigHttpCheckRequestMethodEnum `json:"requestMethod"`
	UseSsl        *bool                                        `json:"useSsl"`
	Path          *string                                      `json:"path"`
	Port          *int64                                       `json:"port"`
	AuthInfo      *UptimeCheckConfigHttpCheckAuthInfo          `json:"authInfo"`
	MaskHeaders   *bool                                        `json:"maskHeaders"`
	Headers       map[string]string                            `json:"headers"`
	ContentType   *UptimeCheckConfigHttpCheckContentTypeEnum   `json:"contentType"`
	ValidateSsl   *bool                                        `json:"validateSsl"`
	Body          *string                                      `json:"body"`
}

// This object is used to assert a desired state where this UptimeCheckConfigHttpCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigHttpCheck *UptimeCheckConfigHttpCheck = &UptimeCheckConfigHttpCheck{empty: true}

func (r *UptimeCheckConfigHttpCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigHttpCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UptimeCheckConfigHttpCheckAuthInfo struct {
	empty    bool    `json:"-"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// This object is used to assert a desired state where this UptimeCheckConfigHttpCheckAuthInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigHttpCheckAuthInfo *UptimeCheckConfigHttpCheckAuthInfo = &UptimeCheckConfigHttpCheckAuthInfo{empty: true}

func (r *UptimeCheckConfigHttpCheckAuthInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigHttpCheckAuthInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UptimeCheckConfigTcpCheck struct {
	empty bool   `json:"-"`
	Port  *int64 `json:"port"`
}

// This object is used to assert a desired state where this UptimeCheckConfigTcpCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigTcpCheck *UptimeCheckConfigTcpCheck = &UptimeCheckConfigTcpCheck{empty: true}

func (r *UptimeCheckConfigTcpCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigTcpCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UptimeCheckConfigContentMatchers struct {
	empty   bool                                         `json:"-"`
	Content *string                                      `json:"content"`
	Matcher *UptimeCheckConfigContentMatchersMatcherEnum `json:"matcher"`
}

// This object is used to assert a desired state where this UptimeCheckConfigContentMatchers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUptimeCheckConfigContentMatchers *UptimeCheckConfigContentMatchers = &UptimeCheckConfigContentMatchers{empty: true}

func (r *UptimeCheckConfigContentMatchers) String() string {
	return dcl.SprintResource(r)
}

func (r *UptimeCheckConfigContentMatchers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *UptimeCheckConfig) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "UptimeCheckConfig",
		Version: "monitoring",
	}
}

const UptimeCheckConfigMaxPage = -1

type UptimeCheckConfigList struct {
	Items []*UptimeCheckConfig

	nextToken string

	pageSize int32

	project string
}

func (l *UptimeCheckConfigList) HasNext() bool {
	return l.nextToken != ""
}

func (l *UptimeCheckConfigList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listUptimeCheckConfig(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListUptimeCheckConfig(ctx context.Context, project string) (*UptimeCheckConfigList, error) {

	return c.ListUptimeCheckConfigWithMaxResults(ctx, project, UptimeCheckConfigMaxPage)

}

func (c *Client) ListUptimeCheckConfigWithMaxResults(ctx context.Context, project string, pageSize int32) (*UptimeCheckConfigList, error) {
	items, token, err := c.listUptimeCheckConfig(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &UptimeCheckConfigList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetUptimeCheckConfig(ctx context.Context, r *UptimeCheckConfig) (*UptimeCheckConfig, error) {
	b, err := c.getUptimeCheckConfigRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalUptimeCheckConfig(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	result.Name = r.Name
	if dcl.IsZeroValue(result.Period) {
		result.Period = dcl.String("60s")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeUptimeCheckConfigNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteUptimeCheckConfig(ctx context.Context, r *UptimeCheckConfig) error {
	if r == nil {
		return fmt.Errorf("UptimeCheckConfig resource is nil")
	}
	c.Config.Logger.Info("Deleting UptimeCheckConfig...")
	deleteOp := deleteUptimeCheckConfigOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllUptimeCheckConfig deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllUptimeCheckConfig(ctx context.Context, project string, filter func(*UptimeCheckConfig) bool) error {
	listObj, err := c.ListUptimeCheckConfig(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllUptimeCheckConfig(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllUptimeCheckConfig(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyUptimeCheckConfig(ctx context.Context, rawDesired *UptimeCheckConfig, opts ...dcl.ApplyOption) (*UptimeCheckConfig, error) {
	c.Config.Logger.Info("Beginning ApplyUptimeCheckConfig...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.uptimeCheckConfigDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []uptimeCheckConfigApiOperation
	if create {
		ops = append(ops, &createUptimeCheckConfigOperation{})
	} else if recreate {
		ops = append(ops, &deleteUptimeCheckConfigOperation{})
		ops = append(ops, &createUptimeCheckConfigOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeUptimeCheckConfigDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetUptimeCheckConfig(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeUptimeCheckConfigNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeUptimeCheckConfigDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffUptimeCheckConfig(c, newDesired, newState)
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
