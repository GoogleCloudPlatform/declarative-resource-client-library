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
package appengine

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

func (r *Version) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "runtime"); err != nil {
		return err
	}
	if err := dcl.Required(r, "deployment"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.App, "App"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Service, "Service"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.AutomaticScaling) {
		if err := r.AutomaticScaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BasicScaling) {
		if err := r.BasicScaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManualScaling) {
		if err := r.ManualScaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Network) {
		if err := r.Network.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Resources) {
		if err := r.Resources.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ApiConfig) {
		if err := r.ApiConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deployment) {
		if err := r.Deployment.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HealthCheck) {
		if err := r.HealthCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReadinessCheck) {
		if err := r.ReadinessCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LivenessCheck) {
		if err := r.LivenessCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Entrypoint) {
		if err := r.Entrypoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VPCAccessConnector) {
		if err := r.VPCAccessConnector.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *VersionAutomaticScaling) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CpuUtilization) {
		if err := r.CpuUtilization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestUtilization) {
		if err := r.RequestUtilization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DiskUtilization) {
		if err := r.DiskUtilization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkUtilization) {
		if err := r.NetworkUtilization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StandardSchedulerSettings) {
		if err := r.StandardSchedulerSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *VersionAutomaticScalingCpuUtilization) validate() error {
	return nil
}
func (r *VersionAutomaticScalingRequestUtilization) validate() error {
	return nil
}
func (r *VersionAutomaticScalingDiskUtilization) validate() error {
	return nil
}
func (r *VersionAutomaticScalingNetworkUtilization) validate() error {
	return nil
}
func (r *VersionAutomaticScalingStandardSchedulerSettings) validate() error {
	return nil
}
func (r *VersionBasicScaling) validate() error {
	return nil
}
func (r *VersionManualScaling) validate() error {
	return nil
}
func (r *VersionNetwork) validate() error {
	return nil
}
func (r *VersionResources) validate() error {
	return nil
}
func (r *VersionResourcesVolumes) validate() error {
	return nil
}
func (r *VersionHandlers) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StaticFiles) {
		if err := r.StaticFiles.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Script) {
		if err := r.Script.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ApiEndpoint) {
		if err := r.ApiEndpoint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *VersionHandlersStaticFiles) validate() error {
	return nil
}
func (r *VersionHandlersScript) validate() error {
	return nil
}
func (r *VersionHandlersApiEndpoint) validate() error {
	return nil
}
func (r *VersionErrorHandlers) validate() error {
	return nil
}
func (r *VersionLibraries) validate() error {
	return nil
}
func (r *VersionApiConfig) validate() error {
	return nil
}
func (r *VersionDeployment) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Container) {
		if err := r.Container.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Zip) {
		if err := r.Zip.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudBuildOptions) {
		if err := r.CloudBuildOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *VersionDeploymentFiles) validate() error {
	return nil
}
func (r *VersionDeploymentContainer) validate() error {
	return nil
}
func (r *VersionDeploymentZip) validate() error {
	return nil
}
func (r *VersionDeploymentCloudBuildOptions) validate() error {
	return nil
}
func (r *VersionHealthCheck) validate() error {
	return nil
}
func (r *VersionReadinessCheck) validate() error {
	return nil
}
func (r *VersionLivenessCheck) validate() error {
	return nil
}
func (r *VersionEntrypoint) validate() error {
	return nil
}
func (r *VersionVPCAccessConnector) validate() error {
	return nil
}

func versionGetURL(userBasePath string, r *Version) (string, error) {
	params := map[string]interface{}{
		"app":     dcl.ValueOrEmptyString(r.App),
		"service": dcl.ValueOrEmptyString(r.Service),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("apps/{{app}}/services/{{service}}/versions/{{name}}?view=FULL", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

func versionListURL(userBasePath, app, service string) (string, error) {
	params := map[string]interface{}{
		"app":     app,
		"service": service,
	}
	return dcl.URL("apps/{{app}}/services/{{service}}/versions", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func versionCreateURL(userBasePath, app, service string) (string, error) {
	params := map[string]interface{}{
		"app":     app,
		"service": service,
	}
	return dcl.URL("apps/{{app}}/services/{{service}}/versions", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func versionDeleteURL(userBasePath string, r *Version) (string, error) {
	params := map[string]interface{}{
		"app":     dcl.ValueOrEmptyString(r.App),
		"service": dcl.ValueOrEmptyString(r.Service),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("apps/{{app}}/services/{{service}}/versions/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

// versionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type versionApiOperation interface {
	do(context.Context, *Version, *Client) error
}

// newUpdateVersionPatchVersionRequest creates a request for an
// Version resource's PatchVersion update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateVersionPatchVersionRequest(ctx context.Context, f *Version, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.ConsumerName; !dcl.IsEmptyValueIndirect(v) {
		req["consumerName"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["id"] = v
	}
	if v, err := expandVersionAutomaticScaling(c, f.AutomaticScaling); err != nil {
		return nil, fmt.Errorf("error expanding AutomaticScaling into automaticScaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["automaticScaling"] = v
	}
	if v, err := expandVersionBasicScaling(c, f.BasicScaling); err != nil {
		return nil, fmt.Errorf("error expanding BasicScaling into basicScaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["basicScaling"] = v
	}
	if v, err := expandVersionManualScaling(c, f.ManualScaling); err != nil {
		return nil, fmt.Errorf("error expanding ManualScaling into manualScaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["manualScaling"] = v
	}
	if v := f.InboundServices; !dcl.IsEmptyValueIndirect(v) {
		req["inboundServices"] = v
	}
	if v := f.InstanceClass; !dcl.IsEmptyValueIndirect(v) {
		req["instanceClass"] = v
	}
	if v, err := expandVersionNetwork(c, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["network"] = v
	}
	if v := f.Zones; !dcl.IsEmptyValueIndirect(v) {
		req["zones"] = v
	}
	if v, err := expandVersionResources(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["resources"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		req["runtime"] = v
	}
	if v := f.RuntimeChannel; !dcl.IsEmptyValueIndirect(v) {
		req["runtimeChannel"] = v
	}
	if v := f.Threadsafe; !dcl.IsEmptyValueIndirect(v) {
		req["threadsafe"] = v
	}
	if v := f.Vm; !dcl.IsEmptyValueIndirect(v) {
		req["vm"] = v
	}
	if v := f.BetaSettings; !dcl.IsEmptyValueIndirect(v) {
		req["betaSettings"] = v
	}
	if v := f.Env; !dcl.IsEmptyValueIndirect(v) {
		req["env"] = v
	}
	if v := f.ServingStatus; !dcl.IsEmptyValueIndirect(v) {
		req["servingStatus"] = v
	}
	if v := f.CreatedBy; !dcl.IsEmptyValueIndirect(v) {
		req["createdBy"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		req["createTime"] = v
	}
	if v := f.DiskUsageBytes; !dcl.IsEmptyValueIndirect(v) {
		req["diskUsageBytes"] = v
	}
	if v := f.RuntimeApiVersion; !dcl.IsEmptyValueIndirect(v) {
		req["runtimeApiVersion"] = v
	}
	if v := f.RuntimeMainExecutablePath; !dcl.IsEmptyValueIndirect(v) {
		req["runtimeMainExecutablePath"] = v
	}
	if v, err := expandVersionHandlersSlice(c, f.Handlers); err != nil {
		return nil, fmt.Errorf("error expanding Handlers into handlers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["handlers"] = v
	}
	if v, err := expandVersionErrorHandlersSlice(c, f.ErrorHandlers); err != nil {
		return nil, fmt.Errorf("error expanding ErrorHandlers into errorHandlers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["errorHandlers"] = v
	}
	if v, err := expandVersionLibrariesSlice(c, f.Libraries); err != nil {
		return nil, fmt.Errorf("error expanding Libraries into libraries: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["libraries"] = v
	}
	if v, err := expandVersionApiConfig(c, f.ApiConfig); err != nil {
		return nil, fmt.Errorf("error expanding ApiConfig into apiConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["apiConfig"] = v
	}
	if v := f.EnvVariables; !dcl.IsEmptyValueIndirect(v) {
		req["envVariables"] = v
	}
	if v := f.DefaultExpiration; !dcl.IsEmptyValueIndirect(v) {
		req["defaultExpiration"] = v
	}
	if v, err := expandVersionHealthCheck(c, f.HealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HealthCheck into healthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["healthCheck"] = v
	}
	if v, err := expandVersionReadinessCheck(c, f.ReadinessCheck); err != nil {
		return nil, fmt.Errorf("error expanding ReadinessCheck into readinessCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["readinessCheck"] = v
	}
	if v, err := expandVersionLivenessCheck(c, f.LivenessCheck); err != nil {
		return nil, fmt.Errorf("error expanding LivenessCheck into livenessCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["livenessCheck"] = v
	}
	if v := f.NobuildFilesRegex; !dcl.IsEmptyValueIndirect(v) {
		req["nobuildFilesRegex"] = v
	}
	if v := f.VersionUrl; !dcl.IsEmptyValueIndirect(v) {
		req["versionUrl"] = v
	}
	if v, err := expandVersionEntrypoint(c, f.Entrypoint); err != nil {
		return nil, fmt.Errorf("error expanding Entrypoint into entrypoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["entrypoint"] = v
	}
	if v, err := expandVersionVPCAccessConnector(c, f.VPCAccessConnector); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccessConnector into vpcAccessConnector: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["vpcAccessConnector"] = v
	}
	return req, nil
}

// marshalUpdateVersionPatchVersionRequest converts the update into
// the final JSON request body.
func marshalUpdateVersionPatchVersionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateVersionPatchVersionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateVersionPatchVersionOperation) do(ctx context.Context, r *Version, c *Client) error {
	_, err := c.GetVersion(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchVersion")
	if err != nil {
		return err
	}

	req, err := newUpdateVersionPatchVersionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateVersionPatchVersionRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listVersionRaw(ctx context.Context, app, service, pageToken string, pageSize int32) ([]byte, error) {
	u, err := versionListURL(c.Config.BasePath, app, service)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != VersionMaxPage {
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

type listVersionOperation struct {
	Versions []map[string]interface{} `json:"versions"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listVersion(ctx context.Context, app, service, pageToken string, pageSize int32) ([]*Version, string, error) {
	b, err := c.listVersionRaw(ctx, app, service, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listVersionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Version
	for _, v := range m.Versions {
		res, err := unmarshalMapVersion(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.App = &app
		res.Service = &service
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllVersion(ctx context.Context, f func(*Version) bool, resources []*Version) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteVersion(ctx, res)
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

type deleteVersionOperation struct{}

func (op *deleteVersionOperation) do(ctx context.Context, r *Version, c *Client) error {

	_, err := c.GetVersion(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Version not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetVersion checking for existence. error: %v", err)
		return err
	}

	u, err := versionDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetVersion(ctx, r.urlNormalized())
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
type createVersionOperation struct {
	response map[string]interface{}
}

func (op *createVersionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createVersionOperation) do(ctx context.Context, r *Version, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	app, service := r.createFields()
	u, err := versionCreateURL(c.Config.BasePath, app, service)

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
	if err := o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetVersion(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getVersionRaw(ctx context.Context, r *Version) ([]byte, error) {

	u, err := versionGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) versionDiffsForRawDesired(ctx context.Context, rawDesired *Version, opts ...dcl.ApplyOption) (initial, desired *Version, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Version
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Version); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Version, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetVersion(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Version resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Version resource: %v", err)
		}
		c.Config.Logger.Info("Found that Version resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeVersionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Version: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Version: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeVersionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Version: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeVersionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Version: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffVersion(c, desired, initial, opts...)
	fmt.Printf("newDiffs: %v\n", diffs)
	return initial, desired, diffs, err
}

func canonicalizeVersionInitialState(rawInitial, rawDesired *Version) (*Version, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeVersionDesiredState(rawDesired, rawInitial *Version, opts ...dcl.ApplyOption) (*Version, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.AutomaticScaling = canonicalizeVersionAutomaticScaling(rawDesired.AutomaticScaling, nil, opts...)
		rawDesired.BasicScaling = canonicalizeVersionBasicScaling(rawDesired.BasicScaling, nil, opts...)
		rawDesired.ManualScaling = canonicalizeVersionManualScaling(rawDesired.ManualScaling, nil, opts...)
		rawDesired.Network = canonicalizeVersionNetwork(rawDesired.Network, nil, opts...)
		rawDesired.Resources = canonicalizeVersionResources(rawDesired.Resources, nil, opts...)
		rawDesired.ApiConfig = canonicalizeVersionApiConfig(rawDesired.ApiConfig, nil, opts...)
		rawDesired.Deployment = canonicalizeVersionDeployment(rawDesired.Deployment, nil, opts...)
		rawDesired.HealthCheck = canonicalizeVersionHealthCheck(rawDesired.HealthCheck, nil, opts...)
		rawDesired.ReadinessCheck = canonicalizeVersionReadinessCheck(rawDesired.ReadinessCheck, nil, opts...)
		rawDesired.LivenessCheck = canonicalizeVersionLivenessCheck(rawDesired.LivenessCheck, nil, opts...)
		rawDesired.Entrypoint = canonicalizeVersionEntrypoint(rawDesired.Entrypoint, nil, opts...)
		rawDesired.VPCAccessConnector = canonicalizeVersionVPCAccessConnector(rawDesired.VPCAccessConnector, nil, opts...)

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.ConsumerName, rawInitial.ConsumerName) {
		rawDesired.ConsumerName = rawInitial.ConsumerName
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.AutomaticScaling = canonicalizeVersionAutomaticScaling(rawDesired.AutomaticScaling, rawInitial.AutomaticScaling, opts...)
	rawDesired.BasicScaling = canonicalizeVersionBasicScaling(rawDesired.BasicScaling, rawInitial.BasicScaling, opts...)
	rawDesired.ManualScaling = canonicalizeVersionManualScaling(rawDesired.ManualScaling, rawInitial.ManualScaling, opts...)
	if dcl.IsZeroValue(rawDesired.InboundServices) {
		rawDesired.InboundServices = rawInitial.InboundServices
	}
	if dcl.StringCanonicalize(rawDesired.InstanceClass, rawInitial.InstanceClass) {
		rawDesired.InstanceClass = rawInitial.InstanceClass
	}
	rawDesired.Network = canonicalizeVersionNetwork(rawDesired.Network, rawInitial.Network, opts...)
	if dcl.IsZeroValue(rawDesired.Zones) {
		rawDesired.Zones = rawInitial.Zones
	}
	rawDesired.Resources = canonicalizeVersionResources(rawDesired.Resources, rawInitial.Resources, opts...)
	if dcl.StringCanonicalize(rawDesired.Runtime, rawInitial.Runtime) {
		rawDesired.Runtime = rawInitial.Runtime
	}
	if dcl.StringCanonicalize(rawDesired.RuntimeChannel, rawInitial.RuntimeChannel) {
		rawDesired.RuntimeChannel = rawInitial.RuntimeChannel
	}
	if dcl.BoolCanonicalize(rawDesired.Threadsafe, rawInitial.Threadsafe) {
		rawDesired.Threadsafe = rawInitial.Threadsafe
	}
	if dcl.BoolCanonicalize(rawDesired.Vm, rawInitial.Vm) {
		rawDesired.Vm = rawInitial.Vm
	}
	if dcl.IsZeroValue(rawDesired.BetaSettings) {
		rawDesired.BetaSettings = rawInitial.BetaSettings
	}
	if dcl.StringCanonicalize(rawDesired.Env, rawInitial.Env) {
		rawDesired.Env = rawInitial.Env
	}
	if dcl.IsZeroValue(rawDesired.ServingStatus) {
		rawDesired.ServingStatus = rawInitial.ServingStatus
	}
	if dcl.StringCanonicalize(rawDesired.RuntimeApiVersion, rawInitial.RuntimeApiVersion) {
		rawDesired.RuntimeApiVersion = rawInitial.RuntimeApiVersion
	}
	if dcl.StringCanonicalize(rawDesired.RuntimeMainExecutablePath, rawInitial.RuntimeMainExecutablePath) {
		rawDesired.RuntimeMainExecutablePath = rawInitial.RuntimeMainExecutablePath
	}
	if dcl.IsZeroValue(rawDesired.Handlers) {
		rawDesired.Handlers = rawInitial.Handlers
	}
	if dcl.IsZeroValue(rawDesired.ErrorHandlers) {
		rawDesired.ErrorHandlers = rawInitial.ErrorHandlers
	}
	if dcl.IsZeroValue(rawDesired.Libraries) {
		rawDesired.Libraries = rawInitial.Libraries
	}
	rawDesired.ApiConfig = canonicalizeVersionApiConfig(rawDesired.ApiConfig, rawInitial.ApiConfig, opts...)
	if dcl.IsZeroValue(rawDesired.EnvVariables) {
		rawDesired.EnvVariables = rawInitial.EnvVariables
	}
	if dcl.StringCanonicalize(rawDesired.DefaultExpiration, rawInitial.DefaultExpiration) {
		rawDesired.DefaultExpiration = rawInitial.DefaultExpiration
	}
	rawDesired.Deployment = canonicalizeVersionDeployment(rawDesired.Deployment, rawInitial.Deployment, opts...)
	rawDesired.HealthCheck = canonicalizeVersionHealthCheck(rawDesired.HealthCheck, rawInitial.HealthCheck, opts...)
	rawDesired.ReadinessCheck = canonicalizeVersionReadinessCheck(rawDesired.ReadinessCheck, rawInitial.ReadinessCheck, opts...)
	rawDesired.LivenessCheck = canonicalizeVersionLivenessCheck(rawDesired.LivenessCheck, rawInitial.LivenessCheck, opts...)
	if dcl.StringCanonicalize(rawDesired.NobuildFilesRegex, rawInitial.NobuildFilesRegex) {
		rawDesired.NobuildFilesRegex = rawInitial.NobuildFilesRegex
	}
	rawDesired.Entrypoint = canonicalizeVersionEntrypoint(rawDesired.Entrypoint, rawInitial.Entrypoint, opts...)
	rawDesired.VPCAccessConnector = canonicalizeVersionVPCAccessConnector(rawDesired.VPCAccessConnector, rawInitial.VPCAccessConnector, opts...)
	if dcl.NameToSelfLink(rawDesired.App, rawInitial.App) {
		rawDesired.App = rawInitial.App
	}
	if dcl.NameToSelfLink(rawDesired.Service, rawInitial.Service) {
		rawDesired.Service = rawInitial.Service
	}

	return rawDesired, nil
}

func canonicalizeVersionNewState(c *Client, rawNew, rawDesired *Version) (*Version, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ConsumerName) && dcl.IsEmptyValueIndirect(rawDesired.ConsumerName) {
		rawNew.ConsumerName = rawDesired.ConsumerName
	} else {
		if dcl.StringCanonicalize(rawDesired.ConsumerName, rawNew.ConsumerName) {
			rawNew.ConsumerName = rawDesired.ConsumerName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AutomaticScaling) && dcl.IsEmptyValueIndirect(rawDesired.AutomaticScaling) {
		rawNew.AutomaticScaling = rawDesired.AutomaticScaling
	} else {
		rawNew.AutomaticScaling = canonicalizeNewVersionAutomaticScaling(c, rawDesired.AutomaticScaling, rawNew.AutomaticScaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.BasicScaling) && dcl.IsEmptyValueIndirect(rawDesired.BasicScaling) {
		rawNew.BasicScaling = rawDesired.BasicScaling
	} else {
		rawNew.BasicScaling = canonicalizeNewVersionBasicScaling(c, rawDesired.BasicScaling, rawNew.BasicScaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ManualScaling) && dcl.IsEmptyValueIndirect(rawDesired.ManualScaling) {
		rawNew.ManualScaling = rawDesired.ManualScaling
	} else {
		rawNew.ManualScaling = canonicalizeNewVersionManualScaling(c, rawDesired.ManualScaling, rawNew.ManualScaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InboundServices) && dcl.IsEmptyValueIndirect(rawDesired.InboundServices) {
		rawNew.InboundServices = rawDesired.InboundServices
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceClass) && dcl.IsEmptyValueIndirect(rawDesired.InstanceClass) {
		rawNew.InstanceClass = rawDesired.InstanceClass
	} else {
		if dcl.StringCanonicalize(rawDesired.InstanceClass, rawNew.InstanceClass) {
			rawNew.InstanceClass = rawDesired.InstanceClass
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		rawNew.Network = canonicalizeNewVersionNetwork(c, rawDesired.Network, rawNew.Network)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zones) && dcl.IsEmptyValueIndirect(rawDesired.Zones) {
		rawNew.Zones = rawDesired.Zones
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Resources) && dcl.IsEmptyValueIndirect(rawDesired.Resources) {
		rawNew.Resources = rawDesired.Resources
	} else {
		rawNew.Resources = canonicalizeNewVersionResources(c, rawDesired.Resources, rawNew.Resources)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Runtime) && dcl.IsEmptyValueIndirect(rawDesired.Runtime) {
		rawNew.Runtime = rawDesired.Runtime
	} else {
		if dcl.StringCanonicalize(rawDesired.Runtime, rawNew.Runtime) {
			rawNew.Runtime = rawDesired.Runtime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeChannel) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeChannel) {
		rawNew.RuntimeChannel = rawDesired.RuntimeChannel
	} else {
		if dcl.StringCanonicalize(rawDesired.RuntimeChannel, rawNew.RuntimeChannel) {
			rawNew.RuntimeChannel = rawDesired.RuntimeChannel
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Threadsafe) && dcl.IsEmptyValueIndirect(rawDesired.Threadsafe) {
		rawNew.Threadsafe = rawDesired.Threadsafe
	} else {
		if dcl.BoolCanonicalize(rawDesired.Threadsafe, rawNew.Threadsafe) {
			rawNew.Threadsafe = rawDesired.Threadsafe
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Vm) && dcl.IsEmptyValueIndirect(rawDesired.Vm) {
		rawNew.Vm = rawDesired.Vm
	} else {
		if dcl.BoolCanonicalize(rawDesired.Vm, rawNew.Vm) {
			rawNew.Vm = rawDesired.Vm
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.BetaSettings) && dcl.IsEmptyValueIndirect(rawDesired.BetaSettings) {
		rawNew.BetaSettings = rawDesired.BetaSettings
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Env) && dcl.IsEmptyValueIndirect(rawDesired.Env) {
		rawNew.Env = rawDesired.Env
	} else {
		if dcl.StringCanonicalize(rawDesired.Env, rawNew.Env) {
			rawNew.Env = rawDesired.Env
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServingStatus) && dcl.IsEmptyValueIndirect(rawDesired.ServingStatus) {
		rawNew.ServingStatus = rawDesired.ServingStatus
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedBy) && dcl.IsEmptyValueIndirect(rawDesired.CreatedBy) {
		rawNew.CreatedBy = rawDesired.CreatedBy
	} else {
		if dcl.StringCanonicalize(rawDesired.CreatedBy, rawNew.CreatedBy) {
			rawNew.CreatedBy = rawDesired.CreatedBy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskUsageBytes) && dcl.IsEmptyValueIndirect(rawDesired.DiskUsageBytes) {
		rawNew.DiskUsageBytes = rawDesired.DiskUsageBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeApiVersion) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeApiVersion) {
		rawNew.RuntimeApiVersion = rawDesired.RuntimeApiVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.RuntimeApiVersion, rawNew.RuntimeApiVersion) {
			rawNew.RuntimeApiVersion = rawDesired.RuntimeApiVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeMainExecutablePath) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeMainExecutablePath) {
		rawNew.RuntimeMainExecutablePath = rawDesired.RuntimeMainExecutablePath
	} else {
		if dcl.StringCanonicalize(rawDesired.RuntimeMainExecutablePath, rawNew.RuntimeMainExecutablePath) {
			rawNew.RuntimeMainExecutablePath = rawDesired.RuntimeMainExecutablePath
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Handlers) && dcl.IsEmptyValueIndirect(rawDesired.Handlers) {
		rawNew.Handlers = rawDesired.Handlers
	} else {
		rawNew.Handlers = canonicalizeNewVersionHandlersSlice(c, rawDesired.Handlers, rawNew.Handlers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ErrorHandlers) && dcl.IsEmptyValueIndirect(rawDesired.ErrorHandlers) {
		rawNew.ErrorHandlers = rawDesired.ErrorHandlers
	} else {
		rawNew.ErrorHandlers = canonicalizeNewVersionErrorHandlersSlice(c, rawDesired.ErrorHandlers, rawNew.ErrorHandlers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Libraries) && dcl.IsEmptyValueIndirect(rawDesired.Libraries) {
		rawNew.Libraries = rawDesired.Libraries
	} else {
		rawNew.Libraries = canonicalizeNewVersionLibrariesSlice(c, rawDesired.Libraries, rawNew.Libraries)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ApiConfig) && dcl.IsEmptyValueIndirect(rawDesired.ApiConfig) {
		rawNew.ApiConfig = rawDesired.ApiConfig
	} else {
		rawNew.ApiConfig = canonicalizeNewVersionApiConfig(c, rawDesired.ApiConfig, rawNew.ApiConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnvVariables) && dcl.IsEmptyValueIndirect(rawDesired.EnvVariables) {
		rawNew.EnvVariables = rawDesired.EnvVariables
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultExpiration) && dcl.IsEmptyValueIndirect(rawDesired.DefaultExpiration) {
		rawNew.DefaultExpiration = rawDesired.DefaultExpiration
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultExpiration, rawNew.DefaultExpiration) {
			rawNew.DefaultExpiration = rawDesired.DefaultExpiration
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deployment) && dcl.IsEmptyValueIndirect(rawDesired.Deployment) {
		rawNew.Deployment = rawDesired.Deployment
	} else {
		rawNew.Deployment = canonicalizeNewVersionDeployment(c, rawDesired.Deployment, rawNew.Deployment)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.HealthCheck) {
		rawNew.HealthCheck = rawDesired.HealthCheck
	} else {
		rawNew.HealthCheck = canonicalizeNewVersionHealthCheck(c, rawDesired.HealthCheck, rawNew.HealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReadinessCheck) && dcl.IsEmptyValueIndirect(rawDesired.ReadinessCheck) {
		rawNew.ReadinessCheck = rawDesired.ReadinessCheck
	} else {
		rawNew.ReadinessCheck = canonicalizeNewVersionReadinessCheck(c, rawDesired.ReadinessCheck, rawNew.ReadinessCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LivenessCheck) && dcl.IsEmptyValueIndirect(rawDesired.LivenessCheck) {
		rawNew.LivenessCheck = rawDesired.LivenessCheck
	} else {
		rawNew.LivenessCheck = canonicalizeNewVersionLivenessCheck(c, rawDesired.LivenessCheck, rawNew.LivenessCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NobuildFilesRegex) && dcl.IsEmptyValueIndirect(rawDesired.NobuildFilesRegex) {
		rawNew.NobuildFilesRegex = rawDesired.NobuildFilesRegex
	} else {
		if dcl.StringCanonicalize(rawDesired.NobuildFilesRegex, rawNew.NobuildFilesRegex) {
			rawNew.NobuildFilesRegex = rawDesired.NobuildFilesRegex
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionUrl) && dcl.IsEmptyValueIndirect(rawDesired.VersionUrl) {
		rawNew.VersionUrl = rawDesired.VersionUrl
	} else {
		if dcl.StringCanonicalize(rawDesired.VersionUrl, rawNew.VersionUrl) {
			rawNew.VersionUrl = rawDesired.VersionUrl
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Entrypoint) && dcl.IsEmptyValueIndirect(rawDesired.Entrypoint) {
		rawNew.Entrypoint = rawDesired.Entrypoint
	} else {
		rawNew.Entrypoint = canonicalizeNewVersionEntrypoint(c, rawDesired.Entrypoint, rawNew.Entrypoint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCAccessConnector) && dcl.IsEmptyValueIndirect(rawDesired.VPCAccessConnector) {
		rawNew.VPCAccessConnector = rawDesired.VPCAccessConnector
	} else {
		rawNew.VPCAccessConnector = canonicalizeNewVersionVPCAccessConnector(c, rawDesired.VPCAccessConnector, rawNew.VPCAccessConnector)
	}

	rawNew.App = rawDesired.App

	rawNew.Service = rawDesired.Service

	return rawNew, nil
}

func canonicalizeVersionAutomaticScaling(des, initial *VersionAutomaticScaling, opts ...dcl.ApplyOption) *VersionAutomaticScaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CoolDownPeriod, initial.CoolDownPeriod) || dcl.IsZeroValue(des.CoolDownPeriod) {
		des.CoolDownPeriod = initial.CoolDownPeriod
	}
	des.CpuUtilization = canonicalizeVersionAutomaticScalingCpuUtilization(des.CpuUtilization, initial.CpuUtilization, opts...)
	if dcl.IsZeroValue(des.MaxConcurrentRequests) {
		des.MaxConcurrentRequests = initial.MaxConcurrentRequests
	}
	if dcl.IsZeroValue(des.MaxIdleInstances) {
		des.MaxIdleInstances = initial.MaxIdleInstances
	}
	if dcl.IsZeroValue(des.MaxTotalInstances) {
		des.MaxTotalInstances = initial.MaxTotalInstances
	}
	if dcl.StringCanonicalize(des.MaxPendingLatency, initial.MaxPendingLatency) || dcl.IsZeroValue(des.MaxPendingLatency) {
		des.MaxPendingLatency = initial.MaxPendingLatency
	}
	if dcl.IsZeroValue(des.MinIdleInstances) {
		des.MinIdleInstances = initial.MinIdleInstances
	}
	if dcl.IsZeroValue(des.MinTotalInstances) {
		des.MinTotalInstances = initial.MinTotalInstances
	}
	if dcl.StringCanonicalize(des.MinPendingLatency, initial.MinPendingLatency) || dcl.IsZeroValue(des.MinPendingLatency) {
		des.MinPendingLatency = initial.MinPendingLatency
	}
	des.RequestUtilization = canonicalizeVersionAutomaticScalingRequestUtilization(des.RequestUtilization, initial.RequestUtilization, opts...)
	des.DiskUtilization = canonicalizeVersionAutomaticScalingDiskUtilization(des.DiskUtilization, initial.DiskUtilization, opts...)
	des.NetworkUtilization = canonicalizeVersionAutomaticScalingNetworkUtilization(des.NetworkUtilization, initial.NetworkUtilization, opts...)
	des.StandardSchedulerSettings = canonicalizeVersionAutomaticScalingStandardSchedulerSettings(des.StandardSchedulerSettings, initial.StandardSchedulerSettings, opts...)

	return des
}

func canonicalizeNewVersionAutomaticScaling(c *Client, des, nw *VersionAutomaticScaling) *VersionAutomaticScaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CoolDownPeriod, nw.CoolDownPeriod) {
		nw.CoolDownPeriod = des.CoolDownPeriod
	}
	nw.CpuUtilization = canonicalizeNewVersionAutomaticScalingCpuUtilization(c, des.CpuUtilization, nw.CpuUtilization)
	if dcl.IsZeroValue(nw.MaxConcurrentRequests) {
		nw.MaxConcurrentRequests = des.MaxConcurrentRequests
	}
	if dcl.IsZeroValue(nw.MaxIdleInstances) {
		nw.MaxIdleInstances = des.MaxIdleInstances
	}
	if dcl.IsZeroValue(nw.MaxTotalInstances) {
		nw.MaxTotalInstances = des.MaxTotalInstances
	}
	if dcl.StringCanonicalize(des.MaxPendingLatency, nw.MaxPendingLatency) {
		nw.MaxPendingLatency = des.MaxPendingLatency
	}
	if dcl.IsZeroValue(nw.MinIdleInstances) {
		nw.MinIdleInstances = des.MinIdleInstances
	}
	if dcl.IsZeroValue(nw.MinTotalInstances) {
		nw.MinTotalInstances = des.MinTotalInstances
	}
	if dcl.StringCanonicalize(des.MinPendingLatency, nw.MinPendingLatency) {
		nw.MinPendingLatency = des.MinPendingLatency
	}
	nw.RequestUtilization = canonicalizeNewVersionAutomaticScalingRequestUtilization(c, des.RequestUtilization, nw.RequestUtilization)
	nw.DiskUtilization = canonicalizeNewVersionAutomaticScalingDiskUtilization(c, des.DiskUtilization, nw.DiskUtilization)
	nw.NetworkUtilization = canonicalizeNewVersionAutomaticScalingNetworkUtilization(c, des.NetworkUtilization, nw.NetworkUtilization)
	nw.StandardSchedulerSettings = canonicalizeNewVersionAutomaticScalingStandardSchedulerSettings(c, des.StandardSchedulerSettings, nw.StandardSchedulerSettings)

	return nw
}

func canonicalizeNewVersionAutomaticScalingSet(c *Client, des, nw []VersionAutomaticScaling) []VersionAutomaticScaling {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingSlice(c *Client, des, nw []VersionAutomaticScaling) []VersionAutomaticScaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScaling(c, &d, &n))
	}

	return items
}

func canonicalizeVersionAutomaticScalingCpuUtilization(des, initial *VersionAutomaticScalingCpuUtilization, opts ...dcl.ApplyOption) *VersionAutomaticScalingCpuUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AggregationWindowLength, initial.AggregationWindowLength) || dcl.IsZeroValue(des.AggregationWindowLength) {
		des.AggregationWindowLength = initial.AggregationWindowLength
	}
	if dcl.IsZeroValue(des.TargetUtilization) {
		des.TargetUtilization = initial.TargetUtilization
	}

	return des
}

func canonicalizeNewVersionAutomaticScalingCpuUtilization(c *Client, des, nw *VersionAutomaticScalingCpuUtilization) *VersionAutomaticScalingCpuUtilization {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AggregationWindowLength, nw.AggregationWindowLength) {
		nw.AggregationWindowLength = des.AggregationWindowLength
	}
	if dcl.IsZeroValue(nw.TargetUtilization) {
		nw.TargetUtilization = des.TargetUtilization
	}

	return nw
}

func canonicalizeNewVersionAutomaticScalingCpuUtilizationSet(c *Client, des, nw []VersionAutomaticScalingCpuUtilization) []VersionAutomaticScalingCpuUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScalingCpuUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingCpuUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingCpuUtilizationSlice(c *Client, des, nw []VersionAutomaticScalingCpuUtilization) []VersionAutomaticScalingCpuUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScalingCpuUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScalingCpuUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeVersionAutomaticScalingRequestUtilization(des, initial *VersionAutomaticScalingRequestUtilization, opts ...dcl.ApplyOption) *VersionAutomaticScalingRequestUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TargetRequestCountPerSecond) {
		des.TargetRequestCountPerSecond = initial.TargetRequestCountPerSecond
	}
	if dcl.IsZeroValue(des.TargetConcurrentRequests) {
		des.TargetConcurrentRequests = initial.TargetConcurrentRequests
	}

	return des
}

func canonicalizeNewVersionAutomaticScalingRequestUtilization(c *Client, des, nw *VersionAutomaticScalingRequestUtilization) *VersionAutomaticScalingRequestUtilization {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.TargetRequestCountPerSecond) {
		nw.TargetRequestCountPerSecond = des.TargetRequestCountPerSecond
	}
	if dcl.IsZeroValue(nw.TargetConcurrentRequests) {
		nw.TargetConcurrentRequests = des.TargetConcurrentRequests
	}

	return nw
}

func canonicalizeNewVersionAutomaticScalingRequestUtilizationSet(c *Client, des, nw []VersionAutomaticScalingRequestUtilization) []VersionAutomaticScalingRequestUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScalingRequestUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingRequestUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingRequestUtilizationSlice(c *Client, des, nw []VersionAutomaticScalingRequestUtilization) []VersionAutomaticScalingRequestUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScalingRequestUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScalingRequestUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeVersionAutomaticScalingDiskUtilization(des, initial *VersionAutomaticScalingDiskUtilization, opts ...dcl.ApplyOption) *VersionAutomaticScalingDiskUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TargetWriteBytesPerSecond) {
		des.TargetWriteBytesPerSecond = initial.TargetWriteBytesPerSecond
	}
	if dcl.IsZeroValue(des.TargetWriteOpsPerSecond) {
		des.TargetWriteOpsPerSecond = initial.TargetWriteOpsPerSecond
	}
	if dcl.IsZeroValue(des.TargetReadBytesPerSecond) {
		des.TargetReadBytesPerSecond = initial.TargetReadBytesPerSecond
	}
	if dcl.IsZeroValue(des.TargetReadOpsPerSecond) {
		des.TargetReadOpsPerSecond = initial.TargetReadOpsPerSecond
	}

	return des
}

func canonicalizeNewVersionAutomaticScalingDiskUtilization(c *Client, des, nw *VersionAutomaticScalingDiskUtilization) *VersionAutomaticScalingDiskUtilization {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.TargetWriteBytesPerSecond) {
		nw.TargetWriteBytesPerSecond = des.TargetWriteBytesPerSecond
	}
	if dcl.IsZeroValue(nw.TargetWriteOpsPerSecond) {
		nw.TargetWriteOpsPerSecond = des.TargetWriteOpsPerSecond
	}
	if dcl.IsZeroValue(nw.TargetReadBytesPerSecond) {
		nw.TargetReadBytesPerSecond = des.TargetReadBytesPerSecond
	}
	if dcl.IsZeroValue(nw.TargetReadOpsPerSecond) {
		nw.TargetReadOpsPerSecond = des.TargetReadOpsPerSecond
	}

	return nw
}

func canonicalizeNewVersionAutomaticScalingDiskUtilizationSet(c *Client, des, nw []VersionAutomaticScalingDiskUtilization) []VersionAutomaticScalingDiskUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScalingDiskUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingDiskUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingDiskUtilizationSlice(c *Client, des, nw []VersionAutomaticScalingDiskUtilization) []VersionAutomaticScalingDiskUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScalingDiskUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScalingDiskUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeVersionAutomaticScalingNetworkUtilization(des, initial *VersionAutomaticScalingNetworkUtilization, opts ...dcl.ApplyOption) *VersionAutomaticScalingNetworkUtilization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TargetSentBytesPerSecond) {
		des.TargetSentBytesPerSecond = initial.TargetSentBytesPerSecond
	}
	if dcl.IsZeroValue(des.TargetSentPacketsPerSecond) {
		des.TargetSentPacketsPerSecond = initial.TargetSentPacketsPerSecond
	}
	if dcl.IsZeroValue(des.TargetReceivedBytesPerSecond) {
		des.TargetReceivedBytesPerSecond = initial.TargetReceivedBytesPerSecond
	}
	if dcl.IsZeroValue(des.TargetReceivedPacketsPerSecond) {
		des.TargetReceivedPacketsPerSecond = initial.TargetReceivedPacketsPerSecond
	}

	return des
}

func canonicalizeNewVersionAutomaticScalingNetworkUtilization(c *Client, des, nw *VersionAutomaticScalingNetworkUtilization) *VersionAutomaticScalingNetworkUtilization {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.TargetSentBytesPerSecond) {
		nw.TargetSentBytesPerSecond = des.TargetSentBytesPerSecond
	}
	if dcl.IsZeroValue(nw.TargetSentPacketsPerSecond) {
		nw.TargetSentPacketsPerSecond = des.TargetSentPacketsPerSecond
	}
	if dcl.IsZeroValue(nw.TargetReceivedBytesPerSecond) {
		nw.TargetReceivedBytesPerSecond = des.TargetReceivedBytesPerSecond
	}
	if dcl.IsZeroValue(nw.TargetReceivedPacketsPerSecond) {
		nw.TargetReceivedPacketsPerSecond = des.TargetReceivedPacketsPerSecond
	}

	return nw
}

func canonicalizeNewVersionAutomaticScalingNetworkUtilizationSet(c *Client, des, nw []VersionAutomaticScalingNetworkUtilization) []VersionAutomaticScalingNetworkUtilization {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScalingNetworkUtilization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingNetworkUtilizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingNetworkUtilizationSlice(c *Client, des, nw []VersionAutomaticScalingNetworkUtilization) []VersionAutomaticScalingNetworkUtilization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScalingNetworkUtilization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScalingNetworkUtilization(c, &d, &n))
	}

	return items
}

func canonicalizeVersionAutomaticScalingStandardSchedulerSettings(des, initial *VersionAutomaticScalingStandardSchedulerSettings, opts ...dcl.ApplyOption) *VersionAutomaticScalingStandardSchedulerSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TargetCpuUtilization) {
		des.TargetCpuUtilization = initial.TargetCpuUtilization
	}
	if dcl.IsZeroValue(des.TargetThroughputUtilization) {
		des.TargetThroughputUtilization = initial.TargetThroughputUtilization
	}
	if dcl.IsZeroValue(des.MinInstances) {
		des.MinInstances = initial.MinInstances
	}
	if dcl.IsZeroValue(des.MaxInstances) {
		des.MaxInstances = initial.MaxInstances
	}

	return des
}

func canonicalizeNewVersionAutomaticScalingStandardSchedulerSettings(c *Client, des, nw *VersionAutomaticScalingStandardSchedulerSettings) *VersionAutomaticScalingStandardSchedulerSettings {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.TargetCpuUtilization) {
		nw.TargetCpuUtilization = des.TargetCpuUtilization
	}
	if dcl.IsZeroValue(nw.TargetThroughputUtilization) {
		nw.TargetThroughputUtilization = des.TargetThroughputUtilization
	}
	if dcl.IsZeroValue(nw.MinInstances) {
		nw.MinInstances = des.MinInstances
	}
	if dcl.IsZeroValue(nw.MaxInstances) {
		nw.MaxInstances = des.MaxInstances
	}

	return nw
}

func canonicalizeNewVersionAutomaticScalingStandardSchedulerSettingsSet(c *Client, des, nw []VersionAutomaticScalingStandardSchedulerSettings) []VersionAutomaticScalingStandardSchedulerSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionAutomaticScalingStandardSchedulerSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionAutomaticScalingStandardSchedulerSettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionAutomaticScalingStandardSchedulerSettingsSlice(c *Client, des, nw []VersionAutomaticScalingStandardSchedulerSettings) []VersionAutomaticScalingStandardSchedulerSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionAutomaticScalingStandardSchedulerSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionAutomaticScalingStandardSchedulerSettings(c, &d, &n))
	}

	return items
}

func canonicalizeVersionBasicScaling(des, initial *VersionBasicScaling, opts ...dcl.ApplyOption) *VersionBasicScaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.IdleTimeout, initial.IdleTimeout) || dcl.IsZeroValue(des.IdleTimeout) {
		des.IdleTimeout = initial.IdleTimeout
	}
	if dcl.IsZeroValue(des.MaxInstances) {
		des.MaxInstances = initial.MaxInstances
	}

	return des
}

func canonicalizeNewVersionBasicScaling(c *Client, des, nw *VersionBasicScaling) *VersionBasicScaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IdleTimeout, nw.IdleTimeout) {
		nw.IdleTimeout = des.IdleTimeout
	}
	if dcl.IsZeroValue(nw.MaxInstances) {
		nw.MaxInstances = des.MaxInstances
	}

	return nw
}

func canonicalizeNewVersionBasicScalingSet(c *Client, des, nw []VersionBasicScaling) []VersionBasicScaling {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionBasicScaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionBasicScalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionBasicScalingSlice(c *Client, des, nw []VersionBasicScaling) []VersionBasicScaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionBasicScaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionBasicScaling(c, &d, &n))
	}

	return items
}

func canonicalizeVersionManualScaling(des, initial *VersionManualScaling, opts ...dcl.ApplyOption) *VersionManualScaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Instances) {
		des.Instances = initial.Instances
	}

	return des
}

func canonicalizeNewVersionManualScaling(c *Client, des, nw *VersionManualScaling) *VersionManualScaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Instances) {
		nw.Instances = des.Instances
	}

	return nw
}

func canonicalizeNewVersionManualScalingSet(c *Client, des, nw []VersionManualScaling) []VersionManualScaling {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionManualScaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionManualScalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionManualScalingSlice(c *Client, des, nw []VersionManualScaling) []VersionManualScaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionManualScaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionManualScaling(c, &d, &n))
	}

	return items
}

func canonicalizeVersionNetwork(des, initial *VersionNetwork, opts ...dcl.ApplyOption) *VersionNetwork {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ForwardedPorts) {
		des.ForwardedPorts = initial.ForwardedPorts
	}
	if dcl.StringCanonicalize(des.InstanceTag, initial.InstanceTag) || dcl.IsZeroValue(des.InstanceTag) {
		des.InstanceTag = initial.InstanceTag
	}
	if dcl.NameToSelfLink(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.NameToSelfLink(des.SubnetworkName, initial.SubnetworkName) || dcl.IsZeroValue(des.SubnetworkName) {
		des.SubnetworkName = initial.SubnetworkName
	}
	if dcl.BoolCanonicalize(des.SessionAffinity, initial.SessionAffinity) || dcl.IsZeroValue(des.SessionAffinity) {
		des.SessionAffinity = initial.SessionAffinity
	}

	return des
}

func canonicalizeNewVersionNetwork(c *Client, des, nw *VersionNetwork) *VersionNetwork {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ForwardedPorts) {
		nw.ForwardedPorts = des.ForwardedPorts
	}
	if dcl.StringCanonicalize(des.InstanceTag, nw.InstanceTag) {
		nw.InstanceTag = des.InstanceTag
	}
	if dcl.NameToSelfLink(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.SubnetworkName, nw.SubnetworkName) {
		nw.SubnetworkName = des.SubnetworkName
	}
	if dcl.BoolCanonicalize(des.SessionAffinity, nw.SessionAffinity) {
		nw.SessionAffinity = des.SessionAffinity
	}

	return nw
}

func canonicalizeNewVersionNetworkSet(c *Client, des, nw []VersionNetwork) []VersionNetwork {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionNetwork
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionNetworkNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionNetworkSlice(c *Client, des, nw []VersionNetwork) []VersionNetwork {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionNetwork
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionNetwork(c, &d, &n))
	}

	return items
}

func canonicalizeVersionResources(des, initial *VersionResources, opts ...dcl.ApplyOption) *VersionResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Cpu) {
		des.Cpu = initial.Cpu
	}
	if dcl.IsZeroValue(des.DiskGb) {
		des.DiskGb = initial.DiskGb
	}
	if dcl.IsZeroValue(des.MemoryGb) {
		des.MemoryGb = initial.MemoryGb
	}
	if dcl.IsZeroValue(des.Volumes) {
		des.Volumes = initial.Volumes
	}

	return des
}

func canonicalizeNewVersionResources(c *Client, des, nw *VersionResources) *VersionResources {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Cpu) {
		nw.Cpu = des.Cpu
	}
	if dcl.IsZeroValue(nw.DiskGb) {
		nw.DiskGb = des.DiskGb
	}
	if dcl.IsZeroValue(nw.MemoryGb) {
		nw.MemoryGb = des.MemoryGb
	}
	nw.Volumes = canonicalizeNewVersionResourcesVolumesSlice(c, des.Volumes, nw.Volumes)

	return nw
}

func canonicalizeNewVersionResourcesSet(c *Client, des, nw []VersionResources) []VersionResources {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionResourcesSlice(c *Client, des, nw []VersionResources) []VersionResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionResources(c, &d, &n))
	}

	return items
}

func canonicalizeVersionResourcesVolumes(des, initial *VersionResourcesVolumes, opts ...dcl.ApplyOption) *VersionResourcesVolumes {
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
	if dcl.StringCanonicalize(des.VolumeType, initial.VolumeType) || dcl.IsZeroValue(des.VolumeType) {
		des.VolumeType = initial.VolumeType
	}
	if dcl.IsZeroValue(des.SizeGb) {
		des.SizeGb = initial.SizeGb
	}

	return des
}

func canonicalizeNewVersionResourcesVolumes(c *Client, des, nw *VersionResourcesVolumes) *VersionResourcesVolumes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.VolumeType, nw.VolumeType) {
		nw.VolumeType = des.VolumeType
	}
	if dcl.IsZeroValue(nw.SizeGb) {
		nw.SizeGb = des.SizeGb
	}

	return nw
}

func canonicalizeNewVersionResourcesVolumesSet(c *Client, des, nw []VersionResourcesVolumes) []VersionResourcesVolumes {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionResourcesVolumes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionResourcesVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionResourcesVolumesSlice(c *Client, des, nw []VersionResourcesVolumes) []VersionResourcesVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionResourcesVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionResourcesVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeVersionHandlers(des, initial *VersionHandlers, opts ...dcl.ApplyOption) *VersionHandlers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.UrlRegex, initial.UrlRegex) || dcl.IsZeroValue(des.UrlRegex) {
		des.UrlRegex = initial.UrlRegex
	}
	des.StaticFiles = canonicalizeVersionHandlersStaticFiles(des.StaticFiles, initial.StaticFiles, opts...)
	des.Script = canonicalizeVersionHandlersScript(des.Script, initial.Script, opts...)
	des.ApiEndpoint = canonicalizeVersionHandlersApiEndpoint(des.ApiEndpoint, initial.ApiEndpoint, opts...)
	if dcl.IsZeroValue(des.SecurityLevel) {
		des.SecurityLevel = initial.SecurityLevel
	}
	if dcl.IsZeroValue(des.Login) {
		des.Login = initial.Login
	}
	if dcl.IsZeroValue(des.AuthFailAction) {
		des.AuthFailAction = initial.AuthFailAction
	}
	if dcl.IsZeroValue(des.RedirectHttpResponseCode) {
		des.RedirectHttpResponseCode = initial.RedirectHttpResponseCode
	}

	return des
}

func canonicalizeNewVersionHandlers(c *Client, des, nw *VersionHandlers) *VersionHandlers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.UrlRegex, nw.UrlRegex) {
		nw.UrlRegex = des.UrlRegex
	}
	nw.StaticFiles = canonicalizeNewVersionHandlersStaticFiles(c, des.StaticFiles, nw.StaticFiles)
	nw.Script = canonicalizeNewVersionHandlersScript(c, des.Script, nw.Script)
	nw.ApiEndpoint = canonicalizeNewVersionHandlersApiEndpoint(c, des.ApiEndpoint, nw.ApiEndpoint)
	if dcl.IsZeroValue(nw.SecurityLevel) {
		nw.SecurityLevel = des.SecurityLevel
	}
	if dcl.IsZeroValue(nw.Login) {
		nw.Login = des.Login
	}
	if dcl.IsZeroValue(nw.AuthFailAction) {
		nw.AuthFailAction = des.AuthFailAction
	}
	if dcl.IsZeroValue(nw.RedirectHttpResponseCode) {
		nw.RedirectHttpResponseCode = des.RedirectHttpResponseCode
	}

	return nw
}

func canonicalizeNewVersionHandlersSet(c *Client, des, nw []VersionHandlers) []VersionHandlers {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionHandlers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionHandlersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionHandlersSlice(c *Client, des, nw []VersionHandlers) []VersionHandlers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionHandlers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionHandlers(c, &d, &n))
	}

	return items
}

func canonicalizeVersionHandlersStaticFiles(des, initial *VersionHandlersStaticFiles, opts ...dcl.ApplyOption) *VersionHandlersStaticFiles {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.StringCanonicalize(des.UploadPathRegex, initial.UploadPathRegex) || dcl.IsZeroValue(des.UploadPathRegex) {
		des.UploadPathRegex = initial.UploadPathRegex
	}
	if dcl.IsZeroValue(des.HttpHeaders) {
		des.HttpHeaders = initial.HttpHeaders
	}
	if dcl.StringCanonicalize(des.MimeType, initial.MimeType) || dcl.IsZeroValue(des.MimeType) {
		des.MimeType = initial.MimeType
	}
	if dcl.StringCanonicalize(des.Expiration, initial.Expiration) || dcl.IsZeroValue(des.Expiration) {
		des.Expiration = initial.Expiration
	}
	if dcl.BoolCanonicalize(des.RequireMatchingFile, initial.RequireMatchingFile) || dcl.IsZeroValue(des.RequireMatchingFile) {
		des.RequireMatchingFile = initial.RequireMatchingFile
	}
	if dcl.BoolCanonicalize(des.ApplicationReadable, initial.ApplicationReadable) || dcl.IsZeroValue(des.ApplicationReadable) {
		des.ApplicationReadable = initial.ApplicationReadable
	}

	return des
}

func canonicalizeNewVersionHandlersStaticFiles(c *Client, des, nw *VersionHandlersStaticFiles) *VersionHandlersStaticFiles {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.UploadPathRegex, nw.UploadPathRegex) {
		nw.UploadPathRegex = des.UploadPathRegex
	}
	if dcl.IsZeroValue(nw.HttpHeaders) {
		nw.HttpHeaders = des.HttpHeaders
	}
	if dcl.StringCanonicalize(des.MimeType, nw.MimeType) {
		nw.MimeType = des.MimeType
	}
	if dcl.StringCanonicalize(des.Expiration, nw.Expiration) {
		nw.Expiration = des.Expiration
	}
	if dcl.BoolCanonicalize(des.RequireMatchingFile, nw.RequireMatchingFile) {
		nw.RequireMatchingFile = des.RequireMatchingFile
	}
	if dcl.BoolCanonicalize(des.ApplicationReadable, nw.ApplicationReadable) {
		nw.ApplicationReadable = des.ApplicationReadable
	}

	return nw
}

func canonicalizeNewVersionHandlersStaticFilesSet(c *Client, des, nw []VersionHandlersStaticFiles) []VersionHandlersStaticFiles {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionHandlersStaticFiles
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionHandlersStaticFilesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionHandlersStaticFilesSlice(c *Client, des, nw []VersionHandlersStaticFiles) []VersionHandlersStaticFiles {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionHandlersStaticFiles
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionHandlersStaticFiles(c, &d, &n))
	}

	return items
}

func canonicalizeVersionHandlersScript(des, initial *VersionHandlersScript, opts ...dcl.ApplyOption) *VersionHandlersScript {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ScriptPath, initial.ScriptPath) || dcl.IsZeroValue(des.ScriptPath) {
		des.ScriptPath = initial.ScriptPath
	}

	return des
}

func canonicalizeNewVersionHandlersScript(c *Client, des, nw *VersionHandlersScript) *VersionHandlersScript {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ScriptPath, nw.ScriptPath) {
		nw.ScriptPath = des.ScriptPath
	}

	return nw
}

func canonicalizeNewVersionHandlersScriptSet(c *Client, des, nw []VersionHandlersScript) []VersionHandlersScript {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionHandlersScript
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionHandlersScriptNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionHandlersScriptSlice(c *Client, des, nw []VersionHandlersScript) []VersionHandlersScript {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionHandlersScript
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionHandlersScript(c, &d, &n))
	}

	return items
}

func canonicalizeVersionHandlersApiEndpoint(des, initial *VersionHandlersApiEndpoint, opts ...dcl.ApplyOption) *VersionHandlersApiEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ScriptPath, initial.ScriptPath) || dcl.IsZeroValue(des.ScriptPath) {
		des.ScriptPath = initial.ScriptPath
	}

	return des
}

func canonicalizeNewVersionHandlersApiEndpoint(c *Client, des, nw *VersionHandlersApiEndpoint) *VersionHandlersApiEndpoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ScriptPath, nw.ScriptPath) {
		nw.ScriptPath = des.ScriptPath
	}

	return nw
}

func canonicalizeNewVersionHandlersApiEndpointSet(c *Client, des, nw []VersionHandlersApiEndpoint) []VersionHandlersApiEndpoint {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionHandlersApiEndpoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionHandlersApiEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionHandlersApiEndpointSlice(c *Client, des, nw []VersionHandlersApiEndpoint) []VersionHandlersApiEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionHandlersApiEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionHandlersApiEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeVersionErrorHandlers(des, initial *VersionErrorHandlers, opts ...dcl.ApplyOption) *VersionErrorHandlers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ErrorCode) {
		des.ErrorCode = initial.ErrorCode
	}
	if dcl.StringCanonicalize(des.StaticFile, initial.StaticFile) || dcl.IsZeroValue(des.StaticFile) {
		des.StaticFile = initial.StaticFile
	}
	if dcl.StringCanonicalize(des.MimeType, initial.MimeType) || dcl.IsZeroValue(des.MimeType) {
		des.MimeType = initial.MimeType
	}

	return des
}

func canonicalizeNewVersionErrorHandlers(c *Client, des, nw *VersionErrorHandlers) *VersionErrorHandlers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ErrorCode) {
		nw.ErrorCode = des.ErrorCode
	}
	if dcl.StringCanonicalize(des.StaticFile, nw.StaticFile) {
		nw.StaticFile = des.StaticFile
	}
	if dcl.StringCanonicalize(des.MimeType, nw.MimeType) {
		nw.MimeType = des.MimeType
	}

	return nw
}

func canonicalizeNewVersionErrorHandlersSet(c *Client, des, nw []VersionErrorHandlers) []VersionErrorHandlers {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionErrorHandlers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionErrorHandlersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionErrorHandlersSlice(c *Client, des, nw []VersionErrorHandlers) []VersionErrorHandlers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionErrorHandlers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionErrorHandlers(c, &d, &n))
	}

	return items
}

func canonicalizeVersionLibraries(des, initial *VersionLibraries, opts ...dcl.ApplyOption) *VersionLibraries {
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

	return des
}

func canonicalizeNewVersionLibraries(c *Client, des, nw *VersionLibraries) *VersionLibraries {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewVersionLibrariesSet(c *Client, des, nw []VersionLibraries) []VersionLibraries {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionLibraries
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionLibrariesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionLibrariesSlice(c *Client, des, nw []VersionLibraries) []VersionLibraries {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionLibraries
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionLibraries(c, &d, &n))
	}

	return items
}

func canonicalizeVersionApiConfig(des, initial *VersionApiConfig, opts ...dcl.ApplyOption) *VersionApiConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AuthFailAction) {
		des.AuthFailAction = initial.AuthFailAction
	}
	if dcl.IsZeroValue(des.Login) {
		des.Login = initial.Login
	}
	if dcl.StringCanonicalize(des.Script, initial.Script) || dcl.IsZeroValue(des.Script) {
		des.Script = initial.Script
	}
	if dcl.IsZeroValue(des.SecurityLevel) {
		des.SecurityLevel = initial.SecurityLevel
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewVersionApiConfig(c *Client, des, nw *VersionApiConfig) *VersionApiConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AuthFailAction) {
		nw.AuthFailAction = des.AuthFailAction
	}
	if dcl.IsZeroValue(nw.Login) {
		nw.Login = des.Login
	}
	if dcl.StringCanonicalize(des.Script, nw.Script) {
		nw.Script = des.Script
	}
	if dcl.IsZeroValue(nw.SecurityLevel) {
		nw.SecurityLevel = des.SecurityLevel
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewVersionApiConfigSet(c *Client, des, nw []VersionApiConfig) []VersionApiConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionApiConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionApiConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionApiConfigSlice(c *Client, des, nw []VersionApiConfig) []VersionApiConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionApiConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionApiConfig(c, &d, &n))
	}

	return items
}

func canonicalizeVersionDeployment(des, initial *VersionDeployment, opts ...dcl.ApplyOption) *VersionDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Files) {
		des.Files = initial.Files
	}
	des.Container = canonicalizeVersionDeploymentContainer(des.Container, initial.Container, opts...)
	des.Zip = canonicalizeVersionDeploymentZip(des.Zip, initial.Zip, opts...)
	des.CloudBuildOptions = canonicalizeVersionDeploymentCloudBuildOptions(des.CloudBuildOptions, initial.CloudBuildOptions, opts...)

	return des
}

func canonicalizeNewVersionDeployment(c *Client, des, nw *VersionDeployment) *VersionDeployment {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Files) {
		nw.Files = des.Files
	}
	nw.Container = canonicalizeNewVersionDeploymentContainer(c, des.Container, nw.Container)
	nw.Zip = canonicalizeNewVersionDeploymentZip(c, des.Zip, nw.Zip)
	nw.CloudBuildOptions = canonicalizeNewVersionDeploymentCloudBuildOptions(c, des.CloudBuildOptions, nw.CloudBuildOptions)

	return nw
}

func canonicalizeNewVersionDeploymentSet(c *Client, des, nw []VersionDeployment) []VersionDeployment {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionDeployment
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionDeploymentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionDeploymentSlice(c *Client, des, nw []VersionDeployment) []VersionDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeVersionDeploymentFiles(des, initial *VersionDeploymentFiles, opts ...dcl.ApplyOption) *VersionDeploymentFiles {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SourceUrl, initial.SourceUrl) || dcl.IsZeroValue(des.SourceUrl) {
		des.SourceUrl = initial.SourceUrl
	}
	if dcl.StringCanonicalize(des.Sha1Sum, initial.Sha1Sum) || dcl.IsZeroValue(des.Sha1Sum) {
		des.Sha1Sum = initial.Sha1Sum
	}
	if dcl.StringCanonicalize(des.MimeType, initial.MimeType) || dcl.IsZeroValue(des.MimeType) {
		des.MimeType = initial.MimeType
	}

	return des
}

func canonicalizeNewVersionDeploymentFiles(c *Client, des, nw *VersionDeploymentFiles) *VersionDeploymentFiles {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SourceUrl, nw.SourceUrl) {
		nw.SourceUrl = des.SourceUrl
	}
	if dcl.StringCanonicalize(des.Sha1Sum, nw.Sha1Sum) {
		nw.Sha1Sum = des.Sha1Sum
	}
	if dcl.StringCanonicalize(des.MimeType, nw.MimeType) {
		nw.MimeType = des.MimeType
	}

	return nw
}

func canonicalizeNewVersionDeploymentFilesSet(c *Client, des, nw []VersionDeploymentFiles) []VersionDeploymentFiles {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionDeploymentFiles
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionDeploymentFilesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionDeploymentFilesSlice(c *Client, des, nw []VersionDeploymentFiles) []VersionDeploymentFiles {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionDeploymentFiles
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionDeploymentFiles(c, &d, &n))
	}

	return items
}

func canonicalizeVersionDeploymentContainer(des, initial *VersionDeploymentContainer, opts ...dcl.ApplyOption) *VersionDeploymentContainer {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		des.Image = initial.Image
	}

	return des
}

func canonicalizeNewVersionDeploymentContainer(c *Client, des, nw *VersionDeploymentContainer) *VersionDeploymentContainer {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Image, nw.Image) {
		nw.Image = des.Image
	}

	return nw
}

func canonicalizeNewVersionDeploymentContainerSet(c *Client, des, nw []VersionDeploymentContainer) []VersionDeploymentContainer {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionDeploymentContainer
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionDeploymentContainerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionDeploymentContainerSlice(c *Client, des, nw []VersionDeploymentContainer) []VersionDeploymentContainer {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionDeploymentContainer
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionDeploymentContainer(c, &d, &n))
	}

	return items
}

func canonicalizeVersionDeploymentZip(des, initial *VersionDeploymentZip, opts ...dcl.ApplyOption) *VersionDeploymentZip {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SourceUrl, initial.SourceUrl) || dcl.IsZeroValue(des.SourceUrl) {
		des.SourceUrl = initial.SourceUrl
	}
	if dcl.IsZeroValue(des.FilesCount) {
		des.FilesCount = initial.FilesCount
	}

	return des
}

func canonicalizeNewVersionDeploymentZip(c *Client, des, nw *VersionDeploymentZip) *VersionDeploymentZip {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SourceUrl, nw.SourceUrl) {
		nw.SourceUrl = des.SourceUrl
	}
	if dcl.IsZeroValue(nw.FilesCount) {
		nw.FilesCount = des.FilesCount
	}

	return nw
}

func canonicalizeNewVersionDeploymentZipSet(c *Client, des, nw []VersionDeploymentZip) []VersionDeploymentZip {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionDeploymentZip
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionDeploymentZipNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionDeploymentZipSlice(c *Client, des, nw []VersionDeploymentZip) []VersionDeploymentZip {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionDeploymentZip
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionDeploymentZip(c, &d, &n))
	}

	return items
}

func canonicalizeVersionDeploymentCloudBuildOptions(des, initial *VersionDeploymentCloudBuildOptions, opts ...dcl.ApplyOption) *VersionDeploymentCloudBuildOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AppYamlPath, initial.AppYamlPath) || dcl.IsZeroValue(des.AppYamlPath) {
		des.AppYamlPath = initial.AppYamlPath
	}
	if dcl.StringCanonicalize(des.CloudBuildTimeout, initial.CloudBuildTimeout) || dcl.IsZeroValue(des.CloudBuildTimeout) {
		des.CloudBuildTimeout = initial.CloudBuildTimeout
	}

	return des
}

func canonicalizeNewVersionDeploymentCloudBuildOptions(c *Client, des, nw *VersionDeploymentCloudBuildOptions) *VersionDeploymentCloudBuildOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AppYamlPath, nw.AppYamlPath) {
		nw.AppYamlPath = des.AppYamlPath
	}
	if dcl.StringCanonicalize(des.CloudBuildTimeout, nw.CloudBuildTimeout) {
		nw.CloudBuildTimeout = des.CloudBuildTimeout
	}

	return nw
}

func canonicalizeNewVersionDeploymentCloudBuildOptionsSet(c *Client, des, nw []VersionDeploymentCloudBuildOptions) []VersionDeploymentCloudBuildOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionDeploymentCloudBuildOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionDeploymentCloudBuildOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionDeploymentCloudBuildOptionsSlice(c *Client, des, nw []VersionDeploymentCloudBuildOptions) []VersionDeploymentCloudBuildOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionDeploymentCloudBuildOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionDeploymentCloudBuildOptions(c, &d, &n))
	}

	return items
}

func canonicalizeVersionHealthCheck(des, initial *VersionHealthCheck, opts ...dcl.ApplyOption) *VersionHealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DisableHealthCheck, initial.DisableHealthCheck) || dcl.IsZeroValue(des.DisableHealthCheck) {
		des.DisableHealthCheck = initial.DisableHealthCheck
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.IsZeroValue(des.HealthyThreshold) {
		des.HealthyThreshold = initial.HealthyThreshold
	}
	if dcl.IsZeroValue(des.UnhealthyThreshold) {
		des.UnhealthyThreshold = initial.UnhealthyThreshold
	}
	if dcl.IsZeroValue(des.RestartThreshold) {
		des.RestartThreshold = initial.RestartThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, initial.CheckInterval) || dcl.IsZeroValue(des.CheckInterval) {
		des.CheckInterval = initial.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		des.Timeout = initial.Timeout
	}

	return des
}

func canonicalizeNewVersionHealthCheck(c *Client, des, nw *VersionHealthCheck) *VersionHealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DisableHealthCheck, nw.DisableHealthCheck) {
		nw.DisableHealthCheck = des.DisableHealthCheck
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.IsZeroValue(nw.HealthyThreshold) {
		nw.HealthyThreshold = des.HealthyThreshold
	}
	if dcl.IsZeroValue(nw.UnhealthyThreshold) {
		nw.UnhealthyThreshold = des.UnhealthyThreshold
	}
	if dcl.IsZeroValue(nw.RestartThreshold) {
		nw.RestartThreshold = des.RestartThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, nw.CheckInterval) {
		nw.CheckInterval = des.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}

	return nw
}

func canonicalizeNewVersionHealthCheckSet(c *Client, des, nw []VersionHealthCheck) []VersionHealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionHealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionHealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionHealthCheckSlice(c *Client, des, nw []VersionHealthCheck) []VersionHealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionHealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionHealthCheck(c, &d, &n))
	}

	return items
}

func canonicalizeVersionReadinessCheck(des, initial *VersionReadinessCheck, opts ...dcl.ApplyOption) *VersionReadinessCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.IsZeroValue(des.FailureThreshold) {
		des.FailureThreshold = initial.FailureThreshold
	}
	if dcl.IsZeroValue(des.SuccessThreshold) {
		des.SuccessThreshold = initial.SuccessThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, initial.CheckInterval) || dcl.IsZeroValue(des.CheckInterval) {
		des.CheckInterval = initial.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		des.Timeout = initial.Timeout
	}
	if dcl.StringCanonicalize(des.AppStartTimeout, initial.AppStartTimeout) || dcl.IsZeroValue(des.AppStartTimeout) {
		des.AppStartTimeout = initial.AppStartTimeout
	}

	return des
}

func canonicalizeNewVersionReadinessCheck(c *Client, des, nw *VersionReadinessCheck) *VersionReadinessCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.IsZeroValue(nw.FailureThreshold) {
		nw.FailureThreshold = des.FailureThreshold
	}
	if dcl.IsZeroValue(nw.SuccessThreshold) {
		nw.SuccessThreshold = des.SuccessThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, nw.CheckInterval) {
		nw.CheckInterval = des.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.AppStartTimeout, nw.AppStartTimeout) {
		nw.AppStartTimeout = des.AppStartTimeout
	}

	return nw
}

func canonicalizeNewVersionReadinessCheckSet(c *Client, des, nw []VersionReadinessCheck) []VersionReadinessCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionReadinessCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionReadinessCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionReadinessCheckSlice(c *Client, des, nw []VersionReadinessCheck) []VersionReadinessCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionReadinessCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionReadinessCheck(c, &d, &n))
	}

	return items
}

func canonicalizeVersionLivenessCheck(des, initial *VersionLivenessCheck, opts ...dcl.ApplyOption) *VersionLivenessCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.IsZeroValue(des.FailureThreshold) {
		des.FailureThreshold = initial.FailureThreshold
	}
	if dcl.IsZeroValue(des.SuccessThreshold) {
		des.SuccessThreshold = initial.SuccessThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, initial.CheckInterval) || dcl.IsZeroValue(des.CheckInterval) {
		des.CheckInterval = initial.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		des.Timeout = initial.Timeout
	}
	if dcl.StringCanonicalize(des.InitialDelay, initial.InitialDelay) || dcl.IsZeroValue(des.InitialDelay) {
		des.InitialDelay = initial.InitialDelay
	}

	return des
}

func canonicalizeNewVersionLivenessCheck(c *Client, des, nw *VersionLivenessCheck) *VersionLivenessCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.IsZeroValue(nw.FailureThreshold) {
		nw.FailureThreshold = des.FailureThreshold
	}
	if dcl.IsZeroValue(nw.SuccessThreshold) {
		nw.SuccessThreshold = des.SuccessThreshold
	}
	if dcl.StringCanonicalize(des.CheckInterval, nw.CheckInterval) {
		nw.CheckInterval = des.CheckInterval
	}
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.InitialDelay, nw.InitialDelay) {
		nw.InitialDelay = des.InitialDelay
	}

	return nw
}

func canonicalizeNewVersionLivenessCheckSet(c *Client, des, nw []VersionLivenessCheck) []VersionLivenessCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionLivenessCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionLivenessCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionLivenessCheckSlice(c *Client, des, nw []VersionLivenessCheck) []VersionLivenessCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionLivenessCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionLivenessCheck(c, &d, &n))
	}

	return items
}

func canonicalizeVersionEntrypoint(des, initial *VersionEntrypoint, opts ...dcl.ApplyOption) *VersionEntrypoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Shell, initial.Shell) || dcl.IsZeroValue(des.Shell) {
		des.Shell = initial.Shell
	}

	return des
}

func canonicalizeNewVersionEntrypoint(c *Client, des, nw *VersionEntrypoint) *VersionEntrypoint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Shell, nw.Shell) {
		nw.Shell = des.Shell
	}

	return nw
}

func canonicalizeNewVersionEntrypointSet(c *Client, des, nw []VersionEntrypoint) []VersionEntrypoint {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionEntrypoint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionEntrypointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionEntrypointSlice(c *Client, des, nw []VersionEntrypoint) []VersionEntrypoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionEntrypoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionEntrypoint(c, &d, &n))
	}

	return items
}

func canonicalizeVersionVPCAccessConnector(des, initial *VersionVPCAccessConnector, opts ...dcl.ApplyOption) *VersionVPCAccessConnector {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewVersionVPCAccessConnector(c *Client, des, nw *VersionVPCAccessConnector) *VersionVPCAccessConnector {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewVersionVPCAccessConnectorSet(c *Client, des, nw []VersionVPCAccessConnector) []VersionVPCAccessConnector {
	if des == nil {
		return nw
	}
	var reorderedNew []VersionVPCAccessConnector
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVersionVPCAccessConnectorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVersionVPCAccessConnectorSlice(c *Client, des, nw []VersionVPCAccessConnector) []VersionVPCAccessConnector {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VersionVPCAccessConnector
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVersionVPCAccessConnector(c, &d, &n))
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
func diffVersion(c *Client, desired, actual *Version, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.ConsumerName, actual.ConsumerName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ConsumerName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutomaticScaling, actual.AutomaticScaling, dcl.Info{ObjectFunction: compareVersionAutomaticScalingNewStyle, EmptyObject: EmptyVersionAutomaticScaling, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("AutomaticScaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BasicScaling, actual.BasicScaling, dcl.Info{ObjectFunction: compareVersionBasicScalingNewStyle, EmptyObject: EmptyVersionBasicScaling, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("BasicScaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManualScaling, actual.ManualScaling, dcl.Info{ObjectFunction: compareVersionManualScalingNewStyle, EmptyObject: EmptyVersionManualScaling, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ManualScaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InboundServices, actual.InboundServices, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("InboundServices")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceClass, actual.InstanceClass, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("InstanceClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{ObjectFunction: compareVersionNetworkNewStyle, EmptyObject: EmptyVersionNetwork, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zones, actual.Zones, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Zones")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareVersionResourcesNewStyle, EmptyObject: EmptyVersionResources, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Runtime, actual.Runtime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Runtime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuntimeChannel, actual.RuntimeChannel, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("RuntimeChannel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Threadsafe, actual.Threadsafe, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Threadsafe")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Vm, actual.Vm, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Vm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BetaSettings, actual.BetaSettings, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("BetaSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServingStatus, actual.ServingStatus, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ServingStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedBy, actual.CreatedBy, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("CreatedBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskUsageBytes, actual.DiskUsageBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("DiskUsageBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuntimeApiVersion, actual.RuntimeApiVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("RuntimeApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuntimeMainExecutablePath, actual.RuntimeMainExecutablePath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("RuntimeMainExecutablePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Handlers, actual.Handlers, dcl.Info{ObjectFunction: compareVersionHandlersNewStyle, EmptyObject: EmptyVersionHandlers, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Handlers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ErrorHandlers, actual.ErrorHandlers, dcl.Info{ObjectFunction: compareVersionErrorHandlersNewStyle, EmptyObject: EmptyVersionErrorHandlers, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ErrorHandlers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Libraries, actual.Libraries, dcl.Info{ObjectFunction: compareVersionLibrariesNewStyle, EmptyObject: EmptyVersionLibraries, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Libraries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiConfig, actual.ApiConfig, dcl.Info{ObjectFunction: compareVersionApiConfigNewStyle, EmptyObject: EmptyVersionApiConfig, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ApiConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnvVariables, actual.EnvVariables, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("EnvVariables")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultExpiration, actual.DefaultExpiration, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("DefaultExpiration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deployment, actual.Deployment, dcl.Info{ObjectFunction: compareVersionDeploymentNewStyle, EmptyObject: EmptyVersionDeployment, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthCheck, actual.HealthCheck, dcl.Info{ObjectFunction: compareVersionHealthCheckNewStyle, EmptyObject: EmptyVersionHealthCheck, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("HealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadinessCheck, actual.ReadinessCheck, dcl.Info{ObjectFunction: compareVersionReadinessCheckNewStyle, EmptyObject: EmptyVersionReadinessCheck, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("ReadinessCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LivenessCheck, actual.LivenessCheck, dcl.Info{ObjectFunction: compareVersionLivenessCheckNewStyle, EmptyObject: EmptyVersionLivenessCheck, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("LivenessCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NobuildFilesRegex, actual.NobuildFilesRegex, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("NobuildFilesRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionUrl, actual.VersionUrl, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("VersionUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Entrypoint, actual.Entrypoint, dcl.Info{ObjectFunction: compareVersionEntrypointNewStyle, EmptyObject: EmptyVersionEntrypoint, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("Entrypoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCAccessConnector, actual.VPCAccessConnector, dcl.Info{ObjectFunction: compareVersionVPCAccessConnectorNewStyle, EmptyObject: EmptyVersionVPCAccessConnector, OperationSelector: dcl.TriggersOperation("updateVersionPatchVersionOperation")}, fn.AddNest("VpcAccessConnector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.App, actual.App, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("App")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareVersionAutomaticScalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScaling)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScaling or *VersionAutomaticScaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScaling)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CoolDownPeriod, actual.CoolDownPeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CoolDownPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuUtilization, actual.CpuUtilization, dcl.Info{ObjectFunction: compareVersionAutomaticScalingCpuUtilizationNewStyle, EmptyObject: EmptyVersionAutomaticScalingCpuUtilization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CpuUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxConcurrentRequests, actual.MaxConcurrentRequests, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxConcurrentRequests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxIdleInstances, actual.MaxIdleInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIdleInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxTotalInstances, actual.MaxTotalInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxTotalInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxPendingLatency, actual.MaxPendingLatency, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxPendingLatency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinIdleInstances, actual.MinIdleInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinIdleInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinTotalInstances, actual.MinTotalInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinTotalInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinPendingLatency, actual.MinPendingLatency, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinPendingLatency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestUtilization, actual.RequestUtilization, dcl.Info{ObjectFunction: compareVersionAutomaticScalingRequestUtilizationNewStyle, EmptyObject: EmptyVersionAutomaticScalingRequestUtilization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequestUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskUtilization, actual.DiskUtilization, dcl.Info{ObjectFunction: compareVersionAutomaticScalingDiskUtilizationNewStyle, EmptyObject: EmptyVersionAutomaticScalingDiskUtilization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkUtilization, actual.NetworkUtilization, dcl.Info{ObjectFunction: compareVersionAutomaticScalingNetworkUtilizationNewStyle, EmptyObject: EmptyVersionAutomaticScalingNetworkUtilization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StandardSchedulerSettings, actual.StandardSchedulerSettings, dcl.Info{ObjectFunction: compareVersionAutomaticScalingStandardSchedulerSettingsNewStyle, EmptyObject: EmptyVersionAutomaticScalingStandardSchedulerSettings, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StandardSchedulerSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionAutomaticScalingCpuUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScalingCpuUtilization)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScalingCpuUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingCpuUtilization or *VersionAutomaticScalingCpuUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScalingCpuUtilization)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScalingCpuUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingCpuUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AggregationWindowLength, actual.AggregationWindowLength, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AggregationWindowLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetUtilization, actual.TargetUtilization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionAutomaticScalingRequestUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScalingRequestUtilization)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScalingRequestUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingRequestUtilization or *VersionAutomaticScalingRequestUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScalingRequestUtilization)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScalingRequestUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingRequestUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetRequestCountPerSecond, actual.TargetRequestCountPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetRequestCountPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetConcurrentRequests, actual.TargetConcurrentRequests, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetConcurrentRequests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionAutomaticScalingDiskUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScalingDiskUtilization)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScalingDiskUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingDiskUtilization or *VersionAutomaticScalingDiskUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScalingDiskUtilization)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScalingDiskUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingDiskUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetWriteBytesPerSecond, actual.TargetWriteBytesPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetWriteBytesPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetWriteOpsPerSecond, actual.TargetWriteOpsPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetWriteOpsPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetReadBytesPerSecond, actual.TargetReadBytesPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetReadBytesPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetReadOpsPerSecond, actual.TargetReadOpsPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetReadOpsPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionAutomaticScalingNetworkUtilizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScalingNetworkUtilization)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScalingNetworkUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingNetworkUtilization or *VersionAutomaticScalingNetworkUtilization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScalingNetworkUtilization)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScalingNetworkUtilization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingNetworkUtilization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetSentBytesPerSecond, actual.TargetSentBytesPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetSentBytesPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetSentPacketsPerSecond, actual.TargetSentPacketsPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetSentPacketsPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetReceivedBytesPerSecond, actual.TargetReceivedBytesPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetReceivedBytesPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetReceivedPacketsPerSecond, actual.TargetReceivedPacketsPerSecond, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetReceivedPacketsPerSecond")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionAutomaticScalingStandardSchedulerSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionAutomaticScalingStandardSchedulerSettings)
	if !ok {
		desiredNotPointer, ok := d.(VersionAutomaticScalingStandardSchedulerSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingStandardSchedulerSettings or *VersionAutomaticScalingStandardSchedulerSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionAutomaticScalingStandardSchedulerSettings)
	if !ok {
		actualNotPointer, ok := a.(VersionAutomaticScalingStandardSchedulerSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionAutomaticScalingStandardSchedulerSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetCpuUtilization, actual.TargetCpuUtilization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetCpuUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetThroughputUtilization, actual.TargetThroughputUtilization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetThroughputUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinInstances, actual.MinInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionBasicScalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionBasicScaling)
	if !ok {
		desiredNotPointer, ok := d.(VersionBasicScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionBasicScaling or *VersionBasicScaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionBasicScaling)
	if !ok {
		actualNotPointer, ok := a.(VersionBasicScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionBasicScaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IdleTimeout, actual.IdleTimeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdleTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionManualScalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionManualScaling)
	if !ok {
		desiredNotPointer, ok := d.(VersionManualScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionManualScaling or *VersionManualScaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionManualScaling)
	if !ok {
		actualNotPointer, ok := a.(VersionManualScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionManualScaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Instances, actual.Instances, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Instances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionNetworkNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionNetwork)
	if !ok {
		desiredNotPointer, ok := d.(VersionNetwork)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionNetwork or *VersionNetwork", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionNetwork)
	if !ok {
		actualNotPointer, ok := a.(VersionNetwork)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionNetwork", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ForwardedPorts, actual.ForwardedPorts, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ForwardedPorts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceTag, actual.InstanceTag, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceTag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetworkName, actual.SubnetworkName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetworkName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SessionAffinity, actual.SessionAffinity, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SessionAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionResources)
	if !ok {
		desiredNotPointer, ok := d.(VersionResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionResources or *VersionResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionResources)
	if !ok {
		actualNotPointer, ok := a.(VersionResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Cpu, actual.Cpu, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Cpu")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskGb, actual.DiskGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemoryGb, actual.MemoryGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemoryGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.Info{ObjectFunction: compareVersionResourcesVolumesNewStyle, EmptyObject: EmptyVersionResourcesVolumes, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionResourcesVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionResourcesVolumes)
	if !ok {
		desiredNotPointer, ok := d.(VersionResourcesVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionResourcesVolumes or *VersionResourcesVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionResourcesVolumes)
	if !ok {
		actualNotPointer, ok := a.(VersionResourcesVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionResourcesVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeType, actual.VolumeType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VolumeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SizeGb, actual.SizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionHandlersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionHandlers)
	if !ok {
		desiredNotPointer, ok := d.(VersionHandlers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlers or *VersionHandlers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionHandlers)
	if !ok {
		actualNotPointer, ok := a.(VersionHandlers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UrlRegex, actual.UrlRegex, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UrlRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StaticFiles, actual.StaticFiles, dcl.Info{ObjectFunction: compareVersionHandlersStaticFilesNewStyle, EmptyObject: EmptyVersionHandlersStaticFiles, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StaticFiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Script, actual.Script, dcl.Info{ObjectFunction: compareVersionHandlersScriptNewStyle, EmptyObject: EmptyVersionHandlersScript, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Script")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiEndpoint, actual.ApiEndpoint, dcl.Info{ObjectFunction: compareVersionHandlersApiEndpointNewStyle, EmptyObject: EmptyVersionHandlersApiEndpoint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApiEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityLevel, actual.SecurityLevel, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityLevel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Login, actual.Login, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Login")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthFailAction, actual.AuthFailAction, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthFailAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedirectHttpResponseCode, actual.RedirectHttpResponseCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedirectHttpResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionHandlersStaticFilesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionHandlersStaticFiles)
	if !ok {
		desiredNotPointer, ok := d.(VersionHandlersStaticFiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersStaticFiles or *VersionHandlersStaticFiles", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionHandlersStaticFiles)
	if !ok {
		actualNotPointer, ok := a.(VersionHandlersStaticFiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersStaticFiles", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UploadPathRegex, actual.UploadPathRegex, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UploadPathRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaders, actual.HttpHeaders, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MimeType, actual.MimeType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Expiration, actual.Expiration, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Expiration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireMatchingFile, actual.RequireMatchingFile, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RequireMatchingFile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApplicationReadable, actual.ApplicationReadable, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApplicationReadable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionHandlersScriptNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionHandlersScript)
	if !ok {
		desiredNotPointer, ok := d.(VersionHandlersScript)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersScript or *VersionHandlersScript", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionHandlersScript)
	if !ok {
		actualNotPointer, ok := a.(VersionHandlersScript)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersScript", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ScriptPath, actual.ScriptPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScriptPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionHandlersApiEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionHandlersApiEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(VersionHandlersApiEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersApiEndpoint or *VersionHandlersApiEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionHandlersApiEndpoint)
	if !ok {
		actualNotPointer, ok := a.(VersionHandlersApiEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHandlersApiEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ScriptPath, actual.ScriptPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScriptPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionErrorHandlersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionErrorHandlers)
	if !ok {
		desiredNotPointer, ok := d.(VersionErrorHandlers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionErrorHandlers or *VersionErrorHandlers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionErrorHandlers)
	if !ok {
		actualNotPointer, ok := a.(VersionErrorHandlers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionErrorHandlers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ErrorCode, actual.ErrorCode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ErrorCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StaticFile, actual.StaticFile, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StaticFile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MimeType, actual.MimeType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionLibrariesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionLibraries)
	if !ok {
		desiredNotPointer, ok := d.(VersionLibraries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionLibraries or *VersionLibraries", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionLibraries)
	if !ok {
		actualNotPointer, ok := a.(VersionLibraries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionLibraries", a)
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
	return diffs, nil
}

func compareVersionApiConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionApiConfig)
	if !ok {
		desiredNotPointer, ok := d.(VersionApiConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionApiConfig or *VersionApiConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionApiConfig)
	if !ok {
		actualNotPointer, ok := a.(VersionApiConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionApiConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AuthFailAction, actual.AuthFailAction, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthFailAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Login, actual.Login, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Login")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Script, actual.Script, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Script")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityLevel, actual.SecurityLevel, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityLevel")); len(ds) != 0 || err != nil {
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

func compareVersionDeploymentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionDeployment)
	if !ok {
		desiredNotPointer, ok := d.(VersionDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeployment or *VersionDeployment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionDeployment)
	if !ok {
		actualNotPointer, ok := a.(VersionDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeployment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Container, actual.Container, dcl.Info{ObjectFunction: compareVersionDeploymentContainerNewStyle, EmptyObject: EmptyVersionDeploymentContainer, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Container")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zip, actual.Zip, dcl.Info{ObjectFunction: compareVersionDeploymentZipNewStyle, EmptyObject: EmptyVersionDeploymentZip, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zip")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudBuildOptions, actual.CloudBuildOptions, dcl.Info{ObjectFunction: compareVersionDeploymentCloudBuildOptionsNewStyle, EmptyObject: EmptyVersionDeploymentCloudBuildOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudBuildOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionDeploymentFilesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionDeploymentFiles)
	if !ok {
		desiredNotPointer, ok := d.(VersionDeploymentFiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentFiles or *VersionDeploymentFiles", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionDeploymentFiles)
	if !ok {
		actualNotPointer, ok := a.(VersionDeploymentFiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentFiles", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SourceUrl, actual.SourceUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha1Sum, actual.Sha1Sum, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha1Sum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MimeType, actual.MimeType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionDeploymentContainerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionDeploymentContainer)
	if !ok {
		desiredNotPointer, ok := d.(VersionDeploymentContainer)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentContainer or *VersionDeploymentContainer", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionDeploymentContainer)
	if !ok {
		actualNotPointer, ok := a.(VersionDeploymentContainer)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentContainer", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionDeploymentZipNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionDeploymentZip)
	if !ok {
		desiredNotPointer, ok := d.(VersionDeploymentZip)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentZip or *VersionDeploymentZip", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionDeploymentZip)
	if !ok {
		actualNotPointer, ok := a.(VersionDeploymentZip)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentZip", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SourceUrl, actual.SourceUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FilesCount, actual.FilesCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FilesCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionDeploymentCloudBuildOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionDeploymentCloudBuildOptions)
	if !ok {
		desiredNotPointer, ok := d.(VersionDeploymentCloudBuildOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentCloudBuildOptions or *VersionDeploymentCloudBuildOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionDeploymentCloudBuildOptions)
	if !ok {
		actualNotPointer, ok := a.(VersionDeploymentCloudBuildOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionDeploymentCloudBuildOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AppYamlPath, actual.AppYamlPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AppYamlPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudBuildTimeout, actual.CloudBuildTimeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudBuildTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionHealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionHealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(VersionHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHealthCheck or *VersionHealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionHealthCheck)
	if !ok {
		actualNotPointer, ok := a.(VersionHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionHealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisableHealthCheck, actual.DisableHealthCheck, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisableHealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthyThreshold, actual.HealthyThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnhealthyThreshold, actual.UnhealthyThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UnhealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RestartThreshold, actual.RestartThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RestartThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CheckInterval, actual.CheckInterval, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CheckInterval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionReadinessCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionReadinessCheck)
	if !ok {
		desiredNotPointer, ok := d.(VersionReadinessCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionReadinessCheck or *VersionReadinessCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionReadinessCheck)
	if !ok {
		actualNotPointer, ok := a.(VersionReadinessCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionReadinessCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CheckInterval, actual.CheckInterval, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CheckInterval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppStartTimeout, actual.AppStartTimeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AppStartTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionLivenessCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionLivenessCheck)
	if !ok {
		desiredNotPointer, ok := d.(VersionLivenessCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionLivenessCheck or *VersionLivenessCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionLivenessCheck)
	if !ok {
		actualNotPointer, ok := a.(VersionLivenessCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionLivenessCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CheckInterval, actual.CheckInterval, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CheckInterval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitialDelay, actual.InitialDelay, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InitialDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionEntrypointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionEntrypoint)
	if !ok {
		desiredNotPointer, ok := d.(VersionEntrypoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionEntrypoint or *VersionEntrypoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionEntrypoint)
	if !ok {
		actualNotPointer, ok := a.(VersionEntrypoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionEntrypoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Shell, actual.Shell, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Shell")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareVersionVPCAccessConnectorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VersionVPCAccessConnector)
	if !ok {
		desiredNotPointer, ok := d.(VersionVPCAccessConnector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionVPCAccessConnector or *VersionVPCAccessConnector", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VersionVPCAccessConnector)
	if !ok {
		actualNotPointer, ok := a.(VersionVPCAccessConnector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VersionVPCAccessConnector", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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
func (r *Version) urlNormalized() *Version {
	normalized := dcl.Copy(*r).(Version)
	normalized.ConsumerName = dcl.SelfLinkToName(r.ConsumerName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.InstanceClass = dcl.SelfLinkToName(r.InstanceClass)
	normalized.Runtime = dcl.SelfLinkToName(r.Runtime)
	normalized.RuntimeChannel = dcl.SelfLinkToName(r.RuntimeChannel)
	normalized.Env = dcl.SelfLinkToName(r.Env)
	normalized.CreatedBy = dcl.SelfLinkToName(r.CreatedBy)
	normalized.RuntimeApiVersion = dcl.SelfLinkToName(r.RuntimeApiVersion)
	normalized.RuntimeMainExecutablePath = dcl.SelfLinkToName(r.RuntimeMainExecutablePath)
	normalized.DefaultExpiration = dcl.SelfLinkToName(r.DefaultExpiration)
	normalized.NobuildFilesRegex = dcl.SelfLinkToName(r.NobuildFilesRegex)
	normalized.VersionUrl = dcl.SelfLinkToName(r.VersionUrl)
	normalized.App = dcl.SelfLinkToName(r.App)
	normalized.Service = dcl.SelfLinkToName(r.Service)
	return &normalized
}

func (r *Version) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Service), dcl.ValueOrEmptyString(n.Name)
}

func (r *Version) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Service)
}

func (r *Version) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Service), dcl.ValueOrEmptyString(n.Name)
}

func (r *Version) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "PatchVersion" {
		fields := map[string]interface{}{
			"app":     dcl.ValueOrEmptyString(n.App),
			"service": dcl.ValueOrEmptyString(n.Service),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("apps/{{app}}/services/{{service}}/versions/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Version resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Version) marshal(c *Client) ([]byte, error) {
	m, err := expandVersion(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Version: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalVersion decodes JSON responses into the Version resource schema.
func unmarshalVersion(b []byte, c *Client) (*Version, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapVersion(m, c)
}

func unmarshalMapVersion(m map[string]interface{}, c *Client) (*Version, error) {

	return flattenVersion(c, m), nil
}

// expandVersion expands Version into a JSON request object.
func expandVersion(c *Client, f *Version) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.ConsumerName; !dcl.IsEmptyValueIndirect(v) {
		m["consumerName"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandVersionAutomaticScaling(c, f.AutomaticScaling); err != nil {
		return nil, fmt.Errorf("error expanding AutomaticScaling into automaticScaling: %w", err)
	} else if v != nil {
		m["automaticScaling"] = v
	}
	if v, err := expandVersionBasicScaling(c, f.BasicScaling); err != nil {
		return nil, fmt.Errorf("error expanding BasicScaling into basicScaling: %w", err)
	} else if v != nil {
		m["basicScaling"] = v
	}
	if v, err := expandVersionManualScaling(c, f.ManualScaling); err != nil {
		return nil, fmt.Errorf("error expanding ManualScaling into manualScaling: %w", err)
	} else if v != nil {
		m["manualScaling"] = v
	}
	if v := f.InboundServices; !dcl.IsEmptyValueIndirect(v) {
		m["inboundServices"] = v
	}
	if v := f.InstanceClass; !dcl.IsEmptyValueIndirect(v) {
		m["instanceClass"] = v
	}
	if v, err := expandVersionNetwork(c, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if v != nil {
		m["network"] = v
	}
	if v := f.Zones; !dcl.IsEmptyValueIndirect(v) {
		m["zones"] = v
	}
	if v, err := expandVersionResources(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if v != nil {
		m["resources"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		m["runtime"] = v
	}
	if v := f.RuntimeChannel; !dcl.IsEmptyValueIndirect(v) {
		m["runtimeChannel"] = v
	}
	if v := f.Threadsafe; !dcl.IsEmptyValueIndirect(v) {
		m["threadsafe"] = v
	}
	if v := f.Vm; !dcl.IsEmptyValueIndirect(v) {
		m["vm"] = v
	}
	if v := f.BetaSettings; !dcl.IsEmptyValueIndirect(v) {
		m["betaSettings"] = v
	}
	if v := f.Env; !dcl.IsEmptyValueIndirect(v) {
		m["env"] = v
	}
	if v := f.ServingStatus; !dcl.IsEmptyValueIndirect(v) {
		m["servingStatus"] = v
	}
	if v := f.CreatedBy; !dcl.IsEmptyValueIndirect(v) {
		m["createdBy"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.DiskUsageBytes; !dcl.IsEmptyValueIndirect(v) {
		m["diskUsageBytes"] = v
	}
	if v := f.RuntimeApiVersion; !dcl.IsEmptyValueIndirect(v) {
		m["runtimeApiVersion"] = v
	}
	if v := f.RuntimeMainExecutablePath; !dcl.IsEmptyValueIndirect(v) {
		m["runtimeMainExecutablePath"] = v
	}
	if v, err := expandVersionHandlersSlice(c, f.Handlers); err != nil {
		return nil, fmt.Errorf("error expanding Handlers into handlers: %w", err)
	} else if v != nil {
		m["handlers"] = v
	}
	if v, err := expandVersionErrorHandlersSlice(c, f.ErrorHandlers); err != nil {
		return nil, fmt.Errorf("error expanding ErrorHandlers into errorHandlers: %w", err)
	} else if v != nil {
		m["errorHandlers"] = v
	}
	if v, err := expandVersionLibrariesSlice(c, f.Libraries); err != nil {
		return nil, fmt.Errorf("error expanding Libraries into libraries: %w", err)
	} else if v != nil {
		m["libraries"] = v
	}
	if v, err := expandVersionApiConfig(c, f.ApiConfig); err != nil {
		return nil, fmt.Errorf("error expanding ApiConfig into apiConfig: %w", err)
	} else if v != nil {
		m["apiConfig"] = v
	}
	if v := f.EnvVariables; !dcl.IsEmptyValueIndirect(v) {
		m["envVariables"] = v
	}
	if v := f.DefaultExpiration; !dcl.IsEmptyValueIndirect(v) {
		m["defaultExpiration"] = v
	}
	if v, err := expandVersionDeployment(c, f.Deployment); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployment: %w", err)
	} else if v != nil {
		m["deployment"] = v
	}
	if v, err := expandVersionHealthCheck(c, f.HealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HealthCheck into healthCheck: %w", err)
	} else if v != nil {
		m["healthCheck"] = v
	}
	if v, err := expandVersionReadinessCheck(c, f.ReadinessCheck); err != nil {
		return nil, fmt.Errorf("error expanding ReadinessCheck into readinessCheck: %w", err)
	} else if v != nil {
		m["readinessCheck"] = v
	}
	if v, err := expandVersionLivenessCheck(c, f.LivenessCheck); err != nil {
		return nil, fmt.Errorf("error expanding LivenessCheck into livenessCheck: %w", err)
	} else if v != nil {
		m["livenessCheck"] = v
	}
	if v := f.NobuildFilesRegex; !dcl.IsEmptyValueIndirect(v) {
		m["nobuildFilesRegex"] = v
	}
	if v := f.VersionUrl; !dcl.IsEmptyValueIndirect(v) {
		m["versionUrl"] = v
	}
	if v, err := expandVersionEntrypoint(c, f.Entrypoint); err != nil {
		return nil, fmt.Errorf("error expanding Entrypoint into entrypoint: %w", err)
	} else if v != nil {
		m["entrypoint"] = v
	}
	if v, err := expandVersionVPCAccessConnector(c, f.VPCAccessConnector); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccessConnector into vpcAccessConnector: %w", err)
	} else if v != nil {
		m["vpcAccessConnector"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding App into app: %w", err)
	} else if v != nil {
		m["app"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Service into service: %w", err)
	} else if v != nil {
		m["service"] = v
	}

	return m, nil
}

// flattenVersion flattens Version from a JSON request object into the
// Version type.
func flattenVersion(c *Client, i interface{}) *Version {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Version{}
	res.ConsumerName = dcl.FlattenString(m["consumerName"])
	res.Name = dcl.FlattenString(m["id"])
	res.AutomaticScaling = flattenVersionAutomaticScaling(c, m["automaticScaling"])
	res.BasicScaling = flattenVersionBasicScaling(c, m["basicScaling"])
	res.ManualScaling = flattenVersionManualScaling(c, m["manualScaling"])
	res.InboundServices = flattenVersionInboundServicesEnumSlice(c, m["inboundServices"])
	res.InstanceClass = dcl.FlattenString(m["instanceClass"])
	res.Network = flattenVersionNetwork(c, m["network"])
	res.Zones = dcl.FlattenStringSlice(m["zones"])
	res.Resources = flattenVersionResources(c, m["resources"])
	res.Runtime = dcl.FlattenString(m["runtime"])
	res.RuntimeChannel = dcl.FlattenString(m["runtimeChannel"])
	res.Threadsafe = dcl.FlattenBool(m["threadsafe"])
	res.Vm = dcl.FlattenBool(m["vm"])
	res.BetaSettings = dcl.FlattenKeyValuePairs(m["betaSettings"])
	res.Env = dcl.FlattenString(m["env"])
	res.ServingStatus = flattenVersionServingStatusEnum(m["servingStatus"])
	res.CreatedBy = dcl.FlattenString(m["createdBy"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.DiskUsageBytes = dcl.FlattenInteger(m["diskUsageBytes"])
	res.RuntimeApiVersion = dcl.FlattenString(m["runtimeApiVersion"])
	res.RuntimeMainExecutablePath = dcl.FlattenString(m["runtimeMainExecutablePath"])
	res.Handlers = flattenVersionHandlersSlice(c, m["handlers"])
	res.ErrorHandlers = flattenVersionErrorHandlersSlice(c, m["errorHandlers"])
	res.Libraries = flattenVersionLibrariesSlice(c, m["libraries"])
	res.ApiConfig = flattenVersionApiConfig(c, m["apiConfig"])
	res.EnvVariables = dcl.FlattenKeyValuePairs(m["envVariables"])
	res.DefaultExpiration = dcl.FlattenString(m["defaultExpiration"])
	res.Deployment = flattenVersionDeployment(c, m["deployment"])
	res.HealthCheck = flattenVersionHealthCheck(c, m["healthCheck"])
	res.ReadinessCheck = flattenVersionReadinessCheck(c, m["readinessCheck"])
	res.LivenessCheck = flattenVersionLivenessCheck(c, m["livenessCheck"])
	res.NobuildFilesRegex = dcl.FlattenString(m["nobuildFilesRegex"])
	res.VersionUrl = dcl.FlattenString(m["versionUrl"])
	res.Entrypoint = flattenVersionEntrypoint(c, m["entrypoint"])
	res.VPCAccessConnector = flattenVersionVPCAccessConnector(c, m["vpcAccessConnector"])
	res.App = dcl.FlattenString(m["app"])
	res.Service = dcl.FlattenString(m["service"])

	return res
}

// expandVersionAutomaticScalingMap expands the contents of VersionAutomaticScaling into a JSON
// request object.
func expandVersionAutomaticScalingMap(c *Client, f map[string]VersionAutomaticScaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingSlice expands the contents of VersionAutomaticScaling into a JSON
// request object.
func expandVersionAutomaticScalingSlice(c *Client, f []VersionAutomaticScaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingMap flattens the contents of VersionAutomaticScaling from a JSON
// response object.
func flattenVersionAutomaticScalingMap(c *Client, i interface{}) map[string]VersionAutomaticScaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScaling{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScaling{}
	}

	items := make(map[string]VersionAutomaticScaling)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingSlice flattens the contents of VersionAutomaticScaling from a JSON
// response object.
func flattenVersionAutomaticScalingSlice(c *Client, i interface{}) []VersionAutomaticScaling {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScaling{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScaling{}
	}

	items := make([]VersionAutomaticScaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScaling expands an instance of VersionAutomaticScaling into a JSON
// request object.
func expandVersionAutomaticScaling(c *Client, f *VersionAutomaticScaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CoolDownPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["coolDownPeriod"] = v
	}
	if v, err := expandVersionAutomaticScalingCpuUtilization(c, f.CpuUtilization); err != nil {
		return nil, fmt.Errorf("error expanding CpuUtilization into cpuUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cpuUtilization"] = v
	}
	if v := f.MaxConcurrentRequests; !dcl.IsEmptyValueIndirect(v) {
		m["maxConcurrentRequests"] = v
	}
	if v := f.MaxIdleInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxIdleInstances"] = v
	}
	if v := f.MaxTotalInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxTotalInstances"] = v
	}
	if v := f.MaxPendingLatency; !dcl.IsEmptyValueIndirect(v) {
		m["maxPendingLatency"] = v
	}
	if v := f.MinIdleInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minIdleInstances"] = v
	}
	if v := f.MinTotalInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minTotalInstances"] = v
	}
	if v := f.MinPendingLatency; !dcl.IsEmptyValueIndirect(v) {
		m["minPendingLatency"] = v
	}
	if v, err := expandVersionAutomaticScalingRequestUtilization(c, f.RequestUtilization); err != nil {
		return nil, fmt.Errorf("error expanding RequestUtilization into requestUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestUtilization"] = v
	}
	if v, err := expandVersionAutomaticScalingDiskUtilization(c, f.DiskUtilization); err != nil {
		return nil, fmt.Errorf("error expanding DiskUtilization into diskUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskUtilization"] = v
	}
	if v, err := expandVersionAutomaticScalingNetworkUtilization(c, f.NetworkUtilization); err != nil {
		return nil, fmt.Errorf("error expanding NetworkUtilization into networkUtilization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkUtilization"] = v
	}
	if v, err := expandVersionAutomaticScalingStandardSchedulerSettings(c, f.StandardSchedulerSettings); err != nil {
		return nil, fmt.Errorf("error expanding StandardSchedulerSettings into standardSchedulerSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["standardSchedulerSettings"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScaling flattens an instance of VersionAutomaticScaling from a JSON
// response object.
func flattenVersionAutomaticScaling(c *Client, i interface{}) *VersionAutomaticScaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScaling
	}
	r.CoolDownPeriod = dcl.FlattenString(m["coolDownPeriod"])
	r.CpuUtilization = flattenVersionAutomaticScalingCpuUtilization(c, m["cpuUtilization"])
	r.MaxConcurrentRequests = dcl.FlattenInteger(m["maxConcurrentRequests"])
	r.MaxIdleInstances = dcl.FlattenInteger(m["maxIdleInstances"])
	r.MaxTotalInstances = dcl.FlattenInteger(m["maxTotalInstances"])
	r.MaxPendingLatency = dcl.FlattenString(m["maxPendingLatency"])
	r.MinIdleInstances = dcl.FlattenInteger(m["minIdleInstances"])
	r.MinTotalInstances = dcl.FlattenInteger(m["minTotalInstances"])
	r.MinPendingLatency = dcl.FlattenString(m["minPendingLatency"])
	r.RequestUtilization = flattenVersionAutomaticScalingRequestUtilization(c, m["requestUtilization"])
	r.DiskUtilization = flattenVersionAutomaticScalingDiskUtilization(c, m["diskUtilization"])
	r.NetworkUtilization = flattenVersionAutomaticScalingNetworkUtilization(c, m["networkUtilization"])
	r.StandardSchedulerSettings = flattenVersionAutomaticScalingStandardSchedulerSettings(c, m["standardSchedulerSettings"])

	return r
}

// expandVersionAutomaticScalingCpuUtilizationMap expands the contents of VersionAutomaticScalingCpuUtilization into a JSON
// request object.
func expandVersionAutomaticScalingCpuUtilizationMap(c *Client, f map[string]VersionAutomaticScalingCpuUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScalingCpuUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingCpuUtilizationSlice expands the contents of VersionAutomaticScalingCpuUtilization into a JSON
// request object.
func expandVersionAutomaticScalingCpuUtilizationSlice(c *Client, f []VersionAutomaticScalingCpuUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScalingCpuUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingCpuUtilizationMap flattens the contents of VersionAutomaticScalingCpuUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingCpuUtilizationMap(c *Client, i interface{}) map[string]VersionAutomaticScalingCpuUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScalingCpuUtilization{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScalingCpuUtilization{}
	}

	items := make(map[string]VersionAutomaticScalingCpuUtilization)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScalingCpuUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingCpuUtilizationSlice flattens the contents of VersionAutomaticScalingCpuUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingCpuUtilizationSlice(c *Client, i interface{}) []VersionAutomaticScalingCpuUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScalingCpuUtilization{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScalingCpuUtilization{}
	}

	items := make([]VersionAutomaticScalingCpuUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScalingCpuUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScalingCpuUtilization expands an instance of VersionAutomaticScalingCpuUtilization into a JSON
// request object.
func expandVersionAutomaticScalingCpuUtilization(c *Client, f *VersionAutomaticScalingCpuUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AggregationWindowLength; !dcl.IsEmptyValueIndirect(v) {
		m["aggregationWindowLength"] = v
	}
	if v := f.TargetUtilization; !dcl.IsEmptyValueIndirect(v) {
		m["targetUtilization"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScalingCpuUtilization flattens an instance of VersionAutomaticScalingCpuUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingCpuUtilization(c *Client, i interface{}) *VersionAutomaticScalingCpuUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScalingCpuUtilization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScalingCpuUtilization
	}
	r.AggregationWindowLength = dcl.FlattenString(m["aggregationWindowLength"])
	r.TargetUtilization = dcl.FlattenDouble(m["targetUtilization"])

	return r
}

// expandVersionAutomaticScalingRequestUtilizationMap expands the contents of VersionAutomaticScalingRequestUtilization into a JSON
// request object.
func expandVersionAutomaticScalingRequestUtilizationMap(c *Client, f map[string]VersionAutomaticScalingRequestUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScalingRequestUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingRequestUtilizationSlice expands the contents of VersionAutomaticScalingRequestUtilization into a JSON
// request object.
func expandVersionAutomaticScalingRequestUtilizationSlice(c *Client, f []VersionAutomaticScalingRequestUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScalingRequestUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingRequestUtilizationMap flattens the contents of VersionAutomaticScalingRequestUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingRequestUtilizationMap(c *Client, i interface{}) map[string]VersionAutomaticScalingRequestUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScalingRequestUtilization{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScalingRequestUtilization{}
	}

	items := make(map[string]VersionAutomaticScalingRequestUtilization)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScalingRequestUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingRequestUtilizationSlice flattens the contents of VersionAutomaticScalingRequestUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingRequestUtilizationSlice(c *Client, i interface{}) []VersionAutomaticScalingRequestUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScalingRequestUtilization{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScalingRequestUtilization{}
	}

	items := make([]VersionAutomaticScalingRequestUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScalingRequestUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScalingRequestUtilization expands an instance of VersionAutomaticScalingRequestUtilization into a JSON
// request object.
func expandVersionAutomaticScalingRequestUtilization(c *Client, f *VersionAutomaticScalingRequestUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetRequestCountPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetRequestCountPerSecond"] = v
	}
	if v := f.TargetConcurrentRequests; !dcl.IsEmptyValueIndirect(v) {
		m["targetConcurrentRequests"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScalingRequestUtilization flattens an instance of VersionAutomaticScalingRequestUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingRequestUtilization(c *Client, i interface{}) *VersionAutomaticScalingRequestUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScalingRequestUtilization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScalingRequestUtilization
	}
	r.TargetRequestCountPerSecond = dcl.FlattenInteger(m["targetRequestCountPerSecond"])
	r.TargetConcurrentRequests = dcl.FlattenInteger(m["targetConcurrentRequests"])

	return r
}

// expandVersionAutomaticScalingDiskUtilizationMap expands the contents of VersionAutomaticScalingDiskUtilization into a JSON
// request object.
func expandVersionAutomaticScalingDiskUtilizationMap(c *Client, f map[string]VersionAutomaticScalingDiskUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScalingDiskUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingDiskUtilizationSlice expands the contents of VersionAutomaticScalingDiskUtilization into a JSON
// request object.
func expandVersionAutomaticScalingDiskUtilizationSlice(c *Client, f []VersionAutomaticScalingDiskUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScalingDiskUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingDiskUtilizationMap flattens the contents of VersionAutomaticScalingDiskUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingDiskUtilizationMap(c *Client, i interface{}) map[string]VersionAutomaticScalingDiskUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScalingDiskUtilization{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScalingDiskUtilization{}
	}

	items := make(map[string]VersionAutomaticScalingDiskUtilization)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScalingDiskUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingDiskUtilizationSlice flattens the contents of VersionAutomaticScalingDiskUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingDiskUtilizationSlice(c *Client, i interface{}) []VersionAutomaticScalingDiskUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScalingDiskUtilization{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScalingDiskUtilization{}
	}

	items := make([]VersionAutomaticScalingDiskUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScalingDiskUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScalingDiskUtilization expands an instance of VersionAutomaticScalingDiskUtilization into a JSON
// request object.
func expandVersionAutomaticScalingDiskUtilization(c *Client, f *VersionAutomaticScalingDiskUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetWriteBytesPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetWriteBytesPerSecond"] = v
	}
	if v := f.TargetWriteOpsPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetWriteOpsPerSecond"] = v
	}
	if v := f.TargetReadBytesPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetReadBytesPerSecond"] = v
	}
	if v := f.TargetReadOpsPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetReadOpsPerSecond"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScalingDiskUtilization flattens an instance of VersionAutomaticScalingDiskUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingDiskUtilization(c *Client, i interface{}) *VersionAutomaticScalingDiskUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScalingDiskUtilization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScalingDiskUtilization
	}
	r.TargetWriteBytesPerSecond = dcl.FlattenInteger(m["targetWriteBytesPerSecond"])
	r.TargetWriteOpsPerSecond = dcl.FlattenInteger(m["targetWriteOpsPerSecond"])
	r.TargetReadBytesPerSecond = dcl.FlattenInteger(m["targetReadBytesPerSecond"])
	r.TargetReadOpsPerSecond = dcl.FlattenInteger(m["targetReadOpsPerSecond"])

	return r
}

// expandVersionAutomaticScalingNetworkUtilizationMap expands the contents of VersionAutomaticScalingNetworkUtilization into a JSON
// request object.
func expandVersionAutomaticScalingNetworkUtilizationMap(c *Client, f map[string]VersionAutomaticScalingNetworkUtilization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScalingNetworkUtilization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingNetworkUtilizationSlice expands the contents of VersionAutomaticScalingNetworkUtilization into a JSON
// request object.
func expandVersionAutomaticScalingNetworkUtilizationSlice(c *Client, f []VersionAutomaticScalingNetworkUtilization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScalingNetworkUtilization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingNetworkUtilizationMap flattens the contents of VersionAutomaticScalingNetworkUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingNetworkUtilizationMap(c *Client, i interface{}) map[string]VersionAutomaticScalingNetworkUtilization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScalingNetworkUtilization{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScalingNetworkUtilization{}
	}

	items := make(map[string]VersionAutomaticScalingNetworkUtilization)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScalingNetworkUtilization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingNetworkUtilizationSlice flattens the contents of VersionAutomaticScalingNetworkUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingNetworkUtilizationSlice(c *Client, i interface{}) []VersionAutomaticScalingNetworkUtilization {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScalingNetworkUtilization{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScalingNetworkUtilization{}
	}

	items := make([]VersionAutomaticScalingNetworkUtilization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScalingNetworkUtilization(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScalingNetworkUtilization expands an instance of VersionAutomaticScalingNetworkUtilization into a JSON
// request object.
func expandVersionAutomaticScalingNetworkUtilization(c *Client, f *VersionAutomaticScalingNetworkUtilization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetSentBytesPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetSentBytesPerSecond"] = v
	}
	if v := f.TargetSentPacketsPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetSentPacketsPerSecond"] = v
	}
	if v := f.TargetReceivedBytesPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetReceivedBytesPerSecond"] = v
	}
	if v := f.TargetReceivedPacketsPerSecond; !dcl.IsEmptyValueIndirect(v) {
		m["targetReceivedPacketsPerSecond"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScalingNetworkUtilization flattens an instance of VersionAutomaticScalingNetworkUtilization from a JSON
// response object.
func flattenVersionAutomaticScalingNetworkUtilization(c *Client, i interface{}) *VersionAutomaticScalingNetworkUtilization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScalingNetworkUtilization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScalingNetworkUtilization
	}
	r.TargetSentBytesPerSecond = dcl.FlattenInteger(m["targetSentBytesPerSecond"])
	r.TargetSentPacketsPerSecond = dcl.FlattenInteger(m["targetSentPacketsPerSecond"])
	r.TargetReceivedBytesPerSecond = dcl.FlattenInteger(m["targetReceivedBytesPerSecond"])
	r.TargetReceivedPacketsPerSecond = dcl.FlattenInteger(m["targetReceivedPacketsPerSecond"])

	return r
}

// expandVersionAutomaticScalingStandardSchedulerSettingsMap expands the contents of VersionAutomaticScalingStandardSchedulerSettings into a JSON
// request object.
func expandVersionAutomaticScalingStandardSchedulerSettingsMap(c *Client, f map[string]VersionAutomaticScalingStandardSchedulerSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionAutomaticScalingStandardSchedulerSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionAutomaticScalingStandardSchedulerSettingsSlice expands the contents of VersionAutomaticScalingStandardSchedulerSettings into a JSON
// request object.
func expandVersionAutomaticScalingStandardSchedulerSettingsSlice(c *Client, f []VersionAutomaticScalingStandardSchedulerSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionAutomaticScalingStandardSchedulerSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionAutomaticScalingStandardSchedulerSettingsMap flattens the contents of VersionAutomaticScalingStandardSchedulerSettings from a JSON
// response object.
func flattenVersionAutomaticScalingStandardSchedulerSettingsMap(c *Client, i interface{}) map[string]VersionAutomaticScalingStandardSchedulerSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionAutomaticScalingStandardSchedulerSettings{}
	}

	if len(a) == 0 {
		return map[string]VersionAutomaticScalingStandardSchedulerSettings{}
	}

	items := make(map[string]VersionAutomaticScalingStandardSchedulerSettings)
	for k, item := range a {
		items[k] = *flattenVersionAutomaticScalingStandardSchedulerSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionAutomaticScalingStandardSchedulerSettingsSlice flattens the contents of VersionAutomaticScalingStandardSchedulerSettings from a JSON
// response object.
func flattenVersionAutomaticScalingStandardSchedulerSettingsSlice(c *Client, i interface{}) []VersionAutomaticScalingStandardSchedulerSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionAutomaticScalingStandardSchedulerSettings{}
	}

	if len(a) == 0 {
		return []VersionAutomaticScalingStandardSchedulerSettings{}
	}

	items := make([]VersionAutomaticScalingStandardSchedulerSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionAutomaticScalingStandardSchedulerSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionAutomaticScalingStandardSchedulerSettings expands an instance of VersionAutomaticScalingStandardSchedulerSettings into a JSON
// request object.
func expandVersionAutomaticScalingStandardSchedulerSettings(c *Client, f *VersionAutomaticScalingStandardSchedulerSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetCpuUtilization; !dcl.IsEmptyValueIndirect(v) {
		m["targetCpuUtilization"] = v
	}
	if v := f.TargetThroughputUtilization; !dcl.IsEmptyValueIndirect(v) {
		m["targetThroughputUtilization"] = v
	}
	if v := f.MinInstances; !dcl.IsEmptyValueIndirect(v) {
		m["minInstances"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}

	return m, nil
}

// flattenVersionAutomaticScalingStandardSchedulerSettings flattens an instance of VersionAutomaticScalingStandardSchedulerSettings from a JSON
// response object.
func flattenVersionAutomaticScalingStandardSchedulerSettings(c *Client, i interface{}) *VersionAutomaticScalingStandardSchedulerSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionAutomaticScalingStandardSchedulerSettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionAutomaticScalingStandardSchedulerSettings
	}
	r.TargetCpuUtilization = dcl.FlattenDouble(m["targetCpuUtilization"])
	r.TargetThroughputUtilization = dcl.FlattenDouble(m["targetThroughputUtilization"])
	r.MinInstances = dcl.FlattenInteger(m["minInstances"])
	r.MaxInstances = dcl.FlattenInteger(m["maxInstances"])

	return r
}

// expandVersionBasicScalingMap expands the contents of VersionBasicScaling into a JSON
// request object.
func expandVersionBasicScalingMap(c *Client, f map[string]VersionBasicScaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionBasicScaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionBasicScalingSlice expands the contents of VersionBasicScaling into a JSON
// request object.
func expandVersionBasicScalingSlice(c *Client, f []VersionBasicScaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionBasicScaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionBasicScalingMap flattens the contents of VersionBasicScaling from a JSON
// response object.
func flattenVersionBasicScalingMap(c *Client, i interface{}) map[string]VersionBasicScaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionBasicScaling{}
	}

	if len(a) == 0 {
		return map[string]VersionBasicScaling{}
	}

	items := make(map[string]VersionBasicScaling)
	for k, item := range a {
		items[k] = *flattenVersionBasicScaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionBasicScalingSlice flattens the contents of VersionBasicScaling from a JSON
// response object.
func flattenVersionBasicScalingSlice(c *Client, i interface{}) []VersionBasicScaling {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionBasicScaling{}
	}

	if len(a) == 0 {
		return []VersionBasicScaling{}
	}

	items := make([]VersionBasicScaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionBasicScaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionBasicScaling expands an instance of VersionBasicScaling into a JSON
// request object.
func expandVersionBasicScaling(c *Client, f *VersionBasicScaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IdleTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["idleTimeout"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}

	return m, nil
}

// flattenVersionBasicScaling flattens an instance of VersionBasicScaling from a JSON
// response object.
func flattenVersionBasicScaling(c *Client, i interface{}) *VersionBasicScaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionBasicScaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionBasicScaling
	}
	r.IdleTimeout = dcl.FlattenString(m["idleTimeout"])
	r.MaxInstances = dcl.FlattenInteger(m["maxInstances"])

	return r
}

// expandVersionManualScalingMap expands the contents of VersionManualScaling into a JSON
// request object.
func expandVersionManualScalingMap(c *Client, f map[string]VersionManualScaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionManualScaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionManualScalingSlice expands the contents of VersionManualScaling into a JSON
// request object.
func expandVersionManualScalingSlice(c *Client, f []VersionManualScaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionManualScaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionManualScalingMap flattens the contents of VersionManualScaling from a JSON
// response object.
func flattenVersionManualScalingMap(c *Client, i interface{}) map[string]VersionManualScaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionManualScaling{}
	}

	if len(a) == 0 {
		return map[string]VersionManualScaling{}
	}

	items := make(map[string]VersionManualScaling)
	for k, item := range a {
		items[k] = *flattenVersionManualScaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionManualScalingSlice flattens the contents of VersionManualScaling from a JSON
// response object.
func flattenVersionManualScalingSlice(c *Client, i interface{}) []VersionManualScaling {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionManualScaling{}
	}

	if len(a) == 0 {
		return []VersionManualScaling{}
	}

	items := make([]VersionManualScaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionManualScaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionManualScaling expands an instance of VersionManualScaling into a JSON
// request object.
func expandVersionManualScaling(c *Client, f *VersionManualScaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Instances; !dcl.IsEmptyValueIndirect(v) {
		m["instances"] = v
	}

	return m, nil
}

// flattenVersionManualScaling flattens an instance of VersionManualScaling from a JSON
// response object.
func flattenVersionManualScaling(c *Client, i interface{}) *VersionManualScaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionManualScaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionManualScaling
	}
	r.Instances = dcl.FlattenInteger(m["instances"])

	return r
}

// expandVersionNetworkMap expands the contents of VersionNetwork into a JSON
// request object.
func expandVersionNetworkMap(c *Client, f map[string]VersionNetwork) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionNetwork(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionNetworkSlice expands the contents of VersionNetwork into a JSON
// request object.
func expandVersionNetworkSlice(c *Client, f []VersionNetwork) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionNetwork(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionNetworkMap flattens the contents of VersionNetwork from a JSON
// response object.
func flattenVersionNetworkMap(c *Client, i interface{}) map[string]VersionNetwork {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionNetwork{}
	}

	if len(a) == 0 {
		return map[string]VersionNetwork{}
	}

	items := make(map[string]VersionNetwork)
	for k, item := range a {
		items[k] = *flattenVersionNetwork(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionNetworkSlice flattens the contents of VersionNetwork from a JSON
// response object.
func flattenVersionNetworkSlice(c *Client, i interface{}) []VersionNetwork {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionNetwork{}
	}

	if len(a) == 0 {
		return []VersionNetwork{}
	}

	items := make([]VersionNetwork, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionNetwork(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionNetwork expands an instance of VersionNetwork into a JSON
// request object.
func expandVersionNetwork(c *Client, f *VersionNetwork) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ForwardedPorts; !dcl.IsEmptyValueIndirect(v) {
		m["forwardedPorts"] = v
	}
	if v := f.InstanceTag; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTag"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.SubnetworkName; !dcl.IsEmptyValueIndirect(v) {
		m["subnetworkName"] = v
	}
	if v := f.SessionAffinity; !dcl.IsEmptyValueIndirect(v) {
		m["sessionAffinity"] = v
	}

	return m, nil
}

// flattenVersionNetwork flattens an instance of VersionNetwork from a JSON
// response object.
func flattenVersionNetwork(c *Client, i interface{}) *VersionNetwork {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionNetwork{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionNetwork
	}
	r.ForwardedPorts = dcl.FlattenStringSlice(m["forwardedPorts"])
	r.InstanceTag = dcl.FlattenString(m["instanceTag"])
	r.Name = dcl.FlattenString(m["name"])
	r.SubnetworkName = dcl.FlattenString(m["subnetworkName"])
	r.SessionAffinity = dcl.FlattenBool(m["sessionAffinity"])

	return r
}

// expandVersionResourcesMap expands the contents of VersionResources into a JSON
// request object.
func expandVersionResourcesMap(c *Client, f map[string]VersionResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionResourcesSlice expands the contents of VersionResources into a JSON
// request object.
func expandVersionResourcesSlice(c *Client, f []VersionResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionResourcesMap flattens the contents of VersionResources from a JSON
// response object.
func flattenVersionResourcesMap(c *Client, i interface{}) map[string]VersionResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionResources{}
	}

	if len(a) == 0 {
		return map[string]VersionResources{}
	}

	items := make(map[string]VersionResources)
	for k, item := range a {
		items[k] = *flattenVersionResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionResourcesSlice flattens the contents of VersionResources from a JSON
// response object.
func flattenVersionResourcesSlice(c *Client, i interface{}) []VersionResources {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionResources{}
	}

	if len(a) == 0 {
		return []VersionResources{}
	}

	items := make([]VersionResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionResources expands an instance of VersionResources into a JSON
// request object.
func expandVersionResources(c *Client, f *VersionResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Cpu; !dcl.IsEmptyValueIndirect(v) {
		m["cpu"] = v
	}
	if v := f.DiskGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskGb"] = v
	}
	if v := f.MemoryGb; !dcl.IsEmptyValueIndirect(v) {
		m["memoryGb"] = v
	}
	if v, err := expandVersionResourcesVolumesSlice(c, f.Volumes); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["volumes"] = v
	}

	return m, nil
}

// flattenVersionResources flattens an instance of VersionResources from a JSON
// response object.
func flattenVersionResources(c *Client, i interface{}) *VersionResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionResources
	}
	r.Cpu = dcl.FlattenDouble(m["cpu"])
	r.DiskGb = dcl.FlattenDouble(m["diskGb"])
	r.MemoryGb = dcl.FlattenDouble(m["memoryGb"])
	r.Volumes = flattenVersionResourcesVolumesSlice(c, m["volumes"])

	return r
}

// expandVersionResourcesVolumesMap expands the contents of VersionResourcesVolumes into a JSON
// request object.
func expandVersionResourcesVolumesMap(c *Client, f map[string]VersionResourcesVolumes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionResourcesVolumes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionResourcesVolumesSlice expands the contents of VersionResourcesVolumes into a JSON
// request object.
func expandVersionResourcesVolumesSlice(c *Client, f []VersionResourcesVolumes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionResourcesVolumes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionResourcesVolumesMap flattens the contents of VersionResourcesVolumes from a JSON
// response object.
func flattenVersionResourcesVolumesMap(c *Client, i interface{}) map[string]VersionResourcesVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionResourcesVolumes{}
	}

	if len(a) == 0 {
		return map[string]VersionResourcesVolumes{}
	}

	items := make(map[string]VersionResourcesVolumes)
	for k, item := range a {
		items[k] = *flattenVersionResourcesVolumes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionResourcesVolumesSlice flattens the contents of VersionResourcesVolumes from a JSON
// response object.
func flattenVersionResourcesVolumesSlice(c *Client, i interface{}) []VersionResourcesVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionResourcesVolumes{}
	}

	if len(a) == 0 {
		return []VersionResourcesVolumes{}
	}

	items := make([]VersionResourcesVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionResourcesVolumes(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionResourcesVolumes expands an instance of VersionResourcesVolumes into a JSON
// request object.
func expandVersionResourcesVolumes(c *Client, f *VersionResourcesVolumes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.VolumeType; !dcl.IsEmptyValueIndirect(v) {
		m["volumeType"] = v
	}
	if v := f.SizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGb"] = v
	}

	return m, nil
}

// flattenVersionResourcesVolumes flattens an instance of VersionResourcesVolumes from a JSON
// response object.
func flattenVersionResourcesVolumes(c *Client, i interface{}) *VersionResourcesVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionResourcesVolumes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionResourcesVolumes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.VolumeType = dcl.FlattenString(m["volumeType"])
	r.SizeGb = dcl.FlattenDouble(m["sizeGb"])

	return r
}

// expandVersionHandlersMap expands the contents of VersionHandlers into a JSON
// request object.
func expandVersionHandlersMap(c *Client, f map[string]VersionHandlers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionHandlers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionHandlersSlice expands the contents of VersionHandlers into a JSON
// request object.
func expandVersionHandlersSlice(c *Client, f []VersionHandlers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionHandlers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionHandlersMap flattens the contents of VersionHandlers from a JSON
// response object.
func flattenVersionHandlersMap(c *Client, i interface{}) map[string]VersionHandlers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionHandlers{}
	}

	if len(a) == 0 {
		return map[string]VersionHandlers{}
	}

	items := make(map[string]VersionHandlers)
	for k, item := range a {
		items[k] = *flattenVersionHandlers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionHandlersSlice flattens the contents of VersionHandlers from a JSON
// response object.
func flattenVersionHandlersSlice(c *Client, i interface{}) []VersionHandlers {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlers{}
	}

	if len(a) == 0 {
		return []VersionHandlers{}
	}

	items := make([]VersionHandlers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlers(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionHandlers expands an instance of VersionHandlers into a JSON
// request object.
func expandVersionHandlers(c *Client, f *VersionHandlers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UrlRegex; !dcl.IsEmptyValueIndirect(v) {
		m["urlRegex"] = v
	}
	if v, err := expandVersionHandlersStaticFiles(c, f.StaticFiles); err != nil {
		return nil, fmt.Errorf("error expanding StaticFiles into staticFiles: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["staticFiles"] = v
	}
	if v, err := expandVersionHandlersScript(c, f.Script); err != nil {
		return nil, fmt.Errorf("error expanding Script into script: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["script"] = v
	}
	if v, err := expandVersionHandlersApiEndpoint(c, f.ApiEndpoint); err != nil {
		return nil, fmt.Errorf("error expanding ApiEndpoint into apiEndpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apiEndpoint"] = v
	}
	if v := f.SecurityLevel; !dcl.IsEmptyValueIndirect(v) {
		m["securityLevel"] = v
	}
	if v := f.Login; !dcl.IsEmptyValueIndirect(v) {
		m["login"] = v
	}
	if v := f.AuthFailAction; !dcl.IsEmptyValueIndirect(v) {
		m["authFailAction"] = v
	}
	if v := f.RedirectHttpResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["redirectHttpResponseCode"] = v
	}

	return m, nil
}

// flattenVersionHandlers flattens an instance of VersionHandlers from a JSON
// response object.
func flattenVersionHandlers(c *Client, i interface{}) *VersionHandlers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionHandlers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionHandlers
	}
	r.UrlRegex = dcl.FlattenString(m["urlRegex"])
	r.StaticFiles = flattenVersionHandlersStaticFiles(c, m["staticFiles"])
	r.Script = flattenVersionHandlersScript(c, m["script"])
	r.ApiEndpoint = flattenVersionHandlersApiEndpoint(c, m["apiEndpoint"])
	r.SecurityLevel = flattenVersionHandlersSecurityLevelEnum(m["securityLevel"])
	r.Login = flattenVersionHandlersLoginEnum(m["login"])
	r.AuthFailAction = flattenVersionHandlersAuthFailActionEnum(m["authFailAction"])
	r.RedirectHttpResponseCode = flattenVersionHandlersRedirectHttpResponseCodeEnum(m["redirectHttpResponseCode"])

	return r
}

// expandVersionHandlersStaticFilesMap expands the contents of VersionHandlersStaticFiles into a JSON
// request object.
func expandVersionHandlersStaticFilesMap(c *Client, f map[string]VersionHandlersStaticFiles) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionHandlersStaticFiles(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionHandlersStaticFilesSlice expands the contents of VersionHandlersStaticFiles into a JSON
// request object.
func expandVersionHandlersStaticFilesSlice(c *Client, f []VersionHandlersStaticFiles) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionHandlersStaticFiles(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionHandlersStaticFilesMap flattens the contents of VersionHandlersStaticFiles from a JSON
// response object.
func flattenVersionHandlersStaticFilesMap(c *Client, i interface{}) map[string]VersionHandlersStaticFiles {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionHandlersStaticFiles{}
	}

	if len(a) == 0 {
		return map[string]VersionHandlersStaticFiles{}
	}

	items := make(map[string]VersionHandlersStaticFiles)
	for k, item := range a {
		items[k] = *flattenVersionHandlersStaticFiles(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionHandlersStaticFilesSlice flattens the contents of VersionHandlersStaticFiles from a JSON
// response object.
func flattenVersionHandlersStaticFilesSlice(c *Client, i interface{}) []VersionHandlersStaticFiles {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersStaticFiles{}
	}

	if len(a) == 0 {
		return []VersionHandlersStaticFiles{}
	}

	items := make([]VersionHandlersStaticFiles, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersStaticFiles(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionHandlersStaticFiles expands an instance of VersionHandlersStaticFiles into a JSON
// request object.
func expandVersionHandlersStaticFiles(c *Client, f *VersionHandlersStaticFiles) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.UploadPathRegex; !dcl.IsEmptyValueIndirect(v) {
		m["uploadPathRegex"] = v
	}
	if v := f.HttpHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaders"] = v
	}
	if v := f.MimeType; !dcl.IsEmptyValueIndirect(v) {
		m["mimeType"] = v
	}
	if v := f.Expiration; !dcl.IsEmptyValueIndirect(v) {
		m["expiration"] = v
	}
	if v := f.RequireMatchingFile; !dcl.IsEmptyValueIndirect(v) {
		m["requireMatchingFile"] = v
	}
	if v := f.ApplicationReadable; !dcl.IsEmptyValueIndirect(v) {
		m["applicationReadable"] = v
	}

	return m, nil
}

// flattenVersionHandlersStaticFiles flattens an instance of VersionHandlersStaticFiles from a JSON
// response object.
func flattenVersionHandlersStaticFiles(c *Client, i interface{}) *VersionHandlersStaticFiles {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionHandlersStaticFiles{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionHandlersStaticFiles
	}
	r.Path = dcl.FlattenString(m["path"])
	r.UploadPathRegex = dcl.FlattenString(m["uploadPathRegex"])
	r.HttpHeaders = dcl.FlattenKeyValuePairs(m["httpHeaders"])
	r.MimeType = dcl.FlattenString(m["mimeType"])
	r.Expiration = dcl.FlattenString(m["expiration"])
	r.RequireMatchingFile = dcl.FlattenBool(m["requireMatchingFile"])
	r.ApplicationReadable = dcl.FlattenBool(m["applicationReadable"])

	return r
}

// expandVersionHandlersScriptMap expands the contents of VersionHandlersScript into a JSON
// request object.
func expandVersionHandlersScriptMap(c *Client, f map[string]VersionHandlersScript) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionHandlersScript(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionHandlersScriptSlice expands the contents of VersionHandlersScript into a JSON
// request object.
func expandVersionHandlersScriptSlice(c *Client, f []VersionHandlersScript) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionHandlersScript(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionHandlersScriptMap flattens the contents of VersionHandlersScript from a JSON
// response object.
func flattenVersionHandlersScriptMap(c *Client, i interface{}) map[string]VersionHandlersScript {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionHandlersScript{}
	}

	if len(a) == 0 {
		return map[string]VersionHandlersScript{}
	}

	items := make(map[string]VersionHandlersScript)
	for k, item := range a {
		items[k] = *flattenVersionHandlersScript(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionHandlersScriptSlice flattens the contents of VersionHandlersScript from a JSON
// response object.
func flattenVersionHandlersScriptSlice(c *Client, i interface{}) []VersionHandlersScript {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersScript{}
	}

	if len(a) == 0 {
		return []VersionHandlersScript{}
	}

	items := make([]VersionHandlersScript, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersScript(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionHandlersScript expands an instance of VersionHandlersScript into a JSON
// request object.
func expandVersionHandlersScript(c *Client, f *VersionHandlersScript) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ScriptPath; !dcl.IsEmptyValueIndirect(v) {
		m["scriptPath"] = v
	}

	return m, nil
}

// flattenVersionHandlersScript flattens an instance of VersionHandlersScript from a JSON
// response object.
func flattenVersionHandlersScript(c *Client, i interface{}) *VersionHandlersScript {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionHandlersScript{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionHandlersScript
	}
	r.ScriptPath = dcl.FlattenString(m["scriptPath"])

	return r
}

// expandVersionHandlersApiEndpointMap expands the contents of VersionHandlersApiEndpoint into a JSON
// request object.
func expandVersionHandlersApiEndpointMap(c *Client, f map[string]VersionHandlersApiEndpoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionHandlersApiEndpoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionHandlersApiEndpointSlice expands the contents of VersionHandlersApiEndpoint into a JSON
// request object.
func expandVersionHandlersApiEndpointSlice(c *Client, f []VersionHandlersApiEndpoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionHandlersApiEndpoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionHandlersApiEndpointMap flattens the contents of VersionHandlersApiEndpoint from a JSON
// response object.
func flattenVersionHandlersApiEndpointMap(c *Client, i interface{}) map[string]VersionHandlersApiEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionHandlersApiEndpoint{}
	}

	if len(a) == 0 {
		return map[string]VersionHandlersApiEndpoint{}
	}

	items := make(map[string]VersionHandlersApiEndpoint)
	for k, item := range a {
		items[k] = *flattenVersionHandlersApiEndpoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionHandlersApiEndpointSlice flattens the contents of VersionHandlersApiEndpoint from a JSON
// response object.
func flattenVersionHandlersApiEndpointSlice(c *Client, i interface{}) []VersionHandlersApiEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersApiEndpoint{}
	}

	if len(a) == 0 {
		return []VersionHandlersApiEndpoint{}
	}

	items := make([]VersionHandlersApiEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersApiEndpoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionHandlersApiEndpoint expands an instance of VersionHandlersApiEndpoint into a JSON
// request object.
func expandVersionHandlersApiEndpoint(c *Client, f *VersionHandlersApiEndpoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ScriptPath; !dcl.IsEmptyValueIndirect(v) {
		m["scriptPath"] = v
	}

	return m, nil
}

// flattenVersionHandlersApiEndpoint flattens an instance of VersionHandlersApiEndpoint from a JSON
// response object.
func flattenVersionHandlersApiEndpoint(c *Client, i interface{}) *VersionHandlersApiEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionHandlersApiEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionHandlersApiEndpoint
	}
	r.ScriptPath = dcl.FlattenString(m["scriptPath"])

	return r
}

// expandVersionErrorHandlersMap expands the contents of VersionErrorHandlers into a JSON
// request object.
func expandVersionErrorHandlersMap(c *Client, f map[string]VersionErrorHandlers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionErrorHandlers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionErrorHandlersSlice expands the contents of VersionErrorHandlers into a JSON
// request object.
func expandVersionErrorHandlersSlice(c *Client, f []VersionErrorHandlers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionErrorHandlers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionErrorHandlersMap flattens the contents of VersionErrorHandlers from a JSON
// response object.
func flattenVersionErrorHandlersMap(c *Client, i interface{}) map[string]VersionErrorHandlers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionErrorHandlers{}
	}

	if len(a) == 0 {
		return map[string]VersionErrorHandlers{}
	}

	items := make(map[string]VersionErrorHandlers)
	for k, item := range a {
		items[k] = *flattenVersionErrorHandlers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionErrorHandlersSlice flattens the contents of VersionErrorHandlers from a JSON
// response object.
func flattenVersionErrorHandlersSlice(c *Client, i interface{}) []VersionErrorHandlers {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionErrorHandlers{}
	}

	if len(a) == 0 {
		return []VersionErrorHandlers{}
	}

	items := make([]VersionErrorHandlers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionErrorHandlers(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionErrorHandlers expands an instance of VersionErrorHandlers into a JSON
// request object.
func expandVersionErrorHandlers(c *Client, f *VersionErrorHandlers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ErrorCode; !dcl.IsEmptyValueIndirect(v) {
		m["errorCode"] = v
	}
	if v := f.StaticFile; !dcl.IsEmptyValueIndirect(v) {
		m["staticFile"] = v
	}
	if v := f.MimeType; !dcl.IsEmptyValueIndirect(v) {
		m["mimeType"] = v
	}

	return m, nil
}

// flattenVersionErrorHandlers flattens an instance of VersionErrorHandlers from a JSON
// response object.
func flattenVersionErrorHandlers(c *Client, i interface{}) *VersionErrorHandlers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionErrorHandlers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionErrorHandlers
	}
	r.ErrorCode = flattenVersionErrorHandlersErrorCodeEnum(m["errorCode"])
	r.StaticFile = dcl.FlattenString(m["staticFile"])
	r.MimeType = dcl.FlattenString(m["mimeType"])

	return r
}

// expandVersionLibrariesMap expands the contents of VersionLibraries into a JSON
// request object.
func expandVersionLibrariesMap(c *Client, f map[string]VersionLibraries) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionLibraries(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionLibrariesSlice expands the contents of VersionLibraries into a JSON
// request object.
func expandVersionLibrariesSlice(c *Client, f []VersionLibraries) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionLibraries(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionLibrariesMap flattens the contents of VersionLibraries from a JSON
// response object.
func flattenVersionLibrariesMap(c *Client, i interface{}) map[string]VersionLibraries {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionLibraries{}
	}

	if len(a) == 0 {
		return map[string]VersionLibraries{}
	}

	items := make(map[string]VersionLibraries)
	for k, item := range a {
		items[k] = *flattenVersionLibraries(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionLibrariesSlice flattens the contents of VersionLibraries from a JSON
// response object.
func flattenVersionLibrariesSlice(c *Client, i interface{}) []VersionLibraries {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionLibraries{}
	}

	if len(a) == 0 {
		return []VersionLibraries{}
	}

	items := make([]VersionLibraries, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionLibraries(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionLibraries expands an instance of VersionLibraries into a JSON
// request object.
func expandVersionLibraries(c *Client, f *VersionLibraries) (map[string]interface{}, error) {
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

	return m, nil
}

// flattenVersionLibraries flattens an instance of VersionLibraries from a JSON
// response object.
func flattenVersionLibraries(c *Client, i interface{}) *VersionLibraries {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionLibraries{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionLibraries
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandVersionApiConfigMap expands the contents of VersionApiConfig into a JSON
// request object.
func expandVersionApiConfigMap(c *Client, f map[string]VersionApiConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionApiConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionApiConfigSlice expands the contents of VersionApiConfig into a JSON
// request object.
func expandVersionApiConfigSlice(c *Client, f []VersionApiConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionApiConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionApiConfigMap flattens the contents of VersionApiConfig from a JSON
// response object.
func flattenVersionApiConfigMap(c *Client, i interface{}) map[string]VersionApiConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionApiConfig{}
	}

	if len(a) == 0 {
		return map[string]VersionApiConfig{}
	}

	items := make(map[string]VersionApiConfig)
	for k, item := range a {
		items[k] = *flattenVersionApiConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionApiConfigSlice flattens the contents of VersionApiConfig from a JSON
// response object.
func flattenVersionApiConfigSlice(c *Client, i interface{}) []VersionApiConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionApiConfig{}
	}

	if len(a) == 0 {
		return []VersionApiConfig{}
	}

	items := make([]VersionApiConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionApiConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionApiConfig expands an instance of VersionApiConfig into a JSON
// request object.
func expandVersionApiConfig(c *Client, f *VersionApiConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AuthFailAction; !dcl.IsEmptyValueIndirect(v) {
		m["authFailAction"] = v
	}
	if v := f.Login; !dcl.IsEmptyValueIndirect(v) {
		m["login"] = v
	}
	if v := f.Script; !dcl.IsEmptyValueIndirect(v) {
		m["script"] = v
	}
	if v := f.SecurityLevel; !dcl.IsEmptyValueIndirect(v) {
		m["securityLevel"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenVersionApiConfig flattens an instance of VersionApiConfig from a JSON
// response object.
func flattenVersionApiConfig(c *Client, i interface{}) *VersionApiConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionApiConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionApiConfig
	}
	r.AuthFailAction = flattenVersionApiConfigAuthFailActionEnum(m["authFailAction"])
	r.Login = flattenVersionApiConfigLoginEnum(m["login"])
	r.Script = dcl.FlattenString(m["script"])
	r.SecurityLevel = flattenVersionApiConfigSecurityLevelEnum(m["securityLevel"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandVersionDeploymentMap expands the contents of VersionDeployment into a JSON
// request object.
func expandVersionDeploymentMap(c *Client, f map[string]VersionDeployment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionDeployment(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionDeploymentSlice expands the contents of VersionDeployment into a JSON
// request object.
func expandVersionDeploymentSlice(c *Client, f []VersionDeployment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionDeployment(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionDeploymentMap flattens the contents of VersionDeployment from a JSON
// response object.
func flattenVersionDeploymentMap(c *Client, i interface{}) map[string]VersionDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionDeployment{}
	}

	if len(a) == 0 {
		return map[string]VersionDeployment{}
	}

	items := make(map[string]VersionDeployment)
	for k, item := range a {
		items[k] = *flattenVersionDeployment(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionDeploymentSlice flattens the contents of VersionDeployment from a JSON
// response object.
func flattenVersionDeploymentSlice(c *Client, i interface{}) []VersionDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionDeployment{}
	}

	if len(a) == 0 {
		return []VersionDeployment{}
	}

	items := make([]VersionDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionDeployment(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionDeployment expands an instance of VersionDeployment into a JSON
// request object.
func expandVersionDeployment(c *Client, f *VersionDeployment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandVersionDeploymentFilesMap(c, f.Files); err != nil {
		return nil, fmt.Errorf("error expanding Files into files: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["files"] = v
	}
	if v, err := expandVersionDeploymentContainer(c, f.Container); err != nil {
		return nil, fmt.Errorf("error expanding Container into container: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["container"] = v
	}
	if v, err := expandVersionDeploymentZip(c, f.Zip); err != nil {
		return nil, fmt.Errorf("error expanding Zip into zip: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zip"] = v
	}
	if v, err := expandVersionDeploymentCloudBuildOptions(c, f.CloudBuildOptions); err != nil {
		return nil, fmt.Errorf("error expanding CloudBuildOptions into cloudBuildOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudBuildOptions"] = v
	}

	return m, nil
}

// flattenVersionDeployment flattens an instance of VersionDeployment from a JSON
// response object.
func flattenVersionDeployment(c *Client, i interface{}) *VersionDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionDeployment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionDeployment
	}
	r.Files = flattenVersionDeploymentFilesMap(c, m["files"])
	r.Container = flattenVersionDeploymentContainer(c, m["container"])
	r.Zip = flattenVersionDeploymentZip(c, m["zip"])
	r.CloudBuildOptions = flattenVersionDeploymentCloudBuildOptions(c, m["cloudBuildOptions"])

	return r
}

// expandVersionDeploymentFilesMap expands the contents of VersionDeploymentFiles into a JSON
// request object.
func expandVersionDeploymentFilesMap(c *Client, f map[string]VersionDeploymentFiles) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionDeploymentFiles(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionDeploymentFilesSlice expands the contents of VersionDeploymentFiles into a JSON
// request object.
func expandVersionDeploymentFilesSlice(c *Client, f []VersionDeploymentFiles) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionDeploymentFiles(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionDeploymentFilesMap flattens the contents of VersionDeploymentFiles from a JSON
// response object.
func flattenVersionDeploymentFilesMap(c *Client, i interface{}) map[string]VersionDeploymentFiles {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionDeploymentFiles{}
	}

	if len(a) == 0 {
		return map[string]VersionDeploymentFiles{}
	}

	items := make(map[string]VersionDeploymentFiles)
	for k, item := range a {
		items[k] = *flattenVersionDeploymentFiles(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionDeploymentFilesSlice flattens the contents of VersionDeploymentFiles from a JSON
// response object.
func flattenVersionDeploymentFilesSlice(c *Client, i interface{}) []VersionDeploymentFiles {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionDeploymentFiles{}
	}

	if len(a) == 0 {
		return []VersionDeploymentFiles{}
	}

	items := make([]VersionDeploymentFiles, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionDeploymentFiles(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionDeploymentFiles expands an instance of VersionDeploymentFiles into a JSON
// request object.
func expandVersionDeploymentFiles(c *Client, f *VersionDeploymentFiles) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SourceUrl; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUrl"] = v
	}
	if v := f.Sha1Sum; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Sum"] = v
	}
	if v := f.MimeType; !dcl.IsEmptyValueIndirect(v) {
		m["mimeType"] = v
	}

	return m, nil
}

// flattenVersionDeploymentFiles flattens an instance of VersionDeploymentFiles from a JSON
// response object.
func flattenVersionDeploymentFiles(c *Client, i interface{}) *VersionDeploymentFiles {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionDeploymentFiles{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionDeploymentFiles
	}
	r.SourceUrl = dcl.FlattenString(m["sourceUrl"])
	r.Sha1Sum = dcl.FlattenString(m["sha1Sum"])
	r.MimeType = dcl.FlattenString(m["mimeType"])

	return r
}

// expandVersionDeploymentContainerMap expands the contents of VersionDeploymentContainer into a JSON
// request object.
func expandVersionDeploymentContainerMap(c *Client, f map[string]VersionDeploymentContainer) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionDeploymentContainer(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionDeploymentContainerSlice expands the contents of VersionDeploymentContainer into a JSON
// request object.
func expandVersionDeploymentContainerSlice(c *Client, f []VersionDeploymentContainer) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionDeploymentContainer(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionDeploymentContainerMap flattens the contents of VersionDeploymentContainer from a JSON
// response object.
func flattenVersionDeploymentContainerMap(c *Client, i interface{}) map[string]VersionDeploymentContainer {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionDeploymentContainer{}
	}

	if len(a) == 0 {
		return map[string]VersionDeploymentContainer{}
	}

	items := make(map[string]VersionDeploymentContainer)
	for k, item := range a {
		items[k] = *flattenVersionDeploymentContainer(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionDeploymentContainerSlice flattens the contents of VersionDeploymentContainer from a JSON
// response object.
func flattenVersionDeploymentContainerSlice(c *Client, i interface{}) []VersionDeploymentContainer {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionDeploymentContainer{}
	}

	if len(a) == 0 {
		return []VersionDeploymentContainer{}
	}

	items := make([]VersionDeploymentContainer, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionDeploymentContainer(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionDeploymentContainer expands an instance of VersionDeploymentContainer into a JSON
// request object.
func expandVersionDeploymentContainer(c *Client, f *VersionDeploymentContainer) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["image"] = v
	}

	return m, nil
}

// flattenVersionDeploymentContainer flattens an instance of VersionDeploymentContainer from a JSON
// response object.
func flattenVersionDeploymentContainer(c *Client, i interface{}) *VersionDeploymentContainer {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionDeploymentContainer{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionDeploymentContainer
	}
	r.Image = dcl.FlattenString(m["image"])

	return r
}

// expandVersionDeploymentZipMap expands the contents of VersionDeploymentZip into a JSON
// request object.
func expandVersionDeploymentZipMap(c *Client, f map[string]VersionDeploymentZip) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionDeploymentZip(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionDeploymentZipSlice expands the contents of VersionDeploymentZip into a JSON
// request object.
func expandVersionDeploymentZipSlice(c *Client, f []VersionDeploymentZip) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionDeploymentZip(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionDeploymentZipMap flattens the contents of VersionDeploymentZip from a JSON
// response object.
func flattenVersionDeploymentZipMap(c *Client, i interface{}) map[string]VersionDeploymentZip {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionDeploymentZip{}
	}

	if len(a) == 0 {
		return map[string]VersionDeploymentZip{}
	}

	items := make(map[string]VersionDeploymentZip)
	for k, item := range a {
		items[k] = *flattenVersionDeploymentZip(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionDeploymentZipSlice flattens the contents of VersionDeploymentZip from a JSON
// response object.
func flattenVersionDeploymentZipSlice(c *Client, i interface{}) []VersionDeploymentZip {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionDeploymentZip{}
	}

	if len(a) == 0 {
		return []VersionDeploymentZip{}
	}

	items := make([]VersionDeploymentZip, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionDeploymentZip(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionDeploymentZip expands an instance of VersionDeploymentZip into a JSON
// request object.
func expandVersionDeploymentZip(c *Client, f *VersionDeploymentZip) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SourceUrl; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUrl"] = v
	}
	if v := f.FilesCount; !dcl.IsEmptyValueIndirect(v) {
		m["filesCount"] = v
	}

	return m, nil
}

// flattenVersionDeploymentZip flattens an instance of VersionDeploymentZip from a JSON
// response object.
func flattenVersionDeploymentZip(c *Client, i interface{}) *VersionDeploymentZip {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionDeploymentZip{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionDeploymentZip
	}
	r.SourceUrl = dcl.FlattenString(m["sourceUrl"])
	r.FilesCount = dcl.FlattenInteger(m["filesCount"])

	return r
}

// expandVersionDeploymentCloudBuildOptionsMap expands the contents of VersionDeploymentCloudBuildOptions into a JSON
// request object.
func expandVersionDeploymentCloudBuildOptionsMap(c *Client, f map[string]VersionDeploymentCloudBuildOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionDeploymentCloudBuildOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionDeploymentCloudBuildOptionsSlice expands the contents of VersionDeploymentCloudBuildOptions into a JSON
// request object.
func expandVersionDeploymentCloudBuildOptionsSlice(c *Client, f []VersionDeploymentCloudBuildOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionDeploymentCloudBuildOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionDeploymentCloudBuildOptionsMap flattens the contents of VersionDeploymentCloudBuildOptions from a JSON
// response object.
func flattenVersionDeploymentCloudBuildOptionsMap(c *Client, i interface{}) map[string]VersionDeploymentCloudBuildOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionDeploymentCloudBuildOptions{}
	}

	if len(a) == 0 {
		return map[string]VersionDeploymentCloudBuildOptions{}
	}

	items := make(map[string]VersionDeploymentCloudBuildOptions)
	for k, item := range a {
		items[k] = *flattenVersionDeploymentCloudBuildOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionDeploymentCloudBuildOptionsSlice flattens the contents of VersionDeploymentCloudBuildOptions from a JSON
// response object.
func flattenVersionDeploymentCloudBuildOptionsSlice(c *Client, i interface{}) []VersionDeploymentCloudBuildOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionDeploymentCloudBuildOptions{}
	}

	if len(a) == 0 {
		return []VersionDeploymentCloudBuildOptions{}
	}

	items := make([]VersionDeploymentCloudBuildOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionDeploymentCloudBuildOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionDeploymentCloudBuildOptions expands an instance of VersionDeploymentCloudBuildOptions into a JSON
// request object.
func expandVersionDeploymentCloudBuildOptions(c *Client, f *VersionDeploymentCloudBuildOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AppYamlPath; !dcl.IsEmptyValueIndirect(v) {
		m["appYamlPath"] = v
	}
	if v := f.CloudBuildTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["cloudBuildTimeout"] = v
	}

	return m, nil
}

// flattenVersionDeploymentCloudBuildOptions flattens an instance of VersionDeploymentCloudBuildOptions from a JSON
// response object.
func flattenVersionDeploymentCloudBuildOptions(c *Client, i interface{}) *VersionDeploymentCloudBuildOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionDeploymentCloudBuildOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionDeploymentCloudBuildOptions
	}
	r.AppYamlPath = dcl.FlattenString(m["appYamlPath"])
	r.CloudBuildTimeout = dcl.FlattenString(m["cloudBuildTimeout"])

	return r
}

// expandVersionHealthCheckMap expands the contents of VersionHealthCheck into a JSON
// request object.
func expandVersionHealthCheckMap(c *Client, f map[string]VersionHealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionHealthCheckSlice expands the contents of VersionHealthCheck into a JSON
// request object.
func expandVersionHealthCheckSlice(c *Client, f []VersionHealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionHealthCheckMap flattens the contents of VersionHealthCheck from a JSON
// response object.
func flattenVersionHealthCheckMap(c *Client, i interface{}) map[string]VersionHealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionHealthCheck{}
	}

	if len(a) == 0 {
		return map[string]VersionHealthCheck{}
	}

	items := make(map[string]VersionHealthCheck)
	for k, item := range a {
		items[k] = *flattenVersionHealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionHealthCheckSlice flattens the contents of VersionHealthCheck from a JSON
// response object.
func flattenVersionHealthCheckSlice(c *Client, i interface{}) []VersionHealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHealthCheck{}
	}

	if len(a) == 0 {
		return []VersionHealthCheck{}
	}

	items := make([]VersionHealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionHealthCheck expands an instance of VersionHealthCheck into a JSON
// request object.
func expandVersionHealthCheck(c *Client, f *VersionHealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DisableHealthCheck; !dcl.IsEmptyValueIndirect(v) {
		m["disableHealthCheck"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.HealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["healthyThreshold"] = v
	}
	if v := f.UnhealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["unhealthyThreshold"] = v
	}
	if v := f.RestartThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["restartThreshold"] = v
	}
	if v := f.CheckInterval; !dcl.IsEmptyValueIndirect(v) {
		m["checkInterval"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}

	return m, nil
}

// flattenVersionHealthCheck flattens an instance of VersionHealthCheck from a JSON
// response object.
func flattenVersionHealthCheck(c *Client, i interface{}) *VersionHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionHealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionHealthCheck
	}
	r.DisableHealthCheck = dcl.FlattenBool(m["disableHealthCheck"])
	r.Host = dcl.FlattenString(m["host"])
	r.HealthyThreshold = dcl.FlattenInteger(m["healthyThreshold"])
	r.UnhealthyThreshold = dcl.FlattenInteger(m["unhealthyThreshold"])
	r.RestartThreshold = dcl.FlattenInteger(m["restartThreshold"])
	r.CheckInterval = dcl.FlattenString(m["checkInterval"])
	r.Timeout = dcl.FlattenString(m["timeout"])

	return r
}

// expandVersionReadinessCheckMap expands the contents of VersionReadinessCheck into a JSON
// request object.
func expandVersionReadinessCheckMap(c *Client, f map[string]VersionReadinessCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionReadinessCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionReadinessCheckSlice expands the contents of VersionReadinessCheck into a JSON
// request object.
func expandVersionReadinessCheckSlice(c *Client, f []VersionReadinessCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionReadinessCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionReadinessCheckMap flattens the contents of VersionReadinessCheck from a JSON
// response object.
func flattenVersionReadinessCheckMap(c *Client, i interface{}) map[string]VersionReadinessCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionReadinessCheck{}
	}

	if len(a) == 0 {
		return map[string]VersionReadinessCheck{}
	}

	items := make(map[string]VersionReadinessCheck)
	for k, item := range a {
		items[k] = *flattenVersionReadinessCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionReadinessCheckSlice flattens the contents of VersionReadinessCheck from a JSON
// response object.
func flattenVersionReadinessCheckSlice(c *Client, i interface{}) []VersionReadinessCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionReadinessCheck{}
	}

	if len(a) == 0 {
		return []VersionReadinessCheck{}
	}

	items := make([]VersionReadinessCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionReadinessCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionReadinessCheck expands an instance of VersionReadinessCheck into a JSON
// request object.
func expandVersionReadinessCheck(c *Client, f *VersionReadinessCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.FailureThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["failureThreshold"] = v
	}
	if v := f.SuccessThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["successThreshold"] = v
	}
	if v := f.CheckInterval; !dcl.IsEmptyValueIndirect(v) {
		m["checkInterval"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.AppStartTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["appStartTimeout"] = v
	}

	return m, nil
}

// flattenVersionReadinessCheck flattens an instance of VersionReadinessCheck from a JSON
// response object.
func flattenVersionReadinessCheck(c *Client, i interface{}) *VersionReadinessCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionReadinessCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionReadinessCheck
	}
	r.Path = dcl.FlattenString(m["path"])
	r.Host = dcl.FlattenString(m["host"])
	r.FailureThreshold = dcl.FlattenInteger(m["failureThreshold"])
	r.SuccessThreshold = dcl.FlattenInteger(m["successThreshold"])
	r.CheckInterval = dcl.FlattenString(m["checkInterval"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.AppStartTimeout = dcl.FlattenString(m["appStartTimeout"])

	return r
}

// expandVersionLivenessCheckMap expands the contents of VersionLivenessCheck into a JSON
// request object.
func expandVersionLivenessCheckMap(c *Client, f map[string]VersionLivenessCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionLivenessCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionLivenessCheckSlice expands the contents of VersionLivenessCheck into a JSON
// request object.
func expandVersionLivenessCheckSlice(c *Client, f []VersionLivenessCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionLivenessCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionLivenessCheckMap flattens the contents of VersionLivenessCheck from a JSON
// response object.
func flattenVersionLivenessCheckMap(c *Client, i interface{}) map[string]VersionLivenessCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionLivenessCheck{}
	}

	if len(a) == 0 {
		return map[string]VersionLivenessCheck{}
	}

	items := make(map[string]VersionLivenessCheck)
	for k, item := range a {
		items[k] = *flattenVersionLivenessCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionLivenessCheckSlice flattens the contents of VersionLivenessCheck from a JSON
// response object.
func flattenVersionLivenessCheckSlice(c *Client, i interface{}) []VersionLivenessCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionLivenessCheck{}
	}

	if len(a) == 0 {
		return []VersionLivenessCheck{}
	}

	items := make([]VersionLivenessCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionLivenessCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionLivenessCheck expands an instance of VersionLivenessCheck into a JSON
// request object.
func expandVersionLivenessCheck(c *Client, f *VersionLivenessCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.FailureThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["failureThreshold"] = v
	}
	if v := f.SuccessThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["successThreshold"] = v
	}
	if v := f.CheckInterval; !dcl.IsEmptyValueIndirect(v) {
		m["checkInterval"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.InitialDelay; !dcl.IsEmptyValueIndirect(v) {
		m["initialDelay"] = v
	}

	return m, nil
}

// flattenVersionLivenessCheck flattens an instance of VersionLivenessCheck from a JSON
// response object.
func flattenVersionLivenessCheck(c *Client, i interface{}) *VersionLivenessCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionLivenessCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionLivenessCheck
	}
	r.Path = dcl.FlattenString(m["path"])
	r.Host = dcl.FlattenString(m["host"])
	r.FailureThreshold = dcl.FlattenInteger(m["failureThreshold"])
	r.SuccessThreshold = dcl.FlattenInteger(m["successThreshold"])
	r.CheckInterval = dcl.FlattenString(m["checkInterval"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.InitialDelay = dcl.FlattenString(m["initialDelay"])

	return r
}

// expandVersionEntrypointMap expands the contents of VersionEntrypoint into a JSON
// request object.
func expandVersionEntrypointMap(c *Client, f map[string]VersionEntrypoint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionEntrypoint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionEntrypointSlice expands the contents of VersionEntrypoint into a JSON
// request object.
func expandVersionEntrypointSlice(c *Client, f []VersionEntrypoint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionEntrypoint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionEntrypointMap flattens the contents of VersionEntrypoint from a JSON
// response object.
func flattenVersionEntrypointMap(c *Client, i interface{}) map[string]VersionEntrypoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionEntrypoint{}
	}

	if len(a) == 0 {
		return map[string]VersionEntrypoint{}
	}

	items := make(map[string]VersionEntrypoint)
	for k, item := range a {
		items[k] = *flattenVersionEntrypoint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionEntrypointSlice flattens the contents of VersionEntrypoint from a JSON
// response object.
func flattenVersionEntrypointSlice(c *Client, i interface{}) []VersionEntrypoint {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionEntrypoint{}
	}

	if len(a) == 0 {
		return []VersionEntrypoint{}
	}

	items := make([]VersionEntrypoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionEntrypoint(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionEntrypoint expands an instance of VersionEntrypoint into a JSON
// request object.
func expandVersionEntrypoint(c *Client, f *VersionEntrypoint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Shell; !dcl.IsEmptyValueIndirect(v) {
		m["shell"] = v
	}

	return m, nil
}

// flattenVersionEntrypoint flattens an instance of VersionEntrypoint from a JSON
// response object.
func flattenVersionEntrypoint(c *Client, i interface{}) *VersionEntrypoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionEntrypoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionEntrypoint
	}
	r.Shell = dcl.FlattenString(m["shell"])

	return r
}

// expandVersionVPCAccessConnectorMap expands the contents of VersionVPCAccessConnector into a JSON
// request object.
func expandVersionVPCAccessConnectorMap(c *Client, f map[string]VersionVPCAccessConnector) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVersionVPCAccessConnector(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVersionVPCAccessConnectorSlice expands the contents of VersionVPCAccessConnector into a JSON
// request object.
func expandVersionVPCAccessConnectorSlice(c *Client, f []VersionVPCAccessConnector) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVersionVPCAccessConnector(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVersionVPCAccessConnectorMap flattens the contents of VersionVPCAccessConnector from a JSON
// response object.
func flattenVersionVPCAccessConnectorMap(c *Client, i interface{}) map[string]VersionVPCAccessConnector {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VersionVPCAccessConnector{}
	}

	if len(a) == 0 {
		return map[string]VersionVPCAccessConnector{}
	}

	items := make(map[string]VersionVPCAccessConnector)
	for k, item := range a {
		items[k] = *flattenVersionVPCAccessConnector(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVersionVPCAccessConnectorSlice flattens the contents of VersionVPCAccessConnector from a JSON
// response object.
func flattenVersionVPCAccessConnectorSlice(c *Client, i interface{}) []VersionVPCAccessConnector {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionVPCAccessConnector{}
	}

	if len(a) == 0 {
		return []VersionVPCAccessConnector{}
	}

	items := make([]VersionVPCAccessConnector, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionVPCAccessConnector(c, item.(map[string]interface{})))
	}

	return items
}

// expandVersionVPCAccessConnector expands an instance of VersionVPCAccessConnector into a JSON
// request object.
func expandVersionVPCAccessConnector(c *Client, f *VersionVPCAccessConnector) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenVersionVPCAccessConnector flattens an instance of VersionVPCAccessConnector from a JSON
// response object.
func flattenVersionVPCAccessConnector(c *Client, i interface{}) *VersionVPCAccessConnector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VersionVPCAccessConnector{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVersionVPCAccessConnector
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// flattenVersionInboundServicesEnumSlice flattens the contents of VersionInboundServicesEnum from a JSON
// response object.
func flattenVersionInboundServicesEnumSlice(c *Client, i interface{}) []VersionInboundServicesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionInboundServicesEnum{}
	}

	if len(a) == 0 {
		return []VersionInboundServicesEnum{}
	}

	items := make([]VersionInboundServicesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionInboundServicesEnum(item.(interface{})))
	}

	return items
}

// flattenVersionInboundServicesEnum asserts that an interface is a string, and returns a
// pointer to a *VersionInboundServicesEnum with the same value as that string.
func flattenVersionInboundServicesEnum(i interface{}) *VersionInboundServicesEnum {
	s, ok := i.(string)
	if !ok {
		return VersionInboundServicesEnumRef("")
	}

	return VersionInboundServicesEnumRef(s)
}

// flattenVersionServingStatusEnumSlice flattens the contents of VersionServingStatusEnum from a JSON
// response object.
func flattenVersionServingStatusEnumSlice(c *Client, i interface{}) []VersionServingStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionServingStatusEnum{}
	}

	if len(a) == 0 {
		return []VersionServingStatusEnum{}
	}

	items := make([]VersionServingStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionServingStatusEnum(item.(interface{})))
	}

	return items
}

// flattenVersionServingStatusEnum asserts that an interface is a string, and returns a
// pointer to a *VersionServingStatusEnum with the same value as that string.
func flattenVersionServingStatusEnum(i interface{}) *VersionServingStatusEnum {
	s, ok := i.(string)
	if !ok {
		return VersionServingStatusEnumRef("")
	}

	return VersionServingStatusEnumRef(s)
}

// flattenVersionHandlersSecurityLevelEnumSlice flattens the contents of VersionHandlersSecurityLevelEnum from a JSON
// response object.
func flattenVersionHandlersSecurityLevelEnumSlice(c *Client, i interface{}) []VersionHandlersSecurityLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersSecurityLevelEnum{}
	}

	if len(a) == 0 {
		return []VersionHandlersSecurityLevelEnum{}
	}

	items := make([]VersionHandlersSecurityLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersSecurityLevelEnum(item.(interface{})))
	}

	return items
}

// flattenVersionHandlersSecurityLevelEnum asserts that an interface is a string, and returns a
// pointer to a *VersionHandlersSecurityLevelEnum with the same value as that string.
func flattenVersionHandlersSecurityLevelEnum(i interface{}) *VersionHandlersSecurityLevelEnum {
	s, ok := i.(string)
	if !ok {
		return VersionHandlersSecurityLevelEnumRef("")
	}

	return VersionHandlersSecurityLevelEnumRef(s)
}

// flattenVersionHandlersLoginEnumSlice flattens the contents of VersionHandlersLoginEnum from a JSON
// response object.
func flattenVersionHandlersLoginEnumSlice(c *Client, i interface{}) []VersionHandlersLoginEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersLoginEnum{}
	}

	if len(a) == 0 {
		return []VersionHandlersLoginEnum{}
	}

	items := make([]VersionHandlersLoginEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersLoginEnum(item.(interface{})))
	}

	return items
}

// flattenVersionHandlersLoginEnum asserts that an interface is a string, and returns a
// pointer to a *VersionHandlersLoginEnum with the same value as that string.
func flattenVersionHandlersLoginEnum(i interface{}) *VersionHandlersLoginEnum {
	s, ok := i.(string)
	if !ok {
		return VersionHandlersLoginEnumRef("")
	}

	return VersionHandlersLoginEnumRef(s)
}

// flattenVersionHandlersAuthFailActionEnumSlice flattens the contents of VersionHandlersAuthFailActionEnum from a JSON
// response object.
func flattenVersionHandlersAuthFailActionEnumSlice(c *Client, i interface{}) []VersionHandlersAuthFailActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersAuthFailActionEnum{}
	}

	if len(a) == 0 {
		return []VersionHandlersAuthFailActionEnum{}
	}

	items := make([]VersionHandlersAuthFailActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersAuthFailActionEnum(item.(interface{})))
	}

	return items
}

// flattenVersionHandlersAuthFailActionEnum asserts that an interface is a string, and returns a
// pointer to a *VersionHandlersAuthFailActionEnum with the same value as that string.
func flattenVersionHandlersAuthFailActionEnum(i interface{}) *VersionHandlersAuthFailActionEnum {
	s, ok := i.(string)
	if !ok {
		return VersionHandlersAuthFailActionEnumRef("")
	}

	return VersionHandlersAuthFailActionEnumRef(s)
}

// flattenVersionHandlersRedirectHttpResponseCodeEnumSlice flattens the contents of VersionHandlersRedirectHttpResponseCodeEnum from a JSON
// response object.
func flattenVersionHandlersRedirectHttpResponseCodeEnumSlice(c *Client, i interface{}) []VersionHandlersRedirectHttpResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionHandlersRedirectHttpResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []VersionHandlersRedirectHttpResponseCodeEnum{}
	}

	items := make([]VersionHandlersRedirectHttpResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionHandlersRedirectHttpResponseCodeEnum(item.(interface{})))
	}

	return items
}

// flattenVersionHandlersRedirectHttpResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *VersionHandlersRedirectHttpResponseCodeEnum with the same value as that string.
func flattenVersionHandlersRedirectHttpResponseCodeEnum(i interface{}) *VersionHandlersRedirectHttpResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return VersionHandlersRedirectHttpResponseCodeEnumRef("")
	}

	return VersionHandlersRedirectHttpResponseCodeEnumRef(s)
}

// flattenVersionErrorHandlersErrorCodeEnumSlice flattens the contents of VersionErrorHandlersErrorCodeEnum from a JSON
// response object.
func flattenVersionErrorHandlersErrorCodeEnumSlice(c *Client, i interface{}) []VersionErrorHandlersErrorCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionErrorHandlersErrorCodeEnum{}
	}

	if len(a) == 0 {
		return []VersionErrorHandlersErrorCodeEnum{}
	}

	items := make([]VersionErrorHandlersErrorCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionErrorHandlersErrorCodeEnum(item.(interface{})))
	}

	return items
}

// flattenVersionErrorHandlersErrorCodeEnum asserts that an interface is a string, and returns a
// pointer to a *VersionErrorHandlersErrorCodeEnum with the same value as that string.
func flattenVersionErrorHandlersErrorCodeEnum(i interface{}) *VersionErrorHandlersErrorCodeEnum {
	s, ok := i.(string)
	if !ok {
		return VersionErrorHandlersErrorCodeEnumRef("")
	}

	return VersionErrorHandlersErrorCodeEnumRef(s)
}

// flattenVersionApiConfigAuthFailActionEnumSlice flattens the contents of VersionApiConfigAuthFailActionEnum from a JSON
// response object.
func flattenVersionApiConfigAuthFailActionEnumSlice(c *Client, i interface{}) []VersionApiConfigAuthFailActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionApiConfigAuthFailActionEnum{}
	}

	if len(a) == 0 {
		return []VersionApiConfigAuthFailActionEnum{}
	}

	items := make([]VersionApiConfigAuthFailActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionApiConfigAuthFailActionEnum(item.(interface{})))
	}

	return items
}

// flattenVersionApiConfigAuthFailActionEnum asserts that an interface is a string, and returns a
// pointer to a *VersionApiConfigAuthFailActionEnum with the same value as that string.
func flattenVersionApiConfigAuthFailActionEnum(i interface{}) *VersionApiConfigAuthFailActionEnum {
	s, ok := i.(string)
	if !ok {
		return VersionApiConfigAuthFailActionEnumRef("")
	}

	return VersionApiConfigAuthFailActionEnumRef(s)
}

// flattenVersionApiConfigLoginEnumSlice flattens the contents of VersionApiConfigLoginEnum from a JSON
// response object.
func flattenVersionApiConfigLoginEnumSlice(c *Client, i interface{}) []VersionApiConfigLoginEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionApiConfigLoginEnum{}
	}

	if len(a) == 0 {
		return []VersionApiConfigLoginEnum{}
	}

	items := make([]VersionApiConfigLoginEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionApiConfigLoginEnum(item.(interface{})))
	}

	return items
}

// flattenVersionApiConfigLoginEnum asserts that an interface is a string, and returns a
// pointer to a *VersionApiConfigLoginEnum with the same value as that string.
func flattenVersionApiConfigLoginEnum(i interface{}) *VersionApiConfigLoginEnum {
	s, ok := i.(string)
	if !ok {
		return VersionApiConfigLoginEnumRef("")
	}

	return VersionApiConfigLoginEnumRef(s)
}

// flattenVersionApiConfigSecurityLevelEnumSlice flattens the contents of VersionApiConfigSecurityLevelEnum from a JSON
// response object.
func flattenVersionApiConfigSecurityLevelEnumSlice(c *Client, i interface{}) []VersionApiConfigSecurityLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VersionApiConfigSecurityLevelEnum{}
	}

	if len(a) == 0 {
		return []VersionApiConfigSecurityLevelEnum{}
	}

	items := make([]VersionApiConfigSecurityLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVersionApiConfigSecurityLevelEnum(item.(interface{})))
	}

	return items
}

// flattenVersionApiConfigSecurityLevelEnum asserts that an interface is a string, and returns a
// pointer to a *VersionApiConfigSecurityLevelEnum with the same value as that string.
func flattenVersionApiConfigSecurityLevelEnum(i interface{}) *VersionApiConfigSecurityLevelEnum {
	s, ok := i.(string)
	if !ok {
		return VersionApiConfigSecurityLevelEnumRef("")
	}

	return VersionApiConfigSecurityLevelEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Version) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalVersion(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.App == nil && ncr.App == nil {
			c.Config.Logger.Info("Both App fields null - considering equal.")
		} else if nr.App == nil || ncr.App == nil {
			c.Config.Logger.Info("Only one App field is null - considering unequal.")
			return false
		} else if *nr.App != *ncr.App {
			return false
		}
		if nr.Service == nil && ncr.Service == nil {
			c.Config.Logger.Info("Both Service fields null - considering equal.")
		} else if nr.Service == nil || ncr.Service == nil {
			c.Config.Logger.Info("Only one Service field is null - considering unequal.")
			return false
		} else if *nr.Service != *ncr.Service {
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

type versionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         versionApiOperation
}

func convertFieldDiffToVersionOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]versionDiff, error) {
	var diffs []versionDiff
	for _, op := range ops {
		diff := versionDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToversionApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToversionApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (versionApiOperation, error) {
	switch op {

	case "updateVersionPatchVersionOperation":
		return &updateVersionPatchVersionOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
