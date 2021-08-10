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
package dcl

import (
	"context"
	"fmt"
)

// ReqCtxKey is the key type for storing values in the context.
// Context requires custom type key.
type ReqCtxKey string

// Keys used in context Value.
const (
	DoNotLogRequestsKey ReqCtxKey = "DoNotLogRequestsKey"
)

// ShouldLogRequest returns true if the request should be logged.
func ShouldLogRequest(ctx context.Context) (bool, error) {
	val := ctx.Value(DoNotLogRequestsKey)
	if val == nil {
		return true, nil
	}
	doNotLog, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("could not convert DoNotLogRequests value to bool")
	}
	return !doNotLog, nil
}
