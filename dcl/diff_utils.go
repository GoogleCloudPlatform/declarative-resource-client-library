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

// RequiresRecreate is for Operations that require recreating.
func RequiresRecreate() func(d *FieldDiff) []string {
	return func(d *FieldDiff) []string { return []string{"Recreate"} }
}

// TriggersOperation is used to tell the diff checker to trigger an operation.
func TriggersOperation(op string) func(d *FieldDiff) []string {
	return func(d *FieldDiff) []string { return []string{op} }
}

// DeduplicateOperations takes a list of FieldDiffs and returns a list of operations.
func DeduplicateOperations(ds []*FieldDiff) []string {
	var opTypes []string
	for _, d := range ds {
		for _, op := range d.ResultingOperation {
			// Two operations are considered identical if they have the same type.
			// The type of an operation is derived from the name of the update method.
			if !StringSliceContains(op, opTypes) {
				opTypes = append(opTypes, op)
			}
		}
	}
	return opTypes
}
