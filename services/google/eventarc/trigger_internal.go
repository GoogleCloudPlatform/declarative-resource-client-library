// Copyright 2022 Google LLC. All Rights Reserved.
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
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string(nil)); err != nil {
		return err
	}
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
func (r *Trigger) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://eventarc.googleapis.com/v1/", params)
}

func (r *Trigger) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Trigger) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers", nr.basePath(), userBasePath, params), nil

}

func (r *Trigger) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers?triggerId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Trigger) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", nr.basePath(), userBasePath, params), nil
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
	res := f
	_ = res

	if v, err := expandTriggerEventFiltersSlice(c, f.EventFilters, res); err != nil {
		return nil, fmt.Errorf("error expanding EventFilters into eventFilters: %w", err)
	} else if v != nil {
		req["eventFilters"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		req["serviceAccount"] = v
	}
	if v, err := expandTriggerDestination(c, f.Destination, res); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["destination"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getTriggerRaw(ctx, f)
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
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTriggerUpdateTriggerOperation) do(ctx context.Context, r *Trigger, c *Client) error {
	_, err := c.GetTrigger(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTrigger")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTriggerUpdateTriggerRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTriggerRaw(ctx context.Context, r *Trigger, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
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

func (c *Client) listTrigger(ctx context.Context, r *Trigger, pageToken string, pageSize int32) ([]*Trigger, string, error) {
	b, err := c.listTriggerRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTriggerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Trigger
	for _, v := range m.Triggers {
		res, err := unmarshalMapTrigger(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
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
	r, err := c.GetTrigger(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Trigger not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTrigger checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetTrigger(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
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
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetTrigger(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTriggerRaw(ctx context.Context, r *Trigger) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) triggerDiffsForRawDesired(ctx context.Context, rawDesired *Trigger, opts ...dcl.ApplyOption) (initial, desired *Trigger, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Trigger
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Trigger); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Trigger, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTrigger(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Trigger resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Trigger resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Trigger resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTriggerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Trigger: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Trigger: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTriggerFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTriggerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Trigger: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTriggerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Trigger: %v", desired)

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
	canonicalDesired := &Trigger{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.EventFilters = canonicalizeTriggerEventFiltersSlice(rawDesired.EventFilters, rawInitial.EventFilters, opts...)
	if dcl.IsZeroValue(rawDesired.ServiceAccount) || (dcl.IsEmptyValueIndirect(rawDesired.ServiceAccount) && dcl.IsEmptyValueIndirect(rawInitial.ServiceAccount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ServiceAccount = rawInitial.ServiceAccount
	} else {
		canonicalDesired.ServiceAccount = rawDesired.ServiceAccount
	}
	canonicalDesired.Destination = canonicalizeTriggerDestination(rawDesired.Destination, rawInitial.Destination, opts...)
	canonicalDesired.Transport = canonicalizeTriggerTransport(rawDesired.Transport, rawInitial.Transport, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
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

func canonicalizeTriggerNewState(c *Client, rawNew, rawDesired *Trigger) (*Trigger, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Uid) && dcl.IsNotReturnedByServer(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.EventFilters) && dcl.IsNotReturnedByServer(rawDesired.EventFilters) {
		rawNew.EventFilters = rawDesired.EventFilters
	} else {
		rawNew.EventFilters = canonicalizeNewTriggerEventFiltersSet(c, rawDesired.EventFilters, rawNew.EventFilters)
	}

	if dcl.IsNotReturnedByServer(rawNew.ServiceAccount) && dcl.IsNotReturnedByServer(rawDesired.ServiceAccount) {
		rawNew.ServiceAccount = rawDesired.ServiceAccount
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Destination) && dcl.IsNotReturnedByServer(rawDesired.Destination) {
		rawNew.Destination = rawDesired.Destination
	} else {
		rawNew.Destination = canonicalizeNewTriggerDestination(c, rawDesired.Destination, rawNew.Destination)
	}

	if dcl.IsNotReturnedByServer(rawNew.Transport) && dcl.IsNotReturnedByServer(rawDesired.Transport) {
		rawNew.Transport = rawDesired.Transport
	} else {
		rawNew.Transport = canonicalizeNewTriggerTransport(c, rawDesired.Transport, rawNew.Transport)
	}

	if dcl.IsNotReturnedByServer(rawNew.Labels) && dcl.IsNotReturnedByServer(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
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

	cDes := &TriggerEventFilters{}

	if dcl.StringCanonicalize(des.Attribute, initial.Attribute) || dcl.IsZeroValue(des.Attribute) {
		cDes.Attribute = initial.Attribute
	} else {
		cDes.Attribute = des.Attribute
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Operator, initial.Operator) || dcl.IsZeroValue(des.Operator) {
		cDes.Operator = initial.Operator
	} else {
		cDes.Operator = des.Operator
	}

	return cDes
}

func canonicalizeTriggerEventFiltersSlice(des, initial []TriggerEventFilters, opts ...dcl.ApplyOption) []TriggerEventFilters {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerEventFilters, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerEventFilters(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerEventFilters, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerEventFilters(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerEventFilters(c *Client, des, nw *TriggerEventFilters) *TriggerEventFilters {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerEventFilters while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Attribute, nw.Attribute) {
		nw.Attribute = des.Attribute
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Operator, nw.Operator) {
		nw.Operator = des.Operator
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

	cDes := &TriggerDestination{}

	cDes.CloudRun = canonicalizeTriggerDestinationCloudRun(des.CloudRun, initial.CloudRun, opts...)
	if dcl.PartialSelfLinkToSelfLink(des.CloudFunction, initial.CloudFunction) || dcl.IsZeroValue(des.CloudFunction) {
		cDes.CloudFunction = initial.CloudFunction
	} else {
		cDes.CloudFunction = des.CloudFunction
	}
	cDes.Gke = canonicalizeTriggerDestinationGke(des.Gke, initial.Gke, opts...)
	if dcl.PartialSelfLinkToSelfLink(des.Workflow, initial.Workflow) || dcl.IsZeroValue(des.Workflow) {
		cDes.Workflow = initial.Workflow
	} else {
		cDes.Workflow = des.Workflow
	}

	return cDes
}

func canonicalizeTriggerDestinationSlice(des, initial []TriggerDestination, opts ...dcl.ApplyOption) []TriggerDestination {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerDestination, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerDestination(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerDestination, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerDestination(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerDestination(c *Client, des, nw *TriggerDestination) *TriggerDestination {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerDestination while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.CloudRun = canonicalizeNewTriggerDestinationCloudRun(c, des.CloudRun, nw.CloudRun)
	if dcl.PartialSelfLinkToSelfLink(des.CloudFunction, nw.CloudFunction) {
		nw.CloudFunction = des.CloudFunction
	}
	nw.Gke = canonicalizeNewTriggerDestinationGke(c, des.Gke, nw.Gke)
	if dcl.PartialSelfLinkToSelfLink(des.Workflow, nw.Workflow) {
		nw.Workflow = des.Workflow
	}

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

	cDes := &TriggerDestinationCloudRun{}

	if dcl.IsZeroValue(des.Service) || (dcl.IsEmptyValueIndirect(des.Service) && dcl.IsEmptyValueIndirect(initial.Service)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Region, initial.Region) || dcl.IsZeroValue(des.Region) {
		cDes.Region = initial.Region
	} else {
		cDes.Region = des.Region
	}

	return cDes
}

func canonicalizeTriggerDestinationCloudRunSlice(des, initial []TriggerDestinationCloudRun, opts ...dcl.ApplyOption) []TriggerDestinationCloudRun {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerDestinationCloudRun, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerDestinationCloudRun(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerDestinationCloudRun, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerDestinationCloudRun(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerDestinationCloudRun(c *Client, des, nw *TriggerDestinationCloudRun) *TriggerDestinationCloudRun {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerDestinationCloudRun while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
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

	cDes := &TriggerDestinationGke{}

	if dcl.IsZeroValue(des.Cluster) || (dcl.IsEmptyValueIndirect(des.Cluster) && dcl.IsEmptyValueIndirect(initial.Cluster)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Cluster = initial.Cluster
	} else {
		cDes.Cluster = des.Cluster
	}
	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeTriggerDestinationGkeSlice(des, initial []TriggerDestinationGke, opts ...dcl.ApplyOption) []TriggerDestinationGke {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerDestinationGke, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerDestinationGke(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerDestinationGke, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerDestinationGke(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerDestinationGke(c *Client, des, nw *TriggerDestinationGke) *TriggerDestinationGke {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerDestinationGke while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
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

	cDes := &TriggerTransport{}

	cDes.Pubsub = canonicalizeTriggerTransportPubsub(des.Pubsub, initial.Pubsub, opts...)

	return cDes
}

func canonicalizeTriggerTransportSlice(des, initial []TriggerTransport, opts ...dcl.ApplyOption) []TriggerTransport {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerTransport, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerTransport(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerTransport, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerTransport(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerTransport(c *Client, des, nw *TriggerTransport) *TriggerTransport {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerTransport while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
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

	cDes := &TriggerTransportPubsub{}

	if dcl.PartialSelfLinkToSelfLink(des.Topic, initial.Topic) || dcl.IsZeroValue(des.Topic) {
		cDes.Topic = initial.Topic
	} else {
		cDes.Topic = des.Topic
	}

	return cDes
}

func canonicalizeTriggerTransportPubsubSlice(des, initial []TriggerTransportPubsub, opts ...dcl.ApplyOption) []TriggerTransportPubsub {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TriggerTransportPubsub, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTriggerTransportPubsub(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TriggerTransportPubsub, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTriggerTransportPubsub(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTriggerTransportPubsub(c *Client, des, nw *TriggerTransportPubsub) *TriggerTransportPubsub {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TriggerTransportPubsub while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.PartialSelfLinkToSelfLink(des.Topic, nw.Topic) {
		nw.Topic = des.Topic
	}
	if dcl.PartialSelfLinkToSelfLink(des.Subscription, nw.Subscription) {
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

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTrigger(c *Client, desired, actual *Trigger, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EventFilters, actual.EventFilters, dcl.Info{Type: "Set", ObjectFunction: compareTriggerEventFiltersNewStyle, EmptyObject: EmptyTriggerEventFilters, OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("EventFilters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{ObjectFunction: compareTriggerDestinationNewStyle, EmptyObject: EmptyTriggerDestination, OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Transport, actual.Transport, dcl.Info{ObjectFunction: compareTriggerTransportNewStyle, EmptyObject: EmptyTriggerTransport, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Transport")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Operator, actual.Operator, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Operator")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CloudRun, actual.CloudRun, dcl.Info{ObjectFunction: compareTriggerDestinationCloudRunNewStyle, EmptyObject: EmptyTriggerDestinationCloudRun, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudRun")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Gke, actual.Gke, dcl.Info{ObjectFunction: compareTriggerDestinationGkeNewStyle, EmptyObject: EmptyTriggerDestinationGke, OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Gke")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Workflow, actual.Workflow, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Workflow")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Cluster, actual.Cluster, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Cluster")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTriggerUpdateTriggerOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Pubsub, actual.Pubsub, dcl.Info{ObjectFunction: compareTriggerTransportPubsubNewStyle, EmptyObject: EmptyTriggerTransportPubsub, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Pubsub")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subscription, actual.Subscription, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subscription")); len(ds) != 0 || err != nil {
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

func (r *Trigger) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateTrigger" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", nr.basePath(), userBasePath, fields), nil

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
func unmarshalTrigger(b []byte, c *Client, res *Trigger) (*Trigger, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTrigger(m, c, res)
}

func unmarshalMapTrigger(m map[string]interface{}, c *Client, res *Trigger) (*Trigger, error) {

	flattened := flattenTrigger(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTrigger expands Trigger into a JSON request object.
func expandTrigger(c *Client, f *Trigger) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/triggers/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandTriggerEventFiltersSlice(c, f.EventFilters, res); err != nil {
		return nil, fmt.Errorf("error expanding EventFilters into eventFilters: %w", err)
	} else if v != nil {
		m["eventFilters"] = v
	}
	if v := f.ServiceAccount; dcl.ValueShouldBeSent(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandTriggerDestination(c, f.Destination, res); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v, err := expandTriggerTransport(c, f.Transport, res); err != nil {
		return nil, fmt.Errorf("error expanding Transport into transport: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["transport"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenTrigger flattens Trigger from a JSON request object into the
// Trigger type.
func flattenTrigger(c *Client, i interface{}, res *Trigger) *Trigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Trigger{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.EventFilters = flattenTriggerEventFiltersSlice(c, m["eventFilters"], res)
	resultRes.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	resultRes.Destination = flattenTriggerDestination(c, m["destination"], res)
	resultRes.Transport = flattenTriggerTransport(c, m["transport"], res)
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandTriggerEventFiltersMap expands the contents of TriggerEventFilters into a JSON
// request object.
func expandTriggerEventFiltersMap(c *Client, f map[string]TriggerEventFilters, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerEventFilters(c, &item, res)
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
func expandTriggerEventFiltersSlice(c *Client, f []TriggerEventFilters, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerEventFilters(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerEventFiltersMap flattens the contents of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFiltersMap(c *Client, i interface{}, res *Trigger) map[string]TriggerEventFilters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerEventFilters{}
	}

	if len(a) == 0 {
		return map[string]TriggerEventFilters{}
	}

	items := make(map[string]TriggerEventFilters)
	for k, item := range a {
		items[k] = *flattenTriggerEventFilters(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerEventFiltersSlice flattens the contents of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFiltersSlice(c *Client, i interface{}, res *Trigger) []TriggerEventFilters {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerEventFilters{}
	}

	if len(a) == 0 {
		return []TriggerEventFilters{}
	}

	items := make([]TriggerEventFilters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerEventFilters(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerEventFilters expands an instance of TriggerEventFilters into a JSON
// request object.
func expandTriggerEventFilters(c *Client, f *TriggerEventFilters, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Attribute; !dcl.IsEmptyValueIndirect(v) {
		m["attribute"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Operator; !dcl.IsEmptyValueIndirect(v) {
		m["operator"] = v
	}

	return m, nil
}

// flattenTriggerEventFilters flattens an instance of TriggerEventFilters from a JSON
// response object.
func flattenTriggerEventFilters(c *Client, i interface{}, res *Trigger) *TriggerEventFilters {
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
	r.Operator = dcl.FlattenString(m["operator"])

	return r
}

// expandTriggerDestinationMap expands the contents of TriggerDestination into a JSON
// request object.
func expandTriggerDestinationMap(c *Client, f map[string]TriggerDestination, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestination(c, &item, res)
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
func expandTriggerDestinationSlice(c *Client, f []TriggerDestination, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestination(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationMap flattens the contents of TriggerDestination from a JSON
// response object.
func flattenTriggerDestinationMap(c *Client, i interface{}, res *Trigger) map[string]TriggerDestination {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestination{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestination{}
	}

	items := make(map[string]TriggerDestination)
	for k, item := range a {
		items[k] = *flattenTriggerDestination(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerDestinationSlice flattens the contents of TriggerDestination from a JSON
// response object.
func flattenTriggerDestinationSlice(c *Client, i interface{}, res *Trigger) []TriggerDestination {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestination{}
	}

	if len(a) == 0 {
		return []TriggerDestination{}
	}

	items := make([]TriggerDestination, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestination(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerDestination expands an instance of TriggerDestination into a JSON
// request object.
func expandTriggerDestination(c *Client, f *TriggerDestination, res *Trigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTriggerDestinationCloudRun(c, f.CloudRun, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRun"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/locations/%s/functions/%s", f.CloudFunction, dcl.SelfLinkToName(res.Project), dcl.SelfLinkToName(res.Location), dcl.SelfLinkToName(f.CloudFunction)); err != nil {
		return nil, fmt.Errorf("error expanding CloudFunction into cloudFunction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudFunction"] = v
	}
	if v, err := expandTriggerDestinationGke(c, f.Gke, res); err != nil {
		return nil, fmt.Errorf("error expanding Gke into gke: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gke"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/locations/%s/workflows/%s", f.Workflow, dcl.SelfLinkToName(res.Project), dcl.SelfLinkToName(res.Location), dcl.SelfLinkToName(f.Workflow)); err != nil {
		return nil, fmt.Errorf("error expanding Workflow into workflow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workflow"] = v
	}

	return m, nil
}

// flattenTriggerDestination flattens an instance of TriggerDestination from a JSON
// response object.
func flattenTriggerDestination(c *Client, i interface{}, res *Trigger) *TriggerDestination {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerDestination{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerDestination
	}
	r.CloudRun = flattenTriggerDestinationCloudRun(c, m["cloudRun"], res)
	r.CloudFunction = dcl.FlattenString(m["cloudFunction"])
	r.Gke = flattenTriggerDestinationGke(c, m["gke"], res)
	r.Workflow = dcl.FlattenString(m["workflow"])

	return r
}

// expandTriggerDestinationCloudRunMap expands the contents of TriggerDestinationCloudRun into a JSON
// request object.
func expandTriggerDestinationCloudRunMap(c *Client, f map[string]TriggerDestinationCloudRun, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestinationCloudRun(c, &item, res)
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
func expandTriggerDestinationCloudRunSlice(c *Client, f []TriggerDestinationCloudRun, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestinationCloudRun(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationCloudRunMap flattens the contents of TriggerDestinationCloudRun from a JSON
// response object.
func flattenTriggerDestinationCloudRunMap(c *Client, i interface{}, res *Trigger) map[string]TriggerDestinationCloudRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestinationCloudRun{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestinationCloudRun{}
	}

	items := make(map[string]TriggerDestinationCloudRun)
	for k, item := range a {
		items[k] = *flattenTriggerDestinationCloudRun(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerDestinationCloudRunSlice flattens the contents of TriggerDestinationCloudRun from a JSON
// response object.
func flattenTriggerDestinationCloudRunSlice(c *Client, i interface{}, res *Trigger) []TriggerDestinationCloudRun {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestinationCloudRun{}
	}

	if len(a) == 0 {
		return []TriggerDestinationCloudRun{}
	}

	items := make([]TriggerDestinationCloudRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestinationCloudRun(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerDestinationCloudRun expands an instance of TriggerDestinationCloudRun into a JSON
// request object.
func expandTriggerDestinationCloudRun(c *Client, f *TriggerDestinationCloudRun, res *Trigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.Service); err != nil {
		return nil, fmt.Errorf("error expanding Service into service: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
func flattenTriggerDestinationCloudRun(c *Client, i interface{}, res *Trigger) *TriggerDestinationCloudRun {
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
func expandTriggerDestinationGkeMap(c *Client, f map[string]TriggerDestinationGke, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestinationGke(c, &item, res)
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
func expandTriggerDestinationGkeSlice(c *Client, f []TriggerDestinationGke, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestinationGke(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationGkeMap flattens the contents of TriggerDestinationGke from a JSON
// response object.
func flattenTriggerDestinationGkeMap(c *Client, i interface{}, res *Trigger) map[string]TriggerDestinationGke {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestinationGke{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestinationGke{}
	}

	items := make(map[string]TriggerDestinationGke)
	for k, item := range a {
		items[k] = *flattenTriggerDestinationGke(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerDestinationGkeSlice flattens the contents of TriggerDestinationGke from a JSON
// response object.
func flattenTriggerDestinationGkeSlice(c *Client, i interface{}, res *Trigger) []TriggerDestinationGke {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestinationGke{}
	}

	if len(a) == 0 {
		return []TriggerDestinationGke{}
	}

	items := make([]TriggerDestinationGke, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestinationGke(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerDestinationGke expands an instance of TriggerDestinationGke into a JSON
// request object.
func expandTriggerDestinationGke(c *Client, f *TriggerDestinationGke, res *Trigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.Cluster); err != nil {
		return nil, fmt.Errorf("error expanding Cluster into cluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
func flattenTriggerDestinationGke(c *Client, i interface{}, res *Trigger) *TriggerDestinationGke {
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
func expandTriggerTransportMap(c *Client, f map[string]TriggerTransport, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerTransport(c, &item, res)
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
func expandTriggerTransportSlice(c *Client, f []TriggerTransport, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerTransport(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerTransportMap flattens the contents of TriggerTransport from a JSON
// response object.
func flattenTriggerTransportMap(c *Client, i interface{}, res *Trigger) map[string]TriggerTransport {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerTransport{}
	}

	if len(a) == 0 {
		return map[string]TriggerTransport{}
	}

	items := make(map[string]TriggerTransport)
	for k, item := range a {
		items[k] = *flattenTriggerTransport(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerTransportSlice flattens the contents of TriggerTransport from a JSON
// response object.
func flattenTriggerTransportSlice(c *Client, i interface{}, res *Trigger) []TriggerTransport {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerTransport{}
	}

	if len(a) == 0 {
		return []TriggerTransport{}
	}

	items := make([]TriggerTransport, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerTransport(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerTransport expands an instance of TriggerTransport into a JSON
// request object.
func expandTriggerTransport(c *Client, f *TriggerTransport, res *Trigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTriggerTransportPubsub(c, f.Pubsub, res); err != nil {
		return nil, fmt.Errorf("error expanding Pubsub into pubsub: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubsub"] = v
	}

	return m, nil
}

// flattenTriggerTransport flattens an instance of TriggerTransport from a JSON
// response object.
func flattenTriggerTransport(c *Client, i interface{}, res *Trigger) *TriggerTransport {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerTransport{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTriggerTransport
	}
	r.Pubsub = flattenTriggerTransportPubsub(c, m["pubsub"], res)

	return r
}

// expandTriggerTransportPubsubMap expands the contents of TriggerTransportPubsub into a JSON
// request object.
func expandTriggerTransportPubsubMap(c *Client, f map[string]TriggerTransportPubsub, res *Trigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerTransportPubsub(c, &item, res)
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
func expandTriggerTransportPubsubSlice(c *Client, f []TriggerTransportPubsub, res *Trigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerTransportPubsub(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerTransportPubsubMap flattens the contents of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsubMap(c *Client, i interface{}, res *Trigger) map[string]TriggerTransportPubsub {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerTransportPubsub{}
	}

	if len(a) == 0 {
		return map[string]TriggerTransportPubsub{}
	}

	items := make(map[string]TriggerTransportPubsub)
	for k, item := range a {
		items[k] = *flattenTriggerTransportPubsub(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTriggerTransportPubsubSlice flattens the contents of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsubSlice(c *Client, i interface{}, res *Trigger) []TriggerTransportPubsub {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerTransportPubsub{}
	}

	if len(a) == 0 {
		return []TriggerTransportPubsub{}
	}

	items := make([]TriggerTransportPubsub, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerTransportPubsub(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTriggerTransportPubsub expands an instance of TriggerTransportPubsub into a JSON
// request object.
func expandTriggerTransportPubsub(c *Client, f *TriggerTransportPubsub, res *Trigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/topics/%s", f.Topic, dcl.SelfLinkToName(res.Project), dcl.SelfLinkToName(f.Topic)); err != nil {
		return nil, fmt.Errorf("error expanding Topic into topic: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["topic"] = v
	}

	return m, nil
}

// flattenTriggerTransportPubsub flattens an instance of TriggerTransportPubsub from a JSON
// response object.
func flattenTriggerTransportPubsub(c *Client, i interface{}, res *Trigger) *TriggerTransportPubsub {
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
		cr, err := unmarshalTrigger(b, c, r)
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

type triggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         triggerApiOperation
}

func convertFieldDiffsToTriggerDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]triggerDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []triggerDiff
	// For each operation name, create a triggerDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := triggerDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTriggerApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTriggerApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (triggerApiOperation, error) {
	switch opName {

	case "updateTriggerUpdateTriggerOperation":
		return &updateTriggerUpdateTriggerOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTriggerFields(r *Trigger) error {
	vDestination := r.Destination
	if vDestination == nil {
		// note: explicitly not the empty object.
		vDestination = &TriggerDestination{}
	}
	if err := extractTriggerDestinationFields(r, vDestination); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDestination) {
		r.Destination = vDestination
	}
	vTransport := r.Transport
	if vTransport == nil {
		// note: explicitly not the empty object.
		vTransport = &TriggerTransport{}
	}
	if err := extractTriggerTransportFields(r, vTransport); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTransport) {
		r.Transport = vTransport
	}
	return nil
}
func extractTriggerEventFiltersFields(r *Trigger, o *TriggerEventFilters) error {
	return nil
}
func extractTriggerDestinationFields(r *Trigger, o *TriggerDestination) error {
	vCloudRun := o.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &TriggerDestinationCloudRun{}
	}
	if err := extractTriggerDestinationCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudRun) {
		o.CloudRun = vCloudRun
	}
	vGke := o.Gke
	if vGke == nil {
		// note: explicitly not the empty object.
		vGke = &TriggerDestinationGke{}
	}
	if err := extractTriggerDestinationGkeFields(r, vGke); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGke) {
		o.Gke = vGke
	}
	return nil
}
func extractTriggerDestinationCloudRunFields(r *Trigger, o *TriggerDestinationCloudRun) error {
	return nil
}
func extractTriggerDestinationGkeFields(r *Trigger, o *TriggerDestinationGke) error {
	return nil
}
func extractTriggerTransportFields(r *Trigger, o *TriggerTransport) error {
	vPubsub := o.Pubsub
	if vPubsub == nil {
		// note: explicitly not the empty object.
		vPubsub = &TriggerTransportPubsub{}
	}
	if err := extractTriggerTransportPubsubFields(r, vPubsub); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPubsub) {
		o.Pubsub = vPubsub
	}
	return nil
}
func extractTriggerTransportPubsubFields(r *Trigger, o *TriggerTransportPubsub) error {
	return nil
}

func postReadExtractTriggerFields(r *Trigger) error {
	vDestination := r.Destination
	if vDestination == nil {
		// note: explicitly not the empty object.
		vDestination = &TriggerDestination{}
	}
	if err := postReadExtractTriggerDestinationFields(r, vDestination); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDestination) {
		r.Destination = vDestination
	}
	vTransport := r.Transport
	if vTransport == nil {
		// note: explicitly not the empty object.
		vTransport = &TriggerTransport{}
	}
	if err := postReadExtractTriggerTransportFields(r, vTransport); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTransport) {
		r.Transport = vTransport
	}
	return nil
}
func postReadExtractTriggerEventFiltersFields(r *Trigger, o *TriggerEventFilters) error {
	return nil
}
func postReadExtractTriggerDestinationFields(r *Trigger, o *TriggerDestination) error {
	vCloudRun := o.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &TriggerDestinationCloudRun{}
	}
	if err := extractTriggerDestinationCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudRun) {
		o.CloudRun = vCloudRun
	}
	vGke := o.Gke
	if vGke == nil {
		// note: explicitly not the empty object.
		vGke = &TriggerDestinationGke{}
	}
	if err := extractTriggerDestinationGkeFields(r, vGke); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGke) {
		o.Gke = vGke
	}
	return nil
}
func postReadExtractTriggerDestinationCloudRunFields(r *Trigger, o *TriggerDestinationCloudRun) error {
	return nil
}
func postReadExtractTriggerDestinationGkeFields(r *Trigger, o *TriggerDestinationGke) error {
	return nil
}
func postReadExtractTriggerTransportFields(r *Trigger, o *TriggerTransport) error {
	vPubsub := o.Pubsub
	if vPubsub == nil {
		// note: explicitly not the empty object.
		vPubsub = &TriggerTransportPubsub{}
	}
	if err := extractTriggerTransportPubsubFields(r, vPubsub); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPubsub) {
		o.Pubsub = vPubsub
	}
	return nil
}
func postReadExtractTriggerTransportPubsubFields(r *Trigger, o *TriggerTransportPubsub) error {
	return nil
}
