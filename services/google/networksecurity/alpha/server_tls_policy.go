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
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ServerTlsPolicy struct {
	Name              *string                           `json:"name"`
	Description       *string                           `json:"description"`
	CreateTime        *string                           `json:"createTime"`
	UpdateTime        *string                           `json:"updateTime"`
	Labels            map[string]string                 `json:"labels"`
	AllowOpen         *bool                             `json:"allowOpen"`
	ServerCertificate *ServerTlsPolicyServerCertificate `json:"serverCertificate"`
	MtlsPolicy        *ServerTlsPolicyMtlsPolicy        `json:"mtlsPolicy"`
	Project           *string                           `json:"project"`
	Location          *string                           `json:"location"`
}

func (r *ServerTlsPolicy) String() string {
	return dcl.SprintResource(r)
}

type ServerTlsPolicyServerCertificate struct {
	empty                       bool                                                         `json:"-"`
	LocalFilepath               *ServerTlsPolicyServerCertificateLocalFilepath               `json:"localFilepath"`
	GrpcEndpoint                *ServerTlsPolicyServerCertificateGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ServerTlsPolicyServerCertificateCertificateProviderInstance `json:"certificateProviderInstance"`
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificate *ServerTlsPolicyServerCertificate = &ServerTlsPolicyServerCertificate{empty: true}

func (r *ServerTlsPolicyServerCertificate) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateLocalFilepath struct {
	empty           bool    `json:"-"`
	CertificatePath *string `json:"certificatePath"`
	PrivateKeyPath  *string `json:"privateKeyPath"`
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateLocalFilepath is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateLocalFilepath *ServerTlsPolicyServerCertificateLocalFilepath = &ServerTlsPolicyServerCertificateLocalFilepath{empty: true}

func (r *ServerTlsPolicyServerCertificateLocalFilepath) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateLocalFilepath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateGrpcEndpoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateGrpcEndpoint *ServerTlsPolicyServerCertificateGrpcEndpoint = &ServerTlsPolicyServerCertificateGrpcEndpoint{empty: true}

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateCertificateProviderInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateCertificateProviderInstance *ServerTlsPolicyServerCertificateCertificateProviderInstance = &ServerTlsPolicyServerCertificateCertificateProviderInstance{empty: true}

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicy struct {
	empty              bool                                          `json:"-"`
	ClientValidationCa []ServerTlsPolicyMtlsPolicyClientValidationCa `json:"clientValidationCa"`
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicy *ServerTlsPolicyMtlsPolicy = &ServerTlsPolicyMtlsPolicy{empty: true}

func (r *ServerTlsPolicyMtlsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCa struct {
	empty                       bool                                                                    `json:"-"`
	CaCertPath                  *string                                                                 `json:"caCertPath"`
	GrpcEndpoint                *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance `json:"certificateProviderInstance"`
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCa is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCa *ServerTlsPolicyMtlsPolicyClientValidationCa = &ServerTlsPolicyMtlsPolicyClientValidationCa{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint = &ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance = &ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ServerTlsPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "networksecurity",
		Type:    "ServerTlsPolicy",
		Version: "alpha",
	}
}

const ServerTlsPolicyMaxPage = -1

type ServerTlsPolicyList struct {
	Items []*ServerTlsPolicy

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ServerTlsPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServerTlsPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listServerTlsPolicy(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListServerTlsPolicy(ctx context.Context, project, location string) (*ServerTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListServerTlsPolicyWithMaxResults(ctx, project, location, ServerTlsPolicyMaxPage)

}

func (c *Client) ListServerTlsPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServerTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listServerTlsPolicy(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServerTlsPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetServerTlsPolicy(ctx context.Context, r *ServerTlsPolicy) (*ServerTlsPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getServerTlsPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalServerTlsPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServerTlsPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteServerTlsPolicy(ctx context.Context, r *ServerTlsPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("ServerTlsPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting ServerTlsPolicy...")
	deleteOp := deleteServerTlsPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllServerTlsPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllServerTlsPolicy(ctx context.Context, project, location string, filter func(*ServerTlsPolicy) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListServerTlsPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllServerTlsPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllServerTlsPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyServerTlsPolicy(ctx context.Context, rawDesired *ServerTlsPolicy, opts ...dcl.ApplyOption) (*ServerTlsPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyServerTlsPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.serverTlsPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []serverTlsPolicyApiOperation
	if create {
		ops = append(ops, &createServerTlsPolicyOperation{})
	} else if recreate {

		ops = append(ops, &deleteServerTlsPolicyOperation{})

		ops = append(ops, &createServerTlsPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeServerTlsPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetServerTlsPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServerTlsPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapServerTlsPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServerTlsPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServerTlsPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServerTlsPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffServerTlsPolicy(c, newDesired, newState)
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
func (r *ServerTlsPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
