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
package monitoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Service) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Custom", "AppEngine", "CloudEndpoints", "ClusterIstio", "MeshIstio", "IstioCanonicalService", "CloudRun", "GkeNamespace", "GkeWorkload", "GkeService"}, r.Custom, r.AppEngine, r.CloudEndpoints, r.ClusterIstio, r.MeshIstio, r.IstioCanonicalService, r.CloudRun, r.GkeNamespace, r.GkeWorkload, r.GkeService); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Custom) {
		if err := r.Custom.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AppEngine) {
		if err := r.AppEngine.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudEndpoints) {
		if err := r.CloudEndpoints.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ClusterIstio) {
		if err := r.ClusterIstio.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MeshIstio) {
		if err := r.MeshIstio.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IstioCanonicalService) {
		if err := r.IstioCanonicalService.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudRun) {
		if err := r.CloudRun.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GkeNamespace) {
		if err := r.GkeNamespace.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GkeWorkload) {
		if err := r.GkeWorkload.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GkeService) {
		if err := r.GkeService.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Telemetry) {
		if err := r.Telemetry.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceCustom) validate() error {
	return nil
}
func (r *ServiceAppEngine) validate() error {
	return nil
}
func (r *ServiceCloudEndpoints) validate() error {
	return nil
}
func (r *ServiceClusterIstio) validate() error {
	return nil
}
func (r *ServiceMeshIstio) validate() error {
	return nil
}
func (r *ServiceIstioCanonicalService) validate() error {
	return nil
}
func (r *ServiceCloudRun) validate() error {
	return nil
}
func (r *ServiceGkeNamespace) validate() error {
	return nil
}
func (r *ServiceGkeWorkload) validate() error {
	return nil
}
func (r *ServiceGkeService) validate() error {
	return nil
}
func (r *ServiceTelemetry) validate() error {
	return nil
}
func (r *Service) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://monitoring.googleapis.com/v3/", params)
}

func (r *Service) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Service) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/services", nr.basePath(), userBasePath, params), nil

}

func (r *Service) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/services?serviceId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Service) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

// serviceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceApiOperation interface {
	do(context.Context, *Service, *Client) error
}

// newUpdateServiceUpdateServiceRequest creates a request for an
// Service resource's UpdateService update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceUpdateServiceRequest(ctx context.Context, f *Service, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/services/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandServiceCustom(c, f.Custom); err != nil {
		return nil, fmt.Errorf("error expanding Custom into custom: %w", err)
	} else if v != nil {
		req["custom"] = v
	}
	if v, err := expandServiceAppEngine(c, f.AppEngine); err != nil {
		return nil, fmt.Errorf("error expanding AppEngine into appEngine: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["appEngine"] = v
	}
	if v, err := expandServiceCloudEndpoints(c, f.CloudEndpoints); err != nil {
		return nil, fmt.Errorf("error expanding CloudEndpoints into cloudEndpoints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cloudEndpoints"] = v
	}
	if v, err := expandServiceClusterIstio(c, f.ClusterIstio); err != nil {
		return nil, fmt.Errorf("error expanding ClusterIstio into clusterIstio: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clusterIstio"] = v
	}
	if v, err := expandServiceMeshIstio(c, f.MeshIstio); err != nil {
		return nil, fmt.Errorf("error expanding MeshIstio into meshIstio: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["meshIstio"] = v
	}
	if v, err := expandServiceIstioCanonicalService(c, f.IstioCanonicalService); err != nil {
		return nil, fmt.Errorf("error expanding IstioCanonicalService into istioCanonicalService: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["istioCanonicalService"] = v
	}
	if v, err := expandServiceCloudRun(c, f.CloudRun); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cloudRun"] = v
	}
	if v, err := expandServiceGkeNamespace(c, f.GkeNamespace); err != nil {
		return nil, fmt.Errorf("error expanding GkeNamespace into gkeNamespace: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["gkeNamespace"] = v
	}
	if v, err := expandServiceGkeWorkload(c, f.GkeWorkload); err != nil {
		return nil, fmt.Errorf("error expanding GkeWorkload into gkeWorkload: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["gkeWorkload"] = v
	}
	if v, err := expandServiceGkeService(c, f.GkeService); err != nil {
		return nil, fmt.Errorf("error expanding GkeService into gkeService: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["gkeService"] = v
	}
	if v, err := expandServiceTelemetry(c, f.Telemetry); err != nil {
		return nil, fmt.Errorf("error expanding Telemetry into telemetry: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["telemetry"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		req["userLabels"] = v
	}
	if v := f.Deleted; !dcl.IsEmptyValueIndirect(v) {
		req["deleted"] = v
	}
	return req, nil
}

// marshalUpdateServiceUpdateServiceRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceUpdateServiceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateServiceUpdateServiceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceUpdateServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	_, err := c.GetService(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateService")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateServiceUpdateServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateServiceUpdateServiceRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listServiceRaw(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceMaxPage {
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

type listServiceOperation struct {
	Services []map[string]interface{} `json:"services"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listService(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]*Service, string, error) {
	b, err := c.listServiceRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Service
	for _, v := range m.Services {
		res, err := unmarshalMapService(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllService(ctx context.Context, f func(*Service) bool, resources []*Service) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteService(ctx, res)
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

type deleteServiceOperation struct{}

func (op *deleteServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	r, err := c.GetService(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Service not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetService checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Service: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createServiceOperation struct {
	response map[string]interface{}
}

func (op *createServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetService(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {

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

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Service
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Service); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Service, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetService(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Service resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Service resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Service resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Service: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Service: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Service: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Service: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceInitialState(rawInitial, rawDesired *Service) (*Service, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Custom) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.Custom = EmptyServiceCustom
		}
	}

	if !dcl.IsZeroValue(rawInitial.AppEngine) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.AppEngine = EmptyServiceAppEngine
		}
	}

	if !dcl.IsZeroValue(rawInitial.CloudEndpoints) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.CloudEndpoints = EmptyServiceCloudEndpoints
		}
	}

	if !dcl.IsZeroValue(rawInitial.ClusterIstio) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.ClusterIstio = EmptyServiceClusterIstio
		}
	}

	if !dcl.IsZeroValue(rawInitial.MeshIstio) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.MeshIstio = EmptyServiceMeshIstio
		}
	}

	if !dcl.IsZeroValue(rawInitial.IstioCanonicalService) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.IstioCanonicalService = EmptyServiceIstioCanonicalService
		}
	}

	if !dcl.IsZeroValue(rawInitial.CloudRun) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.GkeNamespace, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.CloudRun = EmptyServiceCloudRun
		}
	}

	if !dcl.IsZeroValue(rawInitial.GkeNamespace) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeWorkload, rawInitial.GkeService) {
			rawInitial.GkeNamespace = EmptyServiceGkeNamespace
		}
	}

	if !dcl.IsZeroValue(rawInitial.GkeWorkload) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeService) {
			rawInitial.GkeWorkload = EmptyServiceGkeWorkload
		}
	}

	if !dcl.IsZeroValue(rawInitial.GkeService) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Custom, rawInitial.AppEngine, rawInitial.CloudEndpoints, rawInitial.ClusterIstio, rawInitial.MeshIstio, rawInitial.IstioCanonicalService, rawInitial.CloudRun, rawInitial.GkeNamespace, rawInitial.GkeWorkload) {
			rawInitial.GkeService = EmptyServiceGkeService
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceDesiredState(rawDesired, rawInitial *Service, opts ...dcl.ApplyOption) (*Service, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Custom = canonicalizeServiceCustom(rawDesired.Custom, nil, opts...)
		rawDesired.AppEngine = canonicalizeServiceAppEngine(rawDesired.AppEngine, nil, opts...)
		rawDesired.CloudEndpoints = canonicalizeServiceCloudEndpoints(rawDesired.CloudEndpoints, nil, opts...)
		rawDesired.ClusterIstio = canonicalizeServiceClusterIstio(rawDesired.ClusterIstio, nil, opts...)
		rawDesired.MeshIstio = canonicalizeServiceMeshIstio(rawDesired.MeshIstio, nil, opts...)
		rawDesired.IstioCanonicalService = canonicalizeServiceIstioCanonicalService(rawDesired.IstioCanonicalService, nil, opts...)
		rawDesired.CloudRun = canonicalizeServiceCloudRun(rawDesired.CloudRun, nil, opts...)
		rawDesired.GkeNamespace = canonicalizeServiceGkeNamespace(rawDesired.GkeNamespace, nil, opts...)
		rawDesired.GkeWorkload = canonicalizeServiceGkeWorkload(rawDesired.GkeWorkload, nil, opts...)
		rawDesired.GkeService = canonicalizeServiceGkeService(rawDesired.GkeService, nil, opts...)
		rawDesired.Telemetry = canonicalizeServiceTelemetry(rawDesired.Telemetry, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.Custom != nil || rawInitial.Custom != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.Custom = nil
			rawInitial.Custom = nil
		}
	}

	if rawDesired.AppEngine != nil || rawInitial.AppEngine != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.AppEngine = nil
			rawInitial.AppEngine = nil
		}
	}

	if rawDesired.CloudEndpoints != nil || rawInitial.CloudEndpoints != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.CloudEndpoints = nil
			rawInitial.CloudEndpoints = nil
		}
	}

	if rawDesired.ClusterIstio != nil || rawInitial.ClusterIstio != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.ClusterIstio = nil
			rawInitial.ClusterIstio = nil
		}
	}

	if rawDesired.MeshIstio != nil || rawInitial.MeshIstio != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.MeshIstio = nil
			rawInitial.MeshIstio = nil
		}
	}

	if rawDesired.IstioCanonicalService != nil || rawInitial.IstioCanonicalService != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.IstioCanonicalService = nil
			rawInitial.IstioCanonicalService = nil
		}
	}

	if rawDesired.CloudRun != nil || rawInitial.CloudRun != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.GkeNamespace, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.CloudRun = nil
			rawInitial.CloudRun = nil
		}
	}

	if rawDesired.GkeNamespace != nil || rawInitial.GkeNamespace != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeWorkload, rawDesired.GkeService) {
			rawDesired.GkeNamespace = nil
			rawInitial.GkeNamespace = nil
		}
	}

	if rawDesired.GkeWorkload != nil || rawInitial.GkeWorkload != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeService) {
			rawDesired.GkeWorkload = nil
			rawInitial.GkeWorkload = nil
		}
	}

	if rawDesired.GkeService != nil || rawInitial.GkeService != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Custom, rawDesired.AppEngine, rawDesired.CloudEndpoints, rawDesired.ClusterIstio, rawDesired.MeshIstio, rawDesired.IstioCanonicalService, rawDesired.CloudRun, rawDesired.GkeNamespace, rawDesired.GkeWorkload) {
			rawDesired.GkeService = nil
			rawInitial.GkeService = nil
		}
	}

	canonicalDesired := &Service{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	canonicalDesired.Custom = canonicalizeServiceCustom(rawDesired.Custom, rawInitial.Custom, opts...)
	canonicalDesired.AppEngine = canonicalizeServiceAppEngine(rawDesired.AppEngine, rawInitial.AppEngine, opts...)
	canonicalDesired.CloudEndpoints = canonicalizeServiceCloudEndpoints(rawDesired.CloudEndpoints, rawInitial.CloudEndpoints, opts...)
	canonicalDesired.ClusterIstio = canonicalizeServiceClusterIstio(rawDesired.ClusterIstio, rawInitial.ClusterIstio, opts...)
	canonicalDesired.MeshIstio = canonicalizeServiceMeshIstio(rawDesired.MeshIstio, rawInitial.MeshIstio, opts...)
	canonicalDesired.IstioCanonicalService = canonicalizeServiceIstioCanonicalService(rawDesired.IstioCanonicalService, rawInitial.IstioCanonicalService, opts...)
	canonicalDesired.CloudRun = canonicalizeServiceCloudRun(rawDesired.CloudRun, rawInitial.CloudRun, opts...)
	canonicalDesired.GkeNamespace = canonicalizeServiceGkeNamespace(rawDesired.GkeNamespace, rawInitial.GkeNamespace, opts...)
	canonicalDesired.GkeWorkload = canonicalizeServiceGkeWorkload(rawDesired.GkeWorkload, rawInitial.GkeWorkload, opts...)
	canonicalDesired.GkeService = canonicalizeServiceGkeService(rawDesired.GkeService, rawInitial.GkeService, opts...)
	canonicalDesired.Telemetry = canonicalizeServiceTelemetry(rawDesired.Telemetry, rawInitial.Telemetry, opts...)
	if dcl.IsZeroValue(rawDesired.UserLabels) {
		canonicalDesired.UserLabels = rawInitial.UserLabels
	} else {
		canonicalDesired.UserLabels = rawDesired.UserLabels
	}
	if dcl.BoolCanonicalize(rawDesired.Deleted, rawInitial.Deleted) {
		canonicalDesired.Deleted = rawInitial.Deleted
	} else {
		canonicalDesired.Deleted = rawDesired.Deleted
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.DisplayName) && dcl.IsNotReturnedByServer(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Custom) && dcl.IsNotReturnedByServer(rawDesired.Custom) {
		rawNew.Custom = rawDesired.Custom
	} else {
		rawNew.Custom = canonicalizeNewServiceCustom(c, rawDesired.Custom, rawNew.Custom)
	}

	if dcl.IsNotReturnedByServer(rawNew.AppEngine) && dcl.IsNotReturnedByServer(rawDesired.AppEngine) {
		rawNew.AppEngine = rawDesired.AppEngine
	} else {
		rawNew.AppEngine = canonicalizeNewServiceAppEngine(c, rawDesired.AppEngine, rawNew.AppEngine)
	}

	if dcl.IsNotReturnedByServer(rawNew.CloudEndpoints) && dcl.IsNotReturnedByServer(rawDesired.CloudEndpoints) {
		rawNew.CloudEndpoints = rawDesired.CloudEndpoints
	} else {
		rawNew.CloudEndpoints = canonicalizeNewServiceCloudEndpoints(c, rawDesired.CloudEndpoints, rawNew.CloudEndpoints)
	}

	if dcl.IsNotReturnedByServer(rawNew.ClusterIstio) && dcl.IsNotReturnedByServer(rawDesired.ClusterIstio) {
		rawNew.ClusterIstio = rawDesired.ClusterIstio
	} else {
		rawNew.ClusterIstio = canonicalizeNewServiceClusterIstio(c, rawDesired.ClusterIstio, rawNew.ClusterIstio)
	}

	if dcl.IsNotReturnedByServer(rawNew.MeshIstio) && dcl.IsNotReturnedByServer(rawDesired.MeshIstio) {
		rawNew.MeshIstio = rawDesired.MeshIstio
	} else {
		rawNew.MeshIstio = canonicalizeNewServiceMeshIstio(c, rawDesired.MeshIstio, rawNew.MeshIstio)
	}

	if dcl.IsNotReturnedByServer(rawNew.IstioCanonicalService) && dcl.IsNotReturnedByServer(rawDesired.IstioCanonicalService) {
		rawNew.IstioCanonicalService = rawDesired.IstioCanonicalService
	} else {
		rawNew.IstioCanonicalService = canonicalizeNewServiceIstioCanonicalService(c, rawDesired.IstioCanonicalService, rawNew.IstioCanonicalService)
	}

	if dcl.IsNotReturnedByServer(rawNew.CloudRun) && dcl.IsNotReturnedByServer(rawDesired.CloudRun) {
		rawNew.CloudRun = rawDesired.CloudRun
	} else {
		rawNew.CloudRun = canonicalizeNewServiceCloudRun(c, rawDesired.CloudRun, rawNew.CloudRun)
	}

	if dcl.IsNotReturnedByServer(rawNew.GkeNamespace) && dcl.IsNotReturnedByServer(rawDesired.GkeNamespace) {
		rawNew.GkeNamespace = rawDesired.GkeNamespace
	} else {
		rawNew.GkeNamespace = canonicalizeNewServiceGkeNamespace(c, rawDesired.GkeNamespace, rawNew.GkeNamespace)
	}

	if dcl.IsNotReturnedByServer(rawNew.GkeWorkload) && dcl.IsNotReturnedByServer(rawDesired.GkeWorkload) {
		rawNew.GkeWorkload = rawDesired.GkeWorkload
	} else {
		rawNew.GkeWorkload = canonicalizeNewServiceGkeWorkload(c, rawDesired.GkeWorkload, rawNew.GkeWorkload)
	}

	if dcl.IsNotReturnedByServer(rawNew.GkeService) && dcl.IsNotReturnedByServer(rawDesired.GkeService) {
		rawNew.GkeService = rawDesired.GkeService
	} else {
		rawNew.GkeService = canonicalizeNewServiceGkeService(c, rawDesired.GkeService, rawNew.GkeService)
	}

	if dcl.IsNotReturnedByServer(rawNew.Telemetry) && dcl.IsNotReturnedByServer(rawDesired.Telemetry) {
		rawNew.Telemetry = rawDesired.Telemetry
	} else {
		rawNew.Telemetry = canonicalizeNewServiceTelemetry(c, rawDesired.Telemetry, rawNew.Telemetry)
	}

	if dcl.IsNotReturnedByServer(rawNew.UserLabels) && dcl.IsNotReturnedByServer(rawDesired.UserLabels) {
		rawNew.UserLabels = rawDesired.UserLabels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Deleted) && dcl.IsNotReturnedByServer(rawDesired.Deleted) {
		rawNew.Deleted = rawDesired.Deleted
	} else {
		if dcl.BoolCanonicalize(rawDesired.Deleted, rawNew.Deleted) {
			rawNew.Deleted = rawDesired.Deleted
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeServiceCustom(des, initial *ServiceCustom, opts ...dcl.ApplyOption) *ServiceCustom {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &ServiceCustom{}

	return cDes
}

func canonicalizeServiceCustomSlice(des, initial []ServiceCustom, opts ...dcl.ApplyOption) []ServiceCustom {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceCustom, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceCustom(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceCustom, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceCustom(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceCustom(c *Client, des, nw *ServiceCustom) *ServiceCustom {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceCustom while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewServiceCustomSet(c *Client, des, nw []ServiceCustom) []ServiceCustom {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceCustom
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceCustomNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceCustomSlice(c *Client, des, nw []ServiceCustom) []ServiceCustom {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceCustom
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceCustom(c, &d, &n))
	}

	return items
}

func canonicalizeServiceAppEngine(des, initial *ServiceAppEngine, opts ...dcl.ApplyOption) *ServiceAppEngine {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceAppEngine{}

	if dcl.StringCanonicalize(des.ModuleId, initial.ModuleId) || dcl.IsZeroValue(des.ModuleId) {
		cDes.ModuleId = initial.ModuleId
	} else {
		cDes.ModuleId = des.ModuleId
	}

	return cDes
}

func canonicalizeServiceAppEngineSlice(des, initial []ServiceAppEngine, opts ...dcl.ApplyOption) []ServiceAppEngine {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceAppEngine, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceAppEngine(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceAppEngine, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceAppEngine(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceAppEngine(c *Client, des, nw *ServiceAppEngine) *ServiceAppEngine {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceAppEngine while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ModuleId, nw.ModuleId) {
		nw.ModuleId = des.ModuleId
	}

	return nw
}

func canonicalizeNewServiceAppEngineSet(c *Client, des, nw []ServiceAppEngine) []ServiceAppEngine {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceAppEngine
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceAppEngineNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceAppEngineSlice(c *Client, des, nw []ServiceAppEngine) []ServiceAppEngine {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceAppEngine
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceAppEngine(c, &d, &n))
	}

	return items
}

func canonicalizeServiceCloudEndpoints(des, initial *ServiceCloudEndpoints, opts ...dcl.ApplyOption) *ServiceCloudEndpoints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceCloudEndpoints{}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}

	return cDes
}

func canonicalizeServiceCloudEndpointsSlice(des, initial []ServiceCloudEndpoints, opts ...dcl.ApplyOption) []ServiceCloudEndpoints {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceCloudEndpoints, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceCloudEndpoints(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceCloudEndpoints, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceCloudEndpoints(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceCloudEndpoints(c *Client, des, nw *ServiceCloudEndpoints) *ServiceCloudEndpoints {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceCloudEndpoints while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}

	return nw
}

func canonicalizeNewServiceCloudEndpointsSet(c *Client, des, nw []ServiceCloudEndpoints) []ServiceCloudEndpoints {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceCloudEndpoints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceCloudEndpointsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceCloudEndpointsSlice(c *Client, des, nw []ServiceCloudEndpoints) []ServiceCloudEndpoints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceCloudEndpoints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceCloudEndpoints(c, &d, &n))
	}

	return items
}

func canonicalizeServiceClusterIstio(des, initial *ServiceClusterIstio, opts ...dcl.ApplyOption) *ServiceClusterIstio {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceClusterIstio{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		cDes.ClusterName = initial.ClusterName
	} else {
		cDes.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.ServiceNamespace, initial.ServiceNamespace) || dcl.IsZeroValue(des.ServiceNamespace) {
		cDes.ServiceNamespace = initial.ServiceNamespace
	} else {
		cDes.ServiceNamespace = des.ServiceNamespace
	}
	if dcl.StringCanonicalize(des.ServiceName, initial.ServiceName) || dcl.IsZeroValue(des.ServiceName) {
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}

	return cDes
}

func canonicalizeServiceClusterIstioSlice(des, initial []ServiceClusterIstio, opts ...dcl.ApplyOption) []ServiceClusterIstio {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceClusterIstio, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceClusterIstio(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceClusterIstio, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceClusterIstio(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceClusterIstio(c *Client, des, nw *ServiceClusterIstio) *ServiceClusterIstio {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceClusterIstio while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.ServiceNamespace, nw.ServiceNamespace) {
		nw.ServiceNamespace = des.ServiceNamespace
	}
	if dcl.StringCanonicalize(des.ServiceName, nw.ServiceName) {
		nw.ServiceName = des.ServiceName
	}

	return nw
}

func canonicalizeNewServiceClusterIstioSet(c *Client, des, nw []ServiceClusterIstio) []ServiceClusterIstio {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceClusterIstio
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceClusterIstioNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceClusterIstioSlice(c *Client, des, nw []ServiceClusterIstio) []ServiceClusterIstio {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceClusterIstio
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceClusterIstio(c, &d, &n))
	}

	return items
}

func canonicalizeServiceMeshIstio(des, initial *ServiceMeshIstio, opts ...dcl.ApplyOption) *ServiceMeshIstio {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceMeshIstio{}

	if dcl.StringCanonicalize(des.MeshUid, initial.MeshUid) || dcl.IsZeroValue(des.MeshUid) {
		cDes.MeshUid = initial.MeshUid
	} else {
		cDes.MeshUid = des.MeshUid
	}
	if dcl.StringCanonicalize(des.ServiceNamespace, initial.ServiceNamespace) || dcl.IsZeroValue(des.ServiceNamespace) {
		cDes.ServiceNamespace = initial.ServiceNamespace
	} else {
		cDes.ServiceNamespace = des.ServiceNamespace
	}
	if dcl.StringCanonicalize(des.ServiceName, initial.ServiceName) || dcl.IsZeroValue(des.ServiceName) {
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}

	return cDes
}

func canonicalizeServiceMeshIstioSlice(des, initial []ServiceMeshIstio, opts ...dcl.ApplyOption) []ServiceMeshIstio {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceMeshIstio, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceMeshIstio(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceMeshIstio, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceMeshIstio(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceMeshIstio(c *Client, des, nw *ServiceMeshIstio) *ServiceMeshIstio {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceMeshIstio while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MeshUid, nw.MeshUid) {
		nw.MeshUid = des.MeshUid
	}
	if dcl.StringCanonicalize(des.ServiceNamespace, nw.ServiceNamespace) {
		nw.ServiceNamespace = des.ServiceNamespace
	}
	if dcl.StringCanonicalize(des.ServiceName, nw.ServiceName) {
		nw.ServiceName = des.ServiceName
	}

	return nw
}

func canonicalizeNewServiceMeshIstioSet(c *Client, des, nw []ServiceMeshIstio) []ServiceMeshIstio {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceMeshIstio
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceMeshIstioNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceMeshIstioSlice(c *Client, des, nw []ServiceMeshIstio) []ServiceMeshIstio {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceMeshIstio
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceMeshIstio(c, &d, &n))
	}

	return items
}

func canonicalizeServiceIstioCanonicalService(des, initial *ServiceIstioCanonicalService, opts ...dcl.ApplyOption) *ServiceIstioCanonicalService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceIstioCanonicalService{}

	if dcl.StringCanonicalize(des.MeshUid, initial.MeshUid) || dcl.IsZeroValue(des.MeshUid) {
		cDes.MeshUid = initial.MeshUid
	} else {
		cDes.MeshUid = des.MeshUid
	}
	if dcl.StringCanonicalize(des.CanonicalServiceNamespace, initial.CanonicalServiceNamespace) || dcl.IsZeroValue(des.CanonicalServiceNamespace) {
		cDes.CanonicalServiceNamespace = initial.CanonicalServiceNamespace
	} else {
		cDes.CanonicalServiceNamespace = des.CanonicalServiceNamespace
	}
	if dcl.StringCanonicalize(des.CanonicalService, initial.CanonicalService) || dcl.IsZeroValue(des.CanonicalService) {
		cDes.CanonicalService = initial.CanonicalService
	} else {
		cDes.CanonicalService = des.CanonicalService
	}

	return cDes
}

func canonicalizeServiceIstioCanonicalServiceSlice(des, initial []ServiceIstioCanonicalService, opts ...dcl.ApplyOption) []ServiceIstioCanonicalService {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceIstioCanonicalService, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceIstioCanonicalService(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceIstioCanonicalService, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceIstioCanonicalService(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceIstioCanonicalService(c *Client, des, nw *ServiceIstioCanonicalService) *ServiceIstioCanonicalService {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceIstioCanonicalService while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MeshUid, nw.MeshUid) {
		nw.MeshUid = des.MeshUid
	}
	if dcl.StringCanonicalize(des.CanonicalServiceNamespace, nw.CanonicalServiceNamespace) {
		nw.CanonicalServiceNamespace = des.CanonicalServiceNamespace
	}
	if dcl.StringCanonicalize(des.CanonicalService, nw.CanonicalService) {
		nw.CanonicalService = des.CanonicalService
	}

	return nw
}

func canonicalizeNewServiceIstioCanonicalServiceSet(c *Client, des, nw []ServiceIstioCanonicalService) []ServiceIstioCanonicalService {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceIstioCanonicalService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceIstioCanonicalServiceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceIstioCanonicalServiceSlice(c *Client, des, nw []ServiceIstioCanonicalService) []ServiceIstioCanonicalService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceIstioCanonicalService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceIstioCanonicalService(c, &d, &n))
	}

	return items
}

func canonicalizeServiceCloudRun(des, initial *ServiceCloudRun, opts ...dcl.ApplyOption) *ServiceCloudRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceCloudRun{}

	if dcl.StringCanonicalize(des.ServiceName, initial.ServiceName) || dcl.IsZeroValue(des.ServiceName) {
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}
	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}

	return cDes
}

func canonicalizeServiceCloudRunSlice(des, initial []ServiceCloudRun, opts ...dcl.ApplyOption) []ServiceCloudRun {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceCloudRun, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceCloudRun(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceCloudRun, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceCloudRun(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceCloudRun(c *Client, des, nw *ServiceCloudRun) *ServiceCloudRun {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceCloudRun while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ServiceName, nw.ServiceName) {
		nw.ServiceName = des.ServiceName
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}

	return nw
}

func canonicalizeNewServiceCloudRunSet(c *Client, des, nw []ServiceCloudRun) []ServiceCloudRun {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceCloudRun
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceCloudRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceCloudRunSlice(c *Client, des, nw []ServiceCloudRun) []ServiceCloudRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceCloudRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceCloudRun(c, &d, &n))
	}

	return items
}

func canonicalizeServiceGkeNamespace(des, initial *ServiceGkeNamespace, opts ...dcl.ApplyOption) *ServiceGkeNamespace {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceGkeNamespace{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		cDes.ClusterName = initial.ClusterName
	} else {
		cDes.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, initial.NamespaceName) || dcl.IsZeroValue(des.NamespaceName) {
		cDes.NamespaceName = initial.NamespaceName
	} else {
		cDes.NamespaceName = des.NamespaceName
	}

	return cDes
}

func canonicalizeServiceGkeNamespaceSlice(des, initial []ServiceGkeNamespace, opts ...dcl.ApplyOption) []ServiceGkeNamespace {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceGkeNamespace, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceGkeNamespace(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceGkeNamespace, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceGkeNamespace(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceGkeNamespace(c *Client, des, nw *ServiceGkeNamespace) *ServiceGkeNamespace {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceGkeNamespace while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, nw.NamespaceName) {
		nw.NamespaceName = des.NamespaceName
	}

	return nw
}

func canonicalizeNewServiceGkeNamespaceSet(c *Client, des, nw []ServiceGkeNamespace) []ServiceGkeNamespace {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceGkeNamespace
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceGkeNamespaceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceGkeNamespaceSlice(c *Client, des, nw []ServiceGkeNamespace) []ServiceGkeNamespace {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceGkeNamespace
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceGkeNamespace(c, &d, &n))
	}

	return items
}

func canonicalizeServiceGkeWorkload(des, initial *ServiceGkeWorkload, opts ...dcl.ApplyOption) *ServiceGkeWorkload {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceGkeWorkload{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		cDes.ClusterName = initial.ClusterName
	} else {
		cDes.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, initial.NamespaceName) || dcl.IsZeroValue(des.NamespaceName) {
		cDes.NamespaceName = initial.NamespaceName
	} else {
		cDes.NamespaceName = des.NamespaceName
	}
	if dcl.StringCanonicalize(des.TopLevelControllerType, initial.TopLevelControllerType) || dcl.IsZeroValue(des.TopLevelControllerType) {
		cDes.TopLevelControllerType = initial.TopLevelControllerType
	} else {
		cDes.TopLevelControllerType = des.TopLevelControllerType
	}
	if dcl.StringCanonicalize(des.TopLevelControllerName, initial.TopLevelControllerName) || dcl.IsZeroValue(des.TopLevelControllerName) {
		cDes.TopLevelControllerName = initial.TopLevelControllerName
	} else {
		cDes.TopLevelControllerName = des.TopLevelControllerName
	}

	return cDes
}

func canonicalizeServiceGkeWorkloadSlice(des, initial []ServiceGkeWorkload, opts ...dcl.ApplyOption) []ServiceGkeWorkload {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceGkeWorkload, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceGkeWorkload(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceGkeWorkload, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceGkeWorkload(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceGkeWorkload(c *Client, des, nw *ServiceGkeWorkload) *ServiceGkeWorkload {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceGkeWorkload while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, nw.NamespaceName) {
		nw.NamespaceName = des.NamespaceName
	}
	if dcl.StringCanonicalize(des.TopLevelControllerType, nw.TopLevelControllerType) {
		nw.TopLevelControllerType = des.TopLevelControllerType
	}
	if dcl.StringCanonicalize(des.TopLevelControllerName, nw.TopLevelControllerName) {
		nw.TopLevelControllerName = des.TopLevelControllerName
	}

	return nw
}

func canonicalizeNewServiceGkeWorkloadSet(c *Client, des, nw []ServiceGkeWorkload) []ServiceGkeWorkload {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceGkeWorkload
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceGkeWorkloadNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceGkeWorkloadSlice(c *Client, des, nw []ServiceGkeWorkload) []ServiceGkeWorkload {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceGkeWorkload
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceGkeWorkload(c, &d, &n))
	}

	return items
}

func canonicalizeServiceGkeService(des, initial *ServiceGkeService, opts ...dcl.ApplyOption) *ServiceGkeService {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceGkeService{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		cDes.ClusterName = initial.ClusterName
	} else {
		cDes.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, initial.NamespaceName) || dcl.IsZeroValue(des.NamespaceName) {
		cDes.NamespaceName = initial.NamespaceName
	} else {
		cDes.NamespaceName = des.NamespaceName
	}
	if dcl.StringCanonicalize(des.ServiceName, initial.ServiceName) || dcl.IsZeroValue(des.ServiceName) {
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}

	return cDes
}

func canonicalizeServiceGkeServiceSlice(des, initial []ServiceGkeService, opts ...dcl.ApplyOption) []ServiceGkeService {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceGkeService, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceGkeService(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceGkeService, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceGkeService(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceGkeService(c *Client, des, nw *ServiceGkeService) *ServiceGkeService {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceGkeService while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.NamespaceName, nw.NamespaceName) {
		nw.NamespaceName = des.NamespaceName
	}
	if dcl.StringCanonicalize(des.ServiceName, nw.ServiceName) {
		nw.ServiceName = des.ServiceName
	}

	return nw
}

func canonicalizeNewServiceGkeServiceSet(c *Client, des, nw []ServiceGkeService) []ServiceGkeService {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceGkeService
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceGkeServiceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceGkeServiceSlice(c *Client, des, nw []ServiceGkeService) []ServiceGkeService {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceGkeService
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceGkeService(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTelemetry(des, initial *ServiceTelemetry, opts ...dcl.ApplyOption) *ServiceTelemetry {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTelemetry{}

	if dcl.StringCanonicalize(des.ResourceName, initial.ResourceName) || dcl.IsZeroValue(des.ResourceName) {
		cDes.ResourceName = initial.ResourceName
	} else {
		cDes.ResourceName = des.ResourceName
	}

	return cDes
}

func canonicalizeServiceTelemetrySlice(des, initial []ServiceTelemetry, opts ...dcl.ApplyOption) []ServiceTelemetry {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTelemetry, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTelemetry(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTelemetry, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTelemetry(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTelemetry(c *Client, des, nw *ServiceTelemetry) *ServiceTelemetry {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTelemetry while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceName, nw.ResourceName) {
		nw.ResourceName = des.ResourceName
	}

	return nw
}

func canonicalizeNewServiceTelemetrySet(c *Client, des, nw []ServiceTelemetry) []ServiceTelemetry {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTelemetry
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTelemetryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTelemetrySlice(c *Client, des, nw []ServiceTelemetry) []ServiceTelemetry {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTelemetry
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTelemetry(c, &d, &n))
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
func diffService(c *Client, desired, actual *Service, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Custom, actual.Custom, dcl.Info{ObjectFunction: compareServiceCustomNewStyle, EmptyObject: EmptyServiceCustom, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Custom")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppEngine, actual.AppEngine, dcl.Info{ObjectFunction: compareServiceAppEngineNewStyle, EmptyObject: EmptyServiceAppEngine, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("AppEngine")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudEndpoints, actual.CloudEndpoints, dcl.Info{ObjectFunction: compareServiceCloudEndpointsNewStyle, EmptyObject: EmptyServiceCloudEndpoints, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CloudEndpoints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterIstio, actual.ClusterIstio, dcl.Info{ObjectFunction: compareServiceClusterIstioNewStyle, EmptyObject: EmptyServiceClusterIstio, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClusterIstio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MeshIstio, actual.MeshIstio, dcl.Info{ObjectFunction: compareServiceMeshIstioNewStyle, EmptyObject: EmptyServiceMeshIstio, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MeshIstio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IstioCanonicalService, actual.IstioCanonicalService, dcl.Info{ObjectFunction: compareServiceIstioCanonicalServiceNewStyle, EmptyObject: EmptyServiceIstioCanonicalService, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("IstioCanonicalService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudRun, actual.CloudRun, dcl.Info{ObjectFunction: compareServiceCloudRunNewStyle, EmptyObject: EmptyServiceCloudRun, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CloudRun")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GkeNamespace, actual.GkeNamespace, dcl.Info{ObjectFunction: compareServiceGkeNamespaceNewStyle, EmptyObject: EmptyServiceGkeNamespace, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("GkeNamespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GkeWorkload, actual.GkeWorkload, dcl.Info{ObjectFunction: compareServiceGkeWorkloadNewStyle, EmptyObject: EmptyServiceGkeWorkload, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("GkeWorkload")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GkeService, actual.GkeService, dcl.Info{ObjectFunction: compareServiceGkeServiceNewStyle, EmptyObject: EmptyServiceGkeService, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("GkeService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Telemetry, actual.Telemetry, dcl.Info{ObjectFunction: compareServiceTelemetryNewStyle, EmptyObject: EmptyServiceTelemetry, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Telemetry")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserLabels, actual.UserLabels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("UserLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deleted, actual.Deleted, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Deleted")); len(ds) != 0 || err != nil {
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
func compareServiceCustomNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareServiceAppEngineNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceAppEngine)
	if !ok {
		desiredNotPointer, ok := d.(ServiceAppEngine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAppEngine or *ServiceAppEngine", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceAppEngine)
	if !ok {
		actualNotPointer, ok := a.(ServiceAppEngine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAppEngine", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ModuleId, actual.ModuleId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ModuleId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceCloudEndpointsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceCloudEndpoints)
	if !ok {
		desiredNotPointer, ok := d.(ServiceCloudEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceCloudEndpoints or *ServiceCloudEndpoints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceCloudEndpoints)
	if !ok {
		actualNotPointer, ok := a.(ServiceCloudEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceCloudEndpoints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceClusterIstioNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceClusterIstio)
	if !ok {
		desiredNotPointer, ok := d.(ServiceClusterIstio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceClusterIstio or *ServiceClusterIstio", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceClusterIstio)
	if !ok {
		actualNotPointer, ok := a.(ServiceClusterIstio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceClusterIstio", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceNamespace, actual.ServiceNamespace, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceNamespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceMeshIstioNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceMeshIstio)
	if !ok {
		desiredNotPointer, ok := d.(ServiceMeshIstio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMeshIstio or *ServiceMeshIstio", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceMeshIstio)
	if !ok {
		actualNotPointer, ok := a.(ServiceMeshIstio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMeshIstio", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MeshUid, actual.MeshUid, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MeshUid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceNamespace, actual.ServiceNamespace, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceNamespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceIstioCanonicalServiceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceIstioCanonicalService)
	if !ok {
		desiredNotPointer, ok := d.(ServiceIstioCanonicalService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceIstioCanonicalService or *ServiceIstioCanonicalService", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceIstioCanonicalService)
	if !ok {
		actualNotPointer, ok := a.(ServiceIstioCanonicalService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceIstioCanonicalService", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MeshUid, actual.MeshUid, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MeshUid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanonicalServiceNamespace, actual.CanonicalServiceNamespace, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CanonicalServiceNamespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanonicalService, actual.CanonicalService, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CanonicalService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceCloudRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceCloudRun)
	if !ok {
		desiredNotPointer, ok := d.(ServiceCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceCloudRun or *ServiceCloudRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceCloudRun)
	if !ok {
		actualNotPointer, ok := a.(ServiceCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceCloudRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceGkeNamespaceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceGkeNamespace)
	if !ok {
		desiredNotPointer, ok := d.(ServiceGkeNamespace)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeNamespace or *ServiceGkeNamespace", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceGkeNamespace)
	if !ok {
		actualNotPointer, ok := a.(ServiceGkeNamespace)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeNamespace", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NamespaceName, actual.NamespaceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("NamespaceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceGkeWorkloadNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceGkeWorkload)
	if !ok {
		desiredNotPointer, ok := d.(ServiceGkeWorkload)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeWorkload or *ServiceGkeWorkload", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceGkeWorkload)
	if !ok {
		actualNotPointer, ok := a.(ServiceGkeWorkload)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeWorkload", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NamespaceName, actual.NamespaceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("NamespaceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TopLevelControllerType, actual.TopLevelControllerType, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("TopLevelControllerType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TopLevelControllerName, actual.TopLevelControllerName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("TopLevelControllerName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceGkeServiceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceGkeService)
	if !ok {
		desiredNotPointer, ok := d.(ServiceGkeService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeService or *ServiceGkeService", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceGkeService)
	if !ok {
		actualNotPointer, ok := a.(ServiceGkeService)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGkeService", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NamespaceName, actual.NamespaceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("NamespaceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTelemetryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTelemetry)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTelemetry)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTelemetry or *ServiceTelemetry", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTelemetry)
	if !ok {
		actualNotPointer, ok := a.(ServiceTelemetry)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTelemetry", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceName, actual.ResourceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ResourceName")); len(ds) != 0 || err != nil {
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
func (r *Service) urlNormalized() *Service {
	normalized := dcl.Copy(*r).(Service)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Service) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateService" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/services/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Service resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Service) marshal(c *Client) ([]byte, error) {
	m, err := expandService(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Service: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalService decodes JSON responses into the Service resource schema.
func unmarshalService(b []byte, c *Client) (*Service, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapService(m, c)
}

func unmarshalMapService(m map[string]interface{}, c *Client) (*Service, error) {

	flattened := flattenService(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/services/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v, err := expandServiceCustom(c, f.Custom); err != nil {
		return nil, fmt.Errorf("error expanding Custom into custom: %w", err)
	} else {
		m["custom"] = v
	}
	if v, err := expandServiceAppEngine(c, f.AppEngine); err != nil {
		return nil, fmt.Errorf("error expanding AppEngine into appEngine: %w", err)
	} else if v != nil {
		m["appEngine"] = v
	}
	if v, err := expandServiceCloudEndpoints(c, f.CloudEndpoints); err != nil {
		return nil, fmt.Errorf("error expanding CloudEndpoints into cloudEndpoints: %w", err)
	} else if v != nil {
		m["cloudEndpoints"] = v
	}
	if v, err := expandServiceClusterIstio(c, f.ClusterIstio); err != nil {
		return nil, fmt.Errorf("error expanding ClusterIstio into clusterIstio: %w", err)
	} else if v != nil {
		m["clusterIstio"] = v
	}
	if v, err := expandServiceMeshIstio(c, f.MeshIstio); err != nil {
		return nil, fmt.Errorf("error expanding MeshIstio into meshIstio: %w", err)
	} else if v != nil {
		m["meshIstio"] = v
	}
	if v, err := expandServiceIstioCanonicalService(c, f.IstioCanonicalService); err != nil {
		return nil, fmt.Errorf("error expanding IstioCanonicalService into istioCanonicalService: %w", err)
	} else if v != nil {
		m["istioCanonicalService"] = v
	}
	if v, err := expandServiceCloudRun(c, f.CloudRun); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if v != nil {
		m["cloudRun"] = v
	}
	if v, err := expandServiceGkeNamespace(c, f.GkeNamespace); err != nil {
		return nil, fmt.Errorf("error expanding GkeNamespace into gkeNamespace: %w", err)
	} else if v != nil {
		m["gkeNamespace"] = v
	}
	if v, err := expandServiceGkeWorkload(c, f.GkeWorkload); err != nil {
		return nil, fmt.Errorf("error expanding GkeWorkload into gkeWorkload: %w", err)
	} else if v != nil {
		m["gkeWorkload"] = v
	}
	if v, err := expandServiceGkeService(c, f.GkeService); err != nil {
		return nil, fmt.Errorf("error expanding GkeService into gkeService: %w", err)
	} else if v != nil {
		m["gkeService"] = v
	}
	if v, err := expandServiceTelemetry(c, f.Telemetry); err != nil {
		return nil, fmt.Errorf("error expanding Telemetry into telemetry: %w", err)
	} else if v != nil {
		m["telemetry"] = v
	}
	if v := f.UserLabels; dcl.ValueShouldBeSent(v) {
		m["userLabels"] = v
	}
	if v := f.Deleted; dcl.ValueShouldBeSent(v) {
		m["deleted"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenService flattens Service from a JSON request object into the
// Service type.
func flattenService(c *Client, i interface{}) *Service {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Service{}
	res.Name = dcl.FlattenString(m["name"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Custom = flattenServiceCustom(c, m["custom"])
	res.AppEngine = flattenServiceAppEngine(c, m["appEngine"])
	res.CloudEndpoints = flattenServiceCloudEndpoints(c, m["cloudEndpoints"])
	res.ClusterIstio = flattenServiceClusterIstio(c, m["clusterIstio"])
	res.MeshIstio = flattenServiceMeshIstio(c, m["meshIstio"])
	res.IstioCanonicalService = flattenServiceIstioCanonicalService(c, m["istioCanonicalService"])
	res.CloudRun = flattenServiceCloudRun(c, m["cloudRun"])
	res.GkeNamespace = flattenServiceGkeNamespace(c, m["gkeNamespace"])
	res.GkeWorkload = flattenServiceGkeWorkload(c, m["gkeWorkload"])
	res.GkeService = flattenServiceGkeService(c, m["gkeService"])
	res.Telemetry = flattenServiceTelemetry(c, m["telemetry"])
	res.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	res.Deleted = dcl.FlattenBool(m["deleted"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandServiceCustomMap expands the contents of ServiceCustom into a JSON
// request object.
func expandServiceCustomMap(c *Client, f map[string]ServiceCustom) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceCustom(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceCustomSlice expands the contents of ServiceCustom into a JSON
// request object.
func expandServiceCustomSlice(c *Client, f []ServiceCustom) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceCustom(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceCustomMap flattens the contents of ServiceCustom from a JSON
// response object.
func flattenServiceCustomMap(c *Client, i interface{}) map[string]ServiceCustom {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceCustom{}
	}

	if len(a) == 0 {
		return map[string]ServiceCustom{}
	}

	items := make(map[string]ServiceCustom)
	for k, item := range a {
		items[k] = *flattenServiceCustom(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceCustomSlice flattens the contents of ServiceCustom from a JSON
// response object.
func flattenServiceCustomSlice(c *Client, i interface{}) []ServiceCustom {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceCustom{}
	}

	if len(a) == 0 {
		return []ServiceCustom{}
	}

	items := make([]ServiceCustom, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceCustom(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceCustom expands an instance of ServiceCustom into a JSON
// request object.
func expandServiceCustom(c *Client, f *ServiceCustom) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenServiceCustom flattens an instance of ServiceCustom from a JSON
// response object.
func flattenServiceCustom(c *Client, i interface{}) *ServiceCustom {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceCustom{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceCustom
	}

	return r
}

// expandServiceAppEngineMap expands the contents of ServiceAppEngine into a JSON
// request object.
func expandServiceAppEngineMap(c *Client, f map[string]ServiceAppEngine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceAppEngine(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceAppEngineSlice expands the contents of ServiceAppEngine into a JSON
// request object.
func expandServiceAppEngineSlice(c *Client, f []ServiceAppEngine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceAppEngine(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceAppEngineMap flattens the contents of ServiceAppEngine from a JSON
// response object.
func flattenServiceAppEngineMap(c *Client, i interface{}) map[string]ServiceAppEngine {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceAppEngine{}
	}

	if len(a) == 0 {
		return map[string]ServiceAppEngine{}
	}

	items := make(map[string]ServiceAppEngine)
	for k, item := range a {
		items[k] = *flattenServiceAppEngine(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceAppEngineSlice flattens the contents of ServiceAppEngine from a JSON
// response object.
func flattenServiceAppEngineSlice(c *Client, i interface{}) []ServiceAppEngine {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceAppEngine{}
	}

	if len(a) == 0 {
		return []ServiceAppEngine{}
	}

	items := make([]ServiceAppEngine, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceAppEngine(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceAppEngine expands an instance of ServiceAppEngine into a JSON
// request object.
func expandServiceAppEngine(c *Client, f *ServiceAppEngine) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ModuleId; !dcl.IsEmptyValueIndirect(v) {
		m["moduleId"] = v
	}

	return m, nil
}

// flattenServiceAppEngine flattens an instance of ServiceAppEngine from a JSON
// response object.
func flattenServiceAppEngine(c *Client, i interface{}) *ServiceAppEngine {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceAppEngine{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceAppEngine
	}
	r.ModuleId = dcl.FlattenString(m["moduleId"])

	return r
}

// expandServiceCloudEndpointsMap expands the contents of ServiceCloudEndpoints into a JSON
// request object.
func expandServiceCloudEndpointsMap(c *Client, f map[string]ServiceCloudEndpoints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceCloudEndpoints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceCloudEndpointsSlice expands the contents of ServiceCloudEndpoints into a JSON
// request object.
func expandServiceCloudEndpointsSlice(c *Client, f []ServiceCloudEndpoints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceCloudEndpoints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceCloudEndpointsMap flattens the contents of ServiceCloudEndpoints from a JSON
// response object.
func flattenServiceCloudEndpointsMap(c *Client, i interface{}) map[string]ServiceCloudEndpoints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceCloudEndpoints{}
	}

	if len(a) == 0 {
		return map[string]ServiceCloudEndpoints{}
	}

	items := make(map[string]ServiceCloudEndpoints)
	for k, item := range a {
		items[k] = *flattenServiceCloudEndpoints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceCloudEndpointsSlice flattens the contents of ServiceCloudEndpoints from a JSON
// response object.
func flattenServiceCloudEndpointsSlice(c *Client, i interface{}) []ServiceCloudEndpoints {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceCloudEndpoints{}
	}

	if len(a) == 0 {
		return []ServiceCloudEndpoints{}
	}

	items := make([]ServiceCloudEndpoints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceCloudEndpoints(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceCloudEndpoints expands an instance of ServiceCloudEndpoints into a JSON
// request object.
func expandServiceCloudEndpoints(c *Client, f *ServiceCloudEndpoints) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}

	return m, nil
}

// flattenServiceCloudEndpoints flattens an instance of ServiceCloudEndpoints from a JSON
// response object.
func flattenServiceCloudEndpoints(c *Client, i interface{}) *ServiceCloudEndpoints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceCloudEndpoints{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceCloudEndpoints
	}
	r.Service = dcl.FlattenString(m["service"])

	return r
}

// expandServiceClusterIstioMap expands the contents of ServiceClusterIstio into a JSON
// request object.
func expandServiceClusterIstioMap(c *Client, f map[string]ServiceClusterIstio) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceClusterIstio(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceClusterIstioSlice expands the contents of ServiceClusterIstio into a JSON
// request object.
func expandServiceClusterIstioSlice(c *Client, f []ServiceClusterIstio) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceClusterIstio(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceClusterIstioMap flattens the contents of ServiceClusterIstio from a JSON
// response object.
func flattenServiceClusterIstioMap(c *Client, i interface{}) map[string]ServiceClusterIstio {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceClusterIstio{}
	}

	if len(a) == 0 {
		return map[string]ServiceClusterIstio{}
	}

	items := make(map[string]ServiceClusterIstio)
	for k, item := range a {
		items[k] = *flattenServiceClusterIstio(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceClusterIstioSlice flattens the contents of ServiceClusterIstio from a JSON
// response object.
func flattenServiceClusterIstioSlice(c *Client, i interface{}) []ServiceClusterIstio {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceClusterIstio{}
	}

	if len(a) == 0 {
		return []ServiceClusterIstio{}
	}

	items := make([]ServiceClusterIstio, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceClusterIstio(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceClusterIstio expands an instance of ServiceClusterIstio into a JSON
// request object.
func expandServiceClusterIstio(c *Client, f *ServiceClusterIstio) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v := f.ServiceNamespace; !dcl.IsEmptyValueIndirect(v) {
		m["serviceNamespace"] = v
	}
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}

	return m, nil
}

// flattenServiceClusterIstio flattens an instance of ServiceClusterIstio from a JSON
// response object.
func flattenServiceClusterIstio(c *Client, i interface{}) *ServiceClusterIstio {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceClusterIstio{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceClusterIstio
	}
	r.Location = dcl.FlattenString(m["location"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.ServiceNamespace = dcl.FlattenString(m["serviceNamespace"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandServiceMeshIstioMap expands the contents of ServiceMeshIstio into a JSON
// request object.
func expandServiceMeshIstioMap(c *Client, f map[string]ServiceMeshIstio) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceMeshIstio(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceMeshIstioSlice expands the contents of ServiceMeshIstio into a JSON
// request object.
func expandServiceMeshIstioSlice(c *Client, f []ServiceMeshIstio) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceMeshIstio(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceMeshIstioMap flattens the contents of ServiceMeshIstio from a JSON
// response object.
func flattenServiceMeshIstioMap(c *Client, i interface{}) map[string]ServiceMeshIstio {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceMeshIstio{}
	}

	if len(a) == 0 {
		return map[string]ServiceMeshIstio{}
	}

	items := make(map[string]ServiceMeshIstio)
	for k, item := range a {
		items[k] = *flattenServiceMeshIstio(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceMeshIstioSlice flattens the contents of ServiceMeshIstio from a JSON
// response object.
func flattenServiceMeshIstioSlice(c *Client, i interface{}) []ServiceMeshIstio {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceMeshIstio{}
	}

	if len(a) == 0 {
		return []ServiceMeshIstio{}
	}

	items := make([]ServiceMeshIstio, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceMeshIstio(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceMeshIstio expands an instance of ServiceMeshIstio into a JSON
// request object.
func expandServiceMeshIstio(c *Client, f *ServiceMeshIstio) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MeshUid; !dcl.IsEmptyValueIndirect(v) {
		m["meshUid"] = v
	}
	if v := f.ServiceNamespace; !dcl.IsEmptyValueIndirect(v) {
		m["serviceNamespace"] = v
	}
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}

	return m, nil
}

// flattenServiceMeshIstio flattens an instance of ServiceMeshIstio from a JSON
// response object.
func flattenServiceMeshIstio(c *Client, i interface{}) *ServiceMeshIstio {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceMeshIstio{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceMeshIstio
	}
	r.MeshUid = dcl.FlattenString(m["meshUid"])
	r.ServiceNamespace = dcl.FlattenString(m["serviceNamespace"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandServiceIstioCanonicalServiceMap expands the contents of ServiceIstioCanonicalService into a JSON
// request object.
func expandServiceIstioCanonicalServiceMap(c *Client, f map[string]ServiceIstioCanonicalService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceIstioCanonicalService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceIstioCanonicalServiceSlice expands the contents of ServiceIstioCanonicalService into a JSON
// request object.
func expandServiceIstioCanonicalServiceSlice(c *Client, f []ServiceIstioCanonicalService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceIstioCanonicalService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceIstioCanonicalServiceMap flattens the contents of ServiceIstioCanonicalService from a JSON
// response object.
func flattenServiceIstioCanonicalServiceMap(c *Client, i interface{}) map[string]ServiceIstioCanonicalService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceIstioCanonicalService{}
	}

	if len(a) == 0 {
		return map[string]ServiceIstioCanonicalService{}
	}

	items := make(map[string]ServiceIstioCanonicalService)
	for k, item := range a {
		items[k] = *flattenServiceIstioCanonicalService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceIstioCanonicalServiceSlice flattens the contents of ServiceIstioCanonicalService from a JSON
// response object.
func flattenServiceIstioCanonicalServiceSlice(c *Client, i interface{}) []ServiceIstioCanonicalService {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceIstioCanonicalService{}
	}

	if len(a) == 0 {
		return []ServiceIstioCanonicalService{}
	}

	items := make([]ServiceIstioCanonicalService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceIstioCanonicalService(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceIstioCanonicalService expands an instance of ServiceIstioCanonicalService into a JSON
// request object.
func expandServiceIstioCanonicalService(c *Client, f *ServiceIstioCanonicalService) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MeshUid; !dcl.IsEmptyValueIndirect(v) {
		m["meshUid"] = v
	}
	if v := f.CanonicalServiceNamespace; !dcl.IsEmptyValueIndirect(v) {
		m["canonicalServiceNamespace"] = v
	}
	if v := f.CanonicalService; !dcl.IsEmptyValueIndirect(v) {
		m["canonicalService"] = v
	}

	return m, nil
}

// flattenServiceIstioCanonicalService flattens an instance of ServiceIstioCanonicalService from a JSON
// response object.
func flattenServiceIstioCanonicalService(c *Client, i interface{}) *ServiceIstioCanonicalService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceIstioCanonicalService{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceIstioCanonicalService
	}
	r.MeshUid = dcl.FlattenString(m["meshUid"])
	r.CanonicalServiceNamespace = dcl.FlattenString(m["canonicalServiceNamespace"])
	r.CanonicalService = dcl.FlattenString(m["canonicalService"])

	return r
}

// expandServiceCloudRunMap expands the contents of ServiceCloudRun into a JSON
// request object.
func expandServiceCloudRunMap(c *Client, f map[string]ServiceCloudRun) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceCloudRun(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceCloudRunSlice expands the contents of ServiceCloudRun into a JSON
// request object.
func expandServiceCloudRunSlice(c *Client, f []ServiceCloudRun) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceCloudRun(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceCloudRunMap flattens the contents of ServiceCloudRun from a JSON
// response object.
func flattenServiceCloudRunMap(c *Client, i interface{}) map[string]ServiceCloudRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceCloudRun{}
	}

	if len(a) == 0 {
		return map[string]ServiceCloudRun{}
	}

	items := make(map[string]ServiceCloudRun)
	for k, item := range a {
		items[k] = *flattenServiceCloudRun(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceCloudRunSlice flattens the contents of ServiceCloudRun from a JSON
// response object.
func flattenServiceCloudRunSlice(c *Client, i interface{}) []ServiceCloudRun {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceCloudRun{}
	}

	if len(a) == 0 {
		return []ServiceCloudRun{}
	}

	items := make([]ServiceCloudRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceCloudRun(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceCloudRun expands an instance of ServiceCloudRun into a JSON
// request object.
func expandServiceCloudRun(c *Client, f *ServiceCloudRun) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenServiceCloudRun flattens an instance of ServiceCloudRun from a JSON
// response object.
func flattenServiceCloudRun(c *Client, i interface{}) *ServiceCloudRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceCloudRun{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceCloudRun
	}
	r.ServiceName = dcl.FlattenString(m["serviceName"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandServiceGkeNamespaceMap expands the contents of ServiceGkeNamespace into a JSON
// request object.
func expandServiceGkeNamespaceMap(c *Client, f map[string]ServiceGkeNamespace) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceGkeNamespace(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceGkeNamespaceSlice expands the contents of ServiceGkeNamespace into a JSON
// request object.
func expandServiceGkeNamespaceSlice(c *Client, f []ServiceGkeNamespace) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceGkeNamespace(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceGkeNamespaceMap flattens the contents of ServiceGkeNamespace from a JSON
// response object.
func flattenServiceGkeNamespaceMap(c *Client, i interface{}) map[string]ServiceGkeNamespace {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGkeNamespace{}
	}

	if len(a) == 0 {
		return map[string]ServiceGkeNamespace{}
	}

	items := make(map[string]ServiceGkeNamespace)
	for k, item := range a {
		items[k] = *flattenServiceGkeNamespace(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceGkeNamespaceSlice flattens the contents of ServiceGkeNamespace from a JSON
// response object.
func flattenServiceGkeNamespaceSlice(c *Client, i interface{}) []ServiceGkeNamespace {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGkeNamespace{}
	}

	if len(a) == 0 {
		return []ServiceGkeNamespace{}
	}

	items := make([]ServiceGkeNamespace, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGkeNamespace(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceGkeNamespace expands an instance of ServiceGkeNamespace into a JSON
// request object.
func expandServiceGkeNamespace(c *Client, f *ServiceGkeNamespace) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v := f.NamespaceName; !dcl.IsEmptyValueIndirect(v) {
		m["namespaceName"] = v
	}

	return m, nil
}

// flattenServiceGkeNamespace flattens an instance of ServiceGkeNamespace from a JSON
// response object.
func flattenServiceGkeNamespace(c *Client, i interface{}) *ServiceGkeNamespace {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceGkeNamespace{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceGkeNamespace
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.Location = dcl.FlattenString(m["location"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.NamespaceName = dcl.FlattenString(m["namespaceName"])

	return r
}

// expandServiceGkeWorkloadMap expands the contents of ServiceGkeWorkload into a JSON
// request object.
func expandServiceGkeWorkloadMap(c *Client, f map[string]ServiceGkeWorkload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceGkeWorkload(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceGkeWorkloadSlice expands the contents of ServiceGkeWorkload into a JSON
// request object.
func expandServiceGkeWorkloadSlice(c *Client, f []ServiceGkeWorkload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceGkeWorkload(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceGkeWorkloadMap flattens the contents of ServiceGkeWorkload from a JSON
// response object.
func flattenServiceGkeWorkloadMap(c *Client, i interface{}) map[string]ServiceGkeWorkload {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGkeWorkload{}
	}

	if len(a) == 0 {
		return map[string]ServiceGkeWorkload{}
	}

	items := make(map[string]ServiceGkeWorkload)
	for k, item := range a {
		items[k] = *flattenServiceGkeWorkload(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceGkeWorkloadSlice flattens the contents of ServiceGkeWorkload from a JSON
// response object.
func flattenServiceGkeWorkloadSlice(c *Client, i interface{}) []ServiceGkeWorkload {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGkeWorkload{}
	}

	if len(a) == 0 {
		return []ServiceGkeWorkload{}
	}

	items := make([]ServiceGkeWorkload, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGkeWorkload(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceGkeWorkload expands an instance of ServiceGkeWorkload into a JSON
// request object.
func expandServiceGkeWorkload(c *Client, f *ServiceGkeWorkload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v := f.NamespaceName; !dcl.IsEmptyValueIndirect(v) {
		m["namespaceName"] = v
	}
	if v := f.TopLevelControllerType; !dcl.IsEmptyValueIndirect(v) {
		m["topLevelControllerType"] = v
	}
	if v := f.TopLevelControllerName; !dcl.IsEmptyValueIndirect(v) {
		m["topLevelControllerName"] = v
	}

	return m, nil
}

// flattenServiceGkeWorkload flattens an instance of ServiceGkeWorkload from a JSON
// response object.
func flattenServiceGkeWorkload(c *Client, i interface{}) *ServiceGkeWorkload {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceGkeWorkload{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceGkeWorkload
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.Location = dcl.FlattenString(m["location"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.NamespaceName = dcl.FlattenString(m["namespaceName"])
	r.TopLevelControllerType = dcl.FlattenString(m["topLevelControllerType"])
	r.TopLevelControllerName = dcl.FlattenString(m["topLevelControllerName"])

	return r
}

// expandServiceGkeServiceMap expands the contents of ServiceGkeService into a JSON
// request object.
func expandServiceGkeServiceMap(c *Client, f map[string]ServiceGkeService) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceGkeService(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceGkeServiceSlice expands the contents of ServiceGkeService into a JSON
// request object.
func expandServiceGkeServiceSlice(c *Client, f []ServiceGkeService) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceGkeService(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceGkeServiceMap flattens the contents of ServiceGkeService from a JSON
// response object.
func flattenServiceGkeServiceMap(c *Client, i interface{}) map[string]ServiceGkeService {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGkeService{}
	}

	if len(a) == 0 {
		return map[string]ServiceGkeService{}
	}

	items := make(map[string]ServiceGkeService)
	for k, item := range a {
		items[k] = *flattenServiceGkeService(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceGkeServiceSlice flattens the contents of ServiceGkeService from a JSON
// response object.
func flattenServiceGkeServiceSlice(c *Client, i interface{}) []ServiceGkeService {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGkeService{}
	}

	if len(a) == 0 {
		return []ServiceGkeService{}
	}

	items := make([]ServiceGkeService, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGkeService(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceGkeService expands an instance of ServiceGkeService into a JSON
// request object.
func expandServiceGkeService(c *Client, f *ServiceGkeService) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v := f.NamespaceName; !dcl.IsEmptyValueIndirect(v) {
		m["namespaceName"] = v
	}
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}

	return m, nil
}

// flattenServiceGkeService flattens an instance of ServiceGkeService from a JSON
// response object.
func flattenServiceGkeService(c *Client, i interface{}) *ServiceGkeService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceGkeService{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceGkeService
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.Location = dcl.FlattenString(m["location"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.NamespaceName = dcl.FlattenString(m["namespaceName"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandServiceTelemetryMap expands the contents of ServiceTelemetry into a JSON
// request object.
func expandServiceTelemetryMap(c *Client, f map[string]ServiceTelemetry) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTelemetry(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTelemetrySlice expands the contents of ServiceTelemetry into a JSON
// request object.
func expandServiceTelemetrySlice(c *Client, f []ServiceTelemetry) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTelemetry(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTelemetryMap flattens the contents of ServiceTelemetry from a JSON
// response object.
func flattenServiceTelemetryMap(c *Client, i interface{}) map[string]ServiceTelemetry {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTelemetry{}
	}

	if len(a) == 0 {
		return map[string]ServiceTelemetry{}
	}

	items := make(map[string]ServiceTelemetry)
	for k, item := range a {
		items[k] = *flattenServiceTelemetry(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTelemetrySlice flattens the contents of ServiceTelemetry from a JSON
// response object.
func flattenServiceTelemetrySlice(c *Client, i interface{}) []ServiceTelemetry {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTelemetry{}
	}

	if len(a) == 0 {
		return []ServiceTelemetry{}
	}

	items := make([]ServiceTelemetry, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTelemetry(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTelemetry expands an instance of ServiceTelemetry into a JSON
// request object.
func expandServiceTelemetry(c *Client, f *ServiceTelemetry) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceName; !dcl.IsEmptyValueIndirect(v) {
		m["resourceName"] = v
	}

	return m, nil
}

// flattenServiceTelemetry flattens an instance of ServiceTelemetry from a JSON
// response object.
func flattenServiceTelemetry(c *Client, i interface{}) *ServiceTelemetry {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTelemetry{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTelemetry
	}
	r.ResourceName = dcl.FlattenString(m["resourceName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Service) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalService(b, c)
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

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
}

func convertFieldDiffsToServiceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceDiff, error) {
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
	var diffs []serviceDiff
	// For each operation name, create a serviceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := serviceDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToServiceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToServiceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceApiOperation, error) {
	switch opName {

	case "updateServiceUpdateServiceOperation":
		return &updateServiceUpdateServiceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractServiceFields(r *Service) error {
	vCustom := r.Custom
	if vCustom == nil {
		// note: explicitly not the empty object.
		vCustom = &ServiceCustom{}
	}
	if err := extractServiceCustomFields(r, vCustom); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCustom) {
		r.Custom = vCustom
	}
	vAppEngine := r.AppEngine
	if vAppEngine == nil {
		// note: explicitly not the empty object.
		vAppEngine = &ServiceAppEngine{}
	}
	if err := extractServiceAppEngineFields(r, vAppEngine); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAppEngine) {
		r.AppEngine = vAppEngine
	}
	vCloudEndpoints := r.CloudEndpoints
	if vCloudEndpoints == nil {
		// note: explicitly not the empty object.
		vCloudEndpoints = &ServiceCloudEndpoints{}
	}
	if err := extractServiceCloudEndpointsFields(r, vCloudEndpoints); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudEndpoints) {
		r.CloudEndpoints = vCloudEndpoints
	}
	vClusterIstio := r.ClusterIstio
	if vClusterIstio == nil {
		// note: explicitly not the empty object.
		vClusterIstio = &ServiceClusterIstio{}
	}
	if err := extractServiceClusterIstioFields(r, vClusterIstio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vClusterIstio) {
		r.ClusterIstio = vClusterIstio
	}
	vMeshIstio := r.MeshIstio
	if vMeshIstio == nil {
		// note: explicitly not the empty object.
		vMeshIstio = &ServiceMeshIstio{}
	}
	if err := extractServiceMeshIstioFields(r, vMeshIstio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMeshIstio) {
		r.MeshIstio = vMeshIstio
	}
	vIstioCanonicalService := r.IstioCanonicalService
	if vIstioCanonicalService == nil {
		// note: explicitly not the empty object.
		vIstioCanonicalService = &ServiceIstioCanonicalService{}
	}
	if err := extractServiceIstioCanonicalServiceFields(r, vIstioCanonicalService); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vIstioCanonicalService) {
		r.IstioCanonicalService = vIstioCanonicalService
	}
	vCloudRun := r.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &ServiceCloudRun{}
	}
	if err := extractServiceCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudRun) {
		r.CloudRun = vCloudRun
	}
	vGkeNamespace := r.GkeNamespace
	if vGkeNamespace == nil {
		// note: explicitly not the empty object.
		vGkeNamespace = &ServiceGkeNamespace{}
	}
	if err := extractServiceGkeNamespaceFields(r, vGkeNamespace); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeNamespace) {
		r.GkeNamespace = vGkeNamespace
	}
	vGkeWorkload := r.GkeWorkload
	if vGkeWorkload == nil {
		// note: explicitly not the empty object.
		vGkeWorkload = &ServiceGkeWorkload{}
	}
	if err := extractServiceGkeWorkloadFields(r, vGkeWorkload); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeWorkload) {
		r.GkeWorkload = vGkeWorkload
	}
	vGkeService := r.GkeService
	if vGkeService == nil {
		// note: explicitly not the empty object.
		vGkeService = &ServiceGkeService{}
	}
	if err := extractServiceGkeServiceFields(r, vGkeService); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeService) {
		r.GkeService = vGkeService
	}
	vTelemetry := r.Telemetry
	if vTelemetry == nil {
		// note: explicitly not the empty object.
		vTelemetry = &ServiceTelemetry{}
	}
	if err := extractServiceTelemetryFields(r, vTelemetry); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTelemetry) {
		r.Telemetry = vTelemetry
	}
	return nil
}
func extractServiceCustomFields(r *Service, o *ServiceCustom) error {
	return nil
}
func extractServiceAppEngineFields(r *Service, o *ServiceAppEngine) error {
	return nil
}
func extractServiceCloudEndpointsFields(r *Service, o *ServiceCloudEndpoints) error {
	return nil
}
func extractServiceClusterIstioFields(r *Service, o *ServiceClusterIstio) error {
	return nil
}
func extractServiceMeshIstioFields(r *Service, o *ServiceMeshIstio) error {
	return nil
}
func extractServiceIstioCanonicalServiceFields(r *Service, o *ServiceIstioCanonicalService) error {
	return nil
}
func extractServiceCloudRunFields(r *Service, o *ServiceCloudRun) error {
	return nil
}
func extractServiceGkeNamespaceFields(r *Service, o *ServiceGkeNamespace) error {
	return nil
}
func extractServiceGkeWorkloadFields(r *Service, o *ServiceGkeWorkload) error {
	return nil
}
func extractServiceGkeServiceFields(r *Service, o *ServiceGkeService) error {
	return nil
}
func extractServiceTelemetryFields(r *Service, o *ServiceTelemetry) error {
	return nil
}

func postReadExtractServiceFields(r *Service) error {
	vCustom := r.Custom
	if vCustom == nil {
		// note: explicitly not the empty object.
		vCustom = &ServiceCustom{}
	}
	if err := postReadExtractServiceCustomFields(r, vCustom); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCustom) {
		r.Custom = vCustom
	}
	vAppEngine := r.AppEngine
	if vAppEngine == nil {
		// note: explicitly not the empty object.
		vAppEngine = &ServiceAppEngine{}
	}
	if err := postReadExtractServiceAppEngineFields(r, vAppEngine); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAppEngine) {
		r.AppEngine = vAppEngine
	}
	vCloudEndpoints := r.CloudEndpoints
	if vCloudEndpoints == nil {
		// note: explicitly not the empty object.
		vCloudEndpoints = &ServiceCloudEndpoints{}
	}
	if err := postReadExtractServiceCloudEndpointsFields(r, vCloudEndpoints); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudEndpoints) {
		r.CloudEndpoints = vCloudEndpoints
	}
	vClusterIstio := r.ClusterIstio
	if vClusterIstio == nil {
		// note: explicitly not the empty object.
		vClusterIstio = &ServiceClusterIstio{}
	}
	if err := postReadExtractServiceClusterIstioFields(r, vClusterIstio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vClusterIstio) {
		r.ClusterIstio = vClusterIstio
	}
	vMeshIstio := r.MeshIstio
	if vMeshIstio == nil {
		// note: explicitly not the empty object.
		vMeshIstio = &ServiceMeshIstio{}
	}
	if err := postReadExtractServiceMeshIstioFields(r, vMeshIstio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMeshIstio) {
		r.MeshIstio = vMeshIstio
	}
	vIstioCanonicalService := r.IstioCanonicalService
	if vIstioCanonicalService == nil {
		// note: explicitly not the empty object.
		vIstioCanonicalService = &ServiceIstioCanonicalService{}
	}
	if err := postReadExtractServiceIstioCanonicalServiceFields(r, vIstioCanonicalService); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vIstioCanonicalService) {
		r.IstioCanonicalService = vIstioCanonicalService
	}
	vCloudRun := r.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &ServiceCloudRun{}
	}
	if err := postReadExtractServiceCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudRun) {
		r.CloudRun = vCloudRun
	}
	vGkeNamespace := r.GkeNamespace
	if vGkeNamespace == nil {
		// note: explicitly not the empty object.
		vGkeNamespace = &ServiceGkeNamespace{}
	}
	if err := postReadExtractServiceGkeNamespaceFields(r, vGkeNamespace); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeNamespace) {
		r.GkeNamespace = vGkeNamespace
	}
	vGkeWorkload := r.GkeWorkload
	if vGkeWorkload == nil {
		// note: explicitly not the empty object.
		vGkeWorkload = &ServiceGkeWorkload{}
	}
	if err := postReadExtractServiceGkeWorkloadFields(r, vGkeWorkload); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeWorkload) {
		r.GkeWorkload = vGkeWorkload
	}
	vGkeService := r.GkeService
	if vGkeService == nil {
		// note: explicitly not the empty object.
		vGkeService = &ServiceGkeService{}
	}
	if err := postReadExtractServiceGkeServiceFields(r, vGkeService); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGkeService) {
		r.GkeService = vGkeService
	}
	vTelemetry := r.Telemetry
	if vTelemetry == nil {
		// note: explicitly not the empty object.
		vTelemetry = &ServiceTelemetry{}
	}
	if err := postReadExtractServiceTelemetryFields(r, vTelemetry); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTelemetry) {
		r.Telemetry = vTelemetry
	}
	return nil
}
func postReadExtractServiceCustomFields(r *Service, o *ServiceCustom) error {
	return nil
}
func postReadExtractServiceAppEngineFields(r *Service, o *ServiceAppEngine) error {
	return nil
}
func postReadExtractServiceCloudEndpointsFields(r *Service, o *ServiceCloudEndpoints) error {
	return nil
}
func postReadExtractServiceClusterIstioFields(r *Service, o *ServiceClusterIstio) error {
	return nil
}
func postReadExtractServiceMeshIstioFields(r *Service, o *ServiceMeshIstio) error {
	return nil
}
func postReadExtractServiceIstioCanonicalServiceFields(r *Service, o *ServiceIstioCanonicalService) error {
	return nil
}
func postReadExtractServiceCloudRunFields(r *Service, o *ServiceCloudRun) error {
	return nil
}
func postReadExtractServiceGkeNamespaceFields(r *Service, o *ServiceGkeNamespace) error {
	return nil
}
func postReadExtractServiceGkeWorkloadFields(r *Service, o *ServiceGkeWorkload) error {
	return nil
}
func postReadExtractServiceGkeServiceFields(r *Service, o *ServiceGkeService) error {
	return nil
}
func postReadExtractServiceTelemetryFields(r *Service, o *ServiceTelemetry) error {
	return nil
}
