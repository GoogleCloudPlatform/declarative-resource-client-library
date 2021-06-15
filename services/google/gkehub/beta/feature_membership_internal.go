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
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *FeatureMembership) validate() error {

	if err := dcl.RequiredParameter(r.Membership, "Membership"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Feature, "Feature"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Configmanagement) {
		if err := r.Configmanagement.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigSync) {
		if err := r.ConfigSync.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PolicyController) {
		if err := r.PolicyController.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Binauthz) {
		if err := r.Binauthz.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HierarchyController) {
		if err := r.HierarchyController.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagementConfigSync) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Git) {
		if err := r.Git.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagementConfigSyncGit) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementPolicyController) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementBinauthz) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementHierarchyController) validate() error {
	return nil
}

func featureMembershipGetURL(userBasePath string, r *FeatureMembership) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"feature":  dcl.ValueOrEmptyString(r.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil
}

func featureMembershipListURL(userBasePath, project, location, feature string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"feature":  feature,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil

}

func featureMembershipCreateURL(userBasePath, project, location, feature string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"feature":  feature,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil

}

func featureMembershipDeleteURL(userBasePath string, r *FeatureMembership) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"feature":  dcl.ValueOrEmptyString(r.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, params), nil
}

// featureMembershipApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type featureMembershipApiOperation interface {
	do(context.Context, *FeatureMembership, *Client) error
}

// newUpdateFeatureMembershipUpdateFeatureMembershipRequest creates a request for an
// FeatureMembership resource's UpdateFeatureMembership update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFeatureMembershipUpdateFeatureMembershipRequest(ctx context.Context, f *FeatureMembership, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandFeatureMembershipConfigmanagement(c, f.Configmanagement); err != nil {
		return nil, fmt.Errorf("error expanding Configmanagement into configmanagement: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["configmanagement"] = v
	}
	return req, nil
}

// marshalUpdateFeatureMembershipUpdateFeatureMembershipRequest converts the update into
// the final JSON request body.
func marshalUpdateFeatureMembershipUpdateFeatureMembershipRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFeatureMembershipUpdateFeatureMembershipOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) deleteAllFeatureMembership(ctx context.Context, f func(*FeatureMembership) bool, resources []*FeatureMembership) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFeatureMembership(ctx, res)
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

type deleteFeatureMembershipOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFeatureMembershipOperation struct {
	response map[string]interface{}
}

func (op *createFeatureMembershipOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) featureMembershipDiffsForRawDesired(ctx context.Context, rawDesired *FeatureMembership, opts ...dcl.ApplyOption) (initial, desired *FeatureMembership, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FeatureMembership
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FeatureMembership); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected FeatureMembership, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFeatureMembership(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a FeatureMembership resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FeatureMembership resource: %v", err)
		}
		c.Config.Logger.Info("Found that FeatureMembership resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for FeatureMembership: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for FeatureMembership: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFeatureMembershipInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for FeatureMembership: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for FeatureMembership: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFeatureMembership(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFeatureMembershipInitialState(rawInitial, rawDesired *FeatureMembership) (*FeatureMembership, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial *FeatureMembership, opts ...dcl.ApplyOption) (*FeatureMembership, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Configmanagement = canonicalizeFeatureMembershipConfigmanagement(rawDesired.Configmanagement, nil, opts...)

		return rawDesired, nil
	}

	if dcl.NameToSelfLink(rawDesired.Membership, rawInitial.Membership) {
		rawDesired.Membership = rawInitial.Membership
	}
	if dcl.NameToSelfLink(rawDesired.Feature, rawInitial.Feature) {
		rawDesired.Feature = rawInitial.Feature
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	rawDesired.Configmanagement = canonicalizeFeatureMembershipConfigmanagement(rawDesired.Configmanagement, rawInitial.Configmanagement, opts...)

	return rawDesired, nil
}

func canonicalizeFeatureMembershipNewState(c *Client, rawNew, rawDesired *FeatureMembership) (*FeatureMembership, error) {

	rawNew.Membership = rawDesired.Membership

	rawNew.Feature = rawDesired.Feature

	rawNew.Location = rawDesired.Location

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Configmanagement) && dcl.IsEmptyValueIndirect(rawDesired.Configmanagement) {
		rawNew.Configmanagement = rawDesired.Configmanagement
	} else {
		rawNew.Configmanagement = canonicalizeNewFeatureMembershipConfigmanagement(c, rawDesired.Configmanagement, rawNew.Configmanagement)
	}

	return rawNew, nil
}

func canonicalizeFeatureMembershipConfigmanagement(des, initial *FeatureMembershipConfigmanagement, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ConfigSync = canonicalizeFeatureMembershipConfigmanagementConfigSync(des.ConfigSync, initial.ConfigSync, opts...)
	des.PolicyController = canonicalizeFeatureMembershipConfigmanagementPolicyController(des.PolicyController, initial.PolicyController, opts...)
	des.Binauthz = canonicalizeFeatureMembershipConfigmanagementBinauthz(des.Binauthz, initial.Binauthz, opts...)
	des.HierarchyController = canonicalizeFeatureMembershipConfigmanagementHierarchyController(des.HierarchyController, initial.HierarchyController, opts...)
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagement(c *Client, des, nw *FeatureMembershipConfigmanagement) *FeatureMembershipConfigmanagement {
	if des == nil || nw == nil {
		return nw
	}

	nw.ConfigSync = canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c, des.ConfigSync, nw.ConfigSync)
	nw.PolicyController = canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c, des.PolicyController, nw.PolicyController)
	nw.Binauthz = canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c, des.Binauthz, nw.Binauthz)
	nw.HierarchyController = canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c, des.HierarchyController, nw.HierarchyController)
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementSet(c *Client, des, nw []FeatureMembershipConfigmanagement) []FeatureMembershipConfigmanagement {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementSlice(c *Client, des, nw []FeatureMembershipConfigmanagement) []FeatureMembershipConfigmanagement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagement(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementConfigSync(des, initial *FeatureMembershipConfigmanagementConfigSync, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Git = canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(des.Git, initial.Git, opts...)
	if dcl.StringCanonicalize(des.SourceFormat, initial.SourceFormat) || dcl.IsZeroValue(des.SourceFormat) {
		des.SourceFormat = initial.SourceFormat
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c *Client, des, nw *FeatureMembershipConfigmanagementConfigSync) *FeatureMembershipConfigmanagementConfigSync {
	if des == nil || nw == nil {
		return nw
	}

	nw.Git = canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c, des.Git, nw.Git)
	if dcl.StringCanonicalize(des.SourceFormat, nw.SourceFormat) {
		nw.SourceFormat = des.SourceFormat
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncSet(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSync) []FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagementConfigSync
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementConfigSyncNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSync) []FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementConfigSync
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(des, initial *FeatureMembershipConfigmanagementConfigSyncGit, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SyncRepo, initial.SyncRepo) || dcl.IsZeroValue(des.SyncRepo) {
		des.SyncRepo = initial.SyncRepo
	}
	if dcl.StringCanonicalize(des.SyncBranch, initial.SyncBranch) || dcl.IsZeroValue(des.SyncBranch) {
		des.SyncBranch = initial.SyncBranch
	}
	if dcl.StringCanonicalize(des.PolicyDir, initial.PolicyDir) || dcl.IsZeroValue(des.PolicyDir) {
		des.PolicyDir = initial.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, initial.SyncWaitSecs) || dcl.IsZeroValue(des.SyncWaitSecs) {
		des.SyncWaitSecs = initial.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SyncRev, initial.SyncRev) || dcl.IsZeroValue(des.SyncRev) {
		des.SyncRev = initial.SyncRev
	}
	if dcl.StringCanonicalize(des.SecretType, initial.SecretType) || dcl.IsZeroValue(des.SecretType) {
		des.SecretType = initial.SecretType
	}
	if dcl.StringCanonicalize(des.HttpsProxy, initial.HttpsProxy) || dcl.IsZeroValue(des.HttpsProxy) {
		des.HttpsProxy = initial.HttpsProxy
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c *Client, des, nw *FeatureMembershipConfigmanagementConfigSyncGit) *FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SyncRepo, nw.SyncRepo) {
		nw.SyncRepo = des.SyncRepo
	}
	if dcl.StringCanonicalize(des.SyncBranch, nw.SyncBranch) {
		nw.SyncBranch = des.SyncBranch
	}
	if dcl.StringCanonicalize(des.PolicyDir, nw.PolicyDir) {
		nw.PolicyDir = des.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, nw.SyncWaitSecs) {
		nw.SyncWaitSecs = des.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SyncRev, nw.SyncRev) {
		nw.SyncRev = des.SyncRev
	}
	if dcl.StringCanonicalize(des.SecretType, nw.SecretType) {
		nw.SecretType = des.SecretType
	}
	if dcl.StringCanonicalize(des.HttpsProxy, nw.HttpsProxy) {
		nw.HttpsProxy = des.HttpsProxy
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGitSet(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncGit) []FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagementConfigSyncGit
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncGit) []FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementConfigSyncGit
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementPolicyController(des, initial *FeatureMembershipConfigmanagementPolicyController, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.IsZeroValue(des.ExemptableNamespaces) {
		des.ExemptableNamespaces = initial.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, initial.ReferentialRulesEnabled) || dcl.IsZeroValue(des.ReferentialRulesEnabled) {
		des.ReferentialRulesEnabled = initial.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, initial.LogDeniesEnabled) || dcl.IsZeroValue(des.LogDeniesEnabled) {
		des.LogDeniesEnabled = initial.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.TemplateLibraryInstalled, initial.TemplateLibraryInstalled) || dcl.IsZeroValue(des.TemplateLibraryInstalled) {
		des.TemplateLibraryInstalled = initial.TemplateLibraryInstalled
	}
	if dcl.StringCanonicalize(des.AuditIntervalSeconds, initial.AuditIntervalSeconds) || dcl.IsZeroValue(des.AuditIntervalSeconds) {
		des.AuditIntervalSeconds = initial.AuditIntervalSeconds
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c *Client, des, nw *FeatureMembershipConfigmanagementPolicyController) *FeatureMembershipConfigmanagementPolicyController {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.IsZeroValue(nw.ExemptableNamespaces) {
		nw.ExemptableNamespaces = des.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, nw.ReferentialRulesEnabled) {
		nw.ReferentialRulesEnabled = des.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, nw.LogDeniesEnabled) {
		nw.LogDeniesEnabled = des.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.TemplateLibraryInstalled, nw.TemplateLibraryInstalled) {
		nw.TemplateLibraryInstalled = des.TemplateLibraryInstalled
	}
	if dcl.StringCanonicalize(des.AuditIntervalSeconds, nw.AuditIntervalSeconds) {
		nw.AuditIntervalSeconds = des.AuditIntervalSeconds
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerSet(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyController) []FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagementPolicyController
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementPolicyControllerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyController) []FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementPolicyController
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementBinauthz(des, initial *FeatureMembershipConfigmanagementBinauthz, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c *Client, des, nw *FeatureMembershipConfigmanagementBinauthz) *FeatureMembershipConfigmanagementBinauthz {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementBinauthzSet(c *Client, des, nw []FeatureMembershipConfigmanagementBinauthz) []FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagementBinauthz
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementBinauthzNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementBinauthzSlice(c *Client, des, nw []FeatureMembershipConfigmanagementBinauthz) []FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementBinauthz
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementHierarchyController(des, initial *FeatureMembershipConfigmanagementHierarchyController, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.BoolCanonicalize(des.EnablePodTreeLabels, initial.EnablePodTreeLabels) || dcl.IsZeroValue(des.EnablePodTreeLabels) {
		des.EnablePodTreeLabels = initial.EnablePodTreeLabels
	}
	if dcl.BoolCanonicalize(des.EnableHierarchicalResourceQuota, initial.EnableHierarchicalResourceQuota) || dcl.IsZeroValue(des.EnableHierarchicalResourceQuota) {
		des.EnableHierarchicalResourceQuota = initial.EnableHierarchicalResourceQuota
	}

	return des
}

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c *Client, des, nw *FeatureMembershipConfigmanagementHierarchyController) *FeatureMembershipConfigmanagementHierarchyController {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.EnablePodTreeLabels, nw.EnablePodTreeLabels) {
		nw.EnablePodTreeLabels = des.EnablePodTreeLabels
	}
	if dcl.BoolCanonicalize(des.EnableHierarchicalResourceQuota, nw.EnableHierarchicalResourceQuota) {
		nw.EnableHierarchicalResourceQuota = des.EnableHierarchicalResourceQuota
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyControllerSet(c *Client, des, nw []FeatureMembershipConfigmanagementHierarchyController) []FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureMembershipConfigmanagementHierarchyController
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, des, nw []FeatureMembershipConfigmanagementHierarchyController) []FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementHierarchyController
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c, &d, &n))
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
func diffFeatureMembership(c *Client, desired, actual *FeatureMembership, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Membership, actual.Membership, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Membership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Feature, actual.Feature, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Feature")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Configmanagement, actual.Configmanagement, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagement, OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("Configmanagement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareFeatureMembershipConfigmanagementNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagement)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagement or *FeatureMembershipConfigmanagement", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagement)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagement", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigSync, actual.ConfigSync, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementConfigSyncNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementConfigSync, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigSync")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyController, actual.PolicyController, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementPolicyControllerNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementPolicyController, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyController")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Binauthz, actual.Binauthz, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementBinauthzNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementBinauthz, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Binauthz")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HierarchyController, actual.HierarchyController, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementHierarchyController, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HierarchyController")); len(ds) != 0 || err != nil {
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
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementConfigSyncNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementConfigSync)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementConfigSync)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSync or *FeatureMembershipConfigmanagementConfigSync", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementConfigSync)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementConfigSync)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSync", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Git, actual.Git, dcl.Info{ObjectFunction: compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementConfigSyncGit, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Git")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceFormat, actual.SourceFormat, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementConfigSyncGit)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementConfigSyncGit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncGit or *FeatureMembershipConfigmanagementConfigSyncGit", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementConfigSyncGit)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementConfigSyncGit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncGit", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SyncRepo, actual.SyncRepo, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncRepo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncBranch, actual.SyncBranch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncBranch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyDir, actual.PolicyDir, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyDir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncWaitSecs, actual.SyncWaitSecs, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncWaitSecs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncRev, actual.SyncRev, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncRev")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretType, actual.SecretType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecretType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsProxy, actual.HttpsProxy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsProxy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementPolicyControllerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementPolicyController)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementPolicyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyController or *FeatureMembershipConfigmanagementPolicyController", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementPolicyController)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementPolicyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyController", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemptableNamespaces, actual.ExemptableNamespaces, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemptableNamespaces")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReferentialRulesEnabled, actual.ReferentialRulesEnabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReferentialRulesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogDeniesEnabled, actual.LogDeniesEnabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LogDeniesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TemplateLibraryInstalled, actual.TemplateLibraryInstalled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TemplateLibraryInstalled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuditIntervalSeconds, actual.AuditIntervalSeconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuditIntervalSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementBinauthzNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementBinauthz)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementBinauthz)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementBinauthz or *FeatureMembershipConfigmanagementBinauthz", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementBinauthz)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementBinauthz)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementBinauthz", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController or *FeatureMembershipConfigmanagementHierarchyController", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnablePodTreeLabels, actual.EnablePodTreeLabels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnablePodTreeLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableHierarchicalResourceQuota, actual.EnableHierarchicalResourceQuota, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableHierarchicalResourceQuota")); len(ds) != 0 || err != nil {
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
func (r *FeatureMembership) urlNormalized() *FeatureMembership {
	normalized := dcl.Copy(*r).(FeatureMembership)
	normalized.Membership = dcl.SelfLinkToName(r.Membership)
	normalized.Feature = dcl.SelfLinkToName(r.Feature)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *FeatureMembership) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Feature)
}

func (r *FeatureMembership) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Feature)
}

func (r *FeatureMembership) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Feature)
}

func (r *FeatureMembership) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateFeatureMembership" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"feature":  dcl.ValueOrEmptyString(n.Feature),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FeatureMembership resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FeatureMembership) marshal(c *Client) ([]byte, error) {
	m, err := expandFeatureMembership(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FeatureMembership: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFeatureMembership decodes JSON responses into the FeatureMembership resource schema.
func unmarshalFeatureMembership(b []byte, c *Client) (*FeatureMembership, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFeatureMembership(m, c)
}

func unmarshalMapFeatureMembership(m map[string]interface{}, c *Client) (*FeatureMembership, error) {

	return flattenFeatureMembership(c, m), nil
}

// expandFeatureMembership expands FeatureMembership into a JSON request object.
func expandFeatureMembership(c *Client, f *FeatureMembership) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Membership into membership: %w", err)
	} else if v != nil {
		m["membership"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Feature into feature: %w", err)
	} else if v != nil {
		m["feature"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagement(c, f.Configmanagement); err != nil {
		return nil, fmt.Errorf("error expanding Configmanagement into configmanagement: %w", err)
	} else if v != nil {
		m["configmanagement"] = v
	}

	return m, nil
}

// flattenFeatureMembership flattens FeatureMembership from a JSON request object into the
// FeatureMembership type.
func flattenFeatureMembership(c *Client, i interface{}) *FeatureMembership {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &FeatureMembership{}
	res.Membership = dcl.FlattenString(m["membership"])
	res.Feature = dcl.FlattenString(m["feature"])
	res.Location = dcl.FlattenString(m["location"])
	res.Project = dcl.FlattenString(m["project"])
	res.Configmanagement = flattenFeatureMembershipConfigmanagement(c, m["configmanagement"])

	return res
}

// expandFeatureMembershipConfigmanagementMap expands the contents of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagementMap(c *Client, f map[string]FeatureMembershipConfigmanagement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementSlice expands the contents of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagementSlice(c *Client, f []FeatureMembershipConfigmanagement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementMap flattens the contents of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagement{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagement{}
	}

	items := make(map[string]FeatureMembershipConfigmanagement)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementSlice flattens the contents of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagement {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagement{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagement{}
	}

	items := make([]FeatureMembershipConfigmanagement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagement(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagement expands an instance of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagement(c *Client, f *FeatureMembershipConfigmanagement) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureMembershipConfigmanagementConfigSync(c, f.ConfigSync); err != nil {
		return nil, fmt.Errorf("error expanding ConfigSync into configSync: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configSync"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementPolicyController(c, f.PolicyController); err != nil {
		return nil, fmt.Errorf("error expanding PolicyController into policyController: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyController"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementBinauthz(c, f.Binauthz); err != nil {
		return nil, fmt.Errorf("error expanding Binauthz into binauthz: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["binauthz"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementHierarchyController(c, f.HierarchyController); err != nil {
		return nil, fmt.Errorf("error expanding HierarchyController into hierarchyController: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hierarchyController"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagement flattens an instance of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagement(c *Client, i interface{}) *FeatureMembershipConfigmanagement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagement{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagement
	}
	r.ConfigSync = flattenFeatureMembershipConfigmanagementConfigSync(c, m["configSync"])
	r.PolicyController = flattenFeatureMembershipConfigmanagementPolicyController(c, m["policyController"])
	r.Binauthz = flattenFeatureMembershipConfigmanagementBinauthz(c, m["binauthz"])
	r.HierarchyController = flattenFeatureMembershipConfigmanagementHierarchyController(c, m["hierarchyController"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandFeatureMembershipConfigmanagementConfigSyncMap expands the contents of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncMap(c *Client, f map[string]FeatureMembershipConfigmanagementConfigSync) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSync(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementConfigSyncSlice expands the contents of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, f []FeatureMembershipConfigmanagementConfigSync) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSync(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncMap flattens the contents of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagementConfigSync {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementConfigSync{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementConfigSync{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementConfigSync)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementConfigSync(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementConfigSyncSlice flattens the contents of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagementConfigSync {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementConfigSync{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementConfigSync{}
	}

	items := make([]FeatureMembershipConfigmanagementConfigSync, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementConfigSync(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagementConfigSync expands an instance of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSync(c *Client, f *FeatureMembershipConfigmanagementConfigSync) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, f.Git); err != nil {
		return nil, fmt.Errorf("error expanding Git into git: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["git"] = v
	}
	if v := f.SourceFormat; !dcl.IsEmptyValueIndirect(v) {
		m["sourceFormat"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementConfigSync flattens an instance of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSync(c *Client, i interface{}) *FeatureMembershipConfigmanagementConfigSync {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementConfigSync{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementConfigSync
	}
	r.Git = flattenFeatureMembershipConfigmanagementConfigSyncGit(c, m["git"])
	r.SourceFormat = dcl.FlattenString(m["sourceFormat"])

	return r
}

// expandFeatureMembershipConfigmanagementConfigSyncGitMap expands the contents of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGitMap(c *Client, f map[string]FeatureMembershipConfigmanagementConfigSyncGit) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementConfigSyncGitSlice expands the contents of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, f []FeatureMembershipConfigmanagementConfigSyncGit) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncGitMap flattens the contents of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGitMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagementConfigSyncGit {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementConfigSyncGit)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementConfigSyncGit(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementConfigSyncGitSlice flattens the contents of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagementConfigSyncGit {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	items := make([]FeatureMembershipConfigmanagementConfigSyncGit, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementConfigSyncGit(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagementConfigSyncGit expands an instance of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGit(c *Client, f *FeatureMembershipConfigmanagementConfigSyncGit) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SyncRepo; !dcl.IsEmptyValueIndirect(v) {
		m["syncRepo"] = v
	}
	if v := f.SyncBranch; !dcl.IsEmptyValueIndirect(v) {
		m["syncBranch"] = v
	}
	if v := f.PolicyDir; !dcl.IsEmptyValueIndirect(v) {
		m["policyDir"] = v
	}
	if v := f.SyncWaitSecs; !dcl.IsEmptyValueIndirect(v) {
		m["syncWaitSecs"] = v
	}
	if v := f.SyncRev; !dcl.IsEmptyValueIndirect(v) {
		m["syncRev"] = v
	}
	if v := f.SecretType; !dcl.IsEmptyValueIndirect(v) {
		m["secretType"] = v
	}
	if v := f.HttpsProxy; !dcl.IsEmptyValueIndirect(v) {
		m["httpsProxy"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncGit flattens an instance of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGit(c *Client, i interface{}) *FeatureMembershipConfigmanagementConfigSyncGit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementConfigSyncGit{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementConfigSyncGit
	}
	r.SyncRepo = dcl.FlattenString(m["syncRepo"])
	r.SyncBranch = dcl.FlattenString(m["syncBranch"])
	r.PolicyDir = dcl.FlattenString(m["policyDir"])
	r.SyncWaitSecs = dcl.FlattenString(m["syncWaitSecs"])
	r.SyncRev = dcl.FlattenString(m["syncRev"])
	r.SecretType = dcl.FlattenString(m["secretType"])
	r.HttpsProxy = dcl.FlattenString(m["httpsProxy"])

	return r
}

// expandFeatureMembershipConfigmanagementPolicyControllerMap expands the contents of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerMap(c *Client, f map[string]FeatureMembershipConfigmanagementPolicyController) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyController(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementPolicyControllerSlice expands the contents of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, f []FeatureMembershipConfigmanagementPolicyController) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyController(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMap flattens the contents of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagementPolicyController {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementPolicyController{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementPolicyController{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementPolicyController)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementPolicyController(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementPolicyControllerSlice flattens the contents of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagementPolicyController {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementPolicyController{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementPolicyController{}
	}

	items := make([]FeatureMembershipConfigmanagementPolicyController, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementPolicyController(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagementPolicyController expands an instance of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyController(c *Client, f *FeatureMembershipConfigmanagementPolicyController) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.ExemptableNamespaces; !dcl.IsEmptyValueIndirect(v) {
		m["exemptableNamespaces"] = v
	}
	if v := f.ReferentialRulesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["referentialRulesEnabled"] = v
	}
	if v := f.LogDeniesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["logDeniesEnabled"] = v
	}
	if v := f.TemplateLibraryInstalled; !dcl.IsEmptyValueIndirect(v) {
		m["templateLibraryInstalled"] = v
	}
	if v := f.AuditIntervalSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["auditIntervalSeconds"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementPolicyController flattens an instance of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyController(c *Client, i interface{}) *FeatureMembershipConfigmanagementPolicyController {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementPolicyController{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementPolicyController
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.ExemptableNamespaces = dcl.FlattenStringSlice(m["exemptableNamespaces"])
	r.ReferentialRulesEnabled = dcl.FlattenBool(m["referentialRulesEnabled"])
	r.LogDeniesEnabled = dcl.FlattenBool(m["logDeniesEnabled"])
	r.TemplateLibraryInstalled = dcl.FlattenBool(m["templateLibraryInstalled"])
	r.AuditIntervalSeconds = dcl.FlattenString(m["auditIntervalSeconds"])

	return r
}

// expandFeatureMembershipConfigmanagementBinauthzMap expands the contents of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthzMap(c *Client, f map[string]FeatureMembershipConfigmanagementBinauthz) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementBinauthz(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementBinauthzSlice expands the contents of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthzSlice(c *Client, f []FeatureMembershipConfigmanagementBinauthz) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementBinauthz(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementBinauthzMap flattens the contents of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthzMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagementBinauthz {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementBinauthz{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementBinauthz{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementBinauthz)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementBinauthz(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementBinauthzSlice flattens the contents of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthzSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagementBinauthz {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementBinauthz{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementBinauthz{}
	}

	items := make([]FeatureMembershipConfigmanagementBinauthz, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementBinauthz(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagementBinauthz expands an instance of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthz(c *Client, f *FeatureMembershipConfigmanagementBinauthz) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementBinauthz flattens an instance of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthz(c *Client, i interface{}) *FeatureMembershipConfigmanagementBinauthz {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementBinauthz{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementBinauthz
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandFeatureMembershipConfigmanagementHierarchyControllerMap expands the contents of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyControllerMap(c *Client, f map[string]FeatureMembershipConfigmanagementHierarchyController) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementHierarchyController(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementHierarchyControllerSlice expands the contents of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, f []FeatureMembershipConfigmanagementHierarchyController) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementHierarchyController(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementHierarchyControllerMap flattens the contents of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyControllerMap(c *Client, i interface{}) map[string]FeatureMembershipConfigmanagementHierarchyController {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementHierarchyController{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementHierarchyController{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementHierarchyController)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementHierarchyController(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementHierarchyControllerSlice flattens the contents of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, i interface{}) []FeatureMembershipConfigmanagementHierarchyController {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementHierarchyController{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementHierarchyController{}
	}

	items := make([]FeatureMembershipConfigmanagementHierarchyController, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementHierarchyController(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureMembershipConfigmanagementHierarchyController expands an instance of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyController(c *Client, f *FeatureMembershipConfigmanagementHierarchyController) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.EnablePodTreeLabels; !dcl.IsEmptyValueIndirect(v) {
		m["enablePodTreeLabels"] = v
	}
	if v := f.EnableHierarchicalResourceQuota; !dcl.IsEmptyValueIndirect(v) {
		m["enableHierarchicalResourceQuota"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementHierarchyController flattens an instance of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyController(c *Client, i interface{}) *FeatureMembershipConfigmanagementHierarchyController {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementHierarchyController{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementHierarchyController
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.EnablePodTreeLabels = dcl.FlattenBool(m["enablePodTreeLabels"])
	r.EnableHierarchicalResourceQuota = dcl.FlattenBool(m["enableHierarchicalResourceQuota"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FeatureMembership) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFeatureMembership(b, c)
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
		if nr.Feature == nil && ncr.Feature == nil {
			c.Config.Logger.Info("Both Feature fields null - considering equal.")
		} else if nr.Feature == nil || ncr.Feature == nil {
			c.Config.Logger.Info("Only one Feature field is null - considering unequal.")
			return false
		} else if *nr.Feature != *ncr.Feature {
			return false
		}
		return true
	}
}

type featureMembershipDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureMembershipApiOperation
}

func convertFieldDiffToFeatureMembershipOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]featureMembershipDiff, error) {
	var diffs []featureMembershipDiff
	for _, op := range ops {
		diff := featureMembershipDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTofeatureMembershipApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTofeatureMembershipApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (featureMembershipApiOperation, error) {
	switch op {

	case "updateFeatureMembershipUpdateFeatureMembershipOperation":
		return &updateFeatureMembershipUpdateFeatureMembershipOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
