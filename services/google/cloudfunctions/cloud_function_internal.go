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
package cloudfunctions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
	"reflect"
	"strings"
)

func (r *CloudFunction) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"SourceArchiveUrl", "SourceRepository"}, r.SourceArchiveUrl, r.SourceRepository); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"EventTrigger", "HttpsTrigger"}, r.EventTrigger, r.HttpsTrigger); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Region, "Region"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SourceRepository) {
		if err := r.SourceRepository.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpsTrigger) {
		if err := r.HttpsTrigger.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EventTrigger) {
		if err := r.EventTrigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CloudFunctionSourceRepository) validate() error {
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *CloudFunctionHttpsTrigger) validate() error {
	return nil
}
func (r *CloudFunctionEventTrigger) validate() error {
	if err := dcl.Required(r, "eventType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resource"); err != nil {
		return err
	}
	return nil
}

func cloudFunctionGetURL(userBasePath string, r *CloudFunction) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil
}

func cloudFunctionListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil

}

func cloudFunctionCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil

}

func cloudFunctionDeleteURL(userBasePath string, r *CloudFunction) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil
}

// cloudFunctionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type cloudFunctionApiOperation interface {
	do(context.Context, *CloudFunction, *Client) error
}

// newUpdateCloudFunctionUpdateRequest creates a request for an
// CloudFunction resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateCloudFunctionUpdateRequest(ctx context.Context, f *CloudFunction, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		req["runtime"] = v
	}
	if v, err := ExpandCloudFunctionTimeout(f, f.Timeout); err != nil {
		return nil, fmt.Errorf("error expanding Timeout into timeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["timeout"] = v
	}
	if v := f.AvailableMemoryMb; !dcl.IsEmptyValueIndirect(v) {
		req["availableMemoryMb"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.EnvironmentVariables; !dcl.IsEmptyValueIndirect(v) {
		req["environmentVariables"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		req["maxInstances"] = v
	}
	if v := f.VPCConnectorEgressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["ingressSettings"] = v
	}
	return req, nil
}

// marshalUpdateCloudFunctionUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateCloudFunctionUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateCloudFunctionUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateCloudFunctionUpdateOperation) do(ctx context.Context, r *CloudFunction, c *Client) error {
	_, err := c.GetCloudFunction(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"description", "runtime", "timeout", "availableMemoryMb", "labels", "environmentVariables", "maxInstances", "vpcConnectorEgressSettings", "ingressSettings"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateCloudFunctionUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateCloudFunctionUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listCloudFunctionRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := cloudFunctionListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != CloudFunctionMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listCloudFunctionOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listCloudFunction(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*CloudFunction, string, error) {
	b, err := c.listCloudFunctionRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listCloudFunctionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*CloudFunction
	for _, v := range m.Instances {
		res := flattenCloudFunction(c, v)
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCloudFunction(ctx context.Context, f func(*CloudFunction) bool, resources []*CloudFunction) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCloudFunction(ctx, res)
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

type deleteCloudFunctionOperation struct{}

func (op *deleteCloudFunctionOperation) do(ctx context.Context, r *CloudFunction, c *Client) error {

	_, err := c.GetCloudFunction(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("CloudFunction not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetCloudFunction checking for existence. error: %v", err)
		return err
	}

	u, err := cloudFunctionDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetCloudFunction(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createCloudFunctionOperation struct{}

func (op *createCloudFunctionOperation) do(ctx context.Context, r *CloudFunction, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := cloudFunctionCreateURL(c.Config.BasePath, project, region)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetCloudFunction(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getCloudFunctionRaw(ctx context.Context, r *CloudFunction) ([]byte, error) {

	u, err := cloudFunctionGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

func (c *Client) cloudFunctionDiffsForRawDesired(ctx context.Context, rawDesired *CloudFunction, opts ...dcl.ApplyOption) (initial, desired *CloudFunction, diffs []cloudFunctionDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *CloudFunction
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*CloudFunction); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected CloudFunction, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCloudFunction(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a CloudFunction resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve CloudFunction resource: %v", err)
		}

		c.Config.Logger.Info("Found that CloudFunction resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeCloudFunctionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for CloudFunction: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for CloudFunction: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeCloudFunctionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for CloudFunction: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeCloudFunctionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for CloudFunction: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCloudFunction(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeCloudFunctionInitialState(rawInitial, rawDesired *CloudFunction) (*CloudFunction, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeCloudFunctionDesiredState(rawDesired, rawInitial *CloudFunction, opts ...dcl.ApplyOption) (*CloudFunction, error) {

	if dcl.IsZeroValue(rawDesired.SourceArchiveUrl) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.SourceRepository) {
			rawDesired.SourceArchiveUrl = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawDesired.SourceRepository) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.SourceArchiveUrl) {
			rawDesired.SourceRepository = EmptyCloudFunctionSourceRepository
		}
	}

	if dcl.IsZeroValue(rawDesired.EventTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HttpsTrigger) {
			rawDesired.EventTrigger = EmptyCloudFunctionEventTrigger
		}
	}

	if dcl.IsZeroValue(rawDesired.HttpsTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.EventTrigger) {
			rawDesired.HttpsTrigger = EmptyCloudFunctionHttpsTrigger
		}
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*CloudFunction); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected CloudFunction, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SourceRepository = canonicalizeCloudFunctionSourceRepository(rawDesired.SourceRepository, nil, opts...)
		rawDesired.HttpsTrigger = canonicalizeCloudFunctionHttpsTrigger(rawDesired.HttpsTrigger, nil, opts...)
		rawDesired.EventTrigger = canonicalizeCloudFunctionEventTrigger(rawDesired.EventTrigger, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.SourceArchiveUrl) {
		rawDesired.SourceArchiveUrl = rawInitial.SourceArchiveUrl
	}
	rawDesired.SourceRepository = canonicalizeCloudFunctionSourceRepository(rawDesired.SourceRepository, rawInitial.SourceRepository, opts...)
	rawDesired.HttpsTrigger = canonicalizeCloudFunctionHttpsTrigger(rawDesired.HttpsTrigger, rawInitial.HttpsTrigger, opts...)
	rawDesired.EventTrigger = canonicalizeCloudFunctionEventTrigger(rawDesired.EventTrigger, rawInitial.EventTrigger, opts...)
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.IsZeroValue(rawDesired.EntryPoint) {
		rawDesired.EntryPoint = rawInitial.EntryPoint
	}
	if dcl.IsZeroValue(rawDesired.Runtime) {
		rawDesired.Runtime = rawInitial.Runtime
	}
	if dcl.IsZeroValue(rawDesired.Timeout) {
		rawDesired.Timeout = rawInitial.Timeout
	}
	if dcl.IsZeroValue(rawDesired.AvailableMemoryMb) {
		rawDesired.AvailableMemoryMb = rawInitial.AvailableMemoryMb
	}
	if dcl.IsZeroValue(rawDesired.ServiceAccountEmail) {
		rawDesired.ServiceAccountEmail = rawInitial.ServiceAccountEmail
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.VersionId) {
		rawDesired.VersionId = rawInitial.VersionId
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.EnvironmentVariables) {
		rawDesired.EnvironmentVariables = rawInitial.EnvironmentVariables
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.MaxInstances) {
		rawDesired.MaxInstances = rawInitial.MaxInstances
	}
	if dcl.NameToSelfLink(rawDesired.VPCConnector, rawInitial.VPCConnector) {
		rawDesired.VPCConnector = rawInitial.VPCConnector
	}
	if dcl.IsZeroValue(rawDesired.VPCConnectorEgressSettings) {
		rawDesired.VPCConnectorEgressSettings = rawInitial.VPCConnectorEgressSettings
	}
	if dcl.IsZeroValue(rawDesired.IngressSettings) {
		rawDesired.IngressSettings = rawInitial.IngressSettings
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeCloudFunctionNewState(c *Client, rawNew, rawDesired *CloudFunction) (*CloudFunction, error) {

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
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceArchiveUrl) && dcl.IsEmptyValueIndirect(rawDesired.SourceArchiveUrl) {
		rawNew.SourceArchiveUrl = rawDesired.SourceArchiveUrl
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceRepository) && dcl.IsEmptyValueIndirect(rawDesired.SourceRepository) {
		rawNew.SourceRepository = rawDesired.SourceRepository
	} else {
		rawNew.SourceRepository = canonicalizeNewCloudFunctionSourceRepository(c, rawDesired.SourceRepository, rawNew.SourceRepository)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpsTrigger) && dcl.IsEmptyValueIndirect(rawDesired.HttpsTrigger) {
		rawNew.HttpsTrigger = rawDesired.HttpsTrigger
	} else {
		rawNew.HttpsTrigger = canonicalizeNewCloudFunctionHttpsTrigger(c, rawDesired.HttpsTrigger, rawNew.HttpsTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EventTrigger) && dcl.IsEmptyValueIndirect(rawDesired.EventTrigger) {
		rawNew.EventTrigger = rawDesired.EventTrigger
	} else {
		rawNew.EventTrigger = canonicalizeNewCloudFunctionEventTrigger(c, rawDesired.EventTrigger, rawNew.EventTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EntryPoint) && dcl.IsEmptyValueIndirect(rawDesired.EntryPoint) {
		rawNew.EntryPoint = rawDesired.EntryPoint
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Runtime) && dcl.IsEmptyValueIndirect(rawDesired.Runtime) {
		rawNew.Runtime = rawDesired.Runtime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Timeout) && dcl.IsEmptyValueIndirect(rawDesired.Timeout) {
		rawNew.Timeout = rawDesired.Timeout
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AvailableMemoryMb) && dcl.IsEmptyValueIndirect(rawDesired.AvailableMemoryMb) {
		rawNew.AvailableMemoryMb = rawDesired.AvailableMemoryMb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmail) {
		rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionId) && dcl.IsEmptyValueIndirect(rawDesired.VersionId) {
		rawNew.VersionId = rawDesired.VersionId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnvironmentVariables) && dcl.IsEmptyValueIndirect(rawDesired.EnvironmentVariables) {
		rawNew.EnvironmentVariables = rawDesired.EnvironmentVariables
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxInstances) && dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) {
		rawNew.MaxInstances = rawDesired.MaxInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnector) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnector) {
		rawNew.VPCConnector = rawDesired.VPCConnector
	} else {
		if dcl.NameToSelfLink(rawDesired.VPCConnector, rawNew.VPCConnector) {
			rawNew.VPCConnector = rawDesired.VPCConnector
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnectorEgressSettings) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnectorEgressSettings) {
		rawNew.VPCConnectorEgressSettings = rawDesired.VPCConnectorEgressSettings
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IngressSettings) && dcl.IsEmptyValueIndirect(rawDesired.IngressSettings) {
		rawNew.IngressSettings = rawDesired.IngressSettings
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.NameToSelfLink(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
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

func canonicalizeCloudFunctionSourceRepository(des, initial *CloudFunctionSourceRepository, opts ...dcl.ApplyOption) *CloudFunctionSourceRepository {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*CloudFunction)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}
	if dcl.IsZeroValue(des.DeployedUrl) {
		des.DeployedUrl = initial.DeployedUrl
	}

	return des
}

func canonicalizeNewCloudFunctionSourceRepository(c *Client, des, nw *CloudFunctionSourceRepository) *CloudFunctionSourceRepository {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewCloudFunctionSourceRepositorySet(c *Client, des, nw []CloudFunctionSourceRepository) []CloudFunctionSourceRepository {
	if des == nil {
		return nw
	}
	var reorderedNew []CloudFunctionSourceRepository
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareCloudFunctionSourceRepository(c, &d, &n) {
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

func canonicalizeCloudFunctionHttpsTrigger(des, initial *CloudFunctionHttpsTrigger, opts ...dcl.ApplyOption) *CloudFunctionHttpsTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*CloudFunction)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewCloudFunctionHttpsTrigger(c *Client, des, nw *CloudFunctionHttpsTrigger) *CloudFunctionHttpsTrigger {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewCloudFunctionHttpsTriggerSet(c *Client, des, nw []CloudFunctionHttpsTrigger) []CloudFunctionHttpsTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []CloudFunctionHttpsTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareCloudFunctionHttpsTrigger(c, &d, &n) {
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

func canonicalizeCloudFunctionEventTrigger(des, initial *CloudFunctionEventTrigger, opts ...dcl.ApplyOption) *CloudFunctionEventTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*CloudFunction)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EventType) {
		des.EventType = initial.EventType
	}
	if dcl.IsZeroValue(des.Resource) {
		des.Resource = initial.Resource
	}
	if dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.IsZeroValue(des.FailurePolicy) {
		des.FailurePolicy = initial.FailurePolicy
	}

	return des
}

func canonicalizeNewCloudFunctionEventTrigger(c *Client, des, nw *CloudFunctionEventTrigger) *CloudFunctionEventTrigger {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewCloudFunctionEventTriggerSet(c *Client, des, nw []CloudFunctionEventTrigger) []CloudFunctionEventTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []CloudFunctionEventTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareCloudFunctionEventTrigger(c, &d, &n) {
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

type cloudFunctionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         cloudFunctionApiOperation
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
func diffCloudFunction(c *Client, desired, actual *CloudFunction, opts ...dcl.ApplyOption) ([]cloudFunctionDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []cloudFunctionDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %#v\nACTUAL: %#v", desired.Description, actual.Description)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.IsZeroValue(desired.SourceArchiveUrl) && (dcl.IsZeroValue(actual.SourceArchiveUrl) || !reflect.DeepEqual(*desired.SourceArchiveUrl, *actual.SourceArchiveUrl)) {
		c.Config.Logger.Infof("Detected diff in SourceArchiveUrl.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceArchiveUrl, actual.SourceArchiveUrl)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "SourceArchiveUrl",
		})
	}
	if compareCloudFunctionSourceRepository(c, desired.SourceRepository, actual.SourceRepository) {
		c.Config.Logger.Infof("Detected diff in SourceRepository.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceRepository, actual.SourceRepository)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "SourceRepository",
		})
	}
	if compareCloudFunctionHttpsTrigger(c, desired.HttpsTrigger, actual.HttpsTrigger) {
		c.Config.Logger.Infof("Detected diff in HttpsTrigger.\nDESIRED: %#v\nACTUAL: %#v", desired.HttpsTrigger, actual.HttpsTrigger)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "HttpsTrigger",
		})
	}
	if compareCloudFunctionEventTrigger(c, desired.EventTrigger, actual.EventTrigger) {
		c.Config.Logger.Infof("Detected diff in EventTrigger.\nDESIRED: %#v\nACTUAL: %#v", desired.EventTrigger, actual.EventTrigger)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "EventTrigger",
		})
	}
	if !dcl.IsZeroValue(desired.EntryPoint) && (dcl.IsZeroValue(actual.EntryPoint) || !reflect.DeepEqual(*desired.EntryPoint, *actual.EntryPoint)) {
		c.Config.Logger.Infof("Detected diff in EntryPoint.\nDESIRED: %#v\nACTUAL: %#v", desired.EntryPoint, actual.EntryPoint)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "EntryPoint",
		})
	}
	if !dcl.IsZeroValue(desired.Runtime) && (dcl.IsZeroValue(actual.Runtime) || !reflect.DeepEqual(*desired.Runtime, *actual.Runtime)) {
		c.Config.Logger.Infof("Detected diff in Runtime.\nDESIRED: %#v\nACTUAL: %#v", desired.Runtime, actual.Runtime)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "Runtime",
		})

	}
	if !dcl.IsZeroValue(desired.Timeout) && (dcl.IsZeroValue(actual.Timeout) || !reflect.DeepEqual(*desired.Timeout, *actual.Timeout)) {
		c.Config.Logger.Infof("Detected diff in Timeout.\nDESIRED: %#v\nACTUAL: %#v", desired.Timeout, actual.Timeout)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "Timeout",
		})

	}
	if !dcl.IsZeroValue(desired.AvailableMemoryMb) && (dcl.IsZeroValue(actual.AvailableMemoryMb) || !reflect.DeepEqual(*desired.AvailableMemoryMb, *actual.AvailableMemoryMb)) {
		c.Config.Logger.Infof("Detected diff in AvailableMemoryMb.\nDESIRED: %#v\nACTUAL: %#v", desired.AvailableMemoryMb, actual.AvailableMemoryMb)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "AvailableMemoryMb",
		})

	}
	if !dcl.IsZeroValue(desired.ServiceAccountEmail) && (dcl.IsZeroValue(actual.ServiceAccountEmail) || !reflect.DeepEqual(*desired.ServiceAccountEmail, *actual.ServiceAccountEmail)) {
		c.Config.Logger.Infof("Detected diff in ServiceAccountEmail.\nDESIRED: %#v\nACTUAL: %#v", desired.ServiceAccountEmail, actual.ServiceAccountEmail)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "ServiceAccountEmail",
		})
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %#v\nACTUAL: %#v", desired.Labels, actual.Labels)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "Labels",
		})

	}
	if !reflect.DeepEqual(desired.EnvironmentVariables, actual.EnvironmentVariables) {
		c.Config.Logger.Infof("Detected diff in EnvironmentVariables.\nDESIRED: %#v\nACTUAL: %#v", desired.EnvironmentVariables, actual.EnvironmentVariables)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "EnvironmentVariables",
		})

	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.NameToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %#v\nACTUAL: %#v", desired.Network, actual.Network)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.IsZeroValue(desired.MaxInstances) && (dcl.IsZeroValue(actual.MaxInstances) || !reflect.DeepEqual(*desired.MaxInstances, *actual.MaxInstances)) {
		c.Config.Logger.Infof("Detected diff in MaxInstances.\nDESIRED: %#v\nACTUAL: %#v", desired.MaxInstances, actual.MaxInstances)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "MaxInstances",
		})

	}
	if !dcl.IsZeroValue(desired.VPCConnector) && !dcl.NameToSelfLink(desired.VPCConnector, actual.VPCConnector) {
		c.Config.Logger.Infof("Detected diff in VPCConnector.\nDESIRED: %#v\nACTUAL: %#v", desired.VPCConnector, actual.VPCConnector)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "VPCConnector",
		})
	}
	if !dcl.IsZeroValue(desired.VPCConnectorEgressSettings) && (dcl.IsZeroValue(actual.VPCConnectorEgressSettings) || !reflect.DeepEqual(*desired.VPCConnectorEgressSettings, *actual.VPCConnectorEgressSettings)) {
		c.Config.Logger.Infof("Detected diff in VPCConnectorEgressSettings.\nDESIRED: %#v\nACTUAL: %#v", desired.VPCConnectorEgressSettings, actual.VPCConnectorEgressSettings)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "VPCConnectorEgressSettings",
		})

	}
	if !dcl.IsZeroValue(desired.IngressSettings) && (dcl.IsZeroValue(actual.IngressSettings) || !reflect.DeepEqual(*desired.IngressSettings, *actual.IngressSettings)) {
		c.Config.Logger.Infof("Detected diff in IngressSettings.\nDESIRED: %#v\nACTUAL: %#v", desired.IngressSettings, actual.IngressSettings)

		diffs = append(diffs, cloudFunctionDiff{
			UpdateOp:  &updateCloudFunctionUpdateOperation{},
			FieldName: "IngressSettings",
		})

	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.NameToSelfLink(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %#v\nACTUAL: %#v", desired.Region, actual.Region)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, cloudFunctionDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
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
	var deduped []cloudFunctionDiff
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
func compareCloudFunctionSourceRepositorySlice(c *Client, desired, actual []CloudFunctionSourceRepository) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionSourceRepository, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionSourceRepository(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionSourceRepository, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionSourceRepository(c *Client, desired, actual *CloudFunctionSourceRepository) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Url == nil && desired.Url != nil && !dcl.IsEmptyValueIndirect(desired.Url) {
		c.Config.Logger.Infof("desired Url %s - but actually nil", dcl.SprintResource(desired.Url))
		return true
	}
	if !reflect.DeepEqual(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) && !(dcl.IsEmptyValueIndirect(desired.Url) && dcl.IsZeroValue(actual.Url)) {
		c.Config.Logger.Infof("Diff in Url. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	return false
}
func compareCloudFunctionHttpsTriggerSlice(c *Client, desired, actual []CloudFunctionHttpsTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionHttpsTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionHttpsTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionHttpsTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionHttpsTrigger(c *Client, desired, actual *CloudFunctionHttpsTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}
func compareCloudFunctionEventTriggerSlice(c *Client, desired, actual []CloudFunctionEventTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionEventTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionEventTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionEventTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionEventTrigger(c *Client, desired, actual *CloudFunctionEventTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EventType == nil && desired.EventType != nil && !dcl.IsEmptyValueIndirect(desired.EventType) {
		c.Config.Logger.Infof("desired EventType %s - but actually nil", dcl.SprintResource(desired.EventType))
		return true
	}
	if !reflect.DeepEqual(desired.EventType, actual.EventType) && !dcl.IsZeroValue(desired.EventType) && !(dcl.IsEmptyValueIndirect(desired.EventType) && dcl.IsZeroValue(actual.EventType)) {
		c.Config.Logger.Infof("Diff in EventType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EventType), dcl.SprintResource(actual.EventType))
		return true
	}
	if actual.Resource == nil && desired.Resource != nil && !dcl.IsEmptyValueIndirect(desired.Resource) {
		c.Config.Logger.Infof("desired Resource %s - but actually nil", dcl.SprintResource(desired.Resource))
		return true
	}
	if !reflect.DeepEqual(desired.Resource, actual.Resource) && !dcl.IsZeroValue(desired.Resource) && !(dcl.IsEmptyValueIndirect(desired.Resource) && dcl.IsZeroValue(actual.Resource)) {
		c.Config.Logger.Infof("Diff in Resource. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Resource), dcl.SprintResource(actual.Resource))
		return true
	}
	if actual.Service == nil && desired.Service != nil && !dcl.IsEmptyValueIndirect(desired.Service) {
		c.Config.Logger.Infof("desired Service %s - but actually nil", dcl.SprintResource(desired.Service))
		return true
	}
	if !reflect.DeepEqual(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) && !(dcl.IsEmptyValueIndirect(desired.Service) && dcl.IsZeroValue(actual.Service)) {
		c.Config.Logger.Infof("Diff in Service. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if actual.FailurePolicy == nil && desired.FailurePolicy != nil && !dcl.IsEmptyValueIndirect(desired.FailurePolicy) {
		c.Config.Logger.Infof("desired FailurePolicy %s - but actually nil", dcl.SprintResource(desired.FailurePolicy))
		return true
	}
	if !reflect.DeepEqual(desired.FailurePolicy, actual.FailurePolicy) && !dcl.IsZeroValue(desired.FailurePolicy) && !(dcl.IsEmptyValueIndirect(desired.FailurePolicy) && dcl.IsZeroValue(actual.FailurePolicy)) {
		c.Config.Logger.Infof("Diff in FailurePolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailurePolicy), dcl.SprintResource(actual.FailurePolicy))
		return true
	}
	return false
}
func compareCloudFunctionStatusEnumSlice(c *Client, desired, actual []CloudFunctionStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionStatusEnum(c *Client, desired, actual *CloudFunctionStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareCloudFunctionVPCConnectorEgressSettingsEnumSlice(c *Client, desired, actual []CloudFunctionVPCConnectorEgressSettingsEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionVPCConnectorEgressSettingsEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionVPCConnectorEgressSettingsEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionVPCConnectorEgressSettingsEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionVPCConnectorEgressSettingsEnum(c *Client, desired, actual *CloudFunctionVPCConnectorEgressSettingsEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareCloudFunctionIngressSettingsEnumSlice(c *Client, desired, actual []CloudFunctionIngressSettingsEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in CloudFunctionIngressSettingsEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareCloudFunctionIngressSettingsEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in CloudFunctionIngressSettingsEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareCloudFunctionIngressSettingsEnum(c *Client, desired, actual *CloudFunctionIngressSettingsEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *CloudFunction) urlNormalized() *CloudFunction {
	normalized := deepcopy.Copy(*r).(CloudFunction)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.VPCConnector = dcl.SelfLinkToName(r.VPCConnector)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *CloudFunction) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *CloudFunction) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *CloudFunction) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *CloudFunction) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the CloudFunction resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *CloudFunction) marshal(c *Client) ([]byte, error) {
	m, err := expandCloudFunction(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling CloudFunction: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCloudFunction decodes JSON responses into the CloudFunction resource schema.
func unmarshalCloudFunction(b []byte, c *Client) (*CloudFunction, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenCloudFunction(c, m), nil
}

// expandCloudFunction expands CloudFunction into a JSON request object.
func expandCloudFunction(c *Client, f *CloudFunction) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/functions/%s", f.Name, f.Project, f.Region, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SourceArchiveUrl; !dcl.IsEmptyValueIndirect(v) {
		m["sourceArchiveUrl"] = v
	}
	if v, err := expandCloudFunctionSourceRepository(c, f.SourceRepository); err != nil {
		return nil, fmt.Errorf("error expanding SourceRepository into sourceRepository: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sourceRepository"] = v
	}
	if v, err := expandCloudFunctionHttpsTrigger(c, f.HttpsTrigger); err != nil {
		return nil, fmt.Errorf("error expanding HttpsTrigger into httpsTrigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpsTrigger"] = v
	}
	if v, err := expandCloudFunctionEventTrigger(c, f.EventTrigger); err != nil {
		return nil, fmt.Errorf("error expanding EventTrigger into eventTrigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["eventTrigger"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.EntryPoint; !dcl.IsEmptyValueIndirect(v) {
		m["entryPoint"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		m["runtime"] = v
	}
	if v, err := ExpandCloudFunctionTimeout(f, f.Timeout); err != nil {
		return nil, fmt.Errorf("error expanding Timeout into timeout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.AvailableMemoryMb; !dcl.IsEmptyValueIndirect(v) {
		m["availableMemoryMb"] = v
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.VersionId; !dcl.IsEmptyValueIndirect(v) {
		m["versionId"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.EnvironmentVariables; !dcl.IsEmptyValueIndirect(v) {
		m["environmentVariables"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}
	if v := f.VPCConnector; !dcl.IsEmptyValueIndirect(v) {
		m["vpcConnector"] = v
	}
	if v := f.VPCConnectorEgressSettings; !dcl.IsEmptyValueIndirect(v) {
		m["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; !dcl.IsEmptyValueIndirect(v) {
		m["ingressSettings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenCloudFunction flattens CloudFunction from a JSON request object into the
// CloudFunction type.
func flattenCloudFunction(c *Client, i interface{}) *CloudFunction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &CloudFunction{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.SourceArchiveUrl = dcl.FlattenString(m["sourceArchiveUrl"])
	r.SourceRepository = flattenCloudFunctionSourceRepository(c, m["sourceRepository"])
	r.HttpsTrigger = flattenCloudFunctionHttpsTrigger(c, m["httpsTrigger"])
	r.EventTrigger = flattenCloudFunctionEventTrigger(c, m["eventTrigger"])
	r.Status = flattenCloudFunctionStatusEnum(m["status"])
	r.EntryPoint = dcl.FlattenString(m["entryPoint"])
	r.Runtime = dcl.FlattenString(m["runtime"])
	r.Timeout = FlattenCloudFunctionTimeout(m["timeout"])
	r.AvailableMemoryMb = dcl.FlattenInteger(m["availableMemoryMb"])
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.VersionId = dcl.FlattenInteger(m["versionId"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.EnvironmentVariables = dcl.FlattenKeyValuePairs(m["environmentVariables"])
	r.Network = dcl.FlattenString(m["network"])
	r.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	r.VPCConnector = dcl.FlattenString(m["vpcConnector"])
	r.VPCConnectorEgressSettings = flattenCloudFunctionVPCConnectorEgressSettingsEnum(m["vpcConnectorEgressSettings"])
	r.IngressSettings = flattenCloudFunctionIngressSettingsEnum(m["ingressSettings"])
	r.Region = dcl.FlattenString(m["region"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandCloudFunctionSourceRepositoryMap expands the contents of CloudFunctionSourceRepository into a JSON
// request object.
func expandCloudFunctionSourceRepositoryMap(c *Client, f map[string]CloudFunctionSourceRepository) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCloudFunctionSourceRepository(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCloudFunctionSourceRepositorySlice expands the contents of CloudFunctionSourceRepository into a JSON
// request object.
func expandCloudFunctionSourceRepositorySlice(c *Client, f []CloudFunctionSourceRepository) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCloudFunctionSourceRepository(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCloudFunctionSourceRepositoryMap flattens the contents of CloudFunctionSourceRepository from a JSON
// response object.
func flattenCloudFunctionSourceRepositoryMap(c *Client, i interface{}) map[string]CloudFunctionSourceRepository {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CloudFunctionSourceRepository{}
	}

	if len(a) == 0 {
		return map[string]CloudFunctionSourceRepository{}
	}

	items := make(map[string]CloudFunctionSourceRepository)
	for k, item := range a {
		items[k] = *flattenCloudFunctionSourceRepository(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCloudFunctionSourceRepositorySlice flattens the contents of CloudFunctionSourceRepository from a JSON
// response object.
func flattenCloudFunctionSourceRepositorySlice(c *Client, i interface{}) []CloudFunctionSourceRepository {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionSourceRepository{}
	}

	if len(a) == 0 {
		return []CloudFunctionSourceRepository{}
	}

	items := make([]CloudFunctionSourceRepository, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionSourceRepository(c, item.(map[string]interface{})))
	}

	return items
}

// expandCloudFunctionSourceRepository expands an instance of CloudFunctionSourceRepository into a JSON
// request object.
func expandCloudFunctionSourceRepository(c *Client, f *CloudFunctionSourceRepository) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.DeployedUrl; !dcl.IsEmptyValueIndirect(v) {
		m["deployedUrl"] = v
	}

	return m, nil
}

// flattenCloudFunctionSourceRepository flattens an instance of CloudFunctionSourceRepository from a JSON
// response object.
func flattenCloudFunctionSourceRepository(c *Client, i interface{}) *CloudFunctionSourceRepository {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CloudFunctionSourceRepository{}
	r.Url = dcl.FlattenString(m["url"])
	r.DeployedUrl = dcl.FlattenString(m["deployedUrl"])

	return r
}

// expandCloudFunctionHttpsTriggerMap expands the contents of CloudFunctionHttpsTrigger into a JSON
// request object.
func expandCloudFunctionHttpsTriggerMap(c *Client, f map[string]CloudFunctionHttpsTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCloudFunctionHttpsTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCloudFunctionHttpsTriggerSlice expands the contents of CloudFunctionHttpsTrigger into a JSON
// request object.
func expandCloudFunctionHttpsTriggerSlice(c *Client, f []CloudFunctionHttpsTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCloudFunctionHttpsTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCloudFunctionHttpsTriggerMap flattens the contents of CloudFunctionHttpsTrigger from a JSON
// response object.
func flattenCloudFunctionHttpsTriggerMap(c *Client, i interface{}) map[string]CloudFunctionHttpsTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CloudFunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return map[string]CloudFunctionHttpsTrigger{}
	}

	items := make(map[string]CloudFunctionHttpsTrigger)
	for k, item := range a {
		items[k] = *flattenCloudFunctionHttpsTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCloudFunctionHttpsTriggerSlice flattens the contents of CloudFunctionHttpsTrigger from a JSON
// response object.
func flattenCloudFunctionHttpsTriggerSlice(c *Client, i interface{}) []CloudFunctionHttpsTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return []CloudFunctionHttpsTrigger{}
	}

	items := make([]CloudFunctionHttpsTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionHttpsTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandCloudFunctionHttpsTrigger expands an instance of CloudFunctionHttpsTrigger into a JSON
// request object.
func expandCloudFunctionHttpsTrigger(c *Client, f *CloudFunctionHttpsTrigger) (map[string]interface{}, error) {
	if f == nil || f.empty {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenCloudFunctionHttpsTrigger flattens an instance of CloudFunctionHttpsTrigger from a JSON
// response object.
func flattenCloudFunctionHttpsTrigger(c *Client, i interface{}) *CloudFunctionHttpsTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CloudFunctionHttpsTrigger{}
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandCloudFunctionEventTriggerMap expands the contents of CloudFunctionEventTrigger into a JSON
// request object.
func expandCloudFunctionEventTriggerMap(c *Client, f map[string]CloudFunctionEventTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCloudFunctionEventTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCloudFunctionEventTriggerSlice expands the contents of CloudFunctionEventTrigger into a JSON
// request object.
func expandCloudFunctionEventTriggerSlice(c *Client, f []CloudFunctionEventTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCloudFunctionEventTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCloudFunctionEventTriggerMap flattens the contents of CloudFunctionEventTrigger from a JSON
// response object.
func flattenCloudFunctionEventTriggerMap(c *Client, i interface{}) map[string]CloudFunctionEventTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CloudFunctionEventTrigger{}
	}

	if len(a) == 0 {
		return map[string]CloudFunctionEventTrigger{}
	}

	items := make(map[string]CloudFunctionEventTrigger)
	for k, item := range a {
		items[k] = *flattenCloudFunctionEventTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCloudFunctionEventTriggerSlice flattens the contents of CloudFunctionEventTrigger from a JSON
// response object.
func flattenCloudFunctionEventTriggerSlice(c *Client, i interface{}) []CloudFunctionEventTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionEventTrigger{}
	}

	if len(a) == 0 {
		return []CloudFunctionEventTrigger{}
	}

	items := make([]CloudFunctionEventTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionEventTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandCloudFunctionEventTrigger expands an instance of CloudFunctionEventTrigger into a JSON
// request object.
func expandCloudFunctionEventTrigger(c *Client, f *CloudFunctionEventTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EventType; !dcl.IsEmptyValueIndirect(v) {
		m["eventType"] = v
	}
	if v := f.Resource; !dcl.IsEmptyValueIndirect(v) {
		m["resource"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v, err := ExpandCloudFunctionEventRetry(f, f.FailurePolicy); err != nil {
		return nil, fmt.Errorf("error expanding FailurePolicy into failurePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failurePolicy"] = v
	}

	return m, nil
}

// flattenCloudFunctionEventTrigger flattens an instance of CloudFunctionEventTrigger from a JSON
// response object.
func flattenCloudFunctionEventTrigger(c *Client, i interface{}) *CloudFunctionEventTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CloudFunctionEventTrigger{}
	r.EventType = dcl.FlattenString(m["eventType"])
	r.Resource = dcl.FlattenString(m["resource"])
	r.Service = dcl.FlattenString(m["service"])
	r.FailurePolicy = FlattenCloudFunctionEventRetry(m["failurePolicy"])

	return r
}

// flattenCloudFunctionStatusEnumSlice flattens the contents of CloudFunctionStatusEnum from a JSON
// response object.
func flattenCloudFunctionStatusEnumSlice(c *Client, i interface{}) []CloudFunctionStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionStatusEnum{}
	}

	if len(a) == 0 {
		return []CloudFunctionStatusEnum{}
	}

	items := make([]CloudFunctionStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionStatusEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenCloudFunctionStatusEnum asserts that an interface is a string, and returns a
// pointer to a *CloudFunctionStatusEnum with the same value as that string.
func flattenCloudFunctionStatusEnum(i interface{}) *CloudFunctionStatusEnum {
	s, ok := i.(string)
	if !ok {
		return CloudFunctionStatusEnumRef("")
	}

	return CloudFunctionStatusEnumRef(s)
}

// flattenCloudFunctionVPCConnectorEgressSettingsEnumSlice flattens the contents of CloudFunctionVPCConnectorEgressSettingsEnum from a JSON
// response object.
func flattenCloudFunctionVPCConnectorEgressSettingsEnumSlice(c *Client, i interface{}) []CloudFunctionVPCConnectorEgressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionVPCConnectorEgressSettingsEnum{}
	}

	if len(a) == 0 {
		return []CloudFunctionVPCConnectorEgressSettingsEnum{}
	}

	items := make([]CloudFunctionVPCConnectorEgressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionVPCConnectorEgressSettingsEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenCloudFunctionVPCConnectorEgressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *CloudFunctionVPCConnectorEgressSettingsEnum with the same value as that string.
func flattenCloudFunctionVPCConnectorEgressSettingsEnum(i interface{}) *CloudFunctionVPCConnectorEgressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return CloudFunctionVPCConnectorEgressSettingsEnumRef("")
	}

	return CloudFunctionVPCConnectorEgressSettingsEnumRef(s)
}

// flattenCloudFunctionIngressSettingsEnumSlice flattens the contents of CloudFunctionIngressSettingsEnum from a JSON
// response object.
func flattenCloudFunctionIngressSettingsEnumSlice(c *Client, i interface{}) []CloudFunctionIngressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CloudFunctionIngressSettingsEnum{}
	}

	if len(a) == 0 {
		return []CloudFunctionIngressSettingsEnum{}
	}

	items := make([]CloudFunctionIngressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCloudFunctionIngressSettingsEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenCloudFunctionIngressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *CloudFunctionIngressSettingsEnum with the same value as that string.
func flattenCloudFunctionIngressSettingsEnum(i interface{}) *CloudFunctionIngressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return CloudFunctionIngressSettingsEnumRef("")
	}

	return CloudFunctionIngressSettingsEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *CloudFunction) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCloudFunction(b, c)
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
