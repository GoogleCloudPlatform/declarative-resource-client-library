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

func (r *Membership) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Endpoint) {
		if err := r.Endpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Authority) {
		if err := r.Authority.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpoint) validate() error {
	if !dcl.IsEmptyValueIndirect(r.GkeCluster) {
		if err := r.GkeCluster.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubernetesMetadata) {
		if err := r.KubernetesMetadata.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubernetesResource) {
		if err := r.KubernetesResource.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpointGkeCluster) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesMetadata) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ResourceOptions) {
		if err := r.ResourceOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpointKubernetesResourceMembershipResources) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResourceConnectResources) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResourceResourceOptions) validate() error {
	return nil
}
func (r *MembershipState) validate() error {
	return nil
}
func (r *MembershipAuthority) validate() error {
	return nil
}

func membershipGetURL(userBasePath string, r *Membership) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil
}

func membershipListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil

}

func membershipCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships?membershipId={{name}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil

}

func membershipDeleteURL(userBasePath string, r *Membership) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil
}

// membershipApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type membershipApiOperation interface {
	do(context.Context, *Membership, *Client) error
}

// newUpdateMembershipUpdateMembershipRequest creates a request for an
// Membership resource's UpdateMembership update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateMembershipUpdateMembershipRequest(ctx context.Context, f *Membership, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandMembershipEndpoint(c, f.Endpoint); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["endpoint"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.ExternalId; !dcl.IsEmptyValueIndirect(v) {
		req["externalId"] = v
	}
	if v, err := expandMembershipAuthority(c, f.Authority); err != nil {
		return nil, fmt.Errorf("error expanding Authority into authority: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["authority"] = v
	}
	if v := f.InfrastructureType; !dcl.IsEmptyValueIndirect(v) {
		req["infrastructureType"] = v
	}
	return req, nil
}

// marshalUpdateMembershipUpdateMembershipRequest converts the update into
// the final JSON request body.
func marshalUpdateMembershipUpdateMembershipRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateMembershipUpdateMembershipOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateMembershipUpdateMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	_, err := c.GetMembership(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateMembership")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateMembershipUpdateMembershipRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateMembershipUpdateMembershipRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listMembershipRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := membershipListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MembershipMaxPage {
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

type listMembershipOperation struct {
	Resources []map[string]interface{} `json:"resources"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listMembership(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Membership, string, error) {
	b, err := c.listMembershipRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMembershipOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Membership
	for _, v := range m.Resources {
		res, err := unmarshalMapMembership(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllMembership(ctx context.Context, f func(*Membership) bool, resources []*Membership) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteMembership(ctx, res)
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

type deleteMembershipOperation struct{}

func (op *deleteMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {

	_, err := c.GetMembership(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Membership not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetMembership checking for existence. error: %v", err)
		return err
	}

	u, err := membershipDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetMembership(ctx, r.urlNormalized())
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
type createMembershipOperation struct {
	response map[string]interface{}
}

func (op *createMembershipOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := membershipCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetMembership(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMembershipRaw(ctx context.Context, r *Membership) ([]byte, error) {

	u, err := membershipGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) membershipDiffsForRawDesired(ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (initial, desired *Membership, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Membership
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Membership); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Membership, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMembership(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Membership resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Membership resource: %v", err)
		}
		c.Config.Logger.Info("Found that Membership resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Membership: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Membership: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMembershipInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Membership: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Membership: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMembership(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMembershipInitialState(rawInitial, rawDesired *Membership) (*Membership, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMembershipDesiredState(rawDesired, rawInitial *Membership, opts ...dcl.ApplyOption) (*Membership, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Endpoint = canonicalizeMembershipEndpoint(rawDesired.Endpoint, nil, opts...)
		rawDesired.State = canonicalizeMembershipState(rawDesired.State, nil, opts...)
		rawDesired.Authority = canonicalizeMembershipAuthority(rawDesired.Authority, nil, opts...)

		return rawDesired, nil
	}

	rawDesired.Endpoint = canonicalizeMembershipEndpoint(rawDesired.Endpoint, rawInitial.Endpoint, opts...)
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.ExternalId, rawInitial.ExternalId) {
		rawDesired.ExternalId = rawInitial.ExternalId
	}
	rawDesired.Authority = canonicalizeMembershipAuthority(rawDesired.Authority, rawInitial.Authority, opts...)
	if dcl.IsZeroValue(rawDesired.InfrastructureType) {
		rawDesired.InfrastructureType = rawInitial.InfrastructureType
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeMembershipNewState(c *Client, rawNew, rawDesired *Membership) (*Membership, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Endpoint) && dcl.IsEmptyValueIndirect(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		rawNew.Endpoint = canonicalizeNewMembershipEndpoint(c, rawDesired.Endpoint, rawNew.Endpoint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewMembershipState(c, rawDesired.State, rawNew.State)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExternalId) && dcl.IsEmptyValueIndirect(rawDesired.ExternalId) {
		rawNew.ExternalId = rawDesired.ExternalId
	} else {
		if dcl.StringCanonicalize(rawDesired.ExternalId, rawNew.ExternalId) {
			rawNew.ExternalId = rawDesired.ExternalId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastConnectionTime) && dcl.IsEmptyValueIndirect(rawDesired.LastConnectionTime) {
		rawNew.LastConnectionTime = rawDesired.LastConnectionTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UniqueId) && dcl.IsEmptyValueIndirect(rawDesired.UniqueId) {
		rawNew.UniqueId = rawDesired.UniqueId
	} else {
		if dcl.StringCanonicalize(rawDesired.UniqueId, rawNew.UniqueId) {
			rawNew.UniqueId = rawDesired.UniqueId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Authority) && dcl.IsEmptyValueIndirect(rawDesired.Authority) {
		rawNew.Authority = rawDesired.Authority
	} else {
		rawNew.Authority = canonicalizeNewMembershipAuthority(c, rawDesired.Authority, rawNew.Authority)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InfrastructureType) && dcl.IsEmptyValueIndirect(rawDesired.InfrastructureType) {
		rawNew.InfrastructureType = rawDesired.InfrastructureType
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeMembershipEndpoint(des, initial *MembershipEndpoint, opts ...dcl.ApplyOption) *MembershipEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.GkeCluster = canonicalizeMembershipEndpointGkeCluster(des.GkeCluster, initial.GkeCluster, opts...)
	des.KubernetesResource = canonicalizeMembershipEndpointKubernetesResource(des.KubernetesResource, initial.KubernetesResource, opts...)

	return des
}

func canonicalizeNewMembershipEndpoint(c *Client, des, nw *MembershipEndpoint) *MembershipEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	nw.GkeCluster = canonicalizeNewMembershipEndpointGkeCluster(c, des.GkeCluster, nw.GkeCluster)
	nw.KubernetesMetadata = canonicalizeNewMembershipEndpointKubernetesMetadata(c, des.KubernetesMetadata, nw.KubernetesMetadata)
	nw.KubernetesResource = canonicalizeNewMembershipEndpointKubernetesResource(c, des.KubernetesResource, nw.KubernetesResource)

	return nw
}

func canonicalizeNewMembershipEndpointSet(c *Client, des, nw []MembershipEndpoint) []MembershipEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointSlice(c *Client, des, nw []MembershipEndpoint) []MembershipEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointGkeCluster(des, initial *MembershipEndpointGkeCluster, opts ...dcl.ApplyOption) *MembershipEndpointGkeCluster {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ResourceLink, initial.ResourceLink) || dcl.IsZeroValue(des.ResourceLink) {
		des.ResourceLink = initial.ResourceLink
	}

	return des
}

func canonicalizeNewMembershipEndpointGkeCluster(c *Client, des, nw *MembershipEndpointGkeCluster) *MembershipEndpointGkeCluster {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ResourceLink, nw.ResourceLink) {
		nw.ResourceLink = des.ResourceLink
	}

	return nw
}

func canonicalizeNewMembershipEndpointGkeClusterSet(c *Client, des, nw []MembershipEndpointGkeCluster) []MembershipEndpointGkeCluster {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointGkeCluster
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointGkeClusterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointGkeClusterSlice(c *Client, des, nw []MembershipEndpointGkeCluster) []MembershipEndpointGkeCluster {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointGkeCluster
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointGkeCluster(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesMetadata(des, initial *MembershipEndpointKubernetesMetadata, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewMembershipEndpointKubernetesMetadata(c *Client, des, nw *MembershipEndpointKubernetesMetadata) *MembershipEndpointKubernetesMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KubernetesApiServerVersion, nw.KubernetesApiServerVersion) {
		nw.KubernetesApiServerVersion = des.KubernetesApiServerVersion
	}
	if dcl.StringCanonicalize(des.NodeProviderId, nw.NodeProviderId) {
		nw.NodeProviderId = des.NodeProviderId
	}
	if dcl.IsZeroValue(nw.NodeCount) {
		nw.NodeCount = des.NodeCount
	}
	if dcl.IsZeroValue(nw.VcpuCount) {
		nw.VcpuCount = des.VcpuCount
	}
	if dcl.IsZeroValue(nw.MemoryMb) {
		nw.MemoryMb = des.MemoryMb
	}
	if dcl.IsZeroValue(nw.UpdateTime) {
		nw.UpdateTime = des.UpdateTime
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesMetadataSet(c *Client, des, nw []MembershipEndpointKubernetesMetadata) []MembershipEndpointKubernetesMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointKubernetesMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointKubernetesMetadataSlice(c *Client, des, nw []MembershipEndpointKubernetesMetadata) []MembershipEndpointKubernetesMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResource(des, initial *MembershipEndpointKubernetesResource, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MembershipCrManifest, initial.MembershipCrManifest) || dcl.IsZeroValue(des.MembershipCrManifest) {
		des.MembershipCrManifest = initial.MembershipCrManifest
	}
	des.ResourceOptions = canonicalizeMembershipEndpointKubernetesResourceResourceOptions(des.ResourceOptions, initial.ResourceOptions, opts...)

	return des
}

func canonicalizeNewMembershipEndpointKubernetesResource(c *Client, des, nw *MembershipEndpointKubernetesResource) *MembershipEndpointKubernetesResource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MembershipCrManifest, nw.MembershipCrManifest) {
		nw.MembershipCrManifest = des.MembershipCrManifest
	}
	nw.MembershipResources = canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSlice(c, des.MembershipResources, nw.MembershipResources)
	nw.ConnectResources = canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSlice(c, des.ConnectResources, nw.ConnectResources)
	nw.ResourceOptions = canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c, des.ResourceOptions, nw.ResourceOptions)

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceSet(c *Client, des, nw []MembershipEndpointKubernetesResource) []MembershipEndpointKubernetesResource {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointKubernetesResource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointKubernetesResourceSlice(c *Client, des, nw []MembershipEndpointKubernetesResource) []MembershipEndpointKubernetesResource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResource(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceMembershipResources(des, initial *MembershipEndpointKubernetesResourceMembershipResources, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Manifest, initial.Manifest) || dcl.IsZeroValue(des.Manifest) {
		des.Manifest = initial.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, initial.ClusterScoped) || dcl.IsZeroValue(des.ClusterScoped) {
		des.ClusterScoped = initial.ClusterScoped
	}

	return des
}

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResources(c *Client, des, nw *MembershipEndpointKubernetesResourceMembershipResources) *MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Manifest, nw.Manifest) {
		nw.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, nw.ClusterScoped) {
		nw.ClusterScoped = des.ClusterScoped
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSet(c *Client, des, nw []MembershipEndpointKubernetesResourceMembershipResources) []MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointKubernetesResourceMembershipResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceMembershipResources) []MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceMembershipResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceMembershipResources(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceConnectResources(des, initial *MembershipEndpointKubernetesResourceConnectResources, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Manifest, initial.Manifest) || dcl.IsZeroValue(des.Manifest) {
		des.Manifest = initial.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, initial.ClusterScoped) || dcl.IsZeroValue(des.ClusterScoped) {
		des.ClusterScoped = initial.ClusterScoped
	}

	return des
}

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResources(c *Client, des, nw *MembershipEndpointKubernetesResourceConnectResources) *MembershipEndpointKubernetesResourceConnectResources {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Manifest, nw.Manifest) {
		nw.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, nw.ClusterScoped) {
		nw.ClusterScoped = des.ClusterScoped
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSet(c *Client, des, nw []MembershipEndpointKubernetesResourceConnectResources) []MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointKubernetesResourceConnectResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceConnectResources) []MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceConnectResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceConnectResources(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceResourceOptions(des, initial *MembershipEndpointKubernetesResourceResourceOptions, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ConnectVersion, initial.ConnectVersion) || dcl.IsZeroValue(des.ConnectVersion) {
		des.ConnectVersion = initial.ConnectVersion
	}
	if dcl.BoolCanonicalize(des.V1Beta1Crd, initial.V1Beta1Crd) || dcl.IsZeroValue(des.V1Beta1Crd) {
		des.V1Beta1Crd = initial.V1Beta1Crd
	}

	return des
}

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c *Client, des, nw *MembershipEndpointKubernetesResourceResourceOptions) *MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ConnectVersion, nw.ConnectVersion) {
		nw.ConnectVersion = des.ConnectVersion
	}
	if dcl.BoolCanonicalize(des.V1Beta1Crd, nw.V1Beta1Crd) {
		nw.V1Beta1Crd = des.V1Beta1Crd
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptionsSet(c *Client, des, nw []MembershipEndpointKubernetesResourceResourceOptions) []MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipEndpointKubernetesResourceResourceOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceResourceOptions) []MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceResourceOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipState(des, initial *MembershipState, opts ...dcl.ApplyOption) *MembershipState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewMembershipState(c *Client, des, nw *MembershipState) *MembershipState {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Code) {
		nw.Code = des.Code
	}

	return nw
}

func canonicalizeNewMembershipStateSet(c *Client, des, nw []MembershipState) []MembershipState {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipStateSlice(c *Client, des, nw []MembershipState) []MembershipState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipState(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipAuthority(des, initial *MembershipAuthority, opts ...dcl.ApplyOption) *MembershipAuthority {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Issuer, initial.Issuer) || dcl.IsZeroValue(des.Issuer) {
		des.Issuer = initial.Issuer
	}

	return des
}

func canonicalizeNewMembershipAuthority(c *Client, des, nw *MembershipAuthority) *MembershipAuthority {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Issuer, nw.Issuer) {
		nw.Issuer = des.Issuer
	}
	if dcl.StringCanonicalize(des.WorkloadIdentityPool, nw.WorkloadIdentityPool) {
		nw.WorkloadIdentityPool = des.WorkloadIdentityPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, nw.IdentityProvider) {
		nw.IdentityProvider = des.IdentityProvider
	}

	return nw
}

func canonicalizeNewMembershipAuthoritySet(c *Client, des, nw []MembershipAuthority) []MembershipAuthority {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipAuthority
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipAuthorityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMembershipAuthoritySlice(c *Client, des, nw []MembershipAuthority) []MembershipAuthority {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipAuthority
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipAuthority(c, &d, &n))
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
func diffMembership(c *Client, desired, actual *Membership, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.Info{ObjectFunction: compareMembershipEndpointNewStyle, EmptyObject: EmptyMembershipEndpoint, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, ObjectFunction: compareMembershipStateNewStyle, EmptyObject: EmptyMembershipState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalId, actual.ExternalId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("ExternalId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastConnectionTime, actual.LastConnectionTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastConnectionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UniqueId, actual.UniqueId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UniqueId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Authority, actual.Authority, dcl.Info{ObjectFunction: compareMembershipAuthorityNewStyle, EmptyObject: EmptyMembershipAuthority, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Authority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InfrastructureType, actual.InfrastructureType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("InfrastructureType")); len(ds) != 0 || err != nil {
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
func compareMembershipEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpoint or *MembershipEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpoint)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GkeCluster, actual.GkeCluster, dcl.Info{ObjectFunction: compareMembershipEndpointGkeClusterNewStyle, EmptyObject: EmptyMembershipEndpointGkeCluster, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GkeCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesMetadata, actual.KubernetesMetadata, dcl.Info{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesMetadataNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesMetadata, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesMetadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesResource, actual.KubernetesResource, dcl.Info{ObjectFunction: compareMembershipEndpointKubernetesResourceNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResource, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesResource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointGkeClusterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointGkeCluster)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointGkeCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointGkeCluster or *MembershipEndpointGkeCluster", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointGkeCluster)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointGkeCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointGkeCluster", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceLink, actual.ResourceLink, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesMetadata)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesMetadata or *MembershipEndpointKubernetesMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesMetadata)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KubernetesApiServerVersion, actual.KubernetesApiServerVersion, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesApiServerVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeProviderId, actual.NodeProviderId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeProviderId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VcpuCount, actual.VcpuCount, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VcpuCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemoryMb, actual.MemoryMb, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemoryMb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResource)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResource or *MembershipEndpointKubernetesResource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResource)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MembershipCrManifest, actual.MembershipCrManifest, dcl.Info{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MembershipCrManifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MembershipResources, actual.MembershipResources, dcl.Info{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceMembershipResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MembershipResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectResources, actual.ConnectResources, dcl.Info{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceConnectResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceOptions, actual.ResourceOptions, dcl.Info{ObjectFunction: compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceResourceOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceMembershipResources)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceMembershipResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceMembershipResources or *MembershipEndpointKubernetesResourceMembershipResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceMembershipResources)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceMembershipResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceMembershipResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Manifest, actual.Manifest, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Manifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterScoped, actual.ClusterScoped, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterScoped")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceConnectResources)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceConnectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceConnectResources or *MembershipEndpointKubernetesResourceConnectResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceConnectResources)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceConnectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceConnectResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Manifest, actual.Manifest, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Manifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterScoped, actual.ClusterScoped, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterScoped")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceResourceOptions)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceResourceOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceResourceOptions or *MembershipEndpointKubernetesResourceResourceOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceResourceOptions)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceResourceOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceResourceOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConnectVersion, actual.ConnectVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V1Beta1Crd, actual.V1Beta1Crd, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("V1beta1Crd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipState)
	if !ok {
		desiredNotPointer, ok := d.(MembershipState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipState or *MembershipState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipState)
	if !ok {
		actualNotPointer, ok := a.(MembershipState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipAuthorityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipAuthority)
	if !ok {
		desiredNotPointer, ok := d.(MembershipAuthority)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipAuthority or *MembershipAuthority", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipAuthority)
	if !ok {
		actualNotPointer, ok := a.(MembershipAuthority)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipAuthority", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Issuer, actual.Issuer, dcl.Info{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Issuer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadIdentityPool, actual.WorkloadIdentityPool, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdentityProvider, actual.IdentityProvider, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdentityProvider")); len(ds) != 0 || err != nil {
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
func (r *Membership) urlNormalized() *Membership {
	normalized := dcl.Copy(*r).(Membership)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.ExternalId = dcl.SelfLinkToName(r.ExternalId)
	normalized.UniqueId = dcl.SelfLinkToName(r.UniqueId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Membership) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Membership) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Membership) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Membership) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateMembership" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Membership resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Membership) marshal(c *Client) ([]byte, error) {
	m, err := expandMembership(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Membership: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMembership decodes JSON responses into the Membership resource schema.
func unmarshalMembership(b []byte, c *Client) (*Membership, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMembership(m, c)
}

func unmarshalMapMembership(m map[string]interface{}, c *Client) (*Membership, error) {

	return flattenMembership(c, m), nil
}

// expandMembership expands Membership into a JSON request object.
func expandMembership(c *Client, f *Membership) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := expandMembershipEndpoint(c, f.Endpoint); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if v != nil {
		m["endpoint"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/locations/%s/memberships/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandMembershipState(c, f.State); err != nil {
		return nil, fmt.Errorf("error expanding State into state: %w", err)
	} else if v != nil {
		m["state"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.DeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["deleteTime"] = v
	}
	if v := f.ExternalId; !dcl.IsEmptyValueIndirect(v) {
		m["externalId"] = v
	}
	if v := f.LastConnectionTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastConnectionTime"] = v
	}
	if v := f.UniqueId; !dcl.IsEmptyValueIndirect(v) {
		m["uniqueId"] = v
	}
	if v, err := expandMembershipAuthority(c, f.Authority); err != nil {
		return nil, fmt.Errorf("error expanding Authority into authority: %w", err)
	} else if v != nil {
		m["authority"] = v
	}
	if v := f.InfrastructureType; !dcl.IsEmptyValueIndirect(v) {
		m["infrastructureType"] = v
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

// flattenMembership flattens Membership from a JSON request object into the
// Membership type.
func flattenMembership(c *Client, i interface{}) *Membership {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Membership{}
	res.Endpoint = flattenMembershipEndpoint(c, m["endpoint"])
	res.Name = dcl.FlattenString(m["name"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Description = dcl.FlattenString(m["description"])
	res.State = flattenMembershipState(c, m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.ExternalId = dcl.FlattenString(m["externalId"])
	res.LastConnectionTime = dcl.FlattenString(m["lastConnectionTime"])
	res.UniqueId = dcl.FlattenString(m["uniqueId"])
	res.Authority = flattenMembershipAuthority(c, m["authority"])
	res.InfrastructureType = flattenMembershipInfrastructureTypeEnum(m["infrastructureType"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandMembershipEndpointMap expands the contents of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpointMap(c *Client, f map[string]MembershipEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointSlice expands the contents of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpointSlice(c *Client, f []MembershipEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointMap flattens the contents of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpointMap(c *Client, i interface{}) map[string]MembershipEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpoint{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpoint{}
	}

	items := make(map[string]MembershipEndpoint)
	for k, item := range a {
		items[k] = *flattenMembershipEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointSlice flattens the contents of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpointSlice(c *Client, i interface{}) []MembershipEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpoint{}
	}

	if len(a) == 0 {
		return []MembershipEndpoint{}
	}

	items := make([]MembershipEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpoint expands an instance of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpoint(c *Client, f *MembershipEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandMembershipEndpointGkeCluster(c, f.GkeCluster); err != nil {
		return nil, fmt.Errorf("error expanding GkeCluster into gkeCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gkeCluster"] = v
	}
	if v, err := expandMembershipEndpointKubernetesMetadata(c, f.KubernetesMetadata); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesMetadata into kubernetesMetadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesMetadata"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResource(c, f.KubernetesResource); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesResource into kubernetesResource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesResource"] = v
	}

	return m, nil
}

// flattenMembershipEndpoint flattens an instance of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpoint(c *Client, i interface{}) *MembershipEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpoint
	}
	r.GkeCluster = flattenMembershipEndpointGkeCluster(c, m["gkeCluster"])
	r.KubernetesMetadata = flattenMembershipEndpointKubernetesMetadata(c, m["kubernetesMetadata"])
	r.KubernetesResource = flattenMembershipEndpointKubernetesResource(c, m["kubernetesResource"])

	return r
}

// expandMembershipEndpointGkeClusterMap expands the contents of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeClusterMap(c *Client, f map[string]MembershipEndpointGkeCluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointGkeCluster(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointGkeClusterSlice expands the contents of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeClusterSlice(c *Client, f []MembershipEndpointGkeCluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointGkeCluster(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointGkeClusterMap flattens the contents of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeClusterMap(c *Client, i interface{}) map[string]MembershipEndpointGkeCluster {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointGkeCluster{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointGkeCluster{}
	}

	items := make(map[string]MembershipEndpointGkeCluster)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointGkeCluster(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointGkeClusterSlice flattens the contents of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeClusterSlice(c *Client, i interface{}) []MembershipEndpointGkeCluster {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointGkeCluster{}
	}

	if len(a) == 0 {
		return []MembershipEndpointGkeCluster{}
	}

	items := make([]MembershipEndpointGkeCluster, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointGkeCluster(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointGkeCluster expands an instance of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeCluster(c *Client, f *MembershipEndpointGkeCluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHubReferenceLink(f, f.ResourceLink); err != nil {
		return nil, fmt.Errorf("error expanding ResourceLink into resourceLink: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceLink"] = v
	}

	return m, nil
}

// flattenMembershipEndpointGkeCluster flattens an instance of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeCluster(c *Client, i interface{}) *MembershipEndpointGkeCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointGkeCluster{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointGkeCluster
	}
	r.ResourceLink = flattenHubReferenceLink(m["resourceLink"])

	return r
}

// expandMembershipEndpointKubernetesMetadataMap expands the contents of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadataMap(c *Client, f map[string]MembershipEndpointKubernetesMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesMetadataSlice expands the contents of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadataSlice(c *Client, f []MembershipEndpointKubernetesMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesMetadataMap flattens the contents of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadataMap(c *Client, i interface{}) map[string]MembershipEndpointKubernetesMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesMetadata{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesMetadata{}
	}

	items := make(map[string]MembershipEndpointKubernetesMetadata)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointKubernetesMetadataSlice flattens the contents of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadataSlice(c *Client, i interface{}) []MembershipEndpointKubernetesMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesMetadata{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesMetadata{}
	}

	items := make([]MembershipEndpointKubernetesMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointKubernetesMetadata expands an instance of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadata(c *Client, f *MembershipEndpointKubernetesMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KubernetesApiServerVersion; !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesApiServerVersion"] = v
	}
	if v := f.NodeProviderId; !dcl.IsEmptyValueIndirect(v) {
		m["nodeProviderId"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["nodeCount"] = v
	}
	if v := f.VcpuCount; !dcl.IsEmptyValueIndirect(v) {
		m["vcpuCount"] = v
	}
	if v := f.MemoryMb; !dcl.IsEmptyValueIndirect(v) {
		m["memoryMb"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesMetadata flattens an instance of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadata(c *Client, i interface{}) *MembershipEndpointKubernetesMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesMetadata
	}
	r.KubernetesApiServerVersion = dcl.FlattenString(m["kubernetesApiServerVersion"])
	r.NodeProviderId = dcl.FlattenString(m["nodeProviderId"])
	r.NodeCount = dcl.FlattenInteger(m["nodeCount"])
	r.VcpuCount = dcl.FlattenInteger(m["vcpuCount"])
	r.MemoryMb = dcl.FlattenInteger(m["memoryMb"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandMembershipEndpointKubernetesResourceMap expands the contents of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMap(c *Client, f map[string]MembershipEndpointKubernetesResource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceSlice expands the contents of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceSlice(c *Client, f []MembershipEndpointKubernetesResource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceMap flattens the contents of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMap(c *Client, i interface{}) map[string]MembershipEndpointKubernetesResource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResource{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResource{}
	}

	items := make(map[string]MembershipEndpointKubernetesResource)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceSlice flattens the contents of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceSlice(c *Client, i interface{}) []MembershipEndpointKubernetesResource {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResource{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResource{}
	}

	items := make([]MembershipEndpointKubernetesResource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResource(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointKubernetesResource expands an instance of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResource(c *Client, f *MembershipEndpointKubernetesResource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MembershipCrManifest; !dcl.IsEmptyValueIndirect(v) {
		m["membershipCrManifest"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResourceMembershipResourcesSlice(c, f.MembershipResources); err != nil {
		return nil, fmt.Errorf("error expanding MembershipResources into membershipResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["membershipResources"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResourceConnectResourcesSlice(c, f.ConnectResources); err != nil {
		return nil, fmt.Errorf("error expanding ConnectResources into connectResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["connectResources"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, f.ResourceOptions); err != nil {
		return nil, fmt.Errorf("error expanding ResourceOptions into resourceOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceOptions"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResource flattens an instance of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResource(c *Client, i interface{}) *MembershipEndpointKubernetesResource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResource
	}
	r.MembershipCrManifest = dcl.FlattenSecretValue(m["membershipCrManifest"])
	r.MembershipResources = flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice(c, m["membershipResources"])
	r.ConnectResources = flattenMembershipEndpointKubernetesResourceConnectResourcesSlice(c, m["connectResources"])
	r.ResourceOptions = flattenMembershipEndpointKubernetesResourceResourceOptions(c, m["resourceOptions"])

	return r
}

// expandMembershipEndpointKubernetesResourceMembershipResourcesMap expands the contents of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResourcesMap(c *Client, f map[string]MembershipEndpointKubernetesResourceMembershipResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceMembershipResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceMembershipResourcesSlice expands the contents of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, f []MembershipEndpointKubernetesResourceMembershipResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceMembershipResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceMembershipResourcesMap flattens the contents of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResourcesMap(c *Client, i interface{}) map[string]MembershipEndpointKubernetesResourceMembershipResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceMembershipResources{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceMembershipResources{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceMembershipResources)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceMembershipResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice flattens the contents of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, i interface{}) []MembershipEndpointKubernetesResourceMembershipResources {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceMembershipResources{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceMembershipResources{}
	}

	items := make([]MembershipEndpointKubernetesResourceMembershipResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceMembershipResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceMembershipResources expands an instance of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResources(c *Client, f *MembershipEndpointKubernetesResourceMembershipResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Manifest; !dcl.IsEmptyValueIndirect(v) {
		m["manifest"] = v
	}
	if v := f.ClusterScoped; !dcl.IsEmptyValueIndirect(v) {
		m["clusterScoped"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceMembershipResources flattens an instance of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResources(c *Client, i interface{}) *MembershipEndpointKubernetesResourceMembershipResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceMembershipResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceMembershipResources
	}
	r.Manifest = dcl.FlattenString(m["manifest"])
	r.ClusterScoped = dcl.FlattenBool(m["clusterScoped"])

	return r
}

// expandMembershipEndpointKubernetesResourceConnectResourcesMap expands the contents of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResourcesMap(c *Client, f map[string]MembershipEndpointKubernetesResourceConnectResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceConnectResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceConnectResourcesSlice expands the contents of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, f []MembershipEndpointKubernetesResourceConnectResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceConnectResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceConnectResourcesMap flattens the contents of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResourcesMap(c *Client, i interface{}) map[string]MembershipEndpointKubernetesResourceConnectResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceConnectResources{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceConnectResources{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceConnectResources)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceConnectResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceConnectResourcesSlice flattens the contents of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, i interface{}) []MembershipEndpointKubernetesResourceConnectResources {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceConnectResources{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceConnectResources{}
	}

	items := make([]MembershipEndpointKubernetesResourceConnectResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceConnectResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceConnectResources expands an instance of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResources(c *Client, f *MembershipEndpointKubernetesResourceConnectResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Manifest; !dcl.IsEmptyValueIndirect(v) {
		m["manifest"] = v
	}
	if v := f.ClusterScoped; !dcl.IsEmptyValueIndirect(v) {
		m["clusterScoped"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceConnectResources flattens an instance of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResources(c *Client, i interface{}) *MembershipEndpointKubernetesResourceConnectResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceConnectResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceConnectResources
	}
	r.Manifest = dcl.FlattenString(m["manifest"])
	r.ClusterScoped = dcl.FlattenBool(m["clusterScoped"])

	return r
}

// expandMembershipEndpointKubernetesResourceResourceOptionsMap expands the contents of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptionsMap(c *Client, f map[string]MembershipEndpointKubernetesResourceResourceOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceResourceOptionsSlice expands the contents of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, f []MembershipEndpointKubernetesResourceResourceOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceResourceOptionsMap flattens the contents of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptionsMap(c *Client, i interface{}) map[string]MembershipEndpointKubernetesResourceResourceOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceResourceOptions{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceResourceOptions{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceResourceOptions)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceResourceOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceResourceOptionsSlice flattens the contents of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, i interface{}) []MembershipEndpointKubernetesResourceResourceOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceResourceOptions{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceResourceOptions{}
	}

	items := make([]MembershipEndpointKubernetesResourceResourceOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceResourceOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceResourceOptions expands an instance of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptions(c *Client, f *MembershipEndpointKubernetesResourceResourceOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConnectVersion; !dcl.IsEmptyValueIndirect(v) {
		m["connectVersion"] = v
	}
	if v := f.V1Beta1Crd; !dcl.IsEmptyValueIndirect(v) {
		m["v1beta1Crd"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceResourceOptions flattens an instance of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptions(c *Client, i interface{}) *MembershipEndpointKubernetesResourceResourceOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceResourceOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceResourceOptions
	}
	r.ConnectVersion = dcl.FlattenString(m["connectVersion"])
	r.V1Beta1Crd = dcl.FlattenBool(m["v1beta1Crd"])

	return r
}

// expandMembershipStateMap expands the contents of MembershipState into a JSON
// request object.
func expandMembershipStateMap(c *Client, f map[string]MembershipState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipStateSlice expands the contents of MembershipState into a JSON
// request object.
func expandMembershipStateSlice(c *Client, f []MembershipState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipStateMap flattens the contents of MembershipState from a JSON
// response object.
func flattenMembershipStateMap(c *Client, i interface{}) map[string]MembershipState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipState{}
	}

	if len(a) == 0 {
		return map[string]MembershipState{}
	}

	items := make(map[string]MembershipState)
	for k, item := range a {
		items[k] = *flattenMembershipState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipStateSlice flattens the contents of MembershipState from a JSON
// response object.
func flattenMembershipStateSlice(c *Client, i interface{}) []MembershipState {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipState{}
	}

	if len(a) == 0 {
		return []MembershipState{}
	}

	items := make([]MembershipState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipState(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipState expands an instance of MembershipState into a JSON
// request object.
func expandMembershipState(c *Client, f *MembershipState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}

	return m, nil
}

// flattenMembershipState flattens an instance of MembershipState from a JSON
// response object.
func flattenMembershipState(c *Client, i interface{}) *MembershipState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipState
	}
	r.Code = flattenMembershipStateCodeEnum(m["code"])

	return r
}

// expandMembershipAuthorityMap expands the contents of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthorityMap(c *Client, f map[string]MembershipAuthority) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipAuthority(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipAuthoritySlice expands the contents of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthoritySlice(c *Client, f []MembershipAuthority) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipAuthority(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipAuthorityMap flattens the contents of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthorityMap(c *Client, i interface{}) map[string]MembershipAuthority {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipAuthority{}
	}

	if len(a) == 0 {
		return map[string]MembershipAuthority{}
	}

	items := make(map[string]MembershipAuthority)
	for k, item := range a {
		items[k] = *flattenMembershipAuthority(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipAuthoritySlice flattens the contents of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthoritySlice(c *Client, i interface{}) []MembershipAuthority {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipAuthority{}
	}

	if len(a) == 0 {
		return []MembershipAuthority{}
	}

	items := make([]MembershipAuthority, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipAuthority(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipAuthority expands an instance of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthority(c *Client, f *MembershipAuthority) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Issuer; !dcl.IsEmptyValueIndirect(v) {
		m["issuer"] = v
	}
	if v := f.WorkloadIdentityPool; !dcl.IsEmptyValueIndirect(v) {
		m["workloadIdentityPool"] = v
	}
	if v := f.IdentityProvider; !dcl.IsEmptyValueIndirect(v) {
		m["identityProvider"] = v
	}

	return m, nil
}

// flattenMembershipAuthority flattens an instance of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthority(c *Client, i interface{}) *MembershipAuthority {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipAuthority{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipAuthority
	}
	r.Issuer = dcl.FlattenString(m["issuer"])
	r.WorkloadIdentityPool = dcl.FlattenString(m["workloadIdentityPool"])
	r.IdentityProvider = dcl.FlattenString(m["identityProvider"])

	return r
}

// flattenMembershipStateCodeEnumSlice flattens the contents of MembershipStateCodeEnum from a JSON
// response object.
func flattenMembershipStateCodeEnumSlice(c *Client, i interface{}) []MembershipStateCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipStateCodeEnum{}
	}

	if len(a) == 0 {
		return []MembershipStateCodeEnum{}
	}

	items := make([]MembershipStateCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipStateCodeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipStateCodeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipStateCodeEnum with the same value as that string.
func flattenMembershipStateCodeEnum(i interface{}) *MembershipStateCodeEnum {
	s, ok := i.(string)
	if !ok {
		return MembershipStateCodeEnumRef("")
	}

	return MembershipStateCodeEnumRef(s)
}

// flattenMembershipInfrastructureTypeEnumSlice flattens the contents of MembershipInfrastructureTypeEnum from a JSON
// response object.
func flattenMembershipInfrastructureTypeEnumSlice(c *Client, i interface{}) []MembershipInfrastructureTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipInfrastructureTypeEnum{}
	}

	if len(a) == 0 {
		return []MembershipInfrastructureTypeEnum{}
	}

	items := make([]MembershipInfrastructureTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipInfrastructureTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipInfrastructureTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipInfrastructureTypeEnum with the same value as that string.
func flattenMembershipInfrastructureTypeEnum(i interface{}) *MembershipInfrastructureTypeEnum {
	s, ok := i.(string)
	if !ok {
		return MembershipInfrastructureTypeEnumRef("")
	}

	return MembershipInfrastructureTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Membership) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMembership(b, c)
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

type membershipDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         membershipApiOperation
}

func convertFieldDiffToMembershipOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]membershipDiff, error) {
	var diffs []membershipDiff
	for _, op := range ops {
		diff := membershipDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTomembershipApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTomembershipApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (membershipApiOperation, error) {
	switch op {

	case "updateMembershipUpdateMembershipOperation":
		return &updateMembershipUpdateMembershipOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
