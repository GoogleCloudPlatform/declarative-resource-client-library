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
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ClientTlsPolicy struct {
	Name               *string                             `json:"name"`
	Description        *string                             `json:"description"`
	CreateTime         *string                             `json:"createTime"`
	UpdateTime         *string                             `json:"updateTime"`
	Labels             map[string]string                   `json:"labels"`
	Sni                *string                             `json:"sni"`
	ClientCertificate  *ClientTlsPolicyClientCertificate   `json:"clientCertificate"`
	ServerValidationCa []ClientTlsPolicyServerValidationCa `json:"serverValidationCa"`
	Project            *string                             `json:"project"`
	Location           *string                             `json:"location"`
}

func (r *ClientTlsPolicy) String() string {
	return dcl.SprintResource(r)
}

type ClientTlsPolicyClientCertificate struct {
	empty                       bool                                                         `json:"-"`
	GrpcEndpoint                *ClientTlsPolicyClientCertificateGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ClientTlsPolicyClientCertificateCertificateProviderInstance `json:"certificateProviderInstance"`
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificate *ClientTlsPolicyClientCertificate = &ClientTlsPolicyClientCertificate{empty: true}

func (r *ClientTlsPolicyClientCertificate) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyClientCertificateGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificateGrpcEndpoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificateGrpcEndpoint *ClientTlsPolicyClientCertificateGrpcEndpoint = &ClientTlsPolicyClientCertificateGrpcEndpoint{empty: true}

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyClientCertificateCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificateCertificateProviderInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificateCertificateProviderInstance *ClientTlsPolicyClientCertificateCertificateProviderInstance = &ClientTlsPolicyClientCertificateCertificateProviderInstance{empty: true}

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCa struct {
	empty                       bool                                                          `json:"-"`
	GrpcEndpoint                *ClientTlsPolicyServerValidationCaGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ClientTlsPolicyServerValidationCaCertificateProviderInstance `json:"certificateProviderInstance"`
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCa is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCa *ClientTlsPolicyServerValidationCa = &ClientTlsPolicyServerValidationCa{empty: true}

func (r *ClientTlsPolicyServerValidationCa) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCa) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCaGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCaGrpcEndpoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCaGrpcEndpoint *ClientTlsPolicyServerValidationCaGrpcEndpoint = &ClientTlsPolicyServerValidationCaGrpcEndpoint{empty: true}

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCaCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCaCertificateProviderInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance *ClientTlsPolicyServerValidationCaCertificateProviderInstance = &ClientTlsPolicyServerValidationCaCertificateProviderInstance{empty: true}

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ClientTlsPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "networksecurity",
		Type:    "ClientTlsPolicy",
		Version: "beta",
	}
}

const ClientTlsPolicyMaxPage = -1

type ClientTlsPolicyList struct {
	Items []*ClientTlsPolicy

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ClientTlsPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ClientTlsPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listClientTlsPolicy(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListClientTlsPolicy(ctx context.Context, project, location string) (*ClientTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListClientTlsPolicyWithMaxResults(ctx, project, location, ClientTlsPolicyMaxPage)

}

func (c *Client) ListClientTlsPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ClientTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listClientTlsPolicy(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ClientTlsPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetClientTlsPolicy(ctx context.Context, r *ClientTlsPolicy) (*ClientTlsPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getClientTlsPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalClientTlsPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeClientTlsPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteClientTlsPolicy(ctx context.Context, r *ClientTlsPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("ClientTlsPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting ClientTlsPolicy...")
	deleteOp := deleteClientTlsPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllClientTlsPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllClientTlsPolicy(ctx context.Context, project, location string, filter func(*ClientTlsPolicy) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListClientTlsPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllClientTlsPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllClientTlsPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyClientTlsPolicy(ctx context.Context, rawDesired *ClientTlsPolicy, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyClientTlsPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.clientTlsPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []clientTlsPolicyApiOperation
	if create {
		ops = append(ops, &createClientTlsPolicyOperation{})
	} else if recreate {

		ops = append(ops, &deleteClientTlsPolicyOperation{})

		ops = append(ops, &createClientTlsPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeClientTlsPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetClientTlsPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createClientTlsPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapClientTlsPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeClientTlsPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeClientTlsPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeClientTlsPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffClientTlsPolicy(c, newDesired, newState)
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
func (r *ClientTlsPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
