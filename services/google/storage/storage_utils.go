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
// Package storage contains handwritten support code for the storage service.
package storage

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// expandStorageBucketLifecycleWithState expands the with_state enum into the
// ternary boolean is_live. It can be true, false, or unset each corresponding
// to a different value.
func expandStorageBucketLifecycleWithState(_ *BucketLifecycleRuleCondition, v *BucketLifecycleRuleConditionWithStateEnum) (interface{}, error) {
	switch *v {
	case "LIVE":
		b := true
		return &b, nil
	case "ARCHIVED":
		b := false
		return &b, nil
	case "ANY":
		return nil, nil
	}

	return nil, fmt.Errorf("unrecognized BucketLifecycleRuleConditionWithStateEnum value: %v", v)
}

// flattenStorageBucketLifecycleWithState flattens the ternary boolean is_live
// into the with_state enum. It can be true, false, or unset each corresponding
// to a different value.
func flattenStorageBucketLifecycleWithState(v interface{}) *BucketLifecycleRuleConditionWithStateEnum {
	b, ok := v.(bool)
	if !ok { // b is unset
		return BucketLifecycleRuleConditionWithStateEnumRef("ANY")
	}

	if b {
		return BucketLifecycleRuleConditionWithStateEnumRef("LIVE")
	}

	// b is false
	return BucketLifecycleRuleConditionWithStateEnumRef("ARCHIVED")
}

// EncodeStorageHmacKeyCreateRequest encapsulates fields into a metadata {}
// block, as expected by https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys
// Note that HmacKeys is an unusual API- the JSON body shown as the "main" body
// is only sent/returned on Create, and the "storage#hmacKeyMetadata" type is
// the one recorded in serviceprotos.
func EncodeStorageHmacKeyCreateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	dcl.PutMapEntry(req, []string{"metadata"}, m)
	return req
}

// The HmacKey is already effectively deleted if it's in DELETED state.
func hmacKeyDeletePrecondition(r *HmacKey) bool {
	return *r.State == *HmacKeyStateEnumRef("DELETED")
}

// HmacKeyEnsureStateInactive ensures that an HmacKey resource is in the
// INACTIVE state prior to deletion.
func (r *HmacKey) HmacKeyEnsureStateInactive(ctx context.Context, c *Client) error {
	// ensure that the HmacKey is up-to-date
	r, err := c.GetHmacKey(ctx, r)
	if err != nil {
		return err
	}

	inactiveState := HmacKeyStateEnumRef("INACTIVE")
	if r.State == inactiveState {
		return nil
	}

	// otherwise, set to INACTIVE
	r.State = inactiveState
	c.Config.Logger.Infof("Performing sub-apply to reach desired state 'INACTIVE' for HmacKey: %v", r)
	_, err = c.ApplyHmacKey(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// retrieves metadata.accessId value
func retrieveAccessIDFromJSON(m map[string]interface{}) (*string, error) {
	metadata, ok := m["metadata"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("metadata key not found")
	}

	accessID, ok := metadata["accessId"].(string)
	if !ok {
		return nil, fmt.Errorf("metadata.AccessId not found")
	}

	return &accessID, nil
}
