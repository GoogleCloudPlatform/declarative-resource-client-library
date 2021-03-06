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

func (r *TargetPool) validate() error {

	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	return nil
}

func targetPoolGetURL(userBasePath string, r *TargetPool) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func targetPoolListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetPoolCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func targetPoolDeleteURL(userBasePath string, r *TargetPool) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// targetPoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type targetPoolApiOperation interface {
	do(context.Context, *TargetPool, *Client) error
}

// newUpdateTargetPoolAddHCRequest creates a request for an
// TargetPool resource's AddHC update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetPoolAddHCRequest(ctx context.Context, f *TargetPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveFieldArray("projects/%s/global/httpHealthChecks/%s", f.HealthChecks, f.Project, f.HealthChecks); err != nil {
		return nil, fmt.Errorf("error expanding HealthChecks into healthChecks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["healthChecks"] = v
	}
	return req, nil
}

// marshalUpdateTargetPoolAddHCRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetPoolAddHCRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = WrapTargetPoolHealthCheck(m)
	return json.Marshal(m)
}

type updateTargetPoolAddHCOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetPoolAddHCOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	_, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "AddHC")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetPoolAddHCRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetPoolAddHCRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

// newUpdateTargetPoolAddInstanceRequest creates a request for an
// TargetPool resource's AddInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetPoolAddInstanceRequest(ctx context.Context, f *TargetPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Instances; !dcl.IsEmptyValueIndirect(v) {
		req["instances"] = v
	}
	return req, nil
}

// marshalUpdateTargetPoolAddInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetPoolAddInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = WrapTargetPoolInstance(m)
	return json.Marshal(m)
}

type updateTargetPoolAddInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetPoolAddInstanceOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	_, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "AddInstance")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetPoolAddInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetPoolAddInstanceRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

// newUpdateTargetPoolRemoveHCRequest creates a request for an
// TargetPool resource's RemoveHC update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetPoolRemoveHCRequest(ctx context.Context, f *TargetPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveFieldArray("projects/%s/global/httpHealthChecks/%s", f.HealthChecks, f.Project, f.HealthChecks); err != nil {
		return nil, fmt.Errorf("error expanding HealthChecks into healthChecks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["healthChecks"] = v
	}
	return req, nil
}

// marshalUpdateTargetPoolRemoveHCRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetPoolRemoveHCRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = WrapTargetPoolHealthCheck(m)
	return json.Marshal(m)
}

type updateTargetPoolRemoveHCOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetPoolRemoveHCOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	_, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "RemoveHC")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetPoolRemoveHCRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetPoolRemoveHCRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

// newUpdateTargetPoolRemoveInstanceRequest creates a request for an
// TargetPool resource's RemoveInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetPoolRemoveInstanceRequest(ctx context.Context, f *TargetPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Instances; !dcl.IsEmptyValueIndirect(v) {
		req["instances"] = v
	}
	return req, nil
}

// marshalUpdateTargetPoolRemoveInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetPoolRemoveInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = WrapTargetPoolInstance(m)
	return json.Marshal(m)
}

type updateTargetPoolRemoveInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetPoolRemoveInstanceOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	_, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "RemoveInstance")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetPoolRemoveInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetPoolRemoveInstanceRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

// newUpdateTargetPoolSetBackupRequest creates a request for an
// TargetPool resource's SetBackup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetPoolSetBackupRequest(ctx context.Context, f *TargetPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.BackupPool; !dcl.IsEmptyValueIndirect(v) {
		req["backupPool"] = v
	}
	return req, nil
}

// marshalUpdateTargetPoolSetBackupRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetPoolSetBackupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTargetPoolSetBackupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetPoolSetBackupOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	_, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "SetBackup")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetPoolSetBackupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetPoolSetBackupRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listTargetPoolRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := targetPoolListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TargetPoolMaxPage {
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

type listTargetPoolOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listTargetPool(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*TargetPool, string, error) {
	b, err := c.listTargetPoolRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTargetPoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TargetPool
	for _, v := range m.Items {
		res, err := unmarshalMapTargetPool(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTargetPool(ctx context.Context, f func(*TargetPool) bool, resources []*TargetPool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTargetPool(ctx, res)
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

type deleteTargetPoolOperation struct{}

func (op *deleteTargetPoolOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	r, err := c.GetTargetPool(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("TargetPool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTargetPool checking for existence. error: %v", err)
		return err
	}

	u, err := targetPoolDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetTargetPool(ctx, r.URLNormalized())
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
type createTargetPoolOperation struct {
	response map[string]interface{}
}

func (op *createTargetPoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTargetPoolOperation) do(ctx context.Context, r *TargetPool, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := targetPoolCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetTargetPool(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTargetPoolRaw(ctx context.Context, r *TargetPool) ([]byte, error) {

	u, err := targetPoolGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) targetPoolDiffsForRawDesired(ctx context.Context, rawDesired *TargetPool, opts ...dcl.ApplyOption) (initial, desired *TargetPool, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TargetPool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TargetPool); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected TargetPool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTargetPool(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a TargetPool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TargetPool resource: %v", err)
		}
		c.Config.Logger.Info("Found that TargetPool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTargetPoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for TargetPool: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for TargetPool: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTargetPoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for TargetPool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTargetPoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for TargetPool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTargetPool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTargetPoolInitialState(rawInitial, rawDesired *TargetPool) (*TargetPool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTargetPoolDesiredState(rawDesired, rawInitial *TargetPool, opts ...dcl.ApplyOption) (*TargetPool, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &TargetPool{}
	if dcl.StringCanonicalize(rawDesired.BackupPool, rawInitial.BackupPool) {
		canonicalDesired.BackupPool = rawInitial.BackupPool
	} else {
		canonicalDesired.BackupPool = rawDesired.BackupPool
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.FailoverRatio) {
		canonicalDesired.FailoverRatio = rawInitial.FailoverRatio
	} else {
		canonicalDesired.FailoverRatio = rawDesired.FailoverRatio
	}
	if dcl.PartialSelfLinkToSelfLinkArray(rawDesired.HealthChecks, rawInitial.HealthChecks) {
		canonicalDesired.HealthChecks = rawInitial.HealthChecks
	} else {
		canonicalDesired.HealthChecks = rawDesired.HealthChecks
	}
	if dcl.IsZeroValue(rawDesired.Instances) {
		canonicalDesired.Instances = rawInitial.Instances
	} else {
		canonicalDesired.Instances = rawDesired.Instances
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		canonicalDesired.Region = rawInitial.Region
	} else {
		canonicalDesired.Region = rawDesired.Region
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		canonicalDesired.SelfLink = rawInitial.SelfLink
	} else {
		canonicalDesired.SelfLink = rawDesired.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.SessionAffinity) {
		canonicalDesired.SessionAffinity = rawInitial.SessionAffinity
	} else {
		canonicalDesired.SessionAffinity = rawDesired.SessionAffinity
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeTargetPoolNewState(c *Client, rawNew, rawDesired *TargetPool) (*TargetPool, error) {

	if dcl.IsEmptyValueIndirect(rawNew.BackupPool) && dcl.IsEmptyValueIndirect(rawDesired.BackupPool) {
		rawNew.BackupPool = rawDesired.BackupPool
	} else {
		if dcl.StringCanonicalize(rawDesired.BackupPool, rawNew.BackupPool) {
			rawNew.BackupPool = rawDesired.BackupPool
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailoverRatio) && dcl.IsEmptyValueIndirect(rawDesired.FailoverRatio) {
		rawNew.FailoverRatio = rawDesired.FailoverRatio
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthChecks) && dcl.IsEmptyValueIndirect(rawDesired.HealthChecks) {
		rawNew.HealthChecks = rawDesired.HealthChecks
	} else {
		if dcl.PartialSelfLinkToSelfLinkArray(rawDesired.HealthChecks, rawNew.HealthChecks) {
			rawNew.HealthChecks = rawDesired.HealthChecks
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Instances) && dcl.IsEmptyValueIndirect(rawDesired.Instances) {
		rawNew.Instances = rawDesired.Instances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SessionAffinity) && dcl.IsEmptyValueIndirect(rawDesired.SessionAffinity) {
		rawNew.SessionAffinity = rawDesired.SessionAffinity
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTargetPool(c *Client, desired, actual *TargetPool, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.BackupPool, actual.BackupPool, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTargetPoolSetBackupOperation")}, fn.AddNest("BackupPool")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.FailoverRatio, actual.FailoverRatio, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailoverRatio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthChecks, actual.HealthChecks, dcl.Info{Type: "Set", OperationSelector: targetPoolHealthCheck()}, fn.AddNest("HealthChecks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Instances, actual.Instances, dcl.Info{Type: "Set", OperationSelector: targetPoolInstances()}, fn.AddNest("Instances")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SessionAffinity, actual.SessionAffinity, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SessionAffinity")); len(ds) != 0 || err != nil {
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

func (r *TargetPool) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetPool) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *TargetPool) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetPool) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "AddHC" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}/addHealthCheck", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "AddInstance" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}/addInstance", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "RemoveHC" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}/removeHealthCheck", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "RemoveInstance" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}/removeInstance", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "SetBackup" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/targetPools/{{name}}/setBackup", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TargetPool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TargetPool) marshal(c *Client) ([]byte, error) {
	m, err := expandTargetPool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TargetPool: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTargetPool decodes JSON responses into the TargetPool resource schema.
func unmarshalTargetPool(b []byte, c *Client) (*TargetPool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTargetPool(m, c)
}

func unmarshalMapTargetPool(m map[string]interface{}, c *Client) (*TargetPool, error) {

	return flattenTargetPool(c, m), nil
}

// expandTargetPool expands TargetPool into a JSON request object.
func expandTargetPool(c *Client, f *TargetPool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.BackupPool; !dcl.IsEmptyValueIndirect(v) {
		m["backupPool"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.FailoverRatio; !dcl.IsEmptyValueIndirect(v) {
		m["failoverRatio"] = v
	}
	if v, err := dcl.DeriveFieldArray("projects/%s/global/httpHealthChecks/%s", f.HealthChecks, f.Project, f.HealthChecks); err != nil {
		return nil, fmt.Errorf("error expanding HealthChecks into healthChecks: %w", err)
	} else if v != nil {
		m["healthChecks"] = v
	}
	if v := f.Instances; !dcl.IsEmptyValueIndirect(v) {
		m["instances"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.SessionAffinity; !dcl.IsEmptyValueIndirect(v) {
		m["sessionAffinity"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenTargetPool flattens TargetPool from a JSON request object into the
// TargetPool type.
func flattenTargetPool(c *Client, i interface{}) *TargetPool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &TargetPool{}
	res.BackupPool = dcl.FlattenString(m["backupPool"])
	res.Description = dcl.FlattenString(m["description"])
	res.FailoverRatio = dcl.FlattenDouble(m["failoverRatio"])
	res.HealthChecks = dcl.FlattenStringSlice(m["healthChecks"])
	res.Instances = dcl.FlattenStringSlice(m["instances"])
	res.Name = dcl.FlattenString(m["name"])
	res.Region = dcl.FlattenString(m["region"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.SessionAffinity = flattenTargetPoolSessionAffinityEnum(m["sessionAffinity"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// flattenTargetPoolSessionAffinityEnumMap flattens the contents of TargetPoolSessionAffinityEnum from a JSON
// response object.
func flattenTargetPoolSessionAffinityEnumMap(c *Client, i interface{}) map[string]TargetPoolSessionAffinityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetPoolSessionAffinityEnum{}
	}

	if len(a) == 0 {
		return map[string]TargetPoolSessionAffinityEnum{}
	}

	items := make(map[string]TargetPoolSessionAffinityEnum)
	for k, item := range a {
		items[k] = *flattenTargetPoolSessionAffinityEnum(item.(interface{}))
	}

	return items
}

// flattenTargetPoolSessionAffinityEnumSlice flattens the contents of TargetPoolSessionAffinityEnum from a JSON
// response object.
func flattenTargetPoolSessionAffinityEnumSlice(c *Client, i interface{}) []TargetPoolSessionAffinityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetPoolSessionAffinityEnum{}
	}

	if len(a) == 0 {
		return []TargetPoolSessionAffinityEnum{}
	}

	items := make([]TargetPoolSessionAffinityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetPoolSessionAffinityEnum(item.(interface{})))
	}

	return items
}

// flattenTargetPoolSessionAffinityEnum asserts that an interface is a string, and returns a
// pointer to a *TargetPoolSessionAffinityEnum with the same value as that string.
func flattenTargetPoolSessionAffinityEnum(i interface{}) *TargetPoolSessionAffinityEnum {
	s, ok := i.(string)
	if !ok {
		return TargetPoolSessionAffinityEnumRef("")
	}

	return TargetPoolSessionAffinityEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TargetPool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTargetPool(b, c)
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

type targetPoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         targetPoolApiOperation
}

func convertFieldDiffsToTargetPoolDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]targetPoolDiff, error) {
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
	var diffs []targetPoolDiff
	// For each operation name, create a targetPoolDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := targetPoolDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTargetPoolApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTargetPoolApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (targetPoolApiOperation, error) {
	switch opName {

	case "updateTargetPoolAddHCOperation":
		return &updateTargetPoolAddHCOperation{FieldDiffs: fieldDiffs}, nil

	case "updateTargetPoolAddInstanceOperation":
		return &updateTargetPoolAddInstanceOperation{FieldDiffs: fieldDiffs}, nil

	case "updateTargetPoolRemoveHCOperation":
		return &updateTargetPoolRemoveHCOperation{FieldDiffs: fieldDiffs}, nil

	case "updateTargetPoolRemoveInstanceOperation":
		return &updateTargetPoolRemoveInstanceOperation{FieldDiffs: fieldDiffs}, nil

	case "updateTargetPoolSetBackupOperation":
		return &updateTargetPoolSetBackupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
