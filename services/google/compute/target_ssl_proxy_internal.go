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

func (r *TargetSslProxy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func targetSslProxyGetURL(userBasePath string, r *TargetSslProxy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/targetSslProxies/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func targetSslProxyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/targetSslProxies", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetSslProxyCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/targetSslProxies", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetSslProxyDeleteURL(userBasePath string, r *TargetSslProxy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/targetSslProxies/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// targetSslProxyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type targetSslProxyApiOperation interface {
	do(context.Context, *TargetSslProxy, *Client) error
}

func (c *Client) listTargetSslProxyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := targetSslProxyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TargetSslProxyMaxPage {
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

type listTargetSslProxyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listTargetSslProxy(ctx context.Context, project, pageToken string, pageSize int32) ([]*TargetSslProxy, string, error) {
	b, err := c.listTargetSslProxyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTargetSslProxyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TargetSslProxy
	for _, v := range m.Items {
		res := flattenTargetSslProxy(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTargetSslProxy(ctx context.Context, f func(*TargetSslProxy) bool, resources []*TargetSslProxy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTargetSslProxy(ctx, res)
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

type deleteTargetSslProxyOperation struct{}

func (op *deleteTargetSslProxyOperation) do(ctx context.Context, r *TargetSslProxy, c *Client) error {

	_, err := c.GetTargetSslProxy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("TargetSslProxy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTargetSslProxy checking for existence. error: %v", err)
		return err
	}

	u, err := targetSslProxyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetTargetSslProxy(ctx, r.urlNormalized())
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
type createTargetSslProxyOperation struct {
	response map[string]interface{}
}

func (op *createTargetSslProxyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTargetSslProxyOperation) do(ctx context.Context, r *TargetSslProxy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := targetSslProxyCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetTargetSslProxy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTargetSslProxyRaw(ctx context.Context, r *TargetSslProxy) ([]byte, error) {
	if dcl.IsZeroValue(r.ProxyHeader) {
		r.ProxyHeader = TargetSslProxyProxyHeaderEnumRef("NONE")
	}

	u, err := targetSslProxyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) targetSslProxyDiffsForRawDesired(ctx context.Context, rawDesired *TargetSslProxy, opts ...dcl.ApplyOption) (initial, desired *TargetSslProxy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TargetSslProxy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TargetSslProxy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected TargetSslProxy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTargetSslProxy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a TargetSslProxy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TargetSslProxy resource: %v", err)
		}
		c.Config.Logger.Info("Found that TargetSslProxy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTargetSslProxyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for TargetSslProxy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for TargetSslProxy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTargetSslProxyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for TargetSslProxy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTargetSslProxyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for TargetSslProxy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTargetSslProxy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTargetSslProxyInitialState(rawInitial, rawDesired *TargetSslProxy) (*TargetSslProxy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTargetSslProxyDesiredState(rawDesired, rawInitial *TargetSslProxy, opts ...dcl.ApplyOption) (*TargetSslProxy, error) {

	if dcl.IsZeroValue(rawDesired.ProxyHeader) {
		rawDesired.ProxyHeader = TargetSslProxyProxyHeaderEnumRef("NONE")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.Service, rawInitial.Service) {
		rawDesired.Service = rawInitial.Service
	}
	if dcl.IsZeroValue(rawDesired.SslCertificates) {
		rawDesired.SslCertificates = rawInitial.SslCertificates
	}
	if dcl.IsZeroValue(rawDesired.ProxyHeader) {
		rawDesired.ProxyHeader = rawInitial.ProxyHeader
	}
	if dcl.StringCanonicalize(rawDesired.SslPolicy, rawInitial.SslPolicy) {
		rawDesired.SslPolicy = rawInitial.SslPolicy
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeTargetSslProxyNewState(c *Client, rawNew, rawDesired *TargetSslProxy) (*TargetSslProxy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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

	if dcl.IsEmptyValueIndirect(rawNew.Service) && dcl.IsEmptyValueIndirect(rawDesired.Service) {
		rawNew.Service = rawDesired.Service
	} else {
		if dcl.StringCanonicalize(rawDesired.Service, rawNew.Service) {
			rawNew.Service = rawDesired.Service
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SslCertificates) && dcl.IsEmptyValueIndirect(rawDesired.SslCertificates) {
		rawNew.SslCertificates = rawDesired.SslCertificates
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProxyHeader) && dcl.IsEmptyValueIndirect(rawDesired.ProxyHeader) {
		rawNew.ProxyHeader = rawDesired.ProxyHeader
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SslPolicy) && dcl.IsEmptyValueIndirect(rawDesired.SslPolicy) {
		rawNew.SslPolicy = rawDesired.SslPolicy
	} else {
		if dcl.StringCanonicalize(rawDesired.SslPolicy, rawNew.SslPolicy) {
			rawNew.SslPolicy = rawDesired.SslPolicy
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
func diffTargetSslProxy(c *Client, desired, actual *TargetSslProxy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SslCertificates, actual.SslCertificates, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SslCertificates")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SslPolicy, actual.SslPolicy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SslPolicy")); len(ds) != 0 || err != nil {
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
func (r *TargetSslProxy) urlNormalized() *TargetSslProxy {
	normalized := dcl.Copy(*r).(TargetSslProxy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Service = dcl.SelfLinkToName(r.Service)
	normalized.SslPolicy = dcl.SelfLinkToName(r.SslPolicy)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *TargetSslProxy) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetSslProxy) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *TargetSslProxy) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetSslProxy) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TargetSslProxy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TargetSslProxy) marshal(c *Client) ([]byte, error) {
	m, err := expandTargetSslProxy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TargetSslProxy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTargetSslProxy decodes JSON responses into the TargetSslProxy resource schema.
func unmarshalTargetSslProxy(b []byte, c *Client) (*TargetSslProxy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTargetSslProxy(m, c)
}

func unmarshalMapTargetSslProxy(m map[string]interface{}, c *Client) (*TargetSslProxy, error) {

	return flattenTargetSslProxy(c, m), nil
}

// expandTargetSslProxy expands TargetSslProxy into a JSON request object.
func expandTargetSslProxy(c *Client, f *TargetSslProxy) (map[string]interface{}, error) {
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
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.SslCertificates; !dcl.IsEmptyValueIndirect(v) {
		m["sslCertificates"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}
	if v := f.SslPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["sslPolicy"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenTargetSslProxy flattens TargetSslProxy from a JSON request object into the
// TargetSslProxy type.
func flattenTargetSslProxy(c *Client, i interface{}) *TargetSslProxy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &TargetSslProxy{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Service = dcl.FlattenString(m["service"])
	res.SslCertificates = dcl.FlattenStringSlice(m["sslCertificates"])
	res.ProxyHeader = flattenTargetSslProxyProxyHeaderEnum(m["proxyHeader"])
	if _, ok := m["proxyHeader"]; !ok {
		c.Config.Logger.Info("Using default value for proxyHeader")
		res.ProxyHeader = TargetSslProxyProxyHeaderEnumRef("NONE")
	}
	res.SslPolicy = dcl.FlattenString(m["sslPolicy"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// flattenTargetSslProxyProxyHeaderEnumSlice flattens the contents of TargetSslProxyProxyHeaderEnum from a JSON
// response object.
func flattenTargetSslProxyProxyHeaderEnumSlice(c *Client, i interface{}) []TargetSslProxyProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetSslProxyProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []TargetSslProxyProxyHeaderEnum{}
	}

	items := make([]TargetSslProxyProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetSslProxyProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenTargetSslProxyProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *TargetSslProxyProxyHeaderEnum with the same value as that string.
func flattenTargetSslProxyProxyHeaderEnum(i interface{}) *TargetSslProxyProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return TargetSslProxyProxyHeaderEnumRef("")
	}

	return TargetSslProxyProxyHeaderEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TargetSslProxy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTargetSslProxy(b, c)
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

type targetSslProxyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         targetSslProxyApiOperation
}

func convertFieldDiffToTargetSslProxyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]targetSslProxyDiff, error) {
	var diffs []targetSslProxyDiff
	for _, op := range ops {
		diff := targetSslProxyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTotargetSslProxyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTotargetSslProxyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (targetSslProxyApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
