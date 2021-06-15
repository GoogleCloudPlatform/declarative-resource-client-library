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
	Diffs        []*dcl.FieldDiff
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
		res, err := unmarshalMapGuestPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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

func (c *Client) guestPolicyDiffsForRawDesired(ctx context.Context, rawDesired *GuestPolicy, opts ...dcl.ApplyOption) (initial, desired *GuestPolicy, diffs []*dcl.FieldDiff, err error) {
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
	if dcl.IsZeroValue(nw.Zones) {
		nw.Zones = des.Zones
	}
	if dcl.IsZeroValue(nw.Instances) {
		nw.Instances = des.Instances
	}
	if dcl.IsZeroValue(nw.InstanceNamePrefixes) {
		nw.InstanceNamePrefixes = des.InstanceNamePrefixes
	}
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
			if diffs, _ := compareGuestPolicyAssignmentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
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
			if diffs, _ := compareGuestPolicyAssignmentGroupLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyAssignmentOsTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.DesiredState) {
		nw.DesiredState = des.DesiredState
	}
	if dcl.IsZeroValue(nw.Manager) {
		nw.Manager = des.Manager
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
			if diffs, _ := compareGuestPolicyPackagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyPackageRepositoriesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

func canonicalizeNewGuestPolicyPackageRepositoriesAptSet(c *Client, des, nw []GuestPolicyPackageRepositoriesApt) []GuestPolicyPackageRepositoriesApt {
	if des == nil {
		return nw
	}
	var reorderedNew []GuestPolicyPackageRepositoriesApt
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGuestPolicyPackageRepositoriesAptNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.GpgKeys) {
		nw.GpgKeys = des.GpgKeys
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
			if diffs, _ := compareGuestPolicyPackageRepositoriesYumNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.GpgKeys) {
		nw.GpgKeys = des.GpgKeys
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
			if diffs, _ := compareGuestPolicyPackageRepositoriesZypperNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyPackageRepositoriesGooNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.DesiredState) {
		nw.DesiredState = des.DesiredState
	}

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
			if diffs, _ := compareGuestPolicyRecipesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesArtifactsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesArtifactsRemoteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Generation) {
		nw.Generation = des.Generation
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
			if diffs, _ := compareGuestPolicyRecipesArtifactsGcsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsFileCopyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsArchiveExtractionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Flags) {
		nw.Flags = des.Flags
	}
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsMsiInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsDpkgInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsRpmInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Args) {
		nw.Args = des.Args
	}
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsFileExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
	}
	if dcl.IsZeroValue(nw.Interpreter) {
		nw.Interpreter = des.Interpreter
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
			if diffs, _ := compareGuestPolicyRecipesInstallStepsScriptRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsFileCopyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsArchiveExtractionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Flags) {
		nw.Flags = des.Flags
	}
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsMsiInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsDpkgInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsRpmInstallationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Args) {
		nw.Args = des.Args
	}
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsFileExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.AllowedExitCodes) {
		nw.AllowedExitCodes = des.AllowedExitCodes
	}
	if dcl.IsZeroValue(nw.Interpreter) {
		nw.Interpreter = des.Interpreter
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
			if diffs, _ := compareGuestPolicyRecipesUpdateStepsScriptRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []GuestPolicyRecipesUpdateStepsScriptRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGuestPolicyRecipesUpdateStepsScriptRun(c, &d, &n))
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
func diffGuestPolicy(c *Client, desired, actual *GuestPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Assignment, actual.Assignment, dcl.Info{ObjectFunction: compareGuestPolicyAssignmentNewStyle, EmptyObject: EmptyGuestPolicyAssignment, OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Assignment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Packages, actual.Packages, dcl.Info{ObjectFunction: compareGuestPolicyPackagesNewStyle, EmptyObject: EmptyGuestPolicyPackages, OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Packages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PackageRepositories, actual.PackageRepositories, dcl.Info{ObjectFunction: compareGuestPolicyPackageRepositoriesNewStyle, EmptyObject: EmptyGuestPolicyPackageRepositories, OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("PackageRepositories")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Recipes, actual.Recipes, dcl.Info{ObjectFunction: compareGuestPolicyRecipesNewStyle, EmptyObject: EmptyGuestPolicyRecipes, OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Recipes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGuestPolicyUpdateGuestPolicyOperation")}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
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
func compareGuestPolicyAssignmentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyAssignment)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyAssignment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignment or *GuestPolicyAssignment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyAssignment)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyAssignment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GroupLabels, actual.GroupLabels, dcl.Info{ObjectFunction: compareGuestPolicyAssignmentGroupLabelsNewStyle, EmptyObject: EmptyGuestPolicyAssignmentGroupLabels, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zones, actual.Zones, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zones")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Instances, actual.Instances, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Instances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceNamePrefixes, actual.InstanceNamePrefixes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceNamePrefixes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsTypes, actual.OsTypes, dcl.Info{ObjectFunction: compareGuestPolicyAssignmentOsTypesNewStyle, EmptyObject: EmptyGuestPolicyAssignmentOsTypes, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OsTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyAssignmentGroupLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyAssignmentGroupLabels)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyAssignmentGroupLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignmentGroupLabels or *GuestPolicyAssignmentGroupLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyAssignmentGroupLabels)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyAssignmentGroupLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignmentGroupLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyAssignmentOsTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyAssignmentOsTypes)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyAssignmentOsTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignmentOsTypes or *GuestPolicyAssignmentOsTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyAssignmentOsTypes)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyAssignmentOsTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyAssignmentOsTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OsShortName, actual.OsShortName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OsShortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsVersion, actual.OsVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OsVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OsArchitecture, actual.OsArchitecture, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OsArchitecture")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackages)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackages or *GuestPolicyPackages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackages)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DesiredState, actual.DesiredState, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DesiredState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Manager, actual.Manager, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Manager")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackageRepositoriesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackageRepositories)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackageRepositories)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositories or *GuestPolicyPackageRepositories", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackageRepositories)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackageRepositories)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositories", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Apt, actual.Apt, dcl.Info{ObjectFunction: compareGuestPolicyPackageRepositoriesAptNewStyle, EmptyObject: EmptyGuestPolicyPackageRepositoriesApt, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Apt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Yum, actual.Yum, dcl.Info{ObjectFunction: compareGuestPolicyPackageRepositoriesYumNewStyle, EmptyObject: EmptyGuestPolicyPackageRepositoriesYum, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Yum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zypper, actual.Zypper, dcl.Info{ObjectFunction: compareGuestPolicyPackageRepositoriesZypperNewStyle, EmptyObject: EmptyGuestPolicyPackageRepositoriesZypper, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zypper")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Goo, actual.Goo, dcl.Info{ObjectFunction: compareGuestPolicyPackageRepositoriesGooNewStyle, EmptyObject: EmptyGuestPolicyPackageRepositoriesGoo, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Goo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackageRepositoriesAptNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackageRepositoriesApt)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackageRepositoriesApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesApt or *GuestPolicyPackageRepositoriesApt", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackageRepositoriesApt)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackageRepositoriesApt)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesApt", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArchiveType, actual.ArchiveType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArchiveType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Distribution, actual.Distribution, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Distribution")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Components, actual.Components, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Components")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKey, actual.GpgKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GpgKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackageRepositoriesYumNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackageRepositoriesYum)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackageRepositoriesYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesYum or *GuestPolicyPackageRepositoriesYum", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackageRepositoriesYum)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackageRepositoriesYum)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesYum", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseUrl, actual.BaseUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKeys, actual.GpgKeys, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GpgKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackageRepositoriesZypperNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackageRepositoriesZypper)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackageRepositoriesZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesZypper or *GuestPolicyPackageRepositoriesZypper", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackageRepositoriesZypper)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackageRepositoriesZypper)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesZypper", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseUrl, actual.BaseUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GpgKeys, actual.GpgKeys, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GpgKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyPackageRepositoriesGooNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyPackageRepositoriesGoo)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyPackageRepositoriesGoo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesGoo or *GuestPolicyPackageRepositoriesGoo", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyPackageRepositoriesGoo)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyPackageRepositoriesGoo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyPackageRepositoriesGoo", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipes)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipes or *GuestPolicyRecipes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipes)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Artifacts, actual.Artifacts, dcl.Info{ObjectFunction: compareGuestPolicyRecipesArtifactsNewStyle, EmptyObject: EmptyGuestPolicyRecipesArtifacts, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Artifacts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstallSteps, actual.InstallSteps, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallSteps, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstallSteps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateSteps, actual.UpdateSteps, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateSteps, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateSteps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DesiredState, actual.DesiredState, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DesiredState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesArtifactsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesArtifacts)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesArtifacts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifacts or *GuestPolicyRecipesArtifacts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesArtifacts)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesArtifacts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifacts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Remote, actual.Remote, dcl.Info{ObjectFunction: compareGuestPolicyRecipesArtifactsRemoteNewStyle, EmptyObject: EmptyGuestPolicyRecipesArtifactsRemote, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Remote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gcs, actual.Gcs, dcl.Info{ObjectFunction: compareGuestPolicyRecipesArtifactsGcsNewStyle, EmptyObject: EmptyGuestPolicyRecipesArtifactsGcs, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Gcs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowInsecure, actual.AllowInsecure, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowInsecure")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesArtifactsRemoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesArtifactsRemote)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesArtifactsRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifactsRemote or *GuestPolicyRecipesArtifactsRemote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesArtifactsRemote)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesArtifactsRemote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifactsRemote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Checksum, actual.Checksum, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Checksum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesArtifactsGcsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesArtifactsGcs)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesArtifactsGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifactsGcs or *GuestPolicyRecipesArtifactsGcs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesArtifactsGcs)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesArtifactsGcs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesArtifactsGcs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Object, actual.Object, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Object")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallSteps)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallSteps or *GuestPolicyRecipesInstallSteps", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallSteps)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallSteps", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FileCopy, actual.FileCopy, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsFileCopyNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsFileCopy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileCopy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArchiveExtraction, actual.ArchiveExtraction, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsArchiveExtractionNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsArchiveExtraction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArchiveExtraction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MsiInstallation, actual.MsiInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsMsiInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsMsiInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MsiInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DpkgInstallation, actual.DpkgInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsDpkgInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsDpkgInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DpkgInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RpmInstallation, actual.RpmInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsRpmInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsRpmInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RpmInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileExec, actual.FileExec, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsFileExecNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsFileExec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileExec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScriptRun, actual.ScriptRun, dcl.Info{ObjectFunction: compareGuestPolicyRecipesInstallStepsScriptRunNewStyle, EmptyObject: EmptyGuestPolicyRecipesInstallStepsScriptRun, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScriptRun")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsFileCopyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsFileCopy)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsFileCopy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsFileCopy or *GuestPolicyRecipesInstallStepsFileCopy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsFileCopy)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsFileCopy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsFileCopy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Overwrite, actual.Overwrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Overwrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Permissions, actual.Permissions, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Permissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsArchiveExtractionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsArchiveExtraction)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsArchiveExtraction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsArchiveExtraction or *GuestPolicyRecipesInstallStepsArchiveExtraction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsArchiveExtraction)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsArchiveExtraction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsArchiveExtraction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsMsiInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsMsiInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsMsiInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsMsiInstallation or *GuestPolicyRecipesInstallStepsMsiInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsMsiInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsMsiInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsMsiInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Flags, actual.Flags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Flags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsDpkgInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsDpkgInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsDpkgInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsDpkgInstallation or *GuestPolicyRecipesInstallStepsDpkgInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsDpkgInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsDpkgInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsDpkgInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsRpmInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsRpmInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsRpmInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsRpmInstallation or *GuestPolicyRecipesInstallStepsRpmInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsRpmInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsRpmInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsRpmInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsFileExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsFileExec)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsFileExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsFileExec or *GuestPolicyRecipesInstallStepsFileExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsFileExec)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsFileExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsFileExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalPath, actual.LocalPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesInstallStepsScriptRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesInstallStepsScriptRun)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesInstallStepsScriptRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsScriptRun or *GuestPolicyRecipesInstallStepsScriptRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesInstallStepsScriptRun)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesInstallStepsScriptRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesInstallStepsScriptRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Script, actual.Script, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Script")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interpreter, actual.Interpreter, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Interpreter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateSteps)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateSteps or *GuestPolicyRecipesUpdateSteps", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateSteps)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateSteps", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FileCopy, actual.FileCopy, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsFileCopyNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsFileCopy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileCopy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArchiveExtraction, actual.ArchiveExtraction, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsArchiveExtractionNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArchiveExtraction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MsiInstallation, actual.MsiInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsMsiInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsMsiInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MsiInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DpkgInstallation, actual.DpkgInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsDpkgInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DpkgInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RpmInstallation, actual.RpmInstallation, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsRpmInstallationNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsRpmInstallation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RpmInstallation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileExec, actual.FileExec, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsFileExecNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsFileExec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileExec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScriptRun, actual.ScriptRun, dcl.Info{ObjectFunction: compareGuestPolicyRecipesUpdateStepsScriptRunNewStyle, EmptyObject: EmptyGuestPolicyRecipesUpdateStepsScriptRun, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScriptRun")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsFileCopyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsFileCopy)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsFileCopy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsFileCopy or *GuestPolicyRecipesUpdateStepsFileCopy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsFileCopy)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsFileCopy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsFileCopy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Overwrite, actual.Overwrite, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Overwrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Permissions, actual.Permissions, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Permissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsArchiveExtractionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsArchiveExtraction)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsArchiveExtraction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsArchiveExtraction or *GuestPolicyRecipesUpdateStepsArchiveExtraction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsArchiveExtraction)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsArchiveExtraction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsArchiveExtraction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsMsiInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsMsiInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsMsiInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsMsiInstallation or *GuestPolicyRecipesUpdateStepsMsiInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsMsiInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsMsiInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsMsiInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Flags, actual.Flags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Flags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsDpkgInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsDpkgInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsDpkgInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsDpkgInstallation or *GuestPolicyRecipesUpdateStepsDpkgInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsDpkgInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsDpkgInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsDpkgInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsRpmInstallationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsRpmInstallation)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsRpmInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsRpmInstallation or *GuestPolicyRecipesUpdateStepsRpmInstallation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsRpmInstallation)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsRpmInstallation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsRpmInstallation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsFileExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsFileExec)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsFileExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsFileExec or *GuestPolicyRecipesUpdateStepsFileExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsFileExec)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsFileExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsFileExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ArtifactId, actual.ArtifactId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalPath, actual.LocalPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGuestPolicyRecipesUpdateStepsScriptRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GuestPolicyRecipesUpdateStepsScriptRun)
	if !ok {
		desiredNotPointer, ok := d.(GuestPolicyRecipesUpdateStepsScriptRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsScriptRun or *GuestPolicyRecipesUpdateStepsScriptRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GuestPolicyRecipesUpdateStepsScriptRun)
	if !ok {
		actualNotPointer, ok := a.(GuestPolicyRecipesUpdateStepsScriptRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GuestPolicyRecipesUpdateStepsScriptRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Script, actual.Script, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Script")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedExitCodes, actual.AllowedExitCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedExitCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interpreter, actual.Interpreter, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Interpreter")); len(ds) != 0 || err != nil {
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
func (r *GuestPolicy) urlNormalized() *GuestPolicy {
	normalized := dcl.Copy(*r).(GuestPolicy)
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
	} else if v != nil {
		m["assignment"] = v
	}
	if v, err := expandGuestPolicyPackagesSlice(c, f.Packages); err != nil {
		return nil, fmt.Errorf("error expanding Packages into packages: %w", err)
	} else if v != nil {
		m["packages"] = v
	}
	if v, err := expandGuestPolicyPackageRepositoriesSlice(c, f.PackageRepositories); err != nil {
		return nil, fmt.Errorf("error expanding PackageRepositories into packageRepositories: %w", err)
	} else if v != nil {
		m["packageRepositories"] = v
	}
	if v, err := expandGuestPolicyRecipesSlice(c, f.Recipes); err != nil {
		return nil, fmt.Errorf("error expanding Recipes into recipes: %w", err)
	} else if v != nil {
		m["recipes"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
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

	res := &GuestPolicy{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Assignment = flattenGuestPolicyAssignment(c, m["assignment"])
	res.Packages = flattenGuestPolicyPackagesSlice(c, m["packages"])
	res.PackageRepositories = flattenGuestPolicyPackageRepositoriesSlice(c, m["packageRepositories"])
	res.Recipes = flattenGuestPolicyRecipesSlice(c, m["recipes"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Project = dcl.FlattenString(m["project"])

	return res
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyAssignment
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyAssignmentGroupLabels
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyAssignmentOsTypes
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackages
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackageRepositories
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackageRepositoriesApt
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackageRepositoriesYum
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackageRepositoriesZypper
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyPackageRepositoriesGoo
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipes
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesArtifacts
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesArtifactsRemote
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesArtifactsGcs
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallSteps
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsFileCopy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsArchiveExtraction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsMsiInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsDpkgInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsRpmInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsFileExec
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesInstallStepsScriptRun
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateSteps
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsFileCopy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsMsiInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsRpmInstallation
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsFileExec
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGuestPolicyRecipesUpdateStepsScriptRun
	}
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

type guestPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         guestPolicyApiOperation
}

func convertFieldDiffToGuestPolicyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]guestPolicyDiff, error) {
	var diffs []guestPolicyDiff
	for _, op := range ops {
		diff := guestPolicyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToguestPolicyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToguestPolicyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (guestPolicyApiOperation, error) {
	switch op {

	case "updateGuestPolicyUpdateGuestPolicyOperation":
		return &updateGuestPolicyUpdateGuestPolicyOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}