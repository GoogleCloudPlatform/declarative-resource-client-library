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
package eventarc

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

func (r *Trigger) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "eventFilters"); err != nil {
		return err
	}
	if err := dcl.Required(r, "destination"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Destination) {
		if err := r.Destination.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Transport) {
		if err := r.Transport.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TriggerEventFilters) validate() error {
	if err := dcl.Required(r, "attribute"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	return nil
}
func (r *TriggerDestination) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CloudRun) {
		if err := r.CloudRun.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Gke) {
		if err := r.Gke.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TriggerDestinationCloudRun) validate() error {
	if err := dcl.Required(r, "service"); err != nil {
		return err
	}
	if err := dcl.Required(r, "region"); err != nil {
		return err
	}
	return nil
}
func (r *TriggerDestinationGke) validate() error {
	if err := dcl.Required(r, "cluster"); err != nil {
		return err
	}
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	if err := dcl.Required(r, "namespace"); err != nil {
		return err
	}
	if err := dcl.Required(r, "service"); err != nil {
		return err
	}
	return nil
}
func (r *TriggerTransport) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Pubsub) {
		if err := r.Pubsub.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TriggerTransportPubsub) validate() error {
	return nil
}

func triggerGetURL(userBasePath string, r *Trigger) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1/", userBasePath, params), nil
}

func triggerListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers", "https://eventarc.googleapis.com/v1/", userBasePath, params), nil

}

func triggerCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers?triggerId={{name}}", "https://eventarc.googleapis.com/v1/", userBasePath, params), nil

}

func triggerDeleteURL(userBasePath string, r *Trigger) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1/", userBasePath, params), nil
}

// triggerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type triggerApiOperation interface {
	do(context.Context, *Trigger, *Client) error
}

// newUpdateTriggerUpdateTriggerRequest creates a request for an
// Trigger resource's UpdateTrigger update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTriggerUpdateTriggerRequest(ctx context.Context, f *Trigger, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/%s/triggers/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandTriggerEventFiltersSlice(c, f.EventFilters); err != nil {
		return nil, fmt.Errorf("error expanding EventFilters into eventFilters: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["eventFilters"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		req["serviceAccount"] = v
	}
	if v, err := expandTriggerDestination(c, f.Destination); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["destination"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getTriggerRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/triggers/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateTriggerUpdateTriggerRequest converts the update into
// the final JSON request body.
func marshalUpdateTriggerUpdateTriggerRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTriggerUpdateTriggerOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTriggerUpdateTriggerOperation) do(ctx context.Context, r *Trigger, c *Client) error {
	_, err := c.GetTrigger(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTrigger")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"name", "eventFilters", "serviceAccount", "destination", "labels", "etag"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTriggerUpdateTriggerRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTriggerUpdateTriggerRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTriggerRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := triggerListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TriggerMaxPage {
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

type listTriggerOperation struct {
	Triggers []map[string]interface{} `json:"triggers"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listTrigger(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Trigger, string, error) {
	b, err := c.listTriggerRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTriggerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Trigger
	for _, v := range m.Triggers {
		res := flattenTrigger(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTrigger(ctx context.Context, f func(*Trigger) bool, resources []*Trigger) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTrigger(ctx, res)
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

type deleteTriggerOperation struct{}

func (op *deleteTriggerOperation) do(ctx context.Context, r *Trigger, c *Client) error {

	_, err := c.GetTrigger(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Trigger not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTrigger checking for existence. error: %v", err)
		return err
	}

	u, err := triggerDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetTrigger(ctx, r.urlNormalized())
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
type createTriggerOperation struct {
	response map[string]interface{}
}

func (op *createTriggerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTriggerOperation) do(ctx context.Context, r *Trigger, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := triggerCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetTrigger(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTriggerRaw(ctx context.Context, r *Trigger) ([]byte, error) {

	u, err := triggerGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) triggerDiffsForRawDesired(ctx context.Context, rawDesired *Trigger, opts ...dcl.ApplyOption) (initial, desired *Trigger, diffs []triggerDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Trigger
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Trigger); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Trigger, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTrigger(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Trigger resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Trigger resource: %v", err)
		}
		c.Config.Logger.Info("Found that Trigger resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTriggerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Trigger: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Trigger: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTriggerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Trigger: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTriggerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Trigger: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTrigger(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTriggerInitialState(rawInitial, rawDesired *Trigger) (*Trigger, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTriggerDesiredState(rawDesired, rawInitial *Trigger, opts ...dcl.ApplyOption) (*Trigger, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Destination = canonicalizeTriggerDestination(rawDesired.Destination, nil, opts...)
		rawDesired.Transport = canonicalizeTriggerTransport(rawDesired.Transport, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.EventFilters) {
		rawDesired.EventFilters = rawInitial.EventFilters
	}
	if dcl.NameToSelfLink(rawDesired.ServiceAccount, rawInitial.ServiceAccount) {
		rawDesired.ServiceAccount = rawInitial.ServiceAccount
	}
	rawDesired.Destination = canonicalizeTriggerDestination(rawDesired.Destination, rawInitial.Destination, opts...)
	rawDesired.Transport = canonicalizeTriggerTransport(rawDesired.Transport, rawInitial.Transport, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeTriggerNewState(c *Client, rawNew, rawDesired *Trigger) (*Trigger, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EventFilters) && dcl.IsEmptyValueIndirect(rawDesired.EventFilters) {
		rawNew.EventFilters = rawDesired.EventFilters
	} else {
		rawNew.EventFilters = canonicalizeNewTriggerEventFiltersSet(c, rawDesired.EventFilters, rawNew.EventFilters)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccount) {
		rawNew.ServiceAccount = rawDesired.ServiceAccount
	} else {
		if dcl.NameToSelfLink(rawDesired.ServiceAccount, rawNew.ServiceAccount) {
			rawNew.ServiceAccount = rawDesired.ServiceAccount
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Destination) && dcl.IsEmptyValueIndirect(rawDesired.Destination) {
		rawNew.Destination = rawDesired.Destination
	} else {
		rawNew.Destination = canonicalizeNewTriggerDestination(c, rawDesired.Destination, rawNew.Destination)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Transport) && dcl.IsEmptyValueIndirect(rawDesired.Transport) {
		rawNew.Transport = rawDesired.Transport
	} else {
		rawNew.Transport = canonicalizeNewTriggerTransport(c, rawDesired.Transport, rawNew.Transport)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeTriggerEventFilters(des, initial *TriggerEventFilters, opts ...dcl.ApplyOption) *TriggerEventFilters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Attribute, initial.Attribute) || dcl.IsZeroValue(des.Attribute) {
		des.Attribute = initial.Attribute
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewTriggerEventFilters(c *Client, des, nw *TriggerEventFilters) *TriggerEventFilters {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Attribute, nw.Attribute) {
		nw.Attribute = des.Attribute
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewTriggerEventFiltersSet(c *Client, des, nw []TriggerEventFilters) []TriggerEventFilters {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerEventFilters
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerEventFiltersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerEventFiltersSlice(c *Client, des, nw []TriggerEventFilters) []TriggerEventFilters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerEventFilters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerEventFilters(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerDestination(des, initial *TriggerDestination, opts ...dcl.ApplyOption) *TriggerDestination {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.CloudRun = canonicalizeTriggerDestinationCloudRun(des.CloudRun, initial.CloudRun, opts...)
	if dcl.NameToSelfLink(des.CloudFunction, initial.CloudFunction) || dcl.IsZeroValue(des.CloudFunction) {
		des.CloudFunction = initial.CloudFunction
	}
	des.Gke = canonicalizeTriggerDestinationGke(des.Gke, initial.Gke, opts...)

	return des
}

func canonicalizeNewTriggerDestination(c *Client, des, nw *TriggerDestination) *TriggerDestination {
	if des == nil || nw == nil {
		return nw
	}

	nw.CloudRun = canonicalizeNewTriggerDestinationCloudRun(c, des.CloudRun, nw.CloudRun)
	if dcl.NameToSelfLink(des.CloudFunction, nw.CloudFunction) {
		nw.CloudFunction = des.CloudFunction
	}
	nw.Gke = canonicalizeNewTriggerDestinationGke(c, des.Gke, nw.Gke)

	return nw
}

func canonicalizeNewTriggerDestinationSet(c *Client, des, nw []TriggerDestination) []TriggerDestination {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerDestination
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerDestinationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerDestinationSlice(c *Client, des, nw []TriggerDestination) []TriggerDestination {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerDestination
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerDestination(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerDestinationCloudRun(des, initial *TriggerDestinationCloudRun, opts ...dcl.ApplyOption) *TriggerDestinationCloudRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.StringCanonicalize(des.Region, initial.Region) || dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewTriggerDestinationCloudRun(c *Client, des, nw *TriggerDestinationCloudRun) *TriggerDestinationCloudRun {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Region, nw.Region) {
		nw.Region = des.Region
	}

	return nw
}

func canonicalizeNewTriggerDestinationCloudRunSet(c *Client, des, nw []TriggerDestinationCloudRun) []TriggerDestinationCloudRun {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerDestinationCloudRun
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerDestinationCloudRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerDestinationCloudRunSlice(c *Client, des, nw []TriggerDestinationCloudRun) []TriggerDestinationCloudRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerDestinationCloudRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerDestinationCloudRun(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerDestinationGke(des, initial *TriggerDestinationGke, opts ...dcl.ApplyOption) *TriggerDestinationGke {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Cluster, initial.Cluster) || dcl.IsZeroValue(des.Cluster) {
		des.Cluster = initial.Cluster
	}
	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		des.Location = initial.Location
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		des.Namespace = initial.Namespace
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}

	return des
}

func canonicalizeNewTriggerDestinationGke(c *Client, des, nw *TriggerDestinationGke) *TriggerDestinationGke {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Cluster, nw.Cluster) {
		nw.Cluster = des.Cluster
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}
	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewTriggerDestinationGkeSet(c *Client, des, nw []TriggerDestinationGke) []TriggerDestinationGke {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerDestinationGke
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerDestinationGkeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerDestinationGkeSlice(c *Client, des, nw []TriggerDestinationGke) []TriggerDestinationGke {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerDestinationGke
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerDestinationGke(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerTransport(des, initial *TriggerTransport, opts ...dcl.ApplyOption) *TriggerTransport {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Pubsub = canonicalizeTriggerTransportPubsub(des.Pubsub, initial.Pubsub, opts...)

	return des
}

func canonicalizeNewTriggerTransport(c *Client, des, nw *TriggerTransport) *TriggerTransport {
	if des == nil || nw == nil {
		return nw
	}

	nw.Pubsub = canonicalizeNewTriggerTransportPubsub(c, des.Pubsub, nw.Pubsub)

	return nw
}

func canonicalizeNewTriggerTransportSet(c *Client, des, nw []TriggerTransport) []TriggerTransport {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerTransport
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerTransportNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerTransportSlice(c *Client, des, nw []TriggerTransport) []TriggerTransport {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerTransport
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerTransport(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerTransportPubsub(des, initial *TriggerTransportPubsub, opts ...dcl.ApplyOption) *TriggerTransportPubsub {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Topic, initial.Topic) || dcl.IsZeroValue(des.Topic) {
		des.Topic = initial.Topic
	}

	return des
}

func canonicalizeNewTriggerTransportPubsub(c *Client, des, nw *TriggerTransportPubsub) *TriggerTransportPubsub {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Topic, nw.Topic) {
		nw.Topic = des.Topic
	}
	if dcl.StringCanonicalize(des.Subscription, nw.Subscription) {
		nw.Subscription = des.Subscription
	}

	return nw
}

func canonicalizeNewTriggerTransportPubsubSet(c *Client, des, nw []TriggerTransportPubsub) []TriggerTransportPubsub {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerTransportPubsub
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTriggerTransportPubsubNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTriggerTransportPubsubSlice(c *Client, des, nw []TriggerTransportPubsub) []TriggerTransportPubsub {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TriggerTransportPubsub
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerTransportPubsub(c, &d, &n))
	}

	return items
}

type triggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         triggerApiOperation
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
func diffTrigger(c *Client, desired, actual *Trigger, opts ...dcl.ApplyOption) ([]triggerDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []triggerDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.EventFilters, actual.EventFilters, dcl.Info{Type: "Set", ObjectFunction: compareTriggerEventFiltersNewStyle, OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("EventFilters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{ObjectFunction: compareTriggerDestinationNewStyle, OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Transport, actual.Transport, dcl.Info{ObjectFunction: compareTriggerTransportNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Transport")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToTriggerDiff(ds, opts...)
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
	var deduped []triggerDiff
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
func compareTriggerEventFiltersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerEventFilters)
	if !ok {
		desiredNotPointer, ok := d.(TriggerEventFilters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerEventFilters or *TriggerEventFilters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerEventFilters)
	if !ok {
		actualNotPointer, ok := a.(TriggerEventFilters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerEventFilters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Attribute, actual.Attribute, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Attribute")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTriggerDestinationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerDestination)
	if !ok {
		desiredNotPointer, ok := d.(TriggerDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestination or *TriggerDestination", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerDestination)
	if !ok {
		actualNotPointer, ok := a.(TriggerDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestination", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CloudRun, actual.CloudRun, dcl.Info{ObjectFunction: compareTriggerDestinationCloudRunNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudRun")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudFunction, actual.CloudFunction, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudFunction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gke, actual.Gke, dcl.Info{ObjectFunction: compareTriggerDestinationGkeNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Gke")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTriggerDestinationCloudRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerDestinationCloudRun)
	if !ok {
		desiredNotPointer, ok := d.(TriggerDestinationCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestinationCloudRun or *TriggerDestinationCloudRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerDestinationCloudRun)
	if !ok {
		actualNotPointer, ok := a.(TriggerDestinationCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestinationCloudRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTriggerDestinationGkeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerDestinationGke)
	if !ok {
		desiredNotPointer, ok := d.(TriggerDestinationGke)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestinationGke or *TriggerDestinationGke", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerDestinationGke)
	if !ok {
		actualNotPointer, ok := a.(TriggerDestinationGke)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerDestinationGke", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Cluster, actual.Cluster, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Cluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTriggerTransportNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerTransport)
	if !ok {
		desiredNotPointer, ok := d.(TriggerTransport)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerTransport or *TriggerTransport", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerTransport)
	if !ok {
		actualNotPointer, ok := a.(TriggerTransport)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerTransport", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pubsub, actual.Pubsub, dcl.Info{ObjectFunction: compareTriggerTransportPubsubNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Pubsub")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTriggerTransportPubsubNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TriggerTransportPubsub)
	if !ok {
		desiredNotPointer, ok := d.(TriggerTransportPubsub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerTransportPubsub or *TriggerTransportPubsub", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TriggerTransportPubsub)
	if !ok {
		actualNotPointer, ok := a.(TriggerTransportPubsub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TriggerTransportPubsub", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subscription, actual.Subscription, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subscription")); len(ds) != 0 || err != nil {
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
func (r *Trigger) urlNormalized() *Trigger {
	normalized := dcl.Copy(*r).(Trigger)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.ServiceAccount = dcl.SelfLinkToName(r.ServiceAccount)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Trigger) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Trigger) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Trigger) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Trigger) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateTrigger" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Trigger resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Trigger) marshal(c *Client) ([]byte, error) {
	m, err := expandTrigger(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Trigger: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTrigger decodes JSON responses into the Trigger resource schema.
func unmarshalTrigger(b []byte, c *Client) (*Trigger, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTrigger(m, c)
}

func unmarshalMapTrigger(m map[string]interface{}, c *Client) (*Trigger, error) {

	return flattenTrigger(c, m), nil
}

// expandTrigger expands Trigger into a JSON request object.
func expandTrigger(c *Client, f *Trigger) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/triggers/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := expandTriggerEventFiltersSlice(c, f.EventFilters); err != nil {
		return nil, fmt.Errorf("error expanding EventFilters into eventFilters: %w", err)
	} else if v != nil {
		m["eventFilters"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandTriggerDestination(c, f.Destination); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if v != nil {
		m["destination"] = v
	}
	if v, err := expandTriggerTransport(c, f.Transport); err != nil {
		return nil, fmt.Errorf("error expanding Transport into transport: %w", err)
	} else if v != nil {
		m["transport"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
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

// flattenTrigger flattens Trigger from a JSON request object into the
// Trigger type.
func flattenTrigger(c *Client, i interface{}) *Trigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Trigger{}
	res.Name = dcl.FlattenString(m["name"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.EventFilters = flattenTriggerEventFiltersSlice(c, m["eventFilters"])
	res.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	res.Destination = flattenTriggerDestination(c, m["destination"])
	res.Transport = flattenTriggerTransport(c, m["transport"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandTriggerEventFiltersMap expands the contents of TriggerEventFilters into a JSON
// request object.
func expandTriggerEventFiltersMap(c *Client, f map[string]TriggerEventFilters) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerEventFilters(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerEventFiltersSlice expands the contents of TriggerEventFilters into a JSON
// request object.
func expandTriggerEventFiltersSlice(c *Client, f []TriggerEventFilters) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerEventFilters(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerEventFiltersMap flattens the contents of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFiltersMap(c *Client, i interface{}) map[string]TriggerEventFilters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerEventFilters{}
	}

	if len(a) == 0 {
		return map[string]TriggerEventFilters{}
	}

	items := make(map[string]TriggerEventFilters)
	for k, item := range a {
		items[k] = *flattenTriggerEventFilters(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerEventFiltersSlice flattens the contents of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFiltersSlice(c *Client, i interface{}) []TriggerEventFilters {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerEventFilters{}
	}

	if len(a) == 0 {
		return []TriggerEventFilters{}
	}

	items := make([]TriggerEventFilters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerEventFilters(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerEventFilters expands an instance of TriggerEventFilters into a JSON
// request object.
func expandTriggerEventFilters(c *Client, f *TriggerEventFilters) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Attribute; !dcl.IsEmptyValueIndirect(v) {
		m["attribute"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenTriggerEventFilters flattens an instance of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFilters(c *Client, i interface{}) *TriggerEventFilters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerEventFilters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerEventFilters
	}
	r.Attribute = dcl.FlattenString(m["attribute"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandTriggerDestinationMap expands the contents of TriggerDestination into a JSON
// request object.
func expandTriggerDestinationMap(c *Client, f map[string]TriggerDestination) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestination(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerDestinationSlice expands the contents of TriggerDestination into a JSON
// request object.
func expandTriggerDestinationSlice(c *Client, f []TriggerDestination) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestination(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationMap flattens the contents of TriggerDestination from a JSON
// response object.
func flattenTriggerDestinationMap(c *Client, i interface{}) map[string]TriggerDestination {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestination{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestination{}
	}

	items := make(map[string]TriggerDestination)
	for k, item := range a {
		items[k] = *flattenTriggerDestination(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerDestinationSlice flattens the contents of TriggerDestination from a JSON
// response object.
func flattenTriggerDestinationSlice(c *Client, i interface{}) []TriggerDestination {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestination{}
	}

	if len(a) == 0 {
		return []TriggerDestination{}
	}

	items := make([]TriggerDestination, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestination(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerDestination expands an instance of TriggerDestination into a JSON
// request object.
func expandTriggerDestination(c *Client, f *TriggerDestination) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTriggerDestinationCloudRun(c, f.CloudRun); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRun"] = v
	}
	if v := f.CloudFunction; !dcl.IsEmptyValueIndirect(v) {
		m["cloudFunction"] = v
	}
	if v, err := expandTriggerDestinationGke(c, f.Gke); err != nil {
		return nil, fmt.Errorf("error expanding Gke into gke: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gke"] = v
	}

	return m, nil
}

// flattenTriggerDestination flattens an instance of TriggerDestination from a JSON
// response object.
func flattenTriggerDestination(c *Client, i interface{}) *TriggerDestination {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerDestination{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerDestination
	}
	r.CloudRun = flattenTriggerDestinationCloudRun(c, m["cloudRun"])
	r.CloudFunction = dcl.FlattenString(m["cloudFunction"])
	r.Gke = flattenTriggerDestinationGke(c, m["gke"])

	return r
}

// expandTriggerDestinationCloudRunMap expands the contents of TriggerDestinationCloudRun into a JSON
// request object.
func expandTriggerDestinationCloudRunMap(c *Client, f map[string]TriggerDestinationCloudRun) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestinationCloudRun(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerDestinationCloudRunSlice expands the contents of TriggerDestinationCloudRun into a JSON
// request object.
func expandTriggerDestinationCloudRunSlice(c *Client, f []TriggerDestinationCloudRun) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestinationCloudRun(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationCloudRunMap flattens the contents of TriggerDestinationCloudRun from a JSON
// response object.
func flattenTriggerDestinationCloudRunMap(c *Client, i interface{}) map[string]TriggerDestinationCloudRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestinationCloudRun{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestinationCloudRun{}
	}

	items := make(map[string]TriggerDestinationCloudRun)
	for k, item := range a {
		items[k] = *flattenTriggerDestinationCloudRun(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerDestinationCloudRunSlice flattens the contents of TriggerDestinationCloudRun from a JSON
// response object.
func flattenTriggerDestinationCloudRunSlice(c *Client, i interface{}) []TriggerDestinationCloudRun {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestinationCloudRun{}
	}

	if len(a) == 0 {
		return []TriggerDestinationCloudRun{}
	}

	items := make([]TriggerDestinationCloudRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestinationCloudRun(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerDestinationCloudRun expands an instance of TriggerDestinationCloudRun into a JSON
// request object.
func expandTriggerDestinationCloudRun(c *Client, f *TriggerDestinationCloudRun) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenTriggerDestinationCloudRun flattens an instance of TriggerDestinationCloudRun from a JSON
// response object.
func flattenTriggerDestinationCloudRun(c *Client, i interface{}) *TriggerDestinationCloudRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerDestinationCloudRun{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerDestinationCloudRun
	}
	r.Service = dcl.FlattenString(m["service"])
	r.Path = dcl.FlattenString(m["path"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandTriggerDestinationGkeMap expands the contents of TriggerDestinationGke into a JSON
// request object.
func expandTriggerDestinationGkeMap(c *Client, f map[string]TriggerDestinationGke) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestinationGke(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerDestinationGkeSlice expands the contents of TriggerDestinationGke into a JSON
// request object.
func expandTriggerDestinationGkeSlice(c *Client, f []TriggerDestinationGke) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestinationGke(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationGkeMap flattens the contents of TriggerDestinationGke from a JSON
// response object.
func flattenTriggerDestinationGkeMap(c *Client, i interface{}) map[string]TriggerDestinationGke {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestinationGke{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestinationGke{}
	}

	items := make(map[string]TriggerDestinationGke)
	for k, item := range a {
		items[k] = *flattenTriggerDestinationGke(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerDestinationGkeSlice flattens the contents of TriggerDestinationGke from a JSON
// response object.
func flattenTriggerDestinationGkeSlice(c *Client, i interface{}) []TriggerDestinationGke {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestinationGke{}
	}

	if len(a) == 0 {
		return []TriggerDestinationGke{}
	}

	items := make([]TriggerDestinationGke, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestinationGke(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerDestinationGke expands an instance of TriggerDestinationGke into a JSON
// request object.
func expandTriggerDestinationGke(c *Client, f *TriggerDestinationGke) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Cluster; !dcl.IsEmptyValueIndirect(v) {
		m["cluster"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenTriggerDestinationGke flattens an instance of TriggerDestinationGke from a JSON
// response object.
func flattenTriggerDestinationGke(c *Client, i interface{}) *TriggerDestinationGke {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerDestinationGke{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerDestinationGke
	}
	r.Cluster = dcl.FlattenString(m["cluster"])
	r.Location = dcl.FlattenString(m["location"])
	r.Namespace = dcl.FlattenString(m["namespace"])
	r.Service = dcl.FlattenString(m["service"])
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandTriggerTransportMap expands the contents of TriggerTransport into a JSON
// request object.
func expandTriggerTransportMap(c *Client, f map[string]TriggerTransport) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerTransport(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerTransportSlice expands the contents of TriggerTransport into a JSON
// request object.
func expandTriggerTransportSlice(c *Client, f []TriggerTransport) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerTransport(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerTransportMap flattens the contents of TriggerTransport from a JSON
// response object.
func flattenTriggerTransportMap(c *Client, i interface{}) map[string]TriggerTransport {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerTransport{}
	}

	if len(a) == 0 {
		return map[string]TriggerTransport{}
	}

	items := make(map[string]TriggerTransport)
	for k, item := range a {
		items[k] = *flattenTriggerTransport(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerTransportSlice flattens the contents of TriggerTransport from a JSON
// response object.
func flattenTriggerTransportSlice(c *Client, i interface{}) []TriggerTransport {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerTransport{}
	}

	if len(a) == 0 {
		return []TriggerTransport{}
	}

	items := make([]TriggerTransport, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerTransport(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerTransport expands an instance of TriggerTransport into a JSON
// request object.
func expandTriggerTransport(c *Client, f *TriggerTransport) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTriggerTransportPubsub(c, f.Pubsub); err != nil {
		return nil, fmt.Errorf("error expanding Pubsub into pubsub: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubsub"] = v
	}

	return m, nil
}

// flattenTriggerTransport flattens an instance of TriggerTransport from a JSON
// response object.
func flattenTriggerTransport(c *Client, i interface{}) *TriggerTransport {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerTransport{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerTransport
	}
	r.Pubsub = flattenTriggerTransportPubsub(c, m["pubsub"])

	return r
}

// expandTriggerTransportPubsubMap expands the contents of TriggerTransportPubsub into a JSON
// request object.
func expandTriggerTransportPubsubMap(c *Client, f map[string]TriggerTransportPubsub) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerTransportPubsub(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerTransportPubsubSlice expands the contents of TriggerTransportPubsub into a JSON
// request object.
func expandTriggerTransportPubsubSlice(c *Client, f []TriggerTransportPubsub) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerTransportPubsub(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerTransportPubsubMap flattens the contents of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsubMap(c *Client, i interface{}) map[string]TriggerTransportPubsub {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerTransportPubsub{}
	}

	if len(a) == 0 {
		return map[string]TriggerTransportPubsub{}
	}

	items := make(map[string]TriggerTransportPubsub)
	for k, item := range a {
		items[k] = *flattenTriggerTransportPubsub(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerTransportPubsubSlice flattens the contents of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsubSlice(c *Client, i interface{}) []TriggerTransportPubsub {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerTransportPubsub{}
	}

	if len(a) == 0 {
		return []TriggerTransportPubsub{}
	}

	items := make([]TriggerTransportPubsub, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerTransportPubsub(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerTransportPubsub expands an instance of TriggerTransportPubsub into a JSON
// request object.
func expandTriggerTransportPubsub(c *Client, f *TriggerTransportPubsub) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Topic; !dcl.IsEmptyValueIndirect(v) {
		m["topic"] = v
	}
	if v := f.Subscription; !dcl.IsEmptyValueIndirect(v) {
		m["subscription"] = v
	}

	return m, nil
}

// flattenTriggerTransportPubsub flattens an instance of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsub(c *Client, i interface{}) *TriggerTransportPubsub {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerTransportPubsub{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerTransportPubsub
	}
	r.Topic = dcl.FlattenString(m["topic"])
	r.Subscription = dcl.FlattenString(m["subscription"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Trigger) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTrigger(b, c)
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

func convertFieldDiffToTriggerDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]triggerDiff, error) {
	var diffs []triggerDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := triggerDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTotriggerApiOperation(op, opts...)
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

func convertOpNameTotriggerApiOperation(op string, opts ...dcl.ApplyOption) (triggerApiOperation, error) {
	switch op {

	case "updateTriggerUpdateTriggerOperation":
		return &updateTriggerUpdateTriggerOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
