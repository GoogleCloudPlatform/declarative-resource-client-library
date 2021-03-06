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

func (r *TargetVpnGateway) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func targetVpnGatewayGetURL(userBasePath string, r *TargetVpnGateway) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func targetVpnGatewayListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetVpnGateways", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetVpnGatewayCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetVpnGateways", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetVpnGatewayDeleteURL(userBasePath string, r *TargetVpnGateway) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// targetVpnGatewayApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type targetVpnGatewayApiOperation interface {
	do(context.Context, *TargetVpnGateway, *Client) error
}

func (c *Client) listTargetVpnGatewayRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := targetVpnGatewayListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TargetVpnGatewayMaxPage {
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

type listTargetVpnGatewayOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listTargetVpnGateway(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*TargetVpnGateway, string, error) {
	b, err := c.listTargetVpnGatewayRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTargetVpnGatewayOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TargetVpnGateway
	for _, v := range m.Items {
		res, err := unmarshalMapTargetVpnGateway(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTargetVpnGateway(ctx context.Context, f func(*TargetVpnGateway) bool, resources []*TargetVpnGateway) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTargetVpnGateway(ctx, res)
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

type deleteTargetVpnGatewayOperation struct{}

func (op *deleteTargetVpnGatewayOperation) do(ctx context.Context, r *TargetVpnGateway, c *Client) error {
	r, err := c.GetTargetVpnGateway(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("TargetVpnGateway not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTargetVpnGateway checking for existence. error: %v", err)
		return err
	}

	u, err := targetVpnGatewayDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetTargetVpnGateway(ctx, r.URLNormalized())
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
type createTargetVpnGatewayOperation struct {
	response map[string]interface{}
}

func (op *createTargetVpnGatewayOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTargetVpnGatewayOperation) do(ctx context.Context, r *TargetVpnGateway, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := targetVpnGatewayCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetTargetVpnGateway(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTargetVpnGatewayRaw(ctx context.Context, r *TargetVpnGateway) ([]byte, error) {

	u, err := targetVpnGatewayGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) targetVpnGatewayDiffsForRawDesired(ctx context.Context, rawDesired *TargetVpnGateway, opts ...dcl.ApplyOption) (initial, desired *TargetVpnGateway, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TargetVpnGateway
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TargetVpnGateway); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected TargetVpnGateway, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTargetVpnGateway(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a TargetVpnGateway resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TargetVpnGateway resource: %v", err)
		}
		c.Config.Logger.Info("Found that TargetVpnGateway resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTargetVpnGatewayDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for TargetVpnGateway: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for TargetVpnGateway: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTargetVpnGatewayInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for TargetVpnGateway: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTargetVpnGatewayDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for TargetVpnGateway: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTargetVpnGateway(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTargetVpnGatewayInitialState(rawInitial, rawDesired *TargetVpnGateway) (*TargetVpnGateway, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTargetVpnGatewayDesiredState(rawDesired, rawInitial *TargetVpnGateway, opts ...dcl.ApplyOption) (*TargetVpnGateway, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &TargetVpnGateway{}
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
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		canonicalDesired.Region = rawInitial.Region
	} else {
		canonicalDesired.Region = rawDesired.Region
	}
	if dcl.StringCanonicalize(rawDesired.Network, rawInitial.Network) {
		canonicalDesired.Network = rawInitial.Network
	} else {
		canonicalDesired.Network = rawDesired.Network
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeTargetVpnGatewayNewState(c *Client, rawNew, rawDesired *TargetVpnGateway) (*TargetVpnGateway, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.StringCanonicalize(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tunnel) && dcl.IsEmptyValueIndirect(rawDesired.Tunnel) {
		rawNew.Tunnel = rawDesired.Tunnel
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ForwardingRule) && dcl.IsEmptyValueIndirect(rawDesired.ForwardingRule) {
		rawNew.ForwardingRule = rawDesired.ForwardingRule
	} else {
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
func diffTargetVpnGateway(c *Client, desired, actual *TargetVpnGateway, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tunnel, actual.Tunnel, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tunnels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ForwardingRule, actual.ForwardingRule, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ForwardingRules")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}

func (r *TargetVpnGateway) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetVpnGateway) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *TargetVpnGateway) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetVpnGateway) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TargetVpnGateway resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TargetVpnGateway) marshal(c *Client) ([]byte, error) {
	m, err := expandTargetVpnGateway(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TargetVpnGateway: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTargetVpnGateway decodes JSON responses into the TargetVpnGateway resource schema.
func unmarshalTargetVpnGateway(b []byte, c *Client) (*TargetVpnGateway, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTargetVpnGateway(m, c)
}

func unmarshalMapTargetVpnGateway(m map[string]interface{}, c *Client) (*TargetVpnGateway, error) {

	return flattenTargetVpnGateway(c, m), nil
}

// expandTargetVpnGateway expands TargetVpnGateway into a JSON request object.
func expandTargetVpnGateway(c *Client, f *TargetVpnGateway) (map[string]interface{}, error) {
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
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Tunnel; !dcl.IsEmptyValueIndirect(v) {
		m["tunnels"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.ForwardingRule; !dcl.IsEmptyValueIndirect(v) {
		m["forwardingRules"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenTargetVpnGateway flattens TargetVpnGateway from a JSON request object into the
// TargetVpnGateway type.
func flattenTargetVpnGateway(c *Client, i interface{}) *TargetVpnGateway {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &TargetVpnGateway{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Region = dcl.FlattenString(m["region"])
	res.Network = dcl.FlattenString(m["network"])
	res.Tunnel = dcl.FlattenStringSlice(m["tunnels"])
	res.Status = flattenTargetVpnGatewayStatusEnum(m["status"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.ForwardingRule = dcl.FlattenStringSlice(m["forwardingRules"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// flattenTargetVpnGatewayStatusEnumMap flattens the contents of TargetVpnGatewayStatusEnum from a JSON
// response object.
func flattenTargetVpnGatewayStatusEnumMap(c *Client, i interface{}) map[string]TargetVpnGatewayStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetVpnGatewayStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]TargetVpnGatewayStatusEnum{}
	}

	items := make(map[string]TargetVpnGatewayStatusEnum)
	for k, item := range a {
		items[k] = *flattenTargetVpnGatewayStatusEnum(item.(interface{}))
	}

	return items
}

// flattenTargetVpnGatewayStatusEnumSlice flattens the contents of TargetVpnGatewayStatusEnum from a JSON
// response object.
func flattenTargetVpnGatewayStatusEnumSlice(c *Client, i interface{}) []TargetVpnGatewayStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetVpnGatewayStatusEnum{}
	}

	if len(a) == 0 {
		return []TargetVpnGatewayStatusEnum{}
	}

	items := make([]TargetVpnGatewayStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetVpnGatewayStatusEnum(item.(interface{})))
	}

	return items
}

// flattenTargetVpnGatewayStatusEnum asserts that an interface is a string, and returns a
// pointer to a *TargetVpnGatewayStatusEnum with the same value as that string.
func flattenTargetVpnGatewayStatusEnum(i interface{}) *TargetVpnGatewayStatusEnum {
	s, ok := i.(string)
	if !ok {
		return TargetVpnGatewayStatusEnumRef("")
	}

	return TargetVpnGatewayStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TargetVpnGateway) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTargetVpnGateway(b, c)
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
		if nr.Region == nil && ncr.Region == nil {
			c.Config.Logger.Info("Both Region fields null - considering equal.")
		} else if nr.Region == nil || ncr.Region == nil {
			c.Config.Logger.Info("Only one Region field is null - considering unequal.")
			return false
		} else if *nr.Region != *ncr.Region {
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

type targetVpnGatewayDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         targetVpnGatewayApiOperation
}

func convertFieldDiffsToTargetVpnGatewayDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]targetVpnGatewayDiff, error) {
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
	var diffs []targetVpnGatewayDiff
	// For each operation name, create a targetVpnGatewayDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := targetVpnGatewayDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTargetVpnGatewayApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTargetVpnGatewayApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (targetVpnGatewayApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
