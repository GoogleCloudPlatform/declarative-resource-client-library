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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *SslCertificate) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SelfManaged) {
		if err := r.SelfManaged.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SslCertificateSelfManaged) validate() error {
	return nil
}

func sslCertificateGetURL(userBasePath string, r *SslCertificate) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/sslCertificates/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/sslCertificates/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func sslCertificateListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/sslCertificates", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/sslCertificates", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func sslCertificateCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/sslCertificates", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/sslCertificates", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func sslCertificateDeleteURL(userBasePath string, r *SslCertificate) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/sslCertificates/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/sslCertificates/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// sslCertificateApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type sslCertificateApiOperation interface {
	do(context.Context, *SslCertificate, *Client) error
}

func (c *Client) listSslCertificateRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := sslCertificateListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SslCertificateMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listSslCertificateOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listSslCertificate(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*SslCertificate, string, error) {
	b, err := c.listSslCertificateRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSslCertificateOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*SslCertificate
	for _, v := range m.Items {
		res, err := unmarshalMapSslCertificate(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSslCertificate(ctx context.Context, f func(*SslCertificate) bool, resources []*SslCertificate) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSslCertificate(ctx, res)
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

type deleteSslCertificateOperation struct{}

func (op *deleteSslCertificateOperation) do(ctx context.Context, r *SslCertificate, c *Client) error {
	r, err := c.GetSslCertificate(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("SslCertificate not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSslCertificate checking for existence. error: %v", err)
		return err
	}

	u, err := sslCertificateDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetSslCertificate(ctx, r.URLNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createSslCertificateOperation struct {
	response map[string]interface{}
}

func (op *createSslCertificateOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSslCertificateOperation) do(ctx context.Context, r *SslCertificate, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := sslCertificateCreateURL(c.Config.BasePath, project, location)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetSslCertificate(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSslCertificateRaw(ctx context.Context, r *SslCertificate) ([]byte, error) {
	if dcl.IsZeroValue(r.Type) {
		r.Type = SslCertificateTypeEnumRef("SELF_MANAGED")
	}

	u, err := sslCertificateGetURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
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

func (c *Client) sslCertificateDiffsForRawDesired(ctx context.Context, rawDesired *SslCertificate, opts ...dcl.ApplyOption) (initial, desired *SslCertificate, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *SslCertificate
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*SslCertificate); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected SslCertificate, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSslCertificate(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a SslCertificate resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve SslCertificate resource: %v", err)
		}
		c.Config.Logger.Info("Found that SslCertificate resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSslCertificateDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for SslCertificate: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for SslCertificate: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSslCertificateInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for SslCertificate: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSslCertificateDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for SslCertificate: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSslCertificate(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSslCertificateInitialState(rawInitial, rawDesired *SslCertificate) (*SslCertificate, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSslCertificateDesiredState(rawDesired, rawInitial *SslCertificate, opts ...dcl.ApplyOption) (*SslCertificate, error) {

	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = SslCertificateTypeEnumRef("SELF_MANAGED")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SelfManaged = canonicalizeSslCertificateSelfManaged(rawDesired.SelfManaged, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &SslCertificate{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.SelfManaged = canonicalizeSslCertificateSelfManaged(rawDesired.SelfManaged, rawInitial.SelfManaged, opts...)
	if dcl.IsZeroValue(rawDesired.Type) {
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeSslCertificateNewState(c *Client, rawNew, rawDesired *SslCertificate) (*SslCertificate, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfManaged) && dcl.IsEmptyValueIndirect(rawDesired.SelfManaged) {
		rawNew.SelfManaged = rawDesired.SelfManaged
	} else {
		rawNew.SelfManaged = canonicalizeNewSslCertificateSelfManaged(c, rawDesired.SelfManaged, rawNew.SelfManaged)
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
		if dcl.StringCanonicalize(rawDesired.ExpireTime, rawNew.ExpireTime) {
			rawNew.ExpireTime = rawDesired.ExpireTime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeSslCertificateSelfManaged(des, initial *SslCertificateSelfManaged, opts ...dcl.ApplyOption) *SslCertificateSelfManaged {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SslCertificateSelfManaged{}

	if dcl.StringCanonicalize(des.Certificate, initial.Certificate) || dcl.IsZeroValue(des.Certificate) {
		cDes.Certificate = initial.Certificate
	} else {
		cDes.Certificate = des.Certificate
	}
	if dcl.StringCanonicalize(des.PrivateKey, initial.PrivateKey) || dcl.IsZeroValue(des.PrivateKey) {
		cDes.PrivateKey = initial.PrivateKey
	} else {
		cDes.PrivateKey = des.PrivateKey
	}

	return cDes
}

func canonicalizeNewSslCertificateSelfManaged(c *Client, des, nw *SslCertificateSelfManaged) *SslCertificateSelfManaged {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Certificate, nw.Certificate) {
		nw.Certificate = des.Certificate
	}
	nw.PrivateKey = des.PrivateKey

	return nw
}

func canonicalizeNewSslCertificateSelfManagedSet(c *Client, des, nw []SslCertificateSelfManaged) []SslCertificateSelfManaged {
	if des == nil {
		return nw
	}
	var reorderedNew []SslCertificateSelfManaged
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSslCertificateSelfManagedNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSslCertificateSelfManagedSlice(c *Client, des, nw []SslCertificateSelfManaged) []SslCertificateSelfManaged {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SslCertificateSelfManaged
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSslCertificateSelfManaged(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffSslCertificate(c *Client, desired, actual *SslCertificate, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfManaged, actual.SelfManaged, dcl.Info{ObjectFunction: compareSslCertificateSelfManagedNewStyle, EmptyObject: EmptySslCertificateSelfManaged, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfManaged")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectAlternativeNames, actual.SubjectAlternativeNames, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectAlternativeNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareSslCertificateSelfManagedNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SslCertificateSelfManaged)
	if !ok {
		desiredNotPointer, ok := d.(SslCertificateSelfManaged)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslCertificateSelfManaged or *SslCertificateSelfManaged", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SslCertificateSelfManaged)
	if !ok {
		actualNotPointer, ok := a.(SslCertificateSelfManaged)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslCertificateSelfManaged", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Certificate, actual.Certificate, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Certificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateKey, actual.PrivateKey, dcl.Info{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *SslCertificate) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslCertificate) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *SslCertificate) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslCertificate) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the SslCertificate resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *SslCertificate) marshal(c *Client) ([]byte, error) {
	m, err := expandSslCertificate(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling SslCertificate: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSslCertificate decodes JSON responses into the SslCertificate resource schema.
func unmarshalSslCertificate(b []byte, c *Client) (*SslCertificate, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSslCertificate(m, c)
}

func unmarshalMapSslCertificate(m map[string]interface{}, c *Client) (*SslCertificate, error) {

	return flattenSslCertificate(c, m), nil
}

// expandSslCertificate expands SslCertificate into a JSON request object.
func expandSslCertificate(c *Client, f *SslCertificate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := expandSslCertificateSelfManaged(c, f.SelfManaged); err != nil {
		return nil, fmt.Errorf("error expanding SelfManaged into selfManaged: %w", err)
	} else if v != nil {
		m["selfManaged"] = v
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
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenSslCertificate flattens SslCertificate from a JSON request object into the
// SslCertificate type.
func flattenSslCertificate(c *Client, i interface{}) *SslCertificate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &SslCertificate{}
	res.Name = dcl.FlattenString(m["name"])
	res.Id = dcl.FlattenInteger(m["id"])
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Description = dcl.FlattenString(m["description"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.SelfManaged = flattenSslCertificateSelfManaged(c, m["selfManaged"])
	res.Type = flattenSslCertificateTypeEnum(m["type"])
	if _, ok := m["type"]; !ok {
		c.Config.Logger.Info("Using default value for type")
		res.Type = SslCertificateTypeEnumRef("SELF_MANAGED")
	}
	res.SubjectAlternativeNames = dcl.FlattenStringSlice(m["subjectAlternativeNames"])
	res.ExpireTime = dcl.FlattenString(m["expireTime"])
	res.Region = dcl.FlattenString(m["region"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandSslCertificateSelfManagedMap expands the contents of SslCertificateSelfManaged into a JSON
// request object.
func expandSslCertificateSelfManagedMap(c *Client, f map[string]SslCertificateSelfManaged) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSslCertificateSelfManaged(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSslCertificateSelfManagedSlice expands the contents of SslCertificateSelfManaged into a JSON
// request object.
func expandSslCertificateSelfManagedSlice(c *Client, f []SslCertificateSelfManaged) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSslCertificateSelfManaged(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSslCertificateSelfManagedMap flattens the contents of SslCertificateSelfManaged from a JSON
// response object.
func flattenSslCertificateSelfManagedMap(c *Client, i interface{}) map[string]SslCertificateSelfManaged {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslCertificateSelfManaged{}
	}

	if len(a) == 0 {
		return map[string]SslCertificateSelfManaged{}
	}

	items := make(map[string]SslCertificateSelfManaged)
	for k, item := range a {
		items[k] = *flattenSslCertificateSelfManaged(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSslCertificateSelfManagedSlice flattens the contents of SslCertificateSelfManaged from a JSON
// response object.
func flattenSslCertificateSelfManagedSlice(c *Client, i interface{}) []SslCertificateSelfManaged {
	a, ok := i.([]interface{})
	if !ok {
		return []SslCertificateSelfManaged{}
	}

	if len(a) == 0 {
		return []SslCertificateSelfManaged{}
	}

	items := make([]SslCertificateSelfManaged, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslCertificateSelfManaged(c, item.(map[string]interface{})))
	}

	return items
}

// expandSslCertificateSelfManaged expands an instance of SslCertificateSelfManaged into a JSON
// request object.
func expandSslCertificateSelfManaged(c *Client, f *SslCertificateSelfManaged) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Certificate; !dcl.IsEmptyValueIndirect(v) {
		m["certificate"] = v
	}
	if v := f.PrivateKey; !dcl.IsEmptyValueIndirect(v) {
		m["privateKey"] = v
	}

	return m, nil
}

// flattenSslCertificateSelfManaged flattens an instance of SslCertificateSelfManaged from a JSON
// response object.
func flattenSslCertificateSelfManaged(c *Client, i interface{}) *SslCertificateSelfManaged {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SslCertificateSelfManaged{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySslCertificateSelfManaged
	}
	r.Certificate = dcl.FlattenString(m["certificate"])
	r.PrivateKey = dcl.FlattenSecretValue(m["privateKey"])

	return r
}

// flattenSslCertificateTypeEnumMap flattens the contents of SslCertificateTypeEnum from a JSON
// response object.
func flattenSslCertificateTypeEnumMap(c *Client, i interface{}) map[string]SslCertificateTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslCertificateTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]SslCertificateTypeEnum{}
	}

	items := make(map[string]SslCertificateTypeEnum)
	for k, item := range a {
		items[k] = *flattenSslCertificateTypeEnum(item.(interface{}))
	}

	return items
}

// flattenSslCertificateTypeEnumSlice flattens the contents of SslCertificateTypeEnum from a JSON
// response object.
func flattenSslCertificateTypeEnumSlice(c *Client, i interface{}) []SslCertificateTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SslCertificateTypeEnum{}
	}

	if len(a) == 0 {
		return []SslCertificateTypeEnum{}
	}

	items := make([]SslCertificateTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslCertificateTypeEnum(item.(interface{})))
	}

	return items
}

// flattenSslCertificateTypeEnum asserts that an interface is a string, and returns a
// pointer to a *SslCertificateTypeEnum with the same value as that string.
func flattenSslCertificateTypeEnum(i interface{}) *SslCertificateTypeEnum {
	s, ok := i.(string)
	if !ok {
		return SslCertificateTypeEnumRef("")
	}

	return SslCertificateTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *SslCertificate) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSslCertificate(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type sslCertificateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         sslCertificateApiOperation
}

func convertFieldDiffsToSslCertificateDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]sslCertificateDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []sslCertificateDiff
	// For each operation name, create a sslCertificateDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := sslCertificateDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToSslCertificateApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToSslCertificateApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (sslCertificateApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
