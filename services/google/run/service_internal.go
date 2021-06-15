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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Service) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
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

// newUpdateServiceReplaceServiceRequest creates a request for an
// Service resource's ReplaceService update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceReplaceServiceRequest(ctx context.Context, f *Service, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.ApiVersion; !dcl.IsEmptyValueIndirect(v) {
		req["apiVersion"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		req["kind"] = v
	}
	if v, err := expandServiceMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["metadata"] = v
	}
	if v, err := expandServiceSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["spec"] = v
	}
	if v, err := expandServiceStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["status"] = v
	}
	return req, nil
}

// marshalUpdateServiceReplaceServiceRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceReplaceServiceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"metadata", "name"},
	)
	return json.Marshal(m)
}

type updateServiceReplaceServiceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceReplaceServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	_, err := c.GetService(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "ReplaceService")
	if err != nil {
		return err
	}

	req, err := newUpdateServiceReplaceServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateServiceReplaceServiceRequest(c, req)
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
		res, err := unmarshalMapService(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetService(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []*dcl.FieldDiff, err error) {
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

	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.ApiVersion, rawInitial.ApiVersion) {
		rawDesired.ApiVersion = rawInitial.ApiVersion
	}
	if dcl.StringCanonicalize(rawDesired.Kind, rawInitial.Kind) {
		rawDesired.Kind = rawInitial.Kind
	}
	rawDesired.Metadata = canonicalizeServiceMetadata(rawDesired.Metadata, rawInitial.Metadata, opts...)
	rawDesired.Spec = canonicalizeServiceSpec(rawDesired.Spec, rawInitial.Spec, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

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

	if dcl.StringCanonicalize(des.GenerateName, initial.GenerateName) || dcl.IsZeroValue(des.GenerateName) {
		des.GenerateName = initial.GenerateName
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		des.Namespace = initial.Namespace
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.Annotations) {
		des.Annotations = initial.Annotations
	}
	if dcl.IsZeroValue(des.OwnerReferences) {
		des.OwnerReferences = initial.OwnerReferences
	}
	if dcl.IsZeroValue(des.Finalizers) {
		des.Finalizers = initial.Finalizers
	}
	if dcl.NameToSelfLink(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		des.ClusterName = initial.ClusterName
	}

	return des
}

func canonicalizeNewServiceMetadata(c *Client, des, nw *ServiceMetadata) *ServiceMetadata {
	if des == nil || nw == nil {
		return nw
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
	if dcl.IsZeroValue(nw.Generation) {
		nw.Generation = des.Generation
	}
	nw.CreateTime = canonicalizeNewServiceMetadataCreateTime(c, des.CreateTime, nw.CreateTime)
	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}
	if dcl.IsZeroValue(nw.Annotations) {
		nw.Annotations = des.Annotations
	}
	nw.OwnerReferences = canonicalizeNewServiceMetadataOwnerReferencesSlice(c, des.OwnerReferences, nw.OwnerReferences)
	nw.DeleteTime = canonicalizeNewServiceMetadataDeleteTime(c, des.DeleteTime, nw.DeleteTime)
	if dcl.IsZeroValue(nw.DeletionGracePeriodSeconds) {
		nw.DeletionGracePeriodSeconds = des.DeletionGracePeriodSeconds
	}
	if dcl.IsZeroValue(nw.Finalizers) {
		nw.Finalizers = des.Finalizers
	}
	if dcl.NameToSelfLink(des.ClusterName, nw.ClusterName) {
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
			if diffs, _ := compareServiceMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewServiceMetadataCreateTime(c *Client, des, nw *ServiceMetadataCreateTime) *ServiceMetadataCreateTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareServiceMetadataCreateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceMetadataOwnerReferencesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewServiceMetadataDeleteTime(c *Client, des, nw *ServiceMetadataDeleteTime) *ServiceMetadataDeleteTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareServiceMetadataDeleteTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.Annotations) {
		des.Annotations = initial.Annotations
	}
	if dcl.IsZeroValue(des.OwnerReferences) {
		des.OwnerReferences = initial.OwnerReferences
	}
	if dcl.IsZeroValue(des.Finalizers) {
		des.Finalizers = initial.Finalizers
	}
	if dcl.NameToSelfLink(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
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
	if dcl.IsZeroValue(nw.Generation) {
		nw.Generation = des.Generation
	}
	nw.CreateTime = canonicalizeNewServiceSpecTemplateMetadataCreateTime(c, des.CreateTime, nw.CreateTime)
	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}
	if dcl.IsZeroValue(nw.Annotations) {
		nw.Annotations = des.Annotations
	}
	nw.OwnerReferences = canonicalizeNewServiceSpecTemplateMetadataOwnerReferencesSlice(c, des.OwnerReferences, nw.OwnerReferences)
	nw.DeleteTime = canonicalizeNewServiceSpecTemplateMetadataDeleteTime(c, des.DeleteTime, nw.DeleteTime)
	if dcl.IsZeroValue(nw.DeletionGracePeriodSeconds) {
		nw.DeletionGracePeriodSeconds = des.DeletionGracePeriodSeconds
	}
	if dcl.IsZeroValue(nw.Finalizers) {
		nw.Finalizers = des.Finalizers
	}
	if dcl.NameToSelfLink(des.ClusterName, nw.ClusterName) {
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
			if diffs, _ := compareServiceSpecTemplateMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewServiceSpecTemplateMetadataCreateTime(c *Client, des, nw *ServiceSpecTemplateMetadataCreateTime) *ServiceSpecTemplateMetadataCreateTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareServiceSpecTemplateMetadataCreateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateMetadataOwnerReferencesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewServiceSpecTemplateMetadataDeleteTime(c *Client, des, nw *ServiceSpecTemplateMetadataDeleteTime) *ServiceSpecTemplateMetadataDeleteTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareServiceSpecTemplateMetadataDeleteTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.ContainerConcurrency) {
		nw.ContainerConcurrency = des.ContainerConcurrency
	}
	if dcl.IsZeroValue(nw.TimeoutSeconds) {
		nw.TimeoutSeconds = des.TimeoutSeconds
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
			if diffs, _ := compareServiceSpecTemplateSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Command) {
		nw.Command = des.Command
	}
	if dcl.IsZeroValue(nw.Args) {
		nw.Args = des.Args
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvValueFromNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.Limits) {
		nw.Limits = des.Limits
	}
	if dcl.IsZeroValue(nw.Requests) {
		nw.Requests = des.Requests
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.ContainerPort) {
		nw.ContainerPort = des.ContainerPort
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvFromNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvFromSecretRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersVolumeMountsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.InitialDelaySeconds) {
		nw.InitialDelaySeconds = des.InitialDelaySeconds
	}
	if dcl.IsZeroValue(nw.TimeoutSeconds) {
		nw.TimeoutSeconds = des.TimeoutSeconds
	}
	if dcl.IsZeroValue(nw.PeriodSeconds) {
		nw.PeriodSeconds = des.PeriodSeconds
	}
	if dcl.IsZeroValue(nw.SuccessThreshold) {
		nw.SuccessThreshold = des.SuccessThreshold
	}
	if dcl.IsZeroValue(nw.FailureThreshold) {
		nw.FailureThreshold = des.FailureThreshold
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersLivenessProbeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersLivenessProbeExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.InitialDelaySeconds) {
		nw.InitialDelaySeconds = des.InitialDelaySeconds
	}
	if dcl.IsZeroValue(nw.TimeoutSeconds) {
		nw.TimeoutSeconds = des.TimeoutSeconds
	}
	if dcl.IsZeroValue(nw.PeriodSeconds) {
		nw.PeriodSeconds = des.PeriodSeconds
	}
	if dcl.IsZeroValue(nw.SuccessThreshold) {
		nw.SuccessThreshold = des.SuccessThreshold
	}
	if dcl.IsZeroValue(nw.FailureThreshold) {
		nw.FailureThreshold = des.FailureThreshold
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersReadinessProbeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersReadinessProbeExecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.RunAsUser) {
		nw.RunAsUser = des.RunAsUser
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
			if diffs, _ := compareServiceSpecTemplateSpecContainersSecurityContextNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceSpecTemplateSpecVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.DefaultMode) {
		nw.DefaultMode = des.DefaultMode
	}
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
			if diffs, _ := compareServiceSpecTemplateSpecVolumesSecretNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Mode) {
		nw.Mode = des.Mode
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
			if diffs, _ := compareServiceSpecTemplateSpecVolumesSecretItemsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.DefaultMode) {
		nw.DefaultMode = des.DefaultMode
	}
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
			if diffs, _ := compareServiceSpecTemplateSpecVolumesConfigMapNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Mode) {
		nw.Mode = des.Mode
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
			if diffs, _ := compareServiceSpecTemplateSpecVolumesConfigMapItemsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
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
			if diffs, _ := compareServiceSpecTrafficNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.ObservedGeneration) {
		nw.ObservedGeneration = des.ObservedGeneration
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
			if diffs, _ := compareServiceStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceStatusConditionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareServiceStatusConditionsLastTransitionTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
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
			if diffs, _ := compareServiceStatusTrafficNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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
			if diffs, _ := compareServiceStatusAddressNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []ServiceStatusAddress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceStatusAddress(c, &d, &n))
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

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{ObjectFunction: compareServiceMetadataNewStyle, EmptyObject: EmptyServiceMetadata, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.Info{ObjectFunction: compareServiceSpecNewStyle, EmptyObject: EmptyServiceSpec, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceStatusNewStyle, EmptyObject: EmptyServiceStatus, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.GenerateName, actual.GenerateName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GenerateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceVersion, actual.ResourceVersion, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceMetadataCreateTimeNewStyle, EmptyObject: EmptyServiceMetadataCreateTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{IgnoredPrefixes: []string{"cloud.googleapis.com/"}, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OwnerReferences, actual.OwnerReferences, dcl.Info{ObjectFunction: compareServiceMetadataOwnerReferencesNewStyle, EmptyObject: EmptyServiceMetadataOwnerReferences, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("OwnerReferences")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceMetadataDeleteTimeNewStyle, EmptyObject: EmptyServiceMetadataDeleteTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeletionTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeletionGracePeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Finalizers, actual.Finalizers, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Finalizers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Controller, actual.Controller, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Controller")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("BlockOwnerDeletion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.Info{ObjectFunction: compareServiceSpecTemplateNewStyle, EmptyObject: EmptyServiceSpecTemplate, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.Info{ObjectFunction: compareServiceSpecTrafficNewStyle, EmptyObject: EmptyServiceSpecTraffic, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataNewStyle, EmptyObject: EmptyServiceSpecTemplateMetadata, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecNewStyle, EmptyObject: EmptyServiceSpecTemplateSpec, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GenerateName, actual.GenerateName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GenerateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceVersion, actual.ResourceVersion, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceSpecTemplateMetadataCreateTimeNewStyle, EmptyObject: EmptyServiceSpecTemplateMetadataCreateTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OwnerReferences, actual.OwnerReferences, dcl.Info{ObjectFunction: compareServiceSpecTemplateMetadataOwnerReferencesNewStyle, EmptyObject: EmptyServiceSpecTemplateMetadataOwnerReferences, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("OwnerReferences")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceSpecTemplateMetadataDeleteTimeNewStyle, EmptyObject: EmptyServiceSpecTemplateMetadataDeleteTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeletionTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeletionGracePeriodSeconds, actual.DeletionGracePeriodSeconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeletionGracePeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Finalizers, actual.Finalizers, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Finalizers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterName, actual.ClusterName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ClusterName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ApiVersion, actual.ApiVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ApiVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Controller, actual.Controller, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Controller")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BlockOwnerDeletion, actual.BlockOwnerDeletion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("BlockOwnerDeletion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ContainerConcurrency, actual.ContainerConcurrency, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ContainerConcurrency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccountName, actual.ServiceAccountName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ServiceAccountName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Containers, actual.Containers, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainers, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Containers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecVolumes, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnv, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersResourcesNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersResources, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkingDir, actual.WorkingDir, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("WorkingDir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersPortsNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersPorts, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnvFrom, actual.EnvFrom, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvFrom, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnvFrom")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeMounts, actual.VolumeMounts, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersVolumeMountsNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersVolumeMounts, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("VolumeMounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LivenessProbe, actual.LivenessProbe, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersLivenessProbe, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LivenessProbe")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadinessProbe, actual.ReadinessProbe, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersReadinessProbe, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ReadinessProbe")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminationMessagePath, actual.TerminationMessagePath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TerminationMessagePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminationMessagePolicy, actual.TerminationMessagePolicy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TerminationMessagePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImagePullPolicy, actual.ImagePullPolicy, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ImagePullPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityContext, actual.SecurityContext, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersSecurityContextNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersSecurityContext, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SecurityContext")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueFrom, actual.ValueFrom, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvValueFrom, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ValueFrom")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ConfigMapKeyRef, actual.ConfigMapKeyRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ConfigMapKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretKeyRef, actual.SecretKeyRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SecretKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Requests, actual.Requests, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Requests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Protocol, actual.Protocol, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Protocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Prefix, actual.Prefix, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Prefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigMapRef, actual.ConfigMapRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRef, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigMapRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretRef, actual.SecretRef, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromSecretRefNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvFromSecretRef, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecretRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.LocalObjectReference, actual.LocalObjectReference, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalObjectReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadOnly, actual.ReadOnly, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ReadOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MountPath, actual.MountPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("MountPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubPath, actual.SubPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SubPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.InitialDelaySeconds, actual.InitialDelaySeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("InitialDelaySeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeriodSeconds, actual.PeriodSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("PeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Exec, actual.Exec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeExecNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersLivenessProbeExec, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Exec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpGet, actual.HttpGet, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGet, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("HttpGet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TcpSocket, actual.TcpSocket, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeTcpSocketNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersLivenessProbeTcpSocket, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TcpSocket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scheme, actual.Scheme, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Scheme")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaders, actual.HttpHeaders, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("HttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.InitialDelaySeconds, actual.InitialDelaySeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("InitialDelaySeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSeconds, actual.TimeoutSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TimeoutSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeriodSeconds, actual.PeriodSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("PeriodSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessThreshold, actual.SuccessThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SuccessThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureThreshold, actual.FailureThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("FailureThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Exec, actual.Exec, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeExecNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersReadinessProbeExec, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Exec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpGet, actual.HttpGet, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGet, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("HttpGet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TcpSocket, actual.TcpSocket, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeTcpSocketNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersReadinessProbeTcpSocket, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("TcpSocket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scheme, actual.Scheme, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Scheme")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaders, actual.HttpHeaders, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("HttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.RunAsUser, actual.RunAsUser, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("RunAsUser")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesSecretNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecVolumesSecret, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigMap, actual.ConfigMap, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesConfigMapNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecVolumesConfigMap, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ConfigMap")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.SecretName, actual.SecretName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("SecretName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesSecretItemsNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecVolumesSecretItems, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.Info{ObjectFunction: compareServiceSpecTemplateSpecVolumesConfigMapItemsNewStyle, EmptyObject: EmptyServiceSpecTemplateSpecVolumesConfigMapItems, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Optional, actual.Optional, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Optional")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ConfigurationName, actual.ConfigurationName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ConfigurationName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionName, actual.RevisionName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("RevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestRevision, actual.LatestRevision, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LatestRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ObservedGeneration, actual.ObservedGeneration, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ObservedGeneration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{ObjectFunction: compareServiceStatusConditionsNewStyle, EmptyObject: EmptyServiceStatusConditions, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestReadyRevisionName, actual.LatestReadyRevisionName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LatestReadyRevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestCreatedRevisionName, actual.LatestCreatedRevisionName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LatestCreatedRevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.Info{ObjectFunction: compareServiceStatusTrafficNewStyle, EmptyObject: EmptyServiceStatusTraffic, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Address, actual.Address, dcl.Info{ObjectFunction: compareServiceStatusAddressNewStyle, EmptyObject: EmptyServiceStatusAddress, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Address")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.Info{ObjectFunction: compareServiceStatusConditionsLastTransitionTimeNewStyle, EmptyObject: EmptyServiceStatusConditionsLastTransitionTime, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.ConfigurationName, actual.ConfigurationName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("ConfigurationName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionName, actual.RevisionName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("RevisionName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestRevision, actual.LatestRevision, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("LatestRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
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

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceReplaceServiceOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
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
	normalized.ApiVersion = dcl.SelfLinkToName(r.ApiVersion)
	normalized.Kind = dcl.SelfLinkToName(r.Kind)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
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
	if updateName == "ReplaceService" {
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
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"metadata", "name"},
	)

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
	dcl.MoveMapEntry(
		m,
		[]string{"metadata", "name"},
		[]string{"name"},
	)

	return flattenService(c, m), nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
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
	res.ApiVersion = dcl.FlattenString(m["apiVersion"])
	if _, ok := m["apiVersion"]; !ok {
		c.Config.Logger.Info("Using default value for apiVersion")
		res.ApiVersion = dcl.String("serving.knative.dev/v1")
	}
	res.Kind = dcl.FlattenString(m["kind"])
	if _, ok := m["kind"]; !ok {
		c.Config.Logger.Info("Using default value for kind")
		res.Kind = dcl.String("Service")
	}
	res.Metadata = flattenServiceMetadata(c, m["metadata"])
	res.Spec = flattenServiceSpec(c, m["spec"])
	res.Status = flattenServiceStatus(c, m["status"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceMetadata
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceMetadataCreateTime
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceMetadataOwnerReferences
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceMetadataDeleteTime
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpec
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplate
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateMetadata
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateMetadataCreateTime
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateMetadataOwnerReferences
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateMetadataDeleteTime
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpec
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainers
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnv
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvValueFrom
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersResources
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersPorts
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvFrom
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRef
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvFromSecretRef
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersVolumeMounts
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersLivenessProbe
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersLivenessProbeExec
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGet
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersLivenessProbeTcpSocket
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersReadinessProbe
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersReadinessProbeExec
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGet
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersReadinessProbeTcpSocket
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecContainersSecurityContext
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecVolumes
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecVolumesSecret
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecVolumesSecretItems
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecVolumesConfigMap
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTemplateSpecVolumesConfigMapItems
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceSpecTraffic
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceStatus
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceStatusConditions
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceStatusConditionsLastTransitionTime
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceStatusTraffic
	}
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceStatusAddress
	}
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

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
}

func convertFieldDiffToServiceOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceDiff, error) {
	var diffs []serviceDiff
	for _, op := range ops {
		diff := serviceDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToserviceApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToserviceApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceApiOperation, error) {
	switch op {

	case "updateServiceReplaceServiceOperation":
		return &updateServiceReplaceServiceOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
