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
package alpha

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

func (r *OsPolicyAssignment) validate() error {

	if err := dcl.Required(r, "osPolicies"); err != nil {
		return err
	}
	if err := dcl.Required(r, "instanceFilter"); err != nil {
		return err
	}
	if err := dcl.Required(r, "rollout"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.InstanceFilter) {
		if err := r.InstanceFilter.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Rollout) {
		if err := r.Rollout.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPolicies) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if err := dcl.Required(r, "mode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resourceGroups"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroups) validate() error {
	if err := dcl.Required(r, "resources"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OsFilter) {
		if err := r.OsFilter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) validate() error {
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResources) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Pkg) {
		if err := r.Pkg.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Repository) {
		if err := r.Repository.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Exec) {
		if err := r.Exec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.File) {
		if err := r.File.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) validate() error {
	if err := dcl.Required(r, "desiredState"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Apt) {
		if err := r.Apt.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deb) {
		if err := r.Deb.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Yum) {
		if err := r.Yum.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Zypper) {
		if err := r.Zypper.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Rpm) {
		if err := r.Rpm.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Googet) {
		if err := r.Googet.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Msi) {
		if err := r.Msi.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) validate() error {
	if err := dcl.Required(r, "source"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Source) {
		if err := r.Source.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Remote) {
		if err := r.Remote.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Gcs) {
		if err := r.Gcs.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) validate() error {
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) validate() error {
	if err := dcl.Required(r, "bucket"); err != nil {
		return err
	}
	if err := dcl.Required(r, "object"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) validate() error {
	if err := dcl.Required(r, "source"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Source) {
		if err := r.Source.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentFile) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Remote) {
		if err := r.Remote.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Gcs) {
		if err := r.Gcs.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentFileRemote) validate() error {
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentFileGcs) validate() error {
	if err := dcl.Required(r, "bucket"); err != nil {
		return err
	}
	if err := dcl.Required(r, "object"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) validate() error {
	if err := dcl.Required(r, "source"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Source) {
		if err := r.Source.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Apt) {
		if err := r.Apt.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Yum) {
		if err := r.Yum.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Zypper) {
		if err := r.Zypper.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Goo) {
		if err := r.Goo.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) validate() error {
	if err := dcl.Required(r, "archiveType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "distribution"); err != nil {
		return err
	}
	if err := dcl.Required(r, "components"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if err := dcl.Required(r, "baseUrl"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if err := dcl.Required(r, "baseUrl"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) validate() error {
	if err := dcl.Required(r, "validate"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Validate) {
		if err := r.Validate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Enforce) {
		if err := r.Enforce.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentExec) validate() error {
	if err := dcl.Required(r, "interpreter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.File) {
		if err := r.File.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	if err := dcl.Required(r, "state"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.File) {
		if err := r.File.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentInstanceFilter) validate() error {
	return nil
}
func (r *OsPolicyAssignmentInstanceFilterInclusionLabels) validate() error {
	return nil
}
func (r *OsPolicyAssignmentInstanceFilterExclusionLabels) validate() error {
	return nil
}
func (r *OsPolicyAssignmentRollout) validate() error {
	if err := dcl.Required(r, "disruptionBudget"); err != nil {
		return err
	}
	if err := dcl.Required(r, "minWaitDuration"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DisruptionBudget) {
		if err := r.DisruptionBudget.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OsPolicyAssignmentRolloutDisruptionBudget) validate() error {
	return nil
}

func osPolicyAssignmentGetURL(userBasePath string, r *OsPolicyAssignment) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/osPolicyAssignments/{{name}}", "https://osconfig.googleapis.com/v1alpha/", userBasePath, params), nil
}

func osPolicyAssignmentListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/osPolicyAssignments", "https://osconfig.googleapis.com/v1alpha/", userBasePath, params), nil

}

func osPolicyAssignmentCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/osPolicyAssignments?osPolicyAssignmentId={{name}}", "https://osconfig.googleapis.com/v1alpha/", userBasePath, params), nil

}

func osPolicyAssignmentDeleteURL(userBasePath string, r *OsPolicyAssignment) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/osPolicyAssignments/{{name}}", "https://osconfig.googleapis.com/v1alpha/", userBasePath, params), nil
}

// osPolicyAssignmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type osPolicyAssignmentApiOperation interface {
	do(context.Context, *OsPolicyAssignment, *Client) error
}

// newUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest creates a request for an
// OsPolicyAssignment resource's UpdateOSPolicyAssignment update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest(ctx context.Context, f *OsPolicyAssignment, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesSlice(c, f.OsPolicies); err != nil {
		return nil, fmt.Errorf("error expanding OsPolicies into osPolicies: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["osPolicies"] = v
	}
	if v, err := expandOsPolicyAssignmentInstanceFilter(c, f.InstanceFilter); err != nil {
		return nil, fmt.Errorf("error expanding InstanceFilter into instanceFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["instanceFilter"] = v
	}
	if v, err := expandOsPolicyAssignmentRollout(c, f.Rollout); err != nil {
		return nil, fmt.Errorf("error expanding Rollout into rollout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["rollout"] = v
	}
	return req, nil
}

// marshalUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest converts the update into
// the final JSON request body.
func marshalUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation) do(ctx context.Context, r *OsPolicyAssignment, c *Client) error {
	_, err := c.GetOsPolicyAssignment(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateOSPolicyAssignment")
	if err != nil {
		return err
	}
	mask := dcl.TopLevelUpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateOsPolicyAssignmentUpdateOSPolicyAssignmentRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://osconfig.googleapis.com/v1alpha/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listOsPolicyAssignmentRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := osPolicyAssignmentListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != OsPolicyAssignmentMaxPage {
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

type listOsPolicyAssignmentOperation struct {
	OsPolicyAssignments []map[string]interface{} `json:"osPolicyAssignments"`
	Token               string                   `json:"nextPageToken"`
}

func (c *Client) listOsPolicyAssignment(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*OsPolicyAssignment, string, error) {
	b, err := c.listOsPolicyAssignmentRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listOsPolicyAssignmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*OsPolicyAssignment
	for _, v := range m.OsPolicyAssignments {
		res, err := unmarshalMapOsPolicyAssignment(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllOsPolicyAssignment(ctx context.Context, f func(*OsPolicyAssignment) bool, resources []*OsPolicyAssignment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteOsPolicyAssignment(ctx, res)
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

type deleteOsPolicyAssignmentOperation struct{}

func (op *deleteOsPolicyAssignmentOperation) do(ctx context.Context, r *OsPolicyAssignment, c *Client) error {
	r, err := c.GetOsPolicyAssignment(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("OsPolicyAssignment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetOsPolicyAssignment checking for existence. error: %v", err)
		return err
	}

	err = r.waitForNotReconciling(ctx, c)
	if err != nil {
		return err
	}
	u, err := osPolicyAssignmentDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	var o operations.OSPolicyAssignmentDeleteOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://osconfig.googleapis.com/v1alpha/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetOsPolicyAssignment(ctx, r.URLNormalized())
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
type createOsPolicyAssignmentOperation struct {
	response map[string]interface{}
}

func (op *createOsPolicyAssignmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createOsPolicyAssignmentOperation) do(ctx context.Context, r *OsPolicyAssignment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := osPolicyAssignmentCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://osconfig.googleapis.com/v1alpha/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetOsPolicyAssignment(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getOsPolicyAssignmentRaw(ctx context.Context, r *OsPolicyAssignment) ([]byte, error) {

	u, err := osPolicyAssignmentGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) osPolicyAssignmentDiffsForRawDesired(ctx context.Context, rawDesired *OsPolicyAssignment, opts ...dcl.ApplyOption) (initial, desired *OsPolicyAssignment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *OsPolicyAssignment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*OsPolicyAssignment); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected OsPolicyAssignment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetOsPolicyAssignment(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a OsPolicyAssignment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve OsPolicyAssignment resource: %v", err)
		}
		c.Config.Logger.Info("Found that OsPolicyAssignment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeOsPolicyAssignmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for OsPolicyAssignment: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for OsPolicyAssignment: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeOsPolicyAssignmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for OsPolicyAssignment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeOsPolicyAssignmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for OsPolicyAssignment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffOsPolicyAssignment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeOsPolicyAssignmentInitialState(rawInitial, rawDesired *OsPolicyAssignment) (*OsPolicyAssignment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeOsPolicyAssignmentDesiredState(rawDesired, rawInitial *OsPolicyAssignment, opts ...dcl.ApplyOption) (*OsPolicyAssignment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.InstanceFilter = canonicalizeOsPolicyAssignmentInstanceFilter(rawDesired.InstanceFilter, nil, opts...)
		rawDesired.Rollout = canonicalizeOsPolicyAssignmentRollout(rawDesired.Rollout, nil, opts...)

		return rawDesired, nil
	}

	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.OsPolicies) {
		rawDesired.OsPolicies = rawInitial.OsPolicies
	}
	rawDesired.InstanceFilter = canonicalizeOsPolicyAssignmentInstanceFilter(rawDesired.InstanceFilter, rawInitial.InstanceFilter, opts...)
	rawDesired.Rollout = canonicalizeOsPolicyAssignmentRollout(rawDesired.Rollout, rawInitial.Rollout, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeOsPolicyAssignmentNewState(c *Client, rawNew, rawDesired *OsPolicyAssignment) (*OsPolicyAssignment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.OsPolicies) && dcl.IsEmptyValueIndirect(rawDesired.OsPolicies) {
		rawNew.OsPolicies = rawDesired.OsPolicies
	} else {
		rawNew.OsPolicies = canonicalizeNewOsPolicyAssignmentOsPoliciesSlice(c, rawDesired.OsPolicies, rawNew.OsPolicies)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceFilter) && dcl.IsEmptyValueIndirect(rawDesired.InstanceFilter) {
		rawNew.InstanceFilter = rawDesired.InstanceFilter
	} else {
		rawNew.InstanceFilter = canonicalizeNewOsPolicyAssignmentInstanceFilter(c, rawDesired.InstanceFilter, rawNew.InstanceFilter)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Rollout) && dcl.IsEmptyValueIndirect(rawDesired.Rollout) {
		rawNew.Rollout = rawDesired.Rollout
	} else {
		rawNew.Rollout = canonicalizeNewOsPolicyAssignmentRollout(c, rawDesired.Rollout, rawNew.Rollout)
	}

	if dcl.IsEmptyValueIndirect(rawNew.RevisionId) && dcl.IsEmptyValueIndirect(rawDesired.RevisionId) {
		rawNew.RevisionId = rawDesired.RevisionId
	} else {
		if dcl.StringCanonicalize(rawDesired.RevisionId, rawNew.RevisionId) {
			rawNew.RevisionId = rawDesired.RevisionId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RevisionCreateTime) && dcl.IsEmptyValueIndirect(rawDesired.RevisionCreateTime) {
		rawNew.RevisionCreateTime = rawDesired.RevisionCreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RolloutState) && dcl.IsEmptyValueIndirect(rawDesired.RolloutState) {
		rawNew.RolloutState = rawDesired.RolloutState
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Baseline) && dcl.IsEmptyValueIndirect(rawDesired.Baseline) {
		rawNew.Baseline = rawDesired.Baseline
	} else {
		if dcl.BoolCanonicalize(rawDesired.Baseline, rawNew.Baseline) {
			rawNew.Baseline = rawDesired.Baseline
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deleted) && dcl.IsEmptyValueIndirect(rawDesired.Deleted) {
		rawNew.Deleted = rawDesired.Deleted
	} else {
		if dcl.BoolCanonicalize(rawDesired.Deleted, rawNew.Deleted) {
			rawNew.Deleted = rawDesired.Deleted
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Reconciling) && dcl.IsEmptyValueIndirect(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeOsPolicyAssignmentOsPolicies(des, initial *OsPolicyAssignmentOsPolicies, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPolicies {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}
	if dcl.IsZeroValue(des.ResourceGroups) {
		des.ResourceGroups = initial.ResourceGroups
	}
	if dcl.BoolCanonicalize(des.AllowNoResourceGroupMatch, initial.AllowNoResourceGroupMatch) || dcl.IsZeroValue(des.AllowNoResourceGroupMatch) {
		des.AllowNoResourceGroupMatch = initial.AllowNoResourceGroupMatch
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPolicies(c *Client, des, nw *OsPolicyAssignmentOsPolicies) *OsPolicyAssignmentOsPolicies {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.IsZeroValue(nw.Mode) {
		nw.Mode = des.Mode
	}
	nw.ResourceGroups = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c, des.ResourceGroups, nw.ResourceGroups)
	if dcl.BoolCanonicalize(des.AllowNoResourceGroupMatch, nw.AllowNoResourceGroupMatch) {
		nw.AllowNoResourceGroupMatch = des.AllowNoResourceGroupMatch
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesSet(c *Client, des, nw []OsPolicyAssignmentOsPolicies) []OsPolicyAssignmentOsPolicies {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPolicies
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesSlice(c *Client, des, nw []OsPolicyAssignmentOsPolicies) []OsPolicyAssignmentOsPolicies {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPolicies
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPolicies(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroups(des, initial *OsPolicyAssignmentOsPoliciesResourceGroups, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroups {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.OsFilter = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(des.OsFilter, initial.OsFilter, opts...)
	if dcl.IsZeroValue(des.Resources) {
		des.Resources = initial.Resources
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroups(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroups) *OsPolicyAssignmentOsPoliciesResourceGroups {
	if des == nil || nw == nil {
		return nw
	}

	nw.OsFilter = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, des.OsFilter, nw.OsFilter)
	nw.Resources = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c, des.Resources, nw.Resources)

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroups) []OsPolicyAssignmentOsPoliciesResourceGroups {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroups
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroups) []OsPolicyAssignmentOsPoliciesResourceGroups {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroups
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroups(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.OsShortName, initial.OsShortName) || dcl.IsZeroValue(des.OsShortName) {
		des.OsShortName = initial.OsShortName
	}
	if dcl.StringCanonicalize(des.OsVersion, initial.OsVersion) || dcl.IsZeroValue(des.OsVersion) {
		des.OsVersion = initial.OsVersion
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.OsShortName, nw.OsShortName) {
		nw.OsShortName = des.OsShortName
	}
	if dcl.StringCanonicalize(des.OsVersion, nw.OsVersion) {
		nw.OsVersion = des.OsVersion
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResources(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResources, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	des.Pkg = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(des.Pkg, initial.Pkg, opts...)
	des.Repository = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(des.Repository, initial.Repository, opts...)
	des.Exec = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(des.Exec, initial.Exec, opts...)
	des.File = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(des.File, initial.File, opts...)

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResources(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResources) *OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	nw.Pkg = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, des.Pkg, nw.Pkg)
	nw.Repository = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, des.Repository, nw.Repository)
	nw.Exec = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, des.Exec, nw.Exec)
	nw.File = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, des.File, nw.File)

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResources) []OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResources) []OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResources(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DesiredState) {
		des.DesiredState = initial.DesiredState
	}
	des.Apt = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(des.Apt, initial.Apt, opts...)
	des.Deb = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(des.Deb, initial.Deb, opts...)
	des.Yum = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(des.Yum, initial.Yum, opts...)
	des.Zypper = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(des.Zypper, initial.Zypper, opts...)
	des.Rpm = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(des.Rpm, initial.Rpm, opts...)
	des.Googet = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(des.Googet, initial.Googet, opts...)
	des.Msi = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(des.Msi, initial.Msi, opts...)

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.DesiredState) {
		nw.DesiredState = des.DesiredState
	}
	nw.Apt = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, des.Apt, nw.Apt)
	nw.Deb = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, des.Deb, nw.Deb)
	nw.Yum = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, des.Yum, nw.Yum)
	nw.Zypper = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, des.Zypper, nw.Zypper)
	nw.Rpm = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, des.Rpm, nw.Rpm)
	nw.Googet = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, des.Googet, nw.Googet)
	nw.Msi = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, des.Msi, nw.Msi)

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
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

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Source = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(des.Source, initial.Source, opts...)
	if dcl.BoolCanonicalize(des.PullDeps, initial.PullDeps) || dcl.IsZeroValue(des.PullDeps) {
		des.PullDeps = initial.PullDeps
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if des == nil || nw == nil {
		return nw
	}

	nw.Source = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, des.Source, nw.Source)
	if dcl.BoolCanonicalize(des.PullDeps, nw.PullDeps) {
		nw.PullDeps = des.PullDeps
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Remote = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(des.Remote, initial.Remote, opts...)
	des.Gcs = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(des.Gcs, initial.Gcs, opts...)
	if dcl.StringCanonicalize(des.LocalPath, initial.LocalPath) || dcl.IsZeroValue(des.LocalPath) {
		des.LocalPath = initial.LocalPath
	}
	if dcl.BoolCanonicalize(des.AllowInsecure, initial.AllowInsecure) || dcl.IsZeroValue(des.AllowInsecure) {
		des.AllowInsecure = initial.AllowInsecure
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if des == nil || nw == nil {
		return nw
	}

	nw.Remote = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, des.Remote, nw.Remote)
	nw.Gcs = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, des.Gcs, nw.Gcs)
	if dcl.StringCanonicalize(des.LocalPath, nw.LocalPath) {
		nw.LocalPath = des.LocalPath
	}
	if dcl.BoolCanonicalize(des.AllowInsecure, nw.AllowInsecure) {
		nw.AllowInsecure = des.AllowInsecure
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		des.Uri = initial.Uri
	}
	if dcl.StringCanonicalize(des.Sha256Checksum, initial.Sha256Checksum) || dcl.IsZeroValue(des.Sha256Checksum) {
		des.Sha256Checksum = initial.Sha256Checksum
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Sha256Checksum, nw.Sha256Checksum) {
		nw.Sha256Checksum = des.Sha256Checksum
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Bucket, initial.Bucket) || dcl.IsZeroValue(des.Bucket) {
		des.Bucket = initial.Bucket
	}
	if dcl.NameToSelfLink(des.Object, initial.Object) || dcl.IsZeroValue(des.Object) {
		des.Object = initial.Object
	}
	if dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Bucket, nw.Bucket) {
		nw.Bucket = des.Bucket
	}
	if dcl.NameToSelfLink(des.Object, nw.Object) {
		nw.Object = des.Object
	}
	if dcl.IsZeroValue(nw.Generation) {
		nw.Generation = des.Generation
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
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

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
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

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Source = canonicalizeOsPolicyAssignmentFile(des.Source, initial.Source, opts...)
	if dcl.BoolCanonicalize(des.PullDeps, initial.PullDeps) || dcl.IsZeroValue(des.PullDeps) {
		des.PullDeps = initial.PullDeps
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if des == nil || nw == nil {
		return nw
	}

	nw.Source = canonicalizeNewOsPolicyAssignmentFile(c, des.Source, nw.Source)
	if dcl.BoolCanonicalize(des.PullDeps, nw.PullDeps) {
		nw.PullDeps = des.PullDeps
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentFile(des, initial *OsPolicyAssignmentFile, opts ...dcl.ApplyOption) *OsPolicyAssignmentFile {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Remote = canonicalizeOsPolicyAssignmentFileRemote(des.Remote, initial.Remote, opts...)
	des.Gcs = canonicalizeOsPolicyAssignmentFileGcs(des.Gcs, initial.Gcs, opts...)
	if dcl.StringCanonicalize(des.LocalPath, initial.LocalPath) || dcl.IsZeroValue(des.LocalPath) {
		des.LocalPath = initial.LocalPath
	}
	if dcl.BoolCanonicalize(des.AllowInsecure, initial.AllowInsecure) || dcl.IsZeroValue(des.AllowInsecure) {
		des.AllowInsecure = initial.AllowInsecure
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentFile(c *Client, des, nw *OsPolicyAssignmentFile) *OsPolicyAssignmentFile {
	if des == nil || nw == nil {
		return nw
	}

	nw.Remote = canonicalizeNewOsPolicyAssignmentFileRemote(c, des.Remote, nw.Remote)
	nw.Gcs = canonicalizeNewOsPolicyAssignmentFileGcs(c, des.Gcs, nw.Gcs)
	if dcl.StringCanonicalize(des.LocalPath, nw.LocalPath) {
		nw.LocalPath = des.LocalPath
	}
	if dcl.BoolCanonicalize(des.AllowInsecure, nw.AllowInsecure) {
		nw.AllowInsecure = des.AllowInsecure
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentFileSet(c *Client, des, nw []OsPolicyAssignmentFile) []OsPolicyAssignmentFile {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentFile
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentFileNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentFileSlice(c *Client, des, nw []OsPolicyAssignmentFile) []OsPolicyAssignmentFile {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentFile
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentFile(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentFileRemote(des, initial *OsPolicyAssignmentFileRemote, opts ...dcl.ApplyOption) *OsPolicyAssignmentFileRemote {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		des.Uri = initial.Uri
	}
	if dcl.StringCanonicalize(des.Sha256Checksum, initial.Sha256Checksum) || dcl.IsZeroValue(des.Sha256Checksum) {
		des.Sha256Checksum = initial.Sha256Checksum
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentFileRemote(c *Client, des, nw *OsPolicyAssignmentFileRemote) *OsPolicyAssignmentFileRemote {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Sha256Checksum, nw.Sha256Checksum) {
		nw.Sha256Checksum = des.Sha256Checksum
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentFileRemoteSet(c *Client, des, nw []OsPolicyAssignmentFileRemote) []OsPolicyAssignmentFileRemote {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentFileRemote
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentFileRemoteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentFileRemoteSlice(c *Client, des, nw []OsPolicyAssignmentFileRemote) []OsPolicyAssignmentFileRemote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentFileRemote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentFileRemote(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentFileGcs(des, initial *OsPolicyAssignmentFileGcs, opts ...dcl.ApplyOption) *OsPolicyAssignmentFileGcs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Bucket, initial.Bucket) || dcl.IsZeroValue(des.Bucket) {
		des.Bucket = initial.Bucket
	}
	if dcl.StringCanonicalize(des.Object, initial.Object) || dcl.IsZeroValue(des.Object) {
		des.Object = initial.Object
	}
	if dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentFileGcs(c *Client, des, nw *OsPolicyAssignmentFileGcs) *OsPolicyAssignmentFileGcs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Bucket, nw.Bucket) {
		nw.Bucket = des.Bucket
	}
	if dcl.StringCanonicalize(des.Object, nw.Object) {
		nw.Object = des.Object
	}
	if dcl.IsZeroValue(nw.Generation) {
		nw.Generation = des.Generation
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentFileGcsSet(c *Client, des, nw []OsPolicyAssignmentFileGcs) []OsPolicyAssignmentFileGcs {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentFileGcs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentFileGcsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentFileGcsSlice(c *Client, des, nw []OsPolicyAssignmentFileGcs) []OsPolicyAssignmentFileGcs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentFileGcs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentFileGcs(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
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

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Source = canonicalizeOsPolicyAssignmentFile(des.Source, initial.Source, opts...)
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if des == nil || nw == nil {
		return nw
	}

	nw.Source = canonicalizeNewOsPolicyAssignmentFile(c, des.Source, nw.Source)
	if dcl.IsZeroValue(nw.Properties) {
		nw.Properties = des.Properties
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Apt = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(des.Apt, initial.Apt, opts...)
	des.Yum = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(des.Yum, initial.Yum, opts...)
	des.Zypper = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(des.Zypper, initial.Zypper, opts...)
	des.Goo = canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(des.Goo, initial.Goo, opts...)

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if des == nil || nw == nil {
		return nw
	}

	nw.Apt = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, des.Apt, nw.Apt)
	nw.Yum = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, des.Yum, nw.Yum)
	nw.Zypper = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, des.Zypper, nw.Zypper)
	nw.Goo = canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, des.Goo, nw.Goo)

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ArchiveType) {
		des.ArchiveType = initial.ArchiveType
	}
	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		des.Uri = initial.Uri
	}
	if dcl.StringCanonicalize(des.Distribution, initial.Distribution) || dcl.IsZeroValue(des.Distribution) {
		des.Distribution = initial.Distribution
	}
	if dcl.IsZeroValue(des.Components) {
		des.Components = initial.Components
	}
	if dcl.StringCanonicalize(des.GpgKey, initial.GpgKey) || dcl.IsZeroValue(des.GpgKey) {
		des.GpgKey = initial.GpgKey
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ArchiveType) {
		nw.ArchiveType = des.ArchiveType
	}
	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Distribution, nw.Distribution) {
		nw.Distribution = des.Distribution
	}
	if dcl.IsZeroValue(nw.Components) {
		nw.Components = des.Components
	}
	if dcl.StringCanonicalize(des.GpgKey, nw.GpgKey) {
		nw.GpgKey = des.GpgKey
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	if dcl.StringCanonicalize(des.BaseUrl, initial.BaseUrl) || dcl.IsZeroValue(des.BaseUrl) {
		des.BaseUrl = initial.BaseUrl
	}
	if dcl.IsZeroValue(des.GpgKeys) {
		des.GpgKeys = initial.GpgKeys
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.BaseUrl, nw.BaseUrl) {
		nw.BaseUrl = des.BaseUrl
	}
	if dcl.IsZeroValue(nw.GpgKeys) {
		nw.GpgKeys = des.GpgKeys
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	if dcl.StringCanonicalize(des.BaseUrl, initial.BaseUrl) || dcl.IsZeroValue(des.BaseUrl) {
		des.BaseUrl = initial.BaseUrl
	}
	if dcl.IsZeroValue(des.GpgKeys) {
		des.GpgKeys = initial.GpgKeys
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.BaseUrl, nw.BaseUrl) {
		nw.BaseUrl = des.BaseUrl
	}
	if dcl.IsZeroValue(nw.GpgKeys) {
		nw.GpgKeys = des.GpgKeys
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
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
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Validate = canonicalizeOsPolicyAssignmentExec(des.Validate, initial.Validate, opts...)
	des.Enforce = canonicalizeOsPolicyAssignmentExec(des.Enforce, initial.Enforce, opts...)

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if des == nil || nw == nil {
		return nw
	}

	nw.Validate = canonicalizeNewOsPolicyAssignmentExec(c, des.Validate, nw.Validate)
	nw.Enforce = canonicalizeNewOsPolicyAssignmentExec(c, des.Enforce, nw.Enforce)

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentExec(des, initial *OsPolicyAssignmentExec, opts ...dcl.ApplyOption) *OsPolicyAssignmentExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.File = canonicalizeOsPolicyAssignmentFile(des.File, initial.File, opts...)
	if dcl.StringCanonicalize(des.Script, initial.Script) || dcl.IsZeroValue(des.Script) {
		des.Script = initial.Script
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.Interpreter) {
		des.Interpreter = initial.Interpreter
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentExec(c *Client, des, nw *OsPolicyAssignmentExec) *OsPolicyAssignmentExec {
	if des == nil || nw == nil {
		return nw
	}

	nw.File = canonicalizeNewOsPolicyAssignmentFile(c, des.File, nw.File)
	if dcl.StringCanonicalize(des.Script, nw.Script) {
		nw.Script = des.Script
	}
	if dcl.IsZeroValue(nw.Args) {
		nw.Args = des.Args
	}
	if dcl.IsZeroValue(nw.Interpreter) {
		nw.Interpreter = des.Interpreter
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentExecSet(c *Client, des, nw []OsPolicyAssignmentExec) []OsPolicyAssignmentExec {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentExecSlice(c *Client, des, nw []OsPolicyAssignmentExec) []OsPolicyAssignmentExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentExec(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(des, initial *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile, opts ...dcl.ApplyOption) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.File = canonicalizeOsPolicyAssignmentFile(des.File, initial.File, opts...)
	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.StringCanonicalize(des.Permissions, initial.Permissions) || dcl.IsZeroValue(des.Permissions) {
		des.Permissions = initial.Permissions
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c *Client, des, nw *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if des == nil || nw == nil {
		return nw
	}

	nw.File = canonicalizeNewOsPolicyAssignmentFile(c, des.File, nw.File)
	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.IsZeroValue(nw.State) {
		nw.State = des.State
	}
	if dcl.StringCanonicalize(des.Permissions, nw.Permissions) {
		nw.Permissions = des.Permissions
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSet(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSlice(c *Client, des, nw []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentInstanceFilter(des, initial *OsPolicyAssignmentInstanceFilter, opts ...dcl.ApplyOption) *OsPolicyAssignmentInstanceFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.All, initial.All) || dcl.IsZeroValue(des.All) {
		des.All = initial.All
	}
	if dcl.IsZeroValue(des.OsShortNames) {
		des.OsShortNames = initial.OsShortNames
	}
	if dcl.IsZeroValue(des.InclusionLabels) {
		des.InclusionLabels = initial.InclusionLabels
	}
	if dcl.IsZeroValue(des.ExclusionLabels) {
		des.ExclusionLabels = initial.ExclusionLabels
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentInstanceFilter(c *Client, des, nw *OsPolicyAssignmentInstanceFilter) *OsPolicyAssignmentInstanceFilter {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.All, nw.All) {
		nw.All = des.All
	}
	if dcl.IsZeroValue(nw.OsShortNames) {
		nw.OsShortNames = des.OsShortNames
	}
	nw.InclusionLabels = canonicalizeNewOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c, des.InclusionLabels, nw.InclusionLabels)
	nw.ExclusionLabels = canonicalizeNewOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c, des.ExclusionLabels, nw.ExclusionLabels)

	return nw
}

func canonicalizeNewOsPolicyAssignmentInstanceFilterSet(c *Client, des, nw []OsPolicyAssignmentInstanceFilter) []OsPolicyAssignmentInstanceFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentInstanceFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentInstanceFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentInstanceFilterSlice(c *Client, des, nw []OsPolicyAssignmentInstanceFilter) []OsPolicyAssignmentInstanceFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentInstanceFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentInstanceFilter(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentInstanceFilterInclusionLabels(des, initial *OsPolicyAssignmentInstanceFilterInclusionLabels, opts ...dcl.ApplyOption) *OsPolicyAssignmentInstanceFilterInclusionLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentInstanceFilterInclusionLabels(c *Client, des, nw *OsPolicyAssignmentInstanceFilterInclusionLabels) *OsPolicyAssignmentInstanceFilterInclusionLabels {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentInstanceFilterInclusionLabelsSet(c *Client, des, nw []OsPolicyAssignmentInstanceFilterInclusionLabels) []OsPolicyAssignmentInstanceFilterInclusionLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentInstanceFilterInclusionLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentInstanceFilterInclusionLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c *Client, des, nw []OsPolicyAssignmentInstanceFilterInclusionLabels) []OsPolicyAssignmentInstanceFilterInclusionLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentInstanceFilterInclusionLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentInstanceFilterInclusionLabels(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentInstanceFilterExclusionLabels(des, initial *OsPolicyAssignmentInstanceFilterExclusionLabels, opts ...dcl.ApplyOption) *OsPolicyAssignmentInstanceFilterExclusionLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentInstanceFilterExclusionLabels(c *Client, des, nw *OsPolicyAssignmentInstanceFilterExclusionLabels) *OsPolicyAssignmentInstanceFilterExclusionLabels {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentInstanceFilterExclusionLabelsSet(c *Client, des, nw []OsPolicyAssignmentInstanceFilterExclusionLabels) []OsPolicyAssignmentInstanceFilterExclusionLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentInstanceFilterExclusionLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentInstanceFilterExclusionLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c *Client, des, nw []OsPolicyAssignmentInstanceFilterExclusionLabels) []OsPolicyAssignmentInstanceFilterExclusionLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentInstanceFilterExclusionLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentInstanceFilterExclusionLabels(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentRollout(des, initial *OsPolicyAssignmentRollout, opts ...dcl.ApplyOption) *OsPolicyAssignmentRollout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.DisruptionBudget = canonicalizeOsPolicyAssignmentRolloutDisruptionBudget(des.DisruptionBudget, initial.DisruptionBudget, opts...)
	if canonicalizeOSPolicyAssignmentRolloutMinWaitDuration(des.MinWaitDuration, initial.MinWaitDuration) || dcl.IsZeroValue(des.MinWaitDuration) {
		des.MinWaitDuration = initial.MinWaitDuration
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentRollout(c *Client, des, nw *OsPolicyAssignmentRollout) *OsPolicyAssignmentRollout {
	if des == nil || nw == nil {
		return nw
	}

	nw.DisruptionBudget = canonicalizeNewOsPolicyAssignmentRolloutDisruptionBudget(c, des.DisruptionBudget, nw.DisruptionBudget)
	if canonicalizeOSPolicyAssignmentRolloutMinWaitDuration(des.MinWaitDuration, nw.MinWaitDuration) {
		nw.MinWaitDuration = des.MinWaitDuration
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentRolloutSet(c *Client, des, nw []OsPolicyAssignmentRollout) []OsPolicyAssignmentRollout {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentRollout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentRolloutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentRolloutSlice(c *Client, des, nw []OsPolicyAssignmentRollout) []OsPolicyAssignmentRollout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentRollout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentRollout(c, &d, &n))
	}

	return items
}

func canonicalizeOsPolicyAssignmentRolloutDisruptionBudget(des, initial *OsPolicyAssignmentRolloutDisruptionBudget, opts ...dcl.ApplyOption) *OsPolicyAssignmentRolloutDisruptionBudget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewOsPolicyAssignmentRolloutDisruptionBudget(c *Client, des, nw *OsPolicyAssignmentRolloutDisruptionBudget) *OsPolicyAssignmentRolloutDisruptionBudget {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Fixed) {
		nw.Fixed = des.Fixed
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewOsPolicyAssignmentRolloutDisruptionBudgetSet(c *Client, des, nw []OsPolicyAssignmentRolloutDisruptionBudget) []OsPolicyAssignmentRolloutDisruptionBudget {
	if des == nil {
		return nw
	}
	var reorderedNew []OsPolicyAssignmentRolloutDisruptionBudget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareOsPolicyAssignmentRolloutDisruptionBudgetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewOsPolicyAssignmentRolloutDisruptionBudgetSlice(c *Client, des, nw []OsPolicyAssignmentRolloutDisruptionBudget) []OsPolicyAssignmentRolloutDisruptionBudget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OsPolicyAssignmentRolloutDisruptionBudget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOsPolicyAssignmentRolloutDisruptionBudget(c, &d, &n))
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
func diffOsPolicyAssignment(c *Client, desired, actual *OsPolicyAssignment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsPolicies, actual.OsPolicies, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPolicies, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("OsPolicies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceFilter, actual.InstanceFilter, dcl.Info{ObjectFunction: compareOsPolicyAssignmentInstanceFilterNewStyle, EmptyObject: EmptyOsPolicyAssignmentInstanceFilter, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("InstanceFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rollout, actual.Rollout, dcl.Info{ObjectFunction: compareOsPolicyAssignmentRolloutNewStyle, EmptyObject: EmptyOsPolicyAssignmentRollout, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Rollout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionId, actual.RevisionId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RevisionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionCreateTime, actual.RevisionCreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RevisionCreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RolloutState, actual.RolloutState, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RolloutState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Baseline, actual.Baseline, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Baseline")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deleted, actual.Deleted, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deleted")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reconciling, actual.Reconciling, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reconciling")); len(ds) != 0 || err != nil {
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
func compareOsPolicyAssignmentOsPoliciesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPolicies)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPolicies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPolicies or *OsPolicyAssignmentOsPolicies", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPolicies)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPolicies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPolicies", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceGroups, actual.ResourceGroups, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroups, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("ResourceGroups")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowNoResourceGroupMatch, actual.AllowNoResourceGroupMatch, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("AllowNoResourceGroupMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroups)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroups or *OsPolicyAssignmentOsPoliciesResourceGroups", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroups)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroups", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OsFilter, actual.OsFilter, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("OsFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResources, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter or *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OsShortName, actual.OsShortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("OsShortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsVersion, actual.OsVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("OsVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResources)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResources or *OsPolicyAssignmentOsPoliciesResourceGroupsResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResources)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Pkg, actual.Pkg, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Pkg")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Repository, actual.Repository, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Repository")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Exec, actual.Exec, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Exec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.File, actual.File, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("File")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DesiredState, actual.DesiredState, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("DesiredState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Apt, actual.Apt, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Apt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deb, actual.Deb, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Deb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Yum, actual.Yum, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Yum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zypper, actual.Zypper, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Zypper")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rpm, actual.Rpm, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Rpm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Googet, actual.Googet, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Googet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Msi, actual.Msi, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Msi")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PullDeps, actual.PullDeps, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("PullDeps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Remote, actual.Remote, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Remote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gcs, actual.Gcs, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Gcs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalPath, actual.LocalPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("LocalPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowInsecure, actual.AllowInsecure, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("AllowInsecure")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256Checksum, actual.Sha256Checksum, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Sha256Checksum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Object, actual.Object, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Object")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileNewStyle, EmptyObject: EmptyOsPolicyAssignmentFile, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PullDeps, actual.PullDeps, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("PullDeps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentFileNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentFile)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentFile)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFile or *OsPolicyAssignmentFile", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentFile)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentFile)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFile", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Remote, actual.Remote, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileRemoteNewStyle, EmptyObject: EmptyOsPolicyAssignmentFileRemote, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Remote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gcs, actual.Gcs, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileGcsNewStyle, EmptyObject: EmptyOsPolicyAssignmentFileGcs, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Gcs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalPath, actual.LocalPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("LocalPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowInsecure, actual.AllowInsecure, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("AllowInsecure")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentFileRemoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentFileRemote)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentFileRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFileRemote or *OsPolicyAssignmentFileRemote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentFileRemote)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentFileRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFileRemote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256Checksum, actual.Sha256Checksum, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Sha256Checksum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentFileGcsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentFileGcs)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentFileGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFileGcs or *OsPolicyAssignmentFileGcs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentFileGcs)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentFileGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentFileGcs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Object, actual.Object, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Object")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileNewStyle, EmptyObject: EmptyOsPolicyAssignmentFile, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Apt, actual.Apt, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Apt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Yum, actual.Yum, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Yum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zypper, actual.Zypper, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Zypper")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Goo, actual.Goo, dcl.Info{ObjectFunction: compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooNewStyle, EmptyObject: EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Goo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArchiveType, actual.ArchiveType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("ArchiveType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Distribution, actual.Distribution, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Distribution")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Components, actual.Components, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Components")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKey, actual.GpgKey, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("GpgKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseUrl, actual.BaseUrl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("BaseUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKeys, actual.GpgKeys, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("GpgKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseUrl, actual.BaseUrl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("BaseUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKeys, actual.GpgKeys, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("GpgKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Validate, actual.Validate, dcl.Info{ObjectFunction: compareOsPolicyAssignmentExecNewStyle, EmptyObject: EmptyOsPolicyAssignmentExec, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Validate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Enforce, actual.Enforce, dcl.Info{ObjectFunction: compareOsPolicyAssignmentExecNewStyle, EmptyObject: EmptyOsPolicyAssignmentExec, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Enforce")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentExec)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentExec or *OsPolicyAssignmentExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentExec)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.File, actual.File, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileNewStyle, EmptyObject: EmptyOsPolicyAssignmentFile, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("File")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Script, actual.Script, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Script")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interpreter, actual.Interpreter, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Interpreter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile or *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.File, actual.File, dcl.Info{ObjectFunction: compareOsPolicyAssignmentFileNewStyle, EmptyObject: EmptyOsPolicyAssignmentFile, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("File")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Permissions, actual.Permissions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Permissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentInstanceFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentInstanceFilter)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentInstanceFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilter or *OsPolicyAssignmentInstanceFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentInstanceFilter)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentInstanceFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.All, actual.All, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("All")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsShortNames, actual.OsShortNames, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("OsShortNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InclusionLabels, actual.InclusionLabels, dcl.Info{ObjectFunction: compareOsPolicyAssignmentInstanceFilterInclusionLabelsNewStyle, EmptyObject: EmptyOsPolicyAssignmentInstanceFilterInclusionLabels, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("InclusionLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExclusionLabels, actual.ExclusionLabels, dcl.Info{ObjectFunction: compareOsPolicyAssignmentInstanceFilterExclusionLabelsNewStyle, EmptyObject: EmptyOsPolicyAssignmentInstanceFilterExclusionLabels, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("ExclusionLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentInstanceFilterInclusionLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentInstanceFilterInclusionLabels)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentInstanceFilterInclusionLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilterInclusionLabels or *OsPolicyAssignmentInstanceFilterInclusionLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentInstanceFilterInclusionLabels)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentInstanceFilterInclusionLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilterInclusionLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentInstanceFilterExclusionLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentInstanceFilterExclusionLabels)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentInstanceFilterExclusionLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilterExclusionLabels or *OsPolicyAssignmentInstanceFilterExclusionLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentInstanceFilterExclusionLabels)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentInstanceFilterExclusionLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentInstanceFilterExclusionLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentRolloutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentRollout)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentRollout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentRollout or *OsPolicyAssignmentRollout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentRollout)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentRollout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentRollout", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisruptionBudget, actual.DisruptionBudget, dcl.Info{ObjectFunction: compareOsPolicyAssignmentRolloutDisruptionBudgetNewStyle, EmptyObject: EmptyOsPolicyAssignmentRolloutDisruptionBudget, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("DisruptionBudget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinWaitDuration, actual.MinWaitDuration, dcl.Info{CustomDiff: canonicalizeOSPolicyAssignmentRolloutMinWaitDuration, OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("MinWaitDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOsPolicyAssignmentRolloutDisruptionBudgetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OsPolicyAssignmentRolloutDisruptionBudget)
	if !ok {
		desiredNotPointer, ok := d.(OsPolicyAssignmentRolloutDisruptionBudget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentRolloutDisruptionBudget or *OsPolicyAssignmentRolloutDisruptionBudget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OsPolicyAssignmentRolloutDisruptionBudget)
	if !ok {
		actualNotPointer, ok := a.(OsPolicyAssignmentRolloutDisruptionBudget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OsPolicyAssignmentRolloutDisruptionBudget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fixed, actual.Fixed, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Fixed")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *OsPolicyAssignment) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *OsPolicyAssignment) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *OsPolicyAssignment) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *OsPolicyAssignment) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateOSPolicyAssignment" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/osPolicyAssignments/{{name}}", "https://osconfig.googleapis.com/v1alpha/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the OsPolicyAssignment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *OsPolicyAssignment) marshal(c *Client) ([]byte, error) {
	m, err := expandOsPolicyAssignment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling OsPolicyAssignment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalOsPolicyAssignment decodes JSON responses into the OsPolicyAssignment resource schema.
func unmarshalOsPolicyAssignment(b []byte, c *Client) (*OsPolicyAssignment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapOsPolicyAssignment(m, c)
}

func unmarshalMapOsPolicyAssignment(m map[string]interface{}, c *Client) (*OsPolicyAssignment, error) {

	return flattenOsPolicyAssignment(c, m), nil
}

// expandOsPolicyAssignment expands OsPolicyAssignment into a JSON request object.
func expandOsPolicyAssignment(c *Client, f *OsPolicyAssignment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/osPolicyAssignments/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesSlice(c, f.OsPolicies); err != nil {
		return nil, fmt.Errorf("error expanding OsPolicies into osPolicies: %w", err)
	} else if v != nil {
		m["osPolicies"] = v
	}
	if v, err := expandOsPolicyAssignmentInstanceFilter(c, f.InstanceFilter); err != nil {
		return nil, fmt.Errorf("error expanding InstanceFilter into instanceFilter: %w", err)
	} else if v != nil {
		m["instanceFilter"] = v
	}
	if v, err := expandOsPolicyAssignmentRollout(c, f.Rollout); err != nil {
		return nil, fmt.Errorf("error expanding Rollout into rollout: %w", err)
	} else if v != nil {
		m["rollout"] = v
	}
	if v := f.RevisionId; !dcl.IsEmptyValueIndirect(v) {
		m["revisionId"] = v
	}
	if v := f.RevisionCreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["revisionCreateTime"] = v
	}
	if v := f.RolloutState; !dcl.IsEmptyValueIndirect(v) {
		m["rolloutState"] = v
	}
	if v := f.Baseline; !dcl.IsEmptyValueIndirect(v) {
		m["baseline"] = v
	}
	if v := f.Deleted; !dcl.IsEmptyValueIndirect(v) {
		m["deleted"] = v
	}
	if v := f.Reconciling; !dcl.IsEmptyValueIndirect(v) {
		m["reconciling"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
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

// flattenOsPolicyAssignment flattens OsPolicyAssignment from a JSON request object into the
// OsPolicyAssignment type.
func flattenOsPolicyAssignment(c *Client, i interface{}) *OsPolicyAssignment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &OsPolicyAssignment{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.OsPolicies = flattenOsPolicyAssignmentOsPoliciesSlice(c, m["osPolicies"])
	res.InstanceFilter = flattenOsPolicyAssignmentInstanceFilter(c, m["instanceFilter"])
	res.Rollout = flattenOsPolicyAssignmentRollout(c, m["rollout"])
	res.RevisionId = dcl.FlattenString(m["revisionId"])
	res.RevisionCreateTime = dcl.FlattenString(m["revisionCreateTime"])
	res.RolloutState = flattenOsPolicyAssignmentRolloutStateEnum(m["rolloutState"])
	res.Baseline = dcl.FlattenBool(m["baseline"])
	res.Deleted = dcl.FlattenBool(m["deleted"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandOsPolicyAssignmentOsPoliciesMap expands the contents of OsPolicyAssignmentOsPolicies into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesMap(c *Client, f map[string]OsPolicyAssignmentOsPolicies) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPolicies(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesSlice expands the contents of OsPolicyAssignmentOsPolicies into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesSlice(c *Client, f []OsPolicyAssignmentOsPolicies) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPolicies(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesMap flattens the contents of OsPolicyAssignmentOsPolicies from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPolicies {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPolicies{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPolicies{}
	}

	items := make(map[string]OsPolicyAssignmentOsPolicies)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPolicies(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesSlice flattens the contents of OsPolicyAssignmentOsPolicies from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPolicies {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPolicies{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPolicies{}
	}

	items := make([]OsPolicyAssignmentOsPolicies, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPolicies(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPolicies expands an instance of OsPolicyAssignmentOsPolicies into a JSON
// request object.
func expandOsPolicyAssignmentOsPolicies(c *Client, f *OsPolicyAssignmentOsPolicies) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c, f.ResourceGroups); err != nil {
		return nil, fmt.Errorf("error expanding ResourceGroups into resourceGroups: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroups"] = v
	}
	if v := f.AllowNoResourceGroupMatch; !dcl.IsEmptyValueIndirect(v) {
		m["allowNoResourceGroupMatch"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPolicies flattens an instance of OsPolicyAssignmentOsPolicies from a JSON
// response object.
func flattenOsPolicyAssignmentOsPolicies(c *Client, i interface{}) *OsPolicyAssignmentOsPolicies {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPolicies{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPolicies
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Description = dcl.FlattenString(m["description"])
	r.Mode = flattenOsPolicyAssignmentOsPoliciesModeEnum(m["mode"])
	r.ResourceGroups = flattenOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c, m["resourceGroups"])
	r.AllowNoResourceGroupMatch = dcl.FlattenBool(m["allowNoResourceGroupMatch"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroups into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroups) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroups(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroups into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroups) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroups(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroups from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroups {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroups{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroups{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroups)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroups(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroups from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroups {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroups{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroups{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroups, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroups(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroups expands an instance of OsPolicyAssignmentOsPoliciesResourceGroups into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroups(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroups) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, f.OsFilter); err != nil {
		return nil, fmt.Errorf("error expanding OsFilter into osFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["osFilter"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroups flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroups from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroups(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroups {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroups{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroups
	}
	r.OsFilter = flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, m["osFilter"])
	r.Resources = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c, m["resources"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.OsShortName; !dcl.IsEmptyValueIndirect(v) {
		m["osShortName"] = v
	}
	if v := f.OsVersion; !dcl.IsEmptyValueIndirect(v) {
		m["osVersion"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter
	}
	r.OsShortName = dcl.FlattenString(m["osShortName"])
	r.OsVersion = dcl.FlattenString(m["osVersion"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResources into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResources into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResources from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResources{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResources{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResources)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResources from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResources{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResources{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResources expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResources into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResources(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, f.Pkg); err != nil {
		return nil, fmt.Errorf("error expanding Pkg into pkg: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pkg"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, f.Repository); err != nil {
		return nil, fmt.Errorf("error expanding Repository into repository: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["repository"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, f.Exec); err != nil {
		return nil, fmt.Errorf("error expanding Exec into exec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exec"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, f.File); err != nil {
		return nil, fmt.Errorf("error expanding File into file: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["file"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResources flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResources from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResources(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResources
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Pkg = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, m["pkg"])
	r.Repository = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, m["repository"])
	r.Exec = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, m["exec"])
	r.File = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, m["file"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DesiredState; !dcl.IsEmptyValueIndirect(v) {
		m["desiredState"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, f.Apt); err != nil {
		return nil, fmt.Errorf("error expanding Apt into apt: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apt"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, f.Deb); err != nil {
		return nil, fmt.Errorf("error expanding Deb into deb: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deb"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, f.Yum); err != nil {
		return nil, fmt.Errorf("error expanding Yum into yum: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yum"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, f.Zypper); err != nil {
		return nil, fmt.Errorf("error expanding Zypper into zypper: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zypper"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, f.Rpm); err != nil {
		return nil, fmt.Errorf("error expanding Rpm into rpm: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rpm"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, f.Googet); err != nil {
		return nil, fmt.Errorf("error expanding Googet into googet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["googet"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, f.Msi); err != nil {
		return nil, fmt.Errorf("error expanding Msi into msi: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["msi"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	}
	r.DesiredState = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(m["desiredState"])
	r.Apt = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, m["apt"])
	r.Deb = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, m["deb"])
	r.Yum = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, m["yum"])
	r.Zypper = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, m["zypper"])
	r.Rpm = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, m["rpm"])
	r.Googet = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, m["googet"])
	r.Msi = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, m["msi"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, f.Source); err != nil {
		return nil, fmt.Errorf("error expanding Source into source: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.PullDeps; !dcl.IsEmptyValueIndirect(v) {
		m["pullDeps"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb
	}
	r.Source = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, m["source"])
	r.PullDeps = dcl.FlattenBool(m["pullDeps"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, f.Remote); err != nil {
		return nil, fmt.Errorf("error expanding Remote into remote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["remote"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, f.Gcs); err != nil {
		return nil, fmt.Errorf("error expanding Gcs into gcs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gcs"] = v
	}
	if v := f.LocalPath; !dcl.IsEmptyValueIndirect(v) {
		m["localPath"] = v
	}
	if v := f.AllowInsecure; !dcl.IsEmptyValueIndirect(v) {
		m["allowInsecure"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource
	}
	r.Remote = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, m["remote"])
	r.Gcs = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, m["gcs"])
	r.LocalPath = dcl.FlattenString(m["localPath"])
	r.AllowInsecure = dcl.FlattenBool(m["allowInsecure"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.Sha256Checksum; !dcl.IsEmptyValueIndirect(v) {
		m["sha256Checksum"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	}
	r.Uri = dcl.FlattenString(m["uri"])
	r.Sha256Checksum = dcl.FlattenString(m["sha256Checksum"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
		m["bucket"] = v
	}
	if v := f.Object; !dcl.IsEmptyValueIndirect(v) {
		m["object"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	}
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Object = dcl.FlattenString(m["object"])
	r.Generation = dcl.FlattenInteger(m["generation"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentFile(c, f.Source); err != nil {
		return nil, fmt.Errorf("error expanding Source into source: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.PullDeps; !dcl.IsEmptyValueIndirect(v) {
		m["pullDeps"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm
	}
	r.Source = flattenOsPolicyAssignmentFile(c, m["source"])
	r.PullDeps = dcl.FlattenBool(m["pullDeps"])

	return r
}

// expandOsPolicyAssignmentFileMap expands the contents of OsPolicyAssignmentFile into a JSON
// request object.
func expandOsPolicyAssignmentFileMap(c *Client, f map[string]OsPolicyAssignmentFile) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentFile(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentFileSlice expands the contents of OsPolicyAssignmentFile into a JSON
// request object.
func expandOsPolicyAssignmentFileSlice(c *Client, f []OsPolicyAssignmentFile) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentFile(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentFileMap flattens the contents of OsPolicyAssignmentFile from a JSON
// response object.
func flattenOsPolicyAssignmentFileMap(c *Client, i interface{}) map[string]OsPolicyAssignmentFile {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentFile{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentFile{}
	}

	items := make(map[string]OsPolicyAssignmentFile)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentFile(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentFileSlice flattens the contents of OsPolicyAssignmentFile from a JSON
// response object.
func flattenOsPolicyAssignmentFileSlice(c *Client, i interface{}) []OsPolicyAssignmentFile {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentFile{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentFile{}
	}

	items := make([]OsPolicyAssignmentFile, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentFile(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentFile expands an instance of OsPolicyAssignmentFile into a JSON
// request object.
func expandOsPolicyAssignmentFile(c *Client, f *OsPolicyAssignmentFile) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentFileRemote(c, f.Remote); err != nil {
		return nil, fmt.Errorf("error expanding Remote into remote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["remote"] = v
	}
	if v, err := expandOsPolicyAssignmentFileGcs(c, f.Gcs); err != nil {
		return nil, fmt.Errorf("error expanding Gcs into gcs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gcs"] = v
	}
	if v := f.LocalPath; !dcl.IsEmptyValueIndirect(v) {
		m["localPath"] = v
	}
	if v := f.AllowInsecure; !dcl.IsEmptyValueIndirect(v) {
		m["allowInsecure"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentFile flattens an instance of OsPolicyAssignmentFile from a JSON
// response object.
func flattenOsPolicyAssignmentFile(c *Client, i interface{}) *OsPolicyAssignmentFile {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentFile{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentFile
	}
	r.Remote = flattenOsPolicyAssignmentFileRemote(c, m["remote"])
	r.Gcs = flattenOsPolicyAssignmentFileGcs(c, m["gcs"])
	r.LocalPath = dcl.FlattenString(m["localPath"])
	r.AllowInsecure = dcl.FlattenBool(m["allowInsecure"])

	return r
}

// expandOsPolicyAssignmentFileRemoteMap expands the contents of OsPolicyAssignmentFileRemote into a JSON
// request object.
func expandOsPolicyAssignmentFileRemoteMap(c *Client, f map[string]OsPolicyAssignmentFileRemote) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentFileRemote(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentFileRemoteSlice expands the contents of OsPolicyAssignmentFileRemote into a JSON
// request object.
func expandOsPolicyAssignmentFileRemoteSlice(c *Client, f []OsPolicyAssignmentFileRemote) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentFileRemote(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentFileRemoteMap flattens the contents of OsPolicyAssignmentFileRemote from a JSON
// response object.
func flattenOsPolicyAssignmentFileRemoteMap(c *Client, i interface{}) map[string]OsPolicyAssignmentFileRemote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentFileRemote{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentFileRemote{}
	}

	items := make(map[string]OsPolicyAssignmentFileRemote)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentFileRemote(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentFileRemoteSlice flattens the contents of OsPolicyAssignmentFileRemote from a JSON
// response object.
func flattenOsPolicyAssignmentFileRemoteSlice(c *Client, i interface{}) []OsPolicyAssignmentFileRemote {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentFileRemote{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentFileRemote{}
	}

	items := make([]OsPolicyAssignmentFileRemote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentFileRemote(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentFileRemote expands an instance of OsPolicyAssignmentFileRemote into a JSON
// request object.
func expandOsPolicyAssignmentFileRemote(c *Client, f *OsPolicyAssignmentFileRemote) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.Sha256Checksum; !dcl.IsEmptyValueIndirect(v) {
		m["sha256Checksum"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentFileRemote flattens an instance of OsPolicyAssignmentFileRemote from a JSON
// response object.
func flattenOsPolicyAssignmentFileRemote(c *Client, i interface{}) *OsPolicyAssignmentFileRemote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentFileRemote{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentFileRemote
	}
	r.Uri = dcl.FlattenString(m["uri"])
	r.Sha256Checksum = dcl.FlattenString(m["sha256Checksum"])

	return r
}

// expandOsPolicyAssignmentFileGcsMap expands the contents of OsPolicyAssignmentFileGcs into a JSON
// request object.
func expandOsPolicyAssignmentFileGcsMap(c *Client, f map[string]OsPolicyAssignmentFileGcs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentFileGcs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentFileGcsSlice expands the contents of OsPolicyAssignmentFileGcs into a JSON
// request object.
func expandOsPolicyAssignmentFileGcsSlice(c *Client, f []OsPolicyAssignmentFileGcs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentFileGcs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentFileGcsMap flattens the contents of OsPolicyAssignmentFileGcs from a JSON
// response object.
func flattenOsPolicyAssignmentFileGcsMap(c *Client, i interface{}) map[string]OsPolicyAssignmentFileGcs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentFileGcs{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentFileGcs{}
	}

	items := make(map[string]OsPolicyAssignmentFileGcs)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentFileGcs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentFileGcsSlice flattens the contents of OsPolicyAssignmentFileGcs from a JSON
// response object.
func flattenOsPolicyAssignmentFileGcsSlice(c *Client, i interface{}) []OsPolicyAssignmentFileGcs {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentFileGcs{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentFileGcs{}
	}

	items := make([]OsPolicyAssignmentFileGcs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentFileGcs(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentFileGcs expands an instance of OsPolicyAssignmentFileGcs into a JSON
// request object.
func expandOsPolicyAssignmentFileGcs(c *Client, f *OsPolicyAssignmentFileGcs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
		m["bucket"] = v
	}
	if v := f.Object; !dcl.IsEmptyValueIndirect(v) {
		m["object"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentFileGcs flattens an instance of OsPolicyAssignmentFileGcs from a JSON
// response object.
func flattenOsPolicyAssignmentFileGcs(c *Client, i interface{}) *OsPolicyAssignmentFileGcs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentFileGcs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentFileGcs
	}
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Object = dcl.FlattenString(m["object"])
	r.Generation = dcl.FlattenInteger(m["generation"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentFile(c, f.Source); err != nil {
		return nil, fmt.Errorf("error expanding Source into source: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi
	}
	r.Source = flattenOsPolicyAssignmentFile(c, m["source"])
	r.Properties = dcl.FlattenStringSlice(m["properties"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositorySlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, f.Apt); err != nil {
		return nil, fmt.Errorf("error expanding Apt into apt: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apt"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, f.Yum); err != nil {
		return nil, fmt.Errorf("error expanding Yum into yum: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yum"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, f.Zypper); err != nil {
		return nil, fmt.Errorf("error expanding Zypper into zypper: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zypper"] = v
	}
	if v, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, f.Goo); err != nil {
		return nil, fmt.Errorf("error expanding Goo into goo: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["goo"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	}
	r.Apt = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, m["apt"])
	r.Yum = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, m["yum"])
	r.Zypper = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, m["zypper"])
	r.Goo = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, m["goo"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArchiveType; !dcl.IsEmptyValueIndirect(v) {
		m["archiveType"] = v
	}
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.Distribution; !dcl.IsEmptyValueIndirect(v) {
		m["distribution"] = v
	}
	if v := f.Components; !dcl.IsEmptyValueIndirect(v) {
		m["components"] = v
	}
	if v := f.GpgKey; !dcl.IsEmptyValueIndirect(v) {
		m["gpgKey"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	}
	r.ArchiveType = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(m["archiveType"])
	r.Uri = dcl.FlattenString(m["uri"])
	r.Distribution = dcl.FlattenString(m["distribution"])
	r.Components = dcl.FlattenStringSlice(m["components"])
	r.GpgKey = dcl.FlattenString(m["gpgKey"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.BaseUrl; !dcl.IsEmptyValueIndirect(v) {
		m["baseUrl"] = v
	}
	if v := f.GpgKeys; !dcl.IsEmptyValueIndirect(v) {
		m["gpgKeys"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	}
	r.Id = dcl.FlattenString(m["id"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.BaseUrl = dcl.FlattenString(m["baseUrl"])
	r.GpgKeys = dcl.FlattenStringSlice(m["gpgKeys"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.BaseUrl; !dcl.IsEmptyValueIndirect(v) {
		m["baseUrl"] = v
	}
	if v := f.GpgKeys; !dcl.IsEmptyValueIndirect(v) {
		m["gpgKeys"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	}
	r.Id = dcl.FlattenString(m["id"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.BaseUrl = dcl.FlattenString(m["baseUrl"])
	r.GpgKeys = dcl.FlattenStringSlice(m["gpgKeys"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentExec(c, f.Validate); err != nil {
		return nil, fmt.Errorf("error expanding Validate into validate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["validate"] = v
	}
	if v, err := expandOsPolicyAssignmentExec(c, f.Enforce); err != nil {
		return nil, fmt.Errorf("error expanding Enforce into enforce: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["enforce"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	}
	r.Validate = flattenOsPolicyAssignmentExec(c, m["validate"])
	r.Enforce = flattenOsPolicyAssignmentExec(c, m["enforce"])

	return r
}

// expandOsPolicyAssignmentExecMap expands the contents of OsPolicyAssignmentExec into a JSON
// request object.
func expandOsPolicyAssignmentExecMap(c *Client, f map[string]OsPolicyAssignmentExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentExecSlice expands the contents of OsPolicyAssignmentExec into a JSON
// request object.
func expandOsPolicyAssignmentExecSlice(c *Client, f []OsPolicyAssignmentExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentExecMap flattens the contents of OsPolicyAssignmentExec from a JSON
// response object.
func flattenOsPolicyAssignmentExecMap(c *Client, i interface{}) map[string]OsPolicyAssignmentExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentExec{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentExec{}
	}

	items := make(map[string]OsPolicyAssignmentExec)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentExecSlice flattens the contents of OsPolicyAssignmentExec from a JSON
// response object.
func flattenOsPolicyAssignmentExecSlice(c *Client, i interface{}) []OsPolicyAssignmentExec {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentExec{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentExec{}
	}

	items := make([]OsPolicyAssignmentExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentExec expands an instance of OsPolicyAssignmentExec into a JSON
// request object.
func expandOsPolicyAssignmentExec(c *Client, f *OsPolicyAssignmentExec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentFile(c, f.File); err != nil {
		return nil, fmt.Errorf("error expanding File into file: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["file"] = v
	}
	if v := f.Script; !dcl.IsEmptyValueIndirect(v) {
		m["script"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.Interpreter; !dcl.IsEmptyValueIndirect(v) {
		m["interpreter"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentExec flattens an instance of OsPolicyAssignmentExec from a JSON
// response object.
func flattenOsPolicyAssignmentExec(c *Client, i interface{}) *OsPolicyAssignmentExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentExec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentExec
	}
	r.File = flattenOsPolicyAssignmentFile(c, m["file"])
	r.Script = dcl.FlattenString(m["script"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Interpreter = flattenOsPolicyAssignmentExecInterpreterEnum(m["interpreter"])

	return r
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileMap expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileMap(c *Client, f map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSlice expands the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSlice(c *Client, f []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileMap flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileMap(c *Client, i interface{}) map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{}
	}

	items := make(map[string]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile expands an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile into a JSON
// request object.
func expandOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c *Client, f *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentFile(c, f.File); err != nil {
		return nil, fmt.Errorf("error expanding File into file: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["file"] = v
	}
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Permissions; !dcl.IsEmptyValueIndirect(v) {
		m["permissions"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile flattens an instance of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(c *Client, i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	}
	r.File = flattenOsPolicyAssignmentFile(c, m["file"])
	r.Content = dcl.FlattenString(m["content"])
	r.Path = dcl.FlattenString(m["path"])
	r.State = flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(m["state"])
	r.Permissions = dcl.FlattenString(m["permissions"])

	return r
}

// expandOsPolicyAssignmentInstanceFilterMap expands the contents of OsPolicyAssignmentInstanceFilter into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterMap(c *Client, f map[string]OsPolicyAssignmentInstanceFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentInstanceFilterSlice expands the contents of OsPolicyAssignmentInstanceFilter into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterSlice(c *Client, f []OsPolicyAssignmentInstanceFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentInstanceFilterMap flattens the contents of OsPolicyAssignmentInstanceFilter from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterMap(c *Client, i interface{}) map[string]OsPolicyAssignmentInstanceFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentInstanceFilter{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentInstanceFilter{}
	}

	items := make(map[string]OsPolicyAssignmentInstanceFilter)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentInstanceFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentInstanceFilterSlice flattens the contents of OsPolicyAssignmentInstanceFilter from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterSlice(c *Client, i interface{}) []OsPolicyAssignmentInstanceFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentInstanceFilter{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentInstanceFilter{}
	}

	items := make([]OsPolicyAssignmentInstanceFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentInstanceFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentInstanceFilter expands an instance of OsPolicyAssignmentInstanceFilter into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilter(c *Client, f *OsPolicyAssignmentInstanceFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.All; !dcl.IsEmptyValueIndirect(v) {
		m["all"] = v
	}
	if v := f.OsShortNames; !dcl.IsEmptyValueIndirect(v) {
		m["osShortNames"] = v
	}
	if v, err := expandOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c, f.InclusionLabels); err != nil {
		return nil, fmt.Errorf("error expanding InclusionLabels into inclusionLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["inclusionLabels"] = v
	}
	if v, err := expandOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c, f.ExclusionLabels); err != nil {
		return nil, fmt.Errorf("error expanding ExclusionLabels into exclusionLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exclusionLabels"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentInstanceFilter flattens an instance of OsPolicyAssignmentInstanceFilter from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilter(c *Client, i interface{}) *OsPolicyAssignmentInstanceFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentInstanceFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentInstanceFilter
	}
	r.All = dcl.FlattenBool(m["all"])
	r.OsShortNames = dcl.FlattenStringSlice(m["osShortNames"])
	r.InclusionLabels = flattenOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c, m["inclusionLabels"])
	r.ExclusionLabels = flattenOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c, m["exclusionLabels"])

	return r
}

// expandOsPolicyAssignmentInstanceFilterInclusionLabelsMap expands the contents of OsPolicyAssignmentInstanceFilterInclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterInclusionLabelsMap(c *Client, f map[string]OsPolicyAssignmentInstanceFilterInclusionLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilterInclusionLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentInstanceFilterInclusionLabelsSlice expands the contents of OsPolicyAssignmentInstanceFilterInclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c *Client, f []OsPolicyAssignmentInstanceFilterInclusionLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilterInclusionLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentInstanceFilterInclusionLabelsMap flattens the contents of OsPolicyAssignmentInstanceFilterInclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterInclusionLabelsMap(c *Client, i interface{}) map[string]OsPolicyAssignmentInstanceFilterInclusionLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentInstanceFilterInclusionLabels{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentInstanceFilterInclusionLabels{}
	}

	items := make(map[string]OsPolicyAssignmentInstanceFilterInclusionLabels)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentInstanceFilterInclusionLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentInstanceFilterInclusionLabelsSlice flattens the contents of OsPolicyAssignmentInstanceFilterInclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterInclusionLabelsSlice(c *Client, i interface{}) []OsPolicyAssignmentInstanceFilterInclusionLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentInstanceFilterInclusionLabels{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentInstanceFilterInclusionLabels{}
	}

	items := make([]OsPolicyAssignmentInstanceFilterInclusionLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentInstanceFilterInclusionLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentInstanceFilterInclusionLabels expands an instance of OsPolicyAssignmentInstanceFilterInclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterInclusionLabels(c *Client, f *OsPolicyAssignmentInstanceFilterInclusionLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentInstanceFilterInclusionLabels flattens an instance of OsPolicyAssignmentInstanceFilterInclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterInclusionLabels(c *Client, i interface{}) *OsPolicyAssignmentInstanceFilterInclusionLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentInstanceFilterInclusionLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentInstanceFilterInclusionLabels
	}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandOsPolicyAssignmentInstanceFilterExclusionLabelsMap expands the contents of OsPolicyAssignmentInstanceFilterExclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterExclusionLabelsMap(c *Client, f map[string]OsPolicyAssignmentInstanceFilterExclusionLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilterExclusionLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentInstanceFilterExclusionLabelsSlice expands the contents of OsPolicyAssignmentInstanceFilterExclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c *Client, f []OsPolicyAssignmentInstanceFilterExclusionLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentInstanceFilterExclusionLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentInstanceFilterExclusionLabelsMap flattens the contents of OsPolicyAssignmentInstanceFilterExclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterExclusionLabelsMap(c *Client, i interface{}) map[string]OsPolicyAssignmentInstanceFilterExclusionLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentInstanceFilterExclusionLabels{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentInstanceFilterExclusionLabels{}
	}

	items := make(map[string]OsPolicyAssignmentInstanceFilterExclusionLabels)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentInstanceFilterExclusionLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentInstanceFilterExclusionLabelsSlice flattens the contents of OsPolicyAssignmentInstanceFilterExclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterExclusionLabelsSlice(c *Client, i interface{}) []OsPolicyAssignmentInstanceFilterExclusionLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentInstanceFilterExclusionLabels{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentInstanceFilterExclusionLabels{}
	}

	items := make([]OsPolicyAssignmentInstanceFilterExclusionLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentInstanceFilterExclusionLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentInstanceFilterExclusionLabels expands an instance of OsPolicyAssignmentInstanceFilterExclusionLabels into a JSON
// request object.
func expandOsPolicyAssignmentInstanceFilterExclusionLabels(c *Client, f *OsPolicyAssignmentInstanceFilterExclusionLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentInstanceFilterExclusionLabels flattens an instance of OsPolicyAssignmentInstanceFilterExclusionLabels from a JSON
// response object.
func flattenOsPolicyAssignmentInstanceFilterExclusionLabels(c *Client, i interface{}) *OsPolicyAssignmentInstanceFilterExclusionLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentInstanceFilterExclusionLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentInstanceFilterExclusionLabels
	}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandOsPolicyAssignmentRolloutMap expands the contents of OsPolicyAssignmentRollout into a JSON
// request object.
func expandOsPolicyAssignmentRolloutMap(c *Client, f map[string]OsPolicyAssignmentRollout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentRollout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentRolloutSlice expands the contents of OsPolicyAssignmentRollout into a JSON
// request object.
func expandOsPolicyAssignmentRolloutSlice(c *Client, f []OsPolicyAssignmentRollout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentRollout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentRolloutMap flattens the contents of OsPolicyAssignmentRollout from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutMap(c *Client, i interface{}) map[string]OsPolicyAssignmentRollout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentRollout{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentRollout{}
	}

	items := make(map[string]OsPolicyAssignmentRollout)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentRollout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentRolloutSlice flattens the contents of OsPolicyAssignmentRollout from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutSlice(c *Client, i interface{}) []OsPolicyAssignmentRollout {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentRollout{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentRollout{}
	}

	items := make([]OsPolicyAssignmentRollout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentRollout(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentRollout expands an instance of OsPolicyAssignmentRollout into a JSON
// request object.
func expandOsPolicyAssignmentRollout(c *Client, f *OsPolicyAssignmentRollout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOsPolicyAssignmentRolloutDisruptionBudget(c, f.DisruptionBudget); err != nil {
		return nil, fmt.Errorf("error expanding DisruptionBudget into disruptionBudget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["disruptionBudget"] = v
	}
	if v := f.MinWaitDuration; !dcl.IsEmptyValueIndirect(v) {
		m["minWaitDuration"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentRollout flattens an instance of OsPolicyAssignmentRollout from a JSON
// response object.
func flattenOsPolicyAssignmentRollout(c *Client, i interface{}) *OsPolicyAssignmentRollout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentRollout{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentRollout
	}
	r.DisruptionBudget = flattenOsPolicyAssignmentRolloutDisruptionBudget(c, m["disruptionBudget"])
	r.MinWaitDuration = dcl.FlattenString(m["minWaitDuration"])

	return r
}

// expandOsPolicyAssignmentRolloutDisruptionBudgetMap expands the contents of OsPolicyAssignmentRolloutDisruptionBudget into a JSON
// request object.
func expandOsPolicyAssignmentRolloutDisruptionBudgetMap(c *Client, f map[string]OsPolicyAssignmentRolloutDisruptionBudget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOsPolicyAssignmentRolloutDisruptionBudget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOsPolicyAssignmentRolloutDisruptionBudgetSlice expands the contents of OsPolicyAssignmentRolloutDisruptionBudget into a JSON
// request object.
func expandOsPolicyAssignmentRolloutDisruptionBudgetSlice(c *Client, f []OsPolicyAssignmentRolloutDisruptionBudget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOsPolicyAssignmentRolloutDisruptionBudget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOsPolicyAssignmentRolloutDisruptionBudgetMap flattens the contents of OsPolicyAssignmentRolloutDisruptionBudget from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutDisruptionBudgetMap(c *Client, i interface{}) map[string]OsPolicyAssignmentRolloutDisruptionBudget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OsPolicyAssignmentRolloutDisruptionBudget{}
	}

	if len(a) == 0 {
		return map[string]OsPolicyAssignmentRolloutDisruptionBudget{}
	}

	items := make(map[string]OsPolicyAssignmentRolloutDisruptionBudget)
	for k, item := range a {
		items[k] = *flattenOsPolicyAssignmentRolloutDisruptionBudget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOsPolicyAssignmentRolloutDisruptionBudgetSlice flattens the contents of OsPolicyAssignmentRolloutDisruptionBudget from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutDisruptionBudgetSlice(c *Client, i interface{}) []OsPolicyAssignmentRolloutDisruptionBudget {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentRolloutDisruptionBudget{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentRolloutDisruptionBudget{}
	}

	items := make([]OsPolicyAssignmentRolloutDisruptionBudget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentRolloutDisruptionBudget(c, item.(map[string]interface{})))
	}

	return items
}

// expandOsPolicyAssignmentRolloutDisruptionBudget expands an instance of OsPolicyAssignmentRolloutDisruptionBudget into a JSON
// request object.
func expandOsPolicyAssignmentRolloutDisruptionBudget(c *Client, f *OsPolicyAssignmentRolloutDisruptionBudget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fixed; !dcl.IsEmptyValueIndirect(v) {
		m["fixed"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenOsPolicyAssignmentRolloutDisruptionBudget flattens an instance of OsPolicyAssignmentRolloutDisruptionBudget from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutDisruptionBudget(c *Client, i interface{}) *OsPolicyAssignmentRolloutDisruptionBudget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OsPolicyAssignmentRolloutDisruptionBudget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOsPolicyAssignmentRolloutDisruptionBudget
	}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])

	return r
}

// flattenOsPolicyAssignmentOsPoliciesModeEnumSlice flattens the contents of OsPolicyAssignmentOsPoliciesModeEnum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesModeEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesModeEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesModeEnum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesModeEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesModeEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentOsPoliciesModeEnum with the same value as that string.
func flattenOsPolicyAssignmentOsPoliciesModeEnum(i interface{}) *OsPolicyAssignmentOsPoliciesModeEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentOsPoliciesModeEnumRef("")
	}

	return OsPolicyAssignmentOsPoliciesModeEnumRef(s)
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum with the same value as that string.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumRef("")
	}

	return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumRef(s)
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum with the same value as that string.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumRef("")
	}

	return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumRef(s)
}

// flattenOsPolicyAssignmentExecInterpreterEnumSlice flattens the contents of OsPolicyAssignmentExecInterpreterEnum from a JSON
// response object.
func flattenOsPolicyAssignmentExecInterpreterEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentExecInterpreterEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentExecInterpreterEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentExecInterpreterEnum{}
	}

	items := make([]OsPolicyAssignmentExecInterpreterEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentExecInterpreterEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentExecInterpreterEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentExecInterpreterEnum with the same value as that string.
func flattenOsPolicyAssignmentExecInterpreterEnum(i interface{}) *OsPolicyAssignmentExecInterpreterEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentExecInterpreterEnumRef("")
	}

	return OsPolicyAssignmentExecInterpreterEnumRef(s)
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumSlice flattens the contents of OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum from a JSON
// response object.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum{}
	}

	items := make([]OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum with the same value as that string.
func flattenOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(i interface{}) *OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumRef("")
	}

	return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumRef(s)
}

// flattenOsPolicyAssignmentRolloutStateEnumSlice flattens the contents of OsPolicyAssignmentRolloutStateEnum from a JSON
// response object.
func flattenOsPolicyAssignmentRolloutStateEnumSlice(c *Client, i interface{}) []OsPolicyAssignmentRolloutStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OsPolicyAssignmentRolloutStateEnum{}
	}

	if len(a) == 0 {
		return []OsPolicyAssignmentRolloutStateEnum{}
	}

	items := make([]OsPolicyAssignmentRolloutStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOsPolicyAssignmentRolloutStateEnum(item.(interface{})))
	}

	return items
}

// flattenOsPolicyAssignmentRolloutStateEnum asserts that an interface is a string, and returns a
// pointer to a *OsPolicyAssignmentRolloutStateEnum with the same value as that string.
func flattenOsPolicyAssignmentRolloutStateEnum(i interface{}) *OsPolicyAssignmentRolloutStateEnum {
	s, ok := i.(string)
	if !ok {
		return OsPolicyAssignmentRolloutStateEnumRef("")
	}

	return OsPolicyAssignmentRolloutStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *OsPolicyAssignment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalOsPolicyAssignment(b, c)
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

type osPolicyAssignmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         osPolicyAssignmentApiOperation
}

func convertFieldDiffToOsPolicyAssignmentOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]osPolicyAssignmentDiff, error) {
	var diffs []osPolicyAssignmentDiff
	for _, op := range ops {
		diff := osPolicyAssignmentDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToosPolicyAssignmentApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToosPolicyAssignmentApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (osPolicyAssignmentApiOperation, error) {
	switch op {

	case "updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation":
		return &updateOsPolicyAssignmentUpdateOSPolicyAssignmentOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
