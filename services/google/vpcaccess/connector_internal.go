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
package vpcaccess

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

func (r *Connector) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Subnet) {
		if err := r.Subnet.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConnectorSubnet) validate() error {
	return nil
}

func connectorGetURL(userBasePath string, r *Connector) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors/{{name}}", "https://vpcaccess.googleapis.com/v1/", userBasePath, params), nil
}

func connectorListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors", "https://vpcaccess.googleapis.com/v1/", userBasePath, params), nil

}

func connectorCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors?connectorId={{name}}", "https://vpcaccess.googleapis.com/v1/", userBasePath, params), nil

}

func connectorDeleteURL(userBasePath string, r *Connector) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors/{{name}}", "https://vpcaccess.googleapis.com/v1/", userBasePath, params), nil
}

// connectorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type connectorApiOperation interface {
	do(context.Context, *Connector, *Client) error
}

func (c *Client) listConnectorRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := connectorListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ConnectorMaxPage {
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

type listConnectorOperation struct {
	Connectors []map[string]interface{} `json:"connectors"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listConnector(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Connector, string, error) {
	b, err := c.listConnectorRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listConnectorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Connector
	for _, v := range m.Connectors {
		res := flattenConnector(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllConnector(ctx context.Context, f func(*Connector) bool, resources []*Connector) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteConnector(ctx, res)
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

type deleteConnectorOperation struct{}

func (op *deleteConnectorOperation) do(ctx context.Context, r *Connector, c *Client) error {

	_, err := c.GetConnector(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Connector not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetConnector checking for existence. error: %v", err)
		return err
	}

	u, err := connectorDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://vpcaccess.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetConnector(ctx, r.urlNormalized())
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
type createConnectorOperation struct {
	response map[string]interface{}
}

func (op *createConnectorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createConnectorOperation) do(ctx context.Context, r *Connector, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := connectorCreateURL(c.Config.BasePath, project, location, name)

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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://vpcaccess.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetConnector(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getConnectorRaw(ctx context.Context, r *Connector) ([]byte, error) {

	u, err := connectorGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) connectorDiffsForRawDesired(ctx context.Context, rawDesired *Connector, opts ...dcl.ApplyOption) (initial, desired *Connector, diffs []connectorDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Connector
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Connector); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Connector, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetConnector(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Connector resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Connector resource: %v", err)
		}
		c.Config.Logger.Info("Found that Connector resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeConnectorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Connector: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Connector: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeConnectorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Connector: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeConnectorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Connector: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffConnector(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeConnectorInitialState(rawInitial, rawDesired *Connector) (*Connector, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeConnectorDesiredState(rawDesired, rawInitial *Connector, opts ...dcl.ApplyOption) (*Connector, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Subnet = canonicalizeConnectorSubnet(rawDesired.Subnet, nil, opts...)

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawInitial.IPCidrRange) {
		rawDesired.IPCidrRange = rawInitial.IPCidrRange
	}
	if dcl.IsZeroValue(rawDesired.MinThroughput) {
		rawDesired.MinThroughput = rawInitial.MinThroughput
	}
	if dcl.IsZeroValue(rawDesired.MaxThroughput) {
		rawDesired.MaxThroughput = rawInitial.MaxThroughput
	}
	rawDesired.Subnet = canonicalizeConnectorSubnet(rawDesired.Subnet, rawInitial.Subnet, opts...)
	if dcl.StringCanonicalize(rawDesired.MachineType, rawInitial.MachineType) {
		rawDesired.MachineType = rawInitial.MachineType
	}
	if dcl.IsZeroValue(rawDesired.MinInstances) {
		rawDesired.MinInstances = rawInitial.MinInstances
	}
	if dcl.IsZeroValue(rawDesired.MaxInstances) {
		rawDesired.MaxInstances = rawInitial.MaxInstances
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeConnectorNewState(c *Client, rawNew, rawDesired *Connector) (*Connector, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPCidrRange) && dcl.IsEmptyValueIndirect(rawDesired.IPCidrRange) {
		rawNew.IPCidrRange = rawDesired.IPCidrRange
	} else {
		if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawNew.IPCidrRange) {
			rawNew.IPCidrRange = rawDesired.IPCidrRange
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinThroughput) && dcl.IsEmptyValueIndirect(rawDesired.MinThroughput) {
		rawNew.MinThroughput = rawDesired.MinThroughput
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxThroughput) && dcl.IsEmptyValueIndirect(rawDesired.MaxThroughput) {
		rawNew.MaxThroughput = rawDesired.MaxThroughput
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectedProjects) && dcl.IsEmptyValueIndirect(rawDesired.ConnectedProjects) {
		rawNew.ConnectedProjects = rawDesired.ConnectedProjects
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnet) && dcl.IsEmptyValueIndirect(rawDesired.Subnet) {
		rawNew.Subnet = rawDesired.Subnet
	} else {
		rawNew.Subnet = canonicalizeNewConnectorSubnet(c, rawDesired.Subnet, rawNew.Subnet)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MachineType) && dcl.IsEmptyValueIndirect(rawDesired.MachineType) {
		rawNew.MachineType = rawDesired.MachineType
	} else {
		if dcl.StringCanonicalize(rawDesired.MachineType, rawNew.MachineType) {
			rawNew.MachineType = rawDesired.MachineType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinInstances) && dcl.IsEmptyValueIndirect(rawDesired.MinInstances) {
		rawNew.MinInstances = rawDesired.MinInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxInstances) && dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) {
		rawNew.MaxInstances = rawDesired.MaxInstances
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeConnectorSubnet(des, initial *ConnectorSubnet, opts ...dcl.ApplyOption) *ConnectorSubnet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		des.ProjectId = initial.ProjectId
	}

	return des
}

func canonicalizeNewConnectorSubnet(c *Client, des, nw *ConnectorSubnet) *ConnectorSubnet {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}

	return nw
}

func canonicalizeNewConnectorSubnetSet(c *Client, des, nw []ConnectorSubnet) []ConnectorSubnet {
	if des == nil {
		return nw
	}
	var reorderedNew []ConnectorSubnet
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareConnectorSubnetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewConnectorSubnetSlice(c *Client, des, nw []ConnectorSubnet) []ConnectorSubnet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConnectorSubnet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConnectorSubnet(c, &d, &n))
	}

	return items
}

type connectorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         connectorApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffConnector(c *Client, desired, actual *Connector, opts ...dcl.ApplyOption) ([]connectorDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []connectorDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.IPCidrRange, actual.IPCidrRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPCidrRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MinThroughput, actual.MinThroughput, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinThroughput")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MaxThroughput, actual.MaxThroughput, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxThroughput")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.ConnectedProjects, actual.ConnectedProjects, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectedProjects")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Subnet, actual.Subnet, dcl.Info{ObjectFunction: compareConnectorSubnetNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subnet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MinInstances, actual.MinInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToConnectorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
	var deduped []connectorDiff
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
func compareConnectorSubnetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConnectorSubnet)
	if !ok {
		desiredNotPointer, ok := d.(ConnectorSubnet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectorSubnet or *ConnectorSubnet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConnectorSubnet)
	if !ok {
		actualNotPointer, ok := a.(ConnectorSubnet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectorSubnet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Connector) urlNormalized() *Connector {
	normalized := dcl.Copy(*r).(Connector)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.IPCidrRange = dcl.SelfLinkToName(r.IPCidrRange)
	normalized.MachineType = dcl.SelfLinkToName(r.MachineType)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Connector) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connector) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connector) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connector) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Connector resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Connector) marshal(c *Client) ([]byte, error) {
	m, err := expandConnector(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Connector: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalConnector decodes JSON responses into the Connector resource schema.
func unmarshalConnector(b []byte, c *Client) (*Connector, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapConnector(m, c)
}

func unmarshalMapConnector(m map[string]interface{}, c *Client) (*Connector, error) {

	return flattenConnector(c, m), nil
}

// expandConnector expands Connector into a JSON request object.
func expandConnector(c *Client, f *Connector) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.IPCidrRange; !dcl.IsEmptyValueIndirect(v) {
		m["ipCidrRange"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.MinThroughput; !dcl.IsEmptyValueIndirect(v) {
		m["minThroughput"] = v
	}
	if v := f.MaxThroughput; !dcl.IsEmptyValueIndirect(v) {
		m["maxThroughput"] = v
	}
	if v := f.ConnectedProjects; !dcl.IsEmptyValueIndirect(v) {
		m["connectedProjects"] = v
	}
	if v, err := expandConnectorSubnet(c, f.Subnet); err != nil {
		return nil, fmt.Errorf("error expanding Subnet into subnet: %w", err)
	} else if v != nil {
		m["subnet"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.MinInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minInstances"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
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

// flattenConnector flattens Connector from a JSON request object into the
// Connector type.
func flattenConnector(c *Client, i interface{}) *Connector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Connector{}
	res.Name = dcl.FlattenString(m["name"])
	res.Network = dcl.FlattenString(m["network"])
	res.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])
	res.State = flattenConnectorStateEnum(m["state"])
	res.MinThroughput = dcl.FlattenInteger(m["minThroughput"])
	res.MaxThroughput = dcl.FlattenInteger(m["maxThroughput"])
	res.ConnectedProjects = dcl.FlattenStringSlice(m["connectedProjects"])
	res.Subnet = flattenConnectorSubnet(c, m["subnet"])
	res.MachineType = dcl.FlattenString(m["machineType"])
	res.MinInstances = dcl.FlattenInteger(m["minInstances"])
	res.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandConnectorSubnetMap expands the contents of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnetMap(c *Client, f map[string]ConnectorSubnet) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConnectorSubnet(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConnectorSubnetSlice expands the contents of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnetSlice(c *Client, f []ConnectorSubnet) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConnectorSubnet(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConnectorSubnetMap flattens the contents of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnetMap(c *Client, i interface{}) map[string]ConnectorSubnet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectorSubnet{}
	}

	if len(a) == 0 {
		return map[string]ConnectorSubnet{}
	}

	items := make(map[string]ConnectorSubnet)
	for k, item := range a {
		items[k] = *flattenConnectorSubnet(c, item.(map[string]interface{}))
	}

	return items
}

// flattenConnectorSubnetSlice flattens the contents of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnetSlice(c *Client, i interface{}) []ConnectorSubnet {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectorSubnet{}
	}

	if len(a) == 0 {
		return []ConnectorSubnet{}
	}

	items := make([]ConnectorSubnet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectorSubnet(c, item.(map[string]interface{})))
	}

	return items
}

// expandConnectorSubnet expands an instance of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnet(c *Client, f *ConnectorSubnet) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}

	return m, nil
}

// flattenConnectorSubnet flattens an instance of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnet(c *Client, i interface{}) *ConnectorSubnet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConnectorSubnet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConnectorSubnet
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ProjectId = dcl.FlattenString(m["projectId"])

	return r
}

// flattenConnectorStateEnumSlice flattens the contents of ConnectorStateEnum from a JSON
// response object.
func flattenConnectorStateEnumSlice(c *Client, i interface{}) []ConnectorStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectorStateEnum{}
	}

	if len(a) == 0 {
		return []ConnectorStateEnum{}
	}

	items := make([]ConnectorStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectorStateEnum(item.(interface{})))
	}

	return items
}

// flattenConnectorStateEnum asserts that an interface is a string, and returns a
// pointer to a *ConnectorStateEnum with the same value as that string.
func flattenConnectorStateEnum(i interface{}) *ConnectorStateEnum {
	s, ok := i.(string)
	if !ok {
		return ConnectorStateEnumRef("")
	}

	return ConnectorStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Connector) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalConnector(b, c)
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

func convertFieldDiffToConnectorDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]connectorDiff, error) {
	var diffs []connectorDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := connectorDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameToconnectorApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameToconnectorApiOperation(op string, opts ...dcl.ApplyOption) (connectorApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
