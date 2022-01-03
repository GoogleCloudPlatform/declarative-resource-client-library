// Copyright 2022 Google LLC. All Rights Reserved.
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
// Package bigqueryreservation defines types and methods for working with bigqueryreservation GCP resources.
package beta

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// If possible, wait for the end time of the capacity commitment resource before returning. Otherwise return an error.
func (r *CapacityCommitment) waitForEndTime(ctx context.Context, _ *Client) error {
	et, err := time.Parse("2006-01-02T15:04:05.000000Z", dcl.ValueOrEmptyString(r.CommitmentEndTime))
	if err != nil {
		return err
	}
	deadline, ok := ctx.Deadline()
	if ok && et.After(deadline) {
		return fmt.Errorf("cannot delete capacity commitment ending at %v before deadline %v", et, deadline)
	}
	time.Sleep(time.Until(et))
	return nil
}
