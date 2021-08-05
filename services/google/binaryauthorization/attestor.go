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
package binaryauthorization

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Attestor struct {
	Name                 *string                       `json:"name"`
	Description          *string                       `json:"description"`
	UserOwnedDrydockNote *AttestorUserOwnedDrydockNote `json:"userOwnedDrydockNote"`
	UpdateTime           *string                       `json:"updateTime"`
	Project              *string                       `json:"project"`
}

func (r *Attestor) String() string {
	return dcl.SprintResource(r)
}

// The enum AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.
type AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum string

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef returns a *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef(s string) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if s == "" {
		return nil
	}

	v := AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(s)
	return &v
}

func (v AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) Validate() error {
	for _, s := range []string{"SIGNATURE_ALGORITHM_UNSPECIFIED", "RSA_PSS_2048_SHA256", "RSA_PSS_3072_SHA256", "RSA_PSS_4096_SHA256", "RSA_PSS_4096_SHA512", "RSA_SIGN_PKCS1_2048_SHA256", "RSA_SIGN_PKCS1_3072_SHA256", "RSA_SIGN_PKCS1_4096_SHA256", "RSA_SIGN_PKCS1_4096_SHA512", "ECDSA_P256_SHA256", "EC_SIGN_P256_SHA256", "ECDSA_P384_SHA384", "EC_SIGN_P384_SHA384", "ECDSA_P521_SHA512", "EC_SIGN_P521_SHA512"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AttestorUserOwnedDrydockNote struct {
	empty                         bool                                     `json:"-"`
	NoteReference                 *string                                  `json:"noteReference"`
	PublicKeys                    []AttestorUserOwnedDrydockNotePublicKeys `json:"publicKeys"`
	DelegationServiceAccountEmail *string                                  `json:"delegationServiceAccountEmail"`
}

type jsonAttestorUserOwnedDrydockNote AttestorUserOwnedDrydockNote

func (r *AttestorUserOwnedDrydockNote) UnmarshalJSON(data []byte) error {
	var res jsonAttestorUserOwnedDrydockNote
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAttestorUserOwnedDrydockNote
	} else {

		r.NoteReference = res.NoteReference

		r.PublicKeys = res.PublicKeys

		r.DelegationServiceAccountEmail = res.DelegationServiceAccountEmail

	}
	return nil
}

// This object is used to assert a desired state where this AttestorUserOwnedDrydockNote is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAttestorUserOwnedDrydockNote *AttestorUserOwnedDrydockNote = &AttestorUserOwnedDrydockNote{empty: true}

func (r *AttestorUserOwnedDrydockNote) Empty() bool {
	return r.empty
}

func (r *AttestorUserOwnedDrydockNote) String() string {
	return dcl.SprintResource(r)
}

func (r *AttestorUserOwnedDrydockNote) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AttestorUserOwnedDrydockNotePublicKeys struct {
	empty                    bool                                                 `json:"-"`
	Comment                  *string                                              `json:"comment"`
	Id                       *string                                              `json:"id"`
	AsciiArmoredPgpPublicKey *string                                              `json:"asciiArmoredPgpPublicKey"`
	PkixPublicKey            *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey `json:"pkixPublicKey"`
}

type jsonAttestorUserOwnedDrydockNotePublicKeys AttestorUserOwnedDrydockNotePublicKeys

func (r *AttestorUserOwnedDrydockNotePublicKeys) UnmarshalJSON(data []byte) error {
	var res jsonAttestorUserOwnedDrydockNotePublicKeys
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAttestorUserOwnedDrydockNotePublicKeys
	} else {

		r.Comment = res.Comment

		r.Id = res.Id

		r.AsciiArmoredPgpPublicKey = res.AsciiArmoredPgpPublicKey

		r.PkixPublicKey = res.PkixPublicKey

	}
	return nil
}

// This object is used to assert a desired state where this AttestorUserOwnedDrydockNotePublicKeys is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAttestorUserOwnedDrydockNotePublicKeys *AttestorUserOwnedDrydockNotePublicKeys = &AttestorUserOwnedDrydockNotePublicKeys{empty: true}

func (r *AttestorUserOwnedDrydockNotePublicKeys) Empty() bool {
	return r.empty
}

func (r *AttestorUserOwnedDrydockNotePublicKeys) String() string {
	return dcl.SprintResource(r)
}

func (r *AttestorUserOwnedDrydockNotePublicKeys) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey struct {
	empty              bool                                                                       `json:"-"`
	PublicKeyPem       *string                                                                    `json:"publicKeyPem"`
	SignatureAlgorithm *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum `json:"signatureAlgorithm"`
}

type jsonAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey

func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) UnmarshalJSON(data []byte) error {
	var res jsonAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	} else {

		r.PublicKeyPem = res.PublicKeyPem

		r.SignatureAlgorithm = res.SignatureAlgorithm

	}
	return nil
}

// This object is used to assert a desired state where this AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey = &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{empty: true}

func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) Empty() bool {
	return r.empty
}

func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) String() string {
	return dcl.SprintResource(r)
}

func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Attestor) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "binary_authorization",
		Type:    "Attestor",
		Version: "binaryauthorization",
	}
}

const AttestorMaxPage = -1

type AttestorList struct {
	Items []*Attestor

	nextToken string

	pageSize int32

	resource *Attestor
}

func (l *AttestorList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AttestorList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAttestor(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAttestor(ctx context.Context, r *Attestor) (*AttestorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAttestorWithMaxResults(ctx, r, AttestorMaxPage)

}

func (c *Client) ListAttestorWithMaxResults(ctx context.Context, r *Attestor, pageSize int32) (*AttestorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAttestor(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AttestorList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAttestor(ctx context.Context, r *Attestor) (*Attestor, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAttestorRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAttestor(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAttestorNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAttestor(ctx context.Context, r *Attestor) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Attestor resource is nil")
	}
	c.Config.Logger.Info("Deleting Attestor...")
	deleteOp := deleteAttestorOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAttestor deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAttestor(ctx context.Context, r *Attestor, filter func(*Attestor) bool) error {
	listObj, err := c.ListAttestor(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAttestor(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAttestor(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAttestor(ctx context.Context, rawDesired *Attestor, opts ...dcl.ApplyOption) (*Attestor, error) {
	var resultNewState *Attestor
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAttestorHelper(c, ctx, rawDesired, opts...)
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

func applyAttestorHelper(c *Client, ctx context.Context, rawDesired *Attestor, opts ...dcl.ApplyOption) (*Attestor, error) {
	c.Config.Logger.Info("Beginning ApplyAttestor...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.attestorDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAttestorDiffs(c.Config, fieldDiffs, opts)
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
	var ops []attestorApiOperation
	if create {
		ops = append(ops, &createAttestorOperation{})
	} else if recreate {
		ops = append(ops, &deleteAttestorOperation{})
		ops = append(ops, &createAttestorOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAttestorDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAttestor(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAttestorOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAttestor(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAttestorNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAttestorNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAttestorDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAttestor(c, newDesired, newState)
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
