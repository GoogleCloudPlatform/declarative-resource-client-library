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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *SslCert) validate() error {

	if err := dcl.Required(r, "commonName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "instance"); err != nil {
		return err
	}
	return nil
}

func sslCertGetURL(userBasePath string, r *SslCert) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/sslCerts/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

func sslCertListURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/sslCerts", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func sslCertCreateURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/sslCerts", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func sslCertDeleteURL(userBasePath string, r *SslCert) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/sslCerts/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

// sslCertApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type sslCertApiOperation interface {
	do(context.Context, *SslCert, *Client) error
}

func (c *Client) listSslCertRaw(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]byte, error) {
	u, err := sslCertListURL(c.Config.BasePath, project, instance)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SslCertMaxPage {
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

type listSslCertOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listSslCert(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]*SslCert, string, error) {
	b, err := c.listSslCertRaw(ctx, project, instance, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSslCertOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*SslCert
	for _, v := range m.Items {
		res, err := unmarshalMapSslCert(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Instance = &instance
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSslCert(ctx context.Context, f func(*SslCert) bool, resources []*SslCert) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSslCert(ctx, res)
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

type deleteSslCertOperation struct{}

func (op *deleteSslCertOperation) do(ctx context.Context, r *SslCert, c *Client) error {

	_, err := c.GetSslCert(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("SslCert not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSslCert checking for existence. error: %v", err)
		return err
	}

	u, err := sslCertDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetSslCert(ctx, r.urlNormalized())
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
type createSslCertOperation struct {
	response map[string]interface{}
}

func (op *createSslCertOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSslCertOperation) do(ctx context.Context, r *SslCert, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, instance := r.createFields()
	u, err := sslCertCreateURL(c.Config.BasePath, project, instance)

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
	var o operations.SQLCreateCertOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	sha1Fingerprint, ok := op.response["sha1Fingerprint"].(string)
	if !ok {
		return fmt.Errorf("expected sha1Fingerprint to be a string, was %T", sha1Fingerprint)
	}
	r.Name = &sha1Fingerprint

	if _, err := c.GetSslCert(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSslCertRaw(ctx context.Context, r *SslCert) ([]byte, error) {

	u, err := sslCertGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) sslCertDiffsForRawDesired(ctx context.Context, rawDesired *SslCert, opts ...dcl.ApplyOption) (initial, desired *SslCert, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *SslCert
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*SslCert); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected SslCert, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeSslCertDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSslCert(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a SslCert resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve SslCert resource: %v", err)
		}
		c.Config.Logger.Info("Found that SslCert resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSslCertDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for SslCert: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for SslCert: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSslCertInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for SslCert: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSslCertDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for SslCert: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSslCert(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSslCertInitialState(rawInitial, rawDesired *SslCert) (*SslCert, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSslCertDesiredState(rawDesired, rawInitial *SslCert, opts ...dcl.ApplyOption) (*SslCert, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.CommonName, rawInitial.CommonName) {
		rawDesired.CommonName = rawInitial.CommonName
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Instance, rawInitial.Instance) {
		rawDesired.Instance = rawInitial.Instance
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeSslCertNewState(c *Client, rawNew, rawDesired *SslCert) (*SslCert, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CertSerialNumber) && dcl.IsEmptyValueIndirect(rawDesired.CertSerialNumber) {
		rawNew.CertSerialNumber = rawDesired.CertSerialNumber
	} else {
		if dcl.StringCanonicalize(rawDesired.CertSerialNumber, rawNew.CertSerialNumber) {
			rawNew.CertSerialNumber = rawDesired.CertSerialNumber
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Cert) && dcl.IsEmptyValueIndirect(rawDesired.Cert) {
		rawNew.Cert = rawDesired.Cert
	} else {
		if dcl.StringCanonicalize(rawDesired.Cert, rawNew.Cert) {
			rawNew.Cert = rawDesired.Cert
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CommonName) && dcl.IsEmptyValueIndirect(rawDesired.CommonName) {
		rawNew.CommonName = rawDesired.CommonName
	} else {
		if dcl.StringCanonicalize(rawDesired.CommonName, rawNew.CommonName) {
			rawNew.CommonName = rawDesired.CommonName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpirationTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpirationTime) {
		rawNew.ExpirationTime = rawDesired.ExpirationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Instance) && dcl.IsEmptyValueIndirect(rawDesired.Instance) {
		rawNew.Instance = rawDesired.Instance
	} else {
		if dcl.StringCanonicalize(rawDesired.Instance, rawNew.Instance) {
			rawNew.Instance = rawDesired.Instance
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffSslCert(c *Client, desired, actual *SslCert, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CertSerialNumber, actual.CertSerialNumber, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertSerialNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Cert, actual.Cert, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Cert")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommonName, actual.CommonName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommonName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationTime, actual.ExpirationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpirationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha1Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *SslCert) urlNormalized() *SslCert {
	normalized := dcl.Copy(*r).(SslCert)
	normalized.CertSerialNumber = dcl.SelfLinkToName(r.CertSerialNumber)
	normalized.Cert = dcl.SelfLinkToName(r.Cert)
	normalized.CommonName = dcl.SelfLinkToName(r.CommonName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Instance = dcl.SelfLinkToName(r.Instance)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *SslCert) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslCert) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance)
}

func (r *SslCert) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslCert) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the SslCert resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *SslCert) marshal(c *Client) ([]byte, error) {
	m, err := expandSslCert(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling SslCert: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSslCert decodes JSON responses into the SslCert resource schema.
func unmarshalSslCert(b []byte, c *Client) (*SslCert, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSslCert(m, c)
}

func unmarshalMapSslCert(m map[string]interface{}, c *Client) (*SslCert, error) {

	return flattenSslCert(c, m), nil
}

// expandSslCert expands SslCert into a JSON request object.
func expandSslCert(c *Client, f *SslCert) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CertSerialNumber; !dcl.IsEmptyValueIndirect(v) {
		m["certSerialNumber"] = v
	}
	if v := f.Cert; !dcl.IsEmptyValueIndirect(v) {
		m["cert"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Fingerprint"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenSslCert flattens SslCert from a JSON request object into the
// SslCert type.
func flattenSslCert(c *Client, i interface{}) *SslCert {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &SslCert{}
	res.CertSerialNumber = dcl.FlattenString(m["certSerialNumber"])
	res.Cert = dcl.FlattenString(m["cert"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.CommonName = dcl.FlattenString(m["commonName"])
	res.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["sha1Fingerprint"]))
	res.Instance = dcl.FlattenString(m["instance"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *SslCert) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSslCert(b, c)
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
		if nr.Instance == nil && ncr.Instance == nil {
			c.Config.Logger.Info("Both Instance fields null - considering equal.")
		} else if nr.Instance == nil || ncr.Instance == nil {
			c.Config.Logger.Info("Only one Instance field is null - considering unequal.")
			return false
		} else if *nr.Instance != *ncr.Instance {
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

type sslCertDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         sslCertApiOperation
}

func convertFieldDiffToSslCertOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]sslCertDiff, error) {
	var diffs []sslCertDiff
	for _, op := range ops {
		diff := sslCertDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTosslCertApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTosslCertApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (sslCertApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
