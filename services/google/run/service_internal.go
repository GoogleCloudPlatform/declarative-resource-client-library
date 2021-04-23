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
package run

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Service) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceMetadata) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CreateTime) {
		if err := r.CreateTime.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DeleteTime) {
		if err := r.DeleteTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceMetadataCreateTime) validate() error {
	return nil
}
func (r *ServiceMetadataOwnerReferences) validate() error {
	return nil
}
func (r *ServiceMetadataDeleteTime) validate() error {
	return nil
}
func (r *ServiceSpec) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Template) {
		if err := r.Template.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateMetadata) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CreateTime) {
		if err := r.CreateTime.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DeleteTime) {
		if err := r.DeleteTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateMetadataCreateTime) validate() error {
	return nil
}
func (r *ServiceSpecTemplateMetadataOwnerReferences) validate() error {
	return nil
}
func (r *ServiceSpecTemplateMetadataDeleteTime) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpec) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainers) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Resources) {
		if err := r.Resources.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LivenessProbe) {
		if err := r.LivenessProbe.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReadinessProbe) {
		if err := r.ReadinessProbe.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecurityContext) {
		if err := r.SecurityContext.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnv) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ValueFrom) {
		if err := r.ValueFrom.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvValueFrom) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigMapKeyRef) {
		if err := r.ConfigMapKeyRef.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecretKeyRef) {
		if err := r.SecretKeyRef.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LocalObjectReference) {
		if err := r.LocalObjectReference.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LocalObjectReference) {
		if err := r.LocalObjectReference.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersResources) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersPorts) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvFrom) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigMapRef) {
		if err := r.ConfigMapRef.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecretRef) {
		if err := r.SecretRef.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LocalObjectReference) {
		if err := r.LocalObjectReference.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRef) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LocalObjectReference) {
		if err := r.LocalObjectReference.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersVolumeMounts) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersLivenessProbe) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Exec) {
		if err := r.Exec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpGet) {
		if err := r.HttpGet.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TcpSocket) {
		if err := r.TcpSocket.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersLivenessProbeExec) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersReadinessProbe) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Exec) {
		if err := r.Exec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpGet) {
		if err := r.HttpGet.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TcpSocket) {
		if err := r.TcpSocket.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecContainersReadinessProbeExec) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecContainersSecurityContext) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecVolumes) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Secret) {
		if err := r.Secret.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConfigMap) {
		if err := r.ConfigMap.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceSpecTemplateSpecVolumesSecret) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecVolumesSecretItems) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecVolumesConfigMap) validate() error {
	return nil
}
func (r *ServiceSpecTemplateSpecVolumesConfigMapItems) validate() error {
	return nil
}
func (r *ServiceSpecTraffic) validate() error {
	return nil
}
func (r *ServiceStatus) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Address) {
		if err := r.Address.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceStatusConditions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LastTransitionTime) {
		if err := r.LastTransitionTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceStatusConditionsLastTransitionTime) validate() error {
	return nil
}
func (r *ServiceStatusTraffic) validate() error {
	return nil
}
func (r *ServiceStatusAddress) validate() error {
	return nil
}

func serviceGetURL(userBasePath string, r *Service) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("/apis/serving.knative.dev/v1/namespaces/{{project}}/services/{{name}}", "https://{{location}}-run.googleapis.com/", userBasePath, params), nil
}

func serviceListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("/apis/serving.knative.dev/v1/namespaces/{{project}}/services", "https://{{location}}-run.googleapis.com/", userBasePath, params), nil

}

func serviceCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("/apis/serving.knative.dev/v1/namespaces/{{project}}/services", "https://{{location}}-run.googleapis.com/", userBasePath, params), nil

}

func serviceDeleteURL(userBasePath string, r *Service) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("/apis/serving.knative.dev/v1/namespaces/{{project}}/services/{{name}}", "https://{{location}}-run.googleapis.com/", userBasePath, params), nil
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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceUpdateServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	_, err := c.GetService(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateService")
	if err != nil {
		return err
	}

	req, err := newUpdateServiceUpdateServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateServiceUpdateServiceRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listServiceRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := serviceListURL(c.Config.BasePath, project, location)
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
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listService(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Service, string, error) {
	b, err := c.listServiceRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Service
	for _, v := range m.Items {
		res := flattenService(c, v)
		res.Project = &project
		res.Location = &location
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

	_, err := c.GetService(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Service not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetService checking for existence. error: %v", err)
		return err
	}

	u, err := serviceDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Service: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetService(ctx, r.urlNormalized())
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
type createServiceOperation struct {
	response map[string]interface{}
}

func (op *createServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := serviceCreateURL(c.Config.BasePath, project, location)

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
	var o operations.KNativeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://{{location}}-run.googleapis.com/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetService(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {
	if dcl.IsZeroValue(r.ApiVersion) {
		r.ApiVersion = dcl.String("serving.knative.dev/v1")
	}
	if dcl.IsZeroValue(r.Kind) {
		r.Kind = dcl.String("Service")
	}

	u, err := serviceGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []serviceDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Service
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Service); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Service, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetService(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Service resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Service resource: %v", err)
		}
		c.Config.Logger.Info("Found that Service resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Service: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Service: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Service: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Service: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceInitialState(rawInitial, rawDesired *Service) (*Service, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceDesiredState(rawDesired, rawInitial *Service, opts ...dcl.ApplyOption) (*Service, error) {

	if dcl.IsZeroValue(rawDesired.ApiVersion) {
		rawDesired.ApiVersion = dcl.String("serving.knative.dev/v1")
	}

	if dcl.IsZeroValue(rawDesired.Kind) {
		rawDesired.Kind = dcl.String("Service")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Metadata = canonicalizeServiceMetadata(rawDesired.Metadata, nil, opts...)
		rawDesired.Spec = canonicalizeServiceSpec(rawDesired.Spec, nil, opts...)
		rawDesired.Status = canonicalizeServiceStatus(rawDesired.Status, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.ApiVersion, rawInitial.ApiVersion) {
		rawDesired.ApiVersion = rawInitial.ApiVersion
	}
	if dcl.StringCanonicalize(rawDesired.Kind, rawInitial.Kind) {
		rawDesired.Kind = rawInitial.Kind
	}
	rawDesired.Metadata = canonicalizeServiceMetadata(rawDesired.Metadata, rawInitial.Metadata, opts...)
	rawDesired.Spec = canonicalizeServiceSpec(rawDesired.Spec, rawInitial.Spec, opts...)
	rawDesired.Status = canonicalizeServiceStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}

	return rawDesired, nil
}

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ApiVersion) && dcl.IsEmptyValueIndirect(rawDesired.ApiVersion) {
		rawNew.ApiVersion = rawDesired.ApiVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.ApiVersion, rawNew.ApiVersion) {
			rawNew.ApiVersion = rawDesired.ApiVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Kind) && dcl.IsEmptyValueIndirect(rawDesired.Kind) {
		rawNew.Kind = rawDesired.Kind
	} else {
		if dcl.StringCanonicalize(rawDesired.Kind, rawNew.Kind) {
			rawNew.Kind = rawDesired.Kind
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Metadata) && dcl.IsEmptyValueIndirect(rawDesired.Metadata) {
		rawNew.Metadata = rawDesired.Metadata
	} else {
		rawNew.Metadata = canonicalizeNewServiceMetadata(c, rawDesired.Metadata, rawNew.Metadata)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Spec) && dcl.IsEmptyValueIndirect(rawDesired.Spec) {
		rawNew.Spec = rawDesired.Spec
	} else {
		rawNew.Spec = canonicalizeNewServiceSpec(c, rawDesired.Spec, rawNew.Spec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewServiceStatus(c, rawDesired.Status, rawNew.Status)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Name = rawDesired.Name

	return rawNew, nil
}

func canonicalizeServiceMetadata(des, initial *ServiceMetadata, opts ...dcl.ApplyOption) *ServiceMetadata {
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
	if dcl.StringCanonicalize(des.GenerateName, initial.GenerateName) || dcl.IsZeroValue(des.GenerateName) {
		des.GenerateName = initial.GenerateName
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		des.Namespace = initial.Namespace
	}
	if dcl.StringCanonicalize(des.SelfLink, initial.SelfLink) || dcl.IsZeroValue(des.SelfLink) {
		des.SelfLink = initial.SelfLink
	}
	if dcl.StringCanonicalize(des.Uid, initial.Uid) || dcl.IsZeroValue(des.Uid) {
		des.Uid = initial.Uid
	}
	if dcl.StringCanonicalize(des.ResourceVersion, initial.ResourceVersion) || dcl.IsZeroValue(des.ResourceVersion) {
		des.ResourceVersion = initial.ResourceVersion
	}
	if dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}
	des.CreateTime = canonicalizeServiceMetadataCreateTime(des.CreateTime, initial.CreateTime, opts...)
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.Annotations) {
		des.Annotations = initial.Annotations
	}
	if dcl.IsZeroValue(des.OwnerReferences) {
		des.OwnerReferences = initial.OwnerReferences
	}
	des.DeleteTime = canonicalizeServiceMetadataDeleteTime(des.DeleteTime, initial.DeleteTime, opts...)
	if dcl.IsZeroValue(des.DeletionGracePeriodSeconds) {
		des.DeletionGracePeriodSeconds = initial.DeletionGracePeriodSeconds
	}
	if dcl.IsZeroValue(des.Finalizers) {
		des.Finalizers = initial.Finalizers
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		des.ClusterName = initial.ClusterName
	}

	return des
}

func canonicalizeNewServiceMetadata(c *Client, des, nw *ServiceMetadata) *ServiceMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.GenerateName, nw.GenerateName) {
		nw.GenerateName = des.GenerateName
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}
	if dcl.StringCanonicalize(des.SelfLink, nw.SelfLink) {
		nw.SelfLink = des.SelfLink
	}
	if dcl.StringCanonicalize(des.Uid, nw.Uid) {
		nw.Uid = des.Uid
	}
	if dcl.StringCanonicalize(des.ResourceVersion, nw.ResourceVersion) {
		nw.ResourceVersion = des.ResourceVersion
	}
	nw.CreateTime = canonicalizeNewServiceMetadataCreateTime(c, des.CreateTime, nw.CreateTime)
	nw.OwnerReferences = canonicalizeNewServiceMetadataOwnerReferencesSlice(c, des.OwnerReferences, nw.OwnerReferences)
	nw.DeleteTime = canonicalizeNewServiceMetadataDeleteTime(c, des.DeleteTime, nw.DeleteTime)
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}

	return nw
}

func canonicalizeNewServiceMetadataSet(c *Client, des, nw []ServiceMetadata) []ServiceMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceMetadata(c, &d, &n) {
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

func canonicalizeNewServiceMetadataSlice(c *Client, des, nw []ServiceMetadata) []ServiceMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeServiceMetadataCreateTime(des, initial *ServiceMetadataCreateTime, opts ...dcl.ApplyOption) *ServiceMetadataCreateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewServiceMetadataCreateTime(c *Client, des, nw *ServiceMetadataCreateTime) *ServiceMetadataCreateTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceMetadataCreateTimeSet(c *Client, des, nw []ServiceMetadataCreateTime) []ServiceMetadataCreateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceMetadataCreateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceMetadataCreateTime(c, &d, &n) {
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

func canonicalizeNewServiceMetadataCreateTimeSlice(c *Client, des, nw []ServiceMetadataCreateTime) []ServiceMetadataCreateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceMetadataCreateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceMetadataCreateTime(c, &d, &n))
	}

	return items
}

func canonicalizeServiceMetadataOwnerReferences(des, initial *ServiceMetadataOwnerReferences, opts ...dcl.ApplyOption) *ServiceMetadataOwnerReferences {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ApiVersion, initial.ApiVersion) || dcl.IsZeroValue(des.ApiVersion) {
		des.ApiVersion = initial.ApiVersion
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Uid, initial.Uid) || dcl.IsZeroValue(des.Uid) {
		des.Uid = initial.Uid
	}
	if dcl.BoolCanonicalize(des.Controller, initial.Controller) || dcl.IsZeroValue(des.Controller) {
		des.Controller = initial.Controller
	}
	if dcl.BoolCanonicalize(des.BlockOwnerDeletion, initial.BlockOwnerDeletion) || dcl.IsZeroValue(des.BlockOwnerDeletion) {
		des.BlockOwnerDeletion = initial.BlockOwnerDeletion
	}

	return des
}

func canonicalizeNewServiceMetadataOwnerReferences(c *Client, des, nw *ServiceMetadataOwnerReferences) *ServiceMetadataOwnerReferences {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ApiVersion, nw.ApiVersion) {
		nw.ApiVersion = des.ApiVersion
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Uid, nw.Uid) {
		nw.Uid = des.Uid
	}
	if dcl.BoolCanonicalize(des.Controller, nw.Controller) {
		nw.Controller = des.Controller
	}
	if dcl.BoolCanonicalize(des.BlockOwnerDeletion, nw.BlockOwnerDeletion) {
		nw.BlockOwnerDeletion = des.BlockOwnerDeletion
	}

	return nw
}

func canonicalizeNewServiceMetadataOwnerReferencesSet(c *Client, des, nw []ServiceMetadataOwnerReferences) []ServiceMetadataOwnerReferences {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceMetadataOwnerReferences
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceMetadataOwnerReferences(c, &d, &n) {
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

func canonicalizeNewServiceMetadataOwnerReferencesSlice(c *Client, des, nw []ServiceMetadataOwnerReferences) []ServiceMetadataOwnerReferences {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceMetadataOwnerReferences
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceMetadataOwnerReferences(c, &d, &n))
	}

	return items
}

func canonicalizeServiceMetadataDeleteTime(des, initial *ServiceMetadataDeleteTime, opts ...dcl.ApplyOption) *ServiceMetadataDeleteTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewServiceMetadataDeleteTime(c *Client, des, nw *ServiceMetadataDeleteTime) *ServiceMetadataDeleteTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceMetadataDeleteTimeSet(c *Client, des, nw []ServiceMetadataDeleteTime) []ServiceMetadataDeleteTime {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceMetadataDeleteTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceMetadataDeleteTime(c, &d, &n) {
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

func canonicalizeNewServiceMetadataDeleteTimeSlice(c *Client, des, nw []ServiceMetadataDeleteTime) []ServiceMetadataDeleteTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceMetadataDeleteTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceMetadataDeleteTime(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpec(des, initial *ServiceSpec, opts ...dcl.ApplyOption) *ServiceSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Template = canonicalizeServiceSpecTemplate(des.Template, initial.Template, opts...)
	if dcl.IsZeroValue(des.Traffic) {
		des.Traffic = initial.Traffic
	}

	return des
}

func canonicalizeNewServiceSpec(c *Client, des, nw *ServiceSpec) *ServiceSpec {
	if des == nil || nw == nil {
		return nw
	}

	nw.Template = canonicalizeNewServiceSpecTemplate(c, des.Template, nw.Template)
	nw.Traffic = canonicalizeNewServiceSpecTrafficSlice(c, des.Traffic, nw.Traffic)

	return nw
}

func canonicalizeNewServiceSpecSet(c *Client, des, nw []ServiceSpec) []ServiceSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpec(c, &d, &n) {
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

func canonicalizeNewServiceSpecSlice(c *Client, des, nw []ServiceSpec) []ServiceSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpec(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplate(des, initial *ServiceSpecTemplate, opts ...dcl.ApplyOption) *ServiceSpecTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Metadata = canonicalizeServiceSpecTemplateMetadata(des.Metadata, initial.Metadata, opts...)
	des.Spec = canonicalizeServiceSpecTemplateSpec(des.Spec, initial.Spec, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplate(c *Client, des, nw *ServiceSpecTemplate) *ServiceSpecTemplate {
	if des == nil || nw == nil {
		return nw
	}

	nw.Metadata = canonicalizeNewServiceSpecTemplateMetadata(c, des.Metadata, nw.Metadata)
	nw.Spec = canonicalizeNewServiceSpecTemplateSpec(c, des.Spec, nw.Spec)

	return nw
}

func canonicalizeNewServiceSpecTemplateSet(c *Client, des, nw []ServiceSpecTemplate) []ServiceSpecTemplate {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplate(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSlice(c *Client, des, nw []ServiceSpecTemplate) []ServiceSpecTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateMetadata(des, initial *ServiceSpecTemplateMetadata, opts ...dcl.ApplyOption) *ServiceSpecTemplateMetadata {
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
	if dcl.StringCanonicalize(des.GenerateName, initial.GenerateName) || dcl.IsZeroValue(des.GenerateName) {
		des.GenerateName = initial.GenerateName
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		des.Namespace = initial.Namespace
	}
	if dcl.StringCanonicalize(des.SelfLink, initial.SelfLink) || dcl.IsZeroValue(des.SelfLink) {
		des.SelfLink = initial.SelfLink
	}
	if dcl.StringCanonicalize(des.Uid, initial.Uid) || dcl.IsZeroValue(des.Uid) {
		des.Uid = initial.Uid
	}
	if dcl.StringCanonicalize(des.ResourceVersion, initial.ResourceVersion) || dcl.IsZeroValue(des.ResourceVersion) {
		des.ResourceVersion = initial.ResourceVersion
	}
	if dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}
	des.CreateTime = canonicalizeServiceSpecTemplateMetadataCreateTime(des.CreateTime, initial.CreateTime, opts...)
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.Annotations) {
		des.Annotations = initial.Annotations
	}
	if dcl.IsZeroValue(des.OwnerReferences) {
		des.OwnerReferences = initial.OwnerReferences
	}
	des.DeleteTime = canonicalizeServiceSpecTemplateMetadataDeleteTime(des.DeleteTime, initial.DeleteTime, opts...)
	if dcl.IsZeroValue(des.DeletionGracePeriodSeconds) {
		des.DeletionGracePeriodSeconds = initial.DeletionGracePeriodSeconds
	}
	if dcl.IsZeroValue(des.Finalizers) {
		des.Finalizers = initial.Finalizers
	}
	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		des.ClusterName = initial.ClusterName
	}

	return des
}

func canonicalizeNewServiceSpecTemplateMetadata(c *Client, des, nw *ServiceSpecTemplateMetadata) *ServiceSpecTemplateMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.GenerateName, nw.GenerateName) {
		nw.GenerateName = des.GenerateName
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}
	if dcl.StringCanonicalize(des.SelfLink, nw.SelfLink) {
		nw.SelfLink = des.SelfLink
	}
	if dcl.StringCanonicalize(des.Uid, nw.Uid) {
		nw.Uid = des.Uid
	}
	if dcl.StringCanonicalize(des.ResourceVersion, nw.ResourceVersion) {
		nw.ResourceVersion = des.ResourceVersion
	}
	nw.CreateTime = canonicalizeNewServiceSpecTemplateMetadataCreateTime(c, des.CreateTime, nw.CreateTime)
	nw.OwnerReferences = canonicalizeNewServiceSpecTemplateMetadataOwnerReferencesSlice(c, des.OwnerReferences, nw.OwnerReferences)
	nw.DeleteTime = canonicalizeNewServiceSpecTemplateMetadataDeleteTime(c, des.DeleteTime, nw.DeleteTime)
	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateMetadataSet(c *Client, des, nw []ServiceSpecTemplateMetadata) []ServiceSpecTemplateMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateMetadata(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateMetadataSlice(c *Client, des, nw []ServiceSpecTemplateMetadata) []ServiceSpecTemplateMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateMetadataCreateTime(des, initial *ServiceSpecTemplateMetadataCreateTime, opts ...dcl.ApplyOption) *ServiceSpecTemplateMetadataCreateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewServiceSpecTemplateMetadataCreateTime(c *Client, des, nw *ServiceSpecTemplateMetadataCreateTime) *ServiceSpecTemplateMetadataCreateTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateMetadataCreateTimeSet(c *Client, des, nw []ServiceSpecTemplateMetadataCreateTime) []ServiceSpecTemplateMetadataCreateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateMetadataCreateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateMetadataCreateTime(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateMetadataCreateTimeSlice(c *Client, des, nw []ServiceSpecTemplateMetadataCreateTime) []ServiceSpecTemplateMetadataCreateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateMetadataCreateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateMetadataCreateTime(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateMetadataOwnerReferences(des, initial *ServiceSpecTemplateMetadataOwnerReferences, opts ...dcl.ApplyOption) *ServiceSpecTemplateMetadataOwnerReferences {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ApiVersion, initial.ApiVersion) || dcl.IsZeroValue(des.ApiVersion) {
		des.ApiVersion = initial.ApiVersion
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Uid, initial.Uid) || dcl.IsZeroValue(des.Uid) {
		des.Uid = initial.Uid
	}
	if dcl.BoolCanonicalize(des.Controller, initial.Controller) || dcl.IsZeroValue(des.Controller) {
		des.Controller = initial.Controller
	}
	if dcl.BoolCanonicalize(des.BlockOwnerDeletion, initial.BlockOwnerDeletion) || dcl.IsZeroValue(des.BlockOwnerDeletion) {
		des.BlockOwnerDeletion = initial.BlockOwnerDeletion
	}

	return des
}

func canonicalizeNewServiceSpecTemplateMetadataOwnerReferences(c *Client, des, nw *ServiceSpecTemplateMetadataOwnerReferences) *ServiceSpecTemplateMetadataOwnerReferences {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ApiVersion, nw.ApiVersion) {
		nw.ApiVersion = des.ApiVersion
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Uid, nw.Uid) {
		nw.Uid = des.Uid
	}
	if dcl.BoolCanonicalize(des.Controller, nw.Controller) {
		nw.Controller = des.Controller
	}
	if dcl.BoolCanonicalize(des.BlockOwnerDeletion, nw.BlockOwnerDeletion) {
		nw.BlockOwnerDeletion = des.BlockOwnerDeletion
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateMetadataOwnerReferencesSet(c *Client, des, nw []ServiceSpecTemplateMetadataOwnerReferences) []ServiceSpecTemplateMetadataOwnerReferences {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateMetadataOwnerReferences
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateMetadataOwnerReferences(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateMetadataOwnerReferencesSlice(c *Client, des, nw []ServiceSpecTemplateMetadataOwnerReferences) []ServiceSpecTemplateMetadataOwnerReferences {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateMetadataOwnerReferences
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateMetadataOwnerReferences(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateMetadataDeleteTime(des, initial *ServiceSpecTemplateMetadataDeleteTime, opts ...dcl.ApplyOption) *ServiceSpecTemplateMetadataDeleteTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewServiceSpecTemplateMetadataDeleteTime(c *Client, des, nw *ServiceSpecTemplateMetadataDeleteTime) *ServiceSpecTemplateMetadataDeleteTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateMetadataDeleteTimeSet(c *Client, des, nw []ServiceSpecTemplateMetadataDeleteTime) []ServiceSpecTemplateMetadataDeleteTime {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateMetadataDeleteTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateMetadataDeleteTime(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateMetadataDeleteTimeSlice(c *Client, des, nw []ServiceSpecTemplateMetadataDeleteTime) []ServiceSpecTemplateMetadataDeleteTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateMetadataDeleteTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateMetadataDeleteTime(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpec(des, initial *ServiceSpecTemplateSpec, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ContainerConcurrency) {
		des.ContainerConcurrency = initial.ContainerConcurrency
	}
	if dcl.IsZeroValue(des.TimeoutSeconds) {
		des.TimeoutSeconds = initial.TimeoutSeconds
	}
	if dcl.StringCanonicalize(des.ServiceAccountName, initial.ServiceAccountName) || dcl.IsZeroValue(des.ServiceAccountName) {
		des.ServiceAccountName = initial.ServiceAccountName
	}
	if dcl.IsZeroValue(des.Containers) {
		des.Containers = initial.Containers
	}
	if dcl.IsZeroValue(des.Volumes) {
		des.Volumes = initial.Volumes
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpec(c *Client, des, nw *ServiceSpecTemplateSpec) *ServiceSpecTemplateSpec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ServiceAccountName, nw.ServiceAccountName) {
		nw.ServiceAccountName = des.ServiceAccountName
	}
	nw.Containers = canonicalizeNewServiceSpecTemplateSpecContainersSlice(c, des.Containers, nw.Containers)
	nw.Volumes = canonicalizeNewServiceSpecTemplateSpecVolumesSlice(c, des.Volumes, nw.Volumes)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecSet(c *Client, des, nw []ServiceSpecTemplateSpec) []ServiceSpecTemplateSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpec(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecSlice(c *Client, des, nw []ServiceSpecTemplateSpec) []ServiceSpecTemplateSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpec(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainers(des, initial *ServiceSpecTemplateSpecContainers, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainers {
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
	if dcl.StringCanonicalize(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		des.Image = initial.Image
	}
	if dcl.IsZeroValue(des.Command) {
		des.Command = initial.Command
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.Env) {
		des.Env = initial.Env
	}
	des.Resources = canonicalizeServiceSpecTemplateSpecContainersResources(des.Resources, initial.Resources, opts...)
	if dcl.StringCanonicalize(des.WorkingDir, initial.WorkingDir) || dcl.IsZeroValue(des.WorkingDir) {
		des.WorkingDir = initial.WorkingDir
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	}
	if dcl.IsZeroValue(des.EnvFrom) {
		des.EnvFrom = initial.EnvFrom
	}
	if dcl.IsZeroValue(des.VolumeMounts) {
		des.VolumeMounts = initial.VolumeMounts
	}
	des.LivenessProbe = canonicalizeServiceSpecTemplateSpecContainersLivenessProbe(des.LivenessProbe, initial.LivenessProbe, opts...)
	des.ReadinessProbe = canonicalizeServiceSpecTemplateSpecContainersReadinessProbe(des.ReadinessProbe, initial.ReadinessProbe, opts...)
	if dcl.StringCanonicalize(des.TerminationMessagePath, initial.TerminationMessagePath) || dcl.IsZeroValue(des.TerminationMessagePath) {
		des.TerminationMessagePath = initial.TerminationMessagePath
	}
	if dcl.StringCanonicalize(des.TerminationMessagePolicy, initial.TerminationMessagePolicy) || dcl.IsZeroValue(des.TerminationMessagePolicy) {
		des.TerminationMessagePolicy = initial.TerminationMessagePolicy
	}
	if dcl.StringCanonicalize(des.ImagePullPolicy, initial.ImagePullPolicy) || dcl.IsZeroValue(des.ImagePullPolicy) {
		des.ImagePullPolicy = initial.ImagePullPolicy
	}
	des.SecurityContext = canonicalizeServiceSpecTemplateSpecContainersSecurityContext(des.SecurityContext, initial.SecurityContext, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainers(c *Client, des, nw *ServiceSpecTemplateSpecContainers) *ServiceSpecTemplateSpecContainers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	nw.Env = canonicalizeNewServiceSpecTemplateSpecContainersEnvSlice(c, des.Env, nw.Env)
	nw.Resources = canonicalizeNewServiceSpecTemplateSpecContainersResources(c, des.Resources, nw.Resources)
	if dcl.StringCanonicalize(des.WorkingDir, nw.WorkingDir) {
		nw.WorkingDir = des.WorkingDir
	}
	nw.Ports = canonicalizeNewServiceSpecTemplateSpecContainersPortsSlice(c, des.Ports, nw.Ports)
	nw.EnvFrom = canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSlice(c, des.EnvFrom, nw.EnvFrom)
	nw.VolumeMounts = canonicalizeNewServiceSpecTemplateSpecContainersVolumeMountsSlice(c, des.VolumeMounts, nw.VolumeMounts)
	nw.LivenessProbe = canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbe(c, des.LivenessProbe, nw.LivenessProbe)
	nw.ReadinessProbe = canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbe(c, des.ReadinessProbe, nw.ReadinessProbe)
	if dcl.StringCanonicalize(des.TerminationMessagePath, nw.TerminationMessagePath) {
		nw.TerminationMessagePath = des.TerminationMessagePath
	}
	if dcl.StringCanonicalize(des.TerminationMessagePolicy, nw.TerminationMessagePolicy) {
		nw.TerminationMessagePolicy = des.TerminationMessagePolicy
	}
	if dcl.StringCanonicalize(des.ImagePullPolicy, nw.ImagePullPolicy) {
		nw.ImagePullPolicy = des.ImagePullPolicy
	}
	nw.SecurityContext = canonicalizeNewServiceSpecTemplateSpecContainersSecurityContext(c, des.SecurityContext, nw.SecurityContext)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersSet(c *Client, des, nw []ServiceSpecTemplateSpecContainers) []ServiceSpecTemplateSpecContainers {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainers(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainers) []ServiceSpecTemplateSpecContainers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainers(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnv(des, initial *ServiceSpecTemplateSpecContainersEnv, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnv {
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
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	des.ValueFrom = canonicalizeServiceSpecTemplateSpecContainersEnvValueFrom(des.ValueFrom, initial.ValueFrom, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnv(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnv) *ServiceSpecTemplateSpecContainersEnv {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	nw.ValueFrom = canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFrom(c, des.ValueFrom, nw.ValueFrom)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnv) []ServiceSpecTemplateSpecContainersEnv {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnv
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnv(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnv) []ServiceSpecTemplateSpecContainersEnv {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnv
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnv(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvValueFrom(des, initial *ServiceSpecTemplateSpecContainersEnvValueFrom, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvValueFrom {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ConfigMapKeyRef = canonicalizeServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(des.ConfigMapKeyRef, initial.ConfigMapKeyRef, opts...)
	des.SecretKeyRef = canonicalizeServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(des.SecretKeyRef, initial.SecretKeyRef, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFrom(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvValueFrom) *ServiceSpecTemplateSpecContainersEnvValueFrom {
	if des == nil || nw == nil {
		return nw
	}

	nw.ConfigMapKeyRef = canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, des.ConfigMapKeyRef, nw.ConfigMapKeyRef)
	nw.SecretKeyRef = canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, des.SecretKeyRef, nw.SecretKeyRef)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFrom) []ServiceSpecTemplateSpecContainersEnvValueFrom {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvValueFrom
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvValueFrom(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFrom) []ServiceSpecTemplateSpecContainersEnvValueFrom {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvValueFrom
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFrom(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(des, initial *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LocalObjectReference = canonicalizeServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(des.LocalObjectReference, initial.LocalObjectReference, opts...)
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if des == nil || nw == nil {
		return nw
	}

	nw.LocalObjectReference = canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, des.LocalObjectReference, nw.LocalObjectReference)
	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(des, initial *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(des, initial *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LocalObjectReference = canonicalizeServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(des.LocalObjectReference, initial.LocalObjectReference, opts...)
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if des == nil || nw == nil {
		return nw
	}

	nw.LocalObjectReference = canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, des.LocalObjectReference, nw.LocalObjectReference)
	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(des, initial *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersResources(des, initial *ServiceSpecTemplateSpecContainersResources, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Limits) {
		des.Limits = initial.Limits
	}
	if dcl.IsZeroValue(des.Requests) {
		des.Requests = initial.Requests
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersResources(c *Client, des, nw *ServiceSpecTemplateSpecContainersResources) *ServiceSpecTemplateSpecContainersResources {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersResourcesSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersResources) []ServiceSpecTemplateSpecContainersResources {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersResources(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersResourcesSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersResources) []ServiceSpecTemplateSpecContainersResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersResources(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersPorts(des, initial *ServiceSpecTemplateSpecContainersPorts, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersPorts {
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
	if dcl.IsZeroValue(des.ContainerPort) {
		des.ContainerPort = initial.ContainerPort
	}
	if dcl.StringCanonicalize(des.Protocol, initial.Protocol) || dcl.IsZeroValue(des.Protocol) {
		des.Protocol = initial.Protocol
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersPorts(c *Client, des, nw *ServiceSpecTemplateSpecContainersPorts) *ServiceSpecTemplateSpecContainersPorts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Protocol, nw.Protocol) {
		nw.Protocol = des.Protocol
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersPortsSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersPorts) []ServiceSpecTemplateSpecContainersPorts {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersPorts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersPorts(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersPortsSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersPorts) []ServiceSpecTemplateSpecContainersPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersPorts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvFrom(des, initial *ServiceSpecTemplateSpecContainersEnvFrom, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvFrom {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Prefix, initial.Prefix) || dcl.IsZeroValue(des.Prefix) {
		des.Prefix = initial.Prefix
	}
	des.ConfigMapRef = canonicalizeServiceSpecTemplateSpecContainersEnvFromConfigMapRef(des.ConfigMapRef, initial.ConfigMapRef, opts...)
	des.SecretRef = canonicalizeServiceSpecTemplateSpecContainersEnvFromSecretRef(des.SecretRef, initial.SecretRef, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFrom(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvFrom) *ServiceSpecTemplateSpecContainersEnvFrom {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Prefix, nw.Prefix) {
		nw.Prefix = des.Prefix
	}
	nw.ConfigMapRef = canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, des.ConfigMapRef, nw.ConfigMapRef)
	nw.SecretRef = canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRef(c, des.SecretRef, nw.SecretRef)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFrom) []ServiceSpecTemplateSpecContainersEnvFrom {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvFrom
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvFrom(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFrom) []ServiceSpecTemplateSpecContainersEnvFrom {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvFrom
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvFrom(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvFromConfigMapRef(des, initial *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LocalObjectReference = canonicalizeServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(des.LocalObjectReference, initial.LocalObjectReference, opts...)
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if des == nil || nw == nil {
		return nw
	}

	nw.LocalObjectReference = canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, des.LocalObjectReference, nw.LocalObjectReference)
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(des, initial *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvFromSecretRef(des, initial *ServiceSpecTemplateSpecContainersEnvFromSecretRef, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LocalObjectReference = canonicalizeServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(des.LocalObjectReference, initial.LocalObjectReference, opts...)
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRef(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvFromSecretRef) *ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if des == nil || nw == nil {
		return nw
	}

	nw.LocalObjectReference = canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, des.LocalObjectReference, nw.LocalObjectReference)
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromSecretRef) []ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvFromSecretRef
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromSecretRef) []ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvFromSecretRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(des, initial *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c *Client, des, nw *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersVolumeMounts(des, initial *ServiceSpecTemplateSpecContainersVolumeMounts, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersVolumeMounts {
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
	if dcl.BoolCanonicalize(des.ReadOnly, initial.ReadOnly) || dcl.IsZeroValue(des.ReadOnly) {
		des.ReadOnly = initial.ReadOnly
	}
	if dcl.StringCanonicalize(des.MountPath, initial.MountPath) || dcl.IsZeroValue(des.MountPath) {
		des.MountPath = initial.MountPath
	}
	if dcl.StringCanonicalize(des.SubPath, initial.SubPath) || dcl.IsZeroValue(des.SubPath) {
		des.SubPath = initial.SubPath
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersVolumeMounts(c *Client, des, nw *ServiceSpecTemplateSpecContainersVolumeMounts) *ServiceSpecTemplateSpecContainersVolumeMounts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.BoolCanonicalize(des.ReadOnly, nw.ReadOnly) {
		nw.ReadOnly = des.ReadOnly
	}
	if dcl.StringCanonicalize(des.MountPath, nw.MountPath) {
		nw.MountPath = des.MountPath
	}
	if dcl.StringCanonicalize(des.SubPath, nw.SubPath) {
		nw.SubPath = des.SubPath
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersVolumeMountsSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersVolumeMounts) []ServiceSpecTemplateSpecContainersVolumeMounts {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersVolumeMounts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersVolumeMounts(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersVolumeMountsSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersVolumeMounts) []ServiceSpecTemplateSpecContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersVolumeMounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersVolumeMounts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersLivenessProbe(des, initial *ServiceSpecTemplateSpecContainersLivenessProbe, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersLivenessProbe {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.InitialDelaySeconds) {
		des.InitialDelaySeconds = initial.InitialDelaySeconds
	}
	if dcl.IsZeroValue(des.TimeoutSeconds) {
		des.TimeoutSeconds = initial.TimeoutSeconds
	}
	if dcl.IsZeroValue(des.PeriodSeconds) {
		des.PeriodSeconds = initial.PeriodSeconds
	}
	if dcl.IsZeroValue(des.SuccessThreshold) {
		des.SuccessThreshold = initial.SuccessThreshold
	}
	if dcl.IsZeroValue(des.FailureThreshold) {
		des.FailureThreshold = initial.FailureThreshold
	}
	des.Exec = canonicalizeServiceSpecTemplateSpecContainersLivenessProbeExec(des.Exec, initial.Exec, opts...)
	des.HttpGet = canonicalizeServiceSpecTemplateSpecContainersLivenessProbeHttpGet(des.HttpGet, initial.HttpGet, opts...)
	des.TcpSocket = canonicalizeServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(des.TcpSocket, initial.TcpSocket, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbe(c *Client, des, nw *ServiceSpecTemplateSpecContainersLivenessProbe) *ServiceSpecTemplateSpecContainersLivenessProbe {
	if des == nil || nw == nil {
		return nw
	}

	nw.Exec = canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeExec(c, des.Exec, nw.Exec)
	nw.HttpGet = canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, des.HttpGet, nw.HttpGet)
	nw.TcpSocket = canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, des.TcpSocket, nw.TcpSocket)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbe) []ServiceSpecTemplateSpecContainersLivenessProbe {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersLivenessProbe
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersLivenessProbe(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbe) []ServiceSpecTemplateSpecContainersLivenessProbe {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersLivenessProbe
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbe(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersLivenessProbeExec(des, initial *ServiceSpecTemplateSpecContainersLivenessProbeExec, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersLivenessProbeExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Command, initial.Command) || dcl.IsZeroValue(des.Command) {
		des.Command = initial.Command
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeExec(c *Client, des, nw *ServiceSpecTemplateSpecContainersLivenessProbeExec) *ServiceSpecTemplateSpecContainersLivenessProbeExec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Command, nw.Command) {
		nw.Command = des.Command
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeExecSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeExec) []ServiceSpecTemplateSpecContainersLivenessProbeExec {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersLivenessProbeExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersLivenessProbeExec(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeExecSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeExec) []ServiceSpecTemplateSpecContainersLivenessProbeExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersLivenessProbeExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeExec(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersLivenessProbeHttpGet(des, initial *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
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
	if dcl.StringCanonicalize(des.Scheme, initial.Scheme) || dcl.IsZeroValue(des.Scheme) {
		des.Scheme = initial.Scheme
	}
	if dcl.IsZeroValue(des.HttpHeaders) {
		des.HttpHeaders = initial.HttpHeaders
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c *Client, des, nw *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.Scheme, nw.Scheme) {
		nw.Scheme = des.Scheme
	}
	nw.HttpHeaders = canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c, des.HttpHeaders, nw.HttpHeaders)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(des, initial *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
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
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c *Client, des, nw *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(des, initial *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c *Client, des, nw *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersReadinessProbe(des, initial *ServiceSpecTemplateSpecContainersReadinessProbe, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersReadinessProbe {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.InitialDelaySeconds) {
		des.InitialDelaySeconds = initial.InitialDelaySeconds
	}
	if dcl.IsZeroValue(des.TimeoutSeconds) {
		des.TimeoutSeconds = initial.TimeoutSeconds
	}
	if dcl.IsZeroValue(des.PeriodSeconds) {
		des.PeriodSeconds = initial.PeriodSeconds
	}
	if dcl.IsZeroValue(des.SuccessThreshold) {
		des.SuccessThreshold = initial.SuccessThreshold
	}
	if dcl.IsZeroValue(des.FailureThreshold) {
		des.FailureThreshold = initial.FailureThreshold
	}
	des.Exec = canonicalizeServiceSpecTemplateSpecContainersReadinessProbeExec(des.Exec, initial.Exec, opts...)
	des.HttpGet = canonicalizeServiceSpecTemplateSpecContainersReadinessProbeHttpGet(des.HttpGet, initial.HttpGet, opts...)
	des.TcpSocket = canonicalizeServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(des.TcpSocket, initial.TcpSocket, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbe(c *Client, des, nw *ServiceSpecTemplateSpecContainersReadinessProbe) *ServiceSpecTemplateSpecContainersReadinessProbe {
	if des == nil || nw == nil {
		return nw
	}

	nw.Exec = canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeExec(c, des.Exec, nw.Exec)
	nw.HttpGet = canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, des.HttpGet, nw.HttpGet)
	nw.TcpSocket = canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, des.TcpSocket, nw.TcpSocket)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbe) []ServiceSpecTemplateSpecContainersReadinessProbe {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersReadinessProbe
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersReadinessProbe(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbe) []ServiceSpecTemplateSpecContainersReadinessProbe {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersReadinessProbe
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbe(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersReadinessProbeExec(des, initial *ServiceSpecTemplateSpecContainersReadinessProbeExec, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersReadinessProbeExec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Command, initial.Command) || dcl.IsZeroValue(des.Command) {
		des.Command = initial.Command
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeExec(c *Client, des, nw *ServiceSpecTemplateSpecContainersReadinessProbeExec) *ServiceSpecTemplateSpecContainersReadinessProbeExec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Command, nw.Command) {
		nw.Command = des.Command
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeExecSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeExec) []ServiceSpecTemplateSpecContainersReadinessProbeExec {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersReadinessProbeExec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersReadinessProbeExec(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeExecSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeExec) []ServiceSpecTemplateSpecContainersReadinessProbeExec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersReadinessProbeExec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeExec(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersReadinessProbeHttpGet(des, initial *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
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
	if dcl.StringCanonicalize(des.Scheme, initial.Scheme) || dcl.IsZeroValue(des.Scheme) {
		des.Scheme = initial.Scheme
	}
	if dcl.IsZeroValue(des.HttpHeaders) {
		des.HttpHeaders = initial.HttpHeaders
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c *Client, des, nw *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.Scheme, nw.Scheme) {
		nw.Scheme = des.Scheme
	}
	nw.HttpHeaders = canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c, des.HttpHeaders, nw.HttpHeaders)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(des, initial *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
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
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c *Client, des, nw *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(des, initial *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c *Client, des, nw *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecContainersSecurityContext(des, initial *ServiceSpecTemplateSpecContainersSecurityContext, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecContainersSecurityContext {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RunAsUser) {
		des.RunAsUser = initial.RunAsUser
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecContainersSecurityContext(c *Client, des, nw *ServiceSpecTemplateSpecContainersSecurityContext) *ServiceSpecTemplateSpecContainersSecurityContext {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecContainersSecurityContextSet(c *Client, des, nw []ServiceSpecTemplateSpecContainersSecurityContext) []ServiceSpecTemplateSpecContainersSecurityContext {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecContainersSecurityContext
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecContainersSecurityContext(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecContainersSecurityContextSlice(c *Client, des, nw []ServiceSpecTemplateSpecContainersSecurityContext) []ServiceSpecTemplateSpecContainersSecurityContext {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecContainersSecurityContext
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecContainersSecurityContext(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecVolumes(des, initial *ServiceSpecTemplateSpecVolumes, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecVolumes {
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
	des.Secret = canonicalizeServiceSpecTemplateSpecVolumesSecret(des.Secret, initial.Secret, opts...)
	des.ConfigMap = canonicalizeServiceSpecTemplateSpecVolumesConfigMap(des.ConfigMap, initial.ConfigMap, opts...)

	return des
}

func canonicalizeNewServiceSpecTemplateSpecVolumes(c *Client, des, nw *ServiceSpecTemplateSpecVolumes) *ServiceSpecTemplateSpecVolumes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Secret = canonicalizeNewServiceSpecTemplateSpecVolumesSecret(c, des.Secret, nw.Secret)
	nw.ConfigMap = canonicalizeNewServiceSpecTemplateSpecVolumesConfigMap(c, des.ConfigMap, nw.ConfigMap)

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecVolumesSet(c *Client, des, nw []ServiceSpecTemplateSpecVolumes) []ServiceSpecTemplateSpecVolumes {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecVolumes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecVolumes(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecVolumesSlice(c *Client, des, nw []ServiceSpecTemplateSpecVolumes) []ServiceSpecTemplateSpecVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecVolumesSecret(des, initial *ServiceSpecTemplateSpecVolumesSecret, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecVolumesSecret {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SecretName, initial.SecretName) || dcl.IsZeroValue(des.SecretName) {
		des.SecretName = initial.SecretName
	}
	if dcl.IsZeroValue(des.Items) {
		des.Items = initial.Items
	}
	if dcl.IsZeroValue(des.DefaultMode) {
		des.DefaultMode = initial.DefaultMode
	}
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecVolumesSecret(c *Client, des, nw *ServiceSpecTemplateSpecVolumesSecret) *ServiceSpecTemplateSpecVolumesSecret {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SecretName, nw.SecretName) {
		nw.SecretName = des.SecretName
	}
	nw.Items = canonicalizeNewServiceSpecTemplateSpecVolumesSecretItemsSlice(c, des.Items, nw.Items)
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecVolumesSecretSet(c *Client, des, nw []ServiceSpecTemplateSpecVolumesSecret) []ServiceSpecTemplateSpecVolumesSecret {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecVolumesSecret
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecVolumesSecret(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecVolumesSecretSlice(c *Client, des, nw []ServiceSpecTemplateSpecVolumesSecret) []ServiceSpecTemplateSpecVolumesSecret {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecVolumesSecret
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecVolumesSecret(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecVolumesSecretItems(des, initial *ServiceSpecTemplateSpecVolumesSecretItems, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecVolumesSecretItems {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecVolumesSecretItems(c *Client, des, nw *ServiceSpecTemplateSpecVolumesSecretItems) *ServiceSpecTemplateSpecVolumesSecretItems {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecVolumesSecretItemsSet(c *Client, des, nw []ServiceSpecTemplateSpecVolumesSecretItems) []ServiceSpecTemplateSpecVolumesSecretItems {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecVolumesSecretItems
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecVolumesSecretItems(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecVolumesSecretItemsSlice(c *Client, des, nw []ServiceSpecTemplateSpecVolumesSecretItems) []ServiceSpecTemplateSpecVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecVolumesSecretItems
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecVolumesSecretItems(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecVolumesConfigMap(des, initial *ServiceSpecTemplateSpecVolumesConfigMap, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecVolumesConfigMap {
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
	if dcl.IsZeroValue(des.Items) {
		des.Items = initial.Items
	}
	if dcl.IsZeroValue(des.DefaultMode) {
		des.DefaultMode = initial.DefaultMode
	}
	if dcl.BoolCanonicalize(des.Optional, initial.Optional) || dcl.IsZeroValue(des.Optional) {
		des.Optional = initial.Optional
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMap(c *Client, des, nw *ServiceSpecTemplateSpecVolumesConfigMap) *ServiceSpecTemplateSpecVolumesConfigMap {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Items = canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c, des.Items, nw.Items)
	if dcl.BoolCanonicalize(des.Optional, nw.Optional) {
		nw.Optional = des.Optional
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapSet(c *Client, des, nw []ServiceSpecTemplateSpecVolumesConfigMap) []ServiceSpecTemplateSpecVolumesConfigMap {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecVolumesConfigMap
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecVolumesConfigMap(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapSlice(c *Client, des, nw []ServiceSpecTemplateSpecVolumesConfigMap) []ServiceSpecTemplateSpecVolumesConfigMap {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecVolumesConfigMap
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecVolumesConfigMap(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTemplateSpecVolumesConfigMapItems(des, initial *ServiceSpecTemplateSpecVolumesConfigMapItems, opts ...dcl.ApplyOption) *ServiceSpecTemplateSpecVolumesConfigMapItems {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}

	return des
}

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapItems(c *Client, des, nw *ServiceSpecTemplateSpecVolumesConfigMapItems) *ServiceSpecTemplateSpecVolumesConfigMapItems {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapItemsSet(c *Client, des, nw []ServiceSpecTemplateSpecVolumesConfigMapItems) []ServiceSpecTemplateSpecVolumesConfigMapItems {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTemplateSpecVolumesConfigMapItems
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTemplateSpecVolumesConfigMapItems(c, &d, &n) {
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

func canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c *Client, des, nw []ServiceSpecTemplateSpecVolumesConfigMapItems) []ServiceSpecTemplateSpecVolumesConfigMapItems {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTemplateSpecVolumesConfigMapItems
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTemplateSpecVolumesConfigMapItems(c, &d, &n))
	}

	return items
}

func canonicalizeServiceSpecTraffic(des, initial *ServiceSpecTraffic, opts ...dcl.ApplyOption) *ServiceSpecTraffic {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ConfigurationName, initial.ConfigurationName) || dcl.IsZeroValue(des.ConfigurationName) {
		des.ConfigurationName = initial.ConfigurationName
	}
	if dcl.StringCanonicalize(des.RevisionName, initial.RevisionName) || dcl.IsZeroValue(des.RevisionName) {
		des.RevisionName = initial.RevisionName
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		des.Tag = initial.Tag
	}
	if dcl.BoolCanonicalize(des.LatestRevision, initial.LatestRevision) || dcl.IsZeroValue(des.LatestRevision) {
		des.LatestRevision = initial.LatestRevision
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewServiceSpecTraffic(c *Client, des, nw *ServiceSpecTraffic) *ServiceSpecTraffic {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ConfigurationName, nw.ConfigurationName) {
		nw.ConfigurationName = des.ConfigurationName
	}
	if dcl.StringCanonicalize(des.RevisionName, nw.RevisionName) {
		nw.RevisionName = des.RevisionName
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.BoolCanonicalize(des.LatestRevision, nw.LatestRevision) {
		nw.LatestRevision = des.LatestRevision
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewServiceSpecTrafficSet(c *Client, des, nw []ServiceSpecTraffic) []ServiceSpecTraffic {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceSpecTraffic
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceSpecTraffic(c, &d, &n) {
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

func canonicalizeNewServiceSpecTrafficSlice(c *Client, des, nw []ServiceSpecTraffic) []ServiceSpecTraffic {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceSpecTraffic
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceSpecTraffic(c, &d, &n))
	}

	return items
}

func canonicalizeServiceStatus(des, initial *ServiceStatus, opts ...dcl.ApplyOption) *ServiceStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObservedGeneration) {
		des.ObservedGeneration = initial.ObservedGeneration
	}
	if dcl.IsZeroValue(des.Conditions) {
		des.Conditions = initial.Conditions
	}
	if dcl.StringCanonicalize(des.LatestReadyRevisionName, initial.LatestReadyRevisionName) || dcl.IsZeroValue(des.LatestReadyRevisionName) {
		des.LatestReadyRevisionName = initial.LatestReadyRevisionName
	}
	if dcl.StringCanonicalize(des.LatestCreatedRevisionName, initial.LatestCreatedRevisionName) || dcl.IsZeroValue(des.LatestCreatedRevisionName) {
		des.LatestCreatedRevisionName = initial.LatestCreatedRevisionName
	}
	if dcl.IsZeroValue(des.Traffic) {
		des.Traffic = initial.Traffic
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}
	des.Address = canonicalizeServiceStatusAddress(des.Address, initial.Address, opts...)

	return des
}

func canonicalizeNewServiceStatus(c *Client, des, nw *ServiceStatus) *ServiceStatus {
	if des == nil || nw == nil {
		return nw
	}

	nw.Conditions = canonicalizeNewServiceStatusConditionsSlice(c, des.Conditions, nw.Conditions)
	if dcl.StringCanonicalize(des.LatestReadyRevisionName, nw.LatestReadyRevisionName) {
		nw.LatestReadyRevisionName = des.LatestReadyRevisionName
	}
	if dcl.StringCanonicalize(des.LatestCreatedRevisionName, nw.LatestCreatedRevisionName) {
		nw.LatestCreatedRevisionName = des.LatestCreatedRevisionName
	}
	nw.Traffic = canonicalizeNewServiceStatusTrafficSlice(c, des.Traffic, nw.Traffic)
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	nw.Address = canonicalizeNewServiceStatusAddress(c, des.Address, nw.Address)

	return nw
}

func canonicalizeNewServiceStatusSet(c *Client, des, nw []ServiceStatus) []ServiceStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceStatus(c, &d, &n) {
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

func canonicalizeNewServiceStatusSlice(c *Client, des, nw []ServiceStatus) []ServiceStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatus(c, &d, &n))
	}

	return items
}

func canonicalizeServiceStatusConditions(des, initial *ServiceStatusConditions, opts ...dcl.ApplyOption) *ServiceStatusConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.StringCanonicalize(des.Status, initial.Status) || dcl.IsZeroValue(des.Status) {
		des.Status = initial.Status
	}
	if dcl.StringCanonicalize(des.Reason, initial.Reason) || dcl.IsZeroValue(des.Reason) {
		des.Reason = initial.Reason
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	des.LastTransitionTime = canonicalizeServiceStatusConditionsLastTransitionTime(des.LastTransitionTime, initial.LastTransitionTime, opts...)
	if dcl.StringCanonicalize(des.Severity, initial.Severity) || dcl.IsZeroValue(des.Severity) {
		des.Severity = initial.Severity
	}

	return des
}

func canonicalizeNewServiceStatusConditions(c *Client, des, nw *ServiceStatusConditions) *ServiceStatusConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Status, nw.Status) {
		nw.Status = des.Status
	}
	if dcl.StringCanonicalize(des.Reason, nw.Reason) {
		nw.Reason = des.Reason
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.LastTransitionTime = canonicalizeNewServiceStatusConditionsLastTransitionTime(c, des.LastTransitionTime, nw.LastTransitionTime)
	if dcl.StringCanonicalize(des.Severity, nw.Severity) {
		nw.Severity = des.Severity
	}

	return nw
}

func canonicalizeNewServiceStatusConditionsSet(c *Client, des, nw []ServiceStatusConditions) []ServiceStatusConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceStatusConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceStatusConditions(c, &d, &n) {
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

func canonicalizeNewServiceStatusConditionsSlice(c *Client, des, nw []ServiceStatusConditions) []ServiceStatusConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceStatusConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatusConditions(c, &d, &n))
	}

	return items
}

func canonicalizeServiceStatusConditionsLastTransitionTime(des, initial *ServiceStatusConditionsLastTransitionTime, opts ...dcl.ApplyOption) *ServiceStatusConditionsLastTransitionTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewServiceStatusConditionsLastTransitionTime(c *Client, des, nw *ServiceStatusConditionsLastTransitionTime) *ServiceStatusConditionsLastTransitionTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewServiceStatusConditionsLastTransitionTimeSet(c *Client, des, nw []ServiceStatusConditionsLastTransitionTime) []ServiceStatusConditionsLastTransitionTime {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceStatusConditionsLastTransitionTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceStatusConditionsLastTransitionTime(c, &d, &n) {
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

func canonicalizeNewServiceStatusConditionsLastTransitionTimeSlice(c *Client, des, nw []ServiceStatusConditionsLastTransitionTime) []ServiceStatusConditionsLastTransitionTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceStatusConditionsLastTransitionTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatusConditionsLastTransitionTime(c, &d, &n))
	}

	return items
}

func canonicalizeServiceStatusTraffic(des, initial *ServiceStatusTraffic, opts ...dcl.ApplyOption) *ServiceStatusTraffic {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ConfigurationName, initial.ConfigurationName) || dcl.IsZeroValue(des.ConfigurationName) {
		des.ConfigurationName = initial.ConfigurationName
	}
	if dcl.StringCanonicalize(des.RevisionName, initial.RevisionName) || dcl.IsZeroValue(des.RevisionName) {
		des.RevisionName = initial.RevisionName
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		des.Tag = initial.Tag
	}
	if dcl.BoolCanonicalize(des.LatestRevision, initial.LatestRevision) || dcl.IsZeroValue(des.LatestRevision) {
		des.LatestRevision = initial.LatestRevision
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewServiceStatusTraffic(c *Client, des, nw *ServiceStatusTraffic) *ServiceStatusTraffic {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ConfigurationName, nw.ConfigurationName) {
		nw.ConfigurationName = des.ConfigurationName
	}
	if dcl.StringCanonicalize(des.RevisionName, nw.RevisionName) {
		nw.RevisionName = des.RevisionName
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.BoolCanonicalize(des.LatestRevision, nw.LatestRevision) {
		nw.LatestRevision = des.LatestRevision
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewServiceStatusTrafficSet(c *Client, des, nw []ServiceStatusTraffic) []ServiceStatusTraffic {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceStatusTraffic
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceStatusTraffic(c, &d, &n) {
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

func canonicalizeNewServiceStatusTrafficSlice(c *Client, des, nw []ServiceStatusTraffic) []ServiceStatusTraffic {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceStatusTraffic
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatusTraffic(c, &d, &n))
	}

	return items
}

func canonicalizeServiceStatusAddress(des, initial *ServiceStatusAddress, opts ...dcl.ApplyOption) *ServiceStatusAddress {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewServiceStatusAddress(c *Client, des, nw *ServiceStatusAddress) *ServiceStatusAddress {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewServiceStatusAddressSet(c *Client, des, nw []ServiceStatusAddress) []ServiceStatusAddress {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceStatusAddress
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareServiceStatusAddress(c, &d, &n) {
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

func canonicalizeNewServiceStatusAddressSlice(c *Client, des, nw []ServiceStatusAddress) []ServiceStatusAddress {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ServiceStatusAddress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatusAddress(c, &d, &n))
	}

	return items
}

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffService(c *Client, desired, actual *Service, opts ...dcl.ApplyOption) ([]serviceDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []serviceDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ApiVersion",
		})
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Kind",
		})
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{ObjectFunction: compareServiceMetadataNewStyle}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Metadata",
		})
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.Info{ObjectFunction: compareServiceSpecNewStyle}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Spec",
		})
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceStatusNewStyle}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Status",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
		})
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
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
	var deduped []serviceDiff
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
func compareServiceMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceMetadata)
	if !ok {
		desiredNotPointer, ok := d.(ServiceMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadata or *ServiceMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceMetadata)
	if !ok {
		actualNotPointer, ok := a.(ServiceMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GenerateName, actual.GenerateName, dcl.Info{}, fn.AddNest("GenerateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceVersion, actual.ResourceVersion, dcl.Info{}, fn.AddNest("ResourceVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{ObjectFunction: compareServiceMetadataCreateTimeNewStyle}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OwnerReferences, actual.OwnerReferences, dcl.Info{ObjectFunction: compareServiceMetadataOwnerReferencesNewStyle}, fn.AddNest("OwnerReferences")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{ObjectFunction: compareServiceMetadataDeleteTimeNewStyle}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds, dcl.Info{}, fn.AddNest("DeletionGracePeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Finalizers, actual.Finalizers, dcl.Info{}, fn.AddNest("Finalizers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceMetadata(c *Client, desired, actual *ServiceMetadata) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.GenerateName, actual.GenerateName) && !dcl.IsZeroValue(desired.GenerateName) {
		c.Config.Logger.Infof("Diff in GenerateName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GenerateName), dcl.SprintResource(actual.GenerateName))
		return true
	}
	if !dcl.StringCanonicalize(desired.Namespace, actual.Namespace) && !dcl.IsZeroValue(desired.Namespace) {
		c.Config.Logger.Infof("Diff in Namespace.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Namespace), dcl.SprintResource(actual.Namespace))
		return true
	}
	if !dcl.StringCanonicalize(desired.SelfLink, actual.SelfLink) && !dcl.IsZeroValue(desired.SelfLink) {
		c.Config.Logger.Infof("Diff in SelfLink.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SelfLink), dcl.SprintResource(actual.SelfLink))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uid, actual.Uid) && !dcl.IsZeroValue(desired.Uid) {
		c.Config.Logger.Infof("Diff in Uid.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uid), dcl.SprintResource(actual.Uid))
		return true
	}
	if !dcl.StringCanonicalize(desired.ResourceVersion, actual.ResourceVersion) && !dcl.IsZeroValue(desired.ResourceVersion) {
		c.Config.Logger.Infof("Diff in ResourceVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceVersion), dcl.SprintResource(actual.ResourceVersion))
		return true
	}
	if !reflect.DeepEqual(desired.Generation, actual.Generation) && !dcl.IsZeroValue(desired.Generation) {
		c.Config.Logger.Infof("Diff in Generation.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Generation), dcl.SprintResource(actual.Generation))
		return true
	}
	if compareServiceMetadataCreateTime(c, desired.CreateTime, actual.CreateTime) && !dcl.IsZeroValue(desired.CreateTime) {
		c.Config.Logger.Infof("Diff in CreateTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateTime), dcl.SprintResource(actual.CreateTime))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Annotations, actual.Annotations, []string(nil)) && !dcl.IsZeroValue(desired.Annotations) {
		c.Config.Logger.Infof("Diff in Annotations.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Annotations), dcl.SprintResource(actual.Annotations))
		return true
	}
	if compareServiceMetadataOwnerReferencesSlice(c, desired.OwnerReferences, actual.OwnerReferences) && !dcl.IsZeroValue(desired.OwnerReferences) {
		c.Config.Logger.Infof("Diff in OwnerReferences.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OwnerReferences), dcl.SprintResource(actual.OwnerReferences))
		return true
	}
	if compareServiceMetadataDeleteTime(c, desired.DeleteTime, actual.DeleteTime) && !dcl.IsZeroValue(desired.DeleteTime) {
		c.Config.Logger.Infof("Diff in DeleteTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeleteTime), dcl.SprintResource(actual.DeleteTime))
		return true
	}
	if !reflect.DeepEqual(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds) && !dcl.IsZeroValue(desired.DeletionGracePeriodSeconds) {
		c.Config.Logger.Infof("Diff in DeletionGracePeriodSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeletionGracePeriodSeconds), dcl.SprintResource(actual.DeletionGracePeriodSeconds))
		return true
	}
	if !dcl.StringSliceEquals(desired.Finalizers, actual.Finalizers) && !dcl.IsZeroValue(desired.Finalizers) {
		c.Config.Logger.Infof("Diff in Finalizers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Finalizers), dcl.SprintResource(actual.Finalizers))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterName, actual.ClusterName) && !dcl.IsZeroValue(desired.ClusterName) {
		c.Config.Logger.Infof("Diff in ClusterName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterName), dcl.SprintResource(actual.ClusterName))
		return true
	}
	return false
}

func compareServiceMetadataSlice(c *Client, desired, actual []ServiceMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadata, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceMetadata(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceMetadata, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceMetadataMap(c *Client, desired, actual map[string]ServiceMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadata, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceMetadata, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceMetadata(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceMetadata, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceMetadataCreateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceMetadataCreateTime)
	if !ok {
		desiredNotPointer, ok := d.(ServiceMetadataCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataCreateTime or *ServiceMetadataCreateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceMetadataCreateTime)
	if !ok {
		actualNotPointer, ok := a.(ServiceMetadataCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataCreateTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceMetadataCreateTime(c *Client, desired, actual *ServiceMetadataCreateTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareServiceMetadataCreateTimeSlice(c *Client, desired, actual []ServiceMetadataCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataCreateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceMetadataCreateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceMetadataCreateTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceMetadataCreateTimeMap(c *Client, desired, actual map[string]ServiceMetadataCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataCreateTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceMetadataCreateTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceMetadataCreateTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceMetadataCreateTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceMetadataOwnerReferencesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceMetadataOwnerReferences)
	if !ok {
		desiredNotPointer, ok := d.(ServiceMetadataOwnerReferences)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataOwnerReferences or *ServiceMetadataOwnerReferences", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceMetadataOwnerReferences)
	if !ok {
		actualNotPointer, ok := a.(ServiceMetadataOwnerReferences)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataOwnerReferences", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Controller, actual.Controller, dcl.Info{}, fn.AddNest("Controller")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion, dcl.Info{}, fn.AddNest("BlockOwnerDeletion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceMetadataOwnerReferences(c *Client, desired, actual *ServiceMetadataOwnerReferences) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.ApiVersion, actual.ApiVersion) && !dcl.IsZeroValue(desired.ApiVersion) {
		c.Config.Logger.Infof("Diff in ApiVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ApiVersion), dcl.SprintResource(actual.ApiVersion))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uid, actual.Uid) && !dcl.IsZeroValue(desired.Uid) {
		c.Config.Logger.Infof("Diff in Uid.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uid), dcl.SprintResource(actual.Uid))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Controller, actual.Controller) && !dcl.IsZeroValue(desired.Controller) {
		c.Config.Logger.Infof("Diff in Controller.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Controller), dcl.SprintResource(actual.Controller))
		return true
	}
	if !dcl.BoolCanonicalize(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion) && !dcl.IsZeroValue(desired.BlockOwnerDeletion) {
		c.Config.Logger.Infof("Diff in BlockOwnerDeletion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BlockOwnerDeletion), dcl.SprintResource(actual.BlockOwnerDeletion))
		return true
	}
	return false
}

func compareServiceMetadataOwnerReferencesSlice(c *Client, desired, actual []ServiceMetadataOwnerReferences) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataOwnerReferences, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceMetadataOwnerReferences(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceMetadataOwnerReferences, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceMetadataOwnerReferencesMap(c *Client, desired, actual map[string]ServiceMetadataOwnerReferences) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataOwnerReferences, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceMetadataOwnerReferences, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceMetadataOwnerReferences(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceMetadataOwnerReferences, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceMetadataDeleteTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceMetadataDeleteTime)
	if !ok {
		desiredNotPointer, ok := d.(ServiceMetadataDeleteTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataDeleteTime or *ServiceMetadataDeleteTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceMetadataDeleteTime)
	if !ok {
		actualNotPointer, ok := a.(ServiceMetadataDeleteTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceMetadataDeleteTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceMetadataDeleteTime(c *Client, desired, actual *ServiceMetadataDeleteTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareServiceMetadataDeleteTimeSlice(c *Client, desired, actual []ServiceMetadataDeleteTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataDeleteTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceMetadataDeleteTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceMetadataDeleteTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceMetadataDeleteTimeMap(c *Client, desired, actual map[string]ServiceMetadataDeleteTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceMetadataDeleteTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceMetadataDeleteTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceMetadataDeleteTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceMetadataDeleteTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpec)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpec or *ServiceSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpec)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.Info{ObjectFunction: compareServiceSpecTemplateNewStyle}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.Info{ObjectFunction: compareServiceSpecTrafficNewStyle}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpec(c *Client, desired, actual *ServiceSpec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplate(c, desired.Template, actual.Template) && !dcl.IsZeroValue(desired.Template) {
		c.Config.Logger.Infof("Diff in Template.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Template), dcl.SprintResource(actual.Template))
		return true
	}
	if compareServiceSpecTrafficSlice(c, desired.Traffic, actual.Traffic) && !dcl.IsZeroValue(desired.Traffic) {
		c.Config.Logger.Infof("Diff in Traffic.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Traffic), dcl.SprintResource(actual.Traffic))
		return true
	}
	return false
}

func compareServiceSpecSlice(c *Client, desired, actual []ServiceSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpec, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecMap(c *Client, desired, actual map[string]ServiceSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpec, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplate or *ServiceSpecTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplate)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataNewStyle}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecNewStyle}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplate(c *Client, desired, actual *ServiceSpecTemplate) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateMetadata(c, desired.Metadata, actual.Metadata) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if compareServiceSpecTemplateSpec(c, desired.Spec, actual.Spec) && !dcl.IsZeroValue(desired.Spec) {
		c.Config.Logger.Infof("Diff in Spec.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Spec), dcl.SprintResource(actual.Spec))
		return true
	}
	return false
}

func compareServiceSpecTemplateSlice(c *Client, desired, actual []ServiceSpecTemplate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplate, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplate(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplate, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMap(c *Client, desired, actual map[string]ServiceSpecTemplate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplate, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplate, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplate(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplate, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateMetadata)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadata or *ServiceSpecTemplateMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateMetadata)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GenerateName, actual.GenerateName, dcl.Info{}, fn.AddNest("GenerateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceVersion, actual.ResourceVersion, dcl.Info{}, fn.AddNest("ResourceVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataCreateTimeNewStyle}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OwnerReferences, actual.OwnerReferences, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataOwnerReferencesNewStyle}, fn.AddNest("OwnerReferences")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataDeleteTimeNewStyle}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds, dcl.Info{}, fn.AddNest("DeletionGracePeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Finalizers, actual.Finalizers, dcl.Info{}, fn.AddNest("Finalizers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateMetadata(c *Client, desired, actual *ServiceSpecTemplateMetadata) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.GenerateName, actual.GenerateName) && !dcl.IsZeroValue(desired.GenerateName) {
		c.Config.Logger.Infof("Diff in GenerateName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GenerateName), dcl.SprintResource(actual.GenerateName))
		return true
	}
	if !dcl.StringCanonicalize(desired.Namespace, actual.Namespace) && !dcl.IsZeroValue(desired.Namespace) {
		c.Config.Logger.Infof("Diff in Namespace.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Namespace), dcl.SprintResource(actual.Namespace))
		return true
	}
	if !dcl.StringCanonicalize(desired.SelfLink, actual.SelfLink) && !dcl.IsZeroValue(desired.SelfLink) {
		c.Config.Logger.Infof("Diff in SelfLink.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SelfLink), dcl.SprintResource(actual.SelfLink))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uid, actual.Uid) && !dcl.IsZeroValue(desired.Uid) {
		c.Config.Logger.Infof("Diff in Uid.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uid), dcl.SprintResource(actual.Uid))
		return true
	}
	if !dcl.StringCanonicalize(desired.ResourceVersion, actual.ResourceVersion) && !dcl.IsZeroValue(desired.ResourceVersion) {
		c.Config.Logger.Infof("Diff in ResourceVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceVersion), dcl.SprintResource(actual.ResourceVersion))
		return true
	}
	if !reflect.DeepEqual(desired.Generation, actual.Generation) && !dcl.IsZeroValue(desired.Generation) {
		c.Config.Logger.Infof("Diff in Generation.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Generation), dcl.SprintResource(actual.Generation))
		return true
	}
	if compareServiceSpecTemplateMetadataCreateTime(c, desired.CreateTime, actual.CreateTime) && !dcl.IsZeroValue(desired.CreateTime) {
		c.Config.Logger.Infof("Diff in CreateTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateTime), dcl.SprintResource(actual.CreateTime))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Annotations, actual.Annotations, []string(nil)) && !dcl.IsZeroValue(desired.Annotations) {
		c.Config.Logger.Infof("Diff in Annotations.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Annotations), dcl.SprintResource(actual.Annotations))
		return true
	}
	if compareServiceSpecTemplateMetadataOwnerReferencesSlice(c, desired.OwnerReferences, actual.OwnerReferences) && !dcl.IsZeroValue(desired.OwnerReferences) {
		c.Config.Logger.Infof("Diff in OwnerReferences.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OwnerReferences), dcl.SprintResource(actual.OwnerReferences))
		return true
	}
	if compareServiceSpecTemplateMetadataDeleteTime(c, desired.DeleteTime, actual.DeleteTime) && !dcl.IsZeroValue(desired.DeleteTime) {
		c.Config.Logger.Infof("Diff in DeleteTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeleteTime), dcl.SprintResource(actual.DeleteTime))
		return true
	}
	if !reflect.DeepEqual(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds) && !dcl.IsZeroValue(desired.DeletionGracePeriodSeconds) {
		c.Config.Logger.Infof("Diff in DeletionGracePeriodSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeletionGracePeriodSeconds), dcl.SprintResource(actual.DeletionGracePeriodSeconds))
		return true
	}
	if !dcl.StringSliceEquals(desired.Finalizers, actual.Finalizers) && !dcl.IsZeroValue(desired.Finalizers) {
		c.Config.Logger.Infof("Diff in Finalizers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Finalizers), dcl.SprintResource(actual.Finalizers))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterName, actual.ClusterName) && !dcl.IsZeroValue(desired.ClusterName) {
		c.Config.Logger.Infof("Diff in ClusterName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterName), dcl.SprintResource(actual.ClusterName))
		return true
	}
	return false
}

func compareServiceSpecTemplateMetadataSlice(c *Client, desired, actual []ServiceSpecTemplateMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadata, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateMetadata(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadata, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataMap(c *Client, desired, actual map[string]ServiceSpecTemplateMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadata, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadata, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateMetadata(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadata, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataCreateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateMetadataCreateTime)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateMetadataCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataCreateTime or *ServiceSpecTemplateMetadataCreateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateMetadataCreateTime)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateMetadataCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataCreateTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateMetadataCreateTime(c *Client, desired, actual *ServiceSpecTemplateMetadataCreateTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareServiceSpecTemplateMetadataCreateTimeSlice(c *Client, desired, actual []ServiceSpecTemplateMetadataCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataCreateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateMetadataCreateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataCreateTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataCreateTimeMap(c *Client, desired, actual map[string]ServiceSpecTemplateMetadataCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataCreateTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataCreateTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateMetadataCreateTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataCreateTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataOwnerReferencesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateMetadataOwnerReferences)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateMetadataOwnerReferences)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataOwnerReferences or *ServiceSpecTemplateMetadataOwnerReferences", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateMetadataOwnerReferences)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateMetadataOwnerReferences)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataOwnerReferences", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Controller, actual.Controller, dcl.Info{}, fn.AddNest("Controller")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion, dcl.Info{}, fn.AddNest("BlockOwnerDeletion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateMetadataOwnerReferences(c *Client, desired, actual *ServiceSpecTemplateMetadataOwnerReferences) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.ApiVersion, actual.ApiVersion) && !dcl.IsZeroValue(desired.ApiVersion) {
		c.Config.Logger.Infof("Diff in ApiVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ApiVersion), dcl.SprintResource(actual.ApiVersion))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Uid, actual.Uid) && !dcl.IsZeroValue(desired.Uid) {
		c.Config.Logger.Infof("Diff in Uid.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uid), dcl.SprintResource(actual.Uid))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Controller, actual.Controller) && !dcl.IsZeroValue(desired.Controller) {
		c.Config.Logger.Infof("Diff in Controller.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Controller), dcl.SprintResource(actual.Controller))
		return true
	}
	if !dcl.BoolCanonicalize(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion) && !dcl.IsZeroValue(desired.BlockOwnerDeletion) {
		c.Config.Logger.Infof("Diff in BlockOwnerDeletion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BlockOwnerDeletion), dcl.SprintResource(actual.BlockOwnerDeletion))
		return true
	}
	return false
}

func compareServiceSpecTemplateMetadataOwnerReferencesSlice(c *Client, desired, actual []ServiceSpecTemplateMetadataOwnerReferences) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataOwnerReferences, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateMetadataOwnerReferences(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataOwnerReferences, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataOwnerReferencesMap(c *Client, desired, actual map[string]ServiceSpecTemplateMetadataOwnerReferences) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataOwnerReferences, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataOwnerReferences, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateMetadataOwnerReferences(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataOwnerReferences, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataDeleteTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateMetadataDeleteTime)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateMetadataDeleteTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataDeleteTime or *ServiceSpecTemplateMetadataDeleteTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateMetadataDeleteTime)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateMetadataDeleteTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateMetadataDeleteTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateMetadataDeleteTime(c *Client, desired, actual *ServiceSpecTemplateMetadataDeleteTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareServiceSpecTemplateMetadataDeleteTimeSlice(c *Client, desired, actual []ServiceSpecTemplateMetadataDeleteTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataDeleteTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateMetadataDeleteTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataDeleteTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateMetadataDeleteTimeMap(c *Client, desired, actual map[string]ServiceSpecTemplateMetadataDeleteTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateMetadataDeleteTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataDeleteTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateMetadataDeleteTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateMetadataDeleteTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpec)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpec or *ServiceSpecTemplateSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpec)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ContainerConcurrency, actual.ContainerConcurrency, dcl.Info{}, fn.AddNest("ContainerConcurrency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccountName, actual.ServiceAccountName, dcl.Info{}, fn.AddNest("ServiceAccountName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Containers, actual.Containers, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersNewStyle}, fn.AddNest("Containers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesNewStyle}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpec(c *Client, desired, actual *ServiceSpecTemplateSpec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.ContainerConcurrency, actual.ContainerConcurrency) && !dcl.IsZeroValue(desired.ContainerConcurrency) {
		c.Config.Logger.Infof("Diff in ContainerConcurrency.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContainerConcurrency), dcl.SprintResource(actual.ContainerConcurrency))
		return true
	}
	if !reflect.DeepEqual(desired.TimeoutSeconds, actual.TimeoutSeconds) && !dcl.IsZeroValue(desired.TimeoutSeconds) {
		c.Config.Logger.Infof("Diff in TimeoutSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeoutSeconds), dcl.SprintResource(actual.TimeoutSeconds))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccountName, actual.ServiceAccountName) && !dcl.IsZeroValue(desired.ServiceAccountName) {
		c.Config.Logger.Infof("Diff in ServiceAccountName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountName), dcl.SprintResource(actual.ServiceAccountName))
		return true
	}
	if compareServiceSpecTemplateSpecContainersSlice(c, desired.Containers, actual.Containers) && !dcl.IsZeroValue(desired.Containers) {
		c.Config.Logger.Infof("Diff in Containers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Containers), dcl.SprintResource(actual.Containers))
		return true
	}
	if compareServiceSpecTemplateSpecVolumesSlice(c, desired.Volumes, actual.Volumes) && !dcl.IsZeroValue(desired.Volumes) {
		c.Config.Logger.Infof("Diff in Volumes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Volumes), dcl.SprintResource(actual.Volumes))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecSlice(c *Client, desired, actual []ServiceSpecTemplateSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpec, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpec, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainers)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainers or *ServiceSpecTemplateSpecContainers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainers)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvNewStyle}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersResourcesNewStyle}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkingDir, actual.WorkingDir, dcl.Info{}, fn.AddNest("WorkingDir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersPortsNewStyle}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnvFrom, actual.EnvFrom, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromNewStyle}, fn.AddNest("EnvFrom")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeMounts, actual.VolumeMounts, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersVolumeMountsNewStyle}, fn.AddNest("VolumeMounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LivenessProbe, actual.LivenessProbe, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeNewStyle}, fn.AddNest("LivenessProbe")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadinessProbe, actual.ReadinessProbe, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeNewStyle}, fn.AddNest("ReadinessProbe")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminationMessagePath, actual.TerminationMessagePath, dcl.Info{}, fn.AddNest("TerminationMessagePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminationMessagePolicy, actual.TerminationMessagePolicy, dcl.Info{}, fn.AddNest("TerminationMessagePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImagePullPolicy, actual.ImagePullPolicy, dcl.Info{}, fn.AddNest("ImagePullPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityContext, actual.SecurityContext, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersSecurityContextNewStyle}, fn.AddNest("SecurityContext")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainers(c *Client, desired, actual *ServiceSpecTemplateSpecContainers) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Image, actual.Image) && !dcl.IsZeroValue(desired.Image) {
		c.Config.Logger.Infof("Diff in Image.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Image), dcl.SprintResource(actual.Image))
		return true
	}
	if !dcl.StringSliceEquals(desired.Command, actual.Command) && !dcl.IsZeroValue(desired.Command) {
		c.Config.Logger.Infof("Diff in Command.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Command), dcl.SprintResource(actual.Command))
		return true
	}
	if !dcl.StringSliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvSlice(c, desired.Env, actual.Env) && !dcl.IsZeroValue(desired.Env) {
		c.Config.Logger.Infof("Diff in Env.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Env), dcl.SprintResource(actual.Env))
		return true
	}
	if compareServiceSpecTemplateSpecContainersResources(c, desired.Resources, actual.Resources) && !dcl.IsZeroValue(desired.Resources) {
		c.Config.Logger.Infof("Diff in Resources.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Resources), dcl.SprintResource(actual.Resources))
		return true
	}
	if !dcl.StringCanonicalize(desired.WorkingDir, actual.WorkingDir) && !dcl.IsZeroValue(desired.WorkingDir) {
		c.Config.Logger.Infof("Diff in WorkingDir.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkingDir), dcl.SprintResource(actual.WorkingDir))
		return true
	}
	if compareServiceSpecTemplateSpecContainersPortsSlice(c, desired.Ports, actual.Ports) && !dcl.IsZeroValue(desired.Ports) {
		c.Config.Logger.Infof("Diff in Ports.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Ports), dcl.SprintResource(actual.Ports))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvFromSlice(c, desired.EnvFrom, actual.EnvFrom) && !dcl.IsZeroValue(desired.EnvFrom) {
		c.Config.Logger.Infof("Diff in EnvFrom.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnvFrom), dcl.SprintResource(actual.EnvFrom))
		return true
	}
	if compareServiceSpecTemplateSpecContainersVolumeMountsSlice(c, desired.VolumeMounts, actual.VolumeMounts) && !dcl.IsZeroValue(desired.VolumeMounts) {
		c.Config.Logger.Infof("Diff in VolumeMounts.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VolumeMounts), dcl.SprintResource(actual.VolumeMounts))
		return true
	}
	if compareServiceSpecTemplateSpecContainersLivenessProbe(c, desired.LivenessProbe, actual.LivenessProbe) && !dcl.IsZeroValue(desired.LivenessProbe) {
		c.Config.Logger.Infof("Diff in LivenessProbe.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LivenessProbe), dcl.SprintResource(actual.LivenessProbe))
		return true
	}
	if compareServiceSpecTemplateSpecContainersReadinessProbe(c, desired.ReadinessProbe, actual.ReadinessProbe) && !dcl.IsZeroValue(desired.ReadinessProbe) {
		c.Config.Logger.Infof("Diff in ReadinessProbe.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReadinessProbe), dcl.SprintResource(actual.ReadinessProbe))
		return true
	}
	if !dcl.StringCanonicalize(desired.TerminationMessagePath, actual.TerminationMessagePath) && !dcl.IsZeroValue(desired.TerminationMessagePath) {
		c.Config.Logger.Infof("Diff in TerminationMessagePath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TerminationMessagePath), dcl.SprintResource(actual.TerminationMessagePath))
		return true
	}
	if !dcl.StringCanonicalize(desired.TerminationMessagePolicy, actual.TerminationMessagePolicy) && !dcl.IsZeroValue(desired.TerminationMessagePolicy) {
		c.Config.Logger.Infof("Diff in TerminationMessagePolicy.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TerminationMessagePolicy), dcl.SprintResource(actual.TerminationMessagePolicy))
		return true
	}
	if !dcl.StringCanonicalize(desired.ImagePullPolicy, actual.ImagePullPolicy) && !dcl.IsZeroValue(desired.ImagePullPolicy) {
		c.Config.Logger.Infof("Diff in ImagePullPolicy.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImagePullPolicy), dcl.SprintResource(actual.ImagePullPolicy))
		return true
	}
	if compareServiceSpecTemplateSpecContainersSecurityContext(c, desired.SecurityContext, actual.SecurityContext) && !dcl.IsZeroValue(desired.SecurityContext) {
		c.Config.Logger.Infof("Diff in SecurityContext.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecurityContext), dcl.SprintResource(actual.SecurityContext))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainers, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainers(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainers, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainers, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainers, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainers(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainers, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnv)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnv or *ServiceSpecTemplateSpecContainersEnv", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnv)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnv", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueFrom, actual.ValueFrom, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromNewStyle}, fn.AddNest("ValueFrom")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnv(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnv) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvValueFrom(c, desired.ValueFrom, actual.ValueFrom) && !dcl.IsZeroValue(desired.ValueFrom) {
		c.Config.Logger.Infof("Diff in ValueFrom.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ValueFrom), dcl.SprintResource(actual.ValueFrom))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnv) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnv, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnv(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnv, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnv) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnv, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnv, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnv(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnv, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvValueFrom)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvValueFrom)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFrom or *ServiceSpecTemplateSpecContainersEnvValueFrom", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvValueFrom)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvValueFrom)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFrom", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigMapKeyRef, actual.ConfigMapKeyRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefNewStyle}, fn.AddNest("ConfigMapKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretKeyRef, actual.SecretKeyRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefNewStyle}, fn.AddNest("SecretKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvValueFrom(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvValueFrom) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, desired.ConfigMapKeyRef, actual.ConfigMapKeyRef) && !dcl.IsZeroValue(desired.ConfigMapKeyRef) {
		c.Config.Logger.Infof("Diff in ConfigMapKeyRef.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigMapKeyRef), dcl.SprintResource(actual.ConfigMapKeyRef))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, desired.SecretKeyRef, actual.SecretKeyRef) && !dcl.IsZeroValue(desired.SecretKeyRef) {
		c.Config.Logger.Infof("Diff in SecretKeyRef.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecretKeyRef), dcl.SprintResource(actual.SecretKeyRef))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvValueFrom) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFrom, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvValueFrom(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFrom, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvValueFrom) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFrom, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFrom, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvValueFrom(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFrom, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef or *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceNewStyle}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, desired.LocalObjectReference, actual.LocalObjectReference) && !dcl.IsZeroValue(desired.LocalObjectReference) {
		c.Config.Logger.Infof("Diff in LocalObjectReference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalObjectReference), dcl.SprintResource(actual.LocalObjectReference))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference or *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef or *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceNewStyle}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, desired.LocalObjectReference, actual.LocalObjectReference) && !dcl.IsZeroValue(desired.LocalObjectReference) {
		c.Config.Logger.Infof("Diff in LocalObjectReference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalObjectReference), dcl.SprintResource(actual.LocalObjectReference))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference or *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersResources or *ServiceSpecTemplateSpecContainersResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.Info{}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Requests, actual.Requests, dcl.Info{}, fn.AddNest("Requests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersResources(c *Client, desired, actual *ServiceSpecTemplateSpecContainersResources) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.MapEquals(desired.Limits, actual.Limits, []string(nil)) && !dcl.IsZeroValue(desired.Limits) {
		c.Config.Logger.Infof("Diff in Limits.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Limits), dcl.SprintResource(actual.Limits))
		return true
	}
	if !dcl.MapEquals(desired.Requests, actual.Requests, []string(nil)) && !dcl.IsZeroValue(desired.Requests) {
		c.Config.Logger.Infof("Diff in Requests.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Requests), dcl.SprintResource(actual.Requests))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersResourcesSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersResources) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersResources, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersResources(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersResources, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersResourcesMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersResources) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersResources, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersResources, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersResources(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersResources, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersPorts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersPorts or *ServiceSpecTemplateSpecContainersPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersPorts)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.Info{}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Protocol, actual.Protocol, dcl.Info{}, fn.AddNest("Protocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersPorts(c *Client, desired, actual *ServiceSpecTemplateSpecContainersPorts) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !reflect.DeepEqual(desired.ContainerPort, actual.ContainerPort) && !dcl.IsZeroValue(desired.ContainerPort) {
		c.Config.Logger.Infof("Diff in ContainerPort.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContainerPort), dcl.SprintResource(actual.ContainerPort))
		return true
	}
	if !dcl.StringCanonicalize(desired.Protocol, actual.Protocol) && !dcl.IsZeroValue(desired.Protocol) {
		c.Config.Logger.Infof("Diff in Protocol.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Protocol), dcl.SprintResource(actual.Protocol))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersPortsSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersPorts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersPorts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersPorts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersPorts, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersPortsMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersPorts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersPorts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersPorts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersPorts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersPorts, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvFrom)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvFrom)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFrom or *ServiceSpecTemplateSpecContainersEnvFrom", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvFrom)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvFrom)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFrom", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Prefix, actual.Prefix, dcl.Info{}, fn.AddNest("Prefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigMapRef, actual.ConfigMapRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefNewStyle}, fn.AddNest("ConfigMapRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretRef, actual.SecretRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromSecretRefNewStyle}, fn.AddNest("SecretRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvFrom(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvFrom) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Prefix, actual.Prefix) && !dcl.IsZeroValue(desired.Prefix) {
		c.Config.Logger.Infof("Diff in Prefix.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Prefix), dcl.SprintResource(actual.Prefix))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, desired.ConfigMapRef, actual.ConfigMapRef) && !dcl.IsZeroValue(desired.ConfigMapRef) {
		c.Config.Logger.Infof("Diff in ConfigMapRef.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigMapRef), dcl.SprintResource(actual.ConfigMapRef))
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvFromSecretRef(c, desired.SecretRef, actual.SecretRef) && !dcl.IsZeroValue(desired.SecretRef) {
		c.Config.Logger.Infof("Diff in SecretRef.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecretRef), dcl.SprintResource(actual.SecretRef))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvFrom) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFrom, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvFrom(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFrom, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvFrom) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFrom, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFrom, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvFrom(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFrom, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvFromConfigMapRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvFromConfigMapRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromConfigMapRef or *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvFromConfigMapRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvFromConfigMapRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromConfigMapRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceNewStyle}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, desired.LocalObjectReference, actual.LocalObjectReference) && !dcl.IsZeroValue(desired.LocalObjectReference) {
		c.Config.Logger.Infof("Diff in LocalObjectReference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalObjectReference), dcl.SprintResource(actual.LocalObjectReference))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference or *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvFromSecretRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvFromSecretRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromSecretRef or *ServiceSpecTemplateSpecContainersEnvFromSecretRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvFromSecretRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvFromSecretRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromSecretRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceNewStyle}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRef(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvFromSecretRef) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, desired.LocalObjectReference, actual.LocalObjectReference) && !dcl.IsZeroValue(desired.LocalObjectReference) {
		c.Config.Logger.Infof("Diff in LocalObjectReference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalObjectReference), dcl.SprintResource(actual.LocalObjectReference))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvFromSecretRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRef, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRef, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRef, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRef, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRef, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference or *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c *Client, desired, actual *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersVolumeMountsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersVolumeMounts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersVolumeMounts or *ServiceSpecTemplateSpecContainersVolumeMounts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersVolumeMounts)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersVolumeMounts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadOnly, actual.ReadOnly, dcl.Info{}, fn.AddNest("ReadOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MountPath, actual.MountPath, dcl.Info{}, fn.AddNest("MountPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubPath, actual.SubPath, dcl.Info{}, fn.AddNest("SubPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersVolumeMounts(c *Client, desired, actual *ServiceSpecTemplateSpecContainersVolumeMounts) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ReadOnly, actual.ReadOnly) && !dcl.IsZeroValue(desired.ReadOnly) {
		c.Config.Logger.Infof("Diff in ReadOnly.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReadOnly), dcl.SprintResource(actual.ReadOnly))
		return true
	}
	if !dcl.StringCanonicalize(desired.MountPath, actual.MountPath) && !dcl.IsZeroValue(desired.MountPath) {
		c.Config.Logger.Infof("Diff in MountPath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MountPath), dcl.SprintResource(actual.MountPath))
		return true
	}
	if !dcl.StringCanonicalize(desired.SubPath, actual.SubPath) && !dcl.IsZeroValue(desired.SubPath) {
		c.Config.Logger.Infof("Diff in SubPath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubPath), dcl.SprintResource(actual.SubPath))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersVolumeMountsSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersVolumeMounts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersVolumeMounts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersVolumeMounts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersVolumeMounts, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersVolumeMountsMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersVolumeMounts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersVolumeMounts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersVolumeMounts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersVolumeMounts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersVolumeMounts, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersLivenessProbe)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersLivenessProbe)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbe or *ServiceSpecTemplateSpecContainersLivenessProbe", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersLivenessProbe)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersLivenessProbe)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbe", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InitialDelaySeconds, actual.InitialDelaySeconds, dcl.Info{}, fn.AddNest("InitialDelaySeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeriodSeconds, actual.PeriodSeconds, dcl.Info{}, fn.AddNest("PeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Exec, actual.Exec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeExecNewStyle}, fn.AddNest("Exec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpGet, actual.HttpGet, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetNewStyle}, fn.AddNest("HttpGet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TcpSocket, actual.TcpSocket, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketNewStyle}, fn.AddNest("TcpSocket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersLivenessProbe(c *Client, desired, actual *ServiceSpecTemplateSpecContainersLivenessProbe) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.InitialDelaySeconds, actual.InitialDelaySeconds) && !dcl.IsZeroValue(desired.InitialDelaySeconds) {
		c.Config.Logger.Infof("Diff in InitialDelaySeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitialDelaySeconds), dcl.SprintResource(actual.InitialDelaySeconds))
		return true
	}
	if !reflect.DeepEqual(desired.TimeoutSeconds, actual.TimeoutSeconds) && !dcl.IsZeroValue(desired.TimeoutSeconds) {
		c.Config.Logger.Infof("Diff in TimeoutSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeoutSeconds), dcl.SprintResource(actual.TimeoutSeconds))
		return true
	}
	if !reflect.DeepEqual(desired.PeriodSeconds, actual.PeriodSeconds) && !dcl.IsZeroValue(desired.PeriodSeconds) {
		c.Config.Logger.Infof("Diff in PeriodSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PeriodSeconds), dcl.SprintResource(actual.PeriodSeconds))
		return true
	}
	if !reflect.DeepEqual(desired.SuccessThreshold, actual.SuccessThreshold) && !dcl.IsZeroValue(desired.SuccessThreshold) {
		c.Config.Logger.Infof("Diff in SuccessThreshold.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SuccessThreshold), dcl.SprintResource(actual.SuccessThreshold))
		return true
	}
	if !reflect.DeepEqual(desired.FailureThreshold, actual.FailureThreshold) && !dcl.IsZeroValue(desired.FailureThreshold) {
		c.Config.Logger.Infof("Diff in FailureThreshold.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailureThreshold), dcl.SprintResource(actual.FailureThreshold))
		return true
	}
	if compareServiceSpecTemplateSpecContainersLivenessProbeExec(c, desired.Exec, actual.Exec) && !dcl.IsZeroValue(desired.Exec) {
		c.Config.Logger.Infof("Diff in Exec.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Exec), dcl.SprintResource(actual.Exec))
		return true
	}
	if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, desired.HttpGet, actual.HttpGet) && !dcl.IsZeroValue(desired.HttpGet) {
		c.Config.Logger.Infof("Diff in HttpGet.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpGet), dcl.SprintResource(actual.HttpGet))
		return true
	}
	if compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, desired.TcpSocket, actual.TcpSocket) && !dcl.IsZeroValue(desired.TcpSocket) {
		c.Config.Logger.Infof("Diff in TcpSocket.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TcpSocket), dcl.SprintResource(actual.TcpSocket))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersLivenessProbe) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbe, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersLivenessProbe(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbe, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersLivenessProbe) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbe, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbe, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersLivenessProbe(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbe, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersLivenessProbeExec)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersLivenessProbeExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeExec or *ServiceSpecTemplateSpecContainersLivenessProbeExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersLivenessProbeExec)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersLivenessProbeExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersLivenessProbeExec(c *Client, desired, actual *ServiceSpecTemplateSpecContainersLivenessProbeExec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Command, actual.Command) && !dcl.IsZeroValue(desired.Command) {
		c.Config.Logger.Infof("Diff in Command.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Command), dcl.SprintResource(actual.Command))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeExecSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersLivenessProbeExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeExec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersLivenessProbeExec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeExec, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeExecMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeExec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeExec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersLivenessProbeExec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeExec, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersLivenessProbeHttpGet)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersLivenessProbeHttpGet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeHttpGet or *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersLivenessProbeHttpGet)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersLivenessProbeHttpGet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeHttpGet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scheme, actual.Scheme, dcl.Info{}, fn.AddNest("Scheme")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaders, actual.HttpHeaders, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersNewStyle}, fn.AddNest("HttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c *Client, desired, actual *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if !dcl.StringCanonicalize(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) {
		c.Config.Logger.Infof("Diff in Host.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	if !dcl.StringCanonicalize(desired.Scheme, actual.Scheme) && !dcl.IsZeroValue(desired.Scheme) {
		c.Config.Logger.Infof("Diff in Scheme.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scheme), dcl.SprintResource(actual.Scheme))
		return true
	}
	if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c, desired.HttpHeaders, actual.HttpHeaders) && !dcl.IsZeroValue(desired.HttpHeaders) {
		c.Config.Logger.Infof("Diff in HttpHeaders.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpHeaders), dcl.SprintResource(actual.HttpHeaders))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders or *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c *Client, desired, actual *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket or *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c *Client, desired, actual *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Port, actual.Port) && !dcl.IsZeroValue(desired.Port) {
		c.Config.Logger.Infof("Diff in Port.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Port), dcl.SprintResource(actual.Port))
		return true
	}
	if !dcl.StringCanonicalize(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) {
		c.Config.Logger.Infof("Diff in Host.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersReadinessProbe)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersReadinessProbe)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbe or *ServiceSpecTemplateSpecContainersReadinessProbe", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersReadinessProbe)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersReadinessProbe)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbe", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InitialDelaySeconds, actual.InitialDelaySeconds, dcl.Info{}, fn.AddNest("InitialDelaySeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeriodSeconds, actual.PeriodSeconds, dcl.Info{}, fn.AddNest("PeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Exec, actual.Exec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeExecNewStyle}, fn.AddNest("Exec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpGet, actual.HttpGet, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetNewStyle}, fn.AddNest("HttpGet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TcpSocket, actual.TcpSocket, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketNewStyle}, fn.AddNest("TcpSocket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersReadinessProbe(c *Client, desired, actual *ServiceSpecTemplateSpecContainersReadinessProbe) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.InitialDelaySeconds, actual.InitialDelaySeconds) && !dcl.IsZeroValue(desired.InitialDelaySeconds) {
		c.Config.Logger.Infof("Diff in InitialDelaySeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitialDelaySeconds), dcl.SprintResource(actual.InitialDelaySeconds))
		return true
	}
	if !reflect.DeepEqual(desired.TimeoutSeconds, actual.TimeoutSeconds) && !dcl.IsZeroValue(desired.TimeoutSeconds) {
		c.Config.Logger.Infof("Diff in TimeoutSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeoutSeconds), dcl.SprintResource(actual.TimeoutSeconds))
		return true
	}
	if !reflect.DeepEqual(desired.PeriodSeconds, actual.PeriodSeconds) && !dcl.IsZeroValue(desired.PeriodSeconds) {
		c.Config.Logger.Infof("Diff in PeriodSeconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PeriodSeconds), dcl.SprintResource(actual.PeriodSeconds))
		return true
	}
	if !reflect.DeepEqual(desired.SuccessThreshold, actual.SuccessThreshold) && !dcl.IsZeroValue(desired.SuccessThreshold) {
		c.Config.Logger.Infof("Diff in SuccessThreshold.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SuccessThreshold), dcl.SprintResource(actual.SuccessThreshold))
		return true
	}
	if !reflect.DeepEqual(desired.FailureThreshold, actual.FailureThreshold) && !dcl.IsZeroValue(desired.FailureThreshold) {
		c.Config.Logger.Infof("Diff in FailureThreshold.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailureThreshold), dcl.SprintResource(actual.FailureThreshold))
		return true
	}
	if compareServiceSpecTemplateSpecContainersReadinessProbeExec(c, desired.Exec, actual.Exec) && !dcl.IsZeroValue(desired.Exec) {
		c.Config.Logger.Infof("Diff in Exec.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Exec), dcl.SprintResource(actual.Exec))
		return true
	}
	if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, desired.HttpGet, actual.HttpGet) && !dcl.IsZeroValue(desired.HttpGet) {
		c.Config.Logger.Infof("Diff in HttpGet.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpGet), dcl.SprintResource(actual.HttpGet))
		return true
	}
	if compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, desired.TcpSocket, actual.TcpSocket) && !dcl.IsZeroValue(desired.TcpSocket) {
		c.Config.Logger.Infof("Diff in TcpSocket.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TcpSocket), dcl.SprintResource(actual.TcpSocket))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersReadinessProbe) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbe, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersReadinessProbe(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbe, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersReadinessProbe) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbe, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbe, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersReadinessProbe(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbe, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeExecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersReadinessProbeExec)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersReadinessProbeExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeExec or *ServiceSpecTemplateSpecContainersReadinessProbeExec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersReadinessProbeExec)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersReadinessProbeExec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeExec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersReadinessProbeExec(c *Client, desired, actual *ServiceSpecTemplateSpecContainersReadinessProbeExec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Command, actual.Command) && !dcl.IsZeroValue(desired.Command) {
		c.Config.Logger.Infof("Diff in Command.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Command), dcl.SprintResource(actual.Command))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeExecSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersReadinessProbeExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeExec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersReadinessProbeExec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeExec, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeExecMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeExec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeExec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersReadinessProbeExec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeExec, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersReadinessProbeHttpGet)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersReadinessProbeHttpGet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeHttpGet or *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersReadinessProbeHttpGet)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersReadinessProbeHttpGet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeHttpGet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scheme, actual.Scheme, dcl.Info{}, fn.AddNest("Scheme")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaders, actual.HttpHeaders, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersNewStyle}, fn.AddNest("HttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c *Client, desired, actual *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if !dcl.StringCanonicalize(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) {
		c.Config.Logger.Infof("Diff in Host.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	if !dcl.StringCanonicalize(desired.Scheme, actual.Scheme) && !dcl.IsZeroValue(desired.Scheme) {
		c.Config.Logger.Infof("Diff in Scheme.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scheme), dcl.SprintResource(actual.Scheme))
		return true
	}
	if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c, desired.HttpHeaders, actual.HttpHeaders) && !dcl.IsZeroValue(desired.HttpHeaders) {
		c.Config.Logger.Infof("Diff in HttpHeaders.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpHeaders), dcl.SprintResource(actual.HttpHeaders))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders or *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c *Client, desired, actual *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket or *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c *Client, desired, actual *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Port, actual.Port) && !dcl.IsZeroValue(desired.Port) {
		c.Config.Logger.Infof("Diff in Port.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Port), dcl.SprintResource(actual.Port))
		return true
	}
	if !dcl.StringCanonicalize(desired.Host, actual.Host) && !dcl.IsZeroValue(desired.Host) {
		c.Config.Logger.Infof("Diff in Host.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Host), dcl.SprintResource(actual.Host))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersSecurityContextNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecContainersSecurityContext)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecContainersSecurityContext)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersSecurityContext or *ServiceSpecTemplateSpecContainersSecurityContext", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecContainersSecurityContext)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecContainersSecurityContext)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecContainersSecurityContext", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RunAsUser, actual.RunAsUser, dcl.Info{}, fn.AddNest("RunAsUser")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecContainersSecurityContext(c *Client, desired, actual *ServiceSpecTemplateSpecContainersSecurityContext) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.RunAsUser, actual.RunAsUser) && !dcl.IsZeroValue(desired.RunAsUser) {
		c.Config.Logger.Infof("Diff in RunAsUser.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RunAsUser), dcl.SprintResource(actual.RunAsUser))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecContainersSecurityContextSlice(c *Client, desired, actual []ServiceSpecTemplateSpecContainersSecurityContext) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersSecurityContext, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecContainersSecurityContext(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersSecurityContext, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecContainersSecurityContextMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecContainersSecurityContext) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecContainersSecurityContext, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersSecurityContext, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecContainersSecurityContext(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecContainersSecurityContext, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecVolumes)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumes or *ServiceSpecTemplateSpecVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecVolumes)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesSecretNewStyle}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigMap, actual.ConfigMap, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesConfigMapNewStyle}, fn.AddNest("ConfigMap")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecVolumes(c *Client, desired, actual *ServiceSpecTemplateSpecVolumes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if compareServiceSpecTemplateSpecVolumesSecret(c, desired.Secret, actual.Secret) && !dcl.IsZeroValue(desired.Secret) {
		c.Config.Logger.Infof("Diff in Secret.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Secret), dcl.SprintResource(actual.Secret))
		return true
	}
	if compareServiceSpecTemplateSpecVolumesConfigMap(c, desired.ConfigMap, actual.ConfigMap) && !dcl.IsZeroValue(desired.ConfigMap) {
		c.Config.Logger.Infof("Diff in ConfigMap.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigMap), dcl.SprintResource(actual.ConfigMap))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSlice(c *Client, desired, actual []ServiceSpecTemplateSpecVolumes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecVolumes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumes, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecVolumes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecVolumes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumes, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecVolumesSecret)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesSecret or *ServiceSpecTemplateSpecVolumesSecret", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecVolumesSecret)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesSecret", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SecretName, actual.SecretName, dcl.Info{}, fn.AddNest("SecretName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesSecretItemsNewStyle}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.Info{}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecVolumesSecret(c *Client, desired, actual *ServiceSpecTemplateSpecVolumesSecret) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.SecretName, actual.SecretName) && !dcl.IsZeroValue(desired.SecretName) {
		c.Config.Logger.Infof("Diff in SecretName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecretName), dcl.SprintResource(actual.SecretName))
		return true
	}
	if compareServiceSpecTemplateSpecVolumesSecretItemsSlice(c, desired.Items, actual.Items) && !dcl.IsZeroValue(desired.Items) {
		c.Config.Logger.Infof("Diff in Items.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Items), dcl.SprintResource(actual.Items))
		return true
	}
	if !reflect.DeepEqual(desired.DefaultMode, actual.DefaultMode) && !dcl.IsZeroValue(desired.DefaultMode) {
		c.Config.Logger.Infof("Diff in DefaultMode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultMode), dcl.SprintResource(actual.DefaultMode))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretSlice(c *Client, desired, actual []ServiceSpecTemplateSpecVolumesSecret) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesSecret, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecVolumesSecret(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecret, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecVolumesSecret) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesSecret, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecret, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecVolumesSecret(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecret, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretItemsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecVolumesSecretItems)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesSecretItems or *ServiceSpecTemplateSpecVolumesSecretItems", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecVolumesSecretItems)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesSecretItems", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecVolumesSecretItems(c *Client, desired, actual *ServiceSpecTemplateSpecVolumesSecretItems) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) {
		c.Config.Logger.Infof("Diff in Mode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Mode), dcl.SprintResource(actual.Mode))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretItemsSlice(c *Client, desired, actual []ServiceSpecTemplateSpecVolumesSecretItems) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesSecretItems, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecVolumesSecretItems(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecretItems, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesSecretItemsMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecVolumesSecretItems) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesSecretItems, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecretItems, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecVolumesSecretItems(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesSecretItems, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecVolumesConfigMap)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecVolumesConfigMap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesConfigMap or *ServiceSpecTemplateSpecVolumesConfigMap", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecVolumesConfigMap)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecVolumesConfigMap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesConfigMap", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesConfigMapItemsNewStyle}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.Info{}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecVolumesConfigMap(c *Client, desired, actual *ServiceSpecTemplateSpecVolumesConfigMap) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if compareServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c, desired.Items, actual.Items) && !dcl.IsZeroValue(desired.Items) {
		c.Config.Logger.Infof("Diff in Items.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Items), dcl.SprintResource(actual.Items))
		return true
	}
	if !reflect.DeepEqual(desired.DefaultMode, actual.DefaultMode) && !dcl.IsZeroValue(desired.DefaultMode) {
		c.Config.Logger.Infof("Diff in DefaultMode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultMode), dcl.SprintResource(actual.DefaultMode))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Optional, actual.Optional) && !dcl.IsZeroValue(desired.Optional) {
		c.Config.Logger.Infof("Diff in Optional.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Optional), dcl.SprintResource(actual.Optional))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapSlice(c *Client, desired, actual []ServiceSpecTemplateSpecVolumesConfigMap) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesConfigMap, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecVolumesConfigMap(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMap, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecVolumesConfigMap) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesConfigMap, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMap, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecVolumesConfigMap(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMap, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapItemsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTemplateSpecVolumesConfigMapItems)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTemplateSpecVolumesConfigMapItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesConfigMapItems or *ServiceSpecTemplateSpecVolumesConfigMapItems", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTemplateSpecVolumesConfigMapItems)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTemplateSpecVolumesConfigMapItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTemplateSpecVolumesConfigMapItems", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTemplateSpecVolumesConfigMapItems(c *Client, desired, actual *ServiceSpecTemplateSpecVolumesConfigMapItems) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) {
		c.Config.Logger.Infof("Diff in Mode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Mode), dcl.SprintResource(actual.Mode))
		return true
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c *Client, desired, actual []ServiceSpecTemplateSpecVolumesConfigMapItems) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesConfigMapItems, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTemplateSpecVolumesConfigMapItems(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMapItems, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTemplateSpecVolumesConfigMapItemsMap(c *Client, desired, actual map[string]ServiceSpecTemplateSpecVolumesConfigMapItems) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTemplateSpecVolumesConfigMapItems, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMapItems, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTemplateSpecVolumesConfigMapItems(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTemplateSpecVolumesConfigMapItems, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceSpecTrafficNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceSpecTraffic)
	if !ok {
		desiredNotPointer, ok := d.(ServiceSpecTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTraffic or *ServiceSpecTraffic", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceSpecTraffic)
	if !ok {
		actualNotPointer, ok := a.(ServiceSpecTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceSpecTraffic", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigurationName, actual.ConfigurationName, dcl.Info{}, fn.AddNest("ConfigurationName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionName, actual.RevisionName, dcl.Info{}, fn.AddNest("RevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestRevision, actual.LatestRevision, dcl.Info{}, fn.AddNest("LatestRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OutputOnly: true}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceSpecTraffic(c *Client, desired, actual *ServiceSpecTraffic) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.ConfigurationName, actual.ConfigurationName) && !dcl.IsZeroValue(desired.ConfigurationName) {
		c.Config.Logger.Infof("Diff in ConfigurationName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigurationName), dcl.SprintResource(actual.ConfigurationName))
		return true
	}
	if !dcl.StringCanonicalize(desired.RevisionName, actual.RevisionName) && !dcl.IsZeroValue(desired.RevisionName) {
		c.Config.Logger.Infof("Diff in RevisionName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RevisionName), dcl.SprintResource(actual.RevisionName))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) {
		c.Config.Logger.Infof("Diff in Percent.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	if !dcl.StringCanonicalize(desired.Tag, actual.Tag) && !dcl.IsZeroValue(desired.Tag) {
		c.Config.Logger.Infof("Diff in Tag.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tag), dcl.SprintResource(actual.Tag))
		return true
	}
	if !dcl.BoolCanonicalize(desired.LatestRevision, actual.LatestRevision) && !dcl.IsZeroValue(desired.LatestRevision) {
		c.Config.Logger.Infof("Diff in LatestRevision.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LatestRevision), dcl.SprintResource(actual.LatestRevision))
		return true
	}
	return false
}

func compareServiceSpecTrafficSlice(c *Client, desired, actual []ServiceSpecTraffic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTraffic, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceSpecTraffic(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceSpecTraffic, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceSpecTrafficMap(c *Client, desired, actual map[string]ServiceSpecTraffic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceSpecTraffic, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceSpecTraffic, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceSpecTraffic(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceSpecTraffic, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceStatus)
	if !ok {
		desiredNotPointer, ok := d.(ServiceStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatus or *ServiceStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceStatus)
	if !ok {
		actualNotPointer, ok := a.(ServiceStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObservedGeneration, actual.ObservedGeneration, dcl.Info{}, fn.AddNest("ObservedGeneration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{ObjectFunction: compareServiceStatusConditionsNewStyle}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestReadyRevisionName, actual.LatestReadyRevisionName, dcl.Info{}, fn.AddNest("LatestReadyRevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestCreatedRevisionName, actual.LatestCreatedRevisionName, dcl.Info{}, fn.AddNest("LatestCreatedRevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.Info{ObjectFunction: compareServiceStatusTrafficNewStyle}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Address, actual.Address, dcl.Info{ObjectFunction: compareServiceStatusAddressNewStyle}, fn.AddNest("Address")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceStatus(c *Client, desired, actual *ServiceStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.ObservedGeneration, actual.ObservedGeneration) && !dcl.IsZeroValue(desired.ObservedGeneration) {
		c.Config.Logger.Infof("Diff in ObservedGeneration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ObservedGeneration), dcl.SprintResource(actual.ObservedGeneration))
		return true
	}
	if compareServiceStatusConditionsSlice(c, desired.Conditions, actual.Conditions) && !dcl.IsZeroValue(desired.Conditions) {
		c.Config.Logger.Infof("Diff in Conditions.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Conditions), dcl.SprintResource(actual.Conditions))
		return true
	}
	if !dcl.StringCanonicalize(desired.LatestReadyRevisionName, actual.LatestReadyRevisionName) && !dcl.IsZeroValue(desired.LatestReadyRevisionName) {
		c.Config.Logger.Infof("Diff in LatestReadyRevisionName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LatestReadyRevisionName), dcl.SprintResource(actual.LatestReadyRevisionName))
		return true
	}
	if !dcl.StringCanonicalize(desired.LatestCreatedRevisionName, actual.LatestCreatedRevisionName) && !dcl.IsZeroValue(desired.LatestCreatedRevisionName) {
		c.Config.Logger.Infof("Diff in LatestCreatedRevisionName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LatestCreatedRevisionName), dcl.SprintResource(actual.LatestCreatedRevisionName))
		return true
	}
	if compareServiceStatusTrafficSlice(c, desired.Traffic, actual.Traffic) && !dcl.IsZeroValue(desired.Traffic) {
		c.Config.Logger.Infof("Diff in Traffic.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Traffic), dcl.SprintResource(actual.Traffic))
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	if compareServiceStatusAddress(c, desired.Address, actual.Address) && !dcl.IsZeroValue(desired.Address) {
		c.Config.Logger.Infof("Diff in Address.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Address), dcl.SprintResource(actual.Address))
		return true
	}
	return false
}

func compareServiceStatusSlice(c *Client, desired, actual []ServiceStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStatus, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStatusMap(c *Client, desired, actual map[string]ServiceStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceStatus, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceStatusConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceStatusConditions)
	if !ok {
		desiredNotPointer, ok := d.(ServiceStatusConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusConditions or *ServiceStatusConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceStatusConditions)
	if !ok {
		actualNotPointer, ok := a.(ServiceStatusConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.Info{}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.Info{ObjectFunction: compareServiceStatusConditionsLastTransitionTimeNewStyle}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.Info{}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceStatusConditions(c *Client, desired, actual *ServiceStatusConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if !dcl.StringCanonicalize(desired.Status, actual.Status) && !dcl.IsZeroValue(desired.Status) {
		c.Config.Logger.Infof("Diff in Status.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Status), dcl.SprintResource(actual.Status))
		return true
	}
	if !dcl.StringCanonicalize(desired.Reason, actual.Reason) && !dcl.IsZeroValue(desired.Reason) {
		c.Config.Logger.Infof("Diff in Reason.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Reason), dcl.SprintResource(actual.Reason))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if compareServiceStatusConditionsLastTransitionTime(c, desired.LastTransitionTime, actual.LastTransitionTime) && !dcl.IsZeroValue(desired.LastTransitionTime) {
		c.Config.Logger.Infof("Diff in LastTransitionTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LastTransitionTime), dcl.SprintResource(actual.LastTransitionTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Severity, actual.Severity) && !dcl.IsZeroValue(desired.Severity) {
		c.Config.Logger.Infof("Diff in Severity.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Severity), dcl.SprintResource(actual.Severity))
		return true
	}
	return false
}

func compareServiceStatusConditionsSlice(c *Client, desired, actual []ServiceStatusConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStatusConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStatusConditions, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStatusConditionsMap(c *Client, desired, actual map[string]ServiceStatusConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusConditions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceStatusConditions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceStatusConditions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceStatusConditions, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceStatusConditionsLastTransitionTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceStatusConditionsLastTransitionTime)
	if !ok {
		desiredNotPointer, ok := d.(ServiceStatusConditionsLastTransitionTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusConditionsLastTransitionTime or *ServiceStatusConditionsLastTransitionTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceStatusConditionsLastTransitionTime)
	if !ok {
		actualNotPointer, ok := a.(ServiceStatusConditionsLastTransitionTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusConditionsLastTransitionTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceStatusConditionsLastTransitionTime(c *Client, desired, actual *ServiceStatusConditionsLastTransitionTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareServiceStatusConditionsLastTransitionTimeSlice(c *Client, desired, actual []ServiceStatusConditionsLastTransitionTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusConditionsLastTransitionTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStatusConditionsLastTransitionTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStatusConditionsLastTransitionTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStatusConditionsLastTransitionTimeMap(c *Client, desired, actual map[string]ServiceStatusConditionsLastTransitionTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusConditionsLastTransitionTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceStatusConditionsLastTransitionTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceStatusConditionsLastTransitionTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceStatusConditionsLastTransitionTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceStatusTrafficNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceStatusTraffic)
	if !ok {
		desiredNotPointer, ok := d.(ServiceStatusTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusTraffic or *ServiceStatusTraffic", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceStatusTraffic)
	if !ok {
		actualNotPointer, ok := a.(ServiceStatusTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusTraffic", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigurationName, actual.ConfigurationName, dcl.Info{}, fn.AddNest("ConfigurationName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionName, actual.RevisionName, dcl.Info{}, fn.AddNest("RevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestRevision, actual.LatestRevision, dcl.Info{}, fn.AddNest("LatestRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OutputOnly: true}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceStatusTraffic(c *Client, desired, actual *ServiceStatusTraffic) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.ConfigurationName, actual.ConfigurationName) && !dcl.IsZeroValue(desired.ConfigurationName) {
		c.Config.Logger.Infof("Diff in ConfigurationName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigurationName), dcl.SprintResource(actual.ConfigurationName))
		return true
	}
	if !dcl.StringCanonicalize(desired.RevisionName, actual.RevisionName) && !dcl.IsZeroValue(desired.RevisionName) {
		c.Config.Logger.Infof("Diff in RevisionName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RevisionName), dcl.SprintResource(actual.RevisionName))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) {
		c.Config.Logger.Infof("Diff in Percent.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	if !dcl.StringCanonicalize(desired.Tag, actual.Tag) && !dcl.IsZeroValue(desired.Tag) {
		c.Config.Logger.Infof("Diff in Tag.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tag), dcl.SprintResource(actual.Tag))
		return true
	}
	if !dcl.BoolCanonicalize(desired.LatestRevision, actual.LatestRevision) && !dcl.IsZeroValue(desired.LatestRevision) {
		c.Config.Logger.Infof("Diff in LatestRevision.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LatestRevision), dcl.SprintResource(actual.LatestRevision))
		return true
	}
	return false
}

func compareServiceStatusTrafficSlice(c *Client, desired, actual []ServiceStatusTraffic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusTraffic, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStatusTraffic(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStatusTraffic, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStatusTrafficMap(c *Client, desired, actual map[string]ServiceStatusTraffic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusTraffic, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceStatusTraffic, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceStatusTraffic(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceStatusTraffic, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareServiceStatusAddressNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceStatusAddress)
	if !ok {
		desiredNotPointer, ok := d.(ServiceStatusAddress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusAddress or *ServiceStatusAddress", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceStatusAddress)
	if !ok {
		actualNotPointer, ok := a.(ServiceStatusAddress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceStatusAddress", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceStatusAddress(c *Client, desired, actual *ServiceStatusAddress) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	return false
}

func compareServiceStatusAddressSlice(c *Client, desired, actual []ServiceStatusAddress) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusAddress, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStatusAddress(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStatusAddress, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStatusAddressMap(c *Client, desired, actual map[string]ServiceStatusAddress) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStatusAddress, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ServiceStatusAddress, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareServiceStatusAddress(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ServiceStatusAddress, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Service) urlNormalized() *Service {
	normalized := dcl.Copy(*r).(Service)
	normalized.ApiVersion = dcl.SelfLinkToName(r.ApiVersion)
	normalized.Kind = dcl.SelfLinkToName(r.Kind)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	return &normalized
}

func (r *Service) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Service) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Service) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Service) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateService" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("/apis/serving.knative.dev/v1/namespaces/{{project}}/services/{{name}}", "https://{{location}}-run.googleapis.com/", userBasePath, fields), nil

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

	return flattenService(c, m), nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.ApiVersion; !dcl.IsEmptyValueIndirect(v) {
		m["apiVersion"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := expandServiceMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if v != nil {
		m["metadata"] = v
	}
	if v, err := expandServiceSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if v != nil {
		m["spec"] = v
	}
	if v, err := expandServiceStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if v != nil {
		m["status"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
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

	r := &Service{}
	r.ApiVersion = dcl.FlattenString(m["apiVersion"])
	if _, ok := m["apiVersion"]; !ok {
		c.Config.Logger.Info("Using default value for apiVersion")
		r.ApiVersion = dcl.String("serving.knative.dev/v1")
	}
	r.Kind = dcl.FlattenString(m["kind"])
	if _, ok := m["kind"]; !ok {
		c.Config.Logger.Info("Using default value for kind")
		r.Kind = dcl.String("Service")
	}
	r.Metadata = flattenServiceMetadata(c, m["metadata"])
	r.Spec = flattenServiceSpec(c, m["spec"])
	r.Status = flattenServiceStatus(c, m["status"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceMetadataMap expands the contents of ServiceMetadata into a JSON
// request object.
func expandServiceMetadataMap(c *Client, f map[string]ServiceMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceMetadataSlice expands the contents of ServiceMetadata into a JSON
// request object.
func expandServiceMetadataSlice(c *Client, f []ServiceMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceMetadataMap flattens the contents of ServiceMetadata from a JSON
// response object.
func flattenServiceMetadataMap(c *Client, i interface{}) map[string]ServiceMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceMetadata{}
	}

	if len(a) == 0 {
		return map[string]ServiceMetadata{}
	}

	items := make(map[string]ServiceMetadata)
	for k, item := range a {
		items[k] = *flattenServiceMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceMetadataSlice flattens the contents of ServiceMetadata from a JSON
// response object.
func flattenServiceMetadataSlice(c *Client, i interface{}) []ServiceMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceMetadata{}
	}

	if len(a) == 0 {
		return []ServiceMetadata{}
	}

	items := make([]ServiceMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceMetadata expands an instance of ServiceMetadata into a JSON
// request object.
func expandServiceMetadata(c *Client, f *ServiceMetadata) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.GenerateName; !dcl.IsEmptyValueIndirect(v) {
		m["generateName"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.ResourceVersion; !dcl.IsEmptyValueIndirect(v) {
		m["resourceVersion"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}
	if v, err := expandServiceMetadataCreateTime(c, f.CreateTime); err != nil {
		return nil, fmt.Errorf("error expanding CreateTime into creationTimestamp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandServiceMetadataOwnerReferencesSlice(c, f.OwnerReferences); err != nil {
		return nil, fmt.Errorf("error expanding OwnerReferences into ownerReferences: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ownerReferences"] = v
	}
	if v, err := expandServiceMetadataDeleteTime(c, f.DeleteTime); err != nil {
		return nil, fmt.Errorf("error expanding DeleteTime into deletionTimestamp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deletionTimestamp"] = v
	}
	if v := f.DeletionGracePeriodSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["deletionGracePeriodSeconds"] = v
	}
	if v := f.Finalizers; !dcl.IsEmptyValueIndirect(v) {
		m["finalizers"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}

	return m, nil
}

// flattenServiceMetadata flattens an instance of ServiceMetadata from a JSON
// response object.
func flattenServiceMetadata(c *Client, i interface{}) *ServiceMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceMetadata{}
	r.Name = dcl.FlattenString(m["name"])
	r.GenerateName = dcl.FlattenString(m["generateName"])
	r.Namespace = dcl.FlattenString(m["namespace"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Uid = dcl.FlattenString(m["uid"])
	r.ResourceVersion = dcl.FlattenString(m["resourceVersion"])
	r.Generation = dcl.FlattenInteger(m["generation"])
	r.CreateTime = flattenServiceMetadataCreateTime(c, m["creationTimestamp"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.OwnerReferences = flattenServiceMetadataOwnerReferencesSlice(c, m["ownerReferences"])
	r.DeleteTime = flattenServiceMetadataDeleteTime(c, m["deletionTimestamp"])
	r.DeletionGracePeriodSeconds = dcl.FlattenInteger(m["deletionGracePeriodSeconds"])
	r.Finalizers = dcl.FlattenStringSlice(m["finalizers"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])

	return r
}

// expandServiceMetadataCreateTimeMap expands the contents of ServiceMetadataCreateTime into a JSON
// request object.
func expandServiceMetadataCreateTimeMap(c *Client, f map[string]ServiceMetadataCreateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceMetadataCreateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceMetadataCreateTimeSlice expands the contents of ServiceMetadataCreateTime into a JSON
// request object.
func expandServiceMetadataCreateTimeSlice(c *Client, f []ServiceMetadataCreateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceMetadataCreateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceMetadataCreateTimeMap flattens the contents of ServiceMetadataCreateTime from a JSON
// response object.
func flattenServiceMetadataCreateTimeMap(c *Client, i interface{}) map[string]ServiceMetadataCreateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceMetadataCreateTime{}
	}

	if len(a) == 0 {
		return map[string]ServiceMetadataCreateTime{}
	}

	items := make(map[string]ServiceMetadataCreateTime)
	for k, item := range a {
		items[k] = *flattenServiceMetadataCreateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceMetadataCreateTimeSlice flattens the contents of ServiceMetadataCreateTime from a JSON
// response object.
func flattenServiceMetadataCreateTimeSlice(c *Client, i interface{}) []ServiceMetadataCreateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceMetadataCreateTime{}
	}

	if len(a) == 0 {
		return []ServiceMetadataCreateTime{}
	}

	items := make([]ServiceMetadataCreateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceMetadataCreateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceMetadataCreateTime expands an instance of ServiceMetadataCreateTime into a JSON
// request object.
func expandServiceMetadataCreateTime(c *Client, f *ServiceMetadataCreateTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenServiceMetadataCreateTime flattens an instance of ServiceMetadataCreateTime from a JSON
// response object.
func flattenServiceMetadataCreateTime(c *Client, i interface{}) *ServiceMetadataCreateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceMetadataCreateTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandServiceMetadataOwnerReferencesMap expands the contents of ServiceMetadataOwnerReferences into a JSON
// request object.
func expandServiceMetadataOwnerReferencesMap(c *Client, f map[string]ServiceMetadataOwnerReferences) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceMetadataOwnerReferences(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceMetadataOwnerReferencesSlice expands the contents of ServiceMetadataOwnerReferences into a JSON
// request object.
func expandServiceMetadataOwnerReferencesSlice(c *Client, f []ServiceMetadataOwnerReferences) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceMetadataOwnerReferences(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceMetadataOwnerReferencesMap flattens the contents of ServiceMetadataOwnerReferences from a JSON
// response object.
func flattenServiceMetadataOwnerReferencesMap(c *Client, i interface{}) map[string]ServiceMetadataOwnerReferences {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceMetadataOwnerReferences{}
	}

	if len(a) == 0 {
		return map[string]ServiceMetadataOwnerReferences{}
	}

	items := make(map[string]ServiceMetadataOwnerReferences)
	for k, item := range a {
		items[k] = *flattenServiceMetadataOwnerReferences(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceMetadataOwnerReferencesSlice flattens the contents of ServiceMetadataOwnerReferences from a JSON
// response object.
func flattenServiceMetadataOwnerReferencesSlice(c *Client, i interface{}) []ServiceMetadataOwnerReferences {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceMetadataOwnerReferences{}
	}

	if len(a) == 0 {
		return []ServiceMetadataOwnerReferences{}
	}

	items := make([]ServiceMetadataOwnerReferences, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceMetadataOwnerReferences(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceMetadataOwnerReferences expands an instance of ServiceMetadataOwnerReferences into a JSON
// request object.
func expandServiceMetadataOwnerReferences(c *Client, f *ServiceMetadataOwnerReferences) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ApiVersion; !dcl.IsEmptyValueIndirect(v) {
		m["apiVersion"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.Controller; !dcl.IsEmptyValueIndirect(v) {
		m["controller"] = v
	}
	if v := f.BlockOwnerDeletion; !dcl.IsEmptyValueIndirect(v) {
		m["blockOwnerDeletion"] = v
	}

	return m, nil
}

// flattenServiceMetadataOwnerReferences flattens an instance of ServiceMetadataOwnerReferences from a JSON
// response object.
func flattenServiceMetadataOwnerReferences(c *Client, i interface{}) *ServiceMetadataOwnerReferences {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceMetadataOwnerReferences{}
	r.ApiVersion = dcl.FlattenString(m["apiVersion"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.Name = dcl.FlattenString(m["name"])
	r.Uid = dcl.FlattenString(m["uid"])
	r.Controller = dcl.FlattenBool(m["controller"])
	r.BlockOwnerDeletion = dcl.FlattenBool(m["blockOwnerDeletion"])

	return r
}

// expandServiceMetadataDeleteTimeMap expands the contents of ServiceMetadataDeleteTime into a JSON
// request object.
func expandServiceMetadataDeleteTimeMap(c *Client, f map[string]ServiceMetadataDeleteTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceMetadataDeleteTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceMetadataDeleteTimeSlice expands the contents of ServiceMetadataDeleteTime into a JSON
// request object.
func expandServiceMetadataDeleteTimeSlice(c *Client, f []ServiceMetadataDeleteTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceMetadataDeleteTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceMetadataDeleteTimeMap flattens the contents of ServiceMetadataDeleteTime from a JSON
// response object.
func flattenServiceMetadataDeleteTimeMap(c *Client, i interface{}) map[string]ServiceMetadataDeleteTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceMetadataDeleteTime{}
	}

	if len(a) == 0 {
		return map[string]ServiceMetadataDeleteTime{}
	}

	items := make(map[string]ServiceMetadataDeleteTime)
	for k, item := range a {
		items[k] = *flattenServiceMetadataDeleteTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceMetadataDeleteTimeSlice flattens the contents of ServiceMetadataDeleteTime from a JSON
// response object.
func flattenServiceMetadataDeleteTimeSlice(c *Client, i interface{}) []ServiceMetadataDeleteTime {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceMetadataDeleteTime{}
	}

	if len(a) == 0 {
		return []ServiceMetadataDeleteTime{}
	}

	items := make([]ServiceMetadataDeleteTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceMetadataDeleteTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceMetadataDeleteTime expands an instance of ServiceMetadataDeleteTime into a JSON
// request object.
func expandServiceMetadataDeleteTime(c *Client, f *ServiceMetadataDeleteTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenServiceMetadataDeleteTime flattens an instance of ServiceMetadataDeleteTime from a JSON
// response object.
func flattenServiceMetadataDeleteTime(c *Client, i interface{}) *ServiceMetadataDeleteTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceMetadataDeleteTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandServiceSpecMap expands the contents of ServiceSpec into a JSON
// request object.
func expandServiceSpecMap(c *Client, f map[string]ServiceSpec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecSlice expands the contents of ServiceSpec into a JSON
// request object.
func expandServiceSpecSlice(c *Client, f []ServiceSpec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecMap flattens the contents of ServiceSpec from a JSON
// response object.
func flattenServiceSpecMap(c *Client, i interface{}) map[string]ServiceSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpec{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpec{}
	}

	items := make(map[string]ServiceSpec)
	for k, item := range a {
		items[k] = *flattenServiceSpec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecSlice flattens the contents of ServiceSpec from a JSON
// response object.
func flattenServiceSpecSlice(c *Client, i interface{}) []ServiceSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpec{}
	}

	if len(a) == 0 {
		return []ServiceSpec{}
	}

	items := make([]ServiceSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpec(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpec expands an instance of ServiceSpec into a JSON
// request object.
func expandServiceSpec(c *Client, f *ServiceSpec) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplate(c, f.Template); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["template"] = v
	}
	if v, err := expandServiceSpecTrafficSlice(c, f.Traffic); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["traffic"] = v
	}

	return m, nil
}

// flattenServiceSpec flattens an instance of ServiceSpec from a JSON
// response object.
func flattenServiceSpec(c *Client, i interface{}) *ServiceSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpec{}
	r.Template = flattenServiceSpecTemplate(c, m["template"])
	r.Traffic = flattenServiceSpecTrafficSlice(c, m["traffic"])

	return r
}

// expandServiceSpecTemplateMap expands the contents of ServiceSpecTemplate into a JSON
// request object.
func expandServiceSpecTemplateMap(c *Client, f map[string]ServiceSpecTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSlice expands the contents of ServiceSpecTemplate into a JSON
// request object.
func expandServiceSpecTemplateSlice(c *Client, f []ServiceSpecTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateMap flattens the contents of ServiceSpecTemplate from a JSON
// response object.
func flattenServiceSpecTemplateMap(c *Client, i interface{}) map[string]ServiceSpecTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplate{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplate{}
	}

	items := make(map[string]ServiceSpecTemplate)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSlice flattens the contents of ServiceSpecTemplate from a JSON
// response object.
func flattenServiceSpecTemplateSlice(c *Client, i interface{}) []ServiceSpecTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplate{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplate{}
	}

	items := make([]ServiceSpecTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplate(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplate expands an instance of ServiceSpecTemplate into a JSON
// request object.
func expandServiceSpecTemplate(c *Client, f *ServiceSpecTemplate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := expandServiceSpecTemplateSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["spec"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplate flattens an instance of ServiceSpecTemplate from a JSON
// response object.
func flattenServiceSpecTemplate(c *Client, i interface{}) *ServiceSpecTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplate{}
	r.Metadata = flattenServiceSpecTemplateMetadata(c, m["metadata"])
	r.Spec = flattenServiceSpecTemplateSpec(c, m["spec"])

	return r
}

// expandServiceSpecTemplateMetadataMap expands the contents of ServiceSpecTemplateMetadata into a JSON
// request object.
func expandServiceSpecTemplateMetadataMap(c *Client, f map[string]ServiceSpecTemplateMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateMetadataSlice expands the contents of ServiceSpecTemplateMetadata into a JSON
// request object.
func expandServiceSpecTemplateMetadataSlice(c *Client, f []ServiceSpecTemplateMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateMetadataMap flattens the contents of ServiceSpecTemplateMetadata from a JSON
// response object.
func flattenServiceSpecTemplateMetadataMap(c *Client, i interface{}) map[string]ServiceSpecTemplateMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateMetadata{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateMetadata{}
	}

	items := make(map[string]ServiceSpecTemplateMetadata)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateMetadataSlice flattens the contents of ServiceSpecTemplateMetadata from a JSON
// response object.
func flattenServiceSpecTemplateMetadataSlice(c *Client, i interface{}) []ServiceSpecTemplateMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateMetadata{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateMetadata{}
	}

	items := make([]ServiceSpecTemplateMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateMetadata expands an instance of ServiceSpecTemplateMetadata into a JSON
// request object.
func expandServiceSpecTemplateMetadata(c *Client, f *ServiceSpecTemplateMetadata) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.GenerateName; !dcl.IsEmptyValueIndirect(v) {
		m["generateName"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.ResourceVersion; !dcl.IsEmptyValueIndirect(v) {
		m["resourceVersion"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}
	if v, err := expandServiceSpecTemplateMetadataCreateTime(c, f.CreateTime); err != nil {
		return nil, fmt.Errorf("error expanding CreateTime into creationTimestamp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandServiceSpecTemplateMetadataOwnerReferencesSlice(c, f.OwnerReferences); err != nil {
		return nil, fmt.Errorf("error expanding OwnerReferences into ownerReferences: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ownerReferences"] = v
	}
	if v, err := expandServiceSpecTemplateMetadataDeleteTime(c, f.DeleteTime); err != nil {
		return nil, fmt.Errorf("error expanding DeleteTime into deletionTimestamp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deletionTimestamp"] = v
	}
	if v := f.DeletionGracePeriodSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["deletionGracePeriodSeconds"] = v
	}
	if v := f.Finalizers; !dcl.IsEmptyValueIndirect(v) {
		m["finalizers"] = v
	}
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateMetadata flattens an instance of ServiceSpecTemplateMetadata from a JSON
// response object.
func flattenServiceSpecTemplateMetadata(c *Client, i interface{}) *ServiceSpecTemplateMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateMetadata{}
	r.Name = dcl.FlattenString(m["name"])
	r.GenerateName = dcl.FlattenString(m["generateName"])
	r.Namespace = dcl.FlattenString(m["namespace"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Uid = dcl.FlattenString(m["uid"])
	r.ResourceVersion = dcl.FlattenString(m["resourceVersion"])
	r.Generation = dcl.FlattenInteger(m["generation"])
	r.CreateTime = flattenServiceSpecTemplateMetadataCreateTime(c, m["creationTimestamp"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.OwnerReferences = flattenServiceSpecTemplateMetadataOwnerReferencesSlice(c, m["ownerReferences"])
	r.DeleteTime = flattenServiceSpecTemplateMetadataDeleteTime(c, m["deletionTimestamp"])
	r.DeletionGracePeriodSeconds = dcl.FlattenInteger(m["deletionGracePeriodSeconds"])
	r.Finalizers = dcl.FlattenStringSlice(m["finalizers"])
	r.ClusterName = dcl.FlattenString(m["clusterName"])

	return r
}

// expandServiceSpecTemplateMetadataCreateTimeMap expands the contents of ServiceSpecTemplateMetadataCreateTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataCreateTimeMap(c *Client, f map[string]ServiceSpecTemplateMetadataCreateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateMetadataCreateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateMetadataCreateTimeSlice expands the contents of ServiceSpecTemplateMetadataCreateTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataCreateTimeSlice(c *Client, f []ServiceSpecTemplateMetadataCreateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateMetadataCreateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateMetadataCreateTimeMap flattens the contents of ServiceSpecTemplateMetadataCreateTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataCreateTimeMap(c *Client, i interface{}) map[string]ServiceSpecTemplateMetadataCreateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateMetadataCreateTime{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateMetadataCreateTime{}
	}

	items := make(map[string]ServiceSpecTemplateMetadataCreateTime)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateMetadataCreateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateMetadataCreateTimeSlice flattens the contents of ServiceSpecTemplateMetadataCreateTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataCreateTimeSlice(c *Client, i interface{}) []ServiceSpecTemplateMetadataCreateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateMetadataCreateTime{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateMetadataCreateTime{}
	}

	items := make([]ServiceSpecTemplateMetadataCreateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateMetadataCreateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateMetadataCreateTime expands an instance of ServiceSpecTemplateMetadataCreateTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataCreateTime(c *Client, f *ServiceSpecTemplateMetadataCreateTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateMetadataCreateTime flattens an instance of ServiceSpecTemplateMetadataCreateTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataCreateTime(c *Client, i interface{}) *ServiceSpecTemplateMetadataCreateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateMetadataCreateTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandServiceSpecTemplateMetadataOwnerReferencesMap expands the contents of ServiceSpecTemplateMetadataOwnerReferences into a JSON
// request object.
func expandServiceSpecTemplateMetadataOwnerReferencesMap(c *Client, f map[string]ServiceSpecTemplateMetadataOwnerReferences) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateMetadataOwnerReferences(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateMetadataOwnerReferencesSlice expands the contents of ServiceSpecTemplateMetadataOwnerReferences into a JSON
// request object.
func expandServiceSpecTemplateMetadataOwnerReferencesSlice(c *Client, f []ServiceSpecTemplateMetadataOwnerReferences) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateMetadataOwnerReferences(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateMetadataOwnerReferencesMap flattens the contents of ServiceSpecTemplateMetadataOwnerReferences from a JSON
// response object.
func flattenServiceSpecTemplateMetadataOwnerReferencesMap(c *Client, i interface{}) map[string]ServiceSpecTemplateMetadataOwnerReferences {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateMetadataOwnerReferences{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateMetadataOwnerReferences{}
	}

	items := make(map[string]ServiceSpecTemplateMetadataOwnerReferences)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateMetadataOwnerReferences(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateMetadataOwnerReferencesSlice flattens the contents of ServiceSpecTemplateMetadataOwnerReferences from a JSON
// response object.
func flattenServiceSpecTemplateMetadataOwnerReferencesSlice(c *Client, i interface{}) []ServiceSpecTemplateMetadataOwnerReferences {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateMetadataOwnerReferences{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateMetadataOwnerReferences{}
	}

	items := make([]ServiceSpecTemplateMetadataOwnerReferences, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateMetadataOwnerReferences(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateMetadataOwnerReferences expands an instance of ServiceSpecTemplateMetadataOwnerReferences into a JSON
// request object.
func expandServiceSpecTemplateMetadataOwnerReferences(c *Client, f *ServiceSpecTemplateMetadataOwnerReferences) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ApiVersion; !dcl.IsEmptyValueIndirect(v) {
		m["apiVersion"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.Controller; !dcl.IsEmptyValueIndirect(v) {
		m["controller"] = v
	}
	if v := f.BlockOwnerDeletion; !dcl.IsEmptyValueIndirect(v) {
		m["blockOwnerDeletion"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateMetadataOwnerReferences flattens an instance of ServiceSpecTemplateMetadataOwnerReferences from a JSON
// response object.
func flattenServiceSpecTemplateMetadataOwnerReferences(c *Client, i interface{}) *ServiceSpecTemplateMetadataOwnerReferences {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateMetadataOwnerReferences{}
	r.ApiVersion = dcl.FlattenString(m["apiVersion"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.Name = dcl.FlattenString(m["name"])
	r.Uid = dcl.FlattenString(m["uid"])
	r.Controller = dcl.FlattenBool(m["controller"])
	r.BlockOwnerDeletion = dcl.FlattenBool(m["blockOwnerDeletion"])

	return r
}

// expandServiceSpecTemplateMetadataDeleteTimeMap expands the contents of ServiceSpecTemplateMetadataDeleteTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataDeleteTimeMap(c *Client, f map[string]ServiceSpecTemplateMetadataDeleteTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateMetadataDeleteTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateMetadataDeleteTimeSlice expands the contents of ServiceSpecTemplateMetadataDeleteTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataDeleteTimeSlice(c *Client, f []ServiceSpecTemplateMetadataDeleteTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateMetadataDeleteTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateMetadataDeleteTimeMap flattens the contents of ServiceSpecTemplateMetadataDeleteTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataDeleteTimeMap(c *Client, i interface{}) map[string]ServiceSpecTemplateMetadataDeleteTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateMetadataDeleteTime{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateMetadataDeleteTime{}
	}

	items := make(map[string]ServiceSpecTemplateMetadataDeleteTime)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateMetadataDeleteTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateMetadataDeleteTimeSlice flattens the contents of ServiceSpecTemplateMetadataDeleteTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataDeleteTimeSlice(c *Client, i interface{}) []ServiceSpecTemplateMetadataDeleteTime {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateMetadataDeleteTime{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateMetadataDeleteTime{}
	}

	items := make([]ServiceSpecTemplateMetadataDeleteTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateMetadataDeleteTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateMetadataDeleteTime expands an instance of ServiceSpecTemplateMetadataDeleteTime into a JSON
// request object.
func expandServiceSpecTemplateMetadataDeleteTime(c *Client, f *ServiceSpecTemplateMetadataDeleteTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateMetadataDeleteTime flattens an instance of ServiceSpecTemplateMetadataDeleteTime from a JSON
// response object.
func flattenServiceSpecTemplateMetadataDeleteTime(c *Client, i interface{}) *ServiceSpecTemplateMetadataDeleteTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateMetadataDeleteTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandServiceSpecTemplateSpecMap expands the contents of ServiceSpecTemplateSpec into a JSON
// request object.
func expandServiceSpecTemplateSpecMap(c *Client, f map[string]ServiceSpecTemplateSpec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecSlice expands the contents of ServiceSpecTemplateSpec into a JSON
// request object.
func expandServiceSpecTemplateSpecSlice(c *Client, f []ServiceSpecTemplateSpec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecMap flattens the contents of ServiceSpecTemplateSpec from a JSON
// response object.
func flattenServiceSpecTemplateSpecMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpec{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpec{}
	}

	items := make(map[string]ServiceSpecTemplateSpec)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecSlice flattens the contents of ServiceSpecTemplateSpec from a JSON
// response object.
func flattenServiceSpecTemplateSpecSlice(c *Client, i interface{}) []ServiceSpecTemplateSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpec{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpec{}
	}

	items := make([]ServiceSpecTemplateSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpec(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpec expands an instance of ServiceSpecTemplateSpec into a JSON
// request object.
func expandServiceSpecTemplateSpec(c *Client, f *ServiceSpecTemplateSpec) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ContainerConcurrency; !dcl.IsEmptyValueIndirect(v) {
		m["containerConcurrency"] = v
	}
	if v := f.TimeoutSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSeconds"] = v
	}
	if v := f.ServiceAccountName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountName"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersSlice(c, f.Containers); err != nil {
		return nil, fmt.Errorf("error expanding Containers into containers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["containers"] = v
	}
	if v, err := expandServiceSpecTemplateSpecVolumesSlice(c, f.Volumes); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["volumes"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpec flattens an instance of ServiceSpecTemplateSpec from a JSON
// response object.
func flattenServiceSpecTemplateSpec(c *Client, i interface{}) *ServiceSpecTemplateSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpec{}
	r.ContainerConcurrency = dcl.FlattenInteger(m["containerConcurrency"])
	r.TimeoutSeconds = dcl.FlattenInteger(m["timeoutSeconds"])
	r.ServiceAccountName = dcl.FlattenString(m["serviceAccountName"])
	r.Containers = flattenServiceSpecTemplateSpecContainersSlice(c, m["containers"])
	r.Volumes = flattenServiceSpecTemplateSpecVolumesSlice(c, m["volumes"])

	return r
}

// expandServiceSpecTemplateSpecContainersMap expands the contents of ServiceSpecTemplateSpecContainers into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersMap(c *Client, f map[string]ServiceSpecTemplateSpecContainers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersSlice expands the contents of ServiceSpecTemplateSpecContainers into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersSlice(c *Client, f []ServiceSpecTemplateSpecContainers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersMap flattens the contents of ServiceSpecTemplateSpecContainers from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainers{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainers{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainers)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersSlice flattens the contents of ServiceSpecTemplateSpecContainers from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainers {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainers{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainers{}
	}

	items := make([]ServiceSpecTemplateSpecContainers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainers(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainers expands an instance of ServiceSpecTemplateSpecContainers into a JSON
// request object.
func expandServiceSpecTemplateSpecContainers(c *Client, f *ServiceSpecTemplateSpecContainers) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["image"] = v
	}
	if v := f.Command; !dcl.IsEmptyValueIndirect(v) {
		m["command"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvSlice(c, f.Env); err != nil {
		return nil, fmt.Errorf("error expanding Env into env: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["env"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersResources(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v := f.WorkingDir; !dcl.IsEmptyValueIndirect(v) {
		m["workingDir"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersPortsSlice(c, f.Ports); err != nil {
		return nil, fmt.Errorf("error expanding Ports into ports: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ports"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvFromSlice(c, f.EnvFrom); err != nil {
		return nil, fmt.Errorf("error expanding EnvFrom into envFrom: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["envFrom"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersVolumeMountsSlice(c, f.VolumeMounts); err != nil {
		return nil, fmt.Errorf("error expanding VolumeMounts into volumeMounts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["volumeMounts"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersLivenessProbe(c, f.LivenessProbe); err != nil {
		return nil, fmt.Errorf("error expanding LivenessProbe into livenessProbe: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["livenessProbe"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersReadinessProbe(c, f.ReadinessProbe); err != nil {
		return nil, fmt.Errorf("error expanding ReadinessProbe into readinessProbe: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["readinessProbe"] = v
	}
	if v := f.TerminationMessagePath; !dcl.IsEmptyValueIndirect(v) {
		m["terminationMessagePath"] = v
	}
	if v := f.TerminationMessagePolicy; !dcl.IsEmptyValueIndirect(v) {
		m["terminationMessagePolicy"] = v
	}
	if v := f.ImagePullPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["imagePullPolicy"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersSecurityContext(c, f.SecurityContext); err != nil {
		return nil, fmt.Errorf("error expanding SecurityContext into securityContext: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["securityContext"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainers flattens an instance of ServiceSpecTemplateSpecContainers from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainers(c *Client, i interface{}) *ServiceSpecTemplateSpecContainers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainers{}
	r.Name = dcl.FlattenString(m["name"])
	r.Image = dcl.FlattenString(m["image"])
	r.Command = dcl.FlattenStringSlice(m["command"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Env = flattenServiceSpecTemplateSpecContainersEnvSlice(c, m["env"])
	r.Resources = flattenServiceSpecTemplateSpecContainersResources(c, m["resources"])
	r.WorkingDir = dcl.FlattenString(m["workingDir"])
	r.Ports = flattenServiceSpecTemplateSpecContainersPortsSlice(c, m["ports"])
	r.EnvFrom = flattenServiceSpecTemplateSpecContainersEnvFromSlice(c, m["envFrom"])
	r.VolumeMounts = flattenServiceSpecTemplateSpecContainersVolumeMountsSlice(c, m["volumeMounts"])
	r.LivenessProbe = flattenServiceSpecTemplateSpecContainersLivenessProbe(c, m["livenessProbe"])
	r.ReadinessProbe = flattenServiceSpecTemplateSpecContainersReadinessProbe(c, m["readinessProbe"])
	r.TerminationMessagePath = dcl.FlattenString(m["terminationMessagePath"])
	r.TerminationMessagePolicy = dcl.FlattenString(m["terminationMessagePolicy"])
	r.ImagePullPolicy = dcl.FlattenString(m["imagePullPolicy"])
	r.SecurityContext = flattenServiceSpecTemplateSpecContainersSecurityContext(c, m["securityContext"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvMap expands the contents of ServiceSpecTemplateSpecContainersEnv into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnv) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnv(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvSlice expands the contents of ServiceSpecTemplateSpecContainersEnv into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnv) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnv(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvMap flattens the contents of ServiceSpecTemplateSpecContainersEnv from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnv {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnv{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnv{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnv)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnv(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvSlice flattens the contents of ServiceSpecTemplateSpecContainersEnv from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnv {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnv{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnv{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnv, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnv(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnv expands an instance of ServiceSpecTemplateSpecContainersEnv into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnv(c *Client, f *ServiceSpecTemplateSpecContainersEnv) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvValueFrom(c, f.ValueFrom); err != nil {
		return nil, fmt.Errorf("error expanding ValueFrom into valueFrom: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["valueFrom"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnv flattens an instance of ServiceSpecTemplateSpecContainersEnv from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnv(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnv {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnv{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])
	r.ValueFrom = flattenServiceSpecTemplateSpecContainersEnvValueFrom(c, m["valueFrom"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvValueFromMap expands the contents of ServiceSpecTemplateSpecContainersEnvValueFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvValueFrom) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFrom(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSlice expands the contents of ServiceSpecTemplateSpecContainersEnvValueFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvValueFrom) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFrom(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromMap flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvValueFrom {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFrom{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFrom{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvValueFrom)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvValueFrom(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvValueFrom {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvValueFrom{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvValueFrom{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvValueFrom, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvValueFrom(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvValueFrom expands an instance of ServiceSpecTemplateSpecContainersEnvValueFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFrom(c *Client, f *ServiceSpecTemplateSpecContainersEnvValueFrom) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, f.ConfigMapKeyRef); err != nil {
		return nil, fmt.Errorf("error expanding ConfigMapKeyRef into configMapKeyRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configMapKeyRef"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, f.SecretKeyRef); err != nil {
		return nil, fmt.Errorf("error expanding SecretKeyRef into secretKeyRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secretKeyRef"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFrom flattens an instance of ServiceSpecTemplateSpecContainersEnvValueFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFrom(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvValueFrom {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvValueFrom{}
	r.ConfigMapKeyRef = flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, m["configMapKeyRef"])
	r.SecretKeyRef = flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, m["secretKeyRef"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefMap expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefMap flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef expands an instance of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c *Client, f *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, f.LocalObjectReference); err != nil {
		return nil, fmt.Errorf("error expanding LocalObjectReference into localObjectReference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localObjectReference"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef flattens an instance of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{}
	r.LocalObjectReference = flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, m["localObjectReference"])
	r.Key = dcl.FlattenString(m["key"])
	r.Optional = dcl.FlattenBool(m["optional"])
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceMap expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceMap flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference expands an instance of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c *Client, f *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference flattens an instance of ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefMap expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefMap flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef expands an instance of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c *Client, f *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, f.LocalObjectReference); err != nil {
		return nil, fmt.Errorf("error expanding LocalObjectReference into localObjectReference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localObjectReference"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef flattens an instance of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{}
	r.LocalObjectReference = flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, m["localObjectReference"])
	r.Key = dcl.FlattenString(m["key"])
	r.Optional = dcl.FlattenBool(m["optional"])
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceMap expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice expands the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceMap flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference expands an instance of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c *Client, f *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference flattens an instance of ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersResourcesMap expands the contents of ServiceSpecTemplateSpecContainersResources into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersResourcesMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersResourcesSlice expands the contents of ServiceSpecTemplateSpecContainersResources into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersResourcesSlice(c *Client, f []ServiceSpecTemplateSpecContainersResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersResourcesMap flattens the contents of ServiceSpecTemplateSpecContainersResources from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersResourcesMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersResources{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersResources)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersResourcesSlice flattens the contents of ServiceSpecTemplateSpecContainersResources from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersResourcesSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersResources{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersResources{}
	}

	items := make([]ServiceSpecTemplateSpecContainersResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersResources expands an instance of ServiceSpecTemplateSpecContainersResources into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersResources(c *Client, f *ServiceSpecTemplateSpecContainersResources) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Limits; !dcl.IsEmptyValueIndirect(v) {
		m["limits"] = v
	}
	if v := f.Requests; !dcl.IsEmptyValueIndirect(v) {
		m["requests"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersResources flattens an instance of ServiceSpecTemplateSpecContainersResources from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersResources(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersResources{}
	r.Limits = dcl.FlattenKeyValuePairs(m["limits"])
	r.Requests = dcl.FlattenKeyValuePairs(m["requests"])

	return r
}

// expandServiceSpecTemplateSpecContainersPortsMap expands the contents of ServiceSpecTemplateSpecContainersPorts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersPortsMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersPorts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersPorts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersPortsSlice expands the contents of ServiceSpecTemplateSpecContainersPorts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersPortsSlice(c *Client, f []ServiceSpecTemplateSpecContainersPorts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersPorts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersPortsMap flattens the contents of ServiceSpecTemplateSpecContainersPorts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersPortsMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersPorts{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersPorts{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersPorts)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersPorts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersPortsSlice flattens the contents of ServiceSpecTemplateSpecContainersPorts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersPortsSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersPorts{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersPorts{}
	}

	items := make([]ServiceSpecTemplateSpecContainersPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersPorts(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersPorts expands an instance of ServiceSpecTemplateSpecContainersPorts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersPorts(c *Client, f *ServiceSpecTemplateSpecContainersPorts) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ContainerPort; !dcl.IsEmptyValueIndirect(v) {
		m["containerPort"] = v
	}
	if v := f.Protocol; !dcl.IsEmptyValueIndirect(v) {
		m["protocol"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersPorts flattens an instance of ServiceSpecTemplateSpecContainersPorts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersPorts(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersPorts{}
	r.Name = dcl.FlattenString(m["name"])
	r.ContainerPort = dcl.FlattenInteger(m["containerPort"])
	r.Protocol = dcl.FlattenString(m["protocol"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvFromMap expands the contents of ServiceSpecTemplateSpecContainersEnvFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvFrom) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFrom(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvFromSlice expands the contents of ServiceSpecTemplateSpecContainersEnvFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvFrom) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFrom(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromMap flattens the contents of ServiceSpecTemplateSpecContainersEnvFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvFrom {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvFrom{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvFrom{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvFrom)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvFrom(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvFromSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvFrom {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvFrom{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvFrom{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvFrom, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvFrom(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvFrom expands an instance of ServiceSpecTemplateSpecContainersEnvFrom into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFrom(c *Client, f *ServiceSpecTemplateSpecContainersEnvFrom) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Prefix; !dcl.IsEmptyValueIndirect(v) {
		m["prefix"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, f.ConfigMapRef); err != nil {
		return nil, fmt.Errorf("error expanding ConfigMapRef into configMapRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configMapRef"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRef(c, f.SecretRef); err != nil {
		return nil, fmt.Errorf("error expanding SecretRef into secretRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secretRef"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFrom flattens an instance of ServiceSpecTemplateSpecContainersEnvFrom from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFrom(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvFrom {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvFrom{}
	r.Prefix = dcl.FlattenString(m["prefix"])
	r.ConfigMapRef = flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, m["configMapRef"])
	r.SecretRef = flattenServiceSpecTemplateSpecContainersEnvFromSecretRef(c, m["secretRef"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefMap expands the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice expands the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefMap flattens the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvFromConfigMapRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRef expands an instance of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c *Client, f *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, f.LocalObjectReference); err != nil {
		return nil, fmt.Errorf("error expanding LocalObjectReference into localObjectReference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localObjectReference"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRef flattens an instance of ServiceSpecTemplateSpecContainersEnvFromConfigMapRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRef(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{}
	r.LocalObjectReference = flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, m["localObjectReference"])
	r.Optional = dcl.FlattenBool(m["optional"])
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceMap expands the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice expands the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceMap flattens the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference expands an instance of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c *Client, f *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference flattens an instance of ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRefMap expands the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRefMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRefSlice expands the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRefSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvFromSecretRef) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRef(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRefMap flattens the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRefMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRef)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvFromSecretRef(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRefSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRefSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvFromSecretRef{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvFromSecretRef{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvFromSecretRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvFromSecretRef(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRef expands an instance of ServiceSpecTemplateSpecContainersEnvFromSecretRef into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRef(c *Client, f *ServiceSpecTemplateSpecContainersEnvFromSecretRef) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, f.LocalObjectReference); err != nil {
		return nil, fmt.Errorf("error expanding LocalObjectReference into localObjectReference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localObjectReference"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRef flattens an instance of ServiceSpecTemplateSpecContainersEnvFromSecretRef from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRef(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvFromSecretRef{}
	r.LocalObjectReference = flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, m["localObjectReference"])
	r.Optional = dcl.FlattenBool(m["optional"])
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceMap expands the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice expands the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice(c *Client, f []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceMap flattens the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice flattens the contents of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{}
	}

	items := make([]ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference expands an instance of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c *Client, f *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference flattens an instance of ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandServiceSpecTemplateSpecContainersVolumeMountsMap expands the contents of ServiceSpecTemplateSpecContainersVolumeMounts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersVolumeMountsMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersVolumeMounts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersVolumeMounts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersVolumeMountsSlice expands the contents of ServiceSpecTemplateSpecContainersVolumeMounts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersVolumeMountsSlice(c *Client, f []ServiceSpecTemplateSpecContainersVolumeMounts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersVolumeMounts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersVolumeMountsMap flattens the contents of ServiceSpecTemplateSpecContainersVolumeMounts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersVolumeMountsMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersVolumeMounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersVolumeMounts{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersVolumeMounts)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersVolumeMounts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersVolumeMountsSlice flattens the contents of ServiceSpecTemplateSpecContainersVolumeMounts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersVolumeMountsSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersVolumeMounts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersVolumeMounts{}
	}

	items := make([]ServiceSpecTemplateSpecContainersVolumeMounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersVolumeMounts(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersVolumeMounts expands an instance of ServiceSpecTemplateSpecContainersVolumeMounts into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersVolumeMounts(c *Client, f *ServiceSpecTemplateSpecContainersVolumeMounts) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ReadOnly; !dcl.IsEmptyValueIndirect(v) {
		m["readOnly"] = v
	}
	if v := f.MountPath; !dcl.IsEmptyValueIndirect(v) {
		m["mountPath"] = v
	}
	if v := f.SubPath; !dcl.IsEmptyValueIndirect(v) {
		m["subPath"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersVolumeMounts flattens an instance of ServiceSpecTemplateSpecContainersVolumeMounts from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersVolumeMounts(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersVolumeMounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersVolumeMounts{}
	r.Name = dcl.FlattenString(m["name"])
	r.ReadOnly = dcl.FlattenBool(m["readOnly"])
	r.MountPath = dcl.FlattenString(m["mountPath"])
	r.SubPath = dcl.FlattenString(m["subPath"])

	return r
}

// expandServiceSpecTemplateSpecContainersLivenessProbeMap expands the contents of ServiceSpecTemplateSpecContainersLivenessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersLivenessProbe) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbe(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersLivenessProbeSlice expands the contents of ServiceSpecTemplateSpecContainersLivenessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeSlice(c *Client, f []ServiceSpecTemplateSpecContainersLivenessProbe) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbe(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeMap flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersLivenessProbe {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbe{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbe{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersLivenessProbe)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersLivenessProbe(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeSlice flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersLivenessProbe {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersLivenessProbe{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersLivenessProbe{}
	}

	items := make([]ServiceSpecTemplateSpecContainersLivenessProbe, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersLivenessProbe(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersLivenessProbe expands an instance of ServiceSpecTemplateSpecContainersLivenessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbe(c *Client, f *ServiceSpecTemplateSpecContainersLivenessProbe) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.InitialDelaySeconds; !dcl.IsEmptyValueIndirect(v) {
		m["initialDelaySeconds"] = v
	}
	if v := f.TimeoutSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSeconds"] = v
	}
	if v := f.PeriodSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["periodSeconds"] = v
	}
	if v := f.SuccessThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["successThreshold"] = v
	}
	if v := f.FailureThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["failureThreshold"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersLivenessProbeExec(c, f.Exec); err != nil {
		return nil, fmt.Errorf("error expanding Exec into exec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exec"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, f.HttpGet); err != nil {
		return nil, fmt.Errorf("error expanding HttpGet into httpGet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpGet"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, f.TcpSocket); err != nil {
		return nil, fmt.Errorf("error expanding TcpSocket into tcpSocket: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tcpSocket"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbe flattens an instance of ServiceSpecTemplateSpecContainersLivenessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbe(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersLivenessProbe {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersLivenessProbe{}
	r.InitialDelaySeconds = dcl.FlattenInteger(m["initialDelaySeconds"])
	r.TimeoutSeconds = dcl.FlattenInteger(m["timeoutSeconds"])
	r.PeriodSeconds = dcl.FlattenInteger(m["periodSeconds"])
	r.SuccessThreshold = dcl.FlattenInteger(m["successThreshold"])
	r.FailureThreshold = dcl.FlattenInteger(m["failureThreshold"])
	r.Exec = flattenServiceSpecTemplateSpecContainersLivenessProbeExec(c, m["exec"])
	r.HttpGet = flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, m["httpGet"])
	r.TcpSocket = flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, m["tcpSocket"])

	return r
}

// expandServiceSpecTemplateSpecContainersLivenessProbeExecMap expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeExecMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersLivenessProbeExecSlice expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeExecSlice(c *Client, f []ServiceSpecTemplateSpecContainersLivenessProbeExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeExecMap flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeExecMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersLivenessProbeExec)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersLivenessProbeExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeExecSlice flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeExecSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersLivenessProbeExec {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersLivenessProbeExec{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersLivenessProbeExec{}
	}

	items := make([]ServiceSpecTemplateSpecContainersLivenessProbeExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersLivenessProbeExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersLivenessProbeExec expands an instance of ServiceSpecTemplateSpecContainersLivenessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeExec(c *Client, f *ServiceSpecTemplateSpecContainersLivenessProbeExec) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Command; !dcl.IsEmptyValueIndirect(v) {
		m["command"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeExec flattens an instance of ServiceSpecTemplateSpecContainersLivenessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeExec(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersLivenessProbeExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersLivenessProbeExec{}
	r.Command = dcl.FlattenString(m["command"])

	return r
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetMap expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice(c *Client, f []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetMap flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{}
	}

	items := make([]ServiceSpecTemplateSpecContainersLivenessProbeHttpGet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGet expands an instance of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c *Client, f *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Scheme; !dcl.IsEmptyValueIndirect(v) {
		m["scheme"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c, f.HttpHeaders); err != nil {
		return nil, fmt.Errorf("error expanding HttpHeaders into httpHeaders: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaders"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGet flattens an instance of ServiceSpecTemplateSpecContainersLivenessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGet(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{}
	r.Path = dcl.FlattenString(m["path"])
	r.Host = dcl.FlattenString(m["host"])
	r.Scheme = dcl.FlattenString(m["scheme"])
	r.HttpHeaders = flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c, m["httpHeaders"])

	return r
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersMap expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c *Client, f []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersMap flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{}
	}

	items := make([]ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders expands an instance of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c *Client, f *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders flattens an instance of ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocketMap expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocketMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice expands the contents of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice(c *Client, f []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocketMap flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocketMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice flattens the contents of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocketSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{}
	}

	items := make([]ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocket expands an instance of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c *Client, f *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocket flattens an instance of ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{}
	r.Port = dcl.FlattenInteger(m["port"])
	r.Host = dcl.FlattenString(m["host"])

	return r
}

// expandServiceSpecTemplateSpecContainersReadinessProbeMap expands the contents of ServiceSpecTemplateSpecContainersReadinessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersReadinessProbe) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbe(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersReadinessProbeSlice expands the contents of ServiceSpecTemplateSpecContainersReadinessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeSlice(c *Client, f []ServiceSpecTemplateSpecContainersReadinessProbe) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbe(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeMap flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersReadinessProbe {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbe{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbe{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersReadinessProbe)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersReadinessProbe(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeSlice flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersReadinessProbe {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersReadinessProbe{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersReadinessProbe{}
	}

	items := make([]ServiceSpecTemplateSpecContainersReadinessProbe, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersReadinessProbe(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersReadinessProbe expands an instance of ServiceSpecTemplateSpecContainersReadinessProbe into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbe(c *Client, f *ServiceSpecTemplateSpecContainersReadinessProbe) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.InitialDelaySeconds; !dcl.IsEmptyValueIndirect(v) {
		m["initialDelaySeconds"] = v
	}
	if v := f.TimeoutSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSeconds"] = v
	}
	if v := f.PeriodSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["periodSeconds"] = v
	}
	if v := f.SuccessThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["successThreshold"] = v
	}
	if v := f.FailureThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["failureThreshold"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersReadinessProbeExec(c, f.Exec); err != nil {
		return nil, fmt.Errorf("error expanding Exec into exec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exec"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, f.HttpGet); err != nil {
		return nil, fmt.Errorf("error expanding HttpGet into httpGet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpGet"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, f.TcpSocket); err != nil {
		return nil, fmt.Errorf("error expanding TcpSocket into tcpSocket: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tcpSocket"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbe flattens an instance of ServiceSpecTemplateSpecContainersReadinessProbe from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbe(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersReadinessProbe {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersReadinessProbe{}
	r.InitialDelaySeconds = dcl.FlattenInteger(m["initialDelaySeconds"])
	r.TimeoutSeconds = dcl.FlattenInteger(m["timeoutSeconds"])
	r.PeriodSeconds = dcl.FlattenInteger(m["periodSeconds"])
	r.SuccessThreshold = dcl.FlattenInteger(m["successThreshold"])
	r.FailureThreshold = dcl.FlattenInteger(m["failureThreshold"])
	r.Exec = flattenServiceSpecTemplateSpecContainersReadinessProbeExec(c, m["exec"])
	r.HttpGet = flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, m["httpGet"])
	r.TcpSocket = flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, m["tcpSocket"])

	return r
}

// expandServiceSpecTemplateSpecContainersReadinessProbeExecMap expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeExecMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeExec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersReadinessProbeExecSlice expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeExecSlice(c *Client, f []ServiceSpecTemplateSpecContainersReadinessProbeExec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeExec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeExecMap flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeExecMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersReadinessProbeExec)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersReadinessProbeExec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeExecSlice flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeExecSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersReadinessProbeExec {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersReadinessProbeExec{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersReadinessProbeExec{}
	}

	items := make([]ServiceSpecTemplateSpecContainersReadinessProbeExec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersReadinessProbeExec(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersReadinessProbeExec expands an instance of ServiceSpecTemplateSpecContainersReadinessProbeExec into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeExec(c *Client, f *ServiceSpecTemplateSpecContainersReadinessProbeExec) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Command; !dcl.IsEmptyValueIndirect(v) {
		m["command"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeExec flattens an instance of ServiceSpecTemplateSpecContainersReadinessProbeExec from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeExec(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersReadinessProbeExec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersReadinessProbeExec{}
	r.Command = dcl.FlattenString(m["command"])

	return r
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetMap expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice(c *Client, f []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetMap flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{}
	}

	items := make([]ServiceSpecTemplateSpecContainersReadinessProbeHttpGet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGet expands an instance of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c *Client, f *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Scheme; !dcl.IsEmptyValueIndirect(v) {
		m["scheme"] = v
	}
	if v, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c, f.HttpHeaders); err != nil {
		return nil, fmt.Errorf("error expanding HttpHeaders into httpHeaders: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaders"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGet flattens an instance of ServiceSpecTemplateSpecContainersReadinessProbeHttpGet from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGet(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{}
	r.Path = dcl.FlattenString(m["path"])
	r.Host = dcl.FlattenString(m["host"])
	r.Scheme = dcl.FlattenString(m["scheme"])
	r.HttpHeaders = flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c, m["httpHeaders"])

	return r
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersMap expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c *Client, f []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersMap flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{}
	}

	items := make([]ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders expands an instance of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c *Client, f *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders flattens an instance of ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocketMap expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocketMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice expands the contents of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice(c *Client, f []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocketMap flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocketMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice flattens the contents of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocketSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{}
	}

	items := make([]ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocket expands an instance of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c *Client, f *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocket flattens an instance of ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{}
	r.Port = dcl.FlattenInteger(m["port"])
	r.Host = dcl.FlattenString(m["host"])

	return r
}

// expandServiceSpecTemplateSpecContainersSecurityContextMap expands the contents of ServiceSpecTemplateSpecContainersSecurityContext into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersSecurityContextMap(c *Client, f map[string]ServiceSpecTemplateSpecContainersSecurityContext) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersSecurityContext(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecContainersSecurityContextSlice expands the contents of ServiceSpecTemplateSpecContainersSecurityContext into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersSecurityContextSlice(c *Client, f []ServiceSpecTemplateSpecContainersSecurityContext) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecContainersSecurityContext(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecContainersSecurityContextMap flattens the contents of ServiceSpecTemplateSpecContainersSecurityContext from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersSecurityContextMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecContainersSecurityContext {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecContainersSecurityContext{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecContainersSecurityContext{}
	}

	items := make(map[string]ServiceSpecTemplateSpecContainersSecurityContext)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecContainersSecurityContext(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecContainersSecurityContextSlice flattens the contents of ServiceSpecTemplateSpecContainersSecurityContext from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersSecurityContextSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecContainersSecurityContext {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecContainersSecurityContext{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecContainersSecurityContext{}
	}

	items := make([]ServiceSpecTemplateSpecContainersSecurityContext, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecContainersSecurityContext(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecContainersSecurityContext expands an instance of ServiceSpecTemplateSpecContainersSecurityContext into a JSON
// request object.
func expandServiceSpecTemplateSpecContainersSecurityContext(c *Client, f *ServiceSpecTemplateSpecContainersSecurityContext) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.RunAsUser; !dcl.IsEmptyValueIndirect(v) {
		m["runAsUser"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecContainersSecurityContext flattens an instance of ServiceSpecTemplateSpecContainersSecurityContext from a JSON
// response object.
func flattenServiceSpecTemplateSpecContainersSecurityContext(c *Client, i interface{}) *ServiceSpecTemplateSpecContainersSecurityContext {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecContainersSecurityContext{}
	r.RunAsUser = dcl.FlattenInteger(m["runAsUser"])

	return r
}

// expandServiceSpecTemplateSpecVolumesMap expands the contents of ServiceSpecTemplateSpecVolumes into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesMap(c *Client, f map[string]ServiceSpecTemplateSpecVolumes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecVolumesSlice expands the contents of ServiceSpecTemplateSpecVolumes into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSlice(c *Client, f []ServiceSpecTemplateSpecVolumes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecVolumesMap flattens the contents of ServiceSpecTemplateSpecVolumes from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecVolumes{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecVolumes{}
	}

	items := make(map[string]ServiceSpecTemplateSpecVolumes)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecVolumes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecVolumesSlice flattens the contents of ServiceSpecTemplateSpecVolumes from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecVolumes{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecVolumes{}
	}

	items := make([]ServiceSpecTemplateSpecVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecVolumes(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecVolumes expands an instance of ServiceSpecTemplateSpecVolumes into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumes(c *Client, f *ServiceSpecTemplateSpecVolumes) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandServiceSpecTemplateSpecVolumesSecret(c, f.Secret); err != nil {
		return nil, fmt.Errorf("error expanding Secret into secret: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandServiceSpecTemplateSpecVolumesConfigMap(c, f.ConfigMap); err != nil {
		return nil, fmt.Errorf("error expanding ConfigMap into configMap: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configMap"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecVolumes flattens an instance of ServiceSpecTemplateSpecVolumes from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumes(c *Client, i interface{}) *ServiceSpecTemplateSpecVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecVolumes{}
	r.Name = dcl.FlattenString(m["name"])
	r.Secret = flattenServiceSpecTemplateSpecVolumesSecret(c, m["secret"])
	r.ConfigMap = flattenServiceSpecTemplateSpecVolumesConfigMap(c, m["configMap"])

	return r
}

// expandServiceSpecTemplateSpecVolumesSecretMap expands the contents of ServiceSpecTemplateSpecVolumesSecret into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecretMap(c *Client, f map[string]ServiceSpecTemplateSpecVolumesSecret) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesSecret(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecVolumesSecretSlice expands the contents of ServiceSpecTemplateSpecVolumesSecret into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecretSlice(c *Client, f []ServiceSpecTemplateSpecVolumesSecret) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesSecret(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecVolumesSecretMap flattens the contents of ServiceSpecTemplateSpecVolumesSecret from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecretMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecVolumesSecret {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecVolumesSecret{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecVolumesSecret{}
	}

	items := make(map[string]ServiceSpecTemplateSpecVolumesSecret)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecVolumesSecret(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecVolumesSecretSlice flattens the contents of ServiceSpecTemplateSpecVolumesSecret from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecretSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecVolumesSecret {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecVolumesSecret{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecVolumesSecret{}
	}

	items := make([]ServiceSpecTemplateSpecVolumesSecret, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecVolumesSecret(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecVolumesSecret expands an instance of ServiceSpecTemplateSpecVolumesSecret into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecret(c *Client, f *ServiceSpecTemplateSpecVolumesSecret) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.SecretName; !dcl.IsEmptyValueIndirect(v) {
		m["secretName"] = v
	}
	if v, err := expandServiceSpecTemplateSpecVolumesSecretItemsSlice(c, f.Items); err != nil {
		return nil, fmt.Errorf("error expanding Items into items: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["items"] = v
	}
	if v := f.DefaultMode; !dcl.IsEmptyValueIndirect(v) {
		m["defaultMode"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecVolumesSecret flattens an instance of ServiceSpecTemplateSpecVolumesSecret from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecret(c *Client, i interface{}) *ServiceSpecTemplateSpecVolumesSecret {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecVolumesSecret{}
	r.SecretName = dcl.FlattenString(m["secretName"])
	r.Items = flattenServiceSpecTemplateSpecVolumesSecretItemsSlice(c, m["items"])
	r.DefaultMode = dcl.FlattenInteger(m["defaultMode"])
	r.Optional = dcl.FlattenBool(m["optional"])

	return r
}

// expandServiceSpecTemplateSpecVolumesSecretItemsMap expands the contents of ServiceSpecTemplateSpecVolumesSecretItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecretItemsMap(c *Client, f map[string]ServiceSpecTemplateSpecVolumesSecretItems) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesSecretItems(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecVolumesSecretItemsSlice expands the contents of ServiceSpecTemplateSpecVolumesSecretItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecretItemsSlice(c *Client, f []ServiceSpecTemplateSpecVolumesSecretItems) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesSecretItems(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecVolumesSecretItemsMap flattens the contents of ServiceSpecTemplateSpecVolumesSecretItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecretItemsMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecVolumesSecretItems {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecVolumesSecretItems{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecVolumesSecretItems{}
	}

	items := make(map[string]ServiceSpecTemplateSpecVolumesSecretItems)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecVolumesSecretItems(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecVolumesSecretItemsSlice flattens the contents of ServiceSpecTemplateSpecVolumesSecretItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecretItemsSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecVolumesSecretItems {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecVolumesSecretItems{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecVolumesSecretItems{}
	}

	items := make([]ServiceSpecTemplateSpecVolumesSecretItems, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecVolumesSecretItems(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecVolumesSecretItems expands an instance of ServiceSpecTemplateSpecVolumesSecretItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesSecretItems(c *Client, f *ServiceSpecTemplateSpecVolumesSecretItems) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecVolumesSecretItems flattens an instance of ServiceSpecTemplateSpecVolumesSecretItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesSecretItems(c *Client, i interface{}) *ServiceSpecTemplateSpecVolumesSecretItems {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecVolumesSecretItems{}
	r.Key = dcl.FlattenString(m["key"])
	r.Path = dcl.FlattenString(m["path"])
	r.Mode = dcl.FlattenInteger(m["mode"])

	return r
}

// expandServiceSpecTemplateSpecVolumesConfigMapMap expands the contents of ServiceSpecTemplateSpecVolumesConfigMap into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMapMap(c *Client, f map[string]ServiceSpecTemplateSpecVolumesConfigMap) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesConfigMap(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecVolumesConfigMapSlice expands the contents of ServiceSpecTemplateSpecVolumesConfigMap into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMapSlice(c *Client, f []ServiceSpecTemplateSpecVolumesConfigMap) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesConfigMap(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecVolumesConfigMapMap flattens the contents of ServiceSpecTemplateSpecVolumesConfigMap from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMapMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecVolumesConfigMap {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecVolumesConfigMap{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecVolumesConfigMap{}
	}

	items := make(map[string]ServiceSpecTemplateSpecVolumesConfigMap)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecVolumesConfigMap(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecVolumesConfigMapSlice flattens the contents of ServiceSpecTemplateSpecVolumesConfigMap from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMapSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecVolumesConfigMap {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecVolumesConfigMap{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecVolumesConfigMap{}
	}

	items := make([]ServiceSpecTemplateSpecVolumesConfigMap, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecVolumesConfigMap(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecVolumesConfigMap expands an instance of ServiceSpecTemplateSpecVolumesConfigMap into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMap(c *Client, f *ServiceSpecTemplateSpecVolumesConfigMap) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c, f.Items); err != nil {
		return nil, fmt.Errorf("error expanding Items into items: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["items"] = v
	}
	if v := f.DefaultMode; !dcl.IsEmptyValueIndirect(v) {
		m["defaultMode"] = v
	}
	if v := f.Optional; !dcl.IsEmptyValueIndirect(v) {
		m["optional"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecVolumesConfigMap flattens an instance of ServiceSpecTemplateSpecVolumesConfigMap from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMap(c *Client, i interface{}) *ServiceSpecTemplateSpecVolumesConfigMap {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecVolumesConfigMap{}
	r.Name = dcl.FlattenString(m["name"])
	r.Items = flattenServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c, m["items"])
	r.DefaultMode = dcl.FlattenInteger(m["defaultMode"])
	r.Optional = dcl.FlattenBool(m["optional"])

	return r
}

// expandServiceSpecTemplateSpecVolumesConfigMapItemsMap expands the contents of ServiceSpecTemplateSpecVolumesConfigMapItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMapItemsMap(c *Client, f map[string]ServiceSpecTemplateSpecVolumesConfigMapItems) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesConfigMapItems(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTemplateSpecVolumesConfigMapItemsSlice expands the contents of ServiceSpecTemplateSpecVolumesConfigMapItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c *Client, f []ServiceSpecTemplateSpecVolumesConfigMapItems) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTemplateSpecVolumesConfigMapItems(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTemplateSpecVolumesConfigMapItemsMap flattens the contents of ServiceSpecTemplateSpecVolumesConfigMapItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMapItemsMap(c *Client, i interface{}) map[string]ServiceSpecTemplateSpecVolumesConfigMapItems {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTemplateSpecVolumesConfigMapItems{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTemplateSpecVolumesConfigMapItems{}
	}

	items := make(map[string]ServiceSpecTemplateSpecVolumesConfigMapItems)
	for k, item := range a {
		items[k] = *flattenServiceSpecTemplateSpecVolumesConfigMapItems(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTemplateSpecVolumesConfigMapItemsSlice flattens the contents of ServiceSpecTemplateSpecVolumesConfigMapItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMapItemsSlice(c *Client, i interface{}) []ServiceSpecTemplateSpecVolumesConfigMapItems {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTemplateSpecVolumesConfigMapItems{}
	}

	if len(a) == 0 {
		return []ServiceSpecTemplateSpecVolumesConfigMapItems{}
	}

	items := make([]ServiceSpecTemplateSpecVolumesConfigMapItems, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTemplateSpecVolumesConfigMapItems(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTemplateSpecVolumesConfigMapItems expands an instance of ServiceSpecTemplateSpecVolumesConfigMapItems into a JSON
// request object.
func expandServiceSpecTemplateSpecVolumesConfigMapItems(c *Client, f *ServiceSpecTemplateSpecVolumesConfigMapItems) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenServiceSpecTemplateSpecVolumesConfigMapItems flattens an instance of ServiceSpecTemplateSpecVolumesConfigMapItems from a JSON
// response object.
func flattenServiceSpecTemplateSpecVolumesConfigMapItems(c *Client, i interface{}) *ServiceSpecTemplateSpecVolumesConfigMapItems {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTemplateSpecVolumesConfigMapItems{}
	r.Key = dcl.FlattenString(m["key"])
	r.Path = dcl.FlattenString(m["path"])
	r.Mode = dcl.FlattenInteger(m["mode"])

	return r
}

// expandServiceSpecTrafficMap expands the contents of ServiceSpecTraffic into a JSON
// request object.
func expandServiceSpecTrafficMap(c *Client, f map[string]ServiceSpecTraffic) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceSpecTraffic(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceSpecTrafficSlice expands the contents of ServiceSpecTraffic into a JSON
// request object.
func expandServiceSpecTrafficSlice(c *Client, f []ServiceSpecTraffic) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceSpecTraffic(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceSpecTrafficMap flattens the contents of ServiceSpecTraffic from a JSON
// response object.
func flattenServiceSpecTrafficMap(c *Client, i interface{}) map[string]ServiceSpecTraffic {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceSpecTraffic{}
	}

	if len(a) == 0 {
		return map[string]ServiceSpecTraffic{}
	}

	items := make(map[string]ServiceSpecTraffic)
	for k, item := range a {
		items[k] = *flattenServiceSpecTraffic(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceSpecTrafficSlice flattens the contents of ServiceSpecTraffic from a JSON
// response object.
func flattenServiceSpecTrafficSlice(c *Client, i interface{}) []ServiceSpecTraffic {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceSpecTraffic{}
	}

	if len(a) == 0 {
		return []ServiceSpecTraffic{}
	}

	items := make([]ServiceSpecTraffic, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceSpecTraffic(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceSpecTraffic expands an instance of ServiceSpecTraffic into a JSON
// request object.
func expandServiceSpecTraffic(c *Client, f *ServiceSpecTraffic) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ConfigurationName; !dcl.IsEmptyValueIndirect(v) {
		m["configurationName"] = v
	}
	if v := f.RevisionName; !dcl.IsEmptyValueIndirect(v) {
		m["revisionName"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.LatestRevision; !dcl.IsEmptyValueIndirect(v) {
		m["latestRevision"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenServiceSpecTraffic flattens an instance of ServiceSpecTraffic from a JSON
// response object.
func flattenServiceSpecTraffic(c *Client, i interface{}) *ServiceSpecTraffic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceSpecTraffic{}
	r.ConfigurationName = dcl.FlattenString(m["configurationName"])
	r.RevisionName = dcl.FlattenString(m["revisionName"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.LatestRevision = dcl.FlattenBool(m["latestRevision"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandServiceStatusMap expands the contents of ServiceStatus into a JSON
// request object.
func expandServiceStatusMap(c *Client, f map[string]ServiceStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceStatusSlice expands the contents of ServiceStatus into a JSON
// request object.
func expandServiceStatusSlice(c *Client, f []ServiceStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceStatusMap flattens the contents of ServiceStatus from a JSON
// response object.
func flattenServiceStatusMap(c *Client, i interface{}) map[string]ServiceStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceStatus{}
	}

	if len(a) == 0 {
		return map[string]ServiceStatus{}
	}

	items := make(map[string]ServiceStatus)
	for k, item := range a {
		items[k] = *flattenServiceStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceStatusSlice flattens the contents of ServiceStatus from a JSON
// response object.
func flattenServiceStatusSlice(c *Client, i interface{}) []ServiceStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStatus{}
	}

	if len(a) == 0 {
		return []ServiceStatus{}
	}

	items := make([]ServiceStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceStatus expands an instance of ServiceStatus into a JSON
// request object.
func expandServiceStatus(c *Client, f *ServiceStatus) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ObservedGeneration; !dcl.IsEmptyValueIndirect(v) {
		m["observedGeneration"] = v
	}
	if v, err := expandServiceStatusConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditions"] = v
	}
	if v := f.LatestReadyRevisionName; !dcl.IsEmptyValueIndirect(v) {
		m["latestReadyRevisionName"] = v
	}
	if v := f.LatestCreatedRevisionName; !dcl.IsEmptyValueIndirect(v) {
		m["latestCreatedRevisionName"] = v
	}
	if v, err := expandServiceStatusTrafficSlice(c, f.Traffic); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["traffic"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v, err := expandServiceStatusAddress(c, f.Address); err != nil {
		return nil, fmt.Errorf("error expanding Address into address: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["address"] = v
	}

	return m, nil
}

// flattenServiceStatus flattens an instance of ServiceStatus from a JSON
// response object.
func flattenServiceStatus(c *Client, i interface{}) *ServiceStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceStatus{}
	r.ObservedGeneration = dcl.FlattenInteger(m["observedGeneration"])
	r.Conditions = flattenServiceStatusConditionsSlice(c, m["conditions"])
	r.LatestReadyRevisionName = dcl.FlattenString(m["latestReadyRevisionName"])
	r.LatestCreatedRevisionName = dcl.FlattenString(m["latestCreatedRevisionName"])
	r.Traffic = flattenServiceStatusTrafficSlice(c, m["traffic"])
	r.Url = dcl.FlattenString(m["url"])
	r.Address = flattenServiceStatusAddress(c, m["address"])

	return r
}

// expandServiceStatusConditionsMap expands the contents of ServiceStatusConditions into a JSON
// request object.
func expandServiceStatusConditionsMap(c *Client, f map[string]ServiceStatusConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceStatusConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceStatusConditionsSlice expands the contents of ServiceStatusConditions into a JSON
// request object.
func expandServiceStatusConditionsSlice(c *Client, f []ServiceStatusConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceStatusConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceStatusConditionsMap flattens the contents of ServiceStatusConditions from a JSON
// response object.
func flattenServiceStatusConditionsMap(c *Client, i interface{}) map[string]ServiceStatusConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceStatusConditions{}
	}

	if len(a) == 0 {
		return map[string]ServiceStatusConditions{}
	}

	items := make(map[string]ServiceStatusConditions)
	for k, item := range a {
		items[k] = *flattenServiceStatusConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceStatusConditionsSlice flattens the contents of ServiceStatusConditions from a JSON
// response object.
func flattenServiceStatusConditionsSlice(c *Client, i interface{}) []ServiceStatusConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStatusConditions{}
	}

	if len(a) == 0 {
		return []ServiceStatusConditions{}
	}

	items := make([]ServiceStatusConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStatusConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceStatusConditions expands an instance of ServiceStatusConditions into a JSON
// request object.
func expandServiceStatusConditions(c *Client, f *ServiceStatusConditions) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.Reason; !dcl.IsEmptyValueIndirect(v) {
		m["reason"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandServiceStatusConditionsLastTransitionTime(c, f.LastTransitionTime); err != nil {
		return nil, fmt.Errorf("error expanding LastTransitionTime into lastTransitionTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["lastTransitionTime"] = v
	}
	if v := f.Severity; !dcl.IsEmptyValueIndirect(v) {
		m["severity"] = v
	}

	return m, nil
}

// flattenServiceStatusConditions flattens an instance of ServiceStatusConditions from a JSON
// response object.
func flattenServiceStatusConditions(c *Client, i interface{}) *ServiceStatusConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceStatusConditions{}
	r.Type = dcl.FlattenString(m["type"])
	r.Status = dcl.FlattenString(m["status"])
	r.Reason = dcl.FlattenString(m["reason"])
	r.Message = dcl.FlattenString(m["message"])
	r.LastTransitionTime = flattenServiceStatusConditionsLastTransitionTime(c, m["lastTransitionTime"])
	r.Severity = dcl.FlattenString(m["severity"])

	return r
}

// expandServiceStatusConditionsLastTransitionTimeMap expands the contents of ServiceStatusConditionsLastTransitionTime into a JSON
// request object.
func expandServiceStatusConditionsLastTransitionTimeMap(c *Client, f map[string]ServiceStatusConditionsLastTransitionTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceStatusConditionsLastTransitionTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceStatusConditionsLastTransitionTimeSlice expands the contents of ServiceStatusConditionsLastTransitionTime into a JSON
// request object.
func expandServiceStatusConditionsLastTransitionTimeSlice(c *Client, f []ServiceStatusConditionsLastTransitionTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceStatusConditionsLastTransitionTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceStatusConditionsLastTransitionTimeMap flattens the contents of ServiceStatusConditionsLastTransitionTime from a JSON
// response object.
func flattenServiceStatusConditionsLastTransitionTimeMap(c *Client, i interface{}) map[string]ServiceStatusConditionsLastTransitionTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceStatusConditionsLastTransitionTime{}
	}

	if len(a) == 0 {
		return map[string]ServiceStatusConditionsLastTransitionTime{}
	}

	items := make(map[string]ServiceStatusConditionsLastTransitionTime)
	for k, item := range a {
		items[k] = *flattenServiceStatusConditionsLastTransitionTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceStatusConditionsLastTransitionTimeSlice flattens the contents of ServiceStatusConditionsLastTransitionTime from a JSON
// response object.
func flattenServiceStatusConditionsLastTransitionTimeSlice(c *Client, i interface{}) []ServiceStatusConditionsLastTransitionTime {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStatusConditionsLastTransitionTime{}
	}

	if len(a) == 0 {
		return []ServiceStatusConditionsLastTransitionTime{}
	}

	items := make([]ServiceStatusConditionsLastTransitionTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStatusConditionsLastTransitionTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceStatusConditionsLastTransitionTime expands an instance of ServiceStatusConditionsLastTransitionTime into a JSON
// request object.
func expandServiceStatusConditionsLastTransitionTime(c *Client, f *ServiceStatusConditionsLastTransitionTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenServiceStatusConditionsLastTransitionTime flattens an instance of ServiceStatusConditionsLastTransitionTime from a JSON
// response object.
func flattenServiceStatusConditionsLastTransitionTime(c *Client, i interface{}) *ServiceStatusConditionsLastTransitionTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceStatusConditionsLastTransitionTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandServiceStatusTrafficMap expands the contents of ServiceStatusTraffic into a JSON
// request object.
func expandServiceStatusTrafficMap(c *Client, f map[string]ServiceStatusTraffic) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceStatusTraffic(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceStatusTrafficSlice expands the contents of ServiceStatusTraffic into a JSON
// request object.
func expandServiceStatusTrafficSlice(c *Client, f []ServiceStatusTraffic) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceStatusTraffic(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceStatusTrafficMap flattens the contents of ServiceStatusTraffic from a JSON
// response object.
func flattenServiceStatusTrafficMap(c *Client, i interface{}) map[string]ServiceStatusTraffic {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceStatusTraffic{}
	}

	if len(a) == 0 {
		return map[string]ServiceStatusTraffic{}
	}

	items := make(map[string]ServiceStatusTraffic)
	for k, item := range a {
		items[k] = *flattenServiceStatusTraffic(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceStatusTrafficSlice flattens the contents of ServiceStatusTraffic from a JSON
// response object.
func flattenServiceStatusTrafficSlice(c *Client, i interface{}) []ServiceStatusTraffic {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStatusTraffic{}
	}

	if len(a) == 0 {
		return []ServiceStatusTraffic{}
	}

	items := make([]ServiceStatusTraffic, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStatusTraffic(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceStatusTraffic expands an instance of ServiceStatusTraffic into a JSON
// request object.
func expandServiceStatusTraffic(c *Client, f *ServiceStatusTraffic) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ConfigurationName; !dcl.IsEmptyValueIndirect(v) {
		m["configurationName"] = v
	}
	if v := f.RevisionName; !dcl.IsEmptyValueIndirect(v) {
		m["revisionName"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.LatestRevision; !dcl.IsEmptyValueIndirect(v) {
		m["latestRevision"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenServiceStatusTraffic flattens an instance of ServiceStatusTraffic from a JSON
// response object.
func flattenServiceStatusTraffic(c *Client, i interface{}) *ServiceStatusTraffic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceStatusTraffic{}
	r.ConfigurationName = dcl.FlattenString(m["configurationName"])
	r.RevisionName = dcl.FlattenString(m["revisionName"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.LatestRevision = dcl.FlattenBool(m["latestRevision"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandServiceStatusAddressMap expands the contents of ServiceStatusAddress into a JSON
// request object.
func expandServiceStatusAddressMap(c *Client, f map[string]ServiceStatusAddress) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceStatusAddress(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceStatusAddressSlice expands the contents of ServiceStatusAddress into a JSON
// request object.
func expandServiceStatusAddressSlice(c *Client, f []ServiceStatusAddress) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceStatusAddress(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceStatusAddressMap flattens the contents of ServiceStatusAddress from a JSON
// response object.
func flattenServiceStatusAddressMap(c *Client, i interface{}) map[string]ServiceStatusAddress {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceStatusAddress{}
	}

	if len(a) == 0 {
		return map[string]ServiceStatusAddress{}
	}

	items := make(map[string]ServiceStatusAddress)
	for k, item := range a {
		items[k] = *flattenServiceStatusAddress(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceStatusAddressSlice flattens the contents of ServiceStatusAddress from a JSON
// response object.
func flattenServiceStatusAddressSlice(c *Client, i interface{}) []ServiceStatusAddress {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStatusAddress{}
	}

	if len(a) == 0 {
		return []ServiceStatusAddress{}
	}

	items := make([]ServiceStatusAddress, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStatusAddress(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceStatusAddress expands an instance of ServiceStatusAddress into a JSON
// request object.
func expandServiceStatusAddress(c *Client, f *ServiceStatusAddress) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenServiceStatusAddress flattens an instance of ServiceStatusAddress from a JSON
// response object.
func flattenServiceStatusAddress(c *Client, i interface{}) *ServiceStatusAddress {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceStatusAddress{}
	r.Url = dcl.FlattenString(m["url"])

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
