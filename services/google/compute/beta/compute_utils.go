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
// Package compute contains handwritten support code for the compute service.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// EncodeImageDeprecateRequest properly encodes the image deprecation request for an image.
func EncodeImageDeprecateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	// Deprecate requests involve removing the "deprecated" key.
	if deprecatedVal, ok := m["deprecated"]; ok {
		deprecatedMap := deprecatedVal.(map[string]interface{})
		for key, value := range deprecatedMap {
			req[key] = value
		}
	}

	return req
}

// WrapTargetPoolInstance wraps the instances provided by AddInstance and RemoveInstance
// in their required format.
func WrapTargetPoolInstance(m map[string]interface{}) map[string]interface{} {
	is, ok := m["instances"].([]string)
	if !ok {
		return nil
	}
	wrapped := make([]interface{}, len(is))
	for idx, i := range is {
		wrapped[idx] = map[string]interface{}{
			"instance": i,
		}
	}
	return map[string]interface{}{
		"instances": wrapped,
	}
}

// WrapTargetPoolHealthCheck wraps the instances provided by AddHC and RemoveHC
// in their required format.
func WrapTargetPoolHealthCheck(m map[string]interface{}) map[string]interface{} {
	hcs, ok := m["healthChecks"].([]string)
	if !ok {
		return nil
	}
	wrapped := make([]interface{}, len(hcs))
	for idx, hc := range hcs {
		wrapped[idx] = map[string]interface{}{
			"healthCheck": hc,
		}
	}
	return map[string]interface{}{
		"healthChecks": wrapped,
	}
}

func canonicalizeReservationCPUPlatform(o, n interface{}) bool {
	oVal, _ := o.(*string)
	nVal, _ := n.(*string)
	return equalReservationCPUPlatform(oVal, nVal)
}

func equalReservationCPUPlatform(o, n *string) bool {
	if o == nil && n == nil {
		return true
	}
	if o == nil || n == nil {
		return false
	}
	if *o == "automatic" {
		return true
	}
	if *n == "automatic" {
		return true
	}

	return *o == *n
}

func deriveAutoscalerTarget(as *Autoscaler, _ *string) (*string, error) {
	if !dcl.IsEmptyValueIndirect(as.Location) {
		if dcl.IsZone(as.Location) {
			return dcl.DeriveField("projects/%s/zones/%s/instanceGroupManagers/%s", as.Target, as.Project, dcl.SelfLinkToName(as.Location), dcl.SelfLinkToName(as.Target))
		} else if dcl.IsRegion(as.Location) {
			return dcl.DeriveField("projects/%s/regions/%s/instanceGroupManagers/%s", as.Target, as.Project, dcl.SelfLinkToName(as.Location), dcl.SelfLinkToName(as.Target))
		}
	}
	return nil, errors.New("could not derive autoscaler target - location was neither zone or region")
}

// Custom methods for firewall policy and firewall policy rule and association which wait on a compute global organization operation.

func (op *createFirewallPolicyOperation) do(ctx context.Context, r *FirewallPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	parent := r.createFields()
	u, err := firewallPolicyCreateURL(c.Config.BasePath, parent)

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
	// Wait for object to be created.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, r.Parent); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	r.Name = &o.BaseOperation.TargetID

	if _, err := c.GetFirewallPolicy(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (op *updateFirewallPolicyPatchOperation) do(ctx context.Context, r *FirewallPolicy, c *Client) error {
	_, err := c.GetFirewallPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Patch")
	if err != nil {
		return err
	}

	req, err := newUpdateFirewallPolicyPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFirewallPolicyPatchRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// Wait for object to be updated.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, r.Parent); err != nil {
		c.Config.Logger.Warningf("Update failed after waiting for operation: %v", err)
		return err
	}

	return nil
}

// do creates a request and sends it to the appropriate URL.

func (op *deleteFirewallPolicyOperation) do(ctx context.Context, r *FirewallPolicy, c *Client) error {

	_, err := c.GetFirewallPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("FirewallPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFirewallPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := firewallPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body.
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// Wait for object to be deleted.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, r.Parent); err != nil {
		return err
	}
	_, err = c.GetFirewallPolicy(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

func (op *createFirewallPolicyRuleOperation) do(ctx context.Context, r *FirewallPolicyRule, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	firewallPolicy := r.createFields()
	u, err := firewallPolicyRuleCreateURL(c.Config.BasePath, firewallPolicy)

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
	// Get firewall policy in order to get its parent, which is needed to get the operation.
	policy, err := c.GetFirewallPolicy(ctx, &FirewallPolicy{Name: r.FirewallPolicy})
	if err != nil {
		return err
	}
	// Wait for object to be created.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, policy.Parent); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetFirewallPolicyRule(ctx, r.urlNormalized()); err != nil {
		return err
	}
	return nil
}

func (op *updateFirewallPolicyRulePatchRuleOperation) do(ctx context.Context, r *FirewallPolicyRule, c *Client) error {
	_, err := c.GetFirewallPolicyRule(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchRule")
	if err != nil {
		return err
	}

	req, err := newUpdateFirewallPolicyRulePatchRuleRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFirewallPolicyRulePatchRuleRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// Get firewall policy in order to get its parent, which is needed to get the operation.
	policy, err := c.GetFirewallPolicy(ctx, &FirewallPolicy{Name: r.FirewallPolicy})
	if err != nil {
		return err
	}
	c.Config.Logger.Infof("policy: %+v", policy)
	// Wait for object to be updated.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, policy.Parent); err != nil {
		c.Config.Logger.Warningf("Update failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetFirewallPolicyRule(ctx, r.urlNormalized()); err != nil {
		return err
	}
	return nil
}

func (op *deleteFirewallPolicyRuleOperation) do(ctx context.Context, r *FirewallPolicyRule, c *Client) error {
	c.Config.Logger.Infof("Attempting to delete %v", r)

	u, err := firewallPolicyRuleDeleteURL(c.Config.BasePath, r)

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
	// Get firewall policy in order to get its parent, which is needed to get the operation.
	policy, err := c.GetFirewallPolicy(ctx, &FirewallPolicy{Name: r.FirewallPolicy})
	if err != nil {
		return err
	}
	c.Config.Logger.Infof("policy: %+v", policy)
	// Wait for object to be deleted.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, policy.Parent); err != nil {
		c.Config.Logger.Warningf("Deletion failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	_, err = c.GetFirewallPolicyRule(ctx, r.urlNormalized())
	if !dcl.IsNotFoundOrCode(err, 400) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

func (op *createFirewallPolicyAssociationOperation) do(ctx context.Context, r *FirewallPolicyAssociation, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	firewallPolicy := r.createFields()
	u, err := firewallPolicyAssociationCreateURL(c.Config.BasePath, firewallPolicy)

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
	// Get firewall policy in order to get its parent, which is needed to get the operation.
	policy, err := c.GetFirewallPolicy(ctx, &FirewallPolicy{Name: r.FirewallPolicy})
	if err != nil {
		return err
	}
	// Wait for object to be created.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, policy.Parent); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetFirewallPolicyAssociation(ctx, r.urlNormalized()); err != nil {
		return err
	}
	return nil
}

func (op *deleteFirewallPolicyAssociationOperation) do(ctx context.Context, r *FirewallPolicyAssociation, c *Client) error {
	c.Config.Logger.Infof("Attempting to delete %v", r)

	u, err := firewallPolicyAssociationDeleteURL(c.Config.BasePath, r)

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
	// Get firewall policy in order to get its parent, which is needed to get the operation.
	policy, err := c.GetFirewallPolicy(ctx, &FirewallPolicy{Name: r.FirewallPolicy})
	if err != nil {
		return err
	}
	c.Config.Logger.Infof("policy: %+v", policy)
	// Wait for object to be deleted.
	var o operations.ComputeGlobalOrganizationOperation
	if err := dcl.ParseResponse(resp.Response, &o.BaseOperation); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, policy.Parent); err != nil {
		c.Config.Logger.Warningf("Deletion failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	_, err = c.GetFirewallPolicyAssociation(ctx, r.urlNormalized())
	if !dcl.IsNotFoundOrCode(err, 400) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

func encodeNetworkEndpointRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{}, 1)
	endpoints := make([]map[string]interface{}, 1)
	endpoints = append(endpoints, m)
	if err := dcl.PutMapEntry(req, []string{"networkEndpoints"}, endpoints); err != nil {
		return nil
	}
	return req
}

func customUnmarshalNetworkEndpoint(b []byte, c *Client) (*NetworkEndpoint, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return flattenNetworkEndpoint(c, m["networkEndpoint"]), nil
}

func (r *NetworkEndpoint) customMatcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := customUnmarshalNetworkEndpoint(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()

		if nr.Port == nil && ncr.Port == nil {
			c.Config.Logger.Info("Both Port fields null - considering equal.")
		} else if nr.Port == nil || ncr.Port == nil {
			c.Config.Logger.Info("Only one Port field is null - considering unequal.")
			return false
		} else if *nr.Port != *ncr.Port {
			return false
		}
		return true
	}
}

func (c *Client) GetNetworkEndpoint(ctx context.Context, r *NetworkEndpoint) (*NetworkEndpoint, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getNetworkEndpointRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	b, err = dcl.ExtractElementFromList(b, "items", r.customMatcher(c))
	if err != nil {
		return nil, err
	}
	result, err := customUnmarshalNetworkEndpoint(b, c)
	if err != nil {
		return nil, err
	}

	result.Project = r.Project
	result.Location = r.Location
	result.Group = r.Group
	result.Port = r.Port

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNetworkEndpointNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (op *deleteNetworkEndpointOperation) do(ctx context.Context, r *NetworkEndpoint, c *Client) error {
	c.Config.Logger.Infof("Attempting to delete %v", r)

	u, err := networkEndpointDeleteURL(c.Config.BasePath, r.urlNormalized())

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
	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Deletion failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	_, err = c.GetNetworkEndpoint(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

func machineTypeOperations() func(fd *dcl.FieldDiff) []string {
	return func(fd *dcl.FieldDiff) []string {
		// We're assuming that the instance is currently running. If it isn't, this will lead to a no-op stop operation.
		return []string{"updateInstanceStopOperation", "updateInstanceMachineTypeOperation", "updateInstanceStartOperation"}
	}
}

func targetPoolHealthCheck() func(fd *dcl.FieldDiff) []string {
	return func(fd *dcl.FieldDiff) []string {
		var ops []string
		if !dcl.IsZeroValue(fd.ToAdd) {
			ops = append(ops, "updateTargetPoolAddHCOperation")
		}
		if !dcl.IsZeroValue(fd.ToRemove) {
			ops = append(ops, "updateTargetPoolRemoveHCOperation")
		}
		return ops
	}
}

func targetPoolInstances() func(fd *dcl.FieldDiff) []string {
	return func(fd *dcl.FieldDiff) []string {
		var ops []string
		if !dcl.IsZeroValue(fd.ToAdd) {
			ops = append(ops, "updateTargetPoolAddInstanceOperation")
		}
		if !dcl.IsZeroValue(fd.ToRemove) {
			ops = append(ops, "updateTargetPoolRemoveInstanceOperation")
		}
		return ops
	}
}
