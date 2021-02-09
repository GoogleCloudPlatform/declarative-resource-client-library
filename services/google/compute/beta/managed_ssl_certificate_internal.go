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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *ManagedSslCertificate) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Managed) {
		if err := r.Managed.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ManagedSslCertificateManaged) validate() error {
	return nil
}

func managedSslCertificateGetURL(userBasePath string, r *ManagedSslCertificate) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/sslCertificates/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func managedSslCertificateListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/sslCertificates", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func managedSslCertificateCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/sslCertificates", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func managedSslCertificateDeleteURL(userBasePath string, r *ManagedSslCertificate) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/sslCertificates/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// managedSslCertificateApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type managedSslCertificateApiOperation interface {
	do(context.Context, *ManagedSslCertificate, *Client) error
}

func (c *Client) listManagedSslCertificateRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := managedSslCertificateListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ManagedSslCertificateMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listManagedSslCertificateOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listManagedSslCertificate(ctx context.Context, project, pageToken string, pageSize int32) ([]*ManagedSslCertificate, string, error) {
	b, err := c.listManagedSslCertificateRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listManagedSslCertificateOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ManagedSslCertificate
	for _, v := range m.Items {
		res := flattenManagedSslCertificate(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllManagedSslCertificate(ctx context.Context, f func(*ManagedSslCertificate) bool, resources []*ManagedSslCertificate) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteManagedSslCertificate(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteManagedSslCertificateOperation struct{}

func (op *deleteManagedSslCertificateOperation) do(ctx context.Context, r *ManagedSslCertificate, c *Client) error {

	_, err := c.GetManagedSslCertificate(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ManagedSslCertificate not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetManagedSslCertificate checking for existence. error: %v", err)
		return err
	}

	u, err := managedSslCertificateDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		return err
	}
	_, err = c.GetManagedSslCertificate(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createManagedSslCertificateOperation struct{}

func (op *createManagedSslCertificateOperation) do(ctx context.Context, r *ManagedSslCertificate, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := managedSslCertificateCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetManagedSslCertificate(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getManagedSslCertificateRaw(ctx context.Context, r *ManagedSslCertificate) ([]byte, error) {
	if dcl.IsZeroValue(r.Type) {
		r.Type = ManagedSslCertificateTypeEnumRef("MANAGED")
	}

	u, err := managedSslCertificateGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) managedSslCertificateDiffsForRawDesired(ctx context.Context, rawDesired *ManagedSslCertificate, opts ...dcl.ApplyOption) (initial, desired *ManagedSslCertificate, diffs []managedSslCertificateDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ManagedSslCertificate
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ManagedSslCertificate); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ManagedSslCertificate, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetManagedSslCertificate(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ManagedSslCertificate resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ManagedSslCertificate resource: %v", err)
		}
		c.Config.Logger.Info("Found that ManagedSslCertificate resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeManagedSslCertificateDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ManagedSslCertificate: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ManagedSslCertificate: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeManagedSslCertificateInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ManagedSslCertificate: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeManagedSslCertificateDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ManagedSslCertificate: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffManagedSslCertificate(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeManagedSslCertificateInitialState(rawInitial, rawDesired *ManagedSslCertificate) (*ManagedSslCertificate, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeManagedSslCertificateDesiredState(rawDesired, rawInitial *ManagedSslCertificate, opts ...dcl.ApplyOption) (*ManagedSslCertificate, error) {

	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = ManagedSslCertificateTypeEnumRef("MANAGED")
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ManagedSslCertificate); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected ManagedSslCertificate, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Managed = canonicalizeManagedSslCertificateManaged(rawDesired.Managed, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	rawDesired.Managed = canonicalizeManagedSslCertificateManaged(rawDesired.Managed, rawInitial.Managed, opts...)
	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.IsZeroValue(rawDesired.SubjectAlternativeNames) {
		rawDesired.SubjectAlternativeNames = rawInitial.SubjectAlternativeNames
	}
	if dcl.IsZeroValue(rawDesired.ExpireTime) {
		rawDesired.ExpireTime = rawInitial.ExpireTime
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeManagedSslCertificateNewState(c *Client, rawNew, rawDesired *ManagedSslCertificate) (*ManagedSslCertificate, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Managed) && dcl.IsEmptyValueIndirect(rawDesired.Managed) {
		rawNew.Managed = rawDesired.Managed
	} else {
		rawNew.Managed = canonicalizeNewManagedSslCertificateManaged(c, rawDesired.Managed, rawNew.Managed)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubjectAlternativeNames) && dcl.IsEmptyValueIndirect(rawDesired.SubjectAlternativeNames) {
		rawNew.SubjectAlternativeNames = rawDesired.SubjectAlternativeNames
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpireTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeManagedSslCertificateManaged(des, initial *ManagedSslCertificateManaged, opts ...dcl.ApplyOption) *ManagedSslCertificateManaged {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*ManagedSslCertificate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.WithoutTrailingDotArray(des.Domains, initial.Domains) || dcl.IsZeroValue(des.Domains) {
		des.Domains = initial.Domains
	}
	if dcl.IsZeroValue(des.Status) {
		des.Status = initial.Status
	}
	if dcl.IsZeroValue(des.DomainStatus) {
		des.DomainStatus = initial.DomainStatus
	}

	return des
}

func canonicalizeNewManagedSslCertificateManaged(c *Client, des, nw *ManagedSslCertificateManaged) *ManagedSslCertificateManaged {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.WithoutTrailingDotArray(des.Domains, nw.Domains) || dcl.IsZeroValue(des.Domains) {
		nw.Domains = des.Domains
	}

	return nw
}

func canonicalizeNewManagedSslCertificateManagedSet(c *Client, des, nw []ManagedSslCertificateManaged) []ManagedSslCertificateManaged {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedSslCertificateManaged
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedSslCertificateManaged(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

type managedSslCertificateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         managedSslCertificateApiOperation
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffManagedSslCertificate(c *Client, desired, actual *ManagedSslCertificate, opts ...dcl.ApplyOption) ([]managedSslCertificateDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []managedSslCertificateDiff
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, managedSslCertificateDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, managedSslCertificateDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if compareManagedSslCertificateManaged(c, desired.Managed, actual.Managed) {
		c.Config.Logger.Infof("Detected diff in Managed.\nDESIRED: %v\nACTUAL: %v", desired.Managed, actual.Managed)
		diffs = append(diffs, managedSslCertificateDiff{
			RequiresRecreate: true,
			FieldName:        "Managed",
		})
	}
	if !dcl.IsZeroValue(desired.Type) && (dcl.IsZeroValue(actual.Type) || !reflect.DeepEqual(*desired.Type, *actual.Type)) {
		c.Config.Logger.Infof("Detected diff in Type.\nDESIRED: %v\nACTUAL: %v", desired.Type, actual.Type)
		diffs = append(diffs, managedSslCertificateDiff{
			RequiresRecreate: true,
			FieldName:        "Type",
		})
	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []managedSslCertificateDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareManagedSslCertificateManagedSlice(c *Client, desired, actual []ManagedSslCertificateManaged) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedSslCertificateManaged, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedSslCertificateManaged(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedSslCertificateManaged, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedSslCertificateManaged(c *Client, desired, actual *ManagedSslCertificateManaged) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Domains == nil && desired.Domains != nil && !dcl.IsEmptyValueIndirect(desired.Domains) {
		c.Config.Logger.Infof("desired Domains %s - but actually nil", dcl.SprintResource(desired.Domains))
		return true
	}
	if !dcl.WithoutTrailingDotArray(desired.Domains, actual.Domains) && !dcl.IsZeroValue(desired.Domains) {
		c.Config.Logger.Infof("Diff in Domains. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Domains), dcl.SprintResource(actual.Domains))
		return true
	}
	return false
}
func compareManagedSslCertificateManagedStatusEnumSlice(c *Client, desired, actual []ManagedSslCertificateManagedStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedSslCertificateManagedStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedSslCertificateManagedStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedSslCertificateManagedStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedSslCertificateManagedStatusEnum(c *Client, desired, actual *ManagedSslCertificateManagedStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedSslCertificateTypeEnumSlice(c *Client, desired, actual []ManagedSslCertificateTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedSslCertificateTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedSslCertificateTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedSslCertificateTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedSslCertificateTypeEnum(c *Client, desired, actual *ManagedSslCertificateTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ManagedSslCertificate) urlNormalized() *ManagedSslCertificate {
	normalized := deepcopy.Copy(*r).(ManagedSslCertificate)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *ManagedSslCertificate) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ManagedSslCertificate) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *ManagedSslCertificate) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ManagedSslCertificate) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ManagedSslCertificate resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ManagedSslCertificate) marshal(c *Client) ([]byte, error) {
	m, err := expandManagedSslCertificate(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ManagedSslCertificate: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalManagedSslCertificate decodes JSON responses into the ManagedSslCertificate resource schema.
func unmarshalManagedSslCertificate(b []byte, c *Client) (*ManagedSslCertificate, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenManagedSslCertificate(c, m), nil
}

// expandManagedSslCertificate expands ManagedSslCertificate into a JSON request object.
func expandManagedSslCertificate(c *Client, f *ManagedSslCertificate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := expandManagedSslCertificateManaged(c, f.Managed); err != nil {
		return nil, fmt.Errorf("error expanding Managed into managed: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["managed"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.SubjectAlternativeNames; !dcl.IsEmptyValueIndirect(v) {
		m["subjectAlternativeNames"] = v
	}
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenManagedSslCertificate flattens ManagedSslCertificate from a JSON request object into the
// ManagedSslCertificate type.
func flattenManagedSslCertificate(c *Client, i interface{}) *ManagedSslCertificate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &ManagedSslCertificate{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Managed = flattenManagedSslCertificateManaged(c, m["managed"])
	r.Type = flattenManagedSslCertificateTypeEnum(m["type"])
	if _, ok := m["type"]; !ok {
		c.Config.Logger.Info("Using default value for type")
		r.Type = ManagedSslCertificateTypeEnumRef("MANAGED")
	}
	r.SubjectAlternativeNames = dcl.FlattenStringSlice(m["subjectAlternativeNames"])
	r.ExpireTime = dcl.FlattenString(m["expireTime"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandManagedSslCertificateManagedMap expands the contents of ManagedSslCertificateManaged into a JSON
// request object.
func expandManagedSslCertificateManagedMap(c *Client, f map[string]ManagedSslCertificateManaged) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedSslCertificateManaged(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedSslCertificateManagedSlice expands the contents of ManagedSslCertificateManaged into a JSON
// request object.
func expandManagedSslCertificateManagedSlice(c *Client, f []ManagedSslCertificateManaged) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedSslCertificateManaged(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedSslCertificateManagedMap flattens the contents of ManagedSslCertificateManaged from a JSON
// response object.
func flattenManagedSslCertificateManagedMap(c *Client, i interface{}) map[string]ManagedSslCertificateManaged {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedSslCertificateManaged{}
	}

	if len(a) == 0 {
		return map[string]ManagedSslCertificateManaged{}
	}

	items := make(map[string]ManagedSslCertificateManaged)
	for k, item := range a {
		items[k] = *flattenManagedSslCertificateManaged(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedSslCertificateManagedSlice flattens the contents of ManagedSslCertificateManaged from a JSON
// response object.
func flattenManagedSslCertificateManagedSlice(c *Client, i interface{}) []ManagedSslCertificateManaged {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedSslCertificateManaged{}
	}

	if len(a) == 0 {
		return []ManagedSslCertificateManaged{}
	}

	items := make([]ManagedSslCertificateManaged, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedSslCertificateManaged(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedSslCertificateManaged expands an instance of ManagedSslCertificateManaged into a JSON
// request object.
func expandManagedSslCertificateManaged(c *Client, f *ManagedSslCertificateManaged) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Domains; !dcl.IsEmptyValueIndirect(v) {
		m["domains"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.DomainStatus; !dcl.IsEmptyValueIndirect(v) {
		m["domainStatus"] = v
	}

	return m, nil
}

// flattenManagedSslCertificateManaged flattens an instance of ManagedSslCertificateManaged from a JSON
// response object.
func flattenManagedSslCertificateManaged(c *Client, i interface{}) *ManagedSslCertificateManaged {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedSslCertificateManaged{}
	r.Domains = dcl.FlattenStringSlice(m["domains"])
	r.Status = flattenManagedSslCertificateManagedStatusEnum(m["status"])
	r.DomainStatus = dcl.FlattenKeyValuePairs(m["domainStatus"])

	return r
}

// flattenManagedSslCertificateManagedStatusEnumSlice flattens the contents of ManagedSslCertificateManagedStatusEnum from a JSON
// response object.
func flattenManagedSslCertificateManagedStatusEnumSlice(c *Client, i interface{}) []ManagedSslCertificateManagedStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedSslCertificateManagedStatusEnum{}
	}

	if len(a) == 0 {
		return []ManagedSslCertificateManagedStatusEnum{}
	}

	items := make([]ManagedSslCertificateManagedStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedSslCertificateManagedStatusEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenManagedSslCertificateManagedStatusEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedSslCertificateManagedStatusEnum with the same value as that string.
func flattenManagedSslCertificateManagedStatusEnum(i interface{}) *ManagedSslCertificateManagedStatusEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedSslCertificateManagedStatusEnumRef("")
	}

	return ManagedSslCertificateManagedStatusEnumRef(s)
}

// flattenManagedSslCertificateTypeEnumSlice flattens the contents of ManagedSslCertificateTypeEnum from a JSON
// response object.
func flattenManagedSslCertificateTypeEnumSlice(c *Client, i interface{}) []ManagedSslCertificateTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedSslCertificateTypeEnum{}
	}

	if len(a) == 0 {
		return []ManagedSslCertificateTypeEnum{}
	}

	items := make([]ManagedSslCertificateTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedSslCertificateTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenManagedSslCertificateTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedSslCertificateTypeEnum with the same value as that string.
func flattenManagedSslCertificateTypeEnum(i interface{}) *ManagedSslCertificateTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedSslCertificateTypeEnumRef("")
	}

	return ManagedSslCertificateTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ManagedSslCertificate) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalManagedSslCertificate(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}
