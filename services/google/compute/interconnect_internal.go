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
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Interconnect) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	if err := dcl.Required(r, "linkType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "interconnectType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "customerName"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *InterconnectExpectedOutages) validate() error {
	return nil
}
func (r *InterconnectCircuitInfos) validate() error {
	return nil
}

func interconnectGetURL(userBasePath string, r *Interconnect) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/interconnects/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func interconnectListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/interconnects", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func interconnectCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/interconnects", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func interconnectDeleteURL(userBasePath string, r *Interconnect) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/interconnects/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// interconnectApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type interconnectApiOperation interface {
	do(context.Context, *Interconnect, *Client) error
}

// newUpdateInterconnectPatchRequest creates a request for an
// Interconnect resource's Patch update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInterconnectPatchRequest(ctx context.Context, f *Interconnect, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		req["location"] = v
	}
	if v := f.LinkType; !dcl.IsEmptyValueIndirect(v) {
		req["linkType"] = v
	}
	if v := f.RequestedLinkCount; !dcl.IsEmptyValueIndirect(v) {
		req["requestedLinkCount"] = v
	}
	if v := f.InterconnectType; !dcl.IsEmptyValueIndirect(v) {
		req["interconnectType"] = v
	}
	if v := f.AdminEnabled; !dcl.IsEmptyValueIndirect(v) {
		req["adminEnabled"] = v
	}
	if v := f.NocContactEmail; !dcl.IsEmptyValueIndirect(v) {
		req["nocContactEmail"] = v
	}
	if v := f.CustomerName; !dcl.IsEmptyValueIndirect(v) {
		req["customerName"] = v
	}
	return req, nil
}

// marshalUpdateInterconnectPatchRequest converts the update into
// the final JSON request body.
func marshalUpdateInterconnectPatchRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInterconnectPatchOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInterconnectPatchOperation) do(ctx context.Context, r *Interconnect, c *Client) error {
	_, err := c.GetInterconnect(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Patch")
	if err != nil {
		return err
	}

	req, err := newUpdateInterconnectPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInterconnectPatchRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInterconnectRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := interconnectListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InterconnectMaxPage {
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

type listInterconnectOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listInterconnect(ctx context.Context, project, pageToken string, pageSize int32) ([]*Interconnect, string, error) {
	b, err := c.listInterconnectRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInterconnectOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Interconnect
	for _, v := range m.Items {
		res := flattenInterconnect(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInterconnect(ctx context.Context, f func(*Interconnect) bool, resources []*Interconnect) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInterconnect(ctx, res)
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

type deleteInterconnectOperation struct{}

func (op *deleteInterconnectOperation) do(ctx context.Context, r *Interconnect, c *Client) error {

	_, err := c.GetInterconnect(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Interconnect not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInterconnect checking for existence. error: %v", err)
		return err
	}

	u, err := interconnectDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetInterconnect(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createInterconnectOperation struct{}

func (op *createInterconnectOperation) do(ctx context.Context, r *Interconnect, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := interconnectCreateURL(c.Config.BasePath, project)

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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetInterconnect(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getInterconnectRaw(ctx context.Context, r *Interconnect) ([]byte, error) {

	u, err := interconnectGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) interconnectDiffsForRawDesired(ctx context.Context, rawDesired *Interconnect, opts ...dcl.ApplyOption) (initial, desired *Interconnect, diffs []interconnectDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Interconnect
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Interconnect); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Interconnect, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInterconnect(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Interconnect resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Interconnect resource: %v", err)
		}
		c.Config.Logger.Info("Found that Interconnect resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInterconnectDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Interconnect: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Interconnect: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInterconnectInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Interconnect: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInterconnectDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Interconnect: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInterconnect(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInterconnectInitialState(rawInitial, rawDesired *Interconnect) (*Interconnect, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInterconnectDesiredState(rawDesired, rawInitial *Interconnect, opts ...dcl.ApplyOption) (*Interconnect, error) {

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Interconnect); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Interconnect, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.IsZeroValue(rawDesired.LinkType) {
		rawDesired.LinkType = rawInitial.LinkType
	}
	if dcl.IsZeroValue(rawDesired.RequestedLinkCount) {
		rawDesired.RequestedLinkCount = rawInitial.RequestedLinkCount
	}
	if dcl.IsZeroValue(rawDesired.InterconnectType) {
		rawDesired.InterconnectType = rawInitial.InterconnectType
	}
	if dcl.IsZeroValue(rawDesired.AdminEnabled) {
		rawDesired.AdminEnabled = rawInitial.AdminEnabled
	}
	if dcl.IsZeroValue(rawDesired.NocContactEmail) {
		rawDesired.NocContactEmail = rawInitial.NocContactEmail
	}
	if dcl.IsZeroValue(rawDesired.CustomerName) {
		rawDesired.CustomerName = rawInitial.CustomerName
	}
	if dcl.IsZeroValue(rawDesired.OperationalStatus) {
		rawDesired.OperationalStatus = rawInitial.OperationalStatus
	}
	if dcl.IsZeroValue(rawDesired.ProvisionedLinkCount) {
		rawDesired.ProvisionedLinkCount = rawInitial.ProvisionedLinkCount
	}
	if dcl.IsZeroValue(rawDesired.InterconnectAttachments) {
		rawDesired.InterconnectAttachments = rawInitial.InterconnectAttachments
	}
	if dcl.IsZeroValue(rawDesired.PeerIPAddress) {
		rawDesired.PeerIPAddress = rawInitial.PeerIPAddress
	}
	if dcl.IsZeroValue(rawDesired.GoogleIPAddress) {
		rawDesired.GoogleIPAddress = rawInitial.GoogleIPAddress
	}
	if dcl.IsZeroValue(rawDesired.GoogleReferenceId) {
		rawDesired.GoogleReferenceId = rawInitial.GoogleReferenceId
	}
	if dcl.IsZeroValue(rawDesired.ExpectedOutages) {
		rawDesired.ExpectedOutages = rawInitial.ExpectedOutages
	}
	if dcl.IsZeroValue(rawDesired.CircuitInfos) {
		rawDesired.CircuitInfos = rawInitial.CircuitInfos
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeInterconnectNewState(c *Client, rawNew, rawDesired *Interconnect) (*Interconnect, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LinkType) && dcl.IsEmptyValueIndirect(rawDesired.LinkType) {
		rawNew.LinkType = rawDesired.LinkType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RequestedLinkCount) && dcl.IsEmptyValueIndirect(rawDesired.RequestedLinkCount) {
		rawNew.RequestedLinkCount = rawDesired.RequestedLinkCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.InterconnectType) && dcl.IsEmptyValueIndirect(rawDesired.InterconnectType) {
		rawNew.InterconnectType = rawDesired.InterconnectType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdminEnabled) && dcl.IsEmptyValueIndirect(rawDesired.AdminEnabled) {
		rawNew.AdminEnabled = rawDesired.AdminEnabled
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.NocContactEmail) && dcl.IsEmptyValueIndirect(rawDesired.NocContactEmail) {
		rawNew.NocContactEmail = rawDesired.NocContactEmail
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomerName) && dcl.IsEmptyValueIndirect(rawDesired.CustomerName) {
		rawNew.CustomerName = rawDesired.CustomerName
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.OperationalStatus) && dcl.IsEmptyValueIndirect(rawDesired.OperationalStatus) {
		rawNew.OperationalStatus = rawDesired.OperationalStatus
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProvisionedLinkCount) && dcl.IsEmptyValueIndirect(rawDesired.ProvisionedLinkCount) {
		rawNew.ProvisionedLinkCount = rawDesired.ProvisionedLinkCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.InterconnectAttachments) && dcl.IsEmptyValueIndirect(rawDesired.InterconnectAttachments) {
		rawNew.InterconnectAttachments = rawDesired.InterconnectAttachments
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerIPAddress) && dcl.IsEmptyValueIndirect(rawDesired.PeerIPAddress) {
		rawNew.PeerIPAddress = rawDesired.PeerIPAddress
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GoogleIPAddress) && dcl.IsEmptyValueIndirect(rawDesired.GoogleIPAddress) {
		rawNew.GoogleIPAddress = rawDesired.GoogleIPAddress
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GoogleReferenceId) && dcl.IsEmptyValueIndirect(rawDesired.GoogleReferenceId) {
		rawNew.GoogleReferenceId = rawDesired.GoogleReferenceId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpectedOutages) && dcl.IsEmptyValueIndirect(rawDesired.ExpectedOutages) {
		rawNew.ExpectedOutages = rawDesired.ExpectedOutages
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CircuitInfos) && dcl.IsEmptyValueIndirect(rawDesired.CircuitInfos) {
		rawNew.CircuitInfos = rawDesired.CircuitInfos
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeInterconnectExpectedOutages(des, initial *InterconnectExpectedOutages, opts ...dcl.ApplyOption) *InterconnectExpectedOutages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Interconnect)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.Source) {
		des.Source = initial.Source
	}
	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.IssueType) {
		des.IssueType = initial.IssueType
	}
	if dcl.IsZeroValue(des.AffectedCircuits) {
		des.AffectedCircuits = initial.AffectedCircuits
	}
	if dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.IsZeroValue(des.EndTime) {
		des.EndTime = initial.EndTime
	}

	return des
}

func canonicalizeNewInterconnectExpectedOutages(c *Client, des, nw *InterconnectExpectedOutages) *InterconnectExpectedOutages {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInterconnectExpectedOutagesSet(c *Client, des, nw []InterconnectExpectedOutages) []InterconnectExpectedOutages {
	if des == nil {
		return nw
	}
	var reorderedNew []InterconnectExpectedOutages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInterconnectExpectedOutages(c, &d, &n) {
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

func canonicalizeInterconnectCircuitInfos(des, initial *InterconnectCircuitInfos, opts ...dcl.ApplyOption) *InterconnectCircuitInfos {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Interconnect)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.GoogleCircuitId) {
		des.GoogleCircuitId = initial.GoogleCircuitId
	}
	if dcl.IsZeroValue(des.GoogleDemarcId) {
		des.GoogleDemarcId = initial.GoogleDemarcId
	}
	if dcl.IsZeroValue(des.CustomerDemarcId) {
		des.CustomerDemarcId = initial.CustomerDemarcId
	}

	return des
}

func canonicalizeNewInterconnectCircuitInfos(c *Client, des, nw *InterconnectCircuitInfos) *InterconnectCircuitInfos {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInterconnectCircuitInfosSet(c *Client, des, nw []InterconnectCircuitInfos) []InterconnectCircuitInfos {
	if des == nil {
		return nw
	}
	var reorderedNew []InterconnectCircuitInfos
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInterconnectCircuitInfos(c, &d, &n) {
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

type interconnectDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         interconnectApiOperation
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
func diffInterconnect(c *Client, desired, actual *Interconnect, opts ...dcl.ApplyOption) ([]interconnectDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []interconnectDiff
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.Location) && (dcl.IsZeroValue(actual.Location) || !reflect.DeepEqual(*desired.Location, *actual.Location)) {
		c.Config.Logger.Infof("Detected diff in Location.\nDESIRED: %v\nACTUAL: %v", desired.Location, actual.Location)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "Location",
		})

	}
	if !dcl.IsZeroValue(desired.LinkType) && (dcl.IsZeroValue(actual.LinkType) || !reflect.DeepEqual(*desired.LinkType, *actual.LinkType)) {
		c.Config.Logger.Infof("Detected diff in LinkType.\nDESIRED: %v\nACTUAL: %v", desired.LinkType, actual.LinkType)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "LinkType",
		})

	}
	if !dcl.IsZeroValue(desired.RequestedLinkCount) && (dcl.IsZeroValue(actual.RequestedLinkCount) || !reflect.DeepEqual(*desired.RequestedLinkCount, *actual.RequestedLinkCount)) {
		c.Config.Logger.Infof("Detected diff in RequestedLinkCount.\nDESIRED: %v\nACTUAL: %v", desired.RequestedLinkCount, actual.RequestedLinkCount)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "RequestedLinkCount",
		})

	}
	if !dcl.IsZeroValue(desired.InterconnectType) && (dcl.IsZeroValue(actual.InterconnectType) || !reflect.DeepEqual(*desired.InterconnectType, *actual.InterconnectType)) {
		c.Config.Logger.Infof("Detected diff in InterconnectType.\nDESIRED: %v\nACTUAL: %v", desired.InterconnectType, actual.InterconnectType)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "InterconnectType",
		})

	}
	if !dcl.IsZeroValue(desired.AdminEnabled) && (dcl.IsZeroValue(actual.AdminEnabled) || !reflect.DeepEqual(*desired.AdminEnabled, *actual.AdminEnabled)) {
		c.Config.Logger.Infof("Detected diff in AdminEnabled.\nDESIRED: %v\nACTUAL: %v", desired.AdminEnabled, actual.AdminEnabled)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "AdminEnabled",
		})

	}
	if !dcl.IsZeroValue(desired.NocContactEmail) && (dcl.IsZeroValue(actual.NocContactEmail) || !reflect.DeepEqual(*desired.NocContactEmail, *actual.NocContactEmail)) {
		c.Config.Logger.Infof("Detected diff in NocContactEmail.\nDESIRED: %v\nACTUAL: %v", desired.NocContactEmail, actual.NocContactEmail)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "NocContactEmail",
		})

	}
	if !dcl.IsZeroValue(desired.CustomerName) && (dcl.IsZeroValue(actual.CustomerName) || !reflect.DeepEqual(*desired.CustomerName, *actual.CustomerName)) {
		c.Config.Logger.Infof("Detected diff in CustomerName.\nDESIRED: %v\nACTUAL: %v", desired.CustomerName, actual.CustomerName)

		diffs = append(diffs, interconnectDiff{
			UpdateOp:  &updateInterconnectPatchOperation{},
			FieldName: "CustomerName",
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
	var deduped []interconnectDiff
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
func compareInterconnectExpectedOutagesSlice(c *Client, desired, actual []InterconnectExpectedOutages) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectExpectedOutages, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectExpectedOutages(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectExpectedOutages, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectExpectedOutages(c *Client, desired, actual *InterconnectExpectedOutages) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.Source == nil && desired.Source != nil && !dcl.IsEmptyValueIndirect(desired.Source) {
		c.Config.Logger.Infof("desired Source %s - but actually nil", dcl.SprintResource(desired.Source))
		return true
	}
	if !reflect.DeepEqual(desired.Source, actual.Source) && !dcl.IsZeroValue(desired.Source) && !(dcl.IsEmptyValueIndirect(desired.Source) && dcl.IsZeroValue(actual.Source)) {
		c.Config.Logger.Infof("Diff in Source. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Source), dcl.SprintResource(actual.Source))
		return true
	}
	if actual.State == nil && desired.State != nil && !dcl.IsEmptyValueIndirect(desired.State) {
		c.Config.Logger.Infof("desired State %s - but actually nil", dcl.SprintResource(desired.State))
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) && !(dcl.IsEmptyValueIndirect(desired.State) && dcl.IsZeroValue(actual.State)) {
		c.Config.Logger.Infof("Diff in State. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	if actual.IssueType == nil && desired.IssueType != nil && !dcl.IsEmptyValueIndirect(desired.IssueType) {
		c.Config.Logger.Infof("desired IssueType %s - but actually nil", dcl.SprintResource(desired.IssueType))
		return true
	}
	if !reflect.DeepEqual(desired.IssueType, actual.IssueType) && !dcl.IsZeroValue(desired.IssueType) && !(dcl.IsEmptyValueIndirect(desired.IssueType) && dcl.IsZeroValue(actual.IssueType)) {
		c.Config.Logger.Infof("Diff in IssueType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IssueType), dcl.SprintResource(actual.IssueType))
		return true
	}
	if actual.AffectedCircuits == nil && desired.AffectedCircuits != nil && !dcl.IsEmptyValueIndirect(desired.AffectedCircuits) {
		c.Config.Logger.Infof("desired AffectedCircuits %s - but actually nil", dcl.SprintResource(desired.AffectedCircuits))
		return true
	}
	if !dcl.SliceEquals(desired.AffectedCircuits, actual.AffectedCircuits) && !dcl.IsZeroValue(desired.AffectedCircuits) {
		c.Config.Logger.Infof("Diff in AffectedCircuits. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AffectedCircuits), dcl.SprintResource(actual.AffectedCircuits))
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if !reflect.DeepEqual(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) && !(dcl.IsEmptyValueIndirect(desired.StartTime) && dcl.IsZeroValue(actual.StartTime)) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.EndTime == nil && desired.EndTime != nil && !dcl.IsEmptyValueIndirect(desired.EndTime) {
		c.Config.Logger.Infof("desired EndTime %s - but actually nil", dcl.SprintResource(desired.EndTime))
		return true
	}
	if !reflect.DeepEqual(desired.EndTime, actual.EndTime) && !dcl.IsZeroValue(desired.EndTime) && !(dcl.IsEmptyValueIndirect(desired.EndTime) && dcl.IsZeroValue(actual.EndTime)) {
		c.Config.Logger.Infof("Diff in EndTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndTime), dcl.SprintResource(actual.EndTime))
		return true
	}
	return false
}
func compareInterconnectCircuitInfosSlice(c *Client, desired, actual []InterconnectCircuitInfos) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectCircuitInfos, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectCircuitInfos(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectCircuitInfos, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectCircuitInfos(c *Client, desired, actual *InterconnectCircuitInfos) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.GoogleCircuitId == nil && desired.GoogleCircuitId != nil && !dcl.IsEmptyValueIndirect(desired.GoogleCircuitId) {
		c.Config.Logger.Infof("desired GoogleCircuitId %s - but actually nil", dcl.SprintResource(desired.GoogleCircuitId))
		return true
	}
	if !reflect.DeepEqual(desired.GoogleCircuitId, actual.GoogleCircuitId) && !dcl.IsZeroValue(desired.GoogleCircuitId) && !(dcl.IsEmptyValueIndirect(desired.GoogleCircuitId) && dcl.IsZeroValue(actual.GoogleCircuitId)) {
		c.Config.Logger.Infof("Diff in GoogleCircuitId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GoogleCircuitId), dcl.SprintResource(actual.GoogleCircuitId))
		return true
	}
	if actual.GoogleDemarcId == nil && desired.GoogleDemarcId != nil && !dcl.IsEmptyValueIndirect(desired.GoogleDemarcId) {
		c.Config.Logger.Infof("desired GoogleDemarcId %s - but actually nil", dcl.SprintResource(desired.GoogleDemarcId))
		return true
	}
	if !reflect.DeepEqual(desired.GoogleDemarcId, actual.GoogleDemarcId) && !dcl.IsZeroValue(desired.GoogleDemarcId) && !(dcl.IsEmptyValueIndirect(desired.GoogleDemarcId) && dcl.IsZeroValue(actual.GoogleDemarcId)) {
		c.Config.Logger.Infof("Diff in GoogleDemarcId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GoogleDemarcId), dcl.SprintResource(actual.GoogleDemarcId))
		return true
	}
	if actual.CustomerDemarcId == nil && desired.CustomerDemarcId != nil && !dcl.IsEmptyValueIndirect(desired.CustomerDemarcId) {
		c.Config.Logger.Infof("desired CustomerDemarcId %s - but actually nil", dcl.SprintResource(desired.CustomerDemarcId))
		return true
	}
	if !reflect.DeepEqual(desired.CustomerDemarcId, actual.CustomerDemarcId) && !dcl.IsZeroValue(desired.CustomerDemarcId) && !(dcl.IsEmptyValueIndirect(desired.CustomerDemarcId) && dcl.IsZeroValue(actual.CustomerDemarcId)) {
		c.Config.Logger.Infof("Diff in CustomerDemarcId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CustomerDemarcId), dcl.SprintResource(actual.CustomerDemarcId))
		return true
	}
	return false
}
func compareInterconnectLinkTypeEnumSlice(c *Client, desired, actual []InterconnectLinkTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectLinkTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectLinkTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectLinkTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectLinkTypeEnum(c *Client, desired, actual *InterconnectLinkTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectInterconnectTypeEnumSlice(c *Client, desired, actual []InterconnectInterconnectTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectInterconnectTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectInterconnectTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectInterconnectTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectInterconnectTypeEnum(c *Client, desired, actual *InterconnectInterconnectTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectOperationalStatusEnumSlice(c *Client, desired, actual []InterconnectOperationalStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectOperationalStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectOperationalStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectOperationalStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectOperationalStatusEnum(c *Client, desired, actual *InterconnectOperationalStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectExpectedOutagesSourceEnumSlice(c *Client, desired, actual []InterconnectExpectedOutagesSourceEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectExpectedOutagesSourceEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectExpectedOutagesSourceEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectExpectedOutagesSourceEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectExpectedOutagesSourceEnum(c *Client, desired, actual *InterconnectExpectedOutagesSourceEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectExpectedOutagesStateEnumSlice(c *Client, desired, actual []InterconnectExpectedOutagesStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectExpectedOutagesStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectExpectedOutagesStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectExpectedOutagesStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectExpectedOutagesStateEnum(c *Client, desired, actual *InterconnectExpectedOutagesStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectExpectedOutagesIssueTypeEnumSlice(c *Client, desired, actual []InterconnectExpectedOutagesIssueTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectExpectedOutagesIssueTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectExpectedOutagesIssueTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectExpectedOutagesIssueTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectExpectedOutagesIssueTypeEnum(c *Client, desired, actual *InterconnectExpectedOutagesIssueTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInterconnectStateEnumSlice(c *Client, desired, actual []InterconnectStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InterconnectStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInterconnectStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InterconnectStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInterconnectStateEnum(c *Client, desired, actual *InterconnectStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Interconnect) urlNormalized() *Interconnect {
	normalized := deepcopy.Copy(*r).(Interconnect)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Interconnect) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Interconnect) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Interconnect) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Interconnect) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Patch" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/interconnects/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Interconnect resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Interconnect) marshal(c *Client) ([]byte, error) {
	m, err := expandInterconnect(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Interconnect: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInterconnect decodes JSON responses into the Interconnect resource schema.
func unmarshalInterconnect(b []byte, c *Client) (*Interconnect, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenInterconnect(c, m), nil
}

// expandInterconnect expands Interconnect into a JSON request object.
func expandInterconnect(c *Client, f *Interconnect) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.LinkType; !dcl.IsEmptyValueIndirect(v) {
		m["linkType"] = v
	}
	if v := f.RequestedLinkCount; !dcl.IsEmptyValueIndirect(v) {
		m["requestedLinkCount"] = v
	}
	if v := f.InterconnectType; !dcl.IsEmptyValueIndirect(v) {
		m["interconnectType"] = v
	}
	if v := f.AdminEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["adminEnabled"] = v
	}
	if v := f.NocContactEmail; !dcl.IsEmptyValueIndirect(v) {
		m["nocContactEmail"] = v
	}
	if v := f.CustomerName; !dcl.IsEmptyValueIndirect(v) {
		m["customerName"] = v
	}
	if v := f.OperationalStatus; !dcl.IsEmptyValueIndirect(v) {
		m["operationalStatus"] = v
	}
	if v := f.ProvisionedLinkCount; !dcl.IsEmptyValueIndirect(v) {
		m["provisionedLinkCount"] = v
	}
	if v := f.InterconnectAttachments; !dcl.IsEmptyValueIndirect(v) {
		m["interconnectAttachments"] = v
	}
	if v := f.PeerIPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["peerIpAddress"] = v
	}
	if v := f.GoogleIPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["googleIpAddress"] = v
	}
	if v := f.GoogleReferenceId; !dcl.IsEmptyValueIndirect(v) {
		m["googleReferenceId"] = v
	}
	if v, err := expandInterconnectExpectedOutagesSlice(c, f.ExpectedOutages); err != nil {
		return nil, fmt.Errorf("error expanding ExpectedOutages into expectedOutages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["expectedOutages"] = v
	}
	if v, err := expandInterconnectCircuitInfosSlice(c, f.CircuitInfos); err != nil {
		return nil, fmt.Errorf("error expanding CircuitInfos into circuitInfos: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["circuitInfos"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenInterconnect flattens Interconnect from a JSON request object into the
// Interconnect type.
func flattenInterconnect(c *Client, i interface{}) *Interconnect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Interconnect{}
	r.Description = dcl.FlattenString(m["description"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Location = dcl.FlattenString(m["location"])
	r.LinkType = flattenInterconnectLinkTypeEnum(m["linkType"])
	r.RequestedLinkCount = dcl.FlattenInteger(m["requestedLinkCount"])
	r.InterconnectType = flattenInterconnectInterconnectTypeEnum(m["interconnectType"])
	r.AdminEnabled = dcl.FlattenBool(m["adminEnabled"])
	r.NocContactEmail = dcl.FlattenString(m["nocContactEmail"])
	r.CustomerName = dcl.FlattenString(m["customerName"])
	r.OperationalStatus = flattenInterconnectOperationalStatusEnum(m["operationalStatus"])
	r.ProvisionedLinkCount = dcl.FlattenInteger(m["provisionedLinkCount"])
	r.InterconnectAttachments = dcl.FlattenStringSlice(m["interconnectAttachments"])
	r.PeerIPAddress = dcl.FlattenString(m["peerIpAddress"])
	r.GoogleIPAddress = dcl.FlattenString(m["googleIpAddress"])
	r.GoogleReferenceId = dcl.FlattenString(m["googleReferenceId"])
	r.ExpectedOutages = flattenInterconnectExpectedOutagesSlice(c, m["expectedOutages"])
	r.CircuitInfos = flattenInterconnectCircuitInfosSlice(c, m["circuitInfos"])
	r.State = flattenInterconnectStateEnum(m["state"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandInterconnectExpectedOutagesMap expands the contents of InterconnectExpectedOutages into a JSON
// request object.
func expandInterconnectExpectedOutagesMap(c *Client, f map[string]InterconnectExpectedOutages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInterconnectExpectedOutages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInterconnectExpectedOutagesSlice expands the contents of InterconnectExpectedOutages into a JSON
// request object.
func expandInterconnectExpectedOutagesSlice(c *Client, f []InterconnectExpectedOutages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInterconnectExpectedOutages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInterconnectExpectedOutagesMap flattens the contents of InterconnectExpectedOutages from a JSON
// response object.
func flattenInterconnectExpectedOutagesMap(c *Client, i interface{}) map[string]InterconnectExpectedOutages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InterconnectExpectedOutages{}
	}

	if len(a) == 0 {
		return map[string]InterconnectExpectedOutages{}
	}

	items := make(map[string]InterconnectExpectedOutages)
	for k, item := range a {
		items[k] = *flattenInterconnectExpectedOutages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInterconnectExpectedOutagesSlice flattens the contents of InterconnectExpectedOutages from a JSON
// response object.
func flattenInterconnectExpectedOutagesSlice(c *Client, i interface{}) []InterconnectExpectedOutages {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectExpectedOutages{}
	}

	if len(a) == 0 {
		return []InterconnectExpectedOutages{}
	}

	items := make([]InterconnectExpectedOutages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectExpectedOutages(c, item.(map[string]interface{})))
	}

	return items
}

// expandInterconnectExpectedOutages expands an instance of InterconnectExpectedOutages into a JSON
// request object.
func expandInterconnectExpectedOutages(c *Client, f *InterconnectExpectedOutages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Source; !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.IssueType; !dcl.IsEmptyValueIndirect(v) {
		m["issueType"] = v
	}
	if v := f.AffectedCircuits; !dcl.IsEmptyValueIndirect(v) {
		m["affectedCircuits"] = v
	}
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.EndTime; !dcl.IsEmptyValueIndirect(v) {
		m["endTime"] = v
	}

	return m, nil
}

// flattenInterconnectExpectedOutages flattens an instance of InterconnectExpectedOutages from a JSON
// response object.
func flattenInterconnectExpectedOutages(c *Client, i interface{}) *InterconnectExpectedOutages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InterconnectExpectedOutages{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Source = flattenInterconnectExpectedOutagesSourceEnum(m["source"])
	r.State = flattenInterconnectExpectedOutagesStateEnum(m["state"])
	r.IssueType = flattenInterconnectExpectedOutagesIssueTypeEnum(m["issueType"])
	r.AffectedCircuits = dcl.FlattenStringSlice(m["affectedCircuits"])
	r.StartTime = dcl.FlattenInteger(m["startTime"])
	r.EndTime = dcl.FlattenInteger(m["endTime"])

	return r
}

// expandInterconnectCircuitInfosMap expands the contents of InterconnectCircuitInfos into a JSON
// request object.
func expandInterconnectCircuitInfosMap(c *Client, f map[string]InterconnectCircuitInfos) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInterconnectCircuitInfos(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInterconnectCircuitInfosSlice expands the contents of InterconnectCircuitInfos into a JSON
// request object.
func expandInterconnectCircuitInfosSlice(c *Client, f []InterconnectCircuitInfos) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInterconnectCircuitInfos(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInterconnectCircuitInfosMap flattens the contents of InterconnectCircuitInfos from a JSON
// response object.
func flattenInterconnectCircuitInfosMap(c *Client, i interface{}) map[string]InterconnectCircuitInfos {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InterconnectCircuitInfos{}
	}

	if len(a) == 0 {
		return map[string]InterconnectCircuitInfos{}
	}

	items := make(map[string]InterconnectCircuitInfos)
	for k, item := range a {
		items[k] = *flattenInterconnectCircuitInfos(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInterconnectCircuitInfosSlice flattens the contents of InterconnectCircuitInfos from a JSON
// response object.
func flattenInterconnectCircuitInfosSlice(c *Client, i interface{}) []InterconnectCircuitInfos {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectCircuitInfos{}
	}

	if len(a) == 0 {
		return []InterconnectCircuitInfos{}
	}

	items := make([]InterconnectCircuitInfos, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectCircuitInfos(c, item.(map[string]interface{})))
	}

	return items
}

// expandInterconnectCircuitInfos expands an instance of InterconnectCircuitInfos into a JSON
// request object.
func expandInterconnectCircuitInfos(c *Client, f *InterconnectCircuitInfos) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.GoogleCircuitId; !dcl.IsEmptyValueIndirect(v) {
		m["googleCircuitId"] = v
	}
	if v := f.GoogleDemarcId; !dcl.IsEmptyValueIndirect(v) {
		m["googleDemarcId"] = v
	}
	if v := f.CustomerDemarcId; !dcl.IsEmptyValueIndirect(v) {
		m["customerDemarcId"] = v
	}

	return m, nil
}

// flattenInterconnectCircuitInfos flattens an instance of InterconnectCircuitInfos from a JSON
// response object.
func flattenInterconnectCircuitInfos(c *Client, i interface{}) *InterconnectCircuitInfos {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InterconnectCircuitInfos{}
	r.GoogleCircuitId = dcl.FlattenString(m["googleCircuitId"])
	r.GoogleDemarcId = dcl.FlattenString(m["googleDemarcId"])
	r.CustomerDemarcId = dcl.FlattenString(m["customerDemarcId"])

	return r
}

// flattenInterconnectLinkTypeEnumSlice flattens the contents of InterconnectLinkTypeEnum from a JSON
// response object.
func flattenInterconnectLinkTypeEnumSlice(c *Client, i interface{}) []InterconnectLinkTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectLinkTypeEnum{}
	}

	if len(a) == 0 {
		return []InterconnectLinkTypeEnum{}
	}

	items := make([]InterconnectLinkTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectLinkTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectLinkTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectLinkTypeEnum with the same value as that string.
func flattenInterconnectLinkTypeEnum(i interface{}) *InterconnectLinkTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectLinkTypeEnumRef("")
	}

	return InterconnectLinkTypeEnumRef(s)
}

// flattenInterconnectInterconnectTypeEnumSlice flattens the contents of InterconnectInterconnectTypeEnum from a JSON
// response object.
func flattenInterconnectInterconnectTypeEnumSlice(c *Client, i interface{}) []InterconnectInterconnectTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectInterconnectTypeEnum{}
	}

	if len(a) == 0 {
		return []InterconnectInterconnectTypeEnum{}
	}

	items := make([]InterconnectInterconnectTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectInterconnectTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectInterconnectTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectInterconnectTypeEnum with the same value as that string.
func flattenInterconnectInterconnectTypeEnum(i interface{}) *InterconnectInterconnectTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectInterconnectTypeEnumRef("")
	}

	return InterconnectInterconnectTypeEnumRef(s)
}

// flattenInterconnectOperationalStatusEnumSlice flattens the contents of InterconnectOperationalStatusEnum from a JSON
// response object.
func flattenInterconnectOperationalStatusEnumSlice(c *Client, i interface{}) []InterconnectOperationalStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectOperationalStatusEnum{}
	}

	if len(a) == 0 {
		return []InterconnectOperationalStatusEnum{}
	}

	items := make([]InterconnectOperationalStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectOperationalStatusEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectOperationalStatusEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectOperationalStatusEnum with the same value as that string.
func flattenInterconnectOperationalStatusEnum(i interface{}) *InterconnectOperationalStatusEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectOperationalStatusEnumRef("")
	}

	return InterconnectOperationalStatusEnumRef(s)
}

// flattenInterconnectExpectedOutagesSourceEnumSlice flattens the contents of InterconnectExpectedOutagesSourceEnum from a JSON
// response object.
func flattenInterconnectExpectedOutagesSourceEnumSlice(c *Client, i interface{}) []InterconnectExpectedOutagesSourceEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectExpectedOutagesSourceEnum{}
	}

	if len(a) == 0 {
		return []InterconnectExpectedOutagesSourceEnum{}
	}

	items := make([]InterconnectExpectedOutagesSourceEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectExpectedOutagesSourceEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectExpectedOutagesSourceEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectExpectedOutagesSourceEnum with the same value as that string.
func flattenInterconnectExpectedOutagesSourceEnum(i interface{}) *InterconnectExpectedOutagesSourceEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectExpectedOutagesSourceEnumRef("")
	}

	return InterconnectExpectedOutagesSourceEnumRef(s)
}

// flattenInterconnectExpectedOutagesStateEnumSlice flattens the contents of InterconnectExpectedOutagesStateEnum from a JSON
// response object.
func flattenInterconnectExpectedOutagesStateEnumSlice(c *Client, i interface{}) []InterconnectExpectedOutagesStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectExpectedOutagesStateEnum{}
	}

	if len(a) == 0 {
		return []InterconnectExpectedOutagesStateEnum{}
	}

	items := make([]InterconnectExpectedOutagesStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectExpectedOutagesStateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectExpectedOutagesStateEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectExpectedOutagesStateEnum with the same value as that string.
func flattenInterconnectExpectedOutagesStateEnum(i interface{}) *InterconnectExpectedOutagesStateEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectExpectedOutagesStateEnumRef("")
	}

	return InterconnectExpectedOutagesStateEnumRef(s)
}

// flattenInterconnectExpectedOutagesIssueTypeEnumSlice flattens the contents of InterconnectExpectedOutagesIssueTypeEnum from a JSON
// response object.
func flattenInterconnectExpectedOutagesIssueTypeEnumSlice(c *Client, i interface{}) []InterconnectExpectedOutagesIssueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectExpectedOutagesIssueTypeEnum{}
	}

	if len(a) == 0 {
		return []InterconnectExpectedOutagesIssueTypeEnum{}
	}

	items := make([]InterconnectExpectedOutagesIssueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectExpectedOutagesIssueTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectExpectedOutagesIssueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectExpectedOutagesIssueTypeEnum with the same value as that string.
func flattenInterconnectExpectedOutagesIssueTypeEnum(i interface{}) *InterconnectExpectedOutagesIssueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectExpectedOutagesIssueTypeEnumRef("")
	}

	return InterconnectExpectedOutagesIssueTypeEnumRef(s)
}

// flattenInterconnectStateEnumSlice flattens the contents of InterconnectStateEnum from a JSON
// response object.
func flattenInterconnectStateEnumSlice(c *Client, i interface{}) []InterconnectStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InterconnectStateEnum{}
	}

	if len(a) == 0 {
		return []InterconnectStateEnum{}
	}

	items := make([]InterconnectStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInterconnectStateEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInterconnectStateEnum asserts that an interface is a string, and returns a
// pointer to a *InterconnectStateEnum with the same value as that string.
func flattenInterconnectStateEnum(i interface{}) *InterconnectStateEnum {
	s, ok := i.(string)
	if !ok {
		return InterconnectStateEnumRef("")
	}

	return InterconnectStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Interconnect) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInterconnect(b, c)
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
