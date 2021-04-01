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

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Trigger) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "destination"); err != nil {
		return err
	}
	if err := dcl.Required(r, "matchingCriteria"); err != nil {
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
func (r *TriggerDestination) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CloudRunService) {
		if err := r.CloudRunService.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TriggerDestinationCloudRunService) validate() error {
	if err := dcl.Required(r, "service"); err != nil {
		return err
	}
	if err := dcl.Required(r, "region"); err != nil {
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
func (r *TriggerMatchingCriteria) validate() error {
	if err := dcl.Required(r, "attribute"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	return nil
}

func triggerGetURL(userBasePath string, r *Trigger) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1beta1/", userBasePath, params), nil
}

func triggerListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers", "https://eventarc.googleapis.com/v1beta1/", userBasePath, params), nil

}

func triggerCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers?triggerId={{name}}", "https://eventarc.googleapis.com/v1beta1/", userBasePath, params), nil

}

func triggerDeleteURL(userBasePath string, r *Trigger) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1beta1/", userBasePath, params), nil
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
	if v, err := expandTriggerMatchingCriteriaSlice(c, f.MatchingCriteria); err != nil {
		return nil, fmt.Errorf("error expanding MatchingCriteria into matchingCriteria: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["matchingCriteria"] = v
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
	mask := strings.Join([]string{"name", "serviceAccount", "destination", "labels", "matchingCriteria", "etag"}, ",")
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
	err = o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1beta1/", "GET")

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
	if err := o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1beta1/", "GET"); err != nil {
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
	if err := o.Wait(ctx, c.Config, "https://eventarc.googleapis.com/v1beta1/", "GET"); err != nil {
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
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.NameToSelfLink(rawDesired.ServiceAccount, rawInitial.ServiceAccount) {
		rawDesired.ServiceAccount = rawInitial.ServiceAccount
	}
	rawDesired.Destination = canonicalizeTriggerDestination(rawDesired.Destination, rawInitial.Destination, opts...)
	rawDesired.Transport = canonicalizeTriggerTransport(rawDesired.Transport, rawInitial.Transport, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.IsZeroValue(rawDesired.MatchingCriteria) {
		rawDesired.MatchingCriteria = rawInitial.MatchingCriteria
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

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
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

	if dcl.IsEmptyValueIndirect(rawNew.MatchingCriteria) && dcl.IsEmptyValueIndirect(rawDesired.MatchingCriteria) {
		rawNew.MatchingCriteria = rawDesired.MatchingCriteria
	} else {
		rawNew.MatchingCriteria = canonicalizeNewTriggerMatchingCriteriaSet(c, rawDesired.MatchingCriteria, rawNew.MatchingCriteria)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
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

	des.CloudRunService = canonicalizeTriggerDestinationCloudRunService(des.CloudRunService, initial.CloudRunService, opts...)

	return des
}

func canonicalizeNewTriggerDestination(c *Client, des, nw *TriggerDestination) *TriggerDestination {
	if des == nil || nw == nil {
		return nw
	}

	nw.CloudRunService = canonicalizeNewTriggerDestinationCloudRunService(c, des.CloudRunService, nw.CloudRunService)

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
			if !compareTriggerDestination(c, &d, &n) {
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
		return des
	}

	var items []TriggerDestination
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerDestination(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerDestinationCloudRunService(des, initial *TriggerDestinationCloudRunService, opts ...dcl.ApplyOption) *TriggerDestinationCloudRunService {
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

func canonicalizeNewTriggerDestinationCloudRunService(c *Client, des, nw *TriggerDestinationCloudRunService) *TriggerDestinationCloudRunService {
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

func canonicalizeNewTriggerDestinationCloudRunServiceSet(c *Client, des, nw []TriggerDestinationCloudRunService) []TriggerDestinationCloudRunService {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerDestinationCloudRunService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTriggerDestinationCloudRunService(c, &d, &n) {
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

func canonicalizeNewTriggerDestinationCloudRunServiceSlice(c *Client, des, nw []TriggerDestinationCloudRunService) []TriggerDestinationCloudRunService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TriggerDestinationCloudRunService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerDestinationCloudRunService(c, &d, &n))
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
			if !compareTriggerTransport(c, &d, &n) {
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
		return des
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
	if dcl.StringCanonicalize(des.Subscription, initial.Subscription) || dcl.IsZeroValue(des.Subscription) {
		des.Subscription = initial.Subscription
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
			if !compareTriggerTransportPubsub(c, &d, &n) {
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
		return des
	}

	var items []TriggerTransportPubsub
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerTransportPubsub(c, &d, &n))
	}

	return items
}

func canonicalizeTriggerMatchingCriteria(des, initial *TriggerMatchingCriteria, opts ...dcl.ApplyOption) *TriggerMatchingCriteria {
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

func canonicalizeNewTriggerMatchingCriteria(c *Client, des, nw *TriggerMatchingCriteria) *TriggerMatchingCriteria {
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

func canonicalizeNewTriggerMatchingCriteriaSet(c *Client, des, nw []TriggerMatchingCriteria) []TriggerMatchingCriteria {
	if des == nil {
		return nw
	}
	var reorderedNew []TriggerMatchingCriteria
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTriggerMatchingCriteria(c, &d, &n) {
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

func canonicalizeNewTriggerMatchingCriteriaSlice(c *Client, des, nw []TriggerMatchingCriteria) []TriggerMatchingCriteria {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TriggerMatchingCriteria
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTriggerMatchingCriteria(c, &d, &n))
	}

	return items
}

type triggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         triggerApiOperation
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
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, triggerDiff{
			UpdateOp:  &updateTriggerUpdateTriggerOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.ServiceAccount) && !dcl.NameToSelfLink(desired.ServiceAccount, actual.ServiceAccount) {
		c.Config.Logger.Infof("Detected diff in ServiceAccount.\nDESIRED: %v\nACTUAL: %v", desired.ServiceAccount, actual.ServiceAccount)

		diffs = append(diffs, triggerDiff{
			UpdateOp:  &updateTriggerUpdateTriggerOperation{},
			FieldName: "ServiceAccount",
		})

	}
	if compareTriggerDestination(c, desired.Destination, actual.Destination) {
		c.Config.Logger.Infof("Detected diff in Destination.\nDESIRED: %v\nACTUAL: %v", desired.Destination, actual.Destination)

		diffs = append(diffs, triggerDiff{
			UpdateOp:  &updateTriggerUpdateTriggerOperation{},
			FieldName: "Destination",
		})

	}
	if compareTriggerTransport(c, desired.Transport, actual.Transport) {
		c.Config.Logger.Infof("Detected diff in Transport.\nDESIRED: %v\nACTUAL: %v", desired.Transport, actual.Transport)
		diffs = append(diffs, triggerDiff{
			RequiresRecreate: true,
			FieldName:        "Transport",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, triggerDiff{
			UpdateOp:  &updateTriggerUpdateTriggerOperation{},
			FieldName: "Labels",
		})

	}
	if compareTriggerMatchingCriteriaSlice(c, desired.MatchingCriteria, actual.MatchingCriteria) {
		c.Config.Logger.Infof("Detected diff in MatchingCriteria.\nDESIRED: %v\nACTUAL: %v", desired.MatchingCriteria, actual.MatchingCriteria)

		toAdd, toRemove := compareTriggerMatchingCriteriaSets(c, desired.MatchingCriteria, actual.MatchingCriteria)
		c.Config.Logger.Infof("diff in MatchingCriteria is a set field - recomparing with set logic. \nto add: %#v\nto remove: %#v", toAdd, toRemove)
		if len(toAdd) != 0 || len(toRemove) != 0 {
			c.Config.Logger.Info("diff in MatchingCriteria persists after set logic analysis.")
			diffs = append(diffs, triggerDiff{
				UpdateOp:  &updateTriggerUpdateTriggerOperation{},
				FieldName: "MatchingCriteria",
			})
		}

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
func compareTriggerDestination(c *Client, desired, actual *TriggerDestination) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CloudRunService == nil && desired.CloudRunService != nil && !dcl.IsEmptyValueIndirect(desired.CloudRunService) {
		c.Config.Logger.Infof("desired CloudRunService %s - but actually nil", dcl.SprintResource(desired.CloudRunService))
		return true
	}
	if compareTriggerDestinationCloudRunService(c, desired.CloudRunService, actual.CloudRunService) && !dcl.IsZeroValue(desired.CloudRunService) {
		c.Config.Logger.Infof("Diff in CloudRunService. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CloudRunService), dcl.SprintResource(actual.CloudRunService))
		return true
	}
	return false
}

func compareTriggerDestinationSlice(c *Client, desired, actual []TriggerDestination) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerDestination, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTriggerDestination(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TriggerDestination, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTriggerDestinationMap(c *Client, desired, actual map[string]TriggerDestination) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerDestination, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TriggerDestination, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTriggerDestination(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TriggerDestination, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTriggerDestinationCloudRunService(c *Client, desired, actual *TriggerDestinationCloudRunService) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Service == nil && desired.Service != nil && !dcl.IsEmptyValueIndirect(desired.Service) {
		c.Config.Logger.Infof("desired Service %s - but actually nil", dcl.SprintResource(desired.Service))
		return true
	}
	if !dcl.NameToSelfLink(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) {
		c.Config.Logger.Infof("Diff in Service. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if actual.Path == nil && desired.Path != nil && !dcl.IsEmptyValueIndirect(desired.Path) {
		c.Config.Logger.Infof("desired Path %s - but actually nil", dcl.SprintResource(desired.Path))
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if actual.Region == nil && desired.Region != nil && !dcl.IsEmptyValueIndirect(desired.Region) {
		c.Config.Logger.Infof("desired Region %s - but actually nil", dcl.SprintResource(desired.Region))
		return true
	}
	if !dcl.StringCanonicalize(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) {
		c.Config.Logger.Infof("Diff in Region. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}

func compareTriggerDestinationCloudRunServiceSlice(c *Client, desired, actual []TriggerDestinationCloudRunService) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerDestinationCloudRunService, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTriggerDestinationCloudRunService(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TriggerDestinationCloudRunService, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTriggerDestinationCloudRunServiceMap(c *Client, desired, actual map[string]TriggerDestinationCloudRunService) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerDestinationCloudRunService, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TriggerDestinationCloudRunService, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTriggerDestinationCloudRunService(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TriggerDestinationCloudRunService, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTriggerTransport(c *Client, desired, actual *TriggerTransport) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Pubsub == nil && desired.Pubsub != nil && !dcl.IsEmptyValueIndirect(desired.Pubsub) {
		c.Config.Logger.Infof("desired Pubsub %s - but actually nil", dcl.SprintResource(desired.Pubsub))
		return true
	}
	if compareTriggerTransportPubsub(c, desired.Pubsub, actual.Pubsub) && !dcl.IsZeroValue(desired.Pubsub) {
		c.Config.Logger.Infof("Diff in Pubsub. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Pubsub), dcl.SprintResource(actual.Pubsub))
		return true
	}
	return false
}

func compareTriggerTransportSlice(c *Client, desired, actual []TriggerTransport) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerTransport, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTriggerTransport(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TriggerTransport, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTriggerTransportMap(c *Client, desired, actual map[string]TriggerTransport) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerTransport, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TriggerTransport, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTriggerTransport(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TriggerTransport, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTriggerTransportPubsub(c *Client, desired, actual *TriggerTransportPubsub) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Topic == nil && desired.Topic != nil && !dcl.IsEmptyValueIndirect(desired.Topic) {
		c.Config.Logger.Infof("desired Topic %s - but actually nil", dcl.SprintResource(desired.Topic))
		return true
	}
	if !dcl.StringCanonicalize(desired.Topic, actual.Topic) && !dcl.IsZeroValue(desired.Topic) {
		c.Config.Logger.Infof("Diff in Topic. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Topic), dcl.SprintResource(actual.Topic))
		return true
	}
	return false
}

func compareTriggerTransportPubsubSlice(c *Client, desired, actual []TriggerTransportPubsub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerTransportPubsub, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTriggerTransportPubsub(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TriggerTransportPubsub, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTriggerTransportPubsubMap(c *Client, desired, actual map[string]TriggerTransportPubsub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerTransportPubsub, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TriggerTransportPubsub, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTriggerTransportPubsub(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TriggerTransportPubsub, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTriggerMatchingCriteria(c *Client, desired, actual *TriggerMatchingCriteria) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Attribute == nil && desired.Attribute != nil && !dcl.IsEmptyValueIndirect(desired.Attribute) {
		c.Config.Logger.Infof("desired Attribute %s - but actually nil", dcl.SprintResource(desired.Attribute))
		return true
	}
	if !dcl.StringCanonicalize(desired.Attribute, actual.Attribute) && !dcl.IsZeroValue(desired.Attribute) {
		c.Config.Logger.Infof("Diff in Attribute. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Attribute), dcl.SprintResource(actual.Attribute))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareTriggerMatchingCriteriaSlice(c *Client, desired, actual []TriggerMatchingCriteria) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerMatchingCriteria, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTriggerMatchingCriteria(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TriggerMatchingCriteria, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTriggerMatchingCriteriaMap(c *Client, desired, actual map[string]TriggerMatchingCriteria) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TriggerMatchingCriteria, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TriggerMatchingCriteria, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTriggerMatchingCriteria(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TriggerMatchingCriteria, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTriggerMatchingCriteriaSets(c *Client, desired, actual []TriggerMatchingCriteria) (toAdd, toRemove []TriggerMatchingCriteria) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]TriggerMatchingCriteria)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]TriggerMatchingCriteria)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]TriggerMatchingCriteria, 0)
	toRemove = make([]TriggerMatchingCriteria, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareTriggerMatchingCriteria(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareTriggerMatchingCriteria(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Trigger) urlNormalized() *Trigger {
	normalized := deepcopy.Copy(*r).(Trigger)
	normalized.Name = dcl.SelfLinkToName(r.Name)
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
		return dcl.URL("projects/{{project}}/locations/{{location}}/triggers/{{name}}", "https://eventarc.googleapis.com/v1beta1/", userBasePath, fields), nil

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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandTriggerDestination(c, f.Destination); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v, err := expandTriggerTransport(c, f.Transport); err != nil {
		return nil, fmt.Errorf("error expanding Transport into transport: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["transport"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v, err := expandTriggerMatchingCriteriaSlice(c, f.MatchingCriteria); err != nil {
		return nil, fmt.Errorf("error expanding MatchingCriteria into matchingCriteria: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["matchingCriteria"] = v
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
func flattenTrigger(c *Client, i interface{}) *Trigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Trigger{}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Destination = flattenTriggerDestination(c, m["destination"])
	r.Transport = flattenTriggerTransport(c, m["transport"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.MatchingCriteria = flattenTriggerMatchingCriteriaSlice(c, m["matchingCriteria"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

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
	if v, err := expandTriggerDestinationCloudRunService(c, f.CloudRunService); err != nil {
		return nil, fmt.Errorf("error expanding CloudRunService into cloudRunService: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRunService"] = v
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
	r.CloudRunService = flattenTriggerDestinationCloudRunService(c, m["cloudRunService"])

	return r
}

// expandTriggerDestinationCloudRunServiceMap expands the contents of TriggerDestinationCloudRunService into a JSON
// request object.
func expandTriggerDestinationCloudRunServiceMap(c *Client, f map[string]TriggerDestinationCloudRunService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerDestinationCloudRunService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerDestinationCloudRunServiceSlice expands the contents of TriggerDestinationCloudRunService into a JSON
// request object.
func expandTriggerDestinationCloudRunServiceSlice(c *Client, f []TriggerDestinationCloudRunService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerDestinationCloudRunService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerDestinationCloudRunServiceMap flattens the contents of TriggerDestinationCloudRunService from a JSON
// response object.
func flattenTriggerDestinationCloudRunServiceMap(c *Client, i interface{}) map[string]TriggerDestinationCloudRunService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerDestinationCloudRunService{}
	}

	if len(a) == 0 {
		return map[string]TriggerDestinationCloudRunService{}
	}

	items := make(map[string]TriggerDestinationCloudRunService)
	for k, item := range a {
		items[k] = *flattenTriggerDestinationCloudRunService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerDestinationCloudRunServiceSlice flattens the contents of TriggerDestinationCloudRunService from a JSON
// response object.
func flattenTriggerDestinationCloudRunServiceSlice(c *Client, i interface{}) []TriggerDestinationCloudRunService {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerDestinationCloudRunService{}
	}

	if len(a) == 0 {
		return []TriggerDestinationCloudRunService{}
	}

	items := make([]TriggerDestinationCloudRunService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerDestinationCloudRunService(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerDestinationCloudRunService expands an instance of TriggerDestinationCloudRunService into a JSON
// request object.
func expandTriggerDestinationCloudRunService(c *Client, f *TriggerDestinationCloudRunService) (map[string]interface{}, error) {
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

// flattenTriggerDestinationCloudRunService flattens an instance of TriggerDestinationCloudRunService from a JSON
// response object.
func flattenTriggerDestinationCloudRunService(c *Client, i interface{}) *TriggerDestinationCloudRunService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerDestinationCloudRunService{}
	r.Service = dcl.FlattenString(m["service"])
	r.Path = dcl.FlattenString(m["path"])
	r.Region = dcl.FlattenString(m["region"])

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
	r.Topic = dcl.FlattenString(m["topic"])
	r.Subscription = dcl.FlattenString(m["subscription"])

	return r
}

// expandTriggerMatchingCriteriaMap expands the contents of TriggerMatchingCriteria into a JSON
// request object.
func expandTriggerMatchingCriteriaMap(c *Client, f map[string]TriggerMatchingCriteria) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTriggerMatchingCriteria(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTriggerMatchingCriteriaSlice expands the contents of TriggerMatchingCriteria into a JSON
// request object.
func expandTriggerMatchingCriteriaSlice(c *Client, f []TriggerMatchingCriteria) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTriggerMatchingCriteria(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTriggerMatchingCriteriaMap flattens the contents of TriggerMatchingCriteria from a JSON
// response object.
func flattenTriggerMatchingCriteriaMap(c *Client, i interface{}) map[string]TriggerMatchingCriteria {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TriggerMatchingCriteria{}
	}

	if len(a) == 0 {
		return map[string]TriggerMatchingCriteria{}
	}

	items := make(map[string]TriggerMatchingCriteria)
	for k, item := range a {
		items[k] = *flattenTriggerMatchingCriteria(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTriggerMatchingCriteriaSlice flattens the contents of TriggerMatchingCriteria from a JSON
// response object.
func flattenTriggerMatchingCriteriaSlice(c *Client, i interface{}) []TriggerMatchingCriteria {
	a, ok := i.([]interface{})
	if !ok {
		return []TriggerMatchingCriteria{}
	}

	if len(a) == 0 {
		return []TriggerMatchingCriteria{}
	}

	items := make([]TriggerMatchingCriteria, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTriggerMatchingCriteria(c, item.(map[string]interface{})))
	}

	return items
}

// expandTriggerMatchingCriteria expands an instance of TriggerMatchingCriteria into a JSON
// request object.
func expandTriggerMatchingCriteria(c *Client, f *TriggerMatchingCriteria) (map[string]interface{}, error) {
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

// flattenTriggerMatchingCriteria flattens an instance of TriggerMatchingCriteria from a JSON
// response object.
func flattenTriggerMatchingCriteria(c *Client, i interface{}) *TriggerMatchingCriteria {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TriggerMatchingCriteria{}
	r.Attribute = dcl.FlattenString(m["attribute"])
	r.Value = dcl.FlattenString(m["value"])

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
