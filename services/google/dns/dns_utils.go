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
// Package dns contains handwritten code for dns.
package dns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// Do creates a Change request and creates a Resource Record Set.
func (op *createResourceRecordSetOperation) do(ctx context.Context, r *ResourceRecordSet, c *Client) error {
	rrset, err := expandResourceRecordSet(c, r)
	if err != nil {
		return err
	}

	chg := map[string]interface{}{
		"additions": []map[string]interface{}{
			rrset,
		},
	}
	project, mz := r.createFields()
	u, err := resourceRecordSetCreateURL(c.Config.BasePath, project, mz)
	if err != nil {
		return err
	}

	req, err := json.Marshal(chg)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.DNSOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}

	return o.Wait(ctx, c.Config, *r.Project, *r.ManagedZone)
}

// Do creates a Change request and deletes a Resource Record Set.
func (op *deleteResourceRecordSetOperation) do(ctx context.Context, r *ResourceRecordSet, c *Client) error {
	// NS records must always have a value, so we can't delete NS records on domains.
	// Unfortunately, you can set NS records on subdomains, and those
	// CAN and MUST be deleted, so we need to retrieve the managed zone,
	// check if what we're looking at is a subdomain, and only not delete
	// if it's not actually a subdomain
	if r.DnsType != nil && *r.DnsType == "NS" {
		currentMZ, err := c.GetManagedZone(ctx, &ManagedZone{
			Project: r.Project,
			Name:    r.ManagedZone,
		})
		if err != nil {
			return err
		}
		if *currentMZ.DnsName == *r.DnsName {
			return fmt.Errorf("NS records cannot be deleted due to API restrictions on domain: %v", *r.DnsName)
		}
	}
	currentRRS, err := c.GetResourceRecordSet(ctx, r)
	if err != nil {
		return err
	}

	rrset, err := expandResourceRecordSet(c, currentRRS)
	if err != nil {
		return err
	}

	chg := map[string]interface{}{
		"deletions": []map[string]interface{}{
			rrset,
		},
	}
	project, mz := r.createFields()
	u, err := resourceRecordSetCreateURL(c.Config.BasePath, project, mz)
	if err != nil {
		return err
	}

	req, err := json.Marshal(chg)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.DNSOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	return o.Wait(ctx, c.Config, *r.Project, *r.ManagedZone)
}

func (c *Client) getResourceRecordSetRaw(ctx context.Context, r *ResourceRecordSet) ([]byte, error) {

	u, err := resourceRecordSetGetURL(c.Config.BasePath, r.urlNormalized())
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

// GetResourceRecordSet fetches a Resource Record Set.
func (c *Client) GetResourceRecordSet(ctx context.Context, r *ResourceRecordSet) (*ResourceRecordSet, error) {
	b, err := c.getResourceRecordSetRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	var rrsets []*ResourceRecordSet
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	for _, v := range m["rrsets"].([]interface{}) {
		rrsets = append(rrsets, flattenResourceRecordSet(c, v))
	}

	if len(rrsets) == 0 {
		return nil, dcl.NotFoundError{}
	}

	result := rrsets[0]

	result.Project = r.Project
	result.ManagedZone = r.ManagedZone
	result.DnsName = r.DnsName
	result.DnsType = r.DnsType
	return result, result.validate()
}

// do creates a Change request and updates a Resource Record Set.
func (op *updateResourceRecordSetUpdateOperation) do(ctx context.Context, r *ResourceRecordSet, c *Client) error {
	var currentRRS *ResourceRecordSet
	if sh := dcl.FetchStateHint(op.ApplyOptions); sh != nil {
		if r, ok := sh.(*ResourceRecordSet); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ResourceRecordSet, got %T", sh)
		} else {
			currentRRS = r
		}
	}
	if currentRRS == nil {
		currentRRS = r
	}

	currentRRS, err := c.GetResourceRecordSet(ctx, currentRRS)
	if err != nil {
		return err
	}

	add, err := expandResourceRecordSet(c, r)
	if err != nil {
		return err
	}

	del, err := expandResourceRecordSet(c, currentRRS)
	if err != nil {
		return err
	}

	chg := map[string]interface{}{
		"additions": []map[string]interface{}{
			add,
		},
		"deletions": []map[string]interface{}{
			del,
		},
	}

	project, mz := r.createFields()
	u, err := resourceRecordSetCreateURL(c.Config.BasePath, project, mz)
	if err != nil {
		return err
	}

	req, err := json.Marshal(chg)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.DNSOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	return o.Wait(ctx, c.Config, *r.Project, *r.ManagedZone)
}

// managedZoneEncodeReverseLookupConfig encodes the ReverseLookup field as an empty object if true, and
// no object if not true or not present.
func expandManagedZoneReverseLookupConfig(f *ManagedZone, val *bool) (interface{}, error) {
	if val != nil && *val {
		return struct{}{}, nil
	}
	return nil, nil
}

func flattenManagedZoneReverseLookupConfig(config interface{}) *bool {
	return dcl.Bool(config != nil)
}
