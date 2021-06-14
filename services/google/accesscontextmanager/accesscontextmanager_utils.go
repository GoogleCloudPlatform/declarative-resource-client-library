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
package accesscontextmanager

import "context"

// GetAccessPolicyName returns the name of the singleton Access Policy
func GetAccessPolicyName(ctx context.Context, c *Client, parent *string) *string {
	if parent == nil {
		return nil
	}
	c.Config.Logger.Infof("parent: %q\n", *parent)

	policies, err := c.ListAccessPolicy(ctx, *parent)
	// All errors can be associated with no access policies existing.
	if err != nil || len(policies.Items) == 0 {
		return nil
	}
	return policies.Items[0].Name
}

// GetAccessPolicyNameResource returns the name if given a full AccessPolicy reosurce.
func GetAccessPolicyNameResource(ctx context.Context, c *Client, r *AccessPolicy) *string {
	if r.Name != nil {
		return r.Name
	}
	return GetAccessPolicyName(ctx, c, r.Parent)
}
