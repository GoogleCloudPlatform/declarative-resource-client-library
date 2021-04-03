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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *GuestPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Assignment) {
		if err := r.Assignment.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GuestPolicyAssignment) validate() error {
	return nil
}
func (r *GuestPolicyAssignmentGroupLabels) validate() error {
	return nil
}
func (r *GuestPolicyAssignmentOsTypes) validate() error {
	return nil
}
func (r *GuestPolicyPackages) validate() error {
	return nil
}
func (r *GuestPolicyPackageRepositories) validate() error {
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
func (r *GuestPolicyPackageRepositoriesApt) validate() error {
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "distribution"); err != nil {
		return err
	}
	return nil
}
func (r *GuestPolicyPackageRepositoriesYum) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if err := dcl.Required(r, "baseUrl"); err != nil {
		return err
	}
	return nil
}
func (r *GuestPolicyPackageRepositoriesZypper) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	if err := dcl.Required(r, "baseUrl"); err != nil {
		return err
	}
	return nil
}
func (r *GuestPolicyPackageRepositoriesGoo) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *GuestPolicyRecipes) validate() error {
	return nil
}
func (r *GuestPolicyRecipesArtifacts) validate() error {
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
func (r *GuestPolicyRecipesArtifactsRemote) validate() error {
	return nil
}
func (r *GuestPolicyRecipesArtifactsGcs) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallSteps) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FileCopy) {
		if err := r.FileCopy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ArchiveExtraction) {
		if err := r.ArchiveExtraction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MsiInstallation) {
		if err := r.MsiInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DpkgInstallation) {
		if err := r.DpkgInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RpmInstallation) {
		if err := r.RpmInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FileExec) {
		if err := r.FileExec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ScriptRun) {
		if err := r.ScriptRun.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GuestPolicyRecipesInstallStepsFileCopy) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsMsiInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsRpmInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsFileExec) validate() error {
	return nil
}
func (r *GuestPolicyRecipesInstallStepsScriptRun) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateSteps) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FileCopy) {
		if err := r.FileCopy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ArchiveExtraction) {
		if err := r.ArchiveExtraction.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MsiInstallation) {
		if err := r.MsiInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DpkgInstallation) {
		if err := r.DpkgInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RpmInstallation) {
		if err := r.RpmInstallation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FileExec) {
		if err := r.FileExec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ScriptRun) {
		if err := r.ScriptRun.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsFileCopy) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsFileExec) validate() error {
	return nil
}
func (r *GuestPolicyRecipesUpdateStepsScriptRun) validate() error {
	return nil
}

func guestPolicyGetURL(userBasePath string, r *GuestPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/guestPolicies/{{name}}", "https://osconfig.googleapis.com/v1beta/", userBasePath, params), nil
}

func guestPolicyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/guestPolicies", "https://osconfig.googleapis.com/v1beta/", userBasePath, params), nil

}

func guestPolicyCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/guestPolicies?guestPolicyId={{name}}", "https://osconfig.googleapis.com/v1beta/", userBasePath, params), nil

}

func guestPolicyDeleteURL(userBasePath string, r *GuestPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/guestPolicies/{{name}}", "https://osconfig.googleapis.com/v1beta/", userBasePath, params), nil
}

// guestPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type guestPolicyApiOperation interface {
	do(context.Context, *GuestPolicy, *Client) error
}

// newUpdateGuestPolicyUpdateGuestPolicyRequest creates a request for an
// GuestPolicy resource's UpdateGuestPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGuestPolicyUpdateGuestPolicyRequest(ctx context.Context, f *GuestPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandGuestPolicyAssignment(c, f.Assignment); err != nil {
		return nil, fmt.Errorf("error expanding Assignment into assignment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["assignment"] = v
	}
	if v, err := expandGuestPolicyPackagesSlice(c, f.Packages); err != nil {
		return nil, fmt.Errorf("error expanding Packages into packages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["packages"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesSlice(c, f.PackageRepositories); err != nil {
		return nil, fmt.Errorf("error expanding PackageRepositories into packageRepositories: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["packageRepositories"] = v
	}
	if v, err := expandGuestPolicyRecipesSlice(c, f.Recipes); err != nil {
		return nil, fmt.Errorf("error expanding Recipes into recipes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["recipes"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		req["etag"] = v
	}
	b, err := c.getGuestPolicyRaw(ctx, f.urlNormalized())
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

// marshalUpdateGuestPolicyUpdateGuestPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateGuestPolicyUpdateGuestPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGuestPolicyUpdateGuestPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGuestPolicyUpdateGuestPolicyOperation) do(ctx context.Context, r *GuestPolicy, c *Client) error {
	_, err := c.GetGuestPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateGuestPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateGuestPolicyUpdateGuestPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateGuestPolicyUpdateGuestPolicyRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listGuestPolicyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := guestPolicyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GuestPolicyMaxPage {
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

type listGuestPolicyOperation struct {
	GuestPolicies []map[string]interface{} `json:"guestPolicies"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listGuestPolicy(ctx context.Context, project, pageToken string, pageSize int32) ([]*GuestPolicy, string, error) {
	b, err := c.listGuestPolicyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGuestPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*GuestPolicy
	for _, v := range m.GuestPolicies {
		res := flattenGuestPolicy(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGuestPolicy(ctx context.Context, f func(*GuestPolicy) bool, resources []*GuestPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGuestPolicy(ctx, res)
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

type deleteGuestPolicyOperation struct{}

func (op *deleteGuestPolicyOperation) do(ctx context.Context, r *GuestPolicy, c *Client) error {

	_, err := c.GetGuestPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("GuestPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetGuestPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := guestPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete GuestPolicy: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetGuestPolicy(ctx, r.urlNormalized())
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
type createGuestPolicyOperation struct {
	response map[string]interface{}
}

func (op *createGuestPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGuestPolicyOperation) do(ctx context.Context, r *GuestPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, name := r.createFields()
	u, err := guestPolicyCreateURL(c.Config.BasePath, project, name)

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

	if _, err := c.GetGuestPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGuestPolicyRaw(ctx context.Context, r *GuestPolicy) ([]byte, error) {

	u, err := guestPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) guestPolicyDiffsForRawDesired(ctx context.Context, rawDesired *GuestPolicy, opts ...dcl.ApplyOption) (initial, desired *GuestPolicy, diffs []guestPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *GuestPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*GuestPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected GuestPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGuestPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a GuestPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve GuestPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that GuestPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGuestPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for GuestPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for GuestPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGuestPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for GuestPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGuestPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for GuestPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGuestPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGuestPolicyInitialState(rawInitial, rawDesired *GuestPolicy) (*GuestPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGuestPolicyDesiredState(rawDesired, rawInitial *GuestPolicy, opts ...dcl.ApplyOption) (*GuestPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Assignment = canonicalizeGuestPolicyAssignment(rawDesired.Assignment, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	rawDesired.Assignment = canonicalizeGuestPolicyAssignment(rawDesired.Assignment, rawInitial.Assignment, opts...)
	if dcl.IsZeroValue(rawDesired.Packages) {
		rawDesired.Packages = rawInitial.Packages
	}
	if dcl.IsZeroValue(rawDesired.PackageRepositories) {
		rawDesired.PackageRepositories = rawInitial.PackageRepositories
	}
	if dcl.IsZeroValue(rawDesired.Recipes) {
		rawDesired.Recipes = rawInitial.Recipes
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeGuestPolicyNewState(c *Client, rawNew, rawDesired *GuestPolicy) (*GuestPolicy, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Assignment) && dcl.IsEmptyValueIndirect(rawDesired.Assignment) {
		rawNew.Assignment = rawDesired.Assignment
	} else {
		rawNew.Assignment = canonicalizeNewGuestPolicyAssignment(c, rawDesired.Assignment, rawNew.Assignment)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Packages) && dcl.IsEmptyValueIndirect(rawDesired.Packages) {
		rawNew.Packages = rawDesired.Packages
	} else {
		rawNew.Packages = canonicalizeNewGuestPolicyPackagesSlice(c, rawDesired.Packages, rawNew.Packages)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PackageRepositories) && dcl.IsEmptyValueIndirect(rawDesired.PackageRepositories) {
		rawNew.PackageRepositories = rawDesired.PackageRepositories
	} else {
		rawNew.PackageRepositories = canonicalizeNewGuestPolicyPackageRepositoriesSlice(c, rawDesired.PackageRepositories, rawNew.PackageRepositories)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Recipes) && dcl.IsEmptyValueIndirect(rawDesired.Recipes) {
		rawNew.Recipes = rawDesired.Recipes
	} else {
		rawNew.Recipes = canonicalizeNewGuestPolicyRecipesSlice(c, rawDesired.Recipes, rawNew.Recipes)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeGuestPolicyAssignment(des, initial *GuestPolicyAssignment, opts ...dcl.ApplyOption) *GuestPolicyAssignment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.GroupLabels) {
		des.GroupLabels = initial.GroupLabels
	}
	if dcl.IsZeroValue(des.Zones) {
		des.Zones = initial.Zones
	}
	if dcl.IsZeroValue(des.Instances) {
		des.Instances = initial.Instances
	}
	if dcl.IsZeroValue(des.InstanceNamePrefixes) {
		des.InstanceNamePrefixes = initial.InstanceNamePrefixes
	}
	if dcl.IsZeroValue(des.OsTypes) {
		des.OsTypes = initial.OsTypes
	}

	return des
}

func canonicalizeNewGuestPolicyAssignment(c *Client, des, nw *GuestPolicyAssignment) *GuestPolicyAssignment {
	if des == nil || nw == nil {
		return nw
	}

	nw.GroupLabels = canonicalizeNewGuestPolicyAssignmentGroupLabelsSlice(c, des.GroupLabels, nw.GroupLabels)
	nw.OsTypes = canonicalizeNewGuestPolicyAssignmentOsTypesSlice(c, des.OsTypes, nw.OsTypes)

	return nw
}

func canonicalizeNewGuestPolicyAssignmentSet(c *Client, des, nw []GuestPolicyAssignment) []GuestPolicyAssignment {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyAssignment
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyAssignment(c, &d, &n) {
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

func canonicalizeNewGuestPolicyAssignmentSlice(c *Client, des, nw []GuestPolicyAssignment) []GuestPolicyAssignment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyAssignment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyAssignment(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyAssignmentGroupLabels(des, initial *GuestPolicyAssignmentGroupLabels, opts ...dcl.ApplyOption) *GuestPolicyAssignmentGroupLabels {
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

func canonicalizeNewGuestPolicyAssignmentGroupLabels(c *Client, des, nw *GuestPolicyAssignmentGroupLabels) *GuestPolicyAssignmentGroupLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewGuestPolicyAssignmentGroupLabelsSet(c *Client, des, nw []GuestPolicyAssignmentGroupLabels) []GuestPolicyAssignmentGroupLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyAssignmentGroupLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyAssignmentGroupLabels(c, &d, &n) {
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

func canonicalizeNewGuestPolicyAssignmentGroupLabelsSlice(c *Client, des, nw []GuestPolicyAssignmentGroupLabels) []GuestPolicyAssignmentGroupLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyAssignmentGroupLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyAssignmentGroupLabels(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyAssignmentOsTypes(des, initial *GuestPolicyAssignmentOsTypes, opts ...dcl.ApplyOption) *GuestPolicyAssignmentOsTypes {
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
	if dcl.StringCanonicalize(des.OsArchitecture, initial.OsArchitecture) || dcl.IsZeroValue(des.OsArchitecture) {
		des.OsArchitecture = initial.OsArchitecture
	}

	return des
}

func canonicalizeNewGuestPolicyAssignmentOsTypes(c *Client, des, nw *GuestPolicyAssignmentOsTypes) *GuestPolicyAssignmentOsTypes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.OsShortName, nw.OsShortName) {
		nw.OsShortName = des.OsShortName
	}
	if dcl.StringCanonicalize(des.OsVersion, nw.OsVersion) {
		nw.OsVersion = des.OsVersion
	}
	if dcl.StringCanonicalize(des.OsArchitecture, nw.OsArchitecture) {
		nw.OsArchitecture = des.OsArchitecture
	}

	return nw
}

func canonicalizeNewGuestPolicyAssignmentOsTypesSet(c *Client, des, nw []GuestPolicyAssignmentOsTypes) []GuestPolicyAssignmentOsTypes {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyAssignmentOsTypes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyAssignmentOsTypes(c, &d, &n) {
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

func canonicalizeNewGuestPolicyAssignmentOsTypesSlice(c *Client, des, nw []GuestPolicyAssignmentOsTypes) []GuestPolicyAssignmentOsTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyAssignmentOsTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyAssignmentOsTypes(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackages(des, initial *GuestPolicyPackages, opts ...dcl.ApplyOption) *GuestPolicyPackages {
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
	if dcl.IsZeroValue(des.DesiredState) {
		des.DesiredState = initial.DesiredState
	}
	if dcl.IsZeroValue(des.Manager) {
		des.Manager = initial.Manager
	}

	return des
}

func canonicalizeNewGuestPolicyPackages(c *Client, des, nw *GuestPolicyPackages) *GuestPolicyPackages {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewGuestPolicyPackagesSet(c *Client, des, nw []GuestPolicyPackages) []GuestPolicyPackages {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackages(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackagesSlice(c *Client, des, nw []GuestPolicyPackages) []GuestPolicyPackages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackages(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackageRepositories(des, initial *GuestPolicyPackageRepositories, opts ...dcl.ApplyOption) *GuestPolicyPackageRepositories {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Apt = canonicalizeGuestPolicyPackageRepositoriesApt(des.Apt, initial.Apt, opts...)
	des.Yum = canonicalizeGuestPolicyPackageRepositoriesYum(des.Yum, initial.Yum, opts...)
	des.Zypper = canonicalizeGuestPolicyPackageRepositoriesZypper(des.Zypper, initial.Zypper, opts...)
	des.Goo = canonicalizeGuestPolicyPackageRepositoriesGoo(des.Goo, initial.Goo, opts...)

	return des
}

func canonicalizeNewGuestPolicyPackageRepositories(c *Client, des, nw *GuestPolicyPackageRepositories) *GuestPolicyPackageRepositories {
	if des == nil || nw == nil {
		return nw
	}

	nw.Apt = canonicalizeNewGuestPolicyPackageRepositoriesApt(c, des.Apt, nw.Apt)
	nw.Yum = canonicalizeNewGuestPolicyPackageRepositoriesYum(c, des.Yum, nw.Yum)
	nw.Zypper = canonicalizeNewGuestPolicyPackageRepositoriesZypper(c, des.Zypper, nw.Zypper)
	nw.Goo = canonicalizeNewGuestPolicyPackageRepositoriesGoo(c, des.Goo, nw.Goo)

	return nw
}

func canonicalizeNewGuestPolicyPackageRepositoriesSet(c *Client, des, nw []GuestPolicyPackageRepositories) []GuestPolicyPackageRepositories {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositories
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackageRepositories(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackageRepositoriesSlice(c *Client, des, nw []GuestPolicyPackageRepositories) []GuestPolicyPackageRepositories {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackageRepositories
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackageRepositories(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackageRepositoriesApt(des, initial *GuestPolicyPackageRepositoriesApt, opts ...dcl.ApplyOption) *GuestPolicyPackageRepositoriesApt {
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

func canonicalizeNewGuestPolicyPackageRepositoriesApt(c *Client, des, nw *GuestPolicyPackageRepositoriesApt) *GuestPolicyPackageRepositoriesApt {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Distribution, nw.Distribution) {
		nw.Distribution = des.Distribution
	}
	if dcl.StringCanonicalize(des.GpgKey, nw.GpgKey) {
		nw.GpgKey = des.GpgKey
	}

	return nw
}

func canonicalizeNewGuestPolicyPackageRepositoriesAptSet(c *Client, des, nw []GuestPolicyPackageRepositoriesApt) []GuestPolicyPackageRepositoriesApt {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositoriesApt
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackageRepositoriesApt(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackageRepositoriesAptSlice(c *Client, des, nw []GuestPolicyPackageRepositoriesApt) []GuestPolicyPackageRepositoriesApt {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackageRepositoriesApt
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackageRepositoriesApt(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackageRepositoriesYum(des, initial *GuestPolicyPackageRepositoriesYum, opts ...dcl.ApplyOption) *GuestPolicyPackageRepositoriesYum {
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

func canonicalizeNewGuestPolicyPackageRepositoriesYum(c *Client, des, nw *GuestPolicyPackageRepositoriesYum) *GuestPolicyPackageRepositoriesYum {
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

	return nw
}

func canonicalizeNewGuestPolicyPackageRepositoriesYumSet(c *Client, des, nw []GuestPolicyPackageRepositoriesYum) []GuestPolicyPackageRepositoriesYum {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositoriesYum
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackageRepositoriesYum(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackageRepositoriesYumSlice(c *Client, des, nw []GuestPolicyPackageRepositoriesYum) []GuestPolicyPackageRepositoriesYum {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackageRepositoriesYum
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackageRepositoriesYum(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackageRepositoriesZypper(des, initial *GuestPolicyPackageRepositoriesZypper, opts ...dcl.ApplyOption) *GuestPolicyPackageRepositoriesZypper {
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

func canonicalizeNewGuestPolicyPackageRepositoriesZypper(c *Client, des, nw *GuestPolicyPackageRepositoriesZypper) *GuestPolicyPackageRepositoriesZypper {
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

	return nw
}

func canonicalizeNewGuestPolicyPackageRepositoriesZypperSet(c *Client, des, nw []GuestPolicyPackageRepositoriesZypper) []GuestPolicyPackageRepositoriesZypper {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositoriesZypper
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackageRepositoriesZypper(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackageRepositoriesZypperSlice(c *Client, des, nw []GuestPolicyPackageRepositoriesZypper) []GuestPolicyPackageRepositoriesZypper {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackageRepositoriesZypper
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackageRepositoriesZypper(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyPackageRepositoriesGoo(des, initial *GuestPolicyPackageRepositoriesGoo, opts ...dcl.ApplyOption) *GuestPolicyPackageRepositoriesGoo {
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

func canonicalizeNewGuestPolicyPackageRepositoriesGoo(c *Client, des, nw *GuestPolicyPackageRepositoriesGoo) *GuestPolicyPackageRepositoriesGoo {
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

func canonicalizeNewGuestPolicyPackageRepositoriesGooSet(c *Client, des, nw []GuestPolicyPackageRepositoriesGoo) []GuestPolicyPackageRepositoriesGoo {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositoriesGoo
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyPackageRepositoriesGoo(c, &d, &n) {
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

func canonicalizeNewGuestPolicyPackageRepositoriesGooSlice(c *Client, des, nw []GuestPolicyPackageRepositoriesGoo) []GuestPolicyPackageRepositoriesGoo {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyPackageRepositoriesGoo
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyPackageRepositoriesGoo(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipes(des, initial *GuestPolicyRecipes, opts ...dcl.ApplyOption) *GuestPolicyRecipes {
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
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}
	if dcl.IsZeroValue(des.Artifacts) {
		des.Artifacts = initial.Artifacts
	}
	if dcl.IsZeroValue(des.InstallSteps) {
		des.InstallSteps = initial.InstallSteps
	}
	if dcl.IsZeroValue(des.UpdateSteps) {
		des.UpdateSteps = initial.UpdateSteps
	}
	if dcl.IsZeroValue(des.DesiredState) {
		des.DesiredState = initial.DesiredState
	}

	return des
}

func canonicalizeNewGuestPolicyRecipes(c *Client, des, nw *GuestPolicyRecipes) *GuestPolicyRecipes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	nw.Artifacts = canonicalizeNewGuestPolicyRecipesArtifactsSlice(c, des.Artifacts, nw.Artifacts)
	nw.InstallSteps = canonicalizeNewGuestPolicyRecipesInstallStepsSlice(c, des.InstallSteps, nw.InstallSteps)
	nw.UpdateSteps = canonicalizeNewGuestPolicyRecipesUpdateStepsSlice(c, des.UpdateSteps, nw.UpdateSteps)

	return nw
}

func canonicalizeNewGuestPolicyRecipesSet(c *Client, des, nw []GuestPolicyRecipes) []GuestPolicyRecipes {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipes(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesSlice(c *Client, des, nw []GuestPolicyRecipes) []GuestPolicyRecipes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipes(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesArtifacts(des, initial *GuestPolicyRecipesArtifacts, opts ...dcl.ApplyOption) *GuestPolicyRecipesArtifacts {
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
	des.Remote = canonicalizeGuestPolicyRecipesArtifactsRemote(des.Remote, initial.Remote, opts...)
	des.Gcs = canonicalizeGuestPolicyRecipesArtifactsGcs(des.Gcs, initial.Gcs, opts...)
	if dcl.BoolCanonicalize(des.AllowInsecure, initial.AllowInsecure) || dcl.IsZeroValue(des.AllowInsecure) {
		des.AllowInsecure = initial.AllowInsecure
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesArtifacts(c *Client, des, nw *GuestPolicyRecipesArtifacts) *GuestPolicyRecipesArtifacts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	nw.Remote = canonicalizeNewGuestPolicyRecipesArtifactsRemote(c, des.Remote, nw.Remote)
	nw.Gcs = canonicalizeNewGuestPolicyRecipesArtifactsGcs(c, des.Gcs, nw.Gcs)
	if dcl.BoolCanonicalize(des.AllowInsecure, nw.AllowInsecure) {
		nw.AllowInsecure = des.AllowInsecure
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesArtifactsSet(c *Client, des, nw []GuestPolicyRecipesArtifacts) []GuestPolicyRecipesArtifacts {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesArtifacts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesArtifacts(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesArtifactsSlice(c *Client, des, nw []GuestPolicyRecipesArtifacts) []GuestPolicyRecipesArtifacts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesArtifacts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesArtifacts(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesArtifactsRemote(des, initial *GuestPolicyRecipesArtifactsRemote, opts ...dcl.ApplyOption) *GuestPolicyRecipesArtifactsRemote {
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
	if dcl.StringCanonicalize(des.Checksum, initial.Checksum) || dcl.IsZeroValue(des.Checksum) {
		des.Checksum = initial.Checksum
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesArtifactsRemote(c *Client, des, nw *GuestPolicyRecipesArtifactsRemote) *GuestPolicyRecipesArtifactsRemote {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Checksum, nw.Checksum) {
		nw.Checksum = des.Checksum
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesArtifactsRemoteSet(c *Client, des, nw []GuestPolicyRecipesArtifactsRemote) []GuestPolicyRecipesArtifactsRemote {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesArtifactsRemote
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesArtifactsRemote(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesArtifactsRemoteSlice(c *Client, des, nw []GuestPolicyRecipesArtifactsRemote) []GuestPolicyRecipesArtifactsRemote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesArtifactsRemote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesArtifactsRemote(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesArtifactsGcs(des, initial *GuestPolicyRecipesArtifactsGcs, opts ...dcl.ApplyOption) *GuestPolicyRecipesArtifactsGcs {
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
	if dcl.StringCanonicalize(des.Object, initial.Object) || dcl.IsZeroValue(des.Object) {
		des.Object = initial.Object
	}
	if dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesArtifactsGcs(c *Client, des, nw *GuestPolicyRecipesArtifactsGcs) *GuestPolicyRecipesArtifactsGcs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Bucket, nw.Bucket) {
		nw.Bucket = des.Bucket
	}
	if dcl.StringCanonicalize(des.Object, nw.Object) {
		nw.Object = des.Object
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesArtifactsGcsSet(c *Client, des, nw []GuestPolicyRecipesArtifactsGcs) []GuestPolicyRecipesArtifactsGcs {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesArtifactsGcs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesArtifactsGcs(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesArtifactsGcsSlice(c *Client, des, nw []GuestPolicyRecipesArtifactsGcs) []GuestPolicyRecipesArtifactsGcs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesArtifactsGcs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesArtifactsGcs(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallSteps(des, initial *GuestPolicyRecipesInstallSteps, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallSteps {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.FileCopy = canonicalizeGuestPolicyRecipesInstallStepsFileCopy(des.FileCopy, initial.FileCopy, opts...)
	des.ArchiveExtraction = canonicalizeGuestPolicyRecipesInstallStepsArchiveExtraction(des.ArchiveExtraction, initial.ArchiveExtraction, opts...)
	des.MsiInstallation = canonicalizeGuestPolicyRecipesInstallStepsMsiInstallation(des.MsiInstallation, initial.MsiInstallation, opts...)
	des.DpkgInstallation = canonicalizeGuestPolicyRecipesInstallStepsDpkgInstallation(des.DpkgInstallation, initial.DpkgInstallation, opts...)
	des.RpmInstallation = canonicalizeGuestPolicyRecipesInstallStepsRpmInstallation(des.RpmInstallation, initial.RpmInstallation, opts...)
	des.FileExec = canonicalizeGuestPolicyRecipesInstallStepsFileExec(des.FileExec, initial.FileExec, opts...)
	des.ScriptRun = canonicalizeGuestPolicyRecipesInstallStepsScriptRun(des.ScriptRun, initial.ScriptRun, opts...)

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallSteps(c *Client, des, nw *GuestPolicyRecipesInstallSteps) *GuestPolicyRecipesInstallSteps {
	if des == nil || nw == nil {
		return nw
	}

	nw.FileCopy = canonicalizeNewGuestPolicyRecipesInstallStepsFileCopy(c, des.FileCopy, nw.FileCopy)
	nw.ArchiveExtraction = canonicalizeNewGuestPolicyRecipesInstallStepsArchiveExtraction(c, des.ArchiveExtraction, nw.ArchiveExtraction)
	nw.MsiInstallation = canonicalizeNewGuestPolicyRecipesInstallStepsMsiInstallation(c, des.MsiInstallation, nw.MsiInstallation)
	nw.DpkgInstallation = canonicalizeNewGuestPolicyRecipesInstallStepsDpkgInstallation(c, des.DpkgInstallation, nw.DpkgInstallation)
	nw.RpmInstallation = canonicalizeNewGuestPolicyRecipesInstallStepsRpmInstallation(c, des.RpmInstallation, nw.RpmInstallation)
	nw.FileExec = canonicalizeNewGuestPolicyRecipesInstallStepsFileExec(c, des.FileExec, nw.FileExec)
	nw.ScriptRun = canonicalizeNewGuestPolicyRecipesInstallStepsScriptRun(c, des.ScriptRun, nw.ScriptRun)

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsSet(c *Client, des, nw []GuestPolicyRecipesInstallSteps) []GuestPolicyRecipesInstallSteps {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallSteps
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallSteps(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsSlice(c *Client, des, nw []GuestPolicyRecipesInstallSteps) []GuestPolicyRecipesInstallSteps {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallSteps
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallSteps(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsFileCopy(des, initial *GuestPolicyRecipesInstallStepsFileCopy, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsFileCopy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, initial.Destination) || dcl.IsZeroValue(des.Destination) {
		des.Destination = initial.Destination
	}
	if dcl.BoolCanonicalize(des.Overwrite, initial.Overwrite) || dcl.IsZeroValue(des.Overwrite) {
		des.Overwrite = initial.Overwrite
	}
	if dcl.StringCanonicalize(des.Permissions, initial.Permissions) || dcl.IsZeroValue(des.Permissions) {
		des.Permissions = initial.Permissions
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsFileCopy(c *Client, des, nw *GuestPolicyRecipesInstallStepsFileCopy) *GuestPolicyRecipesInstallStepsFileCopy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, nw.Destination) {
		nw.Destination = des.Destination
	}
	if dcl.BoolCanonicalize(des.Overwrite, nw.Overwrite) {
		nw.Overwrite = des.Overwrite
	}
	if dcl.StringCanonicalize(des.Permissions, nw.Permissions) {
		nw.Permissions = des.Permissions
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsFileCopySet(c *Client, des, nw []GuestPolicyRecipesInstallStepsFileCopy) []GuestPolicyRecipesInstallStepsFileCopy {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsFileCopy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsFileCopy(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsFileCopySlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsFileCopy) []GuestPolicyRecipesInstallStepsFileCopy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsFileCopy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsFileCopy(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsArchiveExtraction(des, initial *GuestPolicyRecipesInstallStepsArchiveExtraction, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsArchiveExtraction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, initial.Destination) || dcl.IsZeroValue(des.Destination) {
		des.Destination = initial.Destination
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsArchiveExtraction(c *Client, des, nw *GuestPolicyRecipesInstallStepsArchiveExtraction) *GuestPolicyRecipesInstallStepsArchiveExtraction {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, nw.Destination) {
		nw.Destination = des.Destination
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsArchiveExtractionSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsArchiveExtraction) []GuestPolicyRecipesInstallStepsArchiveExtraction {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsArchiveExtraction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsArchiveExtraction(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsArchiveExtractionSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsArchiveExtraction) []GuestPolicyRecipesInstallStepsArchiveExtraction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsArchiveExtraction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsArchiveExtraction(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsMsiInstallation(des, initial *GuestPolicyRecipesInstallStepsMsiInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsMsiInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.IsZeroValue(des.Flags) {
		des.Flags = initial.Flags
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsMsiInstallation(c *Client, des, nw *GuestPolicyRecipesInstallStepsMsiInstallation) *GuestPolicyRecipesInstallStepsMsiInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsMsiInstallationSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsMsiInstallation) []GuestPolicyRecipesInstallStepsMsiInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsMsiInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsMsiInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsMsiInstallationSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsMsiInstallation) []GuestPolicyRecipesInstallStepsMsiInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsMsiInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsMsiInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsDpkgInstallation(des, initial *GuestPolicyRecipesInstallStepsDpkgInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsDpkgInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsDpkgInstallation(c *Client, des, nw *GuestPolicyRecipesInstallStepsDpkgInstallation) *GuestPolicyRecipesInstallStepsDpkgInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsDpkgInstallationSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsDpkgInstallation) []GuestPolicyRecipesInstallStepsDpkgInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsDpkgInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsDpkgInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsDpkgInstallationSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsDpkgInstallation) []GuestPolicyRecipesInstallStepsDpkgInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsDpkgInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsDpkgInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsRpmInstallation(des, initial *GuestPolicyRecipesInstallStepsRpmInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsRpmInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsRpmInstallation(c *Client, des, nw *GuestPolicyRecipesInstallStepsRpmInstallation) *GuestPolicyRecipesInstallStepsRpmInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsRpmInstallationSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsRpmInstallation) []GuestPolicyRecipesInstallStepsRpmInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsRpmInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsRpmInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsRpmInstallationSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsRpmInstallation) []GuestPolicyRecipesInstallStepsRpmInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsRpmInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsRpmInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsFileExec(des, initial *GuestPolicyRecipesInstallStepsFileExec, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsFileExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.LocalPath, initial.LocalPath) || dcl.IsZeroValue(des.LocalPath) {
		des.LocalPath = initial.LocalPath
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsFileExec(c *Client, des, nw *GuestPolicyRecipesInstallStepsFileExec) *GuestPolicyRecipesInstallStepsFileExec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.LocalPath, nw.LocalPath) {
		nw.LocalPath = des.LocalPath
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsFileExecSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsFileExec) []GuestPolicyRecipesInstallStepsFileExec {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsFileExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsFileExec(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsFileExecSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsFileExec) []GuestPolicyRecipesInstallStepsFileExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsFileExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsFileExec(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesInstallStepsScriptRun(des, initial *GuestPolicyRecipesInstallStepsScriptRun, opts ...dcl.ApplyOption) *GuestPolicyRecipesInstallStepsScriptRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Script, initial.Script) || dcl.IsZeroValue(des.Script) {
		des.Script = initial.Script
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}
	if dcl.IsZeroValue(des.Interpreter) {
		des.Interpreter = initial.Interpreter
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesInstallStepsScriptRun(c *Client, des, nw *GuestPolicyRecipesInstallStepsScriptRun) *GuestPolicyRecipesInstallStepsScriptRun {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Script, nw.Script) {
		nw.Script = des.Script
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesInstallStepsScriptRunSet(c *Client, des, nw []GuestPolicyRecipesInstallStepsScriptRun) []GuestPolicyRecipesInstallStepsScriptRun {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesInstallStepsScriptRun
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesInstallStepsScriptRun(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesInstallStepsScriptRunSlice(c *Client, des, nw []GuestPolicyRecipesInstallStepsScriptRun) []GuestPolicyRecipesInstallStepsScriptRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesInstallStepsScriptRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesInstallStepsScriptRun(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateSteps(des, initial *GuestPolicyRecipesUpdateSteps, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateSteps {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.FileCopy = canonicalizeGuestPolicyRecipesUpdateStepsFileCopy(des.FileCopy, initial.FileCopy, opts...)
	des.ArchiveExtraction = canonicalizeGuestPolicyRecipesUpdateStepsArchiveExtraction(des.ArchiveExtraction, initial.ArchiveExtraction, opts...)
	des.MsiInstallation = canonicalizeGuestPolicyRecipesUpdateStepsMsiInstallation(des.MsiInstallation, initial.MsiInstallation, opts...)
	des.DpkgInstallation = canonicalizeGuestPolicyRecipesUpdateStepsDpkgInstallation(des.DpkgInstallation, initial.DpkgInstallation, opts...)
	des.RpmInstallation = canonicalizeGuestPolicyRecipesUpdateStepsRpmInstallation(des.RpmInstallation, initial.RpmInstallation, opts...)
	des.FileExec = canonicalizeGuestPolicyRecipesUpdateStepsFileExec(des.FileExec, initial.FileExec, opts...)
	des.ScriptRun = canonicalizeGuestPolicyRecipesUpdateStepsScriptRun(des.ScriptRun, initial.ScriptRun, opts...)

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateSteps(c *Client, des, nw *GuestPolicyRecipesUpdateSteps) *GuestPolicyRecipesUpdateSteps {
	if des == nil || nw == nil {
		return nw
	}

	nw.FileCopy = canonicalizeNewGuestPolicyRecipesUpdateStepsFileCopy(c, des.FileCopy, nw.FileCopy)
	nw.ArchiveExtraction = canonicalizeNewGuestPolicyRecipesUpdateStepsArchiveExtraction(c, des.ArchiveExtraction, nw.ArchiveExtraction)
	nw.MsiInstallation = canonicalizeNewGuestPolicyRecipesUpdateStepsMsiInstallation(c, des.MsiInstallation, nw.MsiInstallation)
	nw.DpkgInstallation = canonicalizeNewGuestPolicyRecipesUpdateStepsDpkgInstallation(c, des.DpkgInstallation, nw.DpkgInstallation)
	nw.RpmInstallation = canonicalizeNewGuestPolicyRecipesUpdateStepsRpmInstallation(c, des.RpmInstallation, nw.RpmInstallation)
	nw.FileExec = canonicalizeNewGuestPolicyRecipesUpdateStepsFileExec(c, des.FileExec, nw.FileExec)
	nw.ScriptRun = canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRun(c, des.ScriptRun, nw.ScriptRun)

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsSet(c *Client, des, nw []GuestPolicyRecipesUpdateSteps) []GuestPolicyRecipesUpdateSteps {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateSteps
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateSteps(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsSlice(c *Client, des, nw []GuestPolicyRecipesUpdateSteps) []GuestPolicyRecipesUpdateSteps {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateSteps
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateSteps(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsFileCopy(des, initial *GuestPolicyRecipesUpdateStepsFileCopy, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsFileCopy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, initial.Destination) || dcl.IsZeroValue(des.Destination) {
		des.Destination = initial.Destination
	}
	if dcl.BoolCanonicalize(des.Overwrite, initial.Overwrite) || dcl.IsZeroValue(des.Overwrite) {
		des.Overwrite = initial.Overwrite
	}
	if dcl.StringCanonicalize(des.Permissions, initial.Permissions) || dcl.IsZeroValue(des.Permissions) {
		des.Permissions = initial.Permissions
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileCopy(c *Client, des, nw *GuestPolicyRecipesUpdateStepsFileCopy) *GuestPolicyRecipesUpdateStepsFileCopy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, nw.Destination) {
		nw.Destination = des.Destination
	}
	if dcl.BoolCanonicalize(des.Overwrite, nw.Overwrite) {
		nw.Overwrite = des.Overwrite
	}
	if dcl.StringCanonicalize(des.Permissions, nw.Permissions) {
		nw.Permissions = des.Permissions
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileCopySet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsFileCopy) []GuestPolicyRecipesUpdateStepsFileCopy {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsFileCopy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsFileCopy(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileCopySlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsFileCopy) []GuestPolicyRecipesUpdateStepsFileCopy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsFileCopy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsFileCopy(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsArchiveExtraction(des, initial *GuestPolicyRecipesUpdateStepsArchiveExtraction, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, initial.Destination) || dcl.IsZeroValue(des.Destination) {
		des.Destination = initial.Destination
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsArchiveExtraction(c *Client, des, nw *GuestPolicyRecipesUpdateStepsArchiveExtraction) *GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.Destination, nw.Destination) {
		nw.Destination = des.Destination
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsArchiveExtractionSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsArchiveExtraction) []GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsArchiveExtraction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsArchiveExtractionSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsArchiveExtraction) []GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsArchiveExtraction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsMsiInstallation(des, initial *GuestPolicyRecipesUpdateStepsMsiInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsMsiInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.IsZeroValue(des.Flags) {
		des.Flags = initial.Flags
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsMsiInstallation(c *Client, des, nw *GuestPolicyRecipesUpdateStepsMsiInstallation) *GuestPolicyRecipesUpdateStepsMsiInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsMsiInstallationSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsMsiInstallation) []GuestPolicyRecipesUpdateStepsMsiInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsMsiInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsMsiInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsMsiInstallationSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsMsiInstallation) []GuestPolicyRecipesUpdateStepsMsiInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsMsiInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsMsiInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsDpkgInstallation(des, initial *GuestPolicyRecipesUpdateStepsDpkgInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsDpkgInstallation(c *Client, des, nw *GuestPolicyRecipesUpdateStepsDpkgInstallation) *GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsDpkgInstallationSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsDpkgInstallation) []GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsDpkgInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsDpkgInstallationSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsDpkgInstallation) []GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsDpkgInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsRpmInstallation(des, initial *GuestPolicyRecipesUpdateStepsRpmInstallation, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsRpmInstallation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsRpmInstallation(c *Client, des, nw *GuestPolicyRecipesUpdateStepsRpmInstallation) *GuestPolicyRecipesUpdateStepsRpmInstallation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsRpmInstallationSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsRpmInstallation) []GuestPolicyRecipesUpdateStepsRpmInstallation {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsRpmInstallation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsRpmInstallation(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsRpmInstallationSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsRpmInstallation) []GuestPolicyRecipesUpdateStepsRpmInstallation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsRpmInstallation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsRpmInstallation(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsFileExec(des, initial *GuestPolicyRecipesUpdateStepsFileExec, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsFileExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ArtifactId, initial.ArtifactId) || dcl.IsZeroValue(des.ArtifactId) {
		des.ArtifactId = initial.ArtifactId
	}
	if dcl.StringCanonicalize(des.LocalPath, initial.LocalPath) || dcl.IsZeroValue(des.LocalPath) {
		des.LocalPath = initial.LocalPath
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileExec(c *Client, des, nw *GuestPolicyRecipesUpdateStepsFileExec) *GuestPolicyRecipesUpdateStepsFileExec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ArtifactId, nw.ArtifactId) {
		nw.ArtifactId = des.ArtifactId
	}
	if dcl.StringCanonicalize(des.LocalPath, nw.LocalPath) {
		nw.LocalPath = des.LocalPath
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileExecSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsFileExec) []GuestPolicyRecipesUpdateStepsFileExec {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsFileExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsFileExec(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsFileExecSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsFileExec) []GuestPolicyRecipesUpdateStepsFileExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsFileExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsFileExec(c, &d, &n))
	}

	return items
}

func canonicalizeGuestPolicyRecipesUpdateStepsScriptRun(des, initial *GuestPolicyRecipesUpdateStepsScriptRun, opts ...dcl.ApplyOption) *GuestPolicyRecipesUpdateStepsScriptRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Script, initial.Script) || dcl.IsZeroValue(des.Script) {
		des.Script = initial.Script
	}
	if dcl.IsZeroValue(des.AllowedExitCodes) {
		des.AllowedExitCodes = initial.AllowedExitCodes
	}
	if dcl.IsZeroValue(des.Interpreter) {
		des.Interpreter = initial.Interpreter
	}

	return des
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRun(c *Client, des, nw *GuestPolicyRecipesUpdateStepsScriptRun) *GuestPolicyRecipesUpdateStepsScriptRun {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Script, nw.Script) {
		nw.Script = des.Script
	}

	return nw
}

func canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRunSet(c *Client, des, nw []GuestPolicyRecipesUpdateStepsScriptRun) []GuestPolicyRecipesUpdateStepsScriptRun {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyRecipesUpdateStepsScriptRun
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareGuestPolicyRecipesUpdateStepsScriptRun(c, &d, &n) {
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

func canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRunSlice(c *Client, des, nw []GuestPolicyRecipesUpdateStepsScriptRun) []GuestPolicyRecipesUpdateStepsScriptRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []GuestPolicyRecipesUpdateStepsScriptRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRun(c, &d, &n))
	}

	return items
}

type guestPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         guestPolicyApiOperation
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
func diffGuestPolicy(c *Client, desired, actual *GuestPolicy, opts ...dcl.ApplyOption) ([]guestPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []guestPolicyDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, guestPolicyDiff{
			UpdateOp: &updateGuestPolicyUpdateGuestPolicyOperation{}, FieldName: "Name",
		})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, guestPolicyDiff{
			UpdateOp: &updateGuestPolicyUpdateGuestPolicyOperation{}, FieldName: "Description",
		})
	}

	if d, err := dcl.Diff(desired.CreateTime, actual.CreateTime, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, guestPolicyDiff{RequiresRecreate: true, FieldName: "CreateTime"})
	}

	if d, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, guestPolicyDiff{RequiresRecreate: true, FieldName: "UpdateTime"})
	}

	if d, err := dcl.Diff(desired.Etag, actual.Etag, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, guestPolicyDiff{
			UpdateOp: &updateGuestPolicyUpdateGuestPolicyOperation{}, FieldName: "Etag",
		})
	}

	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Description",
		})

	}
	if compareGuestPolicyAssignment(c, desired.Assignment, actual.Assignment) {
		c.Config.Logger.Infof("Detected diff in Assignment.\nDESIRED: %v\nACTUAL: %v", desired.Assignment, actual.Assignment)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Assignment",
		})

	}
	if compareGuestPolicyPackagesSlice(c, desired.Packages, actual.Packages) {
		c.Config.Logger.Infof("Detected diff in Packages.\nDESIRED: %v\nACTUAL: %v", desired.Packages, actual.Packages)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Packages",
		})

	}
	if compareGuestPolicyPackageRepositoriesSlice(c, desired.PackageRepositories, actual.PackageRepositories) {
		c.Config.Logger.Infof("Detected diff in PackageRepositories.\nDESIRED: %v\nACTUAL: %v", desired.PackageRepositories, actual.PackageRepositories)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "PackageRepositories",
		})

	}
	if compareGuestPolicyRecipesSlice(c, desired.Recipes, actual.Recipes) {
		c.Config.Logger.Infof("Detected diff in Recipes.\nDESIRED: %v\nACTUAL: %v", desired.Recipes, actual.Recipes)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Recipes",
		})

	}
	if !dcl.IsZeroValue(desired.Etag) && !dcl.StringCanonicalize(desired.Etag, actual.Etag) {
		c.Config.Logger.Infof("Detected diff in Etag.\nDESIRED: %v\nACTUAL: %v", desired.Etag, actual.Etag)

		diffs = append(diffs, guestPolicyDiff{
			UpdateOp:  &updateGuestPolicyUpdateGuestPolicyOperation{},
			FieldName: "Etag",
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
	var deduped []guestPolicyDiff
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
func compareGuestPolicyAssignment(c *Client, desired, actual *GuestPolicyAssignment) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.GroupLabels == nil && desired.GroupLabels != nil && !dcl.IsEmptyValueIndirect(desired.GroupLabels) {
		c.Config.Logger.Infof("desired GroupLabels %s - but actually nil", dcl.SprintResource(desired.GroupLabels))
		return true
	}
	if compareGuestPolicyAssignmentGroupLabelsSlice(c, desired.GroupLabels, actual.GroupLabels) && !dcl.IsZeroValue(desired.GroupLabels) {
		c.Config.Logger.Infof("Diff in GroupLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GroupLabels), dcl.SprintResource(actual.GroupLabels))
		return true
	}
	if actual.Zones == nil && desired.Zones != nil && !dcl.IsEmptyValueIndirect(desired.Zones) {
		c.Config.Logger.Infof("desired Zones %s - but actually nil", dcl.SprintResource(desired.Zones))
		return true
	}
	if !dcl.StringSliceEquals(desired.Zones, actual.Zones) && !dcl.IsZeroValue(desired.Zones) {
		c.Config.Logger.Infof("Diff in Zones. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zones), dcl.SprintResource(actual.Zones))
		return true
	}
	if actual.Instances == nil && desired.Instances != nil && !dcl.IsEmptyValueIndirect(desired.Instances) {
		c.Config.Logger.Infof("desired Instances %s - but actually nil", dcl.SprintResource(desired.Instances))
		return true
	}
	if !dcl.StringSliceEqualsWithSelfLink(desired.Instances, actual.Instances) && !dcl.IsZeroValue(desired.Instances) {
		c.Config.Logger.Infof("Diff in Instances. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Instances), dcl.SprintResource(actual.Instances))
		return true
	}
	if actual.InstanceNamePrefixes == nil && desired.InstanceNamePrefixes != nil && !dcl.IsEmptyValueIndirect(desired.InstanceNamePrefixes) {
		c.Config.Logger.Infof("desired InstanceNamePrefixes %s - but actually nil", dcl.SprintResource(desired.InstanceNamePrefixes))
		return true
	}
	if !dcl.StringSliceEquals(desired.InstanceNamePrefixes, actual.InstanceNamePrefixes) && !dcl.IsZeroValue(desired.InstanceNamePrefixes) {
		c.Config.Logger.Infof("Diff in InstanceNamePrefixes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstanceNamePrefixes), dcl.SprintResource(actual.InstanceNamePrefixes))
		return true
	}
	if actual.OsTypes == nil && desired.OsTypes != nil && !dcl.IsEmptyValueIndirect(desired.OsTypes) {
		c.Config.Logger.Infof("desired OsTypes %s - but actually nil", dcl.SprintResource(desired.OsTypes))
		return true
	}
	if compareGuestPolicyAssignmentOsTypesSlice(c, desired.OsTypes, actual.OsTypes) && !dcl.IsZeroValue(desired.OsTypes) {
		c.Config.Logger.Infof("Diff in OsTypes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsTypes), dcl.SprintResource(actual.OsTypes))
		return true
	}
	return false
}

func compareGuestPolicyAssignmentSlice(c *Client, desired, actual []GuestPolicyAssignment) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignment, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyAssignment(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignment, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyAssignmentMap(c *Client, desired, actual map[string]GuestPolicyAssignment) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignment, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignment, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyAssignment(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignment, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyAssignmentGroupLabels(c *Client, desired, actual *GuestPolicyAssignmentGroupLabels) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	return false
}

func compareGuestPolicyAssignmentGroupLabelsSlice(c *Client, desired, actual []GuestPolicyAssignmentGroupLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignmentGroupLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyAssignmentGroupLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentGroupLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyAssignmentGroupLabelsMap(c *Client, desired, actual map[string]GuestPolicyAssignmentGroupLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignmentGroupLabels, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentGroupLabels, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyAssignmentGroupLabels(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentGroupLabels, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyAssignmentOsTypes(c *Client, desired, actual *GuestPolicyAssignmentOsTypes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.OsShortName == nil && desired.OsShortName != nil && !dcl.IsEmptyValueIndirect(desired.OsShortName) {
		c.Config.Logger.Infof("desired OsShortName %s - but actually nil", dcl.SprintResource(desired.OsShortName))
		return true
	}
	if !dcl.StringCanonicalize(desired.OsShortName, actual.OsShortName) && !dcl.IsZeroValue(desired.OsShortName) {
		c.Config.Logger.Infof("Diff in OsShortName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsShortName), dcl.SprintResource(actual.OsShortName))
		return true
	}
	if actual.OsVersion == nil && desired.OsVersion != nil && !dcl.IsEmptyValueIndirect(desired.OsVersion) {
		c.Config.Logger.Infof("desired OsVersion %s - but actually nil", dcl.SprintResource(desired.OsVersion))
		return true
	}
	if !dcl.StringCanonicalize(desired.OsVersion, actual.OsVersion) && !dcl.IsZeroValue(desired.OsVersion) {
		c.Config.Logger.Infof("Diff in OsVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsVersion), dcl.SprintResource(actual.OsVersion))
		return true
	}
	if actual.OsArchitecture == nil && desired.OsArchitecture != nil && !dcl.IsEmptyValueIndirect(desired.OsArchitecture) {
		c.Config.Logger.Infof("desired OsArchitecture %s - but actually nil", dcl.SprintResource(desired.OsArchitecture))
		return true
	}
	if !dcl.StringCanonicalize(desired.OsArchitecture, actual.OsArchitecture) && !dcl.IsZeroValue(desired.OsArchitecture) {
		c.Config.Logger.Infof("Diff in OsArchitecture. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsArchitecture), dcl.SprintResource(actual.OsArchitecture))
		return true
	}
	return false
}

func compareGuestPolicyAssignmentOsTypesSlice(c *Client, desired, actual []GuestPolicyAssignmentOsTypes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignmentOsTypes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyAssignmentOsTypes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentOsTypes, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyAssignmentOsTypesMap(c *Client, desired, actual map[string]GuestPolicyAssignmentOsTypes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyAssignmentOsTypes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentOsTypes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyAssignmentOsTypes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyAssignmentOsTypes, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackages(c *Client, desired, actual *GuestPolicyPackages) bool {
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
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.DesiredState == nil && desired.DesiredState != nil && !dcl.IsEmptyValueIndirect(desired.DesiredState) {
		c.Config.Logger.Infof("desired DesiredState %s - but actually nil", dcl.SprintResource(desired.DesiredState))
		return true
	}
	if !reflect.DeepEqual(desired.DesiredState, actual.DesiredState) && !dcl.IsZeroValue(desired.DesiredState) {
		c.Config.Logger.Infof("Diff in DesiredState. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DesiredState), dcl.SprintResource(actual.DesiredState))
		return true
	}
	if actual.Manager == nil && desired.Manager != nil && !dcl.IsEmptyValueIndirect(desired.Manager) {
		c.Config.Logger.Infof("desired Manager %s - but actually nil", dcl.SprintResource(desired.Manager))
		return true
	}
	if !reflect.DeepEqual(desired.Manager, actual.Manager) && !dcl.IsZeroValue(desired.Manager) {
		c.Config.Logger.Infof("Diff in Manager. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Manager), dcl.SprintResource(actual.Manager))
		return true
	}
	return false
}

func compareGuestPolicyPackagesSlice(c *Client, desired, actual []GuestPolicyPackages) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackages, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackages(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackages, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackagesMap(c *Client, desired, actual map[string]GuestPolicyPackages) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackages, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackages, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackages(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackages, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositories(c *Client, desired, actual *GuestPolicyPackageRepositories) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Apt == nil && desired.Apt != nil && !dcl.IsEmptyValueIndirect(desired.Apt) {
		c.Config.Logger.Infof("desired Apt %s - but actually nil", dcl.SprintResource(desired.Apt))
		return true
	}
	if compareGuestPolicyPackageRepositoriesApt(c, desired.Apt, actual.Apt) && !dcl.IsZeroValue(desired.Apt) {
		c.Config.Logger.Infof("Diff in Apt. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Apt), dcl.SprintResource(actual.Apt))
		return true
	}
	if actual.Yum == nil && desired.Yum != nil && !dcl.IsEmptyValueIndirect(desired.Yum) {
		c.Config.Logger.Infof("desired Yum %s - but actually nil", dcl.SprintResource(desired.Yum))
		return true
	}
	if compareGuestPolicyPackageRepositoriesYum(c, desired.Yum, actual.Yum) && !dcl.IsZeroValue(desired.Yum) {
		c.Config.Logger.Infof("Diff in Yum. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Yum), dcl.SprintResource(actual.Yum))
		return true
	}
	if actual.Zypper == nil && desired.Zypper != nil && !dcl.IsEmptyValueIndirect(desired.Zypper) {
		c.Config.Logger.Infof("desired Zypper %s - but actually nil", dcl.SprintResource(desired.Zypper))
		return true
	}
	if compareGuestPolicyPackageRepositoriesZypper(c, desired.Zypper, actual.Zypper) && !dcl.IsZeroValue(desired.Zypper) {
		c.Config.Logger.Infof("Diff in Zypper. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zypper), dcl.SprintResource(actual.Zypper))
		return true
	}
	if actual.Goo == nil && desired.Goo != nil && !dcl.IsEmptyValueIndirect(desired.Goo) {
		c.Config.Logger.Infof("desired Goo %s - but actually nil", dcl.SprintResource(desired.Goo))
		return true
	}
	if compareGuestPolicyPackageRepositoriesGoo(c, desired.Goo, actual.Goo) && !dcl.IsZeroValue(desired.Goo) {
		c.Config.Logger.Infof("Diff in Goo. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Goo), dcl.SprintResource(actual.Goo))
		return true
	}
	return false
}

func compareGuestPolicyPackageRepositoriesSlice(c *Client, desired, actual []GuestPolicyPackageRepositories) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositories, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositories(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositories, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesMap(c *Client, desired, actual map[string]GuestPolicyPackageRepositories) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositories, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositories, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackageRepositories(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositories, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesApt(c *Client, desired, actual *GuestPolicyPackageRepositoriesApt) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArchiveType == nil && desired.ArchiveType != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveType) {
		c.Config.Logger.Infof("desired ArchiveType %s - but actually nil", dcl.SprintResource(desired.ArchiveType))
		return true
	}
	if !reflect.DeepEqual(desired.ArchiveType, actual.ArchiveType) && !dcl.IsZeroValue(desired.ArchiveType) {
		c.Config.Logger.Infof("Diff in ArchiveType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveType), dcl.SprintResource(actual.ArchiveType))
		return true
	}
	if actual.Uri == nil && desired.Uri != nil && !dcl.IsEmptyValueIndirect(desired.Uri) {
		c.Config.Logger.Infof("desired Uri %s - but actually nil", dcl.SprintResource(desired.Uri))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uri, actual.Uri) && !dcl.IsZeroValue(desired.Uri) {
		c.Config.Logger.Infof("Diff in Uri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uri), dcl.SprintResource(actual.Uri))
		return true
	}
	if actual.Distribution == nil && desired.Distribution != nil && !dcl.IsEmptyValueIndirect(desired.Distribution) {
		c.Config.Logger.Infof("desired Distribution %s - but actually nil", dcl.SprintResource(desired.Distribution))
		return true
	}
	if !dcl.StringCanonicalize(desired.Distribution, actual.Distribution) && !dcl.IsZeroValue(desired.Distribution) {
		c.Config.Logger.Infof("Diff in Distribution. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Distribution), dcl.SprintResource(actual.Distribution))
		return true
	}
	if actual.Components == nil && desired.Components != nil && !dcl.IsEmptyValueIndirect(desired.Components) {
		c.Config.Logger.Infof("desired Components %s - but actually nil", dcl.SprintResource(desired.Components))
		return true
	}
	if !dcl.StringSliceEquals(desired.Components, actual.Components) && !dcl.IsZeroValue(desired.Components) {
		c.Config.Logger.Infof("Diff in Components. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Components), dcl.SprintResource(actual.Components))
		return true
	}
	if actual.GpgKey == nil && desired.GpgKey != nil && !dcl.IsEmptyValueIndirect(desired.GpgKey) {
		c.Config.Logger.Infof("desired GpgKey %s - but actually nil", dcl.SprintResource(desired.GpgKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.GpgKey, actual.GpgKey) && !dcl.IsZeroValue(desired.GpgKey) {
		c.Config.Logger.Infof("Diff in GpgKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GpgKey), dcl.SprintResource(actual.GpgKey))
		return true
	}
	return false
}

func compareGuestPolicyPackageRepositoriesAptSlice(c *Client, desired, actual []GuestPolicyPackageRepositoriesApt) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesApt, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositoriesApt(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesApt, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesAptMap(c *Client, desired, actual map[string]GuestPolicyPackageRepositoriesApt) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesApt, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesApt, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackageRepositoriesApt(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesApt, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesYum(c *Client, desired, actual *GuestPolicyPackageRepositoriesYum) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Id == nil && desired.Id != nil && !dcl.IsEmptyValueIndirect(desired.Id) {
		c.Config.Logger.Infof("desired Id %s - but actually nil", dcl.SprintResource(desired.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	if actual.DisplayName == nil && desired.DisplayName != nil && !dcl.IsEmptyValueIndirect(desired.DisplayName) {
		c.Config.Logger.Infof("desired DisplayName %s - but actually nil", dcl.SprintResource(desired.DisplayName))
		return true
	}
	if !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) && !dcl.IsZeroValue(desired.DisplayName) {
		c.Config.Logger.Infof("Diff in DisplayName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DisplayName), dcl.SprintResource(actual.DisplayName))
		return true
	}
	if actual.BaseUrl == nil && desired.BaseUrl != nil && !dcl.IsEmptyValueIndirect(desired.BaseUrl) {
		c.Config.Logger.Infof("desired BaseUrl %s - but actually nil", dcl.SprintResource(desired.BaseUrl))
		return true
	}
	if !dcl.StringCanonicalize(desired.BaseUrl, actual.BaseUrl) && !dcl.IsZeroValue(desired.BaseUrl) {
		c.Config.Logger.Infof("Diff in BaseUrl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BaseUrl), dcl.SprintResource(actual.BaseUrl))
		return true
	}
	if actual.GpgKeys == nil && desired.GpgKeys != nil && !dcl.IsEmptyValueIndirect(desired.GpgKeys) {
		c.Config.Logger.Infof("desired GpgKeys %s - but actually nil", dcl.SprintResource(desired.GpgKeys))
		return true
	}
	if !dcl.StringSliceEquals(desired.GpgKeys, actual.GpgKeys) && !dcl.IsZeroValue(desired.GpgKeys) {
		c.Config.Logger.Infof("Diff in GpgKeys. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GpgKeys), dcl.SprintResource(actual.GpgKeys))
		return true
	}
	return false
}

func compareGuestPolicyPackageRepositoriesYumSlice(c *Client, desired, actual []GuestPolicyPackageRepositoriesYum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesYum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositoriesYum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesYum, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesYumMap(c *Client, desired, actual map[string]GuestPolicyPackageRepositoriesYum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesYum, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesYum, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackageRepositoriesYum(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesYum, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesZypper(c *Client, desired, actual *GuestPolicyPackageRepositoriesZypper) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Id == nil && desired.Id != nil && !dcl.IsEmptyValueIndirect(desired.Id) {
		c.Config.Logger.Infof("desired Id %s - but actually nil", dcl.SprintResource(desired.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	if actual.DisplayName == nil && desired.DisplayName != nil && !dcl.IsEmptyValueIndirect(desired.DisplayName) {
		c.Config.Logger.Infof("desired DisplayName %s - but actually nil", dcl.SprintResource(desired.DisplayName))
		return true
	}
	if !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) && !dcl.IsZeroValue(desired.DisplayName) {
		c.Config.Logger.Infof("Diff in DisplayName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DisplayName), dcl.SprintResource(actual.DisplayName))
		return true
	}
	if actual.BaseUrl == nil && desired.BaseUrl != nil && !dcl.IsEmptyValueIndirect(desired.BaseUrl) {
		c.Config.Logger.Infof("desired BaseUrl %s - but actually nil", dcl.SprintResource(desired.BaseUrl))
		return true
	}
	if !dcl.StringCanonicalize(desired.BaseUrl, actual.BaseUrl) && !dcl.IsZeroValue(desired.BaseUrl) {
		c.Config.Logger.Infof("Diff in BaseUrl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BaseUrl), dcl.SprintResource(actual.BaseUrl))
		return true
	}
	if actual.GpgKeys == nil && desired.GpgKeys != nil && !dcl.IsEmptyValueIndirect(desired.GpgKeys) {
		c.Config.Logger.Infof("desired GpgKeys %s - but actually nil", dcl.SprintResource(desired.GpgKeys))
		return true
	}
	if !dcl.StringSliceEquals(desired.GpgKeys, actual.GpgKeys) && !dcl.IsZeroValue(desired.GpgKeys) {
		c.Config.Logger.Infof("Diff in GpgKeys. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GpgKeys), dcl.SprintResource(actual.GpgKeys))
		return true
	}
	return false
}

func compareGuestPolicyPackageRepositoriesZypperSlice(c *Client, desired, actual []GuestPolicyPackageRepositoriesZypper) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesZypper, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositoriesZypper(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesZypper, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesZypperMap(c *Client, desired, actual map[string]GuestPolicyPackageRepositoriesZypper) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesZypper, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesZypper, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackageRepositoriesZypper(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesZypper, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesGoo(c *Client, desired, actual *GuestPolicyPackageRepositoriesGoo) bool {
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
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Url == nil && desired.Url != nil && !dcl.IsEmptyValueIndirect(desired.Url) {
		c.Config.Logger.Infof("desired Url %s - but actually nil", dcl.SprintResource(desired.Url))
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	return false
}

func compareGuestPolicyPackageRepositoriesGooSlice(c *Client, desired, actual []GuestPolicyPackageRepositoriesGoo) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesGoo, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositoriesGoo(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesGoo, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesGooMap(c *Client, desired, actual map[string]GuestPolicyPackageRepositoriesGoo) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesGoo, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesGoo, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyPackageRepositoriesGoo(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesGoo, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipes(c *Client, desired, actual *GuestPolicyRecipes) bool {
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
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Version == nil && desired.Version != nil && !dcl.IsEmptyValueIndirect(desired.Version) {
		c.Config.Logger.Infof("desired Version %s - but actually nil", dcl.SprintResource(desired.Version))
		return true
	}
	if !dcl.StringCanonicalize(desired.Version, actual.Version) && !dcl.IsZeroValue(desired.Version) {
		c.Config.Logger.Infof("Diff in Version. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Version), dcl.SprintResource(actual.Version))
		return true
	}
	if actual.Artifacts == nil && desired.Artifacts != nil && !dcl.IsEmptyValueIndirect(desired.Artifacts) {
		c.Config.Logger.Infof("desired Artifacts %s - but actually nil", dcl.SprintResource(desired.Artifacts))
		return true
	}
	if compareGuestPolicyRecipesArtifactsSlice(c, desired.Artifacts, actual.Artifacts) && !dcl.IsZeroValue(desired.Artifacts) {
		c.Config.Logger.Infof("Diff in Artifacts. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Artifacts), dcl.SprintResource(actual.Artifacts))
		return true
	}
	if actual.InstallSteps == nil && desired.InstallSteps != nil && !dcl.IsEmptyValueIndirect(desired.InstallSteps) {
		c.Config.Logger.Infof("desired InstallSteps %s - but actually nil", dcl.SprintResource(desired.InstallSteps))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsSlice(c, desired.InstallSteps, actual.InstallSteps) && !dcl.IsZeroValue(desired.InstallSteps) {
		c.Config.Logger.Infof("Diff in InstallSteps. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstallSteps), dcl.SprintResource(actual.InstallSteps))
		return true
	}
	if actual.UpdateSteps == nil && desired.UpdateSteps != nil && !dcl.IsEmptyValueIndirect(desired.UpdateSteps) {
		c.Config.Logger.Infof("desired UpdateSteps %s - but actually nil", dcl.SprintResource(desired.UpdateSteps))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsSlice(c, desired.UpdateSteps, actual.UpdateSteps) && !dcl.IsZeroValue(desired.UpdateSteps) {
		c.Config.Logger.Infof("Diff in UpdateSteps. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpdateSteps), dcl.SprintResource(actual.UpdateSteps))
		return true
	}
	if actual.DesiredState == nil && desired.DesiredState != nil && !dcl.IsEmptyValueIndirect(desired.DesiredState) {
		c.Config.Logger.Infof("desired DesiredState %s - but actually nil", dcl.SprintResource(desired.DesiredState))
		return true
	}
	if !reflect.DeepEqual(desired.DesiredState, actual.DesiredState) && !dcl.IsZeroValue(desired.DesiredState) {
		c.Config.Logger.Infof("Diff in DesiredState. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DesiredState), dcl.SprintResource(actual.DesiredState))
		return true
	}
	return false
}

func compareGuestPolicyRecipesSlice(c *Client, desired, actual []GuestPolicyRecipes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipes, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesMap(c *Client, desired, actual map[string]GuestPolicyRecipes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipes, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifacts(c *Client, desired, actual *GuestPolicyRecipesArtifacts) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Id == nil && desired.Id != nil && !dcl.IsEmptyValueIndirect(desired.Id) {
		c.Config.Logger.Infof("desired Id %s - but actually nil", dcl.SprintResource(desired.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	if actual.Remote == nil && desired.Remote != nil && !dcl.IsEmptyValueIndirect(desired.Remote) {
		c.Config.Logger.Infof("desired Remote %s - but actually nil", dcl.SprintResource(desired.Remote))
		return true
	}
	if compareGuestPolicyRecipesArtifactsRemote(c, desired.Remote, actual.Remote) && !dcl.IsZeroValue(desired.Remote) {
		c.Config.Logger.Infof("Diff in Remote. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Remote), dcl.SprintResource(actual.Remote))
		return true
	}
	if actual.Gcs == nil && desired.Gcs != nil && !dcl.IsEmptyValueIndirect(desired.Gcs) {
		c.Config.Logger.Infof("desired Gcs %s - but actually nil", dcl.SprintResource(desired.Gcs))
		return true
	}
	if compareGuestPolicyRecipesArtifactsGcs(c, desired.Gcs, actual.Gcs) && !dcl.IsZeroValue(desired.Gcs) {
		c.Config.Logger.Infof("Diff in Gcs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Gcs), dcl.SprintResource(actual.Gcs))
		return true
	}
	if actual.AllowInsecure == nil && desired.AllowInsecure != nil && !dcl.IsEmptyValueIndirect(desired.AllowInsecure) {
		c.Config.Logger.Infof("desired AllowInsecure %s - but actually nil", dcl.SprintResource(desired.AllowInsecure))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AllowInsecure, actual.AllowInsecure) && !dcl.IsZeroValue(desired.AllowInsecure) {
		c.Config.Logger.Infof("Diff in AllowInsecure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowInsecure), dcl.SprintResource(actual.AllowInsecure))
		return true
	}
	return false
}

func compareGuestPolicyRecipesArtifactsSlice(c *Client, desired, actual []GuestPolicyRecipesArtifacts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifacts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesArtifacts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifacts, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifactsMap(c *Client, desired, actual map[string]GuestPolicyRecipesArtifacts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifacts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifacts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesArtifacts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifacts, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifactsRemote(c *Client, desired, actual *GuestPolicyRecipesArtifactsRemote) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Uri == nil && desired.Uri != nil && !dcl.IsEmptyValueIndirect(desired.Uri) {
		c.Config.Logger.Infof("desired Uri %s - but actually nil", dcl.SprintResource(desired.Uri))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uri, actual.Uri) && !dcl.IsZeroValue(desired.Uri) {
		c.Config.Logger.Infof("Diff in Uri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uri), dcl.SprintResource(actual.Uri))
		return true
	}
	if actual.Checksum == nil && desired.Checksum != nil && !dcl.IsEmptyValueIndirect(desired.Checksum) {
		c.Config.Logger.Infof("desired Checksum %s - but actually nil", dcl.SprintResource(desired.Checksum))
		return true
	}
	if !dcl.StringCanonicalize(desired.Checksum, actual.Checksum) && !dcl.IsZeroValue(desired.Checksum) {
		c.Config.Logger.Infof("Diff in Checksum. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Checksum), dcl.SprintResource(actual.Checksum))
		return true
	}
	return false
}

func compareGuestPolicyRecipesArtifactsRemoteSlice(c *Client, desired, actual []GuestPolicyRecipesArtifactsRemote) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifactsRemote, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesArtifactsRemote(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsRemote, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifactsRemoteMap(c *Client, desired, actual map[string]GuestPolicyRecipesArtifactsRemote) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifactsRemote, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsRemote, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesArtifactsRemote(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsRemote, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifactsGcs(c *Client, desired, actual *GuestPolicyRecipesArtifactsGcs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bucket == nil && desired.Bucket != nil && !dcl.IsEmptyValueIndirect(desired.Bucket) {
		c.Config.Logger.Infof("desired Bucket %s - but actually nil", dcl.SprintResource(desired.Bucket))
		return true
	}
	if !dcl.NameToSelfLink(desired.Bucket, actual.Bucket) && !dcl.IsZeroValue(desired.Bucket) {
		c.Config.Logger.Infof("Diff in Bucket. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bucket), dcl.SprintResource(actual.Bucket))
		return true
	}
	if actual.Object == nil && desired.Object != nil && !dcl.IsEmptyValueIndirect(desired.Object) {
		c.Config.Logger.Infof("desired Object %s - but actually nil", dcl.SprintResource(desired.Object))
		return true
	}
	if !dcl.StringCanonicalize(desired.Object, actual.Object) && !dcl.IsZeroValue(desired.Object) {
		c.Config.Logger.Infof("Diff in Object. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Object), dcl.SprintResource(actual.Object))
		return true
	}
	if actual.Generation == nil && desired.Generation != nil && !dcl.IsEmptyValueIndirect(desired.Generation) {
		c.Config.Logger.Infof("desired Generation %s - but actually nil", dcl.SprintResource(desired.Generation))
		return true
	}
	if !reflect.DeepEqual(desired.Generation, actual.Generation) && !dcl.IsZeroValue(desired.Generation) {
		c.Config.Logger.Infof("Diff in Generation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Generation), dcl.SprintResource(actual.Generation))
		return true
	}
	return false
}

func compareGuestPolicyRecipesArtifactsGcsSlice(c *Client, desired, actual []GuestPolicyRecipesArtifactsGcs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifactsGcs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesArtifactsGcs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsGcs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesArtifactsGcsMap(c *Client, desired, actual map[string]GuestPolicyRecipesArtifactsGcs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesArtifactsGcs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsGcs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesArtifactsGcs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesArtifactsGcs, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallSteps(c *Client, desired, actual *GuestPolicyRecipesInstallSteps) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FileCopy == nil && desired.FileCopy != nil && !dcl.IsEmptyValueIndirect(desired.FileCopy) {
		c.Config.Logger.Infof("desired FileCopy %s - but actually nil", dcl.SprintResource(desired.FileCopy))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsFileCopy(c, desired.FileCopy, actual.FileCopy) && !dcl.IsZeroValue(desired.FileCopy) {
		c.Config.Logger.Infof("Diff in FileCopy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileCopy), dcl.SprintResource(actual.FileCopy))
		return true
	}
	if actual.ArchiveExtraction == nil && desired.ArchiveExtraction != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveExtraction) {
		c.Config.Logger.Infof("desired ArchiveExtraction %s - but actually nil", dcl.SprintResource(desired.ArchiveExtraction))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsArchiveExtraction(c, desired.ArchiveExtraction, actual.ArchiveExtraction) && !dcl.IsZeroValue(desired.ArchiveExtraction) {
		c.Config.Logger.Infof("Diff in ArchiveExtraction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveExtraction), dcl.SprintResource(actual.ArchiveExtraction))
		return true
	}
	if actual.MsiInstallation == nil && desired.MsiInstallation != nil && !dcl.IsEmptyValueIndirect(desired.MsiInstallation) {
		c.Config.Logger.Infof("desired MsiInstallation %s - but actually nil", dcl.SprintResource(desired.MsiInstallation))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsMsiInstallation(c, desired.MsiInstallation, actual.MsiInstallation) && !dcl.IsZeroValue(desired.MsiInstallation) {
		c.Config.Logger.Infof("Diff in MsiInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MsiInstallation), dcl.SprintResource(actual.MsiInstallation))
		return true
	}
	if actual.DpkgInstallation == nil && desired.DpkgInstallation != nil && !dcl.IsEmptyValueIndirect(desired.DpkgInstallation) {
		c.Config.Logger.Infof("desired DpkgInstallation %s - but actually nil", dcl.SprintResource(desired.DpkgInstallation))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsDpkgInstallation(c, desired.DpkgInstallation, actual.DpkgInstallation) && !dcl.IsZeroValue(desired.DpkgInstallation) {
		c.Config.Logger.Infof("Diff in DpkgInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DpkgInstallation), dcl.SprintResource(actual.DpkgInstallation))
		return true
	}
	if actual.RpmInstallation == nil && desired.RpmInstallation != nil && !dcl.IsEmptyValueIndirect(desired.RpmInstallation) {
		c.Config.Logger.Infof("desired RpmInstallation %s - but actually nil", dcl.SprintResource(desired.RpmInstallation))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsRpmInstallation(c, desired.RpmInstallation, actual.RpmInstallation) && !dcl.IsZeroValue(desired.RpmInstallation) {
		c.Config.Logger.Infof("Diff in RpmInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RpmInstallation), dcl.SprintResource(actual.RpmInstallation))
		return true
	}
	if actual.FileExec == nil && desired.FileExec != nil && !dcl.IsEmptyValueIndirect(desired.FileExec) {
		c.Config.Logger.Infof("desired FileExec %s - but actually nil", dcl.SprintResource(desired.FileExec))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsFileExec(c, desired.FileExec, actual.FileExec) && !dcl.IsZeroValue(desired.FileExec) {
		c.Config.Logger.Infof("Diff in FileExec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileExec), dcl.SprintResource(actual.FileExec))
		return true
	}
	if actual.ScriptRun == nil && desired.ScriptRun != nil && !dcl.IsEmptyValueIndirect(desired.ScriptRun) {
		c.Config.Logger.Infof("desired ScriptRun %s - but actually nil", dcl.SprintResource(desired.ScriptRun))
		return true
	}
	if compareGuestPolicyRecipesInstallStepsScriptRun(c, desired.ScriptRun, actual.ScriptRun) && !dcl.IsZeroValue(desired.ScriptRun) {
		c.Config.Logger.Infof("Diff in ScriptRun. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptRun), dcl.SprintResource(actual.ScriptRun))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsSlice(c *Client, desired, actual []GuestPolicyRecipesInstallSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallSteps, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallSteps(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallSteps, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallSteps, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallSteps, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallSteps(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallSteps, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileCopy(c *Client, desired, actual *GuestPolicyRecipesInstallStepsFileCopy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Destination == nil && desired.Destination != nil && !dcl.IsEmptyValueIndirect(desired.Destination) {
		c.Config.Logger.Infof("desired Destination %s - but actually nil", dcl.SprintResource(desired.Destination))
		return true
	}
	if !dcl.StringCanonicalize(desired.Destination, actual.Destination) && !dcl.IsZeroValue(desired.Destination) {
		c.Config.Logger.Infof("Diff in Destination. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Destination), dcl.SprintResource(actual.Destination))
		return true
	}
	if actual.Overwrite == nil && desired.Overwrite != nil && !dcl.IsEmptyValueIndirect(desired.Overwrite) {
		c.Config.Logger.Infof("desired Overwrite %s - but actually nil", dcl.SprintResource(desired.Overwrite))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Overwrite, actual.Overwrite) && !dcl.IsZeroValue(desired.Overwrite) {
		c.Config.Logger.Infof("Diff in Overwrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Overwrite), dcl.SprintResource(actual.Overwrite))
		return true
	}
	if actual.Permissions == nil && desired.Permissions != nil && !dcl.IsEmptyValueIndirect(desired.Permissions) {
		c.Config.Logger.Infof("desired Permissions %s - but actually nil", dcl.SprintResource(desired.Permissions))
		return true
	}
	if !dcl.StringCanonicalize(desired.Permissions, actual.Permissions) && !dcl.IsZeroValue(desired.Permissions) {
		c.Config.Logger.Infof("Diff in Permissions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Permissions), dcl.SprintResource(actual.Permissions))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileCopySlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsFileCopy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsFileCopy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsFileCopy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileCopy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileCopyMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsFileCopy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsFileCopy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileCopy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsFileCopy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileCopy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsArchiveExtraction(c *Client, desired, actual *GuestPolicyRecipesInstallStepsArchiveExtraction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Destination == nil && desired.Destination != nil && !dcl.IsEmptyValueIndirect(desired.Destination) {
		c.Config.Logger.Infof("desired Destination %s - but actually nil", dcl.SprintResource(desired.Destination))
		return true
	}
	if !dcl.StringCanonicalize(desired.Destination, actual.Destination) && !dcl.IsZeroValue(desired.Destination) {
		c.Config.Logger.Infof("Diff in Destination. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Destination), dcl.SprintResource(actual.Destination))
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsArchiveExtractionSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsArchiveExtraction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsArchiveExtraction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsArchiveExtraction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsArchiveExtraction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsArchiveExtractionMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsArchiveExtraction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsArchiveExtraction, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsArchiveExtraction, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsArchiveExtraction(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsArchiveExtraction, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsMsiInstallation(c *Client, desired, actual *GuestPolicyRecipesInstallStepsMsiInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Flags == nil && desired.Flags != nil && !dcl.IsEmptyValueIndirect(desired.Flags) {
		c.Config.Logger.Infof("desired Flags %s - but actually nil", dcl.SprintResource(desired.Flags))
		return true
	}
	if !dcl.StringSliceEquals(desired.Flags, actual.Flags) && !dcl.IsZeroValue(desired.Flags) {
		c.Config.Logger.Infof("Diff in Flags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Flags), dcl.SprintResource(actual.Flags))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsMsiInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsMsiInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsMsiInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsMsiInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsMsiInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsMsiInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsMsiInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsMsiInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsMsiInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsMsiInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsMsiInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsDpkgInstallation(c *Client, desired, actual *GuestPolicyRecipesInstallStepsDpkgInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsDpkgInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsDpkgInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsDpkgInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsDpkgInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsDpkgInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsDpkgInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsDpkgInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsDpkgInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsDpkgInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsDpkgInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsDpkgInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsRpmInstallation(c *Client, desired, actual *GuestPolicyRecipesInstallStepsRpmInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsRpmInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsRpmInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsRpmInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsRpmInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsRpmInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsRpmInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsRpmInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsRpmInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsRpmInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsRpmInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsRpmInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileExec(c *Client, desired, actual *GuestPolicyRecipesInstallStepsFileExec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.LocalPath == nil && desired.LocalPath != nil && !dcl.IsEmptyValueIndirect(desired.LocalPath) {
		c.Config.Logger.Infof("desired LocalPath %s - but actually nil", dcl.SprintResource(desired.LocalPath))
		return true
	}
	if !dcl.StringCanonicalize(desired.LocalPath, actual.LocalPath) && !dcl.IsZeroValue(desired.LocalPath) {
		c.Config.Logger.Infof("Diff in LocalPath. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalPath), dcl.SprintResource(actual.LocalPath))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.StringSliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileExecSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsFileExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsFileExec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsFileExec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileExec, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsFileExecMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsFileExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsFileExec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileExec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsFileExec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsFileExec, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsScriptRun(c *Client, desired, actual *GuestPolicyRecipesInstallStepsScriptRun) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Script == nil && desired.Script != nil && !dcl.IsEmptyValueIndirect(desired.Script) {
		c.Config.Logger.Infof("desired Script %s - but actually nil", dcl.SprintResource(desired.Script))
		return true
	}
	if !dcl.StringCanonicalize(desired.Script, actual.Script) && !dcl.IsZeroValue(desired.Script) {
		c.Config.Logger.Infof("Diff in Script. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Script), dcl.SprintResource(actual.Script))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	if actual.Interpreter == nil && desired.Interpreter != nil && !dcl.IsEmptyValueIndirect(desired.Interpreter) {
		c.Config.Logger.Infof("desired Interpreter %s - but actually nil", dcl.SprintResource(desired.Interpreter))
		return true
	}
	if !reflect.DeepEqual(desired.Interpreter, actual.Interpreter) && !dcl.IsZeroValue(desired.Interpreter) {
		c.Config.Logger.Infof("Diff in Interpreter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Interpreter), dcl.SprintResource(actual.Interpreter))
		return true
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsScriptRunSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsScriptRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsScriptRun, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsScriptRun(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsScriptRun, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsScriptRunMap(c *Client, desired, actual map[string]GuestPolicyRecipesInstallStepsScriptRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsScriptRun, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsScriptRun, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesInstallStepsScriptRun(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsScriptRun, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateSteps(c *Client, desired, actual *GuestPolicyRecipesUpdateSteps) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FileCopy == nil && desired.FileCopy != nil && !dcl.IsEmptyValueIndirect(desired.FileCopy) {
		c.Config.Logger.Infof("desired FileCopy %s - but actually nil", dcl.SprintResource(desired.FileCopy))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsFileCopy(c, desired.FileCopy, actual.FileCopy) && !dcl.IsZeroValue(desired.FileCopy) {
		c.Config.Logger.Infof("Diff in FileCopy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileCopy), dcl.SprintResource(actual.FileCopy))
		return true
	}
	if actual.ArchiveExtraction == nil && desired.ArchiveExtraction != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveExtraction) {
		c.Config.Logger.Infof("desired ArchiveExtraction %s - but actually nil", dcl.SprintResource(desired.ArchiveExtraction))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsArchiveExtraction(c, desired.ArchiveExtraction, actual.ArchiveExtraction) && !dcl.IsZeroValue(desired.ArchiveExtraction) {
		c.Config.Logger.Infof("Diff in ArchiveExtraction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveExtraction), dcl.SprintResource(actual.ArchiveExtraction))
		return true
	}
	if actual.MsiInstallation == nil && desired.MsiInstallation != nil && !dcl.IsEmptyValueIndirect(desired.MsiInstallation) {
		c.Config.Logger.Infof("desired MsiInstallation %s - but actually nil", dcl.SprintResource(desired.MsiInstallation))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsMsiInstallation(c, desired.MsiInstallation, actual.MsiInstallation) && !dcl.IsZeroValue(desired.MsiInstallation) {
		c.Config.Logger.Infof("Diff in MsiInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MsiInstallation), dcl.SprintResource(actual.MsiInstallation))
		return true
	}
	if actual.DpkgInstallation == nil && desired.DpkgInstallation != nil && !dcl.IsEmptyValueIndirect(desired.DpkgInstallation) {
		c.Config.Logger.Infof("desired DpkgInstallation %s - but actually nil", dcl.SprintResource(desired.DpkgInstallation))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsDpkgInstallation(c, desired.DpkgInstallation, actual.DpkgInstallation) && !dcl.IsZeroValue(desired.DpkgInstallation) {
		c.Config.Logger.Infof("Diff in DpkgInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DpkgInstallation), dcl.SprintResource(actual.DpkgInstallation))
		return true
	}
	if actual.RpmInstallation == nil && desired.RpmInstallation != nil && !dcl.IsEmptyValueIndirect(desired.RpmInstallation) {
		c.Config.Logger.Infof("desired RpmInstallation %s - but actually nil", dcl.SprintResource(desired.RpmInstallation))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsRpmInstallation(c, desired.RpmInstallation, actual.RpmInstallation) && !dcl.IsZeroValue(desired.RpmInstallation) {
		c.Config.Logger.Infof("Diff in RpmInstallation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RpmInstallation), dcl.SprintResource(actual.RpmInstallation))
		return true
	}
	if actual.FileExec == nil && desired.FileExec != nil && !dcl.IsEmptyValueIndirect(desired.FileExec) {
		c.Config.Logger.Infof("desired FileExec %s - but actually nil", dcl.SprintResource(desired.FileExec))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsFileExec(c, desired.FileExec, actual.FileExec) && !dcl.IsZeroValue(desired.FileExec) {
		c.Config.Logger.Infof("Diff in FileExec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileExec), dcl.SprintResource(actual.FileExec))
		return true
	}
	if actual.ScriptRun == nil && desired.ScriptRun != nil && !dcl.IsEmptyValueIndirect(desired.ScriptRun) {
		c.Config.Logger.Infof("desired ScriptRun %s - but actually nil", dcl.SprintResource(desired.ScriptRun))
		return true
	}
	if compareGuestPolicyRecipesUpdateStepsScriptRun(c, desired.ScriptRun, actual.ScriptRun) && !dcl.IsZeroValue(desired.ScriptRun) {
		c.Config.Logger.Infof("Diff in ScriptRun. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptRun), dcl.SprintResource(actual.ScriptRun))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateSteps, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateSteps(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateSteps, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateSteps, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateSteps, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateSteps(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateSteps, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileCopy(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsFileCopy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Destination == nil && desired.Destination != nil && !dcl.IsEmptyValueIndirect(desired.Destination) {
		c.Config.Logger.Infof("desired Destination %s - but actually nil", dcl.SprintResource(desired.Destination))
		return true
	}
	if !dcl.StringCanonicalize(desired.Destination, actual.Destination) && !dcl.IsZeroValue(desired.Destination) {
		c.Config.Logger.Infof("Diff in Destination. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Destination), dcl.SprintResource(actual.Destination))
		return true
	}
	if actual.Overwrite == nil && desired.Overwrite != nil && !dcl.IsEmptyValueIndirect(desired.Overwrite) {
		c.Config.Logger.Infof("desired Overwrite %s - but actually nil", dcl.SprintResource(desired.Overwrite))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Overwrite, actual.Overwrite) && !dcl.IsZeroValue(desired.Overwrite) {
		c.Config.Logger.Infof("Diff in Overwrite. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Overwrite), dcl.SprintResource(actual.Overwrite))
		return true
	}
	if actual.Permissions == nil && desired.Permissions != nil && !dcl.IsEmptyValueIndirect(desired.Permissions) {
		c.Config.Logger.Infof("desired Permissions %s - but actually nil", dcl.SprintResource(desired.Permissions))
		return true
	}
	if !dcl.StringCanonicalize(desired.Permissions, actual.Permissions) && !dcl.IsZeroValue(desired.Permissions) {
		c.Config.Logger.Infof("Diff in Permissions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Permissions), dcl.SprintResource(actual.Permissions))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileCopySlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsFileCopy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsFileCopy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsFileCopy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileCopy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileCopyMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsFileCopy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsFileCopy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileCopy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsFileCopy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileCopy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtraction(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsArchiveExtraction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Destination == nil && desired.Destination != nil && !dcl.IsEmptyValueIndirect(desired.Destination) {
		c.Config.Logger.Infof("desired Destination %s - but actually nil", dcl.SprintResource(desired.Destination))
		return true
	}
	if !dcl.StringCanonicalize(desired.Destination, actual.Destination) && !dcl.IsZeroValue(desired.Destination) {
		c.Config.Logger.Infof("Diff in Destination. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Destination), dcl.SprintResource(actual.Destination))
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtractionSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsArchiveExtraction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsArchiveExtraction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsArchiveExtraction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtractionMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsArchiveExtraction, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsArchiveExtraction, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsArchiveExtraction, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsMsiInstallation(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsMsiInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.Flags == nil && desired.Flags != nil && !dcl.IsEmptyValueIndirect(desired.Flags) {
		c.Config.Logger.Infof("desired Flags %s - but actually nil", dcl.SprintResource(desired.Flags))
		return true
	}
	if !dcl.StringSliceEquals(desired.Flags, actual.Flags) && !dcl.IsZeroValue(desired.Flags) {
		c.Config.Logger.Infof("Diff in Flags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Flags), dcl.SprintResource(actual.Flags))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsMsiInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsMsiInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsMsiInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsMsiInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsMsiInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsMsiInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsMsiInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsMsiInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsMsiInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsMsiInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsMsiInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsDpkgInstallation(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsDpkgInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsDpkgInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsDpkgInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsDpkgInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsDpkgInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsDpkgInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsDpkgInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsDpkgInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsDpkgInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsRpmInstallation(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsRpmInstallation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsRpmInstallationSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsRpmInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsRpmInstallation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsRpmInstallation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsRpmInstallation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsRpmInstallationMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsRpmInstallation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsRpmInstallation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsRpmInstallation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsRpmInstallation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsRpmInstallation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileExec(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsFileExec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ArtifactId == nil && desired.ArtifactId != nil && !dcl.IsEmptyValueIndirect(desired.ArtifactId) {
		c.Config.Logger.Infof("desired ArtifactId %s - but actually nil", dcl.SprintResource(desired.ArtifactId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ArtifactId, actual.ArtifactId) && !dcl.IsZeroValue(desired.ArtifactId) {
		c.Config.Logger.Infof("Diff in ArtifactId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArtifactId), dcl.SprintResource(actual.ArtifactId))
		return true
	}
	if actual.LocalPath == nil && desired.LocalPath != nil && !dcl.IsEmptyValueIndirect(desired.LocalPath) {
		c.Config.Logger.Infof("desired LocalPath %s - but actually nil", dcl.SprintResource(desired.LocalPath))
		return true
	}
	if !dcl.StringCanonicalize(desired.LocalPath, actual.LocalPath) && !dcl.IsZeroValue(desired.LocalPath) {
		c.Config.Logger.Infof("Diff in LocalPath. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalPath), dcl.SprintResource(actual.LocalPath))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.StringSliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileExecSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsFileExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsFileExec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsFileExec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileExec, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsFileExecMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsFileExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsFileExec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileExec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsFileExec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsFileExec, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsScriptRun(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsScriptRun) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Script == nil && desired.Script != nil && !dcl.IsEmptyValueIndirect(desired.Script) {
		c.Config.Logger.Infof("desired Script %s - but actually nil", dcl.SprintResource(desired.Script))
		return true
	}
	if !dcl.StringCanonicalize(desired.Script, actual.Script) && !dcl.IsZeroValue(desired.Script) {
		c.Config.Logger.Infof("Diff in Script. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Script), dcl.SprintResource(actual.Script))
		return true
	}
	if actual.AllowedExitCodes == nil && desired.AllowedExitCodes != nil && !dcl.IsEmptyValueIndirect(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("desired AllowedExitCodes %s - but actually nil", dcl.SprintResource(desired.AllowedExitCodes))
		return true
	}
	if !dcl.IntSliceEquals(desired.AllowedExitCodes, actual.AllowedExitCodes) && !dcl.IsZeroValue(desired.AllowedExitCodes) {
		c.Config.Logger.Infof("Diff in AllowedExitCodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedExitCodes), dcl.SprintResource(actual.AllowedExitCodes))
		return true
	}
	if actual.Interpreter == nil && desired.Interpreter != nil && !dcl.IsEmptyValueIndirect(desired.Interpreter) {
		c.Config.Logger.Infof("desired Interpreter %s - but actually nil", dcl.SprintResource(desired.Interpreter))
		return true
	}
	if !reflect.DeepEqual(desired.Interpreter, actual.Interpreter) && !dcl.IsZeroValue(desired.Interpreter) {
		c.Config.Logger.Infof("Diff in Interpreter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Interpreter), dcl.SprintResource(actual.Interpreter))
		return true
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsScriptRunSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsScriptRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsScriptRun, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsScriptRun(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsScriptRun, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsScriptRunMap(c *Client, desired, actual map[string]GuestPolicyRecipesUpdateStepsScriptRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsScriptRun, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsScriptRun, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareGuestPolicyRecipesUpdateStepsScriptRun(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsScriptRun, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackagesDesiredStateEnumSlice(c *Client, desired, actual []GuestPolicyPackagesDesiredStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackagesDesiredStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackagesDesiredStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackagesDesiredStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackagesDesiredStateEnum(c *Client, desired, actual *GuestPolicyPackagesDesiredStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyPackagesManagerEnumSlice(c *Client, desired, actual []GuestPolicyPackagesManagerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackagesManagerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackagesManagerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackagesManagerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackagesManagerEnum(c *Client, desired, actual *GuestPolicyPackagesManagerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyPackageRepositoriesAptArchiveTypeEnumSlice(c *Client, desired, actual []GuestPolicyPackageRepositoriesAptArchiveTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyPackageRepositoriesAptArchiveTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyPackageRepositoriesAptArchiveTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyPackageRepositoriesAptArchiveTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyPackageRepositoriesAptArchiveTypeEnum(c *Client, desired, actual *GuestPolicyPackageRepositoriesAptArchiveTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(c *Client, desired, actual *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumSlice(c *Client, desired, actual []GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(c *Client, desired, actual *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumSlice(c *Client, desired, actual []GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(c *Client, desired, actual *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareGuestPolicyRecipesDesiredStateEnumSlice(c *Client, desired, actual []GuestPolicyRecipesDesiredStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in GuestPolicyRecipesDesiredStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareGuestPolicyRecipesDesiredStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in GuestPolicyRecipesDesiredStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareGuestPolicyRecipesDesiredStateEnum(c *Client, desired, actual *GuestPolicyRecipesDesiredStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *GuestPolicy) urlNormalized() *GuestPolicy {
	normalized := deepcopy.Copy(*r).(GuestPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *GuestPolicy) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *GuestPolicy) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *GuestPolicy) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *GuestPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateGuestPolicy" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/guestPolicies/{{name}}", "https://osconfig.googleapis.com/v1beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the GuestPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *GuestPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandGuestPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling GuestPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGuestPolicy decodes JSON responses into the GuestPolicy resource schema.
func unmarshalGuestPolicy(b []byte, c *Client) (*GuestPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGuestPolicy(m, c)
}

func unmarshalMapGuestPolicy(m map[string]interface{}, c *Client) (*GuestPolicy, error) {

	return flattenGuestPolicy(c, m), nil
}

// expandGuestPolicy expands GuestPolicy into a JSON request object.
func expandGuestPolicy(c *Client, f *GuestPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := expandGuestPolicyAssignment(c, f.Assignment); err != nil {
		return nil, fmt.Errorf("error expanding Assignment into assignment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["assignment"] = v
	}
	if v, err := expandGuestPolicyPackagesSlice(c, f.Packages); err != nil {
		return nil, fmt.Errorf("error expanding Packages into packages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["packages"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesSlice(c, f.PackageRepositories); err != nil {
		return nil, fmt.Errorf("error expanding PackageRepositories into packageRepositories: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["packageRepositories"] = v
	}
	if v, err := expandGuestPolicyRecipesSlice(c, f.Recipes); err != nil {
		return nil, fmt.Errorf("error expanding Recipes into recipes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["recipes"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenGuestPolicy flattens GuestPolicy from a JSON request object into the
// GuestPolicy type.
func flattenGuestPolicy(c *Client, i interface{}) *GuestPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &GuestPolicy{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Assignment = flattenGuestPolicyAssignment(c, m["assignment"])
	r.Packages = flattenGuestPolicyPackagesSlice(c, m["packages"])
	r.PackageRepositories = flattenGuestPolicyPackageRepositoriesSlice(c, m["packageRepositories"])
	r.Recipes = flattenGuestPolicyRecipesSlice(c, m["recipes"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandGuestPolicyAssignmentMap expands the contents of GuestPolicyAssignment into a JSON
// request object.
func expandGuestPolicyAssignmentMap(c *Client, f map[string]GuestPolicyAssignment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyAssignment(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyAssignmentSlice expands the contents of GuestPolicyAssignment into a JSON
// request object.
func expandGuestPolicyAssignmentSlice(c *Client, f []GuestPolicyAssignment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyAssignment(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyAssignmentMap flattens the contents of GuestPolicyAssignment from a JSON
// response object.
func flattenGuestPolicyAssignmentMap(c *Client, i interface{}) map[string]GuestPolicyAssignment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyAssignment{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyAssignment{}
	}

	items := make(map[string]GuestPolicyAssignment)
	for k, item := range a {
		items[k] = *flattenGuestPolicyAssignment(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyAssignmentSlice flattens the contents of GuestPolicyAssignment from a JSON
// response object.
func flattenGuestPolicyAssignmentSlice(c *Client, i interface{}) []GuestPolicyAssignment {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyAssignment{}
	}

	if len(a) == 0 {
		return []GuestPolicyAssignment{}
	}

	items := make([]GuestPolicyAssignment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyAssignment(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyAssignment expands an instance of GuestPolicyAssignment into a JSON
// request object.
func expandGuestPolicyAssignment(c *Client, f *GuestPolicyAssignment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGuestPolicyAssignmentGroupLabelsSlice(c, f.GroupLabels); err != nil {
		return nil, fmt.Errorf("error expanding GroupLabels into groupLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["groupLabels"] = v
	}
	if v := f.Zones; !dcl.IsEmptyValueIndirect(v) {
		m["zones"] = v
	}
	if v, err := expandGuestPolicyInstances(f, f.Instances); err != nil {
		return nil, fmt.Errorf("error expanding Instances into instances: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["instances"] = v
	}
	if v := f.InstanceNamePrefixes; !dcl.IsEmptyValueIndirect(v) {
		m["instanceNamePrefixes"] = v
	}
	if v, err := expandGuestPolicyAssignmentOsTypesSlice(c, f.OsTypes); err != nil {
		return nil, fmt.Errorf("error expanding OsTypes into osTypes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["osTypes"] = v
	}

	return m, nil
}

// flattenGuestPolicyAssignment flattens an instance of GuestPolicyAssignment from a JSON
// response object.
func flattenGuestPolicyAssignment(c *Client, i interface{}) *GuestPolicyAssignment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyAssignment{}
	r.GroupLabels = flattenGuestPolicyAssignmentGroupLabelsSlice(c, m["groupLabels"])
	r.Zones = dcl.FlattenStringSlice(m["zones"])
	r.Instances = flattenGuestPolicyInstances(m["instances"])
	r.InstanceNamePrefixes = dcl.FlattenStringSlice(m["instanceNamePrefixes"])
	r.OsTypes = flattenGuestPolicyAssignmentOsTypesSlice(c, m["osTypes"])

	return r
}

// expandGuestPolicyAssignmentGroupLabelsMap expands the contents of GuestPolicyAssignmentGroupLabels into a JSON
// request object.
func expandGuestPolicyAssignmentGroupLabelsMap(c *Client, f map[string]GuestPolicyAssignmentGroupLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyAssignmentGroupLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyAssignmentGroupLabelsSlice expands the contents of GuestPolicyAssignmentGroupLabels into a JSON
// request object.
func expandGuestPolicyAssignmentGroupLabelsSlice(c *Client, f []GuestPolicyAssignmentGroupLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyAssignmentGroupLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyAssignmentGroupLabelsMap flattens the contents of GuestPolicyAssignmentGroupLabels from a JSON
// response object.
func flattenGuestPolicyAssignmentGroupLabelsMap(c *Client, i interface{}) map[string]GuestPolicyAssignmentGroupLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyAssignmentGroupLabels{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyAssignmentGroupLabels{}
	}

	items := make(map[string]GuestPolicyAssignmentGroupLabels)
	for k, item := range a {
		items[k] = *flattenGuestPolicyAssignmentGroupLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyAssignmentGroupLabelsSlice flattens the contents of GuestPolicyAssignmentGroupLabels from a JSON
// response object.
func flattenGuestPolicyAssignmentGroupLabelsSlice(c *Client, i interface{}) []GuestPolicyAssignmentGroupLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyAssignmentGroupLabels{}
	}

	if len(a) == 0 {
		return []GuestPolicyAssignmentGroupLabels{}
	}

	items := make([]GuestPolicyAssignmentGroupLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyAssignmentGroupLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyAssignmentGroupLabels expands an instance of GuestPolicyAssignmentGroupLabels into a JSON
// request object.
func expandGuestPolicyAssignmentGroupLabels(c *Client, f *GuestPolicyAssignmentGroupLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenGuestPolicyAssignmentGroupLabels flattens an instance of GuestPolicyAssignmentGroupLabels from a JSON
// response object.
func flattenGuestPolicyAssignmentGroupLabels(c *Client, i interface{}) *GuestPolicyAssignmentGroupLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyAssignmentGroupLabels{}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandGuestPolicyAssignmentOsTypesMap expands the contents of GuestPolicyAssignmentOsTypes into a JSON
// request object.
func expandGuestPolicyAssignmentOsTypesMap(c *Client, f map[string]GuestPolicyAssignmentOsTypes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyAssignmentOsTypes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyAssignmentOsTypesSlice expands the contents of GuestPolicyAssignmentOsTypes into a JSON
// request object.
func expandGuestPolicyAssignmentOsTypesSlice(c *Client, f []GuestPolicyAssignmentOsTypes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyAssignmentOsTypes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyAssignmentOsTypesMap flattens the contents of GuestPolicyAssignmentOsTypes from a JSON
// response object.
func flattenGuestPolicyAssignmentOsTypesMap(c *Client, i interface{}) map[string]GuestPolicyAssignmentOsTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyAssignmentOsTypes{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyAssignmentOsTypes{}
	}

	items := make(map[string]GuestPolicyAssignmentOsTypes)
	for k, item := range a {
		items[k] = *flattenGuestPolicyAssignmentOsTypes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyAssignmentOsTypesSlice flattens the contents of GuestPolicyAssignmentOsTypes from a JSON
// response object.
func flattenGuestPolicyAssignmentOsTypesSlice(c *Client, i interface{}) []GuestPolicyAssignmentOsTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyAssignmentOsTypes{}
	}

	if len(a) == 0 {
		return []GuestPolicyAssignmentOsTypes{}
	}

	items := make([]GuestPolicyAssignmentOsTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyAssignmentOsTypes(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyAssignmentOsTypes expands an instance of GuestPolicyAssignmentOsTypes into a JSON
// request object.
func expandGuestPolicyAssignmentOsTypes(c *Client, f *GuestPolicyAssignmentOsTypes) (map[string]interface{}, error) {
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
	if v := f.OsArchitecture; !dcl.IsEmptyValueIndirect(v) {
		m["osArchitecture"] = v
	}

	return m, nil
}

// flattenGuestPolicyAssignmentOsTypes flattens an instance of GuestPolicyAssignmentOsTypes from a JSON
// response object.
func flattenGuestPolicyAssignmentOsTypes(c *Client, i interface{}) *GuestPolicyAssignmentOsTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyAssignmentOsTypes{}
	r.OsShortName = dcl.FlattenString(m["osShortName"])
	r.OsVersion = dcl.FlattenString(m["osVersion"])
	r.OsArchitecture = dcl.FlattenString(m["osArchitecture"])

	return r
}

// expandGuestPolicyPackagesMap expands the contents of GuestPolicyPackages into a JSON
// request object.
func expandGuestPolicyPackagesMap(c *Client, f map[string]GuestPolicyPackages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackagesSlice expands the contents of GuestPolicyPackages into a JSON
// request object.
func expandGuestPolicyPackagesSlice(c *Client, f []GuestPolicyPackages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackagesMap flattens the contents of GuestPolicyPackages from a JSON
// response object.
func flattenGuestPolicyPackagesMap(c *Client, i interface{}) map[string]GuestPolicyPackages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackages{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackages{}
	}

	items := make(map[string]GuestPolicyPackages)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackagesSlice flattens the contents of GuestPolicyPackages from a JSON
// response object.
func flattenGuestPolicyPackagesSlice(c *Client, i interface{}) []GuestPolicyPackages {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackages{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackages{}
	}

	items := make([]GuestPolicyPackages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackages(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackages expands an instance of GuestPolicyPackages into a JSON
// request object.
func expandGuestPolicyPackages(c *Client, f *GuestPolicyPackages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DesiredState; !dcl.IsEmptyValueIndirect(v) {
		m["desiredState"] = v
	}
	if v := f.Manager; !dcl.IsEmptyValueIndirect(v) {
		m["manager"] = v
	}

	return m, nil
}

// flattenGuestPolicyPackages flattens an instance of GuestPolicyPackages from a JSON
// response object.
func flattenGuestPolicyPackages(c *Client, i interface{}) *GuestPolicyPackages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackages{}
	r.Name = dcl.FlattenString(m["name"])
	r.DesiredState = flattenGuestPolicyPackagesDesiredStateEnum(m["desiredState"])
	r.Manager = flattenGuestPolicyPackagesManagerEnum(m["manager"])

	return r
}

// expandGuestPolicyPackageRepositoriesMap expands the contents of GuestPolicyPackageRepositories into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesMap(c *Client, f map[string]GuestPolicyPackageRepositories) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackageRepositories(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackageRepositoriesSlice expands the contents of GuestPolicyPackageRepositories into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesSlice(c *Client, f []GuestPolicyPackageRepositories) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackageRepositories(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackageRepositoriesMap flattens the contents of GuestPolicyPackageRepositories from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesMap(c *Client, i interface{}) map[string]GuestPolicyPackageRepositories {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackageRepositories{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackageRepositories{}
	}

	items := make(map[string]GuestPolicyPackageRepositories)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackageRepositories(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesSlice flattens the contents of GuestPolicyPackageRepositories from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesSlice(c *Client, i interface{}) []GuestPolicyPackageRepositories {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositories{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositories{}
	}

	items := make([]GuestPolicyPackageRepositories, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositories(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackageRepositories expands an instance of GuestPolicyPackageRepositories into a JSON
// request object.
func expandGuestPolicyPackageRepositories(c *Client, f *GuestPolicyPackageRepositories) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGuestPolicyPackageRepositoriesApt(c, f.Apt); err != nil {
		return nil, fmt.Errorf("error expanding Apt into apt: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apt"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesYum(c, f.Yum); err != nil {
		return nil, fmt.Errorf("error expanding Yum into yum: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yum"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesZypper(c, f.Zypper); err != nil {
		return nil, fmt.Errorf("error expanding Zypper into zypper: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zypper"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesGoo(c, f.Goo); err != nil {
		return nil, fmt.Errorf("error expanding Goo into goo: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["goo"] = v
	}

	return m, nil
}

// flattenGuestPolicyPackageRepositories flattens an instance of GuestPolicyPackageRepositories from a JSON
// response object.
func flattenGuestPolicyPackageRepositories(c *Client, i interface{}) *GuestPolicyPackageRepositories {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackageRepositories{}
	r.Apt = flattenGuestPolicyPackageRepositoriesApt(c, m["apt"])
	r.Yum = flattenGuestPolicyPackageRepositoriesYum(c, m["yum"])
	r.Zypper = flattenGuestPolicyPackageRepositoriesZypper(c, m["zypper"])
	r.Goo = flattenGuestPolicyPackageRepositoriesGoo(c, m["goo"])

	return r
}

// expandGuestPolicyPackageRepositoriesAptMap expands the contents of GuestPolicyPackageRepositoriesApt into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesAptMap(c *Client, f map[string]GuestPolicyPackageRepositoriesApt) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesApt(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackageRepositoriesAptSlice expands the contents of GuestPolicyPackageRepositoriesApt into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesAptSlice(c *Client, f []GuestPolicyPackageRepositoriesApt) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesApt(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackageRepositoriesAptMap flattens the contents of GuestPolicyPackageRepositoriesApt from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesAptMap(c *Client, i interface{}) map[string]GuestPolicyPackageRepositoriesApt {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackageRepositoriesApt{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackageRepositoriesApt{}
	}

	items := make(map[string]GuestPolicyPackageRepositoriesApt)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackageRepositoriesApt(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesAptSlice flattens the contents of GuestPolicyPackageRepositoriesApt from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesAptSlice(c *Client, i interface{}) []GuestPolicyPackageRepositoriesApt {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositoriesApt{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositoriesApt{}
	}

	items := make([]GuestPolicyPackageRepositoriesApt, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositoriesApt(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackageRepositoriesApt expands an instance of GuestPolicyPackageRepositoriesApt into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesApt(c *Client, f *GuestPolicyPackageRepositoriesApt) (map[string]interface{}, error) {
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

// flattenGuestPolicyPackageRepositoriesApt flattens an instance of GuestPolicyPackageRepositoriesApt from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesApt(c *Client, i interface{}) *GuestPolicyPackageRepositoriesApt {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackageRepositoriesApt{}
	r.ArchiveType = flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnum(m["archiveType"])
	r.Uri = dcl.FlattenString(m["uri"])
	r.Distribution = dcl.FlattenString(m["distribution"])
	r.Components = dcl.FlattenStringSlice(m["components"])
	r.GpgKey = dcl.FlattenString(m["gpgKey"])

	return r
}

// expandGuestPolicyPackageRepositoriesYumMap expands the contents of GuestPolicyPackageRepositoriesYum into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesYumMap(c *Client, f map[string]GuestPolicyPackageRepositoriesYum) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesYum(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackageRepositoriesYumSlice expands the contents of GuestPolicyPackageRepositoriesYum into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesYumSlice(c *Client, f []GuestPolicyPackageRepositoriesYum) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesYum(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackageRepositoriesYumMap flattens the contents of GuestPolicyPackageRepositoriesYum from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesYumMap(c *Client, i interface{}) map[string]GuestPolicyPackageRepositoriesYum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackageRepositoriesYum{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackageRepositoriesYum{}
	}

	items := make(map[string]GuestPolicyPackageRepositoriesYum)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackageRepositoriesYum(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesYumSlice flattens the contents of GuestPolicyPackageRepositoriesYum from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesYumSlice(c *Client, i interface{}) []GuestPolicyPackageRepositoriesYum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositoriesYum{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositoriesYum{}
	}

	items := make([]GuestPolicyPackageRepositoriesYum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositoriesYum(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackageRepositoriesYum expands an instance of GuestPolicyPackageRepositoriesYum into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesYum(c *Client, f *GuestPolicyPackageRepositoriesYum) (map[string]interface{}, error) {
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

// flattenGuestPolicyPackageRepositoriesYum flattens an instance of GuestPolicyPackageRepositoriesYum from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesYum(c *Client, i interface{}) *GuestPolicyPackageRepositoriesYum {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackageRepositoriesYum{}
	r.Id = dcl.FlattenString(m["id"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.BaseUrl = dcl.FlattenString(m["baseUrl"])
	r.GpgKeys = dcl.FlattenStringSlice(m["gpgKeys"])

	return r
}

// expandGuestPolicyPackageRepositoriesZypperMap expands the contents of GuestPolicyPackageRepositoriesZypper into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesZypperMap(c *Client, f map[string]GuestPolicyPackageRepositoriesZypper) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesZypper(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackageRepositoriesZypperSlice expands the contents of GuestPolicyPackageRepositoriesZypper into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesZypperSlice(c *Client, f []GuestPolicyPackageRepositoriesZypper) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesZypper(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackageRepositoriesZypperMap flattens the contents of GuestPolicyPackageRepositoriesZypper from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesZypperMap(c *Client, i interface{}) map[string]GuestPolicyPackageRepositoriesZypper {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackageRepositoriesZypper{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackageRepositoriesZypper{}
	}

	items := make(map[string]GuestPolicyPackageRepositoriesZypper)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackageRepositoriesZypper(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesZypperSlice flattens the contents of GuestPolicyPackageRepositoriesZypper from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesZypperSlice(c *Client, i interface{}) []GuestPolicyPackageRepositoriesZypper {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositoriesZypper{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositoriesZypper{}
	}

	items := make([]GuestPolicyPackageRepositoriesZypper, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositoriesZypper(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackageRepositoriesZypper expands an instance of GuestPolicyPackageRepositoriesZypper into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesZypper(c *Client, f *GuestPolicyPackageRepositoriesZypper) (map[string]interface{}, error) {
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

// flattenGuestPolicyPackageRepositoriesZypper flattens an instance of GuestPolicyPackageRepositoriesZypper from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesZypper(c *Client, i interface{}) *GuestPolicyPackageRepositoriesZypper {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackageRepositoriesZypper{}
	r.Id = dcl.FlattenString(m["id"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.BaseUrl = dcl.FlattenString(m["baseUrl"])
	r.GpgKeys = dcl.FlattenStringSlice(m["gpgKeys"])

	return r
}

// expandGuestPolicyPackageRepositoriesGooMap expands the contents of GuestPolicyPackageRepositoriesGoo into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesGooMap(c *Client, f map[string]GuestPolicyPackageRepositoriesGoo) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesGoo(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyPackageRepositoriesGooSlice expands the contents of GuestPolicyPackageRepositoriesGoo into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesGooSlice(c *Client, f []GuestPolicyPackageRepositoriesGoo) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyPackageRepositoriesGoo(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyPackageRepositoriesGooMap flattens the contents of GuestPolicyPackageRepositoriesGoo from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesGooMap(c *Client, i interface{}) map[string]GuestPolicyPackageRepositoriesGoo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyPackageRepositoriesGoo{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyPackageRepositoriesGoo{}
	}

	items := make(map[string]GuestPolicyPackageRepositoriesGoo)
	for k, item := range a {
		items[k] = *flattenGuestPolicyPackageRepositoriesGoo(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesGooSlice flattens the contents of GuestPolicyPackageRepositoriesGoo from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesGooSlice(c *Client, i interface{}) []GuestPolicyPackageRepositoriesGoo {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositoriesGoo{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositoriesGoo{}
	}

	items := make([]GuestPolicyPackageRepositoriesGoo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositoriesGoo(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyPackageRepositoriesGoo expands an instance of GuestPolicyPackageRepositoriesGoo into a JSON
// request object.
func expandGuestPolicyPackageRepositoriesGoo(c *Client, f *GuestPolicyPackageRepositoriesGoo) (map[string]interface{}, error) {
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

// flattenGuestPolicyPackageRepositoriesGoo flattens an instance of GuestPolicyPackageRepositoriesGoo from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesGoo(c *Client, i interface{}) *GuestPolicyPackageRepositoriesGoo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyPackageRepositoriesGoo{}
	r.Name = dcl.FlattenString(m["name"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandGuestPolicyRecipesMap expands the contents of GuestPolicyRecipes into a JSON
// request object.
func expandGuestPolicyRecipesMap(c *Client, f map[string]GuestPolicyRecipes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesSlice expands the contents of GuestPolicyRecipes into a JSON
// request object.
func expandGuestPolicyRecipesSlice(c *Client, f []GuestPolicyRecipes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesMap flattens the contents of GuestPolicyRecipes from a JSON
// response object.
func flattenGuestPolicyRecipesMap(c *Client, i interface{}) map[string]GuestPolicyRecipes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipes{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipes{}
	}

	items := make(map[string]GuestPolicyRecipes)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesSlice flattens the contents of GuestPolicyRecipes from a JSON
// response object.
func flattenGuestPolicyRecipesSlice(c *Client, i interface{}) []GuestPolicyRecipes {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipes{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipes{}
	}

	items := make([]GuestPolicyRecipes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipes(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipes expands an instance of GuestPolicyRecipes into a JSON
// request object.
func expandGuestPolicyRecipes(c *Client, f *GuestPolicyRecipes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v, err := expandGuestPolicyRecipesArtifactsSlice(c, f.Artifacts); err != nil {
		return nil, fmt.Errorf("error expanding Artifacts into artifacts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["artifacts"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsSlice(c, f.InstallSteps); err != nil {
		return nil, fmt.Errorf("error expanding InstallSteps into installSteps: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["installSteps"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsSlice(c, f.UpdateSteps); err != nil {
		return nil, fmt.Errorf("error expanding UpdateSteps into updateSteps: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["updateSteps"] = v
	}
	if v := f.DesiredState; !dcl.IsEmptyValueIndirect(v) {
		m["desiredState"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipes flattens an instance of GuestPolicyRecipes from a JSON
// response object.
func flattenGuestPolicyRecipes(c *Client, i interface{}) *GuestPolicyRecipes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipes{}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])
	r.Artifacts = flattenGuestPolicyRecipesArtifactsSlice(c, m["artifacts"])
	r.InstallSteps = flattenGuestPolicyRecipesInstallStepsSlice(c, m["installSteps"])
	r.UpdateSteps = flattenGuestPolicyRecipesUpdateStepsSlice(c, m["updateSteps"])
	r.DesiredState = flattenGuestPolicyRecipesDesiredStateEnum(m["desiredState"])

	return r
}

// expandGuestPolicyRecipesArtifactsMap expands the contents of GuestPolicyRecipesArtifacts into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsMap(c *Client, f map[string]GuestPolicyRecipesArtifacts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesArtifacts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesArtifactsSlice expands the contents of GuestPolicyRecipesArtifacts into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsSlice(c *Client, f []GuestPolicyRecipesArtifacts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesArtifacts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesArtifactsMap flattens the contents of GuestPolicyRecipesArtifacts from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsMap(c *Client, i interface{}) map[string]GuestPolicyRecipesArtifacts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesArtifacts{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesArtifacts{}
	}

	items := make(map[string]GuestPolicyRecipesArtifacts)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesArtifacts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesArtifactsSlice flattens the contents of GuestPolicyRecipesArtifacts from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsSlice(c *Client, i interface{}) []GuestPolicyRecipesArtifacts {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesArtifacts{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesArtifacts{}
	}

	items := make([]GuestPolicyRecipesArtifacts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesArtifacts(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesArtifacts expands an instance of GuestPolicyRecipesArtifacts into a JSON
// request object.
func expandGuestPolicyRecipesArtifacts(c *Client, f *GuestPolicyRecipesArtifacts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandGuestPolicyRecipesArtifactsRemote(c, f.Remote); err != nil {
		return nil, fmt.Errorf("error expanding Remote into remote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["remote"] = v
	}
	if v, err := expandGuestPolicyRecipesArtifactsGcs(c, f.Gcs); err != nil {
		return nil, fmt.Errorf("error expanding Gcs into gcs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gcs"] = v
	}
	if v := f.AllowInsecure; !dcl.IsEmptyValueIndirect(v) {
		m["allowInsecure"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesArtifacts flattens an instance of GuestPolicyRecipesArtifacts from a JSON
// response object.
func flattenGuestPolicyRecipesArtifacts(c *Client, i interface{}) *GuestPolicyRecipesArtifacts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesArtifacts{}
	r.Id = dcl.FlattenString(m["id"])
	r.Remote = flattenGuestPolicyRecipesArtifactsRemote(c, m["remote"])
	r.Gcs = flattenGuestPolicyRecipesArtifactsGcs(c, m["gcs"])
	r.AllowInsecure = dcl.FlattenBool(m["allowInsecure"])

	return r
}

// expandGuestPolicyRecipesArtifactsRemoteMap expands the contents of GuestPolicyRecipesArtifactsRemote into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsRemoteMap(c *Client, f map[string]GuestPolicyRecipesArtifactsRemote) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesArtifactsRemote(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesArtifactsRemoteSlice expands the contents of GuestPolicyRecipesArtifactsRemote into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsRemoteSlice(c *Client, f []GuestPolicyRecipesArtifactsRemote) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesArtifactsRemote(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesArtifactsRemoteMap flattens the contents of GuestPolicyRecipesArtifactsRemote from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsRemoteMap(c *Client, i interface{}) map[string]GuestPolicyRecipesArtifactsRemote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesArtifactsRemote{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesArtifactsRemote{}
	}

	items := make(map[string]GuestPolicyRecipesArtifactsRemote)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesArtifactsRemote(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesArtifactsRemoteSlice flattens the contents of GuestPolicyRecipesArtifactsRemote from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsRemoteSlice(c *Client, i interface{}) []GuestPolicyRecipesArtifactsRemote {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesArtifactsRemote{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesArtifactsRemote{}
	}

	items := make([]GuestPolicyRecipesArtifactsRemote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesArtifactsRemote(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesArtifactsRemote expands an instance of GuestPolicyRecipesArtifactsRemote into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsRemote(c *Client, f *GuestPolicyRecipesArtifactsRemote) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.Checksum; !dcl.IsEmptyValueIndirect(v) {
		m["checksum"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesArtifactsRemote flattens an instance of GuestPolicyRecipesArtifactsRemote from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsRemote(c *Client, i interface{}) *GuestPolicyRecipesArtifactsRemote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesArtifactsRemote{}
	r.Uri = dcl.FlattenString(m["uri"])
	r.Checksum = dcl.FlattenString(m["checksum"])

	return r
}

// expandGuestPolicyRecipesArtifactsGcsMap expands the contents of GuestPolicyRecipesArtifactsGcs into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsGcsMap(c *Client, f map[string]GuestPolicyRecipesArtifactsGcs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesArtifactsGcs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesArtifactsGcsSlice expands the contents of GuestPolicyRecipesArtifactsGcs into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsGcsSlice(c *Client, f []GuestPolicyRecipesArtifactsGcs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesArtifactsGcs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesArtifactsGcsMap flattens the contents of GuestPolicyRecipesArtifactsGcs from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsGcsMap(c *Client, i interface{}) map[string]GuestPolicyRecipesArtifactsGcs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesArtifactsGcs{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesArtifactsGcs{}
	}

	items := make(map[string]GuestPolicyRecipesArtifactsGcs)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesArtifactsGcs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesArtifactsGcsSlice flattens the contents of GuestPolicyRecipesArtifactsGcs from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsGcsSlice(c *Client, i interface{}) []GuestPolicyRecipesArtifactsGcs {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesArtifactsGcs{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesArtifactsGcs{}
	}

	items := make([]GuestPolicyRecipesArtifactsGcs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesArtifactsGcs(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesArtifactsGcs expands an instance of GuestPolicyRecipesArtifactsGcs into a JSON
// request object.
func expandGuestPolicyRecipesArtifactsGcs(c *Client, f *GuestPolicyRecipesArtifactsGcs) (map[string]interface{}, error) {
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

// flattenGuestPolicyRecipesArtifactsGcs flattens an instance of GuestPolicyRecipesArtifactsGcs from a JSON
// response object.
func flattenGuestPolicyRecipesArtifactsGcs(c *Client, i interface{}) *GuestPolicyRecipesArtifactsGcs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesArtifactsGcs{}
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Object = dcl.FlattenString(m["object"])
	r.Generation = dcl.FlattenInteger(m["generation"])

	return r
}

// expandGuestPolicyRecipesInstallStepsMap expands the contents of GuestPolicyRecipesInstallSteps into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsMap(c *Client, f map[string]GuestPolicyRecipesInstallSteps) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallSteps(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsSlice expands the contents of GuestPolicyRecipesInstallSteps into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsSlice(c *Client, f []GuestPolicyRecipesInstallSteps) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallSteps(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsMap flattens the contents of GuestPolicyRecipesInstallSteps from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallSteps {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallSteps{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallSteps{}
	}

	items := make(map[string]GuestPolicyRecipesInstallSteps)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallSteps(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsSlice flattens the contents of GuestPolicyRecipesInstallSteps from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallSteps {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallSteps{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallSteps{}
	}

	items := make([]GuestPolicyRecipesInstallSteps, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallSteps(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallSteps expands an instance of GuestPolicyRecipesInstallSteps into a JSON
// request object.
func expandGuestPolicyRecipesInstallSteps(c *Client, f *GuestPolicyRecipesInstallSteps) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGuestPolicyRecipesInstallStepsFileCopy(c, f.FileCopy); err != nil {
		return nil, fmt.Errorf("error expanding FileCopy into fileCopy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fileCopy"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsArchiveExtraction(c, f.ArchiveExtraction); err != nil {
		return nil, fmt.Errorf("error expanding ArchiveExtraction into archiveExtraction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["archiveExtraction"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsMsiInstallation(c, f.MsiInstallation); err != nil {
		return nil, fmt.Errorf("error expanding MsiInstallation into msiInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["msiInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsDpkgInstallation(c, f.DpkgInstallation); err != nil {
		return nil, fmt.Errorf("error expanding DpkgInstallation into dpkgInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dpkgInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsRpmInstallation(c, f.RpmInstallation); err != nil {
		return nil, fmt.Errorf("error expanding RpmInstallation into rpmInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rpmInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsFileExec(c, f.FileExec); err != nil {
		return nil, fmt.Errorf("error expanding FileExec into fileExec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fileExec"] = v
	}
	if v, err := expandGuestPolicyRecipesInstallStepsScriptRun(c, f.ScriptRun); err != nil {
		return nil, fmt.Errorf("error expanding ScriptRun into scriptRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scriptRun"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallSteps flattens an instance of GuestPolicyRecipesInstallSteps from a JSON
// response object.
func flattenGuestPolicyRecipesInstallSteps(c *Client, i interface{}) *GuestPolicyRecipesInstallSteps {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallSteps{}
	r.FileCopy = flattenGuestPolicyRecipesInstallStepsFileCopy(c, m["fileCopy"])
	r.ArchiveExtraction = flattenGuestPolicyRecipesInstallStepsArchiveExtraction(c, m["archiveExtraction"])
	r.MsiInstallation = flattenGuestPolicyRecipesInstallStepsMsiInstallation(c, m["msiInstallation"])
	r.DpkgInstallation = flattenGuestPolicyRecipesInstallStepsDpkgInstallation(c, m["dpkgInstallation"])
	r.RpmInstallation = flattenGuestPolicyRecipesInstallStepsRpmInstallation(c, m["rpmInstallation"])
	r.FileExec = flattenGuestPolicyRecipesInstallStepsFileExec(c, m["fileExec"])
	r.ScriptRun = flattenGuestPolicyRecipesInstallStepsScriptRun(c, m["scriptRun"])

	return r
}

// expandGuestPolicyRecipesInstallStepsFileCopyMap expands the contents of GuestPolicyRecipesInstallStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileCopyMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsFileCopy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsFileCopy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsFileCopySlice expands the contents of GuestPolicyRecipesInstallStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileCopySlice(c *Client, f []GuestPolicyRecipesInstallStepsFileCopy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsFileCopy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsFileCopyMap flattens the contents of GuestPolicyRecipesInstallStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileCopyMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsFileCopy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsFileCopy{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsFileCopy{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsFileCopy)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsFileCopy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsFileCopySlice flattens the contents of GuestPolicyRecipesInstallStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileCopySlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsFileCopy {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsFileCopy{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsFileCopy{}
	}

	items := make([]GuestPolicyRecipesInstallStepsFileCopy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsFileCopy(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsFileCopy expands an instance of GuestPolicyRecipesInstallStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileCopy(c *Client, f *GuestPolicyRecipesInstallStepsFileCopy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Destination; !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v := f.Overwrite; !dcl.IsEmptyValueIndirect(v) {
		m["overwrite"] = v
	}
	if v := f.Permissions; !dcl.IsEmptyValueIndirect(v) {
		m["permissions"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsFileCopy flattens an instance of GuestPolicyRecipesInstallStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileCopy(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsFileCopy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsFileCopy{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Destination = dcl.FlattenString(m["destination"])
	r.Overwrite = dcl.FlattenBool(m["overwrite"])
	r.Permissions = dcl.FlattenString(m["permissions"])

	return r
}

// expandGuestPolicyRecipesInstallStepsArchiveExtractionMap expands the contents of GuestPolicyRecipesInstallStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsArchiveExtractionMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsArchiveExtraction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsArchiveExtraction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsArchiveExtractionSlice expands the contents of GuestPolicyRecipesInstallStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsArchiveExtractionSlice(c *Client, f []GuestPolicyRecipesInstallStepsArchiveExtraction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsArchiveExtraction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsArchiveExtractionMap flattens the contents of GuestPolicyRecipesInstallStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsArchiveExtractionMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsArchiveExtraction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsArchiveExtraction{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsArchiveExtraction{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsArchiveExtraction)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsArchiveExtraction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsArchiveExtractionSlice flattens the contents of GuestPolicyRecipesInstallStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsArchiveExtractionSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsArchiveExtraction {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsArchiveExtraction{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsArchiveExtraction{}
	}

	items := make([]GuestPolicyRecipesInstallStepsArchiveExtraction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsArchiveExtraction(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsArchiveExtraction expands an instance of GuestPolicyRecipesInstallStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsArchiveExtraction(c *Client, f *GuestPolicyRecipesInstallStepsArchiveExtraction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Destination; !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsArchiveExtraction flattens an instance of GuestPolicyRecipesInstallStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsArchiveExtraction(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsArchiveExtraction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsArchiveExtraction{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Destination = dcl.FlattenString(m["destination"])
	r.Type = flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(m["type"])

	return r
}

// expandGuestPolicyRecipesInstallStepsMsiInstallationMap expands the contents of GuestPolicyRecipesInstallStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsMsiInstallationMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsMsiInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsMsiInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsMsiInstallationSlice expands the contents of GuestPolicyRecipesInstallStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsMsiInstallationSlice(c *Client, f []GuestPolicyRecipesInstallStepsMsiInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsMsiInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsMsiInstallationMap flattens the contents of GuestPolicyRecipesInstallStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsMsiInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsMsiInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsMsiInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsMsiInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsMsiInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsMsiInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsMsiInstallationSlice flattens the contents of GuestPolicyRecipesInstallStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsMsiInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsMsiInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsMsiInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsMsiInstallation{}
	}

	items := make([]GuestPolicyRecipesInstallStepsMsiInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsMsiInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsMsiInstallation expands an instance of GuestPolicyRecipesInstallStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsMsiInstallation(c *Client, f *GuestPolicyRecipesInstallStepsMsiInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Flags; !dcl.IsEmptyValueIndirect(v) {
		m["flags"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsMsiInstallation flattens an instance of GuestPolicyRecipesInstallStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsMsiInstallation(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsMsiInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsMsiInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Flags = dcl.FlattenStringSlice(m["flags"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])

	return r
}

// expandGuestPolicyRecipesInstallStepsDpkgInstallationMap expands the contents of GuestPolicyRecipesInstallStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsDpkgInstallationMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsDpkgInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsDpkgInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsDpkgInstallationSlice expands the contents of GuestPolicyRecipesInstallStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsDpkgInstallationSlice(c *Client, f []GuestPolicyRecipesInstallStepsDpkgInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsDpkgInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsDpkgInstallationMap flattens the contents of GuestPolicyRecipesInstallStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsDpkgInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsDpkgInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsDpkgInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsDpkgInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsDpkgInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsDpkgInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsDpkgInstallationSlice flattens the contents of GuestPolicyRecipesInstallStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsDpkgInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsDpkgInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsDpkgInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsDpkgInstallation{}
	}

	items := make([]GuestPolicyRecipesInstallStepsDpkgInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsDpkgInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsDpkgInstallation expands an instance of GuestPolicyRecipesInstallStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsDpkgInstallation(c *Client, f *GuestPolicyRecipesInstallStepsDpkgInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsDpkgInstallation flattens an instance of GuestPolicyRecipesInstallStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsDpkgInstallation(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsDpkgInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsDpkgInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])

	return r
}

// expandGuestPolicyRecipesInstallStepsRpmInstallationMap expands the contents of GuestPolicyRecipesInstallStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsRpmInstallationMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsRpmInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsRpmInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsRpmInstallationSlice expands the contents of GuestPolicyRecipesInstallStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsRpmInstallationSlice(c *Client, f []GuestPolicyRecipesInstallStepsRpmInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsRpmInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsRpmInstallationMap flattens the contents of GuestPolicyRecipesInstallStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsRpmInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsRpmInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsRpmInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsRpmInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsRpmInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsRpmInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsRpmInstallationSlice flattens the contents of GuestPolicyRecipesInstallStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsRpmInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsRpmInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsRpmInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsRpmInstallation{}
	}

	items := make([]GuestPolicyRecipesInstallStepsRpmInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsRpmInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsRpmInstallation expands an instance of GuestPolicyRecipesInstallStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsRpmInstallation(c *Client, f *GuestPolicyRecipesInstallStepsRpmInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsRpmInstallation flattens an instance of GuestPolicyRecipesInstallStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsRpmInstallation(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsRpmInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsRpmInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])

	return r
}

// expandGuestPolicyRecipesInstallStepsFileExecMap expands the contents of GuestPolicyRecipesInstallStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileExecMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsFileExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsFileExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsFileExecSlice expands the contents of GuestPolicyRecipesInstallStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileExecSlice(c *Client, f []GuestPolicyRecipesInstallStepsFileExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsFileExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsFileExecMap flattens the contents of GuestPolicyRecipesInstallStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileExecMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsFileExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsFileExec{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsFileExec{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsFileExec)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsFileExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsFileExecSlice flattens the contents of GuestPolicyRecipesInstallStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileExecSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsFileExec {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsFileExec{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsFileExec{}
	}

	items := make([]GuestPolicyRecipesInstallStepsFileExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsFileExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsFileExec expands an instance of GuestPolicyRecipesInstallStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsFileExec(c *Client, f *GuestPolicyRecipesInstallStepsFileExec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.LocalPath; !dcl.IsEmptyValueIndirect(v) {
		m["localPath"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsFileExec flattens an instance of GuestPolicyRecipesInstallStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsFileExec(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsFileExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsFileExec{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.LocalPath = dcl.FlattenString(m["localPath"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])

	return r
}

// expandGuestPolicyRecipesInstallStepsScriptRunMap expands the contents of GuestPolicyRecipesInstallStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsScriptRunMap(c *Client, f map[string]GuestPolicyRecipesInstallStepsScriptRun) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsScriptRun(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesInstallStepsScriptRunSlice expands the contents of GuestPolicyRecipesInstallStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsScriptRunSlice(c *Client, f []GuestPolicyRecipesInstallStepsScriptRun) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesInstallStepsScriptRun(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesInstallStepsScriptRunMap flattens the contents of GuestPolicyRecipesInstallStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsScriptRunMap(c *Client, i interface{}) map[string]GuestPolicyRecipesInstallStepsScriptRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesInstallStepsScriptRun{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesInstallStepsScriptRun{}
	}

	items := make(map[string]GuestPolicyRecipesInstallStepsScriptRun)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesInstallStepsScriptRun(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsScriptRunSlice flattens the contents of GuestPolicyRecipesInstallStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsScriptRunSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsScriptRun {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsScriptRun{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsScriptRun{}
	}

	items := make([]GuestPolicyRecipesInstallStepsScriptRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsScriptRun(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesInstallStepsScriptRun expands an instance of GuestPolicyRecipesInstallStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesInstallStepsScriptRun(c *Client, f *GuestPolicyRecipesInstallStepsScriptRun) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Script; !dcl.IsEmptyValueIndirect(v) {
		m["script"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}
	if v := f.Interpreter; !dcl.IsEmptyValueIndirect(v) {
		m["interpreter"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesInstallStepsScriptRun flattens an instance of GuestPolicyRecipesInstallStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsScriptRun(c *Client, i interface{}) *GuestPolicyRecipesInstallStepsScriptRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesInstallStepsScriptRun{}
	r.Script = dcl.FlattenString(m["script"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])
	r.Interpreter = flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(m["interpreter"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsMap expands the contents of GuestPolicyRecipesUpdateSteps into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsMap(c *Client, f map[string]GuestPolicyRecipesUpdateSteps) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateSteps(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsSlice expands the contents of GuestPolicyRecipesUpdateSteps into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsSlice(c *Client, f []GuestPolicyRecipesUpdateSteps) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateSteps(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsMap flattens the contents of GuestPolicyRecipesUpdateSteps from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateSteps {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateSteps{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateSteps{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateSteps)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateSteps(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsSlice flattens the contents of GuestPolicyRecipesUpdateSteps from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateSteps {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateSteps{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateSteps{}
	}

	items := make([]GuestPolicyRecipesUpdateSteps, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateSteps(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateSteps expands an instance of GuestPolicyRecipesUpdateSteps into a JSON
// request object.
func expandGuestPolicyRecipesUpdateSteps(c *Client, f *GuestPolicyRecipesUpdateSteps) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGuestPolicyRecipesUpdateStepsFileCopy(c, f.FileCopy); err != nil {
		return nil, fmt.Errorf("error expanding FileCopy into fileCopy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fileCopy"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsArchiveExtraction(c, f.ArchiveExtraction); err != nil {
		return nil, fmt.Errorf("error expanding ArchiveExtraction into archiveExtraction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["archiveExtraction"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsMsiInstallation(c, f.MsiInstallation); err != nil {
		return nil, fmt.Errorf("error expanding MsiInstallation into msiInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["msiInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsDpkgInstallation(c, f.DpkgInstallation); err != nil {
		return nil, fmt.Errorf("error expanding DpkgInstallation into dpkgInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dpkgInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsRpmInstallation(c, f.RpmInstallation); err != nil {
		return nil, fmt.Errorf("error expanding RpmInstallation into rpmInstallation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rpmInstallation"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsFileExec(c, f.FileExec); err != nil {
		return nil, fmt.Errorf("error expanding FileExec into fileExec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fileExec"] = v
	}
	if v, err := expandGuestPolicyRecipesUpdateStepsScriptRun(c, f.ScriptRun); err != nil {
		return nil, fmt.Errorf("error expanding ScriptRun into scriptRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scriptRun"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateSteps flattens an instance of GuestPolicyRecipesUpdateSteps from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateSteps(c *Client, i interface{}) *GuestPolicyRecipesUpdateSteps {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateSteps{}
	r.FileCopy = flattenGuestPolicyRecipesUpdateStepsFileCopy(c, m["fileCopy"])
	r.ArchiveExtraction = flattenGuestPolicyRecipesUpdateStepsArchiveExtraction(c, m["archiveExtraction"])
	r.MsiInstallation = flattenGuestPolicyRecipesUpdateStepsMsiInstallation(c, m["msiInstallation"])
	r.DpkgInstallation = flattenGuestPolicyRecipesUpdateStepsDpkgInstallation(c, m["dpkgInstallation"])
	r.RpmInstallation = flattenGuestPolicyRecipesUpdateStepsRpmInstallation(c, m["rpmInstallation"])
	r.FileExec = flattenGuestPolicyRecipesUpdateStepsFileExec(c, m["fileExec"])
	r.ScriptRun = flattenGuestPolicyRecipesUpdateStepsScriptRun(c, m["scriptRun"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsFileCopyMap expands the contents of GuestPolicyRecipesUpdateStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileCopyMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsFileCopy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsFileCopy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsFileCopySlice expands the contents of GuestPolicyRecipesUpdateStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileCopySlice(c *Client, f []GuestPolicyRecipesUpdateStepsFileCopy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsFileCopy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsFileCopyMap flattens the contents of GuestPolicyRecipesUpdateStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileCopyMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsFileCopy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsFileCopy{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsFileCopy{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsFileCopy)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsFileCopy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsFileCopySlice flattens the contents of GuestPolicyRecipesUpdateStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileCopySlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsFileCopy {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsFileCopy{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsFileCopy{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsFileCopy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsFileCopy(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsFileCopy expands an instance of GuestPolicyRecipesUpdateStepsFileCopy into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileCopy(c *Client, f *GuestPolicyRecipesUpdateStepsFileCopy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Destination; !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v := f.Overwrite; !dcl.IsEmptyValueIndirect(v) {
		m["overwrite"] = v
	}
	if v := f.Permissions; !dcl.IsEmptyValueIndirect(v) {
		m["permissions"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsFileCopy flattens an instance of GuestPolicyRecipesUpdateStepsFileCopy from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileCopy(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsFileCopy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsFileCopy{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Destination = dcl.FlattenString(m["destination"])
	r.Overwrite = dcl.FlattenBool(m["overwrite"])
	r.Permissions = dcl.FlattenString(m["permissions"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsArchiveExtractionMap expands the contents of GuestPolicyRecipesUpdateStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsArchiveExtractionMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsArchiveExtractionSlice expands the contents of GuestPolicyRecipesUpdateStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsArchiveExtractionSlice(c *Client, f []GuestPolicyRecipesUpdateStepsArchiveExtraction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsArchiveExtraction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsArchiveExtractionMap flattens the contents of GuestPolicyRecipesUpdateStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsArchiveExtractionMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsArchiveExtraction)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsArchiveExtraction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsArchiveExtractionSlice flattens the contents of GuestPolicyRecipesUpdateStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsArchiveExtractionSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsArchiveExtraction {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsArchiveExtraction{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsArchiveExtraction{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsArchiveExtraction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsArchiveExtraction(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsArchiveExtraction expands an instance of GuestPolicyRecipesUpdateStepsArchiveExtraction into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsArchiveExtraction(c *Client, f *GuestPolicyRecipesUpdateStepsArchiveExtraction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Destination; !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsArchiveExtraction flattens an instance of GuestPolicyRecipesUpdateStepsArchiveExtraction from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsArchiveExtraction(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsArchiveExtraction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsArchiveExtraction{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Destination = dcl.FlattenString(m["destination"])
	r.Type = flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(m["type"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsMsiInstallationMap expands the contents of GuestPolicyRecipesUpdateStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsMsiInstallationMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsMsiInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsMsiInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsMsiInstallationSlice expands the contents of GuestPolicyRecipesUpdateStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsMsiInstallationSlice(c *Client, f []GuestPolicyRecipesUpdateStepsMsiInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsMsiInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsMsiInstallationMap flattens the contents of GuestPolicyRecipesUpdateStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsMsiInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsMsiInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsMsiInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsMsiInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsMsiInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsMsiInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsMsiInstallationSlice flattens the contents of GuestPolicyRecipesUpdateStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsMsiInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsMsiInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsMsiInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsMsiInstallation{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsMsiInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsMsiInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsMsiInstallation expands an instance of GuestPolicyRecipesUpdateStepsMsiInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsMsiInstallation(c *Client, f *GuestPolicyRecipesUpdateStepsMsiInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.Flags; !dcl.IsEmptyValueIndirect(v) {
		m["flags"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsMsiInstallation flattens an instance of GuestPolicyRecipesUpdateStepsMsiInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsMsiInstallation(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsMsiInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsMsiInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.Flags = dcl.FlattenStringSlice(m["flags"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsDpkgInstallationMap expands the contents of GuestPolicyRecipesUpdateStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsDpkgInstallationMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsDpkgInstallationSlice expands the contents of GuestPolicyRecipesUpdateStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsDpkgInstallationSlice(c *Client, f []GuestPolicyRecipesUpdateStepsDpkgInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsDpkgInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsDpkgInstallationMap flattens the contents of GuestPolicyRecipesUpdateStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsDpkgInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsDpkgInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsDpkgInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsDpkgInstallationSlice flattens the contents of GuestPolicyRecipesUpdateStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsDpkgInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsDpkgInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsDpkgInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsDpkgInstallation{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsDpkgInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsDpkgInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsDpkgInstallation expands an instance of GuestPolicyRecipesUpdateStepsDpkgInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsDpkgInstallation(c *Client, f *GuestPolicyRecipesUpdateStepsDpkgInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsDpkgInstallation flattens an instance of GuestPolicyRecipesUpdateStepsDpkgInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsDpkgInstallation(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsDpkgInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsDpkgInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsRpmInstallationMap expands the contents of GuestPolicyRecipesUpdateStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsRpmInstallationMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsRpmInstallation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsRpmInstallation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsRpmInstallationSlice expands the contents of GuestPolicyRecipesUpdateStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsRpmInstallationSlice(c *Client, f []GuestPolicyRecipesUpdateStepsRpmInstallation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsRpmInstallation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsRpmInstallationMap flattens the contents of GuestPolicyRecipesUpdateStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsRpmInstallationMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsRpmInstallation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsRpmInstallation{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsRpmInstallation{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsRpmInstallation)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsRpmInstallation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsRpmInstallationSlice flattens the contents of GuestPolicyRecipesUpdateStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsRpmInstallationSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsRpmInstallation {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsRpmInstallation{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsRpmInstallation{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsRpmInstallation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsRpmInstallation(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsRpmInstallation expands an instance of GuestPolicyRecipesUpdateStepsRpmInstallation into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsRpmInstallation(c *Client, f *GuestPolicyRecipesUpdateStepsRpmInstallation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsRpmInstallation flattens an instance of GuestPolicyRecipesUpdateStepsRpmInstallation from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsRpmInstallation(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsRpmInstallation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsRpmInstallation{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsFileExecMap expands the contents of GuestPolicyRecipesUpdateStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileExecMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsFileExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsFileExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsFileExecSlice expands the contents of GuestPolicyRecipesUpdateStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileExecSlice(c *Client, f []GuestPolicyRecipesUpdateStepsFileExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsFileExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsFileExecMap flattens the contents of GuestPolicyRecipesUpdateStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileExecMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsFileExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsFileExec{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsFileExec{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsFileExec)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsFileExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsFileExecSlice flattens the contents of GuestPolicyRecipesUpdateStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileExecSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsFileExec {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsFileExec{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsFileExec{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsFileExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsFileExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsFileExec expands an instance of GuestPolicyRecipesUpdateStepsFileExec into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsFileExec(c *Client, f *GuestPolicyRecipesUpdateStepsFileExec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ArtifactId; !dcl.IsEmptyValueIndirect(v) {
		m["artifactId"] = v
	}
	if v := f.LocalPath; !dcl.IsEmptyValueIndirect(v) {
		m["localPath"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsFileExec flattens an instance of GuestPolicyRecipesUpdateStepsFileExec from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsFileExec(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsFileExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsFileExec{}
	r.ArtifactId = dcl.FlattenString(m["artifactId"])
	r.LocalPath = dcl.FlattenString(m["localPath"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])

	return r
}

// expandGuestPolicyRecipesUpdateStepsScriptRunMap expands the contents of GuestPolicyRecipesUpdateStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsScriptRunMap(c *Client, f map[string]GuestPolicyRecipesUpdateStepsScriptRun) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsScriptRun(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGuestPolicyRecipesUpdateStepsScriptRunSlice expands the contents of GuestPolicyRecipesUpdateStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsScriptRunSlice(c *Client, f []GuestPolicyRecipesUpdateStepsScriptRun) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGuestPolicyRecipesUpdateStepsScriptRun(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGuestPolicyRecipesUpdateStepsScriptRunMap flattens the contents of GuestPolicyRecipesUpdateStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsScriptRunMap(c *Client, i interface{}) map[string]GuestPolicyRecipesUpdateStepsScriptRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GuestPolicyRecipesUpdateStepsScriptRun{}
	}

	if len(a) == 0 {
		return map[string]GuestPolicyRecipesUpdateStepsScriptRun{}
	}

	items := make(map[string]GuestPolicyRecipesUpdateStepsScriptRun)
	for k, item := range a {
		items[k] = *flattenGuestPolicyRecipesUpdateStepsScriptRun(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsScriptRunSlice flattens the contents of GuestPolicyRecipesUpdateStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsScriptRunSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsScriptRun {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsScriptRun{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsScriptRun{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsScriptRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsScriptRun(c, item.(map[string]interface{})))
	}

	return items
}

// expandGuestPolicyRecipesUpdateStepsScriptRun expands an instance of GuestPolicyRecipesUpdateStepsScriptRun into a JSON
// request object.
func expandGuestPolicyRecipesUpdateStepsScriptRun(c *Client, f *GuestPolicyRecipesUpdateStepsScriptRun) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Script; !dcl.IsEmptyValueIndirect(v) {
		m["script"] = v
	}
	if v := f.AllowedExitCodes; !dcl.IsEmptyValueIndirect(v) {
		m["allowedExitCodes"] = v
	}
	if v := f.Interpreter; !dcl.IsEmptyValueIndirect(v) {
		m["interpreter"] = v
	}

	return m, nil
}

// flattenGuestPolicyRecipesUpdateStepsScriptRun flattens an instance of GuestPolicyRecipesUpdateStepsScriptRun from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsScriptRun(c *Client, i interface{}) *GuestPolicyRecipesUpdateStepsScriptRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GuestPolicyRecipesUpdateStepsScriptRun{}
	r.Script = dcl.FlattenString(m["script"])
	r.AllowedExitCodes = dcl.FlattenIntSlice(m["allowedExitCodes"])
	r.Interpreter = flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(m["interpreter"])

	return r
}

// flattenGuestPolicyPackagesDesiredStateEnumSlice flattens the contents of GuestPolicyPackagesDesiredStateEnum from a JSON
// response object.
func flattenGuestPolicyPackagesDesiredStateEnumSlice(c *Client, i interface{}) []GuestPolicyPackagesDesiredStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackagesDesiredStateEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackagesDesiredStateEnum{}
	}

	items := make([]GuestPolicyPackagesDesiredStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackagesDesiredStateEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyPackagesDesiredStateEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyPackagesDesiredStateEnum with the same value as that string.
func flattenGuestPolicyPackagesDesiredStateEnum(i interface{}) *GuestPolicyPackagesDesiredStateEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyPackagesDesiredStateEnumRef("")
	}

	return GuestPolicyPackagesDesiredStateEnumRef(s)
}

// flattenGuestPolicyPackagesManagerEnumSlice flattens the contents of GuestPolicyPackagesManagerEnum from a JSON
// response object.
func flattenGuestPolicyPackagesManagerEnumSlice(c *Client, i interface{}) []GuestPolicyPackagesManagerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackagesManagerEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackagesManagerEnum{}
	}

	items := make([]GuestPolicyPackagesManagerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackagesManagerEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyPackagesManagerEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyPackagesManagerEnum with the same value as that string.
func flattenGuestPolicyPackagesManagerEnum(i interface{}) *GuestPolicyPackagesManagerEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyPackagesManagerEnumRef("")
	}

	return GuestPolicyPackagesManagerEnumRef(s)
}

// flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnumSlice flattens the contents of GuestPolicyPackageRepositoriesAptArchiveTypeEnum from a JSON
// response object.
func flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnumSlice(c *Client, i interface{}) []GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyPackageRepositoriesAptArchiveTypeEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyPackageRepositoriesAptArchiveTypeEnum{}
	}

	items := make([]GuestPolicyPackageRepositoriesAptArchiveTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyPackageRepositoriesAptArchiveTypeEnum with the same value as that string.
func flattenGuestPolicyPackageRepositoriesAptArchiveTypeEnum(i interface{}) *GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef("")
	}

	return GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef(s)
}

// flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumSlice flattens the contents of GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum{}
	}

	items := make([]GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum with the same value as that string.
func flattenGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(i interface{}) *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef("")
	}

	return GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef(s)
}

// flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumSlice flattens the contents of GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum from a JSON
// response object.
func flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumSlice(c *Client, i interface{}) []GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum{}
	}

	items := make([]GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum with the same value as that string.
func flattenGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(i interface{}) *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef("")
	}

	return GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef(s)
}

// flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumSlice flattens the contents of GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum with the same value as that string.
func flattenGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(i interface{}) *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef("")
	}

	return GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef(s)
}

// flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumSlice flattens the contents of GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum from a JSON
// response object.
func flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumSlice(c *Client, i interface{}) []GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum{}
	}

	items := make([]GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum with the same value as that string.
func flattenGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(i interface{}) *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef("")
	}

	return GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef(s)
}

// flattenGuestPolicyRecipesDesiredStateEnumSlice flattens the contents of GuestPolicyRecipesDesiredStateEnum from a JSON
// response object.
func flattenGuestPolicyRecipesDesiredStateEnumSlice(c *Client, i interface{}) []GuestPolicyRecipesDesiredStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GuestPolicyRecipesDesiredStateEnum{}
	}

	if len(a) == 0 {
		return []GuestPolicyRecipesDesiredStateEnum{}
	}

	items := make([]GuestPolicyRecipesDesiredStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGuestPolicyRecipesDesiredStateEnum(item.(interface{})))
	}

	return items
}

// flattenGuestPolicyRecipesDesiredStateEnum asserts that an interface is a string, and returns a
// pointer to a *GuestPolicyRecipesDesiredStateEnum with the same value as that string.
func flattenGuestPolicyRecipesDesiredStateEnum(i interface{}) *GuestPolicyRecipesDesiredStateEnum {
	s, ok := i.(string)
	if !ok {
		return GuestPolicyRecipesDesiredStateEnumRef("")
	}

	return GuestPolicyRecipesDesiredStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *GuestPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGuestPolicy(b, c)
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
