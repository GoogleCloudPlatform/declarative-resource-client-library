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
package bigquery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Dataset) validate() error {

	if !dcl.IsEmptyValueIndirect(r.DefaultEncryptionConfiguration) {
		if err := r.DefaultEncryptionConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DatasetAccess) validate() error {
	if err := dcl.Required(r, "role"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.View) {
		if err := r.View.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Routine) {
		if err := r.Routine.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DatasetAccessView) validate() error {
	if err := dcl.Required(r, "projectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "datasetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "tableId"); err != nil {
		return err
	}
	return nil
}
func (r *DatasetAccessRoutine) validate() error {
	if err := dcl.Required(r, "projectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "datasetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "routineId"); err != nil {
		return err
	}
	return nil
}
func (r *DatasetDefaultEncryptionConfiguration) validate() error {
	return nil
}

func datasetGetURL(userBasePath string, r *Dataset) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{name}}", "https://bigquery.googleapis.com/bigquery/v2/", userBasePath, params), nil
}

func datasetListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/datasets", "https://bigquery.googleapis.com/bigquery/v2/", userBasePath, params), nil

}

func datasetCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/datasets", "https://bigquery.googleapis.com/bigquery/v2/", userBasePath, params), nil

}

func datasetDeleteURL(userBasePath string, r *Dataset) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{name}}", "https://bigquery.googleapis.com/bigquery/v2/", userBasePath, params), nil
}

// datasetApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type datasetApiOperation interface {
	do(context.Context, *Dataset, *Client) error
}

// newUpdateDatasetPatchDatasetRequest creates a request for an
// Dataset resource's PatchDataset update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDatasetPatchDatasetRequest(ctx context.Context, f *Dataset, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		req["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.DefaultTableExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		req["defaultTableExpirationMs"] = v
	}
	if v := f.DefaultPartitionExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		req["defaultPartitionExpirationMs"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandDatasetAccessSlice(c, f.Access); err != nil {
		return nil, fmt.Errorf("error expanding Access into access: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["access"] = v
	}
	b, err := c.getDatasetRaw(ctx, f.urlNormalized())
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

// marshalUpdateDatasetPatchDatasetRequest converts the update into
// the final JSON request body.
func marshalUpdateDatasetPatchDatasetRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"datasetReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"datasetReference", "projectId"},
	)
	return json.Marshal(m)
}

type updateDatasetPatchDatasetOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDatasetPatchDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {
	_, err := c.GetDataset(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchDataset")
	if err != nil {
		return err
	}

	req, err := newUpdateDatasetPatchDatasetRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDatasetPatchDatasetRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDatasetRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := datasetListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DatasetMaxPage {
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

type listDatasetOperation struct {
	Datasets []map[string]interface{} `json:"datasets"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listDataset(ctx context.Context, project, pageToken string, pageSize int32) ([]*Dataset, string, error) {
	b, err := c.listDatasetRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDatasetOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Dataset
	for _, v := range m.Datasets {
		res, err := unmarshalMapDataset(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDataset(ctx context.Context, f func(*Dataset) bool, resources []*Dataset) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDataset(ctx, res)
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

type deleteDatasetOperation struct{}

func (op *deleteDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {

	_, err := c.GetDataset(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Dataset not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDataset checking for existence. error: %v", err)
		return err
	}

	u, err := datasetDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Dataset: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetDataset(ctx, r.urlNormalized())
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
type createDatasetOperation struct {
	response map[string]interface{}
}

func (op *createDatasetOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := datasetCreateURL(c.Config.BasePath, project)

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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetDataset(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDatasetRaw(ctx context.Context, r *Dataset) ([]byte, error) {

	u, err := datasetGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) datasetDiffsForRawDesired(ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (initial, desired *Dataset, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Dataset
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Dataset); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Dataset, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDataset(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Dataset resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Dataset resource: %v", err)
		}
		c.Config.Logger.Info("Found that Dataset resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDatasetDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Dataset: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Dataset: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDatasetInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Dataset: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDatasetDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Dataset: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDataset(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDatasetInitialState(rawInitial, rawDesired *Dataset) (*Dataset, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDatasetDesiredState(rawDesired, rawInitial *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultEncryptionConfiguration = canonicalizeDatasetDefaultEncryptionConfiguration(rawDesired.DefaultEncryptionConfiguration, nil, opts...)

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.FriendlyName, rawInitial.FriendlyName) {
		rawDesired.FriendlyName = rawInitial.FriendlyName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.DefaultTableExpirationMs, rawInitial.DefaultTableExpirationMs) {
		rawDesired.DefaultTableExpirationMs = rawInitial.DefaultTableExpirationMs
	}
	if dcl.StringCanonicalize(rawDesired.DefaultPartitionExpirationMs, rawInitial.DefaultPartitionExpirationMs) {
		rawDesired.DefaultPartitionExpirationMs = rawInitial.DefaultPartitionExpirationMs
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.Access) {
		rawDesired.Access = rawInitial.Access
	}
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.BoolCanonicalize(rawDesired.Published, rawInitial.Published) {
		rawDesired.Published = rawInitial.Published
	}
	rawDesired.DefaultEncryptionConfiguration = canonicalizeDatasetDefaultEncryptionConfiguration(rawDesired.DefaultEncryptionConfiguration, rawInitial.DefaultEncryptionConfiguration, opts...)

	return rawDesired, nil
}

func canonicalizeDatasetNewState(c *Client, rawNew, rawDesired *Dataset) (*Dataset, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.FriendlyName) && dcl.IsEmptyValueIndirect(rawDesired.FriendlyName) {
		rawNew.FriendlyName = rawDesired.FriendlyName
	} else {
		if dcl.StringCanonicalize(rawDesired.FriendlyName, rawNew.FriendlyName) {
			rawNew.FriendlyName = rawDesired.FriendlyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultTableExpirationMs) && dcl.IsEmptyValueIndirect(rawDesired.DefaultTableExpirationMs) {
		rawNew.DefaultTableExpirationMs = rawDesired.DefaultTableExpirationMs
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultTableExpirationMs, rawNew.DefaultTableExpirationMs) {
			rawNew.DefaultTableExpirationMs = rawDesired.DefaultTableExpirationMs
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultPartitionExpirationMs) && dcl.IsEmptyValueIndirect(rawDesired.DefaultPartitionExpirationMs) {
		rawNew.DefaultPartitionExpirationMs = rawDesired.DefaultPartitionExpirationMs
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultPartitionExpirationMs, rawNew.DefaultPartitionExpirationMs) {
			rawNew.DefaultPartitionExpirationMs = rawDesired.DefaultPartitionExpirationMs
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Access) && dcl.IsEmptyValueIndirect(rawDesired.Access) {
		rawNew.Access = rawDesired.Access
	} else {
		rawNew.Access = canonicalizeNewDatasetAccessSet(c, rawDesired.Access, rawNew.Access)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTime) && dcl.IsEmptyValueIndirect(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedTime) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedTime) {
		rawNew.LastModifiedTime = rawDesired.LastModifiedTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Published) && dcl.IsEmptyValueIndirect(rawDesired.Published) {
		rawNew.Published = rawDesired.Published
	} else {
		if dcl.BoolCanonicalize(rawDesired.Published, rawNew.Published) {
			rawNew.Published = rawDesired.Published
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultEncryptionConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.DefaultEncryptionConfiguration) {
		rawNew.DefaultEncryptionConfiguration = rawDesired.DefaultEncryptionConfiguration
	} else {
		rawNew.DefaultEncryptionConfiguration = canonicalizeNewDatasetDefaultEncryptionConfiguration(c, rawDesired.DefaultEncryptionConfiguration, rawNew.DefaultEncryptionConfiguration)
	}

	return rawNew, nil
}

func canonicalizeDatasetAccess(des, initial *DatasetAccess, opts ...dcl.ApplyOption) *DatasetAccess {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if canonicalizeDatasetAccessRole(des.Role, initial.Role) || dcl.IsZeroValue(des.Role) {
		des.Role = initial.Role
	}
	if dcl.StringCanonicalize(des.UserByEmail, initial.UserByEmail) || dcl.IsZeroValue(des.UserByEmail) {
		des.UserByEmail = initial.UserByEmail
	}
	if dcl.StringCanonicalize(des.GroupByEmail, initial.GroupByEmail) || dcl.IsZeroValue(des.GroupByEmail) {
		des.GroupByEmail = initial.GroupByEmail
	}
	if dcl.StringCanonicalize(des.Domain, initial.Domain) || dcl.IsZeroValue(des.Domain) {
		des.Domain = initial.Domain
	}
	if dcl.StringCanonicalize(des.SpecialGroup, initial.SpecialGroup) || dcl.IsZeroValue(des.SpecialGroup) {
		des.SpecialGroup = initial.SpecialGroup
	}
	if dcl.StringCanonicalize(des.IamMember, initial.IamMember) || dcl.IsZeroValue(des.IamMember) {
		des.IamMember = initial.IamMember
	}
	des.View = canonicalizeDatasetAccessView(des.View, initial.View, opts...)
	des.Routine = canonicalizeDatasetAccessRoutine(des.Routine, initial.Routine, opts...)

	return des
}

func canonicalizeNewDatasetAccess(c *Client, des, nw *DatasetAccess) *DatasetAccess {
	if des == nil || nw == nil {
		return nw
	}

	if canonicalizeDatasetAccessRole(des.Role, nw.Role) {
		nw.Role = des.Role
	}
	if dcl.StringCanonicalize(des.UserByEmail, nw.UserByEmail) {
		nw.UserByEmail = des.UserByEmail
	}
	if dcl.StringCanonicalize(des.GroupByEmail, nw.GroupByEmail) {
		nw.GroupByEmail = des.GroupByEmail
	}
	if dcl.StringCanonicalize(des.Domain, nw.Domain) {
		nw.Domain = des.Domain
	}
	if dcl.StringCanonicalize(des.SpecialGroup, nw.SpecialGroup) {
		nw.SpecialGroup = des.SpecialGroup
	}
	if dcl.StringCanonicalize(des.IamMember, nw.IamMember) {
		nw.IamMember = des.IamMember
	}
	nw.View = canonicalizeNewDatasetAccessView(c, des.View, nw.View)
	nw.Routine = canonicalizeNewDatasetAccessRoutine(c, des.Routine, nw.Routine)

	return nw
}

func canonicalizeNewDatasetAccessSet(c *Client, des, nw []DatasetAccess) []DatasetAccess {
	if des == nil {
		return nw
	}
	var reorderedNew []DatasetAccess
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDatasetAccessNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDatasetAccessSlice(c *Client, des, nw []DatasetAccess) []DatasetAccess {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccess
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccess(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetAccessView(des, initial *DatasetAccessView, opts ...dcl.ApplyOption) *DatasetAccessView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		des.ProjectId = initial.ProjectId
	}
	if dcl.NameToSelfLink(des.DatasetId, initial.DatasetId) || dcl.IsZeroValue(des.DatasetId) {
		des.DatasetId = initial.DatasetId
	}
	if dcl.NameToSelfLink(des.TableId, initial.TableId) || dcl.IsZeroValue(des.TableId) {
		des.TableId = initial.TableId
	}

	return des
}

func canonicalizeNewDatasetAccessView(c *Client, des, nw *DatasetAccessView) *DatasetAccessView {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.NameToSelfLink(des.DatasetId, nw.DatasetId) {
		nw.DatasetId = des.DatasetId
	}
	if dcl.NameToSelfLink(des.TableId, nw.TableId) {
		nw.TableId = des.TableId
	}

	return nw
}

func canonicalizeNewDatasetAccessViewSet(c *Client, des, nw []DatasetAccessView) []DatasetAccessView {
	if des == nil {
		return nw
	}
	var reorderedNew []DatasetAccessView
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDatasetAccessViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDatasetAccessViewSlice(c *Client, des, nw []DatasetAccessView) []DatasetAccessView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccessView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccessView(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetAccessRoutine(des, initial *DatasetAccessRoutine, opts ...dcl.ApplyOption) *DatasetAccessRoutine {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		des.ProjectId = initial.ProjectId
	}
	if dcl.NameToSelfLink(des.DatasetId, initial.DatasetId) || dcl.IsZeroValue(des.DatasetId) {
		des.DatasetId = initial.DatasetId
	}
	if dcl.NameToSelfLink(des.RoutineId, initial.RoutineId) || dcl.IsZeroValue(des.RoutineId) {
		des.RoutineId = initial.RoutineId
	}

	return des
}

func canonicalizeNewDatasetAccessRoutine(c *Client, des, nw *DatasetAccessRoutine) *DatasetAccessRoutine {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.NameToSelfLink(des.DatasetId, nw.DatasetId) {
		nw.DatasetId = des.DatasetId
	}
	if dcl.NameToSelfLink(des.RoutineId, nw.RoutineId) {
		nw.RoutineId = des.RoutineId
	}

	return nw
}

func canonicalizeNewDatasetAccessRoutineSet(c *Client, des, nw []DatasetAccessRoutine) []DatasetAccessRoutine {
	if des == nil {
		return nw
	}
	var reorderedNew []DatasetAccessRoutine
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDatasetAccessRoutineNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDatasetAccessRoutineSlice(c *Client, des, nw []DatasetAccessRoutine) []DatasetAccessRoutine {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccessRoutine
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccessRoutine(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetDefaultEncryptionConfiguration(des, initial *DatasetDefaultEncryptionConfiguration, opts ...dcl.ApplyOption) *DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}

	return des
}

func canonicalizeNewDatasetDefaultEncryptionConfiguration(c *Client, des, nw *DatasetDefaultEncryptionConfiguration) *DatasetDefaultEncryptionConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}

	return nw
}

func canonicalizeNewDatasetDefaultEncryptionConfigurationSet(c *Client, des, nw []DatasetDefaultEncryptionConfiguration) []DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []DatasetDefaultEncryptionConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDatasetDefaultEncryptionConfigurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDatasetDefaultEncryptionConfigurationSlice(c *Client, des, nw []DatasetDefaultEncryptionConfiguration) []DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetDefaultEncryptionConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetDefaultEncryptionConfiguration(c, &d, &n))
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
func diffDataset(c *Client, desired, actual *Dataset, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FriendlyName, actual.FriendlyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("FriendlyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultTableExpirationMs, actual.DefaultTableExpirationMs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("DefaultTableExpirationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultPartitionExpirationMs, actual.DefaultPartitionExpirationMs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("DefaultPartitionExpirationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Access, actual.Access, dcl.Info{Type: "Set", ObjectFunction: compareDatasetAccessNewStyle, EmptyObject: EmptyDatasetAccess, OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Access")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedTime, actual.LastModifiedTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Published, actual.Published, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Published")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultEncryptionConfiguration, actual.DefaultEncryptionConfiguration, dcl.Info{ObjectFunction: compareDatasetDefaultEncryptionConfigurationNewStyle, EmptyObject: EmptyDatasetDefaultEncryptionConfiguration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultEncryptionConfiguration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareDatasetAccessNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccess)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccess or *DatasetAccess", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccess)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccess", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Role, actual.Role, dcl.Info{CustomDiff: canonicalizeDatasetAccessRole, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Role")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserByEmail, actual.UserByEmail, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserByEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByEmail, actual.GroupByEmail, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Domain, actual.Domain, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Domain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SpecialGroup, actual.SpecialGroup, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SpecialGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IamMember, actual.IamMember, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IamMember")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.View, actual.View, dcl.Info{ObjectFunction: compareDatasetAccessViewNewStyle, EmptyObject: EmptyDatasetAccessView, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("View")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Routine, actual.Routine, dcl.Info{ObjectFunction: compareDatasetAccessRoutineNewStyle, EmptyObject: EmptyDatasetAccessRoutine, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Routine")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetAccessViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccessView)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccessView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessView or *DatasetAccessView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccessView)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccessView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableId, actual.TableId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TableId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetAccessRoutineNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccessRoutine)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccessRoutine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessRoutine or *DatasetAccessRoutine", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccessRoutine)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccessRoutine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessRoutine", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RoutineId, actual.RoutineId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RoutineId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetDefaultEncryptionConfigurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetDefaultEncryptionConfiguration)
	if !ok {
		desiredNotPointer, ok := d.(DatasetDefaultEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetDefaultEncryptionConfiguration or *DatasetDefaultEncryptionConfiguration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetDefaultEncryptionConfiguration)
	if !ok {
		actualNotPointer, ok := a.(DatasetDefaultEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetDefaultEncryptionConfiguration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
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
func (r *Dataset) urlNormalized() *Dataset {
	normalized := dcl.Copy(*r).(Dataset)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.FriendlyName = dcl.SelfLinkToName(r.FriendlyName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DefaultTableExpirationMs = dcl.SelfLinkToName(r.DefaultTableExpirationMs)
	normalized.DefaultPartitionExpirationMs = dcl.SelfLinkToName(r.DefaultPartitionExpirationMs)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Dataset) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Dataset) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Dataset) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Dataset) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "PatchDataset" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/datasets/{{name}}", "https://bigquery.googleapis.com/bigquery/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Dataset resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Dataset) marshal(c *Client) ([]byte, error) {
	m, err := expandDataset(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Dataset: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"datasetReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"datasetReference", "projectId"},
	)

	return json.Marshal(m)
}

// unmarshalDataset decodes JSON responses into the Dataset resource schema.
func unmarshalDataset(b []byte, c *Client) (*Dataset, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDataset(m, c)
}

func unmarshalMapDataset(m map[string]interface{}, c *Client) (*Dataset, error) {
	dcl.MoveMapEntry(
		m,
		[]string{"datasetReference", "datasetId"},
		[]string{"name"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"datasetReference", "projectId"},
		[]string{"project"},
	)

	return flattenDataset(c, m), nil
}

// expandDataset expands Dataset into a JSON request object.
func expandDataset(c *Client, f *Dataset) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		m["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DefaultTableExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		m["defaultTableExpirationMs"] = v
	}
	if v := f.DefaultPartitionExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		m["defaultPartitionExpirationMs"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandDatasetAccessSlice(c, f.Access); err != nil {
		return nil, fmt.Errorf("error expanding Access into access: %w", err)
	} else if v != nil {
		m["access"] = v
	}
	if v := f.CreationTime; !dcl.IsEmptyValueIndirect(v) {
		m["creationTime"] = v
	}
	if v := f.LastModifiedTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastModifiedTime"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.Published; !dcl.IsEmptyValueIndirect(v) {
		m["published"] = v
	}
	if v, err := expandDatasetDefaultEncryptionConfiguration(c, f.DefaultEncryptionConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding DefaultEncryptionConfiguration into defaultEncryptionConfiguration: %w", err)
	} else if v != nil {
		m["defaultEncryptionConfiguration"] = v
	}

	return m, nil
}

// flattenDataset flattens Dataset from a JSON request object into the
// Dataset type.
func flattenDataset(c *Client, i interface{}) *Dataset {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Dataset{}
	res.Etag = dcl.FlattenString(m["etag"])
	res.Id = dcl.FlattenString(m["id"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Name = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.FriendlyName = dcl.FlattenString(m["friendlyName"])
	res.Description = dcl.FlattenString(m["description"])
	res.DefaultTableExpirationMs = dcl.FlattenString(m["defaultTableExpirationMs"])
	res.DefaultPartitionExpirationMs = dcl.FlattenString(m["defaultPartitionExpirationMs"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Access = flattenDatasetAccessSlice(c, m["access"])
	res.CreationTime = dcl.FlattenInteger(m["creationTime"])
	res.LastModifiedTime = dcl.FlattenInteger(m["lastModifiedTime"])
	res.Location = dcl.FlattenString(m["location"])
	res.Published = dcl.FlattenBool(m["published"])
	res.DefaultEncryptionConfiguration = flattenDatasetDefaultEncryptionConfiguration(c, m["defaultEncryptionConfiguration"])

	return res
}

// expandDatasetAccessMap expands the contents of DatasetAccess into a JSON
// request object.
func expandDatasetAccessMap(c *Client, f map[string]DatasetAccess) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccess(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessSlice expands the contents of DatasetAccess into a JSON
// request object.
func expandDatasetAccessSlice(c *Client, f []DatasetAccess) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccess(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessMap flattens the contents of DatasetAccess from a JSON
// response object.
func flattenDatasetAccessMap(c *Client, i interface{}) map[string]DatasetAccess {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccess{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccess{}
	}

	items := make(map[string]DatasetAccess)
	for k, item := range a {
		items[k] = *flattenDatasetAccess(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDatasetAccessSlice flattens the contents of DatasetAccess from a JSON
// response object.
func flattenDatasetAccessSlice(c *Client, i interface{}) []DatasetAccess {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccess{}
	}

	if len(a) == 0 {
		return []DatasetAccess{}
	}

	items := make([]DatasetAccess, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccess(c, item.(map[string]interface{})))
	}

	return items
}

// expandDatasetAccess expands an instance of DatasetAccess into a JSON
// request object.
func expandDatasetAccess(c *Client, f *DatasetAccess) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		m["role"] = v
	}
	if v := f.UserByEmail; !dcl.IsEmptyValueIndirect(v) {
		m["userByEmail"] = v
	}
	if v := f.GroupByEmail; !dcl.IsEmptyValueIndirect(v) {
		m["groupByEmail"] = v
	}
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}
	if v := f.SpecialGroup; !dcl.IsEmptyValueIndirect(v) {
		m["specialGroup"] = v
	}
	if v := f.IamMember; !dcl.IsEmptyValueIndirect(v) {
		m["iamMember"] = v
	}
	if v, err := expandDatasetAccessView(c, f.View); err != nil {
		return nil, fmt.Errorf("error expanding View into view: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["view"] = v
	}
	if v, err := expandDatasetAccessRoutine(c, f.Routine); err != nil {
		return nil, fmt.Errorf("error expanding Routine into routine: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["routine"] = v
	}

	return m, nil
}

// flattenDatasetAccess flattens an instance of DatasetAccess from a JSON
// response object.
func flattenDatasetAccess(c *Client, i interface{}) *DatasetAccess {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccess{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccess
	}
	r.Role = dcl.FlattenString(m["role"])
	r.UserByEmail = dcl.FlattenString(m["userByEmail"])
	r.GroupByEmail = dcl.FlattenString(m["groupByEmail"])
	r.Domain = dcl.FlattenString(m["domain"])
	r.SpecialGroup = dcl.FlattenString(m["specialGroup"])
	r.IamMember = dcl.FlattenString(m["iamMember"])
	r.View = flattenDatasetAccessView(c, m["view"])
	r.Routine = flattenDatasetAccessRoutine(c, m["routine"])

	return r
}

// expandDatasetAccessViewMap expands the contents of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessViewMap(c *Client, f map[string]DatasetAccessView) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccessView(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessViewSlice expands the contents of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessViewSlice(c *Client, f []DatasetAccessView) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccessView(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessViewMap flattens the contents of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessViewMap(c *Client, i interface{}) map[string]DatasetAccessView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccessView{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccessView{}
	}

	items := make(map[string]DatasetAccessView)
	for k, item := range a {
		items[k] = *flattenDatasetAccessView(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDatasetAccessViewSlice flattens the contents of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessViewSlice(c *Client, i interface{}) []DatasetAccessView {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccessView{}
	}

	if len(a) == 0 {
		return []DatasetAccessView{}
	}

	items := make([]DatasetAccessView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccessView(c, item.(map[string]interface{})))
	}

	return items
}

// expandDatasetAccessView expands an instance of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessView(c *Client, f *DatasetAccessView) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.DatasetId; !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v := f.TableId; !dcl.IsEmptyValueIndirect(v) {
		m["tableId"] = v
	}

	return m, nil
}

// flattenDatasetAccessView flattens an instance of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessView(c *Client, i interface{}) *DatasetAccessView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccessView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccessView
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.TableId = dcl.FlattenString(m["tableId"])

	return r
}

// expandDatasetAccessRoutineMap expands the contents of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutineMap(c *Client, f map[string]DatasetAccessRoutine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccessRoutine(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessRoutineSlice expands the contents of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutineSlice(c *Client, f []DatasetAccessRoutine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccessRoutine(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessRoutineMap flattens the contents of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutineMap(c *Client, i interface{}) map[string]DatasetAccessRoutine {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccessRoutine{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccessRoutine{}
	}

	items := make(map[string]DatasetAccessRoutine)
	for k, item := range a {
		items[k] = *flattenDatasetAccessRoutine(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDatasetAccessRoutineSlice flattens the contents of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutineSlice(c *Client, i interface{}) []DatasetAccessRoutine {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccessRoutine{}
	}

	if len(a) == 0 {
		return []DatasetAccessRoutine{}
	}

	items := make([]DatasetAccessRoutine, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccessRoutine(c, item.(map[string]interface{})))
	}

	return items
}

// expandDatasetAccessRoutine expands an instance of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutine(c *Client, f *DatasetAccessRoutine) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.DatasetId; !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v := f.RoutineId; !dcl.IsEmptyValueIndirect(v) {
		m["routineId"] = v
	}

	return m, nil
}

// flattenDatasetAccessRoutine flattens an instance of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutine(c *Client, i interface{}) *DatasetAccessRoutine {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccessRoutine{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccessRoutine
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.RoutineId = dcl.FlattenString(m["routineId"])

	return r
}

// expandDatasetDefaultEncryptionConfigurationMap expands the contents of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfigurationMap(c *Client, f map[string]DatasetDefaultEncryptionConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetDefaultEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetDefaultEncryptionConfigurationSlice expands the contents of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfigurationSlice(c *Client, f []DatasetDefaultEncryptionConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetDefaultEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetDefaultEncryptionConfigurationMap flattens the contents of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfigurationMap(c *Client, i interface{}) map[string]DatasetDefaultEncryptionConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetDefaultEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return map[string]DatasetDefaultEncryptionConfiguration{}
	}

	items := make(map[string]DatasetDefaultEncryptionConfiguration)
	for k, item := range a {
		items[k] = *flattenDatasetDefaultEncryptionConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDatasetDefaultEncryptionConfigurationSlice flattens the contents of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfigurationSlice(c *Client, i interface{}) []DatasetDefaultEncryptionConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetDefaultEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return []DatasetDefaultEncryptionConfiguration{}
	}

	items := make([]DatasetDefaultEncryptionConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetDefaultEncryptionConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandDatasetDefaultEncryptionConfiguration expands an instance of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfiguration(c *Client, f *DatasetDefaultEncryptionConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenDatasetDefaultEncryptionConfiguration flattens an instance of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfiguration(c *Client, i interface{}) *DatasetDefaultEncryptionConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetDefaultEncryptionConfiguration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetDefaultEncryptionConfiguration
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Dataset) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDataset(b, c)
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

type datasetDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         datasetApiOperation
}

func convertFieldDiffToDatasetOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]datasetDiff, error) {
	var diffs []datasetDiff
	for _, op := range ops {
		diff := datasetDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTodatasetApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTodatasetApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (datasetApiOperation, error) {
	switch op {

	case "updateDatasetPatchDatasetOperation":
		return &updateDatasetPatchDatasetOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
