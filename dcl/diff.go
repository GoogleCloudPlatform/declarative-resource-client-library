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
	"fmt"
	"reflect"
)

// Info is a struct that contains all information about the diff that's about to occur.
type Info struct {
	// Ignore + OutputOnly cause the diff checker to always return no-diff.
	Ignore          bool
	OutputOnly      bool
	IgnoredPrefixes []string
	Type            string
}

// Diff takes in two interfaces and diffs them according to Info.
func Diff(desired, actual interface{}, info *Info) (bool, error) {
	// All Output-only fields should not be diffed.
	if info.OutputOnly || info.Ignore {
		return false, nil
	}

	// If desired is set to nil, we do not care about the field.
	if desired == nil {
		return false, nil
	}

	desiredType := valueType(desired)

	if desiredType == "invalid" {
		return false, nil
	}

	if info.Type == "ReferenceType" {
		dStr, aStr, err := strs(desired, actual)
		if err != nil {
			return false, err
		}
		return !StringEqualsWithSelfLink(dStr, aStr), nil
	}

	switch desiredType {
	case "string":
		dStr, aStr, err := strs(desired, actual)
		if err != nil {
			return false, err
		}
		return !StringCanonicalize(dStr, aStr), nil

	case "map":
		return !MapEquals(desired, actual, info.IgnoredPrefixes), nil

	case "int", "float64", "int64":
		return !reflect.DeepEqual(desired, actual), nil

	case "bool":
		dBool, aBool, err := bools(desired, actual)
		if err != nil {
			return false, err
		}
		return !BoolCanonicalize(dBool, aBool), nil

	default:
		return false, fmt.Errorf("no diffing logic exists for type: %q", desiredType)
	}
}

func valueType(i interface{}) string {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		return reflect.Indirect(reflect.ValueOf(i)).Kind().String()
	}
	return reflect.ValueOf(i).Kind().String()
}

func strs(d, i interface{}) (*string, *string, error) {
	dStr, err := str(d)
	if err != nil {
		return nil, nil, err
	}

	iStr, err := str(i)
	if err != nil {
		return nil, nil, err
	}

	return dStr, iStr, nil
}

func str(d interface{}) (*string, error) {
	dPtr, dOk := d.(*string)
	if !dOk {
		dStr, dOk2 := d.(string)
		if !dOk2 {
			return nil, fmt.Errorf("was given non string %v", d)
		}
		dPtr = String(dStr)
	}
	return dPtr, nil
}

func bools(d, i interface{}) (*bool, *bool, error) {
	dBool, err := boolean(d)
	if err != nil {
		return nil, nil, err
	}

	iBool, err := boolean(i)
	if err != nil {
		return nil, nil, err
	}

	return dBool, iBool, nil
}

func boolean(d interface{}) (*bool, error) {
	dPtr, dOk := d.(*bool)
	if !dOk {
		return nil, nil
	}
	return dPtr, nil
}

func mapVal(d interface{}) (map[string]string, error) {
	dMap, dOk := d.(map[string]string)
	if !dOk {
		return nil, fmt.Errorf("value is not a map[string]string")
	}
	return dMap, nil
}
